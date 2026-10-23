package main

// opsctl validate-chain -- chain-vs-database audit for the Cosmic Signature
// backend (the Go successor of scripts/validate-chain.py).
//
// The SQL validators (validate-db.sql, validate-stats.sql,
// validate-durations.sql) prove the database is INTERNALLY consistent. This
// command closes the remaining gap: it reconciles the database against live
// Arbitrum contract state.
//
// Checks:
//   - log parity (OFF by default; event completeness is normally covered by
//     the tx-collector backup plus `opsctl tx-collector verify`):
//     --archive-dir matches per-transaction per-contract log counts against a
//     tx-collector backup tree in both directions, locally and fast;
//     --full-log-parity scans the chain with eth_getLogs instead (slow; the
//     scan is checkpointed to --state-file and resumes incrementally).
//   - CST totalSupply() vs mints-minus-burns in cg_erc20_transfer.
//   - CS NFT totalSupply() and per-token ownerOf() vs cg_mint_event.
//   - numStakedNfts() on both staking wallets vs stake/unstake events.
//   - PrizesWallet ETH balance vs deposits minus withdrawals.
//   - V3 roundStats(r) vs cg_bid recomputation: numBids for every round (the
//     V2->V3 storage reinterpretation makes it valid for legacy rounds too);
//     the remaining fields only for rounds started under V3.
//
// Live-chain races: bids can land between a chain read and a DB read, so
// volatile checks are retried a few times and pass if chain and DB agree in
// any attempt.
//
// Configuration: RPC_URL and PGSQL_*/DATABASE_URL from the environment (the
// usual opsctl conventions); flags default from ARCHIVE_DIR, ARCHIVE_CONFIG,
// FULL_LOG_PARITY, VALIDATE_CHAIN_STATE, LOG_WORKERS and LOG_CHUNK for
// compatibility with the retired Python script, whose JSON state-file format
// is reused as-is so existing scan checkpoints carry over.

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/cobra"

	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/toolutil"
)

// ErrChainValidationFailed is returned when at least one check finds a mismatch.
var ErrChainValidationFailed = errors.New("chain validation failed")

const (
	chainValidateVolatileRetries = 4
	chainValidateVolatileDelay   = 2 * time.Second
	// getLogs bloom scans over millions of blocks are slow; one long attempt
	// instead of stacking retried queries onto the node.
	chainValidateLogsTimeout = 15 * time.Minute
)

func defaultChainStateFile() string {
	if p := os.Getenv("VALIDATE_CHAIN_STATE"); p != "" {
		return p
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "validate-chain-state.json"
	}
	return filepath.Join(home, ".cache", "cosmic-signature", "validate-chain-state.json")
}

func envUint(name string, def uint64) uint64 {
	if v := os.Getenv(name); v != "" {
		if n, err := strconv.ParseUint(v, 10, 64); err == nil && n > 0 {
			return n
		}
	}
	return def
}

func newValidateChainCmd() *cobra.Command {
	var (
		archiveDir    string
		archiveConfig string
		fullLogParity bool
		stateFile     string
		logWorkers    uint64
		logChunk      uint64
	)
	cmd := &cobra.Command{
		Use:   "validate-chain",
		Short: "Reconcile the SQL database against live contract state",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			rpcURL := os.Getenv("RPC_URL")
			if rpcURL == "" {
				return errors.New("RPC_URL must be set (Arbitrum RPC endpoint)")
			}
			dbConn, err := toolutil.PostgresConnStringFromEnv()
			if err != nil {
				return fmt.Errorf("database config: %w", err)
			}
			ctx := cmd.Context()
			pool, err := pgxpool.New(ctx, dbConn)
			if err != nil {
				return fmt.Errorf("database connect: %w", err)
			}
			defer pool.Close()
			client, err := dialBoundedEthClient(ctx, rpcURL)
			if err != nil {
				return fmt.Errorf("rpc connect: %w", err)
			}
			defer client.Close()

			v := &chainValidator{
				cmd:           cmd,
				ctx:           ctx,
				client:        client,
				pool:          pool,
				rpcURL:        rpcURL,
				archiveDir:    archiveDir,
				archiveConfig: archiveConfig,
				fullLogParity: fullLogParity,
				stateFile:     stateFile,
				logWorkers:    int(logWorkers),
				logChunk:      logChunk,
			}
			return v.run()
		},
	}
	cmd.Flags().StringVar(&archiveDir, "archive-dir", os.Getenv("ARCHIVE_DIR"),
		"tx-collector backup tree; enables the local log-parity match")
	cmd.Flags().StringVar(&archiveConfig, "archive-config", os.Getenv("ARCHIVE_CONFIG"),
		"tx-collector JSON config naming the contracts the archive covers (required with --archive-dir)")
	cmd.Flags().BoolVar(&fullLogParity, "full-log-parity", os.Getenv("FULL_LOG_PARITY") != "",
		"enable the slow eth_getLogs log-parity network scan")
	cmd.Flags().StringVar(&stateFile, "state-file", defaultChainStateFile(),
		"checkpoint file for the network log scan")
	cmd.Flags().Uint64Var(&logWorkers, "log-workers", envUint("LOG_WORKERS", 4),
		"concurrent eth_getLogs requests for the network scan")
	cmd.Flags().Uint64Var(&logChunk, "log-chunk", envUint("LOG_CHUNK", 2_000_000),
		"eth_getLogs block span per request")
	return cmd
}

