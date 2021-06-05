package middleware

import "net/http"

type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *StatusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func NewStatusRecoder(w http.ResponseWriter) *StatusRecorder {
	return &StatusRecorder{
		ResponseWriter: w,
		Status:         http.StatusOK,
	}
}
