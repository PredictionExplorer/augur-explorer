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

// SportsLinkMarketFactoryMarketDetails is an auto generated low-level Go binding around an user-defined struct.
type SportsLinkMarketFactoryMarketDetails struct {
	EventId            *big.Int
	HomeTeamId         *big.Int
	AwayTeamId         *big.Int
	EstimatedStartTime *big.Int
	MarketType         uint8
	EventStatus        uint8
	Value0             *big.Int
}

// SportsLinkMarketFactoryMetaData contains all meta data concerning the SportsLinkMarketFactory contract.
var SportsLinkMarketFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Full\",\"name\":\"_collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_shareFactor\",\"type\":\"uint256\"},{\"internalType\":\"contractFeePot\",\"name\":\"_feePot\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stakerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_settlementFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_protocol\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_protocolFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_linkNode\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sportId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newLinkNode\",\"type\":\"address\"}],\"name\":\"LinkNodeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumSportsLinkMarketFactory.MarketType\",\"name\":\"marketType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"eventId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"homeTeamId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"awayTeamId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"estimatedStartTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"score\",\"type\":\"int256\"}],\"name\":\"MarketCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"MarketResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"protocol\",\"type\":\"address\"}],\"name\":\"ProtocolChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"ProtocolFeeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"protocol\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ProtocolFeeClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"SettlementFeeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"settlementAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SettlementFeeClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SharesBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SharesMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"StakerFeeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winningOutcome\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"settlementFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payout\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"WinningsClaimed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"accumulatedProtocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accumulatedSettlementFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sharesToBurn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"burnShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"calcCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_collateralIn\",\"type\":\"uint256\"}],\"name\":\"calcShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claimManyWinnings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimProtocolFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claimSettlementFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claimWinnings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateral\",\"outputs\":[{\"internalType\":\"contractIERC20Full\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_payload\",\"type\":\"bytes32\"}],\"name\":\"createMarket\",\"outputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"_ids\",\"type\":\"uint256[3]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_payload\",\"type\":\"bytes32\"}],\"name\":\"decodeCreation\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"_eventId\",\"type\":\"uint128\"},{\"internalType\":\"uint16\",\"name\":\"_homeTeamId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_awayTeamId\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"_startTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"int16\",\"name\":\"_homeSpread\",\"type\":\"int16\"},{\"internalType\":\"uint16\",\"name\":\"_totalScore\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"_createSpread\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_createTotal\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_payload\",\"type\":\"bytes32\"}],\"name\":\"decodeResolution\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"_eventId\",\"type\":\"uint128\"},{\"internalType\":\"uint8\",\"name\":\"_eventStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"_homeScore\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_awayScore\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_eventId\",\"type\":\"uint128\"},{\"internalType\":\"uint16\",\"name\":\"_homeTeamId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_awayTeamId\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"_startTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"int16\",\"name\":\"_homeSpread\",\"type\":\"int16\"},{\"internalType\":\"uint16\",\"name\":\"_totalScore\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"_createSpread\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_createTotal\",\"type\":\"bool\"}],\"name\":\"encodeCreation\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_payload\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_eventId\",\"type\":\"uint128\"},{\"internalType\":\"uint8\",\"name\":\"_eventStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"_homeScore\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_awayScore\",\"type\":\"uint16\"}],\"name\":\"encodeResolution\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_payload\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"events\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"homeScore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"awayScore\",\"type\":\"uint256\"},{\"internalType\":\"enumSportsLinkMarketFactory.EventStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"resolutionTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"finalized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePot\",\"outputs\":[{\"internalType\":\"contractFeePot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_eventId\",\"type\":\"uint256\"}],\"name\":\"getEventMarkets\",\"outputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"\",\"type\":\"uint256[3]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getMarket\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"settlementAddress\",\"type\":\"address\"},{\"internalType\":\"contractOwnedERC20[]\",\"name\":\"shareTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"contractOwnedERC20\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"settlementFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structAbstractMarketFactory.Market\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_marketId\",\"type\":\"uint256\"}],\"name\":\"getMarketDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"eventId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"homeTeamId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"awayTeamId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"estimatedStartTime\",\"type\":\"uint256\"},{\"internalType\":\"enumSportsLinkMarketFactory.MarketType\",\"name\":\"marketType\",\"type\":\"uint8\"},{\"internalType\":\"enumSportsLinkMarketFactory.EventStatus\",\"name\":\"eventStatus\",\"type\":\"uint8\"},{\"internalType\":\"int256\",\"name\":\"value0\",\"type\":\"int256\"}],\"internalType\":\"structSportsLinkMarketFactory.MarketDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_eventId\",\"type\":\"uint256\"}],\"name\":\"isEventRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_eventId\",\"type\":\"uint256\"}],\"name\":\"isEventResolved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"isMarketResolved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkNode\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listOfEvents\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listResolvableEvents\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listUnresolvedMarkets\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_shareToMint\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"mintShares\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocol\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"resolveMarket\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newLinkNode\",\"type\":\"address\"}],\"name\":\"setLinkNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newProtocol\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_claimFirst\",\"type\":\"bool\"}],\"name\":\"setProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFee\",\"type\":\"uint256\"}],\"name\":\"setProtocolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFee\",\"type\":\"uint256\"}],\"name\":\"setSettlementFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFee\",\"type\":\"uint256\"}],\"name\":\"setStakerFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settlementFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shareFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sportId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_payload\",\"type\":\"bytes32\"}],\"name\":\"trustedResolveMarkets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SportsLinkMarketFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use SportsLinkMarketFactoryMetaData.ABI instead.
var SportsLinkMarketFactoryABI = SportsLinkMarketFactoryMetaData.ABI

// SportsLinkMarketFactory is an auto generated Go binding around an Ethereum contract.
type SportsLinkMarketFactory struct {
	SportsLinkMarketFactoryCaller     // Read-only binding to the contract
	SportsLinkMarketFactoryTransactor // Write-only binding to the contract
	SportsLinkMarketFactoryFilterer   // Log filterer for contract events
}

// SportsLinkMarketFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SportsLinkMarketFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SportsLinkMarketFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SportsLinkMarketFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SportsLinkMarketFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SportsLinkMarketFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SportsLinkMarketFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SportsLinkMarketFactorySession struct {
	Contract     *SportsLinkMarketFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SportsLinkMarketFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SportsLinkMarketFactoryCallerSession struct {
	Contract *SportsLinkMarketFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// SportsLinkMarketFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SportsLinkMarketFactoryTransactorSession struct {
	Contract     *SportsLinkMarketFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// SportsLinkMarketFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SportsLinkMarketFactoryRaw struct {
	Contract *SportsLinkMarketFactory // Generic contract binding to access the raw methods on
}

// SportsLinkMarketFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SportsLinkMarketFactoryCallerRaw struct {
	Contract *SportsLinkMarketFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// SportsLinkMarketFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SportsLinkMarketFactoryTransactorRaw struct {
	Contract *SportsLinkMarketFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSportsLinkMarketFactory creates a new instance of SportsLinkMarketFactory, bound to a specific deployed contract.
func NewSportsLinkMarketFactory(address common.Address, backend bind.ContractBackend) (*SportsLinkMarketFactory, error) {
	contract, err := bindSportsLinkMarketFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SportsLinkMarketFactory{SportsLinkMarketFactoryCaller: SportsLinkMarketFactoryCaller{contract: contract}, SportsLinkMarketFactoryTransactor: SportsLinkMarketFactoryTransactor{contract: contract}, SportsLinkMarketFactoryFilterer: SportsLinkMarketFactoryFilterer{contract: contract}}, nil
}

// NewSportsLinkMarketFactoryCaller creates a new read-only instance of SportsLinkMarketFactory, bound to a specific deployed contract.
func NewSportsLinkMarketFactoryCaller(address common.Address, caller bind.ContractCaller) (*SportsLinkMarketFactoryCaller, error) {
	contract, err := bindSportsLinkMarketFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SportsLinkMarketFactoryCaller{contract: contract}, nil
}

// NewSportsLinkMarketFactoryTransactor creates a new write-only instance of SportsLinkMarketFactory, bound to a specific deployed contract.
func NewSportsLinkMarketFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*SportsLinkMarketFactoryTransactor, error) {
	contract, err := bindSportsLinkMarketFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SportsLinkMarketFactoryTransactor{contract: contract}, nil
}

// NewSportsLinkMarketFactoryFilterer creates a new log filterer instance of SportsLinkMarketFactory, bound to a specific deployed contract.
func NewSportsLinkMarketFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*SportsLinkMarketFactoryFilterer, error) {
	contract, err := bindSportsLinkMarketFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SportsLinkMarketFactoryFilterer{contract: contract}, nil
}

// bindSportsLinkMarketFactory binds a generic wrapper to an already deployed contract.
func bindSportsLinkMarketFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SportsLinkMarketFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SportsLinkMarketFactory.Contract.SportsLinkMarketFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.SportsLinkMarketFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.SportsLinkMarketFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SportsLinkMarketFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.contract.Transact(opts, method, params...)
}

// AccumulatedProtocolFee is a free data retrieval call binding the contract method 0xa544a62c.
//
// Solidity: function accumulatedProtocolFee() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) AccumulatedProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "accumulatedProtocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedProtocolFee is a free data retrieval call binding the contract method 0xa544a62c.
//
// Solidity: function accumulatedProtocolFee() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) AccumulatedProtocolFee() (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.AccumulatedProtocolFee(&_SportsLinkMarketFactory.CallOpts)
}

// AccumulatedProtocolFee is a free data retrieval call binding the contract method 0xa544a62c.
//
// Solidity: function accumulatedProtocolFee() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) AccumulatedProtocolFee() (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.AccumulatedProtocolFee(&_SportsLinkMarketFactory.CallOpts)
}

// AccumulatedSettlementFees is a free data retrieval call binding the contract method 0x49a4d934.
//
// Solidity: function accumulatedSettlementFees(address ) view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) AccumulatedSettlementFees(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "accumulatedSettlementFees", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedSettlementFees is a free data retrieval call binding the contract method 0x49a4d934.
//
// Solidity: function accumulatedSettlementFees(address ) view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) AccumulatedSettlementFees(arg0 common.Address) (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.AccumulatedSettlementFees(&_SportsLinkMarketFactory.CallOpts, arg0)
}

// AccumulatedSettlementFees is a free data retrieval call binding the contract method 0x49a4d934.
//
// Solidity: function accumulatedSettlementFees(address ) view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) AccumulatedSettlementFees(arg0 common.Address) (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.AccumulatedSettlementFees(&_SportsLinkMarketFactory.CallOpts, arg0)
}

