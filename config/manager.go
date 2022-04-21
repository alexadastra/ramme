package config

import "sync"

// Manager is a simple interface that allows us to switch out both manager implementations
type Manager interface {
	Set(*Config)
	Get() *Config
	Close()
}

// MutexConfigManager manages the configuration instance by preforming locking around access to the Config struct.
type MutexConfigManager struct {
	conf  *Config
	mutex *sync.Mutex
}

// NewMutexConfigManager constructs new MutexConfigManager
func NewMutexConfigManager(conf *Config) *MutexConfigManager {
	return &MutexConfigManager{conf, &sync.Mutex{}}
}

// Set sets new config
func (m *MutexConfigManager) Set(conf *Config) {
	m.mutex.Lock()
	m.conf = conf
	m.mutex.Unlock()
}

// Get returns current config
func (m *MutexConfigManager) Get() *Config {
	m.mutex.Lock()
	temp := m.conf
	m.mutex.Unlock()
	return temp
}

// Close shuts manager down
func (m *MutexConfigManager) Close() {
	//Do Nothing
}

// ChannelConfigManager manages the configuration instance by feeding a
//pointer through a channel whenever the user calls Get()
type ChannelConfigManager struct {
	conf *Config
	get  chan *Config
	set  chan *Config
	done chan bool
}

// NewChannelConfigManager constructs new ChannelConfigManager
func NewChannelConfigManager(conf *Config) *ChannelConfigManager {
	parser := &ChannelConfigManager{conf, make(chan *Config), make(chan *Config), make(chan bool)}
	parser.Start()
	return parser
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
func (m *ChannelConfigManager) Set(conf *Config) {
	m.set <- conf
}

// Get returns current config
func (m *ChannelConfigManager) Get() *Config {
	return <-m.get
}
