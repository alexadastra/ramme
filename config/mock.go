package config

import (
	"sync"
	"time"
)

// MockConfig mocks real service values for other platform componrnts tests
type MockConfig struct {
	bM       *sync.Mutex
	aM       *sync.Mutex
	basic    map[Name]Entry
	advanced map[Name]Entry
}

// NewMockConfig creates new MockConfig
func NewMockConfig() Config {
	return &MockConfig{
		bM: &sync.Mutex{},
		aM: &sync.Mutex{},
		basic: map[Name]Entry{
			"host":                     {Val: "0.0.0.0", T: "string"},
			"grpc_port":                {Val: 6560, T: "int"},
			"http_port":                {Val: 8080, T: "int"},
			"http_write_timeout":       {Val: 15 * time.Second, T: "duration"},
			"http_admin_port":          {Val: 8081, T: "int"},
			"http_admin_read_timeout":  {Val: 15 * time.Second, T: "duration"},
			"http_admin_write_timeout": {Val: 15 * time.Second, T: "duration"},
			"log_level":                {Val: 1, T: "int"},
			"is_local_environment":     {Val: true, T: "bool"},
			"http_read_timeout":        {Val: 15 * time.Second, T: "duration"},
		},
		advanced: map[Name]Entry{},
	}
}

// Get fetches config entry from advanced config
func (c *MockConfig) Get(key Name) Entry {
	if _, ok := BasicConfigMapping[key]; ok {
		c.bM.Lock()
		defer c.bM.Unlock()

		tmp := c.basic[key]
		return tmp
	}

	if _, ok := AdvancedConfigMapping[key]; ok {
		c.aM.Lock()
		defer c.aM.Unlock()

		tmp := c.advanced[key]
		return tmp
	}

	return Entry{}
}

// Set sets new config from given file path
func (c *MockConfig) Set(t *Target) {}
