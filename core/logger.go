package core

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.SugaredLogger
)

type Logger interface {
	DebugLogger
	InfoLogger
	ErrorLogger
}

type DebugLogger interface {
	Debug(message string)
	Debugf(template string, args ...interface{})
}

type InfoLogger interface {
	Info(message string)
	Infof(template string, args ...interface{})
}

type ErrorLogger interface {
	Error(message string)
	Errorf(template string, args ...interface{})
}

func init() {
	zapConfig := zap.NewDevelopmentConfig()

	zapConfig.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	zapLogger, _ := zapConfig.Build()

	logger = zapLogger.Sugar()
}

func Debug(args ...interface{}) {
	logger.Debug(args)
}

func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args)
}
