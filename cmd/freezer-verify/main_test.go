package main

import (
	"context"
	"strings"
	"testing"
)

func TestRunFlagErrors(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("bad flag", func(t *testing.T) {
		t.Parallel()
		var out, errOut strings.Builder
		if _, err := run(ctx, []string{"-nonsense"}, &out, &errOut); err == nil {
			t.Fatal("bad flag must error")
		}
	})

	t.Run("missing input", func(t *testing.T) {
		t.Parallel()
		var out, errOut strings.Builder
		_, err := run(ctx, nil, &out, &errOut)
		if err == nil || !strings.Contains(err.Error(), "--input is required") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("invalid table name", func(t *testing.T) {
		t.Parallel()
		var out, errOut strings.Builder
		_, err := run(ctx, []string{"--input", "x.jsonl", "--table", "bad;name"}, &out, &errOut)
		if err == nil || !strings.Contains(err.Error(), "invalid table name") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("missing input file", func(t *testing.T) {
		t.Parallel()
		var out, errOut strings.Builder
		_, err := run(ctx, []string{"--input", t.TempDir() + "/missing.jsonl"}, &out, &errOut)
		if err == nil || !strings.Contains(err.Error(), "opening input") {
			t.Fatalf("err = %v", err)
		}
	})
}
