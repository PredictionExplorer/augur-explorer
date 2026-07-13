package main

// Credential-file loading. Both files live under $HOME/configs and their
// names come from the TWITTER_KEYS_FILE / DISCORD_KEYS_FILE configuration
// values, exactly like the legacy bot.

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/PredictionExplorer/augur-explorer/internal/notify/tweets"
)

// discordKeys is the JSON shape of the Discord credentials file. The tags
// pin the legacy key spellings operators already have on disk.
type discordKeys struct {
	TokenKey          string
	ChannelID         uint64 `json:"ChannelId"`     // notifications channel
	MainChannelID     uint64 `json:"MainChannelId"` // main chat channel
	MintStatsChanID   uint64 `json:"MintStatsChanId"`
	PriceStatsChanID  uint64 `json:"PriceStatsChanId"`
	DateStatsChanID   uint64 `json:"DateStatsChanId"`
	RewardStatsChanID uint64 `json:"RewardStatsChanId"`
}

// keysPath resolves the credential file name under home/configs.
func keysPath(home, name string) string {
	return filepath.Join(home, "configs", name)
}

// readTwitterKeys loads the Twitter API credentials.
func readTwitterKeys(home, name string) (tweets.TwitterKeys, error) {
	var keys tweets.TwitterKeys
	path := keysPath(home, name)
	b, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return keys, fmt.Errorf("can't read twitter keys file %s: %w", path, err)
	}
	if err := json.Unmarshal(b, &keys); err != nil {
		return keys, fmt.Errorf("can't parse twitter keys file %s: %w", path, err)
	}
	return keys, nil
}

// readDiscordKeys loads the Discord bot token and channel ids.
func readDiscordKeys(home, name string) (discordKeys, error) {
	var keys discordKeys
	path := keysPath(home, name)
	b, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return keys, fmt.Errorf("can't read discord keys file %s: %w", path, err)
	}
	if err := json.Unmarshal(b, &keys); err != nil {
		return keys, fmt.Errorf("can't parse discord keys file %s: %w", path, err)
	}
	return keys, nil
}
