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
		TimeKey:        "TK",
		LevelKey:       "LK",
		NameKey:        "NK",
		CallerKey:      "CK",
		MessageKey:     "MK",
		StacktraceKey:  "SK",
		LineEnding:     "\n",
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     EncodeTime, //zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
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
		DisableCaller:     false,
	}
	// Logger
	zapLogger, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	// SugarLogger
	zapSugarLogger = zapLogger.Sugar()
}

func EncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.999999999"))
}
