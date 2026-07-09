package indexer

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"
)

func TestDualLogHandlerRouting(t *testing.T) {
	var info, errBuf bytes.Buffer
	logger := slog.New(NewDualLogHandler(&info, &errBuf))

	logger.Info("batch complete", "blocks", 42)
	logger.Warn("slow batch")
	logger.Error("batch failed", "stage", "fetch")

	infoOut := info.String()
	for _, want := range []string{"batch complete", "blocks=42", "slow batch", "batch failed", "stage=fetch"} {
		if !strings.Contains(infoOut, want) {
			t.Errorf("info log missing %q:\n%s", want, infoOut)
		}
	}

	errOut := errBuf.String()
	if !strings.Contains(errOut, "batch failed") || !strings.Contains(errOut, "stage=fetch") {
		t.Errorf("error log missing the error record:\n%s", errOut)
	}
	for _, unwanted := range []string{"batch complete", "slow batch"} {
		if strings.Contains(errOut, unwanted) {
			t.Errorf("error log contains non-error record %q:\n%s", unwanted, errOut)
		}
	}
}

func TestDualLogHandlerWithAttrsAndGroup(t *testing.T) {
	var info, errBuf bytes.Buffer
	logger := slog.New(NewDualLogHandler(&info, &errBuf)).
		With("etl", "cosmicgame").
		WithGroup("batch")

	logger.Error("failed", "stage", "process")

	for name, out := range map[string]string{"info": info.String(), "error": errBuf.String()} {
		if !strings.Contains(out, "etl=cosmicgame") {
			t.Errorf("%s log missing the With attribute: %s", name, out)
		}
		if !strings.Contains(out, "batch.stage=process") {
			t.Errorf("%s log missing the grouped attribute: %s", name, out)
		}
	}
}
