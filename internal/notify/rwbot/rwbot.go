// Package rwbot is the RandomWalk notification bot engine shared by
// cmd/notibot (Twitter + Discord) and rwctl notify-bot / tweet-mints
// (Twitter only). It replaces the two divergent copies of the same monitor:
// one injected Engine polls the store for new mint/offer/purchase events and
// floor-price changes, fetches the generated token media, and announces each
// event through the configured sinks, persisting the rw_messaging_status
// watermark after every processed event so restarts resume exactly where the
// previous run stopped.
//
// Everything external is a seam: the database (DataSource, satisfied by
// *randomwalk.Repo), Twitter (Tweeter), Discord (Discord), media downloads
// (Fetcher), video resampling (ResampleFunc, ffmpeg in production) and the
// contract withdrawal-amount read (WithdrawalReader). The loop is therefore
// fully testable without network, chain or Postgres access.
package rwbot

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync/atomic"
	"time"

	rwmodel "github.com/PredictionExplorer/augur-explorer/internal/model/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// Production media endpoints, overridable through Config for tests.
const (
	// DefaultImagesBase is the base URL of the generated token images.
	DefaultImagesBase = "https://api.randomwalknft.com:1443/images/randomwalk"
	// DefaultVideosBase is the base URL of the generated token videos.
	DefaultVideosBase = "https://api.randomwalknft.com:1443/images/randomwalk"
	// DefaultDetailBase is the base URL of the token detail pages.
	DefaultDetailBase = "https://randomwalknft.com/detail"
)

// Default loop timings, preserved from the legacy bots.
const (
	// DefaultPollInterval separates polls when no new events exist.
	DefaultPollInterval = 30 * time.Second
	// DefaultNotReadyDelay is the wait between retries while a token's
	// media has not been generated yet (HTTP 404).
	DefaultNotReadyDelay = 60 * time.Second
	// DefaultRPCRetryDelay is the wait between withdrawal-amount RPC retries.
	DefaultRPCRetryDelay = time.Second
	// DefaultLastMintedInterval spaces the Discord "last mint ago" channel
	// renames (Discord rate-limits channel updates to roughly one per
	// several minutes).
	DefaultLastMintedInterval = 9 * time.Minute
	// DefaultVideoErrorDelay is the pause after a failed video reply tweet,
	// preserved from the legacy bot to avoid hammering Twitter.
	DefaultVideoErrorDelay = 30 * time.Second
	// DefaultMaxAttempts bounds the media-not-ready and RPC retry loops.
	DefaultMaxAttempts = 1000
)

// DataSource is the store surface the engine reads and writes. It is
// satisfied by *randomwalk.Repo.
type DataSource interface {
	MessagingStatus(ctx context.Context) (rwmodel.MsgStatus, error)
	UpdateMessagingStatus(ctx context.Context, status *rwmodel.MsgStatus) error
	AllEventsForNotificationSinceEvtlog(ctx context.Context, rwalkAid, startEvtlogID int64) ([]rwmodel.NotificationEvent2, error)
	FloorPrice(ctx context.Context, rwalkAid, marketAid int64) (noOffers bool, floorPrice float64, offerID, tokenID int64, err error)
	LastMintTimestamp(ctx context.Context) (int64, error)
}

// Tweeter posts notifications to Twitter.
type Tweeter interface {
	// TweetWithImage posts msg with an attached image and returns the tweet
	// id (for threading replies).
	TweetWithImage(ctx context.Context, msg string, image []byte) (tweetID string, err error)
	// TweetWithVideo posts msg with an attached video, threaded under
	// replyToID when non-empty.
	TweetWithVideo(ctx context.Context, msg string, video []byte, replyToID string) error
}

// StatChannel identifies one of the Discord statistics channels the bot
// renames.
type StatChannel int

// The Discord statistics channels.
const (
	// ChannelMints shows the total number of mints.
	ChannelMints StatChannel = iota
	// ChannelPrice shows the current mint price.
	ChannelPrice
	// ChannelLastDate shows how long ago the last mint happened.
	ChannelLastDate
	// ChannelLastReward shows the current withdrawal amount.
	ChannelLastReward
)

// Discord posts notifications and statistics-channel renames to Discord.
type Discord interface {
	// SendMessage posts text with an attached image and an embed linking to
	// detailURL into the notification channel.
	SendMessage(ctx context.Context, text string, image []byte, detailURL string) error
	// SetChannelName renames one statistics channel (best-effort; the
	// implementation owns rate-limit handling).
	SetChannelName(ctx context.Context, ch StatChannel, name string) error
}

