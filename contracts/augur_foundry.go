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

// AugurFoundryMetaData contains all meta data concerning the AugurFoundry contract.
var AugurFoundryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIShareToken\",\"name\":\"_shareToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_cash\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"WrapperCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"cash\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"name\":\"newERC20Wrapper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"_names\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_symbols\",\"type\":\"string[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_decimals\",\"type\":\"uint8[]\"}],\"name\":\"newERC20Wrappers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shareToken\",\"outputs\":[{\"internalType\":\"contractIShareToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"unWrapMultipleTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unWrapTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"wrapMultipleTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"wrapTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"wrappers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AugurFoundryABI is the input ABI used to generate the binding from.
// Deprecated: Use AugurFoundryMetaData.ABI instead.
var AugurFoundryABI = AugurFoundryMetaData.ABI

// AugurFoundry is an auto generated Go binding around an Ethereum contract.
type AugurFoundry struct {
	AugurFoundryCaller     // Read-only binding to the contract
	AugurFoundryTransactor // Write-only binding to the contract
	AugurFoundryFilterer   // Log filterer for contract events
}

// AugurFoundryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AugurFoundryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AugurFoundryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AugurFoundryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AugurFoundryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AugurFoundryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AugurFoundrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AugurFoundrySession struct {
	Contract     *AugurFoundry     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AugurFoundryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AugurFoundryCallerSession struct {
	Contract *AugurFoundryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AugurFoundryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AugurFoundryTransactorSession struct {
	Contract     *AugurFoundryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AugurFoundryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AugurFoundryRaw struct {
	Contract *AugurFoundry // Generic contract binding to access the raw methods on
}

// AugurFoundryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AugurFoundryCallerRaw struct {
	Contract *AugurFoundryCaller // Generic read-only contract binding to access the raw methods on
}

// AugurFoundryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AugurFoundryTransactorRaw struct {
	Contract *AugurFoundryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAugurFoundry creates a new instance of AugurFoundry, bound to a specific deployed contract.
func NewAugurFoundry(address common.Address, backend bind.ContractBackend) (*AugurFoundry, error) {
	contract, err := bindAugurFoundry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AugurFoundry{AugurFoundryCaller: AugurFoundryCaller{contract: contract}, AugurFoundryTransactor: AugurFoundryTransactor{contract: contract}, AugurFoundryFilterer: AugurFoundryFilterer{contract: contract}}, nil
}

// NewAugurFoundryCaller creates a new read-only instance of AugurFoundry, bound to a specific deployed contract.
func NewAugurFoundryCaller(address common.Address, caller bind.ContractCaller) (*AugurFoundryCaller, error) {
	contract, err := bindAugurFoundry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AugurFoundryCaller{contract: contract}, nil
}

// NewAugurFoundryTransactor creates a new write-only instance of AugurFoundry, bound to a specific deployed contract.
func NewAugurFoundryTransactor(address common.Address, transactor bind.ContractTransactor) (*AugurFoundryTransactor, error) {
	contract, err := bindAugurFoundry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AugurFoundryTransactor{contract: contract}, nil
}

// NewAugurFoundryFilterer creates a new log filterer instance of AugurFoundry, bound to a specific deployed contract.
func NewAugurFoundryFilterer(address common.Address, filterer bind.ContractFilterer) (*AugurFoundryFilterer, error) {
	contract, err := bindAugurFoundry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AugurFoundryFilterer{contract: contract}, nil
}

// bindAugurFoundry binds a generic wrapper to an already deployed contract.
func bindAugurFoundry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AugurFoundryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AugurFoundry *AugurFoundryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AugurFoundry.Contract.AugurFoundryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AugurFoundry *AugurFoundryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AugurFoundry.Contract.AugurFoundryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AugurFoundry *AugurFoundryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AugurFoundry.Contract.AugurFoundryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AugurFoundry *AugurFoundryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AugurFoundry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AugurFoundry *AugurFoundryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AugurFoundry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AugurFoundry *AugurFoundryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AugurFoundry.Contract.contract.Transact(opts, method, params...)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_AugurFoundry *AugurFoundryCaller) Cash(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AugurFoundry.contract.Call(opts, &out, "cash")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_AugurFoundry *AugurFoundrySession) Cash() (common.Address, error) {
	return _AugurFoundry.Contract.Cash(&_AugurFoundry.CallOpts)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_AugurFoundry *AugurFoundryCallerSession) Cash() (common.Address, error) {
	return _AugurFoundry.Contract.Cash(&_AugurFoundry.CallOpts)
}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() view returns(address)
func (_AugurFoundry *AugurFoundryCaller) ShareToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AugurFoundry.contract.Call(opts, &out, "shareToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() view returns(address)
func (_AugurFoundry *AugurFoundrySession) ShareToken() (common.Address, error) {
	return _AugurFoundry.Contract.ShareToken(&_AugurFoundry.CallOpts)
}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() view returns(address)
func (_AugurFoundry *AugurFoundryCallerSession) ShareToken() (common.Address, error) {
	return _AugurFoundry.Contract.ShareToken(&_AugurFoundry.CallOpts)
}

// Wrappers is a free data retrieval call binding the contract method 0xcba45df2.
//
// Solidity: function wrappers(uint256 ) view returns(address)
func (_AugurFoundry *AugurFoundryCaller) Wrappers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AugurFoundry.contract.Call(opts, &out, "wrappers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wrappers is a free data retrieval call binding the contract method 0xcba45df2.
//
// Solidity: function wrappers(uint256 ) view returns(address)
func (_AugurFoundry *AugurFoundrySession) Wrappers(arg0 *big.Int) (common.Address, error) {
	return _AugurFoundry.Contract.Wrappers(&_AugurFoundry.CallOpts, arg0)
}

// Wrappers is a free data retrieval call binding the contract method 0xcba45df2.
//
// Solidity: function wrappers(uint256 ) view returns(address)
func (_AugurFoundry *AugurFoundryCallerSession) Wrappers(arg0 *big.Int) (common.Address, error) {
	return _AugurFoundry.Contract.Wrappers(&_AugurFoundry.CallOpts, arg0)
}

// NewERC20Wrapper is a paid mutator transaction binding the contract method 0x21c9fedb.
//
// Solidity: function newERC20Wrapper(uint256 _tokenId, string _name, string _symbol, uint8 _decimals) returns()
func (_AugurFoundry *AugurFoundryTransactor) NewERC20Wrapper(opts *bind.TransactOpts, _tokenId *big.Int, _name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _AugurFoundry.contract.Transact(opts, "newERC20Wrapper", _tokenId, _name, _symbol, _decimals)
}

// NewERC20Wrapper is a paid mutator transaction binding the contract method 0x21c9fedb.
//
// Solidity: function newERC20Wrapper(uint256 _tokenId, string _name, string _symbol, uint8 _decimals) returns()
func (_AugurFoundry *AugurFoundrySession) NewERC20Wrapper(_tokenId *big.Int, _name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _AugurFoundry.Contract.NewERC20Wrapper(&_AugurFoundry.TransactOpts, _tokenId, _name, _symbol, _decimals)
}

// NewERC20Wrapper is a paid mutator transaction binding the contract method 0x21c9fedb.
//
// Solidity: function newERC20Wrapper(uint256 _tokenId, string _name, string _symbol, uint8 _decimals) returns()
func (_AugurFoundry *AugurFoundryTransactorSession) NewERC20Wrapper(_tokenId *big.Int, _name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _AugurFoundry.Contract.NewERC20Wrapper(&_AugurFoundry.TransactOpts, _tokenId, _name, _symbol, _decimals)
}

// NewERC20Wrappers is a paid mutator transaction binding the contract method 0x74ca6745.
//
// Solidity: function newERC20Wrappers(uint256[] _tokenIds, string[] _names, string[] _symbols, uint8[] _decimals) returns()
func (_AugurFoundry *AugurFoundryTransactor) NewERC20Wrappers(opts *bind.TransactOpts, _tokenIds []*big.Int, _names []string, _symbols []string, _decimals []uint8) (*types.Transaction, error) {
	return _AugurFoundry.contract.Transact(opts, "newERC20Wrappers", _tokenIds, _names, _symbols, _decimals)
}

// NewERC20Wrappers is a paid mutator transaction binding the contract method 0x74ca6745.
//
// Solidity: function newERC20Wrappers(uint256[] _tokenIds, string[] _names, string[] _symbols, uint8[] _decimals) returns()
func (_AugurFoundry *AugurFoundrySession) NewERC20Wrappers(_tokenIds []*big.Int, _names []string, _symbols []string, _decimals []uint8) (*types.Transaction, error) {
	return _AugurFoundry.Contract.NewERC20Wrappers(&_AugurFoundry.TransactOpts, _tokenIds, _names, _symbols, _decimals)
}

// NewERC20Wrappers is a paid mutator transaction binding the contract method 0x74ca6745.
//
// Solidity: function newERC20Wrappers(uint256[] _tokenIds, string[] _names, string[] _symbols, uint8[] _decimals) returns()
func (_AugurFoundry *AugurFoundryTransactorSession) NewERC20Wrappers(_tokenIds []*big.Int, _names []string, _symbols []string, _decimals []uint8) (*types.Transaction, error) {
	return _AugurFoundry.Contract.NewERC20Wrappers(&_AugurFoundry.TransactOpts, _tokenIds, _names, _symbols, _decimals)
}

// UnWrapMultipleTokens is a paid mutator transaction binding the contract method 0x62a3b176.
//
// Solidity: function unWrapMultipleTokens(uint256[] _tokenIds, uint256[] _amounts) returns()
func (_AugurFoundry *AugurFoundryTransactor) UnWrapMultipleTokens(opts *bind.TransactOpts, _tokenIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _AugurFoundry.contract.Transact(opts, "unWrapMultipleTokens", _tokenIds, _amounts)
}

// UnWrapMultipleTokens is a paid mutator transaction binding the contract method 0x62a3b176.
//
// Solidity: function unWrapMultipleTokens(uint256[] _tokenIds, uint256[] _amounts) returns()
func (_AugurFoundry *AugurFoundrySession) UnWrapMultipleTokens(_tokenIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _AugurFoundry.Contract.UnWrapMultipleTokens(&_AugurFoundry.TransactOpts, _tokenIds, _amounts)
}

// UnWrapMultipleTokens is a paid mutator transaction binding the contract method 0x62a3b176.
//
// Solidity: function unWrapMultipleTokens(uint256[] _tokenIds, uint256[] _amounts) returns()
func (_AugurFoundry *AugurFoundryTransactorSession) UnWrapMultipleTokens(_tokenIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _AugurFoundry.Contract.UnWrapMultipleTokens(&_AugurFoundry.TransactOpts, _tokenIds, _amounts)
}

// UnWrapTokens is a paid mutator transaction binding the contract method 0x6398b5e5.
//
// Solidity: function unWrapTokens(uint256 _tokenId, uint256 _amount) returns()
func (_AugurFoundry *AugurFoundryTransactor) UnWrapTokens(opts *bind.TransactOpts, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _AugurFoundry.contract.Transact(opts, "unWrapTokens", _tokenId, _amount)
}

// UnWrapTokens is a paid mutator transaction binding the contract method 0x6398b5e5.
//
// Solidity: function unWrapTokens(uint256 _tokenId, uint256 _amount) returns()
func (_AugurFoundry *AugurFoundrySession) UnWrapTokens(_tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _AugurFoundry.Contract.UnWrapTokens(&_AugurFoundry.TransactOpts, _tokenId, _amount)
}

// UnWrapTokens is a paid mutator transaction binding the contract method 0x6398b5e5.
//
// Solidity: function unWrapTokens(uint256 _tokenId, uint256 _amount) returns()
func (_AugurFoundry *AugurFoundryTransactorSession) UnWrapTokens(_tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _AugurFoundry.Contract.UnWrapTokens(&_AugurFoundry.TransactOpts, _tokenId, _amount)
}

// WrapMultipleTokens is a paid mutator transaction binding the contract method 0x591193a5.
//
// Solidity: function wrapMultipleTokens(uint256[] _tokenIds, address _account, uint256[] _amounts) returns()
func (_AugurFoundry *AugurFoundryTransactor) WrapMultipleTokens(opts *bind.TransactOpts, _tokenIds []*big.Int, _account common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _AugurFoundry.contract.Transact(opts, "wrapMultipleTokens", _tokenIds, _account, _amounts)
}

// WrapMultipleTokens is a paid mutator transaction binding the contract method 0x591193a5.
//
// Solidity: function wrapMultipleTokens(uint256[] _tokenIds, address _account, uint256[] _amounts) returns()
func (_AugurFoundry *AugurFoundrySession) WrapMultipleTokens(_tokenIds []*big.Int, _account common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _AugurFoundry.Contract.WrapMultipleTokens(&_AugurFoundry.TransactOpts, _tokenIds, _account, _amounts)
}

// WrapMultipleTokens is a paid mutator transaction binding the contract method 0x591193a5.
//
// Solidity: function wrapMultipleTokens(uint256[] _tokenIds, address _account, uint256[] _amounts) returns()
func (_AugurFoundry *AugurFoundryTransactorSession) WrapMultipleTokens(_tokenIds []*big.Int, _account common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _AugurFoundry.Contract.WrapMultipleTokens(&_AugurFoundry.TransactOpts, _tokenIds, _account, _amounts)
}

// WrapTokens is a paid mutator transaction binding the contract method 0x062e6b87.
//
// Solidity: function wrapTokens(uint256 _tokenId, address _account, uint256 _amount) returns()
func (_AugurFoundry *AugurFoundryTransactor) WrapTokens(opts *bind.TransactOpts, _tokenId *big.Int, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AugurFoundry.contract.Transact(opts, "wrapTokens", _tokenId, _account, _amount)
}

// WrapTokens is a paid mutator transaction binding the contract method 0x062e6b87.
//
// Solidity: function wrapTokens(uint256 _tokenId, address _account, uint256 _amount) returns()
func (_AugurFoundry *AugurFoundrySession) WrapTokens(_tokenId *big.Int, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AugurFoundry.Contract.WrapTokens(&_AugurFoundry.TransactOpts, _tokenId, _account, _amount)
}

// WrapTokens is a paid mutator transaction binding the contract method 0x062e6b87.
//
// Solidity: function wrapTokens(uint256 _tokenId, address _account, uint256 _amount) returns()
func (_AugurFoundry *AugurFoundryTransactorSession) WrapTokens(_tokenId *big.Int, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AugurFoundry.Contract.WrapTokens(&_AugurFoundry.TransactOpts, _tokenId, _account, _amount)
}

// AugurFoundryWrapperCreatedIterator is returned from FilterWrapperCreated and is used to iterate over the raw logs and unpacked data for WrapperCreated events raised by the AugurFoundry contract.
type AugurFoundryWrapperCreatedIterator struct {
	Event *AugurFoundryWrapperCreated // Event containing the contract specifics and raw log

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
func (it *AugurFoundryWrapperCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurFoundryWrapperCreated)
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
		it.Event = new(AugurFoundryWrapperCreated)
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
func (it *AugurFoundryWrapperCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurFoundryWrapperCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurFoundryWrapperCreated represents a WrapperCreated event raised by the AugurFoundry contract.
type AugurFoundryWrapperCreated struct {
	TokenId      *big.Int
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWrapperCreated is a free log retrieval operation binding the contract event 0x7dcd5c8018a7a1acb9c51e302d6e79b9d8cc5020a31ba21d009b001c7b236fac.
//
// Solidity: event WrapperCreated(uint256 indexed tokenId, address tokenAddress)
func (_AugurFoundry *AugurFoundryFilterer) FilterWrapperCreated(opts *bind.FilterOpts, tokenId []*big.Int) (*AugurFoundryWrapperCreatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AugurFoundry.contract.FilterLogs(opts, "WrapperCreated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AugurFoundryWrapperCreatedIterator{contract: _AugurFoundry.contract, event: "WrapperCreated", logs: logs, sub: sub}, nil
}

// WatchWrapperCreated is a free log subscription operation binding the contract event 0x7dcd5c8018a7a1acb9c51e302d6e79b9d8cc5020a31ba21d009b001c7b236fac.
//
// Solidity: event WrapperCreated(uint256 indexed tokenId, address tokenAddress)
func (_AugurFoundry *AugurFoundryFilterer) WatchWrapperCreated(opts *bind.WatchOpts, sink chan<- *AugurFoundryWrapperCreated, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AugurFoundry.contract.WatchLogs(opts, "WrapperCreated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurFoundryWrapperCreated)
				if err := _AugurFoundry.contract.UnpackLog(event, "WrapperCreated", log); err != nil {
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

// ParseWrapperCreated is a log parse operation binding the contract event 0x7dcd5c8018a7a1acb9c51e302d6e79b9d8cc5020a31ba21d009b001c7b236fac.
//
// Solidity: event WrapperCreated(uint256 indexed tokenId, address tokenAddress)
func (_AugurFoundry *AugurFoundryFilterer) ParseWrapperCreated(log types.Log) (*AugurFoundryWrapperCreated, error) {
	event := new(AugurFoundryWrapperCreated)
	if err := _AugurFoundry.contract.UnpackLog(event, "WrapperCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
