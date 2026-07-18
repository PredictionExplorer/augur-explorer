package cosmicgame

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

const maxStatisticsPageLimit = 200

// GlobalStatisticsRecord is the exact-value projection used by API v2.
// Legacy CGStatistics remains unchanged for the frozen v1 dashboard.
type GlobalStatisticsRecord struct {
	TotalBids                          uint64
	CurrentRoundBids                   uint64
	CompletedRounds                    uint64
	TotalPrizeAwards                   uint64
	PrizeRegistryRows                  uint64
	UniqueBidders                      uint64
	UniqueWinners                      uint64
	UniqueDonors                       int64
	UniqueCSTStakers                   uint64
	UniqueRandomWalkStakers            uint64
	UniqueDualStakers                  uint64
	TotalPrizesPaidWei                 string
	TotalEthDonatedWei                 string
	VoluntaryDonationCount             uint64
	VoluntaryDonationsTotalWei         string
	CosmicGameDonationCount            uint64
	CosmicGameDonationsTotalWei        string
	DirectDonationCount                int64
	DirectDonationsTotalWei            string
	CharityWithdrawalCount             uint64
	CharityWithdrawalsTotalWei         string
	RandomWalkTokensUsedInBids         uint64
	DonatedNFTCount                    uint64
	CosmicSignatureMints               uint64
	NamedTokens                        int64
	RaffleEthDepositsTotalWei          string
	RaffleEthWithdrawnTotalWei         string
	ChronoWarriorEthDepositsTotalWei   string
	CSTGivenInPrizesTotalWei           string
	WinnersWithPendingRaffleWithdrawal int64
	CSTConsumedTotalWei                string
	CSTBidCount                        int64
	MarketingRewardsTotalWei           string
	MarketingRewardCount               int64
	DonatedTokenDistribution           []cgmodel.CGDonatedTokenDistrRec
	CSTStaking                         cgmodel.CGStakeStatsCST
	RandomWalkStaking                  cgmodel.CGStakeStatsRWalk
}

// CosmicGameGlobalStatistics returns exact global statistics for API v2.
func (r *Repo) CosmicGameGlobalStatistics(ctx context.Context) (GlobalStatisticsRecord, error) {
	legacy, err := r.CosmicGameStatistics(ctx)
	if err != nil {
		return GlobalStatisticsRecord{}, err
	}
	record := GlobalStatisticsRecord{
		TotalBids:                          legacy.TotalBids,
		CurrentRoundBids:                   legacy.CurNumBids,
		CompletedRounds:                    legacy.TotalPrizes,
		TotalPrizeAwards:                   legacy.TotalPrizeAwards,
		PrizeRegistryRows:                  legacy.CgPrizeRowCount,
		UniqueBidders:                      legacy.NumUniqueBidders,
		UniqueWinners:                      legacy.NumUniqueWinners,
		UniqueDonors:                       legacy.NumUniqueDonors,
		UniqueCSTStakers:                   legacy.NumUniqueStakersCST,
		UniqueRandomWalkStakers:            legacy.NumUniqueStakersRWalk,
		UniqueDualStakers:                  legacy.NumUniqueStakersBoth,
		TotalPrizesPaidWei:                 decimalOrZero(legacy.TotalPrizesPaidAmountWei),
		TotalEthDonatedWei:                 decimalOrZero(legacy.TotalEthDonatedAmount),
		VoluntaryDonationCount:             legacy.NumVoluntaryDonations,
		CosmicGameDonationCount:            legacy.NumCosmicGameDonations,
		DirectDonationCount:                legacy.NumDirectDonations,
		CharityWithdrawalCount:             legacy.NumWithdrawals,
		RandomWalkTokensUsedInBids:         legacy.NumRwalkTokensUsed,
		DonatedNFTCount:                    legacy.NumDonatedNFTs,
		CosmicSignatureMints:               legacy.NumCSTokenMints,
		NamedTokens:                        legacy.TotalNamedTokens,
		WinnersWithPendingRaffleWithdrawal: legacy.NumWinnersWithPendingRaffleWithdrawal,
		CSTConsumedTotalWei:                decimalOrZero(legacy.TotalCSTConsumed),
		CSTBidCount:                        legacy.NumBidsCST,
		MarketingRewardsTotalWei:           decimalOrZero(legacy.TotalMktRewards),
		MarketingRewardCount:               legacy.NumMktRewards,
		DonatedTokenDistribution:           legacy.DonatedTokenDistribution,
		CSTStaking:                         legacy.StakeStatisticsCST,
		RandomWalkStaking:                  legacy.StakeStatisticsRWalk,
		VoluntaryDonationsTotalWei:         "0",
		CosmicGameDonationsTotalWei:        "0",
		DirectDonationsTotalWei:            "0",
		CharityWithdrawalsTotalWei:         "0",
		RaffleEthDepositsTotalWei:          "0",
		RaffleEthWithdrawnTotalWei:         "0",
		ChronoWarriorEthDepositsTotalWei:   "0",
		CSTGivenInPrizesTotalWei:           "0",
	}
	const query = `SELECT
			vol_donations_total::TEXT,
			cg_donations_total::TEXT,
			direct_donations::TEXT,
			sum_withdrawals::TEXT,
			total_raffle_eth_deposits::TEXT,
			total_raffle_eth_withdrawn::TEXT,
			total_chrono_warrior_eth_deposits::TEXT,
			total_cst_given_in_prizes::TEXT
		FROM cg_glob_stats
		LIMIT 1`
	err = r.q(ctx).QueryRow(ctx, query).Scan(
		&record.VoluntaryDonationsTotalWei,
		&record.CosmicGameDonationsTotalWei,
		&record.DirectDonationsTotalWei,
		&record.CharityWithdrawalsTotalWei,
		&record.RaffleEthDepositsTotalWei,
		&record.RaffleEthWithdrawnTotalWei,
		&record.ChronoWarriorEthDepositsTotalWei,
		&record.CSTGivenInPrizesTotalWei,
	)
	if err != nil {
		return GlobalStatisticsRecord{}, store.WrapError("cosmic game global statistics: exact totals", err)
	}
	return record, nil
}

