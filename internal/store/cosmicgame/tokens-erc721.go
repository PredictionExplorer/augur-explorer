package cosmicgame

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// nftListSelectSQL is the unified SELECT for Cosmic Signature NFT lists
// (alias m = cg_mint_event); callers append WHERE/ORDER BY/paging.
const nftListSelectSQL = `SELECT
		m.evtlog_id,
		m.block_num,
		t.id,
		t.tx_hash,
		EXTRACT(EPOCH FROM m.time_stamp)::BIGINT,
		m.time_stamp,
		m.owner_aid,
		wa.addr,
		m.cur_owner_aid,
		oa.addr,
		m.seed,
		m.token_id,
		m.token_name,
		m.round_num,
		p.round_num,
		sa.action_id,
		EXTRACT(EPOCH FROM sa.time_stamp)::BIGINT,
		sa.time_stamp,
		ua.action_id,
		EXTRACT(EPOCH FROM ua.time_stamp)::BIGINT,
		ua.time_stamp,
		cst.erc721_token_id,
		endu.erc721_token_id,
		rnw.is_staker,
		rnw.id
	FROM cg_mint_event m
		LEFT JOIN transaction t ON t.id=tx_id
		LEFT JOIN address wa ON m.owner_aid=wa.address_id
		LEFT JOIN address oa ON m.cur_owner_aid=oa.address_id
		LEFT JOIN cg_prize_claim p
			ON m.token_id>=p.token_id AND m.token_id<p.token_id+p.num_cs_nfts
		LEFT JOIN cg_nft_staked_cst sa ON sa.token_id=m.token_id
		LEFT JOIN cg_nft_unstaked_cst ua ON ua.token_id=m.token_id
		LEFT JOIN cg_lastcst_prize cst ON (m.token_id=cst.erc721_token_id AND m.round_num=cst.round_num)
		LEFT JOIN cg_endurance_prize endu ON (m.token_id=endu.erc721_token_id AND m.round_num=endu.round_num)
		LEFT JOIN cg_raffle_nft_prize rnw ON (m.token_id=rnw.token_id AND m.round_num=rnw.round_num)`

// scanNFTListRecord scans one nftListSelectSQL row, deriving the prize
// record type and staking status exactly like the legacy layer: raffle and
// staker-raffle wins override the main-prize default, endurance and last-CST
// prizes override those; a stake action without a matching unstake means the
// token is still staked.
func scanNFTListRecord(rows pgx.Rows, rec *cgmodel.CGCosmicSignatureMintRec) error {
	var nullPrizeNum, nullRaffleID sql.NullInt64
	var nullStaked sql.NullBool
	var nullEnduTokenID, nullStelTokenID sql.NullInt64
	var nullStakeActionID, nullStakeTimestamp sql.NullInt64
	var nullUnstakeActionID, nullUnstakeTimestamp sql.NullInt64

	err := rows.Scan(
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		store.TimeText(&rec.Tx.DateTime),
		&rec.WinnerAid,
		&rec.WinnerAddr,
		&rec.CurOwnerAid,
		&rec.CurOwnerAddr,
		&rec.Seed,
		&rec.TokenId,
		&rec.TokenName,
		&rec.RoundNum,
		&nullPrizeNum,
		&nullStakeActionID,
		&nullStakeTimestamp,
		store.NullTimeText(&rec.StakeDateTime),
		&nullUnstakeActionID,
		&nullUnstakeTimestamp,
		store.NullTimeText(&rec.ActualUnstakeDateTime),
		&nullStelTokenID,
		&nullEnduTokenID,
		&nullStaked,
		&nullRaffleID,
	)
	if err != nil {
		return err
	}

	rec.RecordType = 3 // Main prize (default)
	if nullRaffleID.Valid {
		rec.RecordType = 1 // Raffle NFT winner
	}
	if nullStaked.Valid && nullStaked.Bool {
		rec.RecordType = 2 // NFT won due to staking (RWalk)
	}
	if nullEnduTokenID.Valid {
		rec.RecordType = 4 // Endurance champion
	}
	if nullStelTokenID.Valid {
		rec.RecordType = 5 // Last CST bidder
	}

	if nullStakeActionID.Valid {
		rec.StakeActionId = nullStakeActionID.Int64
	}
	if nullStakeTimestamp.Valid {
		rec.StakeTimeStamp = nullStakeTimestamp.Int64
	}
	if nullUnstakeActionID.Valid {
		rec.UnstakeActionId = nullUnstakeActionID.Int64
	}
	if nullUnstakeTimestamp.Valid {
		rec.ActualUnstakeTimeStamp = nullUnstakeTimestamp.Int64
	}
	if rec.StakeActionId > 0 && rec.UnstakeActionId > 0 {
		rec.WasUnstaked = true
	}
	if rec.StakeActionId > 0 && rec.UnstakeActionId == 0 {
		rec.Staked = true
	}
	return nil
}

