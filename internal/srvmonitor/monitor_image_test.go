package srvmonitor

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
)

// rwalkAddr is the fake RandomWalk contract address used by image tests.
var rwalkAddr = common.HexToAddress("0x1111111111111111111111111111111111111111")

// imageTestSetup wires an image monitor to a fake mint table, an httptest
// image server and a fake chain whose RWalk contract reports nextTokenId.
type imageTestSetup struct {
	monitor *ImageMonitor
	server  *httptest.Server
	present map[int64]bool // token id -> image exists
	mu      sync.Mutex
}

func newImageTestSetup(t *testing.T, tokenIDs []int64, nextTokenID int64) *imageTestSetup {
	t.Helper()
	s := &imageTestSetup{present: make(map[int64]bool)}

	s.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var tokenID int64
		if _, err := fmt.Sscanf(r.URL.Path, "/images/%d_black_thumb.jpg", &tokenID); err != nil {
			http.NotFound(w, r)
			return
		}
		s.mu.Lock()
		ok := s.present[tokenID]
		s.mu.Unlock()
		if !ok {
			http.NotFound(w, r)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(s.server.Close)

	chain := testchain.New(t)
	chain.EnsureBlock(1)
	stub := testchain.MustContractStub(randomwalk.RWalkABI).
		Return("nextTokenId", big.NewInt(nextTokenID))
	chain.RegisterCall(rwalkAddr, stub.Handler())

	rows := make([][]int64, 0, len(tokenIDs))
	for _, id := range tokenIDs {
		rows = append(rows, []int64{id})
	}
	conn := &fakeConn{queryRows: rows}

	m := NewImageMonitor(ImageServerConfig{
		Name:         "RWalk Thumbnails",
		URL:          s.server.URL + "/images",
		ContractAddr: rwalkAddr.Hex(),
		RPCURL:       chain.URL(),
	}, dbConfig("rwalk"), 35, testIntervals())
	m.connect = connector(conn, nil)
	m.randInt64N = func(n int64) int64 { return 2 } // deterministic random pick

	s.monitor = m
	return s
}

func (s *imageTestSetup) setPresent(ids ...int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, id := range ids {
		s.present[id] = true
	}
}

func TestImageMonitorAllPresentAndMatching(t *testing.T) {
	t.Parallel()
	// Latest mints 42, 41, 40; contract nextTokenId 43 => last minted 42.
	s := newImageTestSetup(t, []int64{42, 41, 40}, 43)
	s.setPresent(42, 41, 40, 2)

	disp := newFakeDisplay()
	errCh := make(chan string, 10)
	s.monitor.check(context.Background(), disp, errCh)

	if msgs := drain(errCh); len(msgs) != 0 {
		t.Fatalf("unexpected errors: %v", msgs)
	}
	d := s.monitor.data
	if d.DBTokenID != 42 || d.ContractTokenID != 42 || !d.TokensMatch {
		t.Fatalf("data = %+v, want matching token 42", d)
	}
	for i, want := range []int64{42, 41, 40} {
		if d.LatestTokens[i].TokenID != want || !d.LatestTokens[i].IsPresent {
			t.Fatalf("latest[%d] = %+v", i, d.LatestTokens[i])
		}
	}
	if d.RandomToken.TokenID != 2 || !d.RandomToken.IsPresent {
		t.Fatalf("random = %+v", d.RandomToken)
	}

	row1 := disp.Row(36)
	if !strings.Contains(row1, "000042:Ok") || !strings.Contains(row1, "Rnd 000002:Ok") {
		t.Fatalf("row1 = %q", row1)
	}
	row2 := disp.Row(37)
	if !strings.Contains(row2, "Last token DB: 000042") ||
		!strings.Contains(row2, "Last token ctrct: 000042") ||
		!strings.Contains(row2, "Match: Ok") {
		t.Fatalf("row2 = %q", row2)
	}
}

func TestImageMonitorMissingImageAndMismatch(t *testing.T) {
	t.Parallel()
	// Contract says last minted is 45, DB says 42: mismatch. Token 41's
	// image is missing.
	s := newImageTestSetup(t, []int64{42, 41, 40}, 46)
	s.setPresent(42, 40, 2)

	disp := newFakeDisplay()
	errCh := make(chan string, 10)
	s.monitor.check(context.Background(), disp, errCh)

	msgs := drain(errCh)
	// One mismatch error + one missing-image error.
	if len(msgs) != 2 {
		t.Fatalf("errors = %v, want 2", msgs)
	}
	var sawMismatch, sawMissing bool
	for _, msg := range msgs {
		if strings.Contains(msg, "Token ID mismatch: DB=42, Contract=45") {
			sawMismatch = true
		}
		if strings.Contains(msg, "RWalk image for token 41 is not present") {
			sawMissing = true
		}
	}
	if !sawMismatch || !sawMissing {
		t.Fatalf("errors = %v", msgs)
	}

	d := s.monitor.data
	if d.TokensMatch || d.LatestTokens[1].IsPresent {
		t.Fatalf("data = %+v", d)
	}
	if row2 := disp.Row(37); !strings.Contains(row2, "Match: Fail") {
		t.Fatalf("row2 = %q", row2)
	}
	if row1 := disp.Row(36); !strings.Contains(row1, "000041:Fail") {
		t.Fatalf("row1 = %q", row1)
	}
}

func TestImageMonitorTokenQueryFailures(t *testing.T) {
	t.Parallel()
	for name, conn := range map[string]*fakeConn{
		"query error": {queryErr: errors.New("db gone")},
		"rows error":  {queryRows: [][]int64{{5}}, rowsErr: errors.New("stream cut")},
		"scan error":  {queryRows: [][]int64{{5}}, rowsScanErr: errors.New("bad column")},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			m := NewImageMonitor(ImageServerConfig{URL: "http://img.invalid", RPCURL: "http://127.0.0.1:1"},
				dbConfig("rwalk"), 35, testIntervals())
			m.connect = connector(conn, nil)

			errCh := make(chan string, 10)
			m.check(context.Background(), newFakeDisplay(), errCh)

			msgs := drain(errCh)
			if len(msgs) != 1 || !strings.Contains(msgs[0], "DB Error:") {
				t.Fatalf("errors = %v", msgs)
			}
			if m.data.DBTokenID != -1 {
				t.Fatalf("data = %+v", m.data)
			}
		})
	}
}

