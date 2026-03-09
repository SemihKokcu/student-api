package middleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5/middleware"
)

func StructuredLogger(l *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			next.ServeHTTP(ww, r)

			l.Info("HTTP Request",
				"method", r.Method,
				"path", r.URL.Path,
				"duration", time.Since(start),
				"status", ww.Status(),
				"remote_addr", r.RemoteAddr,
			)
		})
	}
}

// GinStructuredLogger: Gin için slog tabanlı middleware
func GinStructuredLogger(l *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// İsteği işle
		c.Next()

		// İstek bittikten sonra logla
		end := time.Since(start)
		status := c.Writer.Status()

		l.Info("HTTP Request",
			"method", c.Request.Method,
			"path", path,
			"query", query,
			"status", status,
			"duration", end,
			"ip", c.ClientIP(),
			"error", c.Errors.String(),
		)
	}
}
