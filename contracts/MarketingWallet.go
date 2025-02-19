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

// ICosmicSignatureTokenMintSpec is an auto generated low-level Go binding around an user-defined struct.
type ICosmicSignatureTokenMintSpec struct {
	Account common.Address
	Value   *big.Int
}

// IMarketingWalletMetaData contains all meta data concerning the IMarketingWallet contract.
var IMarketingWalletMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"marketerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structICosmicSignatureToken.MintSpec[]\",\"name\":\"specs_\",\"type\":\"tuple[]\"}],\"name\":\"payManyRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"marketerAddresses_\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"payManyRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"marketerAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"payReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractCosmicSignatureToken\",\"name\":\"token_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"marketerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structICosmicSignatureToken.MintSpec[]\",\"name\":\"specs_\",\"type\":\"tuple[]\"}],\"name\":\"payManyRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"marketerAddresses_\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"payManyRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"marketerAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"payReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractCosmicSignatureToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523461004d5761001961001461011e565b61025e565b610021610052565b610ddf610404823960805181818161030b0152818161062b015281816109750152610bfa0152610ddf90f35b610058565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906100849061005c565b810190811060018060401b0382111761009c57604052565b610066565b906100b46100ad610052565b928361007a565b565b5f80fd5b60018060a01b031690565b6100ce906100ba565b90565b6100da906100c5565b90565b6100e6816100d1565b036100ed57565b5f80fd5b905051906100fe826100dd565b565b9060208282031261011957610116915f016100f1565b90565b6100b6565b61013c6111e380380380610131816100a1565b928339810190610100565b90565b90565b61015661015161015b926100ba565b61013f565b6100ba565b90565b61016790610142565b90565b6101739061015e565b90565b90565b61018d61018861019292610176565b61013f565b6100ba565b90565b61019e90610179565b90565b60209181520190565b5f7f5468652070726f76696465642061646472657373206973207a65726f2e000000910152565b6101de601d6020926101a1565b6101e7816101aa565b0190565b6102009060208101905f8183039101526101d1565b90565b61020c8161016a565b61022661022061021b5f610195565b6100c5565b916100c5565b146102365761023490610259565b565b61023e610052565b63eac0d38960e01b815280610255600482016101eb565b0390fd5b608052565b6102779061027261026d610312565b610279565b610203565b565b61028290610284565b565b61028d9061028f565b565b610298906102bc565b565b6102a3906100c5565b9052565b91906102ba905f6020850194019061029a565b565b806102d76102d16102cc5f610195565b6100c5565b916100c5565b146102e7576102e5906103a4565b565b61030a6102f35f610195565b5f918291631e4fbdf760e01b8352600483016102a7565b0390fd5b5f90565b61031a61030e565b503390565b5f1c90565b60018060a01b031690565b61033b6103409161031f565b610324565b90565b61034d905461032f565b90565b5f1b90565b9061036660018060a01b0391610350565b9181191691161790565b6103799061015e565b90565b90565b9061039461038f61039b92610370565b61037c565b8254610355565b9055565b5f0190565b6103ad5f610343565b6103b7825f61037f565b906103eb6103e57f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e093610370565b91610370565b916103f4610052565b806103fe8161039f565b0390a356fe60806040526004361015610013575b610473565b61001d5f3561008c565b8063715018a6146100875780638b8afcd7146100825780638da5cb5b1461007d578063c603317f14610078578063f2fde38b14610073578063fc0c546a1461006e5763ffcbf0b40361000e5761043f565b610386565b6102d6565b610284565b6101ce565b610178565b6100b4565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f9103126100aa57565b61009c565b5f0190565b346100e2576100c43660046100a0565b6100cc6104c7565b6100d4610092565b806100de816100af565b0390f35b610098565b5f80fd5b60018060a01b031690565b6100ff906100eb565b90565b61010b816100f6565b0361011257565b5f80fd5b9050359061012382610102565b565b90565b61013181610125565b0361013857565b5f80fd5b9050359061014982610128565b565b91906040838203126101735780610167610170925f8601610116565b9360200161013c565b90565b61009c565b346101a75761019161018b36600461014b565b906106c0565b610199610092565b806101a3816100af565b0390f35b610098565b6101b5906100f6565b9052565b91906101cc905f602085019401906101ac565b565b346101fe576101de3660046100a0565b6101fa6101e9610701565b6101f1610092565b918291826101b9565b0390f35b610098565b5f80fd5b5f80fd5b5f80fd5b909182601f830112156102495781359167ffffffffffffffff831161024457602001926040830284011161023f57565b61020b565b610207565b610203565b9060208282031261027f575f82013567ffffffffffffffff811161027a57610276920161020f565b9091565b6100e7565b61009c565b346102b35761029d61029736600461024e565b90610a19565b6102a5610092565b806102af816100af565b0390f35b610098565b906020828203126102d1576102ce915f01610116565b90565b61009c565b34610304576102ee6102e93660046102b8565b610a8a565b6102f6610092565b80610300816100af565b0390f35b610098565b7f000000000000000000000000000000000000000000000000000000000000000090565b90565b61034461033f610349926100eb565b61032d565b6100eb565b90565b61035590610330565b90565b6103619061034c565b90565b61036d90610358565b9052565b9190610384905f60208501940190610364565b565b346103b6576103963660046100a0565b6103b26103a1610309565b6103a9610092565b91829182610371565b0390f35b610098565b909182601f830112156103f55781359167ffffffffffffffff83116103f05760200192602083028401116103eb57565b61020b565b610207565b610203565b9160408383031261043a575f83013567ffffffffffffffff811161043557610427836104329286016103bb565b93909460200161013c565b90565b6100e7565b61009c565b3461046e576104586104523660046103fa565b91610c9f565b610460610092565b8061046a816100af565b0390f35b610098565b5f80fd5b61047f610cac565b6104876104b4565b565b90565b6104a061049b6104a592610489565b61032d565b6100eb565b90565b6104b19061048c565b90565b6104c56104c05f6104a8565b610d3d565b565b6104cf610477565b565b906104e3916104de610cac565b6105dd565b565b6104ee9061034c565b90565b6104fa90610125565b9052565b9190610511905f602085019401906104f1565b565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b9061053f90610517565b810190811067ffffffffffffffff82111761055957604052565b610521565b60e01b90565b151590565b61057281610564565b0361057957565b5f80fd5b9050519061058a82610569565b565b906020828203126105a5576105a2915f0161057d565b90565b61009c565b9160206105cb9294936105c460408201965f8301906101ac565b01906104f1565b565b6105d5610092565b3d5f823e3d90fd5b906020908281906106236106117fe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486926104e5565b9261061a610092565b918291826104fe565b0390a261064f7f0000000000000000000000000000000000000000000000000000000000000000610358565b6106725f63a9059cbb95939561067d610666610092565b9788968795869461055e565b8452600484016105aa565b03925af180156106bb5761068f575b50565b6106af9060203d81116106b4575b6106a78183610535565b81019061058c565b61068c565b503d61069d565b6105cd565b906106ca916104d1565b565b5f90565b5f1c90565b60018060a01b031690565b6106ec6106f1916106d0565b6106d5565b90565b6106fe90546106e0565b90565b6107096106cc565b506107135f6106f4565b90565b9061072891610723610cac565b6108d4565b565b5090565b61074261073d61074792610489565b61032d565b610125565b90565b634e487b7160e01b5f52601160045260245ffd5b61076790610125565b5f8114610775576001900390565b61074a565b634e487b7160e01b5f52603260045260245ffd5b919081101561079e576040020190565b61077a565b356107ad81610102565b90565b356107ba81610128565b90565b5f9103126107c757565b61009c565b60209181520190565b90565b506107e7906020810190610116565b90565b6107f3906100f6565b9052565b5061080690602081019061013c565b90565b61081290610125565b9052565b9060206108416108499361083861082f5f8301836107d8565b5f8601906107ea565b828101906107f7565b910190610809565b565b9061085881604093610816565b0190565b5090565b60400190565b916108748261087a926107cc565b926107d5565b90815f905b82821061088d575050505090565b909192936108af6108a96001926108a4888661085c565b61084b565b95610860565b92019092919261087f565b90916108d19260208301925f818503910152610866565b90565b6108df81839061072a565b5b806108f36108ed5f61072e565b91610125565b111561096f576109029061075e565b9261090f8284869161078e565b610926602061091f5f84016107a3565b92016107b0565b906109666109547fe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486926104e5565b9261095d610092565b918291826104fe565b0390a2926108e0565b506109997f0000000000000000000000000000000000000000000000000000000000000000610358565b91638d6bd91c919092803b15610a14576109c65f80946109d16109ba610092565b9788968795869461055e565b8452600484016108ba565b03925af18015610a0f576109e3575b50565b610a02905f3d8111610a08575b6109fa8183610535565b8101906107bd565b5f6109e0565b503d6109f0565b6105cd565b610513565b90610a2391610716565b565b610a3690610a31610cac565b610a38565b565b80610a53610a4d610a485f6104a8565b6100f6565b916100f6565b14610a6357610a6190610d3d565b565b610a86610a6f5f6104a8565b5f918291631e4fbdf760e01b8352600483016101b9565b0390fd5b610a9390610a25565b565b90610aa89291610aa3610cac565b610b62565b565b5090565b9190811015610abe576020020190565b61077a565b60209181520190565b90565b90610adc816020936107ea565b0190565b60200190565b91610af482610afa92610ac3565b92610acc565b90815f905b828210610b0d575050505090565b90919293610b2f610b29600192610b2488866107d8565b610acf565b95610ae0565b920190929192610aff565b939290610b58602091610b609460408801918883035f8a0152610ae6565b9401906104f1565b565b909192610b70828490610aaa565b5b80610b84610b7e5f61072e565b91610125565b1115610bf257610b939061075e565b90610ba8610ba384868591610aae565b6107a3565b8590610be9610bd77fe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486926104e5565b92610be0610092565b918291826104fe565b0390a290610b71565b509291610c1e7f0000000000000000000000000000000000000000000000000000000000000000610358565b630a8e30d992919392813b15610c9a575f610c4c91610c578296610c40610092565b9889978896879561055e565b855260048501610b3a565b03925af18015610c9557610c69575b50565b610c88905f3d8111610c8e575b610c808183610535565b8101906107bd565b5f610c66565b503d610c76565b6105cd565b610513565b90610caa9291610a95565b565b610cb4610701565b610ccd610cc7610cc2610d9c565b6100f6565b916100f6565b03610cd457565b610cf6610cdf610d9c565b5f91829163118cdaa760e01b8352600483016101b9565b0390fd5b5f1b90565b90610d1060018060a01b0391610cfa565b9181191691161790565b90565b90610d32610d2d610d39926104e5565b610d1a565b8254610cff565b9055565b610d465f6106f4565b610d50825f610d1d565b90610d84610d7e7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0936104e5565b916104e5565b91610d8d610092565b80610d97816100af565b0390a3565b610da46106cc565b50339056fea26469706673582212204b511164a946372f6e52a60db43b9beb3ff6b4284bcae23382ec65323095e5fc64736f6c634300081c0033",
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
