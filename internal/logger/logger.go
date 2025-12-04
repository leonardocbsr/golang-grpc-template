// Package logger provides logging configuration and initialization.
package logger

import (
	"github.com/leonardocbsr/golang-grpc-template/internal/config"
	"github.com/leonardocbsr/golang-grpc-template/internal/logger/formatters"
	"github.com/sirupsen/logrus"
)

// New creates a new logger with the configured format and level.
func New(config config.Config) *logrus.Logger {
	_logger := logrus.New()

	switch config.GetLoggerConfig().Format {
	case "json":
		_logger.SetFormatter(formatters.NewJSONFormatter())
	case "text":
		_logger.SetFormatter(formatters.NewTextFormatter())
	case "pretty":
		_logger.SetFormatter(formatters.NewPrettyFormatter())
	default:
		_logger.SetFormatter(formatters.NewJSONFormatter())
	}

	level, err := logrus.ParseLevel(config.GetLoggerConfig().Level)
	if err != nil {
		panic(err)
	}

	_logger.SetLevel(level)
	return _logger
}
