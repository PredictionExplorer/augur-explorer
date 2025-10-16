// main.go
package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

// ----------------------------
// Flags / config
// ----------------------------

var (
	beaconCSV      = flag.String("beacon", "", "Comma-separated Beacon endpoints (e.g. http://b1,https://b2)")
	outDir         = flag.String("out", "./archive", "Output directory for archived blobs")
	timeout        = flag.Duration("timeout", 15*time.Second, "HTTP timeout per request")
	pollEvery      = flag.Duration("poll", 4*time.Second, "Polling interval to re-check head")
	maxRetries     = flag.Int("retries", 3, "Retries per HTTP request before trying next endpoint")
	sweepEvery     = flag.Duration("sweep-every", 30*time.Second, "How often the gap sweeper runs")
	sweepHorizon   = flag.Int("sweep-horizon", 2048, "How many slots behind head to check & repair")
	maxPerTick     = flag.Int("max-per-tick", 64, "Max slots processed per follow tick")
	maxParallelDL  = flag.Int("max-parallel", 2, "Max parallel blob downloads per slot")
	reportEvery    = flag.Duration("report-every", 1*time.Minute, "Progress report interval")
)
// U64 unmarshals a uint64 that may arrive as a JSON number or as a quoted string.
type U64 uint64

func (u *U64) UnmarshalJSON(b []byte) error {
	// Try number first
	var num uint64
	if err := json.Unmarshal(b, &num); err == nil {
		*u = U64(num)
		return nil
	}
	// Then try string
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	// Empty string -> 0
	if s == "" {
		*u = 0
		return nil
	}
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return err
	}
	*u = U64(v)
	return nil
}

// ----------------------------
// HTTP client & helpers
// ----------------------------

type endpoint struct{ Base string }

type httpClient struct {
	C  *http.Client
	EP []endpoint
}


func newHTTPClient(beacons []string, to time.Duration) *httpClient {
    trimmed := make([]endpoint, 0, len(beacons))
    for _, b := range beacons {
        b = strings.TrimSpace(b)
        if b == "" { continue }
        trimmed = append(trimmed, endpoint{Base: strings.TrimRight(b, "/")})
    }
    if len(trimmed) == 0 { log.Fatal("no beacon endpoints provided; set -beacon") }

    tr := &http.Transport{
        Proxy:                 http.ProxyFromEnvironment,
        MaxIdleConns:          200,
        MaxIdleConnsPerHost:   200,
        IdleConnTimeout:       90 * time.Second,
        ResponseHeaderTimeout: 30 * time.Second,
        ExpectContinueTimeout: 1 * time.Second,
        DialContext: (&net.Dialer{
            Timeout:   10 * time.Second,
            KeepAlive: 30 * time.Second,
        }).DialContext,
        ForceAttemptHTTP2:     false, // your endpoint is http://; keeps it simple
    }

    return &httpClient{
        C: &http.Client{
            Timeout:   to,       // overall per-request cap; adjust via --timeout
            Transport: tr,
        },
        EP: trimmed,
    }
}


func (hc *httpClient) getJSONAny(ctx context.Context, path string, v any) (string, int, error) {
	var lastErr error
	var lastCode int
	for _, ep := range hc.EP {
		url := ep.Base + path
		for attempt := 0; attempt < *maxRetries; attempt++ {
			req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			resp, err := hc.C.Do(req)
			if err != nil {
				lastErr = err
				log.Printf("GET %s attempt %d/%d error: %v", url, attempt+1, *maxRetries, err)
				time.Sleep(backoff(attempt))
				continue
			}
			func() {
				defer resp.Body.Close()
				lastCode = resp.StatusCode
				if resp.StatusCode != http.StatusOK {
					slurp, _ := io.ReadAll(io.LimitReader(resp.Body, 4<<10))
					log.Printf("GET %s -> HTTP %d body: %q", url, resp.StatusCode, string(slurp))
					if retriable(resp.StatusCode) {
						time.Sleep(backoff(attempt))
						return
					}
					lastErr = fmt.Errorf("http %d", resp.StatusCode)
					// stop retrying this URL
					attempt = *maxRetries
					return
				}
				if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
					lastErr = fmt.Errorf("decode %s: %w", url, err)
					log.Printf("GET %s decode error: %v", url, err)
					time.Sleep(backoff(attempt))
					return
				}
				lastErr = nil
			}()
			if lastErr == nil {
				return url, lastCode, nil
			}
		}
	}
	if lastErr == nil {
		lastErr = fmt.Errorf("all endpoints failed with HTTP %d", lastCode)
	}
	return "", lastCode, lastErr
}

func backoff(attempt int) time.Duration {
	ms := 200 * (1 << attempt)
	if ms > 2000 {
		ms = 2000
	}
	return time.Duration(ms) * time.Millisecond
}

