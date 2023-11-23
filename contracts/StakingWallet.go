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

// StakingWalletMetaData contains all meta data concerning the StakingWallet contract.
var StakingWalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractCosmicSignature\",\"name\":\"nft_\",\"type\":\"address\"},{\"internalType\":\"contractCosmicGame\",\"name\":\"game_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevRoundReminder\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountPerHolder\",\"type\":\"uint256\"}],\"name\":\"StakingDepositEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"amountInRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"game\",\"outputs\":[{\"internalType\":\"contractCosmicGame\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isPaid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nft\",\"outputs\":[{\"internalType\":\"contractCosmicSignature\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"numWinnersInRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"previousRoundReminder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260006002553480156200001657600080fd5b50604051620013403803806200134083398181016040528101906200003c919062000289565b6200005c62000050620000e660201b60201c565b620000ee60201b60201c565b81600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050620002d0565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620001e482620001b7565b9050919050565b6000620001f882620001d7565b9050919050565b6200020a81620001eb565b81146200021657600080fd5b50565b6000815190506200022a81620001ff565b92915050565b60006200023d82620001b7565b9050919050565b6000620002518262000230565b9050919050565b620002638162000244565b81146200026f57600080fd5b50565b600081519050620002838162000258565b92915050565b60008060408385031215620002a357620002a2620001b2565b5b6000620002b38582860162000219565b9250506020620002c68582860162000272565b9150509250929050565b61106080620002e06000396000f3fe60806040526004361061009c5760003560e01c80638da5cb5b116100645780638da5cb5b14610174578063b6b55f251461019f578063c0450764146101bb578063c3fe3e28146101f8578063d823e4a614610223578063f2fde38b146102605761009c565b80631eb2874c146100a1578063441a3e70146100cc57806347ccca02146100f55780635658072f14610120578063715018a61461015d575b600080fd5b3480156100ad57600080fd5b506100b6610289565b6040516100c391906109bc565b60405180910390f35b3480156100d857600080fd5b506100f360048036038101906100ee9190610a08565b61028f565b005b34801561010157600080fd5b5061010a610486565b6040516101179190610ac7565b60405180910390f35b34801561012c57600080fd5b5061014760048036038101906101429190610ae2565b6104ac565b60405161015491906109bc565b60405180910390f35b34801561016957600080fd5b506101726104c4565b005b34801561018057600080fd5b5061018961054c565b6040516101969190610b30565b60405180910390f35b6101b960048036038101906101b49190610ae2565b610575565b005b3480156101c757600080fd5b506101e260048036038101906101dd9190610ae2565b610773565b6040516101ef91906109bc565b60405180910390f35b34801561020457600080fd5b5061020d61078b565b60405161021a9190610b7e565b60405180910390f35b34801561022f57600080fd5b5061024a60048036038101906102459190610a08565b6107b1565b6040516102579190610bb4565b60405180910390f35b34801561026c57600080fd5b5061028760048036038101906102829190610bfb565b6107e0565b005b60025481565b600460008381526020019081526020016000205481106102ae57600080fd5b60036000838152602001908152602001600020600082815260200190815260200160002060009054906101000a900460ff16156102ea57600080fd5b600160036000848152602001908152602001600020600083815260200190815260200160002060006101000a81548160ff0219169083151502179055506000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636352211e836040518263ffffffff1660e01b815260040161038491906109bc565b602060405180830381865afa1580156103a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103c59190610c3d565b73ffffffffffffffffffffffffffffffffffffffff1660016000858152602001908152602001600020546040516103fb90610c9b565b60006040518083038185875af1925050503d8060008114610438576040519150601f19603f3d011682016040523d82523d6000602084013e61043d565b606091505b5050905080610481576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047890610d0d565b60405180910390fd5b505050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016020528060005260406000206000915090505481565b6104cc6108d7565b73ffffffffffffffffffffffffffffffffffffffff166104ea61054c565b73ffffffffffffffffffffffffffffffffffffffff1614610540576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053790610d79565b60405180910390fd5b61054a60006108df565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610605576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105fc90610e0b565b60405180910390fd5b600060025490506000813461061a9190610e5a565b9050600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610689573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106ad9190610ea3565b60046000858152602001908152602001600020819055506004600084815260200190815260200160002054816106e39190610eff565b60016000858152602001908152602001600020819055506004600084815260200190815260200160002054816107199190610f30565b600281905550827f75d6e5cc6a300aeeb9df0947f3e2d0a5befe2734b052191b923fb1480cbf82d93484600160008881526020019081526020016000205460405161076693929190610f61565b60405180910390a2505050565b60046020528060005260406000206000915090505481565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60036020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b6107e86108d7565b73ffffffffffffffffffffffffffffffffffffffff1661080661054c565b73ffffffffffffffffffffffffffffffffffffffff161461085c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161085390610d79565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036108cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108c29061100a565b60405180910390fd5b6108d4816108df565b50565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000819050919050565b6109b6816109a3565b82525050565b60006020820190506109d160008301846109ad565b92915050565b600080fd5b6109e5816109a3565b81146109f057600080fd5b50565b600081359050610a02816109dc565b92915050565b60008060408385031215610a1f57610a1e6109d7565b5b6000610a2d858286016109f3565b9250506020610a3e858286016109f3565b9150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000610a8d610a88610a8384610a48565b610a68565b610a48565b9050919050565b6000610a9f82610a72565b9050919050565b6000610ab182610a94565b9050919050565b610ac181610aa6565b82525050565b6000602082019050610adc6000830184610ab8565b92915050565b600060208284031215610af857610af76109d7565b5b6000610b06848285016109f3565b91505092915050565b6000610b1a82610a48565b9050919050565b610b2a81610b0f565b82525050565b6000602082019050610b456000830184610b21565b92915050565b6000610b5682610a72565b9050919050565b6000610b6882610b4b565b9050919050565b610b7881610b5d565b82525050565b6000602082019050610b936000830184610b6f565b92915050565b60008115159050919050565b610bae81610b99565b82525050565b6000602082019050610bc96000830184610ba5565b92915050565b610bd881610b0f565b8114610be357600080fd5b50565b600081359050610bf581610bcf565b92915050565b600060208284031215610c1157610c106109d7565b5b6000610c1f84828501610be6565b91505092915050565b600081519050610c3781610bcf565b92915050565b600060208284031215610c5357610c526109d7565b5b6000610c6184828501610c28565b91505092915050565b600081905092915050565b50565b6000610c85600083610c6a565b9150610c9082610c75565b600082019050919050565b6000610ca682610c78565b9150819050919050565b600082825260208201905092915050565b7f5769746864726177616c206661696c65642e0000000000000000000000000000600082015250565b6000610cf7601283610cb0565b9150610d0282610cc1565b602082019050919050565b60006020820190508181036000830152610d2681610cea565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000610d63602083610cb0565b9150610d6e82610d2d565b602082019050919050565b60006020820190508181036000830152610d9281610d56565b9050919050565b7f4f6e6c792074686520436f736d696347616d6520636f6e74726163742063616e60008201527f206465706f7369742e0000000000000000000000000000000000000000000000602082015250565b6000610df5602983610cb0565b9150610e0082610d99565b604082019050919050565b60006020820190508181036000830152610e2481610de8565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610e65826109a3565b9150610e70836109a3565b9250828201905080821115610e8857610e87610e2b565b5b92915050565b600081519050610e9d816109dc565b92915050565b600060208284031215610eb957610eb86109d7565b5b6000610ec784828501610e8e565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000610f0a826109a3565b9150610f15836109a3565b925082610f2557610f24610ed0565b5b828204905092915050565b6000610f3b826109a3565b9150610f46836109a3565b925082610f5657610f55610ed0565b5b828206905092915050565b6000606082019050610f7660008301866109ad565b610f8360208301856109ad565b610f9060408301846109ad565b949350505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000610ff4602683610cb0565b9150610fff82610f98565b604082019050919050565b6000602082019050818103600083015261102381610fe7565b905091905056fea2646970667358221220be080de8adc81d02250fbbe9ac017e00f2312f996412e778ffe3a64a3d3d708064736f6c63430008130033",
}

// StakingWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingWalletMetaData.ABI instead.
var StakingWalletABI = StakingWalletMetaData.ABI

// StakingWalletBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingWalletMetaData.Bin instead.
var StakingWalletBin = StakingWalletMetaData.Bin

// DeployStakingWallet deploys a new Ethereum contract, binding an instance of StakingWallet to it.
func DeployStakingWallet(auth *bind.TransactOpts, backend bind.ContractBackend, nft_ common.Address, game_ common.Address) (common.Address, *types.Transaction, *StakingWallet, error) {
	parsed, err := StakingWalletMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingWalletBin), backend, nft_, game_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StakingWallet{StakingWalletCaller: StakingWalletCaller{contract: contract}, StakingWalletTransactor: StakingWalletTransactor{contract: contract}, StakingWalletFilterer: StakingWalletFilterer{contract: contract}}, nil
}

// StakingWallet is an auto generated Go binding around an Ethereum contract.
type StakingWallet struct {
	StakingWalletCaller     // Read-only binding to the contract
	StakingWalletTransactor // Write-only binding to the contract
	StakingWalletFilterer   // Log filterer for contract events
}

// StakingWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingWalletSession struct {
	Contract     *StakingWallet    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingWalletCallerSession struct {
	Contract *StakingWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StakingWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingWalletTransactorSession struct {
	Contract     *StakingWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StakingWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingWalletRaw struct {
	Contract *StakingWallet // Generic contract binding to access the raw methods on
}

// StakingWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingWalletCallerRaw struct {
	Contract *StakingWalletCaller // Generic read-only contract binding to access the raw methods on
}

// StakingWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingWalletTransactorRaw struct {
	Contract *StakingWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingWallet creates a new instance of StakingWallet, bound to a specific deployed contract.
func NewStakingWallet(address common.Address, backend bind.ContractBackend) (*StakingWallet, error) {
	contract, err := bindStakingWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingWallet{StakingWalletCaller: StakingWalletCaller{contract: contract}, StakingWalletTransactor: StakingWalletTransactor{contract: contract}, StakingWalletFilterer: StakingWalletFilterer{contract: contract}}, nil
}

// NewStakingWalletCaller creates a new read-only instance of StakingWallet, bound to a specific deployed contract.
func NewStakingWalletCaller(address common.Address, caller bind.ContractCaller) (*StakingWalletCaller, error) {
	contract, err := bindStakingWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingWalletCaller{contract: contract}, nil
}

// NewStakingWalletTransactor creates a new write-only instance of StakingWallet, bound to a specific deployed contract.
func NewStakingWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingWalletTransactor, error) {
	contract, err := bindStakingWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingWalletTransactor{contract: contract}, nil
}