// CosmicSignatureTokens returns the minted Cosmic Signature NFTs, newest
// first, with prize provenance and CST staking status attached.
// limit 0 means no effective limit.
func (r *Repo) CosmicSignatureTokens(ctx context.Context, offset, limit int) ([]cgmodel.CGCosmicSignatureMintRec, error) {
	if limit == 0 {
		limit = 1000000
	}
	query := nftListSelectSQL + `
		ORDER BY m.id DESC
		OFFSET $1 LIMIT $2`
	return queryList(ctx, r, "cosmic signature tokens", 64, query, scanNFTListRecord, offset, limit)
}

// CosmicSignatureTokenInfo returns one minted token with prize provenance
// and its full CST staking state (including the staked-owner address while
// staked), or store.ErrNotFound for an unknown token id.
func (r *Repo) CosmicSignatureTokenInfo(ctx context.Context, tokenID int64) (cgmodel.CGCosmicSignatureMintRec, error) {
	query := `SELECT
			m.evtlog_id,
			m.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM m.time_stamp)::BIGINT,
			m.time_stamp,
			m.owner_aid,
			wa.addr,
			m.cur_owner_aid,
			oa.addr,
			m.seed,
			m.token_id,
			m.round_num,
			p.round_num,
			m.token_name,
			st.staker_aid,
			sta.addr,
			sa.action_id,
			EXTRACT(EPOCH FROM sa.time_stamp)::BIGINT,
			sa.time_stamp,
			u.id,
			EXTRACT(EPOCH FROM u.time_stamp)::BIGINT,
			u.time_stamp,
			cst.erc721_token_id,
			endu.erc721_token_id,
			rnw.is_staker,
			rnw.id
		FROM cg_mint_event m
			LEFT JOIN transaction t ON t.id=tx_id
			LEFT JOIN address wa ON m.owner_aid=wa.address_id
			LEFT JOIN address oa ON m.cur_owner_aid=oa.address_id
			LEFT JOIN cg_prize_claim p
				ON m.token_id>=p.token_id AND m.token_id<p.token_id+p.num_cs_nfts
			LEFT JOIN cg_staked_token_cst st ON (m.token_id=st.token_id)
			LEFT JOIN cg_nft_staked_cst sa ON sa.token_id = m.token_id
			LEFT JOIN cg_nft_unstaked_cst u ON u.token_id=m.token_id
			LEFT JOIN cg_lastcst_prize cst ON m.token_id=cst.erc721_token_id
			LEFT JOIN cg_endurance_prize endu ON m.token_id=endu.erc721_token_id
			LEFT JOIN cg_raffle_nft_prize rnw ON m.token_id=rnw.token_id
			LEFT JOIN address sta ON st.staker_aid = sta.address_id
		WHERE m.token_id=$1`

	var rec cgmodel.CGCosmicSignatureMintRec
	var nullPrizeNum, nullUnstakeID, nullActionID, nullStakerAid, nullRaffleID sql.NullInt64
	var nullAuTimestamp, nullSaTimestamp sql.NullInt64
	var nullStakedOwnerAddr sql.NullString
	var nullStaked sql.NullBool
	var nullEnduTokenID, nullStelTokenID sql.NullInt64
	err := r.q(ctx).QueryRow(ctx, query, tokenID).Scan(
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		store.TimeText(&rec.Tx.DateTime),
		&rec.WinnerAid,
		&rec.WinnerAddr,
		&rec.CurOwnerAid,
		&rec.CurOwnerAddr,
		&rec.Seed,
		&rec.TokenId,
		&rec.RoundNum,
		&nullPrizeNum,
		&rec.TokenName,
		&nullStakerAid,
		&nullStakedOwnerAddr,
		&nullActionID,
		&nullSaTimestamp,
		store.NullTimeText(&rec.StakeDateTime),
		&nullUnstakeID,
		&nullAuTimestamp,
		store.NullTimeText(&rec.ActualUnstakeDateTime),
		&nullStelTokenID,
		&nullEnduTokenID,
		&nullStaked,
		&nullRaffleID,
	)
	if err != nil {
		return cgmodel.CGCosmicSignatureMintRec{}, store.WrapError("cosmic signature token info", err)
	}

	rec.RecordType = 3 // main prize
	if nullRaffleID.Valid {
		rec.RecordType = 1 // raffle NFT winner
	}
	if nullStaked.Valid && nullStaked.Bool {
		rec.RecordType = 2 // nft won due to staking (RWalk)
	}
	if nullEnduTokenID.Valid {
		rec.RecordType = 4 // endurance champion
	}
	if nullStelTokenID.Valid {
		rec.RecordType = 5 // stellar spender
	}
	if nullUnstakeID.Valid {
		rec.UnstakeActionId = nullUnstakeID.Int64
	} else {
		rec.UnstakeActionId = -1
	}
	if nullActionID.Valid {
		rec.StakeActionId = nullActionID.Int64
	} else {
		rec.StakeActionId = -1
	}
	if nullAuTimestamp.Valid {
		rec.ActualUnstakeTimeStamp = nullAuTimestamp.Int64
	}
	if nullStakedOwnerAddr.Valid {
		rec.StakedOwnerAddr = nullStakedOwnerAddr.String
	}
	// Note: nullStaked comes from rnw.is_staker which means "winner WAS a
	// staker when they won". It does NOT indicate whether the token is
	// currently staked, so we don't use it here.
	if nullStakerAid.Valid {
		rec.StakedOwnerAid = nullStakerAid.Int64
	} else {
		rec.StakedOwnerAid = -1
	}
	if nullSaTimestamp.Valid {
		rec.StakeTimeStamp = nullSaTimestamp.Int64
	} else {
		rec.StakeTimeStamp = -1
	}
	if rec.StakeActionId > -1 && rec.UnstakeActionId > -1 {
		rec.WasUnstaked = true
	}
	// Token is staked if it's in cg_staked_token_cst (nullStakerAid.Valid)
	// OR if there's a stake action with no corresponding unstake action.
	switch {
	case nullStakerAid.Valid:
		rec.Staked = true
	case rec.StakeActionId > -1 && rec.UnstakeActionId == -1:
		rec.Staked = true
	default:
		rec.Staked = false
	}
	return rec, nil
}

