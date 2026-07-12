package scan

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// ErrorLog appends per-block scan errors to a file, framed by session
// start/end markers. A nil *ErrorLog is a valid no-op sink.
type ErrorLog struct {
	file *os.File
	mu   sync.Mutex
}

// NewErrorLog opens (or creates) the error log and writes the session
// header.
func NewErrorLog(path string) (*ErrorLog, error) {
	f, err := os.OpenFile(filepath.Clean(path), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)
	if err != nil {
		return nil, err
	}
	if _, err := fmt.Fprintf(f, "\n=== Scan session started at %s ===\n", time.Now().Format(time.RFC3339)); err != nil {
		_ = f.Close() // best-effort cleanup on error path
		return nil, err
	}
	return &ErrorLog{file: f}, nil
}

// Log writes one error entry (thread-safe). A failed write must not abort
// the scan, so it is best-effort.
func (el *ErrorLog) Log(blockNum uint64, errType, errMsg string) {
	if el == nil || el.file == nil {
		return
	}
	el.mu.Lock()
	defer el.mu.Unlock()
	_, _ = fmt.Fprintf(el.file, "[%s] block=%d type=%s error=%s\n",
		time.Now().Format("2006-01-02 15:04:05"), blockNum, errType, errMsg)
}

// Close writes the session trailer and closes the file.
func (el *ErrorLog) Close() error {
	if el == nil || el.file == nil {
		return nil
	}
	el.mu.Lock()
	defer el.mu.Unlock()
	// Best-effort trailer; Close reports the real error, if any.
	_, _ = fmt.Fprintf(el.file, "=== Scan session ended at %s ===\n", time.Now().Format(time.RFC3339))
	return el.file.Close()
}
