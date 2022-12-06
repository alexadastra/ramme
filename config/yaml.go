package config

import (
	"context"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

const fileWatcherInterval = time.Second

// NewConfigFromYAML creates new Config with YAML file watcher from file path
func NewConfigFromYAML(path string) (Config, func(ctx context.Context) error, func() error, error) {
	c := &SyncMapConfig{}

	setFunc := func() error {
		t, err := newTargetFromYAML(path)
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

// newTargetFromYAML creates new Target
func newTargetFromYAML(filePath string) (*Target, error) {
	bytes, err := loadFileData(filePath)
	if err != nil {
		return nil, err
	}

	t, err := unmarshalTargetFromYAML(bytes)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// unmarshalTargetFromYAML unmatshalls YAML file bytes into Config
func unmarshalTargetFromYAML(data []byte) (*Target, error) {
	type tmpEntry struct {
		Val interface{} `yaml:"value"`
		T   string      `yaml:"type"`
	}
	tmp := struct {
		Basic    map[Name]*tmpEntry `yaml:"basic"`
		Advanced map[Name]*tmpEntry `yaml:"advanced"`
	}{}
	err := yaml.Unmarshal(data, &tmp)
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

// loadFileData loads YAML config data from file loader int bytes
func loadFileData(configFile string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Clean(configFile))
}