// TokenNameHistory returns every rename of one token, newest first, with the
// address that performed each rename.
func (r *Repo) TokenNameHistory(ctx context.Context, tokenID int64) ([]cgmodel.CGTokenName, error) {
	query := `SELECT
			n.evtlog_id,
			n.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM n.time_stamp)::BIGINT,
			n.time_stamp,
			n.token_id,
			n.token_name,
			t.from_aid,
			a.addr
		FROM cg_token_name n
			LEFT JOIN transaction t ON t.id=n.tx_id
			LEFT JOIN address a ON a.address_id=t.from_aid
		WHERE n.token_id=$1
		ORDER BY n.id DESC`
	scan := func(rows pgx.Rows, rec *cgmodel.CGTokenName) error {
		return rows.Scan(
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.TokenId,
			&rec.TokenName,
			&rec.ChangedByAid,
			&rec.ChangedBy,
		)
	}
	return queryList(ctx, r, "token name history", 64, query, scan, tokenID)
}

// TokenOwnershipTransfers returns the ERC-721 transfer history of one token,
// newest first. limit 0 means no effective limit.
func (r *Repo) TokenOwnershipTransfers(ctx context.Context, tokenID int64, offset, limit int) ([]cgmodel.CGTransfer, error) {
	if limit == 0 {
		limit = 1000000
	}
	query := `SELECT
			t.id,
			t.evtlog_id,
			t.block_num,
			tx.id,
			tx.tx_hash,
			EXTRACT(EPOCH FROM t.time_stamp)::BIGINT,
			t.time_stamp,
			t.from_aid,
			fa.addr,
			t.to_aid,
			ta.addr,
			t.token_id,
			t.otype
		FROM cg_erc721_transfer t
			LEFT JOIN transaction tx ON tx.id=t.tx_id
			LEFT JOIN address fa ON t.from_aid=fa.address_id
			LEFT JOIN address ta ON t.to_aid=ta.address_id
		WHERE t.token_id=$1
		ORDER BY t.id DESC
		OFFSET $2 LIMIT $3`
	scan := func(rows pgx.Rows, rec *cgmodel.CGTransfer) error {
		return rows.Scan(
			&rec.RecordId,
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.FromAid,
			&rec.FromAddr,
			&rec.ToAid,
			&rec.ToAddr,
			&rec.TokenId,
			&rec.TransferType,
		)
	}
	return queryList(ctx, r, "token ownership transfers", 256, query, scan, tokenID, offset, limit)
}