// NewStakingWalletFilterer creates a new log filterer instance of StakingWallet, bound to a specific deployed contract.
func NewStakingWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingWalletFilterer, error) {
	contract, err := bindStakingWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingWalletFilterer{contract: contract}, nil
}

// bindStakingWallet binds a generic wrapper to an already deployed contract.
func bindStakingWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingWallet *StakingWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingWallet.Contract.StakingWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingWallet *StakingWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingWallet.Contract.StakingWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingWallet *StakingWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingWallet.Contract.StakingWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingWallet *StakingWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingWallet *StakingWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingWallet *StakingWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingWallet.Contract.contract.Transact(opts, method, params...)
}

// AmountInRound is a free data retrieval call binding the contract method 0x5658072f.
//
// Solidity: function amountInRound(uint256 ) view returns(uint256)
func (_StakingWallet *StakingWalletCaller) AmountInRound(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingWallet.contract.Call(opts, &out, "amountInRound", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountInRound is a free data retrieval call binding the contract method 0x5658072f.
//
// Solidity: function amountInRound(uint256 ) view returns(uint256)
func (_StakingWallet *StakingWalletSession) AmountInRound(arg0 *big.Int) (*big.Int, error) {
	return _StakingWallet.Contract.AmountInRound(&_StakingWallet.CallOpts, arg0)
}

// AmountInRound is a free data retrieval call binding the contract method 0x5658072f.
//
// Solidity: function amountInRound(uint256 ) view returns(uint256)
func (_StakingWallet *StakingWalletCallerSession) AmountInRound(arg0 *big.Int) (*big.Int, error) {
	return _StakingWallet.Contract.AmountInRound(&_StakingWallet.CallOpts, arg0)
}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_StakingWallet *StakingWalletCaller) Game(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingWallet.contract.Call(opts, &out, "game")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_StakingWallet *StakingWalletSession) Game() (common.Address, error) {
	return _StakingWallet.Contract.Game(&_StakingWallet.CallOpts)
}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_StakingWallet *StakingWalletCallerSession) Game() (common.Address, error) {
	return _StakingWallet.Contract.Game(&_StakingWallet.CallOpts)
}

