//go:build integration

package apitest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

// stubBlockJSON is a minimal but fully valid Ethereum block header (all fields
// go-ethereum requires to unmarshal types.Header), pinned to block 0x64 with
// timestamp 0x6955b900 (1767225600 = 2026-01-01 00:00:00 UTC), so
// time-derived responses snapshot deterministically.
const stubBlockJSON = `{
	"parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
	"sha3Uncles": "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
	"miner": "0x0000000000000000000000000000000000000000",
	"stateRoot": "0x0000000000000000000000000000000000000000000000000000000000000000",
	"transactionsRoot": "0x0000000000000000000000000000000000000000000000000000000000000000",
	"receiptsRoot": "0x0000000000000000000000000000000000000000000000000000000000000000",
	"logsBloom": "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
	"difficulty": "0x1",
	"number": "0x64",
	"gasLimit": "0x1c9c380",
	"gasUsed": "0x0",
	"timestamp": "0x6955b900",
	"extraData": "0x",
	"mixHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
	"nonce": "0x0000000000000000",
	"hash": "0x00000000000000000000000000000000000000000000000000000000000000b0"
}`

type rpcRequest struct {
	ID     json.RawMessage   `json:"id"`
	Method string            `json:"method"`
	Params []json.RawMessage `json:"params"`
}

// newEthStub returns a deterministic JSON-RPC server standing in for an
// Ethereum node. It answers just enough for the API server to initialize and
// for DB-backed routes to respond deterministically:
//
//   - eth_getBalance        -> 0 wei
//   - eth_getBlockByNumber  -> fixed header (timestamp fixedBlockTimestamp)
//   - eth_getCode           -> 0x (no code)
//   - everything else       -> a fixed JSON-RPC error
//
// Contract reads (eth_call) intentionally fail with a stable message: their
// happy paths depend on live chain state and are pinned in the "RPC
// unavailable" shape until Phase 2 makes contract state injectable.
func newEthStub() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var reqs []rpcRequest
		single := false

		body := json.NewDecoder(r.Body)
		var raw json.RawMessage
		if err := body.Decode(&raw); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		if len(raw) > 0 && raw[0] == '[' {
			if err := json.Unmarshal(raw, &reqs); err != nil {
				http.Error(w, "bad batch", http.StatusBadRequest)
				return
			}
		} else {
			var one rpcRequest
			if err := json.Unmarshal(raw, &one); err != nil {
				http.Error(w, "bad request", http.StatusBadRequest)
				return
			}
			reqs = []rpcRequest{one}
			single = true
		}

		resps := make([]map[string]any, 0, len(reqs))
		for _, req := range reqs {
			resp := map[string]any{"jsonrpc": "2.0", "id": req.ID}
			switch req.Method {
			case "eth_getBalance":
				resp["result"] = "0x0"
			case "eth_getCode":
				resp["result"] = "0x"
			case "eth_getBlockByNumber":
				resp["result"] = json.RawMessage(stubBlockJSON)
			default:
				resp["error"] = map[string]any{
					"code":    -32601,
					"message": "apitest eth stub: method " + req.Method + " not supported",
				}
			}
			resps = append(resps, resp)
		}

		w.Header().Set("Content-Type", "application/json")
		if single {
			_ = json.NewEncoder(w).Encode(resps[0])
			return
		}
		_ = json.NewEncoder(w).Encode(resps)
	}))
}

// newFAQStub returns a fake FAQ bot upstream with fixed responses, so the
// /api/cosmicgame/faq/* proxy routes can be snapshotted deterministically.
func newFAQStub() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"ok","service":"faq-bot-stub"}`))
	})
	mux.HandleFunc("/api/query", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"answer":"stub answer","sources":[]}`))
	})
	mux.HandleFunc("/api/reindex", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"reindexed"}`))
	})
	return httptest.NewServer(mux)
}
