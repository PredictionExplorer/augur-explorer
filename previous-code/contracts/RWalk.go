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

// RWalkMetaData contains all meta data concerning the RWalk contract.
var RWalkMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"seed\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"MintEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newName\",\"type\":\"string\"}],\"name\":\"TokenNameEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entropy\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastMintTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastMinter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numWithdrawals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"saleTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"seeds\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"seedsOfOwner\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setTokenName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeUntilSale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeUntilWithdrawal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenGenerationScript\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenNames\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"walletOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawalAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawalNums\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalWaitSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RWalkABI is the input ABI used to generate the binding from.
// Deprecated: Use RWalkMetaData.ABI instead.
var RWalkABI = RWalkMetaData.ABI

// RWalk is an auto generated Go binding around an Ethereum contract.
type RWalk struct {
	RWalkCaller     // Read-only binding to the contract
	RWalkTransactor // Write-only binding to the contract
	RWalkFilterer   // Log filterer for contract events
}

// RWalkCaller is an auto generated read-only Go binding around an Ethereum contract.
type RWalkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RWalkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RWalkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RWalkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RWalkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RWalkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RWalkSession struct {
	Contract     *RWalk            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RWalkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RWalkCallerSession struct {
	Contract *RWalkCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RWalkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RWalkTransactorSession struct {
	Contract     *RWalkTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RWalkRaw is an auto generated low-level Go binding around an Ethereum contract.
type RWalkRaw struct {
	Contract *RWalk // Generic contract binding to access the raw methods on
}

// RWalkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RWalkCallerRaw struct {
	Contract *RWalkCaller // Generic read-only contract binding to access the raw methods on
}

// RWalkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RWalkTransactorRaw struct {
	Contract *RWalkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRWalk creates a new instance of RWalk, bound to a specific deployed contract.
func NewRWalk(address common.Address, backend bind.ContractBackend) (*RWalk, error) {
	contract, err := bindRWalk(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RWalk{RWalkCaller: RWalkCaller{contract: contract}, RWalkTransactor: RWalkTransactor{contract: contract}, RWalkFilterer: RWalkFilterer{contract: contract}}, nil
}

// NewRWalkCaller creates a new read-only instance of RWalk, bound to a specific deployed contract.
func NewRWalkCaller(address common.Address, caller bind.ContractCaller) (*RWalkCaller, error) {
	contract, err := bindRWalk(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RWalkCaller{contract: contract}, nil
}

// NewRWalkTransactor creates a new write-only instance of RWalk, bound to a specific deployed contract.
func NewRWalkTransactor(address common.Address, transactor bind.ContractTransactor) (*RWalkTransactor, error) {
	contract, err := bindRWalk(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RWalkTransactor{contract: contract}, nil
}

// NewRWalkFilterer creates a new log filterer instance of RWalk, bound to a specific deployed contract.
func NewRWalkFilterer(address common.Address, filterer bind.ContractFilterer) (*RWalkFilterer, error) {
	contract, err := bindRWalk(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RWalkFilterer{contract: contract}, nil
}

// bindRWalk binds a generic wrapper to an already deployed contract.
func bindRWalk(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RWalkABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RWalk *RWalkRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RWalk.Contract.RWalkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RWalk *RWalkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RWalk.Contract.RWalkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RWalk *RWalkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RWalk.Contract.RWalkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RWalk *RWalkCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RWalk.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RWalk *RWalkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RWalk.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RWalk *RWalkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RWalk.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_RWalk *RWalkCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_RWalk *RWalkSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _RWalk.Contract.BalanceOf(&_RWalk.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_RWalk *RWalkCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _RWalk.Contract.BalanceOf(&_RWalk.CallOpts, owner)
}

// Entropy is a free data retrieval call binding the contract method 0x47ce07cc.
//
// Solidity: function entropy() view returns(bytes32)
func (_RWalk *RWalkCaller) Entropy(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "entropy")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Entropy is a free data retrieval call binding the contract method 0x47ce07cc.
//
// Solidity: function entropy() view returns(bytes32)
func (_RWalk *RWalkSession) Entropy() ([32]byte, error) {
	return _RWalk.Contract.Entropy(&_RWalk.CallOpts)
}

// Entropy is a free data retrieval call binding the contract method 0x47ce07cc.
//
// Solidity: function entropy() view returns(bytes32)
func (_RWalk *RWalkCallerSession) Entropy() ([32]byte, error) {
	return _RWalk.Contract.Entropy(&_RWalk.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_RWalk *RWalkCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_RWalk *RWalkSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _RWalk.Contract.GetApproved(&_RWalk.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_RWalk *RWalkCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _RWalk.Contract.GetApproved(&_RWalk.CallOpts, tokenId)
}

// GetMintPrice is a free data retrieval call binding the contract method 0xa7f93ebd.
//
// Solidity: function getMintPrice() view returns(uint256)
func (_RWalk *RWalkCaller) GetMintPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "getMintPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMintPrice is a free data retrieval call binding the contract method 0xa7f93ebd.
//
// Solidity: function getMintPrice() view returns(uint256)
func (_RWalk *RWalkSession) GetMintPrice() (*big.Int, error) {
	return _RWalk.Contract.GetMintPrice(&_RWalk.CallOpts)
}

// GetMintPrice is a free data retrieval call binding the contract method 0xa7f93ebd.
//
// Solidity: function getMintPrice() view returns(uint256)
func (_RWalk *RWalkCallerSession) GetMintPrice() (*big.Int, error) {
	return _RWalk.Contract.GetMintPrice(&_RWalk.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_RWalk *RWalkCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_RWalk *RWalkSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _RWalk.Contract.IsApprovedForAll(&_RWalk.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_RWalk *RWalkCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _RWalk.Contract.IsApprovedForAll(&_RWalk.CallOpts, owner, operator)
}

// LastMintTime is a free data retrieval call binding the contract method 0x9d463520.
//
// Solidity: function lastMintTime() view returns(uint256)
func (_RWalk *RWalkCaller) LastMintTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "lastMintTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastMintTime is a free data retrieval call binding the contract method 0x9d463520.
//
// Solidity: function lastMintTime() view returns(uint256)
func (_RWalk *RWalkSession) LastMintTime() (*big.Int, error) {
	return _RWalk.Contract.LastMintTime(&_RWalk.CallOpts)
}

// LastMintTime is a free data retrieval call binding the contract method 0x9d463520.
//
// Solidity: function lastMintTime() view returns(uint256)
func (_RWalk *RWalkCallerSession) LastMintTime() (*big.Int, error) {
	return _RWalk.Contract.LastMintTime(&_RWalk.CallOpts)
}

// LastMinter is a free data retrieval call binding the contract method 0x4cd609bc.
//
// Solidity: function lastMinter() view returns(address)
func (_RWalk *RWalkCaller) LastMinter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "lastMinter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LastMinter is a free data retrieval call binding the contract method 0x4cd609bc.
//
// Solidity: function lastMinter() view returns(address)
func (_RWalk *RWalkSession) LastMinter() (common.Address, error) {
	return _RWalk.Contract.LastMinter(&_RWalk.CallOpts)
}

// LastMinter is a free data retrieval call binding the contract method 0x4cd609bc.
//
// Solidity: function lastMinter() view returns(address)
func (_RWalk *RWalkCallerSession) LastMinter() (common.Address, error) {
	return _RWalk.Contract.LastMinter(&_RWalk.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RWalk *RWalkCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RWalk *RWalkSession) Name() (string, error) {
	return _RWalk.Contract.Name(&_RWalk.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RWalk *RWalkCallerSession) Name() (string, error) {
	return _RWalk.Contract.Name(&_RWalk.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_RWalk *RWalkCaller) NextTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "nextTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_RWalk *RWalkSession) NextTokenId() (*big.Int, error) {
	return _RWalk.Contract.NextTokenId(&_RWalk.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_RWalk *RWalkCallerSession) NextTokenId() (*big.Int, error) {
	return _RWalk.Contract.NextTokenId(&_RWalk.CallOpts)
}

// NumWithdrawals is a free data retrieval call binding the contract method 0x157e3945.
//
// Solidity: function numWithdrawals() view returns(uint256)
func (_RWalk *RWalkCaller) NumWithdrawals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "numWithdrawals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumWithdrawals is a free data retrieval call binding the contract method 0x157e3945.
//
// Solidity: function numWithdrawals() view returns(uint256)
func (_RWalk *RWalkSession) NumWithdrawals() (*big.Int, error) {
	return _RWalk.Contract.NumWithdrawals(&_RWalk.CallOpts)
}

// NumWithdrawals is a free data retrieval call binding the contract method 0x157e3945.
//
// Solidity: function numWithdrawals() view returns(uint256)
func (_RWalk *RWalkCallerSession) NumWithdrawals() (*big.Int, error) {
	return _RWalk.Contract.NumWithdrawals(&_RWalk.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RWalk *RWalkCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RWalk *RWalkSession) Owner() (common.Address, error) {
	return _RWalk.Contract.Owner(&_RWalk.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RWalk *RWalkCallerSession) Owner() (common.Address, error) {
	return _RWalk.Contract.Owner(&_RWalk.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_RWalk *RWalkCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_RWalk *RWalkSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _RWalk.Contract.OwnerOf(&_RWalk.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_RWalk *RWalkCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _RWalk.Contract.OwnerOf(&_RWalk.CallOpts, tokenId)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_RWalk *RWalkCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_RWalk *RWalkSession) Price() (*big.Int, error) {
	return _RWalk.Contract.Price(&_RWalk.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_RWalk *RWalkCallerSession) Price() (*big.Int, error) {
	return _RWalk.Contract.Price(&_RWalk.CallOpts)
}

// SaleTime is a free data retrieval call binding the contract method 0x1596facb.
//
// Solidity: function saleTime() view returns(uint256)
func (_RWalk *RWalkCaller) SaleTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "saleTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SaleTime is a free data retrieval call binding the contract method 0x1596facb.
//
// Solidity: function saleTime() view returns(uint256)
func (_RWalk *RWalkSession) SaleTime() (*big.Int, error) {
	return _RWalk.Contract.SaleTime(&_RWalk.CallOpts)
}

// SaleTime is a free data retrieval call binding the contract method 0x1596facb.
//
// Solidity: function saleTime() view returns(uint256)
func (_RWalk *RWalkCallerSession) SaleTime() (*big.Int, error) {
	return _RWalk.Contract.SaleTime(&_RWalk.CallOpts)
}

// Seeds is a free data retrieval call binding the contract method 0xf0503e80.
//
// Solidity: function seeds(uint256 ) view returns(bytes32)
func (_RWalk *RWalkCaller) Seeds(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "seeds", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Seeds is a free data retrieval call binding the contract method 0xf0503e80.
//
// Solidity: function seeds(uint256 ) view returns(bytes32)
func (_RWalk *RWalkSession) Seeds(arg0 *big.Int) ([32]byte, error) {
	return _RWalk.Contract.Seeds(&_RWalk.CallOpts, arg0)
}

// Seeds is a free data retrieval call binding the contract method 0xf0503e80.
//
// Solidity: function seeds(uint256 ) view returns(bytes32)
func (_RWalk *RWalkCallerSession) Seeds(arg0 *big.Int) ([32]byte, error) {
	return _RWalk.Contract.Seeds(&_RWalk.CallOpts, arg0)
}

// SeedsOfOwner is a free data retrieval call binding the contract method 0xcb8efe95.
//
// Solidity: function seedsOfOwner(address _owner) view returns(bytes32[])
func (_RWalk *RWalkCaller) SeedsOfOwner(opts *bind.CallOpts, _owner common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "seedsOfOwner", _owner)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// SeedsOfOwner is a free data retrieval call binding the contract method 0xcb8efe95.
//
// Solidity: function seedsOfOwner(address _owner) view returns(bytes32[])
func (_RWalk *RWalkSession) SeedsOfOwner(_owner common.Address) ([][32]byte, error) {
	return _RWalk.Contract.SeedsOfOwner(&_RWalk.CallOpts, _owner)
}

// SeedsOfOwner is a free data retrieval call binding the contract method 0xcb8efe95.
//
// Solidity: function seedsOfOwner(address _owner) view returns(bytes32[])
func (_RWalk *RWalkCallerSession) SeedsOfOwner(_owner common.Address) ([][32]byte, error) {
	return _RWalk.Contract.SeedsOfOwner(&_RWalk.CallOpts, _owner)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RWalk *RWalkCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RWalk *RWalkSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RWalk.Contract.SupportsInterface(&_RWalk.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RWalk *RWalkCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RWalk.Contract.SupportsInterface(&_RWalk.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RWalk *RWalkCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RWalk *RWalkSession) Symbol() (string, error) {
	return _RWalk.Contract.Symbol(&_RWalk.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RWalk *RWalkCallerSession) Symbol() (string, error) {
	return _RWalk.Contract.Symbol(&_RWalk.CallOpts)
}

// TimeUntilSale is a free data retrieval call binding the contract method 0x661a1cfe.
//
// Solidity: function timeUntilSale() view returns(uint256)
func (_RWalk *RWalkCaller) TimeUntilSale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "timeUntilSale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeUntilSale is a free data retrieval call binding the contract method 0x661a1cfe.
//
// Solidity: function timeUntilSale() view returns(uint256)
func (_RWalk *RWalkSession) TimeUntilSale() (*big.Int, error) {
	return _RWalk.Contract.TimeUntilSale(&_RWalk.CallOpts)
}

// TimeUntilSale is a free data retrieval call binding the contract method 0x661a1cfe.
//
// Solidity: function timeUntilSale() view returns(uint256)
func (_RWalk *RWalkCallerSession) TimeUntilSale() (*big.Int, error) {
	return _RWalk.Contract.TimeUntilSale(&_RWalk.CallOpts)
}

// TimeUntilWithdrawal is a free data retrieval call binding the contract method 0x3e8af4da.
//
// Solidity: function timeUntilWithdrawal() view returns(uint256)
func (_RWalk *RWalkCaller) TimeUntilWithdrawal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "timeUntilWithdrawal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeUntilWithdrawal is a free data retrieval call binding the contract method 0x3e8af4da.
//
// Solidity: function timeUntilWithdrawal() view returns(uint256)
func (_RWalk *RWalkSession) TimeUntilWithdrawal() (*big.Int, error) {
	return _RWalk.Contract.TimeUntilWithdrawal(&_RWalk.CallOpts)
}

// TimeUntilWithdrawal is a free data retrieval call binding the contract method 0x3e8af4da.
//
// Solidity: function timeUntilWithdrawal() view returns(uint256)
func (_RWalk *RWalkCallerSession) TimeUntilWithdrawal() (*big.Int, error) {
	return _RWalk.Contract.TimeUntilWithdrawal(&_RWalk.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_RWalk *RWalkCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_RWalk *RWalkSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _RWalk.Contract.TokenByIndex(&_RWalk.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_RWalk *RWalkCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _RWalk.Contract.TokenByIndex(&_RWalk.CallOpts, index)
}

// TokenGenerationScript is a free data retrieval call binding the contract method 0xee67d6a0.
//
// Solidity: function tokenGenerationScript() view returns(string)
func (_RWalk *RWalkCaller) TokenGenerationScript(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "tokenGenerationScript")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenGenerationScript is a free data retrieval call binding the contract method 0xee67d6a0.
//
// Solidity: function tokenGenerationScript() view returns(string)
func (_RWalk *RWalkSession) TokenGenerationScript() (string, error) {
	return _RWalk.Contract.TokenGenerationScript(&_RWalk.CallOpts)
}

// TokenGenerationScript is a free data retrieval call binding the contract method 0xee67d6a0.
//
// Solidity: function tokenGenerationScript() view returns(string)
func (_RWalk *RWalkCallerSession) TokenGenerationScript() (string, error) {
	return _RWalk.Contract.TokenGenerationScript(&_RWalk.CallOpts)
}

// TokenNames is a free data retrieval call binding the contract method 0x310495ab.
//
// Solidity: function tokenNames(uint256 ) view returns(string)
func (_RWalk *RWalkCaller) TokenNames(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "tokenNames", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenNames is a free data retrieval call binding the contract method 0x310495ab.
//
// Solidity: function tokenNames(uint256 ) view returns(string)
func (_RWalk *RWalkSession) TokenNames(arg0 *big.Int) (string, error) {
	return _RWalk.Contract.TokenNames(&_RWalk.CallOpts, arg0)
}

// TokenNames is a free data retrieval call binding the contract method 0x310495ab.
//
// Solidity: function tokenNames(uint256 ) view returns(string)
func (_RWalk *RWalkCallerSession) TokenNames(arg0 *big.Int) (string, error) {
	return _RWalk.Contract.TokenNames(&_RWalk.CallOpts, arg0)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_RWalk *RWalkCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_RWalk *RWalkSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _RWalk.Contract.TokenOfOwnerByIndex(&_RWalk.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_RWalk *RWalkCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _RWalk.Contract.TokenOfOwnerByIndex(&_RWalk.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_RWalk *RWalkCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_RWalk *RWalkSession) TokenURI(tokenId *big.Int) (string, error) {
	return _RWalk.Contract.TokenURI(&_RWalk.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_RWalk *RWalkCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _RWalk.Contract.TokenURI(&_RWalk.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RWalk *RWalkCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RWalk *RWalkSession) TotalSupply() (*big.Int, error) {
	return _RWalk.Contract.TotalSupply(&_RWalk.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RWalk *RWalkCallerSession) TotalSupply() (*big.Int, error) {
	return _RWalk.Contract.TotalSupply(&_RWalk.CallOpts)
}

// WalletOfOwner is a free data retrieval call binding the contract method 0x438b6300.
//
// Solidity: function walletOfOwner(address _owner) view returns(uint256[])
func (_RWalk *RWalkCaller) WalletOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "walletOfOwner", _owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// WalletOfOwner is a free data retrieval call binding the contract method 0x438b6300.
//
// Solidity: function walletOfOwner(address _owner) view returns(uint256[])
func (_RWalk *RWalkSession) WalletOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _RWalk.Contract.WalletOfOwner(&_RWalk.CallOpts, _owner)
}

// WalletOfOwner is a free data retrieval call binding the contract method 0x438b6300.
//
// Solidity: function walletOfOwner(address _owner) view returns(uint256[])
func (_RWalk *RWalkCallerSession) WalletOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _RWalk.Contract.WalletOfOwner(&_RWalk.CallOpts, _owner)
}

// WithdrawalAmount is a free data retrieval call binding the contract method 0xe6e268f4.
//
// Solidity: function withdrawalAmount() view returns(uint256)
func (_RWalk *RWalkCaller) WithdrawalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "withdrawalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalAmount is a free data retrieval call binding the contract method 0xe6e268f4.
//
// Solidity: function withdrawalAmount() view returns(uint256)
func (_RWalk *RWalkSession) WithdrawalAmount() (*big.Int, error) {
	return _RWalk.Contract.WithdrawalAmount(&_RWalk.CallOpts)
}

// WithdrawalAmount is a free data retrieval call binding the contract method 0xe6e268f4.
//
// Solidity: function withdrawalAmount() view returns(uint256)
func (_RWalk *RWalkCallerSession) WithdrawalAmount() (*big.Int, error) {
	return _RWalk.Contract.WithdrawalAmount(&_RWalk.CallOpts)
}

// WithdrawalAmounts is a free data retrieval call binding the contract method 0xf454aae1.
//
// Solidity: function withdrawalAmounts(uint256 ) view returns(uint256)
func (_RWalk *RWalkCaller) WithdrawalAmounts(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "withdrawalAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalAmounts is a free data retrieval call binding the contract method 0xf454aae1.
//
// Solidity: function withdrawalAmounts(uint256 ) view returns(uint256)
func (_RWalk *RWalkSession) WithdrawalAmounts(arg0 *big.Int) (*big.Int, error) {
	return _RWalk.Contract.WithdrawalAmounts(&_RWalk.CallOpts, arg0)
}

// WithdrawalAmounts is a free data retrieval call binding the contract method 0xf454aae1.
//
// Solidity: function withdrawalAmounts(uint256 ) view returns(uint256)
func (_RWalk *RWalkCallerSession) WithdrawalAmounts(arg0 *big.Int) (*big.Int, error) {
	return _RWalk.Contract.WithdrawalAmounts(&_RWalk.CallOpts, arg0)
}

// WithdrawalNums is a free data retrieval call binding the contract method 0x6e56f6f9.
//
// Solidity: function withdrawalNums(uint256 ) view returns(uint256)
func (_RWalk *RWalkCaller) WithdrawalNums(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "withdrawalNums", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalNums is a free data retrieval call binding the contract method 0x6e56f6f9.
//
// Solidity: function withdrawalNums(uint256 ) view returns(uint256)
func (_RWalk *RWalkSession) WithdrawalNums(arg0 *big.Int) (*big.Int, error) {
	return _RWalk.Contract.WithdrawalNums(&_RWalk.CallOpts, arg0)
}

// WithdrawalNums is a free data retrieval call binding the contract method 0x6e56f6f9.
//
// Solidity: function withdrawalNums(uint256 ) view returns(uint256)
func (_RWalk *RWalkCallerSession) WithdrawalNums(arg0 *big.Int) (*big.Int, error) {
	return _RWalk.Contract.WithdrawalNums(&_RWalk.CallOpts, arg0)
}

// WithdrawalWaitSeconds is a free data retrieval call binding the contract method 0x17d209ab.
//
// Solidity: function withdrawalWaitSeconds() view returns(uint256)
func (_RWalk *RWalkCaller) WithdrawalWaitSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RWalk.contract.Call(opts, &out, "withdrawalWaitSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalWaitSeconds is a free data retrieval call binding the contract method 0x17d209ab.
//
// Solidity: function withdrawalWaitSeconds() view returns(uint256)
func (_RWalk *RWalkSession) WithdrawalWaitSeconds() (*big.Int, error) {
	return _RWalk.Contract.WithdrawalWaitSeconds(&_RWalk.CallOpts)
}

// WithdrawalWaitSeconds is a free data retrieval call binding the contract method 0x17d209ab.
//
// Solidity: function withdrawalWaitSeconds() view returns(uint256)
func (_RWalk *RWalkCallerSession) WithdrawalWaitSeconds() (*big.Int, error) {
	return _RWalk.Contract.WithdrawalWaitSeconds(&_RWalk.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_RWalk *RWalkTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RWalk.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_RWalk *RWalkSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RWalk.Contract.Approve(&_RWalk.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_RWalk *RWalkTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RWalk.Contract.Approve(&_RWalk.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_RWalk *RWalkTransactor) Mint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RWalk.contract.Transact(opts, "mint")
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_RWalk *RWalkSession) Mint() (*types.Transaction, error) {
	return _RWalk.Contract.Mint(&_RWalk.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_RWalk *RWalkTransactorSession) Mint() (*types.Transaction, error) {
	return _RWalk.Contract.Mint(&_RWalk.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RWalk *RWalkTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RWalk.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RWalk *RWalkSession) RenounceOwnership() (*types.Transaction, error) {
	return _RWalk.Contract.RenounceOwnership(&_RWalk.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RWalk *RWalkTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RWalk.Contract.RenounceOwnership(&_RWalk.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_RWalk *RWalkTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RWalk.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_RWalk *RWalkSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RWalk.Contract.SafeTransferFrom(&_RWalk.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_RWalk *RWalkTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RWalk.Contract.SafeTransferFrom(&_RWalk.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_RWalk *RWalkTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _RWalk.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_RWalk *RWalkSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _RWalk.Contract.SafeTransferFrom0(&_RWalk.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_RWalk *RWalkTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _RWalk.Contract.SafeTransferFrom0(&_RWalk.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_RWalk *RWalkTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _RWalk.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_RWalk *RWalkSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _RWalk.Contract.SetApprovalForAll(&_RWalk.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_RWalk *RWalkTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _RWalk.Contract.SetApprovalForAll(&_RWalk.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_RWalk *RWalkTransactor) SetBaseURI(opts *bind.TransactOpts, baseURI string) (*types.Transaction, error) {
	return _RWalk.contract.Transact(opts, "setBaseURI", baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_RWalk *RWalkSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _RWalk.Contract.SetBaseURI(&_RWalk.TransactOpts, baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_RWalk *RWalkTransactorSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _RWalk.Contract.SetBaseURI(&_RWalk.TransactOpts, baseURI)
}

// SetTokenName is a paid mutator transaction binding the contract method 0xcdb0e89e.
//
// Solidity: function setTokenName(uint256 tokenId, string name) returns()
func (_RWalk *RWalkTransactor) SetTokenName(opts *bind.TransactOpts, tokenId *big.Int, name string) (*types.Transaction, error) {
	return _RWalk.contract.Transact(opts, "setTokenName", tokenId, name)
}

// SetTokenName is a paid mutator transaction binding the contract method 0xcdb0e89e.
//
// Solidity: function setTokenName(uint256 tokenId, string name) returns()
func (_RWalk *RWalkSession) SetTokenName(tokenId *big.Int, name string) (*types.Transaction, error) {
	return _RWalk.Contract.SetTokenName(&_RWalk.TransactOpts, tokenId, name)
}

// SetTokenName is a paid mutator transaction binding the contract method 0xcdb0e89e.
//
// Solidity: function setTokenName(uint256 tokenId, string name) returns()
func (_RWalk *RWalkTransactorSession) SetTokenName(tokenId *big.Int, name string) (*types.Transaction, error) {
	return _RWalk.Contract.SetTokenName(&_RWalk.TransactOpts, tokenId, name)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_RWalk *RWalkTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RWalk.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_RWalk *RWalkSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RWalk.Contract.TransferFrom(&_RWalk.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_RWalk *RWalkTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RWalk.Contract.TransferFrom(&_RWalk.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RWalk *RWalkTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RWalk.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RWalk *RWalkSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RWalk.Contract.TransferOwnership(&_RWalk.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RWalk *RWalkTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RWalk.Contract.TransferOwnership(&_RWalk.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_RWalk *RWalkTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RWalk.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_RWalk *RWalkSession) Withdraw() (*types.Transaction, error) {
	return _RWalk.Contract.Withdraw(&_RWalk.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_RWalk *RWalkTransactorSession) Withdraw() (*types.Transaction, error) {
	return _RWalk.Contract.Withdraw(&_RWalk.TransactOpts)
}

// RWalkApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the RWalk contract.
type RWalkApprovalIterator struct {
	Event *RWalkApproval // Event containing the contract specifics and raw log

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
func (it *RWalkApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWalkApproval)
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
		it.Event = new(RWalkApproval)
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
func (it *RWalkApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWalkApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWalkApproval represents a Approval event raised by the RWalk contract.
type RWalkApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_RWalk *RWalkFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*RWalkApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _RWalk.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &RWalkApprovalIterator{contract: _RWalk.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_RWalk *RWalkFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *RWalkApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _RWalk.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWalkApproval)
				if err := _RWalk.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_RWalk *RWalkFilterer) ParseApproval(log types.Log) (*RWalkApproval, error) {
	event := new(RWalkApproval)
	if err := _RWalk.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWalkApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the RWalk contract.
type RWalkApprovalForAllIterator struct {
	Event *RWalkApprovalForAll // Event containing the contract specifics and raw log

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
func (it *RWalkApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWalkApprovalForAll)
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
		it.Event = new(RWalkApprovalForAll)
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
func (it *RWalkApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWalkApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWalkApprovalForAll represents a ApprovalForAll event raised by the RWalk contract.
type RWalkApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_RWalk *RWalkFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*RWalkApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _RWalk.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &RWalkApprovalForAllIterator{contract: _RWalk.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_RWalk *RWalkFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *RWalkApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _RWalk.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWalkApprovalForAll)
				if err := _RWalk.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_RWalk *RWalkFilterer) ParseApprovalForAll(log types.Log) (*RWalkApprovalForAll, error) {
	event := new(RWalkApprovalForAll)
	if err := _RWalk.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWalkMintEventIterator is returned from FilterMintEvent and is used to iterate over the raw logs and unpacked data for MintEvent events raised by the RWalk contract.
type RWalkMintEventIterator struct {
	Event *RWalkMintEvent // Event containing the contract specifics and raw log

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
func (it *RWalkMintEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWalkMintEvent)
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
		it.Event = new(RWalkMintEvent)
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
func (it *RWalkMintEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWalkMintEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWalkMintEvent represents a MintEvent event raised by the RWalk contract.
type RWalkMintEvent struct {
	TokenId *big.Int
	Owner   common.Address
	Seed    [32]byte
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMintEvent is a free log retrieval operation binding the contract event 0xad2bc79f659de022c64ef55c71f16d0cf125452ed5fc5757b2edc331f58565ec.
//
// Solidity: event MintEvent(uint256 indexed tokenId, address indexed owner, bytes32 seed, uint256 price)
func (_RWalk *RWalkFilterer) FilterMintEvent(opts *bind.FilterOpts, tokenId []*big.Int, owner []common.Address) (*RWalkMintEventIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _RWalk.contract.FilterLogs(opts, "MintEvent", tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &RWalkMintEventIterator{contract: _RWalk.contract, event: "MintEvent", logs: logs, sub: sub}, nil
}

// WatchMintEvent is a free log subscription operation binding the contract event 0xad2bc79f659de022c64ef55c71f16d0cf125452ed5fc5757b2edc331f58565ec.
//
// Solidity: event MintEvent(uint256 indexed tokenId, address indexed owner, bytes32 seed, uint256 price)
func (_RWalk *RWalkFilterer) WatchMintEvent(opts *bind.WatchOpts, sink chan<- *RWalkMintEvent, tokenId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _RWalk.contract.WatchLogs(opts, "MintEvent", tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWalkMintEvent)
				if err := _RWalk.contract.UnpackLog(event, "MintEvent", log); err != nil {
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

// ParseMintEvent is a log parse operation binding the contract event 0xad2bc79f659de022c64ef55c71f16d0cf125452ed5fc5757b2edc331f58565ec.
//
// Solidity: event MintEvent(uint256 indexed tokenId, address indexed owner, bytes32 seed, uint256 price)
func (_RWalk *RWalkFilterer) ParseMintEvent(log types.Log) (*RWalkMintEvent, error) {
	event := new(RWalkMintEvent)
	if err := _RWalk.contract.UnpackLog(event, "MintEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWalkOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RWalk contract.
type RWalkOwnershipTransferredIterator struct {
	Event *RWalkOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RWalkOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWalkOwnershipTransferred)
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
		it.Event = new(RWalkOwnershipTransferred)
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
func (it *RWalkOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWalkOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWalkOwnershipTransferred represents a OwnershipTransferred event raised by the RWalk contract.
type RWalkOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RWalk *RWalkFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RWalkOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RWalk.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RWalkOwnershipTransferredIterator{contract: _RWalk.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RWalk *RWalkFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RWalkOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RWalk.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWalkOwnershipTransferred)
				if err := _RWalk.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RWalk *RWalkFilterer) ParseOwnershipTransferred(log types.Log) (*RWalkOwnershipTransferred, error) {
	event := new(RWalkOwnershipTransferred)
	if err := _RWalk.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWalkTokenNameEventIterator is returned from FilterTokenNameEvent and is used to iterate over the raw logs and unpacked data for TokenNameEvent events raised by the RWalk contract.
type RWalkTokenNameEventIterator struct {
	Event *RWalkTokenNameEvent // Event containing the contract specifics and raw log

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
func (it *RWalkTokenNameEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWalkTokenNameEvent)
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
		it.Event = new(RWalkTokenNameEvent)
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
func (it *RWalkTokenNameEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWalkTokenNameEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWalkTokenNameEvent represents a TokenNameEvent event raised by the RWalk contract.
type RWalkTokenNameEvent struct {
	TokenId *big.Int
	NewName string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenNameEvent is a free log retrieval operation binding the contract event 0x8ad5e159ff95649c8a9f323ac5a457e741897cf44ce07dfce0e98b84ef9d5f12.
//
// Solidity: event TokenNameEvent(uint256 tokenId, string newName)
func (_RWalk *RWalkFilterer) FilterTokenNameEvent(opts *bind.FilterOpts) (*RWalkTokenNameEventIterator, error) {

	logs, sub, err := _RWalk.contract.FilterLogs(opts, "TokenNameEvent")
	if err != nil {
		return nil, err
	}
	return &RWalkTokenNameEventIterator{contract: _RWalk.contract, event: "TokenNameEvent", logs: logs, sub: sub}, nil
}

// WatchTokenNameEvent is a free log subscription operation binding the contract event 0x8ad5e159ff95649c8a9f323ac5a457e741897cf44ce07dfce0e98b84ef9d5f12.
//
// Solidity: event TokenNameEvent(uint256 tokenId, string newName)
func (_RWalk *RWalkFilterer) WatchTokenNameEvent(opts *bind.WatchOpts, sink chan<- *RWalkTokenNameEvent) (event.Subscription, error) {

	logs, sub, err := _RWalk.contract.WatchLogs(opts, "TokenNameEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWalkTokenNameEvent)
				if err := _RWalk.contract.UnpackLog(event, "TokenNameEvent", log); err != nil {
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

// ParseTokenNameEvent is a log parse operation binding the contract event 0x8ad5e159ff95649c8a9f323ac5a457e741897cf44ce07dfce0e98b84ef9d5f12.
//
// Solidity: event TokenNameEvent(uint256 tokenId, string newName)
func (_RWalk *RWalkFilterer) ParseTokenNameEvent(log types.Log) (*RWalkTokenNameEvent, error) {
	event := new(RWalkTokenNameEvent)
	if err := _RWalk.contract.UnpackLog(event, "TokenNameEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWalkTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the RWalk contract.
type RWalkTransferIterator struct {
	Event *RWalkTransfer // Event containing the contract specifics and raw log

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
func (it *RWalkTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWalkTransfer)
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
		it.Event = new(RWalkTransfer)
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
func (it *RWalkTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWalkTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWalkTransfer represents a Transfer event raised by the RWalk contract.
type RWalkTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_RWalk *RWalkFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*RWalkTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _RWalk.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &RWalkTransferIterator{contract: _RWalk.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_RWalk *RWalkFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RWalkTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _RWalk.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWalkTransfer)
				if err := _RWalk.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_RWalk *RWalkFilterer) ParseTransfer(log types.Log) (*RWalkTransfer, error) {
	event := new(RWalkTransfer)
	if err := _RWalk.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWalkWithdrawalEventIterator is returned from FilterWithdrawalEvent and is used to iterate over the raw logs and unpacked data for WithdrawalEvent events raised by the RWalk contract.
type RWalkWithdrawalEventIterator struct {
	Event *RWalkWithdrawalEvent // Event containing the contract specifics and raw log

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
func (it *RWalkWithdrawalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWalkWithdrawalEvent)
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
		it.Event = new(RWalkWithdrawalEvent)
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
func (it *RWalkWithdrawalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWalkWithdrawalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWalkWithdrawalEvent represents a WithdrawalEvent event raised by the RWalk contract.
type RWalkWithdrawalEvent struct {
	TokenId     *big.Int
	Destination common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalEvent is a free log retrieval operation binding the contract event 0xa11b556ace4b11a5cae8675a293b51e8cde3a06387d34010861789dfd9e9abc7.
//
// Solidity: event WithdrawalEvent(uint256 indexed tokenId, address destination, uint256 amount)
func (_RWalk *RWalkFilterer) FilterWithdrawalEvent(opts *bind.FilterOpts, tokenId []*big.Int) (*RWalkWithdrawalEventIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _RWalk.contract.FilterLogs(opts, "WithdrawalEvent", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &RWalkWithdrawalEventIterator{contract: _RWalk.contract, event: "WithdrawalEvent", logs: logs, sub: sub}, nil
}

// WatchWithdrawalEvent is a free log subscription operation binding the contract event 0xa11b556ace4b11a5cae8675a293b51e8cde3a06387d34010861789dfd9e9abc7.
//
// Solidity: event WithdrawalEvent(uint256 indexed tokenId, address destination, uint256 amount)
func (_RWalk *RWalkFilterer) WatchWithdrawalEvent(opts *bind.WatchOpts, sink chan<- *RWalkWithdrawalEvent, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _RWalk.contract.WatchLogs(opts, "WithdrawalEvent", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWalkWithdrawalEvent)
				if err := _RWalk.contract.UnpackLog(event, "WithdrawalEvent", log); err != nil {
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

// ParseWithdrawalEvent is a log parse operation binding the contract event 0xa11b556ace4b11a5cae8675a293b51e8cde3a06387d34010861789dfd9e9abc7.
//
// Solidity: event WithdrawalEvent(uint256 indexed tokenId, address destination, uint256 amount)
func (_RWalk *RWalkFilterer) ParseWithdrawalEvent(log types.Log) (*RWalkWithdrawalEvent, error) {
	event := new(RWalkWithdrawalEvent)
	if err := _RWalk.contract.UnpackLog(event, "WithdrawalEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
