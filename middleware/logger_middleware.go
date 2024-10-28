package middleware

import (
	"errors"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sidiqPratomo/DJKI-Pengaduan/appconstant"
	"github.com/sidiqPratomo/DJKI-Pengaduan/apperror"
	"github.com/sirupsen/logrus"
)

func Logger(log *logrus.Logger) func(c *gin.Context) {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		if raw != "" {
			path = path + "?" + raw
		}

		statusCode := c.Writer.Status()

		requestId, exist := c.Get(appconstant.RequestId)
		if !exist {
			requestId = ""
		}

		entry := log.WithFields(logrus.Fields{
			"latency":     time.Since(start),
			"method":      c.Request.Method,
			"path":        path,
			"request_id":  requestId,
			"status_code": statusCode,
		})

		if statusCode >= 400 && statusCode <= 599 {
			var appErr *apperror.AppError
			for _, err := range c.Errors {
				if errors.As(err, &appErr) {
					entry.WithFields(logrus.Fields{
						"app error": appErr,
						"stack":     string(appErr.GetStackTrace()),
					}).Error()
				} else {
					entry.WithFields(logrus.Fields{
						"error": err,
						"stack": string(debug.Stack()),
					}).Error()
				}
			}
			return
		}
		entry.Info()
	}
}
