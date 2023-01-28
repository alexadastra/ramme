// Package config contains additional info about app configuration
package config

import "github.com/alexadastra/ramme/config"

// SetConfig sets config
func SetConfig() {
	config.AdvancedConfigMapping = map[config.Name]struct{}{
		PingMessage: {},
	}
}

const (
	// PingMessage is config exmple
	PingMessage config.Name = "ping_message"
)
