package toolutil

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

const (
	ProjectCosmicGame = "cosmicgame"
	ProjectRandomWalk = "randomwalk"
)

// GetContractAddressIds returns address_id values for a project (cosmicgame or randomwalk).
func GetContractAddressIds(db *sql.DB, projectType string) ([]int64, error) {
	var query string
	switch projectType {
	case ProjectRandomWalk:
		query = `
			SELECT a.address_id
			FROM address a
			JOIN rw_contracts rc ON a.addr = rc.randomwalk_addr OR a.addr = rc.marketplace_addr
		`
	case ProjectCosmicGame:
		query = `
			SELECT a.address_id
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
		`
	default:
		return nil, fmt.Errorf("unknown project %q (want %s or %s)", projectType, ProjectCosmicGame, ProjectRandomWalk)
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var aids []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		aids = append(aids, id)
	}
	return aids, rows.Err()
}

// GetContractAddrsByAids resolves hex addresses for address_id values.
func GetContractAddrsByAids(db *sql.DB, contractAids []int64) ([]string, error) {
	rows, err := db.Query(
		`SELECT addr FROM address WHERE address_id = ANY($1) ORDER BY address_id`,
		pq.Array(contractAids),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addrs []string
	for rows.Next() {
		var addr string
		if err := rows.Scan(&addr); err != nil {
			return nil, err
		}
		addrs = append(addrs, addr)
	}
	return addrs, rows.Err()
}
