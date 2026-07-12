package main

// Credential-file loading. Both files live under $HOME/configs and their
// names come from the TWITTER_KEYS_FILE / DISCORD_KEYS_FILE environment
// variables, exactly like the legacy bot.

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/PredictionExplorer/augur-explorer/internal/notify/tweets"
)

// discordKeys is the JSON shape of the Discord credentials file.
type discordKeys struct {
	TokenKey          string
	ChannelId         uint64 // notifications channel
	MainChannelId     uint64 // main chat channel
	MintStatsChanId   uint64
	PriceStatsChanId  uint64
	DateStatsChanId   uint64
	RewardStatsChanId uint64
}

// keysPath resolves the credential file named by envVar under $HOME/configs.
func keysPath(envVar string) string {
	return filepath.Join(os.Getenv("HOME"), "configs", os.Getenv(envVar))
}

// readTwitterKeys loads the Twitter API credentials.
func readTwitterKeys() (tweets.TwitterKeys, error) {
	var keys tweets.TwitterKeys
	path := keysPath("TWITTER_KEYS_FILE")
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
func readDiscordKeys() (discordKeys, error) {
	var keys discordKeys
	path := keysPath("DISCORD_KEYS_FILE")
	b, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return keys, fmt.Errorf("can't read discord keys file %s: %w", path, err)
	}
	if err := json.Unmarshal(b, &keys); err != nil {
		return keys, fmt.Errorf("can't parse discord keys file %s: %w", path, err)
	}
	return keys, nil
}
