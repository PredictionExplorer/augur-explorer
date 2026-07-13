package cosmicgame

import (
	"context"
	"math"
	"strconv"

	"github.com/jackc/pgx/v5"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// SystemModeChanges returns round boundaries (first bid, prize claim) so
// "View Events" can show which admin/configuration events apply to each
// round. When offset is -1 a synthetic "Deployment" row is appended so
// pre-round-0 state is visible even when there are no bids yet.
func (r *Repo) SystemModeChanges(ctx context.Context, offset, limit int) ([]cgmodel.CGSystemModeRec, error) {
	const op = "system mode change event list"
	if limit == 0 {
		limit = 1000000
	}
	addDeploymentEvents := false
	if offset == -1 {
		addDeploymentEvents = true
		offset = 0
	}
	query := `(
			SELECT
				s.evtlog_id,
				s.block_num,
				EXTRACT(EPOCH FROM s.time_stamp)::BIGINT ts,
				s.time_stamp date_time,
				s.round_num,
				0 AS rec_type
			FROM cg_first_bid s
		) UNION ALL (
			SELECT
				p.evtlog_id,
				p.block_num,
				EXTRACT(EPOCH FROM p.time_stamp)::BIGINT ts,
				p.time_stamp date_time,
				-1 AS round_num,
				1 AS rec_type
			FROM cg_prize_claim p
		)
		ORDER BY evtlog_id DESC
		OFFSET $1 LIMIT $2`

	rows, err := r.pool().Query(ctx, query, offset, limit)
	if err != nil {
		return nil, store.WrapError(op, err)
	}
	defer rows.Close()
	records := make([]cgmodel.CGSystemModeRec, 0, 256)
	// A prize-claim row (rec_type 1) closes the round opened by the first-bid
	// row (rec_type 0) seen just before it in evtlog order; only the closed
	// spans are returned.
	var evtlogHi int64 = math.MaxInt64
	var roundNum int64
	for rows.Next() {
		var rec cgmodel.CGSystemModeRec
		var recType int64
		err = rows.Scan(
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TimeStamp,
			store.TimeText(&rec.DateTime),
			&rec.RoundNum,
			&recType,
		)
		if err != nil {
			return nil, store.WrapError(op, err)
		}
		if recType == 1 {
			rec.NextEvtLogId = evtlogHi
			rec.RoundNum = roundNum
			records = append(records, rec)
		} else {
			evtlogHi = rec.EvtLogId
			roundNum = rec.RoundNum
		}
	}
	if err := rows.Err(); err != nil {
		return nil, store.WrapError(op, err)
	}
	if addDeploymentEvents {
		// Always add the "Deployment" row when offset=-1 so pre-round-0
		// config is visible even with no bids.
		records = append(records, cgmodel.CGSystemModeRec{
			EvtLogId:     -1,
			BlockNum:     -1,
			RoundNum:     0,
			NextEvtLogId: evtlogHi,
		})
	}
	return records, nil
}

// adminEventBranch renders one arm of the AdminEventsInRange UNION: table
// determines the source, recordType tags the rows, and the four value
// expressions land in addr_value / int_value / float_value / string_value.
// All inputs are compile-time literals from the table below — nothing
// user-supplied is interpolated.
type adminEventBranch struct {
	recordType int
	table      string
	addrValue  string // SQL expression; "" means no address join
	addrJoinFK string // FK column joined to address when addrValue is set
	intValue   string
	floatValue string
	stringVal  string
}

func (b adminEventBranch) sql() string {
	addrExpr := "'' AS addr_value"
	addrJoin := ""
	if b.addrValue != "" {
		addrExpr = b.addrValue + " AS addr_value"
		if b.addrJoinFK != "" {
			addrJoin = " LEFT JOIN address a ON a.address_id = r." + b.addrJoinFK
		}
	}
	intExpr := "0"
	if b.intValue != "" {
		intExpr = b.intValue
	}
	floatExpr := "0"
	if b.floatValue != "" {
		floatExpr = b.floatValue
	}
	strExpr := "''"
	if b.stringVal != "" {
		strExpr = b.stringVal
	}
	return `(
		SELECT
			` + strconv.Itoa(b.recordType) + ` AS record_type,
			r.id record_id,
			r.evtlog_id,
			r.block_num,
			t.id tx_id,
			t.tx_hash,
			EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,
			r.time_stamp AS date_time,
			` + addrExpr + `,
			` + intExpr + ` AS int_value,
			` + floatExpr + ` AS float_value,
			` + strExpr + ` AS string_value
		FROM ` + b.table + ` r
		LEFT JOIN transaction t ON t.id=r.tx_id` + addrJoin + `
		WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2)
	)`
}

// adminEventBranches is the registry of all admin event tables surfaced by
// AdminEventsInRange, in record_type order (the numbering is part of the API
// contract; see cgmodel.CGAdminEvent).
var adminEventBranches = []adminEventBranch{
	{recordType: 1, table: "cg_adm_charity_pcent", intValue: "r.percentage"},          // CharityPercentageChanged
	{recordType: 2, table: "cg_adm_main_prize_pcent", intValue: "r.percentage"},       // PrizePercentageChanged
	{recordType: 3, table: "cg_adm_raffle_pcent", intValue: "r.percentage"},           // RafflePercentageChanged
	{recordType: 4, table: "cg_adm_stake_pcent", intValue: "r.percentage"},            // StakingPercentageChanged
	{recordType: 5, table: "cg_adm_raf_eth_bidding", intValue: "r.num_winners"},       // NumRaffleETHWinnersBiddingChanged
	{recordType: 6, table: "cg_adm_raf_nft_bidding", intValue: "r.num_winners"},       // NumRaffleNFTWinnersBiddingChanged
	{recordType: 7, table: "cg_delay_duration", intValue: "r.new_value"},              // DelayDurationBeforeRoundActivationChanged
	{recordType: 8, table: "cg_adm_raf_nft_staking_rwalk", intValue: "r.num_winners"}, // NumRaffleNFTWinnersStakingRWalkChanged
	{recordType: 9, table: "cg_adm_charity_wallet", addrValue: "a.addr", addrJoinFK: "new_charity_aid"},
	{recordType: 10, table: "cg_adm_rwalk_addr", addrValue: "a.addr", addrJoinFK: "new_rwalk_aid"},
	{recordType: 11, table: "cg_adm_prizes_wallet_addr", addrValue: "a.addr", addrJoinFK: "new_wallet_aid"},
	{recordType: 12, table: "cg_adm_staking_cst_addr", addrValue: "a.addr", addrJoinFK: "new_staking_aid"},
	{recordType: 13, table: "cg_adm_staking_rwalk_addr", addrValue: "a.addr", addrJoinFK: "new_staking_aid"},
	{recordType: 14, table: "cg_adm_marketing_addr", addrValue: "a.addr", addrJoinFK: "new_marketing_aid"},
	{recordType: 15, table: "cg_adm_costok_addr", addrValue: "a.addr", addrJoinFK: "new_costok_aid"},
	{recordType: 16, table: "cg_adm_cossig_addr", addrValue: "a.addr", addrJoinFK: "new_cossig_aid"},
	{recordType: 17, table: "cg_adm_upgraded", addrValue: "a.addr", addrJoinFK: "implementation_aid"},                // Upgraded
	{recordType: 18, table: "cg_adm_time_inc", intValue: "r.new_time_inc"},                                           // TimeIncreaseChanged
	{recordType: 19, table: "cg_adm_timeout_claimprize", intValue: "r.new_timeout"},                                  // TimeoutClaimPrizeChanged
	{recordType: 20, table: "cg_adm_price_inc", intValue: "r.new_price_increase"},                                    // PriceIncreaseChanged
	{recordType: 21, table: "cg_adm_prize_microsec", intValue: "r.new_microseconds"},                                 // MainPrizeTimeIncrementChanged
	{recordType: 22, table: "cg_adm_inisecprize", intValue: "r.new_inisec"},                                          // InitialSecondsUntilPrizeChanged
	{recordType: 23, table: "cg_adm_treasurer_addr", addrValue: "a.addr", addrJoinFK: "new_treasurer_aid"},           // TreasurerAddressChanged
	{recordType: 24, table: "cg_adm_acttime", intValue: "r.new_atime"},                                               // ActivationTimeChanged
	{recordType: 25, table: "cg_adm_cst_auclen", intValue: "r.new_len"},                                              // CstDutchAuctionDurationChanged
	{recordType: 26, table: "cg_adm_erc_rwd_mul", floatValue: "r.new_reward/1e18", stringVal: "r.new_reward::TEXT"},  // Erc20RewardMultiplierChanged
	{recordType: 27, table: "cg_adm_cst_min_limit", floatValue: "r.min_limit/1e18", stringVal: "r.min_limit::TEXT"},  // StartingBidPriceCSTMinLimitChanged
	{recordType: 28, table: "cg_adm_mkt_reward", floatValue: "r.new_reward/1e18", stringVal: "r.new_reward::TEXT"},   // MarketingRewardChanged
	{recordType: 29, table: "cg_adm_erc20_reward", floatValue: "r.new_reward/1e18", stringVal: "r.new_reward::TEXT"}, // TokenRewardChanged
	{recordType: 30, table: "cg_adm_msg_len", intValue: "r.new_length"},                                              // MaxMessageLengthChanged
	{recordType: 31, table: "cg_adm_script_url", stringVal: "new_url"},                                               // TokenGenerationScriptURLEvent
	{recordType: 32, table: "cg_adm_base_uri_cs", stringVal: "new_uri"},                                              // BaseURIEvent
	{recordType: 33, table: "cg_adm_initialized", intValue: "version"},                                               // Initialized
	{recordType: 34, table: "cg_adm_ownership"},                                                                      // OwnershipTransferred (two address joins; see ownershipBranchSQL)
	{recordType: 35, table: "cg_adm_timeout_withdraw", intValue: "r.new_timeout"},                                    // TimeoutDurationToWithdrawPrizesChanged
	{recordType: 36, table: "cg_adm_eth_auclen", intValue: "r.new_len"},                                              // EthDutchAuctionDurationDivisorChanged
	{recordType: 37, table: "cg_adm_eth_auc_endprice", intValue: "r.new_len"},                                        // EthDutchAuctionEndingBidPriceDivisorChanged
	{recordType: 38, table: "cg_adm_chrono_pcent", intValue: "r.percentage"},                                         // ChronoWarriorPercentageChanged
	{recordType: 39, table: "cg_adm_cst_auclen_chg_div", intValue: "r.new_len"},                                      // CstDutchAuctionDurationChangeDivisorChanged
}

// ownershipBranchSQL handles record_type 34, whose two address joins (previous
// and new owner) do not fit the single-join template.
const ownershipBranchSQL = `(
		SELECT
			34 AS record_type,
			r.id record_id,
			r.evtlog_id,
			r.block_num,
			t.id tx_id,
			t.tx_hash,
			EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,
			r.time_stamp AS date_time,
			noa.addr AS addr_value,
			0 AS int_value,
			0 AS float_value,
			poa.addr AS string_value
		FROM cg_adm_ownership r
			LEFT JOIN transaction t ON t.id=r.tx_id
			LEFT JOIN address poa ON r.prev_owner_aid=poa.address_id
			LEFT JOIN address noa ON r.new_owner_aid=noa.address_id
		WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2)
	)`

// adminEventsQuery assembles the full UNION over every admin event table.
func adminEventsQuery() string {
	sql := "SELECT record_type, record_id, evtlog_id, block_num, tx_id, tx_hash, ts, date_time, addr_value, int_value, float_value, string_value FROM ("
	for i, b := range adminEventBranches {
		if i > 0 {
			sql += " UNION ALL "
		}
		if b.recordType == 34 {
			sql += ownershipBranchSQL
			continue
		}
		sql += b.sql()
	}
	sql += ") everything ORDER BY evtlog_id"
	return sql
}

// AdminEventsInRange returns every admin/configuration event with
// evtlog_start < evtlog_id < evtlog_end, across all 39 admin event tables,
// ordered by evtlog_id.
func (r *Repo) AdminEventsInRange(ctx context.Context, evtlogStart, evtlogEnd int64) ([]cgmodel.CGAdminEvent, error) {
	scan := func(rows pgx.Rows, rec *cgmodel.CGAdminEvent) error {
		return rows.Scan(
			&rec.RecordType,
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			store.TimeText(&rec.DateTime),
			&rec.AddressValue,
			&rec.IntegerValue,
			&rec.FloatValue,
			&rec.StringValue,
		)
	}
	return queryList(ctx, r, "admin events in range", 256, adminEventsQuery(), scan, evtlogStart, evtlogEnd)
}
