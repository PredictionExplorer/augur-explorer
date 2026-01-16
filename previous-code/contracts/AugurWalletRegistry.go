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

// AugurWalletRegistryMetaData contains all meta data concerning the AugurWalletRegistry contract.
var AugurWalletRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"fundingSuccess\",\"type\":\"bool\"}],\"name\":\"ExecuteTransactionStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldRelayHub\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRelayHub\",\"type\":\"address\"}],\"name\":\"RelayHubChanged\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_encodedFunction\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_maxPossibleCharge\",\"type\":\"uint256\"}],\"name\":\"acceptRelayedCall\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_reason\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_context\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"affiliates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"augur\",\"outputs\":[{\"internalType\":\"contractIAugur\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"augurTrading\",\"outputs\":[{\"internalType\":\"contractIAugurTrading\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cash\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"createOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ethExchange\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Exchange\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_payment\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_referralAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_fingerprint\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_desiredSignerBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxExchangeRateInDai\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_revertOnFailure\",\"type\":\"bool\"}],\"name\":\"executeWalletTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fillOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getCreate2WalletAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getHubAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"relay\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"encodedFunction\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"transactionFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"getRelayMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getWallet\",\"outputs\":[{\"internalType\":\"contractIAugurWallet\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIAugur\",\"name\":\"_augur\",\"type\":\"address\"},{\"internalType\":\"contractIAugurTrading\",\"name\":\"_augurTrading\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"actualCharge\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"preRetVal\",\"type\":\"bytes32\"}],\"name\":\"postRelayedCall\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"}],\"name\":\"preRelayedCall\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"relayHubVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"shareToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token0IsCash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wallets\",\"outputs\":[{\"internalType\":\"contractIAugurWallet\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"zeroXTrade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AugurWalletRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use AugurWalletRegistryMetaData.ABI instead.
var AugurWalletRegistryABI = AugurWalletRegistryMetaData.ABI

// AugurWalletRegistry is an auto generated Go binding around an Ethereum contract.
type AugurWalletRegistry struct {
	AugurWalletRegistryCaller     // Read-only binding to the contract
	AugurWalletRegistryTransactor // Write-only binding to the contract
	AugurWalletRegistryFilterer   // Log filterer for contract events
}

// AugurWalletRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AugurWalletRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AugurWalletRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AugurWalletRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AugurWalletRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AugurWalletRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AugurWalletRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AugurWalletRegistrySession struct {
	Contract     *AugurWalletRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AugurWalletRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AugurWalletRegistryCallerSession struct {
	Contract *AugurWalletRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// AugurWalletRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AugurWalletRegistryTransactorSession struct {
	Contract     *AugurWalletRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// AugurWalletRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AugurWalletRegistryRaw struct {
	Contract *AugurWalletRegistry // Generic contract binding to access the raw methods on
}

// AugurWalletRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AugurWalletRegistryCallerRaw struct {
	Contract *AugurWalletRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// AugurWalletRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AugurWalletRegistryTransactorRaw struct {
	Contract *AugurWalletRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAugurWalletRegistry creates a new instance of AugurWalletRegistry, bound to a specific deployed contract.
func NewAugurWalletRegistry(address common.Address, backend bind.ContractBackend) (*AugurWalletRegistry, error) {
	contract, err := bindAugurWalletRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AugurWalletRegistry{AugurWalletRegistryCaller: AugurWalletRegistryCaller{contract: contract}, AugurWalletRegistryTransactor: AugurWalletRegistryTransactor{contract: contract}, AugurWalletRegistryFilterer: AugurWalletRegistryFilterer{contract: contract}}, nil
}

// NewAugurWalletRegistryCaller creates a new read-only instance of AugurWalletRegistry, bound to a specific deployed contract.
func NewAugurWalletRegistryCaller(address common.Address, caller bind.ContractCaller) (*AugurWalletRegistryCaller, error) {
	contract, err := bindAugurWalletRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AugurWalletRegistryCaller{contract: contract}, nil
}

// NewAugurWalletRegistryTransactor creates a new write-only instance of AugurWalletRegistry, bound to a specific deployed contract.
func NewAugurWalletRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*AugurWalletRegistryTransactor, error) {
	contract, err := bindAugurWalletRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AugurWalletRegistryTransactor{contract: contract}, nil
}

// NewAugurWalletRegistryFilterer creates a new log filterer instance of AugurWalletRegistry, bound to a specific deployed contract.
func NewAugurWalletRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*AugurWalletRegistryFilterer, error) {
	contract, err := bindAugurWalletRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AugurWalletRegistryFilterer{contract: contract}, nil
}

