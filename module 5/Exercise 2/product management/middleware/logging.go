package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		rr := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(rr, r)
		endTime := time.Now()
		log.Printf("Method: %s, Path: %s, Status: %d, Duration: %v",
			r.Method, r.URL.Path, rr.statusCode, endTime.Sub(startTime))
	})
}