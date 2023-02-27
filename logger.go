package xins

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.SugaredLogger
)

func init() {
	zapConfig := zap.NewDevelopmentConfig()

	zapConfig.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	zapLogger, _ := zapConfig.Build()

	logger = zapLogger.Sugar()
}

// func Logger() *zap.SugaredLogger {
// 	return logger
// }

var _ Logger = logger

// TODO
type Logger interface {
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Errorf(template string, args ...interface{})
}

func DefaultLogger() Logger {
	return logger
}
