package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"order/api"
)

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {

				log.WithFields(log.Fields{
					"host":       r.Host,
					"path":       r.URL.Path,
					"method":     r.Method,
					"statuscode": http.StatusInternalServerError,
					"error":      err,
				}).Error(err)
				api.WriteErrorResponse(w, http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
