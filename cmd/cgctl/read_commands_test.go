package main

// Command-level tests for the read-only subcommands (info, owner,
// donation-records, erc20/nft reads) against scripted contract state on the
// fake chain.

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
)

var (
	testPrizesAddr    = common.HexToAddress("0x4000000000000000000000000000000000000004")
	testCharityAddr   = common.HexToAddress("0x6000000000000000000000000000000000000006")
	testRWalkNFTAddr  = common.HexToAddress("0x7000000000000000000000000000000000000007")
	testStakingCst    = common.HexToAddress("0x8000000000000000000000000000000000000008")
	testStakingRWalk  = common.HexToAddress("0x9000000000000000000000000000000000000009")
	testMarketingAddr = common.HexToAddress("0xa00000000000000000000000000000000000000a")
)

// startReadChain boots a fake chain with only RPC_URL set (read commands
// need no signer).
func startReadChain(t *testing.T) *testchain.Chain {
	t.Helper()
	chain := testchain.New(t)
	chain.EnsureBlock(100)
	chain.SetGasPrice(big.NewInt(1_000_000_000))
	t.Setenv("RPC_URL", chain.URL())
	t.Setenv("PKEY_HEX", "")
	return chain
}

// infoGameStub scripts every read the info command makes against a V1 game.
func infoGameStub() *testchain.ContractStub {
	stub := testchain.MustContractStub(cgc.CosmicSignatureGameABI, cgc.CosmicSignatureGameV2ABI)
	blockTime := int64(testchain.BlockTime(100))

	// Round status.
	stub.Return("roundNum", big.NewInt(3))
	stub.Return("roundActivationTime", big.NewInt(blockTime-600)) // active for 10 min
	stub.Return("delayDurationBeforeRoundActivation", big.NewInt(300))
	stub.Return("getTotalNumBids", big.NewInt(12))
	stub.Return("bidderAddresses", big.NewInt(4))
	// Timing.
	stub.Return("getDurationUntilMainPrize", big.NewInt(90))
	stub.Return("getDurationUntilMainPrizeRaw", big.NewInt(90))
	stub.Return("mainPrizeTime", big.NewInt(blockTime+90))
	stub.Return("timeoutDurationToClaimMainPrize", big.NewInt(86400))
	stub.Return("mainPrizeTimeIncrementInMicroSeconds", big.NewInt(3_600_000_000))
	stub.Return("getMainPrizeTimeIncrement", big.NewInt(3600))
	stub.Return("mainPrizeTimeIncrementIncreaseDivisor", big.NewInt(50))
	stub.Return("initialDurationUntilMainPrizeDivisor", big.NewInt(200))
	stub.Return("getInitialDurationUntilMainPrize", big.NewInt(18))
	// Bidding / prices.
	stub.Return("nextEthBidPrice", eth(0.05))
	stub.Return("getNextEthBidPrice", eth(0.0505))
	stub.Return("getNextCstBidPrice", eth(11))
	stub.Return("ethBidPriceIncreaseDivisor", big.NewInt(100))
	stub.Return("cstRewardAmountForBidding", eth(100)) // V1 fixed reward
	stub.Return("getCstDutchAuctionDurations", big.NewInt(1800), big.NewInt(400))
	stub.Return("getEthDutchAuctionDurations", big.NewInt(3600), big.NewInt(500))
	stub.Return("cstDutchAuctionBeginningBidPrice", eth(22))
	stub.Return("ethDutchAuctionBeginningBidPrice", eth(0.1))
	stub.Return("ethDutchAuctionEndingBidPriceDivisor", big.NewInt(10))
	// Champions.
	stub.Return("lastBidderAddress", otherAddr)
	stub.Return("lastCstBidderAddress", common.Address{})
	stub.Return("tryGetCurrentChampions", otherAddr, big.NewInt(120), testSignerAddr, big.NewInt(60))
	stub.Return("enduranceChampionStartTimeStamp", big.NewInt(blockTime-500))
	stub.Return("prevEnduranceChampionDuration", big.NewInt(45))
	// Prize distribution.
	stub.Return("getMainEthPrizeAmount", eth(2))
	stub.Return("mainEthPrizeAmountPercentage", big.NewInt(25))
	stub.Return("getCharityEthDonationAmount", eth(0.8))
	stub.Return("charityEthDonationAmountPercentage", big.NewInt(10))
	stub.Return("getRaffleTotalEthPrizeAmountForBidders", eth(0.4))
	stub.Return("raffleTotalEthPrizeAmountForBiddersPercentage", big.NewInt(5))
	stub.Return("getChronoWarriorEthPrizeAmount", eth(0.6))
	stub.Return("chronoWarriorEthPrizeAmountPercentage", big.NewInt(8))
	stub.Return("getCosmicSignatureNftStakingTotalEthRewardAmount", eth(0.7))
	stub.Return("cosmicSignatureNftStakingTotalEthRewardAmountPercentage", big.NewInt(9))
	stub.Return("cstPrizeAmount", eth(300))
	// Raffle config.
	stub.Return("numRaffleEthPrizesForBidders", big.NewInt(3))
	stub.Return("numRaffleCosmicSignatureNftsForBidders", big.NewInt(5))
	stub.Return("numRaffleCosmicSignatureNftsForRandomWalkNftStakers", big.NewInt(4))
	// Addresses.
	stub.Return("prizesWallet", testPrizesAddr)
	stub.Return("charityAddress", testCharityAddr)
	stub.Return("nft", testNFTAddr)
	stub.Return("token", testTokenAddr)
	stub.Return("randomWalkNft", testRWalkNFTAddr)
	stub.Return("stakingWalletCosmicSignatureNft", testStakingCst)
	stub.Return("stakingWalletRandomWalkNft", testStakingRWalk)
	stub.Return("marketingWallet", testMarketingAddr)
	stub.Return("marketingWalletCstContributionAmount", eth(50))
	stub.Return("owner", testSignerAddr)
	// Config.
	stub.Return("bidMessageLengthMaxLimit", big.NewInt(280))
	return stub
}