// chainValidator carries the shared handles and accumulates check results.
type chainValidator struct {
	cmd    *cobra.Command
	ctx    context.Context
	client *ethclient.Client
	pool   *pgxpool.Pool
	rpcURL string

	archiveDir    string
	archiveConfig string
	fullLogParity bool
	stateFile     string
	logWorkers    int
	logChunk      uint64

	checks, fails, warns int
}

func (v *chainValidator) add(status, name, details string) {
	v.checks++
	switch status {
	case "FAIL":
		v.fails++
	case "WARN":
		v.warns++
	}
	line := fmt.Sprintf(" %-5s| %s", status, name)
	if details != "" {
		line += " | " + details
	}
	fmt.Fprintln(v.cmd.OutOrStdout(), line)
}

func (v *chainValidator) progressf(format string, args ...any) {
	fmt.Fprintf(v.cmd.ErrOrStderr(), format+"\n", args...)
}

// scanBigInt runs a single-value query and parses the result as a big.Int.
// Cast DECIMAL sums to ::NUMERIC(80,0)::TEXT in the SQL; counts work as-is.
func (v *chainValidator) scanBigInt(sql string, args ...any) (*big.Int, error) {
	var raw string
	if err := v.pool.QueryRow(v.ctx, sql, args...).Scan(&raw); err != nil {
		return nil, fmt.Errorf("query %q: %w", sql, err)
	}
	n, ok := new(big.Int).SetString(strings.TrimSpace(raw), 10)
	if !ok {
		return nil, fmt.Errorf("query %q: non-numeric result %q", sql, raw)
	}
	return n, nil
}

// volatileMatch compares live chain values with DB values, retrying through
// ETL tailing lag. Both callbacks return a value vector; equal vectors in any
// attempt is a pass. The last attempt's vectors are returned for reporting.
func (v *chainValidator) volatileMatch(
	chainFn, dbFn func() ([]*big.Int, error),
) (ok bool, chainVals, dbVals []*big.Int, err error) {
	for attempt := 0; attempt < chainValidateVolatileRetries; attempt++ {
		if attempt > 0 {
			select {
			case <-v.ctx.Done():
				return false, nil, nil, v.ctx.Err()
			case <-time.After(chainValidateVolatileDelay):
			}
		}
		if chainVals, err = chainFn(); err != nil {
			return false, nil, nil, err
		}
		if dbVals, err = dbFn(); err != nil {
			return false, nil, nil, err
		}
		if bigIntSliceEqual(chainVals, dbVals) {
			return true, chainVals, dbVals, nil
		}
	}
	return false, chainVals, dbVals, nil
}

func bigIntSliceEqual(a, b []*big.Int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Cmp(b[i]) != 0 {
			return false
		}
	}
	return true
}

func one(n *big.Int, err error) func() ([]*big.Int, error) {
	return func() ([]*big.Int, error) { return []*big.Int{n}, err }
}

// labeledContract is one tracked emitter: the registry contracts plus every
// implementation the proxy has pointed at.
type labeledContract struct {
	name string
	addr string // lowercase 0x hex
}

