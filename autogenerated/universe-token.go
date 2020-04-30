// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"inputs\":[{\"internalType\":\"contractIAugur\",\"name\":\"_augur\",\"type\":\"address\"},{\"internalType\":\"contractIUniverse\",\"name\":\"_parentUniverse\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_parentPayoutDistributionHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"_payoutNumerators\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"DEFAULT_NUM_OUTCOMES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DEFAULT_NUM_TICKS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INITIAL_WINDOW_ID_BUFFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"augur\",\"outputs\":[{\"internalType\":\"contractIAugur\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cash\",\"outputs\":[{\"internalType\":\"contractICash\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feePerCashInAttoCash\",\"type\":\"uint256\"},{\"internalType\":\"contractIAffiliateValidator\",\"name\":\"_affiliateValidator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_affiliateFeeDivisor\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_designatedReporterAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"_outcomes\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"_extraInfo\",\"type\":\"string\"}],\"name\":\"createCategoricalMarket\",\"outputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_newMarket\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_parentPayoutNumerators\",\"type\":\"uint256[]\"}],\"name\":\"createChildUniverse\",\"outputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feePerCashInAttoCash\",\"type\":\"uint256\"},{\"internalType\":\"contractIAffiliateValidator\",\"name\":\"_affiliateValidator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_affiliateFeeDivisor\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_designatedReporterAddress\",\"type\":\"address\"},{\"internalType\":\"int256[]\",\"name\":\"_prices\",\"type\":\"int256[]\"},{\"internalType\":\"uint256\",\"name\":\"_numTicks\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_extraInfo\",\"type\":\"string\"}],\"name\":\"createScalarMarket\",\"outputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_newMarket\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feePerCashInAttoCash\",\"type\":\"uint256\"},{\"internalType\":\"contractIAffiliateValidator\",\"name\":\"_affiliateValidator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_affiliateFeeDivisor\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_designatedReporterAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_extraInfo\",\"type\":\"string\"}],\"name\":\"createYesNoMarket\",\"outputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_newMarket\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"creationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"daiJoin\",\"outputs\":[{\"internalType\":\"contractIDaiJoin\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"daiPot\",\"outputs\":[{\"internalType\":\"contractIDaiPot\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"daiVat\",\"outputs\":[{\"internalType\":\"contractIDaiVat\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"decrementOpenInterest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"}],\"name\":\"decrementOpenInterestFromMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"designatedReportNoShowBondInAttoRep\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"designatedReportStakeInAttoRep\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"disputeWindowFactory\",\"outputs\":[{\"internalType\":\"contractIDisputeWindowFactory\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"disputeWindows\",\"outputs\":[{\"internalType\":\"contractIDisputeWindow\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"fork\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"formulas\",\"outputs\":[{\"internalType\":\"contractIFormulas\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parentPayoutDistributionHash\",\"type\":\"bytes32\"}],\"name\":\"getChildUniverse\",\"outputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_initial\",\"type\":\"bool\"}],\"name\":\"getCurrentDisputeWindow\",\"outputs\":[{\"internalType\":\"contractIDisputeWindow\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_initial\",\"type\":\"bool\"}],\"name\":\"getDisputeRoundDurationInSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDisputeThresholdForDisputePacing\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDisputeThresholdForFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeWindowId\",\"type\":\"uint256\"}],\"name\":\"getDisputeWindow\",\"outputs\":[{\"internalType\":\"contractIDisputeWindow\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_initial\",\"type\":\"bool\"}],\"name\":\"getDisputeWindowByTimestamp\",\"outputs\":[{\"internalType\":\"contractIDisputeWindow\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_initial\",\"type\":\"bool\"}],\"name\":\"getDisputeWindowId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_initial\",\"type\":\"bool\"}],\"name\":\"getDisputeWindowStartTimeAndDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getForkEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getForkReputationGoal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getForkingMarket\",\"outputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitialReportMinValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOpenInterestInAttoCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getOrCacheDesignatedReportNoShowBond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getOrCacheDesignatedReportStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getOrCacheMarketRepBond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getOrCacheReportingFeeDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getOrCacheValidityBond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_initial\",\"type\":\"bool\"}],\"name\":\"getOrCreateCurrentDisputeWindow\",\"outputs\":[{\"internalType\":\"contractIDisputeWindow\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_initial\",\"type\":\"bool\"}],\"name\":\"getOrCreateDisputeWindowByTimestamp\",\"outputs\":[{\"internalType\":\"contractIDisputeWindow\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_initial\",\"type\":\"bool\"}],\"name\":\"getOrCreateNextDisputeWindow\",\"outputs\":[{\"internalType\":\"contractIDisputeWindow\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_initial\",\"type\":\"bool\"}],\"name\":\"getOrCreatePreviousDisputeWindow\",\"outputs\":[{\"internalType\":\"contractIDisputeWindow\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_initial\",\"type\":\"bool\"}],\"name\":\"getOrCreatePreviousPreviousDisputeWindow\",\"outputs\":[{\"internalType\":\"contractIDisputeWindow\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getParentPayoutDistributionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getParentUniverse\",\"outputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_outcome\",\"type\":\"uint256\"}],\"name\":\"getPayoutNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPayoutNumerators\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReportingFeeDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReputationToken\",\"outputs\":[{\"internalType\":\"contractIV2ReputationToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTargetRepMarketCapInAttoCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_outcome\",\"type\":\"uint256\"}],\"name\":\"getWinningChildPayoutNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getWinningChildUniverse\",\"outputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"incrementOpenInterest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIDisputeWindow\",\"name\":\"_shadyDisputeWindow\",\"type\":\"address\"}],\"name\":\"isContainerForDisputeWindow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_shadyMarket\",\"type\":\"address\"}],\"name\":\"isContainerForMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIReportingParticipant\",\"name\":\"_shadyReportingParticipant\",\"type\":\"address\"}],\"name\":\"isContainerForReportingParticipant\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isForking\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isForkingMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isOpenInterestCash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_shadyChild\",\"type\":\"address\"}],\"name\":\"isParentOf\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastSweep\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"marketBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"marketFactory\",\"outputs\":[{\"internalType\":\"contractIMarketFactory\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_cashBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_marketOI\",\"type\":\"uint256\"}],\"name\":\"migrateMarketIn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_destinationUniverse\",\"type\":\"address\"}],\"name\":\"migrateMarketOut\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"openInterestCash\",\"outputs\":[{\"internalType\":\"contractIOICash\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"payoutNumerators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pokeRepMarketCapInAttoCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"previousDesignatedReportNoShowBondInAttoRep\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"previousDesignatedReportStakeInAttoRep\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"previousReportingFeeDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"previousValidityBondInAttoCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"repOracle\",\"outputs\":[{\"internalType\":\"contractIRepOracle\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"runPeriodicals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"shareSettlementFeeDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"shareToken\",\"outputs\":[{\"internalType\":\"contractIShareToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"sweepInterest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateForkValues\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parentPayoutDistributionHash\",\"type\":\"bytes32\"}],\"name\":\"updateTentativeWinningChildUniverse\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validityBondInAttoCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTNUMOUTCOMES is a free data retrieval call binding the contract method 0xe4372c3c.
//
// Solidity: function DEFAULT_NUM_OUTCOMES() view returns(uint256)
func (_Token *TokenCaller) DEFAULTNUMOUTCOMES(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "DEFAULT_NUM_OUTCOMES")
	return *ret0, err
}

// DEFAULTNUMOUTCOMES is a free data retrieval call binding the contract method 0xe4372c3c.
//
// Solidity: function DEFAULT_NUM_OUTCOMES() view returns(uint256)
func (_Token *TokenSession) DEFAULTNUMOUTCOMES() (*big.Int, error) {
	return _Token.Contract.DEFAULTNUMOUTCOMES(&_Token.CallOpts)
}

// DEFAULTNUMOUTCOMES is a free data retrieval call binding the contract method 0xe4372c3c.
//
// Solidity: function DEFAULT_NUM_OUTCOMES() view returns(uint256)
func (_Token *TokenCallerSession) DEFAULTNUMOUTCOMES() (*big.Int, error) {
	return _Token.Contract.DEFAULTNUMOUTCOMES(&_Token.CallOpts)
}

// DEFAULTNUMTICKS is a free data retrieval call binding the contract method 0x705123dc.
//
// Solidity: function DEFAULT_NUM_TICKS() view returns(uint256)
func (_Token *TokenCaller) DEFAULTNUMTICKS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "DEFAULT_NUM_TICKS")
	return *ret0, err
}

// DEFAULTNUMTICKS is a free data retrieval call binding the contract method 0x705123dc.
//
// Solidity: function DEFAULT_NUM_TICKS() view returns(uint256)
func (_Token *TokenSession) DEFAULTNUMTICKS() (*big.Int, error) {
	return _Token.Contract.DEFAULTNUMTICKS(&_Token.CallOpts)
}

// DEFAULTNUMTICKS is a free data retrieval call binding the contract method 0x705123dc.
//
// Solidity: function DEFAULT_NUM_TICKS() view returns(uint256)
func (_Token *TokenCallerSession) DEFAULTNUMTICKS() (*big.Int, error) {
	return _Token.Contract.DEFAULTNUMTICKS(&_Token.CallOpts)
}

// INITIALWINDOWIDBUFFER is a free data retrieval call binding the contract method 0xf2cf48f6.
//
// Solidity: function INITIAL_WINDOW_ID_BUFFER() view returns(uint256)
func (_Token *TokenCaller) INITIALWINDOWIDBUFFER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "INITIAL_WINDOW_ID_BUFFER")
	return *ret0, err
}

// INITIALWINDOWIDBUFFER is a free data retrieval call binding the contract method 0xf2cf48f6.
//
// Solidity: function INITIAL_WINDOW_ID_BUFFER() view returns(uint256)
func (_Token *TokenSession) INITIALWINDOWIDBUFFER() (*big.Int, error) {
	return _Token.Contract.INITIALWINDOWIDBUFFER(&_Token.CallOpts)
}

// INITIALWINDOWIDBUFFER is a free data retrieval call binding the contract method 0xf2cf48f6.
//
// Solidity: function INITIAL_WINDOW_ID_BUFFER() view returns(uint256)
func (_Token *TokenCallerSession) INITIALWINDOWIDBUFFER() (*big.Int, error) {
	return _Token.Contract.INITIALWINDOWIDBUFFER(&_Token.CallOpts)
}

// RAY is a free data retrieval call binding the contract method 0x552033c4.
//
// Solidity: function RAY() view returns(uint256)
func (_Token *TokenCaller) RAY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "RAY")
	return *ret0, err
}

// RAY is a free data retrieval call binding the contract method 0x552033c4.
//
// Solidity: function RAY() view returns(uint256)
func (_Token *TokenSession) RAY() (*big.Int, error) {
	return _Token.Contract.RAY(&_Token.CallOpts)
}

// RAY is a free data retrieval call binding the contract method 0x552033c4.
//
// Solidity: function RAY() view returns(uint256)
func (_Token *TokenCallerSession) RAY() (*big.Int, error) {
	return _Token.Contract.RAY(&_Token.CallOpts)
}

// Augur is a free data retrieval call binding the contract method 0x7a0d8f8a.
//
// Solidity: function augur() view returns(address)
func (_Token *TokenCaller) Augur(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "augur")
	return *ret0, err
}

// Augur is a free data retrieval call binding the contract method 0x7a0d8f8a.
//
// Solidity: function augur() view returns(address)
func (_Token *TokenSession) Augur() (common.Address, error) {
	return _Token.Contract.Augur(&_Token.CallOpts)
}

// Augur is a free data retrieval call binding the contract method 0x7a0d8f8a.
//
// Solidity: function augur() view returns(address)
func (_Token *TokenCallerSession) Augur() (common.Address, error) {
	return _Token.Contract.Augur(&_Token.CallOpts)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_Token *TokenCaller) Cash(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "cash")
	return *ret0, err
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_Token *TokenSession) Cash() (common.Address, error) {
	return _Token.Contract.Cash(&_Token.CallOpts)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_Token *TokenCallerSession) Cash() (common.Address, error) {
	return _Token.Contract.Cash(&_Token.CallOpts)
}

// CreationTime is a free data retrieval call binding the contract method 0xd8270dce.
//
// Solidity: function creationTime() view returns(uint256)
func (_Token *TokenCaller) CreationTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "creationTime")
	return *ret0, err
}

// CreationTime is a free data retrieval call binding the contract method 0xd8270dce.
//
// Solidity: function creationTime() view returns(uint256)
func (_Token *TokenSession) CreationTime() (*big.Int, error) {
	return _Token.Contract.CreationTime(&_Token.CallOpts)
}

