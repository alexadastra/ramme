// Package config defines basic app settings
package config

import (
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/pkg/errors"

	"gopkg.in/yaml.v2"
)

// InitBasicConfig inits config
func InitBasicConfig() (Manager, *FileWatcher, error) {
	configData, err := LoadFileData(File)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to load file")
	}
	config, err := UnmarshalConfig(configData)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to unmarshal config")
	}

	confManager := NewMutexConfigManager(&config.Basic)

	// Watch the file for modification and update the config manager with the new config when it's available
	watcher, err := WatchFile(File, time.Second, func() error {
		var configData []byte
		configData, err = LoadFileData(File)
		if err != nil {
			return errors.Wrap(err, "failed to load file")
		}
		err = confManager.Set(configData)
		if err != nil {
			return errors.Wrap(err, "failed to reset config")
		}
		return nil
	})
	if err != nil {
		return nil, nil, err
	}
	return confManager, watcher, nil
}

// LoadFileData loads config data from file loader
func LoadFileData(configFile string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Clean(configFile))
}

// UnmarshalConfig unmarshalls file bytes to advanced config
func UnmarshalConfig(configData []byte) (*Config, error) {
	conf := &Config{}
	err := yaml.Unmarshal(configData, conf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal ")
	}
	return conf, nil
}
