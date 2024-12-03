package log

import (
	"go.uber.org/zap"
)

var (
	exportUseLogger      *zap.Logger
	exportUseSugarLogger *zap.SugaredLogger
)

func Debugf(template string, args ...interface{}) {
	exportUseSugarLogger.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	exportUseSugarLogger.Infof(template, args...)
}

func Fatal(msg string, fields ...zap.Field) {
	exportUseLogger.Fatal(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	exportUseLogger.Info(msg, fields...)
}
