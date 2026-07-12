package rwbot

// Pure formatting helpers: notification texts, media URLs and the Discord
// statistics-channel names. Kept free of I/O so every wire-visible string the
// bot produces is pinned by unit tests.

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/primitives"
)

// Notification event types, matching the values produced by the store's
// AllEventsForNotificationSinceEvtlog union query.
const (
	// EventMint is a token mint.
	EventMint = int64(1)
	// EventSellOffer is a new marketplace SELL offer.
	EventSellOffer = int64(2)
	// EventBought is a marketplace purchase.
	EventBought = int64(3)
	// EventFloorPrice is a floor-price change (synthesized, never stored).
	EventFloorPrice = int64(4)
	// EventBuyOffer is a new marketplace BUY offer.
	EventBuyOffer = int64(5)
)

// ethSign is the ETH currency symbol used in channel names and messages.
const ethSign = "Ξ"

// ImageURL returns the generated-image URL of a token.
func ImageURL(base string, tokenID int64) string {
	return fmt.Sprintf("%v/%06d_black.png", base, tokenID)
}

// VideoURL returns the generated-video URL of a token.
func VideoURL(base string, tokenID int64) string {
	return fmt.Sprintf("%v/%06d_black_single.mp4", base, tokenID)
}

// DetailURL returns the token detail page URL.
func DetailURL(base string, tokenID int64) string {
	return fmt.Sprintf("%v/%v", base, tokenID)
}

// NotificationMessage formats the notification text for an event. detailURL
// is appended after a blank line when non-empty (the Twitter form); Discord
// messages pass "" because the link travels in the embed. withdrawalEth is
// only used by mint events.
func NotificationMessage(evtType, tokenID int64, price, withdrawalEth float64, detailURL string) string {
	var suffix string
	if detailURL != "" {
		suffix = "\n\n" + detailURL
	}
	switch evtType {
	case EventMint:
		return fmt.Sprintf(
			"#%v Minted for %.4f%v.\nLast minter would get %.2f%v if there is no other mint for 30 days.%v",
			tokenID, price, ethSign, withdrawalEth, ethSign, suffix,
		)
	case EventSellOffer:
		return fmt.Sprintf("#%v On sale for %.4f%v.%v", tokenID, price, ethSign, suffix)
	case EventBought:
		return fmt.Sprintf("#%v Bought for %.4f%v.%v", tokenID, price, ethSign, suffix)
	case EventFloorPrice:
		return fmt.Sprintf("Floor price changed to %.4f%v.%v", price, ethSign, suffix)
	case EventBuyOffer:
		return fmt.Sprintf("Buy offer for %.4f%v.%v", price, ethSign, suffix)
	}
	return ""
}

// VideoReplyMessage formats the text of the video tweet posted as a reply to
// a mint notification.
func VideoReplyMessage(tokenID int64) string {
	return fmt.Sprintf("Video for token %v", tokenID)
}

// PriceChannelName formats the Discord "current price" statistics channel
// name.
func PriceChannelName(price float64) string {
	return fmt.Sprintf("Cur. price %v : %.4f", "💲", price)
}

// MintsChannelName formats the Discord "number of mints" statistics channel
// name. The count is tokenID+1 at the call sites (token ids start at 0).
func MintsChannelName(numMints int64) string {
	return fmt.Sprintf("Num. mints %v : %d", "🪙 ", numMints)
}

// LastMintChannelName formats the Discord "last mint ago" statistics channel
// name for a mint at lastMint observed at now.
func LastMintChannelName(lastMint, now time.Time) string {
	duration := primitives.DurationToString(primitives.TimeDifference(lastMint, now))
	return fmt.Sprintf("Last mint: %v ago", duration)
}

// RewardChannelName formats the Discord "last reward" statistics channel
// name.
func RewardChannelName(amountEth float64) string {
	return fmt.Sprintf("Last reward: %.2f%v", amountEth, ethSign)
}

// retryAfterRe extracts the integer second component of a Discord rate-limit
// "retry_after" field embedded in an error string.
var retryAfterRe = regexp.MustCompile(`retry_after"?\s*:\s*(\d+)(?:\.\d+)?`)

// ParseRetryAfterSeconds extracts the whole-second delay from a Discord
// rate-limit error text ("... retry_after": 2.35 ..."). Callers add their own
// safety margin. Returns false when the text carries no retry_after field.
func ParseRetryAfterSeconds(errText string) (int64, bool) {
	m := retryAfterRe.FindStringSubmatch(errText)
	if len(m) != 2 {
		return 0, false
	}
	sec, err := strconv.ParseInt(m[1], 10, 64)
	if err != nil {
		return 0, false
	}
	return sec, true
}
