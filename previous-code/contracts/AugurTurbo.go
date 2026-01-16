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

// AugurTurboMetaData contains all meta data concerning the AugurTurbo contract.
var AugurTurboMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"settlementAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SettlementFeeClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"protocol\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ProtocolFeeClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"SettlementFeeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"protocol\",\"type\":\"address\"}],\"name\":\"ProtocolChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"marketFactory\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"marketId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"lpTokenRecipient\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"marketFactory\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"marketId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"collateral\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"lpTokens\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"sharesReturned\",\"type\":\"uint256[]\"}],\"name\":\"LiquidityChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"MarketResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SharesBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SharesMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"marketFactory\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"marketId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outcome\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"collateral\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"shares\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"SharesSwapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winningOutcome\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"settlementFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payout\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"WinningsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"}],\"name\":\"PriceMarketCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"outcomes\",\"type\":\"string[]\"}],\"name\":\"TrustedMarketCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumSportsLinkMarketFactory.MarketType\",\"name\":\"marketType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"eventId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"homeTeamId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"awayTeamId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"estimatedStarTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"score\",\"type\":\"int256\"}],\"name\":\"SportsMarketCreated\",\"type\":\"event\"}]",
}

// AugurTurboABI is the input ABI used to generate the binding from.
// Deprecated: Use AugurTurboMetaData.ABI instead.
var AugurTurboABI = AugurTurboMetaData.ABI

// AugurTurbo is an auto generated Go binding around an Ethereum contract.
type AugurTurbo struct {
	AugurTurboCaller     // Read-only binding to the contract
	AugurTurboTransactor // Write-only binding to the contract
	AugurTurboFilterer   // Log filterer for contract events
}