// registerInfoWorld installs the game stub plus the prizes/charity wallets.
func registerInfoWorld(t *testing.T, chain *testchain.Chain, game *testchain.ContractStub) {
	t.Helper()
	chain.RegisterCall(testGameAddr, game.Handler())

	prizes := testchain.MustContractStub(cgc.PrizesWalletABI)
	prizes.Return("nextDonatedNftIndex", big.NewInt(6))
	prizes.Return("timeoutDurationToWithdrawPrizes", big.NewInt(600))
	chain.RegisterCall(testPrizesAddr, prizes.Handler())

	charity := testchain.MustContractStub(cgc.CharityWalletABI)
	charity.Return("charityAddress", otherAddr)
	chain.RegisterCall(testCharityAddr, charity.Handler())
}

func TestInfoCommandDumpsV1State(t *testing.T) {
	chain := startReadChain(t)
	registerInfoWorld(t, chain, infoGameStub())

	out, err := executeCmd(t, newInfoCmd(), testGameAddr.Hex())
	if err != nil {
		t.Fatalf("info: %v\noutput: %s", err, out)
	}
	for _, want := range []string{
		"ROUND STATUS",
		"RoundNum                    = 3",
		"= ACTIVE - started 10 min 0 sec ago",
		"Total bids this round       = 12",
		"TIMING / COUNTDOWN",
		"Duration until prize (clamped)= 90 (1 min 30 sec)",
		"BIDDING / PRICES",
		"CST reward per bid (fixed)", // the V1 branch
		"CURRENT BIDDERS / CHAMPIONS",
		"LastCSTBidder prize (this round)= NO",
		"Chrono Warrior              = " + testSignerAddr.String(),
		"PRIZE DISTRIBUTION",
		"Main prize amount           = 2.000000000000000000 ETH (25%)",
		"RAFFLE CONFIG",
		"PRIZES WALLET",
		"Num donated NFTs            = 6",
		"CHARITY",
		"Charity donation receiver   = " + otherAddr.String(),
		"CONTRACT ADDRESSES",
		"RandomWalkNft               = " + testRWalkNFTAddr.String(),
		"CONFIG PARAMETERS",
		"Bid message max length      = 280",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("info output missing %q", want)
		}
	}
	// The V2-only section must not print for a V1 contract.
	if strings.Contains(out, "V2 INITIALIZED PARAMETERS") {
		t.Errorf("V1 info printed the V2 section:\n%s", out)
	}
}