// CreationTime is a free data retrieval call binding the contract method 0xd8270dce.
//
// Solidity: function creationTime() view returns(uint256)
func (_Token *TokenCallerSession) CreationTime() (*big.Int, error) {
	return _Token.Contract.CreationTime(&_Token.CallOpts)
}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_Token *TokenCaller) DaiJoin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "daiJoin")
	return *ret0, err
}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_Token *TokenSession) DaiJoin() (common.Address, error) {
	return _Token.Contract.DaiJoin(&_Token.CallOpts)
}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_Token *TokenCallerSession) DaiJoin() (common.Address, error) {
	return _Token.Contract.DaiJoin(&_Token.CallOpts)
}

// DaiPot is a free data retrieval call binding the contract method 0x2cb1fd20.
//
// Solidity: function daiPot() view returns(address)
func (_Token *TokenCaller) DaiPot(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "daiPot")
	return *ret0, err
}

// DaiPot is a free data retrieval call binding the contract method 0x2cb1fd20.
//
// Solidity: function daiPot() view returns(address)
func (_Token *TokenSession) DaiPot() (common.Address, error) {
	return _Token.Contract.DaiPot(&_Token.CallOpts)
}

// DaiPot is a free data retrieval call binding the contract method 0x2cb1fd20.
//
// Solidity: function daiPot() view returns(address)
func (_Token *TokenCallerSession) DaiPot() (common.Address, error) {
	return _Token.Contract.DaiPot(&_Token.CallOpts)
}

// DaiVat is a free data retrieval call binding the contract method 0x3bd831f4.
//
// Solidity: function daiVat() view returns(address)
func (_Token *TokenCaller) DaiVat(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "daiVat")
	return *ret0, err
}

// DaiVat is a free data retrieval call binding the contract method 0x3bd831f4.
//
// Solidity: function daiVat() view returns(address)
func (_Token *TokenSession) DaiVat() (common.Address, error) {
	return _Token.Contract.DaiVat(&_Token.CallOpts)
}

// DaiVat is a free data retrieval call binding the contract method 0x3bd831f4.
//
// Solidity: function daiVat() view returns(address)
func (_Token *TokenCallerSession) DaiVat() (common.Address, error) {
	return _Token.Contract.DaiVat(&_Token.CallOpts)
}

// DesignatedReportNoShowBondInAttoRep is a free data retrieval call binding the contract method 0xb53b321e.
//
// Solidity: function designatedReportNoShowBondInAttoRep(address ) view returns(uint256)
func (_Token *TokenCaller) DesignatedReportNoShowBondInAttoRep(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "designatedReportNoShowBondInAttoRep", arg0)
	return *ret0, err
}

// DesignatedReportNoShowBondInAttoRep is a free data retrieval call binding the contract method 0xb53b321e.
//
// Solidity: function designatedReportNoShowBondInAttoRep(address ) view returns(uint256)
func (_Token *TokenSession) DesignatedReportNoShowBondInAttoRep(arg0 common.Address) (*big.Int, error) {
	return _Token.Contract.DesignatedReportNoShowBondInAttoRep(&_Token.CallOpts, arg0)
}

// DesignatedReportNoShowBondInAttoRep is a free data retrieval call binding the contract method 0xb53b321e.
//
// Solidity: function designatedReportNoShowBondInAttoRep(address ) view returns(uint256)
func (_Token *TokenCallerSession) DesignatedReportNoShowBondInAttoRep(arg0 common.Address) (*big.Int, error) {
	return _Token.Contract.DesignatedReportNoShowBondInAttoRep(&_Token.CallOpts, arg0)
}

// DesignatedReportStakeInAttoRep is a free data retrieval call binding the contract method 0x3ec6b034.
//
// Solidity: function designatedReportStakeInAttoRep(address ) view returns(uint256)
func (_Token *TokenCaller) DesignatedReportStakeInAttoRep(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "designatedReportStakeInAttoRep", arg0)
	return *ret0, err
}

// DesignatedReportStakeInAttoRep is a free data retrieval call binding the contract method 0x3ec6b034.
//
// Solidity: function designatedReportStakeInAttoRep(address ) view returns(uint256)
func (_Token *TokenSession) DesignatedReportStakeInAttoRep(arg0 common.Address) (*big.Int, error) {
	return _Token.Contract.DesignatedReportStakeInAttoRep(&_Token.CallOpts, arg0)
}

// DesignatedReportStakeInAttoRep is a free data retrieval call binding the contract method 0x3ec6b034.
//
// Solidity: function designatedReportStakeInAttoRep(address ) view returns(uint256)
func (_Token *TokenCallerSession) DesignatedReportStakeInAttoRep(arg0 common.Address) (*big.Int, error) {
	return _Token.Contract.DesignatedReportStakeInAttoRep(&_Token.CallOpts, arg0)
}

// DisputeWindowFactory is a free data retrieval call binding the contract method 0xa9f4185d.
//
// Solidity: function disputeWindowFactory() view returns(address)
func (_Token *TokenCaller) DisputeWindowFactory(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "disputeWindowFactory")
	return *ret0, err
}

// DisputeWindowFactory is a free data retrieval call binding the contract method 0xa9f4185d.
//
// Solidity: function disputeWindowFactory() view returns(address)
func (_Token *TokenSession) DisputeWindowFactory() (common.Address, error) {
	return _Token.Contract.DisputeWindowFactory(&_Token.CallOpts)
}

// DisputeWindowFactory is a free data retrieval call binding the contract method 0xa9f4185d.
//
// Solidity: function disputeWindowFactory() view returns(address)
func (_Token *TokenCallerSession) DisputeWindowFactory() (common.Address, error) {
	return _Token.Contract.DisputeWindowFactory(&_Token.CallOpts)
}

// DisputeWindows is a free data retrieval call binding the contract method 0x57d7a3c4.
//
// Solidity: function disputeWindows(uint256 ) view returns(address)
func (_Token *TokenCaller) DisputeWindows(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "disputeWindows", arg0)
	return *ret0, err
}

// DisputeWindows is a free data retrieval call binding the contract method 0x57d7a3c4.
//
// Solidity: function disputeWindows(uint256 ) view returns(address)
func (_Token *TokenSession) DisputeWindows(arg0 *big.Int) (common.Address, error) {
	return _Token.Contract.DisputeWindows(&_Token.CallOpts, arg0)
}

// DisputeWindows is a free data retrieval call binding the contract method 0x57d7a3c4.
//
// Solidity: function disputeWindows(uint256 ) view returns(address)
func (_Token *TokenCallerSession) DisputeWindows(arg0 *big.Int) (common.Address, error) {
	return _Token.Contract.DisputeWindows(&_Token.CallOpts, arg0)
}

// Formulas is a free data retrieval call binding the contract method 0x7f7c390a.
//
// Solidity: function formulas() view returns(address)
func (_Token *TokenCaller) Formulas(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "formulas")
	return *ret0, err
}

// Formulas is a free data retrieval call binding the contract method 0x7f7c390a.
//
// Solidity: function formulas() view returns(address)
func (_Token *TokenSession) Formulas() (common.Address, error) {
	return _Token.Contract.Formulas(&_Token.CallOpts)
}

// Formulas is a free data retrieval call binding the contract method 0x7f7c390a.
//
// Solidity: function formulas() view returns(address)
func (_Token *TokenCallerSession) Formulas() (common.Address, error) {
	return _Token.Contract.Formulas(&_Token.CallOpts)
}

// GetChildUniverse is a free data retrieval call binding the contract method 0xeceba876.
//
// Solidity: function getChildUniverse(bytes32 _parentPayoutDistributionHash) view returns(address)
func (_Token *TokenCaller) GetChildUniverse(opts *bind.CallOpts, _parentPayoutDistributionHash [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getChildUniverse", _parentPayoutDistributionHash)
	return *ret0, err
}

// GetChildUniverse is a free data retrieval call binding the contract method 0xeceba876.
//
// Solidity: function getChildUniverse(bytes32 _parentPayoutDistributionHash) view returns(address)
func (_Token *TokenSession) GetChildUniverse(_parentPayoutDistributionHash [32]byte) (common.Address, error) {
	return _Token.Contract.GetChildUniverse(&_Token.CallOpts, _parentPayoutDistributionHash)
}

// GetChildUniverse is a free data retrieval call binding the contract method 0xeceba876.
//
// Solidity: function getChildUniverse(bytes32 _parentPayoutDistributionHash) view returns(address)
func (_Token *TokenCallerSession) GetChildUniverse(_parentPayoutDistributionHash [32]byte) (common.Address, error) {
	return _Token.Contract.GetChildUniverse(&_Token.CallOpts, _parentPayoutDistributionHash)
}

// GetCurrentDisputeWindow is a free data retrieval call binding the contract method 0x8699d434.
//
// Solidity: function getCurrentDisputeWindow(bool _initial) view returns(address)
func (_Token *TokenCaller) GetCurrentDisputeWindow(opts *bind.CallOpts, _initial bool) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getCurrentDisputeWindow", _initial)
	return *ret0, err
}

// GetCurrentDisputeWindow is a free data retrieval call binding the contract method 0x8699d434.
//
// Solidity: function getCurrentDisputeWindow(bool _initial) view returns(address)
func (_Token *TokenSession) GetCurrentDisputeWindow(_initial bool) (common.Address, error) {
	return _Token.Contract.GetCurrentDisputeWindow(&_Token.CallOpts, _initial)
}

// GetCurrentDisputeWindow is a free data retrieval call binding the contract method 0x8699d434.
//
// Solidity: function getCurrentDisputeWindow(bool _initial) view returns(address)
func (_Token *TokenCallerSession) GetCurrentDisputeWindow(_initial bool) (common.Address, error) {
	return _Token.Contract.GetCurrentDisputeWindow(&_Token.CallOpts, _initial)
}

// GetDisputeRoundDurationInSeconds is a free data retrieval call binding the contract method 0x879eefa5.
//
// Solidity: function getDisputeRoundDurationInSeconds(bool _initial) view returns(uint256)
func (_Token *TokenCaller) GetDisputeRoundDurationInSeconds(opts *bind.CallOpts, _initial bool) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getDisputeRoundDurationInSeconds", _initial)
	return *ret0, err
}

// GetDisputeRoundDurationInSeconds is a free data retrieval call binding the contract method 0x879eefa5.
//
// Solidity: function getDisputeRoundDurationInSeconds(bool _initial) view returns(uint256)
func (_Token *TokenSession) GetDisputeRoundDurationInSeconds(_initial bool) (*big.Int, error) {
	return _Token.Contract.GetDisputeRoundDurationInSeconds(&_Token.CallOpts, _initial)
}

// GetDisputeRoundDurationInSeconds is a free data retrieval call binding the contract method 0x879eefa5.
//
// Solidity: function getDisputeRoundDurationInSeconds(bool _initial) view returns(uint256)
func (_Token *TokenCallerSession) GetDisputeRoundDurationInSeconds(_initial bool) (*big.Int, error) {
	return _Token.Contract.GetDisputeRoundDurationInSeconds(&_Token.CallOpts, _initial)
}

// GetDisputeThresholdForDisputePacing is a free data retrieval call binding the contract method 0x047825c7.
//
// Solidity: function getDisputeThresholdForDisputePacing() view returns(uint256)
func (_Token *TokenCaller) GetDisputeThresholdForDisputePacing(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getDisputeThresholdForDisputePacing")
	return *ret0, err
}

// GetDisputeThresholdForDisputePacing is a free data retrieval call binding the contract method 0x047825c7.
//
// Solidity: function getDisputeThresholdForDisputePacing() view returns(uint256)
func (_Token *TokenSession) GetDisputeThresholdForDisputePacing() (*big.Int, error) {
	return _Token.Contract.GetDisputeThresholdForDisputePacing(&_Token.CallOpts)
}

// GetDisputeThresholdForDisputePacing is a free data retrieval call binding the contract method 0x047825c7.
//
// Solidity: function getDisputeThresholdForDisputePacing() view returns(uint256)
func (_Token *TokenCallerSession) GetDisputeThresholdForDisputePacing() (*big.Int, error) {
	return _Token.Contract.GetDisputeThresholdForDisputePacing(&_Token.CallOpts)
}

// GetDisputeThresholdForFork is a free data retrieval call binding the contract method 0xfb03eaea.
//
// Solidity: function getDisputeThresholdForFork() view returns(uint256)
func (_Token *TokenCaller) GetDisputeThresholdForFork(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getDisputeThresholdForFork")
	return *ret0, err
}

// GetDisputeThresholdForFork is a free data retrieval call binding the contract method 0xfb03eaea.
//
// Solidity: function getDisputeThresholdForFork() view returns(uint256)
func (_Token *TokenSession) GetDisputeThresholdForFork() (*big.Int, error) {
	return _Token.Contract.GetDisputeThresholdForFork(&_Token.CallOpts)
}

