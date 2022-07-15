// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// IUniswapV3PoolEventsMetaData contains all meta data concerning the IUniswapV3PoolEvents contract.
var IUniswapV3PoolEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"name\":\"Collect\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"name\":\"CollectProtocol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"slot0Tick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityDelta\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityBefore\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"}],\"name\":\"DBG_MOD_POS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtStartPrice\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimit\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"slot0sqrtpricex96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"DBG_SWAP_HEAD\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceStartX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceNextX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickCumulative\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stepAmountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stepAmountOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountProcessed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"exactInput\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGrowthGlobalX128\",\"type\":\"uint256\"}],\"name\":\"DBG_SWAP_LOOP\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGrowthGlobal0X128\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGrowthGlobal1X128\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityDiff\",\"type\":\"uint128\"}],\"name\":\"DBG_SWAP_TAIL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"liquidityDelta\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGrowthGlobal0X128Before\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGrowthGlobal1X128Before\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGrowthInside0X128\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGrowthInside1X128\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flippedLower\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flippedUpper\",\"type\":\"bool\"}],\"name\":\"DBG_UPD_POS\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid1\",\"type\":\"uint256\"}],\"name\":\"Flash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"observationCardinalityNextOld\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"observationCardinalityNextNew\",\"type\":\"uint16\"}],\"name\":\"IncreaseObservationCardinalityNext\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Initialize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol0Old\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol1Old\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol0New\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeProtocol1New\",\"type\":\"uint8\"}],\"name\":\"SetFeeProtocol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Swap\",\"type\":\"event\"}]",
}

// IUniswapV3PoolEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use IUniswapV3PoolEventsMetaData.ABI instead.
var IUniswapV3PoolEventsABI = IUniswapV3PoolEventsMetaData.ABI

// IUniswapV3PoolEvents is an auto generated Go binding around an Ethereum contract.
type IUniswapV3PoolEvents struct {
	IUniswapV3PoolEventsCaller     // Read-only binding to the contract
	IUniswapV3PoolEventsTransactor // Write-only binding to the contract
	IUniswapV3PoolEventsFilterer   // Log filterer for contract events
}

// IUniswapV3PoolEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUniswapV3PoolEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUniswapV3PoolEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUniswapV3PoolEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV3PoolEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUniswapV3PoolEventsSession struct {
	Contract     *IUniswapV3PoolEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IUniswapV3PoolEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUniswapV3PoolEventsCallerSession struct {
	Contract *IUniswapV3PoolEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IUniswapV3PoolEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUniswapV3PoolEventsTransactorSession struct {
	Contract     *IUniswapV3PoolEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IUniswapV3PoolEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUniswapV3PoolEventsRaw struct {
	Contract *IUniswapV3PoolEvents // Generic contract binding to access the raw methods on
}

// IUniswapV3PoolEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUniswapV3PoolEventsCallerRaw struct {
	Contract *IUniswapV3PoolEventsCaller // Generic read-only contract binding to access the raw methods on
}

// IUniswapV3PoolEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUniswapV3PoolEventsTransactorRaw struct {
	Contract *IUniswapV3PoolEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUniswapV3PoolEvents creates a new instance of IUniswapV3PoolEvents, bound to a specific deployed contract.
func NewIUniswapV3PoolEvents(address common.Address, backend bind.ContractBackend) (*IUniswapV3PoolEvents, error) {
	contract, err := bindIUniswapV3PoolEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEvents{IUniswapV3PoolEventsCaller: IUniswapV3PoolEventsCaller{contract: contract}, IUniswapV3PoolEventsTransactor: IUniswapV3PoolEventsTransactor{contract: contract}, IUniswapV3PoolEventsFilterer: IUniswapV3PoolEventsFilterer{contract: contract}}, nil
}

// NewIUniswapV3PoolEventsCaller creates a new read-only instance of IUniswapV3PoolEvents, bound to a specific deployed contract.
func NewIUniswapV3PoolEventsCaller(address common.Address, caller bind.ContractCaller) (*IUniswapV3PoolEventsCaller, error) {
	contract, err := bindIUniswapV3PoolEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsCaller{contract: contract}, nil
}

// NewIUniswapV3PoolEventsTransactor creates a new write-only instance of IUniswapV3PoolEvents, bound to a specific deployed contract.
func NewIUniswapV3PoolEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*IUniswapV3PoolEventsTransactor, error) {
	contract, err := bindIUniswapV3PoolEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsTransactor{contract: contract}, nil
}

// NewIUniswapV3PoolEventsFilterer creates a new log filterer instance of IUniswapV3PoolEvents, bound to a specific deployed contract.
func NewIUniswapV3PoolEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*IUniswapV3PoolEventsFilterer, error) {
	contract, err := bindIUniswapV3PoolEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsFilterer{contract: contract}, nil
}

