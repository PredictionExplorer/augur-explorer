package freezerscanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestUint48ToBytes(t *testing.T) {
	tests := []struct {
		input    uint64
		expected []byte
	}{
		{0, []byte{0, 0, 0, 0, 0, 0}},
		{1, []byte{0, 0, 0, 0, 0, 1}},
		{255, []byte{0, 0, 0, 0, 0, 255}},
		{256, []byte{0, 0, 0, 0, 1, 0}},
		{0x123456789ABC, []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC}},
		{0xFFFFFFFFFFFF, []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
	}

	for _, tt := range tests {
		got := Uint48ToBytes(tt.input)
		if len(got) != 6 {
			t.Errorf("Uint48ToBytes(%d) returned %d bytes, want 6", tt.input, len(got))
			continue
		}
		for i := range got {
			if got[i] != tt.expected[i] {
				t.Errorf("Uint48ToBytes(%d) = %v, want %v", tt.input, got, tt.expected)
				break
			}
		}
	}
}

func TestIndexEntryParsing(t *testing.T) {
	// Create a temporary directory with mock index file
	tmpDir := t.TempDir()
	cidxPath := filepath.Join(tmpDir, "receipts.cidx")
	cdatDir := filepath.Join(tmpDir, "ancient")
	os.MkdirAll(cdatDir, 0755)

	// Create mock cidx file with 10 entries
	// Each entry is 6 bytes, offsets: 0, 10, 20, 30, 40, 50, 60, 70, 80, 90
	var cidxData []byte
	for i := 0; i < 10; i++ {
		offset := uint64(i * 10)
		cidxData = append(cidxData, Uint48ToBytes(offset)...)
	}
	if err := os.WriteFile(cidxPath, cidxData, 0644); err != nil {
		t.Fatalf("Failed to create cidx file: %v", err)
	}

	// Create mock cdat file
	cdatPath := filepath.Join(cdatDir, "receipts.0000.cdat")
	cdatData := make([]byte, 100)
	for i := range cdatData {
		cdatData[i] = byte(i)
	}
	if err := os.WriteFile(cdatPath, cdatData, 0644); err != nil {
		t.Fatalf("Failed to create cdat file: %v", err)
	}

	// Open reader
	reader, err := NewFreezerReader(tmpDir)
	if err != nil {
		t.Fatalf("Failed to create reader: %v", err)
	}
	defer reader.Close()

	// Verify item count
	if reader.ItemCount() != 10 {
		t.Errorf("ItemCount() = %d, want 10", reader.ItemCount())
	}

	// Verify offsets
	for i := uint64(0); i < 10; i++ {
		offset, err := reader.ReadOffset(i)
		if err != nil {
			t.Errorf("ReadOffset(%d) error: %v", i, err)
			continue
		}
		expected := i * 10
		if offset != expected {
			t.Errorf("ReadOffset(%d) = %d, want %d", i, offset, expected)
		}
	}

	// Test out of bounds
	_, err = reader.ReadOffset(10)
	if err == nil {
		t.Error("ReadOffset(10) should return error for out of bounds")
	}
}

func TestValidateIndexRange(t *testing.T) {
	tmpDir := t.TempDir()
	cidxPath := filepath.Join(tmpDir, "receipts.cidx")
	cdatDir := filepath.Join(tmpDir, "ancient")
	os.MkdirAll(cdatDir, 0755)

	// Create valid monotonic index
	var cidxData []byte
	offsets := []uint64{0, 10, 20, 30, 40, 50}
	for _, offset := range offsets {
		cidxData = append(cidxData, Uint48ToBytes(offset)...)
	}
	if err := os.WriteFile(cidxPath, cidxData, 0644); err != nil {
		t.Fatalf("Failed to create cidx file: %v", err)
	}

	// Create mock cdat file
	cdatPath := filepath.Join(cdatDir, "receipts.0000.cdat")
	if err := os.WriteFile(cdatPath, make([]byte, 100), 0644); err != nil {
		t.Fatalf("Failed to create cdat file: %v", err)
	}

	reader, err := NewFreezerReader(tmpDir)
	if err != nil {
		t.Fatalf("Failed to create reader: %v", err)
	}
	defer reader.Close()

	// Validate should succeed for monotonic offsets
	if err := reader.ValidateIndexRange(0, 6); err != nil {
		t.Errorf("ValidateIndexRange failed for valid data: %v", err)
	}
}