func decimalOrZero(value string) string {
	if value == "" {
		return "0"
	}
	return value
}

// ROILeaderboardSort is the public v2 ROI ordering.
type ROILeaderboardSort string

// The six ROI leaderboard orderings.
const (
	ROILeaderboardNetProfit ROILeaderboardSort = "netProfit"
	ROILeaderboardROI       ROILeaderboardSort = "roi"
	ROILeaderboardWinRate   ROILeaderboardSort = "winRate"
	ROILeaderboardSpent     ROILeaderboardSort = "spent"
	ROILeaderboardNFTs      ROILeaderboardSort = "nfts"
	ROILeaderboardBids      ROILeaderboardSort = "bids"
)

// ROILeaderboardRecord is the exact-value v2 leaderboard projection.
type ROILeaderboardRecord struct {
	BidderAid          int64
	BidderAddr         string
	NumBids            int64
	RoundsParticipated int64
	RoundsWon          int64
	WinRateRatio       string
	TotalEthSpentWei   string
	TotalCSTSpentWei   string
	EthWonWei          string
	PrizesCount        int64
	CSTPrizesCount     int64
	NFTPrizesCount     int64
	NetProfitWei       string
	ROIRatio           string
}

// ROILeaderboardPageCursor identifies the final row of a v2 ROI page.
type ROILeaderboardPageCursor struct {
	Sort      ROILeaderboardSort
	MinBids   int
	SortValue string
	Secondary int64
	BidderAid int64
}

const roiLeaderboardV2Base = `WITH rounds_part AS (
			SELECT bidder_aid, COUNT(DISTINCT round_num) AS rounds_participated
			FROM cg_bid GROUP BY bidder_aid
		), rounds_won AS (
			SELECT aid, COUNT(DISTINCT round_num) AS rounds_won FROM (
				SELECT winner_aid AS aid, round_num FROM cg_prize_claim
				UNION SELECT winner_aid, round_num FROM cg_raffle_eth_prize
				UNION SELECT winner_aid, round_num FROM cg_raffle_nft_prize
				UNION SELECT winner_aid, round_num FROM cg_endurance_prize
				UNION SELECT winner_aid, round_num FROM cg_lastcst_prize
				UNION SELECT winner_aid, round_num FROM cg_chrono_warrior_prize
			) u GROUP BY aid
		), leaderboard AS (
			SELECT
				b.bidder_aid,
				a.addr AS bidder_addr,
				b.num_bids,
				COALESCE(rp.rounds_participated,0)::BIGINT AS rounds_participated,
				COALESCE(rw.rounds_won,0)::BIGINT AS rounds_won,
				CASE WHEN COALESCE(rp.rounds_participated,0) > 0
					THEN COALESCE(rw.rounds_won,0)::NUMERIC / rp.rounds_participated
					ELSE 0 END AS win_rate_ratio,
				b.total_eth_spent,
				b.total_cst_spent,
				COALESCE(w.prizes_sum,0) AS eth_won,
				COALESCE(w.prizes_count,0)::BIGINT AS prizes_count,
				COALESCE(w.erc20_count,0)::BIGINT AS cst_prizes_count,
				COALESCE(w.erc721_count,0)::BIGINT AS nft_prizes_count,
				(COALESCE(w.prizes_sum,0) - b.total_eth_spent) AS net_profit_wei,
				CASE WHEN b.total_eth_spent > 0
					THEN (COALESCE(w.prizes_sum,0) - b.total_eth_spent)::NUMERIC / b.total_eth_spent
					ELSE 0 END AS roi_ratio
			FROM cg_bidder b
				LEFT JOIN address a ON b.bidder_aid=a.address_id
				LEFT JOIN cg_winner w ON b.bidder_aid=w.winner_aid
				LEFT JOIN rounds_part rp ON b.bidder_aid=rp.bidder_aid
				LEFT JOIN rounds_won rw ON b.bidder_aid=rw.aid
			WHERE b.num_bids >= $1
		)
		SELECT
			bidder_aid,
			bidder_addr,
			num_bids,
			rounds_participated,
			rounds_won,
			win_rate_ratio::TEXT,
			total_eth_spent::TEXT,
			total_cst_spent::TEXT,
			eth_won::TEXT,
			prizes_count,
			cst_prizes_count,
			nft_prizes_count,
			net_profit_wei::TEXT,
			roi_ratio::TEXT
		FROM leaderboard`