// GetDisputeThresholdForFork is a free data retrieval call binding the contract method 0xfb03eaea.
//
// Solidity: function getDisputeThresholdForFork() view returns(uint256)
func (_Token *TokenCallerSession) GetDisputeThresholdForFork() (*big.Int, error) {
	return _Token.Contract.GetDisputeThresholdForFork(&_Token.CallOpts)
}

// GetDisputeWindow is a free data retrieval call binding the contract method 0x6c23f723.
//
// Solidity: function getDisputeWindow(uint256 _disputeWindowId) view returns(address)
func (_Token *TokenCaller) GetDisputeWindow(opts *bind.CallOpts, _disputeWindowId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getDisputeWindow", _disputeWindowId)
	return *ret0, err
}

// GetDisputeWindow is a free data retrieval call binding the contract method 0x6c23f723.
//
// Solidity: function getDisputeWindow(uint256 _disputeWindowId) view returns(address)
func (_Token *TokenSession) GetDisputeWindow(_disputeWindowId *big.Int) (common.Address, error) {
	return _Token.Contract.GetDisputeWindow(&_Token.CallOpts, _disputeWindowId)
}

// GetDisputeWindow is a free data retrieval call binding the contract method 0x6c23f723.
//
// Solidity: function getDisputeWindow(uint256 _disputeWindowId) view returns(address)
func (_Token *TokenCallerSession) GetDisputeWindow(_disputeWindowId *big.Int) (common.Address, error) {
	return _Token.Contract.GetDisputeWindow(&_Token.CallOpts, _disputeWindowId)
}

// GetDisputeWindowByTimestamp is a free data retrieval call binding the contract method 0x622ae175.
//
// Solidity: function getDisputeWindowByTimestamp(uint256 _timestamp, bool _initial) view returns(address)
func (_Token *TokenCaller) GetDisputeWindowByTimestamp(opts *bind.CallOpts, _timestamp *big.Int, _initial bool) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getDisputeWindowByTimestamp", _timestamp, _initial)
	return *ret0, err
}

// GetDisputeWindowByTimestamp is a free data retrieval call binding the contract method 0x622ae175.
//
// Solidity: function getDisputeWindowByTimestamp(uint256 _timestamp, bool _initial) view returns(address)
func (_Token *TokenSession) GetDisputeWindowByTimestamp(_timestamp *big.Int, _initial bool) (common.Address, error) {
	return _Token.Contract.GetDisputeWindowByTimestamp(&_Token.CallOpts, _timestamp, _initial)
}

// GetDisputeWindowByTimestamp is a free data retrieval call binding the contract method 0x622ae175.
//
// Solidity: function getDisputeWindowByTimestamp(uint256 _timestamp, bool _initial) view returns(address)
func (_Token *TokenCallerSession) GetDisputeWindowByTimestamp(_timestamp *big.Int, _initial bool) (common.Address, error) {
	return _Token.Contract.GetDisputeWindowByTimestamp(&_Token.CallOpts, _timestamp, _initial)
}

// GetDisputeWindowId is a free data retrieval call binding the contract method 0x2c7faa4d.
//
// Solidity: function getDisputeWindowId(uint256 _timestamp, bool _initial) view returns(uint256)
func (_Token *TokenCaller) GetDisputeWindowId(opts *bind.CallOpts, _timestamp *big.Int, _initial bool) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getDisputeWindowId", _timestamp, _initial)
	return *ret0, err
}

// GetDisputeWindowId is a free data retrieval call binding the contract method 0x2c7faa4d.
//
// Solidity: function getDisputeWindowId(uint256 _timestamp, bool _initial) view returns(uint256)
func (_Token *TokenSession) GetDisputeWindowId(_timestamp *big.Int, _initial bool) (*big.Int, error) {
	return _Token.Contract.GetDisputeWindowId(&_Token.CallOpts, _timestamp, _initial)
}

// GetDisputeWindowId is a free data retrieval call binding the contract method 0x2c7faa4d.
//
// Solidity: function getDisputeWindowId(uint256 _timestamp, bool _initial) view returns(uint256)
func (_Token *TokenCallerSession) GetDisputeWindowId(_timestamp *big.Int, _initial bool) (*big.Int, error) {
	return _Token.Contract.GetDisputeWindowId(&_Token.CallOpts, _timestamp, _initial)
}

// GetDisputeWindowStartTimeAndDuration is a free data retrieval call binding the contract method 0x5449aed5.
//
// Solidity: function getDisputeWindowStartTimeAndDuration(uint256 _timestamp, bool _initial) view returns(uint256 _startTime, uint256 _duration)
func (_Token *TokenCaller) GetDisputeWindowStartTimeAndDuration(opts *bind.CallOpts, _timestamp *big.Int, _initial bool) (struct {
	StartTime *big.Int
	Duration  *big.Int
}, error) {
	ret := new(struct {
		StartTime *big.Int
		Duration  *big.Int
	})
	out := ret
	err := _Token.contract.Call(opts, out, "getDisputeWindowStartTimeAndDuration", _timestamp, _initial)
	return *ret, err
}

// GetDisputeWindowStartTimeAndDuration is a free data retrieval call binding the contract method 0x5449aed5.
//
// Solidity: function getDisputeWindowStartTimeAndDuration(uint256 _timestamp, bool _initial) view returns(uint256 _startTime, uint256 _duration)
func (_Token *TokenSession) GetDisputeWindowStartTimeAndDuration(_timestamp *big.Int, _initial bool) (struct {
	StartTime *big.Int
	Duration  *big.Int
}, error) {
	return _Token.Contract.GetDisputeWindowStartTimeAndDuration(&_Token.CallOpts, _timestamp, _initial)
}

// GetDisputeWindowStartTimeAndDuration is a free data retrieval call binding the contract method 0x5449aed5.
//
// Solidity: function getDisputeWindowStartTimeAndDuration(uint256 _timestamp, bool _initial) view returns(uint256 _startTime, uint256 _duration)
func (_Token *TokenCallerSession) GetDisputeWindowStartTimeAndDuration(_timestamp *big.Int, _initial bool) (struct {
	StartTime *big.Int
	Duration  *big.Int
}, error) {
	return _Token.Contract.GetDisputeWindowStartTimeAndDuration(&_Token.CallOpts, _timestamp, _initial)
}

// GetForkEndTime is a free data retrieval call binding the contract method 0x77e71ee5.
//
// Solidity: function getForkEndTime() view returns(uint256)
func (_Token *TokenCaller) GetForkEndTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getForkEndTime")
	return *ret0, err
}

// GetForkEndTime is a free data retrieval call binding the contract method 0x77e71ee5.
//
// Solidity: function getForkEndTime() view returns(uint256)
func (_Token *TokenSession) GetForkEndTime() (*big.Int, error) {
	return _Token.Contract.GetForkEndTime(&_Token.CallOpts)
}

// GetForkEndTime is a free data retrieval call binding the contract method 0x77e71ee5.
//
// Solidity: function getForkEndTime() view returns(uint256)
func (_Token *TokenCallerSession) GetForkEndTime() (*big.Int, error) {
	return _Token.Contract.GetForkEndTime(&_Token.CallOpts)
}

// GetForkReputationGoal is a free data retrieval call binding the contract method 0x7c377d74.
//
// Solidity: function getForkReputationGoal() view returns(uint256)
func (_Token *TokenCaller) GetForkReputationGoal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getForkReputationGoal")
	return *ret0, err
}

// GetForkReputationGoal is a free data retrieval call binding the contract method 0x7c377d74.
//
// Solidity: function getForkReputationGoal() view returns(uint256)
func (_Token *TokenSession) GetForkReputationGoal() (*big.Int, error) {
	return _Token.Contract.GetForkReputationGoal(&_Token.CallOpts)
}

// GetForkReputationGoal is a free data retrieval call binding the contract method 0x7c377d74.
//
// Solidity: function getForkReputationGoal() view returns(uint256)
func (_Token *TokenCallerSession) GetForkReputationGoal() (*big.Int, error) {
	return _Token.Contract.GetForkReputationGoal(&_Token.CallOpts)
}

// GetForkingMarket is a free data retrieval call binding the contract method 0xcb1d8418.
//
// Solidity: function getForkingMarket() view returns(address)
func (_Token *TokenCaller) GetForkingMarket(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getForkingMarket")
	return *ret0, err
}

// GetForkingMarket is a free data retrieval call binding the contract method 0xcb1d8418.
//
// Solidity: function getForkingMarket() view returns(address)
func (_Token *TokenSession) GetForkingMarket() (common.Address, error) {
	return _Token.Contract.GetForkingMarket(&_Token.CallOpts)
}

// GetForkingMarket is a free data retrieval call binding the contract method 0xcb1d8418.
//
// Solidity: function getForkingMarket() view returns(address)
func (_Token *TokenCallerSession) GetForkingMarket() (common.Address, error) {
	return _Token.Contract.GetForkingMarket(&_Token.CallOpts)
}

// GetInitialReportMinValue is a free data retrieval call binding the contract method 0x5f723b50.
//
// Solidity: function getInitialReportMinValue() view returns(uint256)
func (_Token *TokenCaller) GetInitialReportMinValue(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getInitialReportMinValue")
	return *ret0, err
}

// GetInitialReportMinValue is a free data retrieval call binding the contract method 0x5f723b50.
//
// Solidity: function getInitialReportMinValue() view returns(uint256)
func (_Token *TokenSession) GetInitialReportMinValue() (*big.Int, error) {
	return _Token.Contract.GetInitialReportMinValue(&_Token.CallOpts)
}

// GetInitialReportMinValue is a free data retrieval call binding the contract method 0x5f723b50.
//
// Solidity: function getInitialReportMinValue() view returns(uint256)
func (_Token *TokenCallerSession) GetInitialReportMinValue() (*big.Int, error) {
	return _Token.Contract.GetInitialReportMinValue(&_Token.CallOpts)
}

// GetOpenInterestInAttoCash is a free data retrieval call binding the contract method 0xc675f222.
//
// Solidity: function getOpenInterestInAttoCash() view returns(uint256)
func (_Token *TokenCaller) GetOpenInterestInAttoCash(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getOpenInterestInAttoCash")
	return *ret0, err
}

// GetOpenInterestInAttoCash is a free data retrieval call binding the contract method 0xc675f222.
//
// Solidity: function getOpenInterestInAttoCash() view returns(uint256)
func (_Token *TokenSession) GetOpenInterestInAttoCash() (*big.Int, error) {
	return _Token.Contract.GetOpenInterestInAttoCash(&_Token.CallOpts)
}

// GetOpenInterestInAttoCash is a free data retrieval call binding the contract method 0xc675f222.
//
// Solidity: function getOpenInterestInAttoCash() view returns(uint256)
func (_Token *TokenCallerSession) GetOpenInterestInAttoCash() (*big.Int, error) {
	return _Token.Contract.GetOpenInterestInAttoCash(&_Token.CallOpts)
}

// GetParentPayoutDistributionHash is a free data retrieval call binding the contract method 0xc38c0fa7.
//
// Solidity: function getParentPayoutDistributionHash() view returns(bytes32)
func (_Token *TokenCaller) GetParentPayoutDistributionHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getParentPayoutDistributionHash")
	return *ret0, err
}

// GetParentPayoutDistributionHash is a free data retrieval call binding the contract method 0xc38c0fa7.
//
// Solidity: function getParentPayoutDistributionHash() view returns(bytes32)
func (_Token *TokenSession) GetParentPayoutDistributionHash() ([32]byte, error) {
	return _Token.Contract.GetParentPayoutDistributionHash(&_Token.CallOpts)
}

// GetParentPayoutDistributionHash is a free data retrieval call binding the contract method 0xc38c0fa7.
//
// Solidity: function getParentPayoutDistributionHash() view returns(bytes32)
func (_Token *TokenCallerSession) GetParentPayoutDistributionHash() ([32]byte, error) {
	return _Token.Contract.GetParentPayoutDistributionHash(&_Token.CallOpts)
}

// GetParentUniverse is a free data retrieval call binding the contract method 0xa63f1350.
//
// Solidity: function getParentUniverse() view returns(address)
func (_Token *TokenCaller) GetParentUniverse(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getParentUniverse")
	return *ret0, err
}

// GetParentUniverse is a free data retrieval call binding the contract method 0xa63f1350.
//
// Solidity: function getParentUniverse() view returns(address)
func (_Token *TokenSession) GetParentUniverse() (common.Address, error) {
	return _Token.Contract.GetParentUniverse(&_Token.CallOpts)
}

