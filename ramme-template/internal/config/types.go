package config

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Config is the struct that holds our application's configuration
type Config struct {
	Advanced AdvancedConfig `yaml:"advanced"`
}

// AdvancedConfig holds advanced application's configuration
type AdvancedConfig struct {
	PingMessage string `yaml:"ping_message"`
}

// UnmarshalConfig unmarshalls file bytes to advanced config
func UnmarshalConfig(configData []byte) (*Config, error) {
	conf := &Config{}
	err := yaml.Unmarshal(configData, conf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal config")
	}
	return conf, nil
}
