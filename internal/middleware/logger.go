package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("[%s] %s %s | %v\n", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
	})
}
