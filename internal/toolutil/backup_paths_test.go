package toolutil

import "testing"

func TestIsHexTxHash(t *testing.T) {
	h := "0x8aaaa2ba9c9abd092feabe1129658361bee156c00d542d8a59297fa2dc2b45be"
	if !IsHexTxHash(h) {
		t.Fatalf("expected valid tx hash")
	}
	if IsHexTxHash("0x6a714Ae7B5b6eA520F6BCA23d2E609C4Fd5863F2") {
		t.Fatalf("address is not a tx hash")
	}
}
