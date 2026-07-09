package httpx

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWrapResponseWriterIsIdempotent(t *testing.T) {
	w := httptest.NewRecorder()
	first := WrapResponseWriter(w)
	second := WrapResponseWriter(first)
	if first != second {
		t.Fatal("wrapping an already-wrapped writer must return it unchanged")
	}
	if first.Unwrap() != w {
		t.Fatal("Unwrap must return the original writer")
	}
}

func TestResponseWriterRecordsImplicitStatus(t *testing.T) {
	w := httptest.NewRecorder()
	rw := WrapResponseWriter(w)

	if rw.Written() || rw.Status() != 0 {
		t.Fatalf("fresh writer: Written=%v Status=%d", rw.Written(), rw.Status())
	}

	n, err := rw.Write([]byte("hello"))
	if err != nil || n != 5 {
		t.Fatalf("Write = (%d, %v)", n, err)
	}
	if !rw.Written() || rw.Status() != http.StatusOK {
		t.Errorf("after body write: Written=%v Status=%d, want true/200", rw.Written(), rw.Status())
	}
	if rw.Size() != 5 {
		t.Errorf("Size = %d, want 5", rw.Size())
	}
}

func TestResponseWriterRecordsExplicitStatus(t *testing.T) {
	w := httptest.NewRecorder()
	rw := WrapResponseWriter(w)

	rw.WriteHeader(http.StatusTeapot)
	if rw.Status() != http.StatusTeapot || !rw.Written() {
		t.Errorf("Status=%d Written=%v, want 418/true", rw.Status(), rw.Written())
	}

	// A second WriteHeader must not change the recorded status (net/http
	// ignores superfluous calls; the record mirrors what went on the wire).
	rw.WriteHeader(http.StatusInternalServerError)
	if rw.Status() != http.StatusTeapot {
		t.Errorf("Status after second WriteHeader = %d, want 418", rw.Status())
	}
}
