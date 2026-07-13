package v2

import (
	"errors"
	"fmt"
	"math/big"
	"sort"
	"time"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
)

func mapBidFrequencyBuckets(
	records []cgmodel.CGBidFrequencyBucket,
	to int64,
) ([]BidFrequencyBucket, []cgmodel.CGBidFrequencyBucket, error) {
	mapped := make([]BidFrequencyBucket, 0, len(records))
	normalized := make([]cgmodel.CGBidFrequencyBucket, 0, len(records))
	var previous int64
	havePrevious := false
	terminalSeen := false
	for i := range records {
		record := records[i]
		if terminalSeen {
			return nil, nil, errors.New("frequency bucket follows terminal boundary")
		}
		if record.BucketTs == to {
			terminalSeen = true
			continue
		}
		if record.BucketTs < 0 || record.BucketTs > to {
			return nil, nil, fmt.Errorf("frequency bucket %d is outside the requested window", i)
		}
		if havePrevious && record.BucketTs <= previous {
			return nil, nil, fmt.Errorf("frequency bucket %d is not strictly ordered", i)
		}
		if record.NumBids < 0 || record.UniqueBidders < 0 || record.UniqueBidders > record.NumBids {
			return nil, nil, fmt.Errorf("frequency bucket %d has inconsistent counts", i)
		}
		mapped = append(mapped, BidFrequencyBucket{
			BidCount:          record.NumBids,
			BucketTimestamp:   record.BucketTs,
			UniqueBidderCount: record.UniqueBidders,
		})
		normalized = append(normalized, record)
		previous = record.BucketTs
		havePrevious = true
	}
	return mapped, normalized, nil
}

func mapBidSpikes(records []cgmodel.CGBidSpike, to int64) ([]BidSpike, error) {
	mapped := make([]BidSpike, 0, len(records))
	var previousStart int64
	for i := range records {
		record := records[i]
		if record.Index != i {
			return nil, fmt.Errorf("spike %d has non-sequential index %d", i, record.Index)
		}
		if record.StartTs < 0 || record.StartTs >= to ||
			record.EndTs <= record.StartTs ||
			record.PeakTs < record.StartTs || record.PeakTs >= record.EndTs {
			return nil, fmt.Errorf("spike %d has inconsistent timestamps", i)
		}
		if i > 0 && record.StartTs <= previousStart {
			return nil, fmt.Errorf("spike %d is not strictly ordered", i)
		}
		if record.PeakNumBids < 0 || record.TotalBids < record.PeakNumBids || record.BucketCount < 1 {
			return nil, fmt.Errorf("spike %d has inconsistent counts", i)
		}
		endTimestamp := record.EndTs
		if endTimestamp > to {
			endTimestamp = to
		}
		mapped = append(mapped, BidSpike{
			BucketCount:    record.BucketCount,
			EndTimestamp:   endTimestamp,
			Index:          record.Index,
			PeakBidCount:   record.PeakNumBids,
			PeakTimestamp:  record.PeakTs,
			StartTimestamp: record.StartTs,
			TotalBidCount:  record.TotalBids,
		})
		previousStart = record.StartTs
	}
	return mapped, nil
}

func recentBidSpikeIndex(spikes []BidSpike, now time.Time) *int {
	nowTimestamp := now.Unix()
	cutoff := nowTimestamp - recentBidSpikeWindowSeconds
	var recent *int
	for i := range spikes {
		if spikes[i].StartTimestamp < cutoff || spikes[i].StartTimestamp > nowTimestamp {
			continue
		}
		index := spikes[i].Index
		recent = &index
	}
	return recent
}

func mapBidTypeRatioBuckets(
	records []cgmodel.CGBidTypeRatioBucket,
	to int64,
) ([]BidTypeRatioBucket, error) {
	mapped := make([]BidTypeRatioBucket, 0, len(records))
	var previous int64
	havePrevious := false
	terminalSeen := false
	for i := range records {
		record := records[i]
		if terminalSeen {
			return nil, errors.New("bid-type bucket follows terminal boundary")
		}
		if record.BucketTs == to {
			terminalSeen = true
			continue
		}
		if record.BucketTs < 0 || record.BucketTs > to {
			return nil, fmt.Errorf("bid-type bucket %d is outside the requested window", i)
		}
		if havePrevious && record.BucketTs <= previous {
			return nil, fmt.Errorf("bid-type bucket %d is not strictly ordered", i)
		}
		if !validBidTypeCounts(record) {
			return nil, fmt.Errorf("bid-type bucket %d has inconsistent counts", i)
		}
		ethPercentage, err := percentageString(record.EthBids, record.TotalBids)
		if err != nil {
			return nil, fmt.Errorf("bid-type bucket %d ETH percentage: %w", i, err)
		}
		randomWalkPercentage, err := percentageString(record.RwalkBids, record.TotalBids)
		if err != nil {
			return nil, fmt.Errorf("bid-type bucket %d RandomWalk percentage: %w", i, err)
		}
		cstPercentage, err := percentageString(record.CstBids, record.TotalBids)
		if err != nil {
			return nil, fmt.Errorf("bid-type bucket %d CST percentage: %w", i, err)
		}
		mapped = append(mapped, BidTypeRatioBucket{
			BucketTimestamp:      record.BucketTs,
			CstBidCount:          record.CstBids,
			CstPercentage:        cstPercentage,
			EthBidCount:          record.EthBids,
			EthPercentage:        ethPercentage,
			RandomWalkBidCount:   record.RwalkBids,
			RandomWalkPercentage: randomWalkPercentage,
			TotalBidCount:        record.TotalBids,
		})
		previous = record.BucketTs
		havePrevious = true
	}
	return mapped, nil
}

