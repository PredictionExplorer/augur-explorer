package main

// discordSink adapts disgord to the rwbot.Discord seam: notification
// messages with the token image attached, and statistics-channel renames
// with one retry after Discord's rate-limit delay.

import (
	"bytes"
	"context"
	"log/slog"
	"time"

	"github.com/andersfylling/disgord"

	"github.com/PredictionExplorer/augur-explorer/internal/notify/rwbot"
)

// discordCallTimeout bounds one Discord REST call (D22: no outbound call may
// wait forever). Message posts carry one token PNG at most; thirty seconds
// tolerates a slow API while a wedged one fails the notification instead of
// stalling the loop. Rate-limit waits between retries are deliberately not
// under this bound — only the calls themselves.
const discordCallTimeout = 30 * time.Second

type discordSink struct {
	client        *disgord.Client
	notifyChannel disgord.Snowflake
	channels      map[rwbot.StatChannel]disgord.Snowflake
	log           *slog.Logger
	sleep         func(context.Context, time.Duration) error // injected in tests
	callTimeout   time.Duration                              // tests shorten it
}

func newDiscordSink(client *disgord.Client, keys discordKeys, log *slog.Logger) *discordSink {
	return &discordSink{
		client:        client,
		notifyChannel: disgord.Snowflake(keys.ChannelID),
		channels: map[rwbot.StatChannel]disgord.Snowflake{
			rwbot.ChannelMints:      disgord.Snowflake(keys.MintStatsChanID),
			rwbot.ChannelPrice:      disgord.Snowflake(keys.PriceStatsChanID),
			rwbot.ChannelLastDate:   disgord.Snowflake(keys.DateStatsChanID),
			rwbot.ChannelLastReward: disgord.Snowflake(keys.RewardStatsChanID),
		},
		log:         log,
		sleep:       sleepContext,
		callTimeout: discordCallTimeout,
	}
}

func sleepContext(ctx context.Context, delay time.Duration) error {
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

// SendMessage posts the notification text with the token image attached and
// an embed linking to the detail page.
func (d *discordSink) SendMessage(ctx context.Context, text string, image []byte, detailURL string) error {
	ctx, cancel := context.WithTimeout(ctx, d.callTimeout)
	defer cancel()
	_, err := d.client.Channel(d.notifyChannel).WithContext(ctx).CreateMessage(&disgord.CreateMessage{
		Content: text,
		Files: []disgord.CreateMessageFile{
			{Reader: bytes.NewReader(image), FileName: "token.png", SpoilerTag: false},
		},
		Embed: &disgord.Embed{
			Description: detailURL,
			URL:         detailURL,
		},
	})
	return err
}

// SetChannelName renames a statistics channel. Channels without a configured
// id are skipped; a rate-limited rename is retried once after the delay
// Discord reports.
func (d *discordSink) SetChannelName(ctx context.Context, ch rwbot.StatChannel, name string) error {
	id, ok := d.channels[ch]
	if !ok || id == 0 {
		return nil
	}
	rename := func() error {
		callCtx, cancel := context.WithTimeout(ctx, d.callTimeout)
		defer cancel()
		_, err := d.client.Channel(id).WithContext(callCtx).Update(&disgord.UpdateChannel{Name: &name})
		return err
	}
	return renameWithRetry(ctx, rename, d.sleep)
}

// renameWithRetry executes rename and, when Discord answers with a
// rate-limit error carrying retry_after, waits that many seconds plus one
// for safety and retries once.
func renameWithRetry(ctx context.Context, rename func() error, sleep func(context.Context, time.Duration) error) error {
	err := rename()
	if err == nil {
		return nil
	}
	delaySec, ok := rwbot.ParseRetryAfterSeconds(err.Error())
	if !ok {
		return err
	}
	if err := sleep(ctx, time.Duration(delaySec)*time.Second+time.Second); err != nil {
		return err
	}
	return rename()
}
