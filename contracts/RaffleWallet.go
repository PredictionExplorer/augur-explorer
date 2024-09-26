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
	_ = abi.ConvertType
)

// IRaffleWalletMetaData contains all meta data concerning the IRaffleWallet contract.
var IRaffleWalletMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RaffleDepositEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RaffleWithdrawalEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IRaffleWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use IRaffleWalletMetaData.ABI instead.
var IRaffleWalletABI = IRaffleWalletMetaData.ABI

// IRaffleWallet is an auto generated Go binding around an Ethereum contract.
type IRaffleWallet struct {
	IRaffleWalletCaller     // Read-only binding to the contract
	IRaffleWalletTransactor // Write-only binding to the contract
	IRaffleWalletFilterer   // Log filterer for contract events
}

// IRaffleWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRaffleWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRaffleWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRaffleWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRaffleWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRaffleWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRaffleWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRaffleWalletSession struct {
	Contract     *IRaffleWallet    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRaffleWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRaffleWalletCallerSession struct {
	Contract *IRaffleWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IRaffleWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRaffleWalletTransactorSession struct {
	Contract     *IRaffleWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IRaffleWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRaffleWalletRaw struct {
	Contract *IRaffleWallet // Generic contract binding to access the raw methods on
}

// IRaffleWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRaffleWalletCallerRaw struct {
	Contract *IRaffleWalletCaller // Generic read-only contract binding to access the raw methods on
}

// IRaffleWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRaffleWalletTransactorRaw struct {
	Contract *IRaffleWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRaffleWallet creates a new instance of IRaffleWallet, bound to a specific deployed contract.
func NewIRaffleWallet(address common.Address, backend bind.ContractBackend) (*IRaffleWallet, error) {
	contract, err := bindIRaffleWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRaffleWallet{IRaffleWalletCaller: IRaffleWalletCaller{contract: contract}, IRaffleWalletTransactor: IRaffleWalletTransactor{contract: contract}, IRaffleWalletFilterer: IRaffleWalletFilterer{contract: contract}}, nil
}

// NewIRaffleWalletCaller creates a new read-only instance of IRaffleWallet, bound to a specific deployed contract.
func NewIRaffleWalletCaller(address common.Address, caller bind.ContractCaller) (*IRaffleWalletCaller, error) {
	contract, err := bindIRaffleWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRaffleWalletCaller{contract: contract}, nil
}

// NewIRaffleWalletTransactor creates a new write-only instance of IRaffleWallet, bound to a specific deployed contract.
func NewIRaffleWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*IRaffleWalletTransactor, error) {
	contract, err := bindIRaffleWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRaffleWalletTransactor{contract: contract}, nil
}

// NewIRaffleWalletFilterer creates a new log filterer instance of IRaffleWallet, bound to a specific deployed contract.
func NewIRaffleWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*IRaffleWalletFilterer, error) {
	contract, err := bindIRaffleWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRaffleWalletFilterer{contract: contract}, nil
}

// bindIRaffleWallet binds a generic wrapper to an already deployed contract.
func bindIRaffleWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IRaffleWalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRaffleWallet *IRaffleWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRaffleWallet.Contract.IRaffleWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRaffleWallet *IRaffleWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRaffleWallet.Contract.IRaffleWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRaffleWallet *IRaffleWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRaffleWallet.Contract.IRaffleWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRaffleWallet *IRaffleWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRaffleWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRaffleWallet *IRaffleWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRaffleWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRaffleWallet *IRaffleWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRaffleWallet.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address winner) payable returns()
func (_IRaffleWallet *IRaffleWalletTransactor) Deposit(opts *bind.TransactOpts, winner common.Address) (*types.Transaction, error) {
	return _IRaffleWallet.contract.Transact(opts, "deposit", winner)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address winner) payable returns()
func (_IRaffleWallet *IRaffleWalletSession) Deposit(winner common.Address) (*types.Transaction, error) {
	return _IRaffleWallet.Contract.Deposit(&_IRaffleWallet.TransactOpts, winner)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address winner) payable returns()
