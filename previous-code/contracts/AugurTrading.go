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

// AugurTradingMetaData contains all meta data concerning the AugurTrading contract.
var AugurTradingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIAugur\",\"name\":\"_augur\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outcome\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"CancelZeroXOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"volume\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"outcomeVolumes\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalTrades\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"MarketVolumeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumAugurTrading.OrderEventType\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"tradeGroupId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addressData\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"uint256Data\",\"type\":\"uint256[]\"}],\"name\":\"OrderEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"universe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outcome\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netPosition\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"avgPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"realizedProfit\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"frozenFunds\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"realizedCost\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ProfitLossChanged\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"augur\",\"outputs\":[{\"internalType\":\"contractIAugur\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIMarket[]\",\"name\":\"_markets\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_shareHolder\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_fingerprint\",\"type\":\"bytes32\"}],\"name\":\"claimMarketsProceeds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_shareHolder\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_fingerprint\",\"type\":\"bytes32\"}],\"name\":\"claimTradingProceeds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"doApprovals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishDeployment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_volume\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_outcomeVolumes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_totalTrades\",\"type\":\"uint256\"}],\"name\":\"logMarketVolumeChanged\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenRefund\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sharesRefund\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"logOrderCanceled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_orderId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_tradeGroupId\",\"type\":\"bytes32\"}],\"name\":\"logOrderCreated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_filler\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountFilled\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_orderId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_tradeGroupId\",\"type\":\"bytes32\"}],\"name\":\"logOrderFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outcome\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_netPosition\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"_avgPrice\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_realizedProfit\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_frozenFunds\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_realizedCost\",\"type\":\"int256\"}],\"name\":\"logProfitLossChanged\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outcome\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_orderHash\",\"type\":\"bytes32\"}],\"name\":\"logZeroXOrderCanceled\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_orderHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_tradeGroupId\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_orderType\",\"type\":\"uint8\"},{\"internalType\":\"address[]\",\"name\":\"_addressData\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_uint256Data\",\"type\":\"uint256[]\"}],\"name\":\"logZeroXOrderFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"lookup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"registerContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"shareToken\",\"outputs\":[{\"internalType\":\"contractIShareToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"trustedSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"uploader\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AugurTradingABI is the input ABI used to generate the binding from.
// Deprecated: Use AugurTradingMetaData.ABI instead.
var AugurTradingABI = AugurTradingMetaData.ABI

// AugurTrading is an auto generated Go binding around an Ethereum contract.
type AugurTrading struct {
	AugurTradingCaller     // Read-only binding to the contract
	AugurTradingTransactor // Write-only binding to the contract
	AugurTradingFilterer   // Log filterer for contract events
}

