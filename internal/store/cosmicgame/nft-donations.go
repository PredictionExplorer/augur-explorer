package cosmicgame

import (
	"context"

	"github.com/jackc/pgx/v5"

	p "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// scanNFTDonationDonorFirst matches the SELECT order of the global and
// per-user donation lists (donor columns before round_num, idx after the
// token contract address id).
func scanNFTDonationDonorFirst(rows pgx.Rows, rec *p.CGNFTDonation) error {
	return rows.Scan(
		&rec.RecordId,
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		store.TimeText(&rec.Tx.DateTime),
		&rec.DonorAid,
		&rec.DonorAddr,
		&rec.NFTTokenId,
		&rec.RoundNum,
		&rec.TokenAddressId,
		&rec.Index,
		&rec.TokenAddr,
		&rec.NFTTokenURI,
	)
}

// scanNFTDonationRoundFirst matches the SELECT order of the per-round and
// per-token donation lists (round_num first, idx before the token contract
// address id).
func scanNFTDonationRoundFirst(rows pgx.Rows, rec *p.CGNFTDonation) error {
	return rows.Scan(
		&rec.RecordId,
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		store.TimeText(&rec.Tx.DateTime),
		&rec.RoundNum,
		&rec.DonorAid,
		&rec.DonorAddr,
		&rec.NFTTokenId,
		&rec.Index,
		&rec.TokenAddressId,
		&rec.TokenAddr,
		&rec.NFTTokenURI,
	)
}

// NFTDonations returns every NFT donated to the game, newest first.
// limit 0 means no effective limit.
func (r *Repo) NFTDonations(ctx context.Context, offset, limit int) ([]p.CGNFTDonation, error) {
	if limit == 0 {
		limit = 1000000
	}
	query := `SELECT
			d.id,
			d.evtlog_id,
			d.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,
			d.time_stamp,
			d.donor_aid,
			da.addr,
			d.token_id,
			d.round_num,
			nft.address_id,
			d.idx,
			nft.addr,
			d.token_uri
		FROM cg_nft_donation d
			LEFT JOIN transaction t ON t.id=tx_id
			LEFT JOIN address da ON d.donor_aid=da.address_id
			LEFT JOIN address nft ON d.token_aid=nft.address_id
		ORDER BY d.id DESC
		OFFSET $1 LIMIT $2`
	return queryList(ctx, r, "nft donations", 256, query, scanNFTDonationDonorFirst, offset, limit)
}

// NFTDonationInfo returns one NFT donation by record id, or
// store.ErrNotFound when the id does not exist. The returned record does not
// carry round_num or idx (the legacy query never selected them).
func (r *Repo) NFTDonationInfo(ctx context.Context, id int64) (p.CGNFTDonation, error) {
	query := `SELECT
			d.evtlog_id,
			d.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,
			d.time_stamp,
			d.donor_aid,
			da.addr,
			d.token_id,
			nft.address_id,
			nft.addr,
			d.token_uri
		FROM cg_nft_donation d
			LEFT JOIN transaction t ON t.id=tx_id
			LEFT JOIN address da ON d.donor_aid=da.address_id
			LEFT JOIN address nft ON d.token_aid=nft.address_id
		WHERE d.id=$1`
	var rec p.CGNFTDonation
	rec.RecordId = id
	err := r.pool().QueryRow(ctx, query, id).Scan(
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		store.TimeText(&rec.Tx.DateTime),
		&rec.DonorAid,
		&rec.DonorAddr,
		&rec.NFTTokenId,
		&rec.TokenAddressId,
		&rec.TokenAddr,
		&rec.NFTTokenURI,
	)
	if err != nil {
		return p.CGNFTDonation{RecordId: id}, store.WrapError("nft donation info", err)
	}
	return rec, nil
}

// DonatedNFTClaims returns the claim events of donated NFTs, newest first.
// limit 0 means no effective limit.
func (r *Repo) DonatedNFTClaims(ctx context.Context, offset, limit int) ([]p.CGDonatedNFTClaimRec, error) {
	if limit == 0 {
		limit = 1000000
	}
	query := `SELECT
			c.id,
			c.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM c.time_stamp)::BIGINT,
			c.time_stamp,
			c.round_num,
			ta.addr,
			c.token_id,
			c.idx,
			c.winner_aid,
			wa.addr,
			da.addr,
			d.token_uri
		FROM cg_donated_nft_claimed c
			LEFT JOIN transaction t ON t.id=c.tx_id
			LEFT JOIN address ta ON c.token_aid=ta.address_id
			LEFT JOIN address wa ON c.winner_aid=wa.address_id
			LEFT JOIN cg_nft_donation d ON c.idx=d.idx
			LEFT JOIN address da ON d.donor_aid=da.address_id
		ORDER BY c.id DESC
		OFFSET $1 LIMIT $2`
	scan := func(rows pgx.Rows, rec *p.CGDonatedNFTClaimRec) error {
		return rows.Scan(
			&rec.RecordId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.RoundNum,
			&rec.TokenAddr,
			&rec.NFTTokenId,
			&rec.WinnerIndex,
			&rec.WinnerAid,
			&rec.WinnerAddr,
			&rec.DonorAddr,
			&rec.NFTTokenURI,
		)
	}
	return queryList(ctx, r, "donated nft claims", 256, query, scan, offset, limit)
}

// nftDonationsRoundFirstSQL is the shared SELECT of the per-round and
// per-token donation queries; callers append their WHERE/ORDER BY.
const nftDonationsRoundFirstSQL = `SELECT
			d.id,
			d.evtlog_id,
			d.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,
			d.time_stamp,
			d.round_num,
			d.donor_aid,
			da.addr,
			d.token_id,
			d.idx,
			nft.address_id,
			nft.addr,
			d.token_uri
		FROM cg_nft_donation d
			LEFT JOIN transaction t ON t.id=tx_id
			LEFT JOIN address da ON d.donor_aid=da.address_id
			LEFT JOIN address nft ON d.token_aid=nft.address_id`

// NFTDonationsByRound returns the NFTs donated during one round, newest
// first.
func (r *Repo) NFTDonationsByRound(ctx context.Context, roundNum int64) ([]p.CGNFTDonation, error) {
	query := nftDonationsRoundFirstSQL + `
		WHERE d.round_num= $1
		ORDER BY d.id DESC`
	return queryList(ctx, r, "nft donations by round", 256, query, scanNFTDonationRoundFirst, roundNum)
}

// NFTDonationsByToken returns every donation of one NFT contract (by its
// address id), newest round first.
func (r *Repo) NFTDonationsByToken(ctx context.Context, tokenAid int64) ([]p.CGNFTDonation, error) {
	query := nftDonationsRoundFirstSQL + `
		WHERE d.token_aid=$1
		ORDER BY d.round_num DESC, d.id DESC`
	return queryList(ctx, r, "nft donations by token", 256, query, scanNFTDonationRoundFirst, tokenAid)
}

// UnclaimedDonatedNFTsByRound returns the donated NFTs of one round that
// have not been claimed yet, newest first.
func (r *Repo) UnclaimedDonatedNFTsByRound(ctx context.Context, roundNum int64) ([]p.CGNFTDonation, error) {
	query := nftDonationsRoundFirstSQL + `
			LEFT JOIN cg_donated_nft_claimed dc ON d.idx = dc.idx
		WHERE d.round_num= $1 AND dc.idx IS NULL
		ORDER BY d.id DESC`
	return queryList(ctx, r, "unclaimed donated nfts by round", 256, query, scanNFTDonationRoundFirst, roundNum)
}

// DonatedTokenDistribution returns how many NFTs each contract donated,
// most active contract first.
func (r *Repo) DonatedTokenDistribution(ctx context.Context) ([]p.CGDonatedTokenDistrRec, error) {
	query := `SELECT
			ca.addr,
			ns.num_donated
		FROM cg_nft_stats ns
			LEFT JOIN address ca ON ns.contract_aid=ca.address_id
		ORDER BY ns.num_donated DESC`
	scan := func(rows pgx.Rows, rec *p.CGDonatedTokenDistrRec) error {
		return rows.Scan(&rec.ContractAddr, &rec.NumDonatedTokens)
	}
	return queryList(ctx, r, "donated token distribution", 16, query, scan)
}

// NFTDonationsByUser returns every NFT donated by one donor, newest first.
func (r *Repo) NFTDonationsByUser(ctx context.Context, donorAid int64) ([]p.CGNFTDonation, error) {
	query := `SELECT
			d.id,
			d.evtlog_id,
			d.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,
			d.time_stamp,
			d.donor_aid,
			da.addr,
			d.token_id,
			d.round_num,
			nft.address_id,
			d.idx,
			nft.addr,
			d.token_uri
		FROM cg_nft_donation d
			LEFT JOIN transaction t ON t.id=tx_id
			LEFT JOIN address da ON d.donor_aid=da.address_id
			LEFT JOIN address nft ON d.token_aid=nft.address_id
		WHERE d.donor_aid=$1
		ORDER BY d.id DESC`
	return queryList(ctx, r, "nft donations by user", 256, query, scanNFTDonationDonorFirst, donorAid)
}
