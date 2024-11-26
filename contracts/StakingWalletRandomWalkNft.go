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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumCosmicSignatureConstants.NftTypeCode\",\"name\":\"nftTypeCode\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftUnstaked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"numStakedNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"entropy_\",\"type\":\"bytes32\"}],\"name\":\"pickRandomStakerAddressIfPossible\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nftIds_\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId_\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"stakeActionIds_\",\"type\":\"uint256[]\"}],\"name\":\"unstakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"wasNftUsed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// PickRandomStakerAddressIfPossible is a free data retrieval call binding the contract method 0x0f914941.
//
// Solidity: function pickRandomStakerAddressIfPossible(bytes32 entropy_) view returns(address)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftCaller) PickRandomStakerAddressIfPossible(opts *bind.CallOpts, entropy_ [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IStakingWalletRandomWalkNft.contract.Call(opts, &out, "pickRandomStakerAddressIfPossible", entropy_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PickRandomStakerAddressIfPossible is a free data retrieval call binding the contract method 0x0f914941.
//
// Solidity: function pickRandomStakerAddressIfPossible(bytes32 entropy_) view returns(address)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftSession) PickRandomStakerAddressIfPossible(entropy_ [32]byte) (common.Address, error) {
	return _IStakingWalletRandomWalkNft.Contract.PickRandomStakerAddressIfPossible(&_IStakingWalletRandomWalkNft.CallOpts, entropy_)
}

