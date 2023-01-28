package config

import (
	"context"
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

const envWatcherInterval = time.Second

// NewConfigFromJSONEnv ?
func NewConfigFromJSONEnv(key string) (Config, func(ctx context.Context) error, func() error, error) {
	c := &SyncMapConfig{}

	setFunc := func() error {
		bytes, err := fetchEnv(key)
		if err != nil {
			return err
		}
		t, err := unmarshalTargetFromJSON(bytes)
		if err != nil {
			return err
		}
		c.Set(t)
		return nil
	}

	if err := setFunc(); err != nil {
		return nil, nil, nil, err
	}

	envWatcher := NewEnvWatcher(envWatcherInterval, setFunc)

	return c, envWatcher.Start, envWatcher.Stop, nil
}

// NewConfigFromJSON creates new Config with JSON file watcher from file path
func NewConfigFromJSON(path string) (Config, func(ctx context.Context) error, func() error, error) {
	c := &SyncMapConfig{}

	setFunc := func() error {
		t, err := newTargetFromJSON(path)
		if err != nil {
			return errors.Wrap(err, "failed to set config")
		}
		c.Set(t)
		return nil
	}

	if err := setFunc(); err != nil {
		return nil, nil, nil, err
	}

	fw, err := NewFileWatcher(
		path,
		fileWatcherInterval,
		setFunc)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed to create new file watcher")
	}

	return c, fw.Run, fw.Close, nil
}

// newTargetFromJSON creates new Target
func newTargetFromJSON(filePath string) (*Target, error) {
	bytes, err := loadFileData(filePath)
	if err != nil {
		return nil, err
	}

	t, err := unmarshalTargetFromJSON(bytes)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// unmarshalTargetFromJSON unmatshalls JSON file bytes into Config
func unmarshalTargetFromJSON(data []byte) (*Target, error) {
	type tmpEntry struct {
		Val interface{} `json:"value"`
		T   string      `json:"type"`
	}
	tmp := struct {
		Basic    map[Name]*tmpEntry `json:"basic"`
		Advanced map[Name]*tmpEntry `json:"advanced"`
	}{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal config")
	}
	basic := map[Name]*Entry{}
	for name, tmpEntry := range tmp.Basic {
		entry := &Entry{
			Val: tmpEntry.Val,
			T:   tmpEntry.T,
		}
		if err := entry.Validate(); err != nil {
			return nil, errors.Wrap(err, "failed to validate entry")
		}
		basic[name] = entry
	}
	advanced := map[Name]*Entry{}
	for name, tmpEntry := range tmp.Advanced {
		entry := &Entry{
			Val: tmpEntry.Val,
			T:   tmpEntry.T,
		}
		if err := entry.Validate(); err != nil {
			return nil, errors.Wrap(err, "failed to validate entry")
		}
		advanced[name] = entry
	}
	t := &Target{
		Basic:    basic,
		Advanced: advanced,
	}

	return t, nil
}