// CosmicSignatureTokenDistribution returns how many tokens each current
// owner holds, largest holder first.
func (r *Repo) CosmicSignatureTokenDistribution(ctx context.Context) ([]cgmodel.CGCSTokenDistributionRec, error) {
	query := `SELECT
			m.cur_owner_aid,
			a.addr,
			COUNT(*) AS num_tokens
		FROM cg_mint_event m
			LEFT JOIN address a ON m.cur_owner_aid=a.address_id
		GROUP BY m.cur_owner_aid, a.addr
		ORDER BY num_tokens DESC`
	scan := func(rows pgx.Rows, rec *cgmodel.CGCSTokenDistributionRec) error {
		return rows.Scan(&rec.OwnerAid, &rec.OwnerAddr, &rec.NumTokens)
	}
	return queryList(ctx, r, "cosmic signature token distribution", 32, query, scan)
}

func scanTokenSearchResult(rows pgx.Rows, rec *cgmodel.CGTokenSearchResult) error {
	return rows.Scan(
		&rec.MintTimeStamp,
		store.TimeText(&rec.MintDateTime),
		&rec.TokenId,
		&rec.TokenName,
	)
}

// SearchTokensByName returns the tokens whose name contains name
// (case-insensitive), ordered by token id.
func (r *Repo) SearchTokensByName(ctx context.Context, name string) ([]cgmodel.CGTokenSearchResult, error) {
	query := `SELECT
			EXTRACT(EPOCH FROM t.time_stamp)::BIGINT,
			t.time_stamp,
			t.token_id,
			t.token_name
		FROM cg_mint_event t
		WHERE t.token_name ILIKE  $1
		ORDER BY t.token_id`
	return queryList(ctx, r, "search tokens by name", 256, query, scanTokenSearchResult, "%"+name+"%")
}

// NamedTokens returns every token that has been given a name, ordered by
// name.
func (r *Repo) NamedTokens(ctx context.Context) ([]cgmodel.CGTokenSearchResult, error) {
	query := `SELECT
			EXTRACT(EPOCH FROM t.time_stamp)::BIGINT,
			t.time_stamp,
			t.token_id,
			t.token_name
		FROM cg_mint_event t
		WHERE LENGTH(t.token_name)>0
		ORDER BY t.token_name`
	return queryList(ctx, r, "named tokens", 256, query, scanTokenSearchResult)
}

// CosmicSignatureTokenCount returns the total number of minted Cosmic
// Signature tokens.
func (r *Repo) CosmicSignatureTokenCount(ctx context.Context) (int64, error) {
	var numToks int64
	err := r.q(ctx).QueryRow(ctx, "SELECT COUNT(*) FROM cg_mint_event").Scan(&numToks)
	if err != nil {
		return 0, store.WrapError("cosmic signature token count", err)
	}
	return numToks, nil
}

// CosmicSignatureTokenSeed returns the generation seed of one token, or
// store.ErrNotFound for an unknown token id.
func (r *Repo) CosmicSignatureTokenSeed(ctx context.Context, tokenID int64) (string, error) {
	var seed string
	err := r.q(ctx).QueryRow(ctx, "SELECT seed FROM cg_mint_event WHERE token_id=$1", tokenID).Scan(&seed)
	if err != nil {
		return "", store.WrapError("cosmic signature token seed", err)
	}
	return seed, nil
}