// Fetcher downloads one URL. Status is the HTTP status code when a response
// arrived; err reports transport failures.
type Fetcher interface {
	Fetch(ctx context.Context, url string) (status int, body []byte, err error)
}

// ResampleFunc converts a fetched video for Twitter (production: ffmpeg to
// 640x480, which Twitter's upload size limits require).
type ResampleFunc func(ctx context.Context, video []byte) ([]byte, error)

// WithdrawalReader reads the RandomWalk contract's current withdrawal amount
// in ETH.
type WithdrawalReader interface {
	WithdrawalAmountEth(ctx context.Context) (float64, error)
}

// Config wires an Engine. Data, Media, Withdrawal and at least one of
// Twitter/Discord are required.
type Config struct {
	Data DataSource
	// RWalkAid and MarketAid are the resolved address ids of the RandomWalk
	// and marketplace contracts.
	RWalkAid  int64
	MarketAid int64

	Twitter Tweeter // nil disables Twitter
	Discord Discord // nil disables Discord

	Media      Fetcher
	Resample   ResampleFunc // required when SendVideos
	Withdrawal WithdrawalReader

	// SendVideos also posts the token video as a reply to mint tweets
	// (notibot behavior; rwctl's Twitter-only bot historically sent images
	// only).
	SendVideos bool

	Logger *slog.Logger

	ImagesBase string
	VideosBase string
	DetailBase string

	PollInterval       time.Duration
	NotReadyDelay      time.Duration
	RPCRetryDelay      time.Duration
	LastMintedInterval time.Duration
	VideoErrorDelay    time.Duration
	MaxAttempts        int

	// Now is the clock (tests inject a fixed one).
	Now func() time.Time
}

// Engine is the notification bot. Create with New, run with Run.
type Engine struct {
	cfg Config
	log *slog.Logger

	watermark     rwmodel.MsgStatus
	curFloorPrice float64
	lastMintTS    atomic.Int64
}

// New validates cfg, applies defaults and returns a ready Engine.
func New(cfg Config) (*Engine, error) {
	if cfg.Data == nil {
		return nil, errors.New("rwbot: Config.Data is required")
	}
	if cfg.Media == nil {
		return nil, errors.New("rwbot: Config.Media is required")
	}
	if cfg.Withdrawal == nil {
		return nil, errors.New("rwbot: Config.Withdrawal is required")
	}
	if cfg.Twitter == nil && cfg.Discord == nil {
		return nil, errors.New("rwbot: at least one of Config.Twitter or Config.Discord is required")
	}
	if cfg.SendVideos && cfg.Resample == nil {
		return nil, errors.New("rwbot: Config.Resample is required when SendVideos is set")
	}
	if cfg.Logger == nil {
		cfg.Logger = slog.New(slog.DiscardHandler)
	}
	if cfg.ImagesBase == "" {
		cfg.ImagesBase = DefaultImagesBase
	}
	if cfg.VideosBase == "" {
		cfg.VideosBase = DefaultVideosBase
	}
	if cfg.DetailBase == "" {
		cfg.DetailBase = DefaultDetailBase
	}
	if cfg.PollInterval <= 0 {
		cfg.PollInterval = DefaultPollInterval
	}
	if cfg.NotReadyDelay <= 0 {
		cfg.NotReadyDelay = DefaultNotReadyDelay
	}
	if cfg.RPCRetryDelay <= 0 {
		cfg.RPCRetryDelay = DefaultRPCRetryDelay
	}
	if cfg.LastMintedInterval <= 0 {
		cfg.LastMintedInterval = DefaultLastMintedInterval
	}
	if cfg.VideoErrorDelay <= 0 {
		cfg.VideoErrorDelay = DefaultVideoErrorDelay
	}
	if cfg.MaxAttempts <= 0 {
		cfg.MaxAttempts = DefaultMaxAttempts
	}
	if cfg.Now == nil {
		cfg.Now = time.Now
	}
	return &Engine{cfg: cfg, log: cfg.Logger}, nil
}

