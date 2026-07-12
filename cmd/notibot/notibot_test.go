package main

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func writeKeysFile(t *testing.T, name, contents string) {
	t.Helper()
	home := t.TempDir()
	t.Setenv("HOME", home)
	if err := os.MkdirAll(filepath.Join(home, "configs"), 0o750); err != nil {
		t.Fatalf("mkdir configs: %v", err)
	}
	if err := os.WriteFile(filepath.Join(home, "configs", name), []byte(contents), 0o600); err != nil {
		t.Fatalf("write keys file: %v", err)
	}
}

func TestReadTwitterKeys(t *testing.T) {
	writeKeysFile(t, "tw.json", `{"ApiKey":"a","ApiSecret":"b","TokenKey":"c","TokenSecret":"d"}`)
	t.Setenv("TWITTER_KEYS_FILE", "tw.json")

	keys, err := readTwitterKeys()
	if err != nil {
		t.Fatalf("readTwitterKeys: %v", err)
	}
	if keys.ApiKey != "a" || keys.ApiSecret != "b" || keys.TokenKey != "c" || keys.TokenSecret != "d" {
		t.Errorf("keys = %+v", keys)
	}
}

func TestReadTwitterKeysErrors(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	t.Setenv("TWITTER_KEYS_FILE", "missing.json")
	if _, err := readTwitterKeys(); err == nil || !strings.Contains(err.Error(), "can't read") {
		t.Errorf("missing file error = %v", err)
	}

	writeKeysFile(t, "bad.json", "{not json")
	t.Setenv("TWITTER_KEYS_FILE", "bad.json")
	if _, err := readTwitterKeys(); err == nil || !strings.Contains(err.Error(), "can't parse") {
		t.Errorf("bad json error = %v", err)
	}
}

func TestReadDiscordKeys(t *testing.T) {
	writeKeysFile(t, "dc.json", `{
		"TokenKey":"tok","ChannelId":1,"MainChannelId":2,
		"MintStatsChanId":3,"PriceStatsChanId":4,"DateStatsChanId":5,"RewardStatsChanId":6}`)
	t.Setenv("DISCORD_KEYS_FILE", "dc.json")

	keys, err := readDiscordKeys()
	if err != nil {
		t.Fatalf("readDiscordKeys: %v", err)
	}
	if keys.TokenKey != "tok" || keys.ChannelId != 1 || keys.MintStatsChanId != 3 ||
		keys.PriceStatsChanId != 4 || keys.DateStatsChanId != 5 || keys.RewardStatsChanId != 6 {
		t.Errorf("keys = %+v", keys)
	}
}

func TestReadDiscordKeysErrors(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	t.Setenv("DISCORD_KEYS_FILE", "missing.json")
	if _, err := readDiscordKeys(); err == nil || !strings.Contains(err.Error(), "can't read") {
		t.Errorf("missing file error = %v", err)
	}

	writeKeysFile(t, "bad.json", "[")
	t.Setenv("DISCORD_KEYS_FILE", "bad.json")
	if _, err := readDiscordKeys(); err == nil || !strings.Contains(err.Error(), "can't parse") {
		t.Errorf("bad json error = %v", err)
	}
}

func TestRenameWithRetry(t *testing.T) {
	t.Run("first attempt succeeds", func(t *testing.T) {
		calls := 0
		err := renameWithRetry(func() error { calls++; return nil }, func(time.Duration) { t.Error("slept without rate limit") })
		if err != nil || calls != 1 {
			t.Errorf("err = %v, calls = %d", err, calls)
		}
	})
	t.Run("rate limited then succeeds", func(t *testing.T) {
		calls := 0
		var slept time.Duration
		err := renameWithRetry(func() error {
			calls++
			if calls == 1 {
				return errors.New(`HTTP 429 {"message": "rate limited", "retry_after": 2.35}`)
			}
			return nil
		}, func(d time.Duration) { slept = d })
		if err != nil || calls != 2 {
			t.Errorf("err = %v, calls = %d, want retry", err, calls)
		}
		if slept != 3*time.Second { // 2s reported + 1s safety
			t.Errorf("slept = %v, want 3s", slept)
		}
	})
	t.Run("non-rate-limit error is returned without retry", func(t *testing.T) {
		calls := 0
		err := renameWithRetry(func() error { calls++; return errors.New("permission denied") }, func(time.Duration) {})
		if err == nil || calls != 1 {
			t.Errorf("err = %v, calls = %d, want single failing attempt", err, calls)
		}
	})
	t.Run("rate limited twice returns second error", func(t *testing.T) {
		calls := 0
		err := renameWithRetry(func() error {
			calls++
			return errors.New(`retry_after": 1.0`)
		}, func(time.Duration) {})
		if err == nil || calls != 2 {
			t.Errorf("err = %v, calls = %d, want exactly one retry", err, calls)
		}
	})
}

func TestOpenLogFileCreatesAndAppends(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "x.log")
	f, err := openLogFile(path)
	if err != nil {
		t.Fatalf("openLogFile: %v", err)
	}
	if _, err := f.WriteString("one\n"); err != nil {
		t.Fatalf("write: %v", err)
	}
	_ = f.Close()
	f, err = openLogFile(path)
	if err != nil {
		t.Fatalf("openLogFile again: %v", err)
	}
	if _, err := f.WriteString("two\n"); err != nil {
		t.Fatalf("write: %v", err)
	}
	_ = f.Close()
	data, err := os.ReadFile(path) //nolint:gosec // test path under t.TempDir
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	if string(data) != "one\ntwo\n" {
		t.Errorf("log contents = %q, want appended lines", data)
	}
}

func TestOpenLogFileError(t *testing.T) {
	if _, err := openLogFile(filepath.Join(t.TempDir(), "no", "such", "dir", "x.log")); err == nil {
		t.Error("openLogFile into missing directory succeeded")
	}
}

func TestRunFailsWhenLogDirUnwritable(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	// A file occupies the log-directory path, so MkdirAll must fail.
	if err := os.WriteFile(filepath.Join(home, "ae_logs"), []byte("x"), 0o600); err != nil {
		t.Fatal(err)
	}
	err := run(context.Background(), true, false)
	if err == nil || !strings.Contains(err.Error(), "creating log directory") {
		t.Errorf("run = %v, want log-directory failure", err)
	}
}

func TestRunFailsOnUnreachableRPC(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	t.Setenv("RPC_URL", "")
	err := run(context.Background(), true, false)
	if err == nil || !strings.Contains(err.Error(), "connecting to ETH node") {
		t.Errorf("run = %v, want RPC connection failure", err)
	}
}

func TestRunFailsWhenLogFileIsADirectory(t *testing.T) {
	// A directory squatting on a log-file path makes that openLogFile call
	// fail; each case exercises one of the three open sites.
	for _, name := range []string{"notibot_error.log", "notibot_db.log"} {
		t.Run(name, func(t *testing.T) {
			home := t.TempDir()
			t.Setenv("HOME", home)
			logDir := filepath.Join(home, "ae_logs")
			if err := os.MkdirAll(filepath.Join(logDir, name), 0o750); err != nil {
				t.Fatal(err)
			}
			err := run(context.Background(), true, false)
			if err == nil || !strings.Contains(err.Error(), name) {
				t.Errorf("run = %v, want failure opening %s", err, name)
			}
		})
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