func TestImageMonitorMalformedImageURL(t *testing.T) {
	t.Parallel()
	chain := testchain.New(t)
	chain.EnsureBlock(1)
	stub := testchain.MustContractStub(randomwalk.RWalkABI).Return("nextTokenId", big.NewInt(8))
	chain.RegisterCall(rwalkAddr, stub.Handler())

	// A control character in the base URL makes request construction fail;
	// the check reports the error instead of panicking.
	conn := &fakeConn{queryRows: [][]int64{{7}}}
	m := NewImageMonitor(ImageServerConfig{
		URL:          "http://img/\x7f",
		ContractAddr: rwalkAddr.Hex(),
		RPCURL:       chain.URL(),
	}, dbConfig("rwalk"), 35, testIntervals())
	m.connect = connector(conn, nil)
	m.randInt64N = func(int64) int64 { return 3 }

	errCh := make(chan string, 10)
	m.check(context.Background(), newFakeDisplay(), errCh)

	msgs := drain(errCh)
	if len(msgs) != 2 { // latest token + random token both fail
		t.Fatalf("errors = %v", msgs)
	}
	if !strings.Contains(m.data.LatestTokens[0].ErrStr, "invalid control character") {
		t.Fatalf("latest[0] = %+v", m.data.LatestTokens[0])
	}
}

func TestImageMonitorDBFailure(t *testing.T) {
	t.Parallel()
	m := NewImageMonitor(ImageServerConfig{URL: "http://img.test", RPCURL: "http://127.0.0.1:1"},
		dbConfig("rwalk"), 35, testIntervals())
	m.connect = connector(nil, errors.New("db down"))

	disp := newFakeDisplay()
	errCh := make(chan string, 10)
	m.check(context.Background(), disp, errCh)

	msgs := drain(errCh)
	if len(msgs) != 1 || !strings.Contains(msgs[0], "DB Error: db down") {
		t.Fatalf("errors = %v", msgs)
	}
	d := m.data
	if d.DBTokenID != -1 || d.ContractTokenID != -1 || d.TokensMatch {
		t.Fatalf("data = %+v", d)
	}
	for i := range d.LatestTokens {
		if d.LatestTokens[i].TokenID != -1 {
			t.Fatalf("latest[%d] = %+v", i, d.LatestTokens[i])
		}
	}
}

