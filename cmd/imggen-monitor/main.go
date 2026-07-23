// Command imggen-monitor verifies that every minted Cosmic Signature NFT has
// its generated image and video artifacts available on the artifact server,
// and optionally asks the generator service to (re)create missing ones.
//
// Modes:
//
//	imggen-monitor                          # report presence of artifacts for all tokens
//	imggen-monitor -regenerate              # regenerate any missing artifacts
//	imggen-monitor -token 123 -seed 0xabc   # one-shot generation for a single token
//
// Configuration (environment):
//
//	IM_REQUEST_URL  POST endpoint of the artifact generator service
//	IM_IMAGE_URL    base URL where images are served  (<base><id>.png)
//	IM_VIDEO_URL    base URL where videos are served  (<base><id>.mp4)
//	PGSQL_*         PostgreSQL connection (scan mode only)
//
// The check/generate pipeline lives in internal/ops/imggen; this binary
// parses flags and wires the production repository.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/config"
	"github.com/PredictionExplorer/augur-explorer/internal/ops/imggen"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/version"
)

// httpRequestTimeout bounds each artifact-server or generator request.
const httpRequestTimeout = 30 * time.Second

// Server-side database time bounds (D22 defense in depth): the scan reads
// the token directory in one bounded statement, and the systemd timer runs
// the scan unattended, so a wedged statement must not stall it forever.
const (
	dbStatementTimeout = time.Minute
	dbIdleInTxTimeout  = 5 * time.Minute
)

// osExit is stubbed by tests that drive main through its failure arm.
var osExit = os.Exit

func main() {
	if version.HandleFlag(os.Args[1:], os.Stdout) {
		return
	}
	if err := runMain(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "imggen-monitor: %v\n", err)
		osExit(1)
	}
}

// runMain owns the signal-scoped context so its deferred cleanup always runs
// before main decides the exit code (os.Exit skips deferred calls).
func runMain(args []string) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	return run(ctx, args, os.Getenv, os.Stdout, os.Stderr, os.Stdout)
}

// repoTokenSource adapts the CosmicGame repository to imggen.TokenSource.
type repoTokenSource struct {
	repo *cgstore.Repo
}

func (s repoTokenSource) Tokens(ctx context.Context) ([]imggen.Token, error) {
	recs, err := s.repo.CosmicSignatureTokens(ctx, 0, 100000)
	if err != nil {
		return nil, err
	}
	tokens := make([]imggen.Token, 0, len(recs))
	for _, rec := range recs {
		tokens = append(tokens, imggen.Token{ID: rec.TokenId, Seed: rec.Seed})
	}
	return tokens, nil
}

// run parses flags, loads the typed configuration and executes the selected
// mode. Structured diagnostics go to logOut; report output stays on out.
func run(ctx context.Context, args []string, getenv func(string) string, out, errOut, logOut io.Writer) error {
	flags := flag.NewFlagSet("imggen-monitor", flag.ContinueOnError)
	flags.SetOutput(errOut)
	regenerate := flags.Bool("regenerate", false, "regenerate missing artifacts while scanning")
	tokenID := flags.Int64("token", -1, "one-shot: generate artifacts for this token id (requires -seed)")
	seed := flags.String("seed", "", "one-shot: seed of the token passed with -token")
	if err := flags.Parse(args); err != nil {
		return err
	}

	cfg, err := config.LoadImggenMonitor(getenv)
	if err != nil {
		return err
	}
	logger := cfg.Log.NewLogger(logOut)
	logger.LogAttrs(ctx, slog.LevelInfo, "build info", version.LogAttrs()...)

	client := &imggen.Client{
		RequestURL: cfg.RequestURL,
		ImageURL:   cfg.ImageURL,
		VideoURL:   cfg.VideoURL,
		HTTPClient: &http.Client{Timeout: httpRequestTimeout},
	}

	// One-shot mode: request generation for a single token.
	if *tokenID >= 0 {
		if *seed == "" {
			return errors.New("-token requires -seed")
		}
		return client.Generate(ctx, *tokenID, *seed)
	}

	// Scan mode: iterate every minted token from the database.
	storeCfg := cfg.DB.StoreConfig()
	storeCfg.Logger = logger.With("component", "db")
	storeCfg.StatementTimeout = dbStatementTimeout
	storeCfg.IdleInTxSessionTimeout = dbIdleInTxTimeout
	st, err := store.New(ctx, storeCfg)
	if err != nil {
		return fmt.Errorf("failed to connect to storage: %w", err)
	}
	defer st.Close()

	return imggen.Scan(ctx, imggen.ScanOptions{
		Source:     repoTokenSource{repo: cgstore.NewRepo(st)},
		Client:     client,
		Regenerate: *regenerate,
		Out:        out,
	})
}