// PickRandomStakerAddressIfPossible is a free data retrieval call binding the contract method 0x0f914941.
//
// Solidity: function pickRandomStakerAddressIfPossible(bytes32 entropy_) view returns(address)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftCallerSession) PickRandomStakerAddressIfPossible(entropy_ [32]byte) (common.Address, error) {
	return _IStakingWalletRandomWalkNft.Contract.PickRandomStakerAddressIfPossible(&_IStakingWalletRandomWalkNft.CallOpts, entropy_)
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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractRandomWalkNFT\",\"name\":\"randomWalkNft_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"name\":\"NftOneTimeStaking\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"requesterAddress\",\"type\":\"address\"}],\"name\":\"NftStakeActionAccessDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"}],\"name\":\"NftStakeActionInvalidId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumCosmicSignatureConstants.NftTypeCode\",\"name\":\"nftTypeCode\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftUnstaked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"actionCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakedNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"entropy_\",\"type\":\"bytes32\"}],\"name\":\"pickRandomStakerAddressIfPossible\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomWalkNft\",\"outputs\":[{\"internalType\":\"contractRandomWalkNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeActionIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeActions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftOwnerAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nftIds_\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId_\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"stakeActionIds_\",\"type\":\"uint256[]\"}],\"name\":\"unstakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"wasNftUsed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040523461004657610019610014610117565b61024c565b61002161004b565b6113ef61027f823960805181818161041201528181610e22015261129501526113ef90f35b610051565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b9061007d90610055565b810190811060018060401b0382111761009557604052565b61005f565b906100ad6100a661004b565b9283610073565b565b5f80fd5b60018060a01b031690565b6100c7906100b3565b90565b6100d3906100be565b90565b6100df816100ca565b036100e657565b5f80fd5b905051906100f7826100d6565b565b906020828203126101125761010f915f016100ea565b90565b6100af565b61013561166e8038038061012a8161009a565b9283398101906100f9565b90565b90565b61014f61014a610154926100b3565b610138565b6100b3565b90565b6101609061013b565b90565b61016c90610157565b90565b90565b61018661018161018b9261016f565b610138565b6100b3565b90565b61019790610172565b90565b60209181520190565b60207f616e646f6d57616c6b4e66745f2e000000000000000000000000000000000000917f5a65726f2d616464726573732077617320676976656e20666f722074686520725f8201520152565b6101fd602e60409261019a565b610206816101a3565b0190565b61021f9060208101905f8183039101526101f0565b90565b1561022957565b61023161004b565b63eac0d38960e01b8152806102486004820161020a565b0390fd5b61027961025882610163565b61027261026c6102675f61018e565b6100be565b916100be565b1415610222565b60805256fe60806040526004361015610013575b6106ef565b61001d5f356100cc565b80630d50c189146100c75780630f914941146100c25780632e17de78146100bd57806357951c74146100b857806360294405146100b3578063755b4ef7146100ae578063a2b136fb146100a9578063a694fc3a146100a4578063ca7c1f921461009f578063fdbd98b01461009a5763fe939afc0361000e576106bb565b610686565b61063a565b610607565b6105cf565b61048d565b6103cc565b6102e9565b610294565b61021b565b61016a565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f80fd5b909182601f8301121561012a5781359167ffffffffffffffff831161012557602001926020830284011161012057565b6100ec565b6100e8565b6100e4565b90602082820312610160575f82013567ffffffffffffffff811161015b5761015792016100f0565b9091565b6100e0565b6100dc565b5f0190565b346101995761018361017d36600461012f565b90610747565b61018b6100d2565b8061019581610165565b0390f35b6100d8565b90565b6101aa8161019e565b036101b157565b5f80fd5b905035906101c2826101a1565b565b906020828203126101dd576101da915f016101b5565b90565b6100dc565b60018060a01b031690565b6101f6906101e2565b90565b610202906101ed565b9052565b9190610219905f602085019401906101f9565b565b3461024b576102476102366102313660046101c4565b610839565b61023e6100d2565b91829182610206565b0390f35b6100d8565b90565b61025c81610250565b0361026357565b5f80fd5b9050359061027482610253565b565b9060208282031261028f5761028c915f01610267565b90565b6100dc565b346102c2576102ac6102a7366004610276565b610c97565b6102b46100d2565b806102be81610165565b0390f35b6100d8565b6102d090610250565b9052565b91906102e7905f602085019401906102c7565b565b34610319576103156103046102ff366004610276565b610f57565b61030c6100d2565b918291826102d4565b0390f35b6100d8565b634e487b7160e01b5f52603260045260245ffd5b506801000000000000000090565b90565b61034c81610332565b8210156103665761035e600191610340565b910201905f90565b61031e565b1c90565b90565b610382906008610387930261036b565b61036f565b90565b906103959154610372565b90565b680400000000000000026103ab81610332565b8210156103c8576103c5916103bf91610343565b9061038a565b90565b5f80fd5b346103fc576103f86103e76103e2366004610276565b610398565b6103ef6100d2565b918291826102d4565b0390f35b6100d8565b5f91031261040b57565b6100dc565b7f000000000000000000000000000000000000000000000000000000000000000090565b90565b61044b610446610450926101e2565b610434565b6101e2565b90565b61045c90610437565b90565b61046890610453565b90565b6104749061045f565b9052565b919061048b905f6020850194019061046b565b565b346104bd5761049d366004610401565b6104b96104a8610410565b6104b06100d2565b91829182610478565b0390f35b6100d8565b506801000000000000000090565b90565b6104dc816104c2565b8210156104f6576104ee6003916104d0565b910201905f90565b61031e565b5f1c90565b61050c610511916104fb565b61036f565b90565b61051e9054610500565b90565b60018060a01b031690565b61053861053d916104fb565b610521565b90565b61054a905461052c565b90565b6801000000000000000290610561826104c2565b81101561059957610571916104d3565b5061057d5f8201610514565b91610596600261058f60018501610514565b9301610540565b90565b5f80fd5b6040906105c66105cd94969593966105bc60608401985f8501906102c7565b60208301906102c7565b01906101f9565b565b34610602576105fe6105ea6105e5366004610276565b61054d565b6105f59391936100d2565b9384938461059d565b0390f35b6100d8565b346106355761061f61061a366004610276565b611122565b6106276100d2565b8061063181610165565b0390f35b6100d8565b3461066a5761064a366004610401565b610666610655611344565b61065d6100d2565b918291826102d4565b0390f35b6100d8565b610683680100000000000000015f9061038a565b90565b346106b657610696366004610401565b6106b26106a161066f565b6106a96100d2565b918291826102d4565b0390f35b6100d8565b346106ea576106d46106ce36600461012f565b90611359565b6106dc6100d2565b806106e681610165565b0390f35b6100d8565b5f80fd5b90565b61070a61070561070f926106f3565b610434565b610250565b90565b600161071e9101610250565b90565b5090565b9190811015610735576020020190565b61031e565b3561074481610253565b90565b9190916107535f6106f6565b5b8061077161076b610766858890610721565b610250565b91610250565b10156107a15761079c9061079761079261078d85888591610725565b61073a565b610c97565b610712565b610754565b50509050565b5f90565b6107bf6107ba6107c4926106f3565b610434565b6101e2565b90565b6107d0906107ab565b90565b6107e76107e26107ec92610250565b610434565b610250565b90565b6107fb610800916104fb565b6107d3565b90565b634e487b7160e01b5f52601260045260245ffd5b61082361082991610250565b91610250565b908115610834570690565b610803565b6108416107a7565b5061084b5f610514565b908161085f6108595f6106f6565b91610250565b146108b2576108a86108996108936108846108af9561087f6002966107ef565b610817565b68040000000000000002610343565b9061038a565b680100000000000000026104d3565b5001610540565b90565b50506108bd5f6107c7565b90565b90565b906108cd90610250565b9052565b906108db906101ed565b9052565b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b90610907906108df565b810190811067ffffffffffffffff82111761092157604052565b6108e9565b906109396109326100d2565b92836108fd565b565b6109456060610926565b90565b9061099761098e600261095961093b565b946109706109685f8301610514565b5f88016108c3565b61098861097f60018301610514565b602088016108c3565b01610540565b604084016108d1565b565b6109a290610948565b90565b6109af90516101ed565b90565b60209181520190565b5f7f496e76616c6964204e4654207374616b6520616374696f6e2049442e00000000910152565b6109ef601c6020926109b2565b6109f8816109bb565b0190565b9190610a1f906020610a17604086018681035f8801526109e2565b9401906102c7565b565b60207f6e7374616b652069742e00000000000000000000000000000000000000000000917f4f6e6c79204e4654206f776e6572206973207065726d697474656420746f20755f8201520152565b610a7b602a6040926109b2565b610a8481610a21565b0190565b916040610ab9929493610ab2610aa7606083018381035f850152610a6e565b9660208301906102c7565b01906101f9565b565b90565b610ad2610acd610ad792610abb565b610434565b610250565b90565b634e487b7160e01b5f52601160045260245ffd5b610afd610b0391939293610250565b92610250565b8203918211610b0e57565b610ada565b5f1b90565b90610b245f1991610b13565b9181191691161790565b90565b90610b46610b41610b4d926107d3565b610b2e565b8254610b18565b9055565b610b5b9051610250565b90565b1b90565b91906008610b7d910291610b775f1984610b5e565b92610b5e565b9181191691161790565b9190610b9d610b98610ba5936107d3565b610b2e565b908354610b62565b9055565b5f90565b610bbf91610bb9610ba9565b91610b87565b565b91906008610be1910291610bdb60018060a01b0384610b5e565b92610b5e565b9181191691161790565b610bf490610453565b90565b90565b9190610c10610c0b610c1893610beb565b610bf7565b908354610bc1565b9055565b610c2e91610c286107a7565b91610bfa565b565b610c3990610453565b90565b5f80fd5b60e01b90565b5f910312610c5057565b6100dc565b604090610c7e610c859496959396610c7460608401985f8501906101f9565b60208301906101f9565b01906102c7565b565b610c8f6100d2565b3d5f823e3d90fd5b610cb4610cae6801000000000000000283906104d3565b506108c0565b610cbd81610999565b9133610cdc610cd6610cd1604087016109a5565b6101ed565b916101ed565b03610edc57610d9c610d00610cf05f610514565b610cfa6001610abe565b90610aee565b92610d0b845f610b31565b610d785f6002610d2e610d28680400000000000000028990610343565b9061038a565b93610d59610d3d848b01610b51565b84610d526801000000000000000289906104d3565b5001610b31565b610d6583808301610bad565b610d728360018301610bad565b01610c1c565b610d9668040000000000000002610d905f8801610b51565b90610343565b90610b87565b610db9610db3680400000000000000028490610343565b90610bad565b90610dc660208401610b51565b339192610e1a610e08610e02610dfc7f1792a7a9d5e2f53a8d06f2ae40b4446d91b503e84dd7f6307f40cdeb11541668946107d3565b946107d3565b94610beb565b94610e116100d2565b918291826102d4565b0390a4610e467f000000000000000000000000000000000000000000000000000000000000000061045f565b6323b872dd90610e5530610c30565b90610e636020339501610b51565b92813b15610ed7575f610e8991610e948296610e7d6100d2565b98899788968795610c40565b855260048501610c55565b03925af18015610ed257610ea6575b50565b610ec5905f3d8111610ecb575b610ebd81836108fd565b810190610c46565b5f610ea3565b503d610eb3565b610c87565b610c3c565b610ee8604084016109a5565b610f02610efc610ef75f6107c7565b6101ed565b916101ed565b14155f14610f31573390610f2d610f176100d2565b9283926348aca7ef60e11b845260048401610a88565b0390fd5b610f5390610f3d6100d2565b91829163023df6b160e21b8352600483016109fc565b0390fd5b610f6e610f7491610f66610ba9565b506001610343565b9061038a565b90565b60407f206f6e6c79206f6e63652e000000000000000000000000000000000000000000917f54686973204e46542068617320616c7265616479206265656e207374616b65645f8201527f2e20416e204e465420697320616c6c6f77656420746f206265207374616b656460208201520152565b610ff7604b6060926109b2565b61100081610f77565b0190565b919061102790602061101f604086018681035f880152610fea565b9401906102c7565b565b156110315750565b6110539061103d6100d2565b918291633b471d1f60e21b835260048301611004565b0390fd5b61106661106c91939293610250565b92610250565b820180921161107757565b610ada565b9061108d60018060a01b0391610b13565b9181191691161790565b906110ac6110a76110b392610beb565b610bf7565b825461107c565b9055565b634e487b7160e01b5f52602160045260245ffd5b600311156110d557565b6110b7565b906110e4826110cb565b565b6110ef906110da565b90565b6110fb906110e6565b9052565b91602061112092949361111960408201965f8301906110f2565b01906102c7565b565b61115461113a61113460018490610343565b9061038a565b61114c6111465f6106f6565b91610250565b148290611029565b6111736111616001610abe565b61116d60018490610343565b90610b87565b61119861118868010000000000000001610514565b6111926001610abe565b90611057565b6111ab8168010000000000000001610b31565b6112296111cb6111c56801000000000000000284906104d3565b506108c0565b6111d88460018301610b31565b6111e53360028301611097565b6111fb6111f15f610514565b915f839101610b31565b61121983611213680400000000000000028490610343565b90610b87565b6112236001610abe565b90611057565b90611234825f610b31565b906002918390339261127861127261126c7fcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f00829946107d3565b946107d3565b94610beb565b9461128d6112846100d2565b928392836110ff565b0390a46112b97f000000000000000000000000000000000000000000000000000000000000000061045f565b6323b872dd9033906112ca30610c30565b9392813b1561133f575f6112f1916112fc82966112e56100d2565b98899788968795610c40565b855260048501610c55565b03925af1801561133a5761130e575b50565b61132d905f3d8111611333575b61132581836108fd565b810190610c46565b5f61130b565b503d61131b565b610c87565b610c3c565b61134c610ba9565b506113565f610514565b90565b9190916113655f6106f6565b5b8061138361137d611378858890610721565b610250565b91610250565b10156113b3576113ae906113a96113a461139f85888591610725565b61073a565b611122565b610712565b611366565b5050905056fea2646970667358221220541cb1715259b5f3329567be9038b4ac5ce9dcc1bb631c21df3cdaf8c3f3d2af64736f6c634300081b0033",
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