func validBidTypeCounts(record cgmodel.CGBidTypeRatioBucket) bool {
	if record.EthBids < 0 || record.RwalkBids < 0 || record.CstBids < 0 || record.TotalBids < 0 {
		return false
	}
	if record.EthBids > record.TotalBids {
		return false
	}
	remaining := record.TotalBids - record.EthBids
	if record.RwalkBids > remaining {
		return false
	}
	return record.CstBids == remaining-record.RwalkBids
}

func percentageString(part, total int64) (string, error) {
	if part < 0 || total < 0 || part > total {
		return "", errors.New("invalid percentage counts")
	}
	if total == 0 {
		if part != 0 {
			return "", errors.New("nonzero part with zero total")
		}
		return "0", nil
	}
	value := new(big.Rat).SetFrac(big.NewInt(part), big.NewInt(total))
	value.Mul(value, big.NewRat(100, 1))
	return canonicalDecimal(value.FloatString(2), false)
}

func mapTopBidderActivePeriods(
	bidderRecords []cgmodel.CGTopBidderInfo,
	periodRecords []cgmodel.CGBidderActivePeriod,
	from int64,
	to int64,
	top int,
	minBids int,
) ([]TopBidder, []BidderActivePeriod, error) {
	if len(bidderRecords) > top {
		return nil, nil, errors.New("repository returned more top bidders than requested")
	}
	sortedBidders := append([]cgmodel.CGTopBidderInfo(nil), bidderRecords...)
	sort.Slice(sortedBidders, func(i, j int) bool {
		if sortedBidders[i].NumBids != sortedBidders[j].NumBids {
			return sortedBidders[i].NumBids > sortedBidders[j].NumBids
		}
		return sortedBidders[i].BidderAid < sortedBidders[j].BidderAid
	})
	bidders := make([]TopBidder, 0, len(sortedBidders))
	addresses := make(map[int64]string, len(sortedBidders))
	for i := range sortedBidders {
		record := sortedBidders[i]
		if record.NumBids < 1 {
			return nil, nil, fmt.Errorf("top bidder %d has no bids", i)
		}
		address, err := participantAddress(record.BidderAid, record.BidderAddr)
		if err != nil {
			return nil, nil, fmt.Errorf("top bidder %d: %w", i, err)
		}
		if _, duplicate := addresses[record.BidderAid]; duplicate {
			return nil, nil, fmt.Errorf("top bidder %d is duplicated", i)
		}
		addresses[record.BidderAid] = address
		bidders = append(bidders, TopBidder{
			BidderAddress:    address,
			LifetimeBidCount: record.NumBids,
		})
	}

	sortedPeriods := append([]cgmodel.CGBidderActivePeriod(nil), periodRecords...)
	sort.Slice(sortedPeriods, func(i, j int) bool {
		if sortedPeriods[i].PeriodStart != sortedPeriods[j].PeriodStart {
			return sortedPeriods[i].PeriodStart < sortedPeriods[j].PeriodStart
		}
		if sortedPeriods[i].BidderAid != sortedPeriods[j].BidderAid {
			return sortedPeriods[i].BidderAid < sortedPeriods[j].BidderAid
		}
		return sortedPeriods[i].PeriodEnd < sortedPeriods[j].PeriodEnd
	})
	periods := make([]BidderActivePeriod, 0, len(sortedPeriods))
	for i := range sortedPeriods {
		record := sortedPeriods[i]
		topAddress, exists := addresses[record.BidderAid]
		if !exists {
			return nil, nil, fmt.Errorf("active period %d belongs to a non-top bidder", i)
		}
		address, err := participantAddress(record.BidderAid, record.BidderAddr)
		if err != nil {
			return nil, nil, fmt.Errorf("active period %d: %w", i, err)
		}
		if address != topAddress {
			return nil, nil, fmt.Errorf("active period %d has a mismatched bidder address", i)
		}
		if record.PeriodStart < from || record.PeriodStart >= to ||
			record.PeriodEnd < record.PeriodStart || record.PeriodEnd >= to {
			return nil, nil, fmt.Errorf("active period %d is outside the requested window", i)
		}
		if record.DurationSecs != record.PeriodEnd-record.PeriodStart {
			return nil, nil, fmt.Errorf("active period %d has an inconsistent duration", i)
		}
		if record.NumBids < int64(minBids) {
			return nil, nil, fmt.Errorf("active period %d has fewer bids than requested", i)
		}
		periods = append(periods, BidderActivePeriod{
			BidCount:        record.NumBids,
			BidderAddress:   address,
			DurationSeconds: record.DurationSecs,
			EndTimestamp:    record.PeriodEnd,
			StartTimestamp:  record.PeriodStart,
		})
	}
	return bidders, periods, nil
}
