package logger

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	instance *Logger
	once     sync.Once
)

type Logger struct {
	zap *zap.SugaredLogger
}

func NewLogger(level string) *Logger {
	once.Do(func() {
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.TimeKey = "timestamp"
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

		var zapLevel zapcore.Level
		switch level {
		case "debug":
			zapLevel = zapcore.DebugLevel
		case "info":
			zapLevel = zapcore.InfoLevel
		case "warn":
			zapLevel = zapcore.WarnLevel
		case "error":
			zapLevel = zapcore.ErrorLevel
		default:
			zapLevel = zapcore.InfoLevel
		}

		logFile, _ := os.OpenFile("log/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		core := zapcore.NewTee(
			zapcore.NewCore(
				zapcore.NewJSONEncoder(encoderConfig),
				zapcore.AddSync(logFile),
				zapLevel,
			),
			zapcore.NewCore(
				zapcore.NewConsoleEncoder(encoderConfig),
				zapcore.AddSync(os.Stdout),
				zapLevel,
			),
		)

		logger := zap.New(core)
		instance = &Logger{
			zap: logger.Sugar(),
		}
	})

	return instance
}

func (l *Logger) Debug(msg string, keysAndValues ...interface{}) {
	l.zap.Debugw(msg, keysAndValues...)
}

func (l *Logger) Info(msg string, keysAndValues ...interface{}) {
	l.zap.Infow(msg, keysAndValues...)
}

func (l *Logger) Warn(msg string, keysAndValues ...interface{}) {
	l.zap.Warnw(msg, keysAndValues...)
}

func (l *Logger) Error(msg string, keysAndValues ...interface{}) {
	l.zap.Errorw(msg, keysAndValues...)
}

func (l *Logger) Fatal(msg string, keysAndValues ...interface{}) {
	l.zap.Fatalw(msg, keysAndValues...)
}