package middlewares

import (
	"ddCode-server/global"
	logger "ddCode-server/pkg/zlogs"
	"github.com/gin-gonic/gin"
	"time"
)

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		clientIP := c.ClientIP()
		statusCode := c.Writer.Status()
		reqMethod := c.Request.Method
		c.Next()
		cost := time.Since(start)

		global.Logger.Info(path,
			logger.MakeField("status", statusCode),
			logger.MakeField("method", reqMethod),
			logger.MakeField("path", path),
			logger.MakeField("query", query),
			logger.MakeField("ip", clientIP),
			logger.MakeField("user-agent", c.Request.UserAgent()),
			logger.MakeField("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			logger.MakeField("cost", cost),
		)
	}
}
