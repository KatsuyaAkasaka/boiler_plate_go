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
	var conf zap.Config
	var err error
	// var option zap.Option
	conf = zap.NewDevelopmentConfig()
	if env == "prd" {
		conf = zap.NewProductionConfig()
		// Infoからtraceは表示させる
		option := zap.AddStacktrace(zapcore.InfoLevel)
		logger, err = conf.Build(option)
		if err != nil {
			panic(err)
		}
	} else if env == "test" {
		// テスト中はログを出力させない
		conf.Level = zap.NewAtomicLevelAt(zapcore.FatalLevel)
		// conf.DisableStacktrace = true
		logger, err = conf.Build()
		if err != nil {
			panic(err)
		}
	} else {
		conf = zap.NewDevelopmentConfig()
		option := zap.AddStacktrace(zapcore.InfoLevel)
		logger, err = conf.Build(option)
		if err != nil {
			panic(err)
		}
	}
	defer logger.Sync()
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

func Infof(msg string, arg ...interface{}) {
	Logger.Infof(msg, arg...)
}

func Debug(msg interface{}) {
	Logger.Debug(msg)
}

func Debugf(msg string, arg ...interface{}) {
	Logger.Debugf(msg, arg...)
}

func Warn(msg interface{}) {
	Logger.Warn(msg)
}

func Warnf(msg string, arg ...interface{}) {
	Logger.Warnf(msg, arg...)
}

func Error(msg interface{}) {
	Logger.Error(msg)
}

func Errorf(msg string, arg ...interface{}) {
	Logger.Errorf(msg, arg...)
}

func Fatal(msg interface{}) {
	Logger.Fatal(msg)
}

func Fatalf(msg string, arg ...interface{}) {
	Logger.Fatalf(msg, arg...)
}

func Panic(msg interface{}) {
	Logger.Panic(msg)
}
