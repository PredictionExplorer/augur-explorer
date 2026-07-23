package rwbot

// Engine loop tests against scripted fakes: every seam (store, Twitter,
// Discord, media, resampler, contract read) is replaced, so the tests pin
// the orchestration semantics — watermark persistence, retry/skip policy,
// floor-price dedup and shutdown — without any network or database.

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"sync"
	"testing"
	"time"

	rwmodel "github.com/PredictionExplorer/augur-explorer/internal/model/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// --- fakes ---

type fakeData struct {
	mu          sync.Mutex
	status      rwmodel.MsgStatus
	statusErr   error
	events      []rwmodel.NotificationEvent2
	eventsErr   error
	eventsWait  <-chan struct{}
	updateErr   error
	updates     []rwmodel.MsgStatus
	noOffers    bool
	floor       float64
	floorTok    int64
	floorErr    error
	lastMint    int64
	lastMintErr error
}

func (d *fakeData) MessagingStatus(context.Context) (rwmodel.MsgStatus, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.status, d.statusErr
}

func (d *fakeData) UpdateMessagingStatus(_ context.Context, status *rwmodel.MsgStatus) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.updateErr != nil {
		return d.updateErr
	}
	d.updates = append(d.updates, *status)
	d.status = *status
	return nil
}

func (d *fakeData) AllEventsForNotificationSinceEvtlog(ctx context.Context, _, since int64) ([]rwmodel.NotificationEvent2, error) {
	d.mu.Lock()
	wait := d.eventsWait
	d.mu.Unlock()
	if wait != nil {
		select {
		case <-wait:
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.eventsErr != nil {
		return nil, d.eventsErr
	}
	var out []rwmodel.NotificationEvent2
	for _, e := range d.events {
		if e.EvtLogId > since {
			out = append(out, e)
		}
	}
	return out, nil
}

func (d *fakeData) FloorPrice(context.Context, int64, int64) (bool, float64, int64, int64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.noOffers, d.floor, 0, d.floorTok, d.floorErr
}

func (d *fakeData) LastMintTimestamp(context.Context) (int64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.lastMint, d.lastMintErr
}

func (d *fakeData) setFloor(price float64, tokenID int64) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.floor, d.floorTok = price, tokenID
}

func (d *fakeData) watermarks() []rwmodel.MsgStatus {
	d.mu.Lock()
	defer d.mu.Unlock()
	return append([]rwmodel.MsgStatus(nil), d.updates...)
}

type sentTweet struct {
	msg     string
	image   []byte
	video   []byte
	replyTo string
}

type fakeTwitter struct {
	mu       sync.Mutex
	tweets   []sentTweet
	imageErr []error // FIFO of errors for TweetWithImage; nil entries succeed
	videoErr []error
	seq      int
	signal   chan struct{}
	onImage  func() // invoked at the start of every TweetWithImage
}

func newFakeTwitter() *fakeTwitter {
	return &fakeTwitter{signal: make(chan struct{}, 100)}
}

func (f *fakeTwitter) TweetWithImage(_ context.Context, msg string, image []byte) (string, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.onImage != nil {
		f.onImage()
	}
	var err error
	if len(f.imageErr) > 0 {
		err, f.imageErr = f.imageErr[0], f.imageErr[1:]
	}
	f.signal <- struct{}{}
	if err != nil {
		return "", err
	}
	f.seq++
	f.tweets = append(f.tweets, sentTweet{msg: msg, image: image})
	return fmt.Sprintf("tweet-%d", f.seq), nil
}

func (f *fakeTwitter) TweetWithVideo(_ context.Context, msg string, video []byte, replyTo string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	var err error
	if len(f.videoErr) > 0 {
		err, f.videoErr = f.videoErr[0], f.videoErr[1:]
	}
	f.signal <- struct{}{}
	if err != nil {
		return err
	}
	f.tweets = append(f.tweets, sentTweet{msg: msg, video: video, replyTo: replyTo})
	return nil
}

func (f *fakeTwitter) sent() []sentTweet {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]sentTweet(nil), f.tweets...)
}

type discordMessage struct {
	text      string
	image     []byte
	detailURL string
}

type channelRename struct {
	ch   StatChannel
	name string
}

type fakeDiscord struct {
	mu        sync.Mutex
	messages  []discordMessage
	renames   []channelRename
	sendErr   error
	renameErr error
	signal    chan struct{}
}

func newFakeDiscord() *fakeDiscord {
	return &fakeDiscord{signal: make(chan struct{}, 100)}
}

func (f *fakeDiscord) SendMessage(_ context.Context, text string, image []byte, detailURL string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.signal <- struct{}{}
	if f.sendErr != nil {
		return f.sendErr
	}
	f.messages = append(f.messages, discordMessage{text: text, image: image, detailURL: detailURL})
	return nil
}

func (f *fakeDiscord) SetChannelName(_ context.Context, ch StatChannel, name string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.renameErr != nil {
		return f.renameErr
	}
	f.renames = append(f.renames, channelRename{ch: ch, name: name})
	return nil
}

func (f *fakeDiscord) sent() []discordMessage {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]discordMessage(nil), f.messages...)
}

func (f *fakeDiscord) renamed() []channelRename {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]channelRename(nil), f.renames...)
}

type lifecycleDiscord struct {
	lastDateStarted chan struct{}
	lastDateStopped chan struct{}
	startOnce       sync.Once
	stopOnce        sync.Once
}

func newLifecycleDiscord() *lifecycleDiscord {
	return &lifecycleDiscord{
		lastDateStarted: make(chan struct{}),
		lastDateStopped: make(chan struct{}),
	}
}

func (*lifecycleDiscord) SendMessage(context.Context, string, []byte, string) error {
	return nil
}

func (d *lifecycleDiscord) SetChannelName(ctx context.Context, ch StatChannel, _ string) error {
	if ch != ChannelLastDate {
		return nil
	}
	d.startOnce.Do(func() { close(d.lastDateStarted) })
	<-ctx.Done()
	d.stopOnce.Do(func() { close(d.lastDateStopped) })
	return ctx.Err()
}

