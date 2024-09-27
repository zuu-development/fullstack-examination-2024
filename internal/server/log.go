package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

func writeRequestLogJSON(_ echo.Context, v middleware.RequestLoggerValues) error {
	log.WithFields(log.Fields{
		"method":         v.Method,
		"host":           v.Host,
		"path":           v.URIPath,
		"remote_ip":      v.RemoteIP,
		"user_agent":     v.UserAgent,
		"protocol":       v.Protocol,
		"status":         v.Status,
		"latency":        v.Latency,
		"content_length": v.ContentLength, // ContentLengthの型はstringで、GETの場合は空文字列
		"response_size":  v.ResponseSize,
	}).Info("finished")
	return nil
}

func requestLogger() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogValuesFunc:    writeRequestLogJSON,
		LogMethod:        true,
		LogHost:          true,
		LogURIPath:       true,
		LogRemoteIP:      true,
		LogUserAgent:     true,
		LogProtocol:      true,
		LogStatus:        true,
		LogLatency:       true,
		LogContentLength: true,
		LogResponseSize:  true,
	})
}
