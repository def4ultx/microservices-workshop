package middleware

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		recorder := NewStatusRecoder(w)

		next.ServeHTTP(recorder, r)
		elapsed := time.Since(start)

		log.WithFields(log.Fields{
			"host":       r.Host,
			"path":       r.URL.Path,
			"method":     r.Method,
			"statuscode": recorder.Status,
			"elapsed":    elapsed.Milliseconds(),
		}).Info("Incoming request")
	})
}