func TestInfoCommandV2Sections(t *testing.T) {
	chain := startReadChain(t)
	stub := infoGameStub()
	// V2 mechanics: dynamic CST reward and initializeV2 parameters answer.
	stub.Return("getBidCstRewardAmount", eth(7))
	stub.Return("bidCstRewardAmountMultiplier", big.NewInt(10))
	stub.Return("cstDutchAuctionDurationChangeDivisor", big.NewInt(11))
	stub.Return("cstDutchAuctionDuration", big.NewInt(1800))
	registerInfoWorld(t, chain, stub)

	out, err := executeCmd(t, newInfoCmd(), testGameAddr.Hex())
	if err != nil {
		t.Fatalf("info (V2): %v\noutput: %s", err, out)
	}
	for _, want := range []string{
		"getBidCstRewardAmount (next bid)",
		"V2 INITIALIZED PARAMETERS (initializeV2)",
		"cstDutchAuctionDurationChangeDivisor",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("V2 info output missing %q", want)
		}
	}
	if strings.Contains(out, "CST reward per bid (fixed)") {
		t.Errorf("V2 info printed the V1 fixed-reward line:\n%s", out)
	}
}

func TestInfoCommandInactiveRoundAndDefaultAddress(t *testing.T) {
	chain := startReadChain(t)
	stub := infoGameStub()
	stub.Return("roundActivationTime", big.NewInt(int64(testchain.BlockTime(100))+120))
	registerInfoWorld(t, chain, stub)
	// The default local Hardhat address is used when no arg is given.
	chain.RegisterCall(common.HexToAddress(defaultLocalGameAddr), stub.Handler())

	out, err := executeCmd(t, newInfoCmd())
	if err != nil {
		t.Fatalf("info: %v\noutput: %s", err, out)
	}
	if !strings.Contains(out, "Setting default cosmic game contract address to "+defaultLocalGameAddr) {
		t.Errorf("output missing default-address notice:\n%s", out)
	}
	if !strings.Contains(out, "= INACTIVE - activates in 2 min 0 sec") {
		t.Errorf("output missing inactive-round line:\n%s", out)
	}
}

func TestInfoCommandReadFailure(t *testing.T) {
	chain := startReadChain(t)
	// Only roundNum answers; the next read fails like a revert.
	stub := testchain.MustContractStub(cgc.CosmicSignatureGameABI)
	stub.Return("roundNum", big.NewInt(1))
	chain.RegisterCall(testGameAddr, stub.Handler())

	_, err := executeCmd(t, newInfoCmd(), testGameAddr.Hex())
	if err == nil || !strings.Contains(err.Error(), "RoundActivationTime()") {
		t.Errorf("info with failing reads = %v", err)
	}
}

func TestOwnerCommand(t *testing.T) {
	chain := startReadChain(t)
	stub := testchain.MustContractStub(cgc.OwnableABI)
	stub.Return("owner", otherAddr)
	chain.RegisterCall(testGameAddr, stub.Handler())
	chain.SetBalance(otherAddr, eth(3))

	out, err := executeCmd(t, newOwnerCmd(), testGameAddr.Hex())
	if err != nil {
		t.Fatalf("owner: %v\noutput: %s", err, out)
	}
	for _, want := range []string{
		"OWNERSHIP INFO",
		"Owner Address               = " + otherAddr.String(),
		"Owner Balance               = 3.000000000000000000 ETH",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("owner output missing %q:\n%s", want, out)
		}
	}
}

func TestOwnerCommandReadFailure(t *testing.T) {
	chain := startReadChain(t)
	chain.RegisterCall(testGameAddr, testchain.MustContractStub(cgc.OwnableABI).Handler())
	_, err := executeCmd(t, newOwnerCmd(), testGameAddr.Hex())
	if err == nil || !strings.Contains(err.Error(), "Owner()") {
		t.Errorf("owner with revert = %v", err)
	}
}