func (v *chainValidator) run() error {
	fmt.Fprintf(v.cmd.OutOrStdout(), "=== Cosmic Signature chain-vs-DB validator ===\n\n")

	// -- contract addresses --------------------------------------------------
	var row [11]string
	if err := v.pool.QueryRow(v.ctx,
		`SELECT cosmic_game_addr, cosmic_signature_addr, cosmic_token_addr,
		        cosmic_dao_addr, charity_wallet_addr, prizes_wallet_addr,
		        random_walk_addr, staking_wallet_cst_addr,
		        staking_wallet_rwalk_addr, marketing_wallet_addr,
		        implementation_addr
		 FROM cg_contracts`,
	).Scan(&row[0], &row[1], &row[2], &row[3], &row[4], &row[5], &row[6],
		&row[7], &row[8], &row[9], &row[10]); err != nil {
		return fmt.Errorf("cg_contracts: %w", err)
	}
	for i := range row {
		row[i] = strings.ToLower(row[i])
	}
	game, nft, cst, dao := row[0], row[1], row[2], row[3]
	charity, prizes, rwalk := row[4], row[5], row[6]
	stakingCst, stakingRwalk, marketing, implV1 := row[7], row[8], row[9], row[10]

	labeled := []labeledContract{
		{"game proxy", game}, {"CosmicSignatureNft", nft},
		{"CosmicSignatureToken", cst}, {"CosmicSignatureDao", dao},
		{"CharityWallet", charity}, {"PrizesWallet", prizes},
		{"RandomWalkNFT", rwalk},
		{"StakingWallet CST", stakingCst},
		{"StakingWallet RWalk", stakingRwalk},
		{"MarketingWallet", marketing},
	}
	impls := map[string]struct{}{implV1: {}}
	rows, err := v.pool.Query(v.ctx,
		`SELECT DISTINCT LOWER(a.addr) FROM cg_adm_upgraded u
		 JOIN address a ON a.address_id = u.implementation_aid`)
	if err != nil {
		return fmt.Errorf("cg_adm_upgraded: %w", err)
	}
	for rows.Next() {
		var a string
		if err := rows.Scan(&a); err != nil {
			rows.Close()
			return err
		}
		impls[a] = struct{}{}
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return err
	}
	seen := map[string]struct{}{}
	for _, lc := range labeled {
		seen[lc.addr] = struct{}{}
	}
	implList := make([]string, 0, len(impls))
	for a := range impls {
		implList = append(implList, a)
	}
	sort.Strings(implList)
	for _, a := range implList {
		if _, dup := seen[a]; !dup {
			labeled = append(labeled, labeledContract{"game impl " + a[:10], a})
			seen[a] = struct{}{}
		}
	}

	// -- DB head: the upper bound for log scans and DB-side counts. Not a
	// health check (monitor cg_proc_status directly for indexer lag); the
	// GREATEST covers databases where the checkpoint was never initialized,
	// falling back to the newest block that contained a tracked event.
	dbHeadBig, err := v.scanBigInt(
		`SELECT GREATEST((SELECT COALESCE(MAX(block_num), 0) FROM block),
		                 (SELECT COALESCE(MAX(last_block_num), 0) FROM cg_proc_status))::TEXT`)
	if err != nil {
		return err
	}
	dbHead := dbHeadBig.Uint64()

	// -- log parity -----------------------------------------------------------
	switch {
	case v.archiveDir != "":
		if err := v.archiveLogParity(labeled, dbHead); err != nil {
			return err
		}
	case v.fullLogParity:
		if err := v.networkLogParity(labeled, rwalk, dbHead); err != nil {
			return err
		}
	default:
		fmt.Fprintln(v.cmd.OutOrStdout(),
			"(log parity skipped; pass --archive-dir <tx-collector backup dir>"+
				" or --full-log-parity to enable)")
	}

	callOpts := func() *bind.CallOpts { return &bind.CallOpts{Context: v.ctx} }

	// -- CST total supply -----------------------------------------------------
	cstCaller, err := cgcontracts.NewCosmicSignatureTokenCaller(
		ethcommon.HexToAddress(cst), v.client)
	if err != nil {
		return err
	}
	ok, cv, dv, err := v.volatileMatch(
		func() ([]*big.Int, error) {
			return wrapOne(cstCaller.TotalSupply(callOpts()))
		},
		func() ([]*big.Int, error) {
			return wrapOne(v.scanBigInt(
				`SELECT COALESCE(SUM(CASE otype WHEN 1 THEN value
				                      WHEN 2 THEN -value ELSE 0 END), 0)::NUMERIC(80,0)::TEXT
				 FROM cg_erc20_transfer e
				 JOIN address a ON a.address_id = e.contract_aid
				 WHERE LOWER(a.addr) = $1`, cst))
		})
	if err != nil {
		return err
	}
	v.addMatch(ok, "CST totalSupply vs mints-burns",
		fmt.Sprintf("chain=%s db=%s", cv[0], dv[0]))

	// -- NFT supply and per-token ownership ------------------------------------
	nftCaller, err := cgcontracts.NewCosmicSignatureNftCaller(
		ethcommon.HexToAddress(nft), v.client)
	if err != nil {
		return err
	}
	ok, cv, dv, err = v.volatileMatch(
		func() ([]*big.Int, error) { return wrapOne(nftCaller.TotalSupply(callOpts())) },
		func() ([]*big.Int, error) {
			return wrapOne(v.scanBigInt(`SELECT COUNT(*)::TEXT FROM cg_mint_event`))
		})
	if err != nil {
		return err
	}
	v.addMatch(ok, "CS NFT totalSupply vs cg_mint_event count",
		fmt.Sprintf("chain=%s db=%s", cv[0], dv[0]))

	type ownedToken struct {
		id    int64
		owner string
	}
	var owners []ownedToken
	rows, err = v.pool.Query(v.ctx,
		`SELECT m.token_id, LOWER(a.addr) FROM cg_mint_event m
		 JOIN address a ON a.address_id = m.cur_owner_aid ORDER BY m.token_id`)
	if err != nil {
		return err
	}
	for rows.Next() {
		var t ownedToken
		if err := rows.Scan(&t.id, &t.owner); err != nil {
			rows.Close()
			return err
		}
		owners = append(owners, t)
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return err
	}
	var ownerMismatches []string
	for _, t := range owners {
		chainOwner, err := nftCaller.OwnerOf(callOpts(), big.NewInt(t.id))
		if err != nil {
			return fmt.Errorf("ownerOf(%d): %w", t.id, err)
		}
		if strings.ToLower(chainOwner.Hex()) != t.owner {
			ownerMismatches = append(ownerMismatches, fmt.Sprintf(
				"token %d: chain=%s db=%s", t.id,
				strings.ToLower(chainOwner.Hex()), t.owner))
		}
	}
	detail := ""
	if len(ownerMismatches) > 0 {
		detail = strings.Join(truncateList(ownerMismatches, 5), "; ")
	}
	v.addMatch(len(ownerMismatches) == 0,
		fmt.Sprintf("ownerOf() vs cur_owner for %d NFTs", len(owners)), detail)

	// -- staked-NFT counts ------------------------------------------------------
	type stakingCheck struct {
		name                  string
		wallet                string
		stakedTbl, unstakeTbl string
	}
	for _, sc := range []stakingCheck{
		{"CST", stakingCst, "cg_nft_staked_cst", "cg_nft_unstaked_cst"},
		{"RandomWalk", stakingRwalk, "cg_nft_staked_rwalk", "cg_nft_unstaked_rwalk"},
	} {
		caller, err := cgcontracts.NewStakingWalletNftBaseCaller(
			ethcommon.HexToAddress(sc.wallet), v.client)
		if err != nil {
			return err
		}
		// Table names come from the fixed list above, never user input.
		query := fmt.Sprintf(
			`SELECT ((SELECT COUNT(*) FROM %s) - (SELECT COUNT(*) FROM %s))::TEXT`,
			sc.stakedTbl, sc.unstakeTbl)
		ok, cv, dv, err = v.volatileMatch(
			func() ([]*big.Int, error) { return wrapOne(caller.NumStakedNfts(callOpts())) },
			func() ([]*big.Int, error) { return wrapOne(v.scanBigInt(query)) })
		if err != nil {
			return err
		}
		v.addMatch(ok, fmt.Sprintf("numStakedNfts (%s) vs stake-unstake events", sc.name),
			fmt.Sprintf("chain=%s db=%s", cv[0], dv[0]))
	}

	// -- PrizesWallet balance ----------------------------------------------------
	prizesAddr := ethcommon.HexToAddress(prizes)
	ok, cv, dv, err = v.volatileMatch(
		func() ([]*big.Int, error) {
			return wrapOne(v.client.BalanceAt(v.ctx, prizesAddr, nil))
		},
		func() ([]*big.Int, error) {
			return wrapOne(v.scanBigInt(
				`SELECT ((SELECT COALESCE(SUM(amount), 0) FROM cg_prize_deposit)
				       - (SELECT COALESCE(SUM(amount), 0) FROM cg_prize_withdrawal))::NUMERIC(80,0)::TEXT`))
		})
	if err != nil {
		return err
	}
	v.addMatch(ok, "PrizesWallet ETH balance vs deposits-withdrawals",
		fmt.Sprintf("chain=%s db=%s wei", cv[0], dv[0]))

	// -- roundStats(r) vs cg_bid recomputation -------------------------------------
	if err := v.checkRoundStats(game); err != nil {
		return err
	}

	// -- summary ---------------------------------------------------------------
	fmt.Fprintf(v.cmd.OutOrStdout(),
		"\nchecks: %d  passed: %d  failed: %d  warnings: %d\n",
		v.checks, v.checks-v.fails-v.warns, v.fails, v.warns)
	if v.fails > 0 {
		fmt.Fprintln(v.cmd.OutOrStdout(),
			"CHAIN VALIDATION FAILED: at least one check found a mismatch.")
		return ErrChainValidationFailed
	}
	fmt.Fprintln(v.cmd.OutOrStdout(),
		"CHAIN VALIDATION OK: database matches on-chain state.")
	return nil
}