func roiLeaderboardV2Order(sortBy ROILeaderboardSort) (string, bool) {
	switch sortBy {
	case ROILeaderboardNetProfit:
		return "leaderboard.net_profit_wei DESC, leaderboard.bidder_aid ASC", true
	case ROILeaderboardROI:
		return "leaderboard.roi_ratio DESC, leaderboard.bidder_aid ASC", true
	case ROILeaderboardWinRate:
		return "leaderboard.win_rate_ratio DESC, leaderboard.rounds_participated DESC, leaderboard.bidder_aid ASC", true
	case ROILeaderboardSpent:
		return "leaderboard.total_eth_spent DESC, leaderboard.bidder_aid ASC", true
	case ROILeaderboardNFTs:
		return "leaderboard.nft_prizes_count DESC, leaderboard.bidder_aid ASC", true
	case ROILeaderboardBids:
		return "leaderboard.num_bids DESC, leaderboard.bidder_aid ASC", true
	default:
		return "", false
	}
}

func roiLeaderboardV2Column(sortBy ROILeaderboardSort) string {
	switch sortBy {
	case ROILeaderboardNetProfit:
		return "leaderboard.net_profit_wei"
	case ROILeaderboardROI:
		return "leaderboard.roi_ratio"
	case ROILeaderboardWinRate:
		return "leaderboard.win_rate_ratio"
	case ROILeaderboardSpent:
		return "leaderboard.total_eth_spent"
	case ROILeaderboardNFTs:
		return "leaderboard.nft_prizes_count"
	case ROILeaderboardBids:
		return "leaderboard.num_bids"
	default:
		return ""
	}
}

func scanROILeaderboardRecord(rows pgx.Rows, rec *ROILeaderboardRecord) error {
	return rows.Scan(
		&rec.BidderAid,
		&rec.BidderAddr,
		&rec.NumBids,
		&rec.RoundsParticipated,
		&rec.RoundsWon,
		&rec.WinRateRatio,
		&rec.TotalEthSpentWei,
		&rec.TotalCSTSpentWei,
		&rec.EthWonWei,
		&rec.PrizesCount,
		&rec.CSTPrizesCount,
		&rec.NFTPrizesCount,
		&rec.NetProfitWei,
		&rec.ROIRatio,
	)
}

// ROILeaderboardPage returns at most limit rows after the supplied
// sort-scoped keyset cursor.
func (r *Repo) ROILeaderboardPage(
	ctx context.Context,
	minBids int,
	sortBy ROILeaderboardSort,
	after *ROILeaderboardPageCursor,
	limit int,
) (records []ROILeaderboardRecord, hasMore bool, err error) {
	const op = "roi leaderboard page"
	order, ok := roiLeaderboardV2Order(sortBy)
	if !ok {
		return nil, false, fmt.Errorf("%s: invalid sort", op)
	}
	if minBids < 0 {
		return nil, false, fmt.Errorf("%s: min bids must be non-negative", op)
	}
	if limit <= 0 || limit > maxStatisticsPageLimit {
		return nil, false, fmt.Errorf("%s: invalid limit", op)
	}

	query := roiLeaderboardV2Base
	args := []any{minBids, limit + 1}
	if after == nil {
		query += " ORDER BY " + order + " LIMIT $2"
	} else {
		if after.Sort != sortBy || after.MinBids != minBids ||
			after.SortValue == "" || after.BidderAid < 1 || after.Secondary < 0 {
			return nil, false, fmt.Errorf("%s: invalid cursor", op)
		}
		column := roiLeaderboardV2Column(sortBy)
		if sortBy == ROILeaderboardWinRate {
			query += " WHERE (" + column + " < $2::NUMERIC" +
				" OR (" + column + " = $2::NUMERIC AND rounds_participated < $3)" +
				" OR (" + column + " = $2::NUMERIC AND rounds_participated = $3 AND bidder_aid > $4))" +
				" ORDER BY " + order + " LIMIT $5"
			args = []any{minBids, after.SortValue, after.Secondary, after.BidderAid, limit + 1}
		} else {
			query += " WHERE (" + column + " < $2::NUMERIC" +
				" OR (" + column + " = $2::NUMERIC AND bidder_aid > $3))" +
				" ORDER BY " + order + " LIMIT $4"
			args = []any{minBids, after.SortValue, after.BidderAid, limit + 1}
		}
	}
	records, err = queryList(ctx, r, op, limit+1, query, scanROILeaderboardRecord, args...)
	if err != nil {
		return nil, false, err
	}
	if len(records) > limit {
		records = records[:limit]
		hasMore = true
	}
	return records, hasMore, nil
}