// GetParentUniverse is a free data retrieval call binding the contract method 0xa63f1350.
//
// Solidity: function getParentUniverse() view returns(address)
func (_Token *TokenCallerSession) GetParentUniverse() (common.Address, error) {
	return _Token.Contract.GetParentUniverse(&_Token.CallOpts)
}

// GetPayoutNumerator is a free data retrieval call binding the contract method 0xda834ac4.
//
// Solidity: function getPayoutNumerator(uint256 _outcome) view returns(uint256)
func (_Token *TokenCaller) GetPayoutNumerator(opts *bind.CallOpts, _outcome *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getPayoutNumerator", _outcome)
	return *ret0, err
}

// GetPayoutNumerator is a free data retrieval call binding the contract method 0xda834ac4.
//
// Solidity: function getPayoutNumerator(uint256 _outcome) view returns(uint256)
func (_Token *TokenSession) GetPayoutNumerator(_outcome *big.Int) (*big.Int, error) {
	return _Token.Contract.GetPayoutNumerator(&_Token.CallOpts, _outcome)
}

// GetPayoutNumerator is a free data retrieval call binding the contract method 0xda834ac4.
//
// Solidity: function getPayoutNumerator(uint256 _outcome) view returns(uint256)
func (_Token *TokenCallerSession) GetPayoutNumerator(_outcome *big.Int) (*big.Int, error) {
	return _Token.Contract.GetPayoutNumerator(&_Token.CallOpts, _outcome)
}

// GetPayoutNumerators is a free data retrieval call binding the contract method 0x6f84676e.
//
// Solidity: function getPayoutNumerators() view returns(uint256[])
func (_Token *TokenCaller) GetPayoutNumerators(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getPayoutNumerators")
	return *ret0, err
}

// GetPayoutNumerators is a free data retrieval call binding the contract method 0x6f84676e.
//
// Solidity: function getPayoutNumerators() view returns(uint256[])
func (_Token *TokenSession) GetPayoutNumerators() ([]*big.Int, error) {
	return _Token.Contract.GetPayoutNumerators(&_Token.CallOpts)
}

// GetPayoutNumerators is a free data retrieval call binding the contract method 0x6f84676e.
//
// Solidity: function getPayoutNumerators() view returns(uint256[])
func (_Token *TokenCallerSession) GetPayoutNumerators() ([]*big.Int, error) {
	return _Token.Contract.GetPayoutNumerators(&_Token.CallOpts)
}

// GetReportingFeeDivisor is a free data retrieval call binding the contract method 0x0dcde5f5.
//
// Solidity: function getReportingFeeDivisor() view returns(uint256)
func (_Token *TokenCaller) GetReportingFeeDivisor(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getReportingFeeDivisor")
	return *ret0, err
}

// GetReportingFeeDivisor is a free data retrieval call binding the contract method 0x0dcde5f5.
//
// Solidity: function getReportingFeeDivisor() view returns(uint256)
func (_Token *TokenSession) GetReportingFeeDivisor() (*big.Int, error) {
	return _Token.Contract.GetReportingFeeDivisor(&_Token.CallOpts)
}

// GetReportingFeeDivisor is a free data retrieval call binding the contract method 0x0dcde5f5.
//
// Solidity: function getReportingFeeDivisor() view returns(uint256)
func (_Token *TokenCallerSession) GetReportingFeeDivisor() (*big.Int, error) {
	return _Token.Contract.GetReportingFeeDivisor(&_Token.CallOpts)
}

// GetReputationToken is a free data retrieval call binding the contract method 0xb80907f2.
//
// Solidity: function getReputationToken() view returns(address)
func (_Token *TokenCaller) GetReputationToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getReputationToken")
	return *ret0, err
}

// GetReputationToken is a free data retrieval call binding the contract method 0xb80907f2.
//
// Solidity: function getReputationToken() view returns(address)
func (_Token *TokenSession) GetReputationToken() (common.Address, error) {
	return _Token.Contract.GetReputationToken(&_Token.CallOpts)
}

// GetReputationToken is a free data retrieval call binding the contract method 0xb80907f2.
//
// Solidity: function getReputationToken() view returns(address)
func (_Token *TokenCallerSession) GetReputationToken() (common.Address, error) {
	return _Token.Contract.GetReputationToken(&_Token.CallOpts)
}

// GetTargetRepMarketCapInAttoCash is a free data retrieval call binding the contract method 0xdf9fde7e.
//
// Solidity: function getTargetRepMarketCapInAttoCash() view returns(uint256)
func (_Token *TokenCaller) GetTargetRepMarketCapInAttoCash(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getTargetRepMarketCapInAttoCash")
	return *ret0, err
}

// GetTargetRepMarketCapInAttoCash is a free data retrieval call binding the contract method 0xdf9fde7e.
//
// Solidity: function getTargetRepMarketCapInAttoCash() view returns(uint256)
func (_Token *TokenSession) GetTargetRepMarketCapInAttoCash() (*big.Int, error) {
	return _Token.Contract.GetTargetRepMarketCapInAttoCash(&_Token.CallOpts)
}

// GetTargetRepMarketCapInAttoCash is a free data retrieval call binding the contract method 0xdf9fde7e.
//
// Solidity: function getTargetRepMarketCapInAttoCash() view returns(uint256)
func (_Token *TokenCallerSession) GetTargetRepMarketCapInAttoCash() (*big.Int, error) {
	return _Token.Contract.GetTargetRepMarketCapInAttoCash(&_Token.CallOpts)
}

// GetWinningChildPayoutNumerator is a free data retrieval call binding the contract method 0x7262f993.
//
// Solidity: function getWinningChildPayoutNumerator(uint256 _outcome) view returns(uint256)
func (_Token *TokenCaller) GetWinningChildPayoutNumerator(opts *bind.CallOpts, _outcome *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getWinningChildPayoutNumerator", _outcome)
	return *ret0, err
}

// GetWinningChildPayoutNumerator is a free data retrieval call binding the contract method 0x7262f993.
//
// Solidity: function getWinningChildPayoutNumerator(uint256 _outcome) view returns(uint256)
func (_Token *TokenSession) GetWinningChildPayoutNumerator(_outcome *big.Int) (*big.Int, error) {
	return _Token.Contract.GetWinningChildPayoutNumerator(&_Token.CallOpts, _outcome)
}

// GetWinningChildPayoutNumerator is a free data retrieval call binding the contract method 0x7262f993.
//
// Solidity: function getWinningChildPayoutNumerator(uint256 _outcome) view returns(uint256)
func (_Token *TokenCallerSession) GetWinningChildPayoutNumerator(_outcome *big.Int) (*big.Int, error) {
	return _Token.Contract.GetWinningChildPayoutNumerator(&_Token.CallOpts, _outcome)
}

// GetWinningChildUniverse is a free data retrieval call binding the contract method 0x6f70b9cb.
//
// Solidity: function getWinningChildUniverse() view returns(address)
func (_Token *TokenCaller) GetWinningChildUniverse(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getWinningChildUniverse")
	return *ret0, err
}

// GetWinningChildUniverse is a free data retrieval call binding the contract method 0x6f70b9cb.
//
// Solidity: function getWinningChildUniverse() view returns(address)
func (_Token *TokenSession) GetWinningChildUniverse() (common.Address, error) {
	return _Token.Contract.GetWinningChildUniverse(&_Token.CallOpts)
}

// GetWinningChildUniverse is a free data retrieval call binding the contract method 0x6f70b9cb.
//
// Solidity: function getWinningChildUniverse() view returns(address)
func (_Token *TokenCallerSession) GetWinningChildUniverse() (common.Address, error) {
	return _Token.Contract.GetWinningChildUniverse(&_Token.CallOpts)
}

// IsContainerForDisputeWindow is a free data retrieval call binding the contract method 0x01ba1fa3.
//
// Solidity: function isContainerForDisputeWindow(address _shadyDisputeWindow) view returns(bool)
func (_Token *TokenCaller) IsContainerForDisputeWindow(opts *bind.CallOpts, _shadyDisputeWindow common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "isContainerForDisputeWindow", _shadyDisputeWindow)
	return *ret0, err
}

// IsContainerForDisputeWindow is a free data retrieval call binding the contract method 0x01ba1fa3.
//
// Solidity: function isContainerForDisputeWindow(address _shadyDisputeWindow) view returns(bool)
func (_Token *TokenSession) IsContainerForDisputeWindow(_shadyDisputeWindow common.Address) (bool, error) {
	return _Token.Contract.IsContainerForDisputeWindow(&_Token.CallOpts, _shadyDisputeWindow)
}

// IsContainerForDisputeWindow is a free data retrieval call binding the contract method 0x01ba1fa3.
//
// Solidity: function isContainerForDisputeWindow(address _shadyDisputeWindow) view returns(bool)
func (_Token *TokenCallerSession) IsContainerForDisputeWindow(_shadyDisputeWindow common.Address) (bool, error) {
	return _Token.Contract.IsContainerForDisputeWindow(&_Token.CallOpts, _shadyDisputeWindow)
}

// IsContainerForMarket is a free data retrieval call binding the contract method 0x9f7e1bf6.
//
// Solidity: function isContainerForMarket(address _shadyMarket) view returns(bool)
func (_Token *TokenCaller) IsContainerForMarket(opts *bind.CallOpts, _shadyMarket common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "isContainerForMarket", _shadyMarket)
	return *ret0, err
}

// IsContainerForMarket is a free data retrieval call binding the contract method 0x9f7e1bf6.
//
// Solidity: function isContainerForMarket(address _shadyMarket) view returns(bool)
func (_Token *TokenSession) IsContainerForMarket(_shadyMarket common.Address) (bool, error) {
	return _Token.Contract.IsContainerForMarket(&_Token.CallOpts, _shadyMarket)
}

// IsContainerForMarket is a free data retrieval call binding the contract method 0x9f7e1bf6.
//
// Solidity: function isContainerForMarket(address _shadyMarket) view returns(bool)
func (_Token *TokenCallerSession) IsContainerForMarket(_shadyMarket common.Address) (bool, error) {
	return _Token.Contract.IsContainerForMarket(&_Token.CallOpts, _shadyMarket)
}

// IsContainerForReportingParticipant is a free data retrieval call binding the contract method 0xf76514c7.
//
// Solidity: function isContainerForReportingParticipant(address _shadyReportingParticipant) view returns(bool)
func (_Token *TokenCaller) IsContainerForReportingParticipant(opts *bind.CallOpts, _shadyReportingParticipant common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "isContainerForReportingParticipant", _shadyReportingParticipant)
	return *ret0, err
}

// IsContainerForReportingParticipant is a free data retrieval call binding the contract method 0xf76514c7.
//
// Solidity: function isContainerForReportingParticipant(address _shadyReportingParticipant) view returns(bool)
func (_Token *TokenSession) IsContainerForReportingParticipant(_shadyReportingParticipant common.Address) (bool, error) {
	return _Token.Contract.IsContainerForReportingParticipant(&_Token.CallOpts, _shadyReportingParticipant)
}

// IsContainerForReportingParticipant is a free data retrieval call binding the contract method 0xf76514c7.
//
// Solidity: function isContainerForReportingParticipant(address _shadyReportingParticipant) view returns(bool)
func (_Token *TokenCallerSession) IsContainerForReportingParticipant(_shadyReportingParticipant common.Address) (bool, error) {
	return _Token.Contract.IsContainerForReportingParticipant(&_Token.CallOpts, _shadyReportingParticipant)
}

// IsForking is a free data retrieval call binding the contract method 0xbecb1f35.
//
// Solidity: function isForking() view returns(bool)
func (_Token *TokenCaller) IsForking(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "isForking")
	return *ret0, err
}

// IsForking is a free data retrieval call binding the contract method 0xbecb1f35.
//
// Solidity: function isForking() view returns(bool)
func (_Token *TokenSession) IsForking() (bool, error) {
	return _Token.Contract.IsForking(&_Token.CallOpts)
}

// IsForking is a free data retrieval call binding the contract method 0xbecb1f35.
//
// Solidity: function isForking() view returns(bool)
func (_Token *TokenCallerSession) IsForking() (bool, error) {
	return _Token.Contract.IsForking(&_Token.CallOpts)
}

// IsForkingMarket is a free data retrieval call binding the contract method 0xd372fbcd.
//
// Solidity: function isForkingMarket() view returns(bool)
func (_Token *TokenCaller) IsForkingMarket(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "isForkingMarket")
	return *ret0, err
}

// IsForkingMarket is a free data retrieval call binding the contract method 0xd372fbcd.
//
// Solidity: function isForkingMarket() view returns(bool)
func (_Token *TokenSession) IsForkingMarket() (bool, error) {
	return _Token.Contract.IsForkingMarket(&_Token.CallOpts)
}

