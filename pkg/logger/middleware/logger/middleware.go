package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/santirufiner/slicerwise/pkg/logger/context"
	"github.com/sirupsen/logrus"
)

type LoggerMiddleware struct {
	l *logrus.Logger
}

func NewLoggerMiddleware(l *logrus.Logger) *LoggerMiddleware {
	return &LoggerMiddleware{
		l: l,
	}
}

func (lm *LoggerMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		newLogger := logrus.New()
		newLogger.SetLevel(lm.l.GetLevel())
		newLogger.SetOutput(lm.l.Out)
		newLogger.SetReportCaller(false)
		newLogger.SetFormatter(&logrus.TextFormatter{})
		newLogger.AddHook(newLogHook())
		c.Request = c.Request.WithContext(context.SetLogger(c.Request.Context(), newLogger))
	}
}
