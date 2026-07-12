package toolutil

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestIsHexTxHash(t *testing.T) {
	h := "0x8aaaa2ba9c9abd092feabe1129658361bee156c00d542d8a59297fa2dc2b45be"
	if !IsHexTxHash(h) {
		t.Fatalf("expected valid tx hash")
	}
	if IsHexTxHash("0x6a714Ae7B5b6eA520F6BCA23d2E609C4Fd5863F2") {
		t.Fatalf("address is not a tx hash")
	}
}

func TestBackupPathsAndDiscovery(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	hash := "0x8aaaa2ba9c9abd092feabe1129658361bee156c00d542d8a59297fa2dc2b45be"
	txPath := TxRLPPath(dir, 42, hash)
	receiptPath := ReceiptRLPPath(dir, 42, hash)
	if txPath != filepath.Join(dir, "42", hash+"_tx.rlp") ||
		receiptPath != filepath.Join(dir, "42", hash+"_receipt.rlp") {
		t.Fatalf("paths = %q / %q", txPath, receiptPath)
	}
	if err := os.MkdirAll(filepath.Dir(txPath), 0o750); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(txPath, []byte("tx"), 0o600); err != nil {
		t.Fatal(err)
	}
	for _, ignored := range []string{
		filepath.Join(dir, "42", hash+"_receipt.rlp"),
		filepath.Join(dir, "bad", hash+"_tx.rlp"),
		filepath.Join(dir, "42", "bad_tx.rlp"),
		filepath.Join(dir, "42", "nested", hash+"_tx.rlp"),
	} {
		if err := os.MkdirAll(filepath.Dir(ignored), 0o750); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(ignored, []byte("ignored"), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	found, err := BackupTxOnDisk(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(found) != 1 || found[NormalizeTxHash(hash)] != 42 {
		t.Fatalf("found = %#v", found)
	}
	if _, err := BackupTxOnDisk(filepath.Join(dir, "missing")); err == nil {
		t.Fatal("missing backup directory did not fail")
	}
}

func TestParseBackupFileRejectsUnexpectedShapes(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	for _, path := range []string{
		filepath.Join(dir, "42", "file.txt"),
		filepath.Join(dir, "bad", "0x"+strings.Repeat("0", 64)+"_tx.rlp"),
		filepath.Join(dir, "42", "bad_tx.rlp"),
		filepath.Join(dir, "42", "nested", "bad_tx.rlp"),
	} {
		if _, _, ok := parseBackupFile(path, dir, "_tx.rlp"); ok {
			t.Errorf("parseBackupFile(%q) succeeded", path)
		}
	}
}
