package logging

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/h4yfans/case-study/common/environment"
	"go.elastic.co/apm/module/apmzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Initialize() {
	var cfg zap.Config
	if environment.Debug() {
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		cfg = zap.NewProductionConfig()
	}

	cfg.Level.SetLevel(getLogLevel(environment.LogLevel()))
	apm := zap.WrapCore((&apmzap.Core{}).WrapCore)
	log, err := cfg.Build(apm)
	if err != nil {
		panic(fmt.Sprintf("Zap initialization failed: %v", err))
	}

	zap.ReplaceGlobals(log)
}

func Close() {
	defer func() {
		_ = zap.L().Sync()
	}()
	defer func() {
		sentry.Recover()
	}()
}

func getLogLevel(level string) zapcore.Level {
	switch level {
	case "DEBUG":
		return zapcore.DebugLevel
	case "INFO":
		return zapcore.InfoLevel
	case "WARN":
		return zapcore.WarnLevel
	case "WARNING":
		return zapcore.WarnLevel
	case "ERROR":
		return zapcore.ErrorLevel
	default:
		panic(fmt.Sprintf("Unknown log level %v", level))
	}
}
