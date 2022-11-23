package config

import (
	"sync"
	"time"
)

// MockConfig mocks real service values for other platform componrnts tests
type MockConfig struct {
	bM       *sync.Mutex
	aM       *sync.Mutex
	Basic    map[Name]Entry `yaml:"basic"`
	Advanced map[Name]Entry `yaml:"advanced"`
}

// NewMockConfig creates new MockConfig
func NewMockConfig() Config {
	return &MockConfig{
		bM: &sync.Mutex{},
		aM: &sync.Mutex{},
		Basic: map[Name]Entry{
			"host":                     "0.0.0.0",
			"grpc_port":                6560,
			"http_port":                8080,
			"http_write_timeout":       15 * time.Second,
			"http_admin_port":          8081,
			"http_admin_read_timeout":  15 * time.Second,
			"http_admin_write_timeout": 15 * time.Second,
			"log_level":                1,
			"is_local_environment":     true,
			"http_read_timeout":        15 * time.Second,
		},
		Advanced: map[Name]Entry{},
	}
}

// Get fetches config entry from advanced config
func (c *MockConfig) Get(key Name) Entry {
	if _, ok := BasicConfigMapping[key]; ok {
		c.bM.Lock()
		defer c.bM.Unlock()

		tmp := c.Basic[key]
		return tmp
	}

	if _, ok := AdvancedConfigMapping[key]; ok {
		c.aM.Lock()
		defer c.aM.Unlock()

		tmp := c.Advanced[key]
		return tmp
	}

	return nil
}

// Set sets new config from given file path
func (c *MockConfig) Set(filePath string) error {
	return nil
}
