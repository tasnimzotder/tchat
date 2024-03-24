package middleware

import (
	"log"
	"net/http"
	"time"
)

type wrapperWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrapperWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := &wrapperWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapped, r)

		log.Printf("%d %s %s %v", wrapped.statusCode, r.Method, r.URL.Path, time.Since(start))
	})
}