type mediaResp struct {
	status int
	body   []byte
	err    error
}

type fakeMedia struct {
	mu        sync.Mutex
	responses map[string][]mediaResp // FIFO per URL; the last entry repeats
	onFetch   func(url string)       // invoked before answering
}

func newFakeMedia() *fakeMedia { return &fakeMedia{responses: make(map[string][]mediaResp)} }

func (f *fakeMedia) script(url string, resps ...mediaResp) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.responses[url] = append(f.responses[url], resps...)
}

func (f *fakeMedia) Fetch(_ context.Context, url string) (int, []byte, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.onFetch != nil {
		f.onFetch(url)
	}
	q := f.responses[url]
	if len(q) == 0 {
		return 0, nil, fmt.Errorf("fakeMedia: unscripted URL %s", url)
	}
	resp := q[0]
	if len(q) > 1 {
		f.responses[url] = q[1:]
	}
	return resp.status, resp.body, resp.err
}

type fakeWithdrawal struct {
	mu      sync.Mutex
	amounts []float64
	errs    []error // FIFO; nil entries succeed with amounts[0] (which also pops)
	onCall  func()  // invoked before every read
}

func (f *fakeWithdrawal) WithdrawalAmountEth(context.Context) (float64, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.onCall != nil {
		f.onCall()
	}
	if len(f.errs) > 0 {
		err := f.errs[0]
		f.errs = f.errs[1:]
		if err != nil {
			return 0, err
		}
	}
	if len(f.amounts) == 0 {
		return 3.5, nil
	}
	amount := f.amounts[0]
	if len(f.amounts) > 1 {
		f.amounts = f.amounts[1:]
	}
	return amount, nil
}

// --- helpers ---

// bot bundles the engine under test with all its fakes.
type bot struct {
	data       *fakeData
	twitter    *fakeTwitter
	discord    *fakeDiscord
	media      *fakeMedia
	withdrawal *fakeWithdrawal
	cfg        Config
}

func newBot() *bot {
	b := &bot{
		data:       &fakeData{floor: 1.0, floorTok: 1, lastMint: 0},
		twitter:    newFakeTwitter(),
		discord:    newFakeDiscord(),
		media:      newFakeMedia(),
		withdrawal: &fakeWithdrawal{},
	}
	b.cfg = Config{
		Data:               b.data,
		RWalkAid:           11,
		MarketAid:          22,
		Twitter:            b.twitter,
		Discord:            b.discord,
		Media:              b.media,
		Resample:           func(_ context.Context, video []byte) ([]byte, error) { return append([]byte("r:"), video...), nil },
		Withdrawal:         b.withdrawal,
		SendVideos:         true,
		Logger:             slog.New(slog.DiscardHandler),
		ImagesBase:         "https://img",
		VideosBase:         "https://vid",
		DetailBase:         "https://detail",
		PollInterval:       2 * time.Millisecond,
		NotReadyDelay:      time.Millisecond,
		RPCRetryDelay:      time.Millisecond,
		LastMintedInterval: time.Hour, // immediate update only, unless a test shortens it
		VideoErrorDelay:    time.Millisecond,
		MaxAttempts:        3,
	}
	return b
}

// run starts the engine and returns a cancel-and-wait function yielding
// Run's error.
func (b *bot) run(t *testing.T) (stop func() error) {
	t.Helper()
	engine, err := New(b.cfg)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)
	go func() { done <- engine.Run(ctx) }()
	return func() error {
		cancel()
		select {
		case err := <-done:
			return err
		case <-time.After(5 * time.Second):
			t.Fatal("engine did not stop after cancellation")
			return nil
		}
	}
}

// waitSignals reads n signals from ch, failing the test on timeout.
func waitSignals(t *testing.T, ch chan struct{}, n int) {
	t.Helper()
	for i := range n {
		select {
		case <-ch:
		case <-time.After(5 * time.Second):
			t.Fatalf("timed out waiting for signal %d/%d", i+1, n)
		}
	}
}

func mintEvent(evtlogID, tokenID int64, price float64) rwmodel.NotificationEvent2 {
	return rwmodel.NotificationEvent2{
		TokenId: tokenID, EvtLogId: evtlogID, TimeStampMinted: 1700000000 + evtlogID,
		Price: price, EvtType: EventMint,
	}
}

// --- config validation ---

func TestNewValidatesConfig(t *testing.T) {
	base := func() Config {
		b := newBot()
		return b.cfg
	}
	cases := []struct {
		name   string
		mutate func(*Config)
		want   string
	}{
		{"missing data", func(c *Config) { c.Data = nil }, "Config.Data"},
		{"missing media", func(c *Config) { c.Media = nil }, "Config.Media"},
		{"missing withdrawal", func(c *Config) { c.Withdrawal = nil }, "Config.Withdrawal"},
		{"no sinks", func(c *Config) { c.Twitter = nil; c.Discord = nil }, "at least one"},
		{"videos without resampler", func(c *Config) { c.Resample = nil }, "Config.Resample"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cfg := base()
			tc.mutate(&cfg)
			_, err := New(cfg)
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Errorf("New error = %v, want mention of %q", err, tc.want)
			}
		})
	}
}

func TestNewAppliesDefaults(t *testing.T) {
	b := newBot()
	b.cfg.ImagesBase = ""
	b.cfg.VideosBase = ""
	b.cfg.DetailBase = ""
	b.cfg.PollInterval = 0
	b.cfg.MaxAttempts = 0
	b.cfg.Logger = nil
	b.cfg.Now = nil
	engine, err := New(b.cfg)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	if engine.cfg.ImagesBase != DefaultImagesBase || engine.cfg.DetailBase != DefaultDetailBase {
		t.Error("URL defaults not applied")
	}
	if engine.cfg.PollInterval != DefaultPollInterval || engine.cfg.MaxAttempts != DefaultMaxAttempts {
		t.Error("timing defaults not applied")
	}
	if engine.cfg.Logger == nil || engine.cfg.Now == nil {
		t.Error("logger/clock defaults not applied")
	}
}

// --- event processing ---

