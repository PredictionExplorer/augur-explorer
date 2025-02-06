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

// IMarketingWalletMetaData contains all meta data concerning the IMarketingWallet contract.
var IMarketingWalletMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractICosmicSignatureToken\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"CosmicSignatureTokenAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"marketerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structICosmicSignatureToken.MintSpec[]\",\"name\":\"specs_\",\"type\":\"tuple[]\"}],\"name\":\"payManyRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"marketerAddresses_\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"payManyRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"marketerAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"payReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICosmicSignatureToken\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setCosmicSignatureToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IMarketingWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use IMarketingWalletMetaData.ABI instead.
var IMarketingWalletABI = IMarketingWalletMetaData.ABI

// IMarketingWallet is an auto generated Go binding around an Ethereum contract.
type IMarketingWallet struct {
	IMarketingWalletCaller     // Read-only binding to the contract
	IMarketingWalletTransactor // Write-only binding to the contract
	IMarketingWalletFilterer   // Log filterer for contract events
}

// IMarketingWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMarketingWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMarketingWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMarketingWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMarketingWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMarketingWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMarketingWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMarketingWalletSession struct {
	Contract     *IMarketingWallet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMarketingWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMarketingWalletCallerSession struct {
	Contract *IMarketingWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IMarketingWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMarketingWalletTransactorSession struct {
	Contract     *IMarketingWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IMarketingWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMarketingWalletRaw struct {
	Contract *IMarketingWallet // Generic contract binding to access the raw methods on
}

// IMarketingWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMarketingWalletCallerRaw struct {
	Contract *IMarketingWalletCaller // Generic read-only contract binding to access the raw methods on
}

// IMarketingWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMarketingWalletTransactorRaw struct {
	Contract *IMarketingWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMarketingWallet creates a new instance of IMarketingWallet, bound to a specific deployed contract.
func NewIMarketingWallet(address common.Address, backend bind.ContractBackend) (*IMarketingWallet, error) {
	contract, err := bindIMarketingWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMarketingWallet{IMarketingWalletCaller: IMarketingWalletCaller{contract: contract}, IMarketingWalletTransactor: IMarketingWalletTransactor{contract: contract}, IMarketingWalletFilterer: IMarketingWalletFilterer{contract: contract}}, nil
}

// NewIMarketingWalletCaller creates a new read-only instance of IMarketingWallet, bound to a specific deployed contract.
func NewIMarketingWalletCaller(address common.Address, caller bind.ContractCaller) (*IMarketingWalletCaller, error) {
	contract, err := bindIMarketingWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMarketingWalletCaller{contract: contract}, nil
}

// NewIMarketingWalletTransactor creates a new write-only instance of IMarketingWallet, bound to a specific deployed contract.
func NewIMarketingWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*IMarketingWalletTransactor, error) {
	contract, err := bindIMarketingWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMarketingWalletTransactor{contract: contract}, nil
}

// NewIMarketingWalletFilterer creates a new log filterer instance of IMarketingWallet, bound to a specific deployed contract.
func NewIMarketingWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*IMarketingWalletFilterer, error) {
	contract, err := bindIMarketingWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMarketingWalletFilterer{contract: contract}, nil
}

// bindIMarketingWallet binds a generic wrapper to an already deployed contract.
func bindIMarketingWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IMarketingWalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMarketingWallet *IMarketingWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMarketingWallet.Contract.IMarketingWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMarketingWallet *IMarketingWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMarketingWallet.Contract.IMarketingWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMarketingWallet *IMarketingWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMarketingWallet.Contract.IMarketingWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMarketingWallet *IMarketingWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMarketingWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMarketingWallet *IMarketingWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMarketingWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMarketingWallet *IMarketingWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMarketingWallet.Contract.contract.Transact(opts, method, params...)
}

// PayManyRewards is a paid mutator transaction binding the contract method 0xc603317f.
//
// Solidity: function payManyRewards((address,uint256)[] specs_) returns()
func (_IMarketingWallet *IMarketingWalletTransactor) PayManyRewards(opts *bind.TransactOpts, specs_ []ICosmicSignatureTokenMintSpec) (*types.Transaction, error) {
	return _IMarketingWallet.contract.Transact(opts, "payManyRewards", specs_)
}

// PayManyRewards is a paid mutator transaction binding the contract method 0xc603317f.
//
// Solidity: function payManyRewards((address,uint256)[] specs_) returns()
func (_IMarketingWallet *IMarketingWalletSession) PayManyRewards(specs_ []ICosmicSignatureTokenMintSpec) (*types.Transaction, error) {
	return _IMarketingWallet.Contract.PayManyRewards(&_IMarketingWallet.TransactOpts, specs_)
}

// PayManyRewards is a paid mutator transaction binding the contract method 0xc603317f.
//
// Solidity: function payManyRewards((address,uint256)[] specs_) returns()
func (_IMarketingWallet *IMarketingWalletTransactorSession) PayManyRewards(specs_ []ICosmicSignatureTokenMintSpec) (*types.Transaction, error) {
	return _IMarketingWallet.Contract.PayManyRewards(&_IMarketingWallet.TransactOpts, specs_)
}

// PayManyRewards0 is a paid mutator transaction binding the contract method 0xffcbf0b4.
//
// Solidity: function payManyRewards(address[] marketerAddresses_, uint256 amount_) returns()
func (_IMarketingWallet *IMarketingWalletTransactor) PayManyRewards0(opts *bind.TransactOpts, marketerAddresses_ []common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IMarketingWallet.contract.Transact(opts, "payManyRewards0", marketerAddresses_, amount_)
}

// PayManyRewards0 is a paid mutator transaction binding the contract method 0xffcbf0b4.
//
// Solidity: function payManyRewards(address[] marketerAddresses_, uint256 amount_) returns()
func (_IMarketingWallet *IMarketingWalletSession) PayManyRewards0(marketerAddresses_ []common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IMarketingWallet.Contract.PayManyRewards0(&_IMarketingWallet.TransactOpts, marketerAddresses_, amount_)
}

// PayManyRewards0 is a paid mutator transaction binding the contract method 0xffcbf0b4.
//
// Solidity: function payManyRewards(address[] marketerAddresses_, uint256 amount_) returns()
func (_IMarketingWallet *IMarketingWalletTransactorSession) PayManyRewards0(marketerAddresses_ []common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IMarketingWallet.Contract.PayManyRewards0(&_IMarketingWallet.TransactOpts, marketerAddresses_, amount_)
}

