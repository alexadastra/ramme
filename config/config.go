package config

import (
	"git.miem.hse.ru/786/ramme/logger"
	"github.com/kelseyhightower/envconfig"
)

const (
	// SERVICENAME contains a service name prefix which used in ENV variables
	SERVICENAME = "RAMME-SERVICE-NAME"
)

// Config contains ENV variables
type Config struct {
	// Local service host
	Host string `split_words:"true"`
	// Local service GRPC port
	GRPCPort int `split_words:"true"`
	// Local service HTTP port
	HTTPPort int `split_words:"true"`
	// Local secondary service HTTP port (for monitoring, tracing, health/readiness check etc.)
	HTTPSecondaryPort int `split_words:"true"`
	// Logging level in logger.Level notation
	LogLevel logger.Level `split_words:"true"`
}

// Load settles ENV variables into Config structure
func (c *Config) Load(serviceName string) error {
	return envconfig.Process(serviceName, c)
}
