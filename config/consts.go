package config

// Name is alias to string that points to config entry
type Name string

var (
	// BasicConfigMapping defines config names that are at basic storage
	BasicConfigMapping = map[Name]struct{}{
		Host:                  {},
		GRPCPort:              {},
		HTTPPort:              {},
		HTTPReadTimeout:       {},
		HTTPWriteTimeout:      {},
		HTTPAdminReadTimeout:  {},
		HTTPAdminWriteTimeout: {},
		HTTPAdminPort:         {},
		LogLevel:              {},
		IsLocalEnvironment:    {},
	}

	// AdvancedConfigMapping defines config names that are at advanced storage
	AdvancedConfigMapping = map[Name]struct{}{}
)

const (
	// Host defines the host app listens requests from
	Host Name = "host"
	// GRPCPort defines the port app listens gRPC requests from
	GRPCPort Name = "grpc_port"
	// HTTPPort defines the port app listens http requests from
	HTTPPort Name = "http_port"
	// HTTPReadTimeout defines http read timeout
	HTTPReadTimeout Name = "http_read_timeout"
	// HTTPWriteTimeout defines http write timeout
	HTTPWriteTimeout Name = "http_write_timeout" // nolint
	// HTTPAdminReadTimeout defines http read timeout
	HTTPAdminReadTimeout Name = "http_admin_read_timeout"
	// HTTPAdminWriteTimeout defines http write timeout
	HTTPAdminWriteTimeout Name = "http_admin_write_timeout" // nolint
	// HTTPAdminPort defines the port app listens admin http requests from
	HTTPAdminPort Name = "http_admin_port"
	// LogLevel defines log level (ERROR, WARN, INFO, DEBUG)
	LogLevel Name = "log_level"
	// IsLocalEnvironment defines whether the app serves http by default or custom prefix
	IsLocalEnvironment Name = "is_local_environment"
)
