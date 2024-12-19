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

// IStakingWalletRandomWalkNftMetaData contains all meta data concerning the IStakingWalletRandomWalkNft contract.
var IStakingWalletRandomWalkNftMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumCosmicSignatureConstants.NftTypeCode\",\"name\":\"nftTypeCode\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftUnstaked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"numStakedNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"randomNumber_\",\"type\":\"uint256\"}],\"name\":\"pickRandomStakerAddressIfPossible\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nftIds_\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId_\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"stakeActionIds_\",\"type\":\"uint256[]\"}],\"name\":\"unstakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"wasNftUsed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IStakingWalletRandomWalkNftABI is the input ABI used to generate the binding from.
// Deprecated: Use IStakingWalletRandomWalkNftMetaData.ABI instead.
var IStakingWalletRandomWalkNftABI = IStakingWalletRandomWalkNftMetaData.ABI

// IStakingWalletRandomWalkNft is an auto generated Go binding around an Ethereum contract.
type IStakingWalletRandomWalkNft struct {
	IStakingWalletRandomWalkNftCaller     // Read-only binding to the contract
	IStakingWalletRandomWalkNftTransactor // Write-only binding to the contract
	IStakingWalletRandomWalkNftFilterer   // Log filterer for contract events
}

// IStakingWalletRandomWalkNftCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStakingWalletRandomWalkNftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingWalletRandomWalkNftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStakingWalletRandomWalkNftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingWalletRandomWalkNftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStakingWalletRandomWalkNftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingWalletRandomWalkNftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStakingWalletRandomWalkNftSession struct {
	Contract     *IStakingWalletRandomWalkNft // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IStakingWalletRandomWalkNftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStakingWalletRandomWalkNftCallerSession struct {
	Contract *IStakingWalletRandomWalkNftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// IStakingWalletRandomWalkNftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStakingWalletRandomWalkNftTransactorSession struct {
	Contract     *IStakingWalletRandomWalkNftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// IStakingWalletRandomWalkNftRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStakingWalletRandomWalkNftRaw struct {
	Contract *IStakingWalletRandomWalkNft // Generic contract binding to access the raw methods on
}

// IStakingWalletRandomWalkNftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStakingWalletRandomWalkNftCallerRaw struct {
	Contract *IStakingWalletRandomWalkNftCaller // Generic read-only contract binding to access the raw methods on
}

// IStakingWalletRandomWalkNftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStakingWalletRandomWalkNftTransactorRaw struct {
	Contract *IStakingWalletRandomWalkNftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStakingWalletRandomWalkNft creates a new instance of IStakingWalletRandomWalkNft, bound to a specific deployed contract.
func NewIStakingWalletRandomWalkNft(address common.Address, backend bind.ContractBackend) (*IStakingWalletRandomWalkNft, error) {
	contract, err := bindIStakingWalletRandomWalkNft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletRandomWalkNft{IStakingWalletRandomWalkNftCaller: IStakingWalletRandomWalkNftCaller{contract: contract}, IStakingWalletRandomWalkNftTransactor: IStakingWalletRandomWalkNftTransactor{contract: contract}, IStakingWalletRandomWalkNftFilterer: IStakingWalletRandomWalkNftFilterer{contract: contract}}, nil
}

// NewIStakingWalletRandomWalkNftCaller creates a new read-only instance of IStakingWalletRandomWalkNft, bound to a specific deployed contract.
func NewIStakingWalletRandomWalkNftCaller(address common.Address, caller bind.ContractCaller) (*IStakingWalletRandomWalkNftCaller, error) {
	contract, err := bindIStakingWalletRandomWalkNft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletRandomWalkNftCaller{contract: contract}, nil
}

// NewIStakingWalletRandomWalkNftTransactor creates a new write-only instance of IStakingWalletRandomWalkNft, bound to a specific deployed contract.
func NewIStakingWalletRandomWalkNftTransactor(address common.Address, transactor bind.ContractTransactor) (*IStakingWalletRandomWalkNftTransactor, error) {
	contract, err := bindIStakingWalletRandomWalkNft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletRandomWalkNftTransactor{contract: contract}, nil
}

// NewIStakingWalletRandomWalkNftFilterer creates a new log filterer instance of IStakingWalletRandomWalkNft, bound to a specific deployed contract.
func NewIStakingWalletRandomWalkNftFilterer(address common.Address, filterer bind.ContractFilterer) (*IStakingWalletRandomWalkNftFilterer, error) {
	contract, err := bindIStakingWalletRandomWalkNft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletRandomWalkNftFilterer{contract: contract}, nil
}

// bindIStakingWalletRandomWalkNft binds a generic wrapper to an already deployed contract.
func bindIStakingWalletRandomWalkNft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IStakingWalletRandomWalkNftMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStakingWalletRandomWalkNft.Contract.IStakingWalletRandomWalkNftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.IStakingWalletRandomWalkNftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.IStakingWalletRandomWalkNftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStakingWalletRandomWalkNft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.contract.Transact(opts, method, params...)
}

// NumStakedNfts is a free data retrieval call binding the contract method 0xca7c1f92.
//
// Solidity: function numStakedNfts() view returns(uint256)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftCaller) NumStakedNfts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStakingWalletRandomWalkNft.contract.Call(opts, &out, "numStakedNfts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumStakedNfts is a free data retrieval call binding the contract method 0xca7c1f92.
//
// Solidity: function numStakedNfts() view returns(uint256)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftSession) NumStakedNfts() (*big.Int, error) {
	return _IStakingWalletRandomWalkNft.Contract.NumStakedNfts(&_IStakingWalletRandomWalkNft.CallOpts)
}

// NumStakedNfts is a free data retrieval call binding the contract method 0xca7c1f92.
//
// Solidity: function numStakedNfts() view returns(uint256)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftCallerSession) NumStakedNfts() (*big.Int, error) {
	return _IStakingWalletRandomWalkNft.Contract.NumStakedNfts(&_IStakingWalletRandomWalkNft.CallOpts)
}

