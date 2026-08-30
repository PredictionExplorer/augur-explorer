package cosmicgame

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"
	"time"

	cgdb "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

// awaitClosed asserts the ingest lifecycle channel closes promptly. The
// deadline is an outer guard: an already-closed channel returns immediately
// and a started loop stops as soon as its context is cancelled.
func awaitClosed(tb testing.TB, done <-chan struct{}) {
	tb.Helper()
	select {
	case <-done:
	case <-time.After(10 * time.Second):
		tb.Fatal("the trait ingest lifecycle channel never closed")
	}
}

// moduleWithRepo returns a module carrying the one field
// StartTraitIngestion inspects before deciding whether to run. The repo has
// no pool behind it, which is safe precisely because every case here is
// declined before the loop starts; the started path is covered end to end
// against a real database in internal/api/apitest.
func moduleWithRepo(logger *slog.Logger) *API {
	a := NewBare()
	a.repo = cgdb.NewRepo(nil)
	if logger != nil {
		a.logger = logger
	}
	return a
}

// TestStartTraitIngestionDeclinedCases pins the no-op paths. None of them is
// a failure: with no trait rows every token serves the fallback metadata,
// which is exactly the behaviour before the generator published anything.
func TestStartTraitIngestionDeclinedCases(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		module     func(*slog.Logger) *API
		cfg        TraitIngestConfig
		wantLog    string
		wantNudger bool
	}{
		{
			name:   "no asset host configured",
			module: moduleWithRepo,
			cfg:    TraitIngestConfig{},
			// Operators need to know why tokens stay on the fallback.
			wantLog: "NFT_TRAITS_SOURCE_BASE unset",
		},
		{
			name:   "explicitly disabled",
			module: moduleWithRepo,
			cfg:    TraitIngestConfig{SourceBase: "https://assets.example.com", Disabled: true},
		},
		{
			name:   "whitespace-only asset host",
			module: moduleWithRepo,
			cfg:    TraitIngestConfig{SourceBase: "   "},
		},
		{
			name:   "no database link",
			module: func(*slog.Logger) *API { return NewBare() },
			cfg:    TraitIngestConfig{SourceBase: "https://assets.example.com"},
		},
		{
			name:    "malformed asset host",
			module:  moduleWithRepo,
			cfg:     TraitIngestConfig{SourceBase: "nfts.cosmicsignature.com"},
			wantLog: "trait ingestion disabled",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var logged bytes.Buffer
			a := tc.module(slog.New(slog.NewTextHandler(&logged, nil)))

			done := a.StartTraitIngestion(t.Context(), tc.cfg)
			awaitClosed(t, done)

			if a.traitsIngester != nil {
				t.Error("a declined start still published the nudge hook")
			}
			if tc.wantLog != "" && !strings.Contains(logged.String(), tc.wantLog) {
				t.Errorf("log = %q, want it to mention %q", logged.String(), tc.wantLog)
			}
		})
	}
}