// Run polls for events until ctx is cancelled (clean nil return) or a
// database read/write fails (error return: the binaries keep the legacy
// crash-and-restart semantics, resuming from the persisted watermark).
func (e *Engine) Run(ctx context.Context) error {
	status, err := e.cfg.Data.MessagingStatus(ctx)
	if err != nil {
		return e.runErr(ctx, fmt.Errorf("reading messaging status: %w", err))
	}
	e.watermark = status

	lastMint, err := e.cfg.Data.LastMintTimestamp(ctx)
	switch {
	case errors.Is(err, store.ErrNotFound):
		lastMint = 0 // fresh database: no mints yet
	case err != nil:
		return e.runErr(ctx, fmt.Errorf("reading last mint timestamp: %w", err))
	}
	e.lastMintTS.Store(lastMint)

	// Seed the floor-price dedup with the current value so restarts do not
	// re-announce an unchanged floor price.
	_, floorPrice, _, _, err := e.cfg.Data.FloorPrice(ctx, e.cfg.RWalkAid, e.cfg.MarketAid)
	if err != nil {
		return e.runErr(ctx, fmt.Errorf("reading initial floor price: %w", err))
	}
	e.curFloorPrice = floorPrice

	e.log.Info("rwbot starting",
		"evtlog_watermark", e.watermark.EvtLogId,
		"timestamp_watermark", e.watermark.TimeStamp,
		"floor_price", floorPrice,
		"twitter", e.cfg.Twitter != nil,
		"discord", e.cfg.Discord != nil,
	)

	if e.cfg.Discord != nil {
		go e.lastMintedLoop(ctx)
		if amount, werr := e.withdrawalBounded(ctx); werr == nil {
			e.setChannelName(ctx, ChannelLastReward, RewardChannelName(amount))
		}
	}

	for {
		if ctx.Err() != nil {
			return nil
		}
		e.checkFloorPrice(ctx)

		records, err := e.cfg.Data.AllEventsForNotificationSinceEvtlog(ctx, e.cfg.RWalkAid, e.watermark.EvtLogId)
		if err != nil {
			return e.runErr(ctx, fmt.Errorf("fetching notification events: %w", err))
		}
		if len(records) > 0 {
			e.log.Info("processing notification events", "count", len(records), "since_evtlog", e.watermark.EvtLogId)
		}
		retry := false
		for i := range records {
			if ctx.Err() != nil {
				return nil
			}
			var perr error
			retry, perr = e.processEvent(ctx, &records[i])
			if perr != nil {
				return e.runErr(ctx, perr)
			}
			if retry {
				break
			}
		}
		if len(records) == 0 || retry {
			// Idle poll spacing; a media failure also backs off here
			// instead of immediately re-polling.
			if !sleepCtx(ctx, e.cfg.PollInterval) {
				return nil
			}
		}
	}
}

// runErr converts errors caused by context cancellation into the clean
// shutdown nil.
func (e *Engine) runErr(ctx context.Context, err error) error {
	if ctx.Err() != nil {
		return nil
	}
	return err
}

// processEvent handles one notification record. retry=true asks the caller
// to stop the batch and poll again later (transient media/RPC trouble); a
// non-nil error is fatal (watermark persistence failed).
func (e *Engine) processEvent(ctx context.Context, rec *rwmodel.NotificationEvent2) (retry bool, err error) {
	e.log.Info("processing event",
		"evt_type", rec.EvtType, "token_id", rec.TokenId, "price", rec.Price, "evtlog_id", rec.EvtLogId)

	var withdrawalEth float64
	if rec.EvtType == EventMint {
		amount, werr := e.withdrawalBounded(ctx)
		if werr != nil {
			if ctx.Err() != nil {
				return false, ctx.Err()
			}
			e.log.Error("couldn't get withdrawal amount, retrying later", "err", werr)
			return true, nil
		}
		withdrawalEth = amount
		e.setChannelName(ctx, ChannelPrice, PriceChannelName(rec.Price))
		e.setChannelName(ctx, ChannelMints, MintsChannelName(rec.TokenId+1))
		e.lastMintTS.Store(rec.TimeStampMinted)
		e.setChannelName(ctx, ChannelLastReward, RewardChannelName(withdrawalEth))
	}

	image, skip, ferr := e.fetchUntilReady(ctx, ImageURL(e.cfg.ImagesBase, rec.TokenId))
	if ferr != nil {
		if ctx.Err() != nil {
			return false, ctx.Err()
		}
		e.log.Error("couldn't fetch token image, retrying later", "token_id", rec.TokenId, "err", ferr)
		return true, nil
	}
	if skip {
		// The image server answers 403 for some early tokens; the event can
		// never be announced, so it is skipped permanently.
		e.log.Info("image server returned 403, skipping event", "token_id", rec.TokenId)
		return false, e.advanceWatermark(ctx, rec)
	}

	var video []byte
	if rec.EvtType == EventMint && e.cfg.SendVideos && e.cfg.Twitter != nil {
		videoData, vskip, verr := e.fetchUntilReady(ctx, VideoURL(e.cfg.VideosBase, rec.TokenId))
		switch {
		case verr != nil:
			if ctx.Err() != nil {
				return false, ctx.Err()
			}
			e.log.Error("couldn't fetch token video, retrying later", "token_id", rec.TokenId, "err", verr)
			return true, nil
		case vskip:
			e.log.Info("video server returned 403, announcing without video", "token_id", rec.TokenId)
		default:
			video, verr = e.cfg.Resample(ctx, videoData)
			if verr != nil {
				if ctx.Err() != nil {
					return false, ctx.Err()
				}
				e.log.Error("video resampling failed, retrying later", "token_id", rec.TokenId, "err", verr)
				return true, nil
			}
		}
	}

	e.notifyTwitter(ctx, rec, withdrawalEth, image, video)
	e.notifyDiscord(ctx, rec.EvtType, rec.TokenId, rec.Price, withdrawalEth, image)

	return false, e.advanceWatermark(ctx, rec)
}