// CalcCost is a free data retrieval call binding the contract method 0x473a6d52.
//
// Solidity: function calcCost(uint256 _shares) view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) CalcCost(opts *bind.CallOpts, _shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "calcCost", _shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcCost is a free data retrieval call binding the contract method 0x473a6d52.
//
// Solidity: function calcCost(uint256 _shares) view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) CalcCost(_shares *big.Int) (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.CalcCost(&_SportsLinkMarketFactory.CallOpts, _shares)
}

// CalcCost is a free data retrieval call binding the contract method 0x473a6d52.
//
// Solidity: function calcCost(uint256 _shares) view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) CalcCost(_shares *big.Int) (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.CalcCost(&_SportsLinkMarketFactory.CallOpts, _shares)
}

// CalcShares is a free data retrieval call binding the contract method 0xcc87adea.
//
// Solidity: function calcShares(uint256 _collateralIn) view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) CalcShares(opts *bind.CallOpts, _collateralIn *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "calcShares", _collateralIn)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcShares is a free data retrieval call binding the contract method 0xcc87adea.
//
// Solidity: function calcShares(uint256 _collateralIn) view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) CalcShares(_collateralIn *big.Int) (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.CalcShares(&_SportsLinkMarketFactory.CallOpts, _collateralIn)
}

// CalcShares is a free data retrieval call binding the contract method 0xcc87adea.
//
// Solidity: function calcShares(uint256 _collateralIn) view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) CalcShares(_collateralIn *big.Int) (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.CalcShares(&_SportsLinkMarketFactory.CallOpts, _collateralIn)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) Collateral(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "collateral")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) Collateral() (common.Address, error) {
	return _SportsLinkMarketFactory.Contract.Collateral(&_SportsLinkMarketFactory.CallOpts)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) Collateral() (common.Address, error) {
	return _SportsLinkMarketFactory.Contract.Collateral(&_SportsLinkMarketFactory.CallOpts)
}

// DecodeCreation is a free data retrieval call binding the contract method 0xdb747ac9.
//
// Solidity: function decodeCreation(bytes32 _payload) pure returns(uint128 _eventId, uint16 _homeTeamId, uint16 _awayTeamId, uint32 _startTimestamp, int16 _homeSpread, uint16 _totalScore, bool _createSpread, bool _createTotal)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) DecodeCreation(opts *bind.CallOpts, _payload [32]byte) (struct {
	EventId        *big.Int
	HomeTeamId     uint16
	AwayTeamId     uint16
	StartTimestamp uint32
	HomeSpread     int16
	TotalScore     uint16
	CreateSpread   bool
	CreateTotal    bool
}, error) {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "decodeCreation", _payload)

	outstruct := new(struct {
		EventId        *big.Int
		HomeTeamId     uint16
		AwayTeamId     uint16
		StartTimestamp uint32
		HomeSpread     int16
		TotalScore     uint16
		CreateSpread   bool
		CreateTotal    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EventId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.HomeTeamId = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.AwayTeamId = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.StartTimestamp = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.HomeSpread = *abi.ConvertType(out[4], new(int16)).(*int16)
	outstruct.TotalScore = *abi.ConvertType(out[5], new(uint16)).(*uint16)
	outstruct.CreateSpread = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.CreateTotal = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// DecodeCreation is a free data retrieval call binding the contract method 0xdb747ac9.
//
// Solidity: function decodeCreation(bytes32 _payload) pure returns(uint128 _eventId, uint16 _homeTeamId, uint16 _awayTeamId, uint32 _startTimestamp, int16 _homeSpread, uint16 _totalScore, bool _createSpread, bool _createTotal)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) DecodeCreation(_payload [32]byte) (struct {
	EventId        *big.Int
	HomeTeamId     uint16
	AwayTeamId     uint16
	StartTimestamp uint32
	HomeSpread     int16
	TotalScore     uint16
	CreateSpread   bool
	CreateTotal    bool
}, error) {
	return _SportsLinkMarketFactory.Contract.DecodeCreation(&_SportsLinkMarketFactory.CallOpts, _payload)
}

// DecodeCreation is a free data retrieval call binding the contract method 0xdb747ac9.
//
// Solidity: function decodeCreation(bytes32 _payload) pure returns(uint128 _eventId, uint16 _homeTeamId, uint16 _awayTeamId, uint32 _startTimestamp, int16 _homeSpread, uint16 _totalScore, bool _createSpread, bool _createTotal)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) DecodeCreation(_payload [32]byte) (struct {
	EventId        *big.Int
	HomeTeamId     uint16
	AwayTeamId     uint16
	StartTimestamp uint32
	HomeSpread     int16
	TotalScore     uint16
	CreateSpread   bool
	CreateTotal    bool
}, error) {
	return _SportsLinkMarketFactory.Contract.DecodeCreation(&_SportsLinkMarketFactory.CallOpts, _payload)
}

// DecodeResolution is a free data retrieval call binding the contract method 0x27459179.
//
// Solidity: function decodeResolution(bytes32 _payload) pure returns(uint128 _eventId, uint8 _eventStatus, uint16 _homeScore, uint16 _awayScore)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) DecodeResolution(opts *bind.CallOpts, _payload [32]byte) (struct {
	EventId     *big.Int
	EventStatus uint8
	HomeScore   uint16
	AwayScore   uint16
}, error) {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "decodeResolution", _payload)

	outstruct := new(struct {
		EventId     *big.Int
		EventStatus uint8
		HomeScore   uint16
		AwayScore   uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EventId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EventStatus = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.HomeScore = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.AwayScore = *abi.ConvertType(out[3], new(uint16)).(*uint16)

	return *outstruct, err

}

// DecodeResolution is a free data retrieval call binding the contract method 0x27459179.
//
// Solidity: function decodeResolution(bytes32 _payload) pure returns(uint128 _eventId, uint8 _eventStatus, uint16 _homeScore, uint16 _awayScore)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) DecodeResolution(_payload [32]byte) (struct {
	EventId     *big.Int
	EventStatus uint8
	HomeScore   uint16
	AwayScore   uint16
}, error) {
	return _SportsLinkMarketFactory.Contract.DecodeResolution(&_SportsLinkMarketFactory.CallOpts, _payload)
}

// DecodeResolution is a free data retrieval call binding the contract method 0x27459179.
//
// Solidity: function decodeResolution(bytes32 _payload) pure returns(uint128 _eventId, uint8 _eventStatus, uint16 _homeScore, uint16 _awayScore)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) DecodeResolution(_payload [32]byte) (struct {
	EventId     *big.Int
	EventStatus uint8
	HomeScore   uint16
	AwayScore   uint16
}, error) {
	return _SportsLinkMarketFactory.Contract.DecodeResolution(&_SportsLinkMarketFactory.CallOpts, _payload)
}

// EncodeCreation is a free data retrieval call binding the contract method 0x91e22112.
//
// Solidity: function encodeCreation(uint128 _eventId, uint16 _homeTeamId, uint16 _awayTeamId, uint32 _startTimestamp, int16 _homeSpread, uint16 _totalScore, bool _createSpread, bool _createTotal) pure returns(bytes32 _payload)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) EncodeCreation(opts *bind.CallOpts, _eventId *big.Int, _homeTeamId uint16, _awayTeamId uint16, _startTimestamp uint32, _homeSpread int16, _totalScore uint16, _createSpread bool, _createTotal bool) ([32]byte, error) {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "encodeCreation", _eventId, _homeTeamId, _awayTeamId, _startTimestamp, _homeSpread, _totalScore, _createSpread, _createTotal)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EncodeCreation is a free data retrieval call binding the contract method 0x91e22112.
//
// Solidity: function encodeCreation(uint128 _eventId, uint16 _homeTeamId, uint16 _awayTeamId, uint32 _startTimestamp, int16 _homeSpread, uint16 _totalScore, bool _createSpread, bool _createTotal) pure returns(bytes32 _payload)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) EncodeCreation(_eventId *big.Int, _homeTeamId uint16, _awayTeamId uint16, _startTimestamp uint32, _homeSpread int16, _totalScore uint16, _createSpread bool, _createTotal bool) ([32]byte, error) {
	return _SportsLinkMarketFactory.Contract.EncodeCreation(&_SportsLinkMarketFactory.CallOpts, _eventId, _homeTeamId, _awayTeamId, _startTimestamp, _homeSpread, _totalScore, _createSpread, _createTotal)
}

// EncodeCreation is a free data retrieval call binding the contract method 0x91e22112.
//
// Solidity: function encodeCreation(uint128 _eventId, uint16 _homeTeamId, uint16 _awayTeamId, uint32 _startTimestamp, int16 _homeSpread, uint16 _totalScore, bool _createSpread, bool _createTotal) pure returns(bytes32 _payload)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) EncodeCreation(_eventId *big.Int, _homeTeamId uint16, _awayTeamId uint16, _startTimestamp uint32, _homeSpread int16, _totalScore uint16, _createSpread bool, _createTotal bool) ([32]byte, error) {
	return _SportsLinkMarketFactory.Contract.EncodeCreation(&_SportsLinkMarketFactory.CallOpts, _eventId, _homeTeamId, _awayTeamId, _startTimestamp, _homeSpread, _totalScore, _createSpread, _createTotal)
}

// EncodeResolution is a free data retrieval call binding the contract method 0x8580c504.
//
// Solidity: function encodeResolution(uint128 _eventId, uint8 _eventStatus, uint16 _homeScore, uint16 _awayScore) pure returns(bytes32 _payload)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) EncodeResolution(opts *bind.CallOpts, _eventId *big.Int, _eventStatus uint8, _homeScore uint16, _awayScore uint16) ([32]byte, error) {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "encodeResolution", _eventId, _eventStatus, _homeScore, _awayScore)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EncodeResolution is a free data retrieval call binding the contract method 0x8580c504.
//
// Solidity: function encodeResolution(uint128 _eventId, uint8 _eventStatus, uint16 _homeScore, uint16 _awayScore) pure returns(bytes32 _payload)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) EncodeResolution(_eventId *big.Int, _eventStatus uint8, _homeScore uint16, _awayScore uint16) ([32]byte, error) {
	return _SportsLinkMarketFactory.Contract.EncodeResolution(&_SportsLinkMarketFactory.CallOpts, _eventId, _eventStatus, _homeScore, _awayScore)
}

// EncodeResolution is a free data retrieval call binding the contract method 0x8580c504.
//
// Solidity: function encodeResolution(uint128 _eventId, uint8 _eventStatus, uint16 _homeScore, uint16 _awayScore) pure returns(bytes32 _payload)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) EncodeResolution(_eventId *big.Int, _eventStatus uint8, _homeScore uint16, _awayScore uint16) ([32]byte, error) {
	return _SportsLinkMarketFactory.Contract.EncodeResolution(&_SportsLinkMarketFactory.CallOpts, _eventId, _eventStatus, _homeScore, _awayScore)
}