// AugurTurboCaller is an auto generated read-only Go binding around an Ethereum contract.
type AugurTurboCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AugurTurboTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AugurTurboTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AugurTurboFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AugurTurboFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AugurTurboSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AugurTurboSession struct {
	Contract     *AugurTurbo       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AugurTurboCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AugurTurboCallerSession struct {
	Contract *AugurTurboCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AugurTurboTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AugurTurboTransactorSession struct {
	Contract     *AugurTurboTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AugurTurboRaw is an auto generated low-level Go binding around an Ethereum contract.
type AugurTurboRaw struct {
	Contract *AugurTurbo // Generic contract binding to access the raw methods on
}

// AugurTurboCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AugurTurboCallerRaw struct {
	Contract *AugurTurboCaller // Generic read-only contract binding to access the raw methods on
}

// AugurTurboTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AugurTurboTransactorRaw struct {
	Contract *AugurTurboTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAugurTurbo creates a new instance of AugurTurbo, bound to a specific deployed contract.
func NewAugurTurbo(address common.Address, backend bind.ContractBackend) (*AugurTurbo, error) {
	contract, err := bindAugurTurbo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AugurTurbo{AugurTurboCaller: AugurTurboCaller{contract: contract}, AugurTurboTransactor: AugurTurboTransactor{contract: contract}, AugurTurboFilterer: AugurTurboFilterer{contract: contract}}, nil
}

// NewAugurTurboCaller creates a new read-only instance of AugurTurbo, bound to a specific deployed contract.
func NewAugurTurboCaller(address common.Address, caller bind.ContractCaller) (*AugurTurboCaller, error) {
	contract, err := bindAugurTurbo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AugurTurboCaller{contract: contract}, nil
}

// NewAugurTurboTransactor creates a new write-only instance of AugurTurbo, bound to a specific deployed contract.
func NewAugurTurboTransactor(address common.Address, transactor bind.ContractTransactor) (*AugurTurboTransactor, error) {
	contract, err := bindAugurTurbo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AugurTurboTransactor{contract: contract}, nil
}

// NewAugurTurboFilterer creates a new log filterer instance of AugurTurbo, bound to a specific deployed contract.
func NewAugurTurboFilterer(address common.Address, filterer bind.ContractFilterer) (*AugurTurboFilterer, error) {
	contract, err := bindAugurTurbo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AugurTurboFilterer{contract: contract}, nil
}

// bindAugurTurbo binds a generic wrapper to an already deployed contract.
func bindAugurTurbo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AugurTurboABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AugurTurbo *AugurTurboRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AugurTurbo.Contract.AugurTurboCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AugurTurbo *AugurTurboRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AugurTurbo.Contract.AugurTurboTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AugurTurbo *AugurTurboRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AugurTurbo.Contract.AugurTurboTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AugurTurbo *AugurTurboCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AugurTurbo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AugurTurbo *AugurTurboTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AugurTurbo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AugurTurbo *AugurTurboTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AugurTurbo.Contract.contract.Transact(opts, method, params...)
}

// AugurTurboLiquidityChangedIterator is returned from FilterLiquidityChanged and is used to iterate over the raw logs and unpacked data for LiquidityChanged events raised by the AugurTurbo contract.
type AugurTurboLiquidityChangedIterator struct {
	Event *AugurTurboLiquidityChanged // Event containing the contract specifics and raw log

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
func (it *AugurTurboLiquidityChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTurboLiquidityChanged)
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
		it.Event = new(AugurTurboLiquidityChanged)
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
func (it *AugurTurboLiquidityChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTurboLiquidityChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTurboLiquidityChanged represents a LiquidityChanged event raised by the AugurTurbo contract.
type AugurTurboLiquidityChanged struct {
	MarketFactory  common.Address
	MarketId       *big.Int
	User           common.Address
	Recipient      common.Address
	Collateral     *big.Int
	LpTokens       *big.Int
	SharesReturned []*big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLiquidityChanged is a free log retrieval operation binding the contract event 0x9a1dccf45b5053e827f262e45fbb5211c2bd99497d340eecaebbd245eb48f4bc.
//
// Solidity: event LiquidityChanged(address indexed marketFactory, uint256 indexed marketId, address indexed user, address recipient, int256 collateral, int256 lpTokens, uint256[] sharesReturned)
func (_AugurTurbo *AugurTurboFilterer) FilterLiquidityChanged(opts *bind.FilterOpts, marketFactory []common.Address, marketId []*big.Int, user []common.Address) (*AugurTurboLiquidityChangedIterator, error) {

	var marketFactoryRule []interface{}
	for _, marketFactoryItem := range marketFactory {
		marketFactoryRule = append(marketFactoryRule, marketFactoryItem)
	}
	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AugurTurbo.contract.FilterLogs(opts, "LiquidityChanged", marketFactoryRule, marketIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AugurTurboLiquidityChangedIterator{contract: _AugurTurbo.contract, event: "LiquidityChanged", logs: logs, sub: sub}, nil
}

// WatchLiquidityChanged is a free log subscription operation binding the contract event 0x9a1dccf45b5053e827f262e45fbb5211c2bd99497d340eecaebbd245eb48f4bc.
//
// Solidity: event LiquidityChanged(address indexed marketFactory, uint256 indexed marketId, address indexed user, address recipient, int256 collateral, int256 lpTokens, uint256[] sharesReturned)
func (_AugurTurbo *AugurTurboFilterer) WatchLiquidityChanged(opts *bind.WatchOpts, sink chan<- *AugurTurboLiquidityChanged, marketFactory []common.Address, marketId []*big.Int, user []common.Address) (event.Subscription, error) {

	var marketFactoryRule []interface{}
	for _, marketFactoryItem := range marketFactory {
		marketFactoryRule = append(marketFactoryRule, marketFactoryItem)
	}
	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AugurTurbo.contract.WatchLogs(opts, "LiquidityChanged", marketFactoryRule, marketIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTurboLiquidityChanged)
				if err := _AugurTurbo.contract.UnpackLog(event, "LiquidityChanged", log); err != nil {
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

// ParseLiquidityChanged is a log parse operation binding the contract event 0x9a1dccf45b5053e827f262e45fbb5211c2bd99497d340eecaebbd245eb48f4bc.
//
// Solidity: event LiquidityChanged(address indexed marketFactory, uint256 indexed marketId, address indexed user, address recipient, int256 collateral, int256 lpTokens, uint256[] sharesReturned)
func (_AugurTurbo *AugurTurboFilterer) ParseLiquidityChanged(log types.Log) (*AugurTurboLiquidityChanged, error) {
	event := new(AugurTurboLiquidityChanged)
	if err := _AugurTurbo.contract.UnpackLog(event, "LiquidityChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurTurboMarketResolvedIterator is returned from FilterMarketResolved and is used to iterate over the raw logs and unpacked data for MarketResolved events raised by the AugurTurbo contract.
type AugurTurboMarketResolvedIterator struct {
	Event *AugurTurboMarketResolved // Event containing the contract specifics and raw log

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
func (it *AugurTurboMarketResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTurboMarketResolved)
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
		it.Event = new(AugurTurboMarketResolved)
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
func (it *AugurTurboMarketResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTurboMarketResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTurboMarketResolved represents a MarketResolved event raised by the AugurTurbo contract.
type AugurTurboMarketResolved struct {
	Id     *big.Int
	Winner common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMarketResolved is a free log retrieval operation binding the contract event 0xc68d106ea6e4bec784925cfd91767212c71ced92adbac107dc364435321113f6.
//
// Solidity: event MarketResolved(uint256 id, address winner)
func (_AugurTurbo *AugurTurboFilterer) FilterMarketResolved(opts *bind.FilterOpts) (*AugurTurboMarketResolvedIterator, error) {

	logs, sub, err := _AugurTurbo.contract.FilterLogs(opts, "MarketResolved")
	if err != nil {
		return nil, err
	}
	return &AugurTurboMarketResolvedIterator{contract: _AugurTurbo.contract, event: "MarketResolved", logs: logs, sub: sub}, nil
}

// WatchMarketResolved is a free log subscription operation binding the contract event 0xc68d106ea6e4bec784925cfd91767212c71ced92adbac107dc364435321113f6.
//
// Solidity: event MarketResolved(uint256 id, address winner)
func (_AugurTurbo *AugurTurboFilterer) WatchMarketResolved(opts *bind.WatchOpts, sink chan<- *AugurTurboMarketResolved) (event.Subscription, error) {

	logs, sub, err := _AugurTurbo.contract.WatchLogs(opts, "MarketResolved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTurboMarketResolved)
				if err := _AugurTurbo.contract.UnpackLog(event, "MarketResolved", log); err != nil {
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

// ParseMarketResolved is a log parse operation binding the contract event 0xc68d106ea6e4bec784925cfd91767212c71ced92adbac107dc364435321113f6.
//
// Solidity: event MarketResolved(uint256 id, address winner)
func (_AugurTurbo *AugurTurboFilterer) ParseMarketResolved(log types.Log) (*AugurTurboMarketResolved, error) {
	event := new(AugurTurboMarketResolved)
	if err := _AugurTurbo.contract.UnpackLog(event, "MarketResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurTurboPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the AugurTurbo contract.
type AugurTurboPoolCreatedIterator struct {
	Event *AugurTurboPoolCreated // Event containing the contract specifics and raw log

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
func (it *AugurTurboPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTurboPoolCreated)
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
		it.Event = new(AugurTurboPoolCreated)
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
func (it *AugurTurboPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTurboPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTurboPoolCreated represents a PoolCreated event raised by the AugurTurbo contract.
type AugurTurboPoolCreated struct {
	Pool             common.Address
	MarketFactory    common.Address
	MarketId         *big.Int
	Creator          common.Address
	LpTokenRecipient common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0xfb27591822deac7913ea973cf845992f243e2ba5634ef8fbe0b6554aca37a428.
//
// Solidity: event PoolCreated(address pool, address indexed marketFactory, uint256 indexed marketId, address indexed creator, address lpTokenRecipient)
func (_AugurTurbo *AugurTurboFilterer) FilterPoolCreated(opts *bind.FilterOpts, marketFactory []common.Address, marketId []*big.Int, creator []common.Address) (*AugurTurboPoolCreatedIterator, error) {

	var marketFactoryRule []interface{}
	for _, marketFactoryItem := range marketFactory {
		marketFactoryRule = append(marketFactoryRule, marketFactoryItem)
	}
	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _AugurTurbo.contract.FilterLogs(opts, "PoolCreated", marketFactoryRule, marketIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &AugurTurboPoolCreatedIterator{contract: _AugurTurbo.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0xfb27591822deac7913ea973cf845992f243e2ba5634ef8fbe0b6554aca37a428.
//
// Solidity: event PoolCreated(address pool, address indexed marketFactory, uint256 indexed marketId, address indexed creator, address lpTokenRecipient)
func (_AugurTurbo *AugurTurboFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *AugurTurboPoolCreated, marketFactory []common.Address, marketId []*big.Int, creator []common.Address) (event.Subscription, error) {

	var marketFactoryRule []interface{}
	for _, marketFactoryItem := range marketFactory {
		marketFactoryRule = append(marketFactoryRule, marketFactoryItem)
	}
	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _AugurTurbo.contract.WatchLogs(opts, "PoolCreated", marketFactoryRule, marketIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTurboPoolCreated)
				if err := _AugurTurbo.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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

// ParsePoolCreated is a log parse operation binding the contract event 0xfb27591822deac7913ea973cf845992f243e2ba5634ef8fbe0b6554aca37a428.
//
// Solidity: event PoolCreated(address pool, address indexed marketFactory, uint256 indexed marketId, address indexed creator, address lpTokenRecipient)
func (_AugurTurbo *AugurTurboFilterer) ParsePoolCreated(log types.Log) (*AugurTurboPoolCreated, error) {
	event := new(AugurTurboPoolCreated)
	if err := _AugurTurbo.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurTurboPriceMarketCreatedIterator is returned from FilterPriceMarketCreated and is used to iterate over the raw logs and unpacked data for PriceMarketCreated events raised by the AugurTurbo contract.
type AugurTurboPriceMarketCreatedIterator struct {
	Event *AugurTurboPriceMarketCreated // Event containing the contract specifics and raw log

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
func (it *AugurTurboPriceMarketCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTurboPriceMarketCreated)
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
		it.Event = new(AugurTurboPriceMarketCreated)
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
func (it *AugurTurboPriceMarketCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTurboPriceMarketCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTurboPriceMarketCreated represents a PriceMarketCreated event raised by the AugurTurbo contract.
type AugurTurboPriceMarketCreated struct {
	Id        *big.Int
	Creator   common.Address
	EndTime   *big.Int
	SpotPrice *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPriceMarketCreated is a free log retrieval operation binding the contract event 0x6426d5b65aa6991a4001c9794fe6a39c73501502ce3f25e434e578b9a372b64b.
//
// Solidity: event PriceMarketCreated(uint256 id, address creator, uint256 endTime, uint256 spotPrice)
func (_AugurTurbo *AugurTurboFilterer) FilterPriceMarketCreated(opts *bind.FilterOpts) (*AugurTurboPriceMarketCreatedIterator, error) {

	logs, sub, err := _AugurTurbo.contract.FilterLogs(opts, "PriceMarketCreated")
	if err != nil {
		return nil, err
	}
	return &AugurTurboPriceMarketCreatedIterator{contract: _AugurTurbo.contract, event: "PriceMarketCreated", logs: logs, sub: sub}, nil
}

// WatchPriceMarketCreated is a free log subscription operation binding the contract event 0x6426d5b65aa6991a4001c9794fe6a39c73501502ce3f25e434e578b9a372b64b.
//
// Solidity: event PriceMarketCreated(uint256 id, address creator, uint256 endTime, uint256 spotPrice)
func (_AugurTurbo *AugurTurboFilterer) WatchPriceMarketCreated(opts *bind.WatchOpts, sink chan<- *AugurTurboPriceMarketCreated) (event.Subscription, error) {

	logs, sub, err := _AugurTurbo.contract.WatchLogs(opts, "PriceMarketCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTurboPriceMarketCreated)
				if err := _AugurTurbo.contract.UnpackLog(event, "PriceMarketCreated", log); err != nil {
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

// ParsePriceMarketCreated is a log parse operation binding the contract event 0x6426d5b65aa6991a4001c9794fe6a39c73501502ce3f25e434e578b9a372b64b.
//
// Solidity: event PriceMarketCreated(uint256 id, address creator, uint256 endTime, uint256 spotPrice)
func (_AugurTurbo *AugurTurboFilterer) ParsePriceMarketCreated(log types.Log) (*AugurTurboPriceMarketCreated, error) {
	event := new(AugurTurboPriceMarketCreated)
	if err := _AugurTurbo.contract.UnpackLog(event, "PriceMarketCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurTurboProtocolChangedIterator is returned from FilterProtocolChanged and is used to iterate over the raw logs and unpacked data for ProtocolChanged events raised by the AugurTurbo contract.
type AugurTurboProtocolChangedIterator struct {
	Event *AugurTurboProtocolChanged // Event containing the contract specifics and raw log

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
func (it *AugurTurboProtocolChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTurboProtocolChanged)
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
		it.Event = new(AugurTurboProtocolChanged)
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
func (it *AugurTurboProtocolChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTurboProtocolChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTurboProtocolChanged represents a ProtocolChanged event raised by the AugurTurbo contract.
type AugurTurboProtocolChanged struct {
	Protocol common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProtocolChanged is a free log retrieval operation binding the contract event 0x15b84596b3c567ae2998116949ae5f2d47f3055c12d9053db4d6e50f4c794dd9.
//
// Solidity: event ProtocolChanged(address protocol)
func (_AugurTurbo *AugurTurboFilterer) FilterProtocolChanged(opts *bind.FilterOpts) (*AugurTurboProtocolChangedIterator, error) {

	logs, sub, err := _AugurTurbo.contract.FilterLogs(opts, "ProtocolChanged")
	if err != nil {
		return nil, err
	}
	return &AugurTurboProtocolChangedIterator{contract: _AugurTurbo.contract, event: "ProtocolChanged", logs: logs, sub: sub}, nil
}

// WatchProtocolChanged is a free log subscription operation binding the contract event 0x15b84596b3c567ae2998116949ae5f2d47f3055c12d9053db4d6e50f4c794dd9.
//
// Solidity: event ProtocolChanged(address protocol)
func (_AugurTurbo *AugurTurboFilterer) WatchProtocolChanged(opts *bind.WatchOpts, sink chan<- *AugurTurboProtocolChanged) (event.Subscription, error) {

	logs, sub, err := _AugurTurbo.contract.WatchLogs(opts, "ProtocolChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTurboProtocolChanged)
				if err := _AugurTurbo.contract.UnpackLog(event, "ProtocolChanged", log); err != nil {
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

// ParseProtocolChanged is a log parse operation binding the contract event 0x15b84596b3c567ae2998116949ae5f2d47f3055c12d9053db4d6e50f4c794dd9.
//
// Solidity: event ProtocolChanged(address protocol)
func (_AugurTurbo *AugurTurboFilterer) ParseProtocolChanged(log types.Log) (*AugurTurboProtocolChanged, error) {
	event := new(AugurTurboProtocolChanged)
	if err := _AugurTurbo.contract.UnpackLog(event, "ProtocolChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurTurboProtocolFeeClaimedIterator is returned from FilterProtocolFeeClaimed and is used to iterate over the raw logs and unpacked data for ProtocolFeeClaimed events raised by the AugurTurbo contract.
type AugurTurboProtocolFeeClaimedIterator struct {
	Event *AugurTurboProtocolFeeClaimed // Event containing the contract specifics and raw log

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
func (it *AugurTurboProtocolFeeClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTurboProtocolFeeClaimed)
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
		it.Event = new(AugurTurboProtocolFeeClaimed)
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
func (it *AugurTurboProtocolFeeClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTurboProtocolFeeClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTurboProtocolFeeClaimed represents a ProtocolFeeClaimed event raised by the AugurTurbo contract.
type AugurTurboProtocolFeeClaimed struct {
	Protocol common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeClaimed is a free log retrieval operation binding the contract event 0x0f7f5b155b0b0ac6890709a2c7bf1b8bb3f675fff1e7840b4dd3c9acde59048b.
//
// Solidity: event ProtocolFeeClaimed(address protocol, uint256 amount)
func (_AugurTurbo *AugurTurboFilterer) FilterProtocolFeeClaimed(opts *bind.FilterOpts) (*AugurTurboProtocolFeeClaimedIterator, error) {

	logs, sub, err := _AugurTurbo.contract.FilterLogs(opts, "ProtocolFeeClaimed")
	if err != nil {
		return nil, err
	}
	return &AugurTurboProtocolFeeClaimedIterator{contract: _AugurTurbo.contract, event: "ProtocolFeeClaimed", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeClaimed is a free log subscription operation binding the contract event 0x0f7f5b155b0b0ac6890709a2c7bf1b8bb3f675fff1e7840b4dd3c9acde59048b.
//
// Solidity: event ProtocolFeeClaimed(address protocol, uint256 amount)
func (_AugurTurbo *AugurTurboFilterer) WatchProtocolFeeClaimed(opts *bind.WatchOpts, sink chan<- *AugurTurboProtocolFeeClaimed) (event.Subscription, error) {

	logs, sub, err := _AugurTurbo.contract.WatchLogs(opts, "ProtocolFeeClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTurboProtocolFeeClaimed)
				if err := _AugurTurbo.contract.UnpackLog(event, "ProtocolFeeClaimed", log); err != nil {
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

// ParseProtocolFeeClaimed is a log parse operation binding the contract event 0x0f7f5b155b0b0ac6890709a2c7bf1b8bb3f675fff1e7840b4dd3c9acde59048b.
//
// Solidity: event ProtocolFeeClaimed(address protocol, uint256 amount)
func (_AugurTurbo *AugurTurboFilterer) ParseProtocolFeeClaimed(log types.Log) (*AugurTurboProtocolFeeClaimed, error) {
	event := new(AugurTurboProtocolFeeClaimed)
	if err := _AugurTurbo.contract.UnpackLog(event, "ProtocolFeeClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurTurboSettlementFeeChangedIterator is returned from FilterSettlementFeeChanged and is used to iterate over the raw logs and unpacked data for SettlementFeeChanged events raised by the AugurTurbo contract.
type AugurTurboSettlementFeeChangedIterator struct {
	Event *AugurTurboSettlementFeeChanged // Event containing the contract specifics and raw log

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
func (it *AugurTurboSettlementFeeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTurboSettlementFeeChanged)
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
		it.Event = new(AugurTurboSettlementFeeChanged)
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
func (it *AugurTurboSettlementFeeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTurboSettlementFeeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTurboSettlementFeeChanged represents a SettlementFeeChanged event raised by the AugurTurbo contract.
type AugurTurboSettlementFeeChanged struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSettlementFeeChanged is a free log retrieval operation binding the contract event 0x92d395c429898992f8532ee7145901513e524c2085fd7fd1da39b8badcd6df31.
//
// Solidity: event SettlementFeeChanged(uint256 fee)
func (_AugurTurbo *AugurTurboFilterer) FilterSettlementFeeChanged(opts *bind.FilterOpts) (*AugurTurboSettlementFeeChangedIterator, error) {

	logs, sub, err := _AugurTurbo.contract.FilterLogs(opts, "SettlementFeeChanged")
	if err != nil {
		return nil, err
	}
	return &AugurTurboSettlementFeeChangedIterator{contract: _AugurTurbo.contract, event: "SettlementFeeChanged", logs: logs, sub: sub}, nil
}

// WatchSettlementFeeChanged is a free log subscription operation binding the contract event 0x92d395c429898992f8532ee7145901513e524c2085fd7fd1da39b8badcd6df31.
//
// Solidity: event SettlementFeeChanged(uint256 fee)
func (_AugurTurbo *AugurTurboFilterer) WatchSettlementFeeChanged(opts *bind.WatchOpts, sink chan<- *AugurTurboSettlementFeeChanged) (event.Subscription, error) {

	logs, sub, err := _AugurTurbo.contract.WatchLogs(opts, "SettlementFeeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTurboSettlementFeeChanged)
				if err := _AugurTurbo.contract.UnpackLog(event, "SettlementFeeChanged", log); err != nil {
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

// ParseSettlementFeeChanged is a log parse operation binding the contract event 0x92d395c429898992f8532ee7145901513e524c2085fd7fd1da39b8badcd6df31.
//
// Solidity: event SettlementFeeChanged(uint256 fee)
func (_AugurTurbo *AugurTurboFilterer) ParseSettlementFeeChanged(log types.Log) (*AugurTurboSettlementFeeChanged, error) {
	event := new(AugurTurboSettlementFeeChanged)
	if err := _AugurTurbo.contract.UnpackLog(event, "SettlementFeeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurTurboSettlementFeeClaimedIterator is returned from FilterSettlementFeeClaimed and is used to iterate over the raw logs and unpacked data for SettlementFeeClaimed events raised by the AugurTurbo contract.
type AugurTurboSettlementFeeClaimedIterator struct {
	Event *AugurTurboSettlementFeeClaimed // Event containing the contract specifics and raw log

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
func (it *AugurTurboSettlementFeeClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTurboSettlementFeeClaimed)
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
		it.Event = new(AugurTurboSettlementFeeClaimed)
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
func (it *AugurTurboSettlementFeeClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTurboSettlementFeeClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTurboSettlementFeeClaimed represents a SettlementFeeClaimed event raised by the AugurTurbo contract.
type AugurTurboSettlementFeeClaimed struct {
	SettlementAddress common.Address
	Amount            *big.Int
	Receiver          common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSettlementFeeClaimed is a free log retrieval operation binding the contract event 0xc9985ad824d943d66367ce5feea26e18979b3e1c9273742926d87e2b0d747387.
//
// Solidity: event SettlementFeeClaimed(address settlementAddress, uint256 amount, address indexed receiver)
func (_AugurTurbo *AugurTurboFilterer) FilterSettlementFeeClaimed(opts *bind.FilterOpts, receiver []common.Address) (*AugurTurboSettlementFeeClaimedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _AugurTurbo.contract.FilterLogs(opts, "SettlementFeeClaimed", receiverRule)
	if err != nil {
		return nil, err
	}
	return &AugurTurboSettlementFeeClaimedIterator{contract: _AugurTurbo.contract, event: "SettlementFeeClaimed", logs: logs, sub: sub}, nil
}

// WatchSettlementFeeClaimed is a free log subscription operation binding the contract event 0xc9985ad824d943d66367ce5feea26e18979b3e1c9273742926d87e2b0d747387.
//
// Solidity: event SettlementFeeClaimed(address settlementAddress, uint256 amount, address indexed receiver)
func (_AugurTurbo *AugurTurboFilterer) WatchSettlementFeeClaimed(opts *bind.WatchOpts, sink chan<- *AugurTurboSettlementFeeClaimed, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _AugurTurbo.contract.WatchLogs(opts, "SettlementFeeClaimed", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTurboSettlementFeeClaimed)
				if err := _AugurTurbo.contract.UnpackLog(event, "SettlementFeeClaimed", log); err != nil {
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

// ParseSettlementFeeClaimed is a log parse operation binding the contract event 0xc9985ad824d943d66367ce5feea26e18979b3e1c9273742926d87e2b0d747387.
//
// Solidity: event SettlementFeeClaimed(address settlementAddress, uint256 amount, address indexed receiver)
func (_AugurTurbo *AugurTurboFilterer) ParseSettlementFeeClaimed(log types.Log) (*AugurTurboSettlementFeeClaimed, error) {
	event := new(AugurTurboSettlementFeeClaimed)
	if err := _AugurTurbo.contract.UnpackLog(event, "SettlementFeeClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurTurboSharesBurnedIterator is returned from FilterSharesBurned and is used to iterate over the raw logs and unpacked data for SharesBurned events raised by the AugurTurbo contract.
type AugurTurboSharesBurnedIterator struct {
	Event *AugurTurboSharesBurned // Event containing the contract specifics and raw log

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
func (it *AugurTurboSharesBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTurboSharesBurned)
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
		it.Event = new(AugurTurboSharesBurned)
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
func (it *AugurTurboSharesBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTurboSharesBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTurboSharesBurned represents a SharesBurned event raised by the AugurTurbo contract.
type AugurTurboSharesBurned struct {
	Id       *big.Int
	Amount   *big.Int
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSharesBurned is a free log retrieval operation binding the contract event 0xb6fdb729b2ed801daf629f0ab713e4a7a73619505790f6f27fd92d6f2c9688d7.
//
// Solidity: event SharesBurned(uint256 id, uint256 amount, address receiver)
func (_AugurTurbo *AugurTurboFilterer) FilterSharesBurned(opts *bind.FilterOpts) (*AugurTurboSharesBurnedIterator, error) {

	logs, sub, err := _AugurTurbo.contract.FilterLogs(opts, "SharesBurned")
	if err != nil {
		return nil, err
	}
	return &AugurTurboSharesBurnedIterator{contract: _AugurTurbo.contract, event: "SharesBurned", logs: logs, sub: sub}, nil
}

// WatchSharesBurned is a free log subscription operation binding the contract event 0xb6fdb729b2ed801daf629f0ab713e4a7a73619505790f6f27fd92d6f2c9688d7.
//
// Solidity: event SharesBurned(uint256 id, uint256 amount, address receiver)
func (_AugurTurbo *AugurTurboFilterer) WatchSharesBurned(opts *bind.WatchOpts, sink chan<- *AugurTurboSharesBurned) (event.Subscription, error) {

	logs, sub, err := _AugurTurbo.contract.WatchLogs(opts, "SharesBurned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTurboSharesBurned)
				if err := _AugurTurbo.contract.UnpackLog(event, "SharesBurned", log); err != nil {
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

// ParseSharesBurned is a log parse operation binding the contract event 0xb6fdb729b2ed801daf629f0ab713e4a7a73619505790f6f27fd92d6f2c9688d7.
//
// Solidity: event SharesBurned(uint256 id, uint256 amount, address receiver)
func (_AugurTurbo *AugurTurboFilterer) ParseSharesBurned(log types.Log) (*AugurTurboSharesBurned, error) {
	event := new(AugurTurboSharesBurned)
	if err := _AugurTurbo.contract.UnpackLog(event, "SharesBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurTurboSharesMintedIterator is returned from FilterSharesMinted and is used to iterate over the raw logs and unpacked data for SharesMinted events raised by the AugurTurbo contract.
type AugurTurboSharesMintedIterator struct {
	Event *AugurTurboSharesMinted // Event containing the contract specifics and raw log

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
func (it *AugurTurboSharesMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTurboSharesMinted)
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
		it.Event = new(AugurTurboSharesMinted)
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
func (it *AugurTurboSharesMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTurboSharesMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTurboSharesMinted represents a SharesMinted event raised by the AugurTurbo contract.
type AugurTurboSharesMinted struct {
	Id       *big.Int
	Amount   *big.Int
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSharesMinted is a free log retrieval operation binding the contract event 0xd81c0442e10068a9818f3aa093c9ccb804584690df572d7df3da2d892a6973f2.
//
// Solidity: event SharesMinted(uint256 id, uint256 amount, address receiver)
func (_AugurTurbo *AugurTurboFilterer) FilterSharesMinted(opts *bind.FilterOpts) (*AugurTurboSharesMintedIterator, error) {

	logs, sub, err := _AugurTurbo.contract.FilterLogs(opts, "SharesMinted")
	if err != nil {
		return nil, err
	}
	return &AugurTurboSharesMintedIterator{contract: _AugurTurbo.contract, event: "SharesMinted", logs: logs, sub: sub}, nil
}

// WatchSharesMinted is a free log subscription operation binding the contract event 0xd81c0442e10068a9818f3aa093c9ccb804584690df572d7df3da2d892a6973f2.
//
// Solidity: event SharesMinted(uint256 id, uint256 amount, address receiver)
func (_AugurTurbo *AugurTurboFilterer) WatchSharesMinted(opts *bind.WatchOpts, sink chan<- *AugurTurboSharesMinted) (event.Subscription, error) {

	logs, sub, err := _AugurTurbo.contract.WatchLogs(opts, "SharesMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTurboSharesMinted)
				if err := _AugurTurbo.contract.UnpackLog(event, "SharesMinted", log); err != nil {
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

// ParseSharesMinted is a log parse operation binding the contract event 0xd81c0442e10068a9818f3aa093c9ccb804584690df572d7df3da2d892a6973f2.
//
// Solidity: event SharesMinted(uint256 id, uint256 amount, address receiver)
func (_AugurTurbo *AugurTurboFilterer) ParseSharesMinted(log types.Log) (*AugurTurboSharesMinted, error) {
	event := new(AugurTurboSharesMinted)
	if err := _AugurTurbo.contract.UnpackLog(event, "SharesMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurTurboSharesSwappedIterator is returned from FilterSharesSwapped and is used to iterate over the raw logs and unpacked data for SharesSwapped events raised by the AugurTurbo contract.
type AugurTurboSharesSwappedIterator struct {
	Event *AugurTurboSharesSwapped // Event containing the contract specifics and raw log

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
func (it *AugurTurboSharesSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTurboSharesSwapped)
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
		it.Event = new(AugurTurboSharesSwapped)
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
func (it *AugurTurboSharesSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTurboSharesSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTurboSharesSwapped represents a SharesSwapped event raised by the AugurTurbo contract.
type AugurTurboSharesSwapped struct {
	MarketFactory common.Address
	MarketId      *big.Int
	User          common.Address
	Outcome       *big.Int
	Collateral    *big.Int
	Shares        *big.Int
	Price         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSharesSwapped is a free log retrieval operation binding the contract event 0xec2a60d57293d00dfe68ab5f1d18738c4600ce39c0c0c623fc086814615f33fa.
//
// Solidity: event SharesSwapped(address indexed marketFactory, uint256 indexed marketId, address indexed user, uint256 outcome, int256 collateral, int256 shares, uint256 price)
func (_AugurTurbo *AugurTurboFilterer) FilterSharesSwapped(opts *bind.FilterOpts, marketFactory []common.Address, marketId []*big.Int, user []common.Address) (*AugurTurboSharesSwappedIterator, error) {

	var marketFactoryRule []interface{}
	for _, marketFactoryItem := range marketFactory {
		marketFactoryRule = append(marketFactoryRule, marketFactoryItem)
	}
	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AugurTurbo.contract.FilterLogs(opts, "SharesSwapped", marketFactoryRule, marketIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AugurTurboSharesSwappedIterator{contract: _AugurTurbo.contract, event: "SharesSwapped", logs: logs, sub: sub}, nil
}

// WatchSharesSwapped is a free log subscription operation binding the contract event 0xec2a60d57293d00dfe68ab5f1d18738c4600ce39c0c0c623fc086814615f33fa.
//
// Solidity: event SharesSwapped(address indexed marketFactory, uint256 indexed marketId, address indexed user, uint256 outcome, int256 collateral, int256 shares, uint256 price)
func (_AugurTurbo *AugurTurboFilterer) WatchSharesSwapped(opts *bind.WatchOpts, sink chan<- *AugurTurboSharesSwapped, marketFactory []common.Address, marketId []*big.Int, user []common.Address) (event.Subscription, error) {

	var marketFactoryRule []interface{}
	for _, marketFactoryItem := range marketFactory {
		marketFactoryRule = append(marketFactoryRule, marketFactoryItem)
	}
	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AugurTurbo.contract.WatchLogs(opts, "SharesSwapped", marketFactoryRule, marketIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTurboSharesSwapped)
				if err := _AugurTurbo.contract.UnpackLog(event, "SharesSwapped", log); err != nil {
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

// ParseSharesSwapped is a log parse operation binding the contract event 0xec2a60d57293d00dfe68ab5f1d18738c4600ce39c0c0c623fc086814615f33fa.
//
// Solidity: event SharesSwapped(address indexed marketFactory, uint256 indexed marketId, address indexed user, uint256 outcome, int256 collateral, int256 shares, uint256 price)
func (_AugurTurbo *AugurTurboFilterer) ParseSharesSwapped(log types.Log) (*AugurTurboSharesSwapped, error) {
	event := new(AugurTurboSharesSwapped)
	if err := _AugurTurbo.contract.UnpackLog(event, "SharesSwapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurTurboSportsMarketCreatedIterator is returned from FilterSportsMarketCreated and is used to iterate over the raw logs and unpacked data for SportsMarketCreated events raised by the AugurTurbo contract.
type AugurTurboSportsMarketCreatedIterator struct {
	Event *AugurTurboSportsMarketCreated // Event containing the contract specifics and raw log

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
func (it *AugurTurboSportsMarketCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTurboSportsMarketCreated)
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
		it.Event = new(AugurTurboSportsMarketCreated)
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
func (it *AugurTurboSportsMarketCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTurboSportsMarketCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTurboSportsMarketCreated represents a SportsMarketCreated event raised by the AugurTurbo contract.
type AugurTurboSportsMarketCreated struct {
	Id                *big.Int
	Creator           common.Address
	EndTime           *big.Int
	MarketType        uint8
	EventId           *big.Int
	HomeTeamId        *big.Int
	AwayTeamId        *big.Int
	EstimatedStarTime *big.Int
	Score             *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSportsMarketCreated is a free log retrieval operation binding the contract event 0x259259419ab9dc829d60f11fbcf41e81cf76ddcd13da470c3a3a11b70460bdb1.
//
// Solidity: event SportsMarketCreated(uint256 id, address creator, uint256 endTime, uint8 marketType, uint256 indexed eventId, uint256 homeTeamId, uint256 awayTeamId, uint256 estimatedStarTime, int256 score)
func (_AugurTurbo *AugurTurboFilterer) FilterSportsMarketCreated(opts *bind.FilterOpts, eventId []*big.Int) (*AugurTurboSportsMarketCreatedIterator, error) {

	var eventIdRule []interface{}
	for _, eventIdItem := range eventId {
		eventIdRule = append(eventIdRule, eventIdItem)
	}

	logs, sub, err := _AugurTurbo.contract.FilterLogs(opts, "SportsMarketCreated", eventIdRule)
	if err != nil {
		return nil, err
	}
	return &AugurTurboSportsMarketCreatedIterator{contract: _AugurTurbo.contract, event: "SportsMarketCreated", logs: logs, sub: sub}, nil
}

// WatchSportsMarketCreated is a free log subscription operation binding the contract event 0x259259419ab9dc829d60f11fbcf41e81cf76ddcd13da470c3a3a11b70460bdb1.
//
// Solidity: event SportsMarketCreated(uint256 id, address creator, uint256 endTime, uint8 marketType, uint256 indexed eventId, uint256 homeTeamId, uint256 awayTeamId, uint256 estimatedStarTime, int256 score)
func (_AugurTurbo *AugurTurboFilterer) WatchSportsMarketCreated(opts *bind.WatchOpts, sink chan<- *AugurTurboSportsMarketCreated, eventId []*big.Int) (event.Subscription, error) {

	var eventIdRule []interface{}
	for _, eventIdItem := range eventId {
		eventIdRule = append(eventIdRule, eventIdItem)
	}

	logs, sub, err := _AugurTurbo.contract.WatchLogs(opts, "SportsMarketCreated", eventIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTurboSportsMarketCreated)
				if err := _AugurTurbo.contract.UnpackLog(event, "SportsMarketCreated", log); err != nil {
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

// ParseSportsMarketCreated is a log parse operation binding the contract event 0x259259419ab9dc829d60f11fbcf41e81cf76ddcd13da470c3a3a11b70460bdb1.
//
// Solidity: event SportsMarketCreated(uint256 id, address creator, uint256 endTime, uint8 marketType, uint256 indexed eventId, uint256 homeTeamId, uint256 awayTeamId, uint256 estimatedStarTime, int256 score)
func (_AugurTurbo *AugurTurboFilterer) ParseSportsMarketCreated(log types.Log) (*AugurTurboSportsMarketCreated, error) {
	event := new(AugurTurboSportsMarketCreated)
	if err := _AugurTurbo.contract.UnpackLog(event, "SportsMarketCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurTurboTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the AugurTurbo contract.
type AugurTurboTransferIterator struct {
	Event *AugurTurboTransfer // Event containing the contract specifics and raw log

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
func (it *AugurTurboTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTurboTransfer)
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
		it.Event = new(AugurTurboTransfer)
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
func (it *AugurTurboTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTurboTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTurboTransfer represents a Transfer event raised by the AugurTurbo contract.
type AugurTurboTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AugurTurbo *AugurTurboFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AugurTurboTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AugurTurbo.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AugurTurboTransferIterator{contract: _AugurTurbo.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AugurTurbo *AugurTurboFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AugurTurboTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AugurTurbo.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTurboTransfer)
				if err := _AugurTurbo.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_AugurTurbo *AugurTurboFilterer) ParseTransfer(log types.Log) (*AugurTurboTransfer, error) {
	event := new(AugurTurboTransfer)
	if err := _AugurTurbo.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurTurboTrustedMarketCreatedIterator is returned from FilterTrustedMarketCreated and is used to iterate over the raw logs and unpacked data for TrustedMarketCreated events raised by the AugurTurbo contract.
type AugurTurboTrustedMarketCreatedIterator struct {
	Event *AugurTurboTrustedMarketCreated // Event containing the contract specifics and raw log

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
func (it *AugurTurboTrustedMarketCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTurboTrustedMarketCreated)
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
		it.Event = new(AugurTurboTrustedMarketCreated)
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
func (it *AugurTurboTrustedMarketCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTurboTrustedMarketCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTurboTrustedMarketCreated represents a TrustedMarketCreated event raised by the AugurTurbo contract.
type AugurTurboTrustedMarketCreated struct {
	Id          *big.Int
	Creator     common.Address
	EndTime     *big.Int
	Description string
	Outcomes    []string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTrustedMarketCreated is a free log retrieval operation binding the contract event 0x7b73997b0a6a20b511ffa2c5421a164febd7014b1f89a989f32998e4276d2a5e.
//
// Solidity: event TrustedMarketCreated(uint256 id, address creator, uint256 _endTime, string description, string[] outcomes)
func (_AugurTurbo *AugurTurboFilterer) FilterTrustedMarketCreated(opts *bind.FilterOpts) (*AugurTurboTrustedMarketCreatedIterator, error) {

	logs, sub, err := _AugurTurbo.contract.FilterLogs(opts, "TrustedMarketCreated")
	if err != nil {
		return nil, err
	}
	return &AugurTurboTrustedMarketCreatedIterator{contract: _AugurTurbo.contract, event: "TrustedMarketCreated", logs: logs, sub: sub}, nil
}

// WatchTrustedMarketCreated is a free log subscription operation binding the contract event 0x7b73997b0a6a20b511ffa2c5421a164febd7014b1f89a989f32998e4276d2a5e.
//
// Solidity: event TrustedMarketCreated(uint256 id, address creator, uint256 _endTime, string description, string[] outcomes)
func (_AugurTurbo *AugurTurboFilterer) WatchTrustedMarketCreated(opts *bind.WatchOpts, sink chan<- *AugurTurboTrustedMarketCreated) (event.Subscription, error) {

	logs, sub, err := _AugurTurbo.contract.WatchLogs(opts, "TrustedMarketCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTurboTrustedMarketCreated)
				if err := _AugurTurbo.contract.UnpackLog(event, "TrustedMarketCreated", log); err != nil {
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

// ParseTrustedMarketCreated is a log parse operation binding the contract event 0x7b73997b0a6a20b511ffa2c5421a164febd7014b1f89a989f32998e4276d2a5e.
//
// Solidity: event TrustedMarketCreated(uint256 id, address creator, uint256 _endTime, string description, string[] outcomes)
func (_AugurTurbo *AugurTurboFilterer) ParseTrustedMarketCreated(log types.Log) (*AugurTurboTrustedMarketCreated, error) {
	event := new(AugurTurboTrustedMarketCreated)
	if err := _AugurTurbo.contract.UnpackLog(event, "TrustedMarketCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AugurTurboWinningsClaimedIterator is returned from FilterWinningsClaimed and is used to iterate over the raw logs and unpacked data for WinningsClaimed events raised by the AugurTurbo contract.
type AugurTurboWinningsClaimedIterator struct {
	Event *AugurTurboWinningsClaimed // Event containing the contract specifics and raw log

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
func (it *AugurTurboWinningsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AugurTurboWinningsClaimed)
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
		it.Event = new(AugurTurboWinningsClaimed)
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
func (it *AugurTurboWinningsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AugurTurboWinningsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AugurTurboWinningsClaimed represents a WinningsClaimed event raised by the AugurTurbo contract.
type AugurTurboWinningsClaimed struct {
	Id             *big.Int
	WinningOutcome common.Address
	Amount         *big.Int
	SettlementFee  *big.Int
	Payout         *big.Int
	Receiver       common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWinningsClaimed is a free log retrieval operation binding the contract event 0xe67bd0100cd3289557430d36098901ba18161e6279c9711d8650b8af10552104.
//
// Solidity: event WinningsClaimed(uint256 id, address winningOutcome, uint256 amount, uint256 settlementFee, uint256 payout, address indexed receiver)
func (_AugurTurbo *AugurTurboFilterer) FilterWinningsClaimed(opts *bind.FilterOpts, receiver []common.Address) (*AugurTurboWinningsClaimedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _AugurTurbo.contract.FilterLogs(opts, "WinningsClaimed", receiverRule)
	if err != nil {
		return nil, err
	}
	return &AugurTurboWinningsClaimedIterator{contract: _AugurTurbo.contract, event: "WinningsClaimed", logs: logs, sub: sub}, nil
}

// WatchWinningsClaimed is a free log subscription operation binding the contract event 0xe67bd0100cd3289557430d36098901ba18161e6279c9711d8650b8af10552104.
//
// Solidity: event WinningsClaimed(uint256 id, address winningOutcome, uint256 amount, uint256 settlementFee, uint256 payout, address indexed receiver)
func (_AugurTurbo *AugurTurboFilterer) WatchWinningsClaimed(opts *bind.WatchOpts, sink chan<- *AugurTurboWinningsClaimed, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _AugurTurbo.contract.WatchLogs(opts, "WinningsClaimed", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AugurTurboWinningsClaimed)
				if err := _AugurTurbo.contract.UnpackLog(event, "WinningsClaimed", log); err != nil {
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

// ParseWinningsClaimed is a log parse operation binding the contract event 0xe67bd0100cd3289557430d36098901ba18161e6279c9711d8650b8af10552104.
//
// Solidity: event WinningsClaimed(uint256 id, address winningOutcome, uint256 amount, uint256 settlementFee, uint256 payout, address indexed receiver)
func (_AugurTurbo *AugurTurboFilterer) ParseWinningsClaimed(log types.Log) (*AugurTurboWinningsClaimed, error) {
	event := new(AugurTurboWinningsClaimed)
	if err := _AugurTurbo.contract.UnpackLog(event, "WinningsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
