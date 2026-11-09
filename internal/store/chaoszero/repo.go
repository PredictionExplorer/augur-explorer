// Package chaoszero provides the ChaosZero domain's database access: a Repo
// of context-first, error-returning, pgx-native query methods over the
// shared store pool (cz-etl writes and the watermark/contract registries).
package chaoszero

import (
	"context"
	"errors"

	czmodel "github.com/PredictionExplorer/augur-explorer/internal/model/chaoszero"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// Repo executes ChaosZero queries on the shared store pool. Methods take a
// context as their first parameter, return errors and run natively on pgx.
type Repo struct {
	store *store.Store
}

// NewRepo returns a Repo backed by st's connection pool.
func NewRepo(st *store.Store) *Repo {
	return &Repo{store: st}
}

// q resolves the querier for ctx — the transaction the context carries when
// one was begun by this repo's Store (store.InTx), otherwise the shared
// pool. Every query method runs through it, which is what lets the indexer
// commit a block's domain writes atomically with its watermark.
func (r *Repo) q(ctx context.Context) store.Querier { return r.store.Querier(ctx) }

// addrID resolves addr to its address_id through the shared Store's
// lookup-or-create cache; the insert methods use it for every foreign-key
// address column.
func (r *Repo) addrID(ctx context.Context, addr string, blockNum, txID int64) (int64, error) {
	return r.store.LookupOrCreateAddress(ctx, addr, blockNum, txID)
}

// =============================================================================
// PROCESSING STATUS
// =============================================================================

// ProcessingStatus returns the ETL watermark (last processed event id and
// block number), lazily creating the singleton cz_proc_status row on a
// database that lost it (the migration seeds one whose last_block is the
// market deployment block).
func (r *Repo) ProcessingStatus(ctx context.Context) (czmodel.ProcStatus, error) {
	const op = "chaoszero processing status"
	var output czmodel.ProcStatus
	var lastID, lastBlock *int64
	err := r.q(ctx).QueryRow(ctx, "SELECT last_evt_id,last_block FROM cz_proc_status").Scan(&lastID, &lastBlock)
	if err != nil {
		wrapped := store.WrapError(op, err)
		if !errors.Is(wrapped, store.ErrNotFound) {
			return output, wrapped
		}
		// Missing singleton row: recreate it with the column defaults (the
		// deployment-block watermark) and report what it holds.
		if _, err := r.q(ctx).Exec(ctx, "INSERT INTO cz_proc_status DEFAULT VALUES"); err != nil {
			return output, store.WrapError(op+": insert default row", err)
		}
		if err := r.q(ctx).QueryRow(ctx, "SELECT last_evt_id,last_block FROM cz_proc_status").Scan(&lastID, &lastBlock); err != nil {
			return output, store.WrapError(op, err)
		}
	}
	if lastID != nil {
		output.LastIdProcessed = *lastID
	}
	if lastBlock != nil {
		output.LastBlockNum = *lastBlock
	}
	return output, nil
}

// UpdateProcessingStatus persists the ETL watermark.
func (r *Repo) UpdateProcessingStatus(ctx context.Context, status *czmodel.ProcStatus) error {
	_, err := r.q(ctx).Exec(ctx, "UPDATE cz_proc_status SET last_evt_id = $1,last_block=$2",
		status.LastIdProcessed, status.LastBlockNum)
	return store.WrapError("update chaoszero processing status", err)
}

// =============================================================================
// CONTRACT CONFIGURATION
// =============================================================================

// ContractAddrs returns the market, outcome source and settlement vault
// addresses with their address ids from the cz_contracts registry (one
// row). A database without the registry row (or with unregistered
// addresses) yields store.ErrNotFound.
func (r *Repo) ContractAddrs(ctx context.Context) (czmodel.ContractAddresses, error) {
	var output czmodel.ContractAddresses
	query := `SELECT
			market_addr,source_addr,vault_addr,
			m_a.address_id,s_a.address_id,v_a.address_id
		FROM cz_contracts cz
			JOIN address m_a ON cz.market_addr=m_a.addr
			JOIN address s_a ON cz.source_addr=s_a.addr
			JOIN address v_a ON cz.vault_addr=v_a.addr`
	err := r.q(ctx).QueryRow(ctx, query).Scan(
		&output.Market,
		&output.Source,
		&output.Vault,
		&output.MarketAid,
		&output.SourceAid,
		&output.VaultAid,
	)
	if err != nil {
		return output, store.WrapError("chaoszero contract addrs", err)
	}
	return output, nil
}

// RawContractAddrs returns the market, source and vault addresses straight
// from cz_contracts without joining the address table — the ETL uses it at
// startup to register the addresses before ContractAddrs can resolve them.
func (r *Repo) RawContractAddrs(ctx context.Context) (market, source, vault string, err error) {
	err = r.q(ctx).QueryRow(ctx,
		"SELECT market_addr, source_addr, vault_addr FROM cz_contracts LIMIT 1").Scan(&market, &source, &vault)
	if err != nil {
		return "", "", "", store.WrapError("raw chaoszero contract addrs", err)
	}
	return market, source, vault, nil
}