// Events is a free data retrieval call binding the contract method 0x0b791430.
//
// Solidity: function events(uint256 ) view returns(uint256 startTime, uint256 homeScore, uint256 awayScore, uint8 status, uint256 resolutionTime, bool finalized)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) Events(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StartTime      *big.Int
	HomeScore      *big.Int
	AwayScore      *big.Int
	Status         uint8
	ResolutionTime *big.Int
	Finalized      bool
}, error) {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "events", arg0)

	outstruct := new(struct {
		StartTime      *big.Int
		HomeScore      *big.Int
		AwayScore      *big.Int
		Status         uint8
		ResolutionTime *big.Int
		Finalized      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.HomeScore = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AwayScore = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.ResolutionTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Finalized = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// Events is a free data retrieval call binding the contract method 0x0b791430.
//
// Solidity: function events(uint256 ) view returns(uint256 startTime, uint256 homeScore, uint256 awayScore, uint8 status, uint256 resolutionTime, bool finalized)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) Events(arg0 *big.Int) (struct {
	StartTime      *big.Int
	HomeScore      *big.Int
	AwayScore      *big.Int
	Status         uint8
	ResolutionTime *big.Int
	Finalized      bool
}, error) {
	return _SportsLinkMarketFactory.Contract.Events(&_SportsLinkMarketFactory.CallOpts, arg0)
}

// Events is a free data retrieval call binding the contract method 0x0b791430.
//
// Solidity: function events(uint256 ) view returns(uint256 startTime, uint256 homeScore, uint256 awayScore, uint8 status, uint256 resolutionTime, bool finalized)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) Events(arg0 *big.Int) (struct {
	StartTime      *big.Int
	HomeScore      *big.Int
	AwayScore      *big.Int
	Status         uint8
	ResolutionTime *big.Int
	Finalized      bool
}, error) {
	return _SportsLinkMarketFactory.Contract.Events(&_SportsLinkMarketFactory.CallOpts, arg0)
}

// FeePot is a free data retrieval call binding the contract method 0x4c9f66c7.
//
// Solidity: function feePot() view returns(address)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) FeePot(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "feePot")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeePot is a free data retrieval call binding the contract method 0x4c9f66c7.
//
// Solidity: function feePot() view returns(address)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) FeePot() (common.Address, error) {
	return _SportsLinkMarketFactory.Contract.FeePot(&_SportsLinkMarketFactory.CallOpts)
}

// FeePot is a free data retrieval call binding the contract method 0x4c9f66c7.
//
// Solidity: function feePot() view returns(address)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) FeePot() (common.Address, error) {
	return _SportsLinkMarketFactory.Contract.FeePot(&_SportsLinkMarketFactory.CallOpts)
}

// GetEventMarkets is a free data retrieval call binding the contract method 0xcdaac862.
//
// Solidity: function getEventMarkets(uint256 _eventId) view returns(uint256[3])
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) GetEventMarkets(opts *bind.CallOpts, _eventId *big.Int) ([3]*big.Int, error) {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "getEventMarkets", _eventId)

	if err != nil {
		return *new([3]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([3]*big.Int)).(*[3]*big.Int)

	return out0, err

}

// GetEventMarkets is a free data retrieval call binding the contract method 0xcdaac862.
//
// Solidity: function getEventMarkets(uint256 _eventId) view returns(uint256[3])
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) GetEventMarkets(_eventId *big.Int) ([3]*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.GetEventMarkets(&_SportsLinkMarketFactory.CallOpts, _eventId)
}

// GetEventMarkets is a free data retrieval call binding the contract method 0xcdaac862.
//
// Solidity: function getEventMarkets(uint256 _eventId) view returns(uint256[3])
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) GetEventMarkets(_eventId *big.Int) ([3]*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.GetEventMarkets(&_SportsLinkMarketFactory.CallOpts, _eventId)
}

// GetMarket is a free data retrieval call binding the contract method 0xeb44fdd3.
//
// Solidity: function getMarket(uint256 _id) view returns((address,address[],uint256,address,uint256,uint256,uint256,uint256))
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) GetMarket(opts *bind.CallOpts, _id *big.Int) (AbstractMarketFactoryMarket, error) {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "getMarket", _id)

	if err != nil {
		return *new(AbstractMarketFactoryMarket), err
	}

	out0 := *abi.ConvertType(out[0], new(AbstractMarketFactoryMarket)).(*AbstractMarketFactoryMarket)

	return out0, err

}

// GetMarket is a free data retrieval call binding the contract method 0xeb44fdd3.
//
// Solidity: function getMarket(uint256 _id) view returns((address,address[],uint256,address,uint256,uint256,uint256,uint256))
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) GetMarket(_id *big.Int) (AbstractMarketFactoryMarket, error) {
	return _SportsLinkMarketFactory.Contract.GetMarket(&_SportsLinkMarketFactory.CallOpts, _id)
}

// GetMarket is a free data retrieval call binding the contract method 0xeb44fdd3.
//
// Solidity: function getMarket(uint256 _id) view returns((address,address[],uint256,address,uint256,uint256,uint256,uint256))
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) GetMarket(_id *big.Int) (AbstractMarketFactoryMarket, error) {
	return _SportsLinkMarketFactory.Contract.GetMarket(&_SportsLinkMarketFactory.CallOpts, _id)
}

// GetMarketDetails is a free data retrieval call binding the contract method 0xb06c1ba3.
//
// Solidity: function getMarketDetails(uint256 _marketId) view returns((uint256,uint256,uint256,uint256,uint8,uint8,int256))
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) GetMarketDetails(opts *bind.CallOpts, _marketId *big.Int) (SportsLinkMarketFactoryMarketDetails, error) {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "getMarketDetails", _marketId)

	if err != nil {
		return *new(SportsLinkMarketFactoryMarketDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(SportsLinkMarketFactoryMarketDetails)).(*SportsLinkMarketFactoryMarketDetails)

	return out0, err

}

// GetMarketDetails is a free data retrieval call binding the contract method 0xb06c1ba3.
//
// Solidity: function getMarketDetails(uint256 _marketId) view returns((uint256,uint256,uint256,uint256,uint8,uint8,int256))
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) GetMarketDetails(_marketId *big.Int) (SportsLinkMarketFactoryMarketDetails, error) {
	return _SportsLinkMarketFactory.Contract.GetMarketDetails(&_SportsLinkMarketFactory.CallOpts, _marketId)
}

// GetMarketDetails is a free data retrieval call binding the contract method 0xb06c1ba3.
//
// Solidity: function getMarketDetails(uint256 _marketId) view returns((uint256,uint256,uint256,uint256,uint8,uint8,int256))
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) GetMarketDetails(_marketId *big.Int) (SportsLinkMarketFactoryMarketDetails, error) {
	return _SportsLinkMarketFactory.Contract.GetMarketDetails(&_SportsLinkMarketFactory.CallOpts, _marketId)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) GetOwner() (common.Address, error) {
	return _SportsLinkMarketFactory.Contract.GetOwner(&_SportsLinkMarketFactory.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) GetOwner() (common.Address, error) {
	return _SportsLinkMarketFactory.Contract.GetOwner(&_SportsLinkMarketFactory.CallOpts)
}

// IsEventRegistered is a free data retrieval call binding the contract method 0xb65ad5fd.
//
// Solidity: function isEventRegistered(uint256 _eventId) view returns(bool)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) IsEventRegistered(opts *bind.CallOpts, _eventId *big.Int) (bool, error) {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "isEventRegistered", _eventId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEventRegistered is a free data retrieval call binding the contract method 0xb65ad5fd.
//
// Solidity: function isEventRegistered(uint256 _eventId) view returns(bool)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) IsEventRegistered(_eventId *big.Int) (bool, error) {
	return _SportsLinkMarketFactory.Contract.IsEventRegistered(&_SportsLinkMarketFactory.CallOpts, _eventId)
}

// IsEventRegistered is a free data retrieval call binding the contract method 0xb65ad5fd.
//
// Solidity: function isEventRegistered(uint256 _eventId) view returns(bool)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) IsEventRegistered(_eventId *big.Int) (bool, error) {
	return _SportsLinkMarketFactory.Contract.IsEventRegistered(&_SportsLinkMarketFactory.CallOpts, _eventId)
}

// IsEventResolved is a free data retrieval call binding the contract method 0xa1794496.
//
// Solidity: function isEventResolved(uint256 _eventId) view returns(bool)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) IsEventResolved(opts *bind.CallOpts, _eventId *big.Int) (bool, error) {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "isEventResolved", _eventId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEventResolved is a free data retrieval call binding the contract method 0xa1794496.
//
// Solidity: function isEventResolved(uint256 _eventId) view returns(bool)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) IsEventResolved(_eventId *big.Int) (bool, error) {
	return _SportsLinkMarketFactory.Contract.IsEventResolved(&_SportsLinkMarketFactory.CallOpts, _eventId)
}

// IsEventResolved is a free data retrieval call binding the contract method 0xa1794496.
//
// Solidity: function isEventResolved(uint256 _eventId) view returns(bool)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) IsEventResolved(_eventId *big.Int) (bool, error) {
	return _SportsLinkMarketFactory.Contract.IsEventResolved(&_SportsLinkMarketFactory.CallOpts, _eventId)
}

// IsMarketResolved is a free data retrieval call binding the contract method 0x53ac55f5.
//
// Solidity: function isMarketResolved(uint256 _id) view returns(bool)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) IsMarketResolved(opts *bind.CallOpts, _id *big.Int) (bool, error) {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "isMarketResolved", _id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMarketResolved is a free data retrieval call binding the contract method 0x53ac55f5.
//
// Solidity: function isMarketResolved(uint256 _id) view returns(bool)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) IsMarketResolved(_id *big.Int) (bool, error) {
	return _SportsLinkMarketFactory.Contract.IsMarketResolved(&_SportsLinkMarketFactory.CallOpts, _id)
}

// IsMarketResolved is a free data retrieval call binding the contract method 0x53ac55f5.
//
// Solidity: function isMarketResolved(uint256 _id) view returns(bool)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) IsMarketResolved(_id *big.Int) (bool, error) {
	return _SportsLinkMarketFactory.Contract.IsMarketResolved(&_SportsLinkMarketFactory.CallOpts, _id)
}

