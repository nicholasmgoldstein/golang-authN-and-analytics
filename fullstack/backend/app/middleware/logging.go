package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs information about each incoming request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		next.ServeHTTP(w, r)

		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(startTime))
	})
}
