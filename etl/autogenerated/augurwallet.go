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

// AugurWalletABI is the input ABI used to generate the binding from.
const AugurWalletABI = "[{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MSG_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_authorizedProxy\",\"type\":\"address\"}],\"name\":\"addAuthorizedProxy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorizedProxies\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cash\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_referralAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_fingerprint\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_augur\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_registryV2\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_cash\",\"type\":\"address\"},{\"internalType\":\"contractIAffiliates\",\"name\":\"_affiliates\",\"type\":\"address\"},{\"internalType\":\"contractIERC1155\",\"name\":\"_shareToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_createOrder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_fillOrder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zeroXTrade\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIAugurWalletRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_authorizedProxy\",\"type\":\"address\"}],\"name\":\"removeAuthorizedProxy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferCash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minExchangeRateInDai\",\"type\":\"uint256\"}],\"name\":\"withdrawAllFundsAsDai\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// AugurWallet is an auto generated Go binding around an Ethereum contract.
type AugurWallet struct {
	AugurWalletCaller     // Read-only binding to the contract
	AugurWalletTransactor // Write-only binding to the contract
	AugurWalletFilterer   // Log filterer for contract events
}

// AugurWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type AugurWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AugurWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AugurWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AugurWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AugurWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AugurWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AugurWalletSession struct {
	Contract     *AugurWallet      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AugurWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AugurWalletCallerSession struct {
	Contract *AugurWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AugurWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AugurWalletTransactorSession struct {
	Contract     *AugurWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AugurWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type AugurWalletRaw struct {
	Contract *AugurWallet // Generic contract binding to access the raw methods on
}

// AugurWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AugurWalletCallerRaw struct {
	Contract *AugurWalletCaller // Generic read-only contract binding to access the raw methods on
}

// AugurWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AugurWalletTransactorRaw struct {
	Contract *AugurWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAugurWallet creates a new instance of AugurWallet, bound to a specific deployed contract.
func NewAugurWallet(address common.Address, backend bind.ContractBackend) (*AugurWallet, error) {
	contract, err := bindAugurWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AugurWallet{AugurWalletCaller: AugurWalletCaller{contract: contract}, AugurWalletTransactor: AugurWalletTransactor{contract: contract}, AugurWalletFilterer: AugurWalletFilterer{contract: contract}}, nil
}

// NewAugurWalletCaller creates a new read-only instance of AugurWallet, bound to a specific deployed contract.
func NewAugurWalletCaller(address common.Address, caller bind.ContractCaller) (*AugurWalletCaller, error) {
	contract, err := bindAugurWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AugurWalletCaller{contract: contract}, nil
}

// NewAugurWalletTransactor creates a new write-only instance of AugurWallet, bound to a specific deployed contract.
func NewAugurWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*AugurWalletTransactor, error) {
	contract, err := bindAugurWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AugurWalletTransactor{contract: contract}, nil
}

// NewAugurWalletFilterer creates a new log filterer instance of AugurWallet, bound to a specific deployed contract.
func NewAugurWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*AugurWalletFilterer, error) {
	contract, err := bindAugurWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AugurWalletFilterer{contract: contract}, nil
}

// bindAugurWallet binds a generic wrapper to an already deployed contract.
func bindAugurWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AugurWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AugurWallet *AugurWalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AugurWallet.Contract.AugurWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AugurWallet *AugurWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AugurWallet.Contract.AugurWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AugurWallet *AugurWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AugurWallet.Contract.AugurWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AugurWallet *AugurWalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AugurWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AugurWallet *AugurWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AugurWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AugurWallet *AugurWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AugurWallet.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATORTYPEHASH is a free data retrieval call binding the contract method 0x1db61b54.
//
// Solidity: function DOMAIN_SEPARATOR_TYPEHASH() view returns(bytes32)
func (_AugurWallet *AugurWalletCaller) DOMAINSEPARATORTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _AugurWallet.contract.Call(opts, out, "DOMAIN_SEPARATOR_TYPEHASH")
	return *ret0, err
}

// DOMAINSEPARATORTYPEHASH is a free data retrieval call binding the contract method 0x1db61b54.
//
// Solidity: function DOMAIN_SEPARATOR_TYPEHASH() view returns(bytes32)
func (_AugurWallet *AugurWalletSession) DOMAINSEPARATORTYPEHASH() ([32]byte, error) {
	return _AugurWallet.Contract.DOMAINSEPARATORTYPEHASH(&_AugurWallet.CallOpts)
}

