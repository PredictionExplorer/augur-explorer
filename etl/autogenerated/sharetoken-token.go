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

// ShareTokenABI is the input ABI used to generate the binding from.
const ShareTokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_operatorApprovals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_supplys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"}],\"name\":\"assertBalances\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"augur\",\"outputs\":[{\"internalType\":\"contractIAugur\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outcome\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOfMarketOutcome\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"buyCompleteSets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_longOutcome\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_longRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_shortRecipient\",\"type\":\"address\"}],\"name\":\"buyCompleteSetsForTrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"calculateCreatorFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outcome\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numberOfShares\",\"type\":\"uint256\"}],\"name\":\"calculateProceeds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"calculateReportingFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cash\",\"outputs\":[{\"internalType\":\"contractICash\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_shareHolder\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_fingerprint\",\"type\":\"bytes32\"}],\"name\":\"claimTradingProceeds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_outcomeFees\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outcome\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numberOfShares\",\"type\":\"uint256\"}],\"name\":\"divideUpWinnings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_proceeds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_shareHolderShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_creatorShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reporterShare\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getMarket\",\"outputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getOutcome\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outcome\",\"type\":\"uint256\"}],\"name\":\"getTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_outcomes\",\"type\":\"uint256[]\"}],\"name\":\"getTokenIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTypeName\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIAugur\",\"name\":\"_augur\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_numOutcomes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numTicks\",\"type\":\"uint256\"}],\"name\":\"initializeMarket\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_outcomes\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"lowestBalanceOfMarketOutcomes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"publicBuyCompleteSets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"publicSellCompleteSets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_creatorFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reportingFee\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_fingerprint\",\"type\":\"bytes32\"}],\"name\":\"sellCompleteSets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_creatorFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reportingFee\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outcome\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_shortParticipant\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_longParticipant\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_shortRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_longRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sourceAccount\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_fingerprint\",\"type\":\"bytes32\"}],\"name\":\"sellCompleteSetsForTrade\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_creatorFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reportingFee\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outcome\",\"type\":\"uint256\"}],\"name\":\"totalSupplyForMarketOutcome\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"unpackTokenId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outcome\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_values\",\"type\":\"uint256[]\"}],\"name\":\"unsafeBatchTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"unsafeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ShareToken is an auto generated Go binding around an Ethereum contract.
type ShareToken struct {
	ShareTokenCaller     // Read-only binding to the contract
	ShareTokenTransactor // Write-only binding to the contract
	ShareTokenFilterer   // Log filterer for contract events
}

// ShareTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ShareTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShareTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ShareTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShareTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ShareTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShareTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ShareTokenSession struct {
	Contract     *ShareToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ShareTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ShareTokenCallerSession struct {
	Contract *ShareTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ShareTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ShareTokenTransactorSession struct {
	Contract     *ShareTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ShareTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ShareTokenRaw struct {
	Contract *ShareToken // Generic contract binding to access the raw methods on
}

// ShareTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ShareTokenCallerRaw struct {
	Contract *ShareTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ShareTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ShareTokenTransactorRaw struct {
	Contract *ShareTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShareToken creates a new instance of ShareToken, bound to a specific deployed contract.
func NewShareToken(address common.Address, backend bind.ContractBackend) (*ShareToken, error) {
	contract, err := bindShareToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ShareToken{ShareTokenCaller: ShareTokenCaller{contract: contract}, ShareTokenTransactor: ShareTokenTransactor{contract: contract}, ShareTokenFilterer: ShareTokenFilterer{contract: contract}}, nil
}

// NewShareTokenCaller creates a new read-only instance of ShareToken, bound to a specific deployed contract.
func NewShareTokenCaller(address common.Address, caller bind.ContractCaller) (*ShareTokenCaller, error) {
	contract, err := bindShareToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShareTokenCaller{contract: contract}, nil
}

// NewShareTokenTransactor creates a new write-only instance of ShareToken, bound to a specific deployed contract.
func NewShareTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ShareTokenTransactor, error) {
	contract, err := bindShareToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShareTokenTransactor{contract: contract}, nil
}

// NewShareTokenFilterer creates a new log filterer instance of ShareToken, bound to a specific deployed contract.
func NewShareTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ShareTokenFilterer, error) {
	contract, err := bindShareToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShareTokenFilterer{contract: contract}, nil
}

// bindShareToken binds a generic wrapper to an already deployed contract.
func bindShareToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ShareTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShareToken *ShareTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ShareToken.Contract.ShareTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShareToken *ShareTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShareToken.Contract.ShareTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShareToken *ShareTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShareToken.Contract.ShareTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShareToken *ShareTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ShareToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShareToken *ShareTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShareToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShareToken *ShareTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShareToken.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0xfc25a4da.
//
// Solidity: function _balances(uint256 , address ) view returns(uint256)
func (_ShareToken *ShareTokenCaller) Balances(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "_balances", arg0, arg1)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0xfc25a4da.
//
// Solidity: function _balances(uint256 , address ) view returns(uint256)
func (_ShareToken *ShareTokenSession) Balances(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _ShareToken.Contract.Balances(&_ShareToken.CallOpts, arg0, arg1)
}

// Balances is a free data retrieval call binding the contract method 0xfc25a4da.
//
// Solidity: function _balances(uint256 , address ) view returns(uint256)
func (_ShareToken *ShareTokenCallerSession) Balances(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _ShareToken.Contract.Balances(&_ShareToken.CallOpts, arg0, arg1)
}

// OperatorApprovals is a free data retrieval call binding the contract method 0xedc3bc3f.
//
// Solidity: function _operatorApprovals(address , address ) view returns(bool)
func (_ShareToken *ShareTokenCaller) OperatorApprovals(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "_operatorApprovals", arg0, arg1)
	return *ret0, err
}

// OperatorApprovals is a free data retrieval call binding the contract method 0xedc3bc3f.
//
// Solidity: function _operatorApprovals(address , address ) view returns(bool)
func (_ShareToken *ShareTokenSession) OperatorApprovals(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _ShareToken.Contract.OperatorApprovals(&_ShareToken.CallOpts, arg0, arg1)
}

// OperatorApprovals is a free data retrieval call binding the contract method 0xedc3bc3f.
//
// Solidity: function _operatorApprovals(address , address ) view returns(bool)
func (_ShareToken *ShareTokenCallerSession) OperatorApprovals(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _ShareToken.Contract.OperatorApprovals(&_ShareToken.CallOpts, arg0, arg1)
}

