package cosmicgame

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// charityDonationsSelectSQL is the shared SELECT of the charity-wallet
// inbound donation queries (alias d = cg_donation_received).
const charityDonationsSelectSQL = `SELECT
			d.evtlog_id,
			d.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,
			d.time_stamp,
			d.donor_aid,
			da.addr,
			d.amount,
			d.amount/1e18 amount_eth,
			d.round_num
		FROM cg_donation_received d
			LEFT JOIN transaction t ON t.id=tx_id
			LEFT JOIN address da ON d.donor_aid=da.address_id`

func scanCharityDonation(rows pgx.Rows, rec *cgmodel.CGCharityDonation) error {
	return rows.Scan(
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		store.TimeText(&rec.Tx.DateTime),
		&rec.DonorAid,
		&rec.DonorAddr,
		&rec.Amount,
		&rec.AmountEth,
		&rec.RoundNum,
	)
}

// CharityDonations returns every donation received by the charity wallet,
// newest first, marking donations that did not come from the game contract
// (cosmicGameAid) as voluntary.
func (r *Repo) CharityDonations(ctx context.Context, cosmicGameAid int64) ([]cgmodel.CGCharityDonation, error) {
	query := charityDonationsSelectSQL + `
		ORDER BY d.id DESC`
	scan := func(rows pgx.Rows, rec *cgmodel.CGCharityDonation) error {
		if err := scanCharityDonation(rows, rec); err != nil {
			return err
		}
		rec.IsVoluntary = rec.DonorAid != cosmicGameAid
		return nil
	}
	return queryList(ctx, r, "charity donations", 32, query, scan)
}

// CharityDonationsFromCosmicGame returns the charity-wallet donations sent
// by the game contract itself (the per-round charity cut), newest first.
func (r *Repo) CharityDonationsFromCosmicGame(ctx context.Context, cosmicGameAid int64) ([]cgmodel.CGCharityDonation, error) {
	query := charityDonationsSelectSQL + `
		WHERE donor_aid = $1
		ORDER BY d.id DESC`
	return queryList(ctx, r, "charity donations from cosmic game", 32, query, scanCharityDonation, cosmicGameAid)
}

// CharityDonationsVoluntary returns the charity-wallet donations that did
// not come from the game contract, newest first.
func (r *Repo) CharityDonationsVoluntary(ctx context.Context, cosmicGameAid int64) ([]cgmodel.CGCharityDonation, error) {
	query := charityDonationsSelectSQL + `
		WHERE donor_aid != $1
		ORDER BY d.id DESC`
	scan := func(rows pgx.Rows, rec *cgmodel.CGCharityDonation) error {
		if err := scanCharityDonation(rows, rec); err != nil {
			return err
		}
		rec.IsVoluntary = true
		return nil
	}
	return queryList(ctx, r, "charity donations voluntary", 32, query, scan, cosmicGameAid)
}

// CharityWalletWithdrawals returns every outbound transfer of the charity
// wallet, newest first.
func (r *Repo) CharityWalletWithdrawals(ctx context.Context) ([]cgmodel.CGCharityWithdrawal, error) {
	query := `SELECT
			w.id,
			w.evtlog_id,
			w.block_num,
			tx.id,
			tx.tx_hash,
			EXTRACT(EPOCH FROM w.time_stamp)::BIGINT,
			w.time_stamp,
			ca.addr,
			w.amount,
			w.amount/1e18
		FROM cg_donation_sent w
			LEFT JOIN transaction tx ON tx.id=w.tx_id
			LEFT JOIN address ca ON w.charity_aid=ca.address_id
		ORDER BY w.id DESC`
	scan := func(rows pgx.Rows, rec *cgmodel.CGCharityWithdrawal) error {
		return rows.Scan(
			&rec.RecordId,
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.DestinationAddr,
			&rec.Amount,
			&rec.AmountEth,
		)
	}
	return queryList(ctx, r, "charity wallet withdrawals", 256, query, scan)
}

// simpleEthDonationsSelectSQL is the shared SELECT of the plain EthDonated
// queries (alias d = cg_eth_donated).
const simpleEthDonationsSelectSQL = `SELECT
			d.evtlog_id,
			d.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,
			d.time_stamp,
			d.donor_aid,
			da.addr,
			d.amount,
			d.amount/1e18 amount_eth,
			d.round_num
		FROM cg_eth_donated d
			LEFT JOIN transaction t ON t.id=tx_id
			LEFT JOIN address da ON d.donor_aid=da.address_id`