func TestDonationRecordsCommand(t *testing.T) {
	chain := startReadChain(t)
	stub := testchain.MustContractStub(cgc.CosmicSignatureGameABI)
	stub.Return("numEthDonationWithInfoRecords", big.NewInt(2))
	stub.Handle("ethDonationWithInfoRecords", func(args []any) ([]any, error) {
		idx, _ := args[0].(*big.Int)
		return []any{
			big.NewInt(idx.Int64() + 1), // roundNum
			otherAddr,                   // donor
			eth(0.25),                   // amount
			`{"msg":"gm"}`,              // data
		}, nil
	})
	chain.RegisterCall(testGameAddr, stub.Handler())

	out, err := executeCmd(t, newDonationRecordsCmd(), testGameAddr.Hex())
	if err != nil {
		t.Fatalf("donation-records: %v\noutput: %s", err, out)
	}
	for _, want := range []string{
		"DONATION RECORDS",
		"Total Records               = 2",
		"RECORD 0",
		"RECORD 1",
		"Amount                      = 0.250000000000000000 ETH",
		`{"msg":"gm"}`,
	} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q:\n%s", want, out)
		}
	}
}

func TestDonationRecordsCommandEmpty(t *testing.T) {
	chain := startReadChain(t)
	stub := testchain.MustContractStub(cgc.CosmicSignatureGameABI)
	stub.Return("numEthDonationWithInfoRecords", big.NewInt(0))
	chain.RegisterCall(testGameAddr, stub.Handler())

	out, err := executeCmd(t, newDonationRecordsCmd(), testGameAddr.Hex())
	if err != nil {
		t.Fatalf("donation-records: %v\noutput: %s", err, out)
	}
	if !strings.Contains(out, "No ETH donation with info records found") {
		t.Errorf("output missing empty notice:\n%s", out)
	}
}

func TestERC20BalanceCommand(t *testing.T) {
	chain := startReadChain(t)
	stub := testchain.MustContractStub(cgc.ERC20ABI)
	stub.Return("name", "Cosmic Token")
	stub.Return("symbol", "CST")
	stub.Return("decimals", uint8(18))
	stub.Return("totalSupply", eth(1000))
	stub.Return("balanceOf", eth(12))
	chain.RegisterCall(testTokenAddr, stub.Handler())
	chain.SetBalance(otherAddr, eth(1))

	out, err := executeCmd(t, newERC20BalanceCmd(), testTokenAddr.Hex(), otherAddr.Hex())
	if err != nil {
		t.Fatalf("erc20 balance: %v\noutput: %s", err, out)
	}
	for _, want := range []string{
		"TOKEN INFO",
		"Name                        = Cosmic Token",
		"USER BALANCE",
		"Balance                     = 12.000000000000000000 CST",
		"ETH Balance                 = 1.000000000000000000 ETH",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q:\n%s", want, out)
		}
	}
}

func TestERC20BalanceCommandMetadataFallbacks(t *testing.T) {
	chain := startReadChain(t)
	// name/symbol/decimals revert: the command substitutes UNKNOWN/18.
	stub := testchain.MustContractStub(cgc.ERC20ABI)
	stub.Return("totalSupply", eth(1000))
	stub.Return("balanceOf", eth(5))
	chain.RegisterCall(testTokenAddr, stub.Handler())

	out, err := executeCmd(t, newERC20BalanceCmd(), testTokenAddr.Hex(), otherAddr.Hex())
	if err != nil {
		t.Fatalf("erc20 balance: %v\noutput: %s", err, out)
	}
	if !strings.Contains(out, "Name                        = UNKNOWN") {
		t.Errorf("output missing UNKNOWN fallback:\n%s", out)
	}
}

func TestERC20AllowanceCommand(t *testing.T) {
	statusFor := func(t *testing.T, allowance *big.Int, balance *big.Int) string {
		t.Helper()
		chain := startReadChain(t)
		stub := testchain.MustContractStub(cgc.ERC20ABI)
		stub.Return("symbol", "CST")
		stub.Return("decimals", uint8(18))
		stub.Return("allowance", allowance)
		stub.Return("balanceOf", balance)
		chain.RegisterCall(testTokenAddr, stub.Handler())

		out, err := executeCmd(t, newERC20AllowanceCmd(), testTokenAddr.Hex(), otherAddr.Hex(), testSignerAddr.Hex())
		if err != nil {
			t.Fatalf("erc20 allowance: %v\noutput: %s", err, out)
		}
		return out
	}

	unlimited := new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil), big.NewInt(1))
	if out := statusFor(t, unlimited, eth(5)); !strings.Contains(out, "UNLIMITED (MAX_UINT256)") {
		t.Errorf("unlimited allowance status missing:\n%s", out)
	}
	if out := statusFor(t, eth(10), eth(5)); !strings.Contains(out, "Sufficient for full balance") {
		t.Errorf("sufficient allowance status missing:\n%s", out)
	}
	if out := statusFor(t, eth(1), eth(5)); !strings.Contains(out, "Limited allowance") {
		t.Errorf("limited allowance status missing:\n%s", out)
	}
}

