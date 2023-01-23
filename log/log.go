package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(logLevel string) {

	var zapLogLevel zapcore.Level

	switch logLevel {
	case "debug":
		zapLogLevel = zap.DebugLevel
	default:
		zapLogLevel = zap.InfoLevel
	}

	zapConfig := zap.NewDevelopmentConfig()
	zapConfig.Level.SetLevel(zapLogLevel)
	logger, _ := zapConfig.Build()
	zap.ReplaceGlobals(logger)
}
