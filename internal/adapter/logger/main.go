package logger

import (
	"kanban-bot/internal/core/interfaces"

	"go.uber.org/zap"
)

type logger struct {
	zap *zap.SugaredLogger
}

func New(zap *zap.SugaredLogger) interfaces.Logger {
	return &logger{
		zap: zap,
	}
}

func (l *logger) Error(err error, args ...interface{}) {
	l.zap.Error(err, args)
}

func (l *logger) Info(msg string, args ...interface{}) {
	l.zap.Infow(msg, args...)
}

func (l *logger) Warn(msg string, args ...interface{}) {
	l.zap.Warnw(msg, args...)
}
