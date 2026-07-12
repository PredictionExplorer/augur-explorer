package v2

import (
	"encoding/base64"
	"errors"
	"strings"
	"testing"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func TestOpaqueCursorCodec(t *testing.T) {
	t.Parallel()

	type payload struct {
		Version int   `json:"v"`
		Value   int64 `json:"x"`
	}
	invalid := errors.New("test cursor")
	encoded, err := encodeOpaqueCursor(payload{Version: 1, Value: 42}, invalid, "test cursor")
	if err != nil {
		t.Fatalf("encodeOpaqueCursor: %v", err)
	}
	got, err := decodeOpaqueCursor[payload](encoded, invalid)
	if err != nil {
		t.Fatalf("decodeOpaqueCursor: %v", err)
	}
	if got != (payload{Version: 1, Value: 42}) {
		t.Fatalf("decoded payload = %+v", got)
	}

	tests := map[string]struct {
		input string
		want  string
	}{
		"empty": {
			want: "test cursor: invalid length",
		},
		"oversized": {
			input: strings.Repeat("a", maxCursorLength+1),
			want:  "test cursor: invalid length",
		},
		"invalid base64": {
			input: "!",
			want:  "test cursor: invalid base64url",
		},
		"unknown field": {
			input: base64.RawURLEncoding.EncodeToString([]byte(`{"v":1,"x":42,"extra":true}`)),
			want:  "test cursor: invalid payload",
		},
		"trailing JSON": {
			input: base64.RawURLEncoding.EncodeToString([]byte(`{"v":1,"x":42}{}`)),
			want:  "test cursor: trailing payload",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			_, err := decodeOpaqueCursor[payload](test.input, invalid)
			if !errors.Is(err, invalid) || err.Error() != test.want {
				t.Fatalf("decode error = %v, want %q wrapping sentinel", err, test.want)
			}
		})
	}
}

func TestOpaqueCursorCodecReportsMarshalFailure(t *testing.T) {
	t.Parallel()

	invalid := errors.New("test cursor")
	_, err := encodeOpaqueCursor(struct {
		Value chan int `json:"value"`
	}{Value: make(chan int)}, invalid, "test cursor")
	if err == nil || errors.Is(err, invalid) {
		t.Fatalf("marshal error = %v, want non-validation failure", err)
	}

	_, err = encodeOpaqueCursor(strings.Repeat("a", maxCursorLength), invalid, "test cursor")
	if !errors.Is(err, invalid) || err.Error() != "test cursor: encoded value is too long" {
		t.Fatalf("oversized encode error = %v", err)
	}
}

func TestCursorEncodersRejectInvalidFieldsExactly(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		invalid  error
		encode   func() error
		wantText string
	}{
		{
			name:    "participant kind",
			invalid: errInvalidParticipantCursor,
			encode: func() error {
				_, err := encodeParticipantCursor(participantCursor{
					Version:   participantCursorVersion,
					Kind:      cgstore.ParticipantKind("unknown"),
					SortValue: "1",
					AddressID: 1,
				})
				return err
			},
			wantText: "invalid participant cursor: invalid fields",
		},
		{
			name:    "ROI secondary key",
			invalid: errInvalidROILeaderboardCursor,
			encode: func() error {
				_, err := encodeROILeaderboardCursor(roiLeaderboardCursor{
					Version:   roiLeaderboardCursorVersion,
					Sort:      Roi,
					MinBids:   5,
					SortValue: "1",
					Secondary: 1,
					BidderAid: 1,
				})
				return err
			},
			wantText: "invalid ROI leaderboard cursor: invalid fields",
		},
		{
			name:    "claim summary identity",
			invalid: errInvalidClaimSummaryCursor,
			encode: func() error {
				_, err := encodeClaimSummaryCursor(claimSummaryCursor{
					Version:    claimSummaryCursorVersion,
					Round:      -1,
					EventLogID: 1,
				})
				return err
			},
			wantText: "invalid claim summary cursor: invalid fields",
		},
		{
			name:    "claim detail section",
			invalid: errInvalidClaimDetailCursor,
			encode: func() error {
				_, err := encodeClaimDetailCursor(claimDetailCursor{
					Version:    claimDetailCursorVersion,
					Round:      1,
					Section:    claimDetailSection("unknown"),
					EventLogID: 1,
				})
				return err
			},
			wantText: "invalid claim detail cursor: invalid fields",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			err := test.encode()
			if !errors.Is(err, test.invalid) || err.Error() != test.wantText {
				t.Fatalf("encode error = %v, want %q", err, test.wantText)
			}
		})
	}
}