// LinkNode is a free data retrieval call binding the contract method 0xd4b6838e.
//
// Solidity: function linkNode() view returns(address)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) LinkNode(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "linkNode")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LinkNode is a free data retrieval call binding the contract method 0xd4b6838e.
//
// Solidity: function linkNode() view returns(address)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) LinkNode() (common.Address, error) {
	return _SportsLinkMarketFactory.Contract.LinkNode(&_SportsLinkMarketFactory.CallOpts)
}

// LinkNode is a free data retrieval call binding the contract method 0xd4b6838e.
//
// Solidity: function linkNode() view returns(address)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) LinkNode() (common.Address, error) {
	return _SportsLinkMarketFactory.Contract.LinkNode(&_SportsLinkMarketFactory.CallOpts)
}

// ListOfEvents is a free data retrieval call binding the contract method 0x71ba2696.
//
// Solidity: function listOfEvents(uint256 ) view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) ListOfEvents(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "listOfEvents", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ListOfEvents is a free data retrieval call binding the contract method 0x71ba2696.
//
// Solidity: function listOfEvents(uint256 ) view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) ListOfEvents(arg0 *big.Int) (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.ListOfEvents(&_SportsLinkMarketFactory.CallOpts, arg0)
}

// ListOfEvents is a free data retrieval call binding the contract method 0x71ba2696.
//
// Solidity: function listOfEvents(uint256 ) view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) ListOfEvents(arg0 *big.Int) (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.ListOfEvents(&_SportsLinkMarketFactory.CallOpts, arg0)
}

// ListResolvableEvents is a free data retrieval call binding the contract method 0x4a875e0b.
//
// Solidity: function listResolvableEvents() view returns(uint256[])
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) ListResolvableEvents(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "listResolvableEvents")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ListResolvableEvents is a free data retrieval call binding the contract method 0x4a875e0b.
//
// Solidity: function listResolvableEvents() view returns(uint256[])
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) ListResolvableEvents() ([]*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.ListResolvableEvents(&_SportsLinkMarketFactory.CallOpts)
}

// ListResolvableEvents is a free data retrieval call binding the contract method 0x4a875e0b.
//
// Solidity: function listResolvableEvents() view returns(uint256[])
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) ListResolvableEvents() ([]*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.ListResolvableEvents(&_SportsLinkMarketFactory.CallOpts)
}

// ListUnresolvedMarkets is a free data retrieval call binding the contract method 0xd9113f0d.
//
// Solidity: function listUnresolvedMarkets() view returns(uint256[])
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) ListUnresolvedMarkets(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "listUnresolvedMarkets")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ListUnresolvedMarkets is a free data retrieval call binding the contract method 0xd9113f0d.
//
// Solidity: function listUnresolvedMarkets() view returns(uint256[])
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) ListUnresolvedMarkets() ([]*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.ListUnresolvedMarkets(&_SportsLinkMarketFactory.CallOpts)
}

// ListUnresolvedMarkets is a free data retrieval call binding the contract method 0xd9113f0d.
//
// Solidity: function listUnresolvedMarkets() view returns(uint256[])
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) ListUnresolvedMarkets() ([]*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.ListUnresolvedMarkets(&_SportsLinkMarketFactory.CallOpts)
}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) MarketCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "marketCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) MarketCount() (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.MarketCount(&_SportsLinkMarketFactory.CallOpts)
}

// MarketCount is a free data retrieval call binding the contract method 0xec979082.
//
// Solidity: function marketCount() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) MarketCount() (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.MarketCount(&_SportsLinkMarketFactory.CallOpts)
}

// Protocol is a free data retrieval call binding the contract method 0x8ce74426.
//
// Solidity: function protocol() view returns(address)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) Protocol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "protocol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Protocol is a free data retrieval call binding the contract method 0x8ce74426.
//
// Solidity: function protocol() view returns(address)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) Protocol() (common.Address, error) {
	return _SportsLinkMarketFactory.Contract.Protocol(&_SportsLinkMarketFactory.CallOpts)
}

// Protocol is a free data retrieval call binding the contract method 0x8ce74426.
//
// Solidity: function protocol() view returns(address)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) Protocol() (common.Address, error) {
	return _SportsLinkMarketFactory.Contract.Protocol(&_SportsLinkMarketFactory.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) ProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "protocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) ProtocolFee() (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.ProtocolFee(&_SportsLinkMarketFactory.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) ProtocolFee() (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.ProtocolFee(&_SportsLinkMarketFactory.CallOpts)
}

// ResolveMarket is a free data retrieval call binding the contract method 0x6399d03d.
//
// Solidity: function resolveMarket(uint256 ) pure returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) ResolveMarket(opts *bind.CallOpts, arg0 *big.Int) error {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "resolveMarket", arg0)

	if err != nil {
		return err
	}

	return err

}

// ResolveMarket is a free data retrieval call binding the contract method 0x6399d03d.
//
// Solidity: function resolveMarket(uint256 ) pure returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) ResolveMarket(arg0 *big.Int) error {
	return _SportsLinkMarketFactory.Contract.ResolveMarket(&_SportsLinkMarketFactory.CallOpts, arg0)
}

// ResolveMarket is a free data retrieval call binding the contract method 0x6399d03d.
//
// Solidity: function resolveMarket(uint256 ) pure returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) ResolveMarket(arg0 *big.Int) error {
	return _SportsLinkMarketFactory.Contract.ResolveMarket(&_SportsLinkMarketFactory.CallOpts, arg0)
}

// SettlementFee is a free data retrieval call binding the contract method 0x7d1d7fb8.
//
// Solidity: function settlementFee() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) SettlementFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "settlementFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SettlementFee is a free data retrieval call binding the contract method 0x7d1d7fb8.
//
// Solidity: function settlementFee() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) SettlementFee() (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.SettlementFee(&_SportsLinkMarketFactory.CallOpts)
}

// SettlementFee is a free data retrieval call binding the contract method 0x7d1d7fb8.
//
// Solidity: function settlementFee() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) SettlementFee() (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.SettlementFee(&_SportsLinkMarketFactory.CallOpts)
}

// ShareFactor is a free data retrieval call binding the contract method 0x7641ab01.
//
// Solidity: function shareFactor() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) ShareFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "shareFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ShareFactor is a free data retrieval call binding the contract method 0x7641ab01.
//
// Solidity: function shareFactor() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) ShareFactor() (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.ShareFactor(&_SportsLinkMarketFactory.CallOpts)
}

// ShareFactor is a free data retrieval call binding the contract method 0x7641ab01.
//
// Solidity: function shareFactor() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) ShareFactor() (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.ShareFactor(&_SportsLinkMarketFactory.CallOpts)
}

// SportId is a free data retrieval call binding the contract method 0xd5db3efb.
//
// Solidity: function sportId() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) SportId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "sportId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SportId is a free data retrieval call binding the contract method 0xd5db3efb.
//
// Solidity: function sportId() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) SportId() (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.SportId(&_SportsLinkMarketFactory.CallOpts)
}

// SportId is a free data retrieval call binding the contract method 0xd5db3efb.
//
// Solidity: function sportId() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) SportId() (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.SportId(&_SportsLinkMarketFactory.CallOpts)
}

// StakerFee is a free data retrieval call binding the contract method 0x4b2d9ffc.
//
// Solidity: function stakerFee() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCaller) StakerFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SportsLinkMarketFactory.contract.Call(opts, &out, "stakerFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerFee is a free data retrieval call binding the contract method 0x4b2d9ffc.
//
// Solidity: function stakerFee() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) StakerFee() (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.StakerFee(&_SportsLinkMarketFactory.CallOpts)
}

// StakerFee is a free data retrieval call binding the contract method 0x4b2d9ffc.
//
// Solidity: function stakerFee() view returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryCallerSession) StakerFee() (*big.Int, error) {
	return _SportsLinkMarketFactory.Contract.StakerFee(&_SportsLinkMarketFactory.CallOpts)
}

// BurnShares is a paid mutator transaction binding the contract method 0x35a9cdad.
//
// Solidity: function burnShares(uint256 _id, uint256 _sharesToBurn, address _receiver) returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactor) BurnShares(opts *bind.TransactOpts, _id *big.Int, _sharesToBurn *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.contract.Transact(opts, "burnShares", _id, _sharesToBurn, _receiver)
}

// BurnShares is a paid mutator transaction binding the contract method 0x35a9cdad.
//
// Solidity: function burnShares(uint256 _id, uint256 _sharesToBurn, address _receiver) returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) BurnShares(_id *big.Int, _sharesToBurn *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.BurnShares(&_SportsLinkMarketFactory.TransactOpts, _id, _sharesToBurn, _receiver)
}

// BurnShares is a paid mutator transaction binding the contract method 0x35a9cdad.
//
// Solidity: function burnShares(uint256 _id, uint256 _sharesToBurn, address _receiver) returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactorSession) BurnShares(_id *big.Int, _sharesToBurn *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.BurnShares(&_SportsLinkMarketFactory.TransactOpts, _id, _sharesToBurn, _receiver)
}

// ClaimManyWinnings is a paid mutator transaction binding the contract method 0xe5678dfa.
//
// Solidity: function claimManyWinnings(uint256[] _ids, address _receiver) returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactor) ClaimManyWinnings(opts *bind.TransactOpts, _ids []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.contract.Transact(opts, "claimManyWinnings", _ids, _receiver)
}

// ClaimManyWinnings is a paid mutator transaction binding the contract method 0xe5678dfa.
//
// Solidity: function claimManyWinnings(uint256[] _ids, address _receiver) returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) ClaimManyWinnings(_ids []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.ClaimManyWinnings(&_SportsLinkMarketFactory.TransactOpts, _ids, _receiver)
}

// ClaimManyWinnings is a paid mutator transaction binding the contract method 0xe5678dfa.
//
// Solidity: function claimManyWinnings(uint256[] _ids, address _receiver) returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactorSession) ClaimManyWinnings(_ids []*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.ClaimManyWinnings(&_SportsLinkMarketFactory.TransactOpts, _ids, _receiver)
}

// ClaimProtocolFees is a paid mutator transaction binding the contract method 0x4a7d0369.
//
// Solidity: function claimProtocolFees() returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactor) ClaimProtocolFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.contract.Transact(opts, "claimProtocolFees")
}

// ClaimProtocolFees is a paid mutator transaction binding the contract method 0x4a7d0369.
//
// Solidity: function claimProtocolFees() returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) ClaimProtocolFees() (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.ClaimProtocolFees(&_SportsLinkMarketFactory.TransactOpts)
}

// ClaimProtocolFees is a paid mutator transaction binding the contract method 0x4a7d0369.
//
// Solidity: function claimProtocolFees() returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactorSession) ClaimProtocolFees() (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.ClaimProtocolFees(&_SportsLinkMarketFactory.TransactOpts)
}

