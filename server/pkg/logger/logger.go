package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger *zap.SugaredLogger
)

func NewLogger(env string) {
	var logger *zap.Logger
	if env == "prd" {
		logger, _ = zap.NewProduction()
	} else if env == "test" {
		config := zap.Config{
			Level:    zap.NewAtomicLevelAt(zapcore.PanicLevel),
			Encoding: "json",
			EncoderConfig: zapcore.EncoderConfig{
				TimeKey:        "Time",
				LevelKey:       "Level",
				NameKey:        "Name",
				CallerKey:      "Caller",
				MessageKey:     "Msg",
				StacktraceKey:  "St",
				EncodeLevel:    zapcore.CapitalLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.StringDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			},
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
		}
		logger, _ = config.Build()
	} else {
		logger, _ = zap.NewDevelopment()
	}
	defer logger.Sync() // flushes buffer, if any
	Logger = logger.Sugar()
}

func GetLogger(env string) *zap.SugaredLogger {
	if Logger == nil {
		NewLogger(env)
	}
	return Logger
}

func Info(msg interface{}) {
	Logger.Info(msg)
}

func Debug(msg interface{}) {
	Logger.Debug(msg)
}

func Warn(msg interface{}) {
	Logger.Warn(msg)
}

func Error(msg interface{}) {
	Logger.Error(msg)
}
