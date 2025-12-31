package decode

import (
	"database/sql"
	"strings"
	"testing"

	freezerscanner "github.com/PredictionExplorer/augur-explorer/rwcg/freezer-scanner"
	"github.com/lib/pq"
)

// ProjectConfig defines contract addresses and expected event signatures for a project
type ProjectConfig struct {
	Name       string
	Contracts  []string // contract addresses (hex with 0x prefix)
	EventSigs  []string // first 8 chars of topic0 (without 0x)
}

var RandomWalkConfig = ProjectConfig{
	Name: "RandomWalk",
	Contracts: []string{
		"0x895a6F444BE4ba9d124F61DF736605792B35D66b", // RandomWalk NFT
		"0x47eF85Dfb775aCE0934fBa9EEd09D22e6eC0Cc08", // Marketplace
	},
	EventSigs: []string{
		"0ff09947", // ?
		"17307eab", // ApprovalForAll
		"55076e90", // TokenNameEvent
		"8ad5e159", // ?
		"8be0079c", // OwnershipTransferred
		"8c5be1e5", // Approval
		"ad2bc79f", // SaleEvent
		"caacc56f", // ?
		"ddf252ad", // Transfer
	},
}

// CosmicGameConfig - ready for when mainnet data is available
var CosmicGameConfig = ProjectConfig{
	Name: "CosmicGame",
	Contracts: []string{
		// Mainnet addresses - to be filled when deployed
		// These are currently dev addresses from cg_contracts table:
		// "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512", // CosmicGame
		// "0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9", // CosmicSignature
		// "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0", // CosmicToken
		// "0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9", // CosmicDao
	},
	EventSigs: []string{
		// Common CosmicGame events (from the contract):
		// BidEvent, PrizeClaimEvent, DonationEvent, RaffleWinnerEvent, etc.
	},
}

// verifyProject compares freezer events with database events for a project
func verifyProject(t *testing.T, db *sql.DB, fr *freezerscanner.FreezerReader, 
	config ProjectConfig, startBlock, endBlock uint64) (match, missing, extra int) {
	
	if len(config.Contracts) == 0 {
		t.Logf("Skipping %s - no contracts configured", config.Name)
		return 0, 0, 0
	}

	// Build contract lookup map
	contractsLower := make(map[string]bool)
	for _, c := range config.Contracts {
		contractsLower[strings.ToLower(c)] = true
	}

	// Get address IDs for contracts
	var contractAids []int64
	for _, addr := range config.Contracts {
		var aid int64
		err := db.QueryRow("SELECT address_id FROM address WHERE LOWER(addr) = LOWER($1)", addr).Scan(&aid)
		if err != nil {
			t.Logf("%s: Contract %s not found in database", config.Name, addr)
			continue
		}
		contractAids = append(contractAids, aid)
	}

	if len(contractAids) == 0 {
		t.Logf("Skipping %s - no contracts found in database", config.Name)
		return 0, 0, 0
	}

	// Query database events
	rows, err := db.Query(`
		SELECT e.block_num, e.topic0_sig, a.addr, COUNT(*) as cnt
		FROM evt_log e
		JOIN address a ON e.contract_aid = a.address_id
		WHERE e.block_num BETWEEN $1 AND $2
		  AND e.contract_aid = ANY($3)
		GROUP BY e.block_num, e.topic0_sig, a.addr
		ORDER BY e.block_num
	`, startBlock, endBlock, pq.Array(contractAids))
	if err != nil {
		t.Fatalf("Query error: %v", err)
	}
	defer rows.Close()

	type eventKey struct {
		blockNum int64
		topic0   string
		contract string
	}
	dbEvents := make(map[eventKey]int)
	for rows.Next() {
		var blockNum int64
		var topic0, contract string
		var cnt int
		if err := rows.Scan(&blockNum, &topic0, &contract, &cnt); err != nil {
			t.Fatalf("Scan error: %v", err)
		}
		dbEvents[eventKey{blockNum, strings.ToLower(topic0), strings.ToLower(contract)}] = cnt
	}
	t.Logf("%s: Database has %d distinct block+topic0+contract combinations", config.Name, len(dbEvents))

	if len(dbEvents) == 0 {
		t.Logf("Skipping %s - no events in database for block range %d-%d", config.Name, startBlock, endBlock)
		return 0, 0, 0
	}

	// Scan freezer for same block range
	freezerEvents := make(map[eventKey]int)
	blocksScanned := 0

	for blockNum := startBlock; blockNum <= endBlock; blockNum++ {
		data, err := fr.ReadItem(blockNum)
		if err != nil || len(data) == 0 {
			continue
		}

		logs, err := DecodeArbitrumReceipts(data)
		if err != nil {
			continue
		}

		blocksScanned++

		for _, log := range logs {
			addr := strings.ToLower(log.Address.Hex())
			if !contractsLower[addr] {
				continue
			}
			if len(log.Topics) == 0 {
				continue
			}

			topic0 := strings.ToLower(log.Topics[0].Hex()[2:10])
			key := eventKey{int64(blockNum), topic0, addr}
			freezerEvents[key]++
		}
	}

	t.Logf("%s: Freezer scanned %d blocks, found %d distinct events", config.Name, blocksScanned, len(freezerEvents))

	// Compare
	for key, dbCount := range dbEvents {
		if fCount, ok := freezerEvents[key]; ok && fCount == dbCount {
			match++
		} else if !ok {
			missing++
			if missing <= 3 {
				t.Logf("Missing from freezer: block %d topic0 %s contract %s", key.blockNum, key.topic0, key.contract)
			}
		}
	}

	for key := range freezerEvents {
		if _, ok := dbEvents[key]; !ok {
			extra++
			if extra <= 3 {
				t.Logf("Extra in freezer: block %d topic0 %s contract %s", key.blockNum, key.topic0, key.contract)
			}
		}
	}

	return match, missing, extra
}

