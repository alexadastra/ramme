package config

import (
	"sync"

	"github.com/pkg/errors"
)

// Manager is a simple interface that allows us to switch out both manager implementations
type Manager interface {
	Set([]byte) error
	GetBasic() *BasicConfig
	Close()
}

// MutexConfigManager manages the configuration instance by preforming locking around access to the Config struct.
type MutexConfigManager struct {
	conf  *BasicConfig
	mutex *sync.Mutex
}

// NewMutexConfigManager constructs new MutexConfigManager
func NewMutexConfigManager(conf *BasicConfig) Manager {
	return &MutexConfigManager{conf, &sync.Mutex{}}
}

// Set sets new config
func (m *MutexConfigManager) Set(configData []byte) error {
	conf, err := UnmarshalConfig(configData)
	if err != nil {
		return errors.Wrap(err, "failed to load file")
	}
	m.mutex.Lock()
	m.conf = &conf.Basic
	m.mutex.Unlock()
	return nil
}

// GetBasic returns current config
func (m *MutexConfigManager) GetBasic() *BasicConfig {
	m.mutex.Lock()
	temp := m.conf
	m.mutex.Unlock()
	return temp
}

// Close shuts manager down
func (m *MutexConfigManager) Close() {}

// ChannelConfigManager manages the configuration instance by feeding a
//pointer through a channel whenever the user calls Get()
type ChannelConfigManager struct {
	conf *BasicConfig
	get  chan *BasicConfig
	set  chan *BasicConfig
	done chan bool
}

// NewChannelConfigManager constructs new ChannelConfigManager
func NewChannelConfigManager(conf *BasicConfig) Manager {
	manager := &ChannelConfigManager{
		conf,
		make(chan *BasicConfig),
		make(chan *BasicConfig),
		make(chan bool),
	}
	manager.Start()
	return manager
}

// Start starts goroutine listening to file changes
func (m *ChannelConfigManager) Start() {
	go func() {
		defer func() {
			close(m.get)
			close(m.set)
			close(m.done)
		}()
		for {
			select {
			case m.get <- m.conf:
			case value := <-m.set:
				m.conf = value
			case <-m.done:
				return
			}
		}
	}()
}

// Close closes goroutine
func (m *ChannelConfigManager) Close() {
	m.done <- true
}

// Set sets new config
func (m *ChannelConfigManager) Set(configData []byte) error {
	conf, err := UnmarshalConfig(configData)
	if err != nil {
		return errors.Wrap(err, "failed to load file")
	}
	m.set <- &conf.Basic
	return nil
}

// GetBasic returns current config
func (m *ChannelConfigManager) GetBasic() *BasicConfig {
	return <-m.get
}
