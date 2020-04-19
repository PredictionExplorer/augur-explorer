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
const TokenABI = "[{\"inputs\":[{\"internalType\":\"contractIAugur\",\"name\":\"_augur\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outcome\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"CancelZeroXOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"volume\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"outcomeVolumes\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"MarketVolumeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumAugurTrading.OrderEventType\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"tradeGroupId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addressData\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"uint256Data\",\"type\":\"uint256[]\"}],\"name\":\"OrderEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outcome\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netPosition\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"avgPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"realizedProfit\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"frozenFunds\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"realizedCost\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ProfitLossChanged\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"augur\",\"outputs\":[{\"internalType\":\"contractIAugur\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIMarket[]\",\"name\":\"_markets\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_shareHolder\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_fingerprint\",\"type\":\"bytes32\"}],\"name\":\"claimMarketsProceeds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_shareHolder\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_fingerprint\",\"type\":\"bytes32\"}],\"name\":\"claimTradingProceeds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"doApprovals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishDeployment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_volume\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_outcomeVolumes\",\"type\":\"uint256[]\"}],\"name\":\"logMarketVolumeChanged\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenRefund\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sharesRefund\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"logOrderCanceled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_orderId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_tradeGroupId\",\"type\":\"bytes32\"}],\"name\":\"logOrderCreated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_filler\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountFilled\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_orderId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_tradeGroupId\",\"type\":\"bytes32\"}],\"name\":\"logOrderFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outcome\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_netPosition\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"_avgPrice\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_realizedProfit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_frozenFunds\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_realizedCost\",\"type\":\"int256\"}],\"name\":\"logProfitLossChanged\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outcome\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_orderHash\",\"type\":\"bytes32\"}],\"name\":\"logZeroXOrderCanceled\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_orderHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_tradeGroupId\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_orderType\",\"type\":\"uint8\"},{\"internalType\":\"address[]\",\"name\":\"_addressData\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_uint256Data\",\"type\":\"uint256[]\"}],\"name\":\"logZeroXOrderFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"lookup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"registerContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"shareToken\",\"outputs\":[{\"internalType\":\"contractIShareToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"trustedSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"uploader\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

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

// Lookup is a free data retrieval call binding the contract method 0xf39ec1f7.
//
// Solidity: function lookup(bytes32 _key) view returns(address)
func (_Token *TokenCaller) Lookup(opts *bind.CallOpts, _key [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "lookup", _key)
	return *ret0, err
}

// Lookup is a free data retrieval call binding the contract method 0xf39ec1f7.
//
// Solidity: function lookup(bytes32 _key) view returns(address)
func (_Token *TokenSession) Lookup(_key [32]byte) (common.Address, error) {
	return _Token.Contract.Lookup(&_Token.CallOpts, _key)
}

// Lookup is a free data retrieval call binding the contract method 0xf39ec1f7.
//
// Solidity: function lookup(bytes32 _key) view returns(address)
func (_Token *TokenCallerSession) Lookup(_key [32]byte) (common.Address, error) {
	return _Token.Contract.Lookup(&_Token.CallOpts, _key)
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

// TrustedSender is a free data retrieval call binding the contract method 0x3f6ba415.
//
// Solidity: function trustedSender(address ) view returns(bool)
func (_Token *TokenCaller) TrustedSender(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "trustedSender", arg0)
	return *ret0, err
}

// TrustedSender is a free data retrieval call binding the contract method 0x3f6ba415.
//
// Solidity: function trustedSender(address ) view returns(bool)
func (_Token *TokenSession) TrustedSender(arg0 common.Address) (bool, error) {
	return _Token.Contract.TrustedSender(&_Token.CallOpts, arg0)
}

// TrustedSender is a free data retrieval call binding the contract method 0x3f6ba415.
//
// Solidity: function trustedSender(address ) view returns(bool)
func (_Token *TokenCallerSession) TrustedSender(arg0 common.Address) (bool, error) {
	return _Token.Contract.TrustedSender(&_Token.CallOpts, arg0)
}

// Uploader is a free data retrieval call binding the contract method 0x65fe2a0b.
//
// Solidity: function uploader() view returns(address)
func (_Token *TokenCaller) Uploader(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "uploader")
	return *ret0, err
}

// Uploader is a free data retrieval call binding the contract method 0x65fe2a0b.
//
// Solidity: function uploader() view returns(address)
func (_Token *TokenSession) Uploader() (common.Address, error) {
	return _Token.Contract.Uploader(&_Token.CallOpts)
}