// ClaimSettlementFees is a paid mutator transaction binding the contract method 0x8e0ed193.
//
// Solidity: function claimSettlementFees(address _receiver) returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactor) ClaimSettlementFees(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.contract.Transact(opts, "claimSettlementFees", _receiver)
}

// ClaimSettlementFees is a paid mutator transaction binding the contract method 0x8e0ed193.
//
// Solidity: function claimSettlementFees(address _receiver) returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) ClaimSettlementFees(_receiver common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.ClaimSettlementFees(&_SportsLinkMarketFactory.TransactOpts, _receiver)
}

// ClaimSettlementFees is a paid mutator transaction binding the contract method 0x8e0ed193.
//
// Solidity: function claimSettlementFees(address _receiver) returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactorSession) ClaimSettlementFees(_receiver common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.ClaimSettlementFees(&_SportsLinkMarketFactory.TransactOpts, _receiver)
}

// ClaimWinnings is a paid mutator transaction binding the contract method 0x992c9079.
//
// Solidity: function claimWinnings(uint256 _id, address _receiver) returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactor) ClaimWinnings(opts *bind.TransactOpts, _id *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.contract.Transact(opts, "claimWinnings", _id, _receiver)
}

// ClaimWinnings is a paid mutator transaction binding the contract method 0x992c9079.
//
// Solidity: function claimWinnings(uint256 _id, address _receiver) returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) ClaimWinnings(_id *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.ClaimWinnings(&_SportsLinkMarketFactory.TransactOpts, _id, _receiver)
}

// ClaimWinnings is a paid mutator transaction binding the contract method 0x992c9079.
//
// Solidity: function claimWinnings(uint256 _id, address _receiver) returns(uint256)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactorSession) ClaimWinnings(_id *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.ClaimWinnings(&_SportsLinkMarketFactory.TransactOpts, _id, _receiver)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x2243118a.
//
// Solidity: function createMarket(bytes32 _payload) returns(uint256[3] _ids)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactor) CreateMarket(opts *bind.TransactOpts, _payload [32]byte) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.contract.Transact(opts, "createMarket", _payload)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x2243118a.
//
// Solidity: function createMarket(bytes32 _payload) returns(uint256[3] _ids)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) CreateMarket(_payload [32]byte) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.CreateMarket(&_SportsLinkMarketFactory.TransactOpts, _payload)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x2243118a.
//
// Solidity: function createMarket(bytes32 _payload) returns(uint256[3] _ids)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactorSession) CreateMarket(_payload [32]byte) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.CreateMarket(&_SportsLinkMarketFactory.TransactOpts, _payload)
}

// MintShares is a paid mutator transaction binding the contract method 0x221fff81.
//
// Solidity: function mintShares(uint256 _id, uint256 _shareToMint, address _receiver) returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactor) MintShares(opts *bind.TransactOpts, _id *big.Int, _shareToMint *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.contract.Transact(opts, "mintShares", _id, _shareToMint, _receiver)
}

// MintShares is a paid mutator transaction binding the contract method 0x221fff81.
//
// Solidity: function mintShares(uint256 _id, uint256 _shareToMint, address _receiver) returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) MintShares(_id *big.Int, _shareToMint *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.MintShares(&_SportsLinkMarketFactory.TransactOpts, _id, _shareToMint, _receiver)
}

// MintShares is a paid mutator transaction binding the contract method 0x221fff81.
//
// Solidity: function mintShares(uint256 _id, uint256 _shareToMint, address _receiver) returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactorSession) MintShares(_id *big.Int, _shareToMint *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.MintShares(&_SportsLinkMarketFactory.TransactOpts, _id, _shareToMint, _receiver)
}

// SetLinkNode is a paid mutator transaction binding the contract method 0xe2c30b15.
//
// Solidity: function setLinkNode(address _newLinkNode) returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactor) SetLinkNode(opts *bind.TransactOpts, _newLinkNode common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.contract.Transact(opts, "setLinkNode", _newLinkNode)
}

// SetLinkNode is a paid mutator transaction binding the contract method 0xe2c30b15.
//
// Solidity: function setLinkNode(address _newLinkNode) returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) SetLinkNode(_newLinkNode common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.SetLinkNode(&_SportsLinkMarketFactory.TransactOpts, _newLinkNode)
}

// SetLinkNode is a paid mutator transaction binding the contract method 0xe2c30b15.
//
// Solidity: function setLinkNode(address _newLinkNode) returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactorSession) SetLinkNode(_newLinkNode common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.SetLinkNode(&_SportsLinkMarketFactory.TransactOpts, _newLinkNode)
}

// SetProtocol is a paid mutator transaction binding the contract method 0x32ecabe9.
//
// Solidity: function setProtocol(address _newProtocol, bool _claimFirst) returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactor) SetProtocol(opts *bind.TransactOpts, _newProtocol common.Address, _claimFirst bool) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.contract.Transact(opts, "setProtocol", _newProtocol, _claimFirst)
}

// SetProtocol is a paid mutator transaction binding the contract method 0x32ecabe9.
//
// Solidity: function setProtocol(address _newProtocol, bool _claimFirst) returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) SetProtocol(_newProtocol common.Address, _claimFirst bool) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.SetProtocol(&_SportsLinkMarketFactory.TransactOpts, _newProtocol, _claimFirst)
}

// SetProtocol is a paid mutator transaction binding the contract method 0x32ecabe9.
//
// Solidity: function setProtocol(address _newProtocol, bool _claimFirst) returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactorSession) SetProtocol(_newProtocol common.Address, _claimFirst bool) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.SetProtocol(&_SportsLinkMarketFactory.TransactOpts, _newProtocol, _claimFirst)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 _newFee) returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactor) SetProtocolFee(opts *bind.TransactOpts, _newFee *big.Int) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.contract.Transact(opts, "setProtocolFee", _newFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 _newFee) returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) SetProtocolFee(_newFee *big.Int) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.SetProtocolFee(&_SportsLinkMarketFactory.TransactOpts, _newFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 _newFee) returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactorSession) SetProtocolFee(_newFee *big.Int) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.SetProtocolFee(&_SportsLinkMarketFactory.TransactOpts, _newFee)
}

// SetSettlementFee is a paid mutator transaction binding the contract method 0xd5da4f1d.
//
// Solidity: function setSettlementFee(uint256 _newFee) returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactor) SetSettlementFee(opts *bind.TransactOpts, _newFee *big.Int) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.contract.Transact(opts, "setSettlementFee", _newFee)
}

// SetSettlementFee is a paid mutator transaction binding the contract method 0xd5da4f1d.
//
// Solidity: function setSettlementFee(uint256 _newFee) returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) SetSettlementFee(_newFee *big.Int) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.SetSettlementFee(&_SportsLinkMarketFactory.TransactOpts, _newFee)
}

// SetSettlementFee is a paid mutator transaction binding the contract method 0xd5da4f1d.
//
// Solidity: function setSettlementFee(uint256 _newFee) returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactorSession) SetSettlementFee(_newFee *big.Int) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.SetSettlementFee(&_SportsLinkMarketFactory.TransactOpts, _newFee)
}

// SetStakerFee is a paid mutator transaction binding the contract method 0x97eef187.
//
// Solidity: function setStakerFee(uint256 _newFee) returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactor) SetStakerFee(opts *bind.TransactOpts, _newFee *big.Int) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.contract.Transact(opts, "setStakerFee", _newFee)
}

// SetStakerFee is a paid mutator transaction binding the contract method 0x97eef187.
//
// Solidity: function setStakerFee(uint256 _newFee) returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) SetStakerFee(_newFee *big.Int) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.SetStakerFee(&_SportsLinkMarketFactory.TransactOpts, _newFee)
}

// SetStakerFee is a paid mutator transaction binding the contract method 0x97eef187.
//
// Solidity: function setStakerFee(uint256 _newFee) returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactorSession) SetStakerFee(_newFee *big.Int) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.SetStakerFee(&_SportsLinkMarketFactory.TransactOpts, _newFee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.TransferOwnership(&_SportsLinkMarketFactory.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.TransferOwnership(&_SportsLinkMarketFactory.TransactOpts, _newOwner)
}

// TrustedResolveMarkets is a paid mutator transaction binding the contract method 0x575717b0.
//
// Solidity: function trustedResolveMarkets(bytes32 _payload) returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactor) TrustedResolveMarkets(opts *bind.TransactOpts, _payload [32]byte) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.contract.Transact(opts, "trustedResolveMarkets", _payload)
}

// TrustedResolveMarkets is a paid mutator transaction binding the contract method 0x575717b0.
//
// Solidity: function trustedResolveMarkets(bytes32 _payload) returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactorySession) TrustedResolveMarkets(_payload [32]byte) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.TrustedResolveMarkets(&_SportsLinkMarketFactory.TransactOpts, _payload)
}

// TrustedResolveMarkets is a paid mutator transaction binding the contract method 0x575717b0.
//
// Solidity: function trustedResolveMarkets(bytes32 _payload) returns()
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryTransactorSession) TrustedResolveMarkets(_payload [32]byte) (*types.Transaction, error) {
	return _SportsLinkMarketFactory.Contract.TrustedResolveMarkets(&_SportsLinkMarketFactory.TransactOpts, _payload)
}