func TestMintEventNotifiesAllSinksAndPersistsWatermark(t *testing.T) {
	b := newBot()
	b.data.events = []rwmodel.NotificationEvent2{mintEvent(100, 42, 0.5)}
	b.withdrawal.amounts = []float64{3.5}
	b.media.script(ImageURL("https://img", 42), mediaResp{status: 200, body: []byte("img42")})
	b.media.script(VideoURL("https://vid", 42), mediaResp{status: 200, body: []byte("vid42")})

	stop := b.run(t)
	waitSignals(t, b.twitter.signal, 2) // image tweet + video reply
	waitSignals(t, b.discord.signal, 1)
	if err := stop(); err != nil {
		t.Fatalf("Run: %v", err)
	}

	tweets := b.twitter.sent()
	if len(tweets) != 2 {
		t.Fatalf("tweets = %d, want image + video", len(tweets))
	}
	wantMsg := "#42 Minted for 0.5000Ξ.\nLast minter would get 3.50Ξ if there is no other mint for 30 days.\n\nhttps://detail/42"
	if tweets[0].msg != wantMsg || string(tweets[0].image) != "img42" {
		t.Errorf("image tweet = %+v, want msg %q", tweets[0], wantMsg)
	}
	if tweets[1].msg != "Video for token 42" || string(tweets[1].video) != "r:vid42" || tweets[1].replyTo != "tweet-1" {
		t.Errorf("video reply = %+v (resampled video, threaded under tweet-1 expected)", tweets[1])
	}

	msgs := b.discord.sent()
	if len(msgs) != 1 {
		t.Fatalf("discord messages = %d, want 1", len(msgs))
	}
	wantDiscord := "#42 Minted for 0.5000Ξ.\nLast minter would get 3.50Ξ if there is no other mint for 30 days."
	if msgs[0].text != wantDiscord || msgs[0].detailURL != "https://detail/42" {
		t.Errorf("discord message = %+v", msgs[0])
	}

	marks := b.data.watermarks()
	if len(marks) == 0 || marks[len(marks)-1].EvtLogId != 100 {
		t.Errorf("watermarks = %+v, want final EvtLogId 100", marks)
	}

	// Mint statistics channels renamed (price, mints, reward at minimum).
	names := map[StatChannel]string{}
	for _, r := range b.discord.renamed() {
		names[r.ch] = r.name
	}
	if names[ChannelPrice] != "Cur. price 💲 : 0.5000" {
		t.Errorf("price channel = %q", names[ChannelPrice])
	}
	if names[ChannelMints] != "Num. mints 🪙  : 43" {
		t.Errorf("mints channel = %q", names[ChannelMints])
	}
	if !strings.HasPrefix(names[ChannelLastReward], "Last reward: 3.50") {
		t.Errorf("reward channel = %q", names[ChannelLastReward])
	}
}

func TestNonMintEventsSendSingleTweetWithoutVideo(t *testing.T) {
	b := newBot()
	b.data.events = []rwmodel.NotificationEvent2{
		{TokenId: 7, EvtLogId: 10, TimeStampMinted: 1, Price: 2.5, EvtType: EventSellOffer},
		{TokenId: 8, EvtLogId: 11, TimeStampMinted: 2, Price: 1.5, EvtType: EventBought},
		{TokenId: 9, EvtLogId: 12, TimeStampMinted: 3, Price: 0.25, EvtType: EventBuyOffer},
	}
	for _, tok := range []int64{7, 8, 9} {
		b.media.script(ImageURL("https://img", tok), mediaResp{status: 200, body: []byte("img")})
	}

	stop := b.run(t)
	waitSignals(t, b.twitter.signal, 3)
	waitSignals(t, b.discord.signal, 3)
	if err := stop(); err != nil {
		t.Fatalf("Run: %v", err)
	}

	tweets := b.twitter.sent()
	if len(tweets) != 3 {
		t.Fatalf("tweets = %d, want 3 single-image tweets", len(tweets))
	}
	wants := []string{
		"#7 On sale for 2.5000Ξ.\n\nhttps://detail/7",
		"#8 Bought for 1.5000Ξ.\n\nhttps://detail/8",
		"Buy offer for 0.2500Ξ.\n\nhttps://detail/9",
	}
	for i, want := range wants {
		if tweets[i].msg != want {
			t.Errorf("tweet %d = %q, want %q", i, tweets[i].msg, want)
		}
		if tweets[i].video != nil {
			t.Errorf("tweet %d carries a video", i)
		}
	}
	marks := b.data.watermarks()
	if len(marks) != 3 || marks[2].EvtLogId != 12 {
		t.Errorf("watermarks = %+v, want 3 advancing to 12", marks)
	}
}

func TestImage403SkipsEventAndAdvancesWatermark(t *testing.T) {
	b := newBot()
	b.data.events = []rwmodel.NotificationEvent2{
		{TokenId: 100, EvtLogId: 20, TimeStampMinted: 1, Price: 1, EvtType: EventSellOffer},
		{TokenId: 101, EvtLogId: 21, TimeStampMinted: 2, Price: 2, EvtType: EventSellOffer},
	}
	b.media.script(ImageURL("https://img", 100), mediaResp{status: 403})
	b.media.script(ImageURL("https://img", 101), mediaResp{status: 200, body: []byte("ok")})

	stop := b.run(t)
	waitSignals(t, b.twitter.signal, 1) // only token 101 is announced
	if err := stop(); err != nil {
		t.Fatalf("Run: %v", err)
	}

	tweets := b.twitter.sent()
	if len(tweets) != 1 || !strings.Contains(tweets[0].msg, "#101") {
		t.Errorf("tweets = %+v, want only token 101", tweets)
	}
	marks := b.data.watermarks()
	if len(marks) != 2 || marks[0].EvtLogId != 20 || marks[1].EvtLogId != 21 {
		t.Errorf("watermarks = %+v, want the 403 event persisted as processed (20) then 21", marks)
	}
}

