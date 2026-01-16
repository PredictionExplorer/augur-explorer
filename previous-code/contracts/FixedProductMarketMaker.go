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

// FixedProductMarketMakerMetaData contains all meta data concerning the FixedProductMarketMaker contract.
var FixedProductMarketMakerMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"conditionalTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"collateralToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"conditionIds\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"funder\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountsAdded\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"sharesMinted\",\"type\":\"uint256\"}],\"name\":\"FPMMFundingAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"funder\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountsRemoved\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"collateralRemovedFromFeePool\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sharesBurnt\",\"type\":\"uint256\"}],\"name\":\"FPMMFundingRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"investmentAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"outcomeIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"outcomeTokensBought\",\"type\":\"uint256\"}],\"name\":\"FPMMBuy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"outcomeIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"outcomeTokensSold\",\"type\":\"uint256\"}],\"name\":\"FPMMSell\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"collectedFees\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"feesWithdrawableBy\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"withdrawFees\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addedFunds\",\"type\":\"uint256\"},{\"name\":\"distributionHint\",\"type\":\"uint256[]\"}],\"name\":\"addFunding\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sharesToBurn\",\"type\":\"uint256\"}],\"name\":\"removeFunding\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"ids\",\"type\":\"uint256[]\"},{\"name\":\"values\",\"type\":\"uint256[]\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"investmentAmount\",\"type\":\"uint256\"},{\"name\":\"outcomeIndex\",\"type\":\"uint256\"}],\"name\":\"calcBuyAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"name\":\"outcomeIndex\",\"type\":\"uint256\"}],\"name\":\"calcSellAmount\",\"outputs\":[{\"name\":\"outcomeTokenSellAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"investmentAmount\",\"type\":\"uint256\"},{\"name\":\"outcomeIndex\",\"type\":\"uint256\"},{\"name\":\"minOutcomeTokensToBuy\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"name\":\"outcomeIndex\",\"type\":\"uint256\"},{\"name\":\"maxOutcomeTokensToSell\",\"type\":\"uint256\"}],\"name\":\"sell\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FixedProductMarketMakerABI is the input ABI used to generate the binding from.
// Deprecated: Use FixedProductMarketMakerMetaData.ABI instead.
var FixedProductMarketMakerABI = FixedProductMarketMakerMetaData.ABI

// FixedProductMarketMaker is an auto generated Go binding around an Ethereum contract.
type FixedProductMarketMaker struct {
	FixedProductMarketMakerCaller     // Read-only binding to the contract
	FixedProductMarketMakerTransactor // Write-only binding to the contract
	FixedProductMarketMakerFilterer   // Log filterer for contract events
}

// FixedProductMarketMakerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FixedProductMarketMakerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FixedProductMarketMakerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FixedProductMarketMakerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FixedProductMarketMakerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FixedProductMarketMakerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FixedProductMarketMakerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FixedProductMarketMakerSession struct {
	Contract     *FixedProductMarketMaker // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// FixedProductMarketMakerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FixedProductMarketMakerCallerSession struct {
	Contract *FixedProductMarketMakerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// FixedProductMarketMakerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FixedProductMarketMakerTransactorSession struct {
	Contract     *FixedProductMarketMakerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// FixedProductMarketMakerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FixedProductMarketMakerRaw struct {
	Contract *FixedProductMarketMaker // Generic contract binding to access the raw methods on
}

// FixedProductMarketMakerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FixedProductMarketMakerCallerRaw struct {
	Contract *FixedProductMarketMakerCaller // Generic read-only contract binding to access the raw methods on
}

// FixedProductMarketMakerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FixedProductMarketMakerTransactorRaw struct {
	Contract *FixedProductMarketMakerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFixedProductMarketMaker creates a new instance of FixedProductMarketMaker, bound to a specific deployed contract.
func NewFixedProductMarketMaker(address common.Address, backend bind.ContractBackend) (*FixedProductMarketMaker, error) {
	contract, err := bindFixedProductMarketMaker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FixedProductMarketMaker{FixedProductMarketMakerCaller: FixedProductMarketMakerCaller{contract: contract}, FixedProductMarketMakerTransactor: FixedProductMarketMakerTransactor{contract: contract}, FixedProductMarketMakerFilterer: FixedProductMarketMakerFilterer{contract: contract}}, nil
}

// NewFixedProductMarketMakerCaller creates a new read-only instance of FixedProductMarketMaker, bound to a specific deployed contract.
func NewFixedProductMarketMakerCaller(address common.Address, caller bind.ContractCaller) (*FixedProductMarketMakerCaller, error) {
	contract, err := bindFixedProductMarketMaker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FixedProductMarketMakerCaller{contract: contract}, nil
}

// NewFixedProductMarketMakerTransactor creates a new write-only instance of FixedProductMarketMaker, bound to a specific deployed contract.
func NewFixedProductMarketMakerTransactor(address common.Address, transactor bind.ContractTransactor) (*FixedProductMarketMakerTransactor, error) {
	contract, err := bindFixedProductMarketMaker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FixedProductMarketMakerTransactor{contract: contract}, nil
}

// NewFixedProductMarketMakerFilterer creates a new log filterer instance of FixedProductMarketMaker, bound to a specific deployed contract.
func NewFixedProductMarketMakerFilterer(address common.Address, filterer bind.ContractFilterer) (*FixedProductMarketMakerFilterer, error) {
	contract, err := bindFixedProductMarketMaker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FixedProductMarketMakerFilterer{contract: contract}, nil
}