// DOMAINSEPARATORTYPEHASH is a free data retrieval call binding the contract method 0x1db61b54.
//
// Solidity: function DOMAIN_SEPARATOR_TYPEHASH() view returns(bytes32)
func (_AugurWallet *AugurWalletCallerSession) DOMAINSEPARATORTYPEHASH() ([32]byte, error) {
	return _AugurWallet.Contract.DOMAINSEPARATORTYPEHASH(&_AugurWallet.CallOpts)
}

// MSGTYPEHASH is a free data retrieval call binding the contract method 0x04f86da1.
//
// Solidity: function MSG_TYPEHASH() view returns(bytes32)
func (_AugurWallet *AugurWalletCaller) MSGTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _AugurWallet.contract.Call(opts, out, "MSG_TYPEHASH")
	return *ret0, err
}

// MSGTYPEHASH is a free data retrieval call binding the contract method 0x04f86da1.
//
// Solidity: function MSG_TYPEHASH() view returns(bytes32)
func (_AugurWallet *AugurWalletSession) MSGTYPEHASH() ([32]byte, error) {
	return _AugurWallet.Contract.MSGTYPEHASH(&_AugurWallet.CallOpts)
}

// MSGTYPEHASH is a free data retrieval call binding the contract method 0x04f86da1.
//
// Solidity: function MSG_TYPEHASH() view returns(bytes32)
func (_AugurWallet *AugurWalletCallerSession) MSGTYPEHASH() ([32]byte, error) {
	return _AugurWallet.Contract.MSGTYPEHASH(&_AugurWallet.CallOpts)
}

// AuthorizedProxies is a free data retrieval call binding the contract method 0x360d95b6.
//
// Solidity: function authorizedProxies(address ) view returns(bool)
func (_AugurWallet *AugurWalletCaller) AuthorizedProxies(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AugurWallet.contract.Call(opts, out, "authorizedProxies", arg0)
	return *ret0, err
}

// AuthorizedProxies is a free data retrieval call binding the contract method 0x360d95b6.
//
// Solidity: function authorizedProxies(address ) view returns(bool)
func (_AugurWallet *AugurWalletSession) AuthorizedProxies(arg0 common.Address) (bool, error) {
	return _AugurWallet.Contract.AuthorizedProxies(&_AugurWallet.CallOpts, arg0)
}