// IsPaid is a free data retrieval call binding the contract method 0xd823e4a6.
//
// Solidity: function isPaid(uint256 , uint256 ) view returns(bool)
func (_StakingWallet *StakingWalletCaller) IsPaid(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _StakingWallet.contract.Call(opts, &out, "isPaid", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaid is a free data retrieval call binding the contract method 0xd823e4a6.
//
// Solidity: function isPaid(uint256 , uint256 ) view returns(bool)
func (_StakingWallet *StakingWalletSession) IsPaid(arg0 *big.Int, arg1 *big.Int) (bool, error) {
	return _StakingWallet.Contract.IsPaid(&_StakingWallet.CallOpts, arg0, arg1)
}

// IsPaid is a free data retrieval call binding the contract method 0xd823e4a6.
//
// Solidity: function isPaid(uint256 , uint256 ) view returns(bool)
func (_StakingWallet *StakingWalletCallerSession) IsPaid(arg0 *big.Int, arg1 *big.Int) (bool, error) {
	return _StakingWallet.Contract.IsPaid(&_StakingWallet.CallOpts, arg0, arg1)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_StakingWallet *StakingWalletCaller) Nft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingWallet.contract.Call(opts, &out, "nft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_StakingWallet *StakingWalletSession) Nft() (common.Address, error) {
	return _StakingWallet.Contract.Nft(&_StakingWallet.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_StakingWallet *StakingWalletCallerSession) Nft() (common.Address, error) {
	return _StakingWallet.Contract.Nft(&_StakingWallet.CallOpts)
}

// NumWinnersInRound is a free data retrieval call binding the contract method 0xc0450764.
//
// Solidity: function numWinnersInRound(uint256 ) view returns(uint256)
func (_StakingWallet *StakingWalletCaller) NumWinnersInRound(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingWallet.contract.Call(opts, &out, "numWinnersInRound", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumWinnersInRound is a free data retrieval call binding the contract method 0xc0450764.
//
// Solidity: function numWinnersInRound(uint256 ) view returns(uint256)
func (_StakingWallet *StakingWalletSession) NumWinnersInRound(arg0 *big.Int) (*big.Int, error) {
	return _StakingWallet.Contract.NumWinnersInRound(&_StakingWallet.CallOpts, arg0)
}

// NumWinnersInRound is a free data retrieval call binding the contract method 0xc0450764.
//
// Solidity: function numWinnersInRound(uint256 ) view returns(uint256)
func (_StakingWallet *StakingWalletCallerSession) NumWinnersInRound(arg0 *big.Int) (*big.Int, error) {
	return _StakingWallet.Contract.NumWinnersInRound(&_StakingWallet.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingWallet *StakingWalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingWallet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingWallet *StakingWalletSession) Owner() (common.Address, error) {
	return _StakingWallet.Contract.Owner(&_StakingWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingWallet *StakingWalletCallerSession) Owner() (common.Address, error) {
	return _StakingWallet.Contract.Owner(&_StakingWallet.CallOpts)
}

// PreviousRoundReminder is a free data retrieval call binding the contract method 0x1eb2874c.
//
// Solidity: function previousRoundReminder() view returns(uint256)
func (_StakingWallet *StakingWalletCaller) PreviousRoundReminder(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWallet.contract.Call(opts, &out, "previousRoundReminder")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviousRoundReminder is a free data retrieval call binding the contract method 0x1eb2874c.
//
// Solidity: function previousRoundReminder() view returns(uint256)
func (_StakingWallet *StakingWalletSession) PreviousRoundReminder() (*big.Int, error) {
	return _StakingWallet.Contract.PreviousRoundReminder(&_StakingWallet.CallOpts)
}

// PreviousRoundReminder is a free data retrieval call binding the contract method 0x1eb2874c.
//
// Solidity: function previousRoundReminder() view returns(uint256)
func (_StakingWallet *StakingWalletCallerSession) PreviousRoundReminder() (*big.Int, error) {
	return _StakingWallet.Contract.PreviousRoundReminder(&_StakingWallet.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 roundNum) payable returns()
func (_StakingWallet *StakingWalletTransactor) Deposit(opts *bind.TransactOpts, roundNum *big.Int) (*types.Transaction, error) {
	return _StakingWallet.contract.Transact(opts, "deposit", roundNum)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 roundNum) payable returns()
func (_StakingWallet *StakingWalletSession) Deposit(roundNum *big.Int) (*types.Transaction, error) {
	return _StakingWallet.Contract.Deposit(&_StakingWallet.TransactOpts, roundNum)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 roundNum) payable returns()
func (_StakingWallet *StakingWalletTransactorSession) Deposit(roundNum *big.Int) (*types.Transaction, error) {
	return _StakingWallet.Contract.Deposit(&_StakingWallet.TransactOpts, roundNum)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingWallet *StakingWalletTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingWallet.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingWallet *StakingWalletSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingWallet.Contract.RenounceOwnership(&_StakingWallet.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingWallet *StakingWalletTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingWallet.Contract.RenounceOwnership(&_StakingWallet.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingWallet *StakingWalletTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakingWallet.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingWallet *StakingWalletSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingWallet.Contract.TransferOwnership(&_StakingWallet.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingWallet *StakingWalletTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingWallet.Contract.TransferOwnership(&_StakingWallet.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 roundNum, uint256 tokenId) returns()
func (_StakingWallet *StakingWalletTransactor) Withdraw(opts *bind.TransactOpts, roundNum *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _StakingWallet.contract.Transact(opts, "withdraw", roundNum, tokenId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 roundNum, uint256 tokenId) returns()
func (_StakingWallet *StakingWalletSession) Withdraw(roundNum *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _StakingWallet.Contract.Withdraw(&_StakingWallet.TransactOpts, roundNum, tokenId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 roundNum, uint256 tokenId) returns()
func (_StakingWallet *StakingWalletTransactorSession) Withdraw(roundNum *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _StakingWallet.Contract.Withdraw(&_StakingWallet.TransactOpts, roundNum, tokenId)
}

// StakingWalletOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakingWallet contract.
type StakingWalletOwnershipTransferredIterator struct {
	Event *StakingWalletOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingWalletOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletOwnershipTransferred)
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
		it.Event = new(StakingWalletOwnershipTransferred)
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
func (it *StakingWalletOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletOwnershipTransferred represents a OwnershipTransferred event raised by the StakingWallet contract.
type StakingWalletOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingWallet *StakingWalletFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingWalletOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingWallet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletOwnershipTransferredIterator{contract: _StakingWallet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingWallet *StakingWalletFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingWalletOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingWallet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletOwnershipTransferred)
				if err := _StakingWallet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StakingWallet *StakingWalletFilterer) ParseOwnershipTransferred(log types.Log) (*StakingWalletOwnershipTransferred, error) {
	event := new(StakingWalletOwnershipTransferred)
	if err := _StakingWallet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletStakingDepositEventIterator is returned from FilterStakingDepositEvent and is used to iterate over the raw logs and unpacked data for StakingDepositEvent events raised by the StakingWallet contract.
type StakingWalletStakingDepositEventIterator struct {
	Event *StakingWalletStakingDepositEvent // Event containing the contract specifics and raw log

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
func (it *StakingWalletStakingDepositEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletStakingDepositEvent)
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
		it.Event = new(StakingWalletStakingDepositEvent)
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
func (it *StakingWalletStakingDepositEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletStakingDepositEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletStakingDepositEvent represents a StakingDepositEvent event raised by the StakingWallet contract.
type StakingWalletStakingDepositEvent struct {
	Round             *big.Int
	DepositedAmount   *big.Int
	PrevRoundReminder *big.Int
	AmountPerHolder   *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterStakingDepositEvent is a free log retrieval operation binding the contract event 0x75d6e5cc6a300aeeb9df0947f3e2d0a5befe2734b052191b923fb1480cbf82d9.
//
// Solidity: event StakingDepositEvent(uint256 indexed round, uint256 depositedAmount, uint256 prevRoundReminder, uint256 amountPerHolder)
func (_StakingWallet *StakingWalletFilterer) FilterStakingDepositEvent(opts *bind.FilterOpts, round []*big.Int) (*StakingWalletStakingDepositEventIterator, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _StakingWallet.contract.FilterLogs(opts, "StakingDepositEvent", roundRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletStakingDepositEventIterator{contract: _StakingWallet.contract, event: "StakingDepositEvent", logs: logs, sub: sub}, nil
}

// WatchStakingDepositEvent is a free log subscription operation binding the contract event 0x75d6e5cc6a300aeeb9df0947f3e2d0a5befe2734b052191b923fb1480cbf82d9.
//
// Solidity: event StakingDepositEvent(uint256 indexed round, uint256 depositedAmount, uint256 prevRoundReminder, uint256 amountPerHolder)
func (_StakingWallet *StakingWalletFilterer) WatchStakingDepositEvent(opts *bind.WatchOpts, sink chan<- *StakingWalletStakingDepositEvent, round []*big.Int) (event.Subscription, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _StakingWallet.contract.WatchLogs(opts, "StakingDepositEvent", roundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletStakingDepositEvent)
				if err := _StakingWallet.contract.UnpackLog(event, "StakingDepositEvent", log); err != nil {
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

// ParseStakingDepositEvent is a log parse operation binding the contract event 0x75d6e5cc6a300aeeb9df0947f3e2d0a5befe2734b052191b923fb1480cbf82d9.
//
// Solidity: event StakingDepositEvent(uint256 indexed round, uint256 depositedAmount, uint256 prevRoundReminder, uint256 amountPerHolder)
func (_StakingWallet *StakingWalletFilterer) ParseStakingDepositEvent(log types.Log) (*StakingWalletStakingDepositEvent, error) {
	event := new(StakingWalletStakingDepositEvent)
	if err := _StakingWallet.contract.UnpackLog(event, "StakingDepositEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