// Supplys is a free data retrieval call binding the contract method 0xf32ef051.
//
// Solidity: function _supplys(uint256 ) view returns(uint256)
func (_ShareToken *ShareTokenCaller) Supplys(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "_supplys", arg0)
	return *ret0, err
}

// Supplys is a free data retrieval call binding the contract method 0xf32ef051.
//
// Solidity: function _supplys(uint256 ) view returns(uint256)
func (_ShareToken *ShareTokenSession) Supplys(arg0 *big.Int) (*big.Int, error) {
	return _ShareToken.Contract.Supplys(&_ShareToken.CallOpts, arg0)
}

// Supplys is a free data retrieval call binding the contract method 0xf32ef051.
//
// Solidity: function _supplys(uint256 ) view returns(uint256)
func (_ShareToken *ShareTokenCallerSession) Supplys(arg0 *big.Int) (*big.Int, error) {
	return _ShareToken.Contract.Supplys(&_ShareToken.CallOpts, arg0)
}

// AssertBalances is a free data retrieval call binding the contract method 0x4db3b9e3.
//
// Solidity: function assertBalances(address _market) view returns()
func (_ShareToken *ShareTokenCaller) AssertBalances(opts *bind.CallOpts, _market common.Address) error {
	var ()
	out := &[]interface{}{}
	err := _ShareToken.contract.Call(opts, out, "assertBalances", _market)
	return err
}

// AssertBalances is a free data retrieval call binding the contract method 0x4db3b9e3.
//
// Solidity: function assertBalances(address _market) view returns()
func (_ShareToken *ShareTokenSession) AssertBalances(_market common.Address) error {
	return _ShareToken.Contract.AssertBalances(&_ShareToken.CallOpts, _market)
}

// AssertBalances is a free data retrieval call binding the contract method 0x4db3b9e3.
//
// Solidity: function assertBalances(address _market) view returns()
func (_ShareToken *ShareTokenCallerSession) AssertBalances(_market common.Address) error {
	return _ShareToken.Contract.AssertBalances(&_ShareToken.CallOpts, _market)
}

// Augur is a free data retrieval call binding the contract method 0x7a0d8f8a.
//
// Solidity: function augur() view returns(address)
func (_ShareToken *ShareTokenCaller) Augur(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "augur")
	return *ret0, err
}

// Augur is a free data retrieval call binding the contract method 0x7a0d8f8a.
//
// Solidity: function augur() view returns(address)
func (_ShareToken *ShareTokenSession) Augur() (common.Address, error) {
	return _ShareToken.Contract.Augur(&_ShareToken.CallOpts)
}

// Augur is a free data retrieval call binding the contract method 0x7a0d8f8a.
//
// Solidity: function augur() view returns(address)
func (_ShareToken *ShareTokenCallerSession) Augur() (common.Address, error) {
	return _ShareToken.Contract.Augur(&_ShareToken.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ShareToken *ShareTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "balanceOf", account, id)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ShareToken *ShareTokenSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ShareToken.Contract.BalanceOf(&_ShareToken.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ShareToken *ShareTokenCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ShareToken.Contract.BalanceOf(&_ShareToken.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ShareToken *ShareTokenCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "balanceOfBatch", accounts, ids)
	return *ret0, err
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ShareToken *ShareTokenSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ShareToken.Contract.BalanceOfBatch(&_ShareToken.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ShareToken *ShareTokenCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ShareToken.Contract.BalanceOfBatch(&_ShareToken.CallOpts, accounts, ids)
}

// BalanceOfMarketOutcome is a free data retrieval call binding the contract method 0xd64c62f3.
//
// Solidity: function balanceOfMarketOutcome(address _market, uint256 _outcome, address _account) view returns(uint256)
func (_ShareToken *ShareTokenCaller) BalanceOfMarketOutcome(opts *bind.CallOpts, _market common.Address, _outcome *big.Int, _account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "balanceOfMarketOutcome", _market, _outcome, _account)
	return *ret0, err
}

// BalanceOfMarketOutcome is a free data retrieval call binding the contract method 0xd64c62f3.
//
// Solidity: function balanceOfMarketOutcome(address _market, uint256 _outcome, address _account) view returns(uint256)
func (_ShareToken *ShareTokenSession) BalanceOfMarketOutcome(_market common.Address, _outcome *big.Int, _account common.Address) (*big.Int, error) {
	return _ShareToken.Contract.BalanceOfMarketOutcome(&_ShareToken.CallOpts, _market, _outcome, _account)
}

// BalanceOfMarketOutcome is a free data retrieval call binding the contract method 0xd64c62f3.
//
// Solidity: function balanceOfMarketOutcome(address _market, uint256 _outcome, address _account) view returns(uint256)
func (_ShareToken *ShareTokenCallerSession) BalanceOfMarketOutcome(_market common.Address, _outcome *big.Int, _account common.Address) (*big.Int, error) {
	return _ShareToken.Contract.BalanceOfMarketOutcome(&_ShareToken.CallOpts, _market, _outcome, _account)
}

// CalculateCreatorFee is a free data retrieval call binding the contract method 0x06065ced.
//
// Solidity: function calculateCreatorFee(address _market, uint256 _amount) view returns(uint256)
func (_ShareToken *ShareTokenCaller) CalculateCreatorFee(opts *bind.CallOpts, _market common.Address, _amount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "calculateCreatorFee", _market, _amount)
	return *ret0, err
}

// CalculateCreatorFee is a free data retrieval call binding the contract method 0x06065ced.
//
// Solidity: function calculateCreatorFee(address _market, uint256 _amount) view returns(uint256)
func (_ShareToken *ShareTokenSession) CalculateCreatorFee(_market common.Address, _amount *big.Int) (*big.Int, error) {
	return _ShareToken.Contract.CalculateCreatorFee(&_ShareToken.CallOpts, _market, _amount)
}

// CalculateCreatorFee is a free data retrieval call binding the contract method 0x06065ced.
//
// Solidity: function calculateCreatorFee(address _market, uint256 _amount) view returns(uint256)
func (_ShareToken *ShareTokenCallerSession) CalculateCreatorFee(_market common.Address, _amount *big.Int) (*big.Int, error) {
	return _ShareToken.Contract.CalculateCreatorFee(&_ShareToken.CallOpts, _market, _amount)
}

