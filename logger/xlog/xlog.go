// Package xlog contains xlog logger implementation
package xlog

import (
	"os"

	"github.com/alexadastra/ramme/logger"
	"github.com/rs/xlog"
)

// NewLogger creates "github.com/rs/xlog" logger
func NewLogger(config *logger.Config) logger.Logger {
	var out xlog.Output
	switch config.Err {
	// We should find more matches between types of output
	case nil, os.Stderr:
		out = xlog.NewConsoleOutput()
	default:
		out = xlog.NewConsoleOutput()
	}
	return xlog.New(xlog.Config{
		Level:  xlog.Level(config.Level),
		Fields: config.Fields,
		Output: out,
	})
}
