package freezerscanner

import (
	"testing"
)

// TestSequentialReaderDiscovery tests that we can discover cdat files
// Note: Sequential reading without index is experimental; use indexed reader instead
func TestSequentialReaderDiscovery(t *testing.T) {
	sr, err := NewSequentialReader("../../mainnet")
	if err != nil {
		t.Skipf("Skipping: %v", err)
	}
	defer sr.Close()

	files := sr.CdatFiles()
	t.Logf("Found %d cdat files", len(files))

	if len(files) == 0 {
		t.Error("Expected to find cdat files")
	}

	// Verify files are sorted
	for i := 1; i < len(files); i++ {
		if files[i] < files[i-1] {
			t.Errorf("Files not sorted: %s before %s", files[i-1], files[i])
		}
	}
}

