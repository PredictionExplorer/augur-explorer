package rwbot

import (
	"testing"
	"time"
)

func TestMediaAndDetailURLs(t *testing.T) {
	if got := ImageURL("https://x/images", 42); got != "https://x/images/000042_black.png" {
		t.Errorf("ImageURL = %q", got)
	}
	if got := ImageURL("https://x/images", 1234567); got != "https://x/images/1234567_black.png" {
		t.Errorf("ImageURL wide id = %q", got)
	}
	if got := VideoURL("https://x/videos", 3968); got != "https://x/videos/003968_black_single.mp4" {
		t.Errorf("VideoURL = %q", got)
	}
	if got := DetailURL("https://site/detail", 7); got != "https://site/detail/7" {
		t.Errorf("DetailURL = %q", got)
	}
}

func TestNotificationMessageAllTypes(t *testing.T) {
	detail := "https://site/detail/12"
	cases := []struct {
		name    string
		evtType int64
		detail  string
		want    string
	}{
		{
			name:    "mint with url",
			evtType: EventMint,
			detail:  detail,
			want:    "#12 Minted for 1.2346Ξ.\nLast minter would get 3.50Ξ if there is no other mint for 30 days.\n\nhttps://site/detail/12",
		},
		{
			name:    "mint without url",
			evtType: EventMint,
			want:    "#12 Minted for 1.2346Ξ.\nLast minter would get 3.50Ξ if there is no other mint for 30 days.",
		},
		{
			name:    "sell offer",
			evtType: EventSellOffer,
			detail:  detail,
			want:    "#12 On sale for 1.2346Ξ.\n\nhttps://site/detail/12",
		},
		{
			name:    "bought",
			evtType: EventBought,
			detail:  detail,
			want:    "#12 Bought for 1.2346Ξ.\n\nhttps://site/detail/12",
		},
		{
			name:    "floor price",
			evtType: EventFloorPrice,
			detail:  detail,
			want:    "Floor price changed to 1.2346Ξ.\n\nhttps://site/detail/12",
		},
		{
			name:    "buy offer",
			evtType: EventBuyOffer,
			detail:  detail,
			want:    "Buy offer for 1.2346Ξ.\n\nhttps://site/detail/12",
		},
		{
			name:    "unknown type",
			evtType: 99,
			detail:  detail,
			want:    "",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := NotificationMessage(tc.evtType, 12, 1.23456, 3.5, tc.detail)
			if got != tc.want {
				t.Errorf("NotificationMessage = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestVideoReplyMessage(t *testing.T) {
	if got := VideoReplyMessage(42); got != "Video for token 42" {
		t.Errorf("VideoReplyMessage = %q", got)
	}
}

func TestChannelNames(t *testing.T) {
	if got := PriceChannelName(0.51239); got != "Cur. price 💲 : 0.5124" {
		t.Errorf("PriceChannelName = %q", got)
	}
	if got := MintsChannelName(4068); got != "Num. mints 🪙  : 4068" {
		t.Errorf("MintsChannelName = %q", got)
	}
	if got := RewardChannelName(12.345); got != "Last reward: 12.35Ξ" {
		t.Errorf("RewardChannelName = %q", got)
	}
	minted := time.Date(2026, 7, 1, 0, 0, 0, 0, time.UTC)
	now := minted.Add(26 * time.Hour)
	got := LastMintChannelName(minted, now)
	// DurationToString carries a legacy leading space; the channel names in
	// production Discord always looked like this.
	if got != "Last mint:  1 day, 2 hours ago" {
		t.Errorf("LastMintChannelName = %q", got)
	}
}

func TestParseRetryAfterSeconds(t *testing.T) {
	cases := []struct {
		in   string
		want int64
		ok   bool
	}{
		{`{"message": "You are being rate limited.", "retry_after": 2.35, "global": false}`, 2, true},
		{`retry_after": 10.0`, 10, true},
		{`retry_after: 7`, 7, true},
		{`no rate limit info`, 0, false},
		{`retry_after": abc`, 0, false},
		// Overflows int64: the parse fails and the caller gets no delay.
		{`retry_after": 99999999999999999999.5`, 0, false},
	}
	for _, tc := range cases {
		got, ok := ParseRetryAfterSeconds(tc.in)
		if got != tc.want || ok != tc.ok {
			t.Errorf("ParseRetryAfterSeconds(%q) = (%d,%v), want (%d,%v)", tc.in, got, ok, tc.want, tc.ok)
		}
	}
}
