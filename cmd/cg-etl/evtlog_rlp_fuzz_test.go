// Fuzz target for the stored evt_log RLP decode path (MODERNIZATION.md §4.4).
// process_single_event feeds evt_log.log_rlp bytes into rlp.DecodeBytes; the
// bytes come from the database, so a corrupt row must never be able to panic
// the decoder itself (process_single_event reports a decode failure as a
// returned error, which aborts the batch).
package main

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

func FuzzEvtlogRLP(f *testing.F) {
	good, err := rlp.EncodeToBytes(&types.Log{
		Address: common.HexToAddress("0x1234567890123456789012345678901234567890"),
		Topics: []common.Hash{
			common.HexToHash("0x" + BID_EVENT),
			common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001"),
		},
		Data: []byte{0x01, 0x02, 0x03},
	})
	if err != nil {
		f.Fatalf("encode seed log: %v", err)
	}
	f.Add(good)
	f.Add([]byte(nil))
	f.Add([]byte{0xc0})
	f.Add([]byte{0xf8, 0x00})
	f.Add(bytes.Repeat([]byte{0xff}, 64))
	f.Fuzz(func(t *testing.T, data []byte) {
		var log types.Log
		if err := rlp.DecodeBytes(data, &log); err != nil {
			return
		}
		// Decoded logs must survive an encode/decode round trip unchanged in
		// their consensus fields (the guarantee the archive replay relies on).
		reencoded, err := rlp.EncodeToBytes(&log)
		if err != nil {
			t.Fatalf("re-encode of decoded log failed: %v", err)
		}
		var again types.Log
		if err := rlp.DecodeBytes(reencoded, &again); err != nil {
			t.Fatalf("decode of re-encoded log failed: %v", err)
		}
		if again.Address != log.Address || !bytes.Equal(again.Data, log.Data) || len(again.Topics) != len(log.Topics) {
			t.Fatalf("round trip changed log: %+v vs %+v", log, again)
		}
		for i := range log.Topics {
			if again.Topics[i] != log.Topics[i] {
				t.Fatalf("round trip changed topic %d", i)
			}
		}
	})
}
