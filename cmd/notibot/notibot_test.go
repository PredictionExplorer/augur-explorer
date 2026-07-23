package main

import (
	"context"
	"errors"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"testing/synctest"
	"time"
)

func writeKeysFile(t *testing.T, name, contents string) string {
	t.Helper()
	home := t.TempDir()
	if err := os.MkdirAll(filepath.Join(home, "configs"), 0o750); err != nil {
		t.Fatalf("mkdir configs: %v", err)
	}
	if err := os.WriteFile(filepath.Join(home, "configs", name), []byte(contents), 0o600); err != nil {
		t.Fatalf("write keys file: %v", err)
	}
	return home
}

func TestReadTwitterKeys(t *testing.T) {
	home := writeKeysFile(t, "tw.json", `{"ApiKey":"a","ApiSecret":"b","TokenKey":"c","TokenSecret":"d"}`)

	keys, err := readTwitterKeys(home, "tw.json")
	if err != nil {
		t.Fatalf("readTwitterKeys: %v", err)
	}
	if keys.APIKey != "a" || keys.APISecret != "b" || keys.TokenKey != "c" || keys.TokenSecret != "d" {
		t.Errorf("keys = %+v", keys)
	}
}

func TestReadTwitterKeysErrors(t *testing.T) {
	if _, err := readTwitterKeys(t.TempDir(), "missing.json"); err == nil || !strings.Contains(err.Error(), "can't read") {
		t.Errorf("missing file error = %v", err)
	}

	home := writeKeysFile(t, "bad.json", "{not json")
	if _, err := readTwitterKeys(home, "bad.json"); err == nil || !strings.Contains(err.Error(), "can't parse") {
		t.Errorf("bad json error = %v", err)
	}
}

func TestReadDiscordKeys(t *testing.T) {
	home := writeKeysFile(t, "dc.json", `{
		"TokenKey":"tok","ChannelId":1,"MainChannelId":2,
		"MintStatsChanId":3,"PriceStatsChanId":4,"DateStatsChanId":5,"RewardStatsChanId":6}`)

	keys, err := readDiscordKeys(home, "dc.json")
	if err != nil {
		t.Fatalf("readDiscordKeys: %v", err)
	}
	if keys.TokenKey != "tok" || keys.ChannelID != 1 || keys.MintStatsChanID != 3 ||
		keys.PriceStatsChanID != 4 || keys.DateStatsChanID != 5 || keys.RewardStatsChanID != 6 {
		t.Errorf("keys = %+v", keys)
	}
}

func TestReadDiscordKeysErrors(t *testing.T) {
	if _, err := readDiscordKeys(t.TempDir(), "missing.json"); err == nil || !strings.Contains(err.Error(), "can't read") {
		t.Errorf("missing file error = %v", err)
	}

	home := writeKeysFile(t, "bad.json", "[")
	if _, err := readDiscordKeys(home, "bad.json"); err == nil || !strings.Contains(err.Error(), "can't parse") {
		t.Errorf("bad json error = %v", err)
	}
}