// ROILeaderboardSortValue returns the exact cursor key for record.
func ROILeaderboardSortValue(record ROILeaderboardRecord, sortBy ROILeaderboardSort) string {
	switch sortBy {
	case ROILeaderboardNetProfit:
		return record.NetProfitWei
	case ROILeaderboardROI:
		return record.ROIRatio
	case ROILeaderboardWinRate:
		return record.WinRateRatio
	case ROILeaderboardSpent:
		return record.TotalEthSpentWei
	case ROILeaderboardNFTs:
		return fmt.Sprintf("%d", record.NFTPrizesCount)
	case ROILeaderboardBids:
		return fmt.Sprintf("%d", record.NumBids)
	default:
		return ""
	}
}

// ClaimSummaryCursor identifies the final summary in a newest-round-first page.
type ClaimSummaryCursor struct {
	RoundNum   int64
	EventLogID int64
}

// ClaimSummaryRecord is the exact-value v2 claims summary.
type ClaimSummaryRecord struct {
	RoundNum              int64
	EventLogID            int64
	ClaimWindowTimeout    int64
	AwardedTimestamp      int64
	EthAwarded            int64
	EthUnclaimed          int64
	UnclaimedEthAmountWei string
	NFTAwarded            int64
	NFTUnclaimed          int64
	ERC20Awarded          int64
	ERC20Unclaimed        int64
	TotalAwarded          int64
	TotalUnclaimed        int64
	AvgClaimLatencySecs   int64
}

const claimSummaryV2Select = `SELECT
			pc.round_num,
			pc.evtlog_id,
			pc.timeout,
			EXTRACT(EPOCH FROM pc.time_stamp)::BIGINT,
			COALESCE(eth.awarded,0)::BIGINT,
			COALESCE(eth.unclaimed,0)::BIGINT,
			COALESCE(eth.unclaimed_amt,0)::TEXT,
			COALESCE(nft.awarded,0)::BIGINT,
			COALESCE(nft.unclaimed,0)::BIGINT,
			COALESCE(erc.awarded,0)::BIGINT,
			COALESCE(erc.unclaimed,0)::BIGINT,
			COALESCE(cp.avg_secs,0)::BIGINT
		FROM cg_prize_claim pc
		LEFT JOIN (
			SELECT round_num, COUNT(*) awarded,
				COUNT(*) FILTER (WHERE NOT claimed) unclaimed,
				SUM(amount) FILTER (WHERE NOT claimed) unclaimed_amt
			FROM cg_prize_deposit GROUP BY round_num
		) eth ON eth.round_num=pc.round_num
		LEFT JOIN (
			SELECT d.round_num, COUNT(*) awarded,
				COUNT(*) FILTER (WHERE c.round_num IS NULL) unclaimed
			FROM cg_nft_donation d
				LEFT JOIN cg_donated_nft_claimed c
					ON c.round_num=d.round_num AND c.idx=d.idx
			GROUP BY d.round_num
		) nft ON nft.round_num=pc.round_num
		LEFT JOIN (
			SELECT round_num, COUNT(*) awarded,
				COUNT(*) FILTER (WHERE NOT claimed) unclaimed
			FROM cg_erc20_donation_stats GROUP BY round_num
		) erc ON erc.round_num=pc.round_num
		LEFT JOIN (
			SELECT rn round_num, AVG(secs) avg_secs FROM (
				SELECT w.round_num rn,
					EXTRACT(EPOCH FROM (w.time_stamp-pcw.time_stamp)) secs
				FROM cg_prize_withdrawal w
					JOIN cg_prize_claim pcw ON pcw.round_num=w.round_num
				UNION ALL
				SELECT c.round_num,
					EXTRACT(EPOCH FROM (c.time_stamp-pcn.time_stamp))
				FROM cg_donated_nft_claimed c
					JOIN cg_prize_claim pcn ON pcn.round_num=c.round_num
				UNION ALL
				SELECT t.round_num,
					EXTRACT(EPOCH FROM (t.time_stamp-pct.time_stamp))
				FROM cg_donated_tok_claimed t
					JOIN cg_prize_claim pct ON pct.round_num=t.round_num
			) x GROUP BY rn
		) cp ON cp.round_num=pc.round_num`