func TestImage404WaitsUntilGenerated(t *testing.T) {
	b := newBot()
	b.data.events = []rwmodel.NotificationEvent2{{TokenId: 5, EvtLogId: 30, TimeStampMinted: 1, Price: 1, EvtType: EventBought}}
	b.media.script(ImageURL("https://img", 5),
		mediaResp{status: 404}, mediaResp{status: 404}, mediaResp{status: 200, body: []byte("late")})

	stop := b.run(t)
	waitSignals(t, b.twitter.signal, 1)
	if err := stop(); err != nil {
		t.Fatalf("Run: %v", err)
	}
	tweets := b.twitter.sent()
	if len(tweets) != 1 || string(tweets[0].image) != "late" {
		t.Errorf("tweets = %+v, want the late image announced", tweets)
	}
}

func TestImage404BeyondMaxAttemptsRetriesEventLater(t *testing.T) {
	b := newBot()
	b.cfg.MaxAttempts = 1
	b.data.events = []rwmodel.NotificationEvent2{{TokenId: 6, EvtLogId: 31, TimeStampMinted: 1, Price: 1, EvtType: EventBought}}
	// Two 404 rounds exhaust MaxAttempts=1, then the event is retried on the
	// next poll and succeeds.
	b.media.script(ImageURL("https://img", 6),
		mediaResp{status: 404}, mediaResp{status: 404},
		mediaResp{status: 200, body: []byte("eventually")})

	stop := b.run(t)
	waitSignals(t, b.twitter.signal, 1)
	if err := stop(); err != nil {
		t.Fatalf("Run: %v", err)
	}
	if tweets := b.twitter.sent(); len(tweets) != 1 || string(tweets[0].image) != "eventually" {
		t.Errorf("tweets = %+v, want retried event announced", tweets)
	}
}

func TestMediaTransportErrorRetriesWithoutAdvancing(t *testing.T) {
	b := newBot()
	b.data.events = []rwmodel.NotificationEvent2{{TokenId: 3, EvtLogId: 40, TimeStampMinted: 1, Price: 1, EvtType: EventSellOffer}}
	b.media.script(ImageURL("https://img", 3),
		mediaResp{err: errors.New("connection refused")},
		mediaResp{status: 200, body: []byte("recovered")})

	stop := b.run(t)
	waitSignals(t, b.twitter.signal, 1)
	if err := stop(); err != nil {
		t.Fatalf("Run: %v", err)
	}
	tweets := b.twitter.sent()
	if len(tweets) != 1 || string(tweets[0].image) != "recovered" {
		t.Errorf("tweets = %+v, want event announced after transport recovery", tweets)
	}
	marks := b.data.watermarks()
	if len(marks) != 1 || marks[0].EvtLogId != 40 {
		t.Errorf("watermarks = %+v, want exactly one persist after the successful retry", marks)
	}
}

func TestUnexpectedMediaStatusRetries(t *testing.T) {
	b := newBot()
	b.data.events = []rwmodel.NotificationEvent2{{TokenId: 4, EvtLogId: 41, TimeStampMinted: 1, Price: 1, EvtType: EventSellOffer}}
	b.media.script(ImageURL("https://img", 4),
		mediaResp{status: 500},
		mediaResp{status: 200, body: []byte("after 500")})

	stop := b.run(t)
	waitSignals(t, b.twitter.signal, 1)
	if err := stop(); err != nil {
		t.Fatalf("Run: %v", err)
	}
	if tweets := b.twitter.sent(); len(tweets) != 1 || string(tweets[0].image) != "after 500" {
		t.Errorf("tweets = %+v", tweets)
	}
}

func TestVideo403FallsBackToImageOnlyTweet(t *testing.T) {
	b := newBot()
	b.data.events = []rwmodel.NotificationEvent2{mintEvent(50, 9, 1.0)}
	b.media.script(ImageURL("https://img", 9), mediaResp{status: 200, body: []byte("img")})
	b.media.script(VideoURL("https://vid", 9), mediaResp{status: 403})

	stop := b.run(t)
	waitSignals(t, b.twitter.signal, 1)
	if err := stop(); err != nil {
		t.Fatalf("Run: %v", err)
	}
	tweets := b.twitter.sent()
	if len(tweets) != 1 || tweets[0].video != nil {
		t.Errorf("tweets = %+v, want a single image-only tweet", tweets)
	}
	if marks := b.data.watermarks(); len(marks) != 1 || marks[0].EvtLogId != 50 {
		t.Errorf("watermarks = %+v", marks)
	}
}

func TestResampleFailureRetriesEvent(t *testing.T) {
	b := newBot()
	calls := 0
	b.cfg.Resample = func(_ context.Context, video []byte) ([]byte, error) {
		calls++
		if calls == 1 {
			return nil, errors.New("ffmpeg exploded")
		}
		return video, nil
	}
	b.data.events = []rwmodel.NotificationEvent2{mintEvent(60, 12, 1.0)}
	b.media.script(ImageURL("https://img", 12), mediaResp{status: 200, body: []byte("img")})
	b.media.script(VideoURL("https://vid", 12), mediaResp{status: 200, body: []byte("vid")})

	stop := b.run(t)
	waitSignals(t, b.twitter.signal, 2)
	if err := stop(); err != nil {
		t.Fatalf("Run: %v", err)
	}
	if calls < 2 {
		t.Errorf("resample calls = %d, want a retry after the failure", calls)
	}
	tweets := b.twitter.sent()
	if len(tweets) != 2 || string(tweets[1].video) != "vid" {
		t.Errorf("tweets = %+v, want image + video after resample retry", tweets)
	}
}

func TestSendVideosDisabledSkipsVideoFetch(t *testing.T) {
	b := newBot()
	b.cfg.SendVideos = false
	b.cfg.Resample = nil
	b.data.events = []rwmodel.NotificationEvent2{mintEvent(70, 15, 1.0)}
	b.media.script(ImageURL("https://img", 15), mediaResp{status: 200, body: []byte("img")})
	// No video URL scripted: a fetch attempt would fail the fakeMedia.

	stop := b.run(t)
	waitSignals(t, b.twitter.signal, 1)
	if err := stop(); err != nil {
		t.Fatalf("Run: %v", err)
	}
	tweets := b.twitter.sent()
	if len(tweets) != 1 || tweets[0].video != nil {
		t.Errorf("tweets = %+v, want single image tweet", tweets)
	}
}

