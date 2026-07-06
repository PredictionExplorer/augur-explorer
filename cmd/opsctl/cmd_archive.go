package main

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/PredictionExplorer/augur-explorer/internal/toolutil"
	"github.com/spf13/cobra"
)

// archiveCmd groups the tools that maintain the arch_evtlog / arch_tx /
// arch_block archive tables: export from a live database, consistency
// verification, and backfill from an Ethereum node.
var archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Event-log archival tools (arch_evtlog / arch_tx / arch_block)",
}

func init() { register(archiveCmd) }

// resolveProjects expands a --project flag value into the ordered list of
// projects to process. "both" runs cosmicgame first, then randomwalk; either
// order is safe because every archive tool resumes per project.
func resolveProjects(projectType string) ([]string, error) {
	switch strings.ToLower(projectType) {
	case "both":
		return []string{toolutil.ProjectCosmicGame, toolutil.ProjectRandomWalk}, nil
	case toolutil.ProjectRandomWalk, toolutil.ProjectCosmicGame:
		return []string{strings.ToLower(projectType)}, nil
	default:
		return nil, fmt.Errorf("invalid project %q (want randomwalk, cosmicgame, or both)", projectType)
	}
}

// projectContracts resolves the registered contract address ids and hex
// addresses for a project (rw_contracts / cg_contracts tables).
func projectContracts(db *sql.DB, project string) (aids []int64, addrs []string, err error) {
	aids, err = toolutil.GetContractAddressIds(db, project)
	if err != nil {
		return nil, nil, fmt.Errorf("contract aids: %w", err)
	}
	if len(aids) == 0 {
		return nil, nil, fmt.Errorf("no contract addresses found for project %q", project)
	}
	addrs, err = toolutil.GetContractAddrsByAids(db, aids)
	if err != nil {
		return nil, nil, fmt.Errorf("resolve addrs: %w", err)
	}
	if len(addrs) == 0 {
		return nil, nil, fmt.Errorf("no contract addresses resolved for project %q", project)
	}
	return aids, addrs, nil
}
