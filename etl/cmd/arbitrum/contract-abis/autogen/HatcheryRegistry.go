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

// HatcheryRegistryABI is the input ABI used to generate the binding from.
const HatcheryRegistryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"contractIERC20DynamicSymbol\",\"name\":\"_reputationToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"shareTokenFactory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feePot\",\"type\":\"address\"}],\"name\":\"NewHatchery\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_collateral\",\"type\":\"address\"}],\"name\":\"createHatchery\",\"outputs\":[{\"internalType\":\"contractTurboHatchery\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getHatchery\",\"outputs\":[{\"internalType\":\"contractTurboHatchery\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hatcheries\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reputationToken\",\"outputs\":[{\"internalType\":\"contractIERC20DynamicSymbol\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// HatcheryRegistry is an auto generated Go binding around an Ethereum contract.
type HatcheryRegistry struct {
	HatcheryRegistryCaller     // Read-only binding to the contract
	HatcheryRegistryTransactor // Write-only binding to the contract
	HatcheryRegistryFilterer   // Log filterer for contract events
}

// HatcheryRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type HatcheryRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HatcheryRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HatcheryRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HatcheryRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HatcheryRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HatcheryRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HatcheryRegistrySession struct {
	Contract     *HatcheryRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HatcheryRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HatcheryRegistryCallerSession struct {
	Contract *HatcheryRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// HatcheryRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HatcheryRegistryTransactorSession struct {
	Contract     *HatcheryRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// HatcheryRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type HatcheryRegistryRaw struct {
	Contract *HatcheryRegistry // Generic contract binding to access the raw methods on
}

// HatcheryRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HatcheryRegistryCallerRaw struct {
	Contract *HatcheryRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// HatcheryRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HatcheryRegistryTransactorRaw struct {
	Contract *HatcheryRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHatcheryRegistry creates a new instance of HatcheryRegistry, bound to a specific deployed contract.
func NewHatcheryRegistry(address common.Address, backend bind.ContractBackend) (*HatcheryRegistry, error) {
	contract, err := bindHatcheryRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HatcheryRegistry{HatcheryRegistryCaller: HatcheryRegistryCaller{contract: contract}, HatcheryRegistryTransactor: HatcheryRegistryTransactor{contract: contract}, HatcheryRegistryFilterer: HatcheryRegistryFilterer{contract: contract}}, nil
}

// NewHatcheryRegistryCaller creates a new read-only instance of HatcheryRegistry, bound to a specific deployed contract.
func NewHatcheryRegistryCaller(address common.Address, caller bind.ContractCaller) (*HatcheryRegistryCaller, error) {
	contract, err := bindHatcheryRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HatcheryRegistryCaller{contract: contract}, nil
}

// NewHatcheryRegistryTransactor creates a new write-only instance of HatcheryRegistry, bound to a specific deployed contract.
func NewHatcheryRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*HatcheryRegistryTransactor, error) {
	contract, err := bindHatcheryRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HatcheryRegistryTransactor{contract: contract}, nil
}

// NewHatcheryRegistryFilterer creates a new log filterer instance of HatcheryRegistry, bound to a specific deployed contract.
func NewHatcheryRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*HatcheryRegistryFilterer, error) {
	contract, err := bindHatcheryRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HatcheryRegistryFilterer{contract: contract}, nil
}

// bindHatcheryRegistry binds a generic wrapper to an already deployed contract.
func bindHatcheryRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HatcheryRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HatcheryRegistry *HatcheryRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HatcheryRegistry.Contract.HatcheryRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HatcheryRegistry *HatcheryRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HatcheryRegistry.Contract.HatcheryRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HatcheryRegistry *HatcheryRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HatcheryRegistry.Contract.HatcheryRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HatcheryRegistry *HatcheryRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HatcheryRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HatcheryRegistry *HatcheryRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HatcheryRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HatcheryRegistry *HatcheryRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HatcheryRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetHatchery is a free data retrieval call binding the contract method 0x1ee5b776.
//
// Solidity: function getHatchery(address ) view returns(address)
func (_HatcheryRegistry *HatcheryRegistryCaller) GetHatchery(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HatcheryRegistry.contract.Call(opts, out, "getHatchery", arg0)
	return *ret0, err
}