// notifyTwitter announces one event on Twitter: mints with a fetched video
// get an image tweet plus a threaded video reply, everything else a single
// image tweet. Failures are logged, never retried (legacy semantics: the
// watermark advances regardless).
func (e *Engine) notifyTwitter(ctx context.Context, rec *rwmodel.NotificationEvent2, withdrawalEth float64, image, video []byte) {
	if e.cfg.Twitter == nil {
		return
	}
	msg := NotificationMessage(rec.EvtType, rec.TokenId, rec.Price, withdrawalEth, DetailURL(e.cfg.DetailBase, rec.TokenId))
	tweetID, err := e.cfg.Twitter.TweetWithImage(ctx, msg, image)
	if err != nil {
		e.log.Error("error sending tweet", "token_id", rec.TokenId, "err", err)
		return
	}
	e.log.Info("notified event to twitter", "evt_type", rec.EvtType, "token_id", rec.TokenId)
	if len(video) == 0 {
		return
	}
	if err := e.cfg.Twitter.TweetWithVideo(ctx, VideoReplyMessage(rec.TokenId), video, tweetID); err != nil {
		e.log.Error("error sending video reply", "token_id", rec.TokenId, "err", err)
		sleepCtx(ctx, e.cfg.VideoErrorDelay) // avoid hammering Twitter after upload failures
	}
}

// notifyDiscord announces one event in the Discord notification channel.
func (e *Engine) notifyDiscord(ctx context.Context, evtType, tokenID int64, price, withdrawalEth float64, image []byte) {
	if e.cfg.Discord == nil {
		return
	}
	msg := NotificationMessage(evtType, tokenID, price, withdrawalEth, "")
	if err := e.cfg.Discord.SendMessage(ctx, msg, image, DetailURL(e.cfg.DetailBase, tokenID)); err != nil {
		e.log.Error("error sending discord notification", "token_id", tokenID, "err", err)
		return
	}
	e.log.Info("notified event to discord", "evt_type", evtType, "token_id", tokenID)
}

// advanceWatermark records rec as processed, both in memory and in
// rw_messaging_status. A persistence failure is fatal so the process
// restarts rather than silently re-notifying history.
func (e *Engine) advanceWatermark(ctx context.Context, rec *rwmodel.NotificationEvent2) error {
	e.watermark.EvtLogId = rec.EvtLogId
	e.watermark.TimeStamp = rec.TimeStampMinted
	if err := e.cfg.Data.UpdateMessagingStatus(ctx, &e.watermark); err != nil {
		return fmt.Errorf("persisting messaging status: %w", err)
	}
	return nil
}