func TestWithdrawalRetriesThenSucceeds(t *testing.T) {
	b := newBot()
	b.withdrawal.errs = []error{errors.New("rpc blip"), errors.New("rpc blip"), nil}
	b.withdrawal.amounts = []float64{7.25}
	b.data.events = []rwmodel.NotificationEvent2{mintEvent(80, 2, 1.0)}
	b.media.script(ImageURL("https://img", 2), mediaResp{status: 200, body: []byte("img")})
	b.media.script(VideoURL("https://vid", 2), mediaResp{status: 200, body: []byte("vid")})

	stop := b.run(t)
	waitSignals(t, b.twitter.signal, 2)
	if err := stop(); err != nil {
		t.Fatalf("Run: %v", err)
	}
	tweets := b.twitter.sent()
	if len(tweets) == 0 || !strings.Contains(tweets[0].msg, "7.25Ξ") {
		t.Errorf("tweets = %+v, want withdrawal amount 7.25 in the mint text", tweets)
	}
}

func TestTwitterFailureStillAdvancesWatermark(t *testing.T) {
	b := newBot()
	b.data.events = []rwmodel.NotificationEvent2{{TokenId: 1, EvtLogId: 90, TimeStampMinted: 1, Price: 1, EvtType: EventSellOffer}}
	b.twitter.imageErr = []error{errors.New("twitter down")}
	b.media.script(ImageURL("https://img", 1), mediaResp{status: 200, body: []byte("img")})

	stop := b.run(t)
	waitSignals(t, b.twitter.signal, 1)
	waitSignals(t, b.discord.signal, 1)
	if err := stop(); err != nil {
		t.Fatalf("Run: %v", err)
	}
	// Legacy semantics: notification failures are logged, never retried.
	if marks := b.data.watermarks(); len(marks) != 1 || marks[0].EvtLogId != 90 {
		t.Errorf("watermarks = %+v, want the event persisted despite the tweet failure", marks)
	}
	if msgs := b.discord.sent(); len(msgs) != 1 {
		t.Errorf("discord messages = %d, want 1 (independent of twitter failure)", len(msgs))
	}
}

func TestVideoReplyFailureIsLoggedNotRetried(t *testing.T) {
	b := newBot()
	b.data.events = []rwmodel.NotificationEvent2{mintEvent(95, 4, 1.0)}
	b.twitter.videoErr = []error{errors.New("upload failed")}
	b.media.script(ImageURL("https://img", 4), mediaResp{status: 200, body: []byte("img")})
	b.media.script(VideoURL("https://vid", 4), mediaResp{status: 200, body: []byte("vid")})

	stop := b.run(t)
	waitSignals(t, b.twitter.signal, 2)
	if err := stop(); err != nil {
		t.Fatalf("Run: %v", err)
	}
	if marks := b.data.watermarks(); len(marks) != 1 || marks[0].EvtLogId != 95 {
		t.Errorf("watermarks = %+v, want event persisted despite video failure", marks)
	}
}

// --- floor price ---

func TestFloorPriceChangeAnnouncedOnceWithDedup(t *testing.T) {
	b := newBot()
	// A sentinel event marks the moment startup seeding (floor price 1.0) is
	// over; only then is the floor changed.
	b.data.events = []rwmodel.NotificationEvent2{{TokenId: 1, EvtLogId: 5, TimeStampMinted: 1, Price: 1, EvtType: EventSellOffer}}
	b.media.script(ImageURL("https://img", 1), mediaResp{status: 200, body: []byte("img")})
	b.media.script(ImageURL("https://img", 33), mediaResp{status: 200, body: []byte("floorimg")})

	stop := b.run(t)
	waitSignals(t, b.twitter.signal, 1) // sentinel event announced; startup is over
	b.data.setFloor(2.5, 33)
	waitSignals(t, b.twitter.signal, 1) // the floor-price announcement
	waitSignals(t, b.discord.signal, 2)
	if err := stop(); err != nil {
		t.Fatalf("Run: %v", err)
	}

	tweets := b.twitter.sent()
	if len(tweets) != 2 {
		t.Fatalf("tweets = %+v, want sentinel + one floor-price announcement", tweets)
	}
	want := "Floor price changed to 2.5000Ξ.\n\nhttps://detail/33"
	if tweets[1].msg != want {
		t.Errorf("floor tweet = %q, want %q", tweets[1].msg, want)
	}
	msgs := b.discord.sent()
	if len(msgs) != 2 || msgs[1].text != "Floor price changed to 2.5000Ξ." {
		t.Errorf("discord floor messages = %+v", msgs)
	}
}

func TestFloorPriceUnchangedOrNoOffersStaysQuiet(t *testing.T) {
	b := newBot()
	b.data.noOffers = true

	stop := b.run(t)
	time.Sleep(20 * time.Millisecond) // several polls
	if err := stop(); err != nil {
		t.Fatalf("Run: %v", err)
	}
	if tweets := b.twitter.sent(); len(tweets) != 0 {
		t.Errorf("tweets = %+v, want none", tweets)
	}
}

func TestFloorPriceReadErrorIsNonFatal(t *testing.T) {
	b := newBot()
	// The sentinel event marks the end of startup (whose floor read IS
	// fatal); only in-loop floor reads are tolerant.
	b.data.events = []rwmodel.NotificationEvent2{{TokenId: 1, EvtLogId: 5, TimeStampMinted: 1, Price: 1, EvtType: EventSellOffer}}
	b.media.script(ImageURL("https://img", 1), mediaResp{status: 200, body: []byte("img")})

	stop := b.run(t)
	waitSignals(t, b.twitter.signal, 1)
	b.data.mu.Lock()
	b.data.floorErr = errors.New("db flake")
	b.data.mu.Unlock()
	time.Sleep(20 * time.Millisecond) // several polls hit the failing floor read
	if err := stop(); err != nil {
		t.Fatalf("Run returned %v, want nil (floor errors are logged and skipped)", err)
	}
}