// PayReward is a paid mutator transaction binding the contract method 0x8b8afcd7.
//
// Solidity: function payReward(address marketerAddress_, uint256 amount_) returns()
func (_IMarketingWallet *IMarketingWalletTransactor) PayReward(opts *bind.TransactOpts, marketerAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IMarketingWallet.contract.Transact(opts, "payReward", marketerAddress_, amount_)
}

// PayReward is a paid mutator transaction binding the contract method 0x8b8afcd7.
//
// Solidity: function payReward(address marketerAddress_, uint256 amount_) returns()
func (_IMarketingWallet *IMarketingWalletSession) PayReward(marketerAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IMarketingWallet.Contract.PayReward(&_IMarketingWallet.TransactOpts, marketerAddress_, amount_)
}

// PayReward is a paid mutator transaction binding the contract method 0x8b8afcd7.
//
// Solidity: function payReward(address marketerAddress_, uint256 amount_) returns()
func (_IMarketingWallet *IMarketingWalletTransactorSession) PayReward(marketerAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IMarketingWallet.Contract.PayReward(&_IMarketingWallet.TransactOpts, marketerAddress_, amount_)
}

// SetCosmicSignatureToken is a paid mutator transaction binding the contract method 0x9646d758.
//
// Solidity: function setCosmicSignatureToken(address newValue_) returns()
func (_IMarketingWallet *IMarketingWalletTransactor) SetCosmicSignatureToken(opts *bind.TransactOpts, newValue_ common.Address) (*types.Transaction, error) {
	return _IMarketingWallet.contract.Transact(opts, "setCosmicSignatureToken", newValue_)
}

// SetCosmicSignatureToken is a paid mutator transaction binding the contract method 0x9646d758.
//
// Solidity: function setCosmicSignatureToken(address newValue_) returns()
func (_IMarketingWallet *IMarketingWalletSession) SetCosmicSignatureToken(newValue_ common.Address) (*types.Transaction, error) {
	return _IMarketingWallet.Contract.SetCosmicSignatureToken(&_IMarketingWallet.TransactOpts, newValue_)
}

// SetCosmicSignatureToken is a paid mutator transaction binding the contract method 0x9646d758.
//
// Solidity: function setCosmicSignatureToken(address newValue_) returns()
func (_IMarketingWallet *IMarketingWalletTransactorSession) SetCosmicSignatureToken(newValue_ common.Address) (*types.Transaction, error) {
	return _IMarketingWallet.Contract.SetCosmicSignatureToken(&_IMarketingWallet.TransactOpts, newValue_)
}