// PickRandomStakerAddressIfPossible is a free data retrieval call binding the contract method 0x0f914941.
//
// Solidity: function pickRandomStakerAddressIfPossible(bytes32 entropy_) view returns(address)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCaller) PickRandomStakerAddressIfPossible(opts *bind.CallOpts, entropy_ [32]byte) (common.Address, error) {
	var out []interface{}
	err := _StakingWalletRandomWalkNft.contract.Call(opts, &out, "pickRandomStakerAddressIfPossible", entropy_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PickRandomStakerAddressIfPossible is a free data retrieval call binding the contract method 0x0f914941.
//
// Solidity: function pickRandomStakerAddressIfPossible(bytes32 entropy_) view returns(address)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) PickRandomStakerAddressIfPossible(entropy_ [32]byte) (common.Address, error) {
	return _StakingWalletRandomWalkNft.Contract.PickRandomStakerAddressIfPossible(&_StakingWalletRandomWalkNft.CallOpts, entropy_)
}

// PickRandomStakerAddressIfPossible is a free data retrieval call binding the contract method 0x0f914941.
//
// Solidity: function pickRandomStakerAddressIfPossible(bytes32 entropy_) view returns(address)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCallerSession) PickRandomStakerAddressIfPossible(entropy_ [32]byte) (common.Address, error) {
	return _StakingWalletRandomWalkNft.Contract.PickRandomStakerAddressIfPossible(&_StakingWalletRandomWalkNft.CallOpts, entropy_)
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
