package logging

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// 2016-09-27 09:38:21.541541811 +0200 CEST
// 127.0.0.1 - frank [10/Oct/2000:13:55:36 -0700]
// "GET /apache_pb.gif HTTP/1.0" 200 2326
// "http://www.example.com/start.html"
// "Mozilla/4.08 [en] (Win98; I ;Nav)"

// Logger is the logrus logger handler
func Logger(logger *logrus.Logger, notLogged ...string) gin.HandlerFunc {
	var skip map[string]struct{}

	if length := len(notLogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, p := range notLogged {
			skip[p] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		// other handler can change c.Path so:
		path := c.Request.URL.Path
		c.Next()
		statusCode := c.Writer.Status()
		if _, ok := skip[path]; ok {
			return
		}

		entry := logger.WithFields(logrus.Fields{
			"method":      c.Request.Method,
			"path":        path,
			"status_code": statusCode,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			if statusCode >= http.StatusInternalServerError {
				entry.Error()
			} else if statusCode >= http.StatusBadRequest {
				entry.Warn()
			} else {
				entry.Info()
			}
		}
	}
}