// IsForkingMarket is a free data retrieval call binding the contract method 0xd372fbcd.
//
// Solidity: function isForkingMarket() view returns(bool)
func (_Token *TokenCallerSession) IsForkingMarket() (bool, error) {
	return _Token.Contract.IsForkingMarket(&_Token.CallOpts)
}

// IsOpenInterestCash is a free data retrieval call binding the contract method 0x47d20e3b.
//
// Solidity: function isOpenInterestCash(address _address) view returns(bool)
func (_Token *TokenCaller) IsOpenInterestCash(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "isOpenInterestCash", _address)
	return *ret0, err
}

// IsOpenInterestCash is a free data retrieval call binding the contract method 0x47d20e3b.
//
// Solidity: function isOpenInterestCash(address _address) view returns(bool)
func (_Token *TokenSession) IsOpenInterestCash(_address common.Address) (bool, error) {
	return _Token.Contract.IsOpenInterestCash(&_Token.CallOpts, _address)
}

// IsOpenInterestCash is a free data retrieval call binding the contract method 0x47d20e3b.
//
// Solidity: function isOpenInterestCash(address _address) view returns(bool)
func (_Token *TokenCallerSession) IsOpenInterestCash(_address common.Address) (bool, error) {
	return _Token.Contract.IsOpenInterestCash(&_Token.CallOpts, _address)
}

// IsParentOf is a free data retrieval call binding the contract method 0x9517317c.
//
// Solidity: function isParentOf(address _shadyChild) view returns(bool)
func (_Token *TokenCaller) IsParentOf(opts *bind.CallOpts, _shadyChild common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "isParentOf", _shadyChild)
	return *ret0, err
}

// IsParentOf is a free data retrieval call binding the contract method 0x9517317c.
//
// Solidity: function isParentOf(address _shadyChild) view returns(bool)
func (_Token *TokenSession) IsParentOf(_shadyChild common.Address) (bool, error) {
	return _Token.Contract.IsParentOf(&_Token.CallOpts, _shadyChild)
}

// IsParentOf is a free data retrieval call binding the contract method 0x9517317c.
//
// Solidity: function isParentOf(address _shadyChild) view returns(bool)
func (_Token *TokenCallerSession) IsParentOf(_shadyChild common.Address) (bool, error) {
	return _Token.Contract.IsParentOf(&_Token.CallOpts, _shadyChild)
}

// LastSweep is a free data retrieval call binding the contract method 0xe5b4f771.
//
// Solidity: function lastSweep() view returns(uint256)
func (_Token *TokenCaller) LastSweep(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "lastSweep")
	return *ret0, err
}

// LastSweep is a free data retrieval call binding the contract method 0xe5b4f771.
//
// Solidity: function lastSweep() view returns(uint256)
func (_Token *TokenSession) LastSweep() (*big.Int, error) {
	return _Token.Contract.LastSweep(&_Token.CallOpts)
}

// LastSweep is a free data retrieval call binding the contract method 0xe5b4f771.
//
// Solidity: function lastSweep() view returns(uint256)
func (_Token *TokenCallerSession) LastSweep() (*big.Int, error) {
	return _Token.Contract.LastSweep(&_Token.CallOpts)
}

// MarketBalance is a free data retrieval call binding the contract method 0x9672e3ba.
//
// Solidity: function marketBalance(address ) view returns(uint256)
func (_Token *TokenCaller) MarketBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "marketBalance", arg0)
	return *ret0, err
}

// MarketBalance is a free data retrieval call binding the contract method 0x9672e3ba.
//
// Solidity: function marketBalance(address ) view returns(uint256)
func (_Token *TokenSession) MarketBalance(arg0 common.Address) (*big.Int, error) {
	return _Token.Contract.MarketBalance(&_Token.CallOpts, arg0)
}

// MarketBalance is a free data retrieval call binding the contract method 0x9672e3ba.
//
// Solidity: function marketBalance(address ) view returns(uint256)
func (_Token *TokenCallerSession) MarketBalance(arg0 common.Address) (*big.Int, error) {
	return _Token.Contract.MarketBalance(&_Token.CallOpts, arg0)
}

// MarketFactory is a free data retrieval call binding the contract method 0x06ae7095.
//
// Solidity: function marketFactory() view returns(address)
func (_Token *TokenCaller) MarketFactory(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "marketFactory")
	return *ret0, err
}

// MarketFactory is a free data retrieval call binding the contract method 0x06ae7095.
//
// Solidity: function marketFactory() view returns(address)
func (_Token *TokenSession) MarketFactory() (common.Address, error) {
	return _Token.Contract.MarketFactory(&_Token.CallOpts)
}

// MarketFactory is a free data retrieval call binding the contract method 0x06ae7095.
//
// Solidity: function marketFactory() view returns(address)
func (_Token *TokenCallerSession) MarketFactory() (common.Address, error) {
	return _Token.Contract.MarketFactory(&_Token.CallOpts)
}

// OpenInterestCash is a free data retrieval call binding the contract method 0x3940b675.
//
// Solidity: function openInterestCash() view returns(address)
func (_Token *TokenCaller) OpenInterestCash(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "openInterestCash")
	return *ret0, err
}

// OpenInterestCash is a free data retrieval call binding the contract method 0x3940b675.
//
// Solidity: function openInterestCash() view returns(address)
func (_Token *TokenSession) OpenInterestCash() (common.Address, error) {
	return _Token.Contract.OpenInterestCash(&_Token.CallOpts)
}

// OpenInterestCash is a free data retrieval call binding the contract method 0x3940b675.
//
// Solidity: function openInterestCash() view returns(address)
func (_Token *TokenCallerSession) OpenInterestCash() (common.Address, error) {
	return _Token.Contract.OpenInterestCash(&_Token.CallOpts)
}

// PayoutNumerators is a free data retrieval call binding the contract method 0x60d9b489.
//
// Solidity: function payoutNumerators(uint256 ) view returns(uint256)
func (_Token *TokenCaller) PayoutNumerators(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "payoutNumerators", arg0)
	return *ret0, err
}

// PayoutNumerators is a free data retrieval call binding the contract method 0x60d9b489.
//
// Solidity: function payoutNumerators(uint256 ) view returns(uint256)
func (_Token *TokenSession) PayoutNumerators(arg0 *big.Int) (*big.Int, error) {
	return _Token.Contract.PayoutNumerators(&_Token.CallOpts, arg0)
}

// PayoutNumerators is a free data retrieval call binding the contract method 0x60d9b489.
//
// Solidity: function payoutNumerators(uint256 ) view returns(uint256)
func (_Token *TokenCallerSession) PayoutNumerators(arg0 *big.Int) (*big.Int, error) {
	return _Token.Contract.PayoutNumerators(&_Token.CallOpts, arg0)
}

// PreviousDesignatedReportNoShowBondInAttoRep is a free data retrieval call binding the contract method 0x0bcde26d.
//
// Solidity: function previousDesignatedReportNoShowBondInAttoRep() view returns(uint256)
func (_Token *TokenCaller) PreviousDesignatedReportNoShowBondInAttoRep(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "previousDesignatedReportNoShowBondInAttoRep")
	return *ret0, err
}

// PreviousDesignatedReportNoShowBondInAttoRep is a free data retrieval call binding the contract method 0x0bcde26d.
//
// Solidity: function previousDesignatedReportNoShowBondInAttoRep() view returns(uint256)
func (_Token *TokenSession) PreviousDesignatedReportNoShowBondInAttoRep() (*big.Int, error) {
	return _Token.Contract.PreviousDesignatedReportNoShowBondInAttoRep(&_Token.CallOpts)
}

// PreviousDesignatedReportNoShowBondInAttoRep is a free data retrieval call binding the contract method 0x0bcde26d.
//
// Solidity: function previousDesignatedReportNoShowBondInAttoRep() view returns(uint256)
func (_Token *TokenCallerSession) PreviousDesignatedReportNoShowBondInAttoRep() (*big.Int, error) {
	return _Token.Contract.PreviousDesignatedReportNoShowBondInAttoRep(&_Token.CallOpts)
}

// PreviousDesignatedReportStakeInAttoRep is a free data retrieval call binding the contract method 0xb4eeb2d7.
//
// Solidity: function previousDesignatedReportStakeInAttoRep() view returns(uint256)
func (_Token *TokenCaller) PreviousDesignatedReportStakeInAttoRep(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "previousDesignatedReportStakeInAttoRep")
	return *ret0, err
}

// PreviousDesignatedReportStakeInAttoRep is a free data retrieval call binding the contract method 0xb4eeb2d7.
//
// Solidity: function previousDesignatedReportStakeInAttoRep() view returns(uint256)
func (_Token *TokenSession) PreviousDesignatedReportStakeInAttoRep() (*big.Int, error) {
	return _Token.Contract.PreviousDesignatedReportStakeInAttoRep(&_Token.CallOpts)
}

// PreviousDesignatedReportStakeInAttoRep is a free data retrieval call binding the contract method 0xb4eeb2d7.
//
// Solidity: function previousDesignatedReportStakeInAttoRep() view returns(uint256)
func (_Token *TokenCallerSession) PreviousDesignatedReportStakeInAttoRep() (*big.Int, error) {
	return _Token.Contract.PreviousDesignatedReportStakeInAttoRep(&_Token.CallOpts)
}

// PreviousReportingFeeDivisor is a free data retrieval call binding the contract method 0x6f33ebb7.
//
// Solidity: function previousReportingFeeDivisor() view returns(uint256)
func (_Token *TokenCaller) PreviousReportingFeeDivisor(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "previousReportingFeeDivisor")
	return *ret0, err
}

// PreviousReportingFeeDivisor is a free data retrieval call binding the contract method 0x6f33ebb7.
//
// Solidity: function previousReportingFeeDivisor() view returns(uint256)
func (_Token *TokenSession) PreviousReportingFeeDivisor() (*big.Int, error) {
	return _Token.Contract.PreviousReportingFeeDivisor(&_Token.CallOpts)
}

// PreviousReportingFeeDivisor is a free data retrieval call binding the contract method 0x6f33ebb7.
//
// Solidity: function previousReportingFeeDivisor() view returns(uint256)
func (_Token *TokenCallerSession) PreviousReportingFeeDivisor() (*big.Int, error) {
	return _Token.Contract.PreviousReportingFeeDivisor(&_Token.CallOpts)
}

// PreviousValidityBondInAttoCash is a free data retrieval call binding the contract method 0xef7cd234.
//
// Solidity: function previousValidityBondInAttoCash() view returns(uint256)
func (_Token *TokenCaller) PreviousValidityBondInAttoCash(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "previousValidityBondInAttoCash")
	return *ret0, err
}

// PreviousValidityBondInAttoCash is a free data retrieval call binding the contract method 0xef7cd234.
//
// Solidity: function previousValidityBondInAttoCash() view returns(uint256)
func (_Token *TokenSession) PreviousValidityBondInAttoCash() (*big.Int, error) {
	return _Token.Contract.PreviousValidityBondInAttoCash(&_Token.CallOpts)
}

// PreviousValidityBondInAttoCash is a free data retrieval call binding the contract method 0xef7cd234.
//
// Solidity: function previousValidityBondInAttoCash() view returns(uint256)
func (_Token *TokenCallerSession) PreviousValidityBondInAttoCash() (*big.Int, error) {
	return _Token.Contract.PreviousValidityBondInAttoCash(&_Token.CallOpts)
}

// RepOracle is a free data retrieval call binding the contract method 0x48bc4eba.
//
// Solidity: function repOracle() view returns(address)
func (_Token *TokenCaller) RepOracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "repOracle")
	return *ret0, err
}

// RepOracle is a free data retrieval call binding the contract method 0x48bc4eba.
//
// Solidity: function repOracle() view returns(address)
func (_Token *TokenSession) RepOracle() (common.Address, error) {
	return _Token.Contract.RepOracle(&_Token.CallOpts)
}

// RepOracle is a free data retrieval call binding the contract method 0x48bc4eba.
//
// Solidity: function repOracle() view returns(address)
func (_Token *TokenCallerSession) RepOracle() (common.Address, error) {
	return _Token.Contract.RepOracle(&_Token.CallOpts)
}

// ShareSettlementFeeDivisor is a free data retrieval call binding the contract method 0x0d5a9e4d.
//
// Solidity: function shareSettlementFeeDivisor(address ) view returns(uint256)
func (_Token *TokenCaller) ShareSettlementFeeDivisor(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "shareSettlementFeeDivisor", arg0)
	return *ret0, err
}

