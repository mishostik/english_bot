package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var loggerLevelMap = map[string]zapcore.Level{
	"debug": zap.DebugLevel,
	"info":  zap.InfoLevel,
	"warn":  zap.WarnLevel,
	"error": zap.ErrorLevel,
	"panic": zap.PanicLevel,
	"fatal": zap.FatalLevel,
}

const runtimeCaller = 1
