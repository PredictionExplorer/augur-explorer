// Package archive contains the context-aware operational logic used by
// opsctl archive commands.
package archive

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// Querier is the narrow pgx query surface the archive operations use.
// *pgxpool.Pool, *pgx.Conn and pgx.Tx satisfy it.
type Querier interface {
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

// Project selector values accepted by the archive commands.
const (
	ProjectCosmicGame = "cosmicgame"
	ProjectRandomWalk = "randomwalk"
)

// Contracts is the ordered set of database address ids and chain addresses
// registered for a project.
type Contracts struct {
	AddressIDs []int64
	Addresses  []string
}

// ResolveProjects expands the command's project selector. The "both" order is
// intentional: CosmicGame is always processed before RandomWalk.
func ResolveProjects(project string) ([]string, error) {
	switch strings.ToLower(project) {
	case "both":
		return []string{ProjectCosmicGame, ProjectRandomWalk}, nil
	case ProjectCosmicGame, ProjectRandomWalk:
		return []string{strings.ToLower(project)}, nil
	default:
		return nil, fmt.Errorf("invalid project %q (want randomwalk, cosmicgame, or both)", project)
	}
}

// LoadProjectContracts resolves the registered contract address ids and
// addresses for project. Every query is bound to ctx.
func LoadProjectContracts(ctx context.Context, db Querier, project string) (Contracts, error) {
	var query string
	switch project {
	case ProjectRandomWalk:
		query = `
			SELECT DISTINCT a.address_id
			FROM address a
			JOIN rw_contracts rc ON a.addr = rc.randomwalk_addr OR a.addr = rc.marketplace_addr
			ORDER BY a.address_id`
	case ProjectCosmicGame:
		query = `
			SELECT DISTINCT a.address_id
			FROM address a
			JOIN cg_contracts cc ON
				a.addr = cc.cosmic_game_addr OR
				a.addr = cc.cosmic_signature_addr OR
				a.addr = cc.cosmic_token_addr OR
				a.addr = cc.cosmic_dao_addr OR
				a.addr = cc.charity_wallet_addr OR
				a.addr = cc.prizes_wallet_addr OR
				a.addr = cc.random_walk_addr OR
				a.addr = cc.staking_wallet_cst_addr OR
				a.addr = cc.staking_wallet_rwalk_addr OR
				a.addr = cc.marketing_wallet_addr OR
				a.addr = cc.implementation_addr
			ORDER BY a.address_id`
	default:
		return Contracts{}, fmt.Errorf("unknown project %q (want %s or %s)", project, ProjectCosmicGame, ProjectRandomWalk)
	}

	rows, err := db.Query(ctx, query)
	if err != nil {
		return Contracts{}, fmt.Errorf("contract aids: %w", err)
	}
	defer func(contractRows pgx.Rows) { contractRows.Close() }(rows)

	var contracts Contracts
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return Contracts{}, fmt.Errorf("scan contract aid: %w", err)
		}
		contracts.AddressIDs = append(contracts.AddressIDs, id)
	}
	if err := rows.Err(); err != nil {
		return Contracts{}, fmt.Errorf("contract aids: %w", err)
	}
	if len(contracts.AddressIDs) == 0 {
		return Contracts{}, fmt.Errorf("no contract addresses found for project %q", project)
	}

	rows, err = db.Query(ctx,
		`SELECT addr FROM address WHERE address_id = ANY($1) ORDER BY address_id`,
		contracts.AddressIDs,
	)
	if err != nil {
		return Contracts{}, fmt.Errorf("resolve addrs: %w", err)
	}
	defer func(addressRows pgx.Rows) { addressRows.Close() }(rows)
	for rows.Next() {
		var addr string
		if err := rows.Scan(&addr); err != nil {
			return Contracts{}, fmt.Errorf("scan contract address: %w", err)
		}
		contracts.Addresses = append(contracts.Addresses, addr)
	}
	if err := rows.Err(); err != nil {
		return Contracts{}, fmt.Errorf("resolve addrs: %w", err)
	}
	if len(contracts.Addresses) == 0 {
		return Contracts{}, fmt.Errorf("no contract addresses resolved for project %q", project)
	}
	return contracts, nil
}