// Uploader is a free data retrieval call binding the contract method 0x65fe2a0b.
//
// Solidity: function uploader() view returns(address)
func (_Token *TokenCallerSession) Uploader() (common.Address, error) {
	return _Token.Contract.Uploader(&_Token.CallOpts)
}

// ClaimMarketsProceeds is a paid mutator transaction binding the contract method 0xdb754422.
//
// Solidity: function claimMarketsProceeds(address[] _markets, address _shareHolder, bytes32 _fingerprint) returns(bool)
func (_Token *TokenTransactor) ClaimMarketsProceeds(opts *bind.TransactOpts, _markets []common.Address, _shareHolder common.Address, _fingerprint [32]byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "claimMarketsProceeds", _markets, _shareHolder, _fingerprint)
}

// ClaimMarketsProceeds is a paid mutator transaction binding the contract method 0xdb754422.
//
// Solidity: function claimMarketsProceeds(address[] _markets, address _shareHolder, bytes32 _fingerprint) returns(bool)
func (_Token *TokenSession) ClaimMarketsProceeds(_markets []common.Address, _shareHolder common.Address, _fingerprint [32]byte) (*types.Transaction, error) {
	return _Token.Contract.ClaimMarketsProceeds(&_Token.TransactOpts, _markets, _shareHolder, _fingerprint)
}

// ClaimMarketsProceeds is a paid mutator transaction binding the contract method 0xdb754422.
//
// Solidity: function claimMarketsProceeds(address[] _markets, address _shareHolder, bytes32 _fingerprint) returns(bool)
func (_Token *TokenTransactorSession) ClaimMarketsProceeds(_markets []common.Address, _shareHolder common.Address, _fingerprint [32]byte) (*types.Transaction, error) {
	return _Token.Contract.ClaimMarketsProceeds(&_Token.TransactOpts, _markets, _shareHolder, _fingerprint)
}

// ClaimTradingProceeds is a paid mutator transaction binding the contract method 0xefd342c1.
//
// Solidity: function claimTradingProceeds(address _market, address _shareHolder, bytes32 _fingerprint) returns(bool)
func (_Token *TokenTransactor) ClaimTradingProceeds(opts *bind.TransactOpts, _market common.Address, _shareHolder common.Address, _fingerprint [32]byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "claimTradingProceeds", _market, _shareHolder, _fingerprint)
}

// ClaimTradingProceeds is a paid mutator transaction binding the contract method 0xefd342c1.
//
// Solidity: function claimTradingProceeds(address _market, address _shareHolder, bytes32 _fingerprint) returns(bool)
func (_Token *TokenSession) ClaimTradingProceeds(_market common.Address, _shareHolder common.Address, _fingerprint [32]byte) (*types.Transaction, error) {
	return _Token.Contract.ClaimTradingProceeds(&_Token.TransactOpts, _market, _shareHolder, _fingerprint)
}

// ClaimTradingProceeds is a paid mutator transaction binding the contract method 0xefd342c1.
//
// Solidity: function claimTradingProceeds(address _market, address _shareHolder, bytes32 _fingerprint) returns(bool)
func (_Token *TokenTransactorSession) ClaimTradingProceeds(_market common.Address, _shareHolder common.Address, _fingerprint [32]byte) (*types.Transaction, error) {
	return _Token.Contract.ClaimTradingProceeds(&_Token.TransactOpts, _market, _shareHolder, _fingerprint)
}

// DoApprovals is a paid mutator transaction binding the contract method 0x5d4f7742.
//
// Solidity: function doApprovals() returns(bool)
func (_Token *TokenTransactor) DoApprovals(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "doApprovals")
}

// DoApprovals is a paid mutator transaction binding the contract method 0x5d4f7742.
//
// Solidity: function doApprovals() returns(bool)
func (_Token *TokenSession) DoApprovals() (*types.Transaction, error) {
	return _Token.Contract.DoApprovals(&_Token.TransactOpts)
}

// DoApprovals is a paid mutator transaction binding the contract method 0x5d4f7742.
//
// Solidity: function doApprovals() returns(bool)
func (_Token *TokenTransactorSession) DoApprovals() (*types.Transaction, error) {
	return _Token.Contract.DoApprovals(&_Token.TransactOpts)
}

// FinishDeployment is a paid mutator transaction binding the contract method 0xe997fffd.
//
// Solidity: function finishDeployment() returns(bool)
func (_Token *TokenTransactor) FinishDeployment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "finishDeployment")
}

// FinishDeployment is a paid mutator transaction binding the contract method 0xe997fffd.
//
// Solidity: function finishDeployment() returns(bool)
func (_Token *TokenSession) FinishDeployment() (*types.Transaction, error) {
	return _Token.Contract.FinishDeployment(&_Token.TransactOpts)
}