// TestVerifyRandomWalk verifies RandomWalk project events
func TestVerifyRandomWalk(t *testing.T) {
	db, err := sql.Open("postgres", "user=cosmicgame dbname=cosmicgame sslmode=disable")
	if err != nil {
		t.Skipf("Cannot connect to database: %v", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		t.Skipf("Cannot ping database: %v", err)
	}

	fr, err := freezerscanner.NewFreezerReader("../mainnet")
	if err != nil {
		t.Skipf("Cannot open freezer: %v", err)
	}

	// Test first 100K blocks from first event
	match, missing, extra := verifyProject(t, db, fr, RandomWalkConfig, 2910155, 3010156)

	total := match + missing
	if total == 0 {
		t.Skip("No events to verify")
	}

	matchRate := float64(match) / float64(total) * 100
	t.Logf("\nRandomWalk Verification:")
	t.Logf("  Matching: %d", match)
	t.Logf("  Missing: %d", missing)
	t.Logf("  Extra: %d", extra)
	t.Logf("  Match rate: %.1f%%", matchRate)

	if matchRate < 99.0 {
		t.Errorf("Expected at least 99%% match rate, got %.1f%%", matchRate)
	}
}

// TestVerifyCosmicGame verifies CosmicGame project events
func TestVerifyCosmicGame(t *testing.T) {
	if len(CosmicGameConfig.Contracts) == 0 {
		t.Skip("CosmicGame mainnet contracts not configured yet")
	}

	db, err := sql.Open("postgres", "user=cosmicgame dbname=cosmicgame sslmode=disable")
	if err != nil {
		t.Skipf("Cannot connect to database: %v", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		t.Skipf("Cannot ping database: %v", err)
	}

	fr, err := freezerscanner.NewFreezerReader("../mainnet")
	if err != nil {
		t.Skipf("Cannot open freezer: %v", err)
	}

	// Get block range from database
	var minBlock, maxBlock int64
	err = db.QueryRow(`
		SELECT MIN(e.block_num), MAX(e.block_num)
		FROM evt_log e
		JOIN address a ON e.contract_aid = a.address_id
		WHERE LOWER(a.addr) = ANY($1)
	`, pq.Array(CosmicGameConfig.Contracts)).Scan(&minBlock, &maxBlock)
	if err != nil || minBlock == 0 {
		t.Skip("No CosmicGame events in database")
	}

	// Test first 100K blocks
	endBlock := minBlock + 100000
	if endBlock > maxBlock {
		endBlock = maxBlock
	}

	match, missing, extra := verifyProject(t, db, fr, CosmicGameConfig, uint64(minBlock), uint64(endBlock))

	total := match + missing
	if total == 0 {
		t.Skip("No events to verify")
	}

	matchRate := float64(match) / float64(total) * 100
	t.Logf("\nCosmicGame Verification:")
	t.Logf("  Matching: %d", match)
	t.Logf("  Missing: %d", missing)
	t.Logf("  Extra: %d", extra)
	t.Logf("  Match rate: %.1f%%", matchRate)

	if matchRate < 99.0 {
		t.Errorf("Expected at least 99%% match rate, got %.1f%%", matchRate)
	}
}

// TestVerifyAllProjects runs verification for all configured projects
func TestVerifyAllProjects(t *testing.T) {
	db, err := sql.Open("postgres", "user=cosmicgame dbname=cosmicgame sslmode=disable")
	if err != nil {
		t.Skipf("Cannot connect to database: %v", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		t.Skipf("Cannot ping database: %v", err)
	}

	fr, err := freezerscanner.NewFreezerReader("../mainnet")
	if err != nil {
		t.Skipf("Cannot open freezer: %v", err)
	}

	projects := []ProjectConfig{RandomWalkConfig, CosmicGameConfig}

	for _, config := range projects {
		t.Run(config.Name, func(t *testing.T) {
			if len(config.Contracts) == 0 {
				t.Skipf("%s: No contracts configured", config.Name)
				return
			}

			// Get block range from database
			var minBlock, maxBlock sql.NullInt64
			contractsLower := make([]string, len(config.Contracts))
			for i, c := range config.Contracts {
				contractsLower[i] = strings.ToLower(c)
			}

			err := db.QueryRow(`
				SELECT MIN(e.block_num), MAX(e.block_num)
				FROM evt_log e
				JOIN address a ON e.contract_aid = a.address_id
				WHERE LOWER(a.addr) = ANY($1)
			`, pq.Array(contractsLower)).Scan(&minBlock, &maxBlock)
			if err != nil || !minBlock.Valid {
				t.Skipf("%s: No events in database", config.Name)
				return
			}

			// Test first 100K blocks
			endBlock := minBlock.Int64 + 100000
			if endBlock > maxBlock.Int64 {
				endBlock = maxBlock.Int64
			}

			match, missing, extra := verifyProject(t, db, fr, config, uint64(minBlock.Int64), uint64(endBlock))

			total := match + missing
			if total == 0 {
				t.Skipf("%s: No events to verify", config.Name)
				return
			}

			matchRate := float64(match) / float64(total) * 100
			t.Logf("Match: %d, Missing: %d, Extra: %d, Rate: %.1f%%", match, missing, extra, matchRate)

			if matchRate < 99.0 {
				t.Errorf("Expected at least 99%% match rate, got %.1f%%", matchRate)
			}
		})
	}
}
