package logger

import (
	"cbsr.io/golang-grpc-template/config"
	"cbsr.io/golang-grpc-template/logger/formatters"
	"github.com/sirupsen/logrus"
)

func New(config config.IConfig) *logrus.Logger {
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
