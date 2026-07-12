//go:build integration

package rwbot

// End-to-end engine test against a real seeded Postgres: the production
// *randomwalk.Repo satisfies DataSource, the engine announces every fixture
// event exactly once through the fake sinks, and the rw_messaging_status
// watermark persists so a second run announces nothing.

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
	"github.com/PredictionExplorer/augur-explorer/internal/testfixtures"
)

// The production repository must satisfy the engine's store seam.
var _ DataSource = (*rwstore.Repo)(nil)

var (
	sharedRepo   *rwstore.Repo
	errSetupSkip error
)

func TestMain(m *testing.M) {
	os.Exit(runMain(m))
}

func runMain(m *testing.M) int {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	db, stop, err := testdb.Start(ctx)
	if err != nil {
		if errors.Is(err, testdb.ErrContainerUnavailable) {
			errSetupSkip = err
			return m.Run()
		}
		fmt.Fprintf(os.Stderr, "rwbot: starting test database: %v\n", err)
		return 1
	}
	defer stop()

	if err := testfixtures.Apply(ctx, db.SQL); err != nil {
		fmt.Fprintf(os.Stderr, "rwbot: seeding fixture data: %v\n", err)
		return 1
	}
	sharedRepo = rwstore.NewRepo(store.NewFromPool(db.Pool))
	return m.Run()
}

func integrationRepo(t *testing.T) *rwstore.Repo {
	t.Helper()
	if errSetupSkip != nil {
		t.Skipf("skipping: %v", errSetupSkip)
	}
	if sharedRepo == nil {
		t.Fatal("rwbot harness not initialized (TestMain did not run?)")
	}
	return sharedRepo
}

// staticMedia serves the same 200 body for every URL.
type staticMedia struct{ body []byte }

func (s staticMedia) Fetch(context.Context, string) (int, []byte, error) {
	return 200, s.body, nil
}

func TestEngineAgainstSeededDatabase(t *testing.T) {
	repo := integrationRepo(t)
	ctx := context.Background()

	addrs, err := repo.ContractAddrs(ctx)
	if err != nil {
		t.Fatalf("ContractAddrs: %v", err)
	}

	// The fixture events the engine must announce, straight from the
	// production query at the seeded zero watermark.
	expected, err := repo.AllEventsForNotificationSinceEvtlog(ctx, addrs.RandomWalkAid, 0)
	if err != nil {
		t.Fatalf("AllEventsForNotificationSinceEvtlog: %v", err)
	}
	if len(expected) == 0 {
		t.Fatal("fixture dataset has no notification events; the test needs at least one")
	}

	twitter := newFakeTwitter()
	cfg := Config{
		Data:          repo,
		RWalkAid:      addrs.RandomWalkAid,
		MarketAid:     addrs.MarketPlaceAid,
		Twitter:       twitter,
		Media:         staticMedia{body: []byte("img")},
		Withdrawal:    &fakeWithdrawal{amounts: []float64{1.25}},
		Logger:        slog.New(slog.DiscardHandler),
		ImagesBase:    "https://img",
		VideosBase:    "https://vid",
		DetailBase:    "https://detail",
		PollInterval:  5 * time.Millisecond,
		RPCRetryDelay: time.Millisecond,
		MaxAttempts:   3,
	}
	engine, err := New(cfg)
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	runCtx, cancel := context.WithCancel(ctx)
	done := make(chan error, 1)
	go func() { done <- engine.Run(runCtx) }()

	waitSignals(t, twitter.signal, len(expected))
	// Allow the final watermark write to land before stopping.
	deadline := time.Now().Add(5 * time.Second)
	lastEvtlog := expected[len(expected)-1].EvtLogId
	for {
		status, serr := repo.MessagingStatus(ctx)
		if serr == nil && status.EvtLogId == lastEvtlog {
			break
		}
		if time.Now().After(deadline) {
			t.Fatalf("watermark = %+v (err %v), want EvtLogId %d", status, serr, lastEvtlog)
		}
		time.Sleep(time.Millisecond)
	}
	cancel()
	if err := <-done; err != nil {
		t.Fatalf("Run: %v", err)
	}

	sent := twitter.sent()
	if len(sent) != len(expected) {
		t.Fatalf("tweets = %d, want %d (one per fixture event)", len(sent), len(expected))
	}
	// Spot-check the first fixture event's rendered text.
	first := expected[0]
	wantFirst := NotificationMessage(first.EvtType, first.TokenId, first.Price, 1.25, DetailURL("https://detail", first.TokenId))
	if sent[0].msg != wantFirst {
		t.Errorf("first tweet = %q, want %q", sent[0].msg, wantFirst)
	}

	// A second run starts from the persisted watermark and announces nothing.
	engine2, err := New(cfg)
	if err != nil {
		t.Fatalf("New (second run): %v", err)
	}
	runCtx2, cancel2 := context.WithCancel(ctx)
	done2 := make(chan error, 1)
	go func() { done2 <- engine2.Run(runCtx2) }()
	time.Sleep(50 * time.Millisecond)
	cancel2()
	if err := <-done2; err != nil {
		t.Fatalf("second Run: %v", err)
	}
	if again := twitter.sent(); len(again) != len(expected) {
		t.Errorf("second run announced %d extra events; the watermark did not persist", len(again)-len(expected))
	}

	// Restore the seeded watermark so test order cannot affect other suites.
	status, err := repo.MessagingStatus(ctx)
	if err != nil {
		t.Fatalf("MessagingStatus: %v", err)
	}
	status.EvtLogId = 0
	status.TimeStamp = 0
	if err := repo.UpdateMessagingStatus(ctx, &status); err != nil {
		t.Fatalf("restoring messaging status: %v", err)
	}
}