func TestReadItem(t *testing.T) {
	tmpDir := t.TempDir()
	cidxPath := filepath.Join(tmpDir, "receipts.cidx")
	cdatDir := filepath.Join(tmpDir, "ancient")
	os.MkdirAll(cdatDir, 0755)

	// Create index with offsets: 0, 5, 15, 30
	var cidxData []byte
	offsets := []uint64{0, 5, 15, 30}
	for _, offset := range offsets {
		cidxData = append(cidxData, Uint48ToBytes(offset)...)
	}
	if err := os.WriteFile(cidxPath, cidxData, 0644); err != nil {
		t.Fatalf("Failed to create cidx file: %v", err)
	}

	// Create cdat file with data
	// Block 0: bytes 0-4 (5 bytes): [0,1,2,3,4]
	// Block 1: bytes 5-14 (10 bytes): [5,6,7,8,9,10,11,12,13,14]
	// Block 2: bytes 15-29 (15 bytes): [15,16,...,29]
	cdatPath := filepath.Join(cdatDir, "receipts.0000.cdat")
	cdatData := make([]byte, 30)
	for i := range cdatData {
		cdatData[i] = byte(i)
	}
	if err := os.WriteFile(cdatPath, cdatData, 0644); err != nil {
		t.Fatalf("Failed to create cdat file: %v", err)
	}

	reader, err := NewFreezerReader(tmpDir)
	if err != nil {
		t.Fatalf("Failed to create reader: %v", err)
	}
	defer reader.Close()

	// Test reading each block
	tests := []struct {
		block    uint64
		expected []byte
	}{
		{0, []byte{0, 1, 2, 3, 4}},
		{1, []byte{5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
		{2, []byte{15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29}},
	}

	for _, tt := range tests {
		data, err := reader.ReadItem(tt.block)
		if err != nil {
			t.Errorf("ReadItem(%d) error: %v", tt.block, err)
			continue
		}
		if len(data) != len(tt.expected) {
			t.Errorf("ReadItem(%d) returned %d bytes, want %d", tt.block, len(data), len(tt.expected))
			continue
		}
		for i := range data {
			if data[i] != tt.expected[i] {
				t.Errorf("ReadItem(%d)[%d] = %d, want %d", tt.block, i, data[i], tt.expected[i])
				break
			}
		}
	}
}

func TestEmptyBlock(t *testing.T) {
	tmpDir := t.TempDir()
	cidxPath := filepath.Join(tmpDir, "receipts.cidx")
	cdatDir := filepath.Join(tmpDir, "ancient")
	os.MkdirAll(cdatDir, 0755)

	// Create index with repeated offset (empty block at index 1)
	// Block 0: offset 0-5 (5 bytes)
	// Block 1: offset 5-5 (0 bytes - empty!)
	// Block 2: offset 5-10 (5 bytes)
	var cidxData []byte
	offsets := []uint64{0, 5, 5, 10}
	for _, offset := range offsets {
		cidxData = append(cidxData, Uint48ToBytes(offset)...)
	}
	if err := os.WriteFile(cidxPath, cidxData, 0644); err != nil {
		t.Fatalf("Failed to create cidx file: %v", err)
	}

	cdatPath := filepath.Join(cdatDir, "receipts.0000.cdat")
	if err := os.WriteFile(cdatPath, make([]byte, 15), 0644); err != nil {
		t.Fatalf("Failed to create cdat file: %v", err)
	}

	reader, err := NewFreezerReader(tmpDir)
	if err != nil {
		t.Fatalf("Failed to create reader: %v", err)
	}
	defer reader.Close()

	// Block 1 should return empty data (nil or empty slice)
	data, err := reader.ReadItem(1)
	if err != nil {
		t.Errorf("ReadItem(1) error: %v", err)
	}
	if len(data) != 0 {
		t.Errorf("ReadItem(1) returned %d bytes, want 0 (empty block)", len(data))
	}

	// Block 2 should return 5 bytes
	data2, err := reader.ReadItem(2)
	if err != nil {
		t.Errorf("ReadItem(2) error: %v", err)
	}
	if len(data2) != 5 {
		t.Errorf("ReadItem(2) returned %d bytes, want 5", len(data2))
	}
}

