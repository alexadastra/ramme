package config

import "flag"

// TODO: figure out how to pass args more easy

// Args contains data to acess config and start up the app
type Args struct {
	ServiceName string
	ConfigPath  string
}

// ParseFlags parses flags into struct
func ParseFlags() *Args {
	var a Args
	flag.StringVar(&a.ServiceName, "name", "RAMME-TEMPLATE", "defines service name")
	flag.StringVar(&a.ConfigPath, "config_path", "/etc/config/config.yaml",
		"defines the path where the service reads config from")
	flag.Parse()
	return &a
}
