package mylog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var zapLogger *zap.Logger
var zapSugarLogger *zap.SugaredLogger

func init() {
	var err error
	// EncoderConfig
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "severity",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     "\n",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     EncodeTime, //zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	// Config
	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Development:       false,
		DisableStacktrace: true,
		Encoding:          "console", //"json",
		EncoderConfig:     encoderConfig,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stdout"},
	}
	// Logger
	zapLogger, err = config.Build()
	if err != nil {
		panic(err)
	}
	// SugarLogger
	zapSugarLogger = zapLogger.Sugar()
}

func EncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.999999999"))
}