// bindAugurWalletRegistry binds a generic wrapper to an already deployed contract.
func bindAugurWalletRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AugurWalletRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AugurWalletRegistry *AugurWalletRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AugurWalletRegistry.Contract.AugurWalletRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AugurWalletRegistry *AugurWalletRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AugurWalletRegistry.Contract.AugurWalletRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AugurWalletRegistry *AugurWalletRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AugurWalletRegistry.Contract.AugurWalletRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AugurWalletRegistry *AugurWalletRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AugurWalletRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AugurWalletRegistry *AugurWalletRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AugurWalletRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AugurWalletRegistry *AugurWalletRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AugurWalletRegistry.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistryCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AugurWalletRegistry.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistrySession) WETH() (common.Address, error) {
	return _AugurWalletRegistry.Contract.WETH(&_AugurWalletRegistry.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistryCallerSession) WETH() (common.Address, error) {
	return _AugurWalletRegistry.Contract.WETH(&_AugurWalletRegistry.CallOpts)
}

// AcceptRelayedCall is a free data retrieval call binding the contract method 0x83947ea0.
//
// Solidity: function acceptRelayedCall(address , address _from, bytes _encodedFunction, uint256 , uint256 , uint256 , uint256 , bytes , uint256 _maxPossibleCharge) view returns(uint256 _reason, bytes _context)
func (_AugurWalletRegistry *AugurWalletRegistryCaller) AcceptRelayedCall(opts *bind.CallOpts, arg0 common.Address, _from common.Address, _encodedFunction []byte, arg3 *big.Int, arg4 *big.Int, arg5 *big.Int, arg6 *big.Int, arg7 []byte, _maxPossibleCharge *big.Int) (struct {
	Reason  *big.Int
	Context []byte
}, error) {
	var out []interface{}
	err := _AugurWalletRegistry.contract.Call(opts, &out, "acceptRelayedCall", arg0, _from, _encodedFunction, arg3, arg4, arg5, arg6, arg7, _maxPossibleCharge)

	outstruct := new(struct {
		Reason  *big.Int
		Context []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reason = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Context = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// AcceptRelayedCall is a free data retrieval call binding the contract method 0x83947ea0.
//
// Solidity: function acceptRelayedCall(address , address _from, bytes _encodedFunction, uint256 , uint256 , uint256 , uint256 , bytes , uint256 _maxPossibleCharge) view returns(uint256 _reason, bytes _context)
func (_AugurWalletRegistry *AugurWalletRegistrySession) AcceptRelayedCall(arg0 common.Address, _from common.Address, _encodedFunction []byte, arg3 *big.Int, arg4 *big.Int, arg5 *big.Int, arg6 *big.Int, arg7 []byte, _maxPossibleCharge *big.Int) (struct {
	Reason  *big.Int
	Context []byte
}, error) {
	return _AugurWalletRegistry.Contract.AcceptRelayedCall(&_AugurWalletRegistry.CallOpts, arg0, _from, _encodedFunction, arg3, arg4, arg5, arg6, arg7, _maxPossibleCharge)
}

// AcceptRelayedCall is a free data retrieval call binding the contract method 0x83947ea0.
//
// Solidity: function acceptRelayedCall(address , address _from, bytes _encodedFunction, uint256 , uint256 , uint256 , uint256 , bytes , uint256 _maxPossibleCharge) view returns(uint256 _reason, bytes _context)
func (_AugurWalletRegistry *AugurWalletRegistryCallerSession) AcceptRelayedCall(arg0 common.Address, _from common.Address, _encodedFunction []byte, arg3 *big.Int, arg4 *big.Int, arg5 *big.Int, arg6 *big.Int, arg7 []byte, _maxPossibleCharge *big.Int) (struct {
	Reason  *big.Int
	Context []byte
}, error) {
	return _AugurWalletRegistry.Contract.AcceptRelayedCall(&_AugurWalletRegistry.CallOpts, arg0, _from, _encodedFunction, arg3, arg4, arg5, arg6, arg7, _maxPossibleCharge)
}

// Affiliates is a free data retrieval call binding the contract method 0x3f03842a.
//
// Solidity: function affiliates() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistryCaller) Affiliates(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AugurWalletRegistry.contract.Call(opts, &out, "affiliates")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Affiliates is a free data retrieval call binding the contract method 0x3f03842a.
//
// Solidity: function affiliates() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistrySession) Affiliates() (common.Address, error) {
	return _AugurWalletRegistry.Contract.Affiliates(&_AugurWalletRegistry.CallOpts)
}

// Affiliates is a free data retrieval call binding the contract method 0x3f03842a.
//
// Solidity: function affiliates() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistryCallerSession) Affiliates() (common.Address, error) {
	return _AugurWalletRegistry.Contract.Affiliates(&_AugurWalletRegistry.CallOpts)
}

// Augur is a free data retrieval call binding the contract method 0x7a0d8f8a.
//
// Solidity: function augur() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistryCaller) Augur(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AugurWalletRegistry.contract.Call(opts, &out, "augur")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Augur is a free data retrieval call binding the contract method 0x7a0d8f8a.
//
// Solidity: function augur() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistrySession) Augur() (common.Address, error) {
	return _AugurWalletRegistry.Contract.Augur(&_AugurWalletRegistry.CallOpts)
}

// Augur is a free data retrieval call binding the contract method 0x7a0d8f8a.
//
// Solidity: function augur() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistryCallerSession) Augur() (common.Address, error) {
	return _AugurWalletRegistry.Contract.Augur(&_AugurWalletRegistry.CallOpts)
}