func (v *chainValidator) addMatch(ok bool, name, details string) {
	status := "PASS"
	if !ok {
		status = "FAIL"
	}
	v.add(status, name, details)
}

func wrapOne(n *big.Int, err error) ([]*big.Int, error) {
	if err != nil {
		return nil, err
	}
	return []*big.Int{n}, nil
}

func truncateList(items []string, n int) []string {
	if len(items) <= n {
		return items
	}
	out := append([]string{}, items[:n]...)
	return append(out, "...")
}

// --------------------------------------------------------------------------
// roundStats
// --------------------------------------------------------------------------
var roundStatsFieldNames = []string{
	"numBids", "numCstBids", "totalSpentEth", "totalSpentCst",
	"maxEthBidPrice", "maxCstBidPrice",
}

func (v *chainValidator) checkRoundStats(game string) error {
	gameCaller, err := cgcontracts.NewCosmicSignatureGameV3Caller(
		ethcommon.HexToAddress(game), v.client)
	if err != nil {
		return err
	}
	callOpts := &bind.CallOpts{Context: v.ctx}
	curRound, err := gameCaller.RoundNum(callOpts)
	if err != nil {
		return fmt.Errorf("roundNum(): %w", err)
	}

	// The V3 reinitialize block: rounds whose first bid landed at or after it
	// carry full on-chain statistics; earlier rounds only numBids (via the
	// V2->V3 storage reinterpretation).
	var v3Block *int64
	if err := v.pool.QueryRow(v.ctx,
		`SELECT MIN(block_num) FROM cg_adm_initialized WHERE version = 3`,
	).Scan(&v3Block); err != nil {
		return fmt.Errorf("cg_adm_initialized: %w", err)
	}

	for r := int64(0); r <= curRound.Int64(); r++ {
		var firstBidBlock *int64
		if err := v.pool.QueryRow(v.ctx,
			`SELECT MIN(block_num) FROM cg_bid WHERE round_num = $1`, r,
		).Scan(&firstBidBlock); err != nil {
			return err
		}
		v3Round := v3Block != nil && firstBidBlock != nil && *firstBidBlock >= *v3Block
		nFields := 1
		if v3Round {
			nFields = len(roundStatsFieldNames)
		}

		chainFn := func() ([]*big.Int, error) {
			s, err := gameCaller.RoundStats(callOpts, big.NewInt(r))
			if err != nil {
				return nil, fmt.Errorf("roundStats(%d): %w", r, err)
			}
			all := []*big.Int{s.NumBids, s.NumCstBids, s.TotalSpentEthAmount,
				s.TotalSpentCstAmount, s.MaxEthBidPrice, s.MaxCstBidPrice}
			return all[:nFields], nil
		}
		dbFn := func() ([]*big.Int, error) {
			// eth_price is -1 for CST bids, so ETH aggregates exclude them.
			var raw [6]string
			if err := v.pool.QueryRow(v.ctx,
				`SELECT COUNT(*)::TEXT,
				        (COUNT(*) FILTER (WHERE bid_type = 2))::TEXT,
				        COALESCE(SUM(eth_price) FILTER (WHERE bid_type <> 2), 0)::NUMERIC(80,0)::TEXT,
				        COALESCE(SUM(cst_price) FILTER (WHERE bid_type = 2), 0)::NUMERIC(80,0)::TEXT,
				        COALESCE(MAX(eth_price) FILTER (WHERE bid_type <> 2), 0)::NUMERIC(80,0)::TEXT,
				        COALESCE(MAX(cst_price) FILTER (WHERE bid_type = 2), 0)::NUMERIC(80,0)::TEXT
				 FROM cg_bid WHERE round_num = $1`, r,
			).Scan(&raw[0], &raw[1], &raw[2], &raw[3], &raw[4], &raw[5]); err != nil {
				return nil, err
			}
			out := make([]*big.Int, nFields)
			for i := 0; i < nFields; i++ {
				n, okNum := new(big.Int).SetString(raw[i], 10)
				if !okNum {
					return nil, fmt.Errorf("round %d: non-numeric %q", r, raw[i])
				}
				out[i] = n
			}
			return out, nil
		}

		ok, cv, dv, err := v.volatileMatch(chainFn, dbFn)
		if err != nil {
			return err
		}
		scope := "legacy numBids-only"
		if v3Round {
			scope = "V3 full"
		}
		detail := fmt.Sprintf("numBids=%s", cv[0])
		if !ok {
			var diffs []string
			for i := range cv {
				if cv[i].Cmp(dv[i]) != 0 {
					diffs = append(diffs, fmt.Sprintf("%s: chain=%s db=%s",
						roundStatsFieldNames[i], cv[i], dv[i]))
				}
			}
			detail = strings.Join(diffs, "; ")
		}
		v.addMatch(ok, fmt.Sprintf("roundStats(%d) vs cg_bid (%s)", r, scope), detail)
	}
	return nil
}