// --- fatal paths and shutdown ---

func TestMessagingStatusErrorIsFatal(t *testing.T) {
	b := newBot()
	b.data.statusErr = errors.New("no watermark row")
	engine, err := New(b.cfg)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	if err := engine.Run(context.Background()); err == nil || !strings.Contains(err.Error(), "messaging status") {
		t.Errorf("Run = %v, want messaging-status error", err)
	}
}

func TestEventsFetchErrorIsFatal(t *testing.T) {
	b := newBot()
	b.data.eventsErr = errors.New("db gone")
	engine, err := New(b.cfg)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	if err := engine.Run(context.Background()); err == nil || !strings.Contains(err.Error(), "notification events") {
		t.Errorf("Run = %v, want events-fetch error", err)
	}
}

func TestFatalReturnCancelsAndJoinsLastMintedLoop(t *testing.T) {
	t.Parallel()

	b := newBot()
	b.data.lastMint = 1
	b.data.eventsErr = errors.New("db gone")
	discord := newLifecycleDiscord()
	b.cfg.Discord = discord
	b.data.eventsWait = discord.lastDateStarted

	engine, err := New(b.cfg)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	done := make(chan error, 1)
	go func() { done <- engine.Run(context.Background()) }()

	select {
	case err := <-done:
		if err == nil || !strings.Contains(err.Error(), "notification events") {
			t.Fatalf("Run = %v, want notification-events error", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("Run did not return after its fatal data-source error")
	}
	select {
	case <-discord.lastDateStarted:
	default:
		t.Fatal("last-minted sidecar did not start")
	}
	select {
	case <-discord.lastDateStopped:
	default:
		t.Fatal("Run returned before the last-minted sidecar stopped")
	}
}

func TestWatermarkPersistErrorIsFatal(t *testing.T) {
	b := newBot()
	b.data.events = []rwmodel.NotificationEvent2{{TokenId: 1, EvtLogId: 5, TimeStampMinted: 1, Price: 1, EvtType: EventSellOffer}}
	b.data.updateErr = errors.New("write failed")
	b.media.script(ImageURL("https://img", 1), mediaResp{status: 200, body: []byte("img")})
	engine, err := New(b.cfg)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	if err := engine.Run(context.Background()); err == nil || !strings.Contains(err.Error(), "messaging status") {
		t.Errorf("Run = %v, want watermark-persist error", err)
	}
}

func TestInitialFloorPriceErrorIsFatal(t *testing.T) {
	b := newBot()
	b.data.floorErr = errors.New("no db")
	engine, err := New(b.cfg)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	if err := engine.Run(context.Background()); err == nil || !strings.Contains(err.Error(), "floor price") {
		t.Errorf("Run = %v, want initial floor-price error", err)
	}
}

func TestLastMintNotFoundIsTolerated(t *testing.T) {
	b := newBot()
	b.data.lastMintErr = store.ErrNotFound
	stop := b.run(t)
	time.Sleep(10 * time.Millisecond)
	if err := stop(); err != nil {
		t.Fatalf("Run = %v, want nil on an empty-database mint history", err)
	}
}

func TestLastMintOtherErrorIsFatal(t *testing.T) {
	b := newBot()
	b.data.lastMintErr = errors.New("db broken")
	engine, err := New(b.cfg)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	if err := engine.Run(context.Background()); err == nil || !strings.Contains(err.Error(), "last mint") {
		t.Errorf("Run = %v, want last-mint error", err)
	}
}

func TestCancelledContextStopsCleanly(t *testing.T) {
	b := newBot()
	stop := b.run(t)
	time.Sleep(5 * time.Millisecond)
	if err := stop(); err != nil {
		t.Fatalf("Run after cancel = %v, want nil", err)
	}
}

func TestCancelDuringMediaWaitStopsCleanly(t *testing.T) {
	b := newBot()
	b.cfg.NotReadyDelay = time.Hour // cancellation must interrupt the 404 wait
	b.data.events = []rwmodel.NotificationEvent2{{TokenId: 1, EvtLogId: 5, TimeStampMinted: 1, Price: 1, EvtType: EventSellOffer}}
	b.media.script(ImageURL("https://img", 1), mediaResp{status: 404})

	stop := b.run(t)
	time.Sleep(10 * time.Millisecond)
	if err := stop(); err != nil {
		t.Fatalf("Run = %v, want nil when cancelled mid-wait", err)
	}
}

// --- discord statistics channels ---

func TestLastMintedLoopRenamesChannel(t *testing.T) {
	b := newBot()
	b.data.lastMint = 1700000000
	fixed := time.Unix(1700000000, 0).Add(3 * time.Hour)
	b.cfg.Now = func() time.Time { return fixed }

	stop := b.run(t)
	deadline := time.After(5 * time.Second)
	for {
		found := false
		for _, r := range b.discord.renamed() {
			if r.ch == ChannelLastDate && strings.Contains(r.name, "3 hours ago") {
				found = true
			}
		}
		if found {
			break
		}
		select {
		case <-deadline:
			t.Fatalf("last-mint channel never renamed; renames = %+v", b.discord.renamed())
		case <-time.After(time.Millisecond):
		}
	}
	if err := stop(); err != nil {
		t.Fatalf("Run: %v", err)
	}
}

func TestStartupRenamesRewardChannel(t *testing.T) {
	b := newBot()
	b.withdrawal.amounts = []float64{9.5}
	stop := b.run(t)
	deadline := time.After(5 * time.Second)
	for {
		found := false
		for _, r := range b.discord.renamed() {
			if r.ch == ChannelLastReward && r.name == "Last reward: 9.50Ξ" {
				found = true
			}
		}
		if found {
			break
		}
		select {
		case <-deadline:
			t.Fatalf("reward channel never renamed; renames = %+v", b.discord.renamed())
		case <-time.After(time.Millisecond):
		}
	}
	if err := stop(); err != nil {
		t.Fatalf("Run: %v", err)
	}
}

func TestTwitterOnlyBotSkipsDiscordWork(t *testing.T) {
	b := newBot()
	b.cfg.Discord = nil
	b.data.events = []rwmodel.NotificationEvent2{{TokenId: 1, EvtLogId: 5, TimeStampMinted: 1, Price: 1, EvtType: EventSellOffer}}
	b.media.script(ImageURL("https://img", 1), mediaResp{status: 200, body: []byte("img")})

	stop := b.run(t)
	waitSignals(t, b.twitter.signal, 1)
	if err := stop(); err != nil {
		t.Fatalf("Run: %v", err)
	}
	if renames := b.discord.renamed(); len(renames) != 0 {
		t.Errorf("discord renames = %+v, want none for a twitter-only bot", renames)
	}
	if msgs := b.discord.sent(); len(msgs) != 0 {
		t.Errorf("discord messages = %+v, want none", msgs)
	}
}

// --- additional branch coverage ---

func TestMidBatchCancellationStopsBeforeNextEvent(t *testing.T) {
	b := newBot()
	b.data.events = []rwmodel.NotificationEvent2{
		{TokenId: 1, EvtLogId: 5, TimeStampMinted: 1, Price: 1, EvtType: EventSellOffer},
		{TokenId: 2, EvtLogId: 6, TimeStampMinted: 2, Price: 1, EvtType: EventSellOffer},
	}
	b.media.script(ImageURL("https://img", 1), mediaResp{status: 200, body: []byte("img")})
	b.media.script(ImageURL("https://img", 2), mediaResp{status: 200, body: []byte("img")})

	engine, err := New(b.cfg)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	// The first tweet cancels the run: the loop must stop before event 2.
	b.twitter.onImage = cancel
	if err := engine.Run(ctx); err != nil {
		t.Fatalf("Run = %v, want nil after mid-batch cancellation", err)
	}
	if tweets := b.twitter.sent(); len(tweets) != 1 {
		t.Errorf("tweets = %d, want only the first event announced", len(tweets))
	}
}

func TestWithdrawalExhaustionRetriesEventLater(t *testing.T) {
	b := newBot()
	b.cfg.MaxAttempts = 1
	b.cfg.Discord = nil // keep the startup reward-channel read out of the error budget
	// Two failures exhaust the first bounded read (MaxAttempts=1 allows two
	// attempts); the retried event's read fails once more, then succeeds.
	b.withdrawal.errs = []error{errors.New("rpc"), errors.New("rpc"), errors.New("rpc"), nil}
	b.withdrawal.amounts = []float64{4.5}
	b.data.events = []rwmodel.NotificationEvent2{mintEvent(80, 2, 1.0)}
	b.media.script(ImageURL("https://img", 2), mediaResp{status: 200, body: []byte("img")})
	b.media.script(VideoURL("https://vid", 2), mediaResp{status: 200, body: []byte("vid")})

	stop := b.run(t)
	waitSignals(t, b.twitter.signal, 2)
	if err := stop(); err != nil {
		t.Fatalf("Run: %v", err)
	}
	tweets := b.twitter.sent()
	if len(tweets) == 0 || !strings.Contains(tweets[0].msg, "4.50Ξ") {
		t.Errorf("tweets = %+v, want the retried mint announced with 4.50", tweets)
	}
}

func TestWithdrawalCancellationStopsCleanly(t *testing.T) {
	b := newBot()
	b.cfg.Discord = nil // the first withdrawal read must happen inside the event
	b.data.events = []rwmodel.NotificationEvent2{mintEvent(81, 3, 1.0)}

	engine, err := New(b.cfg)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	b.withdrawal.onCall = cancel
	b.withdrawal.errs = []error{errors.New("rpc gone")}
	if err := engine.Run(ctx); err != nil {
		t.Fatalf("Run = %v, want nil when cancelled during the withdrawal retry", err)
	}
}

func TestVideoTransportErrorRetriesEvent(t *testing.T) {
	b := newBot()
	b.data.events = []rwmodel.NotificationEvent2{mintEvent(82, 5, 1.0)}
	b.media.script(ImageURL("https://img", 5), mediaResp{status: 200, body: []byte("img")})
	b.media.script(VideoURL("https://vid", 5),
		mediaResp{err: errors.New("conn reset")},
		mediaResp{status: 200, body: []byte("vid")})

	stop := b.run(t)
	waitSignals(t, b.twitter.signal, 2)
	if err := stop(); err != nil {
		t.Fatalf("Run: %v", err)
	}
	tweets := b.twitter.sent()
	if len(tweets) != 2 || string(tweets[1].video) != "r:vid" {
		t.Errorf("tweets = %+v, want image + video after the transport retry", tweets)
	}
}

func TestVideoFetchCancellationStopsCleanly(t *testing.T) {
	b := newBot()
	b.data.events = []rwmodel.NotificationEvent2{mintEvent(85, 8, 1.0)}
	b.media.script(ImageURL("https://img", 8), mediaResp{status: 200, body: []byte("img")})
	b.media.script(VideoURL("https://vid", 8), mediaResp{err: errors.New("interrupted")})

	engine, err := New(b.cfg)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	b.media.onFetch = func(url string) {
		if strings.Contains(url, "vid") {
			cancel()
		}
	}
	if err := engine.Run(ctx); err != nil {
		t.Fatalf("Run = %v, want nil when cancelled during the video fetch", err)
	}
}

func TestResampleCancellationStopsCleanly(t *testing.T) {
	b := newBot()
	engine := func() *Engine {
		ctxCancel := func() {}
		b.cfg.Resample = func(_ context.Context, _ []byte) ([]byte, error) {
			ctxCancel()
			return nil, errors.New("interrupted")
		}
		e, err := New(b.cfg)
		if err != nil {
			t.Fatalf("New: %v", err)
		}
		ctx, cancel := context.WithCancel(context.Background())
		ctxCancel = cancel
		done := make(chan error, 1)
		go func() { done <- e.Run(ctx) }()
		t.Cleanup(func() {
			cancel()
			if err := <-done; err != nil {
				t.Errorf("Run = %v, want nil when cancelled during resampling", err)
			}
		})
		return e
	}
	b.data.events = []rwmodel.NotificationEvent2{mintEvent(83, 6, 1.0)}
	b.media.script(ImageURL("https://img", 6), mediaResp{status: 200, body: []byte("img")})
	b.media.script(VideoURL("https://vid", 6), mediaResp{status: 200, body: []byte("vid")})
	engine()
	time.Sleep(20 * time.Millisecond)
}

func TestDiscordOnlyBotSkipsTwitter(t *testing.T) {
	b := newBot()
	b.cfg.Twitter = nil
	b.cfg.SendVideos = false
	b.cfg.Resample = nil
	b.data.events = []rwmodel.NotificationEvent2{{TokenId: 1, EvtLogId: 5, TimeStampMinted: 1, Price: 1, EvtType: EventSellOffer}}
	b.media.script(ImageURL("https://img", 1), mediaResp{status: 200, body: []byte("img")})

	stop := b.run(t)
	waitSignals(t, b.discord.signal, 1)
	if err := stop(); err != nil {
		t.Fatalf("Run: %v", err)
	}
	if tweets := b.twitter.sent(); len(tweets) != 0 {
		t.Errorf("tweets = %+v, want none for a discord-only bot", tweets)
	}
	if msgs := b.discord.sent(); len(msgs) != 1 {
		t.Errorf("discord messages = %d, want 1", len(msgs))
	}
}

func TestFloorPriceImageFailureSkipsAnnouncement(t *testing.T) {
	b := newBot()
	b.data.events = []rwmodel.NotificationEvent2{{TokenId: 1, EvtLogId: 5, TimeStampMinted: 1, Price: 1, EvtType: EventSellOffer}}
	b.media.script(ImageURL("https://img", 1), mediaResp{status: 200, body: []byte("img")})
	b.media.script(ImageURL("https://img", 44), mediaResp{status: 403})

	stop := b.run(t)
	waitSignals(t, b.twitter.signal, 1) // sentinel event; startup over
	b.data.setFloor(9.0, 44)
	time.Sleep(20 * time.Millisecond) // floor polls hit the 403 image
	if err := stop(); err != nil {
		t.Fatalf("Run: %v", err)
	}
	for _, tw := range b.twitter.sent() {
		if strings.Contains(tw.msg, "Floor price") {
			t.Errorf("floor price announced despite image failure: %q", tw.msg)
		}
	}
}

func TestFloorPriceTweetFailureIsLogged(t *testing.T) {
	b := newBot()
	b.data.events = []rwmodel.NotificationEvent2{{TokenId: 1, EvtLogId: 5, TimeStampMinted: 1, Price: 1, EvtType: EventSellOffer}}
	b.media.script(ImageURL("https://img", 1), mediaResp{status: 200, body: []byte("img")})
	b.media.script(ImageURL("https://img", 45), mediaResp{status: 200, body: []byte("floor")})

	stop := b.run(t)
	waitSignals(t, b.twitter.signal, 1)
	b.twitter.mu.Lock()
	b.twitter.imageErr = []error{errors.New("twitter down")}
	b.twitter.mu.Unlock()
	b.data.setFloor(9.5, 45)
	waitSignals(t, b.twitter.signal, 1) // the failing floor tweet attempt
	waitSignals(t, b.discord.signal, 2) // sentinel + floor discord message still sent
	if err := stop(); err != nil {
		t.Fatalf("Run: %v", err)
	}
}

func TestDiscordRenameFailureIsNonFatal(t *testing.T) {
	b := newBot()
	b.discord.renameErr = errors.New("rate limited forever")
	b.data.events = []rwmodel.NotificationEvent2{mintEvent(84, 7, 1.0)}
	b.media.script(ImageURL("https://img", 7), mediaResp{status: 200, body: []byte("img")})
	b.media.script(VideoURL("https://vid", 7), mediaResp{status: 200, body: []byte("vid")})

	stop := b.run(t)
	waitSignals(t, b.twitter.signal, 2)
	if err := stop(); err != nil {
		t.Fatalf("Run = %v, want nil despite rename failures", err)
	}
	if marks := b.data.watermarks(); len(marks) != 1 {
		t.Errorf("watermarks = %+v, want the mint persisted", marks)
	}
}

func TestLastMintedTickerKeepsRenaming(t *testing.T) {
	b := newBot()
	b.data.lastMint = 1700000000
	b.cfg.LastMintedInterval = 2 * time.Millisecond
	fixed := time.Unix(1700000000, 0).Add(time.Hour)
	b.cfg.Now = func() time.Time { return fixed }

	stop := b.run(t)
	deadline := time.After(5 * time.Second)
	for {
		count := 0
		for _, r := range b.discord.renamed() {
			if r.ch == ChannelLastDate {
				count++
			}
		}
		if count >= 2 { // immediate update + at least one ticker firing
			break
		}
		select {
		case <-deadline:
			t.Fatalf("ticker renames = %d, want >= 2", count)
		case <-time.After(time.Millisecond):
		}
	}
	if err := stop(); err != nil {
		t.Fatalf("Run: %v", err)
	}
}

func TestSleepCtxNonPositiveDuration(t *testing.T) {
	if !sleepCtx(context.Background(), 0) {
		t.Error("sleepCtx(0) on a live context = false, want true")
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if sleepCtx(ctx, -time.Second) {
		t.Error("sleepCtx(<0) on a cancelled context = true, want false")
	}
}

func TestDiscordSendFailureIsNonFatal(t *testing.T) {
	b := newBot()
	b.discord.sendErr = errors.New("discord api down")
	b.data.events = []rwmodel.NotificationEvent2{{TokenId: 1, EvtLogId: 5, TimeStampMinted: 1, Price: 1, EvtType: EventSellOffer}}
	b.media.script(ImageURL("https://img", 1), mediaResp{status: 200, body: []byte("img")})

	stop := b.run(t)
	waitSignals(t, b.discord.signal, 1)
	if err := stop(); err != nil {
		t.Fatalf("Run = %v, want nil (discord failures are logged)", err)
	}
	if marks := b.data.watermarks(); len(marks) != 1 {
		t.Errorf("watermarks = %+v, want event persisted despite discord failure", marks)
	}
}