// CalculateProceeds is a free data retrieval call binding the contract method 0xf2dc8266.
//
// Solidity: function calculateProceeds(address _market, uint256 _outcome, uint256 _numberOfShares) view returns(uint256)
func (_ShareToken *ShareTokenCaller) CalculateProceeds(opts *bind.CallOpts, _market common.Address, _outcome *big.Int, _numberOfShares *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "calculateProceeds", _market, _outcome, _numberOfShares)
	return *ret0, err
}

// CalculateProceeds is a free data retrieval call binding the contract method 0xf2dc8266.
//
// Solidity: function calculateProceeds(address _market, uint256 _outcome, uint256 _numberOfShares) view returns(uint256)
func (_ShareToken *ShareTokenSession) CalculateProceeds(_market common.Address, _outcome *big.Int, _numberOfShares *big.Int) (*big.Int, error) {
	return _ShareToken.Contract.CalculateProceeds(&_ShareToken.CallOpts, _market, _outcome, _numberOfShares)
}

// CalculateProceeds is a free data retrieval call binding the contract method 0xf2dc8266.
//
// Solidity: function calculateProceeds(address _market, uint256 _outcome, uint256 _numberOfShares) view returns(uint256)
func (_ShareToken *ShareTokenCallerSession) CalculateProceeds(_market common.Address, _outcome *big.Int, _numberOfShares *big.Int) (*big.Int, error) {
	return _ShareToken.Contract.CalculateProceeds(&_ShareToken.CallOpts, _market, _outcome, _numberOfShares)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_ShareToken *ShareTokenCaller) Cash(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "cash")
	return *ret0, err
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_ShareToken *ShareTokenSession) Cash() (common.Address, error) {
	return _ShareToken.Contract.Cash(&_ShareToken.CallOpts)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_ShareToken *ShareTokenCallerSession) Cash() (common.Address, error) {
	return _ShareToken.Contract.Cash(&_ShareToken.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() view returns(bool)
func (_ShareToken *ShareTokenCaller) GetInitialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "getInitialized")
	return *ret0, err
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() view returns(bool)
func (_ShareToken *ShareTokenSession) GetInitialized() (bool, error) {
	return _ShareToken.Contract.GetInitialized(&_ShareToken.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() view returns(bool)
func (_ShareToken *ShareTokenCallerSession) GetInitialized() (bool, error) {
	return _ShareToken.Contract.GetInitialized(&_ShareToken.CallOpts)
}

// GetMarket is a free data retrieval call binding the contract method 0xeb44fdd3.
//
// Solidity: function getMarket(uint256 _tokenId) view returns(address)
func (_ShareToken *ShareTokenCaller) GetMarket(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "getMarket", _tokenId)
	return *ret0, err
}

// GetMarket is a free data retrieval call binding the contract method 0xeb44fdd3.
//
// Solidity: function getMarket(uint256 _tokenId) view returns(address)
func (_ShareToken *ShareTokenSession) GetMarket(_tokenId *big.Int) (common.Address, error) {
	return _ShareToken.Contract.GetMarket(&_ShareToken.CallOpts, _tokenId)
}

// GetMarket is a free data retrieval call binding the contract method 0xeb44fdd3.
//
// Solidity: function getMarket(uint256 _tokenId) view returns(address)
func (_ShareToken *ShareTokenCallerSession) GetMarket(_tokenId *big.Int) (common.Address, error) {
	return _ShareToken.Contract.GetMarket(&_ShareToken.CallOpts, _tokenId)
}

// GetOutcome is a free data retrieval call binding the contract method 0xcd579335.
//
// Solidity: function getOutcome(uint256 _tokenId) view returns(uint256)
func (_ShareToken *ShareTokenCaller) GetOutcome(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "getOutcome", _tokenId)
	return *ret0, err
}

// GetOutcome is a free data retrieval call binding the contract method 0xcd579335.
//
// Solidity: function getOutcome(uint256 _tokenId) view returns(uint256)
func (_ShareToken *ShareTokenSession) GetOutcome(_tokenId *big.Int) (*big.Int, error) {
	return _ShareToken.Contract.GetOutcome(&_ShareToken.CallOpts, _tokenId)
}

// GetOutcome is a free data retrieval call binding the contract method 0xcd579335.
//
// Solidity: function getOutcome(uint256 _tokenId) view returns(uint256)
func (_ShareToken *ShareTokenCallerSession) GetOutcome(_tokenId *big.Int) (*big.Int, error) {
	return _ShareToken.Contract.GetOutcome(&_ShareToken.CallOpts, _tokenId)
}

// GetTokenId is a free data retrieval call binding the contract method 0x29f99b9f.
//
// Solidity: function getTokenId(address _market, uint256 _outcome) pure returns(uint256 _tokenId)
func (_ShareToken *ShareTokenCaller) GetTokenId(opts *bind.CallOpts, _market common.Address, _outcome *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "getTokenId", _market, _outcome)
	return *ret0, err
}

// GetTokenId is a free data retrieval call binding the contract method 0x29f99b9f.
//
// Solidity: function getTokenId(address _market, uint256 _outcome) pure returns(uint256 _tokenId)
func (_ShareToken *ShareTokenSession) GetTokenId(_market common.Address, _outcome *big.Int) (*big.Int, error) {
	return _ShareToken.Contract.GetTokenId(&_ShareToken.CallOpts, _market, _outcome)
}

// GetTokenId is a free data retrieval call binding the contract method 0x29f99b9f.
//
// Solidity: function getTokenId(address _market, uint256 _outcome) pure returns(uint256 _tokenId)
func (_ShareToken *ShareTokenCallerSession) GetTokenId(_market common.Address, _outcome *big.Int) (*big.Int, error) {
	return _ShareToken.Contract.GetTokenId(&_ShareToken.CallOpts, _market, _outcome)
}

// GetTokenIds is a free data retrieval call binding the contract method 0xca3a1e69.
//
// Solidity: function getTokenIds(address _market, uint256[] _outcomes) pure returns(uint256[] _tokenIds)
func (_ShareToken *ShareTokenCaller) GetTokenIds(opts *bind.CallOpts, _market common.Address, _outcomes []*big.Int) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "getTokenIds", _market, _outcomes)
	return *ret0, err
}

// GetTokenIds is a free data retrieval call binding the contract method 0xca3a1e69.
//
// Solidity: function getTokenIds(address _market, uint256[] _outcomes) pure returns(uint256[] _tokenIds)
func (_ShareToken *ShareTokenSession) GetTokenIds(_market common.Address, _outcomes []*big.Int) ([]*big.Int, error) {
	return _ShareToken.Contract.GetTokenIds(&_ShareToken.CallOpts, _market, _outcomes)
}

