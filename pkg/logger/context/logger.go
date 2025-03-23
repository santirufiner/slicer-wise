package context

import (
	"context"
	"github.com/sirupsen/logrus"
)

type key string

const loggerKey = key("logger")

func SetLogger(ctx context.Context, l *logrus.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, l)
}

func GetLogger(ctx context.Context) *logrus.Logger {
	if l, ok := ctx.Value(loggerKey).(*logrus.Logger); ok {
		return l
	}

	newLogger := logrus.StandardLogger()
	newLogger.Errorf("logger not found in context. Using standard logger")
	return newLogger
}
