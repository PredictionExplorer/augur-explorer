package common

import (
	"fmt"
	"testing"
	"time"

	rwp "github.com/PredictionExplorer/augur-explorer/internal/primitives/randomwalk"
)

func TestBuildJSRandomwalkWithdrawalChart(t *testing.T) {
	epoch := time.Unix(0, 0)
	beforeEpoch := time.Unix(-1, 0)

	tests := []struct {
		name    string
		entries []rwp.API_WithdrawalChartEntry
		want    string
	}{
		{name: "nil slice", entries: nil, want: "[]"},
		{name: "empty slice", entries: []rwp.API_WithdrawalChartEntry{}, want: "[]"},
		{
			name: "single entry",
			entries: []rwp.API_WithdrawalChartEntry{
				{TimeStamp: 0, WithdrawalAmount: 1.25},
			},
			want: fmt.Sprintf(
				`[{x:new Date(0 * 1000),y:1.250000000000000000,amount: 1.25,date_str: "%v",timestamp:0}]`,
				epoch,
			),
		},
		{
			name: "multiple entries preserve order and delimiter",
			entries: []rwp.API_WithdrawalChartEntry{
				{TimeStamp: -1, WithdrawalAmount: -0.5},
				{TimeStamp: 2, WithdrawalAmount: 1.0 / 3.0},
			},
			want: fmt.Sprintf(
				`[{x:new Date(-1 * 1000),y:-0.500000000000000000,amount: -0.5,date_str: "%v",timestamp:-1},{x:new Date(2 * 1000),y:0.333333333333333315,amount: 0.3333333333333333,date_str: "%v",timestamp:2}]`,
				beforeEpoch,
				time.Unix(2, 0),
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := string(BuildJSRandomwalkWithdrawalChart(&tt.entries))
			if got != tt.want {
				t.Errorf("chart data = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestBuildJSFloorPriceData(t *testing.T) {
	tests := []struct {
		name    string
		entries []rwp.API_FloorPrice
		want    string
	}{
		{name: "nil slice", entries: nil, want: "[]"},
		{name: "empty slice", entries: []rwp.API_FloorPrice{}, want: "[]"},
		{
			name: "single entry",
			entries: []rwp.API_FloorPrice{
				{TimeStamp: 0, Price: 2.5},
			},
			want: fmt.Sprintf(
				`[{x:new Date(0 * 1000),y:2.500000000000000000,price: 2.5,date_str: "%v"}]`,
				time.Unix(0, 0),
			),
		},
		{
			name: "multiple entries preserve order and delimiter",
			entries: []rwp.API_FloorPrice{
				{TimeStamp: -1, Price: -2.5},
				{TimeStamp: 2, Price: 1.0 / 3.0},
			},
			want: fmt.Sprintf(
				`[{x:new Date(-1 * 1000),y:-2.500000000000000000,price: -2.5,date_str: "%v"},{x:new Date(2 * 1000),y:0.333333333333333315,price: 0.3333333333333333,date_str: "%v"}]`,
				time.Unix(-1, 0),
				time.Unix(2, 0),
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := string(BuildJSFloorPriceData(&tt.entries))
			if got != tt.want {
				t.Errorf("floor price data = %q, want %q", got, tt.want)
			}
		})
	}
}