// AugurTradingCaller is an auto generated read-only Go binding around an Ethereum contract.
type AugurTradingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AugurTradingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AugurTradingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AugurTradingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AugurTradingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AugurTradingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AugurTradingSession struct {
	Contract     *AugurTrading     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AugurTradingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AugurTradingCallerSession struct {
	Contract *AugurTradingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AugurTradingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AugurTradingTransactorSession struct {
	Contract     *AugurTradingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AugurTradingRaw is an auto generated low-level Go binding around an Ethereum contract.
type AugurTradingRaw struct {
	Contract *AugurTrading // Generic contract binding to access the raw methods on
}

// AugurTradingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AugurTradingCallerRaw struct {
	Contract *AugurTradingCaller // Generic read-only contract binding to access the raw methods on
}

// AugurTradingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AugurTradingTransactorRaw struct {
	Contract *AugurTradingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAugurTrading creates a new instance of AugurTrading, bound to a specific deployed contract.
func NewAugurTrading(address common.Address, backend bind.ContractBackend) (*AugurTrading, error) {
	contract, err := bindAugurTrading(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AugurTrading{AugurTradingCaller: AugurTradingCaller{contract: contract}, AugurTradingTransactor: AugurTradingTransactor{contract: contract}, AugurTradingFilterer: AugurTradingFilterer{contract: contract}}, nil
}

// NewAugurTradingCaller creates a new read-only instance of AugurTrading, bound to a specific deployed contract.
func NewAugurTradingCaller(address common.Address, caller bind.ContractCaller) (*AugurTradingCaller, error) {
	contract, err := bindAugurTrading(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AugurTradingCaller{contract: contract}, nil
}

// NewAugurTradingTransactor creates a new write-only instance of AugurTrading, bound to a specific deployed contract.
func NewAugurTradingTransactor(address common.Address, transactor bind.ContractTransactor) (*AugurTradingTransactor, error) {
	contract, err := bindAugurTrading(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AugurTradingTransactor{contract: contract}, nil
}

// NewAugurTradingFilterer creates a new log filterer instance of AugurTrading, bound to a specific deployed contract.
func NewAugurTradingFilterer(address common.Address, filterer bind.ContractFilterer) (*AugurTradingFilterer, error) {
	contract, err := bindAugurTrading(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AugurTradingFilterer{contract: contract}, nil
}

// bindAugurTrading binds a generic wrapper to an already deployed contract.
func bindAugurTrading(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AugurTradingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AugurTrading *AugurTradingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AugurTrading.Contract.AugurTradingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AugurTrading *AugurTradingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AugurTrading.Contract.AugurTradingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AugurTrading *AugurTradingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AugurTrading.Contract.AugurTradingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AugurTrading *AugurTradingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AugurTrading.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AugurTrading *AugurTradingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AugurTrading.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AugurTrading *AugurTradingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AugurTrading.Contract.contract.Transact(opts, method, params...)
}

// Augur is a free data retrieval call binding the contract method 0x7a0d8f8a.
//
// Solidity: function augur() view returns(address)
func (_AugurTrading *AugurTradingCaller) Augur(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AugurTrading.contract.Call(opts, &out, "augur")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Augur is a free data retrieval call binding the contract method 0x7a0d8f8a.
//
// Solidity: function augur() view returns(address)
func (_AugurTrading *AugurTradingSession) Augur() (common.Address, error) {
	return _AugurTrading.Contract.Augur(&_AugurTrading.CallOpts)
}

// Augur is a free data retrieval call binding the contract method 0x7a0d8f8a.
//
// Solidity: function augur() view returns(address)
func (_AugurTrading *AugurTradingCallerSession) Augur() (common.Address, error) {
	return _AugurTrading.Contract.Augur(&_AugurTrading.CallOpts)
}

// Lookup is a free data retrieval call binding the contract method 0xf39ec1f7.
//
// Solidity: function lookup(bytes32 _key) view returns(address)
func (_AugurTrading *AugurTradingCaller) Lookup(opts *bind.CallOpts, _key [32]byte) (common.Address, error) {
	var out []interface{}
	err := _AugurTrading.contract.Call(opts, &out, "lookup", _key)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Lookup is a free data retrieval call binding the contract method 0xf39ec1f7.
//
// Solidity: function lookup(bytes32 _key) view returns(address)
func (_AugurTrading *AugurTradingSession) Lookup(_key [32]byte) (common.Address, error) {
	return _AugurTrading.Contract.Lookup(&_AugurTrading.CallOpts, _key)
}

// Lookup is a free data retrieval call binding the contract method 0xf39ec1f7.
//
// Solidity: function lookup(bytes32 _key) view returns(address)
func (_AugurTrading *AugurTradingCallerSession) Lookup(_key [32]byte) (common.Address, error) {
	return _AugurTrading.Contract.Lookup(&_AugurTrading.CallOpts, _key)
}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() view returns(address)
func (_AugurTrading *AugurTradingCaller) ShareToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AugurTrading.contract.Call(opts, &out, "shareToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() view returns(address)
func (_AugurTrading *AugurTradingSession) ShareToken() (common.Address, error) {
	return _AugurTrading.Contract.ShareToken(&_AugurTrading.CallOpts)
}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() view returns(address)
func (_AugurTrading *AugurTradingCallerSession) ShareToken() (common.Address, error) {
	return _AugurTrading.Contract.ShareToken(&_AugurTrading.CallOpts)
}

// TrustedSender is a free data retrieval call binding the contract method 0x3f6ba415.
//
// Solidity: function trustedSender(address ) view returns(bool)
func (_AugurTrading *AugurTradingCaller) TrustedSender(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AugurTrading.contract.Call(opts, &out, "trustedSender", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TrustedSender is a free data retrieval call binding the contract method 0x3f6ba415.
//
// Solidity: function trustedSender(address ) view returns(bool)
func (_AugurTrading *AugurTradingSession) TrustedSender(arg0 common.Address) (bool, error) {
	return _AugurTrading.Contract.TrustedSender(&_AugurTrading.CallOpts, arg0)
}

// TrustedSender is a free data retrieval call binding the contract method 0x3f6ba415.
//
// Solidity: function trustedSender(address ) view returns(bool)
func (_AugurTrading *AugurTradingCallerSession) TrustedSender(arg0 common.Address) (bool, error) {
	return _AugurTrading.Contract.TrustedSender(&_AugurTrading.CallOpts, arg0)
}

// Uploader is a free data retrieval call binding the contract method 0x65fe2a0b.
//
// Solidity: function uploader() view returns(address)
func (_AugurTrading *AugurTradingCaller) Uploader(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AugurTrading.contract.Call(opts, &out, "uploader")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Uploader is a free data retrieval call binding the contract method 0x65fe2a0b.
//
// Solidity: function uploader() view returns(address)
func (_AugurTrading *AugurTradingSession) Uploader() (common.Address, error) {
	return _AugurTrading.Contract.Uploader(&_AugurTrading.CallOpts)
}

// Uploader is a free data retrieval call binding the contract method 0x65fe2a0b.
//
// Solidity: function uploader() view returns(address)
func (_AugurTrading *AugurTradingCallerSession) Uploader() (common.Address, error) {
	return _AugurTrading.Contract.Uploader(&_AugurTrading.CallOpts)
}

// ClaimMarketsProceeds is a paid mutator transaction binding the contract method 0xdb754422.
//
// Solidity: function claimMarketsProceeds(address[] _markets, address _shareHolder, bytes32 _fingerprint) returns(bool)
func (_AugurTrading *AugurTradingTransactor) ClaimMarketsProceeds(opts *bind.TransactOpts, _markets []common.Address, _shareHolder common.Address, _fingerprint [32]byte) (*types.Transaction, error) {
	return _AugurTrading.contract.Transact(opts, "claimMarketsProceeds", _markets, _shareHolder, _fingerprint)
}

// ClaimMarketsProceeds is a paid mutator transaction binding the contract method 0xdb754422.
//
// Solidity: function claimMarketsProceeds(address[] _markets, address _shareHolder, bytes32 _fingerprint) returns(bool)
func (_AugurTrading *AugurTradingSession) ClaimMarketsProceeds(_markets []common.Address, _shareHolder common.Address, _fingerprint [32]byte) (*types.Transaction, error) {
	return _AugurTrading.Contract.ClaimMarketsProceeds(&_AugurTrading.TransactOpts, _markets, _shareHolder, _fingerprint)
}

// ClaimMarketsProceeds is a paid mutator transaction binding the contract method 0xdb754422.
//
// Solidity: function claimMarketsProceeds(address[] _markets, address _shareHolder, bytes32 _fingerprint) returns(bool)
func (_AugurTrading *AugurTradingTransactorSession) ClaimMarketsProceeds(_markets []common.Address, _shareHolder common.Address, _fingerprint [32]byte) (*types.Transaction, error) {
	return _AugurTrading.Contract.ClaimMarketsProceeds(&_AugurTrading.TransactOpts, _markets, _shareHolder, _fingerprint)
}

// ClaimTradingProceeds is a paid mutator transaction binding the contract method 0xefd342c1.
//
// Solidity: function claimTradingProceeds(address _market, address _shareHolder, bytes32 _fingerprint) returns(bool)
func (_AugurTrading *AugurTradingTransactor) ClaimTradingProceeds(opts *bind.TransactOpts, _market common.Address, _shareHolder common.Address, _fingerprint [32]byte) (*types.Transaction, error) {
	return _AugurTrading.contract.Transact(opts, "claimTradingProceeds", _market, _shareHolder, _fingerprint)
}

// ClaimTradingProceeds is a paid mutator transaction binding the contract method 0xefd342c1.
//
// Solidity: function claimTradingProceeds(address _market, address _shareHolder, bytes32 _fingerprint) returns(bool)
func (_AugurTrading *AugurTradingSession) ClaimTradingProceeds(_market common.Address, _shareHolder common.Address, _fingerprint [32]byte) (*types.Transaction, error) {
	return _AugurTrading.Contract.ClaimTradingProceeds(&_AugurTrading.TransactOpts, _market, _shareHolder, _fingerprint)
}

// ClaimTradingProceeds is a paid mutator transaction binding the contract method 0xefd342c1.
//
// Solidity: function claimTradingProceeds(address _market, address _shareHolder, bytes32 _fingerprint) returns(bool)
func (_AugurTrading *AugurTradingTransactorSession) ClaimTradingProceeds(_market common.Address, _shareHolder common.Address, _fingerprint [32]byte) (*types.Transaction, error) {
	return _AugurTrading.Contract.ClaimTradingProceeds(&_AugurTrading.TransactOpts, _market, _shareHolder, _fingerprint)
}

// DoApprovals is a paid mutator transaction binding the contract method 0x5d4f7742.
//
// Solidity: function doApprovals() returns(bool)
func (_AugurTrading *AugurTradingTransactor) DoApprovals(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AugurTrading.contract.Transact(opts, "doApprovals")
}

// DoApprovals is a paid mutator transaction binding the contract method 0x5d4f7742.
//
// Solidity: function doApprovals() returns(bool)
func (_AugurTrading *AugurTradingSession) DoApprovals() (*types.Transaction, error) {
	return _AugurTrading.Contract.DoApprovals(&_AugurTrading.TransactOpts)
}

// DoApprovals is a paid mutator transaction binding the contract method 0x5d4f7742.
//
// Solidity: function doApprovals() returns(bool)
func (_AugurTrading *AugurTradingTransactorSession) DoApprovals() (*types.Transaction, error) {
	return _AugurTrading.Contract.DoApprovals(&_AugurTrading.TransactOpts)
}

// FinishDeployment is a paid mutator transaction binding the contract method 0xe997fffd.
//
// Solidity: function finishDeployment() returns(bool)
func (_AugurTrading *AugurTradingTransactor) FinishDeployment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AugurTrading.contract.Transact(opts, "finishDeployment")
}

// FinishDeployment is a paid mutator transaction binding the contract method 0xe997fffd.
//
// Solidity: function finishDeployment() returns(bool)
func (_AugurTrading *AugurTradingSession) FinishDeployment() (*types.Transaction, error) {
	return _AugurTrading.Contract.FinishDeployment(&_AugurTrading.TransactOpts)
}

// FinishDeployment is a paid mutator transaction binding the contract method 0xe997fffd.
//
// Solidity: function finishDeployment() returns(bool)
func (_AugurTrading *AugurTradingTransactorSession) FinishDeployment() (*types.Transaction, error) {
	return _AugurTrading.Contract.FinishDeployment(&_AugurTrading.TransactOpts)
}

// LogMarketVolumeChanged is a paid mutator transaction binding the contract method 0x9ae36ee5.
//
// Solidity: function logMarketVolumeChanged(address _universe, address _market, uint256 _volume, uint256[] _outcomeVolumes, uint256 _totalTrades) returns(bool)
func (_AugurTrading *AugurTradingTransactor) LogMarketVolumeChanged(opts *bind.TransactOpts, _universe common.Address, _market common.Address, _volume *big.Int, _outcomeVolumes []*big.Int, _totalTrades *big.Int) (*types.Transaction, error) {
	return _AugurTrading.contract.Transact(opts, "logMarketVolumeChanged", _universe, _market, _volume, _outcomeVolumes, _totalTrades)
}

// LogMarketVolumeChanged is a paid mutator transaction binding the contract method 0x9ae36ee5.
//
// Solidity: function logMarketVolumeChanged(address _universe, address _market, uint256 _volume, uint256[] _outcomeVolumes, uint256 _totalTrades) returns(bool)
func (_AugurTrading *AugurTradingSession) LogMarketVolumeChanged(_universe common.Address, _market common.Address, _volume *big.Int, _outcomeVolumes []*big.Int, _totalTrades *big.Int) (*types.Transaction, error) {
	return _AugurTrading.Contract.LogMarketVolumeChanged(&_AugurTrading.TransactOpts, _universe, _market, _volume, _outcomeVolumes, _totalTrades)
}

// LogMarketVolumeChanged is a paid mutator transaction binding the contract method 0x9ae36ee5.
//
// Solidity: function logMarketVolumeChanged(address _universe, address _market, uint256 _volume, uint256[] _outcomeVolumes, uint256 _totalTrades) returns(bool)
func (_AugurTrading *AugurTradingTransactorSession) LogMarketVolumeChanged(_universe common.Address, _market common.Address, _volume *big.Int, _outcomeVolumes []*big.Int, _totalTrades *big.Int) (*types.Transaction, error) {
	return _AugurTrading.Contract.LogMarketVolumeChanged(&_AugurTrading.TransactOpts, _universe, _market, _volume, _outcomeVolumes, _totalTrades)
}

// LogOrderCanceled is a paid mutator transaction binding the contract method 0xc9e34878.
//
// Solidity: function logOrderCanceled(address _universe, address _market, address _creator, uint256 _tokenRefund, uint256 _sharesRefund, bytes32 _orderId) returns(bool)
func (_AugurTrading *AugurTradingTransactor) LogOrderCanceled(opts *bind.TransactOpts, _universe common.Address, _market common.Address, _creator common.Address, _tokenRefund *big.Int, _sharesRefund *big.Int, _orderId [32]byte) (*types.Transaction, error) {
	return _AugurTrading.contract.Transact(opts, "logOrderCanceled", _universe, _market, _creator, _tokenRefund, _sharesRefund, _orderId)
}

// LogOrderCanceled is a paid mutator transaction binding the contract method 0xc9e34878.
//
// Solidity: function logOrderCanceled(address _universe, address _market, address _creator, uint256 _tokenRefund, uint256 _sharesRefund, bytes32 _orderId) returns(bool)
func (_AugurTrading *AugurTradingSession) LogOrderCanceled(_universe common.Address, _market common.Address, _creator common.Address, _tokenRefund *big.Int, _sharesRefund *big.Int, _orderId [32]byte) (*types.Transaction, error) {
	return _AugurTrading.Contract.LogOrderCanceled(&_AugurTrading.TransactOpts, _universe, _market, _creator, _tokenRefund, _sharesRefund, _orderId)
}

// LogOrderCanceled is a paid mutator transaction binding the contract method 0xc9e34878.
//
// Solidity: function logOrderCanceled(address _universe, address _market, address _creator, uint256 _tokenRefund, uint256 _sharesRefund, bytes32 _orderId) returns(bool)
func (_AugurTrading *AugurTradingTransactorSession) LogOrderCanceled(_universe common.Address, _market common.Address, _creator common.Address, _tokenRefund *big.Int, _sharesRefund *big.Int, _orderId [32]byte) (*types.Transaction, error) {
	return _AugurTrading.Contract.LogOrderCanceled(&_AugurTrading.TransactOpts, _universe, _market, _creator, _tokenRefund, _sharesRefund, _orderId)
}

// LogOrderCreated is a paid mutator transaction binding the contract method 0x98c7cc36.
//
// Solidity: function logOrderCreated(address _universe, bytes32 _orderId, bytes32 _tradeGroupId) returns(bool)
func (_AugurTrading *AugurTradingTransactor) LogOrderCreated(opts *bind.TransactOpts, _universe common.Address, _orderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _AugurTrading.contract.Transact(opts, "logOrderCreated", _universe, _orderId, _tradeGroupId)
}

// LogOrderCreated is a paid mutator transaction binding the contract method 0x98c7cc36.
//
// Solidity: function logOrderCreated(address _universe, bytes32 _orderId, bytes32 _tradeGroupId) returns(bool)
func (_AugurTrading *AugurTradingSession) LogOrderCreated(_universe common.Address, _orderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _AugurTrading.Contract.LogOrderCreated(&_AugurTrading.TransactOpts, _universe, _orderId, _tradeGroupId)
}

// LogOrderCreated is a paid mutator transaction binding the contract method 0x98c7cc36.
//
// Solidity: function logOrderCreated(address _universe, bytes32 _orderId, bytes32 _tradeGroupId) returns(bool)
func (_AugurTrading *AugurTradingTransactorSession) LogOrderCreated(_universe common.Address, _orderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _AugurTrading.Contract.LogOrderCreated(&_AugurTrading.TransactOpts, _universe, _orderId, _tradeGroupId)
}

// LogOrderFilled is a paid mutator transaction binding the contract method 0x5ae5c8e7.
//
// Solidity: function logOrderFilled(address _universe, address _creator, address _filler, uint256 _price, uint256 _fees, uint256 _amountFilled, bytes32 _orderId, bytes32 _tradeGroupId) returns(bool)
func (_AugurTrading *AugurTradingTransactor) LogOrderFilled(opts *bind.TransactOpts, _universe common.Address, _creator common.Address, _filler common.Address, _price *big.Int, _fees *big.Int, _amountFilled *big.Int, _orderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _AugurTrading.contract.Transact(opts, "logOrderFilled", _universe, _creator, _filler, _price, _fees, _amountFilled, _orderId, _tradeGroupId)
}

// LogOrderFilled is a paid mutator transaction binding the contract method 0x5ae5c8e7.
//
// Solidity: function logOrderFilled(address _universe, address _creator, address _filler, uint256 _price, uint256 _fees, uint256 _amountFilled, bytes32 _orderId, bytes32 _tradeGroupId) returns(bool)
func (_AugurTrading *AugurTradingSession) LogOrderFilled(_universe common.Address, _creator common.Address, _filler common.Address, _price *big.Int, _fees *big.Int, _amountFilled *big.Int, _orderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _AugurTrading.Contract.LogOrderFilled(&_AugurTrading.TransactOpts, _universe, _creator, _filler, _price, _fees, _amountFilled, _orderId, _tradeGroupId)
}

// LogOrderFilled is a paid mutator transaction binding the contract method 0x5ae5c8e7.
//
// Solidity: function logOrderFilled(address _universe, address _creator, address _filler, uint256 _price, uint256 _fees, uint256 _amountFilled, bytes32 _orderId, bytes32 _tradeGroupId) returns(bool)
func (_AugurTrading *AugurTradingTransactorSession) LogOrderFilled(_universe common.Address, _creator common.Address, _filler common.Address, _price *big.Int, _fees *big.Int, _amountFilled *big.Int, _orderId [32]byte, _tradeGroupId [32]byte) (*types.Transaction, error) {
	return _AugurTrading.Contract.LogOrderFilled(&_AugurTrading.TransactOpts, _universe, _creator, _filler, _price, _fees, _amountFilled, _orderId, _tradeGroupId)
}

// LogProfitLossChanged is a paid mutator transaction binding the contract method 0x0848a90c.
//
// Solidity: function logProfitLossChanged(address _market, address _account, uint256 _outcome, int256 _netPosition, uint256 _avgPrice, int256 _realizedProfit, int256 _frozenFunds, int256 _realizedCost) returns(bool)
func (_AugurTrading *AugurTradingTransactor) LogProfitLossChanged(opts *bind.TransactOpts, _market common.Address, _account common.Address, _outcome *big.Int, _netPosition *big.Int, _avgPrice *big.Int, _realizedProfit *big.Int, _frozenFunds *big.Int, _realizedCost *big.Int) (*types.Transaction, error) {
	return _AugurTrading.contract.Transact(opts, "logProfitLossChanged", _market, _account, _outcome, _netPosition, _avgPrice, _realizedProfit, _frozenFunds, _realizedCost)
}

// LogProfitLossChanged is a paid mutator transaction binding the contract method 0x0848a90c.
//
// Solidity: function logProfitLossChanged(address _market, address _account, uint256 _outcome, int256 _netPosition, uint256 _avgPrice, int256 _realizedProfit, int256 _frozenFunds, int256 _realizedCost) returns(bool)
func (_AugurTrading *AugurTradingSession) LogProfitLossChanged(_market common.Address, _account common.Address, _outcome *big.Int, _netPosition *big.Int, _avgPrice *big.Int, _realizedProfit *big.Int, _frozenFunds *big.Int, _realizedCost *big.Int) (*types.Transaction, error) {
	return _AugurTrading.Contract.LogProfitLossChanged(&_AugurTrading.TransactOpts, _market, _account, _outcome, _netPosition, _avgPrice, _realizedProfit, _frozenFunds, _realizedCost)
}

// LogProfitLossChanged is a paid mutator transaction binding the contract method 0x0848a90c.
//
// Solidity: function logProfitLossChanged(address _market, address _account, uint256 _outcome, int256 _netPosition, uint256 _avgPrice, int256 _realizedProfit, int256 _frozenFunds, int256 _realizedCost) returns(bool)
func (_AugurTrading *AugurTradingTransactorSession) LogProfitLossChanged(_market common.Address, _account common.Address, _outcome *big.Int, _netPosition *big.Int, _avgPrice *big.Int, _realizedProfit *big.Int, _frozenFunds *big.Int, _realizedCost *big.Int) (*types.Transaction, error) {
	return _AugurTrading.Contract.LogProfitLossChanged(&_AugurTrading.TransactOpts, _market, _account, _outcome, _netPosition, _avgPrice, _realizedProfit, _frozenFunds, _realizedCost)
}

// LogZeroXOrderCanceled is a paid mutator transaction binding the contract method 0x577bb86e.
//
// Solidity: function logZeroXOrderCanceled(address _universe, address _market, address _account, uint256 _outcome, uint256 _price, uint256 _amount, uint8 _type, bytes32 _orderHash) returns()
func (_AugurTrading *AugurTradingTransactor) LogZeroXOrderCanceled(opts *bind.TransactOpts, _universe common.Address, _market common.Address, _account common.Address, _outcome *big.Int, _price *big.Int, _amount *big.Int, _type uint8, _orderHash [32]byte) (*types.Transaction, error) {
	return _AugurTrading.contract.Transact(opts, "logZeroXOrderCanceled", _universe, _market, _account, _outcome, _price, _amount, _type, _orderHash)
}

// LogZeroXOrderCanceled is a paid mutator transaction binding the contract method 0x577bb86e.
//
// Solidity: function logZeroXOrderCanceled(address _universe, address _market, address _account, uint256 _outcome, uint256 _price, uint256 _amount, uint8 _type, bytes32 _orderHash) returns()
func (_AugurTrading *AugurTradingSession) LogZeroXOrderCanceled(_universe common.Address, _market common.Address, _account common.Address, _outcome *big.Int, _price *big.Int, _amount *big.Int, _type uint8, _orderHash [32]byte) (*types.Transaction, error) {
	return _AugurTrading.Contract.LogZeroXOrderCanceled(&_AugurTrading.TransactOpts, _universe, _market, _account, _outcome, _price, _amount, _type, _orderHash)
}

// LogZeroXOrderCanceled is a paid mutator transaction binding the contract method 0x577bb86e.
//
// Solidity: function logZeroXOrderCanceled(address _universe, address _market, address _account, uint256 _outcome, uint256 _price, uint256 _amount, uint8 _type, bytes32 _orderHash) returns()
func (_AugurTrading *AugurTradingTransactorSession) LogZeroXOrderCanceled(_universe common.Address, _market common.Address, _account common.Address, _outcome *big.Int, _price *big.Int, _amount *big.Int, _type uint8, _orderHash [32]byte) (*types.Transaction, error) {
	return _AugurTrading.Contract.LogZeroXOrderCanceled(&_AugurTrading.TransactOpts, _universe, _market, _account, _outcome, _price, _amount, _type, _orderHash)
}

// LogZeroXOrderFilled is a paid mutator transaction binding the contract method 0x99b537a0.
//
// Solidity: function logZeroXOrderFilled(address _universe, address _market, bytes32 _orderHash, bytes32 _tradeGroupId, uint8 _orderType, address[] _addressData, uint256[] _uint256Data) returns(bool)
func (_AugurTrading *AugurTradingTransactor) LogZeroXOrderFilled(opts *bind.TransactOpts, _universe common.Address, _market common.Address, _orderHash [32]byte, _tradeGroupId [32]byte, _orderType uint8, _addressData []common.Address, _uint256Data []*big.Int) (*types.Transaction, error) {
	return _AugurTrading.contract.Transact(opts, "logZeroXOrderFilled", _universe, _market, _orderHash, _tradeGroupId, _orderType, _addressData, _uint256Data)
}

// LogZeroXOrderFilled is a paid mutator transaction binding the contract method 0x99b537a0.
//
// Solidity: function logZeroXOrderFilled(address _universe, address _market, bytes32 _orderHash, bytes32 _tradeGroupId, uint8 _orderType, address[] _addressData, uint256[] _uint256Data) returns(bool)
func (_AugurTrading *AugurTradingSession) LogZeroXOrderFilled(_universe common.Address, _market common.Address, _orderHash [32]byte, _tradeGroupId [32]byte, _orderType uint8, _addressData []common.Address, _uint256Data []*big.Int) (*types.Transaction, error) {
	return _AugurTrading.Contract.LogZeroXOrderFilled(&_AugurTrading.TransactOpts, _universe, _market, _orderHash, _tradeGroupId, _orderType, _addressData, _uint256Data)
}

// LogZeroXOrderFilled is a paid mutator transaction binding the contract method 0x99b537a0.
//
// Solidity: function logZeroXOrderFilled(address _universe, address _market, bytes32 _orderHash, bytes32 _tradeGroupId, uint8 _orderType, address[] _addressData, uint256[] _uint256Data) returns(bool)
func (_AugurTrading *AugurTradingTransactorSession) LogZeroXOrderFilled(_universe common.Address, _market common.Address, _orderHash [32]byte, _tradeGroupId [32]byte, _orderType uint8, _addressData []common.Address, _uint256Data []*big.Int) (*types.Transaction, error) {
	return _AugurTrading.Contract.LogZeroXOrderFilled(&_AugurTrading.TransactOpts, _universe, _market, _orderHash, _tradeGroupId, _orderType, _addressData, _uint256Data)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x645c6fae.
//
// Solidity: function registerContract(bytes32 _key, address _address) returns(bool)
func (_AugurTrading *AugurTradingTransactor) RegisterContract(opts *bind.TransactOpts, _key [32]byte, _address common.Address) (*types.Transaction, error) {
	return _AugurTrading.contract.Transact(opts, "registerContract", _key, _address)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x645c6fae.
//
// Solidity: function registerContract(bytes32 _key, address _address) returns(bool)
func (_AugurTrading *AugurTradingSession) RegisterContract(_key [32]byte, _address common.Address) (*types.Transaction, error) {
	return _AugurTrading.Contract.RegisterContract(&_AugurTrading.TransactOpts, _key, _address)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x645c6fae.
//
// Solidity: function registerContract(bytes32 _key, address _address) returns(bool)
func (_AugurTrading *AugurTradingTransactorSession) RegisterContract(_key [32]byte, _address common.Address) (*types.Transaction, error) {
	return _AugurTrading.Contract.RegisterContract(&_AugurTrading.TransactOpts, _key, _address)
}

// AugurTradingCancelZeroXOrderIterator is returned from FilterCancelZeroXOrder and is used to iterate over the raw logs and unpacked data for CancelZeroXOrder events raised by the AugurTrading contract.
type AugurTradingCancelZeroXOrderIterator struct {
	Event *AugurTradingCancelZeroXOrder // Event containing the contract specifics and raw log

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
func (it *AugurTradingCancelZeroXOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTradingCancelZeroXOrder)
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
		it.Event = new(AugurTradingCancelZeroXOrder)
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
func (it *AugurTradingCancelZeroXOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTradingCancelZeroXOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTradingCancelZeroXOrder represents a CancelZeroXOrder event raised by the AugurTrading contract.
type AugurTradingCancelZeroXOrder struct {
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
func (_AugurTrading *AugurTradingFilterer) FilterCancelZeroXOrder(opts *bind.FilterOpts, universe []common.Address, market []common.Address, account []common.Address) (*AugurTradingCancelZeroXOrderIterator, error) {

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

	logs, sub, err := _AugurTrading.contract.FilterLogs(opts, "CancelZeroXOrder", universeRule, marketRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &AugurTradingCancelZeroXOrderIterator{contract: _AugurTrading.contract, event: "CancelZeroXOrder", logs: logs, sub: sub}, nil
}

// WatchCancelZeroXOrder is a free log subscription operation binding the contract event 0xbe80e5687d7095071b7c4e7a56e0e67bfb9e8a39352f1690fdf74c1ee935c75e.
//
// Solidity: event CancelZeroXOrder(address indexed universe, address indexed market, address indexed account, uint256 outcome, uint256 price, uint256 amount, uint8 orderType, bytes32 orderHash)
func (_AugurTrading *AugurTradingFilterer) WatchCancelZeroXOrder(opts *bind.WatchOpts, sink chan<- *AugurTradingCancelZeroXOrder, universe []common.Address, market []common.Address, account []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AugurTrading.contract.WatchLogs(opts, "CancelZeroXOrder", universeRule, marketRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTradingCancelZeroXOrder)
				if err := _AugurTrading.contract.UnpackLog(event, "CancelZeroXOrder", log); err != nil {
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
func (_AugurTrading *AugurTradingFilterer) ParseCancelZeroXOrder(log types.Log) (*AugurTradingCancelZeroXOrder, error) {
	event := new(AugurTradingCancelZeroXOrder)
	if err := _AugurTrading.contract.UnpackLog(event, "CancelZeroXOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurTradingMarketVolumeChangedIterator is returned from FilterMarketVolumeChanged and is used to iterate over the raw logs and unpacked data for MarketVolumeChanged events raised by the AugurTrading contract.
type AugurTradingMarketVolumeChangedIterator struct {
	Event *AugurTradingMarketVolumeChanged // Event containing the contract specifics and raw log

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
func (it *AugurTradingMarketVolumeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTradingMarketVolumeChanged)
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
		it.Event = new(AugurTradingMarketVolumeChanged)
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
func (it *AugurTradingMarketVolumeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTradingMarketVolumeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTradingMarketVolumeChanged represents a MarketVolumeChanged event raised by the AugurTrading contract.
type AugurTradingMarketVolumeChanged struct {
	Universe       common.Address
	Market         common.Address
	Volume         *big.Int
	OutcomeVolumes []*big.Int
	TotalTrades    *big.Int
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMarketVolumeChanged is a free log retrieval operation binding the contract event 0xcc7cd5af4aead9d3a4a968c74d9063653dccf7110c5ced93fa86b8b03ef5ca24.
//
// Solidity: event MarketVolumeChanged(address indexed universe, address indexed market, uint256 volume, uint256[] outcomeVolumes, uint256 totalTrades, uint256 timestamp)
func (_AugurTrading *AugurTradingFilterer) FilterMarketVolumeChanged(opts *bind.FilterOpts, universe []common.Address, market []common.Address) (*AugurTradingMarketVolumeChangedIterator, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _AugurTrading.contract.FilterLogs(opts, "MarketVolumeChanged", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &AugurTradingMarketVolumeChangedIterator{contract: _AugurTrading.contract, event: "MarketVolumeChanged", logs: logs, sub: sub}, nil
}

// WatchMarketVolumeChanged is a free log subscription operation binding the contract event 0xcc7cd5af4aead9d3a4a968c74d9063653dccf7110c5ced93fa86b8b03ef5ca24.
//
// Solidity: event MarketVolumeChanged(address indexed universe, address indexed market, uint256 volume, uint256[] outcomeVolumes, uint256 totalTrades, uint256 timestamp)
func (_AugurTrading *AugurTradingFilterer) WatchMarketVolumeChanged(opts *bind.WatchOpts, sink chan<- *AugurTradingMarketVolumeChanged, universe []common.Address, market []common.Address) (event.Subscription, error) {

	var universeRule []interface{}
	for _, universeItem := range universe {
		universeRule = append(universeRule, universeItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _AugurTrading.contract.WatchLogs(opts, "MarketVolumeChanged", universeRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTradingMarketVolumeChanged)
				if err := _AugurTrading.contract.UnpackLog(event, "MarketVolumeChanged", log); err != nil {
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

// ParseMarketVolumeChanged is a log parse operation binding the contract event 0xcc7cd5af4aead9d3a4a968c74d9063653dccf7110c5ced93fa86b8b03ef5ca24.
//
// Solidity: event MarketVolumeChanged(address indexed universe, address indexed market, uint256 volume, uint256[] outcomeVolumes, uint256 totalTrades, uint256 timestamp)
func (_AugurTrading *AugurTradingFilterer) ParseMarketVolumeChanged(log types.Log) (*AugurTradingMarketVolumeChanged, error) {
	event := new(AugurTradingMarketVolumeChanged)
	if err := _AugurTrading.contract.UnpackLog(event, "MarketVolumeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurTradingOrderEventIterator is returned from FilterOrderEvent and is used to iterate over the raw logs and unpacked data for OrderEvent events raised by the AugurTrading contract.
type AugurTradingOrderEventIterator struct {
	Event *AugurTradingOrderEvent // Event containing the contract specifics and raw log

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
func (it *AugurTradingOrderEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTradingOrderEvent)
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
		it.Event = new(AugurTradingOrderEvent)
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
func (it *AugurTradingOrderEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTradingOrderEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTradingOrderEvent represents a OrderEvent event raised by the AugurTrading contract.
type AugurTradingOrderEvent struct {
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
func (_AugurTrading *AugurTradingFilterer) FilterOrderEvent(opts *bind.FilterOpts, universe []common.Address, market []common.Address, eventType []uint8) (*AugurTradingOrderEventIterator, error) {

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

	logs, sub, err := _AugurTrading.contract.FilterLogs(opts, "OrderEvent", universeRule, marketRule, eventTypeRule)
	if err != nil {
		return nil, err
	}
	return &AugurTradingOrderEventIterator{contract: _AugurTrading.contract, event: "OrderEvent", logs: logs, sub: sub}, nil
}

// WatchOrderEvent is a free log subscription operation binding the contract event 0x9bab1368a1ed530afaad9c630ba75e6a5c1efa9f6af0139d6cda2b6af6aa801e.
//
// Solidity: event OrderEvent(address indexed universe, address indexed market, uint8 indexed eventType, uint8 orderType, bytes32 orderId, bytes32 tradeGroupId, address[] addressData, uint256[] uint256Data)
func (_AugurTrading *AugurTradingFilterer) WatchOrderEvent(opts *bind.WatchOpts, sink chan<- *AugurTradingOrderEvent, universe []common.Address, market []common.Address, eventType []uint8) (event.Subscription, error) {

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

	logs, sub, err := _AugurTrading.contract.WatchLogs(opts, "OrderEvent", universeRule, marketRule, eventTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTradingOrderEvent)
				if err := _AugurTrading.contract.UnpackLog(event, "OrderEvent", log); err != nil {
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
func (_AugurTrading *AugurTradingFilterer) ParseOrderEvent(log types.Log) (*AugurTradingOrderEvent, error) {
	event := new(AugurTradingOrderEvent)
	if err := _AugurTrading.contract.UnpackLog(event, "OrderEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurTradingProfitLossChangedIterator is returned from FilterProfitLossChanged and is used to iterate over the raw logs and unpacked data for ProfitLossChanged events raised by the AugurTrading contract.
type AugurTradingProfitLossChangedIterator struct {
	Event *AugurTradingProfitLossChanged // Event containing the contract specifics and raw log

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
func (it *AugurTradingProfitLossChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTradingProfitLossChanged)
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
		it.Event = new(AugurTradingProfitLossChanged)
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
func (it *AugurTradingProfitLossChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTradingProfitLossChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTradingProfitLossChanged represents a ProfitLossChanged event raised by the AugurTrading contract.
type AugurTradingProfitLossChanged struct {
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
func (_AugurTrading *AugurTradingFilterer) FilterProfitLossChanged(opts *bind.FilterOpts, universe []common.Address, market []common.Address, account []common.Address) (*AugurTradingProfitLossChangedIterator, error) {

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

	logs, sub, err := _AugurTrading.contract.FilterLogs(opts, "ProfitLossChanged", universeRule, marketRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &AugurTradingProfitLossChangedIterator{contract: _AugurTrading.contract, event: "ProfitLossChanged", logs: logs, sub: sub}, nil
}

// WatchProfitLossChanged is a free log subscription operation binding the contract event 0x59543b7f82735782aa5bdb97dff40ff288d4548a5865da513b40e4088e2ee77e.
//
// Solidity: event ProfitLossChanged(address indexed universe, address indexed market, address indexed account, uint256 outcome, int256 netPosition, uint256 avgPrice, int256 realizedProfit, int256 frozenFunds, int256 realizedCost, uint256 timestamp)
func (_AugurTrading *AugurTradingFilterer) WatchProfitLossChanged(opts *bind.WatchOpts, sink chan<- *AugurTradingProfitLossChanged, universe []common.Address, market []common.Address, account []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AugurTrading.contract.WatchLogs(opts, "ProfitLossChanged", universeRule, marketRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTradingProfitLossChanged)
				if err := _AugurTrading.contract.UnpackLog(event, "ProfitLossChanged", log); err != nil {
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
func (_AugurTrading *AugurTradingFilterer) ParseProfitLossChanged(log types.Log) (*AugurTradingProfitLossChanged, error) {
	event := new(AugurTradingProfitLossChanged)
	if err := _AugurTrading.contract.UnpackLog(event, "ProfitLossChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
