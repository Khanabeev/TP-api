package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Config struct {
	LogLevel string
}

var log *zap.Logger

func init() {

	logLevel, exists := os.LookupEnv("LOG_LEVEL")
	if !exists {
		logLevel = "debug"
	}

	logConfig := Config{
		LogLevel: logLevel,
	}

	var err error

	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()

	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encoderConfig
	config.DisableStacktrace = true

	setLevel(logConfig.LogLevel, &config)

	log, err = config.Build(zap.AddCallerSkip(1))

	if err !=nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}

func setLevel(level string, config *zap.Config) {
	switch level {
	case "debug":
		config.Level.SetLevel(zapcore.DebugLevel)
	case "info":
		config.Level.SetLevel(zapcore.InfoLevel)
	case "error":
		config.Level.SetLevel(zapcore.ErrorLevel)
	case "warn":
		config.Level.SetLevel(zapcore.WarnLevel)
	default:
		config.Level.SetLevel(zapcore.InfoLevel)
	}
}