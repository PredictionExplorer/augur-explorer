package cosmicgame

import (
	"math"
	"math/big"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

func TestClassifyBidRewardMints(t *testing.T) {
	current := ethcommon.HexToAddress("0x2100000000000000000000000000000000000021")
	previous := ethcommon.HexToAddress("0x2200000000000000000000000000000000000022")
	tests := []struct {
		name         string
		mints        []bidRewardMint
		wantCurrent  string
		wantPrevious string
		wantAddress  string
	}{
		{name: "no dynamic reward", wantCurrent: "0", wantPrevious: "0"},
		{
			name:         "legacy single mint",
			mints:        []bidRewardMint{{to: current, amount: big.NewInt(100)}},
			wantCurrent:  "100",
			wantPrevious: "0",
		},
		{
			name: "V3 90/10 in either order",
			mints: []bidRewardMint{
				{to: current, amount: big.NewInt(10)},
				{to: previous, amount: big.NewInt(90)},
			},
			wantCurrent:  "10",
			wantPrevious: "90",
			wantAddress:  previous.String(),
		},
		{
			name: "same address still has two shares",
			mints: []bidRewardMint{
				{to: current, amount: big.NewInt(90)},
				{to: current, amount: big.NewInt(10)},
			},
			wantCurrent:  "10",
			wantPrevious: "90",
			wantAddress:  current.String(),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			currentAmount, previousAmount, address := classifyBidRewardMints(test.mints)
			if currentAmount != test.wantCurrent ||
				previousAmount != test.wantPrevious ||
				address != test.wantAddress {
				t.Fatalf("split = (%s,%s,%s), want (%s,%s,%s)",
					currentAmount, previousAmount, address,
					test.wantCurrent, test.wantPrevious, test.wantAddress)
			}
		})
	}
}

func TestDecodeMainPrizeClaimedV3(t *testing.T) {
	h := newUnitHandlers(t)
	event := h.gameV3ABI.Events["MainPrizeClaimed"]
	data, err := event.Inputs.NonIndexed().Pack(
		big.NewInt(11),
		big.NewInt(22),
		big.NewInt(3),
		big.NewInt(44),
	)
	if err != nil {
		t.Fatal(err)
	}
	winner := ethcommon.HexToAddress("0x2100000000000000000000000000000000000021")
	log := &types.Log{
		Address: h.c.Game,
		Topics: []ethcommon.Hash{
			event.ID,
			ethcommon.BigToHash(big.NewInt(7)),
			ethcommon.BytesToHash(winner.Bytes()),
			ethcommon.BigToHash(big.NewInt(100)),
		},
		Data: data,
	}
	decoded, err := h.decodeMainPrizeClaimedV3(log, &store.EthereumEventLog{
		EvtID: 1, BlockNum: 2, TimeStamp: 3, TxID: 4,
	})
	if err != nil {
		t.Fatalf("decode V3 claim: %v", err)
	}
	if decoded.RoundNum != 7 ||
		decoded.WinnerAddr != winner.String() ||
		decoded.TokenId != 100 ||
		decoded.NumCSNfts != 3 ||
		decoded.Amount != "11" ||
		decoded.CstAmount != "22" ||
		decoded.Timeout != 44 {
		t.Fatalf("decoded V3 claim = %+v", decoded)
	}

	badData, err := event.Inputs.NonIndexed().Pack(
		big.NewInt(11),
		big.NewInt(22),
		big.NewInt(0),
		big.NewInt(44),
	)
	if err != nil {
		t.Fatal(err)
	}
	log.Data = badData
	if _, err := h.decodeMainPrizeClaimedV3(log, &store.EthereumEventLog{}); err == nil {
		t.Fatal("zero V3 main-prize NFT count was accepted")
	}

	overflow := new(big.Int).Lsh(big.NewInt(1), 80)
	badData, err = event.Inputs.NonIndexed().Pack(
		big.NewInt(11),
		big.NewInt(22),
		big.NewInt(3),
		overflow,
	)
	if err != nil {
		t.Fatal(err)
	}
	log.Data = badData
	if _, err := h.decodeMainPrizeClaimedV3(log, &store.EthereumEventLog{}); err == nil {
		t.Fatal("overflowing V3 timeout was accepted")
	}

	badData, err = event.Inputs.NonIndexed().Pack(
		big.NewInt(11),
		big.NewInt(22),
		big.NewInt(2),
		big.NewInt(44),
	)
	if err != nil {
		t.Fatal(err)
	}
	log.Data = badData
	log.Topics[3] = ethcommon.BigToHash(big.NewInt(math.MaxInt64))
	if _, err := h.decodeMainPrizeClaimedV3(log, &store.EthereumEventLog{}); err == nil {
		t.Fatal("overflowing V3 token range was accepted")
	}
}

func TestDecodeV3AdminEvents(t *testing.T) {
	h := newUnitHandlers(t)
	meta := &store.EthereumEventLog{EvtID: 1, BlockNum: 2, TimeStamp: 3, TxID: 4}
	value := big.NewInt(42)
	tests := []struct {
		name   string
		decode func(*types.Log, *store.EthereumEventLog) (string, error)
	}{
		{
			"RoundLateBidDurationDivisorChanged",
			func(log *types.Log, meta *store.EthereumEventLog) (string, error) {
				event, err := h.decodeRoundLateBidDurationDivisorChanged(log, meta)
				if err != nil {
					return "", err
				}
				return event.NewValue, nil
			},
		},
		{
			"RoundLateBidPricePremiumAmountBaseMultiplierChanged",
			func(log *types.Log, meta *store.EthereumEventLog) (string, error) {
				event, err := h.decodeRoundLateBidPremiumBaseMultiplierChanged(log, meta)
				if err != nil {
					return "", err
				}
				return event.NewValue, nil
			},
		},
		{
			"RoundLateBidPricePremiumAmountExponentChanged",
			func(log *types.Log, meta *store.EthereumEventLog) (string, error) {
				event, err := h.decodeRoundLateBidPremiumExponentChanged(log, meta)
				if err != nil {
					return "", err
				}
				return event.NewValue, nil
			},
		},
		{
			"LastBidderBidCstRewardAmountPercentageChanged",
			func(log *types.Log, meta *store.EthereumEventLog) (string, error) {
				event, err := h.decodeLastBidderRewardPercentageChanged(log, meta)
				if err != nil {
					return "", err
				}
				return event.NewValue, nil
			},
		},
		{
			"MainPrizeNumCosmicSignatureNftsChanged",
			func(log *types.Log, meta *store.EthereumEventLog) (string, error) {
				event, err := h.decodeMainPrizeNumNftsChanged(log, meta)
				if err != nil {
					return "", err
				}
				return event.NewValue, nil
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			event := h.gameV3ABI.Events[test.name]
			data, err := event.Inputs.NonIndexed().Pack(value)
			if err != nil {
				t.Fatal(err)
			}
			got, err := test.decode(&types.Log{Address: h.c.Game, Data: data}, meta)
			if err != nil || got != "42" {
				t.Fatalf("decode = %q, %v", got, err)
			}
		})
	}
}