// ShareSettlementFeeDivisor is a free data retrieval call binding the contract method 0x0d5a9e4d.
//
// Solidity: function shareSettlementFeeDivisor(address ) view returns(uint256)
func (_Token *TokenSession) ShareSettlementFeeDivisor(arg0 common.Address) (*big.Int, error) {
	return _Token.Contract.ShareSettlementFeeDivisor(&_Token.CallOpts, arg0)
}

// ShareSettlementFeeDivisor is a free data retrieval call binding the contract method 0x0d5a9e4d.
//
// Solidity: function shareSettlementFeeDivisor(address ) view returns(uint256)
func (_Token *TokenCallerSession) ShareSettlementFeeDivisor(arg0 common.Address) (*big.Int, error) {
	return _Token.Contract.ShareSettlementFeeDivisor(&_Token.CallOpts, arg0)
}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() view returns(address)
func (_Token *TokenCaller) ShareToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "shareToken")
	return *ret0, err
}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() view returns(address)
func (_Token *TokenSession) ShareToken() (common.Address, error) {
	return _Token.Contract.ShareToken(&_Token.CallOpts)
}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() view returns(address)
func (_Token *TokenCallerSession) ShareToken() (common.Address, error) {
	return _Token.Contract.ShareToken(&_Token.CallOpts)
}

// TotalBalance is a free data retrieval call binding the contract method 0xad7a672f.
//
// Solidity: function totalBalance() view returns(uint256)
func (_Token *TokenCaller) TotalBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "totalBalance")
	return *ret0, err
}

// TotalBalance is a free data retrieval call binding the contract method 0xad7a672f.
//
// Solidity: function totalBalance() view returns(uint256)
func (_Token *TokenSession) TotalBalance() (*big.Int, error) {
	return _Token.Contract.TotalBalance(&_Token.CallOpts)
}

// TotalBalance is a free data retrieval call binding the contract method 0xad7a672f.
//
// Solidity: function totalBalance() view returns(uint256)
func (_Token *TokenCallerSession) TotalBalance() (*big.Int, error) {
	return _Token.Contract.TotalBalance(&_Token.CallOpts)
}

// ValidityBondInAttoCash is a free data retrieval call binding the contract method 0xefb106ba.
//
// Solidity: function validityBondInAttoCash(address ) view returns(uint256)
func (_Token *TokenCaller) ValidityBondInAttoCash(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "validityBondInAttoCash", arg0)
	return *ret0, err
}

// ValidityBondInAttoCash is a free data retrieval call binding the contract method 0xefb106ba.
//
// Solidity: function validityBondInAttoCash(address ) view returns(uint256)
func (_Token *TokenSession) ValidityBondInAttoCash(arg0 common.Address) (*big.Int, error) {
	return _Token.Contract.ValidityBondInAttoCash(&_Token.CallOpts, arg0)
}

// ValidityBondInAttoCash is a free data retrieval call binding the contract method 0xefb106ba.
//
// Solidity: function validityBondInAttoCash(address ) view returns(uint256)
func (_Token *TokenCallerSession) ValidityBondInAttoCash(arg0 common.Address) (*big.Int, error) {
	return _Token.Contract.ValidityBondInAttoCash(&_Token.CallOpts, arg0)
}

// CreateCategoricalMarket is a paid mutator transaction binding the contract method 0x45a62887.
//
// Solidity: function createCategoricalMarket(uint256 _endTime, uint256 _feePerCashInAttoCash, address _affiliateValidator, uint256 _affiliateFeeDivisor, address _designatedReporterAddress, bytes32[] _outcomes, string _extraInfo) returns(address _newMarket)
func (_Token *TokenTransactor) CreateCategoricalMarket(opts *bind.TransactOpts, _endTime *big.Int, _feePerCashInAttoCash *big.Int, _affiliateValidator common.Address, _affiliateFeeDivisor *big.Int, _designatedReporterAddress common.Address, _outcomes [][32]byte, _extraInfo string) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "createCategoricalMarket", _endTime, _feePerCashInAttoCash, _affiliateValidator, _affiliateFeeDivisor, _designatedReporterAddress, _outcomes, _extraInfo)
}

// CreateCategoricalMarket is a paid mutator transaction binding the contract method 0x45a62887.
//
// Solidity: function createCategoricalMarket(uint256 _endTime, uint256 _feePerCashInAttoCash, address _affiliateValidator, uint256 _affiliateFeeDivisor, address _designatedReporterAddress, bytes32[] _outcomes, string _extraInfo) returns(address _newMarket)
func (_Token *TokenSession) CreateCategoricalMarket(_endTime *big.Int, _feePerCashInAttoCash *big.Int, _affiliateValidator common.Address, _affiliateFeeDivisor *big.Int, _designatedReporterAddress common.Address, _outcomes [][32]byte, _extraInfo string) (*types.Transaction, error) {
	return _Token.Contract.CreateCategoricalMarket(&_Token.TransactOpts, _endTime, _feePerCashInAttoCash, _affiliateValidator, _affiliateFeeDivisor, _designatedReporterAddress, _outcomes, _extraInfo)
}

// CreateCategoricalMarket is a paid mutator transaction binding the contract method 0x45a62887.
//
// Solidity: function createCategoricalMarket(uint256 _endTime, uint256 _feePerCashInAttoCash, address _affiliateValidator, uint256 _affiliateFeeDivisor, address _designatedReporterAddress, bytes32[] _outcomes, string _extraInfo) returns(address _newMarket)
func (_Token *TokenTransactorSession) CreateCategoricalMarket(_endTime *big.Int, _feePerCashInAttoCash *big.Int, _affiliateValidator common.Address, _affiliateFeeDivisor *big.Int, _designatedReporterAddress common.Address, _outcomes [][32]byte, _extraInfo string) (*types.Transaction, error) {
	return _Token.Contract.CreateCategoricalMarket(&_Token.TransactOpts, _endTime, _feePerCashInAttoCash, _affiliateValidator, _affiliateFeeDivisor, _designatedReporterAddress, _outcomes, _extraInfo)
}

// CreateChildUniverse is a paid mutator transaction binding the contract method 0x3a537176.
//
// Solidity: function createChildUniverse(uint256[] _parentPayoutNumerators) returns(address)
func (_Token *TokenTransactor) CreateChildUniverse(opts *bind.TransactOpts, _parentPayoutNumerators []*big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "createChildUniverse", _parentPayoutNumerators)
}

// CreateChildUniverse is a paid mutator transaction binding the contract method 0x3a537176.
//
// Solidity: function createChildUniverse(uint256[] _parentPayoutNumerators) returns(address)
func (_Token *TokenSession) CreateChildUniverse(_parentPayoutNumerators []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.CreateChildUniverse(&_Token.TransactOpts, _parentPayoutNumerators)
}

// CreateChildUniverse is a paid mutator transaction binding the contract method 0x3a537176.
//
// Solidity: function createChildUniverse(uint256[] _parentPayoutNumerators) returns(address)
func (_Token *TokenTransactorSession) CreateChildUniverse(_parentPayoutNumerators []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.CreateChildUniverse(&_Token.TransactOpts, _parentPayoutNumerators)
}

// CreateScalarMarket is a paid mutator transaction binding the contract method 0x11a80ffc.
//
// Solidity: function createScalarMarket(uint256 _endTime, uint256 _feePerCashInAttoCash, address _affiliateValidator, uint256 _affiliateFeeDivisor, address _designatedReporterAddress, int256[] _prices, uint256 _numTicks, string _extraInfo) returns(address _newMarket)
func (_Token *TokenTransactor) CreateScalarMarket(opts *bind.TransactOpts, _endTime *big.Int, _feePerCashInAttoCash *big.Int, _affiliateValidator common.Address, _affiliateFeeDivisor *big.Int, _designatedReporterAddress common.Address, _prices []*big.Int, _numTicks *big.Int, _extraInfo string) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "createScalarMarket", _endTime, _feePerCashInAttoCash, _affiliateValidator, _affiliateFeeDivisor, _designatedReporterAddress, _prices, _numTicks, _extraInfo)
}

// CreateScalarMarket is a paid mutator transaction binding the contract method 0x11a80ffc.
//
// Solidity: function createScalarMarket(uint256 _endTime, uint256 _feePerCashInAttoCash, address _affiliateValidator, uint256 _affiliateFeeDivisor, address _designatedReporterAddress, int256[] _prices, uint256 _numTicks, string _extraInfo) returns(address _newMarket)
func (_Token *TokenSession) CreateScalarMarket(_endTime *big.Int, _feePerCashInAttoCash *big.Int, _affiliateValidator common.Address, _affiliateFeeDivisor *big.Int, _designatedReporterAddress common.Address, _prices []*big.Int, _numTicks *big.Int, _extraInfo string) (*types.Transaction, error) {
	return _Token.Contract.CreateScalarMarket(&_Token.TransactOpts, _endTime, _feePerCashInAttoCash, _affiliateValidator, _affiliateFeeDivisor, _designatedReporterAddress, _prices, _numTicks, _extraInfo)
}

// CreateScalarMarket is a paid mutator transaction binding the contract method 0x11a80ffc.
//
// Solidity: function createScalarMarket(uint256 _endTime, uint256 _feePerCashInAttoCash, address _affiliateValidator, uint256 _affiliateFeeDivisor, address _designatedReporterAddress, int256[] _prices, uint256 _numTicks, string _extraInfo) returns(address _newMarket)
func (_Token *TokenTransactorSession) CreateScalarMarket(_endTime *big.Int, _feePerCashInAttoCash *big.Int, _affiliateValidator common.Address, _affiliateFeeDivisor *big.Int, _designatedReporterAddress common.Address, _prices []*big.Int, _numTicks *big.Int, _extraInfo string) (*types.Transaction, error) {
	return _Token.Contract.CreateScalarMarket(&_Token.TransactOpts, _endTime, _feePerCashInAttoCash, _affiliateValidator, _affiliateFeeDivisor, _designatedReporterAddress, _prices, _numTicks, _extraInfo)
}

// CreateYesNoMarket is a paid mutator transaction binding the contract method 0xa95b973c.
//
// Solidity: function createYesNoMarket(uint256 _endTime, uint256 _feePerCashInAttoCash, address _affiliateValidator, uint256 _affiliateFeeDivisor, address _designatedReporterAddress, string _extraInfo) returns(address _newMarket)
func (_Token *TokenTransactor) CreateYesNoMarket(opts *bind.TransactOpts, _endTime *big.Int, _feePerCashInAttoCash *big.Int, _affiliateValidator common.Address, _affiliateFeeDivisor *big.Int, _designatedReporterAddress common.Address, _extraInfo string) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "createYesNoMarket", _endTime, _feePerCashInAttoCash, _affiliateValidator, _affiliateFeeDivisor, _designatedReporterAddress, _extraInfo)
}

// CreateYesNoMarket is a paid mutator transaction binding the contract method 0xa95b973c.
//
// Solidity: function createYesNoMarket(uint256 _endTime, uint256 _feePerCashInAttoCash, address _affiliateValidator, uint256 _affiliateFeeDivisor, address _designatedReporterAddress, string _extraInfo) returns(address _newMarket)
func (_Token *TokenSession) CreateYesNoMarket(_endTime *big.Int, _feePerCashInAttoCash *big.Int, _affiliateValidator common.Address, _affiliateFeeDivisor *big.Int, _designatedReporterAddress common.Address, _extraInfo string) (*types.Transaction, error) {
	return _Token.Contract.CreateYesNoMarket(&_Token.TransactOpts, _endTime, _feePerCashInAttoCash, _affiliateValidator, _affiliateFeeDivisor, _designatedReporterAddress, _extraInfo)
}

// CreateYesNoMarket is a paid mutator transaction binding the contract method 0xa95b973c.
//
// Solidity: function createYesNoMarket(uint256 _endTime, uint256 _feePerCashInAttoCash, address _affiliateValidator, uint256 _affiliateFeeDivisor, address _designatedReporterAddress, string _extraInfo) returns(address _newMarket)
func (_Token *TokenTransactorSession) CreateYesNoMarket(_endTime *big.Int, _feePerCashInAttoCash *big.Int, _affiliateValidator common.Address, _affiliateFeeDivisor *big.Int, _designatedReporterAddress common.Address, _extraInfo string) (*types.Transaction, error) {
	return _Token.Contract.CreateYesNoMarket(&_Token.TransactOpts, _endTime, _feePerCashInAttoCash, _affiliateValidator, _affiliateFeeDivisor, _designatedReporterAddress, _extraInfo)
}

// DecrementOpenInterest is a paid mutator transaction binding the contract method 0xb62418a1.
//
// Solidity: function decrementOpenInterest(uint256 _amount) returns(bool)
func (_Token *TokenTransactor) DecrementOpenInterest(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "decrementOpenInterest", _amount)
}

