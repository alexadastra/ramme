package config

import (
	"sync"
	"time"

	"github.com/pkg/errors"
)

const fileWatcherInterval = time.Second

// YAMLConfig fetched from YAML file
type YAMLConfig struct {
	basicST    *sync.Map
	advancedST *sync.Map

	fileWatcher *FileWatcher
}

// NewConfig creates new Config
func NewConfig(path string) (Config, func() error, func() error, error) {
	c := &YAMLConfig{}

	setFunc := func() error {
		t, err := newTarget(path)
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

	c.fileWatcher = fw

	return c, fw.Run, fw.Close, nil
}

// Get fetches config entry from advanced config
func (c *YAMLConfig) Get(key Name) Entry {
	if _, ok := BasicConfigMapping[key]; ok {
		e, ok := c.basicST.Load(key)
		if !ok {
			// TODO: handle error
			panic("basic config entry not found!")
		}
		return e.(Entry)
	}

	if _, ok := AdvancedConfigMapping[key]; ok {
		e, ok := c.advancedST.Load(key)
		if !ok {
			// TODO: handle error
			panic("advanced config entry not found!")
		}
		return e.(Entry)
	}

	return Entry{}
}

// Set sets new config from given file path
func (c *YAMLConfig) Set(t *Target) {
	basicST := &sync.Map{}
	for k, v := range t.Basic {
		basicST.Store(k, v)
	}
	c.basicST = basicST

	advancedST := &sync.Map{}
	for k, v := range t.Advanced {
		advancedST.Store(k, v)
	}
	c.advancedST = advancedST
}
