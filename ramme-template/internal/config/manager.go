package config

import (
	"sync"
	"time"

	"github.com/alexadastra/ramme/config"
	"github.com/pkg/errors"
)

// InitAdvancedConfig inits config
func InitAdvancedConfig() (*MutexConfigManager, *config.FileWatcher, error) {
	configData, err := config.LoadFileData(config.File)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to load file")
	}
	unmarshalledConfig, err := UnmarshalConfig(configData)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to unmarshal unmarshalledConfig")
	}

	confManager := NewMutexConfigManager(&unmarshalledConfig.Advanced)

	// Watch the file for modification and update the unmarshalledConfig manager with the new unmarshalledConfig when it's available
	watcher, err := config.WatchFile(config.File, time.Second, func() error {
		var configData []byte
		configData, err = config.LoadFileData(config.File)
		if err != nil {
			return errors.Wrap(err, "failed to load file")
		}
		err = confManager.Set(configData)
		if err != nil {
			return errors.Wrap(err, "failed to reset unmarshalledConfig")
		}
		return nil
	})
	if err != nil {
		return nil, nil, err
	}
	return confManager, watcher, nil
}

// MutexConfigManager manages the configuration instance by preforming locking around access to the Config struct.
type MutexConfigManager struct {
	conf  *AdvancedConfig
	mutex *sync.Mutex
}

// NewMutexConfigManager constructs new MutexConfigManager
func NewMutexConfigManager(conf *AdvancedConfig) *MutexConfigManager {
	return &MutexConfigManager{conf, &sync.Mutex{}}
}

// Set sets new config
func (m *MutexConfigManager) Set(confBytes []byte) error {
	conf, err := UnmarshalConfig(confBytes)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal config")
	}
	m.mutex.Lock()
	m.conf = &conf.Advanced
	m.mutex.Unlock()
	return nil
}

// Get returns current config
func (m *MutexConfigManager) Get() *AdvancedConfig {
	m.mutex.Lock()
	temp := m.conf
	m.mutex.Unlock()
	return temp
}

// Close shuts manager down
func (m *MutexConfigManager) Close() {
	//Do Nothing
}
