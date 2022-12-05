// Package config defines new config handling implementation
package config

var (
	// ServiceName contains a service name prefix which used in ENV variables
	ServiceName = "RAMME-TEMPLATE"
	// File contains path to .yaml config file
	File = "/etc/config/config.yaml"
)

// Config represents the structure that contains configurations both for logic and middleware
type Config interface {
	Get(key Name) Entry
	Set(t *Target)
}
