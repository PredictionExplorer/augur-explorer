package httpx

import (
	"net/http"
)

// ResponseWriter extends http.ResponseWriter with response state used by
// middleware (status/size accounting) and handlers (has anything been
// written yet). The Router wraps every response writer exactly once before
// the middleware chain runs, so any layer may assert to this interface.
type ResponseWriter interface {
	http.ResponseWriter

	// Status is the response status code, or 0 while nothing has been
	// written yet.
	Status() int
	// Written reports whether the header has been sent (explicitly via
	// WriteHeader or implicitly by the first body write).
	Written() bool
	// Size is the number of body bytes written so far.
	Size() int
	// Unwrap returns the underlying http.ResponseWriter, letting
	// http.ResponseController reach optional interfaces (Flusher, sendfile).
	Unwrap() http.ResponseWriter
}

// responseWriter is the concrete ResponseWriter used by the Router.
type responseWriter struct {
	http.ResponseWriter

	status int
	size   int
}

// WrapResponseWriter returns w as a ResponseWriter, wrapping it when needed.
// Wrapping is idempotent: an already-wrapped writer is returned unchanged so
// nested middleware share one status/size record.
func WrapResponseWriter(w http.ResponseWriter) ResponseWriter {
	if rw, ok := w.(ResponseWriter); ok {
		return rw
	}
	return &responseWriter{ResponseWriter: w}
}

func (w *responseWriter) WriteHeader(status int) {
	if w.status == 0 {
		w.status = status
	}
	w.ResponseWriter.WriteHeader(status)
}

func (w *responseWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = http.StatusOK
	}
	n, err := w.ResponseWriter.Write(b)
	w.size += n
	return n, err
}

func (w *responseWriter) Status() int                 { return w.status }
func (w *responseWriter) Written() bool               { return w.status != 0 }
func (w *responseWriter) Size() int                   { return w.size }
func (w *responseWriter) Unwrap() http.ResponseWriter { return w.ResponseWriter }