func scanSimpleEthDonation(rows pgx.Rows, rec *cgmodel.CGCosmicGameDonationSimple) error {
	return rows.Scan(
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		store.TimeText(&rec.Tx.DateTime),
		&rec.DonorAid,
		&rec.DonorAddr,
		&rec.Amount,
		&rec.AmountEth,
		&rec.RoundNum,
	)
}

// SimpleEthDonations returns the plain (no-info) ETH donations to the game,
// newest first.
func (r *Repo) SimpleEthDonations(ctx context.Context, offset, limit int) ([]cgmodel.CGCosmicGameDonationSimple, error) {
	query := simpleEthDonationsSelectSQL + `
		ORDER BY d.id DESC
		OFFSET $1 LIMIT $2`
	return queryList(ctx, r, "simple eth donations", 32, query, scanSimpleEthDonation, offset, limit)
}

// SimpleEthDonationsByRound returns the plain ETH donations of one round,
// newest first.
func (r *Repo) SimpleEthDonationsByRound(ctx context.Context, roundNum int64) ([]cgmodel.CGCosmicGameDonationSimple, error) {
	query := simpleEthDonationsSelectSQL + `
		WHERE d.round_num = $1
		ORDER BY d.id DESC`
	return queryList(ctx, r, "simple eth donations by round", 32, query, scanSimpleEthDonation, roundNum)
}

// ethDonationsWithInfoSelectSQL is the shared SELECT of the
// EthDonatedWithInfo queries (alias d = cg_eth_donated_wi, dj = the
// donation's JSON payload).
const ethDonationsWithInfoSelectSQL = `SELECT
			d.evtlog_id,
			d.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,
			d.time_stamp,
			d.donor_aid,
			da.addr,
			d.amount,
			d.amount/1e18 amount_eth,
			d.round_num,
			d.record_id,
			dj.data
		FROM cg_eth_donated_wi d
			LEFT JOIN cg_donation_json dj ON dj.record_id=d.record_id
			LEFT JOIN transaction t ON t.id=tx_id
			LEFT JOIN address da ON d.donor_aid=da.address_id`

func scanEthDonationWithInfo(rows pgx.Rows, rec *cgmodel.CGCosmicGameDonationWithInfo) error {
	return rows.Scan(
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		store.TimeText(&rec.Tx.DateTime),
		&rec.DonorAid,
		&rec.DonorAddr,
		&rec.Amount,
		&rec.AmountEth,
		&rec.RoundNum,
		&rec.CGRecordId,
		&rec.DataJson,
	)
}

// EthDonationsWithInfo returns the ETH donations that carry a JSON info
// payload, newest first.
func (r *Repo) EthDonationsWithInfo(ctx context.Context, offset, limit int) ([]cgmodel.CGCosmicGameDonationWithInfo, error) {
	query := ethDonationsWithInfoSelectSQL + `
		ORDER BY d.id DESC
		OFFSET $1 LIMIT $2`
	return queryList(ctx, r, "eth donations with info", 32, query, scanEthDonationWithInfo, offset, limit)
}

// EthDonationsWithInfoByRound returns one round's ETH donations that carry a
// JSON info payload, newest first.
func (r *Repo) EthDonationsWithInfoByRound(ctx context.Context, roundNum int64) ([]cgmodel.CGCosmicGameDonationWithInfo, error) {
	query := ethDonationsWithInfoSelectSQL + `
		WHERE d.round_num=$1
		ORDER BY d.id DESC`
	return queryList(ctx, r, "eth donations with info by round", 32, query, scanEthDonationWithInfo, roundNum)
}

// EthDonationWithInfoRecord returns one info-carrying donation by its
// contract-side record id, or store.ErrNotFound. The legacy layer served
// the zero-value record for unknown ids; callers that need that shape
// translate ErrNotFound back into it.
func (r *Repo) EthDonationWithInfoRecord(ctx context.Context, recordID int64) (cgmodel.CGCosmicGameDonationWithInfo, error) {
	query := ethDonationsWithInfoSelectSQL + `
		WHERE d.record_id=$1`
	var rec cgmodel.CGCosmicGameDonationWithInfo
	err := r.q(ctx).QueryRow(ctx, query, recordID).Scan(
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		store.TimeText(&rec.Tx.DateTime),
		&rec.DonorAid,
		&rec.DonorAddr,
		&rec.Amount,
		&rec.AmountEth,
		&rec.RoundNum,
		&rec.CGRecordId,
		&rec.DataJson,
	)
	if err != nil {
		return cgmodel.CGCosmicGameDonationWithInfo{}, store.WrapError("eth donation with info record", err)
	}
	return rec, nil
}

