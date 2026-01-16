package decode

import (
	"strings"
	"testing"

	freezerscanner "github.com/PredictionExplorer/augur-explorer/rwcg/freezer-scanner"
)

// RandomWalk contract address on Arbitrum
const RandomWalkAddr = "0x895a6F444BE4ba9d124F61DF736605792B35D66b"

// RandomWalk event signatures (first 4 bytes of keccak256)
var RandomWalkEventSigs = map[string]string{
	"0ff09947": "SeedInfoEvent",
	"17307eab": "ApprovalForAll",
	"55076e90": "TokenNameEvent",
	"8ad5e159": "MintEvent",
	"8be0079c": "OwnershipTransferred",
	"8c5be1e5": "Approval",
	"ad2bc79f": "SaleEvent",
	"caacc56f": "WithdrawalEvent",
	"ddf252ad": "Transfer",
}

func TestDecodeArbitrumReceipts(t *testing.T) {
	// Open freezer reader
	fr, err := freezerscanner.NewFreezerReader("../mainnet")
	if err != nil {
		t.Skipf("Skipping: cannot open freezer: %v", err)
	}

	// Test block range where we know RandomWalk events exist
	// Block 2910155 has OwnershipTransferred, blocks 2983675+ have Transfer/SaleEvent
	startBlock := uint64(2_910_000)
	endBlock := uint64(3_000_000)

	t.Logf("Scanning blocks %d to %d for RandomWalk events", startBlock, endBlock)

	rwalkAddrLower := strings.ToLower(RandomWalkAddr)
	totalLogs := 0
	rwalkEvents := 0
	blocksWithEvents := 0

	for blockNum := startBlock; blockNum < endBlock; blockNum++ {
		data, err := fr.ReadItem(blockNum)
		if err != nil || len(data) == 0 {
			continue
		}

		logs, err := DecodeArbitrumReceipts(data)
		if err != nil {
			continue
		}

		totalLogs += len(logs)
		blockHasRwalk := false

		for _, log := range logs {
			if strings.ToLower(log.Address.Hex()) != rwalkAddrLower {
				continue
			}
			if len(log.Topics) == 0 {
				continue
			}

			topic0 := strings.ToLower(log.Topics[0].Hex()[2:10])
			if eventName, ok := RandomWalkEventSigs[topic0]; ok {
				rwalkEvents++
				blockHasRwalk = true
				if rwalkEvents <= 10 {
					t.Logf("Block %d: %s (topic0=%s)", blockNum, eventName, topic0)
				}
			}
		}

		if blockHasRwalk {
			blocksWithEvents++
		}
	}

	t.Logf("\nResults:")
	t.Logf("  Blocks scanned: %d", endBlock-startBlock)
	t.Logf("  Total logs decoded: %d", totalLogs)
	t.Logf("  RandomWalk events found: %d", rwalkEvents)
	t.Logf("  Blocks with RandomWalk events: %d", blocksWithEvents)

	// Verify we found events
	if rwalkEvents == 0 {
		t.Error("Expected to find RandomWalk events in block range 2910000-3000000")
	}
	if totalLogs == 0 {
		t.Error("Expected to decode logs from freezer data")
	}
}

func TestDecodeFirstBlocks(t *testing.T) {
	// Quick sanity test on first few blocks
	fr, err := freezerscanner.NewFreezerReader("../mainnet")
	if err != nil {
		t.Skipf("Skipping: cannot open freezer: %v", err)
	}

	successCount := 0
	for blockNum := uint64(0); blockNum < 100; blockNum++ {
		data, err := fr.ReadItem(blockNum)
		if err != nil || len(data) == 0 {
			continue
		}

		logs, err := DecodeArbitrumReceipts(data)
		if err != nil {
			t.Logf("Block %d: decode error: %v", blockNum, err)
			continue
		}
		successCount++
		if blockNum < 5 {
			t.Logf("Block %d: %d logs", blockNum, len(logs))
		}
	}

	t.Logf("Successfully decoded %d/100 blocks", successCount)
	if successCount < 50 {
		t.Errorf("Expected at least 50 blocks to decode, got %d", successCount)
	}
}