// DecrementOpenInterest is a paid mutator transaction binding the contract method 0xb62418a1.
//
// Solidity: function decrementOpenInterest(uint256 _amount) returns(bool)
func (_Token *TokenSession) DecrementOpenInterest(_amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.DecrementOpenInterest(&_Token.TransactOpts, _amount)
}

// DecrementOpenInterest is a paid mutator transaction binding the contract method 0xb62418a1.
//
// Solidity: function decrementOpenInterest(uint256 _amount) returns(bool)
func (_Token *TokenTransactorSession) DecrementOpenInterest(_amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.DecrementOpenInterest(&_Token.TransactOpts, _amount)
}

// DecrementOpenInterestFromMarket is a paid mutator transaction binding the contract method 0x180ef158.
//
// Solidity: function decrementOpenInterestFromMarket(address _market) returns(bool)
func (_Token *TokenTransactor) DecrementOpenInterestFromMarket(opts *bind.TransactOpts, _market common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "decrementOpenInterestFromMarket", _market)
}

// DecrementOpenInterestFromMarket is a paid mutator transaction binding the contract method 0x180ef158.
//
// Solidity: function decrementOpenInterestFromMarket(address _market) returns(bool)
func (_Token *TokenSession) DecrementOpenInterestFromMarket(_market common.Address) (*types.Transaction, error) {
	return _Token.Contract.DecrementOpenInterestFromMarket(&_Token.TransactOpts, _market)
}

// DecrementOpenInterestFromMarket is a paid mutator transaction binding the contract method 0x180ef158.
//
// Solidity: function decrementOpenInterestFromMarket(address _market) returns(bool)
func (_Token *TokenTransactorSession) DecrementOpenInterestFromMarket(_market common.Address) (*types.Transaction, error) {
	return _Token.Contract.DecrementOpenInterestFromMarket(&_Token.TransactOpts, _market)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address _sender, uint256 _amount, address _market) returns(bool)
func (_Token *TokenTransactor) Deposit(opts *bind.TransactOpts, _sender common.Address, _amount *big.Int, _market common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "deposit", _sender, _amount, _market)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address _sender, uint256 _amount, address _market) returns(bool)
func (_Token *TokenSession) Deposit(_sender common.Address, _amount *big.Int, _market common.Address) (*types.Transaction, error) {
	return _Token.Contract.Deposit(&_Token.TransactOpts, _sender, _amount, _market)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address _sender, uint256 _amount, address _market) returns(bool)
func (_Token *TokenTransactorSession) Deposit(_sender common.Address, _amount *big.Int, _market common.Address) (*types.Transaction, error) {
	return _Token.Contract.Deposit(&_Token.TransactOpts, _sender, _amount, _market)
}

// Fork is a paid mutator transaction binding the contract method 0x4591c060.
//
// Solidity: function fork() returns(bool)
func (_Token *TokenTransactor) Fork(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "fork")
}

// Fork is a paid mutator transaction binding the contract method 0x4591c060.
//
// Solidity: function fork() returns(bool)
func (_Token *TokenSession) Fork() (*types.Transaction, error) {
	return _Token.Contract.Fork(&_Token.TransactOpts)
}

// Fork is a paid mutator transaction binding the contract method 0x4591c060.
//
// Solidity: function fork() returns(bool)
func (_Token *TokenTransactorSession) Fork() (*types.Transaction, error) {
	return _Token.Contract.Fork(&_Token.TransactOpts)
}

// GetOrCacheDesignatedReportNoShowBond is a paid mutator transaction binding the contract method 0xfd1e5e7a.
//
// Solidity: function getOrCacheDesignatedReportNoShowBond() returns(uint256)
func (_Token *TokenTransactor) GetOrCacheDesignatedReportNoShowBond(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "getOrCacheDesignatedReportNoShowBond")
}

// GetOrCacheDesignatedReportNoShowBond is a paid mutator transaction binding the contract method 0xfd1e5e7a.
//
// Solidity: function getOrCacheDesignatedReportNoShowBond() returns(uint256)
func (_Token *TokenSession) GetOrCacheDesignatedReportNoShowBond() (*types.Transaction, error) {
	return _Token.Contract.GetOrCacheDesignatedReportNoShowBond(&_Token.TransactOpts)
}

// GetOrCacheDesignatedReportNoShowBond is a paid mutator transaction binding the contract method 0xfd1e5e7a.
//
// Solidity: function getOrCacheDesignatedReportNoShowBond() returns(uint256)
func (_Token *TokenTransactorSession) GetOrCacheDesignatedReportNoShowBond() (*types.Transaction, error) {
	return _Token.Contract.GetOrCacheDesignatedReportNoShowBond(&_Token.TransactOpts)
}

// GetOrCacheDesignatedReportStake is a paid mutator transaction binding the contract method 0xe79609e2.
//
// Solidity: function getOrCacheDesignatedReportStake() returns(uint256)
func (_Token *TokenTransactor) GetOrCacheDesignatedReportStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "getOrCacheDesignatedReportStake")
}

// GetOrCacheDesignatedReportStake is a paid mutator transaction binding the contract method 0xe79609e2.
//
// Solidity: function getOrCacheDesignatedReportStake() returns(uint256)
func (_Token *TokenSession) GetOrCacheDesignatedReportStake() (*types.Transaction, error) {
	return _Token.Contract.GetOrCacheDesignatedReportStake(&_Token.TransactOpts)
}

// GetOrCacheDesignatedReportStake is a paid mutator transaction binding the contract method 0xe79609e2.
//
// Solidity: function getOrCacheDesignatedReportStake() returns(uint256)
func (_Token *TokenTransactorSession) GetOrCacheDesignatedReportStake() (*types.Transaction, error) {
	return _Token.Contract.GetOrCacheDesignatedReportStake(&_Token.TransactOpts)
}

// GetOrCacheMarketRepBond is a paid mutator transaction binding the contract method 0xa7e8d762.
//
// Solidity: function getOrCacheMarketRepBond() returns(uint256)
func (_Token *TokenTransactor) GetOrCacheMarketRepBond(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "getOrCacheMarketRepBond")
}

// GetOrCacheMarketRepBond is a paid mutator transaction binding the contract method 0xa7e8d762.
//
// Solidity: function getOrCacheMarketRepBond() returns(uint256)
func (_Token *TokenSession) GetOrCacheMarketRepBond() (*types.Transaction, error) {
	return _Token.Contract.GetOrCacheMarketRepBond(&_Token.TransactOpts)
}

// GetOrCacheMarketRepBond is a paid mutator transaction binding the contract method 0xa7e8d762.
//
// Solidity: function getOrCacheMarketRepBond() returns(uint256)
func (_Token *TokenTransactorSession) GetOrCacheMarketRepBond() (*types.Transaction, error) {
	return _Token.Contract.GetOrCacheMarketRepBond(&_Token.TransactOpts)
}

// GetOrCacheReportingFeeDivisor is a paid mutator transaction binding the contract method 0x8f93bffe.
//
// Solidity: function getOrCacheReportingFeeDivisor() returns(uint256)
func (_Token *TokenTransactor) GetOrCacheReportingFeeDivisor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "getOrCacheReportingFeeDivisor")
}

// GetOrCacheReportingFeeDivisor is a paid mutator transaction binding the contract method 0x8f93bffe.
//
// Solidity: function getOrCacheReportingFeeDivisor() returns(uint256)
func (_Token *TokenSession) GetOrCacheReportingFeeDivisor() (*types.Transaction, error) {
	return _Token.Contract.GetOrCacheReportingFeeDivisor(&_Token.TransactOpts)
}

// GetOrCacheReportingFeeDivisor is a paid mutator transaction binding the contract method 0x8f93bffe.
//
// Solidity: function getOrCacheReportingFeeDivisor() returns(uint256)
func (_Token *TokenTransactorSession) GetOrCacheReportingFeeDivisor() (*types.Transaction, error) {
	return _Token.Contract.GetOrCacheReportingFeeDivisor(&_Token.TransactOpts)
}

// GetOrCacheValidityBond is a paid mutator transaction binding the contract method 0xaf4cd457.
//
// Solidity: function getOrCacheValidityBond() returns(uint256)
func (_Token *TokenTransactor) GetOrCacheValidityBond(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "getOrCacheValidityBond")
}

// GetOrCacheValidityBond is a paid mutator transaction binding the contract method 0xaf4cd457.
//
// Solidity: function getOrCacheValidityBond() returns(uint256)
func (_Token *TokenSession) GetOrCacheValidityBond() (*types.Transaction, error) {
	return _Token.Contract.GetOrCacheValidityBond(&_Token.TransactOpts)
}

// GetOrCacheValidityBond is a paid mutator transaction binding the contract method 0xaf4cd457.
//
// Solidity: function getOrCacheValidityBond() returns(uint256)
func (_Token *TokenTransactorSession) GetOrCacheValidityBond() (*types.Transaction, error) {
	return _Token.Contract.GetOrCacheValidityBond(&_Token.TransactOpts)
}

// GetOrCreateCurrentDisputeWindow is a paid mutator transaction binding the contract method 0xe3fa4b04.
//
// Solidity: function getOrCreateCurrentDisputeWindow(bool _initial) returns(address)
func (_Token *TokenTransactor) GetOrCreateCurrentDisputeWindow(opts *bind.TransactOpts, _initial bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "getOrCreateCurrentDisputeWindow", _initial)
}

// GetOrCreateCurrentDisputeWindow is a paid mutator transaction binding the contract method 0xe3fa4b04.
//
// Solidity: function getOrCreateCurrentDisputeWindow(bool _initial) returns(address)
func (_Token *TokenSession) GetOrCreateCurrentDisputeWindow(_initial bool) (*types.Transaction, error) {
	return _Token.Contract.GetOrCreateCurrentDisputeWindow(&_Token.TransactOpts, _initial)
}

// GetOrCreateCurrentDisputeWindow is a paid mutator transaction binding the contract method 0xe3fa4b04.
//
// Solidity: function getOrCreateCurrentDisputeWindow(bool _initial) returns(address)
func (_Token *TokenTransactorSession) GetOrCreateCurrentDisputeWindow(_initial bool) (*types.Transaction, error) {
	return _Token.Contract.GetOrCreateCurrentDisputeWindow(&_Token.TransactOpts, _initial)
}

// GetOrCreateDisputeWindowByTimestamp is a paid mutator transaction binding the contract method 0x8689526b.
//
// Solidity: function getOrCreateDisputeWindowByTimestamp(uint256 _timestamp, bool _initial) returns(address)
func (_Token *TokenTransactor) GetOrCreateDisputeWindowByTimestamp(opts *bind.TransactOpts, _timestamp *big.Int, _initial bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "getOrCreateDisputeWindowByTimestamp", _timestamp, _initial)
}

// GetOrCreateDisputeWindowByTimestamp is a paid mutator transaction binding the contract method 0x8689526b.
//
// Solidity: function getOrCreateDisputeWindowByTimestamp(uint256 _timestamp, bool _initial) returns(address)
func (_Token *TokenSession) GetOrCreateDisputeWindowByTimestamp(_timestamp *big.Int, _initial bool) (*types.Transaction, error) {
	return _Token.Contract.GetOrCreateDisputeWindowByTimestamp(&_Token.TransactOpts, _timestamp, _initial)
}

// GetOrCreateDisputeWindowByTimestamp is a paid mutator transaction binding the contract method 0x8689526b.
//
// Solidity: function getOrCreateDisputeWindowByTimestamp(uint256 _timestamp, bool _initial) returns(address)
func (_Token *TokenTransactorSession) GetOrCreateDisputeWindowByTimestamp(_timestamp *big.Int, _initial bool) (*types.Transaction, error) {
	return _Token.Contract.GetOrCreateDisputeWindowByTimestamp(&_Token.TransactOpts, _timestamp, _initial)
}

// GetOrCreateNextDisputeWindow is a paid mutator transaction binding the contract method 0x92394f32.
//
// Solidity: function getOrCreateNextDisputeWindow(bool _initial) returns(address)
func (_Token *TokenTransactor) GetOrCreateNextDisputeWindow(opts *bind.TransactOpts, _initial bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "getOrCreateNextDisputeWindow", _initial)
}

// GetOrCreateNextDisputeWindow is a paid mutator transaction binding the contract method 0x92394f32.
//
// Solidity: function getOrCreateNextDisputeWindow(bool _initial) returns(address)
func (_Token *TokenSession) GetOrCreateNextDisputeWindow(_initial bool) (*types.Transaction, error) {
	return _Token.Contract.GetOrCreateNextDisputeWindow(&_Token.TransactOpts, _initial)
}