// DonationReceivedEvtIDByTx returns the evtlog_id of the DonationReceived
// event within transaction txID (matched by topic0 signature), or 0 when the
// transaction has none.
func (r *Repo) DonationReceivedEvtIDByTx(ctx context.Context, txID int64, sig string) (int64, error) {
	query := `SELECT
			d.evtlog_id
		FROM
			evt_log e
			LEFT JOIN cg_donation_received d ON e.id=d.evtlog_id
		WHERE
			(e.tx_id=$1) AND
			(e.topic0_sig=$2)
		LIMIT 1`
	var nullID sql.NullInt64
	err := r.q(ctx).QueryRow(ctx, query, txID, sig).Scan(&nullID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, nil
		}
		return 0, store.WrapError("donation received evt id by tx", err)
	}
	return nullID.Int64, nil
}

// combinedEthDonationsSQL builds the UNION of plain and info-carrying ETH
// donations, newest first. filter is one of the whitelisted WHERE clauses
// applied to both branches ("" = no filter).
func combinedEthDonationsSQL(filter string) string {
	return `SELECT
			record_type,
			evtlog_id,
			block_num,
			tx_id,
			tx_hash,
			ts,
			date_time,
			donor_aid,
			donor_addr,
			amount,
			amount_eth,
			round_num,
			record_id,
			json_data
		FROM (
			(
				SELECT
					0 AS record_type,
					d.evtlog_id,
					d.block_num,
					t.id tx_id,
					t.tx_hash,
					EXTRACT(EPOCH FROM d.time_stamp)::BIGINT ts,
					d.time_stamp date_time,
					d.donor_aid,
					da.addr donor_addr,
					d.amount,
					d.amount/1e18 amount_eth,
					d.round_num,
					-1 AS record_id,
					'' AS json_data
				FROM cg_eth_donated d
					LEFT JOIN transaction t ON t.id=tx_id
					LEFT JOIN address da ON d.donor_aid=da.address_id
				` + filter + `
			) UNION ALL (
				SELECT
					1 AS record_type,
					d.evtlog_id,
					d.block_num,
					t.id tx_id,
					t.tx_hash,
					EXTRACT(EPOCH FROM d.time_stamp)::BIGINT ts,
					d.time_stamp date_time,
					d.donor_aid,
					da.addr donor_addr,
					d.amount,
					d.amount/1e18 amount_eth,
					d.round_num,
					d.record_id,
					dj.data json_data
				FROM cg_eth_donated_wi d
					LEFT JOIN cg_donation_json dj ON dj.record_id=d.record_id
					LEFT JOIN transaction t ON t.id=tx_id
					LEFT JOIN address da ON d.donor_aid=da.address_id
				` + filter + `
			)
		) donations
		ORDER BY evtlog_id DESC`
}

func scanCombinedEthDonation(rows pgx.Rows, rec *cgmodel.CGDonationCombinedRec) error {
	return rows.Scan(
		&rec.RecordType,
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		store.TimeText(&rec.Tx.DateTime),
		&rec.DonorAid,
		&rec.DonorAddr,
		&rec.Amount,
		&rec.AmountEth,
		&rec.RoundNum,
		&rec.CGRecordId,
		&rec.DataJson,
	)
}

// EthDonationsByUser returns both donation kinds (plain and with-info) made
// by one donor, newest first.
func (r *Repo) EthDonationsByUser(ctx context.Context, userAid int64) ([]cgmodel.CGDonationCombinedRec, error) {
	query := combinedEthDonationsSQL("WHERE d.donor_aid = $1")
	return queryList(ctx, r, "eth donations by user", 32, query, scanCombinedEthDonation, userAid)
}

// EthDonationsByRound returns both donation kinds (plain and with-info) of
// one round, newest first.
func (r *Repo) EthDonationsByRound(ctx context.Context, roundNum int64) ([]cgmodel.CGDonationCombinedRec, error) {
	query := combinedEthDonationsSQL("WHERE d.round_num = $1")
	return queryList(ctx, r, "eth donations by round", 32, query, scanCombinedEthDonation, roundNum)
}

// EthDonations returns both donation kinds (plain and with-info) over all
// rounds, newest first.
func (r *Repo) EthDonations(ctx context.Context) ([]cgmodel.CGDonationCombinedRec, error) {
	query := combinedEthDonationsSQL("")
	return queryList(ctx, r, "eth donations", 32, query, scanCombinedEthDonation)
}

// RoundEthDonationKind identifies which direct-donation contract event
// produced a v2 round donation record.
type RoundEthDonationKind string

// The two direct-donation record kinds.
const (
	RoundEthDonationPlain    RoundEthDonationKind = "plain"
	RoundEthDonationWithInfo RoundEthDonationKind = "withInfo"
)

