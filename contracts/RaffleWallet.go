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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RaffleDepositEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"round_num\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"winners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposit_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061002d61002261003260201b60201c565b61003a60201b60201c565b6100fe565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b610dba8061010d6000396000f3fe6080604052600436106100705760003560e01c80638da5cb5b1161004e5780638da5cb5b146100d1578063a2fb1175146100fc578063bbd4e8c91461013d578063f2fde38b1461016857610070565b80632e1a7d4d1461007557806347e7ef241461009e578063715018a6146100ba575b600080fd5b34801561008157600080fd5b5061009c60048036038101906100979190610839565b610191565b005b6100b860048036038101906100b391906108c4565b61031f565b005b3480156100c657600080fd5b506100cf610521565b005b3480156100dd57600080fd5b506100e66105a9565b6040516100f39190610913565b60405180910390f35b34801561010857600080fd5b50610123600480360381019061011e9190610839565b6105d2565b604051610134959493929190610958565b60405180910390f35b34801561014957600080fd5b50610152610635565b60405161015f91906109ab565b60405180910390f35b34801561017457600080fd5b5061018f600480360381019061018a91906109c6565b61063b565b005b6001600082815260200190815260200160002060040160009054906101000a900460ff16156101f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101ec90610a50565b60405180910390fd5b600180600083815260200190815260200160002060040160006101000a81548160ff02191690831515021790555060006001600083815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600160008481526020019081526020016000206001015460405161029590610aa1565b60006040518083038185875af1925050503d80600081146102d2576040519150601f19603f3d011682016040523d82523d6000602084013e6102d7565b606091505b505090508061031b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161031290610b02565b60405180910390fd5b5050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361038e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161038590610b6e565b60405180910390fd5b600034116103d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103c890610bda565b60405180910390fd5b6040518060a001604052808373ffffffffffffffffffffffffffffffffffffffff16815260200134815260200160025481526020018281526020016000151581525060016000600254815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155604082015181600201556060820151816003015560808201518160040160006101000a81548160ff021916908315150217905550905050808273ffffffffffffffffffffffffffffffffffffffff167fb1167d0680cf42eff86b3ab041def9a2bc93b1b86dc35ce7d6b8c3060f06ac90600254346040516104fb929190610bfa565b60405180910390a36001600260008282546105169190610c52565b925050819055505050565b610529610732565b73ffffffffffffffffffffffffffffffffffffffff166105476105a9565b73ffffffffffffffffffffffffffffffffffffffff161461059d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161059490610cd2565b60405180910390fd5b6105a7600061073a565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60016020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154908060030154908060040160009054906101000a900460ff16905085565b60025481565b610643610732565b73ffffffffffffffffffffffffffffffffffffffff166106616105a9565b73ffffffffffffffffffffffffffffffffffffffff16146106b7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106ae90610cd2565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610726576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161071d90610d64565b60405180910390fd5b61072f8161073a565b50565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b6000819050919050565b61081681610803565b811461082157600080fd5b50565b6000813590506108338161080d565b92915050565b60006020828403121561084f5761084e6107fe565b5b600061085d84828501610824565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061089182610866565b9050919050565b6108a181610886565b81146108ac57600080fd5b50565b6000813590506108be81610898565b92915050565b600080604083850312156108db576108da6107fe565b5b60006108e9858286016108af565b92505060206108fa85828601610824565b9150509250929050565b61090d81610886565b82525050565b60006020820190506109286000830184610904565b92915050565b61093781610803565b82525050565b60008115159050919050565b6109528161093d565b82525050565b600060a08201905061096d6000830188610904565b61097a602083018761092e565b610987604083018661092e565b610994606083018561092e565b6109a16080830184610949565b9695505050505050565b60006020820190506109c0600083018461092e565b92915050565b6000602082840312156109dc576109db6107fe565b5b60006109ea848285016108af565b91505092915050565b600082825260208201905092915050565b7f526166666c652068617320616c72656479206265656e20636c61696d65642e00600082015250565b6000610a3a601f836109f3565b9150610a4582610a04565b602082019050919050565b60006020820190508181036000830152610a6981610a2d565b9050919050565b600081905092915050565b50565b6000610a8b600083610a70565b9150610a9682610a7b565b600082019050919050565b6000610aac82610a7e565b9150819050919050565b7f5472616e73666572206661696c65642e00000000000000000000000000000000600082015250565b6000610aec6010836109f3565b9150610af782610ab6565b602082019050919050565b60006020820190508181036000830152610b1b81610adf565b9050919050565b7f5a65726f2d616464726573732077617320676976656e2e000000000000000000600082015250565b6000610b586017836109f3565b9150610b6382610b22565b602082019050919050565b60006020820190508181036000830152610b8781610b4b565b9050919050565b7f4e6f2045544820686173206265656e2073656e742e0000000000000000000000600082015250565b6000610bc46015836109f3565b9150610bcf82610b8e565b602082019050919050565b60006020820190508181036000830152610bf381610bb7565b9050919050565b6000604082019050610c0f600083018561092e565b610c1c602083018461092e565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610c5d82610803565b9150610c6883610803565b9250828201905080821115610c8057610c7f610c23565b5b92915050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000610cbc6020836109f3565b9150610cc782610c86565b602082019050919050565b60006020820190508181036000830152610ceb81610caf565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000610d4e6026836109f3565b9150610d5982610cf2565b604082019050919050565b60006020820190508181036000830152610d7d81610d41565b905091905056fea26469706673582212208f4acece0c9c17678f868f2cb30ebaacf043217f5e63ed1f3f7fbe4bd68521f264736f6c63430008130033",
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
// Solidity: function winners(uint256 ) view returns(address destination, uint256 amount, uint256 deposit_id, uint256 round, bool claimed)
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
// Solidity: function winners(uint256 ) view returns(address destination, uint256 amount, uint256 deposit_id, uint256 round, bool claimed)
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
// Solidity: function winners(uint256 ) view returns(address destination, uint256 amount, uint256 deposit_id, uint256 round, bool claimed)
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
// Solidity: function deposit(address winner, uint256 round_num) payable returns()
func (_RaffleWallet *RaffleWalletTransactor) Deposit(opts *bind.TransactOpts, winner common.Address, round_num *big.Int) (*types.Transaction, error) {
	return _RaffleWallet.contract.Transact(opts, "deposit", winner, round_num)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address winner, uint256 round_num) payable returns()
func (_RaffleWallet *RaffleWalletSession) Deposit(winner common.Address, round_num *big.Int) (*types.Transaction, error) {
	return _RaffleWallet.Contract.Deposit(&_RaffleWallet.TransactOpts, winner, round_num)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address winner, uint256 round_num) payable returns()
func (_RaffleWallet *RaffleWalletTransactorSession) Deposit(winner common.Address, round_num *big.Int) (*types.Transaction, error) {
	return _RaffleWallet.Contract.Deposit(&_RaffleWallet.TransactOpts, winner, round_num)
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
// Solidity: event RaffleDepositEvent(address indexed winner, uint256 indexed round, uint256 deposit_id, uint256 amount)
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
// Solidity: event RaffleDepositEvent(address indexed winner, uint256 indexed round, uint256 deposit_id, uint256 amount)
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
// Solidity: event RaffleDepositEvent(address indexed winner, uint256 indexed round, uint256 deposit_id, uint256 amount)
func (_RaffleWallet *RaffleWalletFilterer) ParseRaffleDepositEvent(log types.Log) (*RaffleWalletRaffleDepositEvent, error) {
	event := new(RaffleWalletRaffleDepositEvent)
	if err := _RaffleWallet.contract.UnpackLog(event, "RaffleDepositEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