// AuthorizedProxies is a free data retrieval call binding the contract method 0x360d95b6.
//
// Solidity: function authorizedProxies(address ) view returns(bool)
func (_AugurWallet *AugurWalletCallerSession) AuthorizedProxies(arg0 common.Address) (bool, error) {
	return _AugurWallet.Contract.AuthorizedProxies(&_AugurWallet.CallOpts, arg0)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_AugurWallet *AugurWalletCaller) Cash(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AugurWallet.contract.Call(opts, out, "cash")
	return *ret0, err
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_AugurWallet *AugurWalletSession) Cash() (common.Address, error) {
	return _AugurWallet.Contract.Cash(&_AugurWallet.CallOpts)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_AugurWallet *AugurWalletCallerSession) Cash() (common.Address, error) {
	return _AugurWallet.Contract.Cash(&_AugurWallet.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_AugurWallet *AugurWalletCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _AugurWallet.contract.Call(opts, out, "domainSeparator")
	return *ret0, err
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_AugurWallet *AugurWalletSession) DomainSeparator() ([32]byte, error) {
	return _AugurWallet.Contract.DomainSeparator(&_AugurWallet.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_AugurWallet *AugurWalletCallerSession) DomainSeparator() ([32]byte, error) {
	return _AugurWallet.Contract.DomainSeparator(&_AugurWallet.CallOpts)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_AugurWallet *AugurWalletCaller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AugurWallet.contract.Call(opts, out, "getAmountOut", amountIn, reserveIn, reserveOut)
	return *ret0, err
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_AugurWallet *AugurWalletSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _AugurWallet.Contract.GetAmountOut(&_AugurWallet.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_AugurWallet *AugurWalletCallerSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _AugurWallet.Contract.GetAmountOut(&_AugurWallet.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() view returns(bool)
func (_AugurWallet *AugurWalletCaller) GetInitialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AugurWallet.contract.Call(opts, out, "getInitialized")
	return *ret0, err
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() view returns(bool)
func (_AugurWallet *AugurWalletSession) GetInitialized() (bool, error) {
	return _AugurWallet.Contract.GetInitialized(&_AugurWallet.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() view returns(bool)
func (_AugurWallet *AugurWalletCallerSession) GetInitialized() (bool, error) {
	return _AugurWallet.Contract.GetInitialized(&_AugurWallet.CallOpts)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x0a1028c4.
//
// Solidity: function getMessageHash(bytes _message) view returns(bytes32)
func (_AugurWallet *AugurWalletCaller) GetMessageHash(opts *bind.CallOpts, _message []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _AugurWallet.contract.Call(opts, out, "getMessageHash", _message)
	return *ret0, err
}

// GetMessageHash is a free data retrieval call binding the contract method 0x0a1028c4.
//
// Solidity: function getMessageHash(bytes _message) view returns(bytes32)
func (_AugurWallet *AugurWalletSession) GetMessageHash(_message []byte) ([32]byte, error) {
	return _AugurWallet.Contract.GetMessageHash(&_AugurWallet.CallOpts, _message)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x0a1028c4.
//
// Solidity: function getMessageHash(bytes _message) view returns(bytes32)
func (_AugurWallet *AugurWalletCallerSession) GetMessageHash(_message []byte) ([32]byte, error) {
	return _AugurWallet.Contract.GetMessageHash(&_AugurWallet.CallOpts, _message)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signature) view returns(bytes4)
func (_AugurWallet *AugurWalletCaller) IsValidSignature(opts *bind.CallOpts, _data []byte, _signature []byte) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _AugurWallet.contract.Call(opts, out, "isValidSignature", _data, _signature)
	return *ret0, err
}

// IsValidSignature is a free data retrieval call binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signature) view returns(bytes4)
func (_AugurWallet *AugurWalletSession) IsValidSignature(_data []byte, _signature []byte) ([4]byte, error) {
	return _AugurWallet.Contract.IsValidSignature(&_AugurWallet.CallOpts, _data, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signature) view returns(bytes4)
func (_AugurWallet *AugurWalletCallerSession) IsValidSignature(_data []byte, _signature []byte) ([4]byte, error) {
	return _AugurWallet.Contract.IsValidSignature(&_AugurWallet.CallOpts, _data, _signature)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_AugurWallet *AugurWalletCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AugurWallet.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_AugurWallet *AugurWalletSession) Registry() (common.Address, error) {
	return _AugurWallet.Contract.Registry(&_AugurWallet.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_AugurWallet *AugurWalletCallerSession) Registry() (common.Address, error) {
	return _AugurWallet.Contract.Registry(&_AugurWallet.CallOpts)
}

// AddAuthorizedProxy is a paid mutator transaction binding the contract method 0x0283e758.
//
// Solidity: function addAuthorizedProxy(address _authorizedProxy) returns(bool)
func (_AugurWallet *AugurWalletTransactor) AddAuthorizedProxy(opts *bind.TransactOpts, _authorizedProxy common.Address) (*types.Transaction, error) {
	return _AugurWallet.contract.Transact(opts, "addAuthorizedProxy", _authorizedProxy)
}

// AddAuthorizedProxy is a paid mutator transaction binding the contract method 0x0283e758.
//
// Solidity: function addAuthorizedProxy(address _authorizedProxy) returns(bool)
func (_AugurWallet *AugurWalletSession) AddAuthorizedProxy(_authorizedProxy common.Address) (*types.Transaction, error) {
	return _AugurWallet.Contract.AddAuthorizedProxy(&_AugurWallet.TransactOpts, _authorizedProxy)
}

// AddAuthorizedProxy is a paid mutator transaction binding the contract method 0x0283e758.
//
// Solidity: function addAuthorizedProxy(address _authorizedProxy) returns(bool)
func (_AugurWallet *AugurWalletTransactorSession) AddAuthorizedProxy(_authorizedProxy common.Address) (*types.Transaction, error) {
	return _AugurWallet.Contract.AddAuthorizedProxy(&_AugurWallet.TransactOpts, _authorizedProxy)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x6b8d7a12.
//
// Solidity: function executeTransaction(address _to, bytes _data, uint256 _value) returns(bool)
func (_AugurWallet *AugurWalletTransactor) ExecuteTransaction(opts *bind.TransactOpts, _to common.Address, _data []byte, _value *big.Int) (*types.Transaction, error) {
	return _AugurWallet.contract.Transact(opts, "executeTransaction", _to, _data, _value)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x6b8d7a12.
//
// Solidity: function executeTransaction(address _to, bytes _data, uint256 _value) returns(bool)
func (_AugurWallet *AugurWalletSession) ExecuteTransaction(_to common.Address, _data []byte, _value *big.Int) (*types.Transaction, error) {
	return _AugurWallet.Contract.ExecuteTransaction(&_AugurWallet.TransactOpts, _to, _data, _value)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x6b8d7a12.
//
// Solidity: function executeTransaction(address _to, bytes _data, uint256 _value) returns(bool)
func (_AugurWallet *AugurWalletTransactorSession) ExecuteTransaction(_to common.Address, _data []byte, _value *big.Int) (*types.Transaction, error) {
	return _AugurWallet.Contract.ExecuteTransaction(&_AugurWallet.TransactOpts, _to, _data, _value)
}

// Initialize is a paid mutator transaction binding the contract method 0xd688e8e5.
//
// Solidity: function initialize(address _owner, address _referralAddress, bytes32 _fingerprint, address _augur, address _registry, address _registryV2, address _cash, address _affiliates, address _shareToken, address _createOrder, address _fillOrder, address _zeroXTrade) returns()
func (_AugurWallet *AugurWalletTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _referralAddress common.Address, _fingerprint [32]byte, _augur common.Address, _registry common.Address, _registryV2 common.Address, _cash common.Address, _affiliates common.Address, _shareToken common.Address, _createOrder common.Address, _fillOrder common.Address, _zeroXTrade common.Address) (*types.Transaction, error) {
	return _AugurWallet.contract.Transact(opts, "initialize", _owner, _referralAddress, _fingerprint, _augur, _registry, _registryV2, _cash, _affiliates, _shareToken, _createOrder, _fillOrder, _zeroXTrade)
}

// Initialize is a paid mutator transaction binding the contract method 0xd688e8e5.
//
// Solidity: function initialize(address _owner, address _referralAddress, bytes32 _fingerprint, address _augur, address _registry, address _registryV2, address _cash, address _affiliates, address _shareToken, address _createOrder, address _fillOrder, address _zeroXTrade) returns()
func (_AugurWallet *AugurWalletSession) Initialize(_owner common.Address, _referralAddress common.Address, _fingerprint [32]byte, _augur common.Address, _registry common.Address, _registryV2 common.Address, _cash common.Address, _affiliates common.Address, _shareToken common.Address, _createOrder common.Address, _fillOrder common.Address, _zeroXTrade common.Address) (*types.Transaction, error) {
	return _AugurWallet.Contract.Initialize(&_AugurWallet.TransactOpts, _owner, _referralAddress, _fingerprint, _augur, _registry, _registryV2, _cash, _affiliates, _shareToken, _createOrder, _fillOrder, _zeroXTrade)
}

// Initialize is a paid mutator transaction binding the contract method 0xd688e8e5.
//
// Solidity: function initialize(address _owner, address _referralAddress, bytes32 _fingerprint, address _augur, address _registry, address _registryV2, address _cash, address _affiliates, address _shareToken, address _createOrder, address _fillOrder, address _zeroXTrade) returns()
func (_AugurWallet *AugurWalletTransactorSession) Initialize(_owner common.Address, _referralAddress common.Address, _fingerprint [32]byte, _augur common.Address, _registry common.Address, _registryV2 common.Address, _cash common.Address, _affiliates common.Address, _shareToken common.Address, _createOrder common.Address, _fillOrder common.Address, _zeroXTrade common.Address) (*types.Transaction, error) {
	return _AugurWallet.Contract.Initialize(&_AugurWallet.TransactOpts, _owner, _referralAddress, _fingerprint, _augur, _registry, _registryV2, _cash, _affiliates, _shareToken, _createOrder, _fillOrder, _zeroXTrade)
}

// RemoveAuthorizedProxy is a paid mutator transaction binding the contract method 0xcbdad2a1.
//
// Solidity: function removeAuthorizedProxy(address _authorizedProxy) returns(bool)
func (_AugurWallet *AugurWalletTransactor) RemoveAuthorizedProxy(opts *bind.TransactOpts, _authorizedProxy common.Address) (*types.Transaction, error) {
	return _AugurWallet.contract.Transact(opts, "removeAuthorizedProxy", _authorizedProxy)
}

// RemoveAuthorizedProxy is a paid mutator transaction binding the contract method 0xcbdad2a1.
//
// Solidity: function removeAuthorizedProxy(address _authorizedProxy) returns(bool)
func (_AugurWallet *AugurWalletSession) RemoveAuthorizedProxy(_authorizedProxy common.Address) (*types.Transaction, error) {
	return _AugurWallet.Contract.RemoveAuthorizedProxy(&_AugurWallet.TransactOpts, _authorizedProxy)
}

// RemoveAuthorizedProxy is a paid mutator transaction binding the contract method 0xcbdad2a1.
//
// Solidity: function removeAuthorizedProxy(address _authorizedProxy) returns(bool)
func (_AugurWallet *AugurWalletTransactorSession) RemoveAuthorizedProxy(_authorizedProxy common.Address) (*types.Transaction, error) {
	return _AugurWallet.Contract.RemoveAuthorizedProxy(&_AugurWallet.TransactOpts, _authorizedProxy)
}

// TransferCash is a paid mutator transaction binding the contract method 0xc067e263.
//
// Solidity: function transferCash(address _to, uint256 _amount) returns()
func (_AugurWallet *AugurWalletTransactor) TransferCash(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AugurWallet.contract.Transact(opts, "transferCash", _to, _amount)
}

// TransferCash is a paid mutator transaction binding the contract method 0xc067e263.
//
// Solidity: function transferCash(address _to, uint256 _amount) returns()
func (_AugurWallet *AugurWalletSession) TransferCash(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AugurWallet.Contract.TransferCash(&_AugurWallet.TransactOpts, _to, _amount)
}

// TransferCash is a paid mutator transaction binding the contract method 0xc067e263.
//
// Solidity: function transferCash(address _to, uint256 _amount) returns()
func (_AugurWallet *AugurWalletTransactorSession) TransferCash(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AugurWallet.Contract.TransferCash(&_AugurWallet.TransactOpts, _to, _amount)
}

// WithdrawAllFundsAsDai is a paid mutator transaction binding the contract method 0x4dd78b76.
//
// Solidity: function withdrawAllFundsAsDai(address _destination, uint256 _minExchangeRateInDai) payable returns(bool)
func (_AugurWallet *AugurWalletTransactor) WithdrawAllFundsAsDai(opts *bind.TransactOpts, _destination common.Address, _minExchangeRateInDai *big.Int) (*types.Transaction, error) {
	return _AugurWallet.contract.Transact(opts, "withdrawAllFundsAsDai", _destination, _minExchangeRateInDai)
}

// WithdrawAllFundsAsDai is a paid mutator transaction binding the contract method 0x4dd78b76.
//
// Solidity: function withdrawAllFundsAsDai(address _destination, uint256 _minExchangeRateInDai) payable returns(bool)
func (_AugurWallet *AugurWalletSession) WithdrawAllFundsAsDai(_destination common.Address, _minExchangeRateInDai *big.Int) (*types.Transaction, error) {
	return _AugurWallet.Contract.WithdrawAllFundsAsDai(&_AugurWallet.TransactOpts, _destination, _minExchangeRateInDai)
}

// WithdrawAllFundsAsDai is a paid mutator transaction binding the contract method 0x4dd78b76.
//
// Solidity: function withdrawAllFundsAsDai(address _destination, uint256 _minExchangeRateInDai) payable returns(bool)
func (_AugurWallet *AugurWalletTransactorSession) WithdrawAllFundsAsDai(_destination common.Address, _minExchangeRateInDai *big.Int) (*types.Transaction, error) {
	return _AugurWallet.Contract.WithdrawAllFundsAsDai(&_AugurWallet.TransactOpts, _destination, _minExchangeRateInDai)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_AugurWallet *AugurWalletTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _AugurWallet.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_AugurWallet *AugurWalletSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _AugurWallet.Contract.Fallback(&_AugurWallet.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_AugurWallet *AugurWalletTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _AugurWallet.Contract.Fallback(&_AugurWallet.TransactOpts, calldata)
}
