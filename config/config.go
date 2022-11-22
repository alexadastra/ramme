// Package config defines new config handling implementation
package config

import (
	"io/ioutil"
	"path/filepath"
	"sync"
	"time"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

const fileWatcherInterval = time.Second

var (
	// ServiceName contains a service name prefix which used in ENV variables
	ServiceName = "RAMME-TEMPLATE"
	// File contains path to .yaml config file
	File = "/etc/config/config.yaml"
)

// Config represents the structure that contains configurations both for logic and middleware
type Config interface {
	Get(key string) Entry
	GetBasic(key string) Entry
	Set(filePath string) error
}

// YAMLConfig fetched from YAML file
type YAMLConfig struct {
	bM          *sync.Mutex
	aM          *sync.Mutex
	Basic       map[string]Entry `yaml:"basic"`
	Advanced    map[string]Entry `yaml:"advanced"`
	fileWatcher *FileWatcher
}

// NewConfig creates new Config
func NewConfig(path string) (Config, func() error, func() error, error) {
	c := &YAMLConfig{
		bM:       &sync.Mutex{},
		aM:       &sync.Mutex{},
		Basic:    make(map[string]Entry),
		Advanced: make(map[string]Entry),
	}

	if err := c.Set(path); err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed to initially set config")
	}

	fw, err := NewFileWatcher(path, fileWatcherInterval, func() error { return c.Set(path) })
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed to create new file watcher")
	}

	c.fileWatcher = fw

	return c, fw.Run, fw.Close, nil
}

// GetBasic fetches config entry from basic config
func (c *YAMLConfig) GetBasic(key string) Entry {
	c.bM.Lock()
	defer c.bM.Unlock()

	tmp := c.Basic[key]
	return tmp
}

// Get fetches config entry from advanced config
func (c *YAMLConfig) Get(key string) Entry {
	c.aM.Lock()
	defer c.aM.Unlock()

	tmp := c.Advanced[key]
	return tmp
}

// Set sets new config from given file path
func (c *YAMLConfig) Set(filePath string) error {
	bytes, err := loadFileData(filePath)
	if err != nil {
		return err
	}

	conf, err := unmarshalConfig(bytes)
	if err != nil {
		return err
	}

	c.bM.Lock()
	c.Basic = conf.Basic
	c.bM.Unlock()

	c.aM.Lock()
	c.Advanced = conf.Advanced
	c.aM.Unlock()

	return nil
}

// unmarshalConfig unmatshalls YAML file bytes into Config
func unmarshalConfig(configData []byte) (*YAMLConfig, error) {
	conf := &YAMLConfig{}
	err := yaml.Unmarshal(configData, conf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal config")
	}
	return conf, nil
}

// loadFileData loads YAML config data from file loader int bytes
func loadFileData(configFile string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Clean(configFile))
}