// PickRandomStakerAddressIfPossible is a free data retrieval call binding the contract method 0xd987b38b.
//
// Solidity: function pickRandomStakerAddressIfPossible(uint256 randomNumber_) view returns(address)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftCaller) PickRandomStakerAddressIfPossible(opts *bind.CallOpts, randomNumber_ *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IStakingWalletRandomWalkNft.contract.Call(opts, &out, "pickRandomStakerAddressIfPossible", randomNumber_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PickRandomStakerAddressIfPossible is a free data retrieval call binding the contract method 0xd987b38b.
//
// Solidity: function pickRandomStakerAddressIfPossible(uint256 randomNumber_) view returns(address)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftSession) PickRandomStakerAddressIfPossible(randomNumber_ *big.Int) (common.Address, error) {
	return _IStakingWalletRandomWalkNft.Contract.PickRandomStakerAddressIfPossible(&_IStakingWalletRandomWalkNft.CallOpts, randomNumber_)
}

// PickRandomStakerAddressIfPossible is a free data retrieval call binding the contract method 0xd987b38b.
//
// Solidity: function pickRandomStakerAddressIfPossible(uint256 randomNumber_) view returns(address)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftCallerSession) PickRandomStakerAddressIfPossible(randomNumber_ *big.Int) (common.Address, error) {
	return _IStakingWalletRandomWalkNft.Contract.PickRandomStakerAddressIfPossible(&_IStakingWalletRandomWalkNft.CallOpts, randomNumber_)
}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(uint256)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftCaller) WasNftUsed(opts *bind.CallOpts, nftId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IStakingWalletRandomWalkNft.contract.Call(opts, &out, "wasNftUsed", nftId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(uint256)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftSession) WasNftUsed(nftId_ *big.Int) (*big.Int, error) {
	return _IStakingWalletRandomWalkNft.Contract.WasNftUsed(&_IStakingWalletRandomWalkNft.CallOpts, nftId_)
}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(uint256)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftCallerSession) WasNftUsed(nftId_ *big.Int) (*big.Int, error) {
	return _IStakingWalletRandomWalkNft.Contract.WasNftUsed(&_IStakingWalletRandomWalkNft.CallOpts, nftId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactor) Stake(opts *bind.TransactOpts, nftId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.contract.Transact(opts, "stake", nftId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftSession) Stake(nftId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.Stake(&_IStakingWalletRandomWalkNft.TransactOpts, nftId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactorSession) Stake(nftId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.Stake(&_IStakingWalletRandomWalkNft.TransactOpts, nftId_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactor) StakeMany(opts *bind.TransactOpts, nftIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.contract.Transact(opts, "stakeMany", nftIds_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftSession) StakeMany(nftIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.StakeMany(&_IStakingWalletRandomWalkNft.TransactOpts, nftIds_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactorSession) StakeMany(nftIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.StakeMany(&_IStakingWalletRandomWalkNft.TransactOpts, nftIds_)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactor) Unstake(opts *bind.TransactOpts, stakeActionId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.contract.Transact(opts, "unstake", stakeActionId_)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftSession) Unstake(stakeActionId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.Unstake(&_IStakingWalletRandomWalkNft.TransactOpts, stakeActionId_)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactorSession) Unstake(stakeActionId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.Unstake(&_IStakingWalletRandomWalkNft.TransactOpts, stakeActionId_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactor) UnstakeMany(opts *bind.TransactOpts, stakeActionIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.contract.Transact(opts, "unstakeMany", stakeActionIds_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftSession) UnstakeMany(stakeActionIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.UnstakeMany(&_IStakingWalletRandomWalkNft.TransactOpts, stakeActionIds_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactorSession) UnstakeMany(stakeActionIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.UnstakeMany(&_IStakingWalletRandomWalkNft.TransactOpts, stakeActionIds_)
}

// IStakingWalletRandomWalkNftNftStakedIterator is returned from FilterNftStaked and is used to iterate over the raw logs and unpacked data for NftStaked events raised by the IStakingWalletRandomWalkNft contract.
type IStakingWalletRandomWalkNftNftStakedIterator struct {
	Event *IStakingWalletRandomWalkNftNftStaked // Event containing the contract specifics and raw log

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
func (it *IStakingWalletRandomWalkNftNftStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingWalletRandomWalkNftNftStaked)
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
		it.Event = new(IStakingWalletRandomWalkNftNftStaked)
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
func (it *IStakingWalletRandomWalkNftNftStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingWalletRandomWalkNftNftStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingWalletRandomWalkNftNftStaked represents a NftStaked event raised by the IStakingWalletRandomWalkNft contract.
type IStakingWalletRandomWalkNftNftStaked struct {
	StakeActionId *big.Int
	NftTypeCode   uint8
	NftId         *big.Int
	StakerAddress common.Address
	NumStakedNfts *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNftStaked is a free log retrieval operation binding the contract event 0xcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f00829.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint8 nftTypeCode, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftFilterer) FilterNftStaked(opts *bind.FilterOpts, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (*IStakingWalletRandomWalkNftNftStakedIterator, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}

	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _IStakingWalletRandomWalkNft.contract.FilterLogs(opts, "NftStaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletRandomWalkNftNftStakedIterator{contract: _IStakingWalletRandomWalkNft.contract, event: "NftStaked", logs: logs, sub: sub}, nil
}

// WatchNftStaked is a free log subscription operation binding the contract event 0xcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f00829.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint8 nftTypeCode, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftFilterer) WatchNftStaked(opts *bind.WatchOpts, sink chan<- *IStakingWalletRandomWalkNftNftStaked, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (event.Subscription, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}

	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _IStakingWalletRandomWalkNft.contract.WatchLogs(opts, "NftStaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingWalletRandomWalkNftNftStaked)
				if err := _IStakingWalletRandomWalkNft.contract.UnpackLog(event, "NftStaked", log); err != nil {
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

// ParseNftStaked is a log parse operation binding the contract event 0xcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f00829.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint8 nftTypeCode, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftFilterer) ParseNftStaked(log types.Log) (*IStakingWalletRandomWalkNftNftStaked, error) {
	event := new(IStakingWalletRandomWalkNftNftStaked)
	if err := _IStakingWalletRandomWalkNft.contract.UnpackLog(event, "NftStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingWalletRandomWalkNftNftUnstakedIterator is returned from FilterNftUnstaked and is used to iterate over the raw logs and unpacked data for NftUnstaked events raised by the IStakingWalletRandomWalkNft contract.
type IStakingWalletRandomWalkNftNftUnstakedIterator struct {
	Event *IStakingWalletRandomWalkNftNftUnstaked // Event containing the contract specifics and raw log

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
func (it *IStakingWalletRandomWalkNftNftUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingWalletRandomWalkNftNftUnstaked)
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
		it.Event = new(IStakingWalletRandomWalkNftNftUnstaked)
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
func (it *IStakingWalletRandomWalkNftNftUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingWalletRandomWalkNftNftUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingWalletRandomWalkNftNftUnstaked represents a NftUnstaked event raised by the IStakingWalletRandomWalkNft contract.
type IStakingWalletRandomWalkNftNftUnstaked struct {
	StakeActionId *big.Int
	NftId         *big.Int
	StakerAddress common.Address
	NumStakedNfts *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNftUnstaked is a free log retrieval operation binding the contract event 0x1792a7a9d5e2f53a8d06f2ae40b4446d91b503e84dd7f6307f40cdeb11541668.
//
// Solidity: event NftUnstaked(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftFilterer) FilterNftUnstaked(opts *bind.FilterOpts, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (*IStakingWalletRandomWalkNftNftUnstakedIterator, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _IStakingWalletRandomWalkNft.contract.FilterLogs(opts, "NftUnstaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletRandomWalkNftNftUnstakedIterator{contract: _IStakingWalletRandomWalkNft.contract, event: "NftUnstaked", logs: logs, sub: sub}, nil
}

// WatchNftUnstaked is a free log subscription operation binding the contract event 0x1792a7a9d5e2f53a8d06f2ae40b4446d91b503e84dd7f6307f40cdeb11541668.
//
// Solidity: event NftUnstaked(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftFilterer) WatchNftUnstaked(opts *bind.WatchOpts, sink chan<- *IStakingWalletRandomWalkNftNftUnstaked, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (event.Subscription, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _IStakingWalletRandomWalkNft.contract.WatchLogs(opts, "NftUnstaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingWalletRandomWalkNftNftUnstaked)
				if err := _IStakingWalletRandomWalkNft.contract.UnpackLog(event, "NftUnstaked", log); err != nil {
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

// ParseNftUnstaked is a log parse operation binding the contract event 0x1792a7a9d5e2f53a8d06f2ae40b4446d91b503e84dd7f6307f40cdeb11541668.
//
// Solidity: event NftUnstaked(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftFilterer) ParseNftUnstaked(log types.Log) (*IStakingWalletRandomWalkNftNftUnstaked, error) {
	event := new(IStakingWalletRandomWalkNftNftUnstaked)
	if err := _IStakingWalletRandomWalkNft.contract.UnpackLog(event, "NftUnstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletRandomWalkNftMetaData contains all meta data concerning the StakingWalletRandomWalkNft contract.
var StakingWalletRandomWalkNftMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractRandomWalkNFT\",\"name\":\"randomWalkNft_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"name\":\"NftOneTimeStaking\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"requesterAddress\",\"type\":\"address\"}],\"name\":\"NftStakeActionAccessDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"}],\"name\":\"NftStakeActionInvalidId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumCosmicSignatureConstants.NftTypeCode\",\"name\":\"nftTypeCode\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftUnstaked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"actionCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakedNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"randomNumber_\",\"type\":\"uint256\"}],\"name\":\"pickRandomStakerAddressIfPossible\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomWalkNft\",\"outputs\":[{\"internalType\":\"contractRandomWalkNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeActionIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeActions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftOwnerAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nftIds_\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId_\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"stakeActionIds_\",\"type\":\"uint256[]\"}],\"name\":\"unstakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"wasNftUsed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040523461004657610019610014610117565b610257565b61002161004b565b61138f610263823960805181818161036001528181610d0d0152611180015261138f90f35b610051565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b9061007d90610055565b810190811060018060401b0382111761009557604052565b61005f565b906100ad6100a661004b565b9283610073565b565b5f80fd5b60018060a01b031690565b6100c7906100b3565b90565b6100d3906100be565b90565b6100df816100ca565b036100e657565b5f80fd5b905051906100f7826100d6565b565b906020828203126101125761010f915f016100ea565b90565b6100af565b6101356115f28038038061012a8161009a565b9283398101906100f9565b90565b90565b61014f61014a610154926100b3565b610138565b6100b3565b90565b6101609061013b565b90565b61016c90610157565b90565b90565b61018661018161018b9261016f565b610138565b6100b3565b90565b61019790610172565b90565b60209181520190565b5f7f5468652070726f76696465642061646472657373206973207a65726f2e000000910152565b6101d7601d60209261019a565b6101e0816101a3565b0190565b6101f99060208101905f8183039101526101ca565b90565b61020581610163565b61021f6102196102145f61018e565b6100be565b916100be565b1461022f5761022d90610252565b565b61023761004b565b63eac0d38960e01b81528061024e600482016101e4565b0390fd5b608052565b610260906101fc565b56fe60806040526004361015610013575b6106ab565b61001d5f356100cc565b80630d50c189146100c75780632e17de78146100c257806357951c74146100bd57806360294405146100b8578063755b4ef7146100b3578063a2b136fb146100ae578063a694fc3a146100a9578063ca7c1f92146100a4578063d987b38b1461009f578063fdbd98b01461009a5763fe939afc0361000e57610677565b610642565b6105f6565b6105ac565b610579565b610541565b6103e6565b61031a565b610237565b6101e2565b61016a565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f80fd5b909182601f8301121561012a5781359167ffffffffffffffff831161012557602001926020830284011161012057565b6100ec565b6100e8565b6100e4565b90602082820312610160575f82013567ffffffffffffffff811161015b5761015792016100f0565b9091565b6100e0565b6100dc565b5f0190565b346101995761018361017d36600461012f565b90610703565b61018b6100d2565b8061019581610165565b0390f35b6100d8565b90565b6101aa8161019e565b036101b157565b5f80fd5b905035906101c2826101a1565b565b906020828203126101dd576101da915f016101b5565b90565b6100dc565b34610210576101fa6101f53660046101c4565b610b82565b6102026100d2565b8061020c81610165565b0390f35b6100d8565b61021e9061019e565b9052565b9190610235905f60208501940190610215565b565b346102675761026361025261024d3660046101c4565b610e42565b61025a6100d2565b91829182610222565b0390f35b6100d8565b634e487b7160e01b5f52603260045260245ffd5b506801000000000000000090565b90565b61029a81610280565b8210156102b4576102ac60019161028e565b910201905f90565b61026c565b1c90565b90565b6102d09060086102d593026102b9565b6102bd565b90565b906102e391546102c0565b90565b680400000000000000026102f981610280565b821015610316576103139161030d91610291565b906102d8565b90565b5f80fd5b3461034a576103466103356103303660046101c4565b6102e6565b61033d6100d2565b91829182610222565b0390f35b6100d8565b5f91031261035957565b6100dc565b7f000000000000000000000000000000000000000000000000000000000000000090565b60018060a01b031690565b90565b6103a461039f6103a992610382565b61038d565b610382565b90565b6103b590610390565b90565b6103c1906103ac565b90565b6103cd906103b8565b9052565b91906103e4905f602085019401906103c4565b565b34610416576103f636600461034f565b61041261040161035e565b6104096100d2565b918291826103d1565b0390f35b6100d8565b506801000000000000000090565b90565b6104358161041b565b82101561044f57610447600391610429565b910201905f90565b61026c565b5f1c90565b61046561046a91610454565b6102bd565b90565b6104779054610459565b90565b60018060a01b031690565b61049161049691610454565b61047a565b90565b6104a39054610485565b90565b68010000000000000002906104ba8261041b565b8110156104f2576104ca9161042c565b506104d65f820161046d565b916104ef60026104e86001850161046d565b9301610499565b90565b5f80fd5b6104ff90610382565b90565b61050b906104f6565b9052565b60409061053861053f949695939661052e60608401985f850190610215565b6020830190610215565b0190610502565b565b346105745761057061055c6105573660046101c4565b6104a6565b6105679391936100d2565b9384938461050f565b0390f35b6100d8565b346105a75761059161058c3660046101c4565b61100d565b6105996100d2565b806105a381610165565b0390f35b6100d8565b346105dc576105bc36600461034f565b6105d86105c761122f565b6105cf6100d2565b91829182610222565b0390f35b6100d8565b91906105f4905f60208501940190610502565b565b346106265761062261061161060c3660046101c4565b61127a565b6106196100d2565b918291826105e1565b0390f35b6100d8565b61063f680100000000000000015f906102d8565b90565b346106725761065236600461034f565b61066e61065d61062b565b6106656100d2565b91829182610222565b0390f35b6100d8565b346106a65761069061068a36600461012f565b906112f9565b6106986100d2565b806106a281610165565b0390f35b6100d8565b5f80fd5b90565b6106c66106c16106cb926106af565b61038d565b61019e565b90565b60016106da910161019e565b90565b5090565b91908110156106f1576020020190565b61026c565b35610700816101a1565b90565b91909161070f5f6106b2565b5b8061072d6107276107228588906106dd565b61019e565b9161019e565b101561075d576107589061075361074e610749858885916106e1565b6106f6565b610b82565b6106ce565b610710565b50509050565b90565b906107709061019e565b9052565b9061077e906104f6565b9052565b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906107aa90610782565b810190811067ffffffffffffffff8211176107c457604052565b61078c565b906107dc6107d56100d2565b92836107a0565b565b6107e860606107c9565b90565b9061083a61083160026107fc6107de565b9461081361080b5f830161046d565b5f8801610766565b61082b6108226001830161046d565b60208801610766565b01610499565b60408401610774565b565b610845906107eb565b90565b61085290516104f6565b90565b61086961086461086e926106af565b61038d565b610382565b90565b61087a90610855565b90565b60209181520190565b5f7f496e76616c6964204e4654207374616b6520616374696f6e2049442e00000000910152565b6108ba601c60209261087d565b6108c381610886565b0190565b91906108ea9060206108e2604086018681035f8801526108ad565b940190610215565b565b60207f6e7374616b652069742e00000000000000000000000000000000000000000000917f4f6e6c79204e4654206f776e6572206973207065726d697474656420746f20755f8201520152565b610946602a60409261087d565b61094f816108ec565b0190565b91604061098492949361097d610972606083018381035f850152610939565b966020830190610215565b0190610502565b565b90565b61099d6109986109a292610986565b61038d565b61019e565b90565b634e487b7160e01b5f52601160045260245ffd5b6109c86109ce9193929361019e565b9261019e565b82039182116109d957565b6109a5565b5f1b90565b906109ef5f19916109de565b9181191691161790565b610a0d610a08610a129261019e565b61038d565b61019e565b90565b90565b90610a2d610a28610a34926109f9565b610a15565b82546109e3565b9055565b610a42905161019e565b90565b1b90565b91906008610a64910291610a5e5f1984610a45565b92610a45565b9181191691161790565b9190610a84610a7f610a8c936109f9565b610a15565b908354610a49565b9055565b5f90565b610aa691610aa0610a90565b91610a6e565b565b91906008610ac8910291610ac260018060a01b0384610a45565b92610a45565b9181191691161790565b610adb906103ac565b90565b90565b9190610af7610af2610aff93610ad2565b610ade565b908354610aa8565b9055565b5f90565b610b1991610b13610b03565b91610ae1565b565b610b24906103ac565b90565b5f80fd5b60e01b90565b5f910312610b3b57565b6100dc565b604090610b69610b709496959396610b5f60608401985f850190610502565b6020830190610502565b0190610215565b565b610b7a6100d2565b3d5f823e3d90fd5b610b9f610b9968010000000000000002839061042c565b50610763565b610ba88161083c565b9133610bc7610bc1610bbc60408701610848565b6104f6565b916104f6565b03610dc757610c87610beb610bdb5f61046d565b610be56001610989565b906109b9565b92610bf6845f610a18565b610c635f6002610c19610c13680400000000000000028990610291565b906102d8565b93610c44610c28848b01610a38565b84610c3d68010000000000000002899061042c565b5001610a18565b610c5083808301610a94565b610c5d8360018301610a94565b01610b07565b610c8168040000000000000002610c7b5f8801610a38565b90610291565b90610a6e565b610ca4610c9e680400000000000000028490610291565b90610a94565b90610cb160208401610a38565b339192610d05610cf3610ced610ce77f1792a7a9d5e2f53a8d06f2ae40b4446d91b503e84dd7f6307f40cdeb11541668946109f9565b946109f9565b94610ad2565b94610cfc6100d2565b91829182610222565b0390a4610d317f00000000000000000000000000000000000000000000000000000000000000006103b8565b6323b872dd90610d4030610b1b565b90610d4e6020339501610a38565b92813b15610dc2575f610d7491610d7f8296610d686100d2565b98899788968795610b2b565b855260048501610b40565b03925af18015610dbd57610d91575b50565b610db0905f3d8111610db6575b610da881836107a0565b810190610b31565b5f610d8e565b503d610d9e565b610b72565b610b27565b610dd360408401610848565b610ded610de7610de25f610871565b6104f6565b916104f6565b14155f14610e1c573390610e18610e026100d2565b9283926348aca7ef60e11b845260048401610953565b0390fd5b610e3e90610e286100d2565b91829163023df6b160e21b8352600483016108c7565b0390fd5b610e59610e5f91610e51610a90565b506001610291565b906102d8565b90565b60407f206f6e6c79206f6e63652e000000000000000000000000000000000000000000917f54686973204e46542068617320616c7265616479206265656e207374616b65645f8201527f2e20416e204e465420697320616c6c6f77656420746f206265207374616b656460208201520152565b610ee2604b60609261087d565b610eeb81610e62565b0190565b9190610f12906020610f0a604086018681035f880152610ed5565b940190610215565b565b15610f1c5750565b610f3e90610f286100d2565b918291633b471d1f60e21b835260048301610eef565b0390fd5b610f51610f579193929361019e565b9261019e565b8201809211610f6257565b6109a5565b90610f7860018060a01b03916109de565b9181191691161790565b90610f97610f92610f9e92610ad2565b610ade565b8254610f67565b9055565b634e487b7160e01b5f52602160045260245ffd5b60031115610fc057565b610fa2565b90610fcf82610fb6565b565b610fda90610fc5565b90565b610fe690610fd1565b9052565b91602061100b92949361100460408201965f830190610fdd565b0190610215565b565b61103f61102561101f60018490610291565b906102d8565b6110376110315f6106b2565b9161019e565b148290610f14565b61105e61104c6001610989565b61105860018490610291565b90610a6e565b6110836110736801000000000000000161046d565b61107d6001610989565b90610f42565b6110968168010000000000000001610a18565b6111146110b66110b068010000000000000002849061042c565b50610763565b6110c38460018301610a18565b6110d03360028301610f82565b6110e66110dc5f61046d565b915f839101610a18565b611104836110fe680400000000000000028490610291565b90610a6e565b61110e6001610989565b90610f42565b9061111f825f610a18565b906002918390339261116361115d6111577fcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f00829946109f9565b946109f9565b94610ad2565b9461117861116f6100d2565b92839283610fea565b0390a46111a47f00000000000000000000000000000000000000000000000000000000000000006103b8565b6323b872dd9033906111b530610b1b565b9392813b1561122a575f6111dc916111e782966111d06100d2565b98899788968795610b2b565b855260048501610b40565b03925af18015611225576111f9575b50565b611218905f3d811161121e575b61121081836107a0565b810190610b31565b5f6111f6565b503d611206565b610b72565b610b27565b611237610a90565b506112415f61046d565b90565b634e487b7160e01b5f52601260045260245ffd5b61126461126a9161019e565b9161019e565b908115611275570690565b611244565b611282610b03565b5061128c5f61046d565b90816112a061129a5f6106b2565b9161019e565b146112eb576112e16112d26112cc6112bd6112e895600295611258565b68040000000000000002610291565b906102d8565b6801000000000000000261042c565b5001610499565b90565b50506112f65f610871565b90565b9190916113055f6106b2565b5b8061132361131d6113188588906106dd565b61019e565b9161019e565b10156113535761134e9061134961134461133f858885916106e1565b6106f6565b61100d565b6106ce565b611306565b5050905056fea264697066735822122055b3c31065cd72acde192874f2a61a6de442496e249435a0bcc8443e75e8826564736f6c634300081c0033",
}

// StakingWalletRandomWalkNftABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingWalletRandomWalkNftMetaData.ABI instead.
var StakingWalletRandomWalkNftABI = StakingWalletRandomWalkNftMetaData.ABI

// StakingWalletRandomWalkNftBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingWalletRandomWalkNftMetaData.Bin instead.
var StakingWalletRandomWalkNftBin = StakingWalletRandomWalkNftMetaData.Bin

// DeployStakingWalletRandomWalkNft deploys a new Ethereum contract, binding an instance of StakingWalletRandomWalkNft to it.
func DeployStakingWalletRandomWalkNft(auth *bind.TransactOpts, backend bind.ContractBackend, randomWalkNft_ common.Address) (common.Address, *types.Transaction, *StakingWalletRandomWalkNft, error) {
	parsed, err := StakingWalletRandomWalkNftMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingWalletRandomWalkNftBin), backend, randomWalkNft_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StakingWalletRandomWalkNft{StakingWalletRandomWalkNftCaller: StakingWalletRandomWalkNftCaller{contract: contract}, StakingWalletRandomWalkNftTransactor: StakingWalletRandomWalkNftTransactor{contract: contract}, StakingWalletRandomWalkNftFilterer: StakingWalletRandomWalkNftFilterer{contract: contract}}, nil
}

// StakingWalletRandomWalkNft is an auto generated Go binding around an Ethereum contract.
type StakingWalletRandomWalkNft struct {
	StakingWalletRandomWalkNftCaller     // Read-only binding to the contract
	StakingWalletRandomWalkNftTransactor // Write-only binding to the contract
	StakingWalletRandomWalkNftFilterer   // Log filterer for contract events
}

// StakingWalletRandomWalkNftCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingWalletRandomWalkNftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletRandomWalkNftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingWalletRandomWalkNftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletRandomWalkNftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingWalletRandomWalkNftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletRandomWalkNftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingWalletRandomWalkNftSession struct {
	Contract     *StakingWalletRandomWalkNft // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// StakingWalletRandomWalkNftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingWalletRandomWalkNftCallerSession struct {
	Contract *StakingWalletRandomWalkNftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// StakingWalletRandomWalkNftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingWalletRandomWalkNftTransactorSession struct {
	Contract     *StakingWalletRandomWalkNftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// StakingWalletRandomWalkNftRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingWalletRandomWalkNftRaw struct {
	Contract *StakingWalletRandomWalkNft // Generic contract binding to access the raw methods on
}

// StakingWalletRandomWalkNftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingWalletRandomWalkNftCallerRaw struct {
	Contract *StakingWalletRandomWalkNftCaller // Generic read-only contract binding to access the raw methods on
}

// StakingWalletRandomWalkNftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingWalletRandomWalkNftTransactorRaw struct {
	Contract *StakingWalletRandomWalkNftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingWalletRandomWalkNft creates a new instance of StakingWalletRandomWalkNft, bound to a specific deployed contract.
func NewStakingWalletRandomWalkNft(address common.Address, backend bind.ContractBackend) (*StakingWalletRandomWalkNft, error) {
	contract, err := bindStakingWalletRandomWalkNft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRandomWalkNft{StakingWalletRandomWalkNftCaller: StakingWalletRandomWalkNftCaller{contract: contract}, StakingWalletRandomWalkNftTransactor: StakingWalletRandomWalkNftTransactor{contract: contract}, StakingWalletRandomWalkNftFilterer: StakingWalletRandomWalkNftFilterer{contract: contract}}, nil
}

// NewStakingWalletRandomWalkNftCaller creates a new read-only instance of StakingWalletRandomWalkNft, bound to a specific deployed contract.
func NewStakingWalletRandomWalkNftCaller(address common.Address, caller bind.ContractCaller) (*StakingWalletRandomWalkNftCaller, error) {
	contract, err := bindStakingWalletRandomWalkNft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRandomWalkNftCaller{contract: contract}, nil
}

// NewStakingWalletRandomWalkNftTransactor creates a new write-only instance of StakingWalletRandomWalkNft, bound to a specific deployed contract.
func NewStakingWalletRandomWalkNftTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingWalletRandomWalkNftTransactor, error) {
	contract, err := bindStakingWalletRandomWalkNft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRandomWalkNftTransactor{contract: contract}, nil
}

// NewStakingWalletRandomWalkNftFilterer creates a new log filterer instance of StakingWalletRandomWalkNft, bound to a specific deployed contract.
func NewStakingWalletRandomWalkNftFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingWalletRandomWalkNftFilterer, error) {
	contract, err := bindStakingWalletRandomWalkNft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRandomWalkNftFilterer{contract: contract}, nil
}

// bindStakingWalletRandomWalkNft binds a generic wrapper to an already deployed contract.
func bindStakingWalletRandomWalkNft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingWalletRandomWalkNftMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingWalletRandomWalkNft.Contract.StakingWalletRandomWalkNftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.StakingWalletRandomWalkNftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.StakingWalletRandomWalkNftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingWalletRandomWalkNft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.contract.Transact(opts, method, params...)
}

// ActionCounter is a free data retrieval call binding the contract method 0xfdbd98b0.
//
// Solidity: function actionCounter() view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCaller) ActionCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletRandomWalkNft.contract.Call(opts, &out, "actionCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActionCounter is a free data retrieval call binding the contract method 0xfdbd98b0.
//
// Solidity: function actionCounter() view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) ActionCounter() (*big.Int, error) {
	return _StakingWalletRandomWalkNft.Contract.ActionCounter(&_StakingWalletRandomWalkNft.CallOpts)
}

// ActionCounter is a free data retrieval call binding the contract method 0xfdbd98b0.
//
// Solidity: function actionCounter() view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCallerSession) ActionCounter() (*big.Int, error) {
	return _StakingWalletRandomWalkNft.Contract.ActionCounter(&_StakingWalletRandomWalkNft.CallOpts)
}

// NumStakedNfts is a free data retrieval call binding the contract method 0xca7c1f92.
//
// Solidity: function numStakedNfts() view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCaller) NumStakedNfts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletRandomWalkNft.contract.Call(opts, &out, "numStakedNfts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumStakedNfts is a free data retrieval call binding the contract method 0xca7c1f92.
//
// Solidity: function numStakedNfts() view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) NumStakedNfts() (*big.Int, error) {
	return _StakingWalletRandomWalkNft.Contract.NumStakedNfts(&_StakingWalletRandomWalkNft.CallOpts)
}

// NumStakedNfts is a free data retrieval call binding the contract method 0xca7c1f92.
//
// Solidity: function numStakedNfts() view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCallerSession) NumStakedNfts() (*big.Int, error) {
	return _StakingWalletRandomWalkNft.Contract.NumStakedNfts(&_StakingWalletRandomWalkNft.CallOpts)
}

// PickRandomStakerAddressIfPossible is a free data retrieval call binding the contract method 0xd987b38b.
//
// Solidity: function pickRandomStakerAddressIfPossible(uint256 randomNumber_) view returns(address)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCaller) PickRandomStakerAddressIfPossible(opts *bind.CallOpts, randomNumber_ *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StakingWalletRandomWalkNft.contract.Call(opts, &out, "pickRandomStakerAddressIfPossible", randomNumber_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PickRandomStakerAddressIfPossible is a free data retrieval call binding the contract method 0xd987b38b.
//
// Solidity: function pickRandomStakerAddressIfPossible(uint256 randomNumber_) view returns(address)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) PickRandomStakerAddressIfPossible(randomNumber_ *big.Int) (common.Address, error) {
	return _StakingWalletRandomWalkNft.Contract.PickRandomStakerAddressIfPossible(&_StakingWalletRandomWalkNft.CallOpts, randomNumber_)
}

// PickRandomStakerAddressIfPossible is a free data retrieval call binding the contract method 0xd987b38b.
//
// Solidity: function pickRandomStakerAddressIfPossible(uint256 randomNumber_) view returns(address)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCallerSession) PickRandomStakerAddressIfPossible(randomNumber_ *big.Int) (common.Address, error) {
	return _StakingWalletRandomWalkNft.Contract.PickRandomStakerAddressIfPossible(&_StakingWalletRandomWalkNft.CallOpts, randomNumber_)
}

// RandomWalkNft is a free data retrieval call binding the contract method 0x755b4ef7.
//
// Solidity: function randomWalkNft() view returns(address)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCaller) RandomWalkNft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingWalletRandomWalkNft.contract.Call(opts, &out, "randomWalkNft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RandomWalkNft is a free data retrieval call binding the contract method 0x755b4ef7.
//
// Solidity: function randomWalkNft() view returns(address)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) RandomWalkNft() (common.Address, error) {
	return _StakingWalletRandomWalkNft.Contract.RandomWalkNft(&_StakingWalletRandomWalkNft.CallOpts)
}

// RandomWalkNft is a free data retrieval call binding the contract method 0x755b4ef7.
//
// Solidity: function randomWalkNft() view returns(address)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCallerSession) RandomWalkNft() (common.Address, error) {
	return _StakingWalletRandomWalkNft.Contract.RandomWalkNft(&_StakingWalletRandomWalkNft.CallOpts)
}

// StakeActionIds is a free data retrieval call binding the contract method 0x60294405.
//
// Solidity: function stakeActionIds(uint256 ) view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCaller) StakeActionIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletRandomWalkNft.contract.Call(opts, &out, "stakeActionIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeActionIds is a free data retrieval call binding the contract method 0x60294405.
//
// Solidity: function stakeActionIds(uint256 ) view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) StakeActionIds(arg0 *big.Int) (*big.Int, error) {
	return _StakingWalletRandomWalkNft.Contract.StakeActionIds(&_StakingWalletRandomWalkNft.CallOpts, arg0)
}

// StakeActionIds is a free data retrieval call binding the contract method 0x60294405.
//
// Solidity: function stakeActionIds(uint256 ) view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCallerSession) StakeActionIds(arg0 *big.Int) (*big.Int, error) {
	return _StakingWalletRandomWalkNft.Contract.StakeActionIds(&_StakingWalletRandomWalkNft.CallOpts, arg0)
}