// checkFloorPrice announces marketplace floor-price changes. Read failures
// are logged and skipped (next poll retries); the price is recorded before
// announcing, so a media failure never repeats a floor-price notification.
func (e *Engine) checkFloorPrice(ctx context.Context) {
	noOffers, price, _, tokenID, err := e.cfg.Data.FloorPrice(ctx, e.cfg.RWalkAid, e.cfg.MarketAid)
	if err != nil {
		if ctx.Err() == nil {
			e.log.Error("can't get floor price", "err", err)
		}
		return
	}
	if noOffers || price == e.curFloorPrice {
		return
	}
	e.curFloorPrice = price
	e.log.Info("floor price change detected", "price", price, "token_id", tokenID)

	image, skip, err := e.fetchUntilReady(ctx, ImageURL(e.cfg.ImagesBase, tokenID))
	if err != nil || skip {
		e.log.Error("couldn't fetch image for floor price notification", "token_id", tokenID, "err", err, "skipped_403", skip)
		return
	}
	if e.cfg.Twitter != nil {
		msg := NotificationMessage(EventFloorPrice, tokenID, price, 0, DetailURL(e.cfg.DetailBase, tokenID))
		if _, err := e.cfg.Twitter.TweetWithImage(ctx, msg, image); err != nil {
			e.log.Error("error tweeting floor price change", "err", err)
		} else {
			e.log.Info("notified floor price change to twitter", "price", price)
		}
	}
	e.notifyDiscord(ctx, EventFloorPrice, tokenID, price, 0, image)
}

// fetchUntilReady downloads a media URL, waiting while the file has not been
// generated yet (HTTP 404, bounded by MaxAttempts) and reporting skip=true
// on HTTP 403 (the image server's permanent refusal for some early tokens).
// Any other non-200 status or transport failure is an error.
func (e *Engine) fetchUntilReady(ctx context.Context, url string) (data []byte, skip bool, err error) {
	for attempts := 0; ; attempts++ {
		status, body, ferr := e.cfg.Media.Fetch(ctx, url)
		switch {
		case ferr != nil:
			return nil, false, ferr
		case status == 200:
			return body, false, nil
		case status == 403:
			return nil, true, nil
		case status == 404:
			e.log.Info("media not ready yet, waiting", "url", url)
			if attempts >= e.cfg.MaxAttempts {
				return nil, false, fmt.Errorf("media at %s not ready after %d attempts", url, attempts)
			}
			if !sleepCtx(ctx, e.cfg.NotReadyDelay) {
				return nil, false, ctx.Err()
			}
		default:
			return nil, false, fmt.Errorf("fetching %s: HTTP status %d", url, status)
		}
	}
}

// withdrawalBounded reads the contract withdrawal amount, retrying transient
// RPC failures every RPCRetryDelay, bounded by MaxAttempts.
func (e *Engine) withdrawalBounded(ctx context.Context) (float64, error) {
	var lastErr error
	for attempts := 0; attempts <= e.cfg.MaxAttempts; attempts++ {
		amount, err := e.cfg.Withdrawal.WithdrawalAmountEth(ctx)
		if err == nil {
			return amount, nil
		}
		lastErr = err
		e.log.Error("error getting withdrawal amount", "err", err)
		if !sleepCtx(ctx, e.cfg.RPCRetryDelay) {
			return 0, ctx.Err()
		}
	}
	return 0, fmt.Errorf("withdrawal amount unavailable: %w", lastErr)
}

// setChannelName renames a Discord statistics channel, best-effort.
func (e *Engine) setChannelName(ctx context.Context, ch StatChannel, name string) {
	if e.cfg.Discord == nil {
		return
	}
	if err := e.cfg.Discord.SetChannelName(ctx, ch, name); err != nil {
		e.log.Error("couldn't rename discord channel", "channel", int(ch), "name", name, "err", err)
		return
	}
	e.log.Info("renamed discord channel", "channel", int(ch), "name", name)
}

// lastMintedLoop keeps the Discord "last mint ago" channel fresh: one
// immediate update, then one per LastMintedInterval, until ctx is cancelled.
func (e *Engine) lastMintedLoop(ctx context.Context) {
	e.updateLastMinted(ctx)
	ticker := time.NewTicker(e.cfg.LastMintedInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			e.updateLastMinted(ctx)
		}
	}
}

func (e *Engine) updateLastMinted(ctx context.Context) {
	ts := e.lastMintTS.Load()
	if ts <= 0 {
		return
	}
	name := LastMintChannelName(time.Unix(ts, 0), e.cfg.Now())
	e.setChannelName(ctx, ChannelLastDate, name)
}

// sleepCtx sleeps for d unless ctx is cancelled first; it reports whether
// the full sleep completed.
func sleepCtx(ctx context.Context, d time.Duration) bool {
	if d <= 0 {
		return ctx.Err() == nil
	}
	timer := time.NewTimer(d)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return false
	case <-timer.C:
		return true
	}
}
