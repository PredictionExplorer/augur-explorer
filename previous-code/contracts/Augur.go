// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IAugurCreationDataGetterMarketCreationData is an auto generated low-level Go binding around an user-defined struct.
type IAugurCreationDataGetterMarketCreationData struct {
	ExtraInfo                string
	MarketCreator            common.Address
	Outcomes                 [][32]byte
	DisplayPrices            []*big.Int
	MarketType               uint8
	RecommendedTradeInterval *big.Int
}

// AugurMetaData contains all meta data concerning the Augur contract.
var AugurMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numCompleteSets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CompleteSetsPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numCompleteSets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CompleteSetsSold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"designatedReportStake\",\"type\":\"uint256\"}],\"name\":\"DesignatedReportStakeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"disputeCrowdsourcer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"payoutNumerators\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nextWindowStartTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nextWindowEndTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pacingOn\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalRepStakedInPayout\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalRepStakedInMarket\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"disputeRound\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DisputeCrowdsourcerCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"disputeCrowdsourcer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"payoutNumerators\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentStake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeRemaining\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"disputeRound\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DisputeCrowdsourcerContribution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"disputeCrowdsourcer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"payoutNumerators\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"disputeRound\",\"type\":\"uint256\"}],\"name\":\"DisputeCrowdsourcerCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"disputeCrowdsourcer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountRedeemed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repReceived\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"payoutNumerators\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DisputeCrowdsourcerRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"disputeWindow\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"initial\",\"type\":\"bool\"}],\"name\":\"DisputeWindowCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"FinishDeployment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialReporter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isDesignatedReporter\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"payoutNumerators\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nextWindowStartTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nextWindowEndTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"InitialReportSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialReporter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountRedeemed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repReceived\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"payoutNumerators\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"InitialReporterRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"InitialReporterTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIUniverse\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"extraInfo\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"contractIMarket\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"marketCreator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"designatedReporter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feePerCashInAttoCash\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256[]\",\"name\":\"prices\",\"type\":\"int256[]\"},{\"indexed\":false,\"internalType\":\"enumIMarket.MarketType\",\"name\":\"marketType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numTicks\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"outcomes\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"noShowBond\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"MarketCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"winningPayoutNumerators\",\"type\":\"uint256[]\"}],\"name\":\"MarketFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"originalUniverse\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newUniverse\",\"type\":\"address\"}],\"name\":\"MarketMigrated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marketOI\",\"type\":\"uint256\"}],\"name\":\"MarketOIChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"MarketParticipantsDisavowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"MarketRepBondTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"MarketTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"noShowBond\",\"type\":\"uint256\"}],\"name\":\"NoShowBondChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"disputeWindow\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"attoParticipationTokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feePayoutShare\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ParticipationTokensRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"RegisterContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reportingFee\",\"type\":\"uint256\"}],\"name\":\"ReportingFeeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"reportingParticipant\",\"type\":\"address\"}],\"name\":\"ReportingParticipantDisavowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outcome\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"ShareTokenBalanceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTimestamp\",\"type\":\"uint256\"}],\"name\":\"TimestampSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumAugur.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outcome\",\"type\":\"uint256\"}],\"name\":\"TokenBalanceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumAugur.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"TokensBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumAugur.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"TokensMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumAugur.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"TokensTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outcome\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numShares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numPayoutTokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TradingProceedsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"parentUniverse\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"childUniverse\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"payoutNumerators\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"creationTimestamp\",\"type\":\"uint256\"}],\"name\":\"UniverseCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIMarket\",\"name\":\"forkingMarket\",\"type\":\"address\"}],\"name\":\"UniverseForked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validityBond\",\"type\":\"uint256\"}],\"name\":\"ValidityBondChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"warpSyncHash\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marketEndTime\",\"type\":\"uint256\"}],\"name\":\"WarpSyncDataUpdated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"DEFAULT_RECOMMENDED_TRADE_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_TRADE_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TRADE_INTERVAL_VALUE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cash\",\"outputs\":[{\"internalType\":\"contractICash\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parentPayoutDistributionHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"_parentPayoutNumerators\",\"type\":\"uint256[]\"}],\"name\":\"createChildUniverse\",\"outputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createGenesisUniverse\",\"outputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_payoutNumerators\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_numTicks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numOutcomes\",\"type\":\"uint256\"}],\"name\":\"derivePayoutDistributionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_disputeCrowdsourcer\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_payoutNumerators\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_disputeRound\",\"type\":\"uint256\"}],\"name\":\"disputeCrowdsourcerCreated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishDeployment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"forkCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"genesisUniverse\",\"outputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"}],\"name\":\"getMarketCreationData\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"extraInfo\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"marketCreator\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"outcomes\",\"type\":\"bytes32[]\"},{\"internalType\":\"int256[]\",\"name\":\"displayPrices\",\"type\":\"int256[]\"},{\"internalType\":\"enumIMarket.MarketType\",\"name\":\"marketType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"recommendedTradeInterval\",\"type\":\"uint256\"}],\"internalType\":\"structIAugurCreationDataGetter.MarketCreationData\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"}],\"name\":\"getMarketOutcomes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_outcomes\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"}],\"name\":\"getMarketRecommendedTradeInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"}],\"name\":\"getMarketType\",\"outputs\":[{\"internalType\":\"enumIMarket.MarketType\",\"name\":\"_marketType\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getMaximumMarketEndDate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_displayRange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numTicks\",\"type\":\"uint256\"}],\"name\":\"getTradeInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"}],\"name\":\"getUniverseForkIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIDisputeCrowdsourcer\",\"name\":\"_crowdsourcer\",\"type\":\"address\"}],\"name\":\"isKnownCrowdsourcer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeSender\",\"type\":\"address\"}],\"name\":\"isKnownFeeSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"}],\"name\":\"isKnownMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"}],\"name\":\"isKnownUniverse\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isTrustedSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_numCompleteSets\",\"type\":\"uint256\"}],\"name\":\"logCompleteSetsPurchased\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_numCompleteSets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fees\",\"type\":\"uint256\"}],\"name\":\"logCompleteSetsSold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_designatedReportStake\",\"type\":\"uint256\"}],\"name\":\"logDesignatedReportStakeChanged\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_disputeCrowdsourcer\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_payoutNumerators\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_nextWindowStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nextWindowEndTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_pacingOn\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_totalRepStakedInPayout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalRepStakedInMarket\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_disputeRound\",\"type\":\"uint256\"}],\"name\":\"logDisputeCrowdsourcerCompleted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reporter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_disputeCrowdsourcer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountStaked\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"_payoutNumerators\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_currentStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stakeRemaining\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_disputeRound\",\"type\":\"uint256\"}],\"name\":\"logDisputeCrowdsourcerContribution\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reporter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountRedeemed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_repReceived\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_payoutNumerators\",\"type\":\"uint256[]\"}],\"name\":\"logDisputeCrowdsourcerRedeemed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"logDisputeCrowdsourcerTokensBurned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"logDisputeCrowdsourcerTokensMinted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fromBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_toBalance\",\"type\":\"uint256\"}],\"name\":\"logDisputeCrowdsourcerTokensTransferred\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIDisputeWindow\",\"name\":\"_disputeWindow\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_initial\",\"type\":\"bool\"}],\"name\":\"logDisputeWindowCreated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reporter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialReporter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountStaked\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isDesignatedReporter\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"_payoutNumerators\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_nextWindowStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nextWindowEndTime\",\"type\":\"uint256\"}],\"name\":\"logInitialReportSubmitted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reporter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountRedeemed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_repReceived\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_payoutNumerators\",\"type\":\"uint256[]\"}],\"name\":\"logInitialReporterRedeemed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"logInitialReporterTransferred\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_winningPayoutNumerators\",\"type\":\"uint256[]\"}],\"name\":\"logMarketFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"contractIUniverse\",\"name\":\"_originalUniverse\",\"type\":\"address\"}],\"name\":\"logMarketMigrated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"}],\"name\":\"logMarketOIChanged\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"}],\"name\":\"logMarketParticipantsDisavowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"logMarketRepBondTransferred\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"logMarketTransferred\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_noShowBond\",\"type\":\"uint256\"}],\"name\":\"logNoShowBondChanged\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"logParticipationTokensBurned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"logParticipationTokensMinted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_attoParticipationTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feePayoutShare\",\"type\":\"uint256\"}],\"name\":\"logParticipationTokensRedeemed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fromBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_toBalance\",\"type\":\"uint256\"}],\"name\":\"logParticipationTokensTransferred\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_reportingFee\",\"type\":\"uint256\"}],\"name\":\"logReportingFeeChanged\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"}],\"name\":\"logReportingParticipantDisavowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"logReputationTokensBurned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"logReputationTokensMinted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fromBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_toBalance\",\"type\":\"uint256\"}],\"name\":\"logReputationTokensTransferred\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outcome\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"logShareTokensBalanceChanged\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newTimestamp\",\"type\":\"uint256\"}],\"name\":\"logTimestampSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outcome\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numShares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numPayoutTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fees\",\"type\":\"uint256\"}],\"name\":\"logTradingProceedsClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_forkingMarket\",\"type\":\"address\"}],\"name\":\"logUniverseForked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validityBond\",\"type\":\"uint256\"}],\"name\":\"logValidityBondChanged\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_warpSyncHash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_marketEndTime\",\"type\":\"uint256\"}],\"name\":\"logWarpSyncDataUpdated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"lookup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_extraInfo\",\"type\":\"string\"},{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_marketCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_designatedReporter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_feePerCashInAttoCash\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_outcomes\",\"type\":\"bytes32[]\"}],\"name\":\"onCategoricalMarketCreated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_extraInfo\",\"type\":\"string\"},{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_marketCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_designatedReporter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_feePerCashInAttoCash\",\"type\":\"uint256\"},{\"internalType\":\"int256[]\",\"name\":\"_prices\",\"type\":\"int256[]\"},{\"internalType\":\"uint256\",\"name\":\"_numTicks\",\"type\":\"uint256\"}],\"name\":\"onScalarMarketCreated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_extraInfo\",\"type\":\"string\"},{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_marketCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_designatedReporter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_feePerCashInAttoCash\",\"type\":\"uint256\"}],\"name\":\"onYesNoMarketCreated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"registerContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"time\",\"outputs\":[{\"internalType\":\"contractITime\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"trustedCashTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"upgradeTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"uploader\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractIDaiVat\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AugurABI is the input ABI used to generate the binding from.
// Deprecated: Use AugurMetaData.ABI instead.
var AugurABI = AugurMetaData.ABI

// Augur is an auto generated Go binding around an Ethereum contract.
type Augur struct {
	AugurCaller     // Read-only binding to the contract
	AugurTransactor // Write-only binding to the contract
	AugurFilterer   // Log filterer for contract events
}

