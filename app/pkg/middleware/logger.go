// Reference: https://github.com/gin-contrib/zap

package middleware

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ZapLogger 接收访问日志
func ZapLogger(logger *zap.Logger, skipPath []string) gin.HandlerFunc {
	skipPaths := make(map[string]bool, len(skipPath))
	for _, path := range skipPath {
		skipPaths[path] = true
	}

	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		if _, ok := skipPaths[path]; !ok {
			latency := time.Since(start)

			if len(c.Errors) > 0 {
				for _, err := range c.Errors.Errors() {
					logger.Error(err)
				}
			} else {
				logger.Info(
					path,
					zap.Int("status", c.Writer.Status()),
					zap.String("method", c.Request.Method),
					zap.String("path", path),
					zap.String("query", query),
					zap.String("id", c.ClientIP()),
					zap.String("user-agent", c.Request.UserAgent()),
					zap.Duration("latency", latency),
				)
			}
		}
	}
}

// RecoveryWithZap 接收panic日志并恢复应用
func RecoveryWithZap(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// check for a broken connection, as it is not really a condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
							strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(
						c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)

					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error))
					c.Abort()
					return
				}

				if stack {
					logger.Error(
						"[Recovery from panic]",
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error(
						"[Recovery from panic]",
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
