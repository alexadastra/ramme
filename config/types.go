package config

import "git.miem.hse.ru/786/ramme/logger"

var (
	// ServiceName contains a service name prefix which used in ENV variables
	ServiceName = "RAMME-SERVICE-NAME"
	// File contains path to .yaml config file
	File = "/etc/config/config.yaml"
)

// Config is the struct that holds our application's configuration
type Config struct {
	Basic    BasicConfig    `yaml:"basic"`
	Advanced AdvancedConfig `yaml:"advanced"`
}

// BasicConfig holds basic application's configuration
type BasicConfig struct {
	// Local service host
	Host string `split_words:"true" yaml:"host"`
	// Local service GRPC port
	GRPCPort int `split_words:"true" yaml:"grpc_port"`
	// Local service HTTP port
	HTTPPort int `split_words:"true" yaml:"http_port"`
	// Local secondary service HTTP port (for monitoring, tracing, health/readiness check etc.)
	HTTPSecondaryPort int `split_words:"true" yaml:"http_secondary_port"`
	// Logging level in logger.Level notation
	LogLevel logger.Level `split_words:"true" yaml:"log_level"`
	// is local environment
	IsLocalEnvironment bool `split_words:"true" yaml:"is_local_environment"`
}

// AdvancedConfig holds specific application's configuration
type AdvancedConfig struct {
	PingMessage string `yaml:"ping_message"`
}