// GetTokenIds is a free data retrieval call binding the contract method 0xca3a1e69.
//
// Solidity: function getTokenIds(address _market, uint256[] _outcomes) pure returns(uint256[] _tokenIds)
func (_ShareToken *ShareTokenCallerSession) GetTokenIds(_market common.Address, _outcomes []*big.Int) ([]*big.Int, error) {
	return _ShareToken.Contract.GetTokenIds(&_ShareToken.CallOpts, _market, _outcomes)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() view returns(bytes32)
func (_ShareToken *ShareTokenCaller) GetTypeName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "getTypeName")
	return *ret0, err
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() view returns(bytes32)
func (_ShareToken *ShareTokenSession) GetTypeName() ([32]byte, error) {
	return _ShareToken.Contract.GetTypeName(&_ShareToken.CallOpts)
}

// GetTypeName is a free data retrieval call binding the contract method 0xdb0a087c.
//
// Solidity: function getTypeName() view returns(bytes32)
func (_ShareToken *ShareTokenCallerSession) GetTypeName() ([32]byte, error) {
	return _ShareToken.Contract.GetTypeName(&_ShareToken.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ShareToken *ShareTokenCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "isApprovedForAll", account, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ShareToken *ShareTokenSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ShareToken.Contract.IsApprovedForAll(&_ShareToken.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ShareToken *ShareTokenCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ShareToken.Contract.IsApprovedForAll(&_ShareToken.CallOpts, account, operator)
}

// LowestBalanceOfMarketOutcomes is a free data retrieval call binding the contract method 0xf8a2167b.
//
// Solidity: function lowestBalanceOfMarketOutcomes(address _market, uint256[] _outcomes, address _account) view returns(uint256)
func (_ShareToken *ShareTokenCaller) LowestBalanceOfMarketOutcomes(opts *bind.CallOpts, _market common.Address, _outcomes []*big.Int, _account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "lowestBalanceOfMarketOutcomes", _market, _outcomes, _account)
	return *ret0, err
}

// LowestBalanceOfMarketOutcomes is a free data retrieval call binding the contract method 0xf8a2167b.
//
// Solidity: function lowestBalanceOfMarketOutcomes(address _market, uint256[] _outcomes, address _account) view returns(uint256)
func (_ShareToken *ShareTokenSession) LowestBalanceOfMarketOutcomes(_market common.Address, _outcomes []*big.Int, _account common.Address) (*big.Int, error) {
	return _ShareToken.Contract.LowestBalanceOfMarketOutcomes(&_ShareToken.CallOpts, _market, _outcomes, _account)
}

// LowestBalanceOfMarketOutcomes is a free data retrieval call binding the contract method 0xf8a2167b.
//
// Solidity: function lowestBalanceOfMarketOutcomes(address _market, uint256[] _outcomes, address _account) view returns(uint256)
func (_ShareToken *ShareTokenCallerSession) LowestBalanceOfMarketOutcomes(_market common.Address, _outcomes []*big.Int, _account common.Address) (*big.Int, error) {
	return _ShareToken.Contract.LowestBalanceOfMarketOutcomes(&_ShareToken.CallOpts, _market, _outcomes, _account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ShareToken *ShareTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ShareToken *ShareTokenSession) Name() (string, error) {
	return _ShareToken.Contract.Name(&_ShareToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ShareToken *ShareTokenCallerSession) Name() (string, error) {
	return _ShareToken.Contract.Name(&_ShareToken.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ShareToken *ShareTokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ShareToken *ShareTokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ShareToken.Contract.SupportsInterface(&_ShareToken.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ShareToken *ShareTokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ShareToken.Contract.SupportsInterface(&_ShareToken.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ShareToken *ShareTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ShareToken *ShareTokenSession) Symbol() (string, error) {
	return _ShareToken.Contract.Symbol(&_ShareToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ShareToken *ShareTokenCallerSession) Symbol() (string, error) {
	return _ShareToken.Contract.Symbol(&_ShareToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_ShareToken *ShareTokenCaller) TotalSupply(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "totalSupply", id)
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_ShareToken *ShareTokenSession) TotalSupply(id *big.Int) (*big.Int, error) {
	return _ShareToken.Contract.TotalSupply(&_ShareToken.CallOpts, id)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_ShareToken *ShareTokenCallerSession) TotalSupply(id *big.Int) (*big.Int, error) {
	return _ShareToken.Contract.TotalSupply(&_ShareToken.CallOpts, id)
}

// TotalSupplyForMarketOutcome is a free data retrieval call binding the contract method 0x8c5e2f87.
//
// Solidity: function totalSupplyForMarketOutcome(address _market, uint256 _outcome) view returns(uint256)
func (_ShareToken *ShareTokenCaller) TotalSupplyForMarketOutcome(opts *bind.CallOpts, _market common.Address, _outcome *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ShareToken.contract.Call(opts, out, "totalSupplyForMarketOutcome", _market, _outcome)
	return *ret0, err
}

// TotalSupplyForMarketOutcome is a free data retrieval call binding the contract method 0x8c5e2f87.
//
// Solidity: function totalSupplyForMarketOutcome(address _market, uint256 _outcome) view returns(uint256)
func (_ShareToken *ShareTokenSession) TotalSupplyForMarketOutcome(_market common.Address, _outcome *big.Int) (*big.Int, error) {
	return _ShareToken.Contract.TotalSupplyForMarketOutcome(&_ShareToken.CallOpts, _market, _outcome)
}

// TotalSupplyForMarketOutcome is a free data retrieval call binding the contract method 0x8c5e2f87.
//
// Solidity: function totalSupplyForMarketOutcome(address _market, uint256 _outcome) view returns(uint256)
func (_ShareToken *ShareTokenCallerSession) TotalSupplyForMarketOutcome(_market common.Address, _outcome *big.Int) (*big.Int, error) {
	return _ShareToken.Contract.TotalSupplyForMarketOutcome(&_ShareToken.CallOpts, _market, _outcome)
}

// UnpackTokenId is a free data retrieval call binding the contract method 0x26afd2e8.
//
// Solidity: function unpackTokenId(uint256 _tokenId) pure returns(address _market, uint256 _outcome)
func (_ShareToken *ShareTokenCaller) UnpackTokenId(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	Market  common.Address
	Outcome *big.Int
}, error) {
	ret := new(struct {
		Market  common.Address
		Outcome *big.Int
	})
	out := ret
	err := _ShareToken.contract.Call(opts, out, "unpackTokenId", _tokenId)
	return *ret, err
}

// UnpackTokenId is a free data retrieval call binding the contract method 0x26afd2e8.
//
// Solidity: function unpackTokenId(uint256 _tokenId) pure returns(address _market, uint256 _outcome)
func (_ShareToken *ShareTokenSession) UnpackTokenId(_tokenId *big.Int) (struct {
	Market  common.Address
	Outcome *big.Int
}, error) {
	return _ShareToken.Contract.UnpackTokenId(&_ShareToken.CallOpts, _tokenId)
}

// UnpackTokenId is a free data retrieval call binding the contract method 0x26afd2e8.
//
// Solidity: function unpackTokenId(uint256 _tokenId) pure returns(address _market, uint256 _outcome)
func (_ShareToken *ShareTokenCallerSession) UnpackTokenId(_tokenId *big.Int) (struct {
	Market  common.Address
	Outcome *big.Int
}, error) {
	return _ShareToken.Contract.UnpackTokenId(&_ShareToken.CallOpts, _tokenId)
}

// BuyCompleteSets is a paid mutator transaction binding the contract method 0xaa48db20.
//
// Solidity: function buyCompleteSets(address _market, address _account, uint256 _amount) returns(bool)
func (_ShareToken *ShareTokenTransactor) BuyCompleteSets(opts *bind.TransactOpts, _market common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ShareToken.contract.Transact(opts, "buyCompleteSets", _market, _account, _amount)
}

// BuyCompleteSets is a paid mutator transaction binding the contract method 0xaa48db20.
//
// Solidity: function buyCompleteSets(address _market, address _account, uint256 _amount) returns(bool)
func (_ShareToken *ShareTokenSession) BuyCompleteSets(_market common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.BuyCompleteSets(&_ShareToken.TransactOpts, _market, _account, _amount)
}

// BuyCompleteSets is a paid mutator transaction binding the contract method 0xaa48db20.
//
// Solidity: function buyCompleteSets(address _market, address _account, uint256 _amount) returns(bool)
func (_ShareToken *ShareTokenTransactorSession) BuyCompleteSets(_market common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.BuyCompleteSets(&_ShareToken.TransactOpts, _market, _account, _amount)
}

// BuyCompleteSetsForTrade is a paid mutator transaction binding the contract method 0x431e8878.
//
// Solidity: function buyCompleteSetsForTrade(address _market, uint256 _amount, uint256 _longOutcome, address _longRecipient, address _shortRecipient) returns(bool)
func (_ShareToken *ShareTokenTransactor) BuyCompleteSetsForTrade(opts *bind.TransactOpts, _market common.Address, _amount *big.Int, _longOutcome *big.Int, _longRecipient common.Address, _shortRecipient common.Address) (*types.Transaction, error) {
	return _ShareToken.contract.Transact(opts, "buyCompleteSetsForTrade", _market, _amount, _longOutcome, _longRecipient, _shortRecipient)
}

// BuyCompleteSetsForTrade is a paid mutator transaction binding the contract method 0x431e8878.
//
// Solidity: function buyCompleteSetsForTrade(address _market, uint256 _amount, uint256 _longOutcome, address _longRecipient, address _shortRecipient) returns(bool)
func (_ShareToken *ShareTokenSession) BuyCompleteSetsForTrade(_market common.Address, _amount *big.Int, _longOutcome *big.Int, _longRecipient common.Address, _shortRecipient common.Address) (*types.Transaction, error) {
	return _ShareToken.Contract.BuyCompleteSetsForTrade(&_ShareToken.TransactOpts, _market, _amount, _longOutcome, _longRecipient, _shortRecipient)
}

// BuyCompleteSetsForTrade is a paid mutator transaction binding the contract method 0x431e8878.
//
// Solidity: function buyCompleteSetsForTrade(address _market, uint256 _amount, uint256 _longOutcome, address _longRecipient, address _shortRecipient) returns(bool)
func (_ShareToken *ShareTokenTransactorSession) BuyCompleteSetsForTrade(_market common.Address, _amount *big.Int, _longOutcome *big.Int, _longRecipient common.Address, _shortRecipient common.Address) (*types.Transaction, error) {
	return _ShareToken.Contract.BuyCompleteSetsForTrade(&_ShareToken.TransactOpts, _market, _amount, _longOutcome, _longRecipient, _shortRecipient)
}

// CalculateReportingFee is a paid mutator transaction binding the contract method 0x81894407.
//
// Solidity: function calculateReportingFee(address _market, uint256 _amount) returns(uint256)
func (_ShareToken *ShareTokenTransactor) CalculateReportingFee(opts *bind.TransactOpts, _market common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ShareToken.contract.Transact(opts, "calculateReportingFee", _market, _amount)
}

// CalculateReportingFee is a paid mutator transaction binding the contract method 0x81894407.
//
// Solidity: function calculateReportingFee(address _market, uint256 _amount) returns(uint256)
func (_ShareToken *ShareTokenSession) CalculateReportingFee(_market common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.CalculateReportingFee(&_ShareToken.TransactOpts, _market, _amount)
}

// CalculateReportingFee is a paid mutator transaction binding the contract method 0x81894407.
//
// Solidity: function calculateReportingFee(address _market, uint256 _amount) returns(uint256)
func (_ShareToken *ShareTokenTransactorSession) CalculateReportingFee(_market common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.CalculateReportingFee(&_ShareToken.TransactOpts, _market, _amount)
}

// ClaimTradingProceeds is a paid mutator transaction binding the contract method 0xefd342c1.
//
// Solidity: function claimTradingProceeds(address _market, address _shareHolder, bytes32 _fingerprint) returns(uint256[] _outcomeFees)
func (_ShareToken *ShareTokenTransactor) ClaimTradingProceeds(opts *bind.TransactOpts, _market common.Address, _shareHolder common.Address, _fingerprint [32]byte) (*types.Transaction, error) {
	return _ShareToken.contract.Transact(opts, "claimTradingProceeds", _market, _shareHolder, _fingerprint)
}

// ClaimTradingProceeds is a paid mutator transaction binding the contract method 0xefd342c1.
//
// Solidity: function claimTradingProceeds(address _market, address _shareHolder, bytes32 _fingerprint) returns(uint256[] _outcomeFees)
func (_ShareToken *ShareTokenSession) ClaimTradingProceeds(_market common.Address, _shareHolder common.Address, _fingerprint [32]byte) (*types.Transaction, error) {
	return _ShareToken.Contract.ClaimTradingProceeds(&_ShareToken.TransactOpts, _market, _shareHolder, _fingerprint)
}

// ClaimTradingProceeds is a paid mutator transaction binding the contract method 0xefd342c1.
//
// Solidity: function claimTradingProceeds(address _market, address _shareHolder, bytes32 _fingerprint) returns(uint256[] _outcomeFees)
func (_ShareToken *ShareTokenTransactorSession) ClaimTradingProceeds(_market common.Address, _shareHolder common.Address, _fingerprint [32]byte) (*types.Transaction, error) {
	return _ShareToken.Contract.ClaimTradingProceeds(&_ShareToken.TransactOpts, _market, _shareHolder, _fingerprint)
}

// DivideUpWinnings is a paid mutator transaction binding the contract method 0x9f66cddf.
//
// Solidity: function divideUpWinnings(address _market, uint256 _outcome, uint256 _numberOfShares) returns(uint256 _proceeds, uint256 _shareHolderShare, uint256 _creatorShare, uint256 _reporterShare)
func (_ShareToken *ShareTokenTransactor) DivideUpWinnings(opts *bind.TransactOpts, _market common.Address, _outcome *big.Int, _numberOfShares *big.Int) (*types.Transaction, error) {
	return _ShareToken.contract.Transact(opts, "divideUpWinnings", _market, _outcome, _numberOfShares)
}

// DivideUpWinnings is a paid mutator transaction binding the contract method 0x9f66cddf.
//
// Solidity: function divideUpWinnings(address _market, uint256 _outcome, uint256 _numberOfShares) returns(uint256 _proceeds, uint256 _shareHolderShare, uint256 _creatorShare, uint256 _reporterShare)
func (_ShareToken *ShareTokenSession) DivideUpWinnings(_market common.Address, _outcome *big.Int, _numberOfShares *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.DivideUpWinnings(&_ShareToken.TransactOpts, _market, _outcome, _numberOfShares)
}

// DivideUpWinnings is a paid mutator transaction binding the contract method 0x9f66cddf.
//
// Solidity: function divideUpWinnings(address _market, uint256 _outcome, uint256 _numberOfShares) returns(uint256 _proceeds, uint256 _shareHolderShare, uint256 _creatorShare, uint256 _reporterShare)
func (_ShareToken *ShareTokenTransactorSession) DivideUpWinnings(_market common.Address, _outcome *big.Int, _numberOfShares *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.DivideUpWinnings(&_ShareToken.TransactOpts, _market, _outcome, _numberOfShares)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _augur) returns()
func (_ShareToken *ShareTokenTransactor) Initialize(opts *bind.TransactOpts, _augur common.Address) (*types.Transaction, error) {
	return _ShareToken.contract.Transact(opts, "initialize", _augur)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _augur) returns()
func (_ShareToken *ShareTokenSession) Initialize(_augur common.Address) (*types.Transaction, error) {
	return _ShareToken.Contract.Initialize(&_ShareToken.TransactOpts, _augur)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _augur) returns()
func (_ShareToken *ShareTokenTransactorSession) Initialize(_augur common.Address) (*types.Transaction, error) {
	return _ShareToken.Contract.Initialize(&_ShareToken.TransactOpts, _augur)
}

// InitializeMarket is a paid mutator transaction binding the contract method 0x3b48090e.
//
// Solidity: function initializeMarket(address _market, uint256 _numOutcomes, uint256 _numTicks) returns()
func (_ShareToken *ShareTokenTransactor) InitializeMarket(opts *bind.TransactOpts, _market common.Address, _numOutcomes *big.Int, _numTicks *big.Int) (*types.Transaction, error) {
	return _ShareToken.contract.Transact(opts, "initializeMarket", _market, _numOutcomes, _numTicks)
}

// InitializeMarket is a paid mutator transaction binding the contract method 0x3b48090e.
//
// Solidity: function initializeMarket(address _market, uint256 _numOutcomes, uint256 _numTicks) returns()
func (_ShareToken *ShareTokenSession) InitializeMarket(_market common.Address, _numOutcomes *big.Int, _numTicks *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.InitializeMarket(&_ShareToken.TransactOpts, _market, _numOutcomes, _numTicks)
}

// InitializeMarket is a paid mutator transaction binding the contract method 0x3b48090e.
//
// Solidity: function initializeMarket(address _market, uint256 _numOutcomes, uint256 _numTicks) returns()
func (_ShareToken *ShareTokenTransactorSession) InitializeMarket(_market common.Address, _numOutcomes *big.Int, _numTicks *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.InitializeMarket(&_ShareToken.TransactOpts, _market, _numOutcomes, _numTicks)
}

// PublicBuyCompleteSets is a paid mutator transaction binding the contract method 0xabb60c80.
//
// Solidity: function publicBuyCompleteSets(address _market, uint256 _amount) returns(bool)
func (_ShareToken *ShareTokenTransactor) PublicBuyCompleteSets(opts *bind.TransactOpts, _market common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ShareToken.contract.Transact(opts, "publicBuyCompleteSets", _market, _amount)
}

// PublicBuyCompleteSets is a paid mutator transaction binding the contract method 0xabb60c80.
//
// Solidity: function publicBuyCompleteSets(address _market, uint256 _amount) returns(bool)
func (_ShareToken *ShareTokenSession) PublicBuyCompleteSets(_market common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.PublicBuyCompleteSets(&_ShareToken.TransactOpts, _market, _amount)
}

// PublicBuyCompleteSets is a paid mutator transaction binding the contract method 0xabb60c80.
//
// Solidity: function publicBuyCompleteSets(address _market, uint256 _amount) returns(bool)
func (_ShareToken *ShareTokenTransactorSession) PublicBuyCompleteSets(_market common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.PublicBuyCompleteSets(&_ShareToken.TransactOpts, _market, _amount)
}

// PublicSellCompleteSets is a paid mutator transaction binding the contract method 0xad16158e.
//
// Solidity: function publicSellCompleteSets(address _market, uint256 _amount) returns(uint256 _creatorFee, uint256 _reportingFee)
func (_ShareToken *ShareTokenTransactor) PublicSellCompleteSets(opts *bind.TransactOpts, _market common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ShareToken.contract.Transact(opts, "publicSellCompleteSets", _market, _amount)
}

// PublicSellCompleteSets is a paid mutator transaction binding the contract method 0xad16158e.
//
// Solidity: function publicSellCompleteSets(address _market, uint256 _amount) returns(uint256 _creatorFee, uint256 _reportingFee)
func (_ShareToken *ShareTokenSession) PublicSellCompleteSets(_market common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.PublicSellCompleteSets(&_ShareToken.TransactOpts, _market, _amount)
}

// PublicSellCompleteSets is a paid mutator transaction binding the contract method 0xad16158e.
//
// Solidity: function publicSellCompleteSets(address _market, uint256 _amount) returns(uint256 _creatorFee, uint256 _reportingFee)
func (_ShareToken *ShareTokenTransactorSession) PublicSellCompleteSets(_market common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.PublicSellCompleteSets(&_ShareToken.TransactOpts, _market, _amount)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_ShareToken *ShareTokenTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _ShareToken.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_ShareToken *ShareTokenSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _ShareToken.Contract.SafeBatchTransferFrom(&_ShareToken.TransactOpts, from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_ShareToken *ShareTokenTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _ShareToken.Contract.SafeBatchTransferFrom(&_ShareToken.TransactOpts, from, to, ids, values, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_ShareToken *ShareTokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ShareToken.contract.Transact(opts, "safeTransferFrom", from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_ShareToken *ShareTokenSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ShareToken.Contract.SafeTransferFrom(&_ShareToken.TransactOpts, from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_ShareToken *ShareTokenTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ShareToken.Contract.SafeTransferFrom(&_ShareToken.TransactOpts, from, to, id, value, data)
}

// SellCompleteSets is a paid mutator transaction binding the contract method 0x5990fea6.
//
// Solidity: function sellCompleteSets(address _market, address _holder, address _recipient, uint256 _amount, bytes32 _fingerprint) returns(uint256 _creatorFee, uint256 _reportingFee)
func (_ShareToken *ShareTokenTransactor) SellCompleteSets(opts *bind.TransactOpts, _market common.Address, _holder common.Address, _recipient common.Address, _amount *big.Int, _fingerprint [32]byte) (*types.Transaction, error) {
	return _ShareToken.contract.Transact(opts, "sellCompleteSets", _market, _holder, _recipient, _amount, _fingerprint)
}

// SellCompleteSets is a paid mutator transaction binding the contract method 0x5990fea6.
//
// Solidity: function sellCompleteSets(address _market, address _holder, address _recipient, uint256 _amount, bytes32 _fingerprint) returns(uint256 _creatorFee, uint256 _reportingFee)
func (_ShareToken *ShareTokenSession) SellCompleteSets(_market common.Address, _holder common.Address, _recipient common.Address, _amount *big.Int, _fingerprint [32]byte) (*types.Transaction, error) {
	return _ShareToken.Contract.SellCompleteSets(&_ShareToken.TransactOpts, _market, _holder, _recipient, _amount, _fingerprint)
}

// SellCompleteSets is a paid mutator transaction binding the contract method 0x5990fea6.
//
// Solidity: function sellCompleteSets(address _market, address _holder, address _recipient, uint256 _amount, bytes32 _fingerprint) returns(uint256 _creatorFee, uint256 _reportingFee)
func (_ShareToken *ShareTokenTransactorSession) SellCompleteSets(_market common.Address, _holder common.Address, _recipient common.Address, _amount *big.Int, _fingerprint [32]byte) (*types.Transaction, error) {
	return _ShareToken.Contract.SellCompleteSets(&_ShareToken.TransactOpts, _market, _holder, _recipient, _amount, _fingerprint)
}

// SellCompleteSetsForTrade is a paid mutator transaction binding the contract method 0x53b7a712.
//
// Solidity: function sellCompleteSetsForTrade(address _market, uint256 _outcome, uint256 _amount, address _shortParticipant, address _longParticipant, address _shortRecipient, address _longRecipient, uint256 _price, address _sourceAccount, bytes32 _fingerprint) returns(uint256 _creatorFee, uint256 _reportingFee)
func (_ShareToken *ShareTokenTransactor) SellCompleteSetsForTrade(opts *bind.TransactOpts, _market common.Address, _outcome *big.Int, _amount *big.Int, _shortParticipant common.Address, _longParticipant common.Address, _shortRecipient common.Address, _longRecipient common.Address, _price *big.Int, _sourceAccount common.Address, _fingerprint [32]byte) (*types.Transaction, error) {
	return _ShareToken.contract.Transact(opts, "sellCompleteSetsForTrade", _market, _outcome, _amount, _shortParticipant, _longParticipant, _shortRecipient, _longRecipient, _price, _sourceAccount, _fingerprint)
}

// SellCompleteSetsForTrade is a paid mutator transaction binding the contract method 0x53b7a712.
//
// Solidity: function sellCompleteSetsForTrade(address _market, uint256 _outcome, uint256 _amount, address _shortParticipant, address _longParticipant, address _shortRecipient, address _longRecipient, uint256 _price, address _sourceAccount, bytes32 _fingerprint) returns(uint256 _creatorFee, uint256 _reportingFee)
func (_ShareToken *ShareTokenSession) SellCompleteSetsForTrade(_market common.Address, _outcome *big.Int, _amount *big.Int, _shortParticipant common.Address, _longParticipant common.Address, _shortRecipient common.Address, _longRecipient common.Address, _price *big.Int, _sourceAccount common.Address, _fingerprint [32]byte) (*types.Transaction, error) {
	return _ShareToken.Contract.SellCompleteSetsForTrade(&_ShareToken.TransactOpts, _market, _outcome, _amount, _shortParticipant, _longParticipant, _shortRecipient, _longRecipient, _price, _sourceAccount, _fingerprint)
}

// SellCompleteSetsForTrade is a paid mutator transaction binding the contract method 0x53b7a712.
//
// Solidity: function sellCompleteSetsForTrade(address _market, uint256 _outcome, uint256 _amount, address _shortParticipant, address _longParticipant, address _shortRecipient, address _longRecipient, uint256 _price, address _sourceAccount, bytes32 _fingerprint) returns(uint256 _creatorFee, uint256 _reportingFee)
func (_ShareToken *ShareTokenTransactorSession) SellCompleteSetsForTrade(_market common.Address, _outcome *big.Int, _amount *big.Int, _shortParticipant common.Address, _longParticipant common.Address, _shortRecipient common.Address, _longRecipient common.Address, _price *big.Int, _sourceAccount common.Address, _fingerprint [32]byte) (*types.Transaction, error) {
	return _ShareToken.Contract.SellCompleteSetsForTrade(&_ShareToken.TransactOpts, _market, _outcome, _amount, _shortParticipant, _longParticipant, _shortRecipient, _longRecipient, _price, _sourceAccount, _fingerprint)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ShareToken *ShareTokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ShareToken.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ShareToken *ShareTokenSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ShareToken.Contract.SetApprovalForAll(&_ShareToken.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ShareToken *ShareTokenTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ShareToken.Contract.SetApprovalForAll(&_ShareToken.TransactOpts, operator, approved)
}

// UnsafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x4028e8ea.
//
// Solidity: function unsafeBatchTransferFrom(address _from, address _to, uint256[] _ids, uint256[] _values) returns()
func (_ShareToken *ShareTokenTransactor) UnsafeBatchTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _ids []*big.Int, _values []*big.Int) (*types.Transaction, error) {
	return _ShareToken.contract.Transact(opts, "unsafeBatchTransferFrom", _from, _to, _ids, _values)
}

// UnsafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x4028e8ea.
//
// Solidity: function unsafeBatchTransferFrom(address _from, address _to, uint256[] _ids, uint256[] _values) returns()
func (_ShareToken *ShareTokenSession) UnsafeBatchTransferFrom(_from common.Address, _to common.Address, _ids []*big.Int, _values []*big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.UnsafeBatchTransferFrom(&_ShareToken.TransactOpts, _from, _to, _ids, _values)
}

// UnsafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x4028e8ea.
//
// Solidity: function unsafeBatchTransferFrom(address _from, address _to, uint256[] _ids, uint256[] _values) returns()
func (_ShareToken *ShareTokenTransactorSession) UnsafeBatchTransferFrom(_from common.Address, _to common.Address, _ids []*big.Int, _values []*big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.UnsafeBatchTransferFrom(&_ShareToken.TransactOpts, _from, _to, _ids, _values)
}

// UnsafeTransferFrom is a paid mutator transaction binding the contract method 0xfbc37ca9.
//
// Solidity: function unsafeTransferFrom(address _from, address _to, uint256 _id, uint256 _value) returns()
func (_ShareToken *ShareTokenTransactor) UnsafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _id *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _ShareToken.contract.Transact(opts, "unsafeTransferFrom", _from, _to, _id, _value)
}

// UnsafeTransferFrom is a paid mutator transaction binding the contract method 0xfbc37ca9.
//
// Solidity: function unsafeTransferFrom(address _from, address _to, uint256 _id, uint256 _value) returns()
func (_ShareToken *ShareTokenSession) UnsafeTransferFrom(_from common.Address, _to common.Address, _id *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.UnsafeTransferFrom(&_ShareToken.TransactOpts, _from, _to, _id, _value)
}

// UnsafeTransferFrom is a paid mutator transaction binding the contract method 0xfbc37ca9.
//
// Solidity: function unsafeTransferFrom(address _from, address _to, uint256 _id, uint256 _value) returns()
func (_ShareToken *ShareTokenTransactorSession) UnsafeTransferFrom(_from common.Address, _to common.Address, _id *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _ShareToken.Contract.UnsafeTransferFrom(&_ShareToken.TransactOpts, _from, _to, _id, _value)
}

// ShareTokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ShareToken contract.
type ShareTokenApprovalForAllIterator struct {
	Event *ShareTokenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ShareTokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShareTokenApprovalForAll)
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
		it.Event = new(ShareTokenApprovalForAll)
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
func (it *ShareTokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShareTokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShareTokenApprovalForAll represents a ApprovalForAll event raised by the ShareToken contract.
type ShareTokenApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ShareToken *ShareTokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ShareTokenApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ShareToken.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ShareTokenApprovalForAllIterator{contract: _ShareToken.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ShareToken *ShareTokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ShareTokenApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ShareToken.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShareTokenApprovalForAll)
				if err := _ShareToken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ShareToken *ShareTokenFilterer) ParseApprovalForAll(log types.Log) (*ShareTokenApprovalForAll, error) {
	event := new(ShareTokenApprovalForAll)
	if err := _ShareToken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ShareTokenTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the ShareToken contract.
type ShareTokenTransferBatchIterator struct {
	Event *ShareTokenTransferBatch // Event containing the contract specifics and raw log

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
func (it *ShareTokenTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShareTokenTransferBatch)
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
		it.Event = new(ShareTokenTransferBatch)
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
func (it *ShareTokenTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShareTokenTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShareTokenTransferBatch represents a TransferBatch event raised by the ShareToken contract.
type ShareTokenTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ShareToken *ShareTokenFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ShareTokenTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShareToken.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ShareTokenTransferBatchIterator{contract: _ShareToken.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ShareToken *ShareTokenFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *ShareTokenTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShareToken.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShareTokenTransferBatch)
				if err := _ShareToken.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ShareToken *ShareTokenFilterer) ParseTransferBatch(log types.Log) (*ShareTokenTransferBatch, error) {
	event := new(ShareTokenTransferBatch)
	if err := _ShareToken.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ShareTokenTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the ShareToken contract.
type ShareTokenTransferSingleIterator struct {
	Event *ShareTokenTransferSingle // Event containing the contract specifics and raw log

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
func (it *ShareTokenTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShareTokenTransferSingle)
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
		it.Event = new(ShareTokenTransferSingle)
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
func (it *ShareTokenTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShareTokenTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShareTokenTransferSingle represents a TransferSingle event raised by the ShareToken contract.
type ShareTokenTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ShareToken *ShareTokenFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ShareTokenTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShareToken.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ShareTokenTransferSingleIterator{contract: _ShareToken.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ShareToken *ShareTokenFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *ShareTokenTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShareToken.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShareTokenTransferSingle)
				if err := _ShareToken.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ShareToken *ShareTokenFilterer) ParseTransferSingle(log types.Log) (*ShareTokenTransferSingle, error) {
	event := new(ShareTokenTransferSingle)
	if err := _ShareToken.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ShareTokenURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the ShareToken contract.
type ShareTokenURIIterator struct {
	Event *ShareTokenURI // Event containing the contract specifics and raw log

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
func (it *ShareTokenURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShareTokenURI)
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
		it.Event = new(ShareTokenURI)
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
func (it *ShareTokenURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShareTokenURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShareTokenURI represents a URI event raised by the ShareToken contract.
type ShareTokenURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ShareToken *ShareTokenFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*ShareTokenURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ShareToken.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &ShareTokenURIIterator{contract: _ShareToken.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ShareToken *ShareTokenFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *ShareTokenURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ShareToken.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShareTokenURI)
				if err := _ShareToken.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ShareToken *ShareTokenFilterer) ParseURI(log types.Log) (*ShareTokenURI, error) {
	event := new(ShareTokenURI)
	if err := _ShareToken.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	return event, nil
}