// AugurTrading is a free data retrieval call binding the contract method 0x8e1bfa73.
//
// Solidity: function augurTrading() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistryCaller) AugurTrading(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AugurWalletRegistry.contract.Call(opts, &out, "augurTrading")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AugurTrading is a free data retrieval call binding the contract method 0x8e1bfa73.
//
// Solidity: function augurTrading() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistrySession) AugurTrading() (common.Address, error) {
	return _AugurWalletRegistry.Contract.AugurTrading(&_AugurWalletRegistry.CallOpts)
}

// AugurTrading is a free data retrieval call binding the contract method 0x8e1bfa73.
//
// Solidity: function augurTrading() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistryCallerSession) AugurTrading() (common.Address, error) {
	return _AugurWalletRegistry.Contract.AugurTrading(&_AugurWalletRegistry.CallOpts)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistryCaller) Cash(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AugurWalletRegistry.contract.Call(opts, &out, "cash")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistrySession) Cash() (common.Address, error) {
	return _AugurWalletRegistry.Contract.Cash(&_AugurWalletRegistry.CallOpts)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistryCallerSession) Cash() (common.Address, error) {
	return _AugurWalletRegistry.Contract.Cash(&_AugurWalletRegistry.CallOpts)
}

// CreateOrder is a free data retrieval call binding the contract method 0x6512e6ec.
//
// Solidity: function createOrder() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistryCaller) CreateOrder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AugurWalletRegistry.contract.Call(opts, &out, "createOrder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreateOrder is a free data retrieval call binding the contract method 0x6512e6ec.
//
// Solidity: function createOrder() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistrySession) CreateOrder() (common.Address, error) {
	return _AugurWalletRegistry.Contract.CreateOrder(&_AugurWalletRegistry.CallOpts)
}

// CreateOrder is a free data retrieval call binding the contract method 0x6512e6ec.
//
// Solidity: function createOrder() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistryCallerSession) CreateOrder() (common.Address, error) {
	return _AugurWalletRegistry.Contract.CreateOrder(&_AugurWalletRegistry.CallOpts)
}

// EthExchange is a free data retrieval call binding the contract method 0x53e569a1.
//
// Solidity: function ethExchange() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistryCaller) EthExchange(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AugurWalletRegistry.contract.Call(opts, &out, "ethExchange")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthExchange is a free data retrieval call binding the contract method 0x53e569a1.
//
// Solidity: function ethExchange() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistrySession) EthExchange() (common.Address, error) {
	return _AugurWalletRegistry.Contract.EthExchange(&_AugurWalletRegistry.CallOpts)
}

// EthExchange is a free data retrieval call binding the contract method 0x53e569a1.
//
// Solidity: function ethExchange() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistryCallerSession) EthExchange() (common.Address, error) {
	return _AugurWalletRegistry.Contract.EthExchange(&_AugurWalletRegistry.CallOpts)
}

// FillOrder is a free data retrieval call binding the contract method 0x5c1ad844.
//
// Solidity: function fillOrder() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistryCaller) FillOrder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AugurWalletRegistry.contract.Call(opts, &out, "fillOrder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FillOrder is a free data retrieval call binding the contract method 0x5c1ad844.
//
// Solidity: function fillOrder() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistrySession) FillOrder() (common.Address, error) {
	return _AugurWalletRegistry.Contract.FillOrder(&_AugurWalletRegistry.CallOpts)
}

