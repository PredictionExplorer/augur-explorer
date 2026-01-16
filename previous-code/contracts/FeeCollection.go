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

// FeeCollectionMetaData contains all meta data concerning the FeeCollection contract.
var FeeCollectionMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collectedBase\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collectedBond\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingBase\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingBond\",\"type\":\"uint256\"}],\"name\":\"FeeCollection\",\"type\":\"event\"}]",
}

// FeeCollectionABI is the input ABI used to generate the binding from.
// Deprecated: Use FeeCollectionMetaData.ABI instead.
var FeeCollectionABI = FeeCollectionMetaData.ABI

// FeeCollection is an auto generated Go binding around an Ethereum contract.
type FeeCollection struct {
	FeeCollectionCaller     // Read-only binding to the contract
	FeeCollectionTransactor // Write-only binding to the contract
	FeeCollectionFilterer   // Log filterer for contract events
}

// FeeCollectionCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeeCollectionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeCollectionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeCollectionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeCollectionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeCollectionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeCollectionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeCollectionSession struct {
	Contract     *FeeCollection    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeCollectionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeCollectionCallerSession struct {
	Contract *FeeCollectionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// FeeCollectionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeCollectionTransactorSession struct {
	Contract     *FeeCollectionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// FeeCollectionRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeeCollectionRaw struct {
	Contract *FeeCollection // Generic contract binding to access the raw methods on
}

// FeeCollectionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeCollectionCallerRaw struct {
	Contract *FeeCollectionCaller // Generic read-only contract binding to access the raw methods on
}

// FeeCollectionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeCollectionTransactorRaw struct {
	Contract *FeeCollectionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeCollection creates a new instance of FeeCollection, bound to a specific deployed contract.
func NewFeeCollection(address common.Address, backend bind.ContractBackend) (*FeeCollection, error) {
	contract, err := bindFeeCollection(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeCollection{FeeCollectionCaller: FeeCollectionCaller{contract: contract}, FeeCollectionTransactor: FeeCollectionTransactor{contract: contract}, FeeCollectionFilterer: FeeCollectionFilterer{contract: contract}}, nil
}

// NewFeeCollectionCaller creates a new read-only instance of FeeCollection, bound to a specific deployed contract.
func NewFeeCollectionCaller(address common.Address, caller bind.ContractCaller) (*FeeCollectionCaller, error) {
	contract, err := bindFeeCollection(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeCollectionCaller{contract: contract}, nil
}

// NewFeeCollectionTransactor creates a new write-only instance of FeeCollection, bound to a specific deployed contract.
func NewFeeCollectionTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeCollectionTransactor, error) {
	contract, err := bindFeeCollection(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeCollectionTransactor{contract: contract}, nil
}

// NewFeeCollectionFilterer creates a new log filterer instance of FeeCollection, bound to a specific deployed contract.
func NewFeeCollectionFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeCollectionFilterer, error) {
	contract, err := bindFeeCollection(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeCollectionFilterer{contract: contract}, nil
}

// bindFeeCollection binds a generic wrapper to an already deployed contract.
func bindFeeCollection(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeCollectionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeCollection *FeeCollectionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeCollection.Contract.FeeCollectionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeCollection *FeeCollectionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeCollection.Contract.FeeCollectionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeCollection *FeeCollectionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeCollection.Contract.FeeCollectionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeCollection *FeeCollectionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeCollection.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeCollection *FeeCollectionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeCollection.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeCollection *FeeCollectionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeCollection.Contract.contract.Transact(opts, method, params...)
}

// FeeCollectionFeeCollectionIterator is returned from FilterFeeCollection and is used to iterate over the raw logs and unpacked data for FeeCollection events raised by the FeeCollection contract.
type FeeCollectionFeeCollectionIterator struct {
	Event *FeeCollectionFeeCollection // Event containing the contract specifics and raw log

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
func (it *FeeCollectionFeeCollectionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeCollectionFeeCollection)
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
		it.Event = new(FeeCollectionFeeCollection)
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
func (it *FeeCollectionFeeCollectionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeCollectionFeeCollectionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeCollectionFeeCollection represents a FeeCollection event raised by the FeeCollection contract.
type FeeCollectionFeeCollection struct {
	CollectedBase *big.Int
	CollectedBond *big.Int
	RemainingBase *big.Int
	RemainingBond *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeeCollection is a free log retrieval operation binding the contract event 0x9f878c349b0fc751f12168fdf539db8c1848b81c0751432f28626da5aa7024ee.
//
// Solidity: event FeeCollection(uint256 collectedBase, uint256 collectedBond, uint256 remainingBase, uint256 remainingBond)
func (_FeeCollection *FeeCollectionFilterer) FilterFeeCollection(opts *bind.FilterOpts) (*FeeCollectionFeeCollectionIterator, error) {

	logs, sub, err := _FeeCollection.contract.FilterLogs(opts, "FeeCollection")
	if err != nil {
		return nil, err
	}
	return &FeeCollectionFeeCollectionIterator{contract: _FeeCollection.contract, event: "FeeCollection", logs: logs, sub: sub}, nil
}

// WatchFeeCollection is a free log subscription operation binding the contract event 0x9f878c349b0fc751f12168fdf539db8c1848b81c0751432f28626da5aa7024ee.
//
// Solidity: event FeeCollection(uint256 collectedBase, uint256 collectedBond, uint256 remainingBase, uint256 remainingBond)
func (_FeeCollection *FeeCollectionFilterer) WatchFeeCollection(opts *bind.WatchOpts, sink chan<- *FeeCollectionFeeCollection) (event.Subscription, error) {

	logs, sub, err := _FeeCollection.contract.WatchLogs(opts, "FeeCollection")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeCollectionFeeCollection)
				if err := _FeeCollection.contract.UnpackLog(event, "FeeCollection", log); err != nil {
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

// ParseFeeCollection is a log parse operation binding the contract event 0x9f878c349b0fc751f12168fdf539db8c1848b81c0751432f28626da5aa7024ee.
//
// Solidity: event FeeCollection(uint256 collectedBase, uint256 collectedBond, uint256 remainingBase, uint256 remainingBond)
func (_FeeCollection *FeeCollectionFilterer) ParseFeeCollection(log types.Log) (*FeeCollectionFeeCollection, error) {
	event := new(FeeCollectionFeeCollection)
	if err := _FeeCollection.contract.UnpackLog(event, "FeeCollection", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