// AugurCaller is an auto generated read-only Go binding around an Ethereum contract.
type AugurCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AugurTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AugurTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AugurFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AugurFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AugurSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AugurSession struct {
	Contract     *Augur            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AugurCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AugurCallerSession struct {
	Contract *AugurCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AugurTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AugurTransactorSession struct {
	Contract     *AugurTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AugurRaw is an auto generated low-level Go binding around an Ethereum contract.
type AugurRaw struct {
	Contract *Augur // Generic contract binding to access the raw methods on
}

// AugurCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AugurCallerRaw struct {
	Contract *AugurCaller // Generic read-only contract binding to access the raw methods on
}

// AugurTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AugurTransactorRaw struct {
	Contract *AugurTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAugur creates a new instance of Augur, bound to a specific deployed contract.
func NewAugur(address common.Address, backend bind.ContractBackend) (*Augur, error) {
	contract, err := bindAugur(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Augur{AugurCaller: AugurCaller{contract: contract}, AugurTransactor: AugurTransactor{contract: contract}, AugurFilterer: AugurFilterer{contract: contract}}, nil
}

// NewAugurCaller creates a new read-only instance of Augur, bound to a specific deployed contract.
func NewAugurCaller(address common.Address, caller bind.ContractCaller) (*AugurCaller, error) {
	contract, err := bindAugur(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AugurCaller{contract: contract}, nil
}

// NewAugurTransactor creates a new write-only instance of Augur, bound to a specific deployed contract.
func NewAugurTransactor(address common.Address, transactor bind.ContractTransactor) (*AugurTransactor, error) {
	contract, err := bindAugur(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AugurTransactor{contract: contract}, nil
}

// NewAugurFilterer creates a new log filterer instance of Augur, bound to a specific deployed contract.
func NewAugurFilterer(address common.Address, filterer bind.ContractFilterer) (*AugurFilterer, error) {
	contract, err := bindAugur(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AugurFilterer{contract: contract}, nil
}

// bindAugur binds a generic wrapper to an already deployed contract.
func bindAugur(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AugurABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Augur *AugurRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Augur.Contract.AugurCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Augur *AugurRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Augur.Contract.AugurTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Augur *AugurRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Augur.Contract.AugurTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Augur *AugurCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Augur.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Augur *AugurTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Augur.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Augur *AugurTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Augur.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTRECOMMENDEDTRADEINTERVAL is a free data retrieval call binding the contract method 0xbd33b942.
//
// Solidity: function DEFAULT_RECOMMENDED_TRADE_INTERVAL() view returns(uint256)
func (_Augur *AugurCaller) DEFAULTRECOMMENDEDTRADEINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Augur.contract.Call(opts, &out, "DEFAULT_RECOMMENDED_TRADE_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTRECOMMENDEDTRADEINTERVAL is a free data retrieval call binding the contract method 0xbd33b942.
//
// Solidity: function DEFAULT_RECOMMENDED_TRADE_INTERVAL() view returns(uint256)
func (_Augur *AugurSession) DEFAULTRECOMMENDEDTRADEINTERVAL() (*big.Int, error) {
	return _Augur.Contract.DEFAULTRECOMMENDEDTRADEINTERVAL(&_Augur.CallOpts)
}

// DEFAULTRECOMMENDEDTRADEINTERVAL is a free data retrieval call binding the contract method 0xbd33b942.
//
// Solidity: function DEFAULT_RECOMMENDED_TRADE_INTERVAL() view returns(uint256)
func (_Augur *AugurCallerSession) DEFAULTRECOMMENDEDTRADEINTERVAL() (*big.Int, error) {
	return _Augur.Contract.DEFAULTRECOMMENDEDTRADEINTERVAL(&_Augur.CallOpts)
}

// MINTRADEINTERVAL is a free data retrieval call binding the contract method 0x8f401bf3.
//
// Solidity: function MIN_TRADE_INTERVAL() view returns(uint256)
func (_Augur *AugurCaller) MINTRADEINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Augur.contract.Call(opts, &out, "MIN_TRADE_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTRADEINTERVAL is a free data retrieval call binding the contract method 0x8f401bf3.
//
// Solidity: function MIN_TRADE_INTERVAL() view returns(uint256)
func (_Augur *AugurSession) MINTRADEINTERVAL() (*big.Int, error) {
	return _Augur.Contract.MINTRADEINTERVAL(&_Augur.CallOpts)
}

// MINTRADEINTERVAL is a free data retrieval call binding the contract method 0x8f401bf3.
//
// Solidity: function MIN_TRADE_INTERVAL() view returns(uint256)
func (_Augur *AugurCallerSession) MINTRADEINTERVAL() (*big.Int, error) {
	return _Augur.Contract.MINTRADEINTERVAL(&_Augur.CallOpts)
}

// TRADEINTERVALVALUE is a free data retrieval call binding the contract method 0xff51f367.
//
// Solidity: function TRADE_INTERVAL_VALUE() view returns(uint256)
func (_Augur *AugurCaller) TRADEINTERVALVALUE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Augur.contract.Call(opts, &out, "TRADE_INTERVAL_VALUE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TRADEINTERVALVALUE is a free data retrieval call binding the contract method 0xff51f367.
//
// Solidity: function TRADE_INTERVAL_VALUE() view returns(uint256)
func (_Augur *AugurSession) TRADEINTERVALVALUE() (*big.Int, error) {
	return _Augur.Contract.TRADEINTERVALVALUE(&_Augur.CallOpts)
}

// TRADEINTERVALVALUE is a free data retrieval call binding the contract method 0xff51f367.
//
// Solidity: function TRADE_INTERVAL_VALUE() view returns(uint256)
func (_Augur *AugurCallerSession) TRADEINTERVALVALUE() (*big.Int, error) {
	return _Augur.Contract.TRADEINTERVALVALUE(&_Augur.CallOpts)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_Augur *AugurCaller) Cash(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Augur.contract.Call(opts, &out, "cash")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_Augur *AugurSession) Cash() (common.Address, error) {
	return _Augur.Contract.Cash(&_Augur.CallOpts)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_Augur *AugurCallerSession) Cash() (common.Address, error) {
	return _Augur.Contract.Cash(&_Augur.CallOpts)
}

// DerivePayoutDistributionHash is a free data retrieval call binding the contract method 0x752f57c0.
//
// Solidity: function derivePayoutDistributionHash(uint256[] _payoutNumerators, uint256 _numTicks, uint256 _numOutcomes) view returns(bytes32)
func (_Augur *AugurCaller) DerivePayoutDistributionHash(opts *bind.CallOpts, _payoutNumerators []*big.Int, _numTicks *big.Int, _numOutcomes *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Augur.contract.Call(opts, &out, "derivePayoutDistributionHash", _payoutNumerators, _numTicks, _numOutcomes)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DerivePayoutDistributionHash is a free data retrieval call binding the contract method 0x752f57c0.
//
// Solidity: function derivePayoutDistributionHash(uint256[] _payoutNumerators, uint256 _numTicks, uint256 _numOutcomes) view returns(bytes32)
func (_Augur *AugurSession) DerivePayoutDistributionHash(_payoutNumerators []*big.Int, _numTicks *big.Int, _numOutcomes *big.Int) ([32]byte, error) {
	return _Augur.Contract.DerivePayoutDistributionHash(&_Augur.CallOpts, _payoutNumerators, _numTicks, _numOutcomes)
}

// DerivePayoutDistributionHash is a free data retrieval call binding the contract method 0x752f57c0.
//
// Solidity: function derivePayoutDistributionHash(uint256[] _payoutNumerators, uint256 _numTicks, uint256 _numOutcomes) view returns(bytes32)
func (_Augur *AugurCallerSession) DerivePayoutDistributionHash(_payoutNumerators []*big.Int, _numTicks *big.Int, _numOutcomes *big.Int) ([32]byte, error) {
	return _Augur.Contract.DerivePayoutDistributionHash(&_Augur.CallOpts, _payoutNumerators, _numTicks, _numOutcomes)
}

// ForkCounter is a free data retrieval call binding the contract method 0xce07324e.
//
// Solidity: function forkCounter() view returns(uint256)
func (_Augur *AugurCaller) ForkCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Augur.contract.Call(opts, &out, "forkCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ForkCounter is a free data retrieval call binding the contract method 0xce07324e.
//
// Solidity: function forkCounter() view returns(uint256)
func (_Augur *AugurSession) ForkCounter() (*big.Int, error) {
	return _Augur.Contract.ForkCounter(&_Augur.CallOpts)
}

// ForkCounter is a free data retrieval call binding the contract method 0xce07324e.
//
// Solidity: function forkCounter() view returns(uint256)
func (_Augur *AugurCallerSession) ForkCounter() (*big.Int, error) {
	return _Augur.Contract.ForkCounter(&_Augur.CallOpts)
}

// GenesisUniverse is a free data retrieval call binding the contract method 0x00b946f0.
//
// Solidity: function genesisUniverse() view returns(address)
func (_Augur *AugurCaller) GenesisUniverse(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Augur.contract.Call(opts, &out, "genesisUniverse")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GenesisUniverse is a free data retrieval call binding the contract method 0x00b946f0.
//
// Solidity: function genesisUniverse() view returns(address)
func (_Augur *AugurSession) GenesisUniverse() (common.Address, error) {
	return _Augur.Contract.GenesisUniverse(&_Augur.CallOpts)
}

// GenesisUniverse is a free data retrieval call binding the contract method 0x00b946f0.
//
// Solidity: function genesisUniverse() view returns(address)
func (_Augur *AugurCallerSession) GenesisUniverse() (common.Address, error) {
	return _Augur.Contract.GenesisUniverse(&_Augur.CallOpts)
}

// GetMarketCreationData is a free data retrieval call binding the contract method 0x4891c9ab.
//
// Solidity: function getMarketCreationData(address _market) view returns((string,address,bytes32[],int256[],uint8,uint256))
func (_Augur *AugurCaller) GetMarketCreationData(opts *bind.CallOpts, _market common.Address) (IAugurCreationDataGetterMarketCreationData, error) {
	var out []interface{}
	err := _Augur.contract.Call(opts, &out, "getMarketCreationData", _market)

	if err != nil {
		return *new(IAugurCreationDataGetterMarketCreationData), err
	}

	out0 := *abi.ConvertType(out[0], new(IAugurCreationDataGetterMarketCreationData)).(*IAugurCreationDataGetterMarketCreationData)

	return out0, err

}

// GetMarketCreationData is a free data retrieval call binding the contract method 0x4891c9ab.
//
// Solidity: function getMarketCreationData(address _market) view returns((string,address,bytes32[],int256[],uint8,uint256))
func (_Augur *AugurSession) GetMarketCreationData(_market common.Address) (IAugurCreationDataGetterMarketCreationData, error) {
	return _Augur.Contract.GetMarketCreationData(&_Augur.CallOpts, _market)
}

// GetMarketCreationData is a free data retrieval call binding the contract method 0x4891c9ab.
//
// Solidity: function getMarketCreationData(address _market) view returns((string,address,bytes32[],int256[],uint8,uint256))
func (_Augur *AugurCallerSession) GetMarketCreationData(_market common.Address) (IAugurCreationDataGetterMarketCreationData, error) {
	return _Augur.Contract.GetMarketCreationData(&_Augur.CallOpts, _market)
}

// GetMarketOutcomes is a free data retrieval call binding the contract method 0x887cc5e4.
//
// Solidity: function getMarketOutcomes(address _market) view returns(bytes32[] _outcomes)
func (_Augur *AugurCaller) GetMarketOutcomes(opts *bind.CallOpts, _market common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _Augur.contract.Call(opts, &out, "getMarketOutcomes", _market)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetMarketOutcomes is a free data retrieval call binding the contract method 0x887cc5e4.
//
// Solidity: function getMarketOutcomes(address _market) view returns(bytes32[] _outcomes)
func (_Augur *AugurSession) GetMarketOutcomes(_market common.Address) ([][32]byte, error) {
	return _Augur.Contract.GetMarketOutcomes(&_Augur.CallOpts, _market)
}

// GetMarketOutcomes is a free data retrieval call binding the contract method 0x887cc5e4.
//
// Solidity: function getMarketOutcomes(address _market) view returns(bytes32[] _outcomes)
func (_Augur *AugurCallerSession) GetMarketOutcomes(_market common.Address) ([][32]byte, error) {
	return _Augur.Contract.GetMarketOutcomes(&_Augur.CallOpts, _market)
}

// GetMarketRecommendedTradeInterval is a free data retrieval call binding the contract method 0xf9c45f71.
//
// Solidity: function getMarketRecommendedTradeInterval(address _market) view returns(uint256)
func (_Augur *AugurCaller) GetMarketRecommendedTradeInterval(opts *bind.CallOpts, _market common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Augur.contract.Call(opts, &out, "getMarketRecommendedTradeInterval", _market)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketRecommendedTradeInterval is a free data retrieval call binding the contract method 0xf9c45f71.
//
// Solidity: function getMarketRecommendedTradeInterval(address _market) view returns(uint256)
func (_Augur *AugurSession) GetMarketRecommendedTradeInterval(_market common.Address) (*big.Int, error) {
	return _Augur.Contract.GetMarketRecommendedTradeInterval(&_Augur.CallOpts, _market)
}

// GetMarketRecommendedTradeInterval is a free data retrieval call binding the contract method 0xf9c45f71.
//
// Solidity: function getMarketRecommendedTradeInterval(address _market) view returns(uint256)
func (_Augur *AugurCallerSession) GetMarketRecommendedTradeInterval(_market common.Address) (*big.Int, error) {
	return _Augur.Contract.GetMarketRecommendedTradeInterval(&_Augur.CallOpts, _market)
}

// GetMarketType is a free data retrieval call binding the contract method 0x9eba4b7b.
//
// Solidity: function getMarketType(address _market) view returns(uint8 _marketType)
func (_Augur *AugurCaller) GetMarketType(opts *bind.CallOpts, _market common.Address) (uint8, error) {
	var out []interface{}
	err := _Augur.contract.Call(opts, &out, "getMarketType", _market)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetMarketType is a free data retrieval call binding the contract method 0x9eba4b7b.
//
// Solidity: function getMarketType(address _market) view returns(uint8 _marketType)
func (_Augur *AugurSession) GetMarketType(_market common.Address) (uint8, error) {
	return _Augur.Contract.GetMarketType(&_Augur.CallOpts, _market)
}

// GetMarketType is a free data retrieval call binding the contract method 0x9eba4b7b.
//
// Solidity: function getMarketType(address _market) view returns(uint8 _marketType)
func (_Augur *AugurCallerSession) GetMarketType(_market common.Address) (uint8, error) {
	return _Augur.Contract.GetMarketType(&_Augur.CallOpts, _market)
}

// GetTimestamp is a free data retrieval call binding the contract method 0x188ec356.
//
// Solidity: function getTimestamp() view returns(uint256)
func (_Augur *AugurCaller) GetTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Augur.contract.Call(opts, &out, "getTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestamp is a free data retrieval call binding the contract method 0x188ec356.
//
// Solidity: function getTimestamp() view returns(uint256)
func (_Augur *AugurSession) GetTimestamp() (*big.Int, error) {
	return _Augur.Contract.GetTimestamp(&_Augur.CallOpts)
}

// GetTimestamp is a free data retrieval call binding the contract method 0x188ec356.
//
// Solidity: function getTimestamp() view returns(uint256)
func (_Augur *AugurCallerSession) GetTimestamp() (*big.Int, error) {
	return _Augur.Contract.GetTimestamp(&_Augur.CallOpts)
}

// GetTradeInterval is a free data retrieval call binding the contract method 0x79d60f04.
//
// Solidity: function getTradeInterval(uint256 _displayRange, uint256 _numTicks) pure returns(uint256)
func (_Augur *AugurCaller) GetTradeInterval(opts *bind.CallOpts, _displayRange *big.Int, _numTicks *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Augur.contract.Call(opts, &out, "getTradeInterval", _displayRange, _numTicks)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTradeInterval is a free data retrieval call binding the contract method 0x79d60f04.
//
// Solidity: function getTradeInterval(uint256 _displayRange, uint256 _numTicks) pure returns(uint256)
func (_Augur *AugurSession) GetTradeInterval(_displayRange *big.Int, _numTicks *big.Int) (*big.Int, error) {
	return _Augur.Contract.GetTradeInterval(&_Augur.CallOpts, _displayRange, _numTicks)
}

// GetTradeInterval is a free data retrieval call binding the contract method 0x79d60f04.
//
// Solidity: function getTradeInterval(uint256 _displayRange, uint256 _numTicks) pure returns(uint256)
func (_Augur *AugurCallerSession) GetTradeInterval(_displayRange *big.Int, _numTicks *big.Int) (*big.Int, error) {
	return _Augur.Contract.GetTradeInterval(&_Augur.CallOpts, _displayRange, _numTicks)
}

// GetUniverseForkIndex is a free data retrieval call binding the contract method 0xaaf6da20.
//
// Solidity: function getUniverseForkIndex(address _universe) view returns(uint256)
func (_Augur *AugurCaller) GetUniverseForkIndex(opts *bind.CallOpts, _universe common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Augur.contract.Call(opts, &out, "getUniverseForkIndex", _universe)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUniverseForkIndex is a free data retrieval call binding the contract method 0xaaf6da20.
//
// Solidity: function getUniverseForkIndex(address _universe) view returns(uint256)
func (_Augur *AugurSession) GetUniverseForkIndex(_universe common.Address) (*big.Int, error) {
	return _Augur.Contract.GetUniverseForkIndex(&_Augur.CallOpts, _universe)
}

// GetUniverseForkIndex is a free data retrieval call binding the contract method 0xaaf6da20.
//
// Solidity: function getUniverseForkIndex(address _universe) view returns(uint256)
func (_Augur *AugurCallerSession) GetUniverseForkIndex(_universe common.Address) (*big.Int, error) {
	return _Augur.Contract.GetUniverseForkIndex(&_Augur.CallOpts, _universe)
}

// IsKnownCrowdsourcer is a free data retrieval call binding the contract method 0xb70da7dc.
//
// Solidity: function isKnownCrowdsourcer(address _crowdsourcer) view returns(bool)
func (_Augur *AugurCaller) IsKnownCrowdsourcer(opts *bind.CallOpts, _crowdsourcer common.Address) (bool, error) {
	var out []interface{}
	err := _Augur.contract.Call(opts, &out, "isKnownCrowdsourcer", _crowdsourcer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKnownCrowdsourcer is a free data retrieval call binding the contract method 0xb70da7dc.
//
// Solidity: function isKnownCrowdsourcer(address _crowdsourcer) view returns(bool)
func (_Augur *AugurSession) IsKnownCrowdsourcer(_crowdsourcer common.Address) (bool, error) {
	return _Augur.Contract.IsKnownCrowdsourcer(&_Augur.CallOpts, _crowdsourcer)
}

// IsKnownCrowdsourcer is a free data retrieval call binding the contract method 0xb70da7dc.
//
// Solidity: function isKnownCrowdsourcer(address _crowdsourcer) view returns(bool)
func (_Augur *AugurCallerSession) IsKnownCrowdsourcer(_crowdsourcer common.Address) (bool, error) {
	return _Augur.Contract.IsKnownCrowdsourcer(&_Augur.CallOpts, _crowdsourcer)
}

// IsKnownFeeSender is a free data retrieval call binding the contract method 0x5897e663.
//
// Solidity: function isKnownFeeSender(address _feeSender) view returns(bool)
func (_Augur *AugurCaller) IsKnownFeeSender(opts *bind.CallOpts, _feeSender common.Address) (bool, error) {
	var out []interface{}
	err := _Augur.contract.Call(opts, &out, "isKnownFeeSender", _feeSender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKnownFeeSender is a free data retrieval call binding the contract method 0x5897e663.
//
// Solidity: function isKnownFeeSender(address _feeSender) view returns(bool)
func (_Augur *AugurSession) IsKnownFeeSender(_feeSender common.Address) (bool, error) {
	return _Augur.Contract.IsKnownFeeSender(&_Augur.CallOpts, _feeSender)
}

// IsKnownFeeSender is a free data retrieval call binding the contract method 0x5897e663.
//
// Solidity: function isKnownFeeSender(address _feeSender) view returns(bool)
func (_Augur *AugurCallerSession) IsKnownFeeSender(_feeSender common.Address) (bool, error) {
	return _Augur.Contract.IsKnownFeeSender(&_Augur.CallOpts, _feeSender)
}

// IsKnownMarket is a free data retrieval call binding the contract method 0xe62b8889.
//
// Solidity: function isKnownMarket(address _market) view returns(bool)
func (_Augur *AugurCaller) IsKnownMarket(opts *bind.CallOpts, _market common.Address) (bool, error) {
	var out []interface{}
	err := _Augur.contract.Call(opts, &out, "isKnownMarket", _market)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKnownMarket is a free data retrieval call binding the contract method 0xe62b8889.
//
// Solidity: function isKnownMarket(address _market) view returns(bool)
func (_Augur *AugurSession) IsKnownMarket(_market common.Address) (bool, error) {
	return _Augur.Contract.IsKnownMarket(&_Augur.CallOpts, _market)
}

// IsKnownMarket is a free data retrieval call binding the contract method 0xe62b8889.
//
// Solidity: function isKnownMarket(address _market) view returns(bool)
func (_Augur *AugurCallerSession) IsKnownMarket(_market common.Address) (bool, error) {
	return _Augur.Contract.IsKnownMarket(&_Augur.CallOpts, _market)
}

// IsKnownUniverse is a free data retrieval call binding the contract method 0x8cfb8f21.
//
// Solidity: function isKnownUniverse(address _universe) view returns(bool)
func (_Augur *AugurCaller) IsKnownUniverse(opts *bind.CallOpts, _universe common.Address) (bool, error) {
	var out []interface{}
	err := _Augur.contract.Call(opts, &out, "isKnownUniverse", _universe)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKnownUniverse is a free data retrieval call binding the contract method 0x8cfb8f21.
//
// Solidity: function isKnownUniverse(address _universe) view returns(bool)
func (_Augur *AugurSession) IsKnownUniverse(_universe common.Address) (bool, error) {
	return _Augur.Contract.IsKnownUniverse(&_Augur.CallOpts, _universe)
}

// IsKnownUniverse is a free data retrieval call binding the contract method 0x8cfb8f21.
//
// Solidity: function isKnownUniverse(address _universe) view returns(bool)
func (_Augur *AugurCallerSession) IsKnownUniverse(_universe common.Address) (bool, error) {
	return _Augur.Contract.IsKnownUniverse(&_Augur.CallOpts, _universe)
}

// Lookup is a free data retrieval call binding the contract method 0xf39ec1f7.
//
// Solidity: function lookup(bytes32 _key) view returns(address)
func (_Augur *AugurCaller) Lookup(opts *bind.CallOpts, _key [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Augur.contract.Call(opts, &out, "lookup", _key)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Lookup is a free data retrieval call binding the contract method 0xf39ec1f7.
//
// Solidity: function lookup(bytes32 _key) view returns(address)
func (_Augur *AugurSession) Lookup(_key [32]byte) (common.Address, error) {
	return _Augur.Contract.Lookup(&_Augur.CallOpts, _key)
}

// Lookup is a free data retrieval call binding the contract method 0xf39ec1f7.
//
// Solidity: function lookup(bytes32 _key) view returns(address)
func (_Augur *AugurCallerSession) Lookup(_key [32]byte) (common.Address, error) {
	return _Augur.Contract.Lookup(&_Augur.CallOpts, _key)
}

// Time is a free data retrieval call binding the contract method 0x16ada547.
//
// Solidity: function time() view returns(address)
func (_Augur *AugurCaller) Time(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Augur.contract.Call(opts, &out, "time")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Time is a free data retrieval call binding the contract method 0x16ada547.
//
// Solidity: function time() view returns(address)
func (_Augur *AugurSession) Time() (common.Address, error) {
	return _Augur.Contract.Time(&_Augur.CallOpts)
}

// Time is a free data retrieval call binding the contract method 0x16ada547.
//
// Solidity: function time() view returns(address)
func (_Augur *AugurCallerSession) Time() (common.Address, error) {
	return _Augur.Contract.Time(&_Augur.CallOpts)
}

// UpgradeTimestamp is a free data retrieval call binding the contract method 0x921bd6f0.
//
// Solidity: function upgradeTimestamp() view returns(uint256)
func (_Augur *AugurCaller) UpgradeTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Augur.contract.Call(opts, &out, "upgradeTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpgradeTimestamp is a free data retrieval call binding the contract method 0x921bd6f0.
//
// Solidity: function upgradeTimestamp() view returns(uint256)
func (_Augur *AugurSession) UpgradeTimestamp() (*big.Int, error) {
	return _Augur.Contract.UpgradeTimestamp(&_Augur.CallOpts)
}

// UpgradeTimestamp is a free data retrieval call binding the contract method 0x921bd6f0.
//
// Solidity: function upgradeTimestamp() view returns(uint256)
func (_Augur *AugurCallerSession) UpgradeTimestamp() (*big.Int, error) {
	return _Augur.Contract.UpgradeTimestamp(&_Augur.CallOpts)
}

// Uploader is a free data retrieval call binding the contract method 0x65fe2a0b.
//
// Solidity: function uploader() view returns(address)
func (_Augur *AugurCaller) Uploader(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Augur.contract.Call(opts, &out, "uploader")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Uploader is a free data retrieval call binding the contract method 0x65fe2a0b.
//
// Solidity: function uploader() view returns(address)
func (_Augur *AugurSession) Uploader() (common.Address, error) {
	return _Augur.Contract.Uploader(&_Augur.CallOpts)
}

// Uploader is a free data retrieval call binding the contract method 0x65fe2a0b.
//
// Solidity: function uploader() view returns(address)
func (_Augur *AugurCallerSession) Uploader() (common.Address, error) {
	return _Augur.Contract.Uploader(&_Augur.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Augur *AugurCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Augur.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Augur *AugurSession) Vat() (common.Address, error) {
	return _Augur.Contract.Vat(&_Augur.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Augur *AugurCallerSession) Vat() (common.Address, error) {
	return _Augur.Contract.Vat(&_Augur.CallOpts)
}

// CreateChildUniverse is a paid mutator transaction binding the contract method 0x4f67af24.
//
// Solidity: function createChildUniverse(bytes32 _parentPayoutDistributionHash, uint256[] _parentPayoutNumerators) returns(address)
func (_Augur *AugurTransactor) CreateChildUniverse(opts *bind.TransactOpts, _parentPayoutDistributionHash [32]byte, _parentPayoutNumerators []*big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "createChildUniverse", _parentPayoutDistributionHash, _parentPayoutNumerators)
}

// CreateChildUniverse is a paid mutator transaction binding the contract method 0x4f67af24.
//
// Solidity: function createChildUniverse(bytes32 _parentPayoutDistributionHash, uint256[] _parentPayoutNumerators) returns(address)
func (_Augur *AugurSession) CreateChildUniverse(_parentPayoutDistributionHash [32]byte, _parentPayoutNumerators []*big.Int) (*types.Transaction, error) {
	return _Augur.Contract.CreateChildUniverse(&_Augur.TransactOpts, _parentPayoutDistributionHash, _parentPayoutNumerators)
}

// CreateChildUniverse is a paid mutator transaction binding the contract method 0x4f67af24.
//
// Solidity: function createChildUniverse(bytes32 _parentPayoutDistributionHash, uint256[] _parentPayoutNumerators) returns(address)
func (_Augur *AugurTransactorSession) CreateChildUniverse(_parentPayoutDistributionHash [32]byte, _parentPayoutNumerators []*big.Int) (*types.Transaction, error) {
	return _Augur.Contract.CreateChildUniverse(&_Augur.TransactOpts, _parentPayoutDistributionHash, _parentPayoutNumerators)
}

// CreateGenesisUniverse is a paid mutator transaction binding the contract method 0x9684da1a.
//
// Solidity: function createGenesisUniverse() returns(address)
func (_Augur *AugurTransactor) CreateGenesisUniverse(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "createGenesisUniverse")
}

// CreateGenesisUniverse is a paid mutator transaction binding the contract method 0x9684da1a.
//
// Solidity: function createGenesisUniverse() returns(address)
func (_Augur *AugurSession) CreateGenesisUniverse() (*types.Transaction, error) {
	return _Augur.Contract.CreateGenesisUniverse(&_Augur.TransactOpts)
}

// CreateGenesisUniverse is a paid mutator transaction binding the contract method 0x9684da1a.
//
// Solidity: function createGenesisUniverse() returns(address)
func (_Augur *AugurTransactorSession) CreateGenesisUniverse() (*types.Transaction, error) {
	return _Augur.Contract.CreateGenesisUniverse(&_Augur.TransactOpts)
}

// DisputeCrowdsourcerCreated is a paid mutator transaction binding the contract method 0x5fe4fccd.
//
// Solidity: function disputeCrowdsourcerCreated(address _universe, address _market, address _disputeCrowdsourcer, uint256[] _payoutNumerators, uint256 _size, uint256 _disputeRound) returns(bool)
func (_Augur *AugurTransactor) DisputeCrowdsourcerCreated(opts *bind.TransactOpts, _universe common.Address, _market common.Address, _disputeCrowdsourcer common.Address, _payoutNumerators []*big.Int, _size *big.Int, _disputeRound *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "disputeCrowdsourcerCreated", _universe, _market, _disputeCrowdsourcer, _payoutNumerators, _size, _disputeRound)
}

// DisputeCrowdsourcerCreated is a paid mutator transaction binding the contract method 0x5fe4fccd.
//
// Solidity: function disputeCrowdsourcerCreated(address _universe, address _market, address _disputeCrowdsourcer, uint256[] _payoutNumerators, uint256 _size, uint256 _disputeRound) returns(bool)
func (_Augur *AugurSession) DisputeCrowdsourcerCreated(_universe common.Address, _market common.Address, _disputeCrowdsourcer common.Address, _payoutNumerators []*big.Int, _size *big.Int, _disputeRound *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.DisputeCrowdsourcerCreated(&_Augur.TransactOpts, _universe, _market, _disputeCrowdsourcer, _payoutNumerators, _size, _disputeRound)
}

// DisputeCrowdsourcerCreated is a paid mutator transaction binding the contract method 0x5fe4fccd.
//
// Solidity: function disputeCrowdsourcerCreated(address _universe, address _market, address _disputeCrowdsourcer, uint256[] _payoutNumerators, uint256 _size, uint256 _disputeRound) returns(bool)
func (_Augur *AugurTransactorSession) DisputeCrowdsourcerCreated(_universe common.Address, _market common.Address, _disputeCrowdsourcer common.Address, _payoutNumerators []*big.Int, _size *big.Int, _disputeRound *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.DisputeCrowdsourcerCreated(&_Augur.TransactOpts, _universe, _market, _disputeCrowdsourcer, _payoutNumerators, _size, _disputeRound)
}

// FinishDeployment is a paid mutator transaction binding the contract method 0xe997fffd.
//
// Solidity: function finishDeployment() returns(bool)
func (_Augur *AugurTransactor) FinishDeployment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "finishDeployment")
}

// FinishDeployment is a paid mutator transaction binding the contract method 0xe997fffd.
//
// Solidity: function finishDeployment() returns(bool)
func (_Augur *AugurSession) FinishDeployment() (*types.Transaction, error) {
	return _Augur.Contract.FinishDeployment(&_Augur.TransactOpts)
}

// FinishDeployment is a paid mutator transaction binding the contract method 0xe997fffd.
//
// Solidity: function finishDeployment() returns(bool)
func (_Augur *AugurTransactorSession) FinishDeployment() (*types.Transaction, error) {
	return _Augur.Contract.FinishDeployment(&_Augur.TransactOpts)
}

// GetMaximumMarketEndDate is a paid mutator transaction binding the contract method 0x484f6e3e.
//
// Solidity: function getMaximumMarketEndDate() returns(uint256)
func (_Augur *AugurTransactor) GetMaximumMarketEndDate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "getMaximumMarketEndDate")
}

// GetMaximumMarketEndDate is a paid mutator transaction binding the contract method 0x484f6e3e.
//
// Solidity: function getMaximumMarketEndDate() returns(uint256)
func (_Augur *AugurSession) GetMaximumMarketEndDate() (*types.Transaction, error) {
	return _Augur.Contract.GetMaximumMarketEndDate(&_Augur.TransactOpts)
}

// GetMaximumMarketEndDate is a paid mutator transaction binding the contract method 0x484f6e3e.
//
// Solidity: function getMaximumMarketEndDate() returns(uint256)
func (_Augur *AugurTransactorSession) GetMaximumMarketEndDate() (*types.Transaction, error) {
	return _Augur.Contract.GetMaximumMarketEndDate(&_Augur.TransactOpts)
}

// IsTrustedSender is a paid mutator transaction binding the contract method 0x22415f60.
//
// Solidity: function isTrustedSender(address _address) returns(bool)
func (_Augur *AugurTransactor) IsTrustedSender(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "isTrustedSender", _address)
}

// IsTrustedSender is a paid mutator transaction binding the contract method 0x22415f60.
//
// Solidity: function isTrustedSender(address _address) returns(bool)
func (_Augur *AugurSession) IsTrustedSender(_address common.Address) (*types.Transaction, error) {
	return _Augur.Contract.IsTrustedSender(&_Augur.TransactOpts, _address)
}

// IsTrustedSender is a paid mutator transaction binding the contract method 0x22415f60.
//
// Solidity: function isTrustedSender(address _address) returns(bool)
func (_Augur *AugurTransactorSession) IsTrustedSender(_address common.Address) (*types.Transaction, error) {
	return _Augur.Contract.IsTrustedSender(&_Augur.TransactOpts, _address)
}

// LogCompleteSetsPurchased is a paid mutator transaction binding the contract method 0xc509d0b2.
//
// Solidity: function logCompleteSetsPurchased(address _universe, address _market, address _account, uint256 _numCompleteSets) returns(bool)
func (_Augur *AugurTransactor) LogCompleteSetsPurchased(opts *bind.TransactOpts, _universe common.Address, _market common.Address, _account common.Address, _numCompleteSets *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logCompleteSetsPurchased", _universe, _market, _account, _numCompleteSets)
}

// LogCompleteSetsPurchased is a paid mutator transaction binding the contract method 0xc509d0b2.
//
// Solidity: function logCompleteSetsPurchased(address _universe, address _market, address _account, uint256 _numCompleteSets) returns(bool)
func (_Augur *AugurSession) LogCompleteSetsPurchased(_universe common.Address, _market common.Address, _account common.Address, _numCompleteSets *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogCompleteSetsPurchased(&_Augur.TransactOpts, _universe, _market, _account, _numCompleteSets)
}

// LogCompleteSetsPurchased is a paid mutator transaction binding the contract method 0xc509d0b2.
//
// Solidity: function logCompleteSetsPurchased(address _universe, address _market, address _account, uint256 _numCompleteSets) returns(bool)
func (_Augur *AugurTransactorSession) LogCompleteSetsPurchased(_universe common.Address, _market common.Address, _account common.Address, _numCompleteSets *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogCompleteSetsPurchased(&_Augur.TransactOpts, _universe, _market, _account, _numCompleteSets)
}

// LogCompleteSetsSold is a paid mutator transaction binding the contract method 0x7920757e.
//
// Solidity: function logCompleteSetsSold(address _universe, address _market, address _account, uint256 _numCompleteSets, uint256 _fees) returns(bool)
func (_Augur *AugurTransactor) LogCompleteSetsSold(opts *bind.TransactOpts, _universe common.Address, _market common.Address, _account common.Address, _numCompleteSets *big.Int, _fees *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logCompleteSetsSold", _universe, _market, _account, _numCompleteSets, _fees)
}

// LogCompleteSetsSold is a paid mutator transaction binding the contract method 0x7920757e.
//
// Solidity: function logCompleteSetsSold(address _universe, address _market, address _account, uint256 _numCompleteSets, uint256 _fees) returns(bool)
func (_Augur *AugurSession) LogCompleteSetsSold(_universe common.Address, _market common.Address, _account common.Address, _numCompleteSets *big.Int, _fees *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogCompleteSetsSold(&_Augur.TransactOpts, _universe, _market, _account, _numCompleteSets, _fees)
}

// LogCompleteSetsSold is a paid mutator transaction binding the contract method 0x7920757e.
//
// Solidity: function logCompleteSetsSold(address _universe, address _market, address _account, uint256 _numCompleteSets, uint256 _fees) returns(bool)
func (_Augur *AugurTransactorSession) LogCompleteSetsSold(_universe common.Address, _market common.Address, _account common.Address, _numCompleteSets *big.Int, _fees *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogCompleteSetsSold(&_Augur.TransactOpts, _universe, _market, _account, _numCompleteSets, _fees)
}

// LogDesignatedReportStakeChanged is a paid mutator transaction binding the contract method 0x973df5b9.
//
// Solidity: function logDesignatedReportStakeChanged(uint256 _designatedReportStake) returns(bool)
func (_Augur *AugurTransactor) LogDesignatedReportStakeChanged(opts *bind.TransactOpts, _designatedReportStake *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logDesignatedReportStakeChanged", _designatedReportStake)
}

// LogDesignatedReportStakeChanged is a paid mutator transaction binding the contract method 0x973df5b9.
//
// Solidity: function logDesignatedReportStakeChanged(uint256 _designatedReportStake) returns(bool)
func (_Augur *AugurSession) LogDesignatedReportStakeChanged(_designatedReportStake *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogDesignatedReportStakeChanged(&_Augur.TransactOpts, _designatedReportStake)
}

// LogDesignatedReportStakeChanged is a paid mutator transaction binding the contract method 0x973df5b9.
//
// Solidity: function logDesignatedReportStakeChanged(uint256 _designatedReportStake) returns(bool)
func (_Augur *AugurTransactorSession) LogDesignatedReportStakeChanged(_designatedReportStake *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogDesignatedReportStakeChanged(&_Augur.TransactOpts, _designatedReportStake)
}

// LogDisputeCrowdsourcerCompleted is a paid mutator transaction binding the contract method 0x6a95e6a7.
//
// Solidity: function logDisputeCrowdsourcerCompleted(address _universe, address _market, address _disputeCrowdsourcer, uint256[] _payoutNumerators, uint256 _nextWindowStartTime, uint256 _nextWindowEndTime, bool _pacingOn, uint256 _totalRepStakedInPayout, uint256 _totalRepStakedInMarket, uint256 _disputeRound) returns(bool)
func (_Augur *AugurTransactor) LogDisputeCrowdsourcerCompleted(opts *bind.TransactOpts, _universe common.Address, _market common.Address, _disputeCrowdsourcer common.Address, _payoutNumerators []*big.Int, _nextWindowStartTime *big.Int, _nextWindowEndTime *big.Int, _pacingOn bool, _totalRepStakedInPayout *big.Int, _totalRepStakedInMarket *big.Int, _disputeRound *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logDisputeCrowdsourcerCompleted", _universe, _market, _disputeCrowdsourcer, _payoutNumerators, _nextWindowStartTime, _nextWindowEndTime, _pacingOn, _totalRepStakedInPayout, _totalRepStakedInMarket, _disputeRound)
}

// LogDisputeCrowdsourcerCompleted is a paid mutator transaction binding the contract method 0x6a95e6a7.
//
// Solidity: function logDisputeCrowdsourcerCompleted(address _universe, address _market, address _disputeCrowdsourcer, uint256[] _payoutNumerators, uint256 _nextWindowStartTime, uint256 _nextWindowEndTime, bool _pacingOn, uint256 _totalRepStakedInPayout, uint256 _totalRepStakedInMarket, uint256 _disputeRound) returns(bool)
func (_Augur *AugurSession) LogDisputeCrowdsourcerCompleted(_universe common.Address, _market common.Address, _disputeCrowdsourcer common.Address, _payoutNumerators []*big.Int, _nextWindowStartTime *big.Int, _nextWindowEndTime *big.Int, _pacingOn bool, _totalRepStakedInPayout *big.Int, _totalRepStakedInMarket *big.Int, _disputeRound *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogDisputeCrowdsourcerCompleted(&_Augur.TransactOpts, _universe, _market, _disputeCrowdsourcer, _payoutNumerators, _nextWindowStartTime, _nextWindowEndTime, _pacingOn, _totalRepStakedInPayout, _totalRepStakedInMarket, _disputeRound)
}

// LogDisputeCrowdsourcerCompleted is a paid mutator transaction binding the contract method 0x6a95e6a7.
//
// Solidity: function logDisputeCrowdsourcerCompleted(address _universe, address _market, address _disputeCrowdsourcer, uint256[] _payoutNumerators, uint256 _nextWindowStartTime, uint256 _nextWindowEndTime, bool _pacingOn, uint256 _totalRepStakedInPayout, uint256 _totalRepStakedInMarket, uint256 _disputeRound) returns(bool)
func (_Augur *AugurTransactorSession) LogDisputeCrowdsourcerCompleted(_universe common.Address, _market common.Address, _disputeCrowdsourcer common.Address, _payoutNumerators []*big.Int, _nextWindowStartTime *big.Int, _nextWindowEndTime *big.Int, _pacingOn bool, _totalRepStakedInPayout *big.Int, _totalRepStakedInMarket *big.Int, _disputeRound *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogDisputeCrowdsourcerCompleted(&_Augur.TransactOpts, _universe, _market, _disputeCrowdsourcer, _payoutNumerators, _nextWindowStartTime, _nextWindowEndTime, _pacingOn, _totalRepStakedInPayout, _totalRepStakedInMarket, _disputeRound)
}

// LogDisputeCrowdsourcerContribution is a paid mutator transaction binding the contract method 0x70b78eba.
//
// Solidity: function logDisputeCrowdsourcerContribution(address _universe, address _reporter, address _market, address _disputeCrowdsourcer, uint256 _amountStaked, string _description, uint256[] _payoutNumerators, uint256 _currentStake, uint256 _stakeRemaining, uint256 _disputeRound) returns(bool)
func (_Augur *AugurTransactor) LogDisputeCrowdsourcerContribution(opts *bind.TransactOpts, _universe common.Address, _reporter common.Address, _market common.Address, _disputeCrowdsourcer common.Address, _amountStaked *big.Int, _description string, _payoutNumerators []*big.Int, _currentStake *big.Int, _stakeRemaining *big.Int, _disputeRound *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logDisputeCrowdsourcerContribution", _universe, _reporter, _market, _disputeCrowdsourcer, _amountStaked, _description, _payoutNumerators, _currentStake, _stakeRemaining, _disputeRound)
}

// LogDisputeCrowdsourcerContribution is a paid mutator transaction binding the contract method 0x70b78eba.
//
// Solidity: function logDisputeCrowdsourcerContribution(address _universe, address _reporter, address _market, address _disputeCrowdsourcer, uint256 _amountStaked, string _description, uint256[] _payoutNumerators, uint256 _currentStake, uint256 _stakeRemaining, uint256 _disputeRound) returns(bool)
func (_Augur *AugurSession) LogDisputeCrowdsourcerContribution(_universe common.Address, _reporter common.Address, _market common.Address, _disputeCrowdsourcer common.Address, _amountStaked *big.Int, _description string, _payoutNumerators []*big.Int, _currentStake *big.Int, _stakeRemaining *big.Int, _disputeRound *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogDisputeCrowdsourcerContribution(&_Augur.TransactOpts, _universe, _reporter, _market, _disputeCrowdsourcer, _amountStaked, _description, _payoutNumerators, _currentStake, _stakeRemaining, _disputeRound)
}

// LogDisputeCrowdsourcerContribution is a paid mutator transaction binding the contract method 0x70b78eba.
//
// Solidity: function logDisputeCrowdsourcerContribution(address _universe, address _reporter, address _market, address _disputeCrowdsourcer, uint256 _amountStaked, string _description, uint256[] _payoutNumerators, uint256 _currentStake, uint256 _stakeRemaining, uint256 _disputeRound) returns(bool)
func (_Augur *AugurTransactorSession) LogDisputeCrowdsourcerContribution(_universe common.Address, _reporter common.Address, _market common.Address, _disputeCrowdsourcer common.Address, _amountStaked *big.Int, _description string, _payoutNumerators []*big.Int, _currentStake *big.Int, _stakeRemaining *big.Int, _disputeRound *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogDisputeCrowdsourcerContribution(&_Augur.TransactOpts, _universe, _reporter, _market, _disputeCrowdsourcer, _amountStaked, _description, _payoutNumerators, _currentStake, _stakeRemaining, _disputeRound)
}

// LogDisputeCrowdsourcerRedeemed is a paid mutator transaction binding the contract method 0x80675fdc.
//
// Solidity: function logDisputeCrowdsourcerRedeemed(address _universe, address _reporter, address _market, uint256 _amountRedeemed, uint256 _repReceived, uint256[] _payoutNumerators) returns(bool)
func (_Augur *AugurTransactor) LogDisputeCrowdsourcerRedeemed(opts *bind.TransactOpts, _universe common.Address, _reporter common.Address, _market common.Address, _amountRedeemed *big.Int, _repReceived *big.Int, _payoutNumerators []*big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logDisputeCrowdsourcerRedeemed", _universe, _reporter, _market, _amountRedeemed, _repReceived, _payoutNumerators)
}

// LogDisputeCrowdsourcerRedeemed is a paid mutator transaction binding the contract method 0x80675fdc.
//
// Solidity: function logDisputeCrowdsourcerRedeemed(address _universe, address _reporter, address _market, uint256 _amountRedeemed, uint256 _repReceived, uint256[] _payoutNumerators) returns(bool)
func (_Augur *AugurSession) LogDisputeCrowdsourcerRedeemed(_universe common.Address, _reporter common.Address, _market common.Address, _amountRedeemed *big.Int, _repReceived *big.Int, _payoutNumerators []*big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogDisputeCrowdsourcerRedeemed(&_Augur.TransactOpts, _universe, _reporter, _market, _amountRedeemed, _repReceived, _payoutNumerators)
}

// LogDisputeCrowdsourcerRedeemed is a paid mutator transaction binding the contract method 0x80675fdc.
//
// Solidity: function logDisputeCrowdsourcerRedeemed(address _universe, address _reporter, address _market, uint256 _amountRedeemed, uint256 _repReceived, uint256[] _payoutNumerators) returns(bool)
func (_Augur *AugurTransactorSession) LogDisputeCrowdsourcerRedeemed(_universe common.Address, _reporter common.Address, _market common.Address, _amountRedeemed *big.Int, _repReceived *big.Int, _payoutNumerators []*big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogDisputeCrowdsourcerRedeemed(&_Augur.TransactOpts, _universe, _reporter, _market, _amountRedeemed, _repReceived, _payoutNumerators)
}

// LogDisputeCrowdsourcerTokensBurned is a paid mutator transaction binding the contract method 0x1142e31e.
//
// Solidity: function logDisputeCrowdsourcerTokensBurned(address _universe, address _target, uint256 _amount, uint256 _totalSupply, uint256 _balance) returns(bool)
func (_Augur *AugurTransactor) LogDisputeCrowdsourcerTokensBurned(opts *bind.TransactOpts, _universe common.Address, _target common.Address, _amount *big.Int, _totalSupply *big.Int, _balance *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logDisputeCrowdsourcerTokensBurned", _universe, _target, _amount, _totalSupply, _balance)
}

// LogDisputeCrowdsourcerTokensBurned is a paid mutator transaction binding the contract method 0x1142e31e.
//
// Solidity: function logDisputeCrowdsourcerTokensBurned(address _universe, address _target, uint256 _amount, uint256 _totalSupply, uint256 _balance) returns(bool)
func (_Augur *AugurSession) LogDisputeCrowdsourcerTokensBurned(_universe common.Address, _target common.Address, _amount *big.Int, _totalSupply *big.Int, _balance *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogDisputeCrowdsourcerTokensBurned(&_Augur.TransactOpts, _universe, _target, _amount, _totalSupply, _balance)
}

// LogDisputeCrowdsourcerTokensBurned is a paid mutator transaction binding the contract method 0x1142e31e.
//
// Solidity: function logDisputeCrowdsourcerTokensBurned(address _universe, address _target, uint256 _amount, uint256 _totalSupply, uint256 _balance) returns(bool)
func (_Augur *AugurTransactorSession) LogDisputeCrowdsourcerTokensBurned(_universe common.Address, _target common.Address, _amount *big.Int, _totalSupply *big.Int, _balance *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogDisputeCrowdsourcerTokensBurned(&_Augur.TransactOpts, _universe, _target, _amount, _totalSupply, _balance)
}

// LogDisputeCrowdsourcerTokensMinted is a paid mutator transaction binding the contract method 0x8f669d87.
//
// Solidity: function logDisputeCrowdsourcerTokensMinted(address _universe, address _target, uint256 _amount, uint256 _totalSupply, uint256 _balance) returns(bool)
func (_Augur *AugurTransactor) LogDisputeCrowdsourcerTokensMinted(opts *bind.TransactOpts, _universe common.Address, _target common.Address, _amount *big.Int, _totalSupply *big.Int, _balance *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logDisputeCrowdsourcerTokensMinted", _universe, _target, _amount, _totalSupply, _balance)
}

// LogDisputeCrowdsourcerTokensMinted is a paid mutator transaction binding the contract method 0x8f669d87.
//
// Solidity: function logDisputeCrowdsourcerTokensMinted(address _universe, address _target, uint256 _amount, uint256 _totalSupply, uint256 _balance) returns(bool)
func (_Augur *AugurSession) LogDisputeCrowdsourcerTokensMinted(_universe common.Address, _target common.Address, _amount *big.Int, _totalSupply *big.Int, _balance *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogDisputeCrowdsourcerTokensMinted(&_Augur.TransactOpts, _universe, _target, _amount, _totalSupply, _balance)
}

// LogDisputeCrowdsourcerTokensMinted is a paid mutator transaction binding the contract method 0x8f669d87.
//
// Solidity: function logDisputeCrowdsourcerTokensMinted(address _universe, address _target, uint256 _amount, uint256 _totalSupply, uint256 _balance) returns(bool)
func (_Augur *AugurTransactorSession) LogDisputeCrowdsourcerTokensMinted(_universe common.Address, _target common.Address, _amount *big.Int, _totalSupply *big.Int, _balance *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogDisputeCrowdsourcerTokensMinted(&_Augur.TransactOpts, _universe, _target, _amount, _totalSupply, _balance)
}

// LogDisputeCrowdsourcerTokensTransferred is a paid mutator transaction binding the contract method 0x1902540c.
//
// Solidity: function logDisputeCrowdsourcerTokensTransferred(address _universe, address _from, address _to, uint256 _value, uint256 _fromBalance, uint256 _toBalance) returns(bool)
func (_Augur *AugurTransactor) LogDisputeCrowdsourcerTokensTransferred(opts *bind.TransactOpts, _universe common.Address, _from common.Address, _to common.Address, _value *big.Int, _fromBalance *big.Int, _toBalance *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logDisputeCrowdsourcerTokensTransferred", _universe, _from, _to, _value, _fromBalance, _toBalance)
}

// LogDisputeCrowdsourcerTokensTransferred is a paid mutator transaction binding the contract method 0x1902540c.
//
// Solidity: function logDisputeCrowdsourcerTokensTransferred(address _universe, address _from, address _to, uint256 _value, uint256 _fromBalance, uint256 _toBalance) returns(bool)
func (_Augur *AugurSession) LogDisputeCrowdsourcerTokensTransferred(_universe common.Address, _from common.Address, _to common.Address, _value *big.Int, _fromBalance *big.Int, _toBalance *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogDisputeCrowdsourcerTokensTransferred(&_Augur.TransactOpts, _universe, _from, _to, _value, _fromBalance, _toBalance)
}

// LogDisputeCrowdsourcerTokensTransferred is a paid mutator transaction binding the contract method 0x1902540c.
//
// Solidity: function logDisputeCrowdsourcerTokensTransferred(address _universe, address _from, address _to, uint256 _value, uint256 _fromBalance, uint256 _toBalance) returns(bool)
func (_Augur *AugurTransactorSession) LogDisputeCrowdsourcerTokensTransferred(_universe common.Address, _from common.Address, _to common.Address, _value *big.Int, _fromBalance *big.Int, _toBalance *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogDisputeCrowdsourcerTokensTransferred(&_Augur.TransactOpts, _universe, _from, _to, _value, _fromBalance, _toBalance)
}

// LogDisputeWindowCreated is a paid mutator transaction binding the contract method 0x606755be.
//
// Solidity: function logDisputeWindowCreated(address _disputeWindow, uint256 _id, bool _initial) returns(bool)
func (_Augur *AugurTransactor) LogDisputeWindowCreated(opts *bind.TransactOpts, _disputeWindow common.Address, _id *big.Int, _initial bool) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logDisputeWindowCreated", _disputeWindow, _id, _initial)
}

// LogDisputeWindowCreated is a paid mutator transaction binding the contract method 0x606755be.
//
// Solidity: function logDisputeWindowCreated(address _disputeWindow, uint256 _id, bool _initial) returns(bool)
func (_Augur *AugurSession) LogDisputeWindowCreated(_disputeWindow common.Address, _id *big.Int, _initial bool) (*types.Transaction, error) {
	return _Augur.Contract.LogDisputeWindowCreated(&_Augur.TransactOpts, _disputeWindow, _id, _initial)
}

// LogDisputeWindowCreated is a paid mutator transaction binding the contract method 0x606755be.
//
// Solidity: function logDisputeWindowCreated(address _disputeWindow, uint256 _id, bool _initial) returns(bool)
func (_Augur *AugurTransactorSession) LogDisputeWindowCreated(_disputeWindow common.Address, _id *big.Int, _initial bool) (*types.Transaction, error) {
	return _Augur.Contract.LogDisputeWindowCreated(&_Augur.TransactOpts, _disputeWindow, _id, _initial)
}

// LogInitialReportSubmitted is a paid mutator transaction binding the contract method 0xda9d7a48.
//
// Solidity: function logInitialReportSubmitted(address _universe, address _reporter, address _market, address _initialReporter, uint256 _amountStaked, bool _isDesignatedReporter, uint256[] _payoutNumerators, string _description, uint256 _nextWindowStartTime, uint256 _nextWindowEndTime) returns(bool)
func (_Augur *AugurTransactor) LogInitialReportSubmitted(opts *bind.TransactOpts, _universe common.Address, _reporter common.Address, _market common.Address, _initialReporter common.Address, _amountStaked *big.Int, _isDesignatedReporter bool, _payoutNumerators []*big.Int, _description string, _nextWindowStartTime *big.Int, _nextWindowEndTime *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logInitialReportSubmitted", _universe, _reporter, _market, _initialReporter, _amountStaked, _isDesignatedReporter, _payoutNumerators, _description, _nextWindowStartTime, _nextWindowEndTime)
}

// LogInitialReportSubmitted is a paid mutator transaction binding the contract method 0xda9d7a48.
//
// Solidity: function logInitialReportSubmitted(address _universe, address _reporter, address _market, address _initialReporter, uint256 _amountStaked, bool _isDesignatedReporter, uint256[] _payoutNumerators, string _description, uint256 _nextWindowStartTime, uint256 _nextWindowEndTime) returns(bool)
func (_Augur *AugurSession) LogInitialReportSubmitted(_universe common.Address, _reporter common.Address, _market common.Address, _initialReporter common.Address, _amountStaked *big.Int, _isDesignatedReporter bool, _payoutNumerators []*big.Int, _description string, _nextWindowStartTime *big.Int, _nextWindowEndTime *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogInitialReportSubmitted(&_Augur.TransactOpts, _universe, _reporter, _market, _initialReporter, _amountStaked, _isDesignatedReporter, _payoutNumerators, _description, _nextWindowStartTime, _nextWindowEndTime)
}

// LogInitialReportSubmitted is a paid mutator transaction binding the contract method 0xda9d7a48.
//
// Solidity: function logInitialReportSubmitted(address _universe, address _reporter, address _market, address _initialReporter, uint256 _amountStaked, bool _isDesignatedReporter, uint256[] _payoutNumerators, string _description, uint256 _nextWindowStartTime, uint256 _nextWindowEndTime) returns(bool)
func (_Augur *AugurTransactorSession) LogInitialReportSubmitted(_universe common.Address, _reporter common.Address, _market common.Address, _initialReporter common.Address, _amountStaked *big.Int, _isDesignatedReporter bool, _payoutNumerators []*big.Int, _description string, _nextWindowStartTime *big.Int, _nextWindowEndTime *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogInitialReportSubmitted(&_Augur.TransactOpts, _universe, _reporter, _market, _initialReporter, _amountStaked, _isDesignatedReporter, _payoutNumerators, _description, _nextWindowStartTime, _nextWindowEndTime)
}

// LogInitialReporterRedeemed is a paid mutator transaction binding the contract method 0xb1c094fa.
//
// Solidity: function logInitialReporterRedeemed(address _universe, address _reporter, address _market, uint256 _amountRedeemed, uint256 _repReceived, uint256[] _payoutNumerators) returns(bool)
func (_Augur *AugurTransactor) LogInitialReporterRedeemed(opts *bind.TransactOpts, _universe common.Address, _reporter common.Address, _market common.Address, _amountRedeemed *big.Int, _repReceived *big.Int, _payoutNumerators []*big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logInitialReporterRedeemed", _universe, _reporter, _market, _amountRedeemed, _repReceived, _payoutNumerators)
}

// LogInitialReporterRedeemed is a paid mutator transaction binding the contract method 0xb1c094fa.
//
// Solidity: function logInitialReporterRedeemed(address _universe, address _reporter, address _market, uint256 _amountRedeemed, uint256 _repReceived, uint256[] _payoutNumerators) returns(bool)
func (_Augur *AugurSession) LogInitialReporterRedeemed(_universe common.Address, _reporter common.Address, _market common.Address, _amountRedeemed *big.Int, _repReceived *big.Int, _payoutNumerators []*big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogInitialReporterRedeemed(&_Augur.TransactOpts, _universe, _reporter, _market, _amountRedeemed, _repReceived, _payoutNumerators)
}

// LogInitialReporterRedeemed is a paid mutator transaction binding the contract method 0xb1c094fa.
//
// Solidity: function logInitialReporterRedeemed(address _universe, address _reporter, address _market, uint256 _amountRedeemed, uint256 _repReceived, uint256[] _payoutNumerators) returns(bool)
func (_Augur *AugurTransactorSession) LogInitialReporterRedeemed(_universe common.Address, _reporter common.Address, _market common.Address, _amountRedeemed *big.Int, _repReceived *big.Int, _payoutNumerators []*big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogInitialReporterRedeemed(&_Augur.TransactOpts, _universe, _reporter, _market, _amountRedeemed, _repReceived, _payoutNumerators)
}

// LogInitialReporterTransferred is a paid mutator transaction binding the contract method 0xe3142e90.
//
// Solidity: function logInitialReporterTransferred(address _universe, address _market, address _from, address _to) returns(bool)
func (_Augur *AugurTransactor) LogInitialReporterTransferred(opts *bind.TransactOpts, _universe common.Address, _market common.Address, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logInitialReporterTransferred", _universe, _market, _from, _to)
}

// LogInitialReporterTransferred is a paid mutator transaction binding the contract method 0xe3142e90.
//
// Solidity: function logInitialReporterTransferred(address _universe, address _market, address _from, address _to) returns(bool)
func (_Augur *AugurSession) LogInitialReporterTransferred(_universe common.Address, _market common.Address, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogInitialReporterTransferred(&_Augur.TransactOpts, _universe, _market, _from, _to)
}

// LogInitialReporterTransferred is a paid mutator transaction binding the contract method 0xe3142e90.
//
// Solidity: function logInitialReporterTransferred(address _universe, address _market, address _from, address _to) returns(bool)
func (_Augur *AugurTransactorSession) LogInitialReporterTransferred(_universe common.Address, _market common.Address, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogInitialReporterTransferred(&_Augur.TransactOpts, _universe, _market, _from, _to)
}

// LogMarketFinalized is a paid mutator transaction binding the contract method 0x27d8e850.
//
// Solidity: function logMarketFinalized(address _universe, uint256[] _winningPayoutNumerators) returns(bool)
func (_Augur *AugurTransactor) LogMarketFinalized(opts *bind.TransactOpts, _universe common.Address, _winningPayoutNumerators []*big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logMarketFinalized", _universe, _winningPayoutNumerators)
}

// LogMarketFinalized is a paid mutator transaction binding the contract method 0x27d8e850.
//
// Solidity: function logMarketFinalized(address _universe, uint256[] _winningPayoutNumerators) returns(bool)
func (_Augur *AugurSession) LogMarketFinalized(_universe common.Address, _winningPayoutNumerators []*big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogMarketFinalized(&_Augur.TransactOpts, _universe, _winningPayoutNumerators)
}

// LogMarketFinalized is a paid mutator transaction binding the contract method 0x27d8e850.
//
// Solidity: function logMarketFinalized(address _universe, uint256[] _winningPayoutNumerators) returns(bool)
func (_Augur *AugurTransactorSession) LogMarketFinalized(_universe common.Address, _winningPayoutNumerators []*big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogMarketFinalized(&_Augur.TransactOpts, _universe, _winningPayoutNumerators)
}

// LogMarketMigrated is a paid mutator transaction binding the contract method 0x17674e4d.
//
// Solidity: function logMarketMigrated(address _market, address _originalUniverse) returns(bool)
func (_Augur *AugurTransactor) LogMarketMigrated(opts *bind.TransactOpts, _market common.Address, _originalUniverse common.Address) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logMarketMigrated", _market, _originalUniverse)
}

// LogMarketMigrated is a paid mutator transaction binding the contract method 0x17674e4d.
//
// Solidity: function logMarketMigrated(address _market, address _originalUniverse) returns(bool)
func (_Augur *AugurSession) LogMarketMigrated(_market common.Address, _originalUniverse common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogMarketMigrated(&_Augur.TransactOpts, _market, _originalUniverse)
}

// LogMarketMigrated is a paid mutator transaction binding the contract method 0x17674e4d.
//
// Solidity: function logMarketMigrated(address _market, address _originalUniverse) returns(bool)
func (_Augur *AugurTransactorSession) LogMarketMigrated(_market common.Address, _originalUniverse common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogMarketMigrated(&_Augur.TransactOpts, _market, _originalUniverse)
}

// LogMarketOIChanged is a paid mutator transaction binding the contract method 0x722e3320.
//
// Solidity: function logMarketOIChanged(address _universe, address _market) returns(bool)
func (_Augur *AugurTransactor) LogMarketOIChanged(opts *bind.TransactOpts, _universe common.Address, _market common.Address) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logMarketOIChanged", _universe, _market)
}

// LogMarketOIChanged is a paid mutator transaction binding the contract method 0x722e3320.
//
// Solidity: function logMarketOIChanged(address _universe, address _market) returns(bool)
func (_Augur *AugurSession) LogMarketOIChanged(_universe common.Address, _market common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogMarketOIChanged(&_Augur.TransactOpts, _universe, _market)
}

// LogMarketOIChanged is a paid mutator transaction binding the contract method 0x722e3320.
//
// Solidity: function logMarketOIChanged(address _universe, address _market) returns(bool)
func (_Augur *AugurTransactorSession) LogMarketOIChanged(_universe common.Address, _market common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogMarketOIChanged(&_Augur.TransactOpts, _universe, _market)
}

// LogMarketParticipantsDisavowed is a paid mutator transaction binding the contract method 0xc67af5cc.
//
// Solidity: function logMarketParticipantsDisavowed(address _universe) returns(bool)
func (_Augur *AugurTransactor) LogMarketParticipantsDisavowed(opts *bind.TransactOpts, _universe common.Address) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logMarketParticipantsDisavowed", _universe)
}

// LogMarketParticipantsDisavowed is a paid mutator transaction binding the contract method 0xc67af5cc.
//
// Solidity: function logMarketParticipantsDisavowed(address _universe) returns(bool)
func (_Augur *AugurSession) LogMarketParticipantsDisavowed(_universe common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogMarketParticipantsDisavowed(&_Augur.TransactOpts, _universe)
}

// LogMarketParticipantsDisavowed is a paid mutator transaction binding the contract method 0xc67af5cc.
//
// Solidity: function logMarketParticipantsDisavowed(address _universe) returns(bool)
func (_Augur *AugurTransactorSession) LogMarketParticipantsDisavowed(_universe common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogMarketParticipantsDisavowed(&_Augur.TransactOpts, _universe)
}

// LogMarketRepBondTransferred is a paid mutator transaction binding the contract method 0x08bb7309.
//
// Solidity: function logMarketRepBondTransferred(address _universe, address _from, address _to) returns(bool)
func (_Augur *AugurTransactor) LogMarketRepBondTransferred(opts *bind.TransactOpts, _universe common.Address, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logMarketRepBondTransferred", _universe, _from, _to)
}

// LogMarketRepBondTransferred is a paid mutator transaction binding the contract method 0x08bb7309.
//
// Solidity: function logMarketRepBondTransferred(address _universe, address _from, address _to) returns(bool)
func (_Augur *AugurSession) LogMarketRepBondTransferred(_universe common.Address, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogMarketRepBondTransferred(&_Augur.TransactOpts, _universe, _from, _to)
}

// LogMarketRepBondTransferred is a paid mutator transaction binding the contract method 0x08bb7309.
//
// Solidity: function logMarketRepBondTransferred(address _universe, address _from, address _to) returns(bool)
func (_Augur *AugurTransactorSession) LogMarketRepBondTransferred(_universe common.Address, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogMarketRepBondTransferred(&_Augur.TransactOpts, _universe, _from, _to)
}

// LogMarketTransferred is a paid mutator transaction binding the contract method 0x23290737.
//
// Solidity: function logMarketTransferred(address _universe, address _from, address _to) returns(bool)
func (_Augur *AugurTransactor) LogMarketTransferred(opts *bind.TransactOpts, _universe common.Address, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logMarketTransferred", _universe, _from, _to)
}

// LogMarketTransferred is a paid mutator transaction binding the contract method 0x23290737.
//
// Solidity: function logMarketTransferred(address _universe, address _from, address _to) returns(bool)
func (_Augur *AugurSession) LogMarketTransferred(_universe common.Address, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogMarketTransferred(&_Augur.TransactOpts, _universe, _from, _to)
}

// LogMarketTransferred is a paid mutator transaction binding the contract method 0x23290737.
//
// Solidity: function logMarketTransferred(address _universe, address _from, address _to) returns(bool)
func (_Augur *AugurTransactorSession) LogMarketTransferred(_universe common.Address, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogMarketTransferred(&_Augur.TransactOpts, _universe, _from, _to)
}

// LogNoShowBondChanged is a paid mutator transaction binding the contract method 0x58c4092c.
//
// Solidity: function logNoShowBondChanged(uint256 _noShowBond) returns(bool)
func (_Augur *AugurTransactor) LogNoShowBondChanged(opts *bind.TransactOpts, _noShowBond *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logNoShowBondChanged", _noShowBond)
}

// LogNoShowBondChanged is a paid mutator transaction binding the contract method 0x58c4092c.
//
// Solidity: function logNoShowBondChanged(uint256 _noShowBond) returns(bool)
func (_Augur *AugurSession) LogNoShowBondChanged(_noShowBond *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogNoShowBondChanged(&_Augur.TransactOpts, _noShowBond)
}

// LogNoShowBondChanged is a paid mutator transaction binding the contract method 0x58c4092c.
//
// Solidity: function logNoShowBondChanged(uint256 _noShowBond) returns(bool)
func (_Augur *AugurTransactorSession) LogNoShowBondChanged(_noShowBond *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogNoShowBondChanged(&_Augur.TransactOpts, _noShowBond)
}

// LogParticipationTokensBurned is a paid mutator transaction binding the contract method 0x59967d1f.
//
// Solidity: function logParticipationTokensBurned(address _universe, address _target, uint256 _amount, uint256 _totalSupply, uint256 _balance) returns(bool)
func (_Augur *AugurTransactor) LogParticipationTokensBurned(opts *bind.TransactOpts, _universe common.Address, _target common.Address, _amount *big.Int, _totalSupply *big.Int, _balance *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logParticipationTokensBurned", _universe, _target, _amount, _totalSupply, _balance)
}

// LogParticipationTokensBurned is a paid mutator transaction binding the contract method 0x59967d1f.
//
// Solidity: function logParticipationTokensBurned(address _universe, address _target, uint256 _amount, uint256 _totalSupply, uint256 _balance) returns(bool)
func (_Augur *AugurSession) LogParticipationTokensBurned(_universe common.Address, _target common.Address, _amount *big.Int, _totalSupply *big.Int, _balance *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogParticipationTokensBurned(&_Augur.TransactOpts, _universe, _target, _amount, _totalSupply, _balance)
}

// LogParticipationTokensBurned is a paid mutator transaction binding the contract method 0x59967d1f.
//
// Solidity: function logParticipationTokensBurned(address _universe, address _target, uint256 _amount, uint256 _totalSupply, uint256 _balance) returns(bool)
func (_Augur *AugurTransactorSession) LogParticipationTokensBurned(_universe common.Address, _target common.Address, _amount *big.Int, _totalSupply *big.Int, _balance *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogParticipationTokensBurned(&_Augur.TransactOpts, _universe, _target, _amount, _totalSupply, _balance)
}

// LogParticipationTokensMinted is a paid mutator transaction binding the contract method 0xecdaa595.
//
// Solidity: function logParticipationTokensMinted(address _universe, address _target, uint256 _amount, uint256 _totalSupply, uint256 _balance) returns(bool)
func (_Augur *AugurTransactor) LogParticipationTokensMinted(opts *bind.TransactOpts, _universe common.Address, _target common.Address, _amount *big.Int, _totalSupply *big.Int, _balance *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logParticipationTokensMinted", _universe, _target, _amount, _totalSupply, _balance)
}

// LogParticipationTokensMinted is a paid mutator transaction binding the contract method 0xecdaa595.
//
// Solidity: function logParticipationTokensMinted(address _universe, address _target, uint256 _amount, uint256 _totalSupply, uint256 _balance) returns(bool)
func (_Augur *AugurSession) LogParticipationTokensMinted(_universe common.Address, _target common.Address, _amount *big.Int, _totalSupply *big.Int, _balance *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogParticipationTokensMinted(&_Augur.TransactOpts, _universe, _target, _amount, _totalSupply, _balance)
}

// LogParticipationTokensMinted is a paid mutator transaction binding the contract method 0xecdaa595.
//
// Solidity: function logParticipationTokensMinted(address _universe, address _target, uint256 _amount, uint256 _totalSupply, uint256 _balance) returns(bool)
func (_Augur *AugurTransactorSession) LogParticipationTokensMinted(_universe common.Address, _target common.Address, _amount *big.Int, _totalSupply *big.Int, _balance *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogParticipationTokensMinted(&_Augur.TransactOpts, _universe, _target, _amount, _totalSupply, _balance)
}

// LogParticipationTokensRedeemed is a paid mutator transaction binding the contract method 0xa3aac84d.
//
// Solidity: function logParticipationTokensRedeemed(address _universe, address _account, uint256 _attoParticipationTokens, uint256 _feePayoutShare) returns(bool)
func (_Augur *AugurTransactor) LogParticipationTokensRedeemed(opts *bind.TransactOpts, _universe common.Address, _account common.Address, _attoParticipationTokens *big.Int, _feePayoutShare *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logParticipationTokensRedeemed", _universe, _account, _attoParticipationTokens, _feePayoutShare)
}

// LogParticipationTokensRedeemed is a paid mutator transaction binding the contract method 0xa3aac84d.
//
// Solidity: function logParticipationTokensRedeemed(address _universe, address _account, uint256 _attoParticipationTokens, uint256 _feePayoutShare) returns(bool)
func (_Augur *AugurSession) LogParticipationTokensRedeemed(_universe common.Address, _account common.Address, _attoParticipationTokens *big.Int, _feePayoutShare *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogParticipationTokensRedeemed(&_Augur.TransactOpts, _universe, _account, _attoParticipationTokens, _feePayoutShare)
}

// LogParticipationTokensRedeemed is a paid mutator transaction binding the contract method 0xa3aac84d.
//
// Solidity: function logParticipationTokensRedeemed(address _universe, address _account, uint256 _attoParticipationTokens, uint256 _feePayoutShare) returns(bool)
func (_Augur *AugurTransactorSession) LogParticipationTokensRedeemed(_universe common.Address, _account common.Address, _attoParticipationTokens *big.Int, _feePayoutShare *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogParticipationTokensRedeemed(&_Augur.TransactOpts, _universe, _account, _attoParticipationTokens, _feePayoutShare)
}

// LogParticipationTokensTransferred is a paid mutator transaction binding the contract method 0x26990346.
//
// Solidity: function logParticipationTokensTransferred(address _universe, address _from, address _to, uint256 _value, uint256 _fromBalance, uint256 _toBalance) returns(bool)
func (_Augur *AugurTransactor) LogParticipationTokensTransferred(opts *bind.TransactOpts, _universe common.Address, _from common.Address, _to common.Address, _value *big.Int, _fromBalance *big.Int, _toBalance *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logParticipationTokensTransferred", _universe, _from, _to, _value, _fromBalance, _toBalance)
}

// LogParticipationTokensTransferred is a paid mutator transaction binding the contract method 0x26990346.
//
// Solidity: function logParticipationTokensTransferred(address _universe, address _from, address _to, uint256 _value, uint256 _fromBalance, uint256 _toBalance) returns(bool)
func (_Augur *AugurSession) LogParticipationTokensTransferred(_universe common.Address, _from common.Address, _to common.Address, _value *big.Int, _fromBalance *big.Int, _toBalance *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogParticipationTokensTransferred(&_Augur.TransactOpts, _universe, _from, _to, _value, _fromBalance, _toBalance)
}

// LogParticipationTokensTransferred is a paid mutator transaction binding the contract method 0x26990346.
//
// Solidity: function logParticipationTokensTransferred(address _universe, address _from, address _to, uint256 _value, uint256 _fromBalance, uint256 _toBalance) returns(bool)
func (_Augur *AugurTransactorSession) LogParticipationTokensTransferred(_universe common.Address, _from common.Address, _to common.Address, _value *big.Int, _fromBalance *big.Int, _toBalance *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogParticipationTokensTransferred(&_Augur.TransactOpts, _universe, _from, _to, _value, _fromBalance, _toBalance)
}

// LogReportingFeeChanged is a paid mutator transaction binding the contract method 0xb394ce2c.
//
// Solidity: function logReportingFeeChanged(uint256 _reportingFee) returns(bool)
func (_Augur *AugurTransactor) LogReportingFeeChanged(opts *bind.TransactOpts, _reportingFee *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logReportingFeeChanged", _reportingFee)
}

// LogReportingFeeChanged is a paid mutator transaction binding the contract method 0xb394ce2c.
//
// Solidity: function logReportingFeeChanged(uint256 _reportingFee) returns(bool)
func (_Augur *AugurSession) LogReportingFeeChanged(_reportingFee *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogReportingFeeChanged(&_Augur.TransactOpts, _reportingFee)
}

// LogReportingFeeChanged is a paid mutator transaction binding the contract method 0xb394ce2c.
//
// Solidity: function logReportingFeeChanged(uint256 _reportingFee) returns(bool)
func (_Augur *AugurTransactorSession) LogReportingFeeChanged(_reportingFee *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogReportingFeeChanged(&_Augur.TransactOpts, _reportingFee)
}

// LogReportingParticipantDisavowed is a paid mutator transaction binding the contract method 0x17570e80.
//
// Solidity: function logReportingParticipantDisavowed(address _universe, address _market) returns(bool)
func (_Augur *AugurTransactor) LogReportingParticipantDisavowed(opts *bind.TransactOpts, _universe common.Address, _market common.Address) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logReportingParticipantDisavowed", _universe, _market)
}

// LogReportingParticipantDisavowed is a paid mutator transaction binding the contract method 0x17570e80.
//
// Solidity: function logReportingParticipantDisavowed(address _universe, address _market) returns(bool)
func (_Augur *AugurSession) LogReportingParticipantDisavowed(_universe common.Address, _market common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogReportingParticipantDisavowed(&_Augur.TransactOpts, _universe, _market)
}

// LogReportingParticipantDisavowed is a paid mutator transaction binding the contract method 0x17570e80.
//
// Solidity: function logReportingParticipantDisavowed(address _universe, address _market) returns(bool)
func (_Augur *AugurTransactorSession) LogReportingParticipantDisavowed(_universe common.Address, _market common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogReportingParticipantDisavowed(&_Augur.TransactOpts, _universe, _market)
}

// LogReputationTokensBurned is a paid mutator transaction binding the contract method 0x3dfdce82.
//
// Solidity: function logReputationTokensBurned(address _universe, address _target, uint256 _amount, uint256 _totalSupply, uint256 _balance) returns(bool)
func (_Augur *AugurTransactor) LogReputationTokensBurned(opts *bind.TransactOpts, _universe common.Address, _target common.Address, _amount *big.Int, _totalSupply *big.Int, _balance *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logReputationTokensBurned", _universe, _target, _amount, _totalSupply, _balance)
}

// LogReputationTokensBurned is a paid mutator transaction binding the contract method 0x3dfdce82.
//
// Solidity: function logReputationTokensBurned(address _universe, address _target, uint256 _amount, uint256 _totalSupply, uint256 _balance) returns(bool)
func (_Augur *AugurSession) LogReputationTokensBurned(_universe common.Address, _target common.Address, _amount *big.Int, _totalSupply *big.Int, _balance *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogReputationTokensBurned(&_Augur.TransactOpts, _universe, _target, _amount, _totalSupply, _balance)
}

// LogReputationTokensBurned is a paid mutator transaction binding the contract method 0x3dfdce82.
//
// Solidity: function logReputationTokensBurned(address _universe, address _target, uint256 _amount, uint256 _totalSupply, uint256 _balance) returns(bool)
func (_Augur *AugurTransactorSession) LogReputationTokensBurned(_universe common.Address, _target common.Address, _amount *big.Int, _totalSupply *big.Int, _balance *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogReputationTokensBurned(&_Augur.TransactOpts, _universe, _target, _amount, _totalSupply, _balance)
}

// LogReputationTokensMinted is a paid mutator transaction binding the contract method 0xd673ad96.
//
// Solidity: function logReputationTokensMinted(address _universe, address _target, uint256 _amount, uint256 _totalSupply, uint256 _balance) returns(bool)
func (_Augur *AugurTransactor) LogReputationTokensMinted(opts *bind.TransactOpts, _universe common.Address, _target common.Address, _amount *big.Int, _totalSupply *big.Int, _balance *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logReputationTokensMinted", _universe, _target, _amount, _totalSupply, _balance)
}

// LogReputationTokensMinted is a paid mutator transaction binding the contract method 0xd673ad96.
//
// Solidity: function logReputationTokensMinted(address _universe, address _target, uint256 _amount, uint256 _totalSupply, uint256 _balance) returns(bool)
func (_Augur *AugurSession) LogReputationTokensMinted(_universe common.Address, _target common.Address, _amount *big.Int, _totalSupply *big.Int, _balance *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogReputationTokensMinted(&_Augur.TransactOpts, _universe, _target, _amount, _totalSupply, _balance)
}

// LogReputationTokensMinted is a paid mutator transaction binding the contract method 0xd673ad96.
//
// Solidity: function logReputationTokensMinted(address _universe, address _target, uint256 _amount, uint256 _totalSupply, uint256 _balance) returns(bool)
func (_Augur *AugurTransactorSession) LogReputationTokensMinted(_universe common.Address, _target common.Address, _amount *big.Int, _totalSupply *big.Int, _balance *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogReputationTokensMinted(&_Augur.TransactOpts, _universe, _target, _amount, _totalSupply, _balance)
}

// LogReputationTokensTransferred is a paid mutator transaction binding the contract method 0xe1c678fe.
//
// Solidity: function logReputationTokensTransferred(address _universe, address _from, address _to, uint256 _value, uint256 _fromBalance, uint256 _toBalance) returns(bool)
func (_Augur *AugurTransactor) LogReputationTokensTransferred(opts *bind.TransactOpts, _universe common.Address, _from common.Address, _to common.Address, _value *big.Int, _fromBalance *big.Int, _toBalance *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logReputationTokensTransferred", _universe, _from, _to, _value, _fromBalance, _toBalance)
}

// LogReputationTokensTransferred is a paid mutator transaction binding the contract method 0xe1c678fe.
//
// Solidity: function logReputationTokensTransferred(address _universe, address _from, address _to, uint256 _value, uint256 _fromBalance, uint256 _toBalance) returns(bool)
func (_Augur *AugurSession) LogReputationTokensTransferred(_universe common.Address, _from common.Address, _to common.Address, _value *big.Int, _fromBalance *big.Int, _toBalance *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogReputationTokensTransferred(&_Augur.TransactOpts, _universe, _from, _to, _value, _fromBalance, _toBalance)
}

// LogReputationTokensTransferred is a paid mutator transaction binding the contract method 0xe1c678fe.
//
// Solidity: function logReputationTokensTransferred(address _universe, address _from, address _to, uint256 _value, uint256 _fromBalance, uint256 _toBalance) returns(bool)
func (_Augur *AugurTransactorSession) LogReputationTokensTransferred(_universe common.Address, _from common.Address, _to common.Address, _value *big.Int, _fromBalance *big.Int, _toBalance *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogReputationTokensTransferred(&_Augur.TransactOpts, _universe, _from, _to, _value, _fromBalance, _toBalance)
}

// LogShareTokensBalanceChanged is a paid mutator transaction binding the contract method 0xd015ddbd.
//
// Solidity: function logShareTokensBalanceChanged(address _account, address _market, uint256 _outcome, uint256 _balance) returns(bool)
func (_Augur *AugurTransactor) LogShareTokensBalanceChanged(opts *bind.TransactOpts, _account common.Address, _market common.Address, _outcome *big.Int, _balance *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logShareTokensBalanceChanged", _account, _market, _outcome, _balance)
}

// LogShareTokensBalanceChanged is a paid mutator transaction binding the contract method 0xd015ddbd.
//
// Solidity: function logShareTokensBalanceChanged(address _account, address _market, uint256 _outcome, uint256 _balance) returns(bool)
func (_Augur *AugurSession) LogShareTokensBalanceChanged(_account common.Address, _market common.Address, _outcome *big.Int, _balance *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogShareTokensBalanceChanged(&_Augur.TransactOpts, _account, _market, _outcome, _balance)
}

// LogShareTokensBalanceChanged is a paid mutator transaction binding the contract method 0xd015ddbd.
//
// Solidity: function logShareTokensBalanceChanged(address _account, address _market, uint256 _outcome, uint256 _balance) returns(bool)
func (_Augur *AugurTransactorSession) LogShareTokensBalanceChanged(_account common.Address, _market common.Address, _outcome *big.Int, _balance *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogShareTokensBalanceChanged(&_Augur.TransactOpts, _account, _market, _outcome, _balance)
}

// LogTimestampSet is a paid mutator transaction binding the contract method 0xc8e6b2a8.
//
// Solidity: function logTimestampSet(uint256 _newTimestamp) returns(bool)
func (_Augur *AugurTransactor) LogTimestampSet(opts *bind.TransactOpts, _newTimestamp *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logTimestampSet", _newTimestamp)
}

// LogTimestampSet is a paid mutator transaction binding the contract method 0xc8e6b2a8.
//
// Solidity: function logTimestampSet(uint256 _newTimestamp) returns(bool)
func (_Augur *AugurSession) LogTimestampSet(_newTimestamp *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogTimestampSet(&_Augur.TransactOpts, _newTimestamp)
}

// LogTimestampSet is a paid mutator transaction binding the contract method 0xc8e6b2a8.
//
// Solidity: function logTimestampSet(uint256 _newTimestamp) returns(bool)
func (_Augur *AugurTransactorSession) LogTimestampSet(_newTimestamp *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogTimestampSet(&_Augur.TransactOpts, _newTimestamp)
}

// LogTradingProceedsClaimed is a paid mutator transaction binding the contract method 0xcec4ede1.
//
// Solidity: function logTradingProceedsClaimed(address _universe, address _sender, address _market, uint256 _outcome, uint256 _numShares, uint256 _numPayoutTokens, uint256 _fees) returns(bool)
func (_Augur *AugurTransactor) LogTradingProceedsClaimed(opts *bind.TransactOpts, _universe common.Address, _sender common.Address, _market common.Address, _outcome *big.Int, _numShares *big.Int, _numPayoutTokens *big.Int, _fees *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logTradingProceedsClaimed", _universe, _sender, _market, _outcome, _numShares, _numPayoutTokens, _fees)
}

// LogTradingProceedsClaimed is a paid mutator transaction binding the contract method 0xcec4ede1.
//
// Solidity: function logTradingProceedsClaimed(address _universe, address _sender, address _market, uint256 _outcome, uint256 _numShares, uint256 _numPayoutTokens, uint256 _fees) returns(bool)
func (_Augur *AugurSession) LogTradingProceedsClaimed(_universe common.Address, _sender common.Address, _market common.Address, _outcome *big.Int, _numShares *big.Int, _numPayoutTokens *big.Int, _fees *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogTradingProceedsClaimed(&_Augur.TransactOpts, _universe, _sender, _market, _outcome, _numShares, _numPayoutTokens, _fees)
}

// LogTradingProceedsClaimed is a paid mutator transaction binding the contract method 0xcec4ede1.
//
// Solidity: function logTradingProceedsClaimed(address _universe, address _sender, address _market, uint256 _outcome, uint256 _numShares, uint256 _numPayoutTokens, uint256 _fees) returns(bool)
func (_Augur *AugurTransactorSession) LogTradingProceedsClaimed(_universe common.Address, _sender common.Address, _market common.Address, _outcome *big.Int, _numShares *big.Int, _numPayoutTokens *big.Int, _fees *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogTradingProceedsClaimed(&_Augur.TransactOpts, _universe, _sender, _market, _outcome, _numShares, _numPayoutTokens, _fees)
}

// LogUniverseForked is a paid mutator transaction binding the contract method 0x8ec6a771.
//
// Solidity: function logUniverseForked(address _forkingMarket) returns(bool)
func (_Augur *AugurTransactor) LogUniverseForked(opts *bind.TransactOpts, _forkingMarket common.Address) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logUniverseForked", _forkingMarket)
}

// LogUniverseForked is a paid mutator transaction binding the contract method 0x8ec6a771.
//
// Solidity: function logUniverseForked(address _forkingMarket) returns(bool)
func (_Augur *AugurSession) LogUniverseForked(_forkingMarket common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogUniverseForked(&_Augur.TransactOpts, _forkingMarket)
}

// LogUniverseForked is a paid mutator transaction binding the contract method 0x8ec6a771.
//
// Solidity: function logUniverseForked(address _forkingMarket) returns(bool)
func (_Augur *AugurTransactorSession) LogUniverseForked(_forkingMarket common.Address) (*types.Transaction, error) {
	return _Augur.Contract.LogUniverseForked(&_Augur.TransactOpts, _forkingMarket)
}

// LogValidityBondChanged is a paid mutator transaction binding the contract method 0xb5a28c4e.
//
// Solidity: function logValidityBondChanged(uint256 _validityBond) returns(bool)
func (_Augur *AugurTransactor) LogValidityBondChanged(opts *bind.TransactOpts, _validityBond *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logValidityBondChanged", _validityBond)
}

// LogValidityBondChanged is a paid mutator transaction binding the contract method 0xb5a28c4e.
//
// Solidity: function logValidityBondChanged(uint256 _validityBond) returns(bool)
func (_Augur *AugurSession) LogValidityBondChanged(_validityBond *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogValidityBondChanged(&_Augur.TransactOpts, _validityBond)
}

// LogValidityBondChanged is a paid mutator transaction binding the contract method 0xb5a28c4e.
//
// Solidity: function logValidityBondChanged(uint256 _validityBond) returns(bool)
func (_Augur *AugurTransactorSession) LogValidityBondChanged(_validityBond *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogValidityBondChanged(&_Augur.TransactOpts, _validityBond)
}

// LogWarpSyncDataUpdated is a paid mutator transaction binding the contract method 0xcf53f72f.
//
// Solidity: function logWarpSyncDataUpdated(address _universe, uint256 _warpSyncHash, uint256 _marketEndTime) returns(bool)
func (_Augur *AugurTransactor) LogWarpSyncDataUpdated(opts *bind.TransactOpts, _universe common.Address, _warpSyncHash *big.Int, _marketEndTime *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "logWarpSyncDataUpdated", _universe, _warpSyncHash, _marketEndTime)
}

// LogWarpSyncDataUpdated is a paid mutator transaction binding the contract method 0xcf53f72f.
//
// Solidity: function logWarpSyncDataUpdated(address _universe, uint256 _warpSyncHash, uint256 _marketEndTime) returns(bool)
func (_Augur *AugurSession) LogWarpSyncDataUpdated(_universe common.Address, _warpSyncHash *big.Int, _marketEndTime *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogWarpSyncDataUpdated(&_Augur.TransactOpts, _universe, _warpSyncHash, _marketEndTime)
}

// LogWarpSyncDataUpdated is a paid mutator transaction binding the contract method 0xcf53f72f.
//
// Solidity: function logWarpSyncDataUpdated(address _universe, uint256 _warpSyncHash, uint256 _marketEndTime) returns(bool)
func (_Augur *AugurTransactorSession) LogWarpSyncDataUpdated(_universe common.Address, _warpSyncHash *big.Int, _marketEndTime *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.LogWarpSyncDataUpdated(&_Augur.TransactOpts, _universe, _warpSyncHash, _marketEndTime)
}

// OnCategoricalMarketCreated is a paid mutator transaction binding the contract method 0x3f6d798a.
//
// Solidity: function onCategoricalMarketCreated(uint256 _endTime, string _extraInfo, address _market, address _marketCreator, address _designatedReporter, uint256 _feePerCashInAttoCash, bytes32[] _outcomes) returns(bool)
func (_Augur *AugurTransactor) OnCategoricalMarketCreated(opts *bind.TransactOpts, _endTime *big.Int, _extraInfo string, _market common.Address, _marketCreator common.Address, _designatedReporter common.Address, _feePerCashInAttoCash *big.Int, _outcomes [][32]byte) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "onCategoricalMarketCreated", _endTime, _extraInfo, _market, _marketCreator, _designatedReporter, _feePerCashInAttoCash, _outcomes)
}

// OnCategoricalMarketCreated is a paid mutator transaction binding the contract method 0x3f6d798a.
//
// Solidity: function onCategoricalMarketCreated(uint256 _endTime, string _extraInfo, address _market, address _marketCreator, address _designatedReporter, uint256 _feePerCashInAttoCash, bytes32[] _outcomes) returns(bool)
func (_Augur *AugurSession) OnCategoricalMarketCreated(_endTime *big.Int, _extraInfo string, _market common.Address, _marketCreator common.Address, _designatedReporter common.Address, _feePerCashInAttoCash *big.Int, _outcomes [][32]byte) (*types.Transaction, error) {
	return _Augur.Contract.OnCategoricalMarketCreated(&_Augur.TransactOpts, _endTime, _extraInfo, _market, _marketCreator, _designatedReporter, _feePerCashInAttoCash, _outcomes)
}

// OnCategoricalMarketCreated is a paid mutator transaction binding the contract method 0x3f6d798a.
//
// Solidity: function onCategoricalMarketCreated(uint256 _endTime, string _extraInfo, address _market, address _marketCreator, address _designatedReporter, uint256 _feePerCashInAttoCash, bytes32[] _outcomes) returns(bool)
func (_Augur *AugurTransactorSession) OnCategoricalMarketCreated(_endTime *big.Int, _extraInfo string, _market common.Address, _marketCreator common.Address, _designatedReporter common.Address, _feePerCashInAttoCash *big.Int, _outcomes [][32]byte) (*types.Transaction, error) {
	return _Augur.Contract.OnCategoricalMarketCreated(&_Augur.TransactOpts, _endTime, _extraInfo, _market, _marketCreator, _designatedReporter, _feePerCashInAttoCash, _outcomes)
}

// OnScalarMarketCreated is a paid mutator transaction binding the contract method 0xed4d5244.
//
// Solidity: function onScalarMarketCreated(uint256 _endTime, string _extraInfo, address _market, address _marketCreator, address _designatedReporter, uint256 _feePerCashInAttoCash, int256[] _prices, uint256 _numTicks) returns(bool)
func (_Augur *AugurTransactor) OnScalarMarketCreated(opts *bind.TransactOpts, _endTime *big.Int, _extraInfo string, _market common.Address, _marketCreator common.Address, _designatedReporter common.Address, _feePerCashInAttoCash *big.Int, _prices []*big.Int, _numTicks *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "onScalarMarketCreated", _endTime, _extraInfo, _market, _marketCreator, _designatedReporter, _feePerCashInAttoCash, _prices, _numTicks)
}

// OnScalarMarketCreated is a paid mutator transaction binding the contract method 0xed4d5244.
//
// Solidity: function onScalarMarketCreated(uint256 _endTime, string _extraInfo, address _market, address _marketCreator, address _designatedReporter, uint256 _feePerCashInAttoCash, int256[] _prices, uint256 _numTicks) returns(bool)
func (_Augur *AugurSession) OnScalarMarketCreated(_endTime *big.Int, _extraInfo string, _market common.Address, _marketCreator common.Address, _designatedReporter common.Address, _feePerCashInAttoCash *big.Int, _prices []*big.Int, _numTicks *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.OnScalarMarketCreated(&_Augur.TransactOpts, _endTime, _extraInfo, _market, _marketCreator, _designatedReporter, _feePerCashInAttoCash, _prices, _numTicks)
}

// OnScalarMarketCreated is a paid mutator transaction binding the contract method 0xed4d5244.
//
// Solidity: function onScalarMarketCreated(uint256 _endTime, string _extraInfo, address _market, address _marketCreator, address _designatedReporter, uint256 _feePerCashInAttoCash, int256[] _prices, uint256 _numTicks) returns(bool)
func (_Augur *AugurTransactorSession) OnScalarMarketCreated(_endTime *big.Int, _extraInfo string, _market common.Address, _marketCreator common.Address, _designatedReporter common.Address, _feePerCashInAttoCash *big.Int, _prices []*big.Int, _numTicks *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.OnScalarMarketCreated(&_Augur.TransactOpts, _endTime, _extraInfo, _market, _marketCreator, _designatedReporter, _feePerCashInAttoCash, _prices, _numTicks)
}

// OnYesNoMarketCreated is a paid mutator transaction binding the contract method 0x2280068d.
//
// Solidity: function onYesNoMarketCreated(uint256 _endTime, string _extraInfo, address _market, address _marketCreator, address _designatedReporter, uint256 _feePerCashInAttoCash) returns(bool)
func (_Augur *AugurTransactor) OnYesNoMarketCreated(opts *bind.TransactOpts, _endTime *big.Int, _extraInfo string, _market common.Address, _marketCreator common.Address, _designatedReporter common.Address, _feePerCashInAttoCash *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "onYesNoMarketCreated", _endTime, _extraInfo, _market, _marketCreator, _designatedReporter, _feePerCashInAttoCash)
}

// OnYesNoMarketCreated is a paid mutator transaction binding the contract method 0x2280068d.
//
// Solidity: function onYesNoMarketCreated(uint256 _endTime, string _extraInfo, address _market, address _marketCreator, address _designatedReporter, uint256 _feePerCashInAttoCash) returns(bool)
func (_Augur *AugurSession) OnYesNoMarketCreated(_endTime *big.Int, _extraInfo string, _market common.Address, _marketCreator common.Address, _designatedReporter common.Address, _feePerCashInAttoCash *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.OnYesNoMarketCreated(&_Augur.TransactOpts, _endTime, _extraInfo, _market, _marketCreator, _designatedReporter, _feePerCashInAttoCash)
}

// OnYesNoMarketCreated is a paid mutator transaction binding the contract method 0x2280068d.
//
// Solidity: function onYesNoMarketCreated(uint256 _endTime, string _extraInfo, address _market, address _marketCreator, address _designatedReporter, uint256 _feePerCashInAttoCash) returns(bool)
func (_Augur *AugurTransactorSession) OnYesNoMarketCreated(_endTime *big.Int, _extraInfo string, _market common.Address, _marketCreator common.Address, _designatedReporter common.Address, _feePerCashInAttoCash *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.OnYesNoMarketCreated(&_Augur.TransactOpts, _endTime, _extraInfo, _market, _marketCreator, _designatedReporter, _feePerCashInAttoCash)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x645c6fae.
//
// Solidity: function registerContract(bytes32 _key, address _address) returns(bool)
func (_Augur *AugurTransactor) RegisterContract(opts *bind.TransactOpts, _key [32]byte, _address common.Address) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "registerContract", _key, _address)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x645c6fae.
//
// Solidity: function registerContract(bytes32 _key, address _address) returns(bool)
func (_Augur *AugurSession) RegisterContract(_key [32]byte, _address common.Address) (*types.Transaction, error) {
	return _Augur.Contract.RegisterContract(&_Augur.TransactOpts, _key, _address)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x645c6fae.
//
// Solidity: function registerContract(bytes32 _key, address _address) returns(bool)
func (_Augur *AugurTransactorSession) RegisterContract(_key [32]byte, _address common.Address) (*types.Transaction, error) {
	return _Augur.Contract.RegisterContract(&_Augur.TransactOpts, _key, _address)
}

// TrustedCashTransfer is a paid mutator transaction binding the contract method 0x6743dcaf.
//
// Solidity: function trustedCashTransfer(address _from, address _to, uint256 _amount) returns(bool)
func (_Augur *AugurTransactor) TrustedCashTransfer(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.contract.Transact(opts, "trustedCashTransfer", _from, _to, _amount)
}

// TrustedCashTransfer is a paid mutator transaction binding the contract method 0x6743dcaf.
//
// Solidity: function trustedCashTransfer(address _from, address _to, uint256 _amount) returns(bool)
func (_Augur *AugurSession) TrustedCashTransfer(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.TrustedCashTransfer(&_Augur.TransactOpts, _from, _to, _amount)
}

// TrustedCashTransfer is a paid mutator transaction binding the contract method 0x6743dcaf.
//
// Solidity: function trustedCashTransfer(address _from, address _to, uint256 _amount) returns(bool)
func (_Augur *AugurTransactorSession) TrustedCashTransfer(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Augur.Contract.TrustedCashTransfer(&_Augur.TransactOpts, _from, _to, _amount)
}

// AugurCompleteSetsPurchasedIterator is returned from FilterCompleteSetsPurchased and is used to iterate over the raw logs and unpacked data for CompleteSetsPurchased events raised by the Augur contract.
type AugurCompleteSetsPurchasedIterator struct {
	Event *AugurCompleteSetsPurchased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurCompleteSetsPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurCompleteSetsPurchased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurCompleteSetsPurchased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurCompleteSetsPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurCompleteSetsPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurCompleteSetsPurchased represents a CompleteSetsPurchased event raised by the Augur contract.
type AugurCompleteSetsPurchased struct {
	Universe        common.Address
	Market          common.Address
	Account         common.Address
	NumCompleteSets *big.Int
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCompleteSetsPurchased is a free log retrieval operation binding the contract event 0xfe06587917de7df83a446bcbb889cee699d7fc35b7b53e263282c2acb5a16499.
//
// Solidity: event CompleteSetsPurchased(address indexed universe, address indexed market, address indexed account, uint256 numCompleteSets, uint256 timestamp)
func (_Augur *AugurFilterer) FilterCompleteSetsPurchased(opts *bind.FilterOpts, universe []common.Address, market []common.Address, account []common.Address) (*AugurCompleteSetsPurchasedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "CompleteSetsPurchased", universeRule, marketRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &AugurCompleteSetsPurchasedIterator{contract: _Augur.contract, event: "CompleteSetsPurchased", logs: logs, sub: sub}, nil
}

// WatchCompleteSetsPurchased is a free log subscription operation binding the contract event 0xfe06587917de7df83a446bcbb889cee699d7fc35b7b53e263282c2acb5a16499.
//
// Solidity: event CompleteSetsPurchased(address indexed universe, address indexed market, address indexed account, uint256 numCompleteSets, uint256 timestamp)
func (_Augur *AugurFilterer) WatchCompleteSetsPurchased(opts *bind.WatchOpts, sink chan<- *AugurCompleteSetsPurchased, universe []common.Address, market []common.Address, account []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "CompleteSetsPurchased", universeRule, marketRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurCompleteSetsPurchased)
				if err := _Augur.contract.UnpackLog(event, "CompleteSetsPurchased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCompleteSetsPurchased is a log parse operation binding the contract event 0xfe06587917de7df83a446bcbb889cee699d7fc35b7b53e263282c2acb5a16499.
//
// Solidity: event CompleteSetsPurchased(address indexed universe, address indexed market, address indexed account, uint256 numCompleteSets, uint256 timestamp)
func (_Augur *AugurFilterer) ParseCompleteSetsPurchased(log types.Log) (*AugurCompleteSetsPurchased, error) {
	event := new(AugurCompleteSetsPurchased)
	if err := _Augur.contract.UnpackLog(event, "CompleteSetsPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurCompleteSetsSoldIterator is returned from FilterCompleteSetsSold and is used to iterate over the raw logs and unpacked data for CompleteSetsSold events raised by the Augur contract.
type AugurCompleteSetsSoldIterator struct {
	Event *AugurCompleteSetsSold // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurCompleteSetsSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurCompleteSetsSold)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurCompleteSetsSold)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurCompleteSetsSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurCompleteSetsSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurCompleteSetsSold represents a CompleteSetsSold event raised by the Augur contract.
type AugurCompleteSetsSold struct {
	Universe        common.Address
	Market          common.Address
	Account         common.Address
	NumCompleteSets *big.Int
	Fees            *big.Int
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCompleteSetsSold is a free log retrieval operation binding the contract event 0xdd7dcfa6708112395eb94e9b1889295fb19af21ef290e918256838c979b2dfbd.
//
// Solidity: event CompleteSetsSold(address indexed universe, address indexed market, address indexed account, uint256 numCompleteSets, uint256 fees, uint256 timestamp)
func (_Augur *AugurFilterer) FilterCompleteSetsSold(opts *bind.FilterOpts, universe []common.Address, market []common.Address, account []common.Address) (*AugurCompleteSetsSoldIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "CompleteSetsSold", universeRule, marketRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &AugurCompleteSetsSoldIterator{contract: _Augur.contract, event: "CompleteSetsSold", logs: logs, sub: sub}, nil
}

// WatchCompleteSetsSold is a free log subscription operation binding the contract event 0xdd7dcfa6708112395eb94e9b1889295fb19af21ef290e918256838c979b2dfbd.
//
// Solidity: event CompleteSetsSold(address indexed universe, address indexed market, address indexed account, uint256 numCompleteSets, uint256 fees, uint256 timestamp)
func (_Augur *AugurFilterer) WatchCompleteSetsSold(opts *bind.WatchOpts, sink chan<- *AugurCompleteSetsSold, universe []common.Address, market []common.Address, account []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "CompleteSetsSold", universeRule, marketRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurCompleteSetsSold)
				if err := _Augur.contract.UnpackLog(event, "CompleteSetsSold", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCompleteSetsSold is a log parse operation binding the contract event 0xdd7dcfa6708112395eb94e9b1889295fb19af21ef290e918256838c979b2dfbd.
//
// Solidity: event CompleteSetsSold(address indexed universe, address indexed market, address indexed account, uint256 numCompleteSets, uint256 fees, uint256 timestamp)
func (_Augur *AugurFilterer) ParseCompleteSetsSold(log types.Log) (*AugurCompleteSetsSold, error) {
	event := new(AugurCompleteSetsSold)
	if err := _Augur.contract.UnpackLog(event, "CompleteSetsSold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurDesignatedReportStakeChangedIterator is returned from FilterDesignatedReportStakeChanged and is used to iterate over the raw logs and unpacked data for DesignatedReportStakeChanged events raised by the Augur contract.
type AugurDesignatedReportStakeChangedIterator struct {
	Event *AugurDesignatedReportStakeChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurDesignatedReportStakeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurDesignatedReportStakeChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurDesignatedReportStakeChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurDesignatedReportStakeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurDesignatedReportStakeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurDesignatedReportStakeChanged represents a DesignatedReportStakeChanged event raised by the Augur contract.
type AugurDesignatedReportStakeChanged struct {
	Universe              common.Address
	DesignatedReportStake *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterDesignatedReportStakeChanged is a free log retrieval operation binding the contract event 0x9c75a088fcb0527d67a80a7d0a5006bbabe02f4b23984234ae68b2b146f001bc.
//
// Solidity: event DesignatedReportStakeChanged(address indexed universe, uint256 designatedReportStake)
func (_Augur *AugurFilterer) FilterDesignatedReportStakeChanged(opts *bind.FilterOpts, universe []common.Address) (*AugurDesignatedReportStakeChangedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "DesignatedReportStakeChanged", universeRule)
	if err != nil {
		return nil, err
	}
	return &AugurDesignatedReportStakeChangedIterator{contract: _Augur.contract, event: "DesignatedReportStakeChanged", logs: logs, sub: sub}, nil
}

// WatchDesignatedReportStakeChanged is a free log subscription operation binding the contract event 0x9c75a088fcb0527d67a80a7d0a5006bbabe02f4b23984234ae68b2b146f001bc.
//
// Solidity: event DesignatedReportStakeChanged(address indexed universe, uint256 designatedReportStake)
func (_Augur *AugurFilterer) WatchDesignatedReportStakeChanged(opts *bind.WatchOpts, sink chan<- *AugurDesignatedReportStakeChanged, universe []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "DesignatedReportStakeChanged", universeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurDesignatedReportStakeChanged)
				if err := _Augur.contract.UnpackLog(event, "DesignatedReportStakeChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDesignatedReportStakeChanged is a log parse operation binding the contract event 0x9c75a088fcb0527d67a80a7d0a5006bbabe02f4b23984234ae68b2b146f001bc.
//
// Solidity: event DesignatedReportStakeChanged(address indexed universe, uint256 designatedReportStake)
func (_Augur *AugurFilterer) ParseDesignatedReportStakeChanged(log types.Log) (*AugurDesignatedReportStakeChanged, error) {
	event := new(AugurDesignatedReportStakeChanged)
	if err := _Augur.contract.UnpackLog(event, "DesignatedReportStakeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurDisputeCrowdsourcerCompletedIterator is returned from FilterDisputeCrowdsourcerCompleted and is used to iterate over the raw logs and unpacked data for DisputeCrowdsourcerCompleted events raised by the Augur contract.
type AugurDisputeCrowdsourcerCompletedIterator struct {
	Event *AugurDisputeCrowdsourcerCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurDisputeCrowdsourcerCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurDisputeCrowdsourcerCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurDisputeCrowdsourcerCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurDisputeCrowdsourcerCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurDisputeCrowdsourcerCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurDisputeCrowdsourcerCompleted represents a DisputeCrowdsourcerCompleted event raised by the Augur contract.
type AugurDisputeCrowdsourcerCompleted struct {
	Universe               common.Address
	Market                 common.Address
	DisputeCrowdsourcer    common.Address
	PayoutNumerators       []*big.Int
	NextWindowStartTime    *big.Int
	NextWindowEndTime      *big.Int
	PacingOn               bool
	TotalRepStakedInPayout *big.Int
	TotalRepStakedInMarket *big.Int
	DisputeRound           *big.Int
	Timestamp              *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterDisputeCrowdsourcerCompleted is a free log retrieval operation binding the contract event 0x81afc41f9f2f0d22a52a2ddb3a0b6db83baf39c05544fd25f2751b72b1943bb5.
//
// Solidity: event DisputeCrowdsourcerCompleted(address indexed universe, address indexed market, address disputeCrowdsourcer, uint256[] payoutNumerators, uint256 nextWindowStartTime, uint256 nextWindowEndTime, bool pacingOn, uint256 totalRepStakedInPayout, uint256 totalRepStakedInMarket, uint256 disputeRound, uint256 timestamp)
func (_Augur *AugurFilterer) FilterDisputeCrowdsourcerCompleted(opts *bind.FilterOpts, universe []common.Address, market []common.Address) (*AugurDisputeCrowdsourcerCompletedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "DisputeCrowdsourcerCompleted", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &AugurDisputeCrowdsourcerCompletedIterator{contract: _Augur.contract, event: "DisputeCrowdsourcerCompleted", logs: logs, sub: sub}, nil
}

// WatchDisputeCrowdsourcerCompleted is a free log subscription operation binding the contract event 0x81afc41f9f2f0d22a52a2ddb3a0b6db83baf39c05544fd25f2751b72b1943bb5.
//
// Solidity: event DisputeCrowdsourcerCompleted(address indexed universe, address indexed market, address disputeCrowdsourcer, uint256[] payoutNumerators, uint256 nextWindowStartTime, uint256 nextWindowEndTime, bool pacingOn, uint256 totalRepStakedInPayout, uint256 totalRepStakedInMarket, uint256 disputeRound, uint256 timestamp)
func (_Augur *AugurFilterer) WatchDisputeCrowdsourcerCompleted(opts *bind.WatchOpts, sink chan<- *AugurDisputeCrowdsourcerCompleted, universe []common.Address, market []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "DisputeCrowdsourcerCompleted", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurDisputeCrowdsourcerCompleted)
				if err := _Augur.contract.UnpackLog(event, "DisputeCrowdsourcerCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDisputeCrowdsourcerCompleted is a log parse operation binding the contract event 0x81afc41f9f2f0d22a52a2ddb3a0b6db83baf39c05544fd25f2751b72b1943bb5.
//
// Solidity: event DisputeCrowdsourcerCompleted(address indexed universe, address indexed market, address disputeCrowdsourcer, uint256[] payoutNumerators, uint256 nextWindowStartTime, uint256 nextWindowEndTime, bool pacingOn, uint256 totalRepStakedInPayout, uint256 totalRepStakedInMarket, uint256 disputeRound, uint256 timestamp)
func (_Augur *AugurFilterer) ParseDisputeCrowdsourcerCompleted(log types.Log) (*AugurDisputeCrowdsourcerCompleted, error) {
	event := new(AugurDisputeCrowdsourcerCompleted)
	if err := _Augur.contract.UnpackLog(event, "DisputeCrowdsourcerCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurDisputeCrowdsourcerContributionIterator is returned from FilterDisputeCrowdsourcerContribution and is used to iterate over the raw logs and unpacked data for DisputeCrowdsourcerContribution events raised by the Augur contract.
type AugurDisputeCrowdsourcerContributionIterator struct {
	Event *AugurDisputeCrowdsourcerContribution // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurDisputeCrowdsourcerContributionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurDisputeCrowdsourcerContribution)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurDisputeCrowdsourcerContribution)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurDisputeCrowdsourcerContributionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurDisputeCrowdsourcerContributionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurDisputeCrowdsourcerContribution represents a DisputeCrowdsourcerContribution event raised by the Augur contract.
type AugurDisputeCrowdsourcerContribution struct {
	Universe            common.Address
	Reporter            common.Address
	Market              common.Address
	DisputeCrowdsourcer common.Address
	AmountStaked        *big.Int
	Description         string
	PayoutNumerators    []*big.Int
	CurrentStake        *big.Int
	StakeRemaining      *big.Int
	DisputeRound        *big.Int
	Timestamp           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDisputeCrowdsourcerContribution is a free log retrieval operation binding the contract event 0xe7f47639cdf56ec6c5451df334b73c9ca5cccd20da2c0f4e390e9bb71a6f672a.
//
// Solidity: event DisputeCrowdsourcerContribution(address indexed universe, address indexed reporter, address indexed market, address disputeCrowdsourcer, uint256 amountStaked, string description, uint256[] payoutNumerators, uint256 currentStake, uint256 stakeRemaining, uint256 disputeRound, uint256 timestamp)
func (_Augur *AugurFilterer) FilterDisputeCrowdsourcerContribution(opts *bind.FilterOpts, universe []common.Address, reporter []common.Address, market []common.Address) (*AugurDisputeCrowdsourcerContributionIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "DisputeCrowdsourcerContribution", universeRule, reporterRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &AugurDisputeCrowdsourcerContributionIterator{contract: _Augur.contract, event: "DisputeCrowdsourcerContribution", logs: logs, sub: sub}, nil
}

// WatchDisputeCrowdsourcerContribution is a free log subscription operation binding the contract event 0xe7f47639cdf56ec6c5451df334b73c9ca5cccd20da2c0f4e390e9bb71a6f672a.
//
// Solidity: event DisputeCrowdsourcerContribution(address indexed universe, address indexed reporter, address indexed market, address disputeCrowdsourcer, uint256 amountStaked, string description, uint256[] payoutNumerators, uint256 currentStake, uint256 stakeRemaining, uint256 disputeRound, uint256 timestamp)
func (_Augur *AugurFilterer) WatchDisputeCrowdsourcerContribution(opts *bind.WatchOpts, sink chan<- *AugurDisputeCrowdsourcerContribution, universe []common.Address, reporter []common.Address, market []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "DisputeCrowdsourcerContribution", universeRule, reporterRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurDisputeCrowdsourcerContribution)
				if err := _Augur.contract.UnpackLog(event, "DisputeCrowdsourcerContribution", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDisputeCrowdsourcerContribution is a log parse operation binding the contract event 0xe7f47639cdf56ec6c5451df334b73c9ca5cccd20da2c0f4e390e9bb71a6f672a.
//
// Solidity: event DisputeCrowdsourcerContribution(address indexed universe, address indexed reporter, address indexed market, address disputeCrowdsourcer, uint256 amountStaked, string description, uint256[] payoutNumerators, uint256 currentStake, uint256 stakeRemaining, uint256 disputeRound, uint256 timestamp)
func (_Augur *AugurFilterer) ParseDisputeCrowdsourcerContribution(log types.Log) (*AugurDisputeCrowdsourcerContribution, error) {
	event := new(AugurDisputeCrowdsourcerContribution)
	if err := _Augur.contract.UnpackLog(event, "DisputeCrowdsourcerContribution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurDisputeCrowdsourcerCreatedIterator is returned from FilterDisputeCrowdsourcerCreated and is used to iterate over the raw logs and unpacked data for DisputeCrowdsourcerCreated events raised by the Augur contract.
type AugurDisputeCrowdsourcerCreatedIterator struct {
	Event *AugurDisputeCrowdsourcerCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurDisputeCrowdsourcerCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurDisputeCrowdsourcerCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurDisputeCrowdsourcerCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurDisputeCrowdsourcerCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurDisputeCrowdsourcerCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurDisputeCrowdsourcerCreated represents a DisputeCrowdsourcerCreated event raised by the Augur contract.
type AugurDisputeCrowdsourcerCreated struct {
	Universe            common.Address
	Market              common.Address
	DisputeCrowdsourcer common.Address
	PayoutNumerators    []*big.Int
	Size                *big.Int
	DisputeRound        *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDisputeCrowdsourcerCreated is a free log retrieval operation binding the contract event 0xf9a0b30bcf861874bf36630742f0d56b22648898d7cdd0cd785d74acd17e0d44.
//
// Solidity: event DisputeCrowdsourcerCreated(address indexed universe, address indexed market, address disputeCrowdsourcer, uint256[] payoutNumerators, uint256 size, uint256 disputeRound)
func (_Augur *AugurFilterer) FilterDisputeCrowdsourcerCreated(opts *bind.FilterOpts, universe []common.Address, market []common.Address) (*AugurDisputeCrowdsourcerCreatedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "DisputeCrowdsourcerCreated", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &AugurDisputeCrowdsourcerCreatedIterator{contract: _Augur.contract, event: "DisputeCrowdsourcerCreated", logs: logs, sub: sub}, nil
}

// WatchDisputeCrowdsourcerCreated is a free log subscription operation binding the contract event 0xf9a0b30bcf861874bf36630742f0d56b22648898d7cdd0cd785d74acd17e0d44.
//
// Solidity: event DisputeCrowdsourcerCreated(address indexed universe, address indexed market, address disputeCrowdsourcer, uint256[] payoutNumerators, uint256 size, uint256 disputeRound)
func (_Augur *AugurFilterer) WatchDisputeCrowdsourcerCreated(opts *bind.WatchOpts, sink chan<- *AugurDisputeCrowdsourcerCreated, universe []common.Address, market []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "DisputeCrowdsourcerCreated", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurDisputeCrowdsourcerCreated)
				if err := _Augur.contract.UnpackLog(event, "DisputeCrowdsourcerCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDisputeCrowdsourcerCreated is a log parse operation binding the contract event 0xf9a0b30bcf861874bf36630742f0d56b22648898d7cdd0cd785d74acd17e0d44.
//
// Solidity: event DisputeCrowdsourcerCreated(address indexed universe, address indexed market, address disputeCrowdsourcer, uint256[] payoutNumerators, uint256 size, uint256 disputeRound)
func (_Augur *AugurFilterer) ParseDisputeCrowdsourcerCreated(log types.Log) (*AugurDisputeCrowdsourcerCreated, error) {
	event := new(AugurDisputeCrowdsourcerCreated)
	if err := _Augur.contract.UnpackLog(event, "DisputeCrowdsourcerCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurDisputeCrowdsourcerRedeemedIterator is returned from FilterDisputeCrowdsourcerRedeemed and is used to iterate over the raw logs and unpacked data for DisputeCrowdsourcerRedeemed events raised by the Augur contract.
type AugurDisputeCrowdsourcerRedeemedIterator struct {
	Event *AugurDisputeCrowdsourcerRedeemed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurDisputeCrowdsourcerRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurDisputeCrowdsourcerRedeemed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurDisputeCrowdsourcerRedeemed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurDisputeCrowdsourcerRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurDisputeCrowdsourcerRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurDisputeCrowdsourcerRedeemed represents a DisputeCrowdsourcerRedeemed event raised by the Augur contract.
type AugurDisputeCrowdsourcerRedeemed struct {
	Universe            common.Address
	Reporter            common.Address
	Market              common.Address
	DisputeCrowdsourcer common.Address
	AmountRedeemed      *big.Int
	RepReceived         *big.Int
	PayoutNumerators    []*big.Int
	Timestamp           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDisputeCrowdsourcerRedeemed is a free log retrieval operation binding the contract event 0x6afb0328cf957750be87a6f34b1cd21457ddf1382af65f9592ff2d333945633f.
//
// Solidity: event DisputeCrowdsourcerRedeemed(address indexed universe, address indexed reporter, address indexed market, address disputeCrowdsourcer, uint256 amountRedeemed, uint256 repReceived, uint256[] payoutNumerators, uint256 timestamp)
func (_Augur *AugurFilterer) FilterDisputeCrowdsourcerRedeemed(opts *bind.FilterOpts, universe []common.Address, reporter []common.Address, market []common.Address) (*AugurDisputeCrowdsourcerRedeemedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "DisputeCrowdsourcerRedeemed", universeRule, reporterRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &AugurDisputeCrowdsourcerRedeemedIterator{contract: _Augur.contract, event: "DisputeCrowdsourcerRedeemed", logs: logs, sub: sub}, nil
}

// WatchDisputeCrowdsourcerRedeemed is a free log subscription operation binding the contract event 0x6afb0328cf957750be87a6f34b1cd21457ddf1382af65f9592ff2d333945633f.
//
// Solidity: event DisputeCrowdsourcerRedeemed(address indexed universe, address indexed reporter, address indexed market, address disputeCrowdsourcer, uint256 amountRedeemed, uint256 repReceived, uint256[] payoutNumerators, uint256 timestamp)
func (_Augur *AugurFilterer) WatchDisputeCrowdsourcerRedeemed(opts *bind.WatchOpts, sink chan<- *AugurDisputeCrowdsourcerRedeemed, universe []common.Address, reporter []common.Address, market []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "DisputeCrowdsourcerRedeemed", universeRule, reporterRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurDisputeCrowdsourcerRedeemed)
				if err := _Augur.contract.UnpackLog(event, "DisputeCrowdsourcerRedeemed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDisputeCrowdsourcerRedeemed is a log parse operation binding the contract event 0x6afb0328cf957750be87a6f34b1cd21457ddf1382af65f9592ff2d333945633f.
//
// Solidity: event DisputeCrowdsourcerRedeemed(address indexed universe, address indexed reporter, address indexed market, address disputeCrowdsourcer, uint256 amountRedeemed, uint256 repReceived, uint256[] payoutNumerators, uint256 timestamp)
func (_Augur *AugurFilterer) ParseDisputeCrowdsourcerRedeemed(log types.Log) (*AugurDisputeCrowdsourcerRedeemed, error) {
	event := new(AugurDisputeCrowdsourcerRedeemed)
	if err := _Augur.contract.UnpackLog(event, "DisputeCrowdsourcerRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurDisputeWindowCreatedIterator is returned from FilterDisputeWindowCreated and is used to iterate over the raw logs and unpacked data for DisputeWindowCreated events raised by the Augur contract.
type AugurDisputeWindowCreatedIterator struct {
	Event *AugurDisputeWindowCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurDisputeWindowCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurDisputeWindowCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurDisputeWindowCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurDisputeWindowCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurDisputeWindowCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurDisputeWindowCreated represents a DisputeWindowCreated event raised by the Augur contract.
type AugurDisputeWindowCreated struct {
	Universe      common.Address
	DisputeWindow common.Address
	StartTime     *big.Int
	EndTime       *big.Int
	Id            *big.Int
	Initial       bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDisputeWindowCreated is a free log retrieval operation binding the contract event 0x97f8b399e255f30d56b759b645c86652624ee258937579ff4a747abaeae857c4.
//
// Solidity: event DisputeWindowCreated(address indexed universe, address disputeWindow, uint256 startTime, uint256 endTime, uint256 id, bool initial)
func (_Augur *AugurFilterer) FilterDisputeWindowCreated(opts *bind.FilterOpts, universe []common.Address) (*AugurDisputeWindowCreatedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "DisputeWindowCreated", universeRule)
	if err != nil {
		return nil, err
	}
	return &AugurDisputeWindowCreatedIterator{contract: _Augur.contract, event: "DisputeWindowCreated", logs: logs, sub: sub}, nil
}

// WatchDisputeWindowCreated is a free log subscription operation binding the contract event 0x97f8b399e255f30d56b759b645c86652624ee258937579ff4a747abaeae857c4.
//
// Solidity: event DisputeWindowCreated(address indexed universe, address disputeWindow, uint256 startTime, uint256 endTime, uint256 id, bool initial)
func (_Augur *AugurFilterer) WatchDisputeWindowCreated(opts *bind.WatchOpts, sink chan<- *AugurDisputeWindowCreated, universe []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "DisputeWindowCreated", universeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurDisputeWindowCreated)
				if err := _Augur.contract.UnpackLog(event, "DisputeWindowCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDisputeWindowCreated is a log parse operation binding the contract event 0x97f8b399e255f30d56b759b645c86652624ee258937579ff4a747abaeae857c4.
//
// Solidity: event DisputeWindowCreated(address indexed universe, address disputeWindow, uint256 startTime, uint256 endTime, uint256 id, bool initial)
func (_Augur *AugurFilterer) ParseDisputeWindowCreated(log types.Log) (*AugurDisputeWindowCreated, error) {
	event := new(AugurDisputeWindowCreated)
	if err := _Augur.contract.UnpackLog(event, "DisputeWindowCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurFinishDeploymentIterator is returned from FilterFinishDeployment and is used to iterate over the raw logs and unpacked data for FinishDeployment events raised by the Augur contract.
type AugurFinishDeploymentIterator struct {
	Event *AugurFinishDeployment // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurFinishDeploymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurFinishDeployment)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurFinishDeployment)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurFinishDeploymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurFinishDeploymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurFinishDeployment represents a FinishDeployment event raised by the Augur contract.
type AugurFinishDeployment struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFinishDeployment is a free log retrieval operation binding the contract event 0xf06c142f93fdd00fbcd1e8f3d82e6f22667d52df764b39570061a7dbeea09be0.
//
// Solidity: event FinishDeployment()
func (_Augur *AugurFilterer) FilterFinishDeployment(opts *bind.FilterOpts) (*AugurFinishDeploymentIterator, error) {

	logs, sub, err := _Augur.contract.FilterLogs(opts, "FinishDeployment")
	if err != nil {
		return nil, err
	}
	return &AugurFinishDeploymentIterator{contract: _Augur.contract, event: "FinishDeployment", logs: logs, sub: sub}, nil
}

// WatchFinishDeployment is a free log subscription operation binding the contract event 0xf06c142f93fdd00fbcd1e8f3d82e6f22667d52df764b39570061a7dbeea09be0.
//
// Solidity: event FinishDeployment()
func (_Augur *AugurFilterer) WatchFinishDeployment(opts *bind.WatchOpts, sink chan<- *AugurFinishDeployment) (event.Subscription, error) {

	logs, sub, err := _Augur.contract.WatchLogs(opts, "FinishDeployment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurFinishDeployment)
				if err := _Augur.contract.UnpackLog(event, "FinishDeployment", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFinishDeployment is a log parse operation binding the contract event 0xf06c142f93fdd00fbcd1e8f3d82e6f22667d52df764b39570061a7dbeea09be0.
//
// Solidity: event FinishDeployment()
func (_Augur *AugurFilterer) ParseFinishDeployment(log types.Log) (*AugurFinishDeployment, error) {
	event := new(AugurFinishDeployment)
	if err := _Augur.contract.UnpackLog(event, "FinishDeployment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurInitialReportSubmittedIterator is returned from FilterInitialReportSubmitted and is used to iterate over the raw logs and unpacked data for InitialReportSubmitted events raised by the Augur contract.
type AugurInitialReportSubmittedIterator struct {
	Event *AugurInitialReportSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurInitialReportSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurInitialReportSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurInitialReportSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurInitialReportSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurInitialReportSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurInitialReportSubmitted represents a InitialReportSubmitted event raised by the Augur contract.
type AugurInitialReportSubmitted struct {
	Universe             common.Address
	Reporter             common.Address
	Market               common.Address
	InitialReporter      common.Address
	AmountStaked         *big.Int
	IsDesignatedReporter bool
	PayoutNumerators     []*big.Int
	Description          string
	NextWindowStartTime  *big.Int
	NextWindowEndTime    *big.Int
	Timestamp            *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterInitialReportSubmitted is a free log retrieval operation binding the contract event 0xc3ebb227c22e7644e9bef8822009f746a72c86f239760124d67fdc2c302b3115.
//
// Solidity: event InitialReportSubmitted(address indexed universe, address indexed reporter, address indexed market, address initialReporter, uint256 amountStaked, bool isDesignatedReporter, uint256[] payoutNumerators, string description, uint256 nextWindowStartTime, uint256 nextWindowEndTime, uint256 timestamp)
func (_Augur *AugurFilterer) FilterInitialReportSubmitted(opts *bind.FilterOpts, universe []common.Address, reporter []common.Address, market []common.Address) (*AugurInitialReportSubmittedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "InitialReportSubmitted", universeRule, reporterRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &AugurInitialReportSubmittedIterator{contract: _Augur.contract, event: "InitialReportSubmitted", logs: logs, sub: sub}, nil
}

// WatchInitialReportSubmitted is a free log subscription operation binding the contract event 0xc3ebb227c22e7644e9bef8822009f746a72c86f239760124d67fdc2c302b3115.
//
// Solidity: event InitialReportSubmitted(address indexed universe, address indexed reporter, address indexed market, address initialReporter, uint256 amountStaked, bool isDesignatedReporter, uint256[] payoutNumerators, string description, uint256 nextWindowStartTime, uint256 nextWindowEndTime, uint256 timestamp)
func (_Augur *AugurFilterer) WatchInitialReportSubmitted(opts *bind.WatchOpts, sink chan<- *AugurInitialReportSubmitted, universe []common.Address, reporter []common.Address, market []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "InitialReportSubmitted", universeRule, reporterRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurInitialReportSubmitted)
				if err := _Augur.contract.UnpackLog(event, "InitialReportSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialReportSubmitted is a log parse operation binding the contract event 0xc3ebb227c22e7644e9bef8822009f746a72c86f239760124d67fdc2c302b3115.
//
// Solidity: event InitialReportSubmitted(address indexed universe, address indexed reporter, address indexed market, address initialReporter, uint256 amountStaked, bool isDesignatedReporter, uint256[] payoutNumerators, string description, uint256 nextWindowStartTime, uint256 nextWindowEndTime, uint256 timestamp)
func (_Augur *AugurFilterer) ParseInitialReportSubmitted(log types.Log) (*AugurInitialReportSubmitted, error) {
	event := new(AugurInitialReportSubmitted)
	if err := _Augur.contract.UnpackLog(event, "InitialReportSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurInitialReporterRedeemedIterator is returned from FilterInitialReporterRedeemed and is used to iterate over the raw logs and unpacked data for InitialReporterRedeemed events raised by the Augur contract.
type AugurInitialReporterRedeemedIterator struct {
	Event *AugurInitialReporterRedeemed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurInitialReporterRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurInitialReporterRedeemed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurInitialReporterRedeemed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurInitialReporterRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurInitialReporterRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurInitialReporterRedeemed represents a InitialReporterRedeemed event raised by the Augur contract.
type AugurInitialReporterRedeemed struct {
	Universe         common.Address
	Reporter         common.Address
	Market           common.Address
	InitialReporter  common.Address
	AmountRedeemed   *big.Int
	RepReceived      *big.Int
	PayoutNumerators []*big.Int
	Timestamp        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterInitialReporterRedeemed is a free log retrieval operation binding the contract event 0x3ffffb51f92f91faf4ba8c906f5a0180d1033be93b1e227cd92c872dc234fdf0.
//
// Solidity: event InitialReporterRedeemed(address indexed universe, address indexed reporter, address indexed market, address initialReporter, uint256 amountRedeemed, uint256 repReceived, uint256[] payoutNumerators, uint256 timestamp)
func (_Augur *AugurFilterer) FilterInitialReporterRedeemed(opts *bind.FilterOpts, universe []common.Address, reporter []common.Address, market []common.Address) (*AugurInitialReporterRedeemedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "InitialReporterRedeemed", universeRule, reporterRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &AugurInitialReporterRedeemedIterator{contract: _Augur.contract, event: "InitialReporterRedeemed", logs: logs, sub: sub}, nil
}

// WatchInitialReporterRedeemed is a free log subscription operation binding the contract event 0x3ffffb51f92f91faf4ba8c906f5a0180d1033be93b1e227cd92c872dc234fdf0.
//
// Solidity: event InitialReporterRedeemed(address indexed universe, address indexed reporter, address indexed market, address initialReporter, uint256 amountRedeemed, uint256 repReceived, uint256[] payoutNumerators, uint256 timestamp)
func (_Augur *AugurFilterer) WatchInitialReporterRedeemed(opts *bind.WatchOpts, sink chan<- *AugurInitialReporterRedeemed, universe []common.Address, reporter []common.Address, market []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "InitialReporterRedeemed", universeRule, reporterRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurInitialReporterRedeemed)
				if err := _Augur.contract.UnpackLog(event, "InitialReporterRedeemed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialReporterRedeemed is a log parse operation binding the contract event 0x3ffffb51f92f91faf4ba8c906f5a0180d1033be93b1e227cd92c872dc234fdf0.
//
// Solidity: event InitialReporterRedeemed(address indexed universe, address indexed reporter, address indexed market, address initialReporter, uint256 amountRedeemed, uint256 repReceived, uint256[] payoutNumerators, uint256 timestamp)
func (_Augur *AugurFilterer) ParseInitialReporterRedeemed(log types.Log) (*AugurInitialReporterRedeemed, error) {
	event := new(AugurInitialReporterRedeemed)
	if err := _Augur.contract.UnpackLog(event, "InitialReporterRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurInitialReporterTransferredIterator is returned from FilterInitialReporterTransferred and is used to iterate over the raw logs and unpacked data for InitialReporterTransferred events raised by the Augur contract.
type AugurInitialReporterTransferredIterator struct {
	Event *AugurInitialReporterTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurInitialReporterTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurInitialReporterTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurInitialReporterTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurInitialReporterTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurInitialReporterTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurInitialReporterTransferred represents a InitialReporterTransferred event raised by the Augur contract.
type AugurInitialReporterTransferred struct {
	Universe common.Address
	Market   common.Address
	From     common.Address
	To       common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInitialReporterTransferred is a free log retrieval operation binding the contract event 0xee62c58e2603b92f96a002e012f4f3bd5748102cfa3b711f6d778c6237fcaa96.
//
// Solidity: event InitialReporterTransferred(address indexed universe, address indexed market, address from, address to)
func (_Augur *AugurFilterer) FilterInitialReporterTransferred(opts *bind.FilterOpts, universe []common.Address, market []common.Address) (*AugurInitialReporterTransferredIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "InitialReporterTransferred", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &AugurInitialReporterTransferredIterator{contract: _Augur.contract, event: "InitialReporterTransferred", logs: logs, sub: sub}, nil
}

// WatchInitialReporterTransferred is a free log subscription operation binding the contract event 0xee62c58e2603b92f96a002e012f4f3bd5748102cfa3b711f6d778c6237fcaa96.
//
// Solidity: event InitialReporterTransferred(address indexed universe, address indexed market, address from, address to)
func (_Augur *AugurFilterer) WatchInitialReporterTransferred(opts *bind.WatchOpts, sink chan<- *AugurInitialReporterTransferred, universe []common.Address, market []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "InitialReporterTransferred", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurInitialReporterTransferred)
				if err := _Augur.contract.UnpackLog(event, "InitialReporterTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialReporterTransferred is a log parse operation binding the contract event 0xee62c58e2603b92f96a002e012f4f3bd5748102cfa3b711f6d778c6237fcaa96.
//
// Solidity: event InitialReporterTransferred(address indexed universe, address indexed market, address from, address to)
func (_Augur *AugurFilterer) ParseInitialReporterTransferred(log types.Log) (*AugurInitialReporterTransferred, error) {
	event := new(AugurInitialReporterTransferred)
	if err := _Augur.contract.UnpackLog(event, "InitialReporterTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurMarketCreatedIterator is returned from FilterMarketCreated and is used to iterate over the raw logs and unpacked data for MarketCreated events raised by the Augur contract.
type AugurMarketCreatedIterator struct {
	Event *AugurMarketCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurMarketCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurMarketCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurMarketCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurMarketCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurMarketCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurMarketCreated represents a MarketCreated event raised by the Augur contract.
type AugurMarketCreated struct {
	Universe             common.Address
	EndTime              *big.Int
	ExtraInfo            string
	Market               common.Address
	MarketCreator        common.Address
	DesignatedReporter   common.Address
	FeePerCashInAttoCash *big.Int
	Prices               []*big.Int
	MarketType           uint8
	NumTicks             *big.Int
	Outcomes             [][32]byte
	NoShowBond           *big.Int
	Timestamp            *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterMarketCreated is a free log retrieval operation binding the contract event 0xea17ae24b0d40ea7962a6d832db46d1f81eaec1562946d0830d1c21d4c000ec1.
//
// Solidity: event MarketCreated(address indexed universe, uint256 endTime, string extraInfo, address market, address indexed marketCreator, address designatedReporter, uint256 feePerCashInAttoCash, int256[] prices, uint8 marketType, uint256 numTicks, bytes32[] outcomes, uint256 noShowBond, uint256 timestamp)
func (_Augur *AugurFilterer) FilterMarketCreated(opts *bind.FilterOpts, universe []common.Address, marketCreator []common.Address) (*AugurMarketCreatedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}

	var marketCreatorRule []interface{}
	for _, marketCreatorItem := range marketCreator {
		marketCreatorRule = append(marketCreatorRule, marketCreatorItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "MarketCreated", universeRule, marketCreatorRule)
	if err != nil {
		return nil, err
	}
	return &AugurMarketCreatedIterator{contract: _Augur.contract, event: "MarketCreated", logs: logs, sub: sub}, nil
}

// WatchMarketCreated is a free log subscription operation binding the contract event 0xea17ae24b0d40ea7962a6d832db46d1f81eaec1562946d0830d1c21d4c000ec1.
//
// Solidity: event MarketCreated(address indexed universe, uint256 endTime, string extraInfo, address market, address indexed marketCreator, address designatedReporter, uint256 feePerCashInAttoCash, int256[] prices, uint8 marketType, uint256 numTicks, bytes32[] outcomes, uint256 noShowBond, uint256 timestamp)
func (_Augur *AugurFilterer) WatchMarketCreated(opts *bind.WatchOpts, sink chan<- *AugurMarketCreated, universe []common.Address, marketCreator []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}

	var marketCreatorRule []interface{}
	for _, marketCreatorItem := range marketCreator {
		marketCreatorRule = append(marketCreatorRule, marketCreatorItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "MarketCreated", universeRule, marketCreatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurMarketCreated)
				if err := _Augur.contract.UnpackLog(event, "MarketCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketCreated is a log parse operation binding the contract event 0xea17ae24b0d40ea7962a6d832db46d1f81eaec1562946d0830d1c21d4c000ec1.
//
// Solidity: event MarketCreated(address indexed universe, uint256 endTime, string extraInfo, address market, address indexed marketCreator, address designatedReporter, uint256 feePerCashInAttoCash, int256[] prices, uint8 marketType, uint256 numTicks, bytes32[] outcomes, uint256 noShowBond, uint256 timestamp)
func (_Augur *AugurFilterer) ParseMarketCreated(log types.Log) (*AugurMarketCreated, error) {
	event := new(AugurMarketCreated)
	if err := _Augur.contract.UnpackLog(event, "MarketCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurMarketFinalizedIterator is returned from FilterMarketFinalized and is used to iterate over the raw logs and unpacked data for MarketFinalized events raised by the Augur contract.
type AugurMarketFinalizedIterator struct {
	Event *AugurMarketFinalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurMarketFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurMarketFinalized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurMarketFinalized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurMarketFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurMarketFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurMarketFinalized represents a MarketFinalized event raised by the Augur contract.
type AugurMarketFinalized struct {
	Universe                common.Address
	Market                  common.Address
	Timestamp               *big.Int
	WinningPayoutNumerators []*big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterMarketFinalized is a free log retrieval operation binding the contract event 0x6d39632c2dc10305bf5771cfff4af1851f07c03ea27b821cad382466bdf7a21f.
//
// Solidity: event MarketFinalized(address indexed universe, address indexed market, uint256 timestamp, uint256[] winningPayoutNumerators)
func (_Augur *AugurFilterer) FilterMarketFinalized(opts *bind.FilterOpts, universe []common.Address, market []common.Address) (*AugurMarketFinalizedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "MarketFinalized", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &AugurMarketFinalizedIterator{contract: _Augur.contract, event: "MarketFinalized", logs: logs, sub: sub}, nil
}

// WatchMarketFinalized is a free log subscription operation binding the contract event 0x6d39632c2dc10305bf5771cfff4af1851f07c03ea27b821cad382466bdf7a21f.
//
// Solidity: event MarketFinalized(address indexed universe, address indexed market, uint256 timestamp, uint256[] winningPayoutNumerators)
func (_Augur *AugurFilterer) WatchMarketFinalized(opts *bind.WatchOpts, sink chan<- *AugurMarketFinalized, universe []common.Address, market []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "MarketFinalized", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurMarketFinalized)
				if err := _Augur.contract.UnpackLog(event, "MarketFinalized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketFinalized is a log parse operation binding the contract event 0x6d39632c2dc10305bf5771cfff4af1851f07c03ea27b821cad382466bdf7a21f.
//
// Solidity: event MarketFinalized(address indexed universe, address indexed market, uint256 timestamp, uint256[] winningPayoutNumerators)
func (_Augur *AugurFilterer) ParseMarketFinalized(log types.Log) (*AugurMarketFinalized, error) {
	event := new(AugurMarketFinalized)
	if err := _Augur.contract.UnpackLog(event, "MarketFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurMarketMigratedIterator is returned from FilterMarketMigrated and is used to iterate over the raw logs and unpacked data for MarketMigrated events raised by the Augur contract.
type AugurMarketMigratedIterator struct {
	Event *AugurMarketMigrated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurMarketMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurMarketMigrated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurMarketMigrated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurMarketMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurMarketMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurMarketMigrated represents a MarketMigrated event raised by the Augur contract.
type AugurMarketMigrated struct {
	Market           common.Address
	OriginalUniverse common.Address
	NewUniverse      common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMarketMigrated is a free log retrieval operation binding the contract event 0xc3cf07f8fa0fafc25a9dd0bad2cd6b961c55dad41b42c8ef8f931bc40e41e08c.
//
// Solidity: event MarketMigrated(address indexed market, address indexed originalUniverse, address indexed newUniverse)
func (_Augur *AugurFilterer) FilterMarketMigrated(opts *bind.FilterOpts, market []common.Address, originalUniverse []common.Address, newUniverse []common.Address) (*AugurMarketMigratedIterator, error) {

	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var originalUniverseRule []interface{}
	for _, originalUniverseItem := range originalUniverse {
		originalUniverseRule = append(originalUniverseRule, originalUniverseItem)
	}
	var newUniverseRule []interface{}
	for _, newUniverseItem := range newUniverse {
		newUniverseRule = append(newUniverseRule, newUniverseItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "MarketMigrated", marketRule, originalUniverseRule, newUniverseRule)
	if err != nil {
		return nil, err
	}
	return &AugurMarketMigratedIterator{contract: _Augur.contract, event: "MarketMigrated", logs: logs, sub: sub}, nil
}

// WatchMarketMigrated is a free log subscription operation binding the contract event 0xc3cf07f8fa0fafc25a9dd0bad2cd6b961c55dad41b42c8ef8f931bc40e41e08c.
//
// Solidity: event MarketMigrated(address indexed market, address indexed originalUniverse, address indexed newUniverse)
func (_Augur *AugurFilterer) WatchMarketMigrated(opts *bind.WatchOpts, sink chan<- *AugurMarketMigrated, market []common.Address, originalUniverse []common.Address, newUniverse []common.Address) (event.Subscription, error) {

	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var originalUniverseRule []interface{}
	for _, originalUniverseItem := range originalUniverse {
		originalUniverseRule = append(originalUniverseRule, originalUniverseItem)
	}
	var newUniverseRule []interface{}
	for _, newUniverseItem := range newUniverse {
		newUniverseRule = append(newUniverseRule, newUniverseItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "MarketMigrated", marketRule, originalUniverseRule, newUniverseRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurMarketMigrated)
				if err := _Augur.contract.UnpackLog(event, "MarketMigrated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketMigrated is a log parse operation binding the contract event 0xc3cf07f8fa0fafc25a9dd0bad2cd6b961c55dad41b42c8ef8f931bc40e41e08c.
//
// Solidity: event MarketMigrated(address indexed market, address indexed originalUniverse, address indexed newUniverse)
func (_Augur *AugurFilterer) ParseMarketMigrated(log types.Log) (*AugurMarketMigrated, error) {
	event := new(AugurMarketMigrated)
	if err := _Augur.contract.UnpackLog(event, "MarketMigrated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurMarketOIChangedIterator is returned from FilterMarketOIChanged and is used to iterate over the raw logs and unpacked data for MarketOIChanged events raised by the Augur contract.
type AugurMarketOIChangedIterator struct {
	Event *AugurMarketOIChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurMarketOIChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurMarketOIChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurMarketOIChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurMarketOIChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurMarketOIChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurMarketOIChanged represents a MarketOIChanged event raised by the Augur contract.
type AugurMarketOIChanged struct {
	Universe common.Address
	Market   common.Address
	MarketOI *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketOIChanged is a free log retrieval operation binding the contract event 0x213a05b9ad8567c2f8fa868e7375e5bf30e69add0dbb5913ca8a3e58c815c268.
//
// Solidity: event MarketOIChanged(address indexed universe, address indexed market, uint256 marketOI)
func (_Augur *AugurFilterer) FilterMarketOIChanged(opts *bind.FilterOpts, universe []common.Address, market []common.Address) (*AugurMarketOIChangedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "MarketOIChanged", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &AugurMarketOIChangedIterator{contract: _Augur.contract, event: "MarketOIChanged", logs: logs, sub: sub}, nil
}

// WatchMarketOIChanged is a free log subscription operation binding the contract event 0x213a05b9ad8567c2f8fa868e7375e5bf30e69add0dbb5913ca8a3e58c815c268.
//
// Solidity: event MarketOIChanged(address indexed universe, address indexed market, uint256 marketOI)
func (_Augur *AugurFilterer) WatchMarketOIChanged(opts *bind.WatchOpts, sink chan<- *AugurMarketOIChanged, universe []common.Address, market []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "MarketOIChanged", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurMarketOIChanged)
				if err := _Augur.contract.UnpackLog(event, "MarketOIChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketOIChanged is a log parse operation binding the contract event 0x213a05b9ad8567c2f8fa868e7375e5bf30e69add0dbb5913ca8a3e58c815c268.
//
// Solidity: event MarketOIChanged(address indexed universe, address indexed market, uint256 marketOI)
func (_Augur *AugurFilterer) ParseMarketOIChanged(log types.Log) (*AugurMarketOIChanged, error) {
	event := new(AugurMarketOIChanged)
	if err := _Augur.contract.UnpackLog(event, "MarketOIChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurMarketParticipantsDisavowedIterator is returned from FilterMarketParticipantsDisavowed and is used to iterate over the raw logs and unpacked data for MarketParticipantsDisavowed events raised by the Augur contract.
type AugurMarketParticipantsDisavowedIterator struct {
	Event *AugurMarketParticipantsDisavowed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurMarketParticipantsDisavowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurMarketParticipantsDisavowed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurMarketParticipantsDisavowed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurMarketParticipantsDisavowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurMarketParticipantsDisavowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurMarketParticipantsDisavowed represents a MarketParticipantsDisavowed event raised by the Augur contract.
type AugurMarketParticipantsDisavowed struct {
	Universe common.Address
	Market   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketParticipantsDisavowed is a free log retrieval operation binding the contract event 0x3b4f3db017516414df2695e5b0052661779d7163a6cd4368fd74313be73fa0b8.
//
// Solidity: event MarketParticipantsDisavowed(address indexed universe, address indexed market)
func (_Augur *AugurFilterer) FilterMarketParticipantsDisavowed(opts *bind.FilterOpts, universe []common.Address, market []common.Address) (*AugurMarketParticipantsDisavowedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "MarketParticipantsDisavowed", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &AugurMarketParticipantsDisavowedIterator{contract: _Augur.contract, event: "MarketParticipantsDisavowed", logs: logs, sub: sub}, nil
}

// WatchMarketParticipantsDisavowed is a free log subscription operation binding the contract event 0x3b4f3db017516414df2695e5b0052661779d7163a6cd4368fd74313be73fa0b8.
//
// Solidity: event MarketParticipantsDisavowed(address indexed universe, address indexed market)
func (_Augur *AugurFilterer) WatchMarketParticipantsDisavowed(opts *bind.WatchOpts, sink chan<- *AugurMarketParticipantsDisavowed, universe []common.Address, market []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "MarketParticipantsDisavowed", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurMarketParticipantsDisavowed)
				if err := _Augur.contract.UnpackLog(event, "MarketParticipantsDisavowed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketParticipantsDisavowed is a log parse operation binding the contract event 0x3b4f3db017516414df2695e5b0052661779d7163a6cd4368fd74313be73fa0b8.
//
// Solidity: event MarketParticipantsDisavowed(address indexed universe, address indexed market)
func (_Augur *AugurFilterer) ParseMarketParticipantsDisavowed(log types.Log) (*AugurMarketParticipantsDisavowed, error) {
	event := new(AugurMarketParticipantsDisavowed)
	if err := _Augur.contract.UnpackLog(event, "MarketParticipantsDisavowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurMarketRepBondTransferredIterator is returned from FilterMarketRepBondTransferred and is used to iterate over the raw logs and unpacked data for MarketRepBondTransferred events raised by the Augur contract.
type AugurMarketRepBondTransferredIterator struct {
	Event *AugurMarketRepBondTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurMarketRepBondTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurMarketRepBondTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurMarketRepBondTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurMarketRepBondTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurMarketRepBondTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurMarketRepBondTransferred represents a MarketRepBondTransferred event raised by the Augur contract.
type AugurMarketRepBondTransferred struct {
	Universe common.Address
	Market   common.Address
	From     common.Address
	To       common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketRepBondTransferred is a free log retrieval operation binding the contract event 0x0519ee50d0e6120223e58d0b52824ca4985c524f045a3d6a529936e511d2ba8d.
//
// Solidity: event MarketRepBondTransferred(address indexed universe, address market, address from, address to)
func (_Augur *AugurFilterer) FilterMarketRepBondTransferred(opts *bind.FilterOpts, universe []common.Address) (*AugurMarketRepBondTransferredIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "MarketRepBondTransferred", universeRule)
	if err != nil {
		return nil, err
	}
	return &AugurMarketRepBondTransferredIterator{contract: _Augur.contract, event: "MarketRepBondTransferred", logs: logs, sub: sub}, nil
}

// WatchMarketRepBondTransferred is a free log subscription operation binding the contract event 0x0519ee50d0e6120223e58d0b52824ca4985c524f045a3d6a529936e511d2ba8d.
//
// Solidity: event MarketRepBondTransferred(address indexed universe, address market, address from, address to)
func (_Augur *AugurFilterer) WatchMarketRepBondTransferred(opts *bind.WatchOpts, sink chan<- *AugurMarketRepBondTransferred, universe []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "MarketRepBondTransferred", universeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurMarketRepBondTransferred)
				if err := _Augur.contract.UnpackLog(event, "MarketRepBondTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketRepBondTransferred is a log parse operation binding the contract event 0x0519ee50d0e6120223e58d0b52824ca4985c524f045a3d6a529936e511d2ba8d.
//
// Solidity: event MarketRepBondTransferred(address indexed universe, address market, address from, address to)
func (_Augur *AugurFilterer) ParseMarketRepBondTransferred(log types.Log) (*AugurMarketRepBondTransferred, error) {
	event := new(AugurMarketRepBondTransferred)
	if err := _Augur.contract.UnpackLog(event, "MarketRepBondTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurMarketTransferredIterator is returned from FilterMarketTransferred and is used to iterate over the raw logs and unpacked data for MarketTransferred events raised by the Augur contract.
type AugurMarketTransferredIterator struct {
	Event *AugurMarketTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurMarketTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurMarketTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurMarketTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurMarketTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurMarketTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurMarketTransferred represents a MarketTransferred event raised by the Augur contract.
type AugurMarketTransferred struct {
	Universe common.Address
	Market   common.Address
	From     common.Address
	To       common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketTransferred is a free log retrieval operation binding the contract event 0x55f2a7bfa32e835c3f3c3cff653a3d11c077ce1b00c5a41c6aaf09eedc1ac3b2.
//
// Solidity: event MarketTransferred(address indexed universe, address indexed market, address from, address to)
func (_Augur *AugurFilterer) FilterMarketTransferred(opts *bind.FilterOpts, universe []common.Address, market []common.Address) (*AugurMarketTransferredIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "MarketTransferred", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &AugurMarketTransferredIterator{contract: _Augur.contract, event: "MarketTransferred", logs: logs, sub: sub}, nil
}

// WatchMarketTransferred is a free log subscription operation binding the contract event 0x55f2a7bfa32e835c3f3c3cff653a3d11c077ce1b00c5a41c6aaf09eedc1ac3b2.
//
// Solidity: event MarketTransferred(address indexed universe, address indexed market, address from, address to)
func (_Augur *AugurFilterer) WatchMarketTransferred(opts *bind.WatchOpts, sink chan<- *AugurMarketTransferred, universe []common.Address, market []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "MarketTransferred", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurMarketTransferred)
				if err := _Augur.contract.UnpackLog(event, "MarketTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketTransferred is a log parse operation binding the contract event 0x55f2a7bfa32e835c3f3c3cff653a3d11c077ce1b00c5a41c6aaf09eedc1ac3b2.
//
// Solidity: event MarketTransferred(address indexed universe, address indexed market, address from, address to)
func (_Augur *AugurFilterer) ParseMarketTransferred(log types.Log) (*AugurMarketTransferred, error) {
	event := new(AugurMarketTransferred)
	if err := _Augur.contract.UnpackLog(event, "MarketTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurNoShowBondChangedIterator is returned from FilterNoShowBondChanged and is used to iterate over the raw logs and unpacked data for NoShowBondChanged events raised by the Augur contract.
type AugurNoShowBondChangedIterator struct {
	Event *AugurNoShowBondChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurNoShowBondChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurNoShowBondChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurNoShowBondChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurNoShowBondChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurNoShowBondChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurNoShowBondChanged represents a NoShowBondChanged event raised by the Augur contract.
type AugurNoShowBondChanged struct {
	Universe   common.Address
	NoShowBond *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNoShowBondChanged is a free log retrieval operation binding the contract event 0xd1fc3f2cb1387e602db0e6f8f22649df65df5246eeff281cf6d1ef62feda4ece.
//
// Solidity: event NoShowBondChanged(address indexed universe, uint256 noShowBond)
func (_Augur *AugurFilterer) FilterNoShowBondChanged(opts *bind.FilterOpts, universe []common.Address) (*AugurNoShowBondChangedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "NoShowBondChanged", universeRule)
	if err != nil {
		return nil, err
	}
	return &AugurNoShowBondChangedIterator{contract: _Augur.contract, event: "NoShowBondChanged", logs: logs, sub: sub}, nil
}

// WatchNoShowBondChanged is a free log subscription operation binding the contract event 0xd1fc3f2cb1387e602db0e6f8f22649df65df5246eeff281cf6d1ef62feda4ece.
//
// Solidity: event NoShowBondChanged(address indexed universe, uint256 noShowBond)
func (_Augur *AugurFilterer) WatchNoShowBondChanged(opts *bind.WatchOpts, sink chan<- *AugurNoShowBondChanged, universe []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "NoShowBondChanged", universeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurNoShowBondChanged)
				if err := _Augur.contract.UnpackLog(event, "NoShowBondChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNoShowBondChanged is a log parse operation binding the contract event 0xd1fc3f2cb1387e602db0e6f8f22649df65df5246eeff281cf6d1ef62feda4ece.
//
// Solidity: event NoShowBondChanged(address indexed universe, uint256 noShowBond)
func (_Augur *AugurFilterer) ParseNoShowBondChanged(log types.Log) (*AugurNoShowBondChanged, error) {
	event := new(AugurNoShowBondChanged)
	if err := _Augur.contract.UnpackLog(event, "NoShowBondChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurParticipationTokensRedeemedIterator is returned from FilterParticipationTokensRedeemed and is used to iterate over the raw logs and unpacked data for ParticipationTokensRedeemed events raised by the Augur contract.
type AugurParticipationTokensRedeemedIterator struct {
	Event *AugurParticipationTokensRedeemed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurParticipationTokensRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurParticipationTokensRedeemed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurParticipationTokensRedeemed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurParticipationTokensRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurParticipationTokensRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurParticipationTokensRedeemed represents a ParticipationTokensRedeemed event raised by the Augur contract.
type AugurParticipationTokensRedeemed struct {
	Universe                common.Address
	DisputeWindow           common.Address
	Account                 common.Address
	AttoParticipationTokens *big.Int
	FeePayoutShare          *big.Int
	Timestamp               *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterParticipationTokensRedeemed is a free log retrieval operation binding the contract event 0x18052b5e29020458e154999fa71891a5db3404a5b0b9c5ec60c90adca7d38d63.
//
// Solidity: event ParticipationTokensRedeemed(address indexed universe, address indexed disputeWindow, address indexed account, uint256 attoParticipationTokens, uint256 feePayoutShare, uint256 timestamp)
func (_Augur *AugurFilterer) FilterParticipationTokensRedeemed(opts *bind.FilterOpts, universe []common.Address, disputeWindow []common.Address, account []common.Address) (*AugurParticipationTokensRedeemedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var disputeWindowRule []interface{}
	for _, disputeWindowItem := range disputeWindow {
		disputeWindowRule = append(disputeWindowRule, disputeWindowItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "ParticipationTokensRedeemed", universeRule, disputeWindowRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &AugurParticipationTokensRedeemedIterator{contract: _Augur.contract, event: "ParticipationTokensRedeemed", logs: logs, sub: sub}, nil
}

// WatchParticipationTokensRedeemed is a free log subscription operation binding the contract event 0x18052b5e29020458e154999fa71891a5db3404a5b0b9c5ec60c90adca7d38d63.
//
// Solidity: event ParticipationTokensRedeemed(address indexed universe, address indexed disputeWindow, address indexed account, uint256 attoParticipationTokens, uint256 feePayoutShare, uint256 timestamp)
func (_Augur *AugurFilterer) WatchParticipationTokensRedeemed(opts *bind.WatchOpts, sink chan<- *AugurParticipationTokensRedeemed, universe []common.Address, disputeWindow []common.Address, account []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var disputeWindowRule []interface{}
	for _, disputeWindowItem := range disputeWindow {
		disputeWindowRule = append(disputeWindowRule, disputeWindowItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "ParticipationTokensRedeemed", universeRule, disputeWindowRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurParticipationTokensRedeemed)
				if err := _Augur.contract.UnpackLog(event, "ParticipationTokensRedeemed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseParticipationTokensRedeemed is a log parse operation binding the contract event 0x18052b5e29020458e154999fa71891a5db3404a5b0b9c5ec60c90adca7d38d63.
//
// Solidity: event ParticipationTokensRedeemed(address indexed universe, address indexed disputeWindow, address indexed account, uint256 attoParticipationTokens, uint256 feePayoutShare, uint256 timestamp)
func (_Augur *AugurFilterer) ParseParticipationTokensRedeemed(log types.Log) (*AugurParticipationTokensRedeemed, error) {
	event := new(AugurParticipationTokensRedeemed)
	if err := _Augur.contract.UnpackLog(event, "ParticipationTokensRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurRegisterContractIterator is returned from FilterRegisterContract and is used to iterate over the raw logs and unpacked data for RegisterContract events raised by the Augur contract.
type AugurRegisterContractIterator struct {
	Event *AugurRegisterContract // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurRegisterContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurRegisterContract)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurRegisterContract)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurRegisterContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurRegisterContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurRegisterContract represents a RegisterContract event raised by the Augur contract.
type AugurRegisterContract struct {
	ContractAddress common.Address
	Key             [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRegisterContract is a free log retrieval operation binding the contract event 0xa037dd0e01f0488a530cb17065a6d2f284fae016004fc744ee2a41d5cacf85d5.
//
// Solidity: event RegisterContract(address contractAddress, bytes32 key)
func (_Augur *AugurFilterer) FilterRegisterContract(opts *bind.FilterOpts) (*AugurRegisterContractIterator, error) {

	logs, sub, err := _Augur.contract.FilterLogs(opts, "RegisterContract")
	if err != nil {
		return nil, err
	}
	return &AugurRegisterContractIterator{contract: _Augur.contract, event: "RegisterContract", logs: logs, sub: sub}, nil
}

// WatchRegisterContract is a free log subscription operation binding the contract event 0xa037dd0e01f0488a530cb17065a6d2f284fae016004fc744ee2a41d5cacf85d5.
//
// Solidity: event RegisterContract(address contractAddress, bytes32 key)
func (_Augur *AugurFilterer) WatchRegisterContract(opts *bind.WatchOpts, sink chan<- *AugurRegisterContract) (event.Subscription, error) {

	logs, sub, err := _Augur.contract.WatchLogs(opts, "RegisterContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurRegisterContract)
				if err := _Augur.contract.UnpackLog(event, "RegisterContract", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRegisterContract is a log parse operation binding the contract event 0xa037dd0e01f0488a530cb17065a6d2f284fae016004fc744ee2a41d5cacf85d5.
//
// Solidity: event RegisterContract(address contractAddress, bytes32 key)
func (_Augur *AugurFilterer) ParseRegisterContract(log types.Log) (*AugurRegisterContract, error) {
	event := new(AugurRegisterContract)
	if err := _Augur.contract.UnpackLog(event, "RegisterContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurReportingFeeChangedIterator is returned from FilterReportingFeeChanged and is used to iterate over the raw logs and unpacked data for ReportingFeeChanged events raised by the Augur contract.
type AugurReportingFeeChangedIterator struct {
	Event *AugurReportingFeeChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurReportingFeeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurReportingFeeChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurReportingFeeChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurReportingFeeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurReportingFeeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurReportingFeeChanged represents a ReportingFeeChanged event raised by the Augur contract.
type AugurReportingFeeChanged struct {
	Universe     common.Address
	ReportingFee *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterReportingFeeChanged is a free log retrieval operation binding the contract event 0xadddfaec4505d90a6a211907536944e6e1af7ff5cf6d1873de43e36020f36009.
//
// Solidity: event ReportingFeeChanged(address indexed universe, uint256 reportingFee)
func (_Augur *AugurFilterer) FilterReportingFeeChanged(opts *bind.FilterOpts, universe []common.Address) (*AugurReportingFeeChangedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "ReportingFeeChanged", universeRule)
	if err != nil {
		return nil, err
	}
	return &AugurReportingFeeChangedIterator{contract: _Augur.contract, event: "ReportingFeeChanged", logs: logs, sub: sub}, nil
}

// WatchReportingFeeChanged is a free log subscription operation binding the contract event 0xadddfaec4505d90a6a211907536944e6e1af7ff5cf6d1873de43e36020f36009.
//
// Solidity: event ReportingFeeChanged(address indexed universe, uint256 reportingFee)
func (_Augur *AugurFilterer) WatchReportingFeeChanged(opts *bind.WatchOpts, sink chan<- *AugurReportingFeeChanged, universe []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "ReportingFeeChanged", universeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurReportingFeeChanged)
				if err := _Augur.contract.UnpackLog(event, "ReportingFeeChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReportingFeeChanged is a log parse operation binding the contract event 0xadddfaec4505d90a6a211907536944e6e1af7ff5cf6d1873de43e36020f36009.
//
// Solidity: event ReportingFeeChanged(address indexed universe, uint256 reportingFee)
func (_Augur *AugurFilterer) ParseReportingFeeChanged(log types.Log) (*AugurReportingFeeChanged, error) {
	event := new(AugurReportingFeeChanged)
	if err := _Augur.contract.UnpackLog(event, "ReportingFeeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurReportingParticipantDisavowedIterator is returned from FilterReportingParticipantDisavowed and is used to iterate over the raw logs and unpacked data for ReportingParticipantDisavowed events raised by the Augur contract.
type AugurReportingParticipantDisavowedIterator struct {
	Event *AugurReportingParticipantDisavowed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurReportingParticipantDisavowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurReportingParticipantDisavowed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurReportingParticipantDisavowed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurReportingParticipantDisavowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurReportingParticipantDisavowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurReportingParticipantDisavowed represents a ReportingParticipantDisavowed event raised by the Augur contract.
type AugurReportingParticipantDisavowed struct {
	Universe             common.Address
	Market               common.Address
	ReportingParticipant common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterReportingParticipantDisavowed is a free log retrieval operation binding the contract event 0xb20adf682c8f82b94a135452f54ac4483c9ee8c9b2324e946120696ab1d034b4.
//
// Solidity: event ReportingParticipantDisavowed(address indexed universe, address indexed market, address reportingParticipant)
func (_Augur *AugurFilterer) FilterReportingParticipantDisavowed(opts *bind.FilterOpts, universe []common.Address, market []common.Address) (*AugurReportingParticipantDisavowedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "ReportingParticipantDisavowed", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &AugurReportingParticipantDisavowedIterator{contract: _Augur.contract, event: "ReportingParticipantDisavowed", logs: logs, sub: sub}, nil
}

// WatchReportingParticipantDisavowed is a free log subscription operation binding the contract event 0xb20adf682c8f82b94a135452f54ac4483c9ee8c9b2324e946120696ab1d034b4.
//
// Solidity: event ReportingParticipantDisavowed(address indexed universe, address indexed market, address reportingParticipant)
func (_Augur *AugurFilterer) WatchReportingParticipantDisavowed(opts *bind.WatchOpts, sink chan<- *AugurReportingParticipantDisavowed, universe []common.Address, market []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "ReportingParticipantDisavowed", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurReportingParticipantDisavowed)
				if err := _Augur.contract.UnpackLog(event, "ReportingParticipantDisavowed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReportingParticipantDisavowed is a log parse operation binding the contract event 0xb20adf682c8f82b94a135452f54ac4483c9ee8c9b2324e946120696ab1d034b4.
//
// Solidity: event ReportingParticipantDisavowed(address indexed universe, address indexed market, address reportingParticipant)
func (_Augur *AugurFilterer) ParseReportingParticipantDisavowed(log types.Log) (*AugurReportingParticipantDisavowed, error) {
	event := new(AugurReportingParticipantDisavowed)
	if err := _Augur.contract.UnpackLog(event, "ReportingParticipantDisavowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurShareTokenBalanceChangedIterator is returned from FilterShareTokenBalanceChanged and is used to iterate over the raw logs and unpacked data for ShareTokenBalanceChanged events raised by the Augur contract.
type AugurShareTokenBalanceChangedIterator struct {
	Event *AugurShareTokenBalanceChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurShareTokenBalanceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurShareTokenBalanceChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurShareTokenBalanceChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurShareTokenBalanceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurShareTokenBalanceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurShareTokenBalanceChanged represents a ShareTokenBalanceChanged event raised by the Augur contract.
type AugurShareTokenBalanceChanged struct {
	Universe common.Address
	Account  common.Address
	Market   common.Address
	Outcome  *big.Int
	Balance  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterShareTokenBalanceChanged is a free log retrieval operation binding the contract event 0x350ea32dc29530b9557420816d743c436f8397086f98c96292138edd69e01cb3.
//
// Solidity: event ShareTokenBalanceChanged(address indexed universe, address indexed account, address indexed market, uint256 outcome, uint256 balance)
func (_Augur *AugurFilterer) FilterShareTokenBalanceChanged(opts *bind.FilterOpts, universe []common.Address, account []common.Address, market []common.Address) (*AugurShareTokenBalanceChangedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "ShareTokenBalanceChanged", universeRule, accountRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &AugurShareTokenBalanceChangedIterator{contract: _Augur.contract, event: "ShareTokenBalanceChanged", logs: logs, sub: sub}, nil
}

// WatchShareTokenBalanceChanged is a free log subscription operation binding the contract event 0x350ea32dc29530b9557420816d743c436f8397086f98c96292138edd69e01cb3.
//
// Solidity: event ShareTokenBalanceChanged(address indexed universe, address indexed account, address indexed market, uint256 outcome, uint256 balance)
func (_Augur *AugurFilterer) WatchShareTokenBalanceChanged(opts *bind.WatchOpts, sink chan<- *AugurShareTokenBalanceChanged, universe []common.Address, account []common.Address, market []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "ShareTokenBalanceChanged", universeRule, accountRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurShareTokenBalanceChanged)
				if err := _Augur.contract.UnpackLog(event, "ShareTokenBalanceChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseShareTokenBalanceChanged is a log parse operation binding the contract event 0x350ea32dc29530b9557420816d743c436f8397086f98c96292138edd69e01cb3.
//
// Solidity: event ShareTokenBalanceChanged(address indexed universe, address indexed account, address indexed market, uint256 outcome, uint256 balance)
func (_Augur *AugurFilterer) ParseShareTokenBalanceChanged(log types.Log) (*AugurShareTokenBalanceChanged, error) {
	event := new(AugurShareTokenBalanceChanged)
	if err := _Augur.contract.UnpackLog(event, "ShareTokenBalanceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurTimestampSetIterator is returned from FilterTimestampSet and is used to iterate over the raw logs and unpacked data for TimestampSet events raised by the Augur contract.
type AugurTimestampSetIterator struct {
	Event *AugurTimestampSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurTimestampSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTimestampSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurTimestampSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurTimestampSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTimestampSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTimestampSet represents a TimestampSet event raised by the Augur contract.
type AugurTimestampSet struct {
	NewTimestamp *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTimestampSet is a free log retrieval operation binding the contract event 0x11dda748f0bd3af85a073da0088a0acb827d9584a4fdb825c81f1232a5309538.
//
// Solidity: event TimestampSet(uint256 newTimestamp)
func (_Augur *AugurFilterer) FilterTimestampSet(opts *bind.FilterOpts) (*AugurTimestampSetIterator, error) {

	logs, sub, err := _Augur.contract.FilterLogs(opts, "TimestampSet")
	if err != nil {
		return nil, err
	}
	return &AugurTimestampSetIterator{contract: _Augur.contract, event: "TimestampSet", logs: logs, sub: sub}, nil
}

// WatchTimestampSet is a free log subscription operation binding the contract event 0x11dda748f0bd3af85a073da0088a0acb827d9584a4fdb825c81f1232a5309538.
//
// Solidity: event TimestampSet(uint256 newTimestamp)
func (_Augur *AugurFilterer) WatchTimestampSet(opts *bind.WatchOpts, sink chan<- *AugurTimestampSet) (event.Subscription, error) {

	logs, sub, err := _Augur.contract.WatchLogs(opts, "TimestampSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTimestampSet)
				if err := _Augur.contract.UnpackLog(event, "TimestampSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTimestampSet is a log parse operation binding the contract event 0x11dda748f0bd3af85a073da0088a0acb827d9584a4fdb825c81f1232a5309538.
//
// Solidity: event TimestampSet(uint256 newTimestamp)
func (_Augur *AugurFilterer) ParseTimestampSet(log types.Log) (*AugurTimestampSet, error) {
	event := new(AugurTimestampSet)
	if err := _Augur.contract.UnpackLog(event, "TimestampSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurTokenBalanceChangedIterator is returned from FilterTokenBalanceChanged and is used to iterate over the raw logs and unpacked data for TokenBalanceChanged events raised by the Augur contract.
type AugurTokenBalanceChangedIterator struct {
	Event *AugurTokenBalanceChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurTokenBalanceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTokenBalanceChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurTokenBalanceChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurTokenBalanceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTokenBalanceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTokenBalanceChanged represents a TokenBalanceChanged event raised by the Augur contract.
type AugurTokenBalanceChanged struct {
	Universe  common.Address
	Owner     common.Address
	Token     common.Address
	TokenType uint8
	Market    common.Address
	Balance   *big.Int
	Outcome   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenBalanceChanged is a free log retrieval operation binding the contract event 0x63fd58f559b73fc4da5511c341ec8a7b31c5c48538ef83c6077712b6edf5f7cb.
//
// Solidity: event TokenBalanceChanged(address indexed universe, address indexed owner, address token, uint8 tokenType, address market, uint256 balance, uint256 outcome)
func (_Augur *AugurFilterer) FilterTokenBalanceChanged(opts *bind.FilterOpts, universe []common.Address, owner []common.Address) (*AugurTokenBalanceChangedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "TokenBalanceChanged", universeRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &AugurTokenBalanceChangedIterator{contract: _Augur.contract, event: "TokenBalanceChanged", logs: logs, sub: sub}, nil
}

// WatchTokenBalanceChanged is a free log subscription operation binding the contract event 0x63fd58f559b73fc4da5511c341ec8a7b31c5c48538ef83c6077712b6edf5f7cb.
//
// Solidity: event TokenBalanceChanged(address indexed universe, address indexed owner, address token, uint8 tokenType, address market, uint256 balance, uint256 outcome)
func (_Augur *AugurFilterer) WatchTokenBalanceChanged(opts *bind.WatchOpts, sink chan<- *AugurTokenBalanceChanged, universe []common.Address, owner []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "TokenBalanceChanged", universeRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTokenBalanceChanged)
				if err := _Augur.contract.UnpackLog(event, "TokenBalanceChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenBalanceChanged is a log parse operation binding the contract event 0x63fd58f559b73fc4da5511c341ec8a7b31c5c48538ef83c6077712b6edf5f7cb.
//
// Solidity: event TokenBalanceChanged(address indexed universe, address indexed owner, address token, uint8 tokenType, address market, uint256 balance, uint256 outcome)
func (_Augur *AugurFilterer) ParseTokenBalanceChanged(log types.Log) (*AugurTokenBalanceChanged, error) {
	event := new(AugurTokenBalanceChanged)
	if err := _Augur.contract.UnpackLog(event, "TokenBalanceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurTokensBurnedIterator is returned from FilterTokensBurned and is used to iterate over the raw logs and unpacked data for TokensBurned events raised by the Augur contract.
type AugurTokensBurnedIterator struct {
	Event *AugurTokensBurned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurTokensBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTokensBurned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurTokensBurned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurTokensBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTokensBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTokensBurned represents a TokensBurned event raised by the Augur contract.
type AugurTokensBurned struct {
	Universe    common.Address
	Token       common.Address
	Target      common.Address
	Amount      *big.Int
	TokenType   uint8
	Market      common.Address
	TotalSupply *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensBurned is a free log retrieval operation binding the contract event 0x145a4839b3d82d1e28f6ed93f52622b351892e835530386bb1fe4effba99aeea.
//
// Solidity: event TokensBurned(address indexed universe, address indexed token, address indexed target, uint256 amount, uint8 tokenType, address market, uint256 totalSupply)
func (_Augur *AugurFilterer) FilterTokensBurned(opts *bind.FilterOpts, universe []common.Address, token []common.Address, target []common.Address) (*AugurTokensBurnedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "TokensBurned", universeRule, tokenRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &AugurTokensBurnedIterator{contract: _Augur.contract, event: "TokensBurned", logs: logs, sub: sub}, nil
}

// WatchTokensBurned is a free log subscription operation binding the contract event 0x145a4839b3d82d1e28f6ed93f52622b351892e835530386bb1fe4effba99aeea.
//
// Solidity: event TokensBurned(address indexed universe, address indexed token, address indexed target, uint256 amount, uint8 tokenType, address market, uint256 totalSupply)
func (_Augur *AugurFilterer) WatchTokensBurned(opts *bind.WatchOpts, sink chan<- *AugurTokensBurned, universe []common.Address, token []common.Address, target []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "TokensBurned", universeRule, tokenRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTokensBurned)
				if err := _Augur.contract.UnpackLog(event, "TokensBurned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensBurned is a log parse operation binding the contract event 0x145a4839b3d82d1e28f6ed93f52622b351892e835530386bb1fe4effba99aeea.
//
// Solidity: event TokensBurned(address indexed universe, address indexed token, address indexed target, uint256 amount, uint8 tokenType, address market, uint256 totalSupply)
func (_Augur *AugurFilterer) ParseTokensBurned(log types.Log) (*AugurTokensBurned, error) {
	event := new(AugurTokensBurned)
	if err := _Augur.contract.UnpackLog(event, "TokensBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurTokensMintedIterator is returned from FilterTokensMinted and is used to iterate over the raw logs and unpacked data for TokensMinted events raised by the Augur contract.
type AugurTokensMintedIterator struct {
	Event *AugurTokensMinted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurTokensMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTokensMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurTokensMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurTokensMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTokensMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTokensMinted represents a TokensMinted event raised by the Augur contract.
type AugurTokensMinted struct {
	Universe    common.Address
	Token       common.Address
	Target      common.Address
	Amount      *big.Int
	TokenType   uint8
	Market      common.Address
	TotalSupply *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensMinted is a free log retrieval operation binding the contract event 0x07f766729171db8cc73d96b25cc56784077e26c7ff48b0187877ace391c181a6.
//
// Solidity: event TokensMinted(address indexed universe, address indexed token, address indexed target, uint256 amount, uint8 tokenType, address market, uint256 totalSupply)
func (_Augur *AugurFilterer) FilterTokensMinted(opts *bind.FilterOpts, universe []common.Address, token []common.Address, target []common.Address) (*AugurTokensMintedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "TokensMinted", universeRule, tokenRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &AugurTokensMintedIterator{contract: _Augur.contract, event: "TokensMinted", logs: logs, sub: sub}, nil
}

// WatchTokensMinted is a free log subscription operation binding the contract event 0x07f766729171db8cc73d96b25cc56784077e26c7ff48b0187877ace391c181a6.
//
// Solidity: event TokensMinted(address indexed universe, address indexed token, address indexed target, uint256 amount, uint8 tokenType, address market, uint256 totalSupply)
func (_Augur *AugurFilterer) WatchTokensMinted(opts *bind.WatchOpts, sink chan<- *AugurTokensMinted, universe []common.Address, token []common.Address, target []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "TokensMinted", universeRule, tokenRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTokensMinted)
				if err := _Augur.contract.UnpackLog(event, "TokensMinted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensMinted is a log parse operation binding the contract event 0x07f766729171db8cc73d96b25cc56784077e26c7ff48b0187877ace391c181a6.
//
// Solidity: event TokensMinted(address indexed universe, address indexed token, address indexed target, uint256 amount, uint8 tokenType, address market, uint256 totalSupply)
func (_Augur *AugurFilterer) ParseTokensMinted(log types.Log) (*AugurTokensMinted, error) {
	event := new(AugurTokensMinted)
	if err := _Augur.contract.UnpackLog(event, "TokensMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurTokensTransferredIterator is returned from FilterTokensTransferred and is used to iterate over the raw logs and unpacked data for TokensTransferred events raised by the Augur contract.
type AugurTokensTransferredIterator struct {
	Event *AugurTokensTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurTokensTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTokensTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurTokensTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurTokensTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTokensTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTokensTransferred represents a TokensTransferred event raised by the Augur contract.
type AugurTokensTransferred struct {
	Universe  common.Address
	Token     common.Address
	From      common.Address
	To        common.Address
	Value     *big.Int
	TokenType uint8
	Market    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensTransferred is a free log retrieval operation binding the contract event 0x3c67396e9c55d2fc8ad68875fc5beca1d96ad2a2f23b210ccc1d986551ab6fdf.
//
// Solidity: event TokensTransferred(address indexed universe, address token, address indexed from, address indexed to, uint256 value, uint8 tokenType, address market)
func (_Augur *AugurFilterer) FilterTokensTransferred(opts *bind.FilterOpts, universe []common.Address, from []common.Address, to []common.Address) (*AugurTokensTransferredIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "TokensTransferred", universeRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AugurTokensTransferredIterator{contract: _Augur.contract, event: "TokensTransferred", logs: logs, sub: sub}, nil
}

// WatchTokensTransferred is a free log subscription operation binding the contract event 0x3c67396e9c55d2fc8ad68875fc5beca1d96ad2a2f23b210ccc1d986551ab6fdf.
//
// Solidity: event TokensTransferred(address indexed universe, address token, address indexed from, address indexed to, uint256 value, uint8 tokenType, address market)
func (_Augur *AugurFilterer) WatchTokensTransferred(opts *bind.WatchOpts, sink chan<- *AugurTokensTransferred, universe []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "TokensTransferred", universeRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTokensTransferred)
				if err := _Augur.contract.UnpackLog(event, "TokensTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensTransferred is a log parse operation binding the contract event 0x3c67396e9c55d2fc8ad68875fc5beca1d96ad2a2f23b210ccc1d986551ab6fdf.
//
// Solidity: event TokensTransferred(address indexed universe, address token, address indexed from, address indexed to, uint256 value, uint8 tokenType, address market)
func (_Augur *AugurFilterer) ParseTokensTransferred(log types.Log) (*AugurTokensTransferred, error) {
	event := new(AugurTokensTransferred)
	if err := _Augur.contract.UnpackLog(event, "TokensTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurTradingProceedsClaimedIterator is returned from FilterTradingProceedsClaimed and is used to iterate over the raw logs and unpacked data for TradingProceedsClaimed events raised by the Augur contract.
type AugurTradingProceedsClaimedIterator struct {
	Event *AugurTradingProceedsClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurTradingProceedsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTradingProceedsClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurTradingProceedsClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurTradingProceedsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTradingProceedsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTradingProceedsClaimed represents a TradingProceedsClaimed event raised by the Augur contract.
type AugurTradingProceedsClaimed struct {
	Universe        common.Address
	Sender          common.Address
	Market          common.Address
	Outcome         *big.Int
	NumShares       *big.Int
	NumPayoutTokens *big.Int
	Fees            *big.Int
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTradingProceedsClaimed is a free log retrieval operation binding the contract event 0x95366b7f64c6bb45149f9f7c522403fceebe5170ff76b8ffde2b0ab943ac11ce.
//
// Solidity: event TradingProceedsClaimed(address indexed universe, address indexed sender, address market, uint256 outcome, uint256 numShares, uint256 numPayoutTokens, uint256 fees, uint256 timestamp)
func (_Augur *AugurFilterer) FilterTradingProceedsClaimed(opts *bind.FilterOpts, universe []common.Address, sender []common.Address) (*AugurTradingProceedsClaimedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "TradingProceedsClaimed", universeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AugurTradingProceedsClaimedIterator{contract: _Augur.contract, event: "TradingProceedsClaimed", logs: logs, sub: sub}, nil
}

// WatchTradingProceedsClaimed is a free log subscription operation binding the contract event 0x95366b7f64c6bb45149f9f7c522403fceebe5170ff76b8ffde2b0ab943ac11ce.
//
// Solidity: event TradingProceedsClaimed(address indexed universe, address indexed sender, address market, uint256 outcome, uint256 numShares, uint256 numPayoutTokens, uint256 fees, uint256 timestamp)
func (_Augur *AugurFilterer) WatchTradingProceedsClaimed(opts *bind.WatchOpts, sink chan<- *AugurTradingProceedsClaimed, universe []common.Address, sender []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "TradingProceedsClaimed", universeRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTradingProceedsClaimed)
				if err := _Augur.contract.UnpackLog(event, "TradingProceedsClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTradingProceedsClaimed is a log parse operation binding the contract event 0x95366b7f64c6bb45149f9f7c522403fceebe5170ff76b8ffde2b0ab943ac11ce.
//
// Solidity: event TradingProceedsClaimed(address indexed universe, address indexed sender, address market, uint256 outcome, uint256 numShares, uint256 numPayoutTokens, uint256 fees, uint256 timestamp)
func (_Augur *AugurFilterer) ParseTradingProceedsClaimed(log types.Log) (*AugurTradingProceedsClaimed, error) {
	event := new(AugurTradingProceedsClaimed)
	if err := _Augur.contract.UnpackLog(event, "TradingProceedsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurUniverseCreatedIterator is returned from FilterUniverseCreated and is used to iterate over the raw logs and unpacked data for UniverseCreated events raised by the Augur contract.
type AugurUniverseCreatedIterator struct {
	Event *AugurUniverseCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurUniverseCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurUniverseCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurUniverseCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurUniverseCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurUniverseCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurUniverseCreated represents a UniverseCreated event raised by the Augur contract.
type AugurUniverseCreated struct {
	ParentUniverse    common.Address
	ChildUniverse     common.Address
	PayoutNumerators  []*big.Int
	CreationTimestamp *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUniverseCreated is a free log retrieval operation binding the contract event 0xe36b09d83f9cfa88c37f071fc2cfb5ff30b764cbd98088e70d965573c9ce5bbd.
//
// Solidity: event UniverseCreated(address indexed parentUniverse, address indexed childUniverse, uint256[] payoutNumerators, uint256 creationTimestamp)
func (_Augur *AugurFilterer) FilterUniverseCreated(opts *bind.FilterOpts, parentUniverse []common.Address, childUniverse []common.Address) (*AugurUniverseCreatedIterator, error) {

	var parentUniverseRule []interface{}
	for _, parentUniverseItem := range parentUniverse {
		parentUniverseRule = append(parentUniverseRule, parentUniverseItem)
	}
	var childUniverseRule []interface{}
	for _, childUniverseItem := range childUniverse {
		childUniverseRule = append(childUniverseRule, childUniverseItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "UniverseCreated", parentUniverseRule, childUniverseRule)
	if err != nil {
		return nil, err
	}
	return &AugurUniverseCreatedIterator{contract: _Augur.contract, event: "UniverseCreated", logs: logs, sub: sub}, nil
}

// WatchUniverseCreated is a free log subscription operation binding the contract event 0xe36b09d83f9cfa88c37f071fc2cfb5ff30b764cbd98088e70d965573c9ce5bbd.
//
// Solidity: event UniverseCreated(address indexed parentUniverse, address indexed childUniverse, uint256[] payoutNumerators, uint256 creationTimestamp)
func (_Augur *AugurFilterer) WatchUniverseCreated(opts *bind.WatchOpts, sink chan<- *AugurUniverseCreated, parentUniverse []common.Address, childUniverse []common.Address) (event.Subscription, error) {

	var parentUniverseRule []interface{}
	for _, parentUniverseItem := range parentUniverse {
		parentUniverseRule = append(parentUniverseRule, parentUniverseItem)
	}
	var childUniverseRule []interface{}
	for _, childUniverseItem := range childUniverse {
		childUniverseRule = append(childUniverseRule, childUniverseItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "UniverseCreated", parentUniverseRule, childUniverseRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurUniverseCreated)
				if err := _Augur.contract.UnpackLog(event, "UniverseCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUniverseCreated is a log parse operation binding the contract event 0xe36b09d83f9cfa88c37f071fc2cfb5ff30b764cbd98088e70d965573c9ce5bbd.
//
// Solidity: event UniverseCreated(address indexed parentUniverse, address indexed childUniverse, uint256[] payoutNumerators, uint256 creationTimestamp)
func (_Augur *AugurFilterer) ParseUniverseCreated(log types.Log) (*AugurUniverseCreated, error) {
	event := new(AugurUniverseCreated)
	if err := _Augur.contract.UnpackLog(event, "UniverseCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurUniverseForkedIterator is returned from FilterUniverseForked and is used to iterate over the raw logs and unpacked data for UniverseForked events raised by the Augur contract.
type AugurUniverseForkedIterator struct {
	Event *AugurUniverseForked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurUniverseForkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurUniverseForked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurUniverseForked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurUniverseForkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurUniverseForkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurUniverseForked represents a UniverseForked event raised by the Augur contract.
type AugurUniverseForked struct {
	Universe      common.Address
	ForkingMarket common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUniverseForked is a free log retrieval operation binding the contract event 0xce5b6de2a0053ebc6c04e68bcbb9f0a1f2deeb7049c72881e198f95b5752db82.
//
// Solidity: event UniverseForked(address indexed universe, address forkingMarket)
func (_Augur *AugurFilterer) FilterUniverseForked(opts *bind.FilterOpts, universe []common.Address) (*AugurUniverseForkedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "UniverseForked", universeRule)
	if err != nil {
		return nil, err
	}
	return &AugurUniverseForkedIterator{contract: _Augur.contract, event: "UniverseForked", logs: logs, sub: sub}, nil
}

// WatchUniverseForked is a free log subscription operation binding the contract event 0xce5b6de2a0053ebc6c04e68bcbb9f0a1f2deeb7049c72881e198f95b5752db82.
//
// Solidity: event UniverseForked(address indexed universe, address forkingMarket)
func (_Augur *AugurFilterer) WatchUniverseForked(opts *bind.WatchOpts, sink chan<- *AugurUniverseForked, universe []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "UniverseForked", universeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurUniverseForked)
				if err := _Augur.contract.UnpackLog(event, "UniverseForked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUniverseForked is a log parse operation binding the contract event 0xce5b6de2a0053ebc6c04e68bcbb9f0a1f2deeb7049c72881e198f95b5752db82.
//
// Solidity: event UniverseForked(address indexed universe, address forkingMarket)
func (_Augur *AugurFilterer) ParseUniverseForked(log types.Log) (*AugurUniverseForked, error) {
	event := new(AugurUniverseForked)
	if err := _Augur.contract.UnpackLog(event, "UniverseForked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurValidityBondChangedIterator is returned from FilterValidityBondChanged and is used to iterate over the raw logs and unpacked data for ValidityBondChanged events raised by the Augur contract.
type AugurValidityBondChangedIterator struct {
	Event *AugurValidityBondChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurValidityBondChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurValidityBondChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurValidityBondChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurValidityBondChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurValidityBondChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurValidityBondChanged represents a ValidityBondChanged event raised by the Augur contract.
type AugurValidityBondChanged struct {
	Universe     common.Address
	ValidityBond *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidityBondChanged is a free log retrieval operation binding the contract event 0x69af68e366a0570364e3a086f3b5ac79f08ecc3f93eaccbfcf3864809b12b5d8.
//
// Solidity: event ValidityBondChanged(address indexed universe, uint256 validityBond)
func (_Augur *AugurFilterer) FilterValidityBondChanged(opts *bind.FilterOpts, universe []common.Address) (*AugurValidityBondChangedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "ValidityBondChanged", universeRule)
	if err != nil {
		return nil, err
	}
	return &AugurValidityBondChangedIterator{contract: _Augur.contract, event: "ValidityBondChanged", logs: logs, sub: sub}, nil
}

// WatchValidityBondChanged is a free log subscription operation binding the contract event 0x69af68e366a0570364e3a086f3b5ac79f08ecc3f93eaccbfcf3864809b12b5d8.
//
// Solidity: event ValidityBondChanged(address indexed universe, uint256 validityBond)
func (_Augur *AugurFilterer) WatchValidityBondChanged(opts *bind.WatchOpts, sink chan<- *AugurValidityBondChanged, universe []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "ValidityBondChanged", universeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurValidityBondChanged)
				if err := _Augur.contract.UnpackLog(event, "ValidityBondChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidityBondChanged is a log parse operation binding the contract event 0x69af68e366a0570364e3a086f3b5ac79f08ecc3f93eaccbfcf3864809b12b5d8.
//
// Solidity: event ValidityBondChanged(address indexed universe, uint256 validityBond)
func (_Augur *AugurFilterer) ParseValidityBondChanged(log types.Log) (*AugurValidityBondChanged, error) {
	event := new(AugurValidityBondChanged)
	if err := _Augur.contract.UnpackLog(event, "ValidityBondChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurWarpSyncDataUpdatedIterator is returned from FilterWarpSyncDataUpdated and is used to iterate over the raw logs and unpacked data for WarpSyncDataUpdated events raised by the Augur contract.
type AugurWarpSyncDataUpdatedIterator struct {
	Event *AugurWarpSyncDataUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AugurWarpSyncDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurWarpSyncDataUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AugurWarpSyncDataUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AugurWarpSyncDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurWarpSyncDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurWarpSyncDataUpdated represents a WarpSyncDataUpdated event raised by the Augur contract.
type AugurWarpSyncDataUpdated struct {
	Universe      common.Address
	WarpSyncHash  *big.Int
	MarketEndTime *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWarpSyncDataUpdated is a free log retrieval operation binding the contract event 0x7589653fe5a2ab3ccc12538316852339868efdd9d3bd0b84d055cf224cf96873.
//
// Solidity: event WarpSyncDataUpdated(address indexed universe, uint256 warpSyncHash, uint256 marketEndTime)
func (_Augur *AugurFilterer) FilterWarpSyncDataUpdated(opts *bind.FilterOpts, universe []common.Address) (*AugurWarpSyncDataUpdatedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}

	logs, sub, err := _Augur.contract.FilterLogs(opts, "WarpSyncDataUpdated", universeRule)
	if err != nil {
		return nil, err
	}
	return &AugurWarpSyncDataUpdatedIterator{contract: _Augur.contract, event: "WarpSyncDataUpdated", logs: logs, sub: sub}, nil
}

// WatchWarpSyncDataUpdated is a free log subscription operation binding the contract event 0x7589653fe5a2ab3ccc12538316852339868efdd9d3bd0b84d055cf224cf96873.
//
// Solidity: event WarpSyncDataUpdated(address indexed universe, uint256 warpSyncHash, uint256 marketEndTime)
func (_Augur *AugurFilterer) WatchWarpSyncDataUpdated(opts *bind.WatchOpts, sink chan<- *AugurWarpSyncDataUpdated, universe []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}

	logs, sub, err := _Augur.contract.WatchLogs(opts, "WarpSyncDataUpdated", universeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurWarpSyncDataUpdated)
				if err := _Augur.contract.UnpackLog(event, "WarpSyncDataUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWarpSyncDataUpdated is a log parse operation binding the contract event 0x7589653fe5a2ab3ccc12538316852339868efdd9d3bd0b84d055cf224cf96873.
//
// Solidity: event WarpSyncDataUpdated(address indexed universe, uint256 warpSyncHash, uint256 marketEndTime)
func (_Augur *AugurFilterer) ParseWarpSyncDataUpdated(log types.Log) (*AugurWarpSyncDataUpdated, error) {
	event := new(AugurWarpSyncDataUpdated)
	if err := _Augur.contract.UnpackLog(event, "WarpSyncDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
