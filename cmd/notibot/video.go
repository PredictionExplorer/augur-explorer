package main

// ffmpeg adapter for the rwbot.ResampleFunc seam. Twitter rejects the
// generated videos at native resolution, so they are converted to 640x480
// before upload. ffmpeg works on files, so the bytes round-trip through a
// private temp directory.

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// resampleVideo converts a fetched token video to 640x480 with ffmpeg.
func resampleVideo(_ context.Context, video []byte) ([]byte, error) {
	dir, err := os.MkdirTemp("", "notibot-video-*")
	if err != nil {
		return nil, fmt.Errorf("creating video temp dir: %w", err)
	}
	defer os.RemoveAll(dir) //nolint:errcheck // best-effort temp cleanup

	in := filepath.Join(dir, "in.mp4")
	out := filepath.Join(dir, "out.mp4")
	if err := os.WriteFile(in, video, 0o600); err != nil { //nolint:gosec // G703: path built above in a private temp dir
		return nil, fmt.Errorf("writing video temp file: %w", err)
	}
	if err := ffmpeg.Input(in).
		Output(out, ffmpeg.KwArgs{"s": "640x480"}).
		OverWriteOutput().ErrorToStdOut().Run(); err != nil {
		return nil, fmt.Errorf("converting video with ffmpeg: %w", err)
	}
	data, err := os.ReadFile(out) //nolint:gosec // path built above in a private temp dir
	if err != nil {
		return nil, fmt.Errorf("reading converted video: %w", err)
	}
	return data, nil
}