// --------------------------------------------------------------------------
// Log parity: archive mode
// --------------------------------------------------------------------------
type txAddrKey struct {
	txHash string // lowercase, 0x-prefixed
	addr   string // lowercase, 0x-prefixed
}

func (v *chainValidator) archiveLogParity(labeled []labeledContract, dbHead uint64) error {
	if v.archiveConfig == "" {
		return errors.New("--archive-config (collector JSON) is required with --archive-dir")
	}
	cfg, err := toolutil.LoadCollectorConfig(v.archiveConfig)
	if err != nil {
		return fmt.Errorf("archive config: %w", err)
	}
	scopeAddrs, err := cfg.ResolveContractAddresses()
	if err != nil {
		return fmt.Errorf("archive config contracts: %w", err)
	}
	scope := map[string]struct{}{}
	for _, a := range scopeAddrs {
		scope[strings.ToLower(a)] = struct{}{}
	}
	tracked := map[string]struct{}{}
	for _, lc := range labeled {
		tracked[lc.addr] = struct{}{}
	}

	arch, nReceipts, decodeErrs, err := archiveLogCounts(v.archiveDir, tracked, dbHead)
	if err != nil {
		return fmt.Errorf("archive walk: %w", err)
	}
	v.addMatch(len(decodeErrs) == 0,
		fmt.Sprintf("archive readable (%d receipts <= block %d)", nReceipts, dbHead),
		strings.Join(truncateList(decodeErrs, 3), "; "))

	trackedList := make([]string, 0, len(tracked))
	for a := range tracked {
		trackedList = append(trackedList, a)
	}
	dbGroups := map[txAddrKey]uint64{}
	rows, err := v.pool.Query(v.ctx,
		`SELECT LOWER(t.tx_hash), LOWER(a.addr), COUNT(*)
		 FROM evt_log e
		 JOIN transaction t ON t.id = e.tx_id
		 JOIN address a ON a.address_id = e.contract_aid
		 WHERE LOWER(a.addr) = ANY($1) AND e.block_num <= $2
		 GROUP BY 1, 2`, trackedList, int64(dbHead)) // #nosec G115 -- block numbers fit int64
	if err != nil {
		return err
	}
	for rows.Next() {
		var k txAddrKey
		var n uint64
		if err := rows.Scan(&k.txHash, &k.addr, &n); err != nil {
			rows.Close()
			return err
		}
		dbGroups[k] = n
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return err
	}

	for _, lc := range labeled {
		if _, inScope := scope[lc.addr]; !inScope {
			v.add("WARN", "log parity (archive): "+lc.name,
				"not in collector config; not covered by the archive")
			continue
		}
		txs := map[string]struct{}{}
		for k := range arch {
			if k.addr == lc.addr {
				txs[k.txHash] = struct{}{}
			}
		}
		for k := range dbGroups {
			if k.addr == lc.addr {
				txs[k.txHash] = struct{}{}
			}
		}
		txList := make([]string, 0, len(txs))
		for tx := range txs {
			txList = append(txList, tx)
		}
		sort.Strings(txList)
		var aTotal, dTotal uint64
		var diffs []string
		for _, tx := range txList {
			an := arch[txAddrKey{tx, lc.addr}]
			dn := dbGroups[txAddrKey{tx, lc.addr}]
			aTotal += an
			dTotal += dn
			if an != dn {
				diffs = append(diffs, fmt.Sprintf("tx %s: archive=%d db=%d", tx, an, dn))
			}
		}
		detail := fmt.Sprintf("archive=%d db=%d (..%d)", aTotal, dTotal, dbHead)
		if len(diffs) > 0 {
			detail += "; " + strings.Join(truncateList(diffs, 5), "; ")
		}
		v.addMatch(len(diffs) == 0, "log parity (archive): "+lc.name, detail)
	}
	return nil
}