func scanClaimSummaryRecord(rows pgx.Rows, rec *ClaimSummaryRecord) error {
	if err := rows.Scan(
		&rec.RoundNum,
		&rec.EventLogID,
		&rec.ClaimWindowTimeout,
		&rec.AwardedTimestamp,
		&rec.EthAwarded,
		&rec.EthUnclaimed,
		&rec.UnclaimedEthAmountWei,
		&rec.NFTAwarded,
		&rec.NFTUnclaimed,
		&rec.ERC20Awarded,
		&rec.ERC20Unclaimed,
		&rec.AvgClaimLatencySecs,
	); err != nil {
		return err
	}
	rec.TotalAwarded = rec.EthAwarded + rec.NFTAwarded + rec.ERC20Awarded
	rec.TotalUnclaimed = rec.EthUnclaimed + rec.NFTUnclaimed + rec.ERC20Unclaimed
	return nil
}

// ClaimsSummaryPage returns claimable-asset summaries newest round first.
func (r *Repo) ClaimsSummaryPage(
	ctx context.Context,
	after *ClaimSummaryCursor,
	limit int,
) (records []ClaimSummaryRecord, hasMore bool, err error) {
	const op = "claims summary page"
	if limit <= 0 || limit > maxStatisticsPageLimit {
		return nil, false, fmt.Errorf("%s: invalid limit", op)
	}
	query := claimSummaryV2Select + `
		WHERE (COALESCE(eth.awarded,0)+COALESCE(nft.awarded,0)+COALESCE(erc.awarded,0)) > 0
		ORDER BY pc.round_num DESC, pc.evtlog_id DESC
		LIMIT $1`
	args := []any{limit + 1}
	if after != nil {
		if after.RoundNum < 0 || after.EventLogID < 1 {
			return nil, false, fmt.Errorf("%s: invalid cursor", op)
		}
		query = claimSummaryV2Select + `
			WHERE (COALESCE(eth.awarded,0)+COALESCE(nft.awarded,0)+COALESCE(erc.awarded,0)) > 0
				AND (pc.round_num,pc.evtlog_id) < ($1,$2)
			ORDER BY pc.round_num DESC, pc.evtlog_id DESC
			LIMIT $3`
		args = []any{after.RoundNum, after.EventLogID, limit + 1}
	}
	records, err = queryList(ctx, r, op, limit+1, query, scanClaimSummaryRecord, args...)
	if err != nil {
		return nil, false, err
	}
	if len(records) > limit {
		records = records[:limit]
		hasMore = true
	}
	return records, hasMore, nil
}

// ClaimSummaryByRound returns a completed round's summary, including a
// zero-count summary when the round awarded no claimable assets.
func (r *Repo) ClaimSummaryByRound(ctx context.Context, roundNum int64) (ClaimSummaryRecord, error) {
	if roundNum < 0 {
		return ClaimSummaryRecord{}, errors.New("claim summary by round: round must be non-negative")
	}
	var rec ClaimSummaryRecord
	err := r.q(ctx).QueryRow(ctx, claimSummaryV2Select+" WHERE pc.round_num=$1", roundNum).Scan(
		&rec.RoundNum,
		&rec.EventLogID,
		&rec.ClaimWindowTimeout,
		&rec.AwardedTimestamp,
		&rec.EthAwarded,
		&rec.EthUnclaimed,
		&rec.UnclaimedEthAmountWei,
		&rec.NFTAwarded,
		&rec.NFTUnclaimed,
		&rec.ERC20Awarded,
		&rec.ERC20Unclaimed,
		&rec.AvgClaimLatencySecs,
	)
	if err != nil {
		return ClaimSummaryRecord{}, store.WrapError("claim summary by round", err)
	}
	rec.TotalAwarded = rec.EthAwarded + rec.NFTAwarded + rec.ERC20Awarded
	rec.TotalUnclaimed = rec.EthUnclaimed + rec.NFTUnclaimed + rec.ERC20Unclaimed
	return rec, nil
}

// ClaimAssetType is the stable v2 claim asset taxonomy.
type ClaimAssetType string

// The claimable asset kinds tracked by PrizesWallet.
const (
	ClaimAssetETH    ClaimAssetType = "eth"
	ClaimAssetERC721 ClaimAssetType = "erc721"
	ClaimAssetERC20  ClaimAssetType = "erc20"
)

// ClaimEventCursor identifies the final event in a chronological section page.
type ClaimEventCursor struct {
	EventLogID int64
}

// ClaimTransactionRecord is one exact claim event.
type ClaimTransactionRecord struct {
	EventLogID       int64
	RoundNum         int64
	AssetType        ClaimAssetType
	RecipientAddr    string
	BeneficiaryAddr  string
	EthAmountWei     *string
	TokenAddr        *string
	TokenID          *int64
	AmountBaseUnits  *string
	ClaimedAfterSecs int64
	ClaimedTimestamp int64
	TxHash           string
}