func TestNFTReadCommands(t *testing.T) {
	newNFTStub := func() *testchain.ContractStub {
		stub := testchain.MustContractStub(cgc.CosmicSignatureNftABI, cgc.ERC721ABI)
		stub.Return("getApproved", common.Address{})
		stub.Return("ownerOf", otherAddr)
		stub.Return("balanceOf", big.NewInt(4))
		stub.Return("isApprovedForAll", true)
		return stub
	}

	t.Run("approved not set", func(t *testing.T) {
		chain := startReadChain(t)
		chain.RegisterCall(testNFTAddr, newNFTStub().Handler())
		out, err := executeCmd(t, newNFTApprovedCmd(), testNFTAddr.Hex(), "7")
		if err != nil {
			t.Fatalf("nft approved: %v\noutput: %s", err, out)
		}
		if !strings.Contains(out, "NOT APPROVED - No operator set for this token") {
			t.Errorf("output missing not-approved status:\n%s", out)
		}
	})

	t.Run("approved set", func(t *testing.T) {
		chain := startReadChain(t)
		stub := newNFTStub()
		stub.Return("getApproved", testSignerAddr)
		chain.RegisterCall(testNFTAddr, stub.Handler())
		out, err := executeCmd(t, newNFTApprovedCmd(), testNFTAddr.Hex(), "7")
		if err != nil {
			t.Fatalf("nft approved: %v\noutput: %s", err, out)
		}
		if !strings.Contains(out, "APPROVED - Operator can transfer this token") {
			t.Errorf("output missing approved status:\n%s", out)
		}
	})

	t.Run("is-approved-for-all", func(t *testing.T) {
		chain := startReadChain(t)
		chain.RegisterCall(testNFTAddr, newNFTStub().Handler())
		out, err := executeCmd(t, newNFTIsApprovedForAllCmd(), testNFTAddr.Hex(), otherAddr.Hex(), testSignerAddr.Hex())
		if err != nil {
			t.Fatalf("nft is-approved-for-all: %v\noutput: %s", err, out)
		}
		if !strings.Contains(out, "APPROVED - Operator can transfer all owner's tokens") {
			t.Errorf("output missing approval status:\n%s", out)
		}
	})

	t.Run("owner-of", func(t *testing.T) {
		chain := startReadChain(t)
		chain.RegisterCall(testNFTAddr, newNFTStub().Handler())
		chain.SetBalance(otherAddr, eth(2))
		out, err := executeCmd(t, newNFTOwnerOfCmd(), testNFTAddr.Hex(), "7")
		if err != nil {
			t.Fatalf("nft owner-of: %v\noutput: %s", err, out)
		}
		for _, want := range []string{
			"Owner Address               = " + otherAddr.String(),
			"Owner's Total NFTs          = 4",
			"Owner's ETH Balance         = 2.000000000000000000 ETH",
		} {
			if !strings.Contains(out, want) {
				t.Errorf("output missing %q:\n%s", want, out)
			}
		}
	})

	t.Run("owner-of nonexistent token", func(t *testing.T) {
		chain := startReadChain(t)
		chain.RegisterCall(testNFTAddr, testchain.MustContractStub(cgc.ERC721ABI).Handler())
		_, err := executeCmd(t, newNFTOwnerOfCmd(), testNFTAddr.Hex(), "999")
		if err == nil || !strings.Contains(err.Error(), "token may not exist") {
			t.Errorf("owner-of revert = %v", err)
		}
	})

	t.Run("invalid token id", func(t *testing.T) {
		startReadChain(t)
		_, err := executeCmd(t, newNFTApprovedCmd(), testNFTAddr.Hex(), "seven")
		if err == nil || !strings.Contains(err.Error(), "token-id") {
			t.Errorf("invalid token id = %v", err)
		}
	})
}