// FinishDeployment is a paid mutator transaction binding the contract method 0xe997fffd.
//
// Solidity: function finishDeployment() returns(bool)
func (_Token *TokenTransactorSession) FinishDeployment() (*types.Transaction, error) {
	return _Token.Contract.FinishDeployment(&_Token.TransactOpts)
}

// LogMarketVolumeChanged is a paid mutator transaction binding the contract method 0x74cf9765.
//
// Solidity: function logMarketVolumeChanged(address _universe, address _market, uint256 _volume, uint256[] _outcomeVolumes) returns(bool)
func (_Token *TokenTransactor) LogMarketVolumeChanged(opts *bind.TransactOpts, _universe common.Address, _market common.Address, _volume *big.Int, _outcomeVolumes []*big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "logMarketVolumeChanged", _universe, _market, _volume, _outcomeVolumes)
}

// LogMarketVolumeChanged is a paid mutator transaction binding the contract method 0x74cf9765.
//
// Solidity: function logMarketVolumeChanged(address _universe, address _market, uint256 _volume, uint256[] _outcomeVolumes) returns(bool)
func (_Token *TokenSession) LogMarketVolumeChanged(_universe common.Address, _market common.Address, _volume *big.Int, _outcomeVolumes []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.LogMarketVolumeChanged(&_Token.TransactOpts, _universe, _market, _volume, _outcomeVolumes)
}

// LogMarketVolumeChanged is a paid mutator transaction binding the contract method 0x74cf9765.
//
// Solidity: function logMarketVolumeChanged(address _universe, address _market, uint256 _volume, uint256[] _outcomeVolumes) returns(bool)
func (_Token *TokenTransactorSession) LogMarketVolumeChanged(_universe common.Address, _market common.Address, _volume *big.Int, _outcomeVolumes []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.LogMarketVolumeChanged(&_Token.TransactOpts, _universe, _market, _volume, _outcomeVolumes)
}

// LogOrderCanceled is a paid mutator transaction binding the contract method 0xc9e34878.
//
// Solidity: function logOrderCanceled(address _universe, address _market, address _creator, uint256 _tokenRefund, uint256 _sharesRefund, bytes32 _orderId) returns(bool)
func (_Token *TokenTransactor) LogOrderCanceled(opts *bind.TransactOpts, _universe common.Address, _market common.Address, _creator common.Address, _tokenRefund *big.Int, _sharesRefund *big.Int, _orderId [32]byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "logOrderCanceled", _universe, _market, _creator, _tokenRefund, _sharesRefund, _orderId)
}

// LogOrderCanceled is a paid mutator transaction binding the contract method 0xc9e34878.
//
// Solidity: function logOrderCanceled(address _universe, address _market, address _creator, uint256 _tokenRefund, uint256 _sharesRefund, bytes32 _orderId) returns(bool)
func (_Token *TokenSession) LogOrderCanceled(_universe common.Address, _market common.Address, _creator common.Address, _tokenRefund *big.Int, _sharesRefund *big.Int, _orderId [32]byte) (*types.Transaction, error) {
	return _Token.Contract.LogOrderCanceled(&_Token.TransactOpts, _universe, _market, _creator, _tokenRefund, _sharesRefund, _orderId)
}

// LogOrderCanceled is a paid mutator transaction binding the contract method 0xc9e34878.
//
// Solidity: function logOrderCanceled(address _universe, address _market, address _creator, uint256 _tokenRefund, uint256 _sharesRefund, bytes32 _orderId) returns(bool)
func (_Token *TokenTransactorSession) LogOrderCanceled(_universe common.Address, _market common.Address, _creator common.Address, _tokenRefund *big.Int, _sharesRefund *big.Int, _orderId [32]byte) (*types.Transaction, error) {
	return _Token.Contract.LogOrderCanceled(&_Token.TransactOpts, _universe, _market, _creator, _tokenRefund, _sharesRefund, _orderId)
}

// LogOrderCreated is a paid mutator transaction binding the contract method 0x98c7cc36.
//
// Solidity: function logOrderCreated(address _universe, bytes32 _orderId, bytes32 _tradeGroupId) returns(bool)
func (_Token *TokenTransactor) LogOrderCreated(opts *bind.TransactOpts, _universe common.Address, _orderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "logOrderCreated", _universe, _orderId, _tradeGroupId)
}

// LogOrderCreated is a paid mutator transaction binding the contract method 0x98c7cc36.
//
// Solidity: function logOrderCreated(address _universe, bytes32 _orderId, bytes32 _tradeGroupId) returns(bool)
func (_Token *TokenSession) LogOrderCreated(_universe common.Address, _orderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _Token.Contract.LogOrderCreated(&_Token.TransactOpts, _universe, _orderId, _tradeGroupId)
}

