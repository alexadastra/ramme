package config_new

const (
	// Host defines the host app listens requests from
	Host = "host"
	// GRPCPort defines the port app listens gRPC requests from
	GRPCPort = "grpc_port"
	// HTTPPort defines the port app listens http requests from
	HTTPPort = "http_port"
	// HTTPReadTimeout defines http read timeout
	HTTPReadTimeout = "http_read_timeout"
	// HTTPWriteTimeout defines http write timeout
	HTTPWriteTimeout = "http_write_timeout" // nolint
	// HTTPAdminReadTimeout defines http read timeout
	HTTPAdminReadTimeout = "http_admin_read_timeout"
	// HTTPAdminWriteTimeout defines http write timeout
	HTTPAdminWriteTimeout = "http_admin_write_timeout" // nolint
	// HTTPAdminPort defines the port app listens admin http requests from
	HTTPAdminPort = "http_admin_port"
	// LogLevel defines log level (ERROR, WARN, INFO, DEBUG)
	LogLevel = "log_level"
	// IsLocalEnvironment defines whether the app serves http by default or custom prefix
	IsLocalEnvironment = "is_local_environment"
)