// archiveLogCounts walks a tx-collector backup tree and counts logs per
// (tx, contract) for tracked contracts, limited to blocks <= dbHead.
func archiveLogCounts(
	dir string, tracked map[string]struct{}, dbHead uint64,
) (counts map[txAddrKey]uint64, nReceipts int, decodeErrs []string, err error) {
	counts = map[txAddrKey]uint64{}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, 0, nil, err
	}
	for _, blockDir := range entries {
		if !blockDir.IsDir() {
			continue
		}
		blockNum, convErr := strconv.ParseUint(blockDir.Name(), 10, 64)
		if convErr != nil || blockNum > dbHead {
			continue
		}
		files, err := os.ReadDir(filepath.Join(dir, blockDir.Name()))
		if err != nil {
			return nil, 0, nil, err
		}
		for _, f := range files {
			const suffix = "_receipt.rlp"
			if f.IsDir() || !strings.HasSuffix(f.Name(), suffix) {
				continue
			}
			nReceipts++
			txHash := strings.ToLower(strings.TrimSuffix(f.Name(), suffix))
			path := filepath.Join(dir, blockDir.Name(), f.Name())
			blob, readErr := os.ReadFile(path) // #nosec G304 -- operator-supplied backup tree
			if readErr != nil {
				decodeErrs = append(decodeErrs, fmt.Sprintf("%s: %v", path, readErr))
				continue
			}
			br, _, decErr := toolutil.TryDecodeReceiptRLP(blob)
			if decErr != nil {
				decodeErrs = append(decodeErrs, fmt.Sprintf("%s: %v", path, decErr))
				continue
			}
			for i := range br.Logs {
				addr := strings.ToLower(br.Logs[i].Address.Hex())
				if _, okT := tracked[addr]; okT {
					counts[txAddrKey{txHash, addr}]++
				}
			}
		}
	}
	return counts, nReceipts, decodeErrs, nil
}

