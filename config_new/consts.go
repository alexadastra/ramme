package config_new

const (
	// Host defines the host app listens requests from
	Host = "host"
	// GRPCPort defines the port app listens gRPC requests from
	GRPCPort = "grpc_port"
	// HTTPPort defines the port app listens http requests from
	HTTPPort = "http_port"
	// HTTPSecondaryPort defines the port app listens admin http requests from
	HTTPSecondaryPort = "http_secondary_port"
	// LogLevel defines log level (ERROR, WARN, INFO, DEBUG)
	LogLevel = "log_level"
	// IsLocalEnvironment defines whether the app serves http by default or custom prefix
	IsLocalEnvironment = "is_local_environment"
)
