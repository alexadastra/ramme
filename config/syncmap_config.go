package config

import (
	"sync"
)

// SyncMapConfig is config storage based on sync.Map
type SyncMapConfig struct {
	basicST    *sync.Map
	advancedST *sync.Map
}

// Get fetches config entry from advanced config
func (c *SyncMapConfig) Get(key Name) Entry {
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
func (c *SyncMapConfig) Set(t *Target) {
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