func retriable(code int) bool { return code == 429 || code >= 500 }

// ----------------------------
// BEACON API: types & parsing
// ----------------------------

type headHeadersResp struct {
	Data struct {
		Header struct {
			Message struct {
				Slot string `json:"slot"`
			} `json:"message"`
		} `json:"header"`
	} `json:"data"`
}

type blockV struct {
	Data struct {
		Message struct {
			Body struct {
				ExecutionPayload struct {
					BlobKZGCommitments []string `json:"blob_kzg_commitments"`
				} `json:"execution_payload"`
				BlobKZGCommitments []string `json:"blob_kzg_commitments"`
			} `json:"body"`
		} `json:"message"`
	} `json:"data"`
}
type blobSidecarsResp struct {
	Data []struct {
		Index         U64   `json:"index"`
		KZGCommitment string `json:"kzg_commitment"`
		KZGProof      string `json:"kzg_proof"`
		Blob          string `json:"blob"`
	} `json:"data"`
}

// ----------------------------
// Archiver core
// ----------------------------

const blobBinarySize = 131072 // 128 KiB

var ErrNoBlock = errors.New("no canonical block at slot")

type archiver struct {
	hc        *httpClient
	out       string
	statePath string

	mu       sync.Mutex
	lastSeen uint64

	wgWrites       sync.WaitGroup
	inflightWrites int32 // atomic
	shuttingDown   atomic.Bool

	lastReport time.Time
}

func newArchiver(hc *httpClient, out string) *archiver {
	ap := &archiver{
		hc:        hc,
		out:       out,
		statePath: filepath.Join(out, "state.json"),
		lastReport: time.Now(),
	}
	if err := os.MkdirAll(out, 0o755); err != nil {
		log.Fatalf("mkdir %s: %v", out, err)
	}
	ap.loadState()
	return ap
}

func (a *archiver) loadState() {
	f, err := os.Open(a.statePath)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Printf("state read error: %v", err)
		}
		return
	}
	defer f.Close()
	var s struct {
		LastSeen uint64 `json:"last_seen_slot"`
	}
	if err := json.NewDecoder(f).Decode(&s); err != nil {
		log.Printf("state decode error: %v", err)
		return
	}
	a.lastSeen = s.LastSeen
	log.Printf("follow: starting from lastSeen=%d", a.lastSeen)
}

func (a *archiver) saveState() {
	tmp := a.statePath + ".tmp"
	if err := func() error {
		f, err := os.Create(tmp)
		if err != nil {
			return err
		}
		defer f.Close()
		return json.NewEncoder(f).Encode(struct {
			LastSeen uint64 `json:"last_seen_slot"`
		}{LastSeen: a.lastSeen})
	}(); err != nil {
		log.Printf("state write error: %v", err)
		return
	}
	_ = os.Rename(tmp, a.statePath)
}

func (a *archiver) getHeadSlot(ctx context.Context) (uint64, error) {
	var out headHeadersResp
	if _, _, err := a.hc.getJSONAny(ctx, "/eth/v1/beacon/headers/head", &out); err != nil {
		return 0, err
	}
	s, err := strconv.ParseUint(out.Data.Header.Message.Slot, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parse head slot: %w", err)
	}
	return s, nil
}

func (a *archiver) getCommitmentCount(ctx context.Context, slot uint64) (int, error) {
	var b blockV
	// v2 first
	if _, code, err := a.hc.getJSONAny(ctx, fmt.Sprintf("/eth/v2/beacon/blocks/%d", slot), &b); err == nil {
		if n := len(b.Data.Message.Body.ExecutionPayload.BlobKZGCommitments); n > 0 {
			return n, nil
		}
		if n := len(b.Data.Message.Body.BlobKZGCommitments); n > 0 {
			return n, nil
		}
		return 0, nil
	} else if code != http.StatusNotFound {
		// try v1 below regardless
	}

	// v1
	if _, code, err := a.hc.getJSONAny(ctx, fmt.Sprintf("/eth/v1/beacon/blocks/%d", slot), &b); err != nil {
		if code == http.StatusNotFound {
			return 0, ErrNoBlock
		}
		return 0, err
	}
	if n := len(b.Data.Message.Body.ExecutionPayload.BlobKZGCommitments); n > 0 {
		return n, nil
	}
	if n := len(b.Data.Message.Body.BlobKZGCommitments); n > 0 {
		return n, nil
	}
	return 0, nil
}

