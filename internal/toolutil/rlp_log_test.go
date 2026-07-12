package toolutil

import (
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func TestLogRLPHelpers(t *testing.T) {
	t.Parallel()
	empty := &types.Log{}
	if Topic0Sig(empty) != "" || Topic0SigHex(empty) != "" {
		t.Fatal("empty topics produced a signature")
	}

	topic := ethcommon.HexToHash("0x1234567890abcdef")
	log := &types.Log{
		Address: ethcommon.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"),
		Topics:  []ethcommon.Hash{topic},
		Data:    []byte{1, 2, 3},
		Index:   4,
	}
	if got := Topic0Sig(log); got != "00000000" {
		t.Fatalf("Topic0Sig = %q", got)
	}
	if got := Topic0SigHex(log); len(got) != 64 ||
		got[len(got)-16:] != "1234567890abcdef" {
		t.Fatalf("Topic0SigHex = %q", got)
	}
	encoded, err := EncodeLogRLP(log)
	if err != nil {
		t.Fatal(err)
	}
	equal, err := LogRLPEqual(encoded, log)
	if err != nil || !equal {
		t.Fatalf("LogRLPEqual = %v, %v", equal, err)
	}
	equal, err = LogRLPEqual(encoded[:len(encoded)-1], log)
	if err != nil || equal {
		t.Fatalf("length mismatch = %v, %v", equal, err)
	}
	changed := append([]byte(nil), encoded...)
	changed[len(changed)-1] ^= 0xff
	equal, err = LogRLPEqual(changed, log)
	if err != nil || equal {
		t.Fatalf("content mismatch = %v, %v", equal, err)
	}
	if got := NormalizeAddr("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"); got != ethcommon.HexToAddress(got).Hex() {
		t.Fatalf("NormalizeAddr = %q", got)
	}
}