// --------------------------------------------------------------------------
// Log parity: network scan with checkpointing
// --------------------------------------------------------------------------

// chainScanState is the on-disk checkpoint, format-compatible with the
// retired Python script so existing scan progress carries over.
type chainScanState map[string]*chainScanEntry

type chainScanEntry struct {
	From      uint64            `json:"from"`
	ScannedTo uint64            `json:"scanned_to"`
	Counts    map[string]uint64 `json:"counts"`
}

func loadChainScanState(path string) chainScanState {
	data, err := os.ReadFile(path) // #nosec G304 -- operator-supplied state path
	if err != nil {
		return chainScanState{}
	}
	var s chainScanState
	if json.Unmarshal(data, &s) != nil {
		return chainScanState{}
	}
	return s
}

func saveChainScanState(path string, s chainScanState) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o750); err != nil {
		return err
	}
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, data, 0o600); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}

func (v *chainValidator) networkLogParity(
	labeled []labeledContract, rwalk string, dbHead uint64,
) error {
	chainID, err := v.client.ChainID(v.ctx)
	if err != nil {
		return fmt.Errorf("eth_chainId: %w", err)
	}

	// A dedicated client with a long per-request timeout: bloom scans over
	// millions of blocks are slow, and stacked retries would pile concurrent
	// scans onto the node.
	logsRPC, err := ethrpc.DialOptions(v.ctx, v.rpcURL,
		ethrpc.WithHTTPClient(&http.Client{Timeout: chainValidateLogsTimeout}))
	if err != nil {
		return fmt.Errorf("rpc connect (logs): %w", err)
	}
	logsClient := ethclient.NewClient(logsRPC)
	defer logsClient.Close()

	var nonRwalk []string
	addrSet := make([]ethcommon.Address, 0, len(labeled))
	for _, lc := range labeled {
		addrSet = append(addrSet, ethcommon.HexToAddress(lc.addr))
		if lc.addr != rwalk {
			nonRwalk = append(nonRwalk, lc.addr)
		}
	}
	cosmicFromBig, err := v.scanBigInt(
		`SELECT COALESCE(MIN(e.block_num), 0)::TEXT FROM evt_log e
		 JOIN address a ON a.address_id = e.contract_aid
		 WHERE LOWER(a.addr) = ANY($1)`, nonRwalk)
	if err != nil {
		return err
	}
	cosmicFrom := cosmicFromBig.Uint64()

	state := loadChainScanState(v.stateFile)
	counts, err := v.scanLogsCheckpointed(logsClient, state,
		fmt.Sprintf("%d:platform", chainID), addrSet, cosmicFrom, dbHead)
	if err != nil {
		return err
	}

	var rwFrom *int64
	if err := v.pool.QueryRow(v.ctx,
		`SELECT MIN(e.block_num) FROM evt_log e
		 JOIN address a ON a.address_id = e.contract_aid
		 WHERE LOWER(a.addr) = $1`, rwalk).Scan(&rwFrom); err != nil {
		return err
	}
	var rwPrefix map[string]uint64
	rwNeedsPrefix := rwFrom != nil && uint64(*rwFrom) < cosmicFrom // #nosec G115 -- block numbers are non-negative
	if rwNeedsPrefix {
		rwPrefix, err = v.scanLogsCheckpointed(logsClient, state,
			fmt.Sprintf("%d:randomwalk-prefix", chainID),
			[]ethcommon.Address{ethcommon.HexToAddress(rwalk)},
			uint64(*rwFrom), cosmicFrom-1) // #nosec G115 -- block numbers are non-negative
		if err != nil {
			return err
		}
	}

	for _, lc := range labeled {
		if counts == nil || (lc.addr == rwalk && rwNeedsPrefix && rwPrefix == nil) {
			v.add("WARN", "log parity: "+lc.name,
				"checkpoint ahead of DB head; delete the state file to rescan")
			continue
		}
		if lc.addr == rwalk && rwFrom == nil {
			v.add("WARN", "log parity: "+lc.name, "no indexed logs; skipped")
			continue
		}
		chainCnt := counts[lc.addr]
		if lc.addr == rwalk && rwPrefix != nil {
			chainCnt += rwPrefix[lc.addr]
		}
		dbCntBig, err := v.scanBigInt(
			`SELECT COUNT(*)::TEXT FROM evt_log e
			 JOIN address a ON a.address_id = e.contract_aid
			 WHERE LOWER(a.addr) = $1 AND e.block_num <= $2`,
			lc.addr, int64(dbHead)) // #nosec G115 -- block numbers fit int64
		if err != nil {
			return err
		}
		dbCnt := dbCntBig.Uint64()
		v.addMatch(chainCnt == dbCnt, "log parity: "+lc.name,
			fmt.Sprintf("chain=%d db=%d (..%d)", chainCnt, dbCnt, dbHead))
	}
	return nil
}

