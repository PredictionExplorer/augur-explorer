package cosmicgame

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	. "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	p "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
)

func enrichAdminEventsResolvedValues(events []p.CGAdminEvent) {
	if EthClient == nil || len(events) == 0 {
		return
	}
	gameAddr := contractState.Snapshot().Addrs.CosmicGame
	v1, _ := NewCosmicSignatureGame(gameAddr, EthClient)
	v2, _ := NewCosmicSignatureGameV2(gameAddr, EthClient)
	if v1 == nil {
		return
	}
	for i := range events {
		resolved := resolveAdminEventFromContract(v1, v2, &events[i])
		if resolved != "" {
			events[i].ResolvedValue = resolved
		}
	}
}

func resolveAdminEventFromContract(v1 *CosmicSignatureGame, v2 *CosmicSignatureGameV2, rec *p.CGAdminEvent) string {
	if rec == nil || rec.BlockNum < 0 {
		return ""
	}
	copts := &bind.CallOpts{BlockNumber: big.NewInt(rec.BlockNum)}
	switch rec.RecordType {
	case 20:
		return formatPercentFromDivisor(rec.IntegerValue)
	case 21:
		return formatMicrosecondsAsSeconds(rec.IntegerValue)
	case 22:
		if val, err := v1.GetInitialDurationUntilMainPrize(copts); err == nil && val.Sign() > 0 {
			pct := formatPercentFromDivisor(rec.IntegerValue)
			return fmt.Sprintf("%s (%s)", formatDurationSeconds(val.Int64()), pct)
		}
	case 25:
		// V2 stores duration in seconds directly; V1 stores a divisor-like parameter in the same column.
		if rec.IntegerValue > 0 {
			return formatDurationSeconds(rec.IntegerValue)
		}
		if v2 != nil {
			if dur, err := v2.CstDutchAuctionDuration(copts); err == nil && dur.Sign() > 0 {
				return formatDurationSeconds(dur.Int64())
			}
		}
	case 36:
		if dur, _, err := v1.GetEthDutchAuctionDurations(copts); err == nil && dur.Sign() > 0 {
			return formatDurationSeconds(dur.Int64())
		}
	case 37:
		if begin, err := v1.EthDutchAuctionBeginningBidPrice(copts); err == nil && begin.Sign() > 0 && rec.IntegerValue > 0 {
			end := new(big.Int).Div(begin, big.NewInt(rec.IntegerValue))
			return formatEthFromWei(end.Int64())
		}
		if rec.IntegerValue > 0 {
			if next, err := v1.GetNextEthBidPrice(copts); err == nil && next.Sign() > 0 {
				end := new(big.Int).Div(next, big.NewInt(rec.IntegerValue))
				return formatEthFromWei(end.Int64()) + " (from next bid price)"
			}
		}
	case 39:
		if v2 != nil {
			if dur, err := v2.CstDutchAuctionDuration(copts); err == nil && dur.Sign() > 0 && rec.IntegerValue > 0 {
				delta := dur.Int64() / rec.IntegerValue
				if delta > 0 {
					return fmt.Sprintf("%s change per bid", formatDurationSeconds(delta))
				}
			}
		}
	}
	return ""
}

func formatPercentFromDivisor(divisor int64) string {
	if divisor <= 0 {
		return ""
	}
	return fmt.Sprintf("%.4f%%", 100.0/float64(divisor))
}

func formatMicrosecondsAsSeconds(micros int64) string {
	if micros <= 0 {
		return "0 sec"
	}
	whole := micros / 1_000_000
	frac := micros % 1_000_000
	if frac == 0 {
		return fmt.Sprintf("%d sec", whole)
	}
	return fmt.Sprintf("%.6f sec", float64(micros)/1_000_000.0)
}

func formatDurationSeconds(secs int64) string {
	if secs <= 0 {
		return "0 sec"
	}
	if secs < 120 {
		return fmt.Sprintf("%d sec", secs)
	}
	hours := secs / 3600
	mins := (secs % 3600) / 60
	rem := secs % 60
	if hours > 0 {
		if rem == 0 && mins == 0 {
			return fmt.Sprintf("%dh", hours)
		}
		if rem == 0 {
			return fmt.Sprintf("%dh %dm", hours, mins)
		}
		return fmt.Sprintf("%dh %dm %ds", hours, mins, rem)
	}
	if rem == 0 {
		return fmt.Sprintf("%dm", mins)
	}
	return fmt.Sprintf("%dm %ds", mins, rem)
}

func formatEthFromWei(wei int64) string {
	if wei <= 0 {
		return "0 ETH"
	}
	f := new(big.Float).SetInt(big.NewInt(wei))
	f.Quo(f, big.NewFloat(1e18))
	val, _ := f.Float64()
	if val >= 0.0001 {
		return fmt.Sprintf("%.8f ETH", val)
	}
	return fmt.Sprintf("%.12f ETH", val)
}
