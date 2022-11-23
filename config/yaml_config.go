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

// YAMLConfig fetched from YAML file
type YAMLConfig struct {
	bM          *sync.Mutex
	aM          *sync.Mutex
	Basic       map[Name]Entry `yaml:"basic"`
	Advanced    map[Name]Entry `yaml:"advanced"`
	fileWatcher *FileWatcher
}

// NewConfig creates new Config
func NewConfig(path string) (Config, func() error, func() error, error) {
	c := &YAMLConfig{
		bM:       &sync.Mutex{},
		aM:       &sync.Mutex{},
		Basic:    make(map[Name]Entry),
		Advanced: make(map[Name]Entry),
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

// Get fetches config entry from advanced config
func (c *YAMLConfig) Get(key Name) Entry {
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