// LogOrderCreated is a paid mutator transaction binding the contract method 0x98c7cc36.
//
// Solidity: function logOrderCreated(address _universe, bytes32 _orderId, bytes32 _tradeGroupId) returns(bool)
func (_Token *TokenTransactorSession) LogOrderCreated(_universe common.Address, _orderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _Token.Contract.LogOrderCreated(&_Token.TransactOpts, _universe, _orderId, _tradeGroupId)
}

// LogOrderFilled is a paid mutator transaction binding the contract method 0x5ae5c8e7.
//
// Solidity: function logOrderFilled(address _universe, address _creator, address _filler, uint256 _price, uint256 _fees, uint256 _amountFilled, bytes32 _orderId, bytes32 _tradeGroupId) returns(bool)
func (_Token *TokenTransactor) LogOrderFilled(opts *bind.TransactOpts, _universe common.Address, _creator common.Address, _filler common.Address, _price *big.Int, _fees *big.Int, _amountFilled *big.Int, _orderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "logOrderFilled", _universe, _creator, _filler, _price, _fees, _amountFilled, _orderId, _tradeGroupId)
}

// LogOrderFilled is a paid mutator transaction binding the contract method 0x5ae5c8e7.
//
// Solidity: function logOrderFilled(address _universe, address _creator, address _filler, uint256 _price, uint256 _fees, uint256 _amountFilled, bytes32 _orderId, bytes32 _tradeGroupId) returns(bool)
func (_Token *TokenSession) LogOrderFilled(_universe common.Address, _creator common.Address, _filler common.Address, _price *big.Int, _fees *big.Int, _amountFilled *big.Int, _orderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _Token.Contract.LogOrderFilled(&_Token.TransactOpts, _universe, _creator, _filler, _price, _fees, _amountFilled, _orderId, _tradeGroupId)
}

// LogOrderFilled is a paid mutator transaction binding the contract method 0x5ae5c8e7.
//
// Solidity: function logOrderFilled(address _universe, address _creator, address _filler, uint256 _price, uint256 _fees, uint256 _amountFilled, bytes32 _orderId, bytes32 _tradeGroupId) returns(bool)
func (_Token *TokenTransactorSession) LogOrderFilled(_universe common.Address, _creator common.Address, _filler common.Address, _price *big.Int, _fees *big.Int, _amountFilled *big.Int, _orderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _Token.Contract.LogOrderFilled(&_Token.TransactOpts, _universe, _creator, _filler, _price, _fees, _amountFilled, _orderId, _tradeGroupId)
}

// LogProfitLossChanged is a paid mutator transaction binding the contract method 0x0848a90c.
//
// Solidity: function logProfitLossChanged(address _market, address _account, uint256 _outcome, int256 _netPosition, uint256 _avgPrice, int256 _realizedProfit, int256 _frozenFunds, int256 _realizedCost) returns(bool)
func (_Token *TokenTransactor) LogProfitLossChanged(opts *bind.TransactOpts, _market common.Address, _account common.Address, _outcome *big.Int, _netPosition *big.Int, _avgPrice *big.Int, _realizedProfit *big.Int, _frozenFunds *big.Int, _realizedCost *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "logProfitLossChanged", _market, _account, _outcome, _netPosition, _avgPrice, _realizedProfit, _frozenFunds, _realizedCost)
}

// LogProfitLossChanged is a paid mutator transaction binding the contract method 0x0848a90c.
//
// Solidity: function logProfitLossChanged(address _market, address _account, uint256 _outcome, int256 _netPosition, uint256 _avgPrice, int256 _realizedProfit, int256 _frozenFunds, int256 _realizedCost) returns(bool)
func (_Token *TokenSession) LogProfitLossChanged(_market common.Address, _account common.Address, _outcome *big.Int, _netPosition *big.Int, _avgPrice *big.Int, _realizedProfit *big.Int, _frozenFunds *big.Int, _realizedCost *big.Int) (*types.Transaction, error) {
	return _Token.Contract.LogProfitLossChanged(&_Token.TransactOpts, _market, _account, _outcome, _netPosition, _avgPrice, _realizedProfit, _frozenFunds, _realizedCost)
}