// FillOrder is a free data retrieval call binding the contract method 0x5c1ad844.
//
// Solidity: function fillOrder() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistryCallerSession) FillOrder() (common.Address, error) {
	return _AugurWalletRegistry.Contract.FillOrder(&_AugurWalletRegistry.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_AugurWalletRegistry *AugurWalletRegistryCaller) GetAmountIn(opts *bind.CallOpts, amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AugurWalletRegistry.contract.Call(opts, &out, "getAmountIn", amountOut, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_AugurWalletRegistry *AugurWalletRegistrySession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _AugurWalletRegistry.Contract.GetAmountIn(&_AugurWalletRegistry.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_AugurWalletRegistry *AugurWalletRegistryCallerSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _AugurWalletRegistry.Contract.GetAmountIn(&_AugurWalletRegistry.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_AugurWalletRegistry *AugurWalletRegistryCaller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AugurWalletRegistry.contract.Call(opts, &out, "getAmountOut", amountIn, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_AugurWalletRegistry *AugurWalletRegistrySession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _AugurWalletRegistry.Contract.GetAmountOut(&_AugurWalletRegistry.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_AugurWalletRegistry *AugurWalletRegistryCallerSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _AugurWalletRegistry.Contract.GetAmountOut(&_AugurWalletRegistry.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetCreate2WalletAddress is a free data retrieval call binding the contract method 0x80a35930.
//
// Solidity: function getCreate2WalletAddress(address _owner) view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistryCaller) GetCreate2WalletAddress(opts *bind.CallOpts, _owner common.Address) (common.Address, error) {
	var out []interface{}
	err := _AugurWalletRegistry.contract.Call(opts, &out, "getCreate2WalletAddress", _owner)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCreate2WalletAddress is a free data retrieval call binding the contract method 0x80a35930.
//
// Solidity: function getCreate2WalletAddress(address _owner) view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistrySession) GetCreate2WalletAddress(_owner common.Address) (common.Address, error) {
	return _AugurWalletRegistry.Contract.GetCreate2WalletAddress(&_AugurWalletRegistry.CallOpts, _owner)
}

// GetCreate2WalletAddress is a free data retrieval call binding the contract method 0x80a35930.
//
// Solidity: function getCreate2WalletAddress(address _owner) view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistryCallerSession) GetCreate2WalletAddress(_owner common.Address) (common.Address, error) {
	return _AugurWalletRegistry.Contract.GetCreate2WalletAddress(&_AugurWalletRegistry.CallOpts, _owner)
}

// GetHubAddr is a free data retrieval call binding the contract method 0x74e861d6.
//
// Solidity: function getHubAddr() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistryCaller) GetHubAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AugurWalletRegistry.contract.Call(opts, &out, "getHubAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetHubAddr is a free data retrieval call binding the contract method 0x74e861d6.
//
// Solidity: function getHubAddr() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistrySession) GetHubAddr() (common.Address, error) {
	return _AugurWalletRegistry.Contract.GetHubAddr(&_AugurWalletRegistry.CallOpts)
}

// GetHubAddr is a free data retrieval call binding the contract method 0x74e861d6.
//
// Solidity: function getHubAddr() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistryCallerSession) GetHubAddr() (common.Address, error) {
	return _AugurWalletRegistry.Contract.GetHubAddr(&_AugurWalletRegistry.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() view returns(bool)
func (_AugurWalletRegistry *AugurWalletRegistryCaller) GetInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AugurWalletRegistry.contract.Call(opts, &out, "getInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() view returns(bool)
func (_AugurWalletRegistry *AugurWalletRegistrySession) GetInitialized() (bool, error) {
	return _AugurWalletRegistry.Contract.GetInitialized(&_AugurWalletRegistry.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() view returns(bool)
func (_AugurWalletRegistry *AugurWalletRegistryCallerSession) GetInitialized() (bool, error) {
	return _AugurWalletRegistry.Contract.GetInitialized(&_AugurWalletRegistry.CallOpts)
}

// GetRelayMessageHash is a free data retrieval call binding the contract method 0xd5b9562b.
//
// Solidity: function getRelayMessageHash(address relay, address from, address to, bytes encodedFunction, uint256 transactionFee, uint256 gasPrice, uint256 gasLimit, uint256 nonce) view returns(bytes32)
func (_AugurWalletRegistry *AugurWalletRegistryCaller) GetRelayMessageHash(opts *bind.CallOpts, relay common.Address, from common.Address, to common.Address, encodedFunction []byte, transactionFee *big.Int, gasPrice *big.Int, gasLimit *big.Int, nonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _AugurWalletRegistry.contract.Call(opts, &out, "getRelayMessageHash", relay, from, to, encodedFunction, transactionFee, gasPrice, gasLimit, nonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRelayMessageHash is a free data retrieval call binding the contract method 0xd5b9562b.
//
// Solidity: function getRelayMessageHash(address relay, address from, address to, bytes encodedFunction, uint256 transactionFee, uint256 gasPrice, uint256 gasLimit, uint256 nonce) view returns(bytes32)
func (_AugurWalletRegistry *AugurWalletRegistrySession) GetRelayMessageHash(relay common.Address, from common.Address, to common.Address, encodedFunction []byte, transactionFee *big.Int, gasPrice *big.Int, gasLimit *big.Int, nonce *big.Int) ([32]byte, error) {
	return _AugurWalletRegistry.Contract.GetRelayMessageHash(&_AugurWalletRegistry.CallOpts, relay, from, to, encodedFunction, transactionFee, gasPrice, gasLimit, nonce)
}

// GetRelayMessageHash is a free data retrieval call binding the contract method 0xd5b9562b.
//
// Solidity: function getRelayMessageHash(address relay, address from, address to, bytes encodedFunction, uint256 transactionFee, uint256 gasPrice, uint256 gasLimit, uint256 nonce) view returns(bytes32)
func (_AugurWalletRegistry *AugurWalletRegistryCallerSession) GetRelayMessageHash(relay common.Address, from common.Address, to common.Address, encodedFunction []byte, transactionFee *big.Int, gasPrice *big.Int, gasLimit *big.Int, nonce *big.Int) ([32]byte, error) {
	return _AugurWalletRegistry.Contract.GetRelayMessageHash(&_AugurWalletRegistry.CallOpts, relay, from, to, encodedFunction, transactionFee, gasPrice, gasLimit, nonce)
}

// GetWallet is a free data retrieval call binding the contract method 0x04d0a647.
//
// Solidity: function getWallet(address _account) view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistryCaller) GetWallet(opts *bind.CallOpts, _account common.Address) (common.Address, error) {
	var out []interface{}
	err := _AugurWalletRegistry.contract.Call(opts, &out, "getWallet", _account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWallet is a free data retrieval call binding the contract method 0x04d0a647.
//
// Solidity: function getWallet(address _account) view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistrySession) GetWallet(_account common.Address) (common.Address, error) {
	return _AugurWalletRegistry.Contract.GetWallet(&_AugurWalletRegistry.CallOpts, _account)
}

// GetWallet is a free data retrieval call binding the contract method 0x04d0a647.
//
// Solidity: function getWallet(address _account) view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistryCallerSession) GetWallet(_account common.Address) (common.Address, error) {
	return _AugurWalletRegistry.Contract.GetWallet(&_AugurWalletRegistry.CallOpts, _account)
}

// RelayHubVersion is a free data retrieval call binding the contract method 0xad61ccd5.
//
// Solidity: function relayHubVersion() view returns(string)
func (_AugurWalletRegistry *AugurWalletRegistryCaller) RelayHubVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AugurWalletRegistry.contract.Call(opts, &out, "relayHubVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RelayHubVersion is a free data retrieval call binding the contract method 0xad61ccd5.
//
// Solidity: function relayHubVersion() view returns(string)
func (_AugurWalletRegistry *AugurWalletRegistrySession) RelayHubVersion() (string, error) {
	return _AugurWalletRegistry.Contract.RelayHubVersion(&_AugurWalletRegistry.CallOpts)
}

// RelayHubVersion is a free data retrieval call binding the contract method 0xad61ccd5.
//
// Solidity: function relayHubVersion() view returns(string)
func (_AugurWalletRegistry *AugurWalletRegistryCallerSession) RelayHubVersion() (string, error) {
	return _AugurWalletRegistry.Contract.RelayHubVersion(&_AugurWalletRegistry.CallOpts)
}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistryCaller) ShareToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AugurWalletRegistry.contract.Call(opts, &out, "shareToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistrySession) ShareToken() (common.Address, error) {
	return _AugurWalletRegistry.Contract.ShareToken(&_AugurWalletRegistry.CallOpts)
}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistryCallerSession) ShareToken() (common.Address, error) {
	return _AugurWalletRegistry.Contract.ShareToken(&_AugurWalletRegistry.CallOpts)
}

// Token0IsCash is a free data retrieval call binding the contract method 0x7b66c82a.
//
// Solidity: function token0IsCash() view returns(bool)
func (_AugurWalletRegistry *AugurWalletRegistryCaller) Token0IsCash(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AugurWalletRegistry.contract.Call(opts, &out, "token0IsCash")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Token0IsCash is a free data retrieval call binding the contract method 0x7b66c82a.
//
// Solidity: function token0IsCash() view returns(bool)
func (_AugurWalletRegistry *AugurWalletRegistrySession) Token0IsCash() (bool, error) {
	return _AugurWalletRegistry.Contract.Token0IsCash(&_AugurWalletRegistry.CallOpts)
}

// Token0IsCash is a free data retrieval call binding the contract method 0x7b66c82a.
//
// Solidity: function token0IsCash() view returns(bool)
func (_AugurWalletRegistry *AugurWalletRegistryCallerSession) Token0IsCash() (bool, error) {
	return _AugurWalletRegistry.Contract.Token0IsCash(&_AugurWalletRegistry.CallOpts)
}

// Wallets is a free data retrieval call binding the contract method 0x89b08f11.
//
// Solidity: function wallets(address ) view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistryCaller) Wallets(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _AugurWalletRegistry.contract.Call(opts, &out, "wallets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wallets is a free data retrieval call binding the contract method 0x89b08f11.
//
// Solidity: function wallets(address ) view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistrySession) Wallets(arg0 common.Address) (common.Address, error) {
	return _AugurWalletRegistry.Contract.Wallets(&_AugurWalletRegistry.CallOpts, arg0)
}

// Wallets is a free data retrieval call binding the contract method 0x89b08f11.
//
// Solidity: function wallets(address ) view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistryCallerSession) Wallets(arg0 common.Address) (common.Address, error) {
	return _AugurWalletRegistry.Contract.Wallets(&_AugurWalletRegistry.CallOpts, arg0)
}

// ZeroXTrade is a free data retrieval call binding the contract method 0x1a505b76.
//
// Solidity: function zeroXTrade() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistryCaller) ZeroXTrade(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AugurWalletRegistry.contract.Call(opts, &out, "zeroXTrade")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZeroXTrade is a free data retrieval call binding the contract method 0x1a505b76.
//
// Solidity: function zeroXTrade() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistrySession) ZeroXTrade() (common.Address, error) {
	return _AugurWalletRegistry.Contract.ZeroXTrade(&_AugurWalletRegistry.CallOpts)
}

// ZeroXTrade is a free data retrieval call binding the contract method 0x1a505b76.
//
// Solidity: function zeroXTrade() view returns(address)
func (_AugurWalletRegistry *AugurWalletRegistryCallerSession) ZeroXTrade() (common.Address, error) {
	return _AugurWalletRegistry.Contract.ZeroXTrade(&_AugurWalletRegistry.CallOpts)
}

// ExecuteWalletTransaction is a paid mutator transaction binding the contract method 0x78dc0eed.
//
// Solidity: function executeWalletTransaction(address _to, bytes _data, uint256 _value, uint256 _payment, address _referralAddress, bytes32 _fingerprint, uint256 _desiredSignerBalance, uint256 _maxExchangeRateInDai, bool _revertOnFailure) returns()
func (_AugurWalletRegistry *AugurWalletRegistryTransactor) ExecuteWalletTransaction(opts *bind.TransactOpts, _to common.Address, _data []byte, _value *big.Int, _payment *big.Int, _referralAddress common.Address, _fingerprint [32]byte, _desiredSignerBalance *big.Int, _maxExchangeRateInDai *big.Int, _revertOnFailure bool) (*types.Transaction, error) {
	return _AugurWalletRegistry.contract.Transact(opts, "executeWalletTransaction", _to, _data, _value, _payment, _referralAddress, _fingerprint, _desiredSignerBalance, _maxExchangeRateInDai, _revertOnFailure)
}

// ExecuteWalletTransaction is a paid mutator transaction binding the contract method 0x78dc0eed.
//
// Solidity: function executeWalletTransaction(address _to, bytes _data, uint256 _value, uint256 _payment, address _referralAddress, bytes32 _fingerprint, uint256 _desiredSignerBalance, uint256 _maxExchangeRateInDai, bool _revertOnFailure) returns()
func (_AugurWalletRegistry *AugurWalletRegistrySession) ExecuteWalletTransaction(_to common.Address, _data []byte, _value *big.Int, _payment *big.Int, _referralAddress common.Address, _fingerprint [32]byte, _desiredSignerBalance *big.Int, _maxExchangeRateInDai *big.Int, _revertOnFailure bool) (*types.Transaction, error) {
	return _AugurWalletRegistry.Contract.ExecuteWalletTransaction(&_AugurWalletRegistry.TransactOpts, _to, _data, _value, _payment, _referralAddress, _fingerprint, _desiredSignerBalance, _maxExchangeRateInDai, _revertOnFailure)
}

// ExecuteWalletTransaction is a paid mutator transaction binding the contract method 0x78dc0eed.
//
// Solidity: function executeWalletTransaction(address _to, bytes _data, uint256 _value, uint256 _payment, address _referralAddress, bytes32 _fingerprint, uint256 _desiredSignerBalance, uint256 _maxExchangeRateInDai, bool _revertOnFailure) returns()
func (_AugurWalletRegistry *AugurWalletRegistryTransactorSession) ExecuteWalletTransaction(_to common.Address, _data []byte, _value *big.Int, _payment *big.Int, _referralAddress common.Address, _fingerprint [32]byte, _desiredSignerBalance *big.Int, _maxExchangeRateInDai *big.Int, _revertOnFailure bool) (*types.Transaction, error) {
	return _AugurWalletRegistry.Contract.ExecuteWalletTransaction(&_AugurWalletRegistry.TransactOpts, _to, _data, _value, _payment, _referralAddress, _fingerprint, _desiredSignerBalance, _maxExchangeRateInDai, _revertOnFailure)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _augur, address _augurTrading) payable returns(bool)
func (_AugurWalletRegistry *AugurWalletRegistryTransactor) Initialize(opts *bind.TransactOpts, _augur common.Address, _augurTrading common.Address) (*types.Transaction, error) {
	return _AugurWalletRegistry.contract.Transact(opts, "initialize", _augur, _augurTrading)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _augur, address _augurTrading) payable returns(bool)
func (_AugurWalletRegistry *AugurWalletRegistrySession) Initialize(_augur common.Address, _augurTrading common.Address) (*types.Transaction, error) {
	return _AugurWalletRegistry.Contract.Initialize(&_AugurWalletRegistry.TransactOpts, _augur, _augurTrading)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _augur, address _augurTrading) payable returns(bool)
func (_AugurWalletRegistry *AugurWalletRegistryTransactorSession) Initialize(_augur common.Address, _augurTrading common.Address) (*types.Transaction, error) {
	return _AugurWalletRegistry.Contract.Initialize(&_AugurWalletRegistry.TransactOpts, _augur, _augurTrading)
}

// PostRelayedCall is a paid mutator transaction binding the contract method 0xe06e0e22.
//
// Solidity: function postRelayedCall(bytes context, bool success, uint256 actualCharge, bytes32 preRetVal) returns()
func (_AugurWalletRegistry *AugurWalletRegistryTransactor) PostRelayedCall(opts *bind.TransactOpts, context []byte, success bool, actualCharge *big.Int, preRetVal [32]byte) (*types.Transaction, error) {
	return _AugurWalletRegistry.contract.Transact(opts, "postRelayedCall", context, success, actualCharge, preRetVal)
}

// PostRelayedCall is a paid mutator transaction binding the contract method 0xe06e0e22.
//
// Solidity: function postRelayedCall(bytes context, bool success, uint256 actualCharge, bytes32 preRetVal) returns()
func (_AugurWalletRegistry *AugurWalletRegistrySession) PostRelayedCall(context []byte, success bool, actualCharge *big.Int, preRetVal [32]byte) (*types.Transaction, error) {
	return _AugurWalletRegistry.Contract.PostRelayedCall(&_AugurWalletRegistry.TransactOpts, context, success, actualCharge, preRetVal)
}

// PostRelayedCall is a paid mutator transaction binding the contract method 0xe06e0e22.
//
// Solidity: function postRelayedCall(bytes context, bool success, uint256 actualCharge, bytes32 preRetVal) returns()
func (_AugurWalletRegistry *AugurWalletRegistryTransactorSession) PostRelayedCall(context []byte, success bool, actualCharge *big.Int, preRetVal [32]byte) (*types.Transaction, error) {
	return _AugurWalletRegistry.Contract.PostRelayedCall(&_AugurWalletRegistry.TransactOpts, context, success, actualCharge, preRetVal)
}

// PreRelayedCall is a paid mutator transaction binding the contract method 0x80274db7.
//
// Solidity: function preRelayedCall(bytes context) returns(bytes32)
func (_AugurWalletRegistry *AugurWalletRegistryTransactor) PreRelayedCall(opts *bind.TransactOpts, context []byte) (*types.Transaction, error) {
	return _AugurWalletRegistry.contract.Transact(opts, "preRelayedCall", context)
}

// PreRelayedCall is a paid mutator transaction binding the contract method 0x80274db7.
//
// Solidity: function preRelayedCall(bytes context) returns(bytes32)
func (_AugurWalletRegistry *AugurWalletRegistrySession) PreRelayedCall(context []byte) (*types.Transaction, error) {
	return _AugurWalletRegistry.Contract.PreRelayedCall(&_AugurWalletRegistry.TransactOpts, context)
}

// PreRelayedCall is a paid mutator transaction binding the contract method 0x80274db7.
//
// Solidity: function preRelayedCall(bytes context) returns(bytes32)
func (_AugurWalletRegistry *AugurWalletRegistryTransactorSession) PreRelayedCall(context []byte) (*types.Transaction, error) {
	return _AugurWalletRegistry.Contract.PreRelayedCall(&_AugurWalletRegistry.TransactOpts, context)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_AugurWalletRegistry *AugurWalletRegistryTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _AugurWalletRegistry.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_AugurWalletRegistry *AugurWalletRegistrySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _AugurWalletRegistry.Contract.Fallback(&_AugurWalletRegistry.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_AugurWalletRegistry *AugurWalletRegistryTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _AugurWalletRegistry.Contract.Fallback(&_AugurWalletRegistry.TransactOpts, calldata)
}

// AugurWalletRegistryExecuteTransactionStatusIterator is returned from FilterExecuteTransactionStatus and is used to iterate over the raw logs and unpacked data for ExecuteTransactionStatus events raised by the AugurWalletRegistry contract.
type AugurWalletRegistryExecuteTransactionStatusIterator struct {
	Event *AugurWalletRegistryExecuteTransactionStatus // Event containing the contract specifics and raw log

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
func (it *AugurWalletRegistryExecuteTransactionStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurWalletRegistryExecuteTransactionStatus)
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
		it.Event = new(AugurWalletRegistryExecuteTransactionStatus)
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
func (it *AugurWalletRegistryExecuteTransactionStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurWalletRegistryExecuteTransactionStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurWalletRegistryExecuteTransactionStatus represents a ExecuteTransactionStatus event raised by the AugurWalletRegistry contract.
type AugurWalletRegistryExecuteTransactionStatus struct {
	Success        bool
	FundingSuccess bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterExecuteTransactionStatus is a free log retrieval operation binding the contract event 0xee9c28a7fe7177d351e891cb4ca5b7a4e4aba4974be67fb7665ba1ad0e703439.
//
// Solidity: event ExecuteTransactionStatus(bool success, bool fundingSuccess)
func (_AugurWalletRegistry *AugurWalletRegistryFilterer) FilterExecuteTransactionStatus(opts *bind.FilterOpts) (*AugurWalletRegistryExecuteTransactionStatusIterator, error) {

	logs, sub, err := _AugurWalletRegistry.contract.FilterLogs(opts, "ExecuteTransactionStatus")
	if err != nil {
		return nil, err
	}
	return &AugurWalletRegistryExecuteTransactionStatusIterator{contract: _AugurWalletRegistry.contract, event: "ExecuteTransactionStatus", logs: logs, sub: sub}, nil
}

// WatchExecuteTransactionStatus is a free log subscription operation binding the contract event 0xee9c28a7fe7177d351e891cb4ca5b7a4e4aba4974be67fb7665ba1ad0e703439.
//
// Solidity: event ExecuteTransactionStatus(bool success, bool fundingSuccess)
func (_AugurWalletRegistry *AugurWalletRegistryFilterer) WatchExecuteTransactionStatus(opts *bind.WatchOpts, sink chan<- *AugurWalletRegistryExecuteTransactionStatus) (event.Subscription, error) {

	logs, sub, err := _AugurWalletRegistry.contract.WatchLogs(opts, "ExecuteTransactionStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurWalletRegistryExecuteTransactionStatus)
				if err := _AugurWalletRegistry.contract.UnpackLog(event, "ExecuteTransactionStatus", log); err != nil {
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

// ParseExecuteTransactionStatus is a log parse operation binding the contract event 0xee9c28a7fe7177d351e891cb4ca5b7a4e4aba4974be67fb7665ba1ad0e703439.
//
// Solidity: event ExecuteTransactionStatus(bool success, bool fundingSuccess)
func (_AugurWalletRegistry *AugurWalletRegistryFilterer) ParseExecuteTransactionStatus(log types.Log) (*AugurWalletRegistryExecuteTransactionStatus, error) {
	event := new(AugurWalletRegistryExecuteTransactionStatus)
	if err := _AugurWalletRegistry.contract.UnpackLog(event, "ExecuteTransactionStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurWalletRegistryRelayHubChangedIterator is returned from FilterRelayHubChanged and is used to iterate over the raw logs and unpacked data for RelayHubChanged events raised by the AugurWalletRegistry contract.
type AugurWalletRegistryRelayHubChangedIterator struct {
	Event *AugurWalletRegistryRelayHubChanged // Event containing the contract specifics and raw log

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
func (it *AugurWalletRegistryRelayHubChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurWalletRegistryRelayHubChanged)
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
		it.Event = new(AugurWalletRegistryRelayHubChanged)
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
func (it *AugurWalletRegistryRelayHubChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurWalletRegistryRelayHubChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurWalletRegistryRelayHubChanged represents a RelayHubChanged event raised by the AugurWalletRegistry contract.
type AugurWalletRegistryRelayHubChanged struct {
	OldRelayHub common.Address
	NewRelayHub common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRelayHubChanged is a free log retrieval operation binding the contract event 0xb9f84b8e65164b14439ae3620df0a4d8786d896996c0282b683f9d8c08f046e8.
//
// Solidity: event RelayHubChanged(address indexed oldRelayHub, address indexed newRelayHub)
func (_AugurWalletRegistry *AugurWalletRegistryFilterer) FilterRelayHubChanged(opts *bind.FilterOpts, oldRelayHub []common.Address, newRelayHub []common.Address) (*AugurWalletRegistryRelayHubChangedIterator, error) {

	var oldRelayHubRule []interface{}
	for _, oldRelayHubItem := range oldRelayHub {
		oldRelayHubRule = append(oldRelayHubRule, oldRelayHubItem)
	}
	var newRelayHubRule []interface{}
	for _, newRelayHubItem := range newRelayHub {
		newRelayHubRule = append(newRelayHubRule, newRelayHubItem)
	}

	logs, sub, err := _AugurWalletRegistry.contract.FilterLogs(opts, "RelayHubChanged", oldRelayHubRule, newRelayHubRule)
	if err != nil {
		return nil, err
	}
	return &AugurWalletRegistryRelayHubChangedIterator{contract: _AugurWalletRegistry.contract, event: "RelayHubChanged", logs: logs, sub: sub}, nil
}

// WatchRelayHubChanged is a free log subscription operation binding the contract event 0xb9f84b8e65164b14439ae3620df0a4d8786d896996c0282b683f9d8c08f046e8.
//
// Solidity: event RelayHubChanged(address indexed oldRelayHub, address indexed newRelayHub)
func (_AugurWalletRegistry *AugurWalletRegistryFilterer) WatchRelayHubChanged(opts *bind.WatchOpts, sink chan<- *AugurWalletRegistryRelayHubChanged, oldRelayHub []common.Address, newRelayHub []common.Address) (event.Subscription, error) {

	var oldRelayHubRule []interface{}
	for _, oldRelayHubItem := range oldRelayHub {
		oldRelayHubRule = append(oldRelayHubRule, oldRelayHubItem)
	}
	var newRelayHubRule []interface{}
	for _, newRelayHubItem := range newRelayHub {
		newRelayHubRule = append(newRelayHubRule, newRelayHubItem)
	}

	logs, sub, err := _AugurWalletRegistry.contract.WatchLogs(opts, "RelayHubChanged", oldRelayHubRule, newRelayHubRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurWalletRegistryRelayHubChanged)
				if err := _AugurWalletRegistry.contract.UnpackLog(event, "RelayHubChanged", log); err != nil {
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

// ParseRelayHubChanged is a log parse operation binding the contract event 0xb9f84b8e65164b14439ae3620df0a4d8786d896996c0282b683f9d8c08f046e8.
//
// Solidity: event RelayHubChanged(address indexed oldRelayHub, address indexed newRelayHub)
func (_AugurWalletRegistry *AugurWalletRegistryFilterer) ParseRelayHubChanged(log types.Log) (*AugurWalletRegistryRelayHubChanged, error) {
	event := new(AugurWalletRegistryRelayHubChanged)
	if err := _AugurWalletRegistry.contract.UnpackLog(event, "RelayHubChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
