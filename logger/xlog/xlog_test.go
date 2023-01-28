package xlog

import (
	"os"
	"testing"

	"github.com/alexadastra/ramme/logger"
)

func TestNewXLog(t *testing.T) {
	log1 := NewLogger(&logger.Config{
		Level: logger.LevelDebug,
	})
	if log1 == nil {
		t.Error("Got uninitialized XLog logger")
	}
	log2 := NewLogger(&logger.Config{
		Level: logger.LevelInfo,
		Out:   os.Stdout,
		Err:   os.Stdout,
	})
	if log2 == nil {
		t.Error("Got uninitialized XLog logger")
	}
}
