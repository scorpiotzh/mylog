package mylog

import (
	"fmt"
	"runtime/debug"
)

const (
	LevelNotice = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelPanic
	LevelFatal
)

const (
	colorRed = iota + 91
	colorGreen
	colorYellow
	colorBlue
	colorMagenta
	colorPanic
	colorFatal
)

type logger struct {
	name  string
	level int
}

func NewLogger(name string, level int) *logger {
	return &logger{
		name:  name,
		level: level,
	}
}

func (l *logger) ErrStack() {
	l.Warn(string(debug.Stack()))
}

func (l *logger) Noticef(format string, a ...interface{}) {
	l.Notice(fmt.Sprintf(format, a...))
}

func (l *logger) Debugf(format string, a ...interface{}) {
	l.Debug(fmt.Sprintf(format, a...))
}

func (l *logger) Infof(format string, a ...interface{}) {
	l.Info(fmt.Sprintf(format, a...))
}

func (l *logger) Warnf(format string, a ...interface{}) {
	l.Warn(fmt.Sprintf(format, a...))
}

func (l *logger) Errorf(format string, a ...interface{}) {
	l.Error(fmt.Sprintf(format, a...))
}

func (l *logger) Panicf(format string, a ...interface{}) {
	l.Panic(fmt.Sprintf(format, a...))
}

func (l *logger) Fatalf(format string, a ...interface{}) {
	l.Fatal(fmt.Sprintf(format, a...))
}

func (l *logger) Notice(a ...interface{}) {
	if l.level > LevelNotice {
		return
	}
	msg := fmt.Sprintf("\x1b[%dm▶ [%s] %s\x1b[0m", colorGreen, l.name, fmt.Sprint(a...))
	zapSugarLogger.Debug(msg)
}

func (l *logger) Debug(a ...interface{}) {
	if l.level > LevelDebug {
		return
	}
	msg := fmt.Sprintf("\x1b[%dm▶ [%s] %s\x1b[0m", colorBlue, l.name, fmt.Sprint(a...))
	zapSugarLogger.Debug(msg)
}

func (l *logger) Info(a ...interface{}) {
	if l.level > LevelInfo {
		return
	}
	msg := fmt.Sprintf("\x1b[%dm▶ [%s] %s\x1b[0m", colorYellow, l.name, fmt.Sprint(a...))
	zapSugarLogger.Info(msg)
}

func (l *logger) Warn(a ...interface{}) {
	if l.level > LevelWarn {
		return
	}
	msg := fmt.Sprintf("\x1b[%dm▶ [%s] %s\x1b[0m", colorMagenta, l.name, fmt.Sprint(a...))
	zapSugarLogger.Warn(msg)
}

func (l *logger) Error(a ...interface{}) {
	if l.level > LevelError {
		return
	}
	msg := fmt.Sprintf("\x1b[%dm▶ [%s] %s\x1b[0m", colorRed, l.name, fmt.Sprint(a...))
	zapSugarLogger.Error(msg)
}

func (l *logger) Panic(a ...interface{}) {
	if l.level > LevelPanic {
		return
	}
	msg := fmt.Sprintf("\x1b[%dm▶ [%s] %s\x1b[0m", colorPanic, l.name, fmt.Sprint(a...))
	zapSugarLogger.Panic(msg)
}

func (l *logger) Fatal(a ...interface{}) {
	if l.level > LevelFatal {
		return
	}
	msg := fmt.Sprintf("\x1b[%dm▶ [%s] %s\x1b[0m", colorFatal, l.name, fmt.Sprint(a...))
	zapSugarLogger.Fatal(msg)
}