func TestRenameWithRetry(t *testing.T) {
	t.Run("first attempt succeeds", func(t *testing.T) {
		calls := 0
		err := renameWithRetry(context.Background(), func() error { calls++; return nil },
			func(context.Context, time.Duration) error {
				t.Error("slept without rate limit")
				return nil
			})
		if err != nil || calls != 1 {
			t.Errorf("err = %v, calls = %d", err, calls)
		}
	})
	t.Run("rate limited then succeeds", func(t *testing.T) {
		calls := 0
		var slept time.Duration
		err := renameWithRetry(context.Background(), func() error {
			calls++
			if calls == 1 {
				return errors.New(`HTTP 429 {"message": "rate limited", "retry_after": 2.35}`)
			}
			return nil
		}, func(_ context.Context, d time.Duration) error {
			slept = d
			return nil
		})
		if err != nil || calls != 2 {
			t.Errorf("err = %v, calls = %d, want retry", err, calls)
		}
		if slept != 3*time.Second { // 2s reported + 1s safety
			t.Errorf("slept = %v, want 3s", slept)
		}
	})
	t.Run("non-rate-limit error is returned without retry", func(t *testing.T) {
		calls := 0
		err := renameWithRetry(context.Background(),
			func() error { calls++; return errors.New("permission denied") },
			func(context.Context, time.Duration) error { return nil })
		if err == nil || calls != 1 {
			t.Errorf("err = %v, calls = %d, want single failing attempt", err, calls)
		}
	})
	t.Run("rate limited twice returns second error", func(t *testing.T) {
		calls := 0
		err := renameWithRetry(context.Background(), func() error {
			calls++
			return errors.New(`retry_after": 1.0`)
		}, func(context.Context, time.Duration) error { return nil })
		if err == nil || calls != 2 {
			t.Errorf("err = %v, calls = %d, want exactly one retry", err, calls)
		}
	})
	t.Run("cancellation during rate-limit wait stops retry", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		calls := 0
		err := renameWithRetry(ctx, func() error {
			calls++
			return errors.New(`retry_after": 60.0`)
		}, sleepContext)
		if !errors.Is(err, context.Canceled) || calls != 1 {
			t.Errorf("err = %v, calls = %d, want one attempt and context cancellation", err, calls)
		}
	})
}

func TestSleepContextCompletes(t *testing.T) {
	t.Parallel()
	synctest.Test(t, func(t *testing.T) {
		if err := sleepContext(t.Context(), time.Hour); err != nil {
			t.Fatalf("sleepContext: %v", err)
		}
	})
}

func TestRunFailsWithoutRPCURL(t *testing.T) {
	err := run(context.Background(), func(string) string { return "" }, io.Discard, true, false)
	if err == nil || !strings.Contains(err.Error(), "RPC_URL: required but not set") {
		t.Errorf("run = %v, want the aggregated configuration failure", err)
	}
}

func TestRunFailsOnUnreachableRPC(t *testing.T) {
	env := map[string]string{"RPC_URL": "://not-a-url"}
	err := run(context.Background(), func(k string) string { return env[k] }, io.Discard, true, false)
	if err == nil || !strings.Contains(err.Error(), "connecting to ETH node") {
		t.Errorf("run = %v, want RPC connection failure", err)
	}
}

func TestResampleVideoTempDirFailure(t *testing.T) {
	t.Setenv("TMPDIR", filepath.Join(t.TempDir(), "does", "not", "exist"))
	if _, err := resampleVideo(context.Background(), []byte("x")); err == nil ||
		!strings.Contains(err.Error(), "creating video temp dir") {
		t.Errorf("resampleVideo = %v, want temp-dir failure", err)
	}
}

func TestResampleVideoRejectsGarbage(t *testing.T) {
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		t.Skip("ffmpeg not installed")
	}
	if _, err := resampleVideo(context.Background(), []byte("not a video")); err == nil {
		t.Error("resampling garbage succeeded")
	}
}

func TestResampleVideoConvertsRealVideo(t *testing.T) {
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		t.Skip("ffmpeg not installed")
	}
	// Generate a tiny test-pattern clip, then run it through the production
	// resampler.
	src := filepath.Join(t.TempDir(), "src.mp4")
	gen := exec.Command("ffmpeg", //nolint:gosec // G204: fixed args, path under t.TempDir
		"-f", "lavfi", "-i", "testsrc=duration=0.2:size=128x96:rate=5",
		"-pix_fmt", "yuv420p", "-y", src)
	if out, err := gen.CombinedOutput(); err != nil {
		t.Skipf("ffmpeg could not generate a test clip: %v\n%s", err, out)
	}
	video, err := os.ReadFile(src) //nolint:gosec // test path under t.TempDir
	if err != nil {
		t.Fatalf("reading generated clip: %v", err)
	}

	converted, err := resampleVideo(context.Background(), video)
	if err != nil {
		t.Fatalf("resampleVideo: %v", err)
	}
	if len(converted) == 0 {
		t.Error("converted video is empty")
	}
}