func (a *archiver) fetchSidecar(ctx context.Context, slot uint64, index int) (*blobSidecarsResp, error) {
	var resp blobSidecarsResp
	_, _, err := a.hc.getJSONAny(ctx, fmt.Sprintf("/eth/v1/beacon/blob_sidecars/%d?indices=%d", slot, index), &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (a *archiver) saveBlobBinary(outPath, blobHex string) error {
	if strings.HasPrefix(blobHex, "0x") || strings.HasPrefix(blobHex, "0X") {
		blobHex = blobHex[2:]
	}
	if len(blobHex)%2 != 0 {
		return errors.New("blob hex has odd length")
	}
	raw := make([]byte, len(blobHex)/2)
	if _, err := hex.Decode(raw, []byte(blobHex)); err != nil {
		return fmt.Errorf("hex decode: %w", err)
	}
	if len(raw) != blobBinarySize {
		return fmt.Errorf("unexpected blob size: got %d, want %d", len(raw), blobBinarySize)
	}

	dir := filepath.Dir(outPath)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	tmp := outPath + ".tmp"
	f, err := os.OpenFile(tmp, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer func() { _ = os.Remove(tmp) }()

	if _, err := f.Write(raw); err != nil {
		_ = f.Close()
		return err
	}
	if err := f.Sync(); err != nil {
		_ = f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return os.Rename(tmp, outPath)
}

func (a *archiver) slotDir(slot uint64) string { return filepath.Join(a.out, fmt.Sprintf("%d", slot)) }

func (a *archiver) countSavedBlobs(slot uint64) (int, []int) {
	dir := a.slotDir(slot)
	entries, err := os.ReadDir(dir)
	if err != nil {
		return 0, nil
	}
	seen := make(map[int]bool)
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if !strings.HasSuffix(name, ".blob") {
			continue
		}
		// name: slot_<slot>_index_<i>.blob
		parts := strings.Split(name, "_")
		if len(parts) < 5 {
			continue
		}
		iStr := strings.TrimSuffix(parts[len(parts)-1], ".blob")
		i, err := strconv.Atoi(iStr)
		if err == nil {
			seen[i] = true
		}
	}
	var indices []int
	for i := range seen {
		indices = append(indices, i)
	}
	sort.Ints(indices)
	return len(indices), indices
}

func (a *archiver) downloadIndex(ctx context.Context, slot uint64, idx int) error {
	resp, err := a.fetchSidecar(ctx, slot, idx)
	if err != nil {
		return err
	}
	if len(resp.Data) == 0 || resp.Data[0].Blob == "" {
		return fmt.Errorf("slot %d idx %d: empty sidecar", slot, idx)
	}
	out := filepath.Join(a.slotDir(slot), fmt.Sprintf("slot_%d_index_%d.blob", slot, idx))

	// protect write section
	a.wgWrites.Add(1)
	atomic.AddInt32(&a.inflightWrites, 1)

	saveErr := a.saveBlobBinary(out, resp.Data[0].Blob)

	atomic.AddInt32(&a.inflightWrites, -1)
	a.wgWrites.Done()

	if saveErr != nil {
		return saveErr
	}

	// Best-effort metadata
	meta := map[string]any{
		"slot":           slot,
		"index":          idx,
		"kzg_commitment": resp.Data[0].KZGCommitment,
		"kzg_proof":      resp.Data[0].KZGProof,
		"saved_at_unix":  time.Now().Unix(),
	}
	_ = os.WriteFile(strings.TrimSuffix(out, ".blob")+".json", mustJSON(meta), 0o644)

	log.Printf("saved %s (%d bytes)", out, blobBinarySize)
	return nil
}

func (a *archiver) processSlot(ctx context.Context, slot uint64) error {
	count, err := a.getCommitmentCount(ctx, slot)
	if err != nil {
		if errors.Is(err, ErrNoBlock) {
			log.Printf("slot %d: no canonical block (empty)", slot)
			return nil
		}
		return err
	}
	if count == 0 {
		// legit block but no blobs
		return nil
	}

	// ensure dir
	if err := os.MkdirAll(a.slotDir(slot), 0o755); err != nil {
		return err
	}

	// download missing indices with small concurrency
	saved, have := a.countSavedBlobs(slot)
	if saved == count {
		return nil
	}
	missing := make(map[int]bool)
	for i := 0; i < count; i++ {
		missing[i] = true
	}
	for _, i := range have {
		delete(missing, i)
	}

	sem := make(chan struct{}, *maxParallelDL)
	var wg sync.WaitGroup
	var firstErr atomic.Value

	for i := range missing {
		if a.shuttingDown.Load() {
			break
		}
		wg.Add(1)
		sem <- struct{}{}
		go func(ix int) {
			defer wg.Done()
			defer func() { <-sem }()
			dlCtx, cancel := context.WithTimeout(ctx, *timeout*2)
			defer cancel()
			if err := a.downloadIndex(dlCtx, slot, ix); err != nil {
				firstErr.Store(err)
			}
		}(i)
	}
	wg.Wait()

	if v := firstErr.Load(); v != nil {
		return v.(error)
	}
	return nil
}

func mustJSON(v any) []byte { b, _ := json.MarshalIndent(v, "", "  "); return b }

// ----------------------------
// Follower: moves lastSeen forward
// ----------------------------

func (a *archiver) runFollow(ctx context.Context) {
	t := time.NewTicker(*pollEvery)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			head, err := a.getHeadSlot(ctx)
			if err != nil {
				log.Printf("head error: %v", err)
				continue
			}
			start := a.lastSeen + 1
			if start == 0 {
				start = head
			}
			if start > head {
				a.maybeReport(head)
				continue
			}

			end := start + uint64(*maxPerTick) - 1
			if end > head {
				end = head
			}
			for s := start; s <= end; s++ {
				if a.shuttingDown.Load() {
					return
				}
				slotCtx, cancel := context.WithTimeout(ctx, *timeout*2)
				err := a.processSlot(slotCtx, s)
				cancel()
				if err != nil {
					// transient; try next tick
					log.Printf("follow slot %d: %v", s, err)
					break
				}
				a.mu.Lock()
				a.lastSeen = s
				a.mu.Unlock()
				a.saveState()
			}
			a.maybeReport(head)
		}
	}
}

func (a *archiver) maybeReport(head uint64) {
	if time.Since(a.lastReport) < *reportEvery {
		return
	}
	a.lastReport = time.Now()
	lag := int64(head) - int64(a.lastSeen)
	if lag < 0 {
		lag = 0
	}
	log.Printf("progress: lastSeen=%d head=%d lag=%d slots inflightWrites=%d",
		a.lastSeen, head, lag, atomic.LoadInt32(&a.inflightWrites))
}

// ----------------------------
// Sweeper: auto gap repair
// ----------------------------

func (a *archiver) runSweeper(ctx context.Context) {
	t := time.NewTicker(*sweepEvery)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			head, err := a.getHeadSlot(ctx)
			if err != nil {
				log.Printf("sweeper head error: %v", err)
				continue
			}
			var start uint64
			if a.lastSeen > uint64(*sweepHorizon) {
				start = a.lastSeen - uint64(*sweepHorizon)
			} else {
				start = 0
			}
			end := a.lastSeen
			if end > head {
				end = head
			}
			if start == 0 && end == 0 {
				continue
			}

			for s := start; s <= end; s++ {
				if a.shuttingDown.Load() {
					return
				}
				slotCtx, cancel := context.WithTimeout(ctx, *timeout*2)
				count, err := a.getCommitmentCount(slotCtx, s)
				cancel()
				if err != nil {
					if errors.Is(err, ErrNoBlock) {
						continue
					}
					// transient, skip this round
					continue
				}
				if count == 0 {
					continue
				}
				existing, _ := a.countSavedBlobs(s)
				if existing >= count {
					continue
				}

				// Repair this slot
				slotCtx2, cancel2 := context.WithTimeout(ctx, *timeout*4)
				err = a.processSlot(slotCtx2, s)
				cancel2()
				if err != nil {
					// don’t spam logs; one line is enough
					log.Printf("sweeper: slot %d repair failed: %v", s, err)
				}
			}
		}
	}
}

