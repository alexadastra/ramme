package config

import (
	"io/ioutil"
	"path/filepath"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// Target represents the DTO objects that is passed from infrastructure (filewatcher, ETCD etc.) to in-memory storage
type Target struct {
	Basic    map[Name]*Entry
	Advanced map[Name]*Entry
}

// newTarget creates new Target
func newTarget(filePath string) (*Target, error) {
	bytes, err := loadFileData(filePath)
	if err != nil {
		return nil, err
	}

	t, err := unmarshalTarget(bytes)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// unmarshalConfig unmatshalls YAML file bytes into Config
func unmarshalTarget(data []byte) (*Target, error) {
	tmp := struct {
		Basic    map[Name]*Entry `yaml:"basic"`
		Advanced map[Name]*Entry `yaml:"advanced"`
	}{}
	err := yaml.Unmarshal(data, &tmp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal config")
	}
	for _, entry := range tmp.Basic {
		if err := entry.Validate(); err != nil {
			return nil, errors.Wrap(err, "failed to validate entry")
		}
	}
	for _, entry := range tmp.Advanced {
		if err := entry.Validate(); err != nil {
			return nil, errors.Wrap(err, "failed to validate entry")
		}
	}
	t := &Target{
		Basic:    tmp.Basic,
		Advanced: tmp.Advanced,
	}

	return t, nil
}

// loadFileData loads YAML config data from file loader int bytes
func loadFileData(configFile string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Clean(configFile))
}