func (_IRaffleWallet *IRaffleWalletTransactorSession) Deposit(winner common.Address) (*types.Transaction, error) {
	return _IRaffleWallet.Contract.Deposit(&_IRaffleWallet.TransactOpts, winner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_IRaffleWallet *IRaffleWalletTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRaffleWallet.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_IRaffleWallet *IRaffleWalletSession) Withdraw() (*types.Transaction, error) {
	return _IRaffleWallet.Contract.Withdraw(&_IRaffleWallet.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_IRaffleWallet *IRaffleWalletTransactorSession) Withdraw() (*types.Transaction, error) {
	return _IRaffleWallet.Contract.Withdraw(&_IRaffleWallet.TransactOpts)
}

// IRaffleWalletRaffleDepositEventIterator is returned from FilterRaffleDepositEvent and is used to iterate over the raw logs and unpacked data for RaffleDepositEvent events raised by the IRaffleWallet contract.
type IRaffleWalletRaffleDepositEventIterator struct {
	Event *IRaffleWalletRaffleDepositEvent // Event containing the contract specifics and raw log

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
func (it *IRaffleWalletRaffleDepositEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRaffleWalletRaffleDepositEvent)
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
		it.Event = new(IRaffleWalletRaffleDepositEvent)
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
func (it *IRaffleWalletRaffleDepositEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRaffleWalletRaffleDepositEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRaffleWalletRaffleDepositEvent represents a RaffleDepositEvent event raised by the IRaffleWallet contract.
type IRaffleWalletRaffleDepositEvent struct {
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRaffleDepositEvent is a free log retrieval operation binding the contract event 0xcf6f6dcf9306290f700abd5d57c300a4ad7bf0d4086d5a5e88040fcecb0fef90.
//
// Solidity: event RaffleDepositEvent(address indexed winner, uint256 amount)
func (_IRaffleWallet *IRaffleWalletFilterer) FilterRaffleDepositEvent(opts *bind.FilterOpts, winner []common.Address) (*IRaffleWalletRaffleDepositEventIterator, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _IRaffleWallet.contract.FilterLogs(opts, "RaffleDepositEvent", winnerRule)
	if err != nil {
		return nil, err
	}
	return &IRaffleWalletRaffleDepositEventIterator{contract: _IRaffleWallet.contract, event: "RaffleDepositEvent", logs: logs, sub: sub}, nil
}

// WatchRaffleDepositEvent is a free log subscription operation binding the contract event 0xcf6f6dcf9306290f700abd5d57c300a4ad7bf0d4086d5a5e88040fcecb0fef90.
//
// Solidity: event RaffleDepositEvent(address indexed winner, uint256 amount)
func (_IRaffleWallet *IRaffleWalletFilterer) WatchRaffleDepositEvent(opts *bind.WatchOpts, sink chan<- *IRaffleWalletRaffleDepositEvent, winner []common.Address) (event.Subscription, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _IRaffleWallet.contract.WatchLogs(opts, "RaffleDepositEvent", winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRaffleWalletRaffleDepositEvent)
				if err := _IRaffleWallet.contract.UnpackLog(event, "RaffleDepositEvent", log); err != nil {
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

// ParseRaffleDepositEvent is a log parse operation binding the contract event 0xcf6f6dcf9306290f700abd5d57c300a4ad7bf0d4086d5a5e88040fcecb0fef90.
//
// Solidity: event RaffleDepositEvent(address indexed winner, uint256 amount)
func (_IRaffleWallet *IRaffleWalletFilterer) ParseRaffleDepositEvent(log types.Log) (*IRaffleWalletRaffleDepositEvent, error) {
	event := new(IRaffleWalletRaffleDepositEvent)
	if err := _IRaffleWallet.contract.UnpackLog(event, "RaffleDepositEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRaffleWalletRaffleWithdrawalEventIterator is returned from FilterRaffleWithdrawalEvent and is used to iterate over the raw logs and unpacked data for RaffleWithdrawalEvent events raised by the IRaffleWallet contract.
type IRaffleWalletRaffleWithdrawalEventIterator struct {
	Event *IRaffleWalletRaffleWithdrawalEvent // Event containing the contract specifics and raw log

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
func (it *IRaffleWalletRaffleWithdrawalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRaffleWalletRaffleWithdrawalEvent)
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
		it.Event = new(IRaffleWalletRaffleWithdrawalEvent)
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
func (it *IRaffleWalletRaffleWithdrawalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRaffleWalletRaffleWithdrawalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRaffleWalletRaffleWithdrawalEvent represents a RaffleWithdrawalEvent event raised by the IRaffleWallet contract.
type IRaffleWalletRaffleWithdrawalEvent struct {
	Destination common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRaffleWithdrawalEvent is a free log retrieval operation binding the contract event 0x49775450da95afb818c4651b894f124d05629f06572a18900bf29db74b04a0d8.
//
// Solidity: event RaffleWithdrawalEvent(address indexed destination, uint256 amount)
func (_IRaffleWallet *IRaffleWalletFilterer) FilterRaffleWithdrawalEvent(opts *bind.FilterOpts, destination []common.Address) (*IRaffleWalletRaffleWithdrawalEventIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _IRaffleWallet.contract.FilterLogs(opts, "RaffleWithdrawalEvent", destinationRule)
	if err != nil {
		return nil, err
	}
	return &IRaffleWalletRaffleWithdrawalEventIterator{contract: _IRaffleWallet.contract, event: "RaffleWithdrawalEvent", logs: logs, sub: sub}, nil
}

// WatchRaffleWithdrawalEvent is a free log subscription operation binding the contract event 0x49775450da95afb818c4651b894f124d05629f06572a18900bf29db74b04a0d8.
//
// Solidity: event RaffleWithdrawalEvent(address indexed destination, uint256 amount)
func (_IRaffleWallet *IRaffleWalletFilterer) WatchRaffleWithdrawalEvent(opts *bind.WatchOpts, sink chan<- *IRaffleWalletRaffleWithdrawalEvent, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _IRaffleWallet.contract.WatchLogs(opts, "RaffleWithdrawalEvent", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRaffleWalletRaffleWithdrawalEvent)
				if err := _IRaffleWallet.contract.UnpackLog(event, "RaffleWithdrawalEvent", log); err != nil {
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

// ParseRaffleWithdrawalEvent is a log parse operation binding the contract event 0x49775450da95afb818c4651b894f124d05629f06572a18900bf29db74b04a0d8.
//
// Solidity: event RaffleWithdrawalEvent(address indexed destination, uint256 amount)
func (_IRaffleWallet *IRaffleWalletFilterer) ParseRaffleWithdrawalEvent(log types.Log) (*IRaffleWalletRaffleWithdrawalEvent, error) {
	event := new(IRaffleWalletRaffleWithdrawalEvent)
	if err := _IRaffleWallet.contract.UnpackLog(event, "RaffleWithdrawalEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaffleWalletMetaData contains all meta data concerning the RaffleWallet contract.
var RaffleWalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"game_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"DepositFromUnauthorizedSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"FundTransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"NonZeroValueRequired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroBalance\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RaffleDepositEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RaffleWithdrawalEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"game\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523461002f576100196100146100f4565b610246565b610021610034565b610bfb6103968239610bfb90f35b61003a565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906100669061003e565b810190811060018060401b0382111761007e57604052565b610048565b9061009661008f610034565b928361005c565b565b5f80fd5b60018060a01b031690565b6100b09061009c565b90565b6100bc816100a7565b036100c357565b5f80fd5b905051906100d4826100b3565b565b906020828203126100ef576100ec915f016100c7565b90565b610098565b610112610f918038038061010781610083565b9283398101906100d6565b90565b90565b90565b61012f61012a61013492610115565b610118565b61009c565b90565b6101409061011b565b90565b60209181520190565b5f7f5a65726f2d616464726573732077617320676976656e2e000000000000000000910152565b6101806017602092610143565b6101898161014c565b0190565b6101a29060208101905f818303910152610173565b90565b156101ac57565b6101b4610034565b63eac0d38960e01b8152806101cb6004820161018d565b0390fd5b5f1b90565b906101e560018060a01b03916101cf565b9181191691161790565b6102036101fe6102089261009c565b610118565b61009c565b90565b610214906101ef565b90565b6102209061020b565b90565b90565b9061023b61023661024292610217565b610223565b82546101d4565b9055565b61027f9061025333610281565b6102788161027161026b6102665f610137565b6100a7565b916100a7565b14156101a5565b6001610226565b565b61028a906102ae565b565b610295906100a7565b9052565b91906102ac905f6020850194019061028c565b565b806102c96102c36102be5f610137565b6100a7565b916100a7565b146102d9576102d790610336565b565b6102fc6102e55f610137565b5f918291631e4fbdf760e01b835260048301610299565b0390fd5b5f1c90565b60018060a01b031690565b61031c61032191610300565b610305565b90565b61032e9054610310565b90565b5f0190565b61033f5f610324565b610349825f610226565b9061037d6103777f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e093610217565b91610217565b91610386610034565b8061039081610331565b0390a356fe60806040526004361015610013575b610387565b61001d5f3561008c565b806327e235e3146100875780633ccfd60b14610082578063715018a61461007d5780638da5cb5b14610078578063c3fe3e2814610073578063f2fde38b1461006e5763f340fa010361000e5761035e565b61032b565b6102f6565b610281565b61022c565b6101f9565b6101b0565b60e01c90565b60405190565b5f80fd5b5f80fd5b60018060a01b031690565b6100b4906100a0565b90565b6100c0816100ab565b036100c757565b5f80fd5b905035906100d8826100b7565b565b906020828203126100f3576100f0915f016100cb565b90565b61009c565b90565b61010f61010a610114926100a0565b6100f8565b6100a0565b90565b610120906100fb565b90565b61012c90610117565b90565b9061013990610123565b5f5260205260405f2090565b1c90565b90565b61015c9060086101619302610145565b610149565b90565b9061016f915461014c565b90565b610188906101836002915f9261012f565b610164565b90565b90565b6101979061018b565b9052565b91906101ae905f6020850194019061018e565b565b346101e0576101dc6101cb6101c63660046100da565b610172565b6101d3610092565b9182918261019b565b0390f35b610098565b5f9103126101ef57565b61009c565b5f0190565b34610227576102093660046101e5565b61021161063c565b610219610092565b80610223816101f4565b0390f35b610098565b3461025a5761023c3660046101e5565b61024461074a565b61024c610092565b80610256816101f4565b0390f35b610098565b610268906100ab565b9052565b919061027f905f6020850194019061025f565b565b346102b1576102913660046101e5565b6102ad61029c610779565b6102a4610092565b9182918261026c565b0390f35b610098565b60018060a01b031690565b6102d19060086102d69302610145565b6102b6565b90565b906102e491546102c1565b90565b6102f360015f906102d9565b90565b34610326576103063660046101e5565b6103226103116102e7565b610319610092565b9182918261026c565b0390f35b610098565b346103595761034361033e3660046100da565b6107f3565b61034b610092565b80610355816101f4565b0390f35b610098565b61037161036c3660046100da565b6109f7565b610379610092565b80610383816101f4565b0390f35b5f80fd5b5f1c90565b61039c6103a19161038b565b610149565b90565b6103ae9054610390565b90565b90565b6103c86103c36103cd926103b1565b6100f8565b61018b565b90565b60209181520190565b5f7f596f75722062616c616e636520697320302e0000000000000000000000000000910152565b61040d60126020926103d0565b610416816103d9565b0190565b61042f9060208101905f818303910152610400565b90565b1561043957565b610441610092565b63cb7450f760e01b8152806104586004820161041a565b0390fd5b5f1b90565b9061046d5f199161045c565b9181191691161790565b61048b6104866104909261018b565b6100f8565b61018b565b90565b90565b906104ab6104a66104b292610477565b610493565b8254610461565b9055565b905090565b6104c65f80926104b6565b0190565b6104d3906104bb565b90565b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906104fe906104d6565b810190811067ffffffffffffffff82111761051857604052565b6104e0565b90610530610529610092565b92836104f4565b565b67ffffffffffffffff81116105505761054c6020916104d6565b0190565b6104e0565b9061056761056283610532565b61051d565b918252565b606090565b3d5f1461058c576105813d610555565b903d5f602084013e5b565b61059461056c565b9061058a565b5f7f5472616e73666572206661696c65642e00000000000000000000000000000000910152565b6105ce60106020926103d0565b6105d78161059a565b0190565b91604061060c9294936106056105fa606083018381035f8501526105c1565b96602083019061018e565b019061025f565b565b15610617575050565b610638610622610092565b92839263310a0fbb60e21b8452600484016105db565b0390fd5b61065061064b6002339061012f565b6103a4565b61066c816106666106605f6103b4565b9161018b565b11610432565b6106896106785f6103b4565b6106846002339061012f565b610496565b6106b85f803384610698610092565b90816106a3816104ca565b03925af16106af610571565b5082339161060e565b336106f86106e67f49775450da95afb818c4651b894f124d05629f06572a18900bf29db74b04a0d892610123565b926106ef610092565b9182918261019b565b0390a2565b610705610acd565b61070d610737565b565b61072361071e610728926103b1565b6100f8565b6100a0565b90565b6107349061070f565b90565b6107486107435f61072b565b610b59565b565b6107526106fd565b565b5f90565b6107646107699161038b565b6102b6565b90565b6107769054610758565b90565b610781610754565b5061078b5f61076c565b90565b61079f9061079a610acd565b6107a1565b565b806107bc6107b66107b15f61072b565b6100ab565b916100ab565b146107cc576107ca90610b59565b565b6107ef6107d85f61072b565b5f918291631e4fbdf760e01b83526004830161026c565b0390fd5b6107fc9061078e565b565b5f7f5a65726f2d616464726573732077617320676976656e2e000000000000000000910152565b61083260176020926103d0565b61083b816107fe565b0190565b6108549060208101905f818303910152610825565b90565b1561085e57565b610866610092565b63eac0d38960e01b81528061087d6004820161083f565b0390fd5b5f7f4e6f2045544820686173206265656e2073656e742e0000000000000000000000910152565b6108b560156020926103d0565b6108be81610881565b0190565b6108d79060208101905f8183039101526108a8565b90565b156108e157565b6108e9610092565b63af33979960e01b815280610900600482016108c2565b0390fd5b60207f706f7369742e0000000000000000000000000000000000000000000000000000917f4f6e6c7920436f736d696347616d6520697320616c6c6f77656420746f2064655f8201520152565b61095e60266040926103d0565b61096781610904565b0190565b919061098e906020610986604086018681035f880152610951565b94019061025f565b565b156109985750565b6109ba906109a4610092565b918291637ed5977760e11b83526004830161096b565b0390fd5b634e487b7160e01b5f52601160045260245ffd5b6109e16109e79193929361018b565b9261018b565b82018092116109f257565b6109be565b610a1c81610a15610a0f610a0a5f61072b565b6100ab565b916100ab565b1415610857565b610a3834610a32610a2c5f6103b4565b9161018b565b116108da565b610a5f33610a57610a51610a4c600161076c565b6100ab565b916100ab565b143390610990565b610a8734610a81610a726002859061012f565b91610a7c836103a4565b6109d2565b90610496565b3490610ac8610ab67fcf6f6dcf9306290f700abd5d57c300a4ad7bf0d4086d5a5e88040fcecb0fef9092610123565b92610abf610092565b9182918261019b565b0390a2565b610ad5610779565b610aee610ae8610ae3610bb8565b6100ab565b916100ab565b03610af557565b610b17610b00610bb8565b5f91829163118cdaa760e01b83526004830161026c565b0390fd5b90610b2c60018060a01b039161045c565b9181191691161790565b90565b90610b4e610b49610b5592610123565b610b36565b8254610b1b565b9055565b610b625f61076c565b610b6c825f610b39565b90610ba0610b9a7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e093610123565b91610123565b91610ba9610092565b80610bb3816101f4565b0390a3565b610bc0610754565b50339056fea264697066735822122089020295cddb5abf8e7a440edeb1b9ff26e2692596ae6741abb6609e6d0d5c6c64736f6c634300081a0033",
}

// RaffleWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use RaffleWalletMetaData.ABI instead.
var RaffleWalletABI = RaffleWalletMetaData.ABI

// RaffleWalletBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RaffleWalletMetaData.Bin instead.
var RaffleWalletBin = RaffleWalletMetaData.Bin

// DeployRaffleWallet deploys a new Ethereum contract, binding an instance of RaffleWallet to it.
func DeployRaffleWallet(auth *bind.TransactOpts, backend bind.ContractBackend, game_ common.Address) (common.Address, *types.Transaction, *RaffleWallet, error) {
	parsed, err := RaffleWalletMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RaffleWalletBin), backend, game_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RaffleWallet{RaffleWalletCaller: RaffleWalletCaller{contract: contract}, RaffleWalletTransactor: RaffleWalletTransactor{contract: contract}, RaffleWalletFilterer: RaffleWalletFilterer{contract: contract}}, nil
}

// RaffleWallet is an auto generated Go binding around an Ethereum contract.
type RaffleWallet struct {
	RaffleWalletCaller     // Read-only binding to the contract
	RaffleWalletTransactor // Write-only binding to the contract
	RaffleWalletFilterer   // Log filterer for contract events
}

// RaffleWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type RaffleWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RaffleWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RaffleWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RaffleWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RaffleWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RaffleWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RaffleWalletSession struct {
	Contract     *RaffleWallet     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RaffleWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RaffleWalletCallerSession struct {
	Contract *RaffleWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RaffleWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RaffleWalletTransactorSession struct {
	Contract     *RaffleWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RaffleWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type RaffleWalletRaw struct {
	Contract *RaffleWallet // Generic contract binding to access the raw methods on
}

// RaffleWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RaffleWalletCallerRaw struct {
	Contract *RaffleWalletCaller // Generic read-only contract binding to access the raw methods on
}

// RaffleWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RaffleWalletTransactorRaw struct {
	Contract *RaffleWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRaffleWallet creates a new instance of RaffleWallet, bound to a specific deployed contract.
func NewRaffleWallet(address common.Address, backend bind.ContractBackend) (*RaffleWallet, error) {
	contract, err := bindRaffleWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RaffleWallet{RaffleWalletCaller: RaffleWalletCaller{contract: contract}, RaffleWalletTransactor: RaffleWalletTransactor{contract: contract}, RaffleWalletFilterer: RaffleWalletFilterer{contract: contract}}, nil
}

// NewRaffleWalletCaller creates a new read-only instance of RaffleWallet, bound to a specific deployed contract.
func NewRaffleWalletCaller(address common.Address, caller bind.ContractCaller) (*RaffleWalletCaller, error) {
	contract, err := bindRaffleWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RaffleWalletCaller{contract: contract}, nil
}

// NewRaffleWalletTransactor creates a new write-only instance of RaffleWallet, bound to a specific deployed contract.
func NewRaffleWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*RaffleWalletTransactor, error) {
	contract, err := bindRaffleWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RaffleWalletTransactor{contract: contract}, nil
}

// NewRaffleWalletFilterer creates a new log filterer instance of RaffleWallet, bound to a specific deployed contract.
func NewRaffleWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*RaffleWalletFilterer, error) {
	contract, err := bindRaffleWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RaffleWalletFilterer{contract: contract}, nil
}

// bindRaffleWallet binds a generic wrapper to an already deployed contract.
func bindRaffleWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RaffleWalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RaffleWallet *RaffleWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RaffleWallet.Contract.RaffleWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RaffleWallet *RaffleWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RaffleWallet.Contract.RaffleWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RaffleWallet *RaffleWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RaffleWallet.Contract.RaffleWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RaffleWallet *RaffleWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RaffleWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RaffleWallet *RaffleWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RaffleWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RaffleWallet *RaffleWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RaffleWallet.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_RaffleWallet *RaffleWalletCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RaffleWallet.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_RaffleWallet *RaffleWalletSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _RaffleWallet.Contract.Balances(&_RaffleWallet.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_RaffleWallet *RaffleWalletCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _RaffleWallet.Contract.Balances(&_RaffleWallet.CallOpts, arg0)
}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_RaffleWallet *RaffleWalletCaller) Game(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RaffleWallet.contract.Call(opts, &out, "game")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_RaffleWallet *RaffleWalletSession) Game() (common.Address, error) {
	return _RaffleWallet.Contract.Game(&_RaffleWallet.CallOpts)
}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_RaffleWallet *RaffleWalletCallerSession) Game() (common.Address, error) {
	return _RaffleWallet.Contract.Game(&_RaffleWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RaffleWallet *RaffleWalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RaffleWallet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RaffleWallet *RaffleWalletSession) Owner() (common.Address, error) {
	return _RaffleWallet.Contract.Owner(&_RaffleWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RaffleWallet *RaffleWalletCallerSession) Owner() (common.Address, error) {
	return _RaffleWallet.Contract.Owner(&_RaffleWallet.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address winner) payable returns()
func (_RaffleWallet *RaffleWalletTransactor) Deposit(opts *bind.TransactOpts, winner common.Address) (*types.Transaction, error) {
	return _RaffleWallet.contract.Transact(opts, "deposit", winner)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address winner) payable returns()
func (_RaffleWallet *RaffleWalletSession) Deposit(winner common.Address) (*types.Transaction, error) {
	return _RaffleWallet.Contract.Deposit(&_RaffleWallet.TransactOpts, winner)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address winner) payable returns()
func (_RaffleWallet *RaffleWalletTransactorSession) Deposit(winner common.Address) (*types.Transaction, error) {
	return _RaffleWallet.Contract.Deposit(&_RaffleWallet.TransactOpts, winner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RaffleWallet *RaffleWalletTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RaffleWallet.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RaffleWallet *RaffleWalletSession) RenounceOwnership() (*types.Transaction, error) {
	return _RaffleWallet.Contract.RenounceOwnership(&_RaffleWallet.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RaffleWallet *RaffleWalletTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RaffleWallet.Contract.RenounceOwnership(&_RaffleWallet.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RaffleWallet *RaffleWalletTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RaffleWallet.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RaffleWallet *RaffleWalletSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RaffleWallet.Contract.TransferOwnership(&_RaffleWallet.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RaffleWallet *RaffleWalletTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RaffleWallet.Contract.TransferOwnership(&_RaffleWallet.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_RaffleWallet *RaffleWalletTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RaffleWallet.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_RaffleWallet *RaffleWalletSession) Withdraw() (*types.Transaction, error) {
	return _RaffleWallet.Contract.Withdraw(&_RaffleWallet.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_RaffleWallet *RaffleWalletTransactorSession) Withdraw() (*types.Transaction, error) {
	return _RaffleWallet.Contract.Withdraw(&_RaffleWallet.TransactOpts)
}

// RaffleWalletOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RaffleWallet contract.
type RaffleWalletOwnershipTransferredIterator struct {
	Event *RaffleWalletOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RaffleWalletOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaffleWalletOwnershipTransferred)
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
		it.Event = new(RaffleWalletOwnershipTransferred)
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
func (it *RaffleWalletOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaffleWalletOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaffleWalletOwnershipTransferred represents a OwnershipTransferred event raised by the RaffleWallet contract.
type RaffleWalletOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RaffleWallet *RaffleWalletFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RaffleWalletOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RaffleWallet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RaffleWalletOwnershipTransferredIterator{contract: _RaffleWallet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RaffleWallet *RaffleWalletFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RaffleWalletOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RaffleWallet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaffleWalletOwnershipTransferred)
				if err := _RaffleWallet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RaffleWallet *RaffleWalletFilterer) ParseOwnershipTransferred(log types.Log) (*RaffleWalletOwnershipTransferred, error) {
	event := new(RaffleWalletOwnershipTransferred)
	if err := _RaffleWallet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaffleWalletRaffleDepositEventIterator is returned from FilterRaffleDepositEvent and is used to iterate over the raw logs and unpacked data for RaffleDepositEvent events raised by the RaffleWallet contract.
type RaffleWalletRaffleDepositEventIterator struct {
	Event *RaffleWalletRaffleDepositEvent // Event containing the contract specifics and raw log

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
func (it *RaffleWalletRaffleDepositEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaffleWalletRaffleDepositEvent)
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
		it.Event = new(RaffleWalletRaffleDepositEvent)
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
func (it *RaffleWalletRaffleDepositEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaffleWalletRaffleDepositEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaffleWalletRaffleDepositEvent represents a RaffleDepositEvent event raised by the RaffleWallet contract.
type RaffleWalletRaffleDepositEvent struct {
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRaffleDepositEvent is a free log retrieval operation binding the contract event 0xcf6f6dcf9306290f700abd5d57c300a4ad7bf0d4086d5a5e88040fcecb0fef90.
//
// Solidity: event RaffleDepositEvent(address indexed winner, uint256 amount)
func (_RaffleWallet *RaffleWalletFilterer) FilterRaffleDepositEvent(opts *bind.FilterOpts, winner []common.Address) (*RaffleWalletRaffleDepositEventIterator, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _RaffleWallet.contract.FilterLogs(opts, "RaffleDepositEvent", winnerRule)
	if err != nil {
		return nil, err
	}
	return &RaffleWalletRaffleDepositEventIterator{contract: _RaffleWallet.contract, event: "RaffleDepositEvent", logs: logs, sub: sub}, nil
}

// WatchRaffleDepositEvent is a free log subscription operation binding the contract event 0xcf6f6dcf9306290f700abd5d57c300a4ad7bf0d4086d5a5e88040fcecb0fef90.
//
// Solidity: event RaffleDepositEvent(address indexed winner, uint256 amount)
func (_RaffleWallet *RaffleWalletFilterer) WatchRaffleDepositEvent(opts *bind.WatchOpts, sink chan<- *RaffleWalletRaffleDepositEvent, winner []common.Address) (event.Subscription, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _RaffleWallet.contract.WatchLogs(opts, "RaffleDepositEvent", winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaffleWalletRaffleDepositEvent)
				if err := _RaffleWallet.contract.UnpackLog(event, "RaffleDepositEvent", log); err != nil {
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

// ParseRaffleDepositEvent is a log parse operation binding the contract event 0xcf6f6dcf9306290f700abd5d57c300a4ad7bf0d4086d5a5e88040fcecb0fef90.
//
// Solidity: event RaffleDepositEvent(address indexed winner, uint256 amount)
func (_RaffleWallet *RaffleWalletFilterer) ParseRaffleDepositEvent(log types.Log) (*RaffleWalletRaffleDepositEvent, error) {
	event := new(RaffleWalletRaffleDepositEvent)
	if err := _RaffleWallet.contract.UnpackLog(event, "RaffleDepositEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaffleWalletRaffleWithdrawalEventIterator is returned from FilterRaffleWithdrawalEvent and is used to iterate over the raw logs and unpacked data for RaffleWithdrawalEvent events raised by the RaffleWallet contract.
type RaffleWalletRaffleWithdrawalEventIterator struct {
	Event *RaffleWalletRaffleWithdrawalEvent // Event containing the contract specifics and raw log

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
func (it *RaffleWalletRaffleWithdrawalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaffleWalletRaffleWithdrawalEvent)
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
		it.Event = new(RaffleWalletRaffleWithdrawalEvent)
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
func (it *RaffleWalletRaffleWithdrawalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaffleWalletRaffleWithdrawalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaffleWalletRaffleWithdrawalEvent represents a RaffleWithdrawalEvent event raised by the RaffleWallet contract.
type RaffleWalletRaffleWithdrawalEvent struct {
	Destination common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRaffleWithdrawalEvent is a free log retrieval operation binding the contract event 0x49775450da95afb818c4651b894f124d05629f06572a18900bf29db74b04a0d8.
//
// Solidity: event RaffleWithdrawalEvent(address indexed destination, uint256 amount)
func (_RaffleWallet *RaffleWalletFilterer) FilterRaffleWithdrawalEvent(opts *bind.FilterOpts, destination []common.Address) (*RaffleWalletRaffleWithdrawalEventIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _RaffleWallet.contract.FilterLogs(opts, "RaffleWithdrawalEvent", destinationRule)
	if err != nil {
		return nil, err
	}
	return &RaffleWalletRaffleWithdrawalEventIterator{contract: _RaffleWallet.contract, event: "RaffleWithdrawalEvent", logs: logs, sub: sub}, nil
}

// WatchRaffleWithdrawalEvent is a free log subscription operation binding the contract event 0x49775450da95afb818c4651b894f124d05629f06572a18900bf29db74b04a0d8.
//
// Solidity: event RaffleWithdrawalEvent(address indexed destination, uint256 amount)
func (_RaffleWallet *RaffleWalletFilterer) WatchRaffleWithdrawalEvent(opts *bind.WatchOpts, sink chan<- *RaffleWalletRaffleWithdrawalEvent, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _RaffleWallet.contract.WatchLogs(opts, "RaffleWithdrawalEvent", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaffleWalletRaffleWithdrawalEvent)
				if err := _RaffleWallet.contract.UnpackLog(event, "RaffleWithdrawalEvent", log); err != nil {
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

// ParseRaffleWithdrawalEvent is a log parse operation binding the contract event 0x49775450da95afb818c4651b894f124d05629f06572a18900bf29db74b04a0d8.
//
// Solidity: event RaffleWithdrawalEvent(address indexed destination, uint256 amount)
func (_RaffleWallet *RaffleWalletFilterer) ParseRaffleWithdrawalEvent(log types.Log) (*RaffleWalletRaffleWithdrawalEvent, error) {
	event := new(RaffleWalletRaffleWithdrawalEvent)
	if err := _RaffleWallet.contract.UnpackLog(event, "RaffleWithdrawalEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
