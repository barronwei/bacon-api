package logger

import (
	"log"
	"net/http"
	"time"
)

// Logger for http request
func Logger(h http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		log.Println(
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