// bindIUniswapV3PoolEvents binds a generic wrapper to an already deployed contract.
func bindIUniswapV3PoolEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IUniswapV3PoolEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV3PoolEvents.Contract.IUniswapV3PoolEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV3PoolEvents.Contract.IUniswapV3PoolEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV3PoolEvents.Contract.IUniswapV3PoolEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV3PoolEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV3PoolEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV3PoolEvents.Contract.contract.Transact(opts, method, params...)
}

// IUniswapV3PoolEventsBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsBurnIterator struct {
	Event *IUniswapV3PoolEventsBurn // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolEventsBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolEventsBurn)
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
		it.Event = new(IUniswapV3PoolEventsBurn)
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
func (it *IUniswapV3PoolEventsBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolEventsBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolEventsBurn represents a Burn event raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsBurn struct {
	Owner     common.Address
	TickLower *big.Int
	TickUpper *big.Int
	Amount    *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) FilterBurn(opts *bind.FilterOpts, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (*IUniswapV3PoolEventsBurnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.FilterLogs(opts, "Burn", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsBurnIterator{contract: _IUniswapV3PoolEvents.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolEventsBurn, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.WatchLogs(opts, "Burn", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolEventsBurn)
				if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) ParseBurn(log types.Log) (*IUniswapV3PoolEventsBurn, error) {
	event := new(IUniswapV3PoolEventsBurn)
	if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolEventsCollectIterator is returned from FilterCollect and is used to iterate over the raw logs and unpacked data for Collect events raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsCollectIterator struct {
	Event *IUniswapV3PoolEventsCollect // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolEventsCollectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolEventsCollect)
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
		it.Event = new(IUniswapV3PoolEventsCollect)
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
func (it *IUniswapV3PoolEventsCollectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolEventsCollectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolEventsCollect represents a Collect event raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsCollect struct {
	Owner     common.Address
	Recipient common.Address
	TickLower *big.Int
	TickUpper *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollect is a free log retrieval operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount0, uint128 amount1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) FilterCollect(opts *bind.FilterOpts, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (*IUniswapV3PoolEventsCollectIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.FilterLogs(opts, "Collect", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsCollectIterator{contract: _IUniswapV3PoolEvents.contract, event: "Collect", logs: logs, sub: sub}, nil
}

// WatchCollect is a free log subscription operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount0, uint128 amount1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) WatchCollect(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolEventsCollect, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.WatchLogs(opts, "Collect", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolEventsCollect)
				if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "Collect", log); err != nil {
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

// ParseCollect is a log parse operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount0, uint128 amount1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) ParseCollect(log types.Log) (*IUniswapV3PoolEventsCollect, error) {
	event := new(IUniswapV3PoolEventsCollect)
	if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "Collect", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolEventsCollectProtocolIterator is returned from FilterCollectProtocol and is used to iterate over the raw logs and unpacked data for CollectProtocol events raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsCollectProtocolIterator struct {
	Event *IUniswapV3PoolEventsCollectProtocol // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolEventsCollectProtocolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolEventsCollectProtocol)
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
		it.Event = new(IUniswapV3PoolEventsCollectProtocol)
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
func (it *IUniswapV3PoolEventsCollectProtocolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolEventsCollectProtocolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolEventsCollectProtocol represents a CollectProtocol event raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsCollectProtocol struct {
	Sender    common.Address
	Recipient common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollectProtocol is a free log retrieval operation binding the contract event 0x596b573906218d3411850b26a6b437d6c4522fdb43d2d2386263f86d50b8b151.
//
// Solidity: event CollectProtocol(address indexed sender, address indexed recipient, uint128 amount0, uint128 amount1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) FilterCollectProtocol(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*IUniswapV3PoolEventsCollectProtocolIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.FilterLogs(opts, "CollectProtocol", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsCollectProtocolIterator{contract: _IUniswapV3PoolEvents.contract, event: "CollectProtocol", logs: logs, sub: sub}, nil
}

// WatchCollectProtocol is a free log subscription operation binding the contract event 0x596b573906218d3411850b26a6b437d6c4522fdb43d2d2386263f86d50b8b151.
//
// Solidity: event CollectProtocol(address indexed sender, address indexed recipient, uint128 amount0, uint128 amount1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) WatchCollectProtocol(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolEventsCollectProtocol, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.WatchLogs(opts, "CollectProtocol", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolEventsCollectProtocol)
				if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "CollectProtocol", log); err != nil {
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

// ParseCollectProtocol is a log parse operation binding the contract event 0x596b573906218d3411850b26a6b437d6c4522fdb43d2d2386263f86d50b8b151.
//
// Solidity: event CollectProtocol(address indexed sender, address indexed recipient, uint128 amount0, uint128 amount1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) ParseCollectProtocol(log types.Log) (*IUniswapV3PoolEventsCollectProtocol, error) {
	event := new(IUniswapV3PoolEventsCollectProtocol)
	if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "CollectProtocol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolEventsDBGMODPOSIterator is returned from FilterDBGMODPOS and is used to iterate over the raw logs and unpacked data for DBGMODPOS events raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsDBGMODPOSIterator struct {
	Event *IUniswapV3PoolEventsDBGMODPOS // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolEventsDBGMODPOSIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolEventsDBGMODPOS)
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
		it.Event = new(IUniswapV3PoolEventsDBGMODPOS)
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
func (it *IUniswapV3PoolEventsDBGMODPOSIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolEventsDBGMODPOSIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolEventsDBGMODPOS represents a DBGMODPOS event raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsDBGMODPOS struct {
	Owner           common.Address
	TickLower       *big.Int
	TickUpper       *big.Int
	Slot0Tick       *big.Int
	LiquidityDelta  *big.Int
	LiquidityBefore *big.Int
	Amount0         *big.Int
	Amount1         *big.Int
	SqrtPriceX96    *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDBGMODPOS is a free log retrieval operation binding the contract event 0xb05d03afcde6cc3b7059da906a709c3e3c68c3c665a90c35ae0ade8d0903f666.
//
// Solidity: event DBG_MOD_POS(address owner, int24 tickLower, int24 tickUpper, int24 slot0Tick, uint128 liquidityDelta, uint128 liquidityBefore, int256 amount0, int256 amount1, uint160 sqrtPriceX96)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) FilterDBGMODPOS(opts *bind.FilterOpts) (*IUniswapV3PoolEventsDBGMODPOSIterator, error) {

	logs, sub, err := _IUniswapV3PoolEvents.contract.FilterLogs(opts, "DBG_MOD_POS")
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsDBGMODPOSIterator{contract: _IUniswapV3PoolEvents.contract, event: "DBG_MOD_POS", logs: logs, sub: sub}, nil
}

// WatchDBGMODPOS is a free log subscription operation binding the contract event 0xb05d03afcde6cc3b7059da906a709c3e3c68c3c665a90c35ae0ade8d0903f666.
//
// Solidity: event DBG_MOD_POS(address owner, int24 tickLower, int24 tickUpper, int24 slot0Tick, uint128 liquidityDelta, uint128 liquidityBefore, int256 amount0, int256 amount1, uint160 sqrtPriceX96)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) WatchDBGMODPOS(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolEventsDBGMODPOS) (event.Subscription, error) {

	logs, sub, err := _IUniswapV3PoolEvents.contract.WatchLogs(opts, "DBG_MOD_POS")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolEventsDBGMODPOS)
				if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "DBG_MOD_POS", log); err != nil {
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

// ParseDBGMODPOS is a log parse operation binding the contract event 0xb05d03afcde6cc3b7059da906a709c3e3c68c3c665a90c35ae0ade8d0903f666.
//
// Solidity: event DBG_MOD_POS(address owner, int24 tickLower, int24 tickUpper, int24 slot0Tick, uint128 liquidityDelta, uint128 liquidityBefore, int256 amount0, int256 amount1, uint160 sqrtPriceX96)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) ParseDBGMODPOS(log types.Log) (*IUniswapV3PoolEventsDBGMODPOS, error) {
	event := new(IUniswapV3PoolEventsDBGMODPOS)
	if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "DBG_MOD_POS", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolEventsDBGSWAPHEADIterator is returned from FilterDBGSWAPHEAD and is used to iterate over the raw logs and unpacked data for DBGSWAPHEAD events raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsDBGSWAPHEADIterator struct {
	Event *IUniswapV3PoolEventsDBGSWAPHEAD // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolEventsDBGSWAPHEADIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolEventsDBGSWAPHEAD)
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
		it.Event = new(IUniswapV3PoolEventsDBGSWAPHEAD)
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
func (it *IUniswapV3PoolEventsDBGSWAPHEADIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolEventsDBGSWAPHEADIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolEventsDBGSWAPHEAD represents a DBGSWAPHEAD event raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsDBGSWAPHEAD struct {
	Pool              common.Address
	ZeroForOne        bool
	Amount            *big.Int
	SqrtStartPrice    *big.Int
	SqrtPriceLimit    *big.Int
	Slot0sqrtpricex96 *big.Int
	Tick              *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDBGSWAPHEAD is a free log retrieval operation binding the contract event 0x884bf0b67c5a7a033ec93aeab39846c4e697715e51811c8b9499b2666468c207.
//
// Solidity: event DBG_SWAP_HEAD(address indexed pool, bool zeroForOne, int256 amount, uint160 sqrtStartPrice, uint160 sqrtPriceLimit, uint160 slot0sqrtpricex96, int24 tick)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) FilterDBGSWAPHEAD(opts *bind.FilterOpts, pool []common.Address) (*IUniswapV3PoolEventsDBGSWAPHEADIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.FilterLogs(opts, "DBG_SWAP_HEAD", poolRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsDBGSWAPHEADIterator{contract: _IUniswapV3PoolEvents.contract, event: "DBG_SWAP_HEAD", logs: logs, sub: sub}, nil
}

// WatchDBGSWAPHEAD is a free log subscription operation binding the contract event 0x884bf0b67c5a7a033ec93aeab39846c4e697715e51811c8b9499b2666468c207.
//
// Solidity: event DBG_SWAP_HEAD(address indexed pool, bool zeroForOne, int256 amount, uint160 sqrtStartPrice, uint160 sqrtPriceLimit, uint160 slot0sqrtpricex96, int24 tick)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) WatchDBGSWAPHEAD(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolEventsDBGSWAPHEAD, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.WatchLogs(opts, "DBG_SWAP_HEAD", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolEventsDBGSWAPHEAD)
				if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "DBG_SWAP_HEAD", log); err != nil {
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

// ParseDBGSWAPHEAD is a log parse operation binding the contract event 0x884bf0b67c5a7a033ec93aeab39846c4e697715e51811c8b9499b2666468c207.
//
// Solidity: event DBG_SWAP_HEAD(address indexed pool, bool zeroForOne, int256 amount, uint160 sqrtStartPrice, uint160 sqrtPriceLimit, uint160 slot0sqrtpricex96, int24 tick)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) ParseDBGSWAPHEAD(log types.Log) (*IUniswapV3PoolEventsDBGSWAPHEAD, error) {
	event := new(IUniswapV3PoolEventsDBGSWAPHEAD)
	if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "DBG_SWAP_HEAD", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolEventsDBGSWAPLOOPIterator is returned from FilterDBGSWAPLOOP and is used to iterate over the raw logs and unpacked data for DBGSWAPLOOP events raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsDBGSWAPLOOPIterator struct {
	Event *IUniswapV3PoolEventsDBGSWAPLOOP // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolEventsDBGSWAPLOOPIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolEventsDBGSWAPLOOP)
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
		it.Event = new(IUniswapV3PoolEventsDBGSWAPLOOP)
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
func (it *IUniswapV3PoolEventsDBGSWAPLOOPIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolEventsDBGSWAPLOOPIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolEventsDBGSWAPLOOP represents a DBGSWAPLOOP event raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsDBGSWAPLOOP struct {
	Pool                common.Address
	SqrtPriceX96        *big.Int
	SqrtPriceStartX96   *big.Int
	SqrtPriceNextX96    *big.Int
	Tick                *big.Int
	TickCumulative      *big.Int
	Initialized         bool
	StepAmountIn        *big.Int
	StepAmountOut       *big.Int
	AmountProcessed     *big.Int
	FeeAmount           *big.Int
	Liquidity           *big.Int
	ExactInput          bool
	FeeGrowthGlobalX128 *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDBGSWAPLOOP is a free log retrieval operation binding the contract event 0x5cdbca1f8631135a1302301500a65ffc83d0b7c1177e309dfc6a14e4ea6c58f3.
//
// Solidity: event DBG_SWAP_LOOP(address indexed pool, uint160 sqrtPriceX96, uint160 sqrtPriceStartX96, uint160 sqrtPriceNextX96, int24 tick, int24 tickCumulative, bool initialized, uint256 stepAmountIn, uint256 stepAmountOut, uint256 amountProcessed, uint256 feeAmount, uint128 liquidity, bool exactInput, uint256 feeGrowthGlobalX128)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) FilterDBGSWAPLOOP(opts *bind.FilterOpts, pool []common.Address) (*IUniswapV3PoolEventsDBGSWAPLOOPIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.FilterLogs(opts, "DBG_SWAP_LOOP", poolRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsDBGSWAPLOOPIterator{contract: _IUniswapV3PoolEvents.contract, event: "DBG_SWAP_LOOP", logs: logs, sub: sub}, nil
}

// WatchDBGSWAPLOOP is a free log subscription operation binding the contract event 0x5cdbca1f8631135a1302301500a65ffc83d0b7c1177e309dfc6a14e4ea6c58f3.
//
// Solidity: event DBG_SWAP_LOOP(address indexed pool, uint160 sqrtPriceX96, uint160 sqrtPriceStartX96, uint160 sqrtPriceNextX96, int24 tick, int24 tickCumulative, bool initialized, uint256 stepAmountIn, uint256 stepAmountOut, uint256 amountProcessed, uint256 feeAmount, uint128 liquidity, bool exactInput, uint256 feeGrowthGlobalX128)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) WatchDBGSWAPLOOP(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolEventsDBGSWAPLOOP, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.WatchLogs(opts, "DBG_SWAP_LOOP", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolEventsDBGSWAPLOOP)
				if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "DBG_SWAP_LOOP", log); err != nil {
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

// ParseDBGSWAPLOOP is a log parse operation binding the contract event 0x5cdbca1f8631135a1302301500a65ffc83d0b7c1177e309dfc6a14e4ea6c58f3.
//
// Solidity: event DBG_SWAP_LOOP(address indexed pool, uint160 sqrtPriceX96, uint160 sqrtPriceStartX96, uint160 sqrtPriceNextX96, int24 tick, int24 tickCumulative, bool initialized, uint256 stepAmountIn, uint256 stepAmountOut, uint256 amountProcessed, uint256 feeAmount, uint128 liquidity, bool exactInput, uint256 feeGrowthGlobalX128)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) ParseDBGSWAPLOOP(log types.Log) (*IUniswapV3PoolEventsDBGSWAPLOOP, error) {
	event := new(IUniswapV3PoolEventsDBGSWAPLOOP)
	if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "DBG_SWAP_LOOP", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolEventsDBGSWAPTAILIterator is returned from FilterDBGSWAPTAIL and is used to iterate over the raw logs and unpacked data for DBGSWAPTAIL events raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsDBGSWAPTAILIterator struct {
	Event *IUniswapV3PoolEventsDBGSWAPTAIL // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolEventsDBGSWAPTAILIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolEventsDBGSWAPTAIL)
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
		it.Event = new(IUniswapV3PoolEventsDBGSWAPTAIL)
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
func (it *IUniswapV3PoolEventsDBGSWAPTAILIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolEventsDBGSWAPTAILIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolEventsDBGSWAPTAIL represents a DBGSWAPTAIL event raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsDBGSWAPTAIL struct {
	Pool                 common.Address
	FeeGrowthGlobal0X128 *big.Int
	FeeGrowthGlobal1X128 *big.Int
	Tick                 *big.Int
	SqrtPriceX96         *big.Int
	Liquidity            *big.Int
	LiquidityDiff        *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterDBGSWAPTAIL is a free log retrieval operation binding the contract event 0x507cf18262378ccee2abf0e6b16ba3a87afe50f610f6d98fec691ae6ab12a584.
//
// Solidity: event DBG_SWAP_TAIL(address indexed pool, uint256 feeGrowthGlobal0X128, uint256 feeGrowthGlobal1X128, int24 tick, uint160 sqrtPriceX96, uint128 liquidity, uint128 liquidityDiff)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) FilterDBGSWAPTAIL(opts *bind.FilterOpts, pool []common.Address) (*IUniswapV3PoolEventsDBGSWAPTAILIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.FilterLogs(opts, "DBG_SWAP_TAIL", poolRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsDBGSWAPTAILIterator{contract: _IUniswapV3PoolEvents.contract, event: "DBG_SWAP_TAIL", logs: logs, sub: sub}, nil
}

// WatchDBGSWAPTAIL is a free log subscription operation binding the contract event 0x507cf18262378ccee2abf0e6b16ba3a87afe50f610f6d98fec691ae6ab12a584.
//
// Solidity: event DBG_SWAP_TAIL(address indexed pool, uint256 feeGrowthGlobal0X128, uint256 feeGrowthGlobal1X128, int24 tick, uint160 sqrtPriceX96, uint128 liquidity, uint128 liquidityDiff)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) WatchDBGSWAPTAIL(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolEventsDBGSWAPTAIL, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.WatchLogs(opts, "DBG_SWAP_TAIL", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolEventsDBGSWAPTAIL)
				if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "DBG_SWAP_TAIL", log); err != nil {
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

// ParseDBGSWAPTAIL is a log parse operation binding the contract event 0x507cf18262378ccee2abf0e6b16ba3a87afe50f610f6d98fec691ae6ab12a584.
//
// Solidity: event DBG_SWAP_TAIL(address indexed pool, uint256 feeGrowthGlobal0X128, uint256 feeGrowthGlobal1X128, int24 tick, uint160 sqrtPriceX96, uint128 liquidity, uint128 liquidityDiff)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) ParseDBGSWAPTAIL(log types.Log) (*IUniswapV3PoolEventsDBGSWAPTAIL, error) {
	event := new(IUniswapV3PoolEventsDBGSWAPTAIL)
	if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "DBG_SWAP_TAIL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolEventsDBGUPDPOSIterator is returned from FilterDBGUPDPOS and is used to iterate over the raw logs and unpacked data for DBGUPDPOS events raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsDBGUPDPOSIterator struct {
	Event *IUniswapV3PoolEventsDBGUPDPOS // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolEventsDBGUPDPOSIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolEventsDBGUPDPOS)
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
		it.Event = new(IUniswapV3PoolEventsDBGUPDPOS)
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
func (it *IUniswapV3PoolEventsDBGUPDPOSIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolEventsDBGUPDPOSIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolEventsDBGUPDPOS represents a DBGUPDPOS event raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsDBGUPDPOS struct {
	Owner                      common.Address
	TickLower                  *big.Int
	TickUpper                  *big.Int
	Tick                       *big.Int
	LiquidityDelta             *big.Int
	FeeGrowthGlobal0X128Before *big.Int
	FeeGrowthGlobal1X128Before *big.Int
	FeeGrowthInside0X128       *big.Int
	FeeGrowthInside1X128       *big.Int
	FlippedLower               bool
	FlippedUpper               bool
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterDBGUPDPOS is a free log retrieval operation binding the contract event 0x2e894cce336ea758e6f37d72cafa9ef74e237e6e488a2f6ad4148c3fbca85f24.
//
// Solidity: event DBG_UPD_POS(address owner, int24 tickLower, int24 tickUpper, int24 tick, int128 liquidityDelta, uint256 feeGrowthGlobal0X128Before, uint256 feeGrowthGlobal1X128Before, uint256 feeGrowthInside0X128, uint256 feeGrowthInside1X128, bool flippedLower, bool flippedUpper)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) FilterDBGUPDPOS(opts *bind.FilterOpts) (*IUniswapV3PoolEventsDBGUPDPOSIterator, error) {

	logs, sub, err := _IUniswapV3PoolEvents.contract.FilterLogs(opts, "DBG_UPD_POS")
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsDBGUPDPOSIterator{contract: _IUniswapV3PoolEvents.contract, event: "DBG_UPD_POS", logs: logs, sub: sub}, nil
}

// WatchDBGUPDPOS is a free log subscription operation binding the contract event 0x2e894cce336ea758e6f37d72cafa9ef74e237e6e488a2f6ad4148c3fbca85f24.
//
// Solidity: event DBG_UPD_POS(address owner, int24 tickLower, int24 tickUpper, int24 tick, int128 liquidityDelta, uint256 feeGrowthGlobal0X128Before, uint256 feeGrowthGlobal1X128Before, uint256 feeGrowthInside0X128, uint256 feeGrowthInside1X128, bool flippedLower, bool flippedUpper)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) WatchDBGUPDPOS(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolEventsDBGUPDPOS) (event.Subscription, error) {

	logs, sub, err := _IUniswapV3PoolEvents.contract.WatchLogs(opts, "DBG_UPD_POS")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolEventsDBGUPDPOS)
				if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "DBG_UPD_POS", log); err != nil {
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

// ParseDBGUPDPOS is a log parse operation binding the contract event 0x2e894cce336ea758e6f37d72cafa9ef74e237e6e488a2f6ad4148c3fbca85f24.
//
// Solidity: event DBG_UPD_POS(address owner, int24 tickLower, int24 tickUpper, int24 tick, int128 liquidityDelta, uint256 feeGrowthGlobal0X128Before, uint256 feeGrowthGlobal1X128Before, uint256 feeGrowthInside0X128, uint256 feeGrowthInside1X128, bool flippedLower, bool flippedUpper)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) ParseDBGUPDPOS(log types.Log) (*IUniswapV3PoolEventsDBGUPDPOS, error) {
	event := new(IUniswapV3PoolEventsDBGUPDPOS)
	if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "DBG_UPD_POS", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolEventsFlashIterator is returned from FilterFlash and is used to iterate over the raw logs and unpacked data for Flash events raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsFlashIterator struct {
	Event *IUniswapV3PoolEventsFlash // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolEventsFlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolEventsFlash)
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
		it.Event = new(IUniswapV3PoolEventsFlash)
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
func (it *IUniswapV3PoolEventsFlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolEventsFlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolEventsFlash represents a Flash event raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsFlash struct {
	Sender    common.Address
	Recipient common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Paid0     *big.Int
	Paid1     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFlash is a free log retrieval operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) FilterFlash(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*IUniswapV3PoolEventsFlashIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.FilterLogs(opts, "Flash", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsFlashIterator{contract: _IUniswapV3PoolEvents.contract, event: "Flash", logs: logs, sub: sub}, nil
}

// WatchFlash is a free log subscription operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) WatchFlash(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolEventsFlash, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.WatchLogs(opts, "Flash", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolEventsFlash)
				if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "Flash", log); err != nil {
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

// ParseFlash is a log parse operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) ParseFlash(log types.Log) (*IUniswapV3PoolEventsFlash, error) {
	event := new(IUniswapV3PoolEventsFlash)
	if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "Flash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolEventsIncreaseObservationCardinalityNextIterator is returned from FilterIncreaseObservationCardinalityNext and is used to iterate over the raw logs and unpacked data for IncreaseObservationCardinalityNext events raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsIncreaseObservationCardinalityNextIterator struct {
	Event *IUniswapV3PoolEventsIncreaseObservationCardinalityNext // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolEventsIncreaseObservationCardinalityNextIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolEventsIncreaseObservationCardinalityNext)
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
		it.Event = new(IUniswapV3PoolEventsIncreaseObservationCardinalityNext)
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
func (it *IUniswapV3PoolEventsIncreaseObservationCardinalityNextIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolEventsIncreaseObservationCardinalityNextIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolEventsIncreaseObservationCardinalityNext represents a IncreaseObservationCardinalityNext event raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsIncreaseObservationCardinalityNext struct {
	ObservationCardinalityNextOld uint16
	ObservationCardinalityNextNew uint16
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterIncreaseObservationCardinalityNext is a free log retrieval operation binding the contract event 0xac49e518f90a358f652e4400164f05a5d8f7e35e7747279bc3a93dbf584e125a.
//
// Solidity: event IncreaseObservationCardinalityNext(uint16 observationCardinalityNextOld, uint16 observationCardinalityNextNew)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) FilterIncreaseObservationCardinalityNext(opts *bind.FilterOpts) (*IUniswapV3PoolEventsIncreaseObservationCardinalityNextIterator, error) {

	logs, sub, err := _IUniswapV3PoolEvents.contract.FilterLogs(opts, "IncreaseObservationCardinalityNext")
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsIncreaseObservationCardinalityNextIterator{contract: _IUniswapV3PoolEvents.contract, event: "IncreaseObservationCardinalityNext", logs: logs, sub: sub}, nil
}

// WatchIncreaseObservationCardinalityNext is a free log subscription operation binding the contract event 0xac49e518f90a358f652e4400164f05a5d8f7e35e7747279bc3a93dbf584e125a.
//
// Solidity: event IncreaseObservationCardinalityNext(uint16 observationCardinalityNextOld, uint16 observationCardinalityNextNew)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) WatchIncreaseObservationCardinalityNext(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolEventsIncreaseObservationCardinalityNext) (event.Subscription, error) {

	logs, sub, err := _IUniswapV3PoolEvents.contract.WatchLogs(opts, "IncreaseObservationCardinalityNext")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolEventsIncreaseObservationCardinalityNext)
				if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "IncreaseObservationCardinalityNext", log); err != nil {
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

// ParseIncreaseObservationCardinalityNext is a log parse operation binding the contract event 0xac49e518f90a358f652e4400164f05a5d8f7e35e7747279bc3a93dbf584e125a.
//
// Solidity: event IncreaseObservationCardinalityNext(uint16 observationCardinalityNextOld, uint16 observationCardinalityNextNew)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) ParseIncreaseObservationCardinalityNext(log types.Log) (*IUniswapV3PoolEventsIncreaseObservationCardinalityNext, error) {
	event := new(IUniswapV3PoolEventsIncreaseObservationCardinalityNext)
	if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "IncreaseObservationCardinalityNext", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolEventsInitializeIterator is returned from FilterInitialize and is used to iterate over the raw logs and unpacked data for Initialize events raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsInitializeIterator struct {
	Event *IUniswapV3PoolEventsInitialize // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolEventsInitializeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolEventsInitialize)
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
		it.Event = new(IUniswapV3PoolEventsInitialize)
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
func (it *IUniswapV3PoolEventsInitializeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolEventsInitializeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolEventsInitialize represents a Initialize event raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsInitialize struct {
	SqrtPriceX96 *big.Int
	Tick         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInitialize is a free log retrieval operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 sqrtPriceX96, int24 tick)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) FilterInitialize(opts *bind.FilterOpts) (*IUniswapV3PoolEventsInitializeIterator, error) {

	logs, sub, err := _IUniswapV3PoolEvents.contract.FilterLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsInitializeIterator{contract: _IUniswapV3PoolEvents.contract, event: "Initialize", logs: logs, sub: sub}, nil
}

// WatchInitialize is a free log subscription operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 sqrtPriceX96, int24 tick)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) WatchInitialize(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolEventsInitialize) (event.Subscription, error) {

	logs, sub, err := _IUniswapV3PoolEvents.contract.WatchLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolEventsInitialize)
				if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "Initialize", log); err != nil {
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

// ParseInitialize is a log parse operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 sqrtPriceX96, int24 tick)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) ParseInitialize(log types.Log) (*IUniswapV3PoolEventsInitialize, error) {
	event := new(IUniswapV3PoolEventsInitialize)
	if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "Initialize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolEventsMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsMintIterator struct {
	Event *IUniswapV3PoolEventsMint // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolEventsMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolEventsMint)
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
		it.Event = new(IUniswapV3PoolEventsMint)
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
func (it *IUniswapV3PoolEventsMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolEventsMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolEventsMint represents a Mint event raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsMint struct {
	Sender    common.Address
	Owner     common.Address
	TickLower *big.Int
	TickUpper *big.Int
	Amount    *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) FilterMint(opts *bind.FilterOpts, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (*IUniswapV3PoolEventsMintIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.FilterLogs(opts, "Mint", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsMintIterator{contract: _IUniswapV3PoolEvents.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolEventsMint, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.WatchLogs(opts, "Mint", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolEventsMint)
				if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) ParseMint(log types.Log) (*IUniswapV3PoolEventsMint, error) {
	event := new(IUniswapV3PoolEventsMint)
	if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolEventsSetFeeProtocolIterator is returned from FilterSetFeeProtocol and is used to iterate over the raw logs and unpacked data for SetFeeProtocol events raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsSetFeeProtocolIterator struct {
	Event *IUniswapV3PoolEventsSetFeeProtocol // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolEventsSetFeeProtocolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolEventsSetFeeProtocol)
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
		it.Event = new(IUniswapV3PoolEventsSetFeeProtocol)
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
func (it *IUniswapV3PoolEventsSetFeeProtocolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolEventsSetFeeProtocolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolEventsSetFeeProtocol represents a SetFeeProtocol event raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsSetFeeProtocol struct {
	FeeProtocol0Old uint8
	FeeProtocol1Old uint8
	FeeProtocol0New uint8
	FeeProtocol1New uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetFeeProtocol is a free log retrieval operation binding the contract event 0x973d8d92bb299f4af6ce49b52a8adb85ae46b9f214c4c4fc06ac77401237b133.
//
// Solidity: event SetFeeProtocol(uint8 feeProtocol0Old, uint8 feeProtocol1Old, uint8 feeProtocol0New, uint8 feeProtocol1New)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) FilterSetFeeProtocol(opts *bind.FilterOpts) (*IUniswapV3PoolEventsSetFeeProtocolIterator, error) {

	logs, sub, err := _IUniswapV3PoolEvents.contract.FilterLogs(opts, "SetFeeProtocol")
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsSetFeeProtocolIterator{contract: _IUniswapV3PoolEvents.contract, event: "SetFeeProtocol", logs: logs, sub: sub}, nil
}

// WatchSetFeeProtocol is a free log subscription operation binding the contract event 0x973d8d92bb299f4af6ce49b52a8adb85ae46b9f214c4c4fc06ac77401237b133.
//
// Solidity: event SetFeeProtocol(uint8 feeProtocol0Old, uint8 feeProtocol1Old, uint8 feeProtocol0New, uint8 feeProtocol1New)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) WatchSetFeeProtocol(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolEventsSetFeeProtocol) (event.Subscription, error) {

	logs, sub, err := _IUniswapV3PoolEvents.contract.WatchLogs(opts, "SetFeeProtocol")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolEventsSetFeeProtocol)
				if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "SetFeeProtocol", log); err != nil {
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

// ParseSetFeeProtocol is a log parse operation binding the contract event 0x973d8d92bb299f4af6ce49b52a8adb85ae46b9f214c4c4fc06ac77401237b133.
//
// Solidity: event SetFeeProtocol(uint8 feeProtocol0Old, uint8 feeProtocol1Old, uint8 feeProtocol0New, uint8 feeProtocol1New)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) ParseSetFeeProtocol(log types.Log) (*IUniswapV3PoolEventsSetFeeProtocol, error) {
	event := new(IUniswapV3PoolEventsSetFeeProtocol)
	if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "SetFeeProtocol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IUniswapV3PoolEventsSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsSwapIterator struct {
	Event *IUniswapV3PoolEventsSwap // Event containing the contract specifics and raw log

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
func (it *IUniswapV3PoolEventsSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IUniswapV3PoolEventsSwap)
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
		it.Event = new(IUniswapV3PoolEventsSwap)
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
func (it *IUniswapV3PoolEventsSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IUniswapV3PoolEventsSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IUniswapV3PoolEventsSwap represents a Swap event raised by the IUniswapV3PoolEvents contract.
type IUniswapV3PoolEventsSwap struct {
	Sender       common.Address
	Recipient    common.Address
	Amount0      *big.Int
	Amount1      *big.Int
	SqrtPriceX96 *big.Int
	Liquidity    *big.Int
	Tick         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*IUniswapV3PoolEventsSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.FilterLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &IUniswapV3PoolEventsSwapIterator{contract: _IUniswapV3PoolEvents.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *IUniswapV3PoolEventsSwap, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IUniswapV3PoolEvents.contract.WatchLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IUniswapV3PoolEventsSwap)
				if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick)
func (_IUniswapV3PoolEvents *IUniswapV3PoolEventsFilterer) ParseSwap(log types.Log) (*IUniswapV3PoolEventsSwap, error) {
	event := new(IUniswapV3PoolEventsSwap)
	if err := _IUniswapV3PoolEvents.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
