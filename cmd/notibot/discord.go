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

type discordSink struct {
	client        *disgord.Client
	notifyChannel disgord.Snowflake
	channels      map[rwbot.StatChannel]disgord.Snowflake
	log           *slog.Logger
	sleep         func(time.Duration) // injected in tests
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
		log:   log,
		sleep: time.Sleep,
	}
}

// SendMessage posts the notification text with the token image attached and
// an embed linking to the detail page.
//
//nolint:staticcheck // SA1019: the deprecated disgord builder API is what the legacy bot shipped with; migrating is a separate change
func (d *discordSink) SendMessage(_ context.Context, text string, image []byte, detailURL string) error {
	_, err := d.client.Channel(d.notifyChannel).CreateMessage(&disgord.CreateMessageParams{
		Content: text,
		Files: []disgord.CreateMessageFileParams{
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
func (d *discordSink) SetChannelName(_ context.Context, ch rwbot.StatChannel, name string) error {
	id, ok := d.channels[ch]
	if !ok || id == 0 {
		return nil
	}
	rename := func() error {
		//nolint:staticcheck // SA1019: the deprecated disgord builder API is what the legacy bot shipped with
		_, err := d.client.Channel(id).UpdateBuilder().SetName(name).Execute()
		return err
	}
	return renameWithRetry(rename, d.sleep)
}

// renameWithRetry executes rename and, when Discord answers with a
// rate-limit error carrying retry_after, waits that many seconds plus one
// for safety and retries once.
func renameWithRetry(rename func() error, sleep func(time.Duration)) error {
	err := rename()
	if err == nil {
		return nil
	}
	delaySec, ok := rwbot.ParseRetryAfterSeconds(err.Error())
	if !ok {
		return err
	}
	sleep(time.Duration(delaySec)*time.Second + time.Second)
	return rename()
}
