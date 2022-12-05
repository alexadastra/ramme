// Package env contains storage for handling env data
package env

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/pkg/errors"
)

// TODO: should add more complex JSON structure to explicitly declare entries types

// FIXME: cant set duration

// SecretVariable defines suffix for env variable that should contain secrets
const SecretVariable = "SECRET"

// Storage contains data from env
type Storage struct {
	m *sync.Map
}

// New creates new Storage
func New(appName string) (*Storage, error) {
	s := &Storage{m: &sync.Map{}}
	mapping, err := s.ParseSecrets(appName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse environment")
	}
	for name, value := range mapping {
		s.m.Store(name, value)
	}
	return s, nil
}

// Get fetches env value by key
func (s *Storage) Get(key string) (interface{}, bool) {
	return s.m.Load(key)
}

// ParseSecrets parses env stuff into map
func (s *Storage) ParseSecrets(appName string) (map[string]interface{}, error) {
	val, ok := os.LookupEnv(fmt.Sprintf("%s-%s", appName, SecretVariable))
	if !ok {
		return nil, errors.New("failed to find environment variable")
	}

	mapping := map[string]interface{}{}
	if err := json.Unmarshal([]byte(val), &mapping); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshall value into map")
	}

	return mapping, nil
}
