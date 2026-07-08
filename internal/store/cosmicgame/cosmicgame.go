package cosmicgame

import (
	"context"
	"errors"

	p "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// ContractAddrs returns the CosmicGame contract address registry
// (cg_contracts, one row) plus the RandomWalk marketplace address.
// A database without the registry row yields store.ErrNotFound.
func (r *Repo) ContractAddrs(ctx context.Context) (p.CosmicGameContractAddrs, error) {
	query := `SELECT
			cg.cosmic_game_addr,
			cg.cosmic_signature_addr,
			cg.cosmic_token_addr,
			cg.cosmic_dao_addr,
			cg.charity_wallet_addr,
			cg.prizes_wallet_addr,
			cg.random_walk_addr,
			cg.staking_wallet_cst_addr,
			cg.staking_wallet_rwalk_addr,
			cg.marketing_wallet_addr,
			COALESCE((SELECT marketplace_addr FROM rw_contracts LIMIT 1), '') AS marketplace_addr,
			cg.implementation_addr
		FROM cg_contracts cg`
	var out p.CosmicGameContractAddrs
	err := r.pool().QueryRow(ctx, query).Scan(
		&out.CosmicGameAddr,
		&out.CosmicSignatureAddr,
		&out.CosmicTokenAddr,
		&out.CosmicDaoAddr,
		&out.CharityWalletAddr,
		&out.PrizesWalletAddr,
		&out.RandomWalkAddr,
		&out.StakingWalletCSTAddr,
		&out.StakingWalletRWalkAddr,
		&out.MarketingWalletAddr,
		&out.MarketplaceAddr,
		&out.ImplementationAddr,
	)
	if err != nil {
		return p.CosmicGameContractAddrs{}, store.WrapError("cosmic game contract addrs", err)
	}
	return out, nil
}

// ProcessingStatus returns the ETL watermark (last processed event id and
// block number), lazily creating the singleton cg_proc_status row on a fresh
// database.
func (r *Repo) ProcessingStatus(ctx context.Context) (p.CosmicGameProcStatus, error) {
	const op = "cosmic game processing status"
	var out p.CosmicGameProcStatus
	var lastEvtID, lastBlock *int64
	err := r.pool().QueryRow(ctx, "SELECT last_evt_id, last_block_num FROM cg_proc_status").Scan(&lastEvtID, &lastBlock)
	if err != nil {
		wrapped := store.WrapError(op, err)
		if !errors.Is(wrapped, store.ErrNotFound) {
			return out, wrapped
		}
		// Fresh database: create the singleton row and report the zero
		// watermark it holds.
		if _, err := r.pool().Exec(ctx, "INSERT INTO cg_proc_status DEFAULT VALUES"); err != nil {
			return out, store.WrapError(op+": insert default row", err)
		}
		if err := r.pool().QueryRow(ctx, "SELECT last_evt_id, last_block_num FROM cg_proc_status").Scan(&lastEvtID, &lastBlock); err != nil {
			return out, store.WrapError(op, err)
		}
	}
	if lastEvtID != nil {
		out.LastEvtIdProcessed = *lastEvtID
	}
	if lastBlock != nil {
		out.LastBlockNum = *lastBlock
	}
	return out, nil
}

// UpdateProcessingStatus persists the ETL watermark.
func (r *Repo) UpdateProcessingStatus(ctx context.Context, status *p.CosmicGameProcStatus) error {
	query := "UPDATE cg_proc_status SET last_evt_id = $1, last_block_num = $2"
	_, err := r.pool().Exec(ctx, query, status.LastEvtIdProcessed, status.LastBlockNum)
	return store.WrapError("update cosmic game processing status", err)
}
