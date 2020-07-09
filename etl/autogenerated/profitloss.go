// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package primitives

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

// ProfitLossABI is the input ABI used to generate the binding from.
const ProfitLossABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outcome\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fees\",\"type\":\"uint256\"}],\"name\":\"adjustTraderProfitForFees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"augur\",\"outputs\":[{\"internalType\":\"contractIAugur\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"augurTrading\",\"outputs\":[{\"internalType\":\"contractIAugurTrading\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cancelOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"createOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fillOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outcome\",\"type\":\"uint256\"}],\"name\":\"getAvgPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outcome\",\"type\":\"uint256\"}],\"name\":\"getFrozenFunds\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outcome\",\"type\":\"uint256\"}],\"name\":\"getNetPosition\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outcome\",\"type\":\"uint256\"}],\"name\":\"getRealizedCost\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outcome\",\"type\":\"uint256\"}],\"name\":\"getRealizedProfit\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIAugur\",\"name\":\"_augur\",\"type\":\"address\"},{\"internalType\":\"contractIAugurTrading\",\"name\":\"_augurTrading\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"orders\",\"outputs\":[{\"internalType\":\"contractIOrders\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_outcomeFees\",\"type\":\"uint256[]\"}],\"name\":\"recordClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outcome\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_frozenFundDelta\",\"type\":\"int256\"}],\"name\":\"recordFrozenFundChange\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIUniverse\",\"name\":\"_universe\",\"type\":\"address\"},{\"internalType\":\"contractIMarket\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_longAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_shortAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outcome\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_amount\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_price\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"_numLongTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numShortTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numLongShares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numShortShares\",\"type\":\"uint256\"}],\"name\":\"recordTrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"shareToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ProfitLoss is an auto generated Go binding around an Ethereum contract.
type ProfitLoss struct {
	ProfitLossCaller     // Read-only binding to the contract
	ProfitLossTransactor // Write-only binding to the contract
	ProfitLossFilterer   // Log filterer for contract events
}

// ProfitLossCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProfitLossCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfitLossTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProfitLossTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfitLossFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProfitLossFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfitLossSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProfitLossSession struct {
	Contract     *ProfitLoss       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProfitLossCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProfitLossCallerSession struct {
	Contract *ProfitLossCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ProfitLossTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProfitLossTransactorSession struct {
	Contract     *ProfitLossTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ProfitLossRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProfitLossRaw struct {
	Contract *ProfitLoss // Generic contract binding to access the raw methods on
}

// ProfitLossCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProfitLossCallerRaw struct {
	Contract *ProfitLossCaller // Generic read-only contract binding to access the raw methods on
}

// ProfitLossTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProfitLossTransactorRaw struct {
	Contract *ProfitLossTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProfitLoss creates a new instance of ProfitLoss, bound to a specific deployed contract.
func NewProfitLoss(address common.Address, backend bind.ContractBackend) (*ProfitLoss, error) {
	contract, err := bindProfitLoss(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProfitLoss{ProfitLossCaller: ProfitLossCaller{contract: contract}, ProfitLossTransactor: ProfitLossTransactor{contract: contract}, ProfitLossFilterer: ProfitLossFilterer{contract: contract}}, nil
}

// NewProfitLossCaller creates a new read-only instance of ProfitLoss, bound to a specific deployed contract.
func NewProfitLossCaller(address common.Address, caller bind.ContractCaller) (*ProfitLossCaller, error) {
	contract, err := bindProfitLoss(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProfitLossCaller{contract: contract}, nil
}

// NewProfitLossTransactor creates a new write-only instance of ProfitLoss, bound to a specific deployed contract.
func NewProfitLossTransactor(address common.Address, transactor bind.ContractTransactor) (*ProfitLossTransactor, error) {
	contract, err := bindProfitLoss(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProfitLossTransactor{contract: contract}, nil
}

// NewProfitLossFilterer creates a new log filterer instance of ProfitLoss, bound to a specific deployed contract.
func NewProfitLossFilterer(address common.Address, filterer bind.ContractFilterer) (*ProfitLossFilterer, error) {
	contract, err := bindProfitLoss(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProfitLossFilterer{contract: contract}, nil
}

// bindProfitLoss binds a generic wrapper to an already deployed contract.
func bindProfitLoss(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProfitLossABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProfitLoss *ProfitLossRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProfitLoss.Contract.ProfitLossCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProfitLoss *ProfitLossRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProfitLoss.Contract.ProfitLossTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProfitLoss *ProfitLossRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProfitLoss.Contract.ProfitLossTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProfitLoss *ProfitLossCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProfitLoss.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProfitLoss *ProfitLossTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProfitLoss.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProfitLoss *ProfitLossTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProfitLoss.Contract.contract.Transact(opts, method, params...)
}

// Augur is a free data retrieval call binding the contract method 0x7a0d8f8a.
//
// Solidity: function augur() view returns(address)
func (_ProfitLoss *ProfitLossCaller) Augur(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ProfitLoss.contract.Call(opts, out, "augur")
	return *ret0, err
}

// Augur is a free data retrieval call binding the contract method 0x7a0d8f8a.
//
// Solidity: function augur() view returns(address)
func (_ProfitLoss *ProfitLossSession) Augur() (common.Address, error) {
	return _ProfitLoss.Contract.Augur(&_ProfitLoss.CallOpts)
}

// Augur is a free data retrieval call binding the contract method 0x7a0d8f8a.
//
// Solidity: function augur() view returns(address)
func (_ProfitLoss *ProfitLossCallerSession) Augur() (common.Address, error) {
	return _ProfitLoss.Contract.Augur(&_ProfitLoss.CallOpts)
}

// AugurTrading is a free data retrieval call binding the contract method 0x8e1bfa73.
//
// Solidity: function augurTrading() view returns(address)
func (_ProfitLoss *ProfitLossCaller) AugurTrading(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ProfitLoss.contract.Call(opts, out, "augurTrading")
	return *ret0, err
}

// AugurTrading is a free data retrieval call binding the contract method 0x8e1bfa73.
//
// Solidity: function augurTrading() view returns(address)
func (_ProfitLoss *ProfitLossSession) AugurTrading() (common.Address, error) {
	return _ProfitLoss.Contract.AugurTrading(&_ProfitLoss.CallOpts)
}

// AugurTrading is a free data retrieval call binding the contract method 0x8e1bfa73.
//
// Solidity: function augurTrading() view returns(address)
func (_ProfitLoss *ProfitLossCallerSession) AugurTrading() (common.Address, error) {
	return _ProfitLoss.Contract.AugurTrading(&_ProfitLoss.CallOpts)
}

// CancelOrder is a free data retrieval call binding the contract method 0x6a816548.
//
// Solidity: function cancelOrder() view returns(address)
func (_ProfitLoss *ProfitLossCaller) CancelOrder(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ProfitLoss.contract.Call(opts, out, "cancelOrder")
	return *ret0, err
}

// CancelOrder is a free data retrieval call binding the contract method 0x6a816548.
//
// Solidity: function cancelOrder() view returns(address)
func (_ProfitLoss *ProfitLossSession) CancelOrder() (common.Address, error) {
	return _ProfitLoss.Contract.CancelOrder(&_ProfitLoss.CallOpts)
}

// CancelOrder is a free data retrieval call binding the contract method 0x6a816548.
//
// Solidity: function cancelOrder() view returns(address)
func (_ProfitLoss *ProfitLossCallerSession) CancelOrder() (common.Address, error) {
	return _ProfitLoss.Contract.CancelOrder(&_ProfitLoss.CallOpts)
}

// CreateOrder is a free data retrieval call binding the contract method 0x6512e6ec.
//
// Solidity: function createOrder() view returns(address)
func (_ProfitLoss *ProfitLossCaller) CreateOrder(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ProfitLoss.contract.Call(opts, out, "createOrder")
	return *ret0, err
}

// CreateOrder is a free data retrieval call binding the contract method 0x6512e6ec.
//
// Solidity: function createOrder() view returns(address)
func (_ProfitLoss *ProfitLossSession) CreateOrder() (common.Address, error) {
	return _ProfitLoss.Contract.CreateOrder(&_ProfitLoss.CallOpts)
}

// CreateOrder is a free data retrieval call binding the contract method 0x6512e6ec.
//
// Solidity: function createOrder() view returns(address)
func (_ProfitLoss *ProfitLossCallerSession) CreateOrder() (common.Address, error) {
	return _ProfitLoss.Contract.CreateOrder(&_ProfitLoss.CallOpts)
}

// FillOrder is a free data retrieval call binding the contract method 0x5c1ad844.
//
// Solidity: function fillOrder() view returns(address)
func (_ProfitLoss *ProfitLossCaller) FillOrder(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ProfitLoss.contract.Call(opts, out, "fillOrder")
	return *ret0, err
}

// FillOrder is a free data retrieval call binding the contract method 0x5c1ad844.
//
// Solidity: function fillOrder() view returns(address)
func (_ProfitLoss *ProfitLossSession) FillOrder() (common.Address, error) {
	return _ProfitLoss.Contract.FillOrder(&_ProfitLoss.CallOpts)
}

// FillOrder is a free data retrieval call binding the contract method 0x5c1ad844.
//
// Solidity: function fillOrder() view returns(address)
func (_ProfitLoss *ProfitLossCallerSession) FillOrder() (common.Address, error) {
	return _ProfitLoss.Contract.FillOrder(&_ProfitLoss.CallOpts)
}

// GetAvgPrice is a free data retrieval call binding the contract method 0x60f6e77e.
//
// Solidity: function getAvgPrice(address _market, address _account, uint256 _outcome) view returns(int256)
func (_ProfitLoss *ProfitLossCaller) GetAvgPrice(opts *bind.CallOpts, _market common.Address, _account common.Address, _outcome *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ProfitLoss.contract.Call(opts, out, "getAvgPrice", _market, _account, _outcome)
	return *ret0, err
}

// GetAvgPrice is a free data retrieval call binding the contract method 0x60f6e77e.
//
// Solidity: function getAvgPrice(address _market, address _account, uint256 _outcome) view returns(int256)
func (_ProfitLoss *ProfitLossSession) GetAvgPrice(_market common.Address, _account common.Address, _outcome *big.Int) (*big.Int, error) {
	return _ProfitLoss.Contract.GetAvgPrice(&_ProfitLoss.CallOpts, _market, _account, _outcome)
}

// GetAvgPrice is a free data retrieval call binding the contract method 0x60f6e77e.
//
// Solidity: function getAvgPrice(address _market, address _account, uint256 _outcome) view returns(int256)
func (_ProfitLoss *ProfitLossCallerSession) GetAvgPrice(_market common.Address, _account common.Address, _outcome *big.Int) (*big.Int, error) {
	return _ProfitLoss.Contract.GetAvgPrice(&_ProfitLoss.CallOpts, _market, _account, _outcome)
}

// GetFrozenFunds is a free data retrieval call binding the contract method 0x7b459892.
//
// Solidity: function getFrozenFunds(address _market, address _account, uint256 _outcome) view returns(int256)
func (_ProfitLoss *ProfitLossCaller) GetFrozenFunds(opts *bind.CallOpts, _market common.Address, _account common.Address, _outcome *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ProfitLoss.contract.Call(opts, out, "getFrozenFunds", _market, _account, _outcome)
	return *ret0, err
}

// GetFrozenFunds is a free data retrieval call binding the contract method 0x7b459892.
//
// Solidity: function getFrozenFunds(address _market, address _account, uint256 _outcome) view returns(int256)
func (_ProfitLoss *ProfitLossSession) GetFrozenFunds(_market common.Address, _account common.Address, _outcome *big.Int) (*big.Int, error) {
	return _ProfitLoss.Contract.GetFrozenFunds(&_ProfitLoss.CallOpts, _market, _account, _outcome)
}

// GetFrozenFunds is a free data retrieval call binding the contract method 0x7b459892.
//
// Solidity: function getFrozenFunds(address _market, address _account, uint256 _outcome) view returns(int256)
func (_ProfitLoss *ProfitLossCallerSession) GetFrozenFunds(_market common.Address, _account common.Address, _outcome *big.Int) (*big.Int, error) {
	return _ProfitLoss.Contract.GetFrozenFunds(&_ProfitLoss.CallOpts, _market, _account, _outcome)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() view returns(bool)
func (_ProfitLoss *ProfitLossCaller) GetInitialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ProfitLoss.contract.Call(opts, out, "getInitialized")
	return *ret0, err
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() view returns(bool)
func (_ProfitLoss *ProfitLossSession) GetInitialized() (bool, error) {
	return _ProfitLoss.Contract.GetInitialized(&_ProfitLoss.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() view returns(bool)
func (_ProfitLoss *ProfitLossCallerSession) GetInitialized() (bool, error) {
	return _ProfitLoss.Contract.GetInitialized(&_ProfitLoss.CallOpts)
}

// GetNetPosition is a free data retrieval call binding the contract method 0xc908ffa7.
//
// Solidity: function getNetPosition(address _market, address _account, uint256 _outcome) view returns(int256)
func (_ProfitLoss *ProfitLossCaller) GetNetPosition(opts *bind.CallOpts, _market common.Address, _account common.Address, _outcome *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ProfitLoss.contract.Call(opts, out, "getNetPosition", _market, _account, _outcome)
	return *ret0, err
}

// GetNetPosition is a free data retrieval call binding the contract method 0xc908ffa7.
//
// Solidity: function getNetPosition(address _market, address _account, uint256 _outcome) view returns(int256)
func (_ProfitLoss *ProfitLossSession) GetNetPosition(_market common.Address, _account common.Address, _outcome *big.Int) (*big.Int, error) {
	return _ProfitLoss.Contract.GetNetPosition(&_ProfitLoss.CallOpts, _market, _account, _outcome)
}

// GetNetPosition is a free data retrieval call binding the contract method 0xc908ffa7.
//
// Solidity: function getNetPosition(address _market, address _account, uint256 _outcome) view returns(int256)
func (_ProfitLoss *ProfitLossCallerSession) GetNetPosition(_market common.Address, _account common.Address, _outcome *big.Int) (*big.Int, error) {
	return _ProfitLoss.Contract.GetNetPosition(&_ProfitLoss.CallOpts, _market, _account, _outcome)
}

// GetRealizedCost is a free data retrieval call binding the contract method 0x714e9eb4.
//
// Solidity: function getRealizedCost(address _market, address _account, uint256 _outcome) view returns(int256)
func (_ProfitLoss *ProfitLossCaller) GetRealizedCost(opts *bind.CallOpts, _market common.Address, _account common.Address, _outcome *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ProfitLoss.contract.Call(opts, out, "getRealizedCost", _market, _account, _outcome)
	return *ret0, err
}

// GetRealizedCost is a free data retrieval call binding the contract method 0x714e9eb4.
//
// Solidity: function getRealizedCost(address _market, address _account, uint256 _outcome) view returns(int256)
func (_ProfitLoss *ProfitLossSession) GetRealizedCost(_market common.Address, _account common.Address, _outcome *big.Int) (*big.Int, error) {
	return _ProfitLoss.Contract.GetRealizedCost(&_ProfitLoss.CallOpts, _market, _account, _outcome)
}

// GetRealizedCost is a free data retrieval call binding the contract method 0x714e9eb4.
//
// Solidity: function getRealizedCost(address _market, address _account, uint256 _outcome) view returns(int256)
func (_ProfitLoss *ProfitLossCallerSession) GetRealizedCost(_market common.Address, _account common.Address, _outcome *big.Int) (*big.Int, error) {
	return _ProfitLoss.Contract.GetRealizedCost(&_ProfitLoss.CallOpts, _market, _account, _outcome)
}

// GetRealizedProfit is a free data retrieval call binding the contract method 0x994ad2e1.
//
// Solidity: function getRealizedProfit(address _market, address _account, uint256 _outcome) view returns(int256)
func (_ProfitLoss *ProfitLossCaller) GetRealizedProfit(opts *bind.CallOpts, _market common.Address, _account common.Address, _outcome *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ProfitLoss.contract.Call(opts, out, "getRealizedProfit", _market, _account, _outcome)
	return *ret0, err
}

// GetRealizedProfit is a free data retrieval call binding the contract method 0x994ad2e1.
//
// Solidity: function getRealizedProfit(address _market, address _account, uint256 _outcome) view returns(int256)
func (_ProfitLoss *ProfitLossSession) GetRealizedProfit(_market common.Address, _account common.Address, _outcome *big.Int) (*big.Int, error) {
	return _ProfitLoss.Contract.GetRealizedProfit(&_ProfitLoss.CallOpts, _market, _account, _outcome)
}

// GetRealizedProfit is a free data retrieval call binding the contract method 0x994ad2e1.
//
// Solidity: function getRealizedProfit(address _market, address _account, uint256 _outcome) view returns(int256)
func (_ProfitLoss *ProfitLossCallerSession) GetRealizedProfit(_market common.Address, _account common.Address, _outcome *big.Int) (*big.Int, error) {
	return _ProfitLoss.Contract.GetRealizedProfit(&_ProfitLoss.CallOpts, _market, _account, _outcome)
}

// Orders is a free data retrieval call binding the contract method 0x4fb764c9.
//
// Solidity: function orders() view returns(address)
func (_ProfitLoss *ProfitLossCaller) Orders(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ProfitLoss.contract.Call(opts, out, "orders")
	return *ret0, err
}

// Orders is a free data retrieval call binding the contract method 0x4fb764c9.
//
// Solidity: function orders() view returns(address)
func (_ProfitLoss *ProfitLossSession) Orders() (common.Address, error) {
	return _ProfitLoss.Contract.Orders(&_ProfitLoss.CallOpts)
}

// Orders is a free data retrieval call binding the contract method 0x4fb764c9.
//
// Solidity: function orders() view returns(address)
func (_ProfitLoss *ProfitLossCallerSession) Orders() (common.Address, error) {
	return _ProfitLoss.Contract.Orders(&_ProfitLoss.CallOpts)
}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() view returns(address)
func (_ProfitLoss *ProfitLossCaller) ShareToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ProfitLoss.contract.Call(opts, out, "shareToken")
	return *ret0, err
}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() view returns(address)
func (_ProfitLoss *ProfitLossSession) ShareToken() (common.Address, error) {
	return _ProfitLoss.Contract.ShareToken(&_ProfitLoss.CallOpts)
}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() view returns(address)
func (_ProfitLoss *ProfitLossCallerSession) ShareToken() (common.Address, error) {
	return _ProfitLoss.Contract.ShareToken(&_ProfitLoss.CallOpts)
}

// AdjustTraderProfitForFees is a paid mutator transaction binding the contract method 0x5310dd31.
//
// Solidity: function adjustTraderProfitForFees(address _market, address _trader, uint256 _outcome, uint256 _fees) returns(bool)
func (_ProfitLoss *ProfitLossTransactor) AdjustTraderProfitForFees(opts *bind.TransactOpts, _market common.Address, _trader common.Address, _outcome *big.Int, _fees *big.Int) (*types.Transaction, error) {
	return _ProfitLoss.contract.Transact(opts, "adjustTraderProfitForFees", _market, _trader, _outcome, _fees)
}

// AdjustTraderProfitForFees is a paid mutator transaction binding the contract method 0x5310dd31.
//
// Solidity: function adjustTraderProfitForFees(address _market, address _trader, uint256 _outcome, uint256 _fees) returns(bool)
func (_ProfitLoss *ProfitLossSession) AdjustTraderProfitForFees(_market common.Address, _trader common.Address, _outcome *big.Int, _fees *big.Int) (*types.Transaction, error) {
	return _ProfitLoss.Contract.AdjustTraderProfitForFees(&_ProfitLoss.TransactOpts, _market, _trader, _outcome, _fees)
}

// AdjustTraderProfitForFees is a paid mutator transaction binding the contract method 0x5310dd31.
//
// Solidity: function adjustTraderProfitForFees(address _market, address _trader, uint256 _outcome, uint256 _fees) returns(bool)
func (_ProfitLoss *ProfitLossTransactorSession) AdjustTraderProfitForFees(_market common.Address, _trader common.Address, _outcome *big.Int, _fees *big.Int) (*types.Transaction, error) {
	return _ProfitLoss.Contract.AdjustTraderProfitForFees(&_ProfitLoss.TransactOpts, _market, _trader, _outcome, _fees)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _augur, address _augurTrading) returns()
func (_ProfitLoss *ProfitLossTransactor) Initialize(opts *bind.TransactOpts, _augur common.Address, _augurTrading common.Address) (*types.Transaction, error) {
	return _ProfitLoss.contract.Transact(opts, "initialize", _augur, _augurTrading)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _augur, address _augurTrading) returns()
func (_ProfitLoss *ProfitLossSession) Initialize(_augur common.Address, _augurTrading common.Address) (*types.Transaction, error) {
	return _ProfitLoss.Contract.Initialize(&_ProfitLoss.TransactOpts, _augur, _augurTrading)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _augur, address _augurTrading) returns()
func (_ProfitLoss *ProfitLossTransactorSession) Initialize(_augur common.Address, _augurTrading common.Address) (*types.Transaction, error) {
	return _ProfitLoss.Contract.Initialize(&_ProfitLoss.TransactOpts, _augur, _augurTrading)
}

// RecordClaim is a paid mutator transaction binding the contract method 0x45a684bb.
//
// Solidity: function recordClaim(address _market, address _account, uint256[] _outcomeFees) returns(bool)
func (_ProfitLoss *ProfitLossTransactor) RecordClaim(opts *bind.TransactOpts, _market common.Address, _account common.Address, _outcomeFees []*big.Int) (*types.Transaction, error) {
	return _ProfitLoss.contract.Transact(opts, "recordClaim", _market, _account, _outcomeFees)
}

// RecordClaim is a paid mutator transaction binding the contract method 0x45a684bb.
//
// Solidity: function recordClaim(address _market, address _account, uint256[] _outcomeFees) returns(bool)
func (_ProfitLoss *ProfitLossSession) RecordClaim(_market common.Address, _account common.Address, _outcomeFees []*big.Int) (*types.Transaction, error) {
	return _ProfitLoss.Contract.RecordClaim(&_ProfitLoss.TransactOpts, _market, _account, _outcomeFees)
}

// RecordClaim is a paid mutator transaction binding the contract method 0x45a684bb.
//
// Solidity: function recordClaim(address _market, address _account, uint256[] _outcomeFees) returns(bool)
func (_ProfitLoss *ProfitLossTransactorSession) RecordClaim(_market common.Address, _account common.Address, _outcomeFees []*big.Int) (*types.Transaction, error) {
	return _ProfitLoss.Contract.RecordClaim(&_ProfitLoss.TransactOpts, _market, _account, _outcomeFees)
}

// RecordFrozenFundChange is a paid mutator transaction binding the contract method 0x2e57cab7.
//
// Solidity: function recordFrozenFundChange(address _universe, address _market, address _account, uint256 _outcome, int256 _frozenFundDelta) returns(bool)
func (_ProfitLoss *ProfitLossTransactor) RecordFrozenFundChange(opts *bind.TransactOpts, _universe common.Address, _market common.Address, _account common.Address, _outcome *big.Int, _frozenFundDelta *big.Int) (*types.Transaction, error) {
	return _ProfitLoss.contract.Transact(opts, "recordFrozenFundChange", _universe, _market, _account, _outcome, _frozenFundDelta)
}

// RecordFrozenFundChange is a paid mutator transaction binding the contract method 0x2e57cab7.
//
// Solidity: function recordFrozenFundChange(address _universe, address _market, address _account, uint256 _outcome, int256 _frozenFundDelta) returns(bool)
func (_ProfitLoss *ProfitLossSession) RecordFrozenFundChange(_universe common.Address, _market common.Address, _account common.Address, _outcome *big.Int, _frozenFundDelta *big.Int) (*types.Transaction, error) {
	return _ProfitLoss.Contract.RecordFrozenFundChange(&_ProfitLoss.TransactOpts, _universe, _market, _account, _outcome, _frozenFundDelta)
}

// RecordFrozenFundChange is a paid mutator transaction binding the contract method 0x2e57cab7.
//
// Solidity: function recordFrozenFundChange(address _universe, address _market, address _account, uint256 _outcome, int256 _frozenFundDelta) returns(bool)
func (_ProfitLoss *ProfitLossTransactorSession) RecordFrozenFundChange(_universe common.Address, _market common.Address, _account common.Address, _outcome *big.Int, _frozenFundDelta *big.Int) (*types.Transaction, error) {
	return _ProfitLoss.Contract.RecordFrozenFundChange(&_ProfitLoss.TransactOpts, _universe, _market, _account, _outcome, _frozenFundDelta)
}

// RecordTrade is a paid mutator transaction binding the contract method 0x12922c94.
//
// Solidity: function recordTrade(address _universe, address _market, address _longAddress, address _shortAddress, uint256 _outcome, int256 _amount, int256 _price, uint256 _numLongTokens, uint256 _numShortTokens, uint256 _numLongShares, uint256 _numShortShares) returns(bool)
func (_ProfitLoss *ProfitLossTransactor) RecordTrade(opts *bind.TransactOpts, _universe common.Address, _market common.Address, _longAddress common.Address, _shortAddress common.Address, _outcome *big.Int, _amount *big.Int, _price *big.Int, _numLongTokens *big.Int, _numShortTokens *big.Int, _numLongShares *big.Int, _numShortShares *big.Int) (*types.Transaction, error) {
	return _ProfitLoss.contract.Transact(opts, "recordTrade", _universe, _market, _longAddress, _shortAddress, _outcome, _amount, _price, _numLongTokens, _numShortTokens, _numLongShares, _numShortShares)
}

// RecordTrade is a paid mutator transaction binding the contract method 0x12922c94.
//
// Solidity: function recordTrade(address _universe, address _market, address _longAddress, address _shortAddress, uint256 _outcome, int256 _amount, int256 _price, uint256 _numLongTokens, uint256 _numShortTokens, uint256 _numLongShares, uint256 _numShortShares) returns(bool)
func (_ProfitLoss *ProfitLossSession) RecordTrade(_universe common.Address, _market common.Address, _longAddress common.Address, _shortAddress common.Address, _outcome *big.Int, _amount *big.Int, _price *big.Int, _numLongTokens *big.Int, _numShortTokens *big.Int, _numLongShares *big.Int, _numShortShares *big.Int) (*types.Transaction, error) {
	return _ProfitLoss.Contract.RecordTrade(&_ProfitLoss.TransactOpts, _universe, _market, _longAddress, _shortAddress, _outcome, _amount, _price, _numLongTokens, _numShortTokens, _numLongShares, _numShortShares)
}

// RecordTrade is a paid mutator transaction binding the contract method 0x12922c94.
//
// Solidity: function recordTrade(address _universe, address _market, address _longAddress, address _shortAddress, uint256 _outcome, int256 _amount, int256 _price, uint256 _numLongTokens, uint256 _numShortTokens, uint256 _numLongShares, uint256 _numShortShares) returns(bool)
func (_ProfitLoss *ProfitLossTransactorSession) RecordTrade(_universe common.Address, _market common.Address, _longAddress common.Address, _shortAddress common.Address, _outcome *big.Int, _amount *big.Int, _price *big.Int, _numLongTokens *big.Int, _numShortTokens *big.Int, _numLongShares *big.Int, _numShortShares *big.Int) (*types.Transaction, error) {
	return _ProfitLoss.Contract.RecordTrade(&_ProfitLoss.TransactOpts, _universe, _market, _longAddress, _shortAddress, _outcome, _amount, _price, _numLongTokens, _numShortTokens, _numLongShares, _numShortShares)
}