func claimTransactionsPageSQL(after bool) string {
	filter := "WHERE x.round_num=$1"
	limitPlaceholder := "$2"
	if after {
		filter += " AND x.evtlog_id>$2"
		limitPlaceholder = "$3"
	}
	return fmt.Sprintf(`SELECT
			segment,event_log_id,round_num,asset_type,recipient_addr,beneficiary_addr,
			eth_amount_wei,token_addr,token_id,amount_base_units,
			claimed_after_secs,claimed_at,tx_hash
		FROM (
			(SELECT 0 AS segment,w.evtlog_id AS event_log_id,w.round_num,'eth'::TEXT AS asset_type,
				win.addr AS recipient_addr,ben.addr AS beneficiary_addr,
				w.amount::TEXT AS eth_amount_wei,NULL::TEXT AS token_addr,
				NULL::BIGINT AS token_id,NULL::TEXT AS amount_base_units,
				EXTRACT(EPOCH FROM (w.time_stamp-pc.time_stamp))::BIGINT AS claimed_after_secs,
				EXTRACT(EPOCH FROM w.time_stamp)::BIGINT AS claimed_at,t.tx_hash
			FROM cg_prize_withdrawal w
				JOIN cg_prize_claim pc ON pc.round_num=w.round_num
				JOIN address ben ON ben.address_id=w.beneficiary_aid
				JOIN address win ON win.address_id=w.winner_aid
				LEFT JOIN transaction t ON t.id=w.tx_id
			%s
			ORDER BY w.evtlog_id
			LIMIT %s)
			UNION ALL
			(SELECT 1,dc.evtlog_id,dc.round_num,'erc721'::TEXT,w.addr,w.addr,
				NULL::TEXT,ta.addr,dc.token_id::BIGINT,NULL::TEXT,
				EXTRACT(EPOCH FROM (dc.time_stamp-pc.time_stamp))::BIGINT,
				EXTRACT(EPOCH FROM dc.time_stamp)::BIGINT,t.tx_hash
			FROM cg_donated_nft_claimed dc
				JOIN cg_prize_claim pc ON pc.round_num=dc.round_num
				JOIN address w ON w.address_id=dc.winner_aid
				JOIN address ta ON ta.address_id=dc.token_aid
				LEFT JOIN transaction t ON t.id=dc.tx_id
			%s
			ORDER BY dc.evtlog_id
			LIMIT %s)
			UNION ALL
			(SELECT 2,dc.evtlog_id,dc.round_num,'erc20'::TEXT,w.addr,w.addr,
				NULL::TEXT,ta.addr,NULL::BIGINT,dc.amount::TEXT,
				EXTRACT(EPOCH FROM (dc.time_stamp-pc.time_stamp))::BIGINT,
				EXTRACT(EPOCH FROM dc.time_stamp)::BIGINT,t.tx_hash
			FROM cg_donated_tok_claimed dc
				JOIN cg_prize_claim pc ON pc.round_num=dc.round_num
				JOIN address w ON w.address_id=dc.winner_aid
				JOIN address ta ON ta.address_id=dc.token_aid
				LEFT JOIN transaction t ON t.id=dc.tx_id
			%s
			ORDER BY dc.evtlog_id
			LIMIT %s)
		) claims
		ORDER BY event_log_id
		LIMIT %s`,
		filterForAlias(filter, "x", "w"), limitPlaceholder,
		filterForAlias(filter, "x", "dc"), limitPlaceholder,
		filterForAlias(filter, "x", "dc"), limitPlaceholder,
		limitPlaceholder,
	)
}

func filterForAlias(filter, from, to string) string {
	return strings.ReplaceAll(filter, from+".", to+".")
}

func scanClaimTransactionRecord(rows pgx.Rows, rec *ClaimTransactionRecord) error {
	var (
		segment                               int
		assetType                             string
		ethAmount, tokenAddr, amountBaseUnits sql.NullString
		tokenID                               sql.NullInt64
		txHash                                sql.NullString
	)
	if err := rows.Scan(
		&segment,
		&rec.EventLogID,
		&rec.RoundNum,
		&assetType,
		&rec.RecipientAddr,
		&rec.BeneficiaryAddr,
		&ethAmount,
		&tokenAddr,
		&tokenID,
		&amountBaseUnits,
		&rec.ClaimedAfterSecs,
		&rec.ClaimedTimestamp,
		&txHash,
	); err != nil {
		return err
	}
	rec.AssetType = ClaimAssetType(assetType)
	rec.TxHash = txHash.String
	if ethAmount.Valid {
		value := ethAmount.String
		rec.EthAmountWei = &value
	}
	if tokenAddr.Valid {
		value := tokenAddr.String
		rec.TokenAddr = &value
	}
	if tokenID.Valid {
		value := tokenID.Int64
		rec.TokenID = &value
	}
	if amountBaseUnits.Valid {
		value := amountBaseUnits.String
		rec.AmountBaseUnits = &value
	}
	return nil
}