// LogProfitLossChanged is a paid mutator transaction binding the contract method 0x0848a90c.
//
// Solidity: function logProfitLossChanged(address _market, address _account, uint256 _outcome, int256 _netPosition, uint256 _avgPrice, int256 _realizedProfit, int256 _frozenFunds, int256 _realizedCost) returns(bool)
func (_Token *TokenTransactorSession) LogProfitLossChanged(_market common.Address, _account common.Address, _outcome *big.Int, _netPosition *big.Int, _avgPrice *big.Int, _realizedProfit *big.Int, _frozenFunds *big.Int, _realizedCost *big.Int) (*types.Transaction, error) {
	return _Token.Contract.LogProfitLossChanged(&_Token.TransactOpts, _market, _account, _outcome, _netPosition, _avgPrice, _realizedProfit, _frozenFunds, _realizedCost)
}

// LogZeroXOrderCanceled is a paid mutator transaction binding the contract method 0x577bb86e.
//
// Solidity: function logZeroXOrderCanceled(address _universe, address _market, address _account, uint256 _outcome, uint256 _price, uint256 _amount, uint8 _type, bytes32 _orderHash) returns()
func (_Token *TokenTransactor) LogZeroXOrderCanceled(opts *bind.TransactOpts, _universe common.Address, _market common.Address, _account common.Address, _outcome *big.Int, _price *big.Int, _amount *big.Int, _type uint8, _orderHash [32]byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "logZeroXOrderCanceled", _universe, _market, _account, _outcome, _price, _amount, _type, _orderHash)
}

// LogZeroXOrderCanceled is a paid mutator transaction binding the contract method 0x577bb86e.
//
// Solidity: function logZeroXOrderCanceled(address _universe, address _market, address _account, uint256 _outcome, uint256 _price, uint256 _amount, uint8 _type, bytes32 _orderHash) returns()
func (_Token *TokenSession) LogZeroXOrderCanceled(_universe common.Address, _market common.Address, _account common.Address, _outcome *big.Int, _price *big.Int, _amount *big.Int, _type uint8, _orderHash [32]byte) (*types.Transaction, error) {
	return _Token.Contract.LogZeroXOrderCanceled(&_Token.TransactOpts, _universe, _market, _account, _outcome, _price, _amount, _type, _orderHash)
}

// LogZeroXOrderCanceled is a paid mutator transaction binding the contract method 0x577bb86e.
//
// Solidity: function logZeroXOrderCanceled(address _universe, address _market, address _account, uint256 _outcome, uint256 _price, uint256 _amount, uint8 _type, bytes32 _orderHash) returns()
func (_Token *TokenTransactorSession) LogZeroXOrderCanceled(_universe common.Address, _market common.Address, _account common.Address, _outcome *big.Int, _price *big.Int, _amount *big.Int, _type uint8, _orderHash [32]byte) (*types.Transaction, error) {
	return _Token.Contract.LogZeroXOrderCanceled(&_Token.TransactOpts, _universe, _market, _account, _outcome, _price, _amount, _type, _orderHash)
}

// LogZeroXOrderFilled is a paid mutator transaction binding the contract method 0x99b537a0.
//
// Solidity: function logZeroXOrderFilled(address _universe, address _market, bytes32 _orderHash, bytes32 _tradeGroupId, uint8 _orderType, address[] _addressData, uint256[] _uint256Data) returns(bool)
func (_Token *TokenTransactor) LogZeroXOrderFilled(opts *bind.TransactOpts, _universe common.Address, _market common.Address, _orderHash [32]byte, _tradeGroupId [32]byte, _orderType uint8, _addressData []common.Address, _uint256Data []*big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "logZeroXOrderFilled", _universe, _market, _orderHash, _tradeGroupId, _orderType, _addressData, _uint256Data)
}

// LogZeroXOrderFilled is a paid mutator transaction binding the contract method 0x99b537a0.
//
// Solidity: function logZeroXOrderFilled(address _universe, address _market, bytes32 _orderHash, bytes32 _tradeGroupId, uint8 _orderType, address[] _addressData, uint256[] _uint256Data) returns(bool)
func (_Token *TokenSession) LogZeroXOrderFilled(_universe common.Address, _market common.Address, _orderHash [32]byte, _tradeGroupId [32]byte, _orderType uint8, _addressData []common.Address, _uint256Data []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.LogZeroXOrderFilled(&_Token.TransactOpts, _universe, _market, _orderHash, _tradeGroupId, _orderType, _addressData, _uint256Data)
}