// SportsLinkMarketFactoryLinkNodeChangedIterator is returned from FilterLinkNodeChanged and is used to iterate over the raw logs and unpacked data for LinkNodeChanged events raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactoryLinkNodeChangedIterator struct {
	Event *SportsLinkMarketFactoryLinkNodeChanged // Event containing the contract specifics and raw log

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
func (it *SportsLinkMarketFactoryLinkNodeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SportsLinkMarketFactoryLinkNodeChanged)
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
		it.Event = new(SportsLinkMarketFactoryLinkNodeChanged)
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
func (it *SportsLinkMarketFactoryLinkNodeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SportsLinkMarketFactoryLinkNodeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SportsLinkMarketFactoryLinkNodeChanged represents a LinkNodeChanged event raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactoryLinkNodeChanged struct {
	NewLinkNode common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLinkNodeChanged is a free log retrieval operation binding the contract event 0x6b7517523482c8d89ffbc530829d5decd506cf6dc60874b11fa26c8a53bb9fa9.
//
// Solidity: event LinkNodeChanged(address newLinkNode)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) FilterLinkNodeChanged(opts *bind.FilterOpts) (*SportsLinkMarketFactoryLinkNodeChangedIterator, error) {

	logs, sub, err := _SportsLinkMarketFactory.contract.FilterLogs(opts, "LinkNodeChanged")
	if err != nil {
		return nil, err
	}
	return &SportsLinkMarketFactoryLinkNodeChangedIterator{contract: _SportsLinkMarketFactory.contract, event: "LinkNodeChanged", logs: logs, sub: sub}, nil
}

// WatchLinkNodeChanged is a free log subscription operation binding the contract event 0x6b7517523482c8d89ffbc530829d5decd506cf6dc60874b11fa26c8a53bb9fa9.
//
// Solidity: event LinkNodeChanged(address newLinkNode)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) WatchLinkNodeChanged(opts *bind.WatchOpts, sink chan<- *SportsLinkMarketFactoryLinkNodeChanged) (event.Subscription, error) {

	logs, sub, err := _SportsLinkMarketFactory.contract.WatchLogs(opts, "LinkNodeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SportsLinkMarketFactoryLinkNodeChanged)
				if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "LinkNodeChanged", log); err != nil {
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

// ParseLinkNodeChanged is a log parse operation binding the contract event 0x6b7517523482c8d89ffbc530829d5decd506cf6dc60874b11fa26c8a53bb9fa9.
//
// Solidity: event LinkNodeChanged(address newLinkNode)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) ParseLinkNodeChanged(log types.Log) (*SportsLinkMarketFactoryLinkNodeChanged, error) {
	event := new(SportsLinkMarketFactoryLinkNodeChanged)
	if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "LinkNodeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SportsLinkMarketFactoryMarketCreatedIterator is returned from FilterMarketCreated and is used to iterate over the raw logs and unpacked data for MarketCreated events raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactoryMarketCreatedIterator struct {
	Event *SportsLinkMarketFactoryMarketCreated // Event containing the contract specifics and raw log

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
func (it *SportsLinkMarketFactoryMarketCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SportsLinkMarketFactoryMarketCreated)
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
		it.Event = new(SportsLinkMarketFactoryMarketCreated)
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
func (it *SportsLinkMarketFactoryMarketCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SportsLinkMarketFactoryMarketCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SportsLinkMarketFactoryMarketCreated represents a MarketCreated event raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactoryMarketCreated struct {
	Id                 *big.Int
	Creator            common.Address
	EndTime            *big.Int
	MarketType         uint8
	EventId            *big.Int
	HomeTeamId         *big.Int
	AwayTeamId         *big.Int
	EstimatedStartTime *big.Int
	Score              *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMarketCreated is a free log retrieval operation binding the contract event 0xafad6545e5200f9fdf4db34dfca61a9e7d72518593fd5155d11fd69c20e1555f.
//
// Solidity: event MarketCreated(uint256 id, address creator, uint256 endTime, uint8 marketType, uint256 indexed eventId, uint256 homeTeamId, uint256 awayTeamId, uint256 estimatedStartTime, int256 score)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) FilterMarketCreated(opts *bind.FilterOpts, eventId []*big.Int) (*SportsLinkMarketFactoryMarketCreatedIterator, error) {

	var eventIdRule []interface{}
	for _, eventIdItem := range eventId {
		eventIdRule = append(eventIdRule, eventIdItem)
	}

	logs, sub, err := _SportsLinkMarketFactory.contract.FilterLogs(opts, "MarketCreated", eventIdRule)
	if err != nil {
		return nil, err
	}
	return &SportsLinkMarketFactoryMarketCreatedIterator{contract: _SportsLinkMarketFactory.contract, event: "MarketCreated", logs: logs, sub: sub}, nil
}

// WatchMarketCreated is a free log subscription operation binding the contract event 0xafad6545e5200f9fdf4db34dfca61a9e7d72518593fd5155d11fd69c20e1555f.
//
// Solidity: event MarketCreated(uint256 id, address creator, uint256 endTime, uint8 marketType, uint256 indexed eventId, uint256 homeTeamId, uint256 awayTeamId, uint256 estimatedStartTime, int256 score)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) WatchMarketCreated(opts *bind.WatchOpts, sink chan<- *SportsLinkMarketFactoryMarketCreated, eventId []*big.Int) (event.Subscription, error) {

	var eventIdRule []interface{}
	for _, eventIdItem := range eventId {
		eventIdRule = append(eventIdRule, eventIdItem)
	}

	logs, sub, err := _SportsLinkMarketFactory.contract.WatchLogs(opts, "MarketCreated", eventIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SportsLinkMarketFactoryMarketCreated)
				if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "MarketCreated", log); err != nil {
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

// ParseMarketCreated is a log parse operation binding the contract event 0xafad6545e5200f9fdf4db34dfca61a9e7d72518593fd5155d11fd69c20e1555f.
//
// Solidity: event MarketCreated(uint256 id, address creator, uint256 endTime, uint8 marketType, uint256 indexed eventId, uint256 homeTeamId, uint256 awayTeamId, uint256 estimatedStartTime, int256 score)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) ParseMarketCreated(log types.Log) (*SportsLinkMarketFactoryMarketCreated, error) {
	event := new(SportsLinkMarketFactoryMarketCreated)
	if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "MarketCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SportsLinkMarketFactoryMarketResolvedIterator is returned from FilterMarketResolved and is used to iterate over the raw logs and unpacked data for MarketResolved events raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactoryMarketResolvedIterator struct {
	Event *SportsLinkMarketFactoryMarketResolved // Event containing the contract specifics and raw log

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
func (it *SportsLinkMarketFactoryMarketResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SportsLinkMarketFactoryMarketResolved)
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
		it.Event = new(SportsLinkMarketFactoryMarketResolved)
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
func (it *SportsLinkMarketFactoryMarketResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SportsLinkMarketFactoryMarketResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SportsLinkMarketFactoryMarketResolved represents a MarketResolved event raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactoryMarketResolved struct {
	Id     *big.Int
	Winner common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMarketResolved is a free log retrieval operation binding the contract event 0xc68d106ea6e4bec784925cfd91767212c71ced92adbac107dc364435321113f6.
//
// Solidity: event MarketResolved(uint256 id, address winner)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) FilterMarketResolved(opts *bind.FilterOpts) (*SportsLinkMarketFactoryMarketResolvedIterator, error) {

	logs, sub, err := _SportsLinkMarketFactory.contract.FilterLogs(opts, "MarketResolved")
	if err != nil {
		return nil, err
	}
	return &SportsLinkMarketFactoryMarketResolvedIterator{contract: _SportsLinkMarketFactory.contract, event: "MarketResolved", logs: logs, sub: sub}, nil
}

// WatchMarketResolved is a free log subscription operation binding the contract event 0xc68d106ea6e4bec784925cfd91767212c71ced92adbac107dc364435321113f6.
//
// Solidity: event MarketResolved(uint256 id, address winner)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) WatchMarketResolved(opts *bind.WatchOpts, sink chan<- *SportsLinkMarketFactoryMarketResolved) (event.Subscription, error) {

	logs, sub, err := _SportsLinkMarketFactory.contract.WatchLogs(opts, "MarketResolved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SportsLinkMarketFactoryMarketResolved)
				if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "MarketResolved", log); err != nil {
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

// ParseMarketResolved is a log parse operation binding the contract event 0xc68d106ea6e4bec784925cfd91767212c71ced92adbac107dc364435321113f6.
//
// Solidity: event MarketResolved(uint256 id, address winner)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) ParseMarketResolved(log types.Log) (*SportsLinkMarketFactoryMarketResolved, error) {
	event := new(SportsLinkMarketFactoryMarketResolved)
	if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "MarketResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SportsLinkMarketFactoryProtocolChangedIterator is returned from FilterProtocolChanged and is used to iterate over the raw logs and unpacked data for ProtocolChanged events raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactoryProtocolChangedIterator struct {
	Event *SportsLinkMarketFactoryProtocolChanged // Event containing the contract specifics and raw log

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
func (it *SportsLinkMarketFactoryProtocolChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SportsLinkMarketFactoryProtocolChanged)
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
		it.Event = new(SportsLinkMarketFactoryProtocolChanged)
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
func (it *SportsLinkMarketFactoryProtocolChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SportsLinkMarketFactoryProtocolChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SportsLinkMarketFactoryProtocolChanged represents a ProtocolChanged event raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactoryProtocolChanged struct {
	Protocol common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProtocolChanged is a free log retrieval operation binding the contract event 0x15b84596b3c567ae2998116949ae5f2d47f3055c12d9053db4d6e50f4c794dd9.
//
// Solidity: event ProtocolChanged(address protocol)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) FilterProtocolChanged(opts *bind.FilterOpts) (*SportsLinkMarketFactoryProtocolChangedIterator, error) {

	logs, sub, err := _SportsLinkMarketFactory.contract.FilterLogs(opts, "ProtocolChanged")
	if err != nil {
		return nil, err
	}
	return &SportsLinkMarketFactoryProtocolChangedIterator{contract: _SportsLinkMarketFactory.contract, event: "ProtocolChanged", logs: logs, sub: sub}, nil
}

// WatchProtocolChanged is a free log subscription operation binding the contract event 0x15b84596b3c567ae2998116949ae5f2d47f3055c12d9053db4d6e50f4c794dd9.
//
// Solidity: event ProtocolChanged(address protocol)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) WatchProtocolChanged(opts *bind.WatchOpts, sink chan<- *SportsLinkMarketFactoryProtocolChanged) (event.Subscription, error) {

	logs, sub, err := _SportsLinkMarketFactory.contract.WatchLogs(opts, "ProtocolChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SportsLinkMarketFactoryProtocolChanged)
				if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "ProtocolChanged", log); err != nil {
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

// ParseProtocolChanged is a log parse operation binding the contract event 0x15b84596b3c567ae2998116949ae5f2d47f3055c12d9053db4d6e50f4c794dd9.
//
// Solidity: event ProtocolChanged(address protocol)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) ParseProtocolChanged(log types.Log) (*SportsLinkMarketFactoryProtocolChanged, error) {
	event := new(SportsLinkMarketFactoryProtocolChanged)
	if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "ProtocolChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SportsLinkMarketFactoryProtocolFeeChangedIterator is returned from FilterProtocolFeeChanged and is used to iterate over the raw logs and unpacked data for ProtocolFeeChanged events raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactoryProtocolFeeChangedIterator struct {
	Event *SportsLinkMarketFactoryProtocolFeeChanged // Event containing the contract specifics and raw log

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
func (it *SportsLinkMarketFactoryProtocolFeeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SportsLinkMarketFactoryProtocolFeeChanged)
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
		it.Event = new(SportsLinkMarketFactoryProtocolFeeChanged)
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
func (it *SportsLinkMarketFactoryProtocolFeeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SportsLinkMarketFactoryProtocolFeeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SportsLinkMarketFactoryProtocolFeeChanged represents a ProtocolFeeChanged event raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactoryProtocolFeeChanged struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeChanged is a free log retrieval operation binding the contract event 0xada2cde3c4a561f5c23e2fdbfb223e1f0d1ec7109b9811b32644e6e974d6631f.
//
// Solidity: event ProtocolFeeChanged(uint256 fee)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) FilterProtocolFeeChanged(opts *bind.FilterOpts) (*SportsLinkMarketFactoryProtocolFeeChangedIterator, error) {

	logs, sub, err := _SportsLinkMarketFactory.contract.FilterLogs(opts, "ProtocolFeeChanged")
	if err != nil {
		return nil, err
	}
	return &SportsLinkMarketFactoryProtocolFeeChangedIterator{contract: _SportsLinkMarketFactory.contract, event: "ProtocolFeeChanged", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeChanged is a free log subscription operation binding the contract event 0xada2cde3c4a561f5c23e2fdbfb223e1f0d1ec7109b9811b32644e6e974d6631f.
//
// Solidity: event ProtocolFeeChanged(uint256 fee)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) WatchProtocolFeeChanged(opts *bind.WatchOpts, sink chan<- *SportsLinkMarketFactoryProtocolFeeChanged) (event.Subscription, error) {

	logs, sub, err := _SportsLinkMarketFactory.contract.WatchLogs(opts, "ProtocolFeeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SportsLinkMarketFactoryProtocolFeeChanged)
				if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "ProtocolFeeChanged", log); err != nil {
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

// ParseProtocolFeeChanged is a log parse operation binding the contract event 0xada2cde3c4a561f5c23e2fdbfb223e1f0d1ec7109b9811b32644e6e974d6631f.
//
// Solidity: event ProtocolFeeChanged(uint256 fee)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) ParseProtocolFeeChanged(log types.Log) (*SportsLinkMarketFactoryProtocolFeeChanged, error) {
	event := new(SportsLinkMarketFactoryProtocolFeeChanged)
	if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "ProtocolFeeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SportsLinkMarketFactoryProtocolFeeClaimedIterator is returned from FilterProtocolFeeClaimed and is used to iterate over the raw logs and unpacked data for ProtocolFeeClaimed events raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactoryProtocolFeeClaimedIterator struct {
	Event *SportsLinkMarketFactoryProtocolFeeClaimed // Event containing the contract specifics and raw log

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
func (it *SportsLinkMarketFactoryProtocolFeeClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SportsLinkMarketFactoryProtocolFeeClaimed)
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
		it.Event = new(SportsLinkMarketFactoryProtocolFeeClaimed)
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
func (it *SportsLinkMarketFactoryProtocolFeeClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SportsLinkMarketFactoryProtocolFeeClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SportsLinkMarketFactoryProtocolFeeClaimed represents a ProtocolFeeClaimed event raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactoryProtocolFeeClaimed struct {
	Protocol common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeClaimed is a free log retrieval operation binding the contract event 0x0f7f5b155b0b0ac6890709a2c7bf1b8bb3f675fff1e7840b4dd3c9acde59048b.
//
// Solidity: event ProtocolFeeClaimed(address protocol, uint256 amount)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) FilterProtocolFeeClaimed(opts *bind.FilterOpts) (*SportsLinkMarketFactoryProtocolFeeClaimedIterator, error) {

	logs, sub, err := _SportsLinkMarketFactory.contract.FilterLogs(opts, "ProtocolFeeClaimed")
	if err != nil {
		return nil, err
	}
	return &SportsLinkMarketFactoryProtocolFeeClaimedIterator{contract: _SportsLinkMarketFactory.contract, event: "ProtocolFeeClaimed", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeClaimed is a free log subscription operation binding the contract event 0x0f7f5b155b0b0ac6890709a2c7bf1b8bb3f675fff1e7840b4dd3c9acde59048b.
//
// Solidity: event ProtocolFeeClaimed(address protocol, uint256 amount)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) WatchProtocolFeeClaimed(opts *bind.WatchOpts, sink chan<- *SportsLinkMarketFactoryProtocolFeeClaimed) (event.Subscription, error) {

	logs, sub, err := _SportsLinkMarketFactory.contract.WatchLogs(opts, "ProtocolFeeClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SportsLinkMarketFactoryProtocolFeeClaimed)
				if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "ProtocolFeeClaimed", log); err != nil {
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

// ParseProtocolFeeClaimed is a log parse operation binding the contract event 0x0f7f5b155b0b0ac6890709a2c7bf1b8bb3f675fff1e7840b4dd3c9acde59048b.
//
// Solidity: event ProtocolFeeClaimed(address protocol, uint256 amount)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) ParseProtocolFeeClaimed(log types.Log) (*SportsLinkMarketFactoryProtocolFeeClaimed, error) {
	event := new(SportsLinkMarketFactoryProtocolFeeClaimed)
	if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "ProtocolFeeClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SportsLinkMarketFactorySettlementFeeChangedIterator is returned from FilterSettlementFeeChanged and is used to iterate over the raw logs and unpacked data for SettlementFeeChanged events raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactorySettlementFeeChangedIterator struct {
	Event *SportsLinkMarketFactorySettlementFeeChanged // Event containing the contract specifics and raw log

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
func (it *SportsLinkMarketFactorySettlementFeeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SportsLinkMarketFactorySettlementFeeChanged)
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
		it.Event = new(SportsLinkMarketFactorySettlementFeeChanged)
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
func (it *SportsLinkMarketFactorySettlementFeeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SportsLinkMarketFactorySettlementFeeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SportsLinkMarketFactorySettlementFeeChanged represents a SettlementFeeChanged event raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactorySettlementFeeChanged struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSettlementFeeChanged is a free log retrieval operation binding the contract event 0x92d395c429898992f8532ee7145901513e524c2085fd7fd1da39b8badcd6df31.
//
// Solidity: event SettlementFeeChanged(uint256 fee)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) FilterSettlementFeeChanged(opts *bind.FilterOpts) (*SportsLinkMarketFactorySettlementFeeChangedIterator, error) {

	logs, sub, err := _SportsLinkMarketFactory.contract.FilterLogs(opts, "SettlementFeeChanged")
	if err != nil {
		return nil, err
	}
	return &SportsLinkMarketFactorySettlementFeeChangedIterator{contract: _SportsLinkMarketFactory.contract, event: "SettlementFeeChanged", logs: logs, sub: sub}, nil
}

// WatchSettlementFeeChanged is a free log subscription operation binding the contract event 0x92d395c429898992f8532ee7145901513e524c2085fd7fd1da39b8badcd6df31.
//
// Solidity: event SettlementFeeChanged(uint256 fee)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) WatchSettlementFeeChanged(opts *bind.WatchOpts, sink chan<- *SportsLinkMarketFactorySettlementFeeChanged) (event.Subscription, error) {

	logs, sub, err := _SportsLinkMarketFactory.contract.WatchLogs(opts, "SettlementFeeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SportsLinkMarketFactorySettlementFeeChanged)
				if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "SettlementFeeChanged", log); err != nil {
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

// ParseSettlementFeeChanged is a log parse operation binding the contract event 0x92d395c429898992f8532ee7145901513e524c2085fd7fd1da39b8badcd6df31.
//
// Solidity: event SettlementFeeChanged(uint256 fee)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) ParseSettlementFeeChanged(log types.Log) (*SportsLinkMarketFactorySettlementFeeChanged, error) {
	event := new(SportsLinkMarketFactorySettlementFeeChanged)
	if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "SettlementFeeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SportsLinkMarketFactorySettlementFeeClaimedIterator is returned from FilterSettlementFeeClaimed and is used to iterate over the raw logs and unpacked data for SettlementFeeClaimed events raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactorySettlementFeeClaimedIterator struct {
	Event *SportsLinkMarketFactorySettlementFeeClaimed // Event containing the contract specifics and raw log

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
func (it *SportsLinkMarketFactorySettlementFeeClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SportsLinkMarketFactorySettlementFeeClaimed)
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
		it.Event = new(SportsLinkMarketFactorySettlementFeeClaimed)
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
func (it *SportsLinkMarketFactorySettlementFeeClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SportsLinkMarketFactorySettlementFeeClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SportsLinkMarketFactorySettlementFeeClaimed represents a SettlementFeeClaimed event raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactorySettlementFeeClaimed struct {
	SettlementAddress common.Address
	Amount            *big.Int
	Receiver          common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSettlementFeeClaimed is a free log retrieval operation binding the contract event 0xc9985ad824d943d66367ce5feea26e18979b3e1c9273742926d87e2b0d747387.
//
// Solidity: event SettlementFeeClaimed(address settlementAddress, uint256 amount, address indexed receiver)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) FilterSettlementFeeClaimed(opts *bind.FilterOpts, receiver []common.Address) (*SportsLinkMarketFactorySettlementFeeClaimedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _SportsLinkMarketFactory.contract.FilterLogs(opts, "SettlementFeeClaimed", receiverRule)
	if err != nil {
		return nil, err
	}
	return &SportsLinkMarketFactorySettlementFeeClaimedIterator{contract: _SportsLinkMarketFactory.contract, event: "SettlementFeeClaimed", logs: logs, sub: sub}, nil
}

// WatchSettlementFeeClaimed is a free log subscription operation binding the contract event 0xc9985ad824d943d66367ce5feea26e18979b3e1c9273742926d87e2b0d747387.
//
// Solidity: event SettlementFeeClaimed(address settlementAddress, uint256 amount, address indexed receiver)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) WatchSettlementFeeClaimed(opts *bind.WatchOpts, sink chan<- *SportsLinkMarketFactorySettlementFeeClaimed, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _SportsLinkMarketFactory.contract.WatchLogs(opts, "SettlementFeeClaimed", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SportsLinkMarketFactorySettlementFeeClaimed)
				if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "SettlementFeeClaimed", log); err != nil {
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

// ParseSettlementFeeClaimed is a log parse operation binding the contract event 0xc9985ad824d943d66367ce5feea26e18979b3e1c9273742926d87e2b0d747387.
//
// Solidity: event SettlementFeeClaimed(address settlementAddress, uint256 amount, address indexed receiver)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) ParseSettlementFeeClaimed(log types.Log) (*SportsLinkMarketFactorySettlementFeeClaimed, error) {
	event := new(SportsLinkMarketFactorySettlementFeeClaimed)
	if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "SettlementFeeClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SportsLinkMarketFactorySharesBurnedIterator is returned from FilterSharesBurned and is used to iterate over the raw logs and unpacked data for SharesBurned events raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactorySharesBurnedIterator struct {
	Event *SportsLinkMarketFactorySharesBurned // Event containing the contract specifics and raw log

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
func (it *SportsLinkMarketFactorySharesBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SportsLinkMarketFactorySharesBurned)
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
		it.Event = new(SportsLinkMarketFactorySharesBurned)
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
func (it *SportsLinkMarketFactorySharesBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SportsLinkMarketFactorySharesBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SportsLinkMarketFactorySharesBurned represents a SharesBurned event raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactorySharesBurned struct {
	Id       *big.Int
	Amount   *big.Int
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSharesBurned is a free log retrieval operation binding the contract event 0xb6fdb729b2ed801daf629f0ab713e4a7a73619505790f6f27fd92d6f2c9688d7.
//
// Solidity: event SharesBurned(uint256 id, uint256 amount, address receiver)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) FilterSharesBurned(opts *bind.FilterOpts) (*SportsLinkMarketFactorySharesBurnedIterator, error) {

	logs, sub, err := _SportsLinkMarketFactory.contract.FilterLogs(opts, "SharesBurned")
	if err != nil {
		return nil, err
	}
	return &SportsLinkMarketFactorySharesBurnedIterator{contract: _SportsLinkMarketFactory.contract, event: "SharesBurned", logs: logs, sub: sub}, nil
}

// WatchSharesBurned is a free log subscription operation binding the contract event 0xb6fdb729b2ed801daf629f0ab713e4a7a73619505790f6f27fd92d6f2c9688d7.
//
// Solidity: event SharesBurned(uint256 id, uint256 amount, address receiver)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) WatchSharesBurned(opts *bind.WatchOpts, sink chan<- *SportsLinkMarketFactorySharesBurned) (event.Subscription, error) {

	logs, sub, err := _SportsLinkMarketFactory.contract.WatchLogs(opts, "SharesBurned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SportsLinkMarketFactorySharesBurned)
				if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "SharesBurned", log); err != nil {
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

// ParseSharesBurned is a log parse operation binding the contract event 0xb6fdb729b2ed801daf629f0ab713e4a7a73619505790f6f27fd92d6f2c9688d7.
//
// Solidity: event SharesBurned(uint256 id, uint256 amount, address receiver)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) ParseSharesBurned(log types.Log) (*SportsLinkMarketFactorySharesBurned, error) {
	event := new(SportsLinkMarketFactorySharesBurned)
	if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "SharesBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SportsLinkMarketFactorySharesMintedIterator is returned from FilterSharesMinted and is used to iterate over the raw logs and unpacked data for SharesMinted events raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactorySharesMintedIterator struct {
	Event *SportsLinkMarketFactorySharesMinted // Event containing the contract specifics and raw log

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
func (it *SportsLinkMarketFactorySharesMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SportsLinkMarketFactorySharesMinted)
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
		it.Event = new(SportsLinkMarketFactorySharesMinted)
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
func (it *SportsLinkMarketFactorySharesMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SportsLinkMarketFactorySharesMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SportsLinkMarketFactorySharesMinted represents a SharesMinted event raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactorySharesMinted struct {
	Id       *big.Int
	Amount   *big.Int
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSharesMinted is a free log retrieval operation binding the contract event 0xd81c0442e10068a9818f3aa093c9ccb804584690df572d7df3da2d892a6973f2.
//
// Solidity: event SharesMinted(uint256 id, uint256 amount, address receiver)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) FilterSharesMinted(opts *bind.FilterOpts) (*SportsLinkMarketFactorySharesMintedIterator, error) {

	logs, sub, err := _SportsLinkMarketFactory.contract.FilterLogs(opts, "SharesMinted")
	if err != nil {
		return nil, err
	}
	return &SportsLinkMarketFactorySharesMintedIterator{contract: _SportsLinkMarketFactory.contract, event: "SharesMinted", logs: logs, sub: sub}, nil
}

// WatchSharesMinted is a free log subscription operation binding the contract event 0xd81c0442e10068a9818f3aa093c9ccb804584690df572d7df3da2d892a6973f2.
//
// Solidity: event SharesMinted(uint256 id, uint256 amount, address receiver)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) WatchSharesMinted(opts *bind.WatchOpts, sink chan<- *SportsLinkMarketFactorySharesMinted) (event.Subscription, error) {

	logs, sub, err := _SportsLinkMarketFactory.contract.WatchLogs(opts, "SharesMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SportsLinkMarketFactorySharesMinted)
				if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "SharesMinted", log); err != nil {
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

// ParseSharesMinted is a log parse operation binding the contract event 0xd81c0442e10068a9818f3aa093c9ccb804584690df572d7df3da2d892a6973f2.
//
// Solidity: event SharesMinted(uint256 id, uint256 amount, address receiver)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) ParseSharesMinted(log types.Log) (*SportsLinkMarketFactorySharesMinted, error) {
	event := new(SportsLinkMarketFactorySharesMinted)
	if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "SharesMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SportsLinkMarketFactoryStakerFeeChangedIterator is returned from FilterStakerFeeChanged and is used to iterate over the raw logs and unpacked data for StakerFeeChanged events raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactoryStakerFeeChangedIterator struct {
	Event *SportsLinkMarketFactoryStakerFeeChanged // Event containing the contract specifics and raw log

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
func (it *SportsLinkMarketFactoryStakerFeeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SportsLinkMarketFactoryStakerFeeChanged)
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
		it.Event = new(SportsLinkMarketFactoryStakerFeeChanged)
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
func (it *SportsLinkMarketFactoryStakerFeeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SportsLinkMarketFactoryStakerFeeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SportsLinkMarketFactoryStakerFeeChanged represents a StakerFeeChanged event raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactoryStakerFeeChanged struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStakerFeeChanged is a free log retrieval operation binding the contract event 0xcc4df50442ac32f0142ba4853f617661e0823be0e92148e7e5f36ce56c139825.
//
// Solidity: event StakerFeeChanged(uint256 fee)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) FilterStakerFeeChanged(opts *bind.FilterOpts) (*SportsLinkMarketFactoryStakerFeeChangedIterator, error) {

	logs, sub, err := _SportsLinkMarketFactory.contract.FilterLogs(opts, "StakerFeeChanged")
	if err != nil {
		return nil, err
	}
	return &SportsLinkMarketFactoryStakerFeeChangedIterator{contract: _SportsLinkMarketFactory.contract, event: "StakerFeeChanged", logs: logs, sub: sub}, nil
}

// WatchStakerFeeChanged is a free log subscription operation binding the contract event 0xcc4df50442ac32f0142ba4853f617661e0823be0e92148e7e5f36ce56c139825.
//
// Solidity: event StakerFeeChanged(uint256 fee)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) WatchStakerFeeChanged(opts *bind.WatchOpts, sink chan<- *SportsLinkMarketFactoryStakerFeeChanged) (event.Subscription, error) {

	logs, sub, err := _SportsLinkMarketFactory.contract.WatchLogs(opts, "StakerFeeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SportsLinkMarketFactoryStakerFeeChanged)
				if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "StakerFeeChanged", log); err != nil {
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

// ParseStakerFeeChanged is a log parse operation binding the contract event 0xcc4df50442ac32f0142ba4853f617661e0823be0e92148e7e5f36ce56c139825.
//
// Solidity: event StakerFeeChanged(uint256 fee)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) ParseStakerFeeChanged(log types.Log) (*SportsLinkMarketFactoryStakerFeeChanged, error) {
	event := new(SportsLinkMarketFactoryStakerFeeChanged)
	if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "StakerFeeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SportsLinkMarketFactoryWinningsClaimedIterator is returned from FilterWinningsClaimed and is used to iterate over the raw logs and unpacked data for WinningsClaimed events raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactoryWinningsClaimedIterator struct {
	Event *SportsLinkMarketFactoryWinningsClaimed // Event containing the contract specifics and raw log

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
func (it *SportsLinkMarketFactoryWinningsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SportsLinkMarketFactoryWinningsClaimed)
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
		it.Event = new(SportsLinkMarketFactoryWinningsClaimed)
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
func (it *SportsLinkMarketFactoryWinningsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SportsLinkMarketFactoryWinningsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SportsLinkMarketFactoryWinningsClaimed represents a WinningsClaimed event raised by the SportsLinkMarketFactory contract.
type SportsLinkMarketFactoryWinningsClaimed struct {
	Id             *big.Int
	WinningOutcome common.Address
	Amount         *big.Int
	SettlementFee  *big.Int
	Payout         *big.Int
	Receiver       common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWinningsClaimed is a free log retrieval operation binding the contract event 0xe67bd0100cd3289557430d36098901ba18161e6279c9711d8650b8af10552104.
//
// Solidity: event WinningsClaimed(uint256 id, address winningOutcome, uint256 amount, uint256 settlementFee, uint256 payout, address indexed receiver)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) FilterWinningsClaimed(opts *bind.FilterOpts, receiver []common.Address) (*SportsLinkMarketFactoryWinningsClaimedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _SportsLinkMarketFactory.contract.FilterLogs(opts, "WinningsClaimed", receiverRule)
	if err != nil {
		return nil, err
	}
	return &SportsLinkMarketFactoryWinningsClaimedIterator{contract: _SportsLinkMarketFactory.contract, event: "WinningsClaimed", logs: logs, sub: sub}, nil
}

// WatchWinningsClaimed is a free log subscription operation binding the contract event 0xe67bd0100cd3289557430d36098901ba18161e6279c9711d8650b8af10552104.
//
// Solidity: event WinningsClaimed(uint256 id, address winningOutcome, uint256 amount, uint256 settlementFee, uint256 payout, address indexed receiver)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) WatchWinningsClaimed(opts *bind.WatchOpts, sink chan<- *SportsLinkMarketFactoryWinningsClaimed, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _SportsLinkMarketFactory.contract.WatchLogs(opts, "WinningsClaimed", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SportsLinkMarketFactoryWinningsClaimed)
				if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "WinningsClaimed", log); err != nil {
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

// ParseWinningsClaimed is a log parse operation binding the contract event 0xe67bd0100cd3289557430d36098901ba18161e6279c9711d8650b8af10552104.
//
// Solidity: event WinningsClaimed(uint256 id, address winningOutcome, uint256 amount, uint256 settlementFee, uint256 payout, address indexed receiver)
func (_SportsLinkMarketFactory *SportsLinkMarketFactoryFilterer) ParseWinningsClaimed(log types.Log) (*SportsLinkMarketFactoryWinningsClaimed, error) {
	event := new(SportsLinkMarketFactoryWinningsClaimed)
	if err := _SportsLinkMarketFactory.contract.UnpackLog(event, "WinningsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
