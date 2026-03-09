package middleware

import (
	"log/slog"
	"net/http"
	"time"

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
