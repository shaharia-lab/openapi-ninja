// Package logging provides logging functionality for the application
package logging

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/shahariaazam/openapi-ninja/pkg/config"
)

// NewLogger build new logger for the application
func NewLogger(cfg config.Config) *logrus.Logger {
	logger := logrus.New()

	if cfg.RunningOnAppEngine {
		logger.Out = &AppEngineLogWriter{}
	}

	// Set the logger's formatter to a JSON formatter
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Set the logger's output to stdout
	logger.SetOutput(os.Stdout)

	return logger
}
