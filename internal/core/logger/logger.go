package logger

import (
    "go.uber.org/zap"
)

var Logger *zap.Logger

func Init() {
    var err error
    Logger, err = zap.NewProduction()
    if err != nil {
        panic(err)
    }
    defer Logger.Sync() // Flushes buffer, if any
}

func Info(msg string, fields ...zap.Field) {
    Logger.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
    Logger.Error(msg, fields...)
}