// GetOrCreateNextDisputeWindow is a paid mutator transaction binding the contract method 0x92394f32.
//
// Solidity: function getOrCreateNextDisputeWindow(bool _initial) returns(address)
func (_Token *TokenTransactorSession) GetOrCreateNextDisputeWindow(_initial bool) (*types.Transaction, error) {
	return _Token.Contract.GetOrCreateNextDisputeWindow(&_Token.TransactOpts, _initial)
}

// GetOrCreatePreviousDisputeWindow is a paid mutator transaction binding the contract method 0xe2d8edaf.
//
// Solidity: function getOrCreatePreviousDisputeWindow(bool _initial) returns(address)
func (_Token *TokenTransactor) GetOrCreatePreviousDisputeWindow(opts *bind.TransactOpts, _initial bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "getOrCreatePreviousDisputeWindow", _initial)
}

// GetOrCreatePreviousDisputeWindow is a paid mutator transaction binding the contract method 0xe2d8edaf.
//
// Solidity: function getOrCreatePreviousDisputeWindow(bool _initial) returns(address)
func (_Token *TokenSession) GetOrCreatePreviousDisputeWindow(_initial bool) (*types.Transaction, error) {
	return _Token.Contract.GetOrCreatePreviousDisputeWindow(&_Token.TransactOpts, _initial)
}

// GetOrCreatePreviousDisputeWindow is a paid mutator transaction binding the contract method 0xe2d8edaf.
//
// Solidity: function getOrCreatePreviousDisputeWindow(bool _initial) returns(address)
func (_Token *TokenTransactorSession) GetOrCreatePreviousDisputeWindow(_initial bool) (*types.Transaction, error) {
	return _Token.Contract.GetOrCreatePreviousDisputeWindow(&_Token.TransactOpts, _initial)
}

// GetOrCreatePreviousPreviousDisputeWindow is a paid mutator transaction binding the contract method 0xf28b0956.
//
// Solidity: function getOrCreatePreviousPreviousDisputeWindow(bool _initial) returns(address)
func (_Token *TokenTransactor) GetOrCreatePreviousPreviousDisputeWindow(opts *bind.TransactOpts, _initial bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "getOrCreatePreviousPreviousDisputeWindow", _initial)
}

// GetOrCreatePreviousPreviousDisputeWindow is a paid mutator transaction binding the contract method 0xf28b0956.
//
// Solidity: function getOrCreatePreviousPreviousDisputeWindow(bool _initial) returns(address)
func (_Token *TokenSession) GetOrCreatePreviousPreviousDisputeWindow(_initial bool) (*types.Transaction, error) {
	return _Token.Contract.GetOrCreatePreviousPreviousDisputeWindow(&_Token.TransactOpts, _initial)
}

// GetOrCreatePreviousPreviousDisputeWindow is a paid mutator transaction binding the contract method 0xf28b0956.
//
// Solidity: function getOrCreatePreviousPreviousDisputeWindow(bool _initial) returns(address)
func (_Token *TokenTransactorSession) GetOrCreatePreviousPreviousDisputeWindow(_initial bool) (*types.Transaction, error) {
	return _Token.Contract.GetOrCreatePreviousPreviousDisputeWindow(&_Token.TransactOpts, _initial)
}

// IncrementOpenInterest is a paid mutator transaction binding the contract method 0xce483e88.
//
// Solidity: function incrementOpenInterest(uint256 _amount) returns(bool)
func (_Token *TokenTransactor) IncrementOpenInterest(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "incrementOpenInterest", _amount)
}

// IncrementOpenInterest is a paid mutator transaction binding the contract method 0xce483e88.
//
// Solidity: function incrementOpenInterest(uint256 _amount) returns(bool)
func (_Token *TokenSession) IncrementOpenInterest(_amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.IncrementOpenInterest(&_Token.TransactOpts, _amount)
}

// IncrementOpenInterest is a paid mutator transaction binding the contract method 0xce483e88.
//
// Solidity: function incrementOpenInterest(uint256 _amount) returns(bool)
func (_Token *TokenTransactorSession) IncrementOpenInterest(_amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.IncrementOpenInterest(&_Token.TransactOpts, _amount)
}

// MigrateMarketIn is a paid mutator transaction binding the contract method 0x8d2ecfba.
//
// Solidity: function migrateMarketIn(address _market, uint256 _cashBalance, uint256 _marketOI) returns(bool)
func (_Token *TokenTransactor) MigrateMarketIn(opts *bind.TransactOpts, _market common.Address, _cashBalance *big.Int, _marketOI *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "migrateMarketIn", _market, _cashBalance, _marketOI)
}

// MigrateMarketIn is a paid mutator transaction binding the contract method 0x8d2ecfba.
//
// Solidity: function migrateMarketIn(address _market, uint256 _cashBalance, uint256 _marketOI) returns(bool)
func (_Token *TokenSession) MigrateMarketIn(_market common.Address, _cashBalance *big.Int, _marketOI *big.Int) (*types.Transaction, error) {
	return _Token.Contract.MigrateMarketIn(&_Token.TransactOpts, _market, _cashBalance, _marketOI)
}

// MigrateMarketIn is a paid mutator transaction binding the contract method 0x8d2ecfba.
//
// Solidity: function migrateMarketIn(address _market, uint256 _cashBalance, uint256 _marketOI) returns(bool)
func (_Token *TokenTransactorSession) MigrateMarketIn(_market common.Address, _cashBalance *big.Int, _marketOI *big.Int) (*types.Transaction, error) {
	return _Token.Contract.MigrateMarketIn(&_Token.TransactOpts, _market, _cashBalance, _marketOI)
}

// MigrateMarketOut is a paid mutator transaction binding the contract method 0x11be56d7.
//
// Solidity: function migrateMarketOut(address _destinationUniverse) returns(bool)
func (_Token *TokenTransactor) MigrateMarketOut(opts *bind.TransactOpts, _destinationUniverse common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "migrateMarketOut", _destinationUniverse)
}

// MigrateMarketOut is a paid mutator transaction binding the contract method 0x11be56d7.
//
// Solidity: function migrateMarketOut(address _destinationUniverse) returns(bool)
func (_Token *TokenSession) MigrateMarketOut(_destinationUniverse common.Address) (*types.Transaction, error) {
	return _Token.Contract.MigrateMarketOut(&_Token.TransactOpts, _destinationUniverse)
}

// MigrateMarketOut is a paid mutator transaction binding the contract method 0x11be56d7.
//
// Solidity: function migrateMarketOut(address _destinationUniverse) returns(bool)
func (_Token *TokenTransactorSession) MigrateMarketOut(_destinationUniverse common.Address) (*types.Transaction, error) {
	return _Token.Contract.MigrateMarketOut(&_Token.TransactOpts, _destinationUniverse)
}

// PokeRepMarketCapInAttoCash is a paid mutator transaction binding the contract method 0x13bf24c1.
//
// Solidity: function pokeRepMarketCapInAttoCash() returns(uint256)
func (_Token *TokenTransactor) PokeRepMarketCapInAttoCash(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "pokeRepMarketCapInAttoCash")
}

// PokeRepMarketCapInAttoCash is a paid mutator transaction binding the contract method 0x13bf24c1.
//
// Solidity: function pokeRepMarketCapInAttoCash() returns(uint256)
func (_Token *TokenSession) PokeRepMarketCapInAttoCash() (*types.Transaction, error) {
	return _Token.Contract.PokeRepMarketCapInAttoCash(&_Token.TransactOpts)
}

// PokeRepMarketCapInAttoCash is a paid mutator transaction binding the contract method 0x13bf24c1.
//
// Solidity: function pokeRepMarketCapInAttoCash() returns(uint256)
func (_Token *TokenTransactorSession) PokeRepMarketCapInAttoCash() (*types.Transaction, error) {
	return _Token.Contract.PokeRepMarketCapInAttoCash(&_Token.TransactOpts)
}

// RunPeriodicals is a paid mutator transaction binding the contract method 0x98fedb09.
//
// Solidity: function runPeriodicals() returns(bool)
func (_Token *TokenTransactor) RunPeriodicals(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "runPeriodicals")
}

// RunPeriodicals is a paid mutator transaction binding the contract method 0x98fedb09.
//
// Solidity: function runPeriodicals() returns(bool)
func (_Token *TokenSession) RunPeriodicals() (*types.Transaction, error) {
	return _Token.Contract.RunPeriodicals(&_Token.TransactOpts)
}

// RunPeriodicals is a paid mutator transaction binding the contract method 0x98fedb09.
//
// Solidity: function runPeriodicals() returns(bool)
func (_Token *TokenTransactorSession) RunPeriodicals() (*types.Transaction, error) {
	return _Token.Contract.RunPeriodicals(&_Token.TransactOpts)
}

// SweepInterest is a paid mutator transaction binding the contract method 0x3342f689.
//
// Solidity: function sweepInterest() returns(bool)
func (_Token *TokenTransactor) SweepInterest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "sweepInterest")
}

// SweepInterest is a paid mutator transaction binding the contract method 0x3342f689.
//
// Solidity: function sweepInterest() returns(bool)
func (_Token *TokenSession) SweepInterest() (*types.Transaction, error) {
	return _Token.Contract.SweepInterest(&_Token.TransactOpts)
}

// SweepInterest is a paid mutator transaction binding the contract method 0x3342f689.
//
// Solidity: function sweepInterest() returns(bool)
func (_Token *TokenTransactorSession) SweepInterest() (*types.Transaction, error) {
	return _Token.Contract.SweepInterest(&_Token.TransactOpts)
}

// UpdateForkValues is a paid mutator transaction binding the contract method 0x9ab448d9.
//
// Solidity: function updateForkValues() returns(bool)
func (_Token *TokenTransactor) UpdateForkValues(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "updateForkValues")
}

// UpdateForkValues is a paid mutator transaction binding the contract method 0x9ab448d9.
//
// Solidity: function updateForkValues() returns(bool)
func (_Token *TokenSession) UpdateForkValues() (*types.Transaction, error) {
	return _Token.Contract.UpdateForkValues(&_Token.TransactOpts)
}

// UpdateForkValues is a paid mutator transaction binding the contract method 0x9ab448d9.
//
// Solidity: function updateForkValues() returns(bool)
func (_Token *TokenTransactorSession) UpdateForkValues() (*types.Transaction, error) {
	return _Token.Contract.UpdateForkValues(&_Token.TransactOpts)
}

// UpdateTentativeWinningChildUniverse is a paid mutator transaction binding the contract method 0xf7095d9d.
//
// Solidity: function updateTentativeWinningChildUniverse(bytes32 _parentPayoutDistributionHash) returns(bool)
func (_Token *TokenTransactor) UpdateTentativeWinningChildUniverse(opts *bind.TransactOpts, _parentPayoutDistributionHash [32]byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "updateTentativeWinningChildUniverse", _parentPayoutDistributionHash)
}

// UpdateTentativeWinningChildUniverse is a paid mutator transaction binding the contract method 0xf7095d9d.
//
// Solidity: function updateTentativeWinningChildUniverse(bytes32 _parentPayoutDistributionHash) returns(bool)
func (_Token *TokenSession) UpdateTentativeWinningChildUniverse(_parentPayoutDistributionHash [32]byte) (*types.Transaction, error) {
	return _Token.Contract.UpdateTentativeWinningChildUniverse(&_Token.TransactOpts, _parentPayoutDistributionHash)
}

// UpdateTentativeWinningChildUniverse is a paid mutator transaction binding the contract method 0xf7095d9d.
//
// Solidity: function updateTentativeWinningChildUniverse(bytes32 _parentPayoutDistributionHash) returns(bool)
func (_Token *TokenTransactorSession) UpdateTentativeWinningChildUniverse(_parentPayoutDistributionHash [32]byte) (*types.Transaction, error) {
	return _Token.Contract.UpdateTentativeWinningChildUniverse(&_Token.TransactOpts, _parentPayoutDistributionHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address _recipient, uint256 _amount, address _market) returns(bool)
func (_Token *TokenTransactor) Withdraw(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int, _market common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "withdraw", _recipient, _amount, _market)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address _recipient, uint256 _amount, address _market) returns(bool)
func (_Token *TokenSession) Withdraw(_recipient common.Address, _amount *big.Int, _market common.Address) (*types.Transaction, error) {
	return _Token.Contract.Withdraw(&_Token.TransactOpts, _recipient, _amount, _market)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address _recipient, uint256 _amount, address _market) returns(bool)
func (_Token *TokenTransactorSession) Withdraw(_recipient common.Address, _amount *big.Int, _market common.Address) (*types.Transaction, error) {
	return _Token.Contract.Withdraw(&_Token.TransactOpts, _recipient, _amount, _market)
}