// scanLogsCheckpointed counts logs per address over [from, to] with one
// multi-address eth_getLogs pass, LOG_WORKERS chunks in flight, folding
// completed chunks into the checkpoint strictly in order so an interrupt
// never records a gap. Returns nil (not an error) if the checkpoint is ahead
// of `to` (e.g. the DB was restored from an older backup).
func (v *chainValidator) scanLogsCheckpointed(
	client *ethclient.Client, state chainScanState, scanKey string,
	addrs []ethcommon.Address, from, to uint64,
) (map[string]uint64, error) {
	entry := state[scanKey]
	if entry == nil || entry.From != from {
		entry = &chainScanEntry{From: from, ScannedTo: from - 1,
			Counts: map[string]uint64{}}
	}
	if entry.ScannedTo > to {
		return nil, nil
	}
	counts := map[string]uint64{}
	for a, n := range entry.Counts {
		counts[a] = n
	}

	type chunk struct{ start, end uint64 }
	var chunks []chunk
	for start := entry.ScannedTo + 1; start <= to; start = chunks[len(chunks)-1].end + 1 {
		end := min(start+v.logChunk-1, to)
		chunks = append(chunks, chunk{start, end})
	}
	if len(chunks) == 0 {
		return counts, nil
	}

	type chunkResult struct {
		counts map[string]uint64
		err    error
	}
	results := make([]chan chunkResult, len(chunks))
	sem := make(chan struct{}, max(v.logWorkers, 1))
	for i := range chunks {
		results[i] = make(chan chunkResult, 1)
		go func(i int) {
			sem <- struct{}{}
			defer func() { <-sem }()
			delta, err := scanLogChunk(v.ctx, client, addrs, chunks[i].start, chunks[i].end)
			results[i] <- chunkResult{delta, err}
		}(i)
	}

	t0 := time.Now()
	total := func() uint64 {
		var n uint64
		for _, c := range counts {
			n += c
		}
		return n
	}
	for i := range chunks {
		var res chunkResult
		select {
		case <-v.ctx.Done():
			return nil, v.ctx.Err()
		case res = <-results[i]:
		}
		if res.err != nil {
			return nil, fmt.Errorf("getLogs %d..%d: %w", chunks[i].start, chunks[i].end, res.err)
		}
		for a, n := range res.counts {
			counts[a] += n
		}
		entry = &chainScanEntry{From: from, ScannedTo: chunks[i].end, Counts: counts}
		state[scanKey] = entry
		if err := saveChainScanState(v.stateFile, state); err != nil {
			return nil, fmt.Errorf("save state: %w", err)
		}
		rate := time.Since(t0) / time.Duration(i+1)
		eta := rate * time.Duration(len(chunks)-i-1)
		v.progressf("[scan %s] blocks %d..%d of %d (%d%%, %d logs, ETA %s)",
			scanKey, chunks[i].start, chunks[i].end, to,
			(chunks[i].end-from+1)*100/(to-from+1), total(), eta.Round(time.Second))
	}
	return counts, nil
}

// scanLogChunk counts logs per address in [from, to], bisecting ranges the
// node rejects (result caps, timeouts on oversized scans).
func scanLogChunk(
	ctx context.Context, client *ethclient.Client,
	addrs []ethcommon.Address, from, to uint64,
) (map[string]uint64, error) {
	logs, err := client.FilterLogs(ctx, ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(from),
		ToBlock:   new(big.Int).SetUint64(to),
		Addresses: addrs,
	})
	if err != nil {
		if from == to || ctx.Err() != nil {
			return nil, err
		}
		mid := from + (to-from)/2
		left, err := scanLogChunk(ctx, client, addrs, from, mid)
		if err != nil {
			return nil, err
		}
		right, err := scanLogChunk(ctx, client, addrs, mid+1, to)
		if err != nil {
			return nil, err
		}
		for a, n := range right {
			left[a] += n
		}
		return left, nil
	}
	counts := map[string]uint64{}
	for i := range logs {
		counts[strings.ToLower(logs[i].Address.Hex())]++
	}
	return counts, nil
}
