package config

import (
	"newgo/internal/logger"

	"github.com/rs/zerolog/log"
)

func Init() AppConfig {
	baseLogger := logger.SetupLoggerWriters()
	log.Logger = baseLogger

	LoadEnvOnce()

	log.Logger = logger.ApplyLogLevelFromEnv(baseLogger)
	logger.Success("All services initialized")

	return Config
}

func Shutdown() {
	// ...
}

func SetModule(name string) {
	logger.SetModule(name)
}

func Trace(format string, args ...interface{}) {
	logger.Trace(format, args...)
}

func Debug(format string, args ...interface{}) {
	logger.Debug(format, args...)
}

func Info(format string, args ...interface{}) {
	logger.Info(format, args...)
}

func Success(format string, args ...interface{}) {
	logger.Success(format, args...)
}

func Warn(format string, args ...interface{}) {
	logger.Warn(format, args...)
}

func Fail(format string, args ...interface{}) {
	logger.Fail(format, args...)
}

func Fatal(format string, args ...interface{}) {
	logger.Fatal(format, args...)
}

func Panic(format string, args ...interface{}) {
	logger.Panic(format, args...)
}
