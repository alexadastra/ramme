package config

import (
	"time"

	"gopkg.in/fsnotify.v1"
)

// FileWatcher watches a file on a set interval, and preforms de-duplication of write
// events such that only 1 write event is reported even if multiple writes
// happened during the specified duration.
type FileWatcher struct {
	fsNotify *fsnotify.Watcher
	interval time.Duration
	done     chan struct{}
	callback func()
}

// WatchFile begins watching a file with a specific interval and action
func WatchFile(path string, interval time.Duration, action func()) (*FileWatcher, error) {
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
	// Launch a go thread to watch the file
	go watcher.run()

	return watcher, err
}

func (w *FileWatcher) run() {
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
			w.callback()
			// Reset the last event
			lastWriteEvent = nil
		case <-w.done:
			goto Close
		}
	}
Close:
	close(w.done)
}

// Close shuts watcher down
func (w *FileWatcher) Close() {
	w.done <- struct{}{}
	w.fsNotify.Close()
}