// ClaimTransactionsPage returns a chronological page of all asset claim events.
func (r *Repo) ClaimTransactionsPage(
	ctx context.Context,
	roundNum int64,
	after *ClaimEventCursor,
	limit int,
) (records []ClaimTransactionRecord, hasMore bool, err error) {
	if err := validateClaimEventPage("claim transactions page", roundNum, after, limit); err != nil {
		return nil, false, err
	}
	args := []any{roundNum, limit + 1}
	if after != nil {
		args = []any{roundNum, after.EventLogID, limit + 1}
	}
	records, err = queryList(ctx, r, "claim transactions page", limit+1,
		claimTransactionsPageSQL(after != nil), scanClaimTransactionRecord, args...)
	if err != nil {
		return nil, false, err
	}
	if len(records) > limit {
		records = records[:limit]
		hasMore = true
	}
	return records, hasMore, nil
}

// AttachedTokenRecord is one NFT or ERC-20 attached during a round.
type AttachedTokenRecord struct {
	EventLogID      int64
	RoundNum        int64
	AssetType       ClaimAssetType
	ContributorAddr string
	TokenAddr       string
	TokenID         *int64
	AmountBaseUnits *string
	OccurredAt      int64
	TxHash          string
}

func attachedTokensPageSQL(after bool) string {
	filterNFT := "WHERE d.round_num=$1"
	filterERC := "WHERE e.round_num=$1"
	limitPlaceholder := "$2"
	if after {
		filterNFT += " AND d.evtlog_id>$2"
		filterERC += " AND e.evtlog_id>$2"
		limitPlaceholder = "$3"
	}
	return fmt.Sprintf(`SELECT segment,event_log_id,round_num,asset_type,contributor_addr,
			token_addr,token_id,amount_base_units,occurred_at,tx_hash
		FROM (
			(SELECT 0 AS segment,d.evtlog_id AS event_log_id,d.round_num,'erc721'::TEXT AS asset_type,
				donor.addr AS contributor_addr,token.addr AS token_addr,
				d.token_id::BIGINT AS token_id,NULL::TEXT AS amount_base_units,
				EXTRACT(EPOCH FROM d.time_stamp)::BIGINT AS occurred_at,t.tx_hash
			FROM cg_nft_donation d
				JOIN address donor ON donor.address_id=d.donor_aid
				JOIN address token ON token.address_id=d.token_aid
				LEFT JOIN transaction t ON t.id=d.tx_id
			%s ORDER BY d.evtlog_id LIMIT %s)
			UNION ALL
			(SELECT 1,e.evtlog_id,e.round_num,'erc20'::TEXT,donor.addr,token.addr,
				NULL::BIGINT,e.amount::TEXT,
				EXTRACT(EPOCH FROM e.time_stamp)::BIGINT,t.tx_hash
			FROM cg_erc20_donation e
				JOIN address donor ON donor.address_id=e.donor_aid
				JOIN address token ON token.address_id=e.token_aid
				LEFT JOIN transaction t ON t.id=e.tx_id
			%s ORDER BY e.evtlog_id LIMIT %s)
		) attached
		ORDER BY event_log_id
		LIMIT %s`, filterNFT, limitPlaceholder, filterERC, limitPlaceholder, limitPlaceholder)
}

func scanAttachedTokenRecord(rows pgx.Rows, rec *AttachedTokenRecord) error {
	var (
		segment         int
		assetType       string
		tokenID         sql.NullInt64
		amountBaseUnits sql.NullString
		txHash          sql.NullString
	)
	if err := rows.Scan(
		&segment,
		&rec.EventLogID,
		&rec.RoundNum,
		&assetType,
		&rec.ContributorAddr,
		&rec.TokenAddr,
		&tokenID,
		&amountBaseUnits,
		&rec.OccurredAt,
		&txHash,
	); err != nil {
		return err
	}
	rec.AssetType = ClaimAssetType(assetType)
	rec.TxHash = txHash.String
	if tokenID.Valid {
		value := tokenID.Int64
		rec.TokenID = &value
	}
	if amountBaseUnits.Valid {
		value := amountBaseUnits.String
		rec.AmountBaseUnits = &value
	}
	return nil
}

// AttachedTokensPage returns a chronological page of attached token events.
func (r *Repo) AttachedTokensPage(
	ctx context.Context,
	roundNum int64,
	after *ClaimEventCursor,
	limit int,
) (records []AttachedTokenRecord, hasMore bool, err error) {
	if err := validateClaimEventPage("attached tokens page", roundNum, after, limit); err != nil {
		return nil, false, err
	}
	args := []any{roundNum, limit + 1}
	if after != nil {
		args = []any{roundNum, after.EventLogID, limit + 1}
	}
	records, err = queryList(ctx, r, "attached tokens page", limit+1,
		attachedTokensPageSQL(after != nil), scanAttachedTokenRecord, args...)
	if err != nil {
		return nil, false, err
	}
	if len(records) > limit {
		records = records[:limit]
		hasMore = true
	}
	return records, hasMore, nil
}

// UnclaimedItemCursor identifies the last kind/key pair in an unclaimed page.
type UnclaimedItemCursor struct {
	Segment int
	Key     int64
}

// UnclaimedItemRecord is one currently unclaimed asset.
type UnclaimedItemRecord struct {
	Segment         int
	Key             int64
	RoundNum        int64
	AssetType       ClaimAssetType
	RecipientAddr   string
	EthAmountWei    *string
	TokenAddr       *string
	TokenID         *int64
	AmountBaseUnits *string
}