// IMarketingWalletCosmicSignatureTokenAddressChangedIterator is returned from FilterCosmicSignatureTokenAddressChanged and is used to iterate over the raw logs and unpacked data for CosmicSignatureTokenAddressChanged events raised by the IMarketingWallet contract.
type IMarketingWalletCosmicSignatureTokenAddressChangedIterator struct {
	Event *IMarketingWalletCosmicSignatureTokenAddressChanged // Event containing the contract specifics and raw log

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
func (it *IMarketingWalletCosmicSignatureTokenAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMarketingWalletCosmicSignatureTokenAddressChanged)
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
		it.Event = new(IMarketingWalletCosmicSignatureTokenAddressChanged)
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
func (it *IMarketingWalletCosmicSignatureTokenAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMarketingWalletCosmicSignatureTokenAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMarketingWalletCosmicSignatureTokenAddressChanged represents a CosmicSignatureTokenAddressChanged event raised by the IMarketingWallet contract.
type IMarketingWalletCosmicSignatureTokenAddressChanged struct {
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCosmicSignatureTokenAddressChanged is a free log retrieval operation binding the contract event 0x9b3eda10f1724a2cf9f7dae4ac263c77908df4d00e92f1377b66fc8be37cd8c1.
//
// Solidity: event CosmicSignatureTokenAddressChanged(address newValue)
func (_IMarketingWallet *IMarketingWalletFilterer) FilterCosmicSignatureTokenAddressChanged(opts *bind.FilterOpts) (*IMarketingWalletCosmicSignatureTokenAddressChangedIterator, error) {

	logs, sub, err := _IMarketingWallet.contract.FilterLogs(opts, "CosmicSignatureTokenAddressChanged")
	if err != nil {
		return nil, err
	}
	return &IMarketingWalletCosmicSignatureTokenAddressChangedIterator{contract: _IMarketingWallet.contract, event: "CosmicSignatureTokenAddressChanged", logs: logs, sub: sub}, nil
}

// WatchCosmicSignatureTokenAddressChanged is a free log subscription operation binding the contract event 0x9b3eda10f1724a2cf9f7dae4ac263c77908df4d00e92f1377b66fc8be37cd8c1.
//
// Solidity: event CosmicSignatureTokenAddressChanged(address newValue)
func (_IMarketingWallet *IMarketingWalletFilterer) WatchCosmicSignatureTokenAddressChanged(opts *bind.WatchOpts, sink chan<- *IMarketingWalletCosmicSignatureTokenAddressChanged) (event.Subscription, error) {

	logs, sub, err := _IMarketingWallet.contract.WatchLogs(opts, "CosmicSignatureTokenAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMarketingWalletCosmicSignatureTokenAddressChanged)
				if err := _IMarketingWallet.contract.UnpackLog(event, "CosmicSignatureTokenAddressChanged", log); err != nil {
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

// ParseCosmicSignatureTokenAddressChanged is a log parse operation binding the contract event 0x9b3eda10f1724a2cf9f7dae4ac263c77908df4d00e92f1377b66fc8be37cd8c1.
//
// Solidity: event CosmicSignatureTokenAddressChanged(address newValue)
func (_IMarketingWallet *IMarketingWalletFilterer) ParseCosmicSignatureTokenAddressChanged(log types.Log) (*IMarketingWalletCosmicSignatureTokenAddressChanged, error) {
	event := new(IMarketingWalletCosmicSignatureTokenAddressChanged)
	if err := _IMarketingWallet.contract.UnpackLog(event, "CosmicSignatureTokenAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMarketingWalletRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the IMarketingWallet contract.
type IMarketingWalletRewardPaidIterator struct {
	Event *IMarketingWalletRewardPaid // Event containing the contract specifics and raw log

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
func (it *IMarketingWalletRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMarketingWalletRewardPaid)
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
		it.Event = new(IMarketingWalletRewardPaid)
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
func (it *IMarketingWalletRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMarketingWalletRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMarketingWalletRewardPaid represents a RewardPaid event raised by the IMarketingWallet contract.
type IMarketingWalletRewardPaid struct {
	MarketerAddress common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed marketerAddress, uint256 amount)
func (_IMarketingWallet *IMarketingWalletFilterer) FilterRewardPaid(opts *bind.FilterOpts, marketerAddress []common.Address) (*IMarketingWalletRewardPaidIterator, error) {

	var marketerAddressRule []interface{}
	for _, marketerAddressItem := range marketerAddress {
		marketerAddressRule = append(marketerAddressRule, marketerAddressItem)
	}

	logs, sub, err := _IMarketingWallet.contract.FilterLogs(opts, "RewardPaid", marketerAddressRule)
	if err != nil {
		return nil, err
	}
	return &IMarketingWalletRewardPaidIterator{contract: _IMarketingWallet.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed marketerAddress, uint256 amount)
func (_IMarketingWallet *IMarketingWalletFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *IMarketingWalletRewardPaid, marketerAddress []common.Address) (event.Subscription, error) {

	var marketerAddressRule []interface{}
	for _, marketerAddressItem := range marketerAddress {
		marketerAddressRule = append(marketerAddressRule, marketerAddressItem)
	}

	logs, sub, err := _IMarketingWallet.contract.WatchLogs(opts, "RewardPaid", marketerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMarketingWalletRewardPaid)
				if err := _IMarketingWallet.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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

// ParseRewardPaid is a log parse operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed marketerAddress, uint256 amount)
func (_IMarketingWallet *IMarketingWalletFilterer) ParseRewardPaid(log types.Log) (*IMarketingWalletRewardPaid, error) {
	event := new(IMarketingWalletRewardPaid)
	if err := _IMarketingWallet.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketingWalletMetaData contains all meta data concerning the MarketingWallet contract.
var MarketingWalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractCosmicSignatureToken\",\"name\":\"token_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractICosmicSignatureToken\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"CosmicSignatureTokenAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"marketerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structICosmicSignatureToken.MintSpec[]\",\"name\":\"specs_\",\"type\":\"tuple[]\"}],\"name\":\"payManyRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"marketerAddresses_\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"payManyRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"marketerAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"payReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICosmicSignatureToken\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setCosmicSignatureToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractCosmicSignatureToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523461002f57610019610014610100565b6102a3565b610021610034565b61100d610411823961100d90f35b61003a565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906100669061003e565b810190811060018060401b0382111761007e57604052565b610048565b9061009661008f610034565b928361005c565b565b5f80fd5b60018060a01b031690565b6100b09061009c565b90565b6100bc906100a7565b90565b6100c8816100b3565b036100cf57565b5f80fd5b905051906100e0826100bf565b565b906020828203126100fb576100f8915f016100d3565b90565b610098565b61011e61141e8038038061011381610083565b9283398101906100e2565b90565b90565b61013861013361013d9261009c565b610121565b61009c565b90565b61014990610124565b90565b61015590610140565b90565b90565b61016f61016a61017492610158565b610121565b61009c565b90565b6101809061015b565b90565b60209181520190565b5f7f5468652070726f76696465642061646472657373206973207a65726f2e000000910152565b6101c0601d602092610183565b6101c98161018c565b0190565b6101e29060208101905f8183039101526101b3565b90565b6101ee8161014c565b6102086102026101fd5f610177565b6100a7565b916100a7565b146102185761021690610296565b565b610220610034565b63eac0d38960e01b815280610237600482016101cd565b0390fd5b5f1b90565b9061025160018060a01b039161023b565b9181191691161790565b61026490610124565b90565b6102709061025b565b90565b90565b9061028b61028661029292610267565b610273565b8254610240565b9055565b6102a1906001610276565b565b6102b5906102b0336102b7565b6101e5565b565b6102c0906102c2565b565b6102cb906102cd565b565b6102d6906102fa565b565b6102e1906100a7565b9052565b91906102f8905f602085019401906102d8565b565b8061031561030f61030a5f610177565b6100a7565b916100a7565b1461032557610323906103b1565b565b6103486103315f610177565b5f918291631e4fbdf760e01b8352600483016102e5565b0390fd5b5f1c90565b60018060a01b031690565b61036861036d9161034c565b610351565b90565b61037a905461035c565b90565b61038690610140565b90565b90565b906103a161039c6103a89261037d565b610389565b8254610240565b9055565b5f0190565b6103ba5f610370565b6103c4825f61038c565b906103f86103f27f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09361037d565b9161037d565b91610401610034565b8061040b816103ac565b0390a356fe60806040526004361015610013575b610523565b61001d5f3561009c565b8063715018a6146100975780638b8afcd7146100925780638da5cb5b1461008d5780639646d75814610088578063c603317f14610083578063f2fde38b1461007e578063fc0c546a146100795763ffcbf0b40361000e576104ef565b610436565b610366565b610314565b610260565b6101de565b610188565b6100c4565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f9103126100ba57565b6100ac565b5f0190565b346100f2576100d43660046100b0565b6100dc610577565b6100e46100a2565b806100ee816100bf565b0390f35b6100a8565b5f80fd5b60018060a01b031690565b61010f906100fb565b90565b61011b81610106565b0361012257565b5f80fd5b9050359061013382610112565b565b90565b61014181610135565b0361014857565b5f80fd5b9050359061015982610138565b565b91906040838203126101835780610177610180925f8601610126565b9360200161014c565b90565b6100ac565b346101b7576101a161019b36600461015b565b9061077c565b6101a96100a2565b806101b3816100bf565b0390f35b6100a8565b6101c590610106565b9052565b91906101dc905f602085019401906101bc565b565b3461020e576101ee3660046100b0565b61020a6101f96107b8565b6102016100a2565b918291826101c9565b0390f35b6100a8565b61021c90610106565b90565b61022881610213565b0361022f57565b5f80fd5b905035906102408261021f565b565b9060208282031261025b57610258915f01610233565b90565b6100ac565b3461028e57610278610273366004610242565b610983565b6102806100a2565b8061028a816100bf565b0390f35b6100a8565b5f80fd5b5f80fd5b5f80fd5b909182601f830112156102d95781359167ffffffffffffffff83116102d45760200192604083028401116102cf57565b61029b565b610297565b610293565b9060208282031261030f575f82013567ffffffffffffffff811161030a57610306920161029f565b9091565b6100f7565b6100ac565b346103435761032d6103273660046102de565b90610c7b565b6103356100a2565b8061033f816100bf565b0390f35b6100a8565b906020828203126103615761035e915f01610126565b90565b6100ac565b346103945761037e610379366004610348565b610cec565b6103866100a2565b80610390816100bf565b0390f35b6100a8565b1c90565b60018060a01b031690565b6103b89060086103bd9302610399565b61039d565b90565b906103cb91546103a8565b90565b6103da60015f906103c0565b90565b90565b6103f46103ef6103f9926100fb565b6103dd565b6100fb565b90565b610405906103e0565b90565b610411906103fc565b90565b61041d90610408565b9052565b9190610434905f60208501940190610414565b565b34610466576104463660046100b0565b6104626104516103ce565b6104596100a2565b91829182610421565b0390f35b6100a8565b909182601f830112156104a55781359167ffffffffffffffff83116104a057602001926020830284011161049b57565b61029b565b610297565b610293565b916040838303126104ea575f83013567ffffffffffffffff81116104e5576104d7836104e292860161046b565b93909460200161014c565b90565b6100f7565b6100ac565b3461051e576105086105023660046104aa565b91610eed565b6105106100a2565b8061051a816100bf565b0390f35b6100a8565b5f80fd5b61052f610efa565b610537610564565b565b90565b61055061054b61055592610539565b6103dd565b6100fb565b90565b6105619061053c565b90565b6105756105705f610558565b610f6b565b565b61057f610527565b565b906105939161058e610efa565b6106b3565b565b5f1c90565b6105a66105ab91610595565b61039d565b90565b6105b8905461059a565b90565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906105e7906105bf565b810190811067ffffffffffffffff82111761060157604052565b6105c9565b60e01b90565b151590565b61061a8161060c565b0361062157565b5f80fd5b9050519061063282610611565b565b9060208282031261064d5761064a915f01610625565b90565b6100ac565b61065b90610135565b9052565b91602061068092949361067960408201965f8301906101bc565b0190610652565b565b61068a6100a2565b3d5f823e3d90fd5b61069b906103fc565b90565b91906106b1905f60208501940190610652565b565b6106c56106c060016105ae565b610408565b602063a9059cbb9183906106ec5f87956106f76106e06100a2565b97889687958694610606565b84526004840161065f565b03925af180156107775761074b575b506107466107347fe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e048692610692565b9261073d6100a2565b9182918261069e565b0390a2565b61076b9060203d8111610770575b61076381836105dd565b810190610634565b610706565b503d610759565b610682565b9061078691610581565b565b5f90565b60018060a01b031690565b6107a36107a891610595565b61078c565b90565b6107b59054610797565b90565b6107c0610788565b506107ca5f6107ab565b90565b6107de906107d9610efa565b61084e565b565b6107e9906103fc565b90565b60209181520190565b5f7f5468652070726f76696465642061646472657373206973207a65726f2e000000910152565b610829601d6020926107ec565b610832816107f5565b0190565b61084b9060208101905f81830391015261081c565b90565b610857816107e0565b61087161086b6108665f610558565b610106565b91610106565b146108815761087f9061092d565b565b6108896100a2565b63eac0d38960e01b8152806108a060048201610836565b0390fd5b6108ad906103e0565b90565b6108b9906108a4565b90565b5f1b90565b906108d260018060a01b03916108bc565b9181191691161790565b6108e5906108a4565b90565b90565b906109006108fb610907926108dc565b6108e8565b82546108c1565b9055565b610914906107e0565b9052565b919061092b905f6020850194019061090b565b565b61094861094161093c836107e0565b6108b0565b60016108eb565b61097e7f9b3eda10f1724a2cf9f7dae4ac263c77908df4d00e92f1377b66fc8be37cd8c1916109756100a2565b91829182610918565b0390a1565b61098c906107cd565b565b906109a09161099b610efa565b610b4c565b565b5f9103126109ac57565b6100ac565b60209181520190565b90565b506109cc906020810190610126565b90565b6109d890610106565b9052565b506109eb90602081019061014c565b90565b6109f790610135565b9052565b906020610a26610a2e93610a1d610a145f8301836109bd565b5f8601906109cf565b828101906109dc565b9101906109ee565b565b90610a3d816040936109fb565b0190565b5090565b60400190565b91610a5982610a5f926109b1565b926109ba565b90815f905b828210610a72575050505090565b90919293610a94610a8e600192610a898886610a41565b610a30565b95610a45565b920190929192610a64565b9091610ab69260208301925f818503910152610a4b565b90565b5090565b610ad1610acc610ad692610539565b6103dd565b610135565b90565b634e487b7160e01b5f52601160045260245ffd5b610af690610135565b5f8114610b04576001900390565b610ad9565b634e487b7160e01b5f52603260045260245ffd5b9190811015610b2d576040020190565b610b09565b35610b3c81610112565b90565b35610b4981610138565b90565b610b5e610b5960016105ae565b610408565b638d6bd91c828492803b15610c7657610b8a5f8094610b95610b7e6100a2565b97889687958694610606565b845260048401610a9f565b03925af18015610c7157610c45575b50610bb0818390610ab9565b5b80610bc4610bbe5f610abd565b91610135565b1115610c4057610bd390610aed565b92610be082848691610b1d565b610bf76020610bf05f8401610b32565b9201610b3f565b90610c37610c257fe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e048692610692565b92610c2e6100a2565b9182918261069e565b0390a292610bb1565b505050565b610c64905f3d8111610c6a575b610c5c81836105dd565b8101906109a2565b5f610ba4565b503d610c52565b610682565b6105bb565b90610c859161098e565b565b610c9890610c93610efa565b610c9a565b565b80610cb5610caf610caa5f610558565b610106565b91610106565b14610cc557610cc390610f6b565b565b610ce8610cd15f610558565b5f918291631e4fbdf760e01b8352600483016101c9565b0390fd5b610cf590610c87565b565b90610d0a9291610d05610efa565b610dc4565b565b60209181520190565b90565b90610d25816020936109cf565b0190565b60200190565b91610d3d82610d4392610d0c565b92610d15565b90815f905b828210610d56575050505090565b90919293610d78610d72600192610d6d88866109bd565b610d18565b95610d29565b920190929192610d48565b939290610da1602091610da99460408801918883035f8a0152610d2f565b940190610652565b565b5090565b9190811015610dbf576020020190565b610b09565b909192610dd9610dd460016105ae565b610408565b630a8e30d98385928792813b15610ee8575f610e0891610e138296610dfc6100a2565b98899788968795610606565b855260048501610d83565b03925af18015610ee357610eb7575b50610e2e828490610dab565b5b80610e42610e3c5f610abd565b91610135565b1115610eb057610e5190610aed565b90610e66610e6184868591610daf565b610b32565b8590610ea7610e957fe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e048692610692565b92610e9e6100a2565b9182918261069e565b0390a290610e2f565b5092505050565b610ed6905f3d8111610edc575b610ece81836105dd565b8101906109a2565b5f610e22565b503d610ec4565b610682565b6105bb565b90610ef89291610cf7565b565b610f026107b8565b610f1b610f15610f10610fca565b610106565b91610106565b03610f2257565b610f44610f2d610fca565b5f91829163118cdaa760e01b8352600483016101c9565b0390fd5b90565b90610f60610f5b610f6792610692565b610f48565b82546108c1565b9055565b610f745f6107ab565b610f7e825f610f4b565b90610fb2610fac7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e093610692565b91610692565b91610fbb6100a2565b80610fc5816100bf565b0390a3565b610fd2610788565b50339056fea2646970667358221220c1ecfd71af5af33772fab1089109723dcf7a996fca5f9878a25175a05bf5936964736f6c634300081c0033",
}

// MarketingWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketingWalletMetaData.ABI instead.
var MarketingWalletABI = MarketingWalletMetaData.ABI

// MarketingWalletBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MarketingWalletMetaData.Bin instead.
var MarketingWalletBin = MarketingWalletMetaData.Bin

// DeployMarketingWallet deploys a new Ethereum contract, binding an instance of MarketingWallet to it.
func DeployMarketingWallet(auth *bind.TransactOpts, backend bind.ContractBackend, token_ common.Address) (common.Address, *types.Transaction, *MarketingWallet, error) {
	parsed, err := MarketingWalletMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MarketingWalletBin), backend, token_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MarketingWallet{MarketingWalletCaller: MarketingWalletCaller{contract: contract}, MarketingWalletTransactor: MarketingWalletTransactor{contract: contract}, MarketingWalletFilterer: MarketingWalletFilterer{contract: contract}}, nil
}

// MarketingWallet is an auto generated Go binding around an Ethereum contract.
type MarketingWallet struct {
	MarketingWalletCaller     // Read-only binding to the contract
	MarketingWalletTransactor // Write-only binding to the contract
	MarketingWalletFilterer   // Log filterer for contract events
}

// MarketingWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketingWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketingWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketingWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketingWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketingWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketingWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketingWalletSession struct {
	Contract     *MarketingWallet  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketingWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketingWalletCallerSession struct {
	Contract *MarketingWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// MarketingWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketingWalletTransactorSession struct {
	Contract     *MarketingWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// MarketingWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketingWalletRaw struct {
	Contract *MarketingWallet // Generic contract binding to access the raw methods on
}

// MarketingWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketingWalletCallerRaw struct {
	Contract *MarketingWalletCaller // Generic read-only contract binding to access the raw methods on
}

// MarketingWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketingWalletTransactorRaw struct {
	Contract *MarketingWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketingWallet creates a new instance of MarketingWallet, bound to a specific deployed contract.
func NewMarketingWallet(address common.Address, backend bind.ContractBackend) (*MarketingWallet, error) {
	contract, err := bindMarketingWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MarketingWallet{MarketingWalletCaller: MarketingWalletCaller{contract: contract}, MarketingWalletTransactor: MarketingWalletTransactor{contract: contract}, MarketingWalletFilterer: MarketingWalletFilterer{contract: contract}}, nil
}

// NewMarketingWalletCaller creates a new read-only instance of MarketingWallet, bound to a specific deployed contract.
func NewMarketingWalletCaller(address common.Address, caller bind.ContractCaller) (*MarketingWalletCaller, error) {
	contract, err := bindMarketingWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketingWalletCaller{contract: contract}, nil
}

// NewMarketingWalletTransactor creates a new write-only instance of MarketingWallet, bound to a specific deployed contract.
func NewMarketingWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketingWalletTransactor, error) {
	contract, err := bindMarketingWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketingWalletTransactor{contract: contract}, nil
}

// NewMarketingWalletFilterer creates a new log filterer instance of MarketingWallet, bound to a specific deployed contract.
func NewMarketingWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketingWalletFilterer, error) {
	contract, err := bindMarketingWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketingWalletFilterer{contract: contract}, nil
}

// bindMarketingWallet binds a generic wrapper to an already deployed contract.
func bindMarketingWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MarketingWalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketingWallet *MarketingWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketingWallet.Contract.MarketingWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketingWallet *MarketingWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketingWallet.Contract.MarketingWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketingWallet *MarketingWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketingWallet.Contract.MarketingWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketingWallet *MarketingWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketingWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketingWallet *MarketingWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketingWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketingWallet *MarketingWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketingWallet.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MarketingWallet *MarketingWalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketingWallet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MarketingWallet *MarketingWalletSession) Owner() (common.Address, error) {
	return _MarketingWallet.Contract.Owner(&_MarketingWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MarketingWallet *MarketingWalletCallerSession) Owner() (common.Address, error) {
	return _MarketingWallet.Contract.Owner(&_MarketingWallet.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_MarketingWallet *MarketingWalletCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketingWallet.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_MarketingWallet *MarketingWalletSession) Token() (common.Address, error) {
	return _MarketingWallet.Contract.Token(&_MarketingWallet.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_MarketingWallet *MarketingWalletCallerSession) Token() (common.Address, error) {
	return _MarketingWallet.Contract.Token(&_MarketingWallet.CallOpts)
}

// PayManyRewards is a paid mutator transaction binding the contract method 0xc603317f.
//
// Solidity: function payManyRewards((address,uint256)[] specs_) returns()
func (_MarketingWallet *MarketingWalletTransactor) PayManyRewards(opts *bind.TransactOpts, specs_ []ICosmicSignatureTokenMintSpec) (*types.Transaction, error) {
	return _MarketingWallet.contract.Transact(opts, "payManyRewards", specs_)
}

// PayManyRewards is a paid mutator transaction binding the contract method 0xc603317f.
//
// Solidity: function payManyRewards((address,uint256)[] specs_) returns()
func (_MarketingWallet *MarketingWalletSession) PayManyRewards(specs_ []ICosmicSignatureTokenMintSpec) (*types.Transaction, error) {
	return _MarketingWallet.Contract.PayManyRewards(&_MarketingWallet.TransactOpts, specs_)
}

// PayManyRewards is a paid mutator transaction binding the contract method 0xc603317f.
//
// Solidity: function payManyRewards((address,uint256)[] specs_) returns()
func (_MarketingWallet *MarketingWalletTransactorSession) PayManyRewards(specs_ []ICosmicSignatureTokenMintSpec) (*types.Transaction, error) {
	return _MarketingWallet.Contract.PayManyRewards(&_MarketingWallet.TransactOpts, specs_)
}

// PayManyRewards0 is a paid mutator transaction binding the contract method 0xffcbf0b4.
//
// Solidity: function payManyRewards(address[] marketerAddresses_, uint256 amount_) returns()
func (_MarketingWallet *MarketingWalletTransactor) PayManyRewards0(opts *bind.TransactOpts, marketerAddresses_ []common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _MarketingWallet.contract.Transact(opts, "payManyRewards0", marketerAddresses_, amount_)
}

// PayManyRewards0 is a paid mutator transaction binding the contract method 0xffcbf0b4.
//
// Solidity: function payManyRewards(address[] marketerAddresses_, uint256 amount_) returns()
func (_MarketingWallet *MarketingWalletSession) PayManyRewards0(marketerAddresses_ []common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _MarketingWallet.Contract.PayManyRewards0(&_MarketingWallet.TransactOpts, marketerAddresses_, amount_)
}

// PayManyRewards0 is a paid mutator transaction binding the contract method 0xffcbf0b4.
//
// Solidity: function payManyRewards(address[] marketerAddresses_, uint256 amount_) returns()
func (_MarketingWallet *MarketingWalletTransactorSession) PayManyRewards0(marketerAddresses_ []common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _MarketingWallet.Contract.PayManyRewards0(&_MarketingWallet.TransactOpts, marketerAddresses_, amount_)
}

// PayReward is a paid mutator transaction binding the contract method 0x8b8afcd7.
//
// Solidity: function payReward(address marketerAddress_, uint256 amount_) returns()
func (_MarketingWallet *MarketingWalletTransactor) PayReward(opts *bind.TransactOpts, marketerAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _MarketingWallet.contract.Transact(opts, "payReward", marketerAddress_, amount_)
}

// PayReward is a paid mutator transaction binding the contract method 0x8b8afcd7.
//
// Solidity: function payReward(address marketerAddress_, uint256 amount_) returns()
func (_MarketingWallet *MarketingWalletSession) PayReward(marketerAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _MarketingWallet.Contract.PayReward(&_MarketingWallet.TransactOpts, marketerAddress_, amount_)
}

// PayReward is a paid mutator transaction binding the contract method 0x8b8afcd7.
//
// Solidity: function payReward(address marketerAddress_, uint256 amount_) returns()
func (_MarketingWallet *MarketingWalletTransactorSession) PayReward(marketerAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _MarketingWallet.Contract.PayReward(&_MarketingWallet.TransactOpts, marketerAddress_, amount_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MarketingWallet *MarketingWalletTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketingWallet.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MarketingWallet *MarketingWalletSession) RenounceOwnership() (*types.Transaction, error) {
	return _MarketingWallet.Contract.RenounceOwnership(&_MarketingWallet.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MarketingWallet *MarketingWalletTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MarketingWallet.Contract.RenounceOwnership(&_MarketingWallet.TransactOpts)
}

// SetCosmicSignatureToken is a paid mutator transaction binding the contract method 0x9646d758.
//
// Solidity: function setCosmicSignatureToken(address newValue_) returns()
func (_MarketingWallet *MarketingWalletTransactor) SetCosmicSignatureToken(opts *bind.TransactOpts, newValue_ common.Address) (*types.Transaction, error) {
	return _MarketingWallet.contract.Transact(opts, "setCosmicSignatureToken", newValue_)
}

// SetCosmicSignatureToken is a paid mutator transaction binding the contract method 0x9646d758.
//
// Solidity: function setCosmicSignatureToken(address newValue_) returns()
func (_MarketingWallet *MarketingWalletSession) SetCosmicSignatureToken(newValue_ common.Address) (*types.Transaction, error) {
	return _MarketingWallet.Contract.SetCosmicSignatureToken(&_MarketingWallet.TransactOpts, newValue_)
}

// SetCosmicSignatureToken is a paid mutator transaction binding the contract method 0x9646d758.
//
// Solidity: function setCosmicSignatureToken(address newValue_) returns()
func (_MarketingWallet *MarketingWalletTransactorSession) SetCosmicSignatureToken(newValue_ common.Address) (*types.Transaction, error) {
	return _MarketingWallet.Contract.SetCosmicSignatureToken(&_MarketingWallet.TransactOpts, newValue_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MarketingWallet *MarketingWalletTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MarketingWallet.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MarketingWallet *MarketingWalletSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MarketingWallet.Contract.TransferOwnership(&_MarketingWallet.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MarketingWallet *MarketingWalletTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MarketingWallet.Contract.TransferOwnership(&_MarketingWallet.TransactOpts, newOwner)
}

// MarketingWalletCosmicSignatureTokenAddressChangedIterator is returned from FilterCosmicSignatureTokenAddressChanged and is used to iterate over the raw logs and unpacked data for CosmicSignatureTokenAddressChanged events raised by the MarketingWallet contract.
type MarketingWalletCosmicSignatureTokenAddressChangedIterator struct {
	Event *MarketingWalletCosmicSignatureTokenAddressChanged // Event containing the contract specifics and raw log

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
func (it *MarketingWalletCosmicSignatureTokenAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketingWalletCosmicSignatureTokenAddressChanged)
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
		it.Event = new(MarketingWalletCosmicSignatureTokenAddressChanged)
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
func (it *MarketingWalletCosmicSignatureTokenAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketingWalletCosmicSignatureTokenAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketingWalletCosmicSignatureTokenAddressChanged represents a CosmicSignatureTokenAddressChanged event raised by the MarketingWallet contract.
type MarketingWalletCosmicSignatureTokenAddressChanged struct {
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCosmicSignatureTokenAddressChanged is a free log retrieval operation binding the contract event 0x9b3eda10f1724a2cf9f7dae4ac263c77908df4d00e92f1377b66fc8be37cd8c1.
//
// Solidity: event CosmicSignatureTokenAddressChanged(address newValue)
func (_MarketingWallet *MarketingWalletFilterer) FilterCosmicSignatureTokenAddressChanged(opts *bind.FilterOpts) (*MarketingWalletCosmicSignatureTokenAddressChangedIterator, error) {

	logs, sub, err := _MarketingWallet.contract.FilterLogs(opts, "CosmicSignatureTokenAddressChanged")
	if err != nil {
		return nil, err
	}
	return &MarketingWalletCosmicSignatureTokenAddressChangedIterator{contract: _MarketingWallet.contract, event: "CosmicSignatureTokenAddressChanged", logs: logs, sub: sub}, nil
}

// WatchCosmicSignatureTokenAddressChanged is a free log subscription operation binding the contract event 0x9b3eda10f1724a2cf9f7dae4ac263c77908df4d00e92f1377b66fc8be37cd8c1.
//
// Solidity: event CosmicSignatureTokenAddressChanged(address newValue)
func (_MarketingWallet *MarketingWalletFilterer) WatchCosmicSignatureTokenAddressChanged(opts *bind.WatchOpts, sink chan<- *MarketingWalletCosmicSignatureTokenAddressChanged) (event.Subscription, error) {

	logs, sub, err := _MarketingWallet.contract.WatchLogs(opts, "CosmicSignatureTokenAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketingWalletCosmicSignatureTokenAddressChanged)
				if err := _MarketingWallet.contract.UnpackLog(event, "CosmicSignatureTokenAddressChanged", log); err != nil {
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

// ParseCosmicSignatureTokenAddressChanged is a log parse operation binding the contract event 0x9b3eda10f1724a2cf9f7dae4ac263c77908df4d00e92f1377b66fc8be37cd8c1.
//
// Solidity: event CosmicSignatureTokenAddressChanged(address newValue)
func (_MarketingWallet *MarketingWalletFilterer) ParseCosmicSignatureTokenAddressChanged(log types.Log) (*MarketingWalletCosmicSignatureTokenAddressChanged, error) {
	event := new(MarketingWalletCosmicSignatureTokenAddressChanged)
	if err := _MarketingWallet.contract.UnpackLog(event, "CosmicSignatureTokenAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketingWalletOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MarketingWallet contract.
type MarketingWalletOwnershipTransferredIterator struct {
	Event *MarketingWalletOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MarketingWalletOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketingWalletOwnershipTransferred)
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
		it.Event = new(MarketingWalletOwnershipTransferred)
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
func (it *MarketingWalletOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketingWalletOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketingWalletOwnershipTransferred represents a OwnershipTransferred event raised by the MarketingWallet contract.
type MarketingWalletOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MarketingWallet *MarketingWalletFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MarketingWalletOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MarketingWallet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MarketingWalletOwnershipTransferredIterator{contract: _MarketingWallet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MarketingWallet *MarketingWalletFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MarketingWalletOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MarketingWallet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketingWalletOwnershipTransferred)
				if err := _MarketingWallet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MarketingWallet *MarketingWalletFilterer) ParseOwnershipTransferred(log types.Log) (*MarketingWalletOwnershipTransferred, error) {
	event := new(MarketingWalletOwnershipTransferred)
	if err := _MarketingWallet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketingWalletRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the MarketingWallet contract.
type MarketingWalletRewardPaidIterator struct {
	Event *MarketingWalletRewardPaid // Event containing the contract specifics and raw log

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
func (it *MarketingWalletRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketingWalletRewardPaid)
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
		it.Event = new(MarketingWalletRewardPaid)
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
func (it *MarketingWalletRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketingWalletRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketingWalletRewardPaid represents a RewardPaid event raised by the MarketingWallet contract.
type MarketingWalletRewardPaid struct {
	MarketerAddress common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed marketerAddress, uint256 amount)
func (_MarketingWallet *MarketingWalletFilterer) FilterRewardPaid(opts *bind.FilterOpts, marketerAddress []common.Address) (*MarketingWalletRewardPaidIterator, error) {

	var marketerAddressRule []interface{}
	for _, marketerAddressItem := range marketerAddress {
		marketerAddressRule = append(marketerAddressRule, marketerAddressItem)
	}

	logs, sub, err := _MarketingWallet.contract.FilterLogs(opts, "RewardPaid", marketerAddressRule)
	if err != nil {
		return nil, err
	}
	return &MarketingWalletRewardPaidIterator{contract: _MarketingWallet.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed marketerAddress, uint256 amount)
func (_MarketingWallet *MarketingWalletFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *MarketingWalletRewardPaid, marketerAddress []common.Address) (event.Subscription, error) {

	var marketerAddressRule []interface{}
	for _, marketerAddressItem := range marketerAddress {
		marketerAddressRule = append(marketerAddressRule, marketerAddressItem)
	}

	logs, sub, err := _MarketingWallet.contract.WatchLogs(opts, "RewardPaid", marketerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketingWalletRewardPaid)
				if err := _MarketingWallet.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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

// ParseRewardPaid is a log parse operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed marketerAddress, uint256 amount)
func (_MarketingWallet *MarketingWalletFilterer) ParseRewardPaid(log types.Log) (*MarketingWalletRewardPaid, error) {
	event := new(MarketingWalletRewardPaid)
	if err := _MarketingWallet.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageSlotMetaData contains all meta data concerning the StorageSlot contract.
var StorageSlotMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x608060405234601d57600e6021565b603e602c823930815050603e90f35b6027565b60405190565b5f80fdfe60806040525f80fdfea2646970667358221220b8186e5d264edd4db6aba734d9e4fc5100dbadc5df307aa34e9c95a01c453cfd64736f6c634300081c0033",
}

// StorageSlotABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageSlotMetaData.ABI instead.
var StorageSlotABI = StorageSlotMetaData.ABI

// StorageSlotBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StorageSlotMetaData.Bin instead.
var StorageSlotBin = StorageSlotMetaData.Bin

// DeployStorageSlot deploys a new Ethereum contract, binding an instance of StorageSlot to it.
func DeployStorageSlot(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StorageSlot, error) {
	parsed, err := StorageSlotMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageSlotBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StorageSlot{StorageSlotCaller: StorageSlotCaller{contract: contract}, StorageSlotTransactor: StorageSlotTransactor{contract: contract}, StorageSlotFilterer: StorageSlotFilterer{contract: contract}}, nil
}

// StorageSlot is an auto generated Go binding around an Ethereum contract.
type StorageSlot struct {
	StorageSlotCaller     // Read-only binding to the contract
	StorageSlotTransactor // Write-only binding to the contract
	StorageSlotFilterer   // Log filterer for contract events
}

// StorageSlotCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageSlotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSlotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageSlotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSlotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageSlotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSlotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSlotSession struct {
	Contract     *StorageSlot      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageSlotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageSlotCallerSession struct {
	Contract *StorageSlotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// StorageSlotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageSlotTransactorSession struct {
	Contract     *StorageSlotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// StorageSlotRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageSlotRaw struct {
	Contract *StorageSlot // Generic contract binding to access the raw methods on
}

// StorageSlotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageSlotCallerRaw struct {
	Contract *StorageSlotCaller // Generic read-only contract binding to access the raw methods on
}

// StorageSlotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageSlotTransactorRaw struct {
	Contract *StorageSlotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorageSlot creates a new instance of StorageSlot, bound to a specific deployed contract.
func NewStorageSlot(address common.Address, backend bind.ContractBackend) (*StorageSlot, error) {
	contract, err := bindStorageSlot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StorageSlot{StorageSlotCaller: StorageSlotCaller{contract: contract}, StorageSlotTransactor: StorageSlotTransactor{contract: contract}, StorageSlotFilterer: StorageSlotFilterer{contract: contract}}, nil
}

// NewStorageSlotCaller creates a new read-only instance of StorageSlot, bound to a specific deployed contract.
func NewStorageSlotCaller(address common.Address, caller bind.ContractCaller) (*StorageSlotCaller, error) {
	contract, err := bindStorageSlot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageSlotCaller{contract: contract}, nil
}

// NewStorageSlotTransactor creates a new write-only instance of StorageSlot, bound to a specific deployed contract.
func NewStorageSlotTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageSlotTransactor, error) {
	contract, err := bindStorageSlot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageSlotTransactor{contract: contract}, nil
}

// NewStorageSlotFilterer creates a new log filterer instance of StorageSlot, bound to a specific deployed contract.
func NewStorageSlotFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageSlotFilterer, error) {
	contract, err := bindStorageSlot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageSlotFilterer{contract: contract}, nil
}

// bindStorageSlot binds a generic wrapper to an already deployed contract.
func bindStorageSlot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StorageSlotMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageSlot *StorageSlotRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageSlot.Contract.StorageSlotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageSlot *StorageSlotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageSlot.Contract.StorageSlotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageSlot *StorageSlotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageSlot.Contract.StorageSlotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageSlot *StorageSlotCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageSlot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageSlot *StorageSlotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageSlot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageSlot *StorageSlotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageSlot.Contract.contract.Transact(opts, method, params...)
}

// TimeMetaData contains all meta data concerning the Time contract.
var TimeMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x608060405234601d57600e6021565b603e602c823930815050603e90f35b6027565b60405190565b5f80fdfe60806040525f80fdfea2646970667358221220b5c0af42af6220cea477ba94ec4920b2bfce70c9785976772f91319f68b1a6d964736f6c634300081c0033",
}

// TimeABI is the input ABI used to generate the binding from.
// Deprecated: Use TimeMetaData.ABI instead.
var TimeABI = TimeMetaData.ABI

// TimeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TimeMetaData.Bin instead.
var TimeBin = TimeMetaData.Bin

// DeployTime deploys a new Ethereum contract, binding an instance of Time to it.
func DeployTime(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Time, error) {
	parsed, err := TimeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TimeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Time{TimeCaller: TimeCaller{contract: contract}, TimeTransactor: TimeTransactor{contract: contract}, TimeFilterer: TimeFilterer{contract: contract}}, nil
}

// Time is an auto generated Go binding around an Ethereum contract.
type Time struct {
	TimeCaller     // Read-only binding to the contract
	TimeTransactor // Write-only binding to the contract
	TimeFilterer   // Log filterer for contract events
}

// TimeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TimeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TimeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TimeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TimeSession struct {
	Contract     *Time             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TimeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TimeCallerSession struct {
	Contract *TimeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TimeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TimeTransactorSession struct {
	Contract     *TimeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TimeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TimeRaw struct {
	Contract *Time // Generic contract binding to access the raw methods on
}

// TimeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TimeCallerRaw struct {
	Contract *TimeCaller // Generic read-only contract binding to access the raw methods on
}

// TimeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TimeTransactorRaw struct {
	Contract *TimeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTime creates a new instance of Time, bound to a specific deployed contract.
func NewTime(address common.Address, backend bind.ContractBackend) (*Time, error) {
	contract, err := bindTime(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Time{TimeCaller: TimeCaller{contract: contract}, TimeTransactor: TimeTransactor{contract: contract}, TimeFilterer: TimeFilterer{contract: contract}}, nil
}

// NewTimeCaller creates a new read-only instance of Time, bound to a specific deployed contract.
func NewTimeCaller(address common.Address, caller bind.ContractCaller) (*TimeCaller, error) {
	contract, err := bindTime(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TimeCaller{contract: contract}, nil
}

// NewTimeTransactor creates a new write-only instance of Time, bound to a specific deployed contract.
func NewTimeTransactor(address common.Address, transactor bind.ContractTransactor) (*TimeTransactor, error) {
	contract, err := bindTime(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TimeTransactor{contract: contract}, nil
}

// NewTimeFilterer creates a new log filterer instance of Time, bound to a specific deployed contract.
func NewTimeFilterer(address common.Address, filterer bind.ContractFilterer) (*TimeFilterer, error) {
	contract, err := bindTime(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TimeFilterer{contract: contract}, nil
}

// bindTime binds a generic wrapper to an already deployed contract.
func bindTime(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TimeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Time *TimeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Time.Contract.TimeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Time *TimeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Time.Contract.TimeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Time *TimeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Time.Contract.TimeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Time *TimeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Time.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Time *TimeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Time.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Time *TimeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Time.Contract.contract.Transact(opts, method, params...)
}
