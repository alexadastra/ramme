package config_new

import (
	"time"

	"github.com/pkg/errors"

	"gopkg.in/fsnotify.v1"
)

// FileWatcher watches a file on a set interval, and preforms de-duplication of write
// events such that only 1 write event is reported even if multiple writes
// happened during the specified duration.
type FileWatcher struct {
	fsNotify *fsnotify.Watcher
	interval time.Duration
	done     chan struct{}
	callback func() error
}

// NewFileWatcher begins watching a file with a specific interval and action
func NewFileWatcher(path string, interval time.Duration, action func() error) (*FileWatcher, error) {
	fsWatcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	// Add the file to be watched
	_ = fsWatcher.Add(path)

	watcher := &FileWatcher{
		fsWatcher,
		interval,
		make(chan struct{}, 1),
		action,
	}

	return watcher, nil
}

// Run runs loop for checking for file change
func (w *FileWatcher) Run() error {
	// Check for write events at this interval
	tick := time.NewTicker(w.interval)

	var lastWriteEvent *fsnotify.Event
	for {
		select {
		case event := <-w.fsNotify.Events:
			// When a ConfigMap update occurs kubernetes AtomicWriter() creates a new directory;
			// writing the updated ConfigMap contents to the new directory. Once the writing is
			// complete it removes the original file symlink and replaces it with a new symlink
			// pointing to the contents of the newly created directory. It does this to achieve
			// atomic ConfigMap updates. But it also means the file we were monitoring for write
			// events never got them and was instead deleted.

			// The correct way to handle this would be to monitor the symlink instead of the
			// actual file for events. However, fsnotify.v1 does not allow us to pass in the
			// IN_DONT_FOLLOW flag to inotify which would allow us to monitor the
			// symlink for changes instead of the de-referenced file. This is not likely to
			// change as fsnotify is designed as cross-platform and not all platforms support
			// symlinks.

			if event.Op == fsnotify.Remove {
				// Since the symlink was removed, we must
				// re-register the file to be watched
				_ = w.fsNotify.Remove(event.Name)
				_ = w.fsNotify.Add(event.Name)
				lastWriteEvent = &event
			}

			// If it was a write event
			if event.Op == fsnotify.Write {
				lastWriteEvent = &event
			}
		case <-tick.C:
			// No events during this interval
			if lastWriteEvent == nil {
				continue
			}
			// Execute the callback
			if err := w.callback(); err != nil {
				return errors.Wrap(err, "failed to perform callback")
			}
			// Reset the last event
			lastWriteEvent = nil
		case <-w.done:
			goto Close
		}
	}
Close:
	close(w.done)
	return nil
}

// Close shuts watcher down
func (w *FileWatcher) Close() error {
	// FIXME: this causes panic by SIGTERM
	// w.done <- struct{}{}
	return w.fsNotify.Close()
}