func TestImageMonitorContractFailure(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(server.Close)

	conn := &fakeConn{queryRows: [][]int64{{7}}}
	m := NewImageMonitor(ImageServerConfig{
		URL:          server.URL,
		ContractAddr: rwalkAddr.Hex(),
		RPCURL:       "http://127.0.0.1:1", // unreachable node
	}, dbConfig("rwalk"), 35, testIntervals())
	m.connect = connector(conn, nil)
	m.randInt64N = func(int64) int64 { return 0 }

	errCh := make(chan string, 10)
	m.check(context.Background(), newFakeDisplay(), errCh)

	msgs := drain(errCh)
	if len(msgs) != 1 || !strings.Contains(msgs[0], "Contract error:") {
		t.Fatalf("errors = %v", msgs)
	}
	d := m.data
	if d.ContractTokenID != -1 || d.TokensMatch {
		t.Fatalf("data = %+v", d)
	}
	// The image checks still ran despite the contract failure.
	if !d.LatestTokens[0].IsPresent || d.LatestTokens[0].TokenID != 7 {
		t.Fatalf("latest[0] = %+v", d.LatestTokens[0])
	}
}

func TestImageMonitorEmptyMintTable(t *testing.T) {
	t.Parallel()
	chain := testchain.New(t)
	chain.EnsureBlock(1)
	stub := testchain.MustContractStub(randomwalk.RWalkABI).Return("nextTokenId", big.NewInt(0))
	chain.RegisterCall(rwalkAddr, stub.Handler())

	conn := &fakeConn{queryRows: nil}
	m := NewImageMonitor(ImageServerConfig{
		URL:          "http://img.invalid",
		ContractAddr: rwalkAddr.Hex(),
		RPCURL:       chain.URL(),
	}, dbConfig("rwalk"), 35, testIntervals())
	m.connect = connector(conn, nil)

	errCh := make(chan string, 10)
	m.check(context.Background(), newFakeDisplay(), errCh)

	d := m.data
	if d.DBTokenID != -1 {
		t.Fatalf("DBTokenID = %d, want -1 for empty table", d.DBTokenID)
	}
	// nextTokenId 0 => last minted -1, matching the empty DB sentinel.
	if d.ContractTokenID != -1 || !d.TokensMatch {
		t.Fatalf("data = %+v", d)
	}
	if d.RandomToken.TokenID != 0 {
		t.Fatalf("random = %+v, want untouched zero value", d.RandomToken)
	}
}

func TestImageMonitorImageServerUnreachable(t *testing.T) {
	t.Parallel()
	chain := testchain.New(t)
	chain.EnsureBlock(1)
	stub := testchain.MustContractStub(randomwalk.RWalkABI).Return("nextTokenId", big.NewInt(8))
	chain.RegisterCall(rwalkAddr, stub.Handler())

	conn := &fakeConn{queryRows: [][]int64{{7}}}
	m := NewImageMonitor(ImageServerConfig{
		URL:          "http://127.0.0.1:1", // unreachable image server
		ContractAddr: rwalkAddr.Hex(),
		RPCURL:       chain.URL(),
	}, dbConfig("rwalk"), 35, testIntervals())
	m.connect = connector(conn, nil)
	m.randInt64N = func(int64) int64 { return 3 }

	errCh := make(chan string, 10)
	m.check(context.Background(), newFakeDisplay(), errCh)

	msgs := drain(errCh)
	// Latest token 7 + random token 3, both unreachable.
	if len(msgs) != 2 {
		t.Fatalf("errors = %v, want 2 transport errors", msgs)
	}
	d := m.data
	if d.LatestTokens[0].IsPresent || d.LatestTokens[0].ErrStr == "" {
		t.Fatalf("latest[0] = %+v", d.LatestTokens[0])
	}
	if d.RandomToken.IsPresent || d.RandomToken.ErrStr == "" {
		t.Fatalf("random = %+v", d.RandomToken)
	}
}

func TestImageMonitorStartLoopStopsOnCancel(t *testing.T) {
	t.Parallel()
	s := newImageTestSetup(t, []int64{5}, 6)
	s.setPresent(5, 2)

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	disp := newFakeDisplay()
	go func() {
		s.monitor.Start(ctx, disp, make(chan string, 100))
		close(done)
	}()

	waitFor(t, "two check cycles", func() bool { return disp.Flushes() >= 2 })
	cancel()
	waitFor(t, "loop exit", func() bool {
		select {
		case <-done:
			return true
		default:
			return false
		}
	})
}
