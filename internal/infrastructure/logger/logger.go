package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ---------------------------------------------------------------------------------------------------------------------
// Logger

type Logger interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
}

type logger struct {
	zl *zap.Logger
}

func NewLogger() (Logger, error) {
	config := zap.Config{
		Level:       zap.NewAtomicLevel(),
		Development: false,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			LevelKey:       "level",
			NameKey:        "name",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout", "./log/development.out.log"},
		ErrorOutputPaths: []string{"stderr", "./log/development.err.log"},
	}
	zl, err := config.Build()
	if err != nil {
		return nil, err
	}
	return &logger{zl: zl}, nil
}

// ---------------------------------------------------------------------------------------------------------------------
// logging functions

func (l *logger) Debug(msg string, fields ...zap.Field) {
	l.zl.Debug(msg, fields...)
}

func (l *logger) Info(msg string, fields ...zap.Field) {
	l.zl.Info(msg, fields...)
}

func (l *logger) Warn(msg string, fields ...zap.Field) {
	l.zl.Warn(msg, fields...)
}

func (l *logger) Error(msg string, fields ...zap.Field) {
	l.zl.Error(msg, fields...)
}

func (l *logger) Fatal(msg string, fields ...zap.Field) {
	l.zl.Fatal(msg, fields...)
}

// ---------------------------------------------------------------------------------------------------------------------
// fields

func Error(err error) zap.Field {
	return zap.Error(err)
}

func String(key, val string) zap.Field {
	return zap.String(key, val)
}

func Int(key string, val int) zap.Field {
	return zap.Int(key, val)
}