// LogZeroXOrderFilled is a paid mutator transaction binding the contract method 0x99b537a0.
//
// Solidity: function logZeroXOrderFilled(address _universe, address _market, bytes32 _orderHash, bytes32 _tradeGroupId, uint8 _orderType, address[] _addressData, uint256[] _uint256Data) returns(bool)
func (_Token *TokenTransactorSession) LogZeroXOrderFilled(_universe common.Address, _market common.Address, _orderHash [32]byte, _tradeGroupId [32]byte, _orderType uint8, _addressData []common.Address, _uint256Data []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.LogZeroXOrderFilled(&_Token.TransactOpts, _universe, _market, _orderHash, _tradeGroupId, _orderType, _addressData, _uint256Data)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x645c6fae.
//
// Solidity: function registerContract(bytes32 _key, address _address) returns(bool)
func (_Token *TokenTransactor) RegisterContract(opts *bind.TransactOpts, _key [32]byte, _address common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "registerContract", _key, _address)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x645c6fae.
//
// Solidity: function registerContract(bytes32 _key, address _address) returns(bool)
func (_Token *TokenSession) RegisterContract(_key [32]byte, _address common.Address) (*types.Transaction, error) {
	return _Token.Contract.RegisterContract(&_Token.TransactOpts, _key, _address)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x645c6fae.
//
// Solidity: function registerContract(bytes32 _key, address _address) returns(bool)
func (_Token *TokenTransactorSession) RegisterContract(_key [32]byte, _address common.Address) (*types.Transaction, error) {
	return _Token.Contract.RegisterContract(&_Token.TransactOpts, _key, _address)
}

// TokenCancelZeroXOrderIterator is returned from FilterCancelZeroXOrder and is used to iterate over the raw logs and unpacked data for CancelZeroXOrder events raised by the Token contract.
type TokenCancelZeroXOrderIterator struct {
	Event *TokenCancelZeroXOrder // Event containing the contract specifics and raw log

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
func (it *TokenCancelZeroXOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenCancelZeroXOrder)
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
		it.Event = new(TokenCancelZeroXOrder)
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
func (it *TokenCancelZeroXOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenCancelZeroXOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenCancelZeroXOrder represents a CancelZeroXOrder event raised by the Token contract.
type TokenCancelZeroXOrder struct {
	Universe  common.Address
	Market    common.Address
	Account   common.Address
	Outcome   *big.Int
	Price     *big.Int
	Amount    *big.Int
	OrderType uint8
	OrderHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCancelZeroXOrder is a free log retrieval operation binding the contract event 0xbe80e5687d7095071b7c4e7a56e0e67bfb9e8a39352f1690fdf74c1ee935c75e.
//
// Solidity: event CancelZeroXOrder(address indexed universe, address indexed market, address indexed account, uint256 outcome, uint256 price, uint256 amount, uint8 orderType, bytes32 orderHash)
func (_Token *TokenFilterer) FilterCancelZeroXOrder(opts *bind.FilterOpts, universe []common.Address, market []common.Address, account []common.Address) (*TokenCancelZeroXOrderIterator, error) {

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

	logs, sub, err := _Token.contract.FilterLogs(opts, "CancelZeroXOrder", universeRule, marketRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &TokenCancelZeroXOrderIterator{contract: _Token.contract, event: "CancelZeroXOrder", logs: logs, sub: sub}, nil
}

// WatchCancelZeroXOrder is a free log subscription operation binding the contract event 0xbe80e5687d7095071b7c4e7a56e0e67bfb9e8a39352f1690fdf74c1ee935c75e.
//
// Solidity: event CancelZeroXOrder(address indexed universe, address indexed market, address indexed account, uint256 outcome, uint256 price, uint256 amount, uint8 orderType, bytes32 orderHash)
func (_Token *TokenFilterer) WatchCancelZeroXOrder(opts *bind.WatchOpts, sink chan<- *TokenCancelZeroXOrder, universe []common.Address, market []common.Address, account []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Token.contract.WatchLogs(opts, "CancelZeroXOrder", universeRule, marketRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenCancelZeroXOrder)
				if err := _Token.contract.UnpackLog(event, "CancelZeroXOrder", log); err != nil {
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

// ParseCancelZeroXOrder is a log parse operation binding the contract event 0xbe80e5687d7095071b7c4e7a56e0e67bfb9e8a39352f1690fdf74c1ee935c75e.
//
// Solidity: event CancelZeroXOrder(address indexed universe, address indexed market, address indexed account, uint256 outcome, uint256 price, uint256 amount, uint8 orderType, bytes32 orderHash)
func (_Token *TokenFilterer) ParseCancelZeroXOrder(log types.Log) (*TokenCancelZeroXOrder, error) {
	event := new(TokenCancelZeroXOrder)
	if err := _Token.contract.UnpackLog(event, "CancelZeroXOrder", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenMarketVolumeChangedIterator is returned from FilterMarketVolumeChanged and is used to iterate over the raw logs and unpacked data for MarketVolumeChanged events raised by the Token contract.
type TokenMarketVolumeChangedIterator struct {
	Event *TokenMarketVolumeChanged // Event containing the contract specifics and raw log

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
func (it *TokenMarketVolumeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenMarketVolumeChanged)
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
		it.Event = new(TokenMarketVolumeChanged)
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
func (it *TokenMarketVolumeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenMarketVolumeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenMarketVolumeChanged represents a MarketVolumeChanged event raised by the Token contract.
type TokenMarketVolumeChanged struct {
	Universe       common.Address
	Market         common.Address
	Volume         *big.Int
	OutcomeVolumes []*big.Int
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMarketVolumeChanged is a free log retrieval operation binding the contract event 0xe9f0af820300e73bae76c8e76943abe7fbb4224b49cb133e2dadc6f71acf6370.
//
// Solidity: event MarketVolumeChanged(address indexed universe, address indexed market, uint256 volume, uint256[] outcomeVolumes, uint256 timestamp)
func (_Token *TokenFilterer) FilterMarketVolumeChanged(opts *bind.FilterOpts, universe []common.Address, market []common.Address) (*TokenMarketVolumeChangedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "MarketVolumeChanged", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &TokenMarketVolumeChangedIterator{contract: _Token.contract, event: "MarketVolumeChanged", logs: logs, sub: sub}, nil
}

// WatchMarketVolumeChanged is a free log subscription operation binding the contract event 0xe9f0af820300e73bae76c8e76943abe7fbb4224b49cb133e2dadc6f71acf6370.
//
// Solidity: event MarketVolumeChanged(address indexed universe, address indexed market, uint256 volume, uint256[] outcomeVolumes, uint256 timestamp)
func (_Token *TokenFilterer) WatchMarketVolumeChanged(opts *bind.WatchOpts, sink chan<- *TokenMarketVolumeChanged, universe []common.Address, market []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "MarketVolumeChanged", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenMarketVolumeChanged)
				if err := _Token.contract.UnpackLog(event, "MarketVolumeChanged", log); err != nil {
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

// ParseMarketVolumeChanged is a log parse operation binding the contract event 0xe9f0af820300e73bae76c8e76943abe7fbb4224b49cb133e2dadc6f71acf6370.
//
// Solidity: event MarketVolumeChanged(address indexed universe, address indexed market, uint256 volume, uint256[] outcomeVolumes, uint256 timestamp)
func (_Token *TokenFilterer) ParseMarketVolumeChanged(log types.Log) (*TokenMarketVolumeChanged, error) {
	event := new(TokenMarketVolumeChanged)
	if err := _Token.contract.UnpackLog(event, "MarketVolumeChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenOrderEventIterator is returned from FilterOrderEvent and is used to iterate over the raw logs and unpacked data for OrderEvent events raised by the Token contract.
type TokenOrderEventIterator struct {
	Event *TokenOrderEvent // Event containing the contract specifics and raw log

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
func (it *TokenOrderEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenOrderEvent)
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
		it.Event = new(TokenOrderEvent)
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
func (it *TokenOrderEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenOrderEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenOrderEvent represents a OrderEvent event raised by the Token contract.
type TokenOrderEvent struct {
	Universe     common.Address
	Market       common.Address
	EventType    uint8
	OrderType    uint8
	OrderId      [32]byte
	TradeGroupId [32]byte
	AddressData  []common.Address
	Uint256Data  []*big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOrderEvent is a free log retrieval operation binding the contract event 0x9bab1368a1ed530afaad9c630ba75e6a5c1efa9f6af0139d6cda2b6af6aa801e.
//
// Solidity: event OrderEvent(address indexed universe, address indexed market, uint8 indexed eventType, uint8 orderType, bytes32 orderId, bytes32 tradeGroupId, address[] addressData, uint256[] uint256Data)
func (_Token *TokenFilterer) FilterOrderEvent(opts *bind.FilterOpts, universe []common.Address, market []common.Address, eventType []uint8) (*TokenOrderEventIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var eventTypeRule []interface{}
	for _, eventTypeItem := range eventType {
		eventTypeRule = append(eventTypeRule, eventTypeItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "OrderEvent", universeRule, marketRule, eventTypeRule)
	if err != nil {
		return nil, err
	}
	return &TokenOrderEventIterator{contract: _Token.contract, event: "OrderEvent", logs: logs, sub: sub}, nil
}

// WatchOrderEvent is a free log subscription operation binding the contract event 0x9bab1368a1ed530afaad9c630ba75e6a5c1efa9f6af0139d6cda2b6af6aa801e.
//
// Solidity: event OrderEvent(address indexed universe, address indexed market, uint8 indexed eventType, uint8 orderType, bytes32 orderId, bytes32 tradeGroupId, address[] addressData, uint256[] uint256Data)
func (_Token *TokenFilterer) WatchOrderEvent(opts *bind.WatchOpts, sink chan<- *TokenOrderEvent, universe []common.Address, market []common.Address, eventType []uint8) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var eventTypeRule []interface{}
	for _, eventTypeItem := range eventType {
		eventTypeRule = append(eventTypeRule, eventTypeItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "OrderEvent", universeRule, marketRule, eventTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenOrderEvent)
				if err := _Token.contract.UnpackLog(event, "OrderEvent", log); err != nil {
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

// ParseOrderEvent is a log parse operation binding the contract event 0x9bab1368a1ed530afaad9c630ba75e6a5c1efa9f6af0139d6cda2b6af6aa801e.
//
// Solidity: event OrderEvent(address indexed universe, address indexed market, uint8 indexed eventType, uint8 orderType, bytes32 orderId, bytes32 tradeGroupId, address[] addressData, uint256[] uint256Data)
func (_Token *TokenFilterer) ParseOrderEvent(log types.Log) (*TokenOrderEvent, error) {
	event := new(TokenOrderEvent)
	if err := _Token.contract.UnpackLog(event, "OrderEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenProfitLossChangedIterator is returned from FilterProfitLossChanged and is used to iterate over the raw logs and unpacked data for ProfitLossChanged events raised by the Token contract.
type TokenProfitLossChangedIterator struct {
	Event *TokenProfitLossChanged // Event containing the contract specifics and raw log

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
func (it *TokenProfitLossChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenProfitLossChanged)
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
		it.Event = new(TokenProfitLossChanged)
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
func (it *TokenProfitLossChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenProfitLossChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenProfitLossChanged represents a ProfitLossChanged event raised by the Token contract.
type TokenProfitLossChanged struct {
	Universe       common.Address
	Market         common.Address
	Account        common.Address
	Outcome        *big.Int
	NetPosition    *big.Int
	AvgPrice       *big.Int
	RealizedProfit *big.Int
	FrozenFunds    *big.Int
	RealizedCost   *big.Int
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterProfitLossChanged is a free log retrieval operation binding the contract event 0x59543b7f82735782aa5bdb97dff40ff288d4548a5865da513b40e4088e2ee77e.
//
// Solidity: event ProfitLossChanged(address indexed universe, address indexed market, address indexed account, uint256 outcome, int256 netPosition, uint256 avgPrice, int256 realizedProfit, int256 frozenFunds, int256 realizedCost, uint256 timestamp)
func (_Token *TokenFilterer) FilterProfitLossChanged(opts *bind.FilterOpts, universe []common.Address, market []common.Address, account []common.Address) (*TokenProfitLossChangedIterator, error) {

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

	logs, sub, err := _Token.contract.FilterLogs(opts, "ProfitLossChanged", universeRule, marketRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &TokenProfitLossChangedIterator{contract: _Token.contract, event: "ProfitLossChanged", logs: logs, sub: sub}, nil
}

// WatchProfitLossChanged is a free log subscription operation binding the contract event 0x59543b7f82735782aa5bdb97dff40ff288d4548a5865da513b40e4088e2ee77e.
//
// Solidity: event ProfitLossChanged(address indexed universe, address indexed market, address indexed account, uint256 outcome, int256 netPosition, uint256 avgPrice, int256 realizedProfit, int256 frozenFunds, int256 realizedCost, uint256 timestamp)
func (_Token *TokenFilterer) WatchProfitLossChanged(opts *bind.WatchOpts, sink chan<- *TokenProfitLossChanged, universe []common.Address, market []common.Address, account []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Token.contract.WatchLogs(opts, "ProfitLossChanged", universeRule, marketRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenProfitLossChanged)
				if err := _Token.contract.UnpackLog(event, "ProfitLossChanged", log); err != nil {
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

// ParseProfitLossChanged is a log parse operation binding the contract event 0x59543b7f82735782aa5bdb97dff40ff288d4548a5865da513b40e4088e2ee77e.
//
// Solidity: event ProfitLossChanged(address indexed universe, address indexed market, address indexed account, uint256 outcome, int256 netPosition, uint256 avgPrice, int256 realizedProfit, int256 frozenFunds, int256 realizedCost, uint256 timestamp)
func (_Token *TokenFilterer) ParseProfitLossChanged(log types.Log) (*TokenProfitLossChanged, error) {
	event := new(TokenProfitLossChanged)
	if err := _Token.contract.UnpackLog(event, "ProfitLossChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}