// GetHatchery is a free data retrieval call binding the contract method 0x1ee5b776.
//
// Solidity: function getHatchery(address ) view returns(address)
func (_HatcheryRegistry *HatcheryRegistrySession) GetHatchery(arg0 common.Address) (common.Address, error) {
	return _HatcheryRegistry.Contract.GetHatchery(&_HatcheryRegistry.CallOpts, arg0)
}

// GetHatchery is a free data retrieval call binding the contract method 0x1ee5b776.
//
// Solidity: function getHatchery(address ) view returns(address)
func (_HatcheryRegistry *HatcheryRegistryCallerSession) GetHatchery(arg0 common.Address) (common.Address, error) {
	return _HatcheryRegistry.Contract.GetHatchery(&_HatcheryRegistry.CallOpts, arg0)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_HatcheryRegistry *HatcheryRegistryCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HatcheryRegistry.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_HatcheryRegistry *HatcheryRegistrySession) GetOwner() (common.Address, error) {
	return _HatcheryRegistry.Contract.GetOwner(&_HatcheryRegistry.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_HatcheryRegistry *HatcheryRegistryCallerSession) GetOwner() (common.Address, error) {
	return _HatcheryRegistry.Contract.GetOwner(&_HatcheryRegistry.CallOpts)
}

// Hatcheries is a free data retrieval call binding the contract method 0xcbec31aa.
//
// Solidity: function hatcheries(uint256 ) view returns(address)
func (_HatcheryRegistry *HatcheryRegistryCaller) Hatcheries(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HatcheryRegistry.contract.Call(opts, out, "hatcheries", arg0)
	return *ret0, err
}

// Hatcheries is a free data retrieval call binding the contract method 0xcbec31aa.
//
// Solidity: function hatcheries(uint256 ) view returns(address)
func (_HatcheryRegistry *HatcheryRegistrySession) Hatcheries(arg0 *big.Int) (common.Address, error) {
	return _HatcheryRegistry.Contract.Hatcheries(&_HatcheryRegistry.CallOpts, arg0)
}

// Hatcheries is a free data retrieval call binding the contract method 0xcbec31aa.
//
// Solidity: function hatcheries(uint256 ) view returns(address)
func (_HatcheryRegistry *HatcheryRegistryCallerSession) Hatcheries(arg0 *big.Int) (common.Address, error) {
	return _HatcheryRegistry.Contract.Hatcheries(&_HatcheryRegistry.CallOpts, arg0)
}

// ReputationToken is a free data retrieval call binding the contract method 0x275bbe9b.
//
// Solidity: function reputationToken() view returns(address)
func (_HatcheryRegistry *HatcheryRegistryCaller) ReputationToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HatcheryRegistry.contract.Call(opts, out, "reputationToken")
	return *ret0, err
}

// ReputationToken is a free data retrieval call binding the contract method 0x275bbe9b.
//
// Solidity: function reputationToken() view returns(address)
func (_HatcheryRegistry *HatcheryRegistrySession) ReputationToken() (common.Address, error) {
	return _HatcheryRegistry.Contract.ReputationToken(&_HatcheryRegistry.CallOpts)
}

// ReputationToken is a free data retrieval call binding the contract method 0x275bbe9b.
//
// Solidity: function reputationToken() view returns(address)
func (_HatcheryRegistry *HatcheryRegistryCallerSession) ReputationToken() (common.Address, error) {
	return _HatcheryRegistry.Contract.ReputationToken(&_HatcheryRegistry.CallOpts)
}

// CreateHatchery is a paid mutator transaction binding the contract method 0xe1554256.
//
// Solidity: function createHatchery(address _collateral) returns(address)
func (_HatcheryRegistry *HatcheryRegistryTransactor) CreateHatchery(opts *bind.TransactOpts, _collateral common.Address) (*types.Transaction, error) {
	return _HatcheryRegistry.contract.Transact(opts, "createHatchery", _collateral)
}