// bindFixedProductMarketMaker binds a generic wrapper to an already deployed contract.
func bindFixedProductMarketMaker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FixedProductMarketMakerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FixedProductMarketMaker *FixedProductMarketMakerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FixedProductMarketMaker.Contract.FixedProductMarketMakerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FixedProductMarketMaker *FixedProductMarketMakerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FixedProductMarketMaker.Contract.FixedProductMarketMakerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FixedProductMarketMaker *FixedProductMarketMakerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FixedProductMarketMaker.Contract.FixedProductMarketMakerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FixedProductMarketMaker *FixedProductMarketMakerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FixedProductMarketMaker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FixedProductMarketMaker *FixedProductMarketMakerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FixedProductMarketMaker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FixedProductMarketMaker *FixedProductMarketMakerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FixedProductMarketMaker.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_FixedProductMarketMaker *FixedProductMarketMakerCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FixedProductMarketMaker.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_FixedProductMarketMaker *FixedProductMarketMakerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _FixedProductMarketMaker.Contract.Allowance(&_FixedProductMarketMaker.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_FixedProductMarketMaker *FixedProductMarketMakerCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _FixedProductMarketMaker.Contract.Allowance(&_FixedProductMarketMaker.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_FixedProductMarketMaker *FixedProductMarketMakerCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FixedProductMarketMaker.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_FixedProductMarketMaker *FixedProductMarketMakerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _FixedProductMarketMaker.Contract.BalanceOf(&_FixedProductMarketMaker.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_FixedProductMarketMaker *FixedProductMarketMakerCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _FixedProductMarketMaker.Contract.BalanceOf(&_FixedProductMarketMaker.CallOpts, account)
}

// CalcBuyAmount is a free data retrieval call binding the contract method 0xf55c79d0.
//
// Solidity: function calcBuyAmount(uint256 investmentAmount, uint256 outcomeIndex) view returns(uint256)
func (_FixedProductMarketMaker *FixedProductMarketMakerCaller) CalcBuyAmount(opts *bind.CallOpts, investmentAmount *big.Int, outcomeIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FixedProductMarketMaker.contract.Call(opts, &out, "calcBuyAmount", investmentAmount, outcomeIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcBuyAmount is a free data retrieval call binding the contract method 0xf55c79d0.
//
// Solidity: function calcBuyAmount(uint256 investmentAmount, uint256 outcomeIndex) view returns(uint256)
func (_FixedProductMarketMaker *FixedProductMarketMakerSession) CalcBuyAmount(investmentAmount *big.Int, outcomeIndex *big.Int) (*big.Int, error) {
	return _FixedProductMarketMaker.Contract.CalcBuyAmount(&_FixedProductMarketMaker.CallOpts, investmentAmount, outcomeIndex)
}

// CalcBuyAmount is a free data retrieval call binding the contract method 0xf55c79d0.
//
// Solidity: function calcBuyAmount(uint256 investmentAmount, uint256 outcomeIndex) view returns(uint256)
func (_FixedProductMarketMaker *FixedProductMarketMakerCallerSession) CalcBuyAmount(investmentAmount *big.Int, outcomeIndex *big.Int) (*big.Int, error) {
	return _FixedProductMarketMaker.Contract.CalcBuyAmount(&_FixedProductMarketMaker.CallOpts, investmentAmount, outcomeIndex)
}

// CalcSellAmount is a free data retrieval call binding the contract method 0x4343116a.
//
// Solidity: function calcSellAmount(uint256 returnAmount, uint256 outcomeIndex) view returns(uint256 outcomeTokenSellAmount)
func (_FixedProductMarketMaker *FixedProductMarketMakerCaller) CalcSellAmount(opts *bind.CallOpts, returnAmount *big.Int, outcomeIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FixedProductMarketMaker.contract.Call(opts, &out, "calcSellAmount", returnAmount, outcomeIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcSellAmount is a free data retrieval call binding the contract method 0x4343116a.
//
// Solidity: function calcSellAmount(uint256 returnAmount, uint256 outcomeIndex) view returns(uint256 outcomeTokenSellAmount)
func (_FixedProductMarketMaker *FixedProductMarketMakerSession) CalcSellAmount(returnAmount *big.Int, outcomeIndex *big.Int) (*big.Int, error) {
	return _FixedProductMarketMaker.Contract.CalcSellAmount(&_FixedProductMarketMaker.CallOpts, returnAmount, outcomeIndex)
}

// CalcSellAmount is a free data retrieval call binding the contract method 0x4343116a.
//
// Solidity: function calcSellAmount(uint256 returnAmount, uint256 outcomeIndex) view returns(uint256 outcomeTokenSellAmount)
func (_FixedProductMarketMaker *FixedProductMarketMakerCallerSession) CalcSellAmount(returnAmount *big.Int, outcomeIndex *big.Int) (*big.Int, error) {
	return _FixedProductMarketMaker.Contract.CalcSellAmount(&_FixedProductMarketMaker.CallOpts, returnAmount, outcomeIndex)
}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_FixedProductMarketMaker *FixedProductMarketMakerCaller) CollateralToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FixedProductMarketMaker.contract.Call(opts, &out, "collateralToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_FixedProductMarketMaker *FixedProductMarketMakerSession) CollateralToken() (common.Address, error) {
	return _FixedProductMarketMaker.Contract.CollateralToken(&_FixedProductMarketMaker.CallOpts)
}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_FixedProductMarketMaker *FixedProductMarketMakerCallerSession) CollateralToken() (common.Address, error) {
	return _FixedProductMarketMaker.Contract.CollateralToken(&_FixedProductMarketMaker.CallOpts)
}

// CollectedFees is a free data retrieval call binding the contract method 0x9003adfe.
//
// Solidity: function collectedFees() view returns(uint256)
func (_FixedProductMarketMaker *FixedProductMarketMakerCaller) CollectedFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FixedProductMarketMaker.contract.Call(opts, &out, "collectedFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollectedFees is a free data retrieval call binding the contract method 0x9003adfe.
//
// Solidity: function collectedFees() view returns(uint256)
func (_FixedProductMarketMaker *FixedProductMarketMakerSession) CollectedFees() (*big.Int, error) {
	return _FixedProductMarketMaker.Contract.CollectedFees(&_FixedProductMarketMaker.CallOpts)
}

// CollectedFees is a free data retrieval call binding the contract method 0x9003adfe.
//
// Solidity: function collectedFees() view returns(uint256)
func (_FixedProductMarketMaker *FixedProductMarketMakerCallerSession) CollectedFees() (*big.Int, error) {
	return _FixedProductMarketMaker.Contract.CollectedFees(&_FixedProductMarketMaker.CallOpts)
}

// ConditionIds is a free data retrieval call binding the contract method 0xd8c55af7.
//
// Solidity: function conditionIds(uint256 ) view returns(bytes32)
func (_FixedProductMarketMaker *FixedProductMarketMakerCaller) ConditionIds(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _FixedProductMarketMaker.contract.Call(opts, &out, "conditionIds", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConditionIds is a free data retrieval call binding the contract method 0xd8c55af7.
//
// Solidity: function conditionIds(uint256 ) view returns(bytes32)
func (_FixedProductMarketMaker *FixedProductMarketMakerSession) ConditionIds(arg0 *big.Int) ([32]byte, error) {
	return _FixedProductMarketMaker.Contract.ConditionIds(&_FixedProductMarketMaker.CallOpts, arg0)
}

// ConditionIds is a free data retrieval call binding the contract method 0xd8c55af7.
//
// Solidity: function conditionIds(uint256 ) view returns(bytes32)
func (_FixedProductMarketMaker *FixedProductMarketMakerCallerSession) ConditionIds(arg0 *big.Int) ([32]byte, error) {
	return _FixedProductMarketMaker.Contract.ConditionIds(&_FixedProductMarketMaker.CallOpts, arg0)
}

// ConditionalTokens is a free data retrieval call binding the contract method 0x5bd9e299.
//
// Solidity: function conditionalTokens() view returns(address)
func (_FixedProductMarketMaker *FixedProductMarketMakerCaller) ConditionalTokens(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FixedProductMarketMaker.contract.Call(opts, &out, "conditionalTokens")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConditionalTokens is a free data retrieval call binding the contract method 0x5bd9e299.
//
// Solidity: function conditionalTokens() view returns(address)
func (_FixedProductMarketMaker *FixedProductMarketMakerSession) ConditionalTokens() (common.Address, error) {
	return _FixedProductMarketMaker.Contract.ConditionalTokens(&_FixedProductMarketMaker.CallOpts)
}

// ConditionalTokens is a free data retrieval call binding the contract method 0x5bd9e299.
//
// Solidity: function conditionalTokens() view returns(address)
func (_FixedProductMarketMaker *FixedProductMarketMakerCallerSession) ConditionalTokens() (common.Address, error) {
	return _FixedProductMarketMaker.Contract.ConditionalTokens(&_FixedProductMarketMaker.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_FixedProductMarketMaker *FixedProductMarketMakerCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FixedProductMarketMaker.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_FixedProductMarketMaker *FixedProductMarketMakerSession) Fee() (*big.Int, error) {
	return _FixedProductMarketMaker.Contract.Fee(&_FixedProductMarketMaker.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_FixedProductMarketMaker *FixedProductMarketMakerCallerSession) Fee() (*big.Int, error) {
	return _FixedProductMarketMaker.Contract.Fee(&_FixedProductMarketMaker.CallOpts)
}

// FeesWithdrawableBy is a free data retrieval call binding the contract method 0x16dbd776.
//
// Solidity: function feesWithdrawableBy(address account) view returns(uint256)
func (_FixedProductMarketMaker *FixedProductMarketMakerCaller) FeesWithdrawableBy(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FixedProductMarketMaker.contract.Call(opts, &out, "feesWithdrawableBy", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeesWithdrawableBy is a free data retrieval call binding the contract method 0x16dbd776.
//
// Solidity: function feesWithdrawableBy(address account) view returns(uint256)
func (_FixedProductMarketMaker *FixedProductMarketMakerSession) FeesWithdrawableBy(account common.Address) (*big.Int, error) {
	return _FixedProductMarketMaker.Contract.FeesWithdrawableBy(&_FixedProductMarketMaker.CallOpts, account)
}

// FeesWithdrawableBy is a free data retrieval call binding the contract method 0x16dbd776.
//
// Solidity: function feesWithdrawableBy(address account) view returns(uint256)
func (_FixedProductMarketMaker *FixedProductMarketMakerCallerSession) FeesWithdrawableBy(account common.Address) (*big.Int, error) {
	return _FixedProductMarketMaker.Contract.FeesWithdrawableBy(&_FixedProductMarketMaker.CallOpts, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FixedProductMarketMaker *FixedProductMarketMakerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _FixedProductMarketMaker.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FixedProductMarketMaker *FixedProductMarketMakerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FixedProductMarketMaker.Contract.SupportsInterface(&_FixedProductMarketMaker.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FixedProductMarketMaker *FixedProductMarketMakerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FixedProductMarketMaker.Contract.SupportsInterface(&_FixedProductMarketMaker.CallOpts, interfaceId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FixedProductMarketMaker *FixedProductMarketMakerCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FixedProductMarketMaker.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FixedProductMarketMaker *FixedProductMarketMakerSession) TotalSupply() (*big.Int, error) {
	return _FixedProductMarketMaker.Contract.TotalSupply(&_FixedProductMarketMaker.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FixedProductMarketMaker *FixedProductMarketMakerCallerSession) TotalSupply() (*big.Int, error) {
	return _FixedProductMarketMaker.Contract.TotalSupply(&_FixedProductMarketMaker.CallOpts)
}

// AddFunding is a paid mutator transaction binding the contract method 0xd5f15a46.
//
// Solidity: function addFunding(uint256 addedFunds, uint256[] distributionHint) returns()
func (_FixedProductMarketMaker *FixedProductMarketMakerTransactor) AddFunding(opts *bind.TransactOpts, addedFunds *big.Int, distributionHint []*big.Int) (*types.Transaction, error) {
	return _FixedProductMarketMaker.contract.Transact(opts, "addFunding", addedFunds, distributionHint)
}

// AddFunding is a paid mutator transaction binding the contract method 0xd5f15a46.
//
// Solidity: function addFunding(uint256 addedFunds, uint256[] distributionHint) returns()
func (_FixedProductMarketMaker *FixedProductMarketMakerSession) AddFunding(addedFunds *big.Int, distributionHint []*big.Int) (*types.Transaction, error) {
	return _FixedProductMarketMaker.Contract.AddFunding(&_FixedProductMarketMaker.TransactOpts, addedFunds, distributionHint)
}

// AddFunding is a paid mutator transaction binding the contract method 0xd5f15a46.
//
// Solidity: function addFunding(uint256 addedFunds, uint256[] distributionHint) returns()
func (_FixedProductMarketMaker *FixedProductMarketMakerTransactorSession) AddFunding(addedFunds *big.Int, distributionHint []*big.Int) (*types.Transaction, error) {
	return _FixedProductMarketMaker.Contract.AddFunding(&_FixedProductMarketMaker.TransactOpts, addedFunds, distributionHint)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_FixedProductMarketMaker *FixedProductMarketMakerTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FixedProductMarketMaker.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_FixedProductMarketMaker *FixedProductMarketMakerSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FixedProductMarketMaker.Contract.Approve(&_FixedProductMarketMaker.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_FixedProductMarketMaker *FixedProductMarketMakerTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FixedProductMarketMaker.Contract.Approve(&_FixedProductMarketMaker.TransactOpts, spender, amount)
}

// Buy is a paid mutator transaction binding the contract method 0x40993b26.
//
// Solidity: function buy(uint256 investmentAmount, uint256 outcomeIndex, uint256 minOutcomeTokensToBuy) returns()
func (_FixedProductMarketMaker *FixedProductMarketMakerTransactor) Buy(opts *bind.TransactOpts, investmentAmount *big.Int, outcomeIndex *big.Int, minOutcomeTokensToBuy *big.Int) (*types.Transaction, error) {
	return _FixedProductMarketMaker.contract.Transact(opts, "buy", investmentAmount, outcomeIndex, minOutcomeTokensToBuy)
}

// Buy is a paid mutator transaction binding the contract method 0x40993b26.
//
// Solidity: function buy(uint256 investmentAmount, uint256 outcomeIndex, uint256 minOutcomeTokensToBuy) returns()
func (_FixedProductMarketMaker *FixedProductMarketMakerSession) Buy(investmentAmount *big.Int, outcomeIndex *big.Int, minOutcomeTokensToBuy *big.Int) (*types.Transaction, error) {
	return _FixedProductMarketMaker.Contract.Buy(&_FixedProductMarketMaker.TransactOpts, investmentAmount, outcomeIndex, minOutcomeTokensToBuy)
}

// Buy is a paid mutator transaction binding the contract method 0x40993b26.
//
// Solidity: function buy(uint256 investmentAmount, uint256 outcomeIndex, uint256 minOutcomeTokensToBuy) returns()
func (_FixedProductMarketMaker *FixedProductMarketMakerTransactorSession) Buy(investmentAmount *big.Int, outcomeIndex *big.Int, minOutcomeTokensToBuy *big.Int) (*types.Transaction, error) {
	return _FixedProductMarketMaker.Contract.Buy(&_FixedProductMarketMaker.TransactOpts, investmentAmount, outcomeIndex, minOutcomeTokensToBuy)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_FixedProductMarketMaker *FixedProductMarketMakerTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _FixedProductMarketMaker.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_FixedProductMarketMaker *FixedProductMarketMakerSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _FixedProductMarketMaker.Contract.DecreaseAllowance(&_FixedProductMarketMaker.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_FixedProductMarketMaker *FixedProductMarketMakerTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _FixedProductMarketMaker.Contract.DecreaseAllowance(&_FixedProductMarketMaker.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_FixedProductMarketMaker *FixedProductMarketMakerTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _FixedProductMarketMaker.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_FixedProductMarketMaker *FixedProductMarketMakerSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _FixedProductMarketMaker.Contract.IncreaseAllowance(&_FixedProductMarketMaker.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_FixedProductMarketMaker *FixedProductMarketMakerTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _FixedProductMarketMaker.Contract.IncreaseAllowance(&_FixedProductMarketMaker.TransactOpts, spender, addedValue)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) returns(bytes4)
func (_FixedProductMarketMaker *FixedProductMarketMakerTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _FixedProductMarketMaker.contract.Transact(opts, "onERC1155BatchReceived", operator, from, ids, values, data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) returns(bytes4)
func (_FixedProductMarketMaker *FixedProductMarketMakerSession) OnERC1155BatchReceived(operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _FixedProductMarketMaker.Contract.OnERC1155BatchReceived(&_FixedProductMarketMaker.TransactOpts, operator, from, ids, values, data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) returns(bytes4)
func (_FixedProductMarketMaker *FixedProductMarketMakerTransactorSession) OnERC1155BatchReceived(operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _FixedProductMarketMaker.Contract.OnERC1155BatchReceived(&_FixedProductMarketMaker.TransactOpts, operator, from, ids, values, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) returns(bytes4)
func (_FixedProductMarketMaker *FixedProductMarketMakerTransactor) OnERC1155Received(opts *bind.TransactOpts, operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _FixedProductMarketMaker.contract.Transact(opts, "onERC1155Received", operator, from, id, value, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) returns(bytes4)
func (_FixedProductMarketMaker *FixedProductMarketMakerSession) OnERC1155Received(operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _FixedProductMarketMaker.Contract.OnERC1155Received(&_FixedProductMarketMaker.TransactOpts, operator, from, id, value, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) returns(bytes4)
func (_FixedProductMarketMaker *FixedProductMarketMakerTransactorSession) OnERC1155Received(operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _FixedProductMarketMaker.Contract.OnERC1155Received(&_FixedProductMarketMaker.TransactOpts, operator, from, id, value, data)
}

// RemoveFunding is a paid mutator transaction binding the contract method 0xe03031a6.
//
// Solidity: function removeFunding(uint256 sharesToBurn) returns()
func (_FixedProductMarketMaker *FixedProductMarketMakerTransactor) RemoveFunding(opts *bind.TransactOpts, sharesToBurn *big.Int) (*types.Transaction, error) {
	return _FixedProductMarketMaker.contract.Transact(opts, "removeFunding", sharesToBurn)
}

// RemoveFunding is a paid mutator transaction binding the contract method 0xe03031a6.
//
// Solidity: function removeFunding(uint256 sharesToBurn) returns()
func (_FixedProductMarketMaker *FixedProductMarketMakerSession) RemoveFunding(sharesToBurn *big.Int) (*types.Transaction, error) {
	return _FixedProductMarketMaker.Contract.RemoveFunding(&_FixedProductMarketMaker.TransactOpts, sharesToBurn)
}

// RemoveFunding is a paid mutator transaction binding the contract method 0xe03031a6.
//
// Solidity: function removeFunding(uint256 sharesToBurn) returns()
func (_FixedProductMarketMaker *FixedProductMarketMakerTransactorSession) RemoveFunding(sharesToBurn *big.Int) (*types.Transaction, error) {
	return _FixedProductMarketMaker.Contract.RemoveFunding(&_FixedProductMarketMaker.TransactOpts, sharesToBurn)
}

// Sell is a paid mutator transaction binding the contract method 0xd3c9727c.
//
// Solidity: function sell(uint256 returnAmount, uint256 outcomeIndex, uint256 maxOutcomeTokensToSell) returns()
func (_FixedProductMarketMaker *FixedProductMarketMakerTransactor) Sell(opts *bind.TransactOpts, returnAmount *big.Int, outcomeIndex *big.Int, maxOutcomeTokensToSell *big.Int) (*types.Transaction, error) {
	return _FixedProductMarketMaker.contract.Transact(opts, "sell", returnAmount, outcomeIndex, maxOutcomeTokensToSell)
}

// Sell is a paid mutator transaction binding the contract method 0xd3c9727c.
//
// Solidity: function sell(uint256 returnAmount, uint256 outcomeIndex, uint256 maxOutcomeTokensToSell) returns()
func (_FixedProductMarketMaker *FixedProductMarketMakerSession) Sell(returnAmount *big.Int, outcomeIndex *big.Int, maxOutcomeTokensToSell *big.Int) (*types.Transaction, error) {
	return _FixedProductMarketMaker.Contract.Sell(&_FixedProductMarketMaker.TransactOpts, returnAmount, outcomeIndex, maxOutcomeTokensToSell)
}

// Sell is a paid mutator transaction binding the contract method 0xd3c9727c.
//
// Solidity: function sell(uint256 returnAmount, uint256 outcomeIndex, uint256 maxOutcomeTokensToSell) returns()
func (_FixedProductMarketMaker *FixedProductMarketMakerTransactorSession) Sell(returnAmount *big.Int, outcomeIndex *big.Int, maxOutcomeTokensToSell *big.Int) (*types.Transaction, error) {
	return _FixedProductMarketMaker.Contract.Sell(&_FixedProductMarketMaker.TransactOpts, returnAmount, outcomeIndex, maxOutcomeTokensToSell)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_FixedProductMarketMaker *FixedProductMarketMakerTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FixedProductMarketMaker.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_FixedProductMarketMaker *FixedProductMarketMakerSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FixedProductMarketMaker.Contract.Transfer(&_FixedProductMarketMaker.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_FixedProductMarketMaker *FixedProductMarketMakerTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FixedProductMarketMaker.Contract.Transfer(&_FixedProductMarketMaker.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_FixedProductMarketMaker *FixedProductMarketMakerTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FixedProductMarketMaker.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_FixedProductMarketMaker *FixedProductMarketMakerSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FixedProductMarketMaker.Contract.TransferFrom(&_FixedProductMarketMaker.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_FixedProductMarketMaker *FixedProductMarketMakerTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FixedProductMarketMaker.Contract.TransferFrom(&_FixedProductMarketMaker.TransactOpts, sender, recipient, amount)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x164e68de.
//
// Solidity: function withdrawFees(address account) returns()
func (_FixedProductMarketMaker *FixedProductMarketMakerTransactor) WithdrawFees(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _FixedProductMarketMaker.contract.Transact(opts, "withdrawFees", account)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x164e68de.
//
// Solidity: function withdrawFees(address account) returns()
func (_FixedProductMarketMaker *FixedProductMarketMakerSession) WithdrawFees(account common.Address) (*types.Transaction, error) {
	return _FixedProductMarketMaker.Contract.WithdrawFees(&_FixedProductMarketMaker.TransactOpts, account)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x164e68de.
//
// Solidity: function withdrawFees(address account) returns()
func (_FixedProductMarketMaker *FixedProductMarketMakerTransactorSession) WithdrawFees(account common.Address) (*types.Transaction, error) {
	return _FixedProductMarketMaker.Contract.WithdrawFees(&_FixedProductMarketMaker.TransactOpts, account)
}

// FixedProductMarketMakerApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the FixedProductMarketMaker contract.
type FixedProductMarketMakerApprovalIterator struct {
	Event *FixedProductMarketMakerApproval // Event containing the contract specifics and raw log

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
func (it *FixedProductMarketMakerApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FixedProductMarketMakerApproval)
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
		it.Event = new(FixedProductMarketMakerApproval)
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
func (it *FixedProductMarketMakerApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FixedProductMarketMakerApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FixedProductMarketMakerApproval represents a Approval event raised by the FixedProductMarketMaker contract.
type FixedProductMarketMakerApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_FixedProductMarketMaker *FixedProductMarketMakerFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*FixedProductMarketMakerApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FixedProductMarketMaker.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &FixedProductMarketMakerApprovalIterator{contract: _FixedProductMarketMaker.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_FixedProductMarketMaker *FixedProductMarketMakerFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FixedProductMarketMakerApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FixedProductMarketMaker.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FixedProductMarketMakerApproval)
				if err := _FixedProductMarketMaker.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_FixedProductMarketMaker *FixedProductMarketMakerFilterer) ParseApproval(log types.Log) (*FixedProductMarketMakerApproval, error) {
	event := new(FixedProductMarketMakerApproval)
	if err := _FixedProductMarketMaker.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FixedProductMarketMakerFPMMBuyIterator is returned from FilterFPMMBuy and is used to iterate over the raw logs and unpacked data for FPMMBuy events raised by the FixedProductMarketMaker contract.
type FixedProductMarketMakerFPMMBuyIterator struct {
	Event *FixedProductMarketMakerFPMMBuy // Event containing the contract specifics and raw log

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
func (it *FixedProductMarketMakerFPMMBuyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FixedProductMarketMakerFPMMBuy)
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
		it.Event = new(FixedProductMarketMakerFPMMBuy)
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
func (it *FixedProductMarketMakerFPMMBuyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FixedProductMarketMakerFPMMBuyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FixedProductMarketMakerFPMMBuy represents a FPMMBuy event raised by the FixedProductMarketMaker contract.
type FixedProductMarketMakerFPMMBuy struct {
	Buyer               common.Address
	InvestmentAmount    *big.Int
	FeeAmount           *big.Int
	OutcomeIndex        *big.Int
	OutcomeTokensBought *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterFPMMBuy is a free log retrieval operation binding the contract event 0x4f62630f51608fc8a7603a9391a5101e58bd7c276139366fc107dc3b67c3dcf8.
//
// Solidity: event FPMMBuy(address indexed buyer, uint256 investmentAmount, uint256 feeAmount, uint256 indexed outcomeIndex, uint256 outcomeTokensBought)
func (_FixedProductMarketMaker *FixedProductMarketMakerFilterer) FilterFPMMBuy(opts *bind.FilterOpts, buyer []common.Address, outcomeIndex []*big.Int) (*FixedProductMarketMakerFPMMBuyIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	var outcomeIndexRule []interface{}
	for _, outcomeIndexItem := range outcomeIndex {
		outcomeIndexRule = append(outcomeIndexRule, outcomeIndexItem)
	}

	logs, sub, err := _FixedProductMarketMaker.contract.FilterLogs(opts, "FPMMBuy", buyerRule, outcomeIndexRule)
	if err != nil {
		return nil, err
	}
	return &FixedProductMarketMakerFPMMBuyIterator{contract: _FixedProductMarketMaker.contract, event: "FPMMBuy", logs: logs, sub: sub}, nil
}

// WatchFPMMBuy is a free log subscription operation binding the contract event 0x4f62630f51608fc8a7603a9391a5101e58bd7c276139366fc107dc3b67c3dcf8.
//
// Solidity: event FPMMBuy(address indexed buyer, uint256 investmentAmount, uint256 feeAmount, uint256 indexed outcomeIndex, uint256 outcomeTokensBought)
func (_FixedProductMarketMaker *FixedProductMarketMakerFilterer) WatchFPMMBuy(opts *bind.WatchOpts, sink chan<- *FixedProductMarketMakerFPMMBuy, buyer []common.Address, outcomeIndex []*big.Int) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	var outcomeIndexRule []interface{}
	for _, outcomeIndexItem := range outcomeIndex {
		outcomeIndexRule = append(outcomeIndexRule, outcomeIndexItem)
	}

	logs, sub, err := _FixedProductMarketMaker.contract.WatchLogs(opts, "FPMMBuy", buyerRule, outcomeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FixedProductMarketMakerFPMMBuy)
				if err := _FixedProductMarketMaker.contract.UnpackLog(event, "FPMMBuy", log); err != nil {
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

// ParseFPMMBuy is a log parse operation binding the contract event 0x4f62630f51608fc8a7603a9391a5101e58bd7c276139366fc107dc3b67c3dcf8.
//
// Solidity: event FPMMBuy(address indexed buyer, uint256 investmentAmount, uint256 feeAmount, uint256 indexed outcomeIndex, uint256 outcomeTokensBought)
func (_FixedProductMarketMaker *FixedProductMarketMakerFilterer) ParseFPMMBuy(log types.Log) (*FixedProductMarketMakerFPMMBuy, error) {
	event := new(FixedProductMarketMakerFPMMBuy)
	if err := _FixedProductMarketMaker.contract.UnpackLog(event, "FPMMBuy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FixedProductMarketMakerFPMMFundingAddedIterator is returned from FilterFPMMFundingAdded and is used to iterate over the raw logs and unpacked data for FPMMFundingAdded events raised by the FixedProductMarketMaker contract.
type FixedProductMarketMakerFPMMFundingAddedIterator struct {
	Event *FixedProductMarketMakerFPMMFundingAdded // Event containing the contract specifics and raw log

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
func (it *FixedProductMarketMakerFPMMFundingAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FixedProductMarketMakerFPMMFundingAdded)
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
		it.Event = new(FixedProductMarketMakerFPMMFundingAdded)
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
func (it *FixedProductMarketMakerFPMMFundingAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FixedProductMarketMakerFPMMFundingAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FixedProductMarketMakerFPMMFundingAdded represents a FPMMFundingAdded event raised by the FixedProductMarketMaker contract.
type FixedProductMarketMakerFPMMFundingAdded struct {
	Funder       common.Address
	AmountsAdded []*big.Int
	SharesMinted *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFPMMFundingAdded is a free log retrieval operation binding the contract event 0xec2dc3e5a3bb9aa0a1deb905d2bd23640d07f107e6ceb484024501aad964a951.
//
// Solidity: event FPMMFundingAdded(address indexed funder, uint256[] amountsAdded, uint256 sharesMinted)
func (_FixedProductMarketMaker *FixedProductMarketMakerFilterer) FilterFPMMFundingAdded(opts *bind.FilterOpts, funder []common.Address) (*FixedProductMarketMakerFPMMFundingAddedIterator, error) {

	var funderRule []interface{}
	for _, funderItem := range funder {
		funderRule = append(funderRule, funderItem)
	}

	logs, sub, err := _FixedProductMarketMaker.contract.FilterLogs(opts, "FPMMFundingAdded", funderRule)
	if err != nil {
		return nil, err
	}
	return &FixedProductMarketMakerFPMMFundingAddedIterator{contract: _FixedProductMarketMaker.contract, event: "FPMMFundingAdded", logs: logs, sub: sub}, nil
}

// WatchFPMMFundingAdded is a free log subscription operation binding the contract event 0xec2dc3e5a3bb9aa0a1deb905d2bd23640d07f107e6ceb484024501aad964a951.
//
// Solidity: event FPMMFundingAdded(address indexed funder, uint256[] amountsAdded, uint256 sharesMinted)
func (_FixedProductMarketMaker *FixedProductMarketMakerFilterer) WatchFPMMFundingAdded(opts *bind.WatchOpts, sink chan<- *FixedProductMarketMakerFPMMFundingAdded, funder []common.Address) (event.Subscription, error) {

	var funderRule []interface{}
	for _, funderItem := range funder {
		funderRule = append(funderRule, funderItem)
	}

	logs, sub, err := _FixedProductMarketMaker.contract.WatchLogs(opts, "FPMMFundingAdded", funderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FixedProductMarketMakerFPMMFundingAdded)
				if err := _FixedProductMarketMaker.contract.UnpackLog(event, "FPMMFundingAdded", log); err != nil {
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

// ParseFPMMFundingAdded is a log parse operation binding the contract event 0xec2dc3e5a3bb9aa0a1deb905d2bd23640d07f107e6ceb484024501aad964a951.
//
// Solidity: event FPMMFundingAdded(address indexed funder, uint256[] amountsAdded, uint256 sharesMinted)
func (_FixedProductMarketMaker *FixedProductMarketMakerFilterer) ParseFPMMFundingAdded(log types.Log) (*FixedProductMarketMakerFPMMFundingAdded, error) {
	event := new(FixedProductMarketMakerFPMMFundingAdded)
	if err := _FixedProductMarketMaker.contract.UnpackLog(event, "FPMMFundingAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FixedProductMarketMakerFPMMFundingRemovedIterator is returned from FilterFPMMFundingRemoved and is used to iterate over the raw logs and unpacked data for FPMMFundingRemoved events raised by the FixedProductMarketMaker contract.
type FixedProductMarketMakerFPMMFundingRemovedIterator struct {
	Event *FixedProductMarketMakerFPMMFundingRemoved // Event containing the contract specifics and raw log

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
func (it *FixedProductMarketMakerFPMMFundingRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FixedProductMarketMakerFPMMFundingRemoved)
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
		it.Event = new(FixedProductMarketMakerFPMMFundingRemoved)
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
func (it *FixedProductMarketMakerFPMMFundingRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FixedProductMarketMakerFPMMFundingRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FixedProductMarketMakerFPMMFundingRemoved represents a FPMMFundingRemoved event raised by the FixedProductMarketMaker contract.
type FixedProductMarketMakerFPMMFundingRemoved struct {
	Funder                       common.Address
	AmountsRemoved               []*big.Int
	CollateralRemovedFromFeePool *big.Int
	SharesBurnt                  *big.Int
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterFPMMFundingRemoved is a free log retrieval operation binding the contract event 0x8b4b2c8ebd04c47fc8bce136a85df9b93fcb1f47c8aa296457d4391519d190e7.
//
// Solidity: event FPMMFundingRemoved(address indexed funder, uint256[] amountsRemoved, uint256 collateralRemovedFromFeePool, uint256 sharesBurnt)
func (_FixedProductMarketMaker *FixedProductMarketMakerFilterer) FilterFPMMFundingRemoved(opts *bind.FilterOpts, funder []common.Address) (*FixedProductMarketMakerFPMMFundingRemovedIterator, error) {

	var funderRule []interface{}
	for _, funderItem := range funder {
		funderRule = append(funderRule, funderItem)
	}

	logs, sub, err := _FixedProductMarketMaker.contract.FilterLogs(opts, "FPMMFundingRemoved", funderRule)
	if err != nil {
		return nil, err
	}
	return &FixedProductMarketMakerFPMMFundingRemovedIterator{contract: _FixedProductMarketMaker.contract, event: "FPMMFundingRemoved", logs: logs, sub: sub}, nil
}

// WatchFPMMFundingRemoved is a free log subscription operation binding the contract event 0x8b4b2c8ebd04c47fc8bce136a85df9b93fcb1f47c8aa296457d4391519d190e7.
//
// Solidity: event FPMMFundingRemoved(address indexed funder, uint256[] amountsRemoved, uint256 collateralRemovedFromFeePool, uint256 sharesBurnt)
func (_FixedProductMarketMaker *FixedProductMarketMakerFilterer) WatchFPMMFundingRemoved(opts *bind.WatchOpts, sink chan<- *FixedProductMarketMakerFPMMFundingRemoved, funder []common.Address) (event.Subscription, error) {

	var funderRule []interface{}
	for _, funderItem := range funder {
		funderRule = append(funderRule, funderItem)
	}

	logs, sub, err := _FixedProductMarketMaker.contract.WatchLogs(opts, "FPMMFundingRemoved", funderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FixedProductMarketMakerFPMMFundingRemoved)
				if err := _FixedProductMarketMaker.contract.UnpackLog(event, "FPMMFundingRemoved", log); err != nil {
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

// ParseFPMMFundingRemoved is a log parse operation binding the contract event 0x8b4b2c8ebd04c47fc8bce136a85df9b93fcb1f47c8aa296457d4391519d190e7.
//
// Solidity: event FPMMFundingRemoved(address indexed funder, uint256[] amountsRemoved, uint256 collateralRemovedFromFeePool, uint256 sharesBurnt)
func (_FixedProductMarketMaker *FixedProductMarketMakerFilterer) ParseFPMMFundingRemoved(log types.Log) (*FixedProductMarketMakerFPMMFundingRemoved, error) {
	event := new(FixedProductMarketMakerFPMMFundingRemoved)
	if err := _FixedProductMarketMaker.contract.UnpackLog(event, "FPMMFundingRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FixedProductMarketMakerFPMMSellIterator is returned from FilterFPMMSell and is used to iterate over the raw logs and unpacked data for FPMMSell events raised by the FixedProductMarketMaker contract.
type FixedProductMarketMakerFPMMSellIterator struct {
	Event *FixedProductMarketMakerFPMMSell // Event containing the contract specifics and raw log

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
func (it *FixedProductMarketMakerFPMMSellIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FixedProductMarketMakerFPMMSell)
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
		it.Event = new(FixedProductMarketMakerFPMMSell)
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
func (it *FixedProductMarketMakerFPMMSellIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FixedProductMarketMakerFPMMSellIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FixedProductMarketMakerFPMMSell represents a FPMMSell event raised by the FixedProductMarketMaker contract.
type FixedProductMarketMakerFPMMSell struct {
	Seller            common.Address
	ReturnAmount      *big.Int
	FeeAmount         *big.Int
	OutcomeIndex      *big.Int
	OutcomeTokensSold *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterFPMMSell is a free log retrieval operation binding the contract event 0xadcf2a240ed9300d681d9a3f5382b6c1beed1b7e46643e0c7b42cbe6e2d766b4.
//
// Solidity: event FPMMSell(address indexed seller, uint256 returnAmount, uint256 feeAmount, uint256 indexed outcomeIndex, uint256 outcomeTokensSold)
func (_FixedProductMarketMaker *FixedProductMarketMakerFilterer) FilterFPMMSell(opts *bind.FilterOpts, seller []common.Address, outcomeIndex []*big.Int) (*FixedProductMarketMakerFPMMSellIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	var outcomeIndexRule []interface{}
	for _, outcomeIndexItem := range outcomeIndex {
		outcomeIndexRule = append(outcomeIndexRule, outcomeIndexItem)
	}

	logs, sub, err := _FixedProductMarketMaker.contract.FilterLogs(opts, "FPMMSell", sellerRule, outcomeIndexRule)
	if err != nil {
		return nil, err
	}
	return &FixedProductMarketMakerFPMMSellIterator{contract: _FixedProductMarketMaker.contract, event: "FPMMSell", logs: logs, sub: sub}, nil
}

// WatchFPMMSell is a free log subscription operation binding the contract event 0xadcf2a240ed9300d681d9a3f5382b6c1beed1b7e46643e0c7b42cbe6e2d766b4.
//
// Solidity: event FPMMSell(address indexed seller, uint256 returnAmount, uint256 feeAmount, uint256 indexed outcomeIndex, uint256 outcomeTokensSold)
func (_FixedProductMarketMaker *FixedProductMarketMakerFilterer) WatchFPMMSell(opts *bind.WatchOpts, sink chan<- *FixedProductMarketMakerFPMMSell, seller []common.Address, outcomeIndex []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	var outcomeIndexRule []interface{}
	for _, outcomeIndexItem := range outcomeIndex {
		outcomeIndexRule = append(outcomeIndexRule, outcomeIndexItem)
	}

	logs, sub, err := _FixedProductMarketMaker.contract.WatchLogs(opts, "FPMMSell", sellerRule, outcomeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FixedProductMarketMakerFPMMSell)
				if err := _FixedProductMarketMaker.contract.UnpackLog(event, "FPMMSell", log); err != nil {
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

// ParseFPMMSell is a log parse operation binding the contract event 0xadcf2a240ed9300d681d9a3f5382b6c1beed1b7e46643e0c7b42cbe6e2d766b4.
//
// Solidity: event FPMMSell(address indexed seller, uint256 returnAmount, uint256 feeAmount, uint256 indexed outcomeIndex, uint256 outcomeTokensSold)
func (_FixedProductMarketMaker *FixedProductMarketMakerFilterer) ParseFPMMSell(log types.Log) (*FixedProductMarketMakerFPMMSell, error) {
	event := new(FixedProductMarketMakerFPMMSell)
	if err := _FixedProductMarketMaker.contract.UnpackLog(event, "FPMMSell", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FixedProductMarketMakerTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the FixedProductMarketMaker contract.
type FixedProductMarketMakerTransferIterator struct {
	Event *FixedProductMarketMakerTransfer // Event containing the contract specifics and raw log

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
func (it *FixedProductMarketMakerTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FixedProductMarketMakerTransfer)
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
		it.Event = new(FixedProductMarketMakerTransfer)
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
func (it *FixedProductMarketMakerTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FixedProductMarketMakerTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FixedProductMarketMakerTransfer represents a Transfer event raised by the FixedProductMarketMaker contract.
type FixedProductMarketMakerTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FixedProductMarketMaker *FixedProductMarketMakerFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FixedProductMarketMakerTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FixedProductMarketMaker.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FixedProductMarketMakerTransferIterator{contract: _FixedProductMarketMaker.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FixedProductMarketMaker *FixedProductMarketMakerFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FixedProductMarketMakerTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FixedProductMarketMaker.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FixedProductMarketMakerTransfer)
				if err := _FixedProductMarketMaker.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FixedProductMarketMaker *FixedProductMarketMakerFilterer) ParseTransfer(log types.Log) (*FixedProductMarketMakerTransfer, error) {
	event := new(FixedProductMarketMakerTransfer)
	if err := _FixedProductMarketMaker.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
