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

// RaffleWalletMetaData contains all meta data concerning the RaffleWallet contract.
var RaffleWalletMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RaffleDepositEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RaffleWithdrawalEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"winners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061002c61002161003160201b60201c565b61003860201b60201c565b6100f9565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b610de6806101065f395ff3fe60806040526004361061006f575f3560e01c80638da5cb5b1161004d5780638da5cb5b146100cd578063a2fb1175146100f7578063bbd4e8c914610137578063f2fde38b146101615761006f565b80632e1a7d4d1461007357806347e7ef241461009b578063715018a6146100b7575b5f80fd5b34801561007e575f80fd5b5061009960048036038101906100949190610899565b610189565b005b6100b560048036038101906100b0919061091e565b61039b565b005b3480156100c2575f80fd5b506100cb610593565b005b3480156100d8575f80fd5b506100e161061a565b6040516100ee919061096b565b60405180910390f35b348015610102575f80fd5b5061011d60048036038101906101189190610899565b610641565b60405161012e9594939291906109ad565b60405180910390f35b348015610142575f80fd5b5061014b61069e565b60405161015891906109fe565b60405180910390f35b34801561016c575f80fd5b5061018760048036038101906101829190610a17565b6106a4565b005b60015f8281526020019081526020015f206004015f9054906101000a900460ff16156101ea576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101e190610a9c565b60405180910390fd5b6001805f8381526020019081526020015f206004015f6101000a81548160ff0219169083151502179055505f60015f8381526020019081526020015f205f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660015f8481526020019081526020015f206001015460405161028090610ae7565b5f6040518083038185875af1925050503d805f81146102ba576040519150601f19603f3d011682016040523d82523d5f602084013e6102bf565b606091505b5050905080610303576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102fa90610b45565b60405180910390fd5b60015f8381526020019081526020015f205f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f49775450da95afb818c4651b894f124d05629f06572a18900bf29db74b04a0d860015f8581526020019081526020015f206001015460405161038f91906109fe565b60405180910390a25050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610409576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040090610bad565b60405180910390fd5b5f341161044b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161044290610c15565b60405180910390fd5b6040518060a001604052808373ffffffffffffffffffffffffffffffffffffffff16815260200134815260200160025481526020018281526020015f151581525060015f60025481526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020155606082015181600301556080820151816004015f6101000a81548160ff021916908315150217905550905050808273ffffffffffffffffffffffffffffffffffffffff167fb1167d0680cf42eff86b3ab041def9a2bc93b1b86dc35ce7d6b8c3060f06ac906002543460405161056e929190610c33565b60405180910390a3600160025f8282546105889190610c87565b925050819055505050565b61059b61079a565b73ffffffffffffffffffffffffffffffffffffffff166105b961061a565b73ffffffffffffffffffffffffffffffffffffffff161461060f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060690610d04565b60405180910390fd5b6106185f6107a1565b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6001602052805f5260405f205f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001015490806002015490806003015490806004015f9054906101000a900460ff16905085565b60025481565b6106ac61079a565b73ffffffffffffffffffffffffffffffffffffffff166106ca61061a565b73ffffffffffffffffffffffffffffffffffffffff1614610720576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161071790610d04565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361078e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161078590610d92565b60405180910390fd5b610797816107a1565b50565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f80fd5b5f819050919050565b61087881610866565b8114610882575f80fd5b50565b5f813590506108938161086f565b92915050565b5f602082840312156108ae576108ad610862565b5b5f6108bb84828501610885565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6108ed826108c4565b9050919050565b6108fd816108e3565b8114610907575f80fd5b50565b5f81359050610918816108f4565b92915050565b5f806040838503121561093457610933610862565b5b5f6109418582860161090a565b925050602061095285828601610885565b9150509250929050565b610965816108e3565b82525050565b5f60208201905061097e5f83018461095c565b92915050565b61098d81610866565b82525050565b5f8115159050919050565b6109a781610993565b82525050565b5f60a0820190506109c05f83018861095c565b6109cd6020830187610984565b6109da6040830186610984565b6109e76060830185610984565b6109f4608083018461099e565b9695505050505050565b5f602082019050610a115f830184610984565b92915050565b5f60208284031215610a2c57610a2b610862565b5b5f610a398482850161090a565b91505092915050565b5f82825260208201905092915050565b7f526166666c652068617320616c72656479206265656e20636c61696d65642e005f82015250565b5f610a86601f83610a42565b9150610a9182610a52565b602082019050919050565b5f6020820190508181035f830152610ab381610a7a565b9050919050565b5f81905092915050565b50565b5f610ad25f83610aba565b9150610add82610ac4565b5f82019050919050565b5f610af182610ac7565b9150819050919050565b7f5472616e73666572206661696c65642e000000000000000000000000000000005f82015250565b5f610b2f601083610a42565b9150610b3a82610afb565b602082019050919050565b5f6020820190508181035f830152610b5c81610b23565b9050919050565b7f5a65726f2d616464726573732077617320676976656e2e0000000000000000005f82015250565b5f610b97601783610a42565b9150610ba282610b63565b602082019050919050565b5f6020820190508181035f830152610bc481610b8b565b9050919050565b7f4e6f2045544820686173206265656e2073656e742e00000000000000000000005f82015250565b5f610bff601583610a42565b9150610c0a82610bcb565b602082019050919050565b5f6020820190508181035f830152610c2c81610bf3565b9050919050565b5f604082019050610c465f830185610984565b610c536020830184610984565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610c9182610866565b9150610c9c83610866565b9250828201905080821115610cb457610cb3610c5a565b5b92915050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f610cee602083610a42565b9150610cf982610cba565b602082019050919050565b5f6020820190508181035f830152610d1b81610ce2565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f610d7c602683610a42565b9150610d8782610d22565b604082019050919050565b5f6020820190508181035f830152610da981610d70565b905091905056fea26469706673582212209333483d981e4bc2f971dfb52d163db4c72ce4e2633228c050b64580bd83d9c464736f6c63430008150033",
}

// RaffleWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use RaffleWalletMetaData.ABI instead.
var RaffleWalletABI = RaffleWalletMetaData.ABI

// RaffleWalletBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RaffleWalletMetaData.Bin instead.
var RaffleWalletBin = RaffleWalletMetaData.Bin

// DeployRaffleWallet deploys a new Ethereum contract, binding an instance of RaffleWallet to it.
func DeployRaffleWallet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RaffleWallet, error) {
	parsed, err := RaffleWalletMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RaffleWalletBin), backend)
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
	parsed, err := abi.JSON(strings.NewReader(RaffleWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// NumDeposits is a free data retrieval call binding the contract method 0xbbd4e8c9.
//
// Solidity: function numDeposits() view returns(uint256)
func (_RaffleWallet *RaffleWalletCaller) NumDeposits(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RaffleWallet.contract.Call(opts, &out, "numDeposits")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumDeposits is a free data retrieval call binding the contract method 0xbbd4e8c9.
//
// Solidity: function numDeposits() view returns(uint256)
func (_RaffleWallet *RaffleWalletSession) NumDeposits() (*big.Int, error) {
	return _RaffleWallet.Contract.NumDeposits(&_RaffleWallet.CallOpts)
}

// NumDeposits is a free data retrieval call binding the contract method 0xbbd4e8c9.
//
// Solidity: function numDeposits() view returns(uint256)
func (_RaffleWallet *RaffleWalletCallerSession) NumDeposits() (*big.Int, error) {
	return _RaffleWallet.Contract.NumDeposits(&_RaffleWallet.CallOpts)
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

// Winners is a free data retrieval call binding the contract method 0xa2fb1175.
//
// Solidity: function winners(uint256 ) view returns(address destination, uint256 amount, uint256 depositId, uint256 round, bool claimed)
func (_RaffleWallet *RaffleWalletCaller) Winners(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Destination common.Address
	Amount      *big.Int
	DepositId   *big.Int
	Round       *big.Int
	Claimed     bool
}, error) {
	var out []interface{}
	err := _RaffleWallet.contract.Call(opts, &out, "winners", arg0)

	outstruct := new(struct {
		Destination common.Address
		Amount      *big.Int
		DepositId   *big.Int
		Round       *big.Int
		Claimed     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Destination = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DepositId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Round = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Claimed = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Winners is a free data retrieval call binding the contract method 0xa2fb1175.
//
// Solidity: function winners(uint256 ) view returns(address destination, uint256 amount, uint256 depositId, uint256 round, bool claimed)
func (_RaffleWallet *RaffleWalletSession) Winners(arg0 *big.Int) (struct {
	Destination common.Address
	Amount      *big.Int
	DepositId   *big.Int
	Round       *big.Int
	Claimed     bool
}, error) {
	return _RaffleWallet.Contract.Winners(&_RaffleWallet.CallOpts, arg0)
}

// Winners is a free data retrieval call binding the contract method 0xa2fb1175.
//
// Solidity: function winners(uint256 ) view returns(address destination, uint256 amount, uint256 depositId, uint256 round, bool claimed)
func (_RaffleWallet *RaffleWalletCallerSession) Winners(arg0 *big.Int) (struct {
	Destination common.Address
	Amount      *big.Int
	DepositId   *big.Int
	Round       *big.Int
	Claimed     bool
}, error) {
	return _RaffleWallet.Contract.Winners(&_RaffleWallet.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address winner, uint256 roundNum) payable returns()
func (_RaffleWallet *RaffleWalletTransactor) Deposit(opts *bind.TransactOpts, winner common.Address, roundNum *big.Int) (*types.Transaction, error) {
	return _RaffleWallet.contract.Transact(opts, "deposit", winner, roundNum)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address winner, uint256 roundNum) payable returns()
func (_RaffleWallet *RaffleWalletSession) Deposit(winner common.Address, roundNum *big.Int) (*types.Transaction, error) {
	return _RaffleWallet.Contract.Deposit(&_RaffleWallet.TransactOpts, winner, roundNum)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address winner, uint256 roundNum) payable returns()
func (_RaffleWallet *RaffleWalletTransactorSession) Deposit(winner common.Address, roundNum *big.Int) (*types.Transaction, error) {
	return _RaffleWallet.Contract.Deposit(&_RaffleWallet.TransactOpts, winner, roundNum)
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

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 depositId) returns()
func (_RaffleWallet *RaffleWalletTransactor) Withdraw(opts *bind.TransactOpts, depositId *big.Int) (*types.Transaction, error) {
	return _RaffleWallet.contract.Transact(opts, "withdraw", depositId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 depositId) returns()
func (_RaffleWallet *RaffleWalletSession) Withdraw(depositId *big.Int) (*types.Transaction, error) {
	return _RaffleWallet.Contract.Withdraw(&_RaffleWallet.TransactOpts, depositId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 depositId) returns()
func (_RaffleWallet *RaffleWalletTransactorSession) Withdraw(depositId *big.Int) (*types.Transaction, error) {
	return _RaffleWallet.Contract.Withdraw(&_RaffleWallet.TransactOpts, depositId)
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
	Winner    common.Address
	Round     *big.Int
	DepositId *big.Int
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRaffleDepositEvent is a free log retrieval operation binding the contract event 0xb1167d0680cf42eff86b3ab041def9a2bc93b1b86dc35ce7d6b8c3060f06ac90.
//
// Solidity: event RaffleDepositEvent(address indexed winner, uint256 indexed round, uint256 depositId, uint256 amount)
func (_RaffleWallet *RaffleWalletFilterer) FilterRaffleDepositEvent(opts *bind.FilterOpts, winner []common.Address, round []*big.Int) (*RaffleWalletRaffleDepositEventIterator, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _RaffleWallet.contract.FilterLogs(opts, "RaffleDepositEvent", winnerRule, roundRule)
	if err != nil {
		return nil, err
	}
	return &RaffleWalletRaffleDepositEventIterator{contract: _RaffleWallet.contract, event: "RaffleDepositEvent", logs: logs, sub: sub}, nil
}

// WatchRaffleDepositEvent is a free log subscription operation binding the contract event 0xb1167d0680cf42eff86b3ab041def9a2bc93b1b86dc35ce7d6b8c3060f06ac90.
//
// Solidity: event RaffleDepositEvent(address indexed winner, uint256 indexed round, uint256 depositId, uint256 amount)
func (_RaffleWallet *RaffleWalletFilterer) WatchRaffleDepositEvent(opts *bind.WatchOpts, sink chan<- *RaffleWalletRaffleDepositEvent, winner []common.Address, round []*big.Int) (event.Subscription, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _RaffleWallet.contract.WatchLogs(opts, "RaffleDepositEvent", winnerRule, roundRule)
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

// ParseRaffleDepositEvent is a log parse operation binding the contract event 0xb1167d0680cf42eff86b3ab041def9a2bc93b1b86dc35ce7d6b8c3060f06ac90.
//
// Solidity: event RaffleDepositEvent(address indexed winner, uint256 indexed round, uint256 depositId, uint256 amount)
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
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRaffleWithdrawalEvent is a free log retrieval operation binding the contract event 0x49775450da95afb818c4651b894f124d05629f06572a18900bf29db74b04a0d8.
//
// Solidity: event RaffleWithdrawalEvent(address indexed winner, uint256 amount)
func (_RaffleWallet *RaffleWalletFilterer) FilterRaffleWithdrawalEvent(opts *bind.FilterOpts, winner []common.Address) (*RaffleWalletRaffleWithdrawalEventIterator, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _RaffleWallet.contract.FilterLogs(opts, "RaffleWithdrawalEvent", winnerRule)
	if err != nil {
		return nil, err
	}
	return &RaffleWalletRaffleWithdrawalEventIterator{contract: _RaffleWallet.contract, event: "RaffleWithdrawalEvent", logs: logs, sub: sub}, nil
}

// WatchRaffleWithdrawalEvent is a free log subscription operation binding the contract event 0x49775450da95afb818c4651b894f124d05629f06572a18900bf29db74b04a0d8.
//
// Solidity: event RaffleWithdrawalEvent(address indexed winner, uint256 amount)
func (_RaffleWallet *RaffleWalletFilterer) WatchRaffleWithdrawalEvent(opts *bind.WatchOpts, sink chan<- *RaffleWalletRaffleWithdrawalEvent, winner []common.Address) (event.Subscription, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _RaffleWallet.contract.WatchLogs(opts, "RaffleWithdrawalEvent", winnerRule)
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
// Solidity: event RaffleWithdrawalEvent(address indexed winner, uint256 amount)
func (_RaffleWallet *RaffleWalletFilterer) ParseRaffleWithdrawalEvent(log types.Log) (*RaffleWalletRaffleWithdrawalEvent, error) {
	event := new(RaffleWalletRaffleWithdrawalEvent)
	if err := _RaffleWallet.contract.UnpackLog(event, "RaffleWithdrawalEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
