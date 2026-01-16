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

// BalancerV2SwapFeePercentageChangedMetaData contains all meta data concerning the BalancerV2SwapFeePercentageChanged contract.
var BalancerV2SwapFeePercentageChangedMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"}],\"name\":\"SwapFeePercentageChanged\",\"type\":\"event\"}]",
}

// BalancerV2SwapFeePercentageChangedABI is the input ABI used to generate the binding from.
// Deprecated: Use BalancerV2SwapFeePercentageChangedMetaData.ABI instead.
var BalancerV2SwapFeePercentageChangedABI = BalancerV2SwapFeePercentageChangedMetaData.ABI

// BalancerV2SwapFeePercentageChanged is an auto generated Go binding around an Ethereum contract.
type BalancerV2SwapFeePercentageChanged struct {
	BalancerV2SwapFeePercentageChangedCaller     // Read-only binding to the contract
	BalancerV2SwapFeePercentageChangedTransactor // Write-only binding to the contract
	BalancerV2SwapFeePercentageChangedFilterer   // Log filterer for contract events
}

// BalancerV2SwapFeePercentageChangedCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalancerV2SwapFeePercentageChangedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerV2SwapFeePercentageChangedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalancerV2SwapFeePercentageChangedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerV2SwapFeePercentageChangedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalancerV2SwapFeePercentageChangedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerV2SwapFeePercentageChangedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalancerV2SwapFeePercentageChangedSession struct {
	Contract     *BalancerV2SwapFeePercentageChanged // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                       // Call options to use throughout this session
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// BalancerV2SwapFeePercentageChangedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalancerV2SwapFeePercentageChangedCallerSession struct {
	Contract *BalancerV2SwapFeePercentageChangedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                             // Call options to use throughout this session
}

// BalancerV2SwapFeePercentageChangedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalancerV2SwapFeePercentageChangedTransactorSession struct {
	Contract     *BalancerV2SwapFeePercentageChangedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                             // Transaction auth options to use throughout this session
}

// BalancerV2SwapFeePercentageChangedRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalancerV2SwapFeePercentageChangedRaw struct {
	Contract *BalancerV2SwapFeePercentageChanged // Generic contract binding to access the raw methods on
}

// BalancerV2SwapFeePercentageChangedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalancerV2SwapFeePercentageChangedCallerRaw struct {
	Contract *BalancerV2SwapFeePercentageChangedCaller // Generic read-only contract binding to access the raw methods on
}

// BalancerV2SwapFeePercentageChangedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalancerV2SwapFeePercentageChangedTransactorRaw struct {
	Contract *BalancerV2SwapFeePercentageChangedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalancerV2SwapFeePercentageChanged creates a new instance of BalancerV2SwapFeePercentageChanged, bound to a specific deployed contract.
