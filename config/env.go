package config

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/pkg/errors"
)

// EnvWatcher ?
type EnvWatcher struct {
	interval time.Duration
	wg       *sync.WaitGroup
	cancel   context.CancelFunc
	callback func() error
}

// NewEnvWatcher ?
func NewEnvWatcher(interval time.Duration, action func() error) *EnvWatcher {
	return &EnvWatcher{
		interval: interval,
		callback: action,
		wg:       &sync.WaitGroup{},
	}
}

// Start ?
func (w *EnvWatcher) Start(ctx context.Context) error {
	w.wg.Add(1)
	defer w.wg.Done()
	ctx, w.cancel = context.WithCancel(ctx)
	ticker := time.NewTicker(w.interval)
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			if err := w.callback(); err != nil {
				log.Printf("error while processing callback: %s", err)
			}
		}
	}
}

// Stop ?
func (w *EnvWatcher) Stop() error {
	w.cancel()
	w.wg.Wait()
	return nil
}

// FetchEnv parses env stuff into map
func fetchEnv(key string) ([]byte, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return nil, errors.New("failed to find environment variable")
	}

	return []byte(val), nil
}