const unclaimedItemsSelect = `SELECT segment,stable_key,round_num,asset_type,recipient_addr,
			eth_amount_wei,token_addr,token_id,amount_base_units
		FROM (
			SELECT 0 AS segment,d.evtlog_id AS stable_key,d.round_num,'eth'::TEXT AS asset_type,
				a.addr AS recipient_addr,d.amount::TEXT AS eth_amount_wei,
				NULL::TEXT AS token_addr,NULL::BIGINT AS token_id,
				NULL::TEXT AS amount_base_units
			FROM cg_prize_deposit d
				JOIN address a ON a.address_id=d.winner_aid
			WHERE d.round_num=$1 AND NOT d.claimed
			UNION ALL
			SELECT 1,d.evtlog_id,d.round_num,'erc721'::TEXT,w.addr,NULL::TEXT,
				token.addr,d.token_id::BIGINT,NULL::TEXT
			FROM cg_nft_donation d
				LEFT JOIN cg_donated_nft_claimed c
					ON c.round_num=d.round_num AND c.idx=d.idx
				JOIN address token ON token.address_id=d.token_aid
				LEFT JOIN cg_prize_claim pc ON pc.round_num=d.round_num
				LEFT JOIN address w ON w.address_id=pc.winner_aid
			WHERE d.round_num=$1 AND c.round_num IS NULL
			UNION ALL
			SELECT 2,s.token_aid,s.round_num,'erc20'::TEXT,w.addr,NULL::TEXT,
				token.addr,NULL::BIGINT,s.total_amount::TEXT
			FROM cg_erc20_donation_stats s
				JOIN address token ON token.address_id=s.token_aid
				LEFT JOIN cg_prize_claim pc ON pc.round_num=s.round_num
				LEFT JOIN address w ON w.address_id=pc.winner_aid
			WHERE s.round_num=$1 AND NOT s.claimed
		) unclaimed`

func scanUnclaimedItemRecord(rows pgx.Rows, rec *UnclaimedItemRecord) error {
	var ethAmount, tokenAddr, amountBaseUnits sql.NullString
	var tokenID sql.NullInt64
	if err := rows.Scan(
		&rec.Segment,
		&rec.Key,
		&rec.RoundNum,
		(*string)(&rec.AssetType),
		&rec.RecipientAddr,
		&ethAmount,
		&tokenAddr,
		&tokenID,
		&amountBaseUnits,
	); err != nil {
		return err
	}
	if ethAmount.Valid {
		value := ethAmount.String
		rec.EthAmountWei = &value
	}
	if tokenAddr.Valid {
		value := tokenAddr.String
		rec.TokenAddr = &value
	}
	if tokenID.Valid {
		value := tokenID.Int64
		rec.TokenID = &value
	}
	if amountBaseUnits.Valid {
		value := amountBaseUnits.String
		rec.AmountBaseUnits = &value
	}
	return nil
}

// UnclaimedItemsPage returns a stable kind/key page of unclaimed assets.
func (r *Repo) UnclaimedItemsPage(
	ctx context.Context,
	roundNum int64,
	after *UnclaimedItemCursor,
	limit int,
) (records []UnclaimedItemRecord, hasMore bool, err error) {
	const op = "unclaimed items page"
	if roundNum < 0 {
		return nil, false, fmt.Errorf("%s: round must be non-negative", op)
	}
	if limit <= 0 || limit > maxStatisticsPageLimit {
		return nil, false, fmt.Errorf("%s: invalid limit", op)
	}
	query := unclaimedItemsSelect + " ORDER BY segment,stable_key LIMIT $2"
	args := []any{roundNum, limit + 1}
	if after != nil {
		if after.Segment < 0 || after.Segment > 2 || after.Key < 1 {
			return nil, false, fmt.Errorf("%s: invalid cursor", op)
		}
		query = unclaimedItemsSelect +
			" WHERE (segment,stable_key)>($2,$3) ORDER BY segment,stable_key LIMIT $4"
		args = []any{roundNum, after.Segment, after.Key, limit + 1}
	}
	records, err = queryList(ctx, r, op, limit+1, query, scanUnclaimedItemRecord, args...)
	if err != nil {
		return nil, false, err
	}
	if len(records) > limit {
		records = records[:limit]
		hasMore = true
	}
	return records, hasMore, nil
}

func validateClaimEventPage(
	op string,
	roundNum int64,
	after *ClaimEventCursor,
	limit int,
) error {
	if roundNum < 0 {
		return fmt.Errorf("%s: round must be non-negative", op)
	}
	if limit <= 0 || limit > maxStatisticsPageLimit {
		return fmt.Errorf("%s: invalid limit", op)
	}
	if after != nil && after.EventLogID < 1 {
		return fmt.Errorf("%s: invalid cursor", op)
	}
	return nil
}
