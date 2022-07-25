package logrus

import (
	"github.com/alexadastra/ramme/logger"
	"github.com/sirupsen/logrus"
)

// New creates "github.com/sirupsen/logrus" logger
func New(config *logger.Config) logger.Logger {
	l := logrus.New()
	l.Level = logrusLevelConverter(config.Level)
	l.WithFields(logrus.Fields(config.Fields))
	return l
}

func logrusLevelConverter(level logger.Level) logrus.Level {
	switch level {
	case logger.LevelDebug:
		return logrus.DebugLevel
	case logger.LevelInfo:
		return logrus.InfoLevel
	case logger.LevelWarn:
		return logrus.WarnLevel
	case logger.LevelError:
		return logrus.ErrorLevel
	case logger.LevelFatal:
		return logrus.FatalLevel
	default:
		return logrus.InfoLevel
	}
}