// StakeActions is a free data retrieval call binding the contract method 0xa2b136fb.
//
// Solidity: function stakeActions(uint256 ) view returns(uint256 index, uint256 nftId, address nftOwnerAddress)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCaller) StakeActions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Index           *big.Int
	NftId           *big.Int
	NftOwnerAddress common.Address
}, error) {
	var out []interface{}
	err := _StakingWalletRandomWalkNft.contract.Call(opts, &out, "stakeActions", arg0)

	outstruct := new(struct {
		Index           *big.Int
		NftId           *big.Int
		NftOwnerAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NftId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NftOwnerAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// StakeActions is a free data retrieval call binding the contract method 0xa2b136fb.
//
// Solidity: function stakeActions(uint256 ) view returns(uint256 index, uint256 nftId, address nftOwnerAddress)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) StakeActions(arg0 *big.Int) (struct {
	Index           *big.Int
	NftId           *big.Int
	NftOwnerAddress common.Address
}, error) {
	return _StakingWalletRandomWalkNft.Contract.StakeActions(&_StakingWalletRandomWalkNft.CallOpts, arg0)
}

// StakeActions is a free data retrieval call binding the contract method 0xa2b136fb.
//
// Solidity: function stakeActions(uint256 ) view returns(uint256 index, uint256 nftId, address nftOwnerAddress)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCallerSession) StakeActions(arg0 *big.Int) (struct {
	Index           *big.Int
	NftId           *big.Int
	NftOwnerAddress common.Address
}, error) {
	return _StakingWalletRandomWalkNft.Contract.StakeActions(&_StakingWalletRandomWalkNft.CallOpts, arg0)
}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCaller) WasNftUsed(opts *bind.CallOpts, nftId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletRandomWalkNft.contract.Call(opts, &out, "wasNftUsed", nftId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) WasNftUsed(nftId_ *big.Int) (*big.Int, error) {
	return _StakingWalletRandomWalkNft.Contract.WasNftUsed(&_StakingWalletRandomWalkNft.CallOpts, nftId_)
}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCallerSession) WasNftUsed(nftId_ *big.Int) (*big.Int, error) {
	return _StakingWalletRandomWalkNft.Contract.WasNftUsed(&_StakingWalletRandomWalkNft.CallOpts, nftId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactor) Stake(opts *bind.TransactOpts, nftId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.contract.Transact(opts, "stake", nftId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) Stake(nftId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.Stake(&_StakingWalletRandomWalkNft.TransactOpts, nftId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactorSession) Stake(nftId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.Stake(&_StakingWalletRandomWalkNft.TransactOpts, nftId_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactor) StakeMany(opts *bind.TransactOpts, nftIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.contract.Transact(opts, "stakeMany", nftIds_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) StakeMany(nftIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.StakeMany(&_StakingWalletRandomWalkNft.TransactOpts, nftIds_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactorSession) StakeMany(nftIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.StakeMany(&_StakingWalletRandomWalkNft.TransactOpts, nftIds_)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactor) Unstake(opts *bind.TransactOpts, stakeActionId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.contract.Transact(opts, "unstake", stakeActionId_)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) Unstake(stakeActionId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.Unstake(&_StakingWalletRandomWalkNft.TransactOpts, stakeActionId_)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactorSession) Unstake(stakeActionId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.Unstake(&_StakingWalletRandomWalkNft.TransactOpts, stakeActionId_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactor) UnstakeMany(opts *bind.TransactOpts, stakeActionIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.contract.Transact(opts, "unstakeMany", stakeActionIds_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) UnstakeMany(stakeActionIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.UnstakeMany(&_StakingWalletRandomWalkNft.TransactOpts, stakeActionIds_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactorSession) UnstakeMany(stakeActionIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.UnstakeMany(&_StakingWalletRandomWalkNft.TransactOpts, stakeActionIds_)
}

// StakingWalletRandomWalkNftNftStakedIterator is returned from FilterNftStaked and is used to iterate over the raw logs and unpacked data for NftStaked events raised by the StakingWalletRandomWalkNft contract.
type StakingWalletRandomWalkNftNftStakedIterator struct {
	Event *StakingWalletRandomWalkNftNftStaked // Event containing the contract specifics and raw log

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
func (it *StakingWalletRandomWalkNftNftStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletRandomWalkNftNftStaked)
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
		it.Event = new(StakingWalletRandomWalkNftNftStaked)
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
func (it *StakingWalletRandomWalkNftNftStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletRandomWalkNftNftStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletRandomWalkNftNftStaked represents a NftStaked event raised by the StakingWalletRandomWalkNft contract.
type StakingWalletRandomWalkNftNftStaked struct {
	StakeActionId *big.Int
	NftTypeCode   uint8
	NftId         *big.Int
	StakerAddress common.Address
	NumStakedNfts *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNftStaked is a free log retrieval operation binding the contract event 0xcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f00829.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint8 nftTypeCode, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftFilterer) FilterNftStaked(opts *bind.FilterOpts, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (*StakingWalletRandomWalkNftNftStakedIterator, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}

	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _StakingWalletRandomWalkNft.contract.FilterLogs(opts, "NftStaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRandomWalkNftNftStakedIterator{contract: _StakingWalletRandomWalkNft.contract, event: "NftStaked", logs: logs, sub: sub}, nil
}

// WatchNftStaked is a free log subscription operation binding the contract event 0xcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f00829.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint8 nftTypeCode, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftFilterer) WatchNftStaked(opts *bind.WatchOpts, sink chan<- *StakingWalletRandomWalkNftNftStaked, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (event.Subscription, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}

	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _StakingWalletRandomWalkNft.contract.WatchLogs(opts, "NftStaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletRandomWalkNftNftStaked)
				if err := _StakingWalletRandomWalkNft.contract.UnpackLog(event, "NftStaked", log); err != nil {
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

// ParseNftStaked is a log parse operation binding the contract event 0xcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f00829.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint8 nftTypeCode, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftFilterer) ParseNftStaked(log types.Log) (*StakingWalletRandomWalkNftNftStaked, error) {
	event := new(StakingWalletRandomWalkNftNftStaked)
	if err := _StakingWalletRandomWalkNft.contract.UnpackLog(event, "NftStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletRandomWalkNftNftUnstakedIterator is returned from FilterNftUnstaked and is used to iterate over the raw logs and unpacked data for NftUnstaked events raised by the StakingWalletRandomWalkNft contract.
type StakingWalletRandomWalkNftNftUnstakedIterator struct {
	Event *StakingWalletRandomWalkNftNftUnstaked // Event containing the contract specifics and raw log

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
func (it *StakingWalletRandomWalkNftNftUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletRandomWalkNftNftUnstaked)
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
		it.Event = new(StakingWalletRandomWalkNftNftUnstaked)
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
func (it *StakingWalletRandomWalkNftNftUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletRandomWalkNftNftUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletRandomWalkNftNftUnstaked represents a NftUnstaked event raised by the StakingWalletRandomWalkNft contract.
type StakingWalletRandomWalkNftNftUnstaked struct {
	StakeActionId *big.Int
	NftId         *big.Int
	StakerAddress common.Address
	NumStakedNfts *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNftUnstaked is a free log retrieval operation binding the contract event 0x1792a7a9d5e2f53a8d06f2ae40b4446d91b503e84dd7f6307f40cdeb11541668.
//
// Solidity: event NftUnstaked(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftFilterer) FilterNftUnstaked(opts *bind.FilterOpts, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (*StakingWalletRandomWalkNftNftUnstakedIterator, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _StakingWalletRandomWalkNft.contract.FilterLogs(opts, "NftUnstaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRandomWalkNftNftUnstakedIterator{contract: _StakingWalletRandomWalkNft.contract, event: "NftUnstaked", logs: logs, sub: sub}, nil
}

// WatchNftUnstaked is a free log subscription operation binding the contract event 0x1792a7a9d5e2f53a8d06f2ae40b4446d91b503e84dd7f6307f40cdeb11541668.
//
// Solidity: event NftUnstaked(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftFilterer) WatchNftUnstaked(opts *bind.WatchOpts, sink chan<- *StakingWalletRandomWalkNftNftUnstaked, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (event.Subscription, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _StakingWalletRandomWalkNft.contract.WatchLogs(opts, "NftUnstaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletRandomWalkNftNftUnstaked)
				if err := _StakingWalletRandomWalkNft.contract.UnpackLog(event, "NftUnstaked", log); err != nil {
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

// ParseNftUnstaked is a log parse operation binding the contract event 0x1792a7a9d5e2f53a8d06f2ae40b4446d91b503e84dd7f6307f40cdeb11541668.
//
// Solidity: event NftUnstaked(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftFilterer) ParseNftUnstaked(log types.Log) (*StakingWalletRandomWalkNftNftUnstaked, error) {
	event := new(StakingWalletRandomWalkNftNftUnstaked)
	if err := _StakingWalletRandomWalkNft.contract.UnpackLog(event, "NftUnstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