// ----------------------------
// main
// ----------------------------

func main() {
	flag.Parse()

	beacons := strings.Split(*beaconCSV, ",")
	hc := newHTTPClient(beacons, *timeout)
	arc := newArchiver(hc, *outDir)

	// Signal handling with graceful write completion.
	ctx, stop := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer stop()

	// Sanity probe (non-fatal)
	var gen map[string]any
	if _, code, err := hc.getJSONAny(ctx, "/eth/v1/beacon/genesis", &gen); err != nil {
		log.Printf("warning: /eth/v1/beacon/genesis probe failed (HTTP %d): %v", code, err)
	} else {
		log.Printf("genesis probe OK")
	}

	// Run follower + sweeper
	var wg sync.WaitGroup
	wg.Add(2)
	go func() { defer wg.Done(); arc.runFollow(ctx) }()
	go func() { defer wg.Done(); arc.runSweeper(ctx) }()

	// Wait for signal
	<-ctx.Done()
	log.Printf("signal received, initiating graceful shutdown...")
	arc.shuttingDown.Store(true)

	// Wait both loops to stop
	wg.Wait()

	// Wait for any in-flight file writes to complete.
	if n := atomic.LoadInt32(&arc.inflightWrites); n > 0 {
		log.Printf("waiting for %d inflight write(s) to finish...", n)
	}
	arc.wgWrites.Wait()

	// Persist final state
	arc.saveState()
	log.Printf("shutdown complete.")
}