func NewBalancerV2SwapFeePercentageChanged(address common.Address, backend bind.ContractBackend) (*BalancerV2SwapFeePercentageChanged, error) {
	contract, err := bindBalancerV2SwapFeePercentageChanged(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalancerV2SwapFeePercentageChanged{BalancerV2SwapFeePercentageChangedCaller: BalancerV2SwapFeePercentageChangedCaller{contract: contract}, BalancerV2SwapFeePercentageChangedTransactor: BalancerV2SwapFeePercentageChangedTransactor{contract: contract}, BalancerV2SwapFeePercentageChangedFilterer: BalancerV2SwapFeePercentageChangedFilterer{contract: contract}}, nil
}

// NewBalancerV2SwapFeePercentageChangedCaller creates a new read-only instance of BalancerV2SwapFeePercentageChanged, bound to a specific deployed contract.
func NewBalancerV2SwapFeePercentageChangedCaller(address common.Address, caller bind.ContractCaller) (*BalancerV2SwapFeePercentageChangedCaller, error) {
	contract, err := bindBalancerV2SwapFeePercentageChanged(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerV2SwapFeePercentageChangedCaller{contract: contract}, nil
}

// NewBalancerV2SwapFeePercentageChangedTransactor creates a new write-only instance of BalancerV2SwapFeePercentageChanged, bound to a specific deployed contract.
func NewBalancerV2SwapFeePercentageChangedTransactor(address common.Address, transactor bind.ContractTransactor) (*BalancerV2SwapFeePercentageChangedTransactor, error) {
	contract, err := bindBalancerV2SwapFeePercentageChanged(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerV2SwapFeePercentageChangedTransactor{contract: contract}, nil
}

// NewBalancerV2SwapFeePercentageChangedFilterer creates a new log filterer instance of BalancerV2SwapFeePercentageChanged, bound to a specific deployed contract.
func NewBalancerV2SwapFeePercentageChangedFilterer(address common.Address, filterer bind.ContractFilterer) (*BalancerV2SwapFeePercentageChangedFilterer, error) {
	contract, err := bindBalancerV2SwapFeePercentageChanged(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalancerV2SwapFeePercentageChangedFilterer{contract: contract}, nil
}

// bindBalancerV2SwapFeePercentageChanged binds a generic wrapper to an already deployed contract.
func bindBalancerV2SwapFeePercentageChanged(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalancerV2SwapFeePercentageChangedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerV2SwapFeePercentageChanged *BalancerV2SwapFeePercentageChangedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalancerV2SwapFeePercentageChanged.Contract.BalancerV2SwapFeePercentageChangedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerV2SwapFeePercentageChanged *BalancerV2SwapFeePercentageChangedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerV2SwapFeePercentageChanged.Contract.BalancerV2SwapFeePercentageChangedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerV2SwapFeePercentageChanged *BalancerV2SwapFeePercentageChangedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerV2SwapFeePercentageChanged.Contract.BalancerV2SwapFeePercentageChangedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerV2SwapFeePercentageChanged *BalancerV2SwapFeePercentageChangedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalancerV2SwapFeePercentageChanged.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerV2SwapFeePercentageChanged *BalancerV2SwapFeePercentageChangedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerV2SwapFeePercentageChanged.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerV2SwapFeePercentageChanged *BalancerV2SwapFeePercentageChangedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerV2SwapFeePercentageChanged.Contract.contract.Transact(opts, method, params...)
}

// BalancerV2SwapFeePercentageChangedSwapFeePercentageChangedIterator is returned from FilterSwapFeePercentageChanged and is used to iterate over the raw logs and unpacked data for SwapFeePercentageChanged events raised by the BalancerV2SwapFeePercentageChanged contract.
type BalancerV2SwapFeePercentageChangedSwapFeePercentageChangedIterator struct {
	Event *BalancerV2SwapFeePercentageChangedSwapFeePercentageChanged // Event containing the contract specifics and raw log

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
func (it *BalancerV2SwapFeePercentageChangedSwapFeePercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerV2SwapFeePercentageChangedSwapFeePercentageChanged)
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
		it.Event = new(BalancerV2SwapFeePercentageChangedSwapFeePercentageChanged)
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
func (it *BalancerV2SwapFeePercentageChangedSwapFeePercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerV2SwapFeePercentageChangedSwapFeePercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerV2SwapFeePercentageChangedSwapFeePercentageChanged represents a SwapFeePercentageChanged event raised by the BalancerV2SwapFeePercentageChanged contract.
type BalancerV2SwapFeePercentageChangedSwapFeePercentageChanged struct {
	SwapFeePercentage *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSwapFeePercentageChanged is a free log retrieval operation binding the contract event 0xa9ba3ffe0b6c366b81232caab38605a0699ad5398d6cce76f91ee809e322dafc.
//
// Solidity: event SwapFeePercentageChanged(uint256 swapFeePercentage)
func (_BalancerV2SwapFeePercentageChanged *BalancerV2SwapFeePercentageChangedFilterer) FilterSwapFeePercentageChanged(opts *bind.FilterOpts) (*BalancerV2SwapFeePercentageChangedSwapFeePercentageChangedIterator, error) {

	logs, sub, err := _BalancerV2SwapFeePercentageChanged.contract.FilterLogs(opts, "SwapFeePercentageChanged")
	if err != nil {
		return nil, err
	}
	return &BalancerV2SwapFeePercentageChangedSwapFeePercentageChangedIterator{contract: _BalancerV2SwapFeePercentageChanged.contract, event: "SwapFeePercentageChanged", logs: logs, sub: sub}, nil
}

// WatchSwapFeePercentageChanged is a free log subscription operation binding the contract event 0xa9ba3ffe0b6c366b81232caab38605a0699ad5398d6cce76f91ee809e322dafc.
//
// Solidity: event SwapFeePercentageChanged(uint256 swapFeePercentage)
func (_BalancerV2SwapFeePercentageChanged *BalancerV2SwapFeePercentageChangedFilterer) WatchSwapFeePercentageChanged(opts *bind.WatchOpts, sink chan<- *BalancerV2SwapFeePercentageChangedSwapFeePercentageChanged) (event.Subscription, error) {

	logs, sub, err := _BalancerV2SwapFeePercentageChanged.contract.WatchLogs(opts, "SwapFeePercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerV2SwapFeePercentageChangedSwapFeePercentageChanged)
				if err := _BalancerV2SwapFeePercentageChanged.contract.UnpackLog(event, "SwapFeePercentageChanged", log); err != nil {
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

// ParseSwapFeePercentageChanged is a log parse operation binding the contract event 0xa9ba3ffe0b6c366b81232caab38605a0699ad5398d6cce76f91ee809e322dafc.
//
// Solidity: event SwapFeePercentageChanged(uint256 swapFeePercentage)
func (_BalancerV2SwapFeePercentageChanged *BalancerV2SwapFeePercentageChangedFilterer) ParseSwapFeePercentageChanged(log types.Log) (*BalancerV2SwapFeePercentageChangedSwapFeePercentageChanged, error) {
	event := new(BalancerV2SwapFeePercentageChangedSwapFeePercentageChanged)
	if err := _BalancerV2SwapFeePercentageChanged.contract.UnpackLog(event, "SwapFeePercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