// CreateHatchery is a paid mutator transaction binding the contract method 0xe1554256.
//
// Solidity: function createHatchery(address _collateral) returns(address)
func (_HatcheryRegistry *HatcheryRegistrySession) CreateHatchery(_collateral common.Address) (*types.Transaction, error) {
	return _HatcheryRegistry.Contract.CreateHatchery(&_HatcheryRegistry.TransactOpts, _collateral)
}

// CreateHatchery is a paid mutator transaction binding the contract method 0xe1554256.
//
// Solidity: function createHatchery(address _collateral) returns(address)
func (_HatcheryRegistry *HatcheryRegistryTransactorSession) CreateHatchery(_collateral common.Address) (*types.Transaction, error) {
	return _HatcheryRegistry.Contract.CreateHatchery(&_HatcheryRegistry.TransactOpts, _collateral)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_HatcheryRegistry *HatcheryRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _HatcheryRegistry.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_HatcheryRegistry *HatcheryRegistrySession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _HatcheryRegistry.Contract.TransferOwnership(&_HatcheryRegistry.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns(bool)
func (_HatcheryRegistry *HatcheryRegistryTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _HatcheryRegistry.Contract.TransferOwnership(&_HatcheryRegistry.TransactOpts, _newOwner)
}

// HatcheryRegistryNewHatcheryIterator is returned from FilterNewHatchery and is used to iterate over the raw logs and unpacked data for NewHatchery events raised by the HatcheryRegistry contract.
type HatcheryRegistryNewHatcheryIterator struct {
	Event *HatcheryRegistryNewHatchery // Event containing the contract specifics and raw log

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
func (it *HatcheryRegistryNewHatcheryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HatcheryRegistryNewHatchery)
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
		it.Event = new(HatcheryRegistryNewHatchery)
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
func (it *HatcheryRegistryNewHatcheryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HatcheryRegistryNewHatcheryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HatcheryRegistryNewHatchery represents a NewHatchery event raised by the HatcheryRegistry contract.
type HatcheryRegistryNewHatchery struct {
	Id                common.Address
	Collateral        common.Address
	ShareTokenFactory common.Address
	FeePot            common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewHatchery is a free log retrieval operation binding the contract event 0x08afdadd49d632c11dbde177a7ab47701b5adaac8f633beedd892c8da8d4393f.
//
// Solidity: event NewHatchery(address id, address indexed collateral, address shareTokenFactory, address feePot)
func (_HatcheryRegistry *HatcheryRegistryFilterer) FilterNewHatchery(opts *bind.FilterOpts, collateral []common.Address) (*HatcheryRegistryNewHatcheryIterator, error) {

	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _HatcheryRegistry.contract.FilterLogs(opts, "NewHatchery", collateralRule)
	if err != nil {
		return nil, err
	}
	return &HatcheryRegistryNewHatcheryIterator{contract: _HatcheryRegistry.contract, event: "NewHatchery", logs: logs, sub: sub}, nil
}

// WatchNewHatchery is a free log subscription operation binding the contract event 0x08afdadd49d632c11dbde177a7ab47701b5adaac8f633beedd892c8da8d4393f.
//
// Solidity: event NewHatchery(address id, address indexed collateral, address shareTokenFactory, address feePot)
func (_HatcheryRegistry *HatcheryRegistryFilterer) WatchNewHatchery(opts *bind.WatchOpts, sink chan<- *HatcheryRegistryNewHatchery, collateral []common.Address) (event.Subscription, error) {

	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _HatcheryRegistry.contract.WatchLogs(opts, "NewHatchery", collateralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HatcheryRegistryNewHatchery)
				if err := _HatcheryRegistry.contract.UnpackLog(event, "NewHatchery", log); err != nil {
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

// ParseNewHatchery is a log parse operation binding the contract event 0x08afdadd49d632c11dbde177a7ab47701b5adaac8f633beedd892c8da8d4393f.
//
// Solidity: event NewHatchery(address id, address indexed collateral, address shareTokenFactory, address feePot)
func (_HatcheryRegistry *HatcheryRegistryFilterer) ParseNewHatchery(log types.Log) (*HatcheryRegistryNewHatchery, error) {
	event := new(HatcheryRegistryNewHatchery)
	if err := _HatcheryRegistry.contract.UnpackLog(event, "NewHatchery", log); err != nil {
		return nil, err
	}
	return event, nil
}