// RoundEthDonationRecord is the exact-wei store projection used by the v2
// round donation collection. The legacy records retain float amounts and
// sentinel fields for frozen v1 responses.
type RoundEthDonationRecord struct {
	Tx               cgmodel.Transaction
	RoundNum         int64
	DonorAddr        string
	EthAmountWei     string
	Kind             RoundEthDonationKind
	ContractRecordID *int64
	Data             *string
}

func roundEthDonationsPageSQL(after bool) string {
	filter := "WHERE d.round_num = $1"
	limitPlaceholder := "$2"
	if after {
		filter += " AND d.evtlog_id < $2"
		limitPlaceholder = "$3"
	}
	return fmt.Sprintf(`SELECT
			donation_kind,
			evtlog_id,
			block_num,
			tx_id,
			tx_hash,
			ts,
			date_time,
			donor_addr,
			amount_wei,
			round_num,
			contract_record_id,
			data
		FROM (
			(SELECT
				'plain'::TEXT AS donation_kind,
				d.evtlog_id,
				d.block_num,
				t.id AS tx_id,
				t.tx_hash,
				EXTRACT(EPOCH FROM d.time_stamp)::BIGINT AS ts,
				d.time_stamp AS date_time,
				da.addr AS donor_addr,
				d.amount::TEXT AS amount_wei,
				d.round_num,
				NULL::BIGINT AS contract_record_id,
				NULL::TEXT AS data
			FROM cg_eth_donated d
				LEFT JOIN transaction t ON t.id=d.tx_id
				LEFT JOIN address da ON da.address_id=d.donor_aid
			%s
			ORDER BY d.evtlog_id DESC
			LIMIT %s)
			UNION ALL
			(SELECT
				'withInfo'::TEXT AS donation_kind,
				d.evtlog_id,
				d.block_num,
				t.id AS tx_id,
				t.tx_hash,
				EXTRACT(EPOCH FROM d.time_stamp)::BIGINT AS ts,
				d.time_stamp AS date_time,
				da.addr AS donor_addr,
				d.amount::TEXT AS amount_wei,
				d.round_num,
				d.record_id AS contract_record_id,
				dj.data
			FROM cg_eth_donated_wi d
				LEFT JOIN cg_donation_json dj ON dj.record_id=d.record_id
				LEFT JOIN transaction t ON t.id=d.tx_id
				LEFT JOIN address da ON da.address_id=d.donor_aid
			%s
			ORDER BY d.evtlog_id DESC
			LIMIT %s)
		) donations
		ORDER BY evtlog_id DESC
		LIMIT %s`, filter, limitPlaceholder, filter, limitPlaceholder, limitPlaceholder)
}

func scanRoundEthDonation(rows pgx.Rows, rec *RoundEthDonationRecord) error {
	var (
		kind     string
		recordID sql.NullInt64
		data     sql.NullString
	)
	if err := rows.Scan(
		&kind,
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		store.TimeText(&rec.Tx.DateTime),
		&rec.DonorAddr,
		&rec.EthAmountWei,
		&rec.RoundNum,
		&recordID,
		&data,
	); err != nil {
		return err
	}
	rec.Kind = RoundEthDonationKind(kind)
	if recordID.Valid {
		value := recordID.Int64
		rec.ContractRecordID = &value
	}
	if data.Valid {
		value := data.String
		rec.Data = &value
	}
	return nil
}

// EthDonationsByRoundPage returns at most limit plain and info-carrying ETH
// donations before the supplied newest-first event cursor.
func (r *Repo) EthDonationsByRoundPage(
	ctx context.Context,
	roundNum int64,
	after *DonationPageCursor,
	limit int,
) (records []RoundEthDonationRecord, hasMore bool, err error) {
	const op = "eth donations by round page"
	if roundNum < 0 {
		return nil, false, fmt.Errorf("%s: round must be non-negative", op)
	}
	if limit <= 0 || limit > maxDonationPageLimit {
		return nil, false, fmt.Errorf("%s: limit must be between 1 and %d", op, maxDonationPageLimit)
	}
	args := []any{roundNum, limit + 1}
	if after != nil {
		if after.EventLogID < 1 {
			return nil, false, fmt.Errorf("%s: invalid cursor", op)
		}
		args = []any{roundNum, after.EventLogID, limit + 1}
	}
	records, err = queryList(
		ctx,
		r,
		op,
		limit+1,
		roundEthDonationsPageSQL(after != nil),
		scanRoundEthDonation,
		args...,
	)
	if err != nil {
		return nil, false, err
	}
	if len(records) > limit {
		records = records[:limit]
		hasMore = true
	}
	return records, hasMore, nil
}
