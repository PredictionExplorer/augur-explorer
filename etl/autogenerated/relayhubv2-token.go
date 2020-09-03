// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IPaymasterGasLimits is an auto generated low-level Go binding around an user-defined struct.
type IPaymasterGasLimits struct {
	AcceptRelayedCallGasLimit *big.Int
	PreRelayedCallGasLimit    *big.Int
	PostRelayedCallGasLimit   *big.Int
}

// ISignatureVerifierGasData is an auto generated low-level Go binding around an user-defined struct.
type ISignatureVerifierGasData struct {
	GasLimit     *big.Int
	GasPrice     *big.Int
	PctRelayFee  *big.Int
	BaseRelayFee *big.Int
}

// ISignatureVerifierRelayData is an auto generated low-level Go binding around an user-defined struct.
type ISignatureVerifierRelayData struct {
	SenderAddress common.Address
	SenderNonce   *big.Int
	RelayWorker   common.Address
	Paymaster     common.Address
	Forwarder     common.Address
}

// ISignatureVerifierRelayRequest is an auto generated low-level Go binding around an user-defined struct.
type ISignatureVerifierRelayRequest struct {
	Target          common.Address
	EncodedFunction []byte
	GasData         ISignatureVerifierGasData
	RelayData       ISignatureVerifierRelayData
}

// RelayHubV2ABI is the input ABI used to generate the binding from.
const RelayHubV2ABI = "[{\"inputs\":[{\"internalType\":\"contractStakeManager\",\"name\":\"_stakeManager\",\"type\":\"address\"},{\"internalType\":\"contractPenalizer\",\"name\":\"_penalizer\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayWorker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"Penalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseRelayFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pctRelayFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"relayUrl\",\"type\":\"string\"}],\"name\":\"RelayServerRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"newRelayWorkers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"workersCount\",\"type\":\"uint256\"}],\"name\":\"RelayWorkersAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayWorker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"TransactionRejectedByPaymaster\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayWorker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"enumIRelayHub.RelayCallStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"charge\",\"type\":\"uint256\"}],\"name\":\"TransactionRelayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_WORKER_COUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newRelayWorkers\",\"type\":\"address[]\"}],\"name\":\"addRelayWorkers\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pctRelayFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRelayFee\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureVerifier.GasData\",\"name\":\"gasData\",\"type\":\"tuple\"}],\"name\":\"calculateCharge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"encodedFunction\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pctRelayFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRelayFee\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureVerifier.GasData\",\"name\":\"gasData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"senderNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayWorker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"internalType\":\"structISignatureVerifier.RelayData\",\"name\":\"relayData\",\"type\":\"tuple\"}],\"internalType\":\"structISignatureVerifier.RelayRequest\",\"name\":\"relayRequest\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"initialGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approvalData\",\"type\":\"bytes\"}],\"name\":\"canRelay\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnValue\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"acceptRelayedCallGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preRelayedCallGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"postRelayedCallGasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structIPaymaster.GasLimits\",\"name\":\"gasLimits\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"depositFor\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getHubOverhead\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStakeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"encodedFunction\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pctRelayFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRelayFee\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureVerifier.GasData\",\"name\":\"gasData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"senderNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayWorker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"internalType\":\"structISignatureVerifier.RelayData\",\"name\":\"relayData\",\"type\":\"tuple\"}],\"internalType\":\"structISignatureVerifier.RelayRequest\",\"name\":\"relayRequest\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"acceptRelayedCallGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preRelayedCallGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"postRelayedCallGasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structIPaymaster.GasLimits\",\"name\":\"gasLimits\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"totalInitialGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipientContext\",\"type\":\"bytes\"}],\"name\":\"innerRelayCall\",\"outputs\":[{\"internalType\":\"enumIRelayHub.RelayCallStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayWorker\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"penalize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"penalizer\",\"outputs\":[{\"internalType\":\"contractPenalizer\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"baseRelayFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pctRelayFee\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"registerRelayServer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"encodedFunction\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pctRelayFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRelayFee\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureVerifier.GasData\",\"name\":\"gasData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"senderNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayWorker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"internalType\":\"structISignatureVerifier.RelayData\",\"name\":\"relayData\",\"type\":\"tuple\"}],\"internalType\":\"structISignatureVerifier.RelayRequest\",\"name\":\"relayRequest\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"approvalData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"externalGasLimit\",\"type\":\"uint256\"}],\"name\":\"relayCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"paymasterAccepted\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"revertReason\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeManager\",\"outputs\":[{\"internalType\":\"contractStakeManager\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"versionHub\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"dest\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"workerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"workerToManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RelayHubV2 is an auto generated Go binding around an Ethereum contract.
type RelayHubV2 struct {
	RelayHubV2Caller     // Read-only binding to the contract
	RelayHubV2Transactor // Write-only binding to the contract
	RelayHubV2Filterer   // Log filterer for contract events
}

// RelayHubV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type RelayHubV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayHubV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type RelayHubV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayHubV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RelayHubV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayHubV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelayHubV2Session struct {
	Contract     *RelayHubV2       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelayHubV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelayHubV2CallerSession struct {
	Contract *RelayHubV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RelayHubV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelayHubV2TransactorSession struct {
	Contract     *RelayHubV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RelayHubV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type RelayHubV2Raw struct {
	Contract *RelayHubV2 // Generic contract binding to access the raw methods on
}

// RelayHubV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelayHubV2CallerRaw struct {
	Contract *RelayHubV2Caller // Generic read-only contract binding to access the raw methods on
}

// RelayHubV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelayHubV2TransactorRaw struct {
	Contract *RelayHubV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewRelayHubV2 creates a new instance of RelayHubV2, bound to a specific deployed contract.
func NewRelayHubV2(address common.Address, backend bind.ContractBackend) (*RelayHubV2, error) {
	contract, err := bindRelayHubV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RelayHubV2{RelayHubV2Caller: RelayHubV2Caller{contract: contract}, RelayHubV2Transactor: RelayHubV2Transactor{contract: contract}, RelayHubV2Filterer: RelayHubV2Filterer{contract: contract}}, nil
}

// NewRelayHubV2Caller creates a new read-only instance of RelayHubV2, bound to a specific deployed contract.
func NewRelayHubV2Caller(address common.Address, caller bind.ContractCaller) (*RelayHubV2Caller, error) {
	contract, err := bindRelayHubV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RelayHubV2Caller{contract: contract}, nil
}

// NewRelayHubV2Transactor creates a new write-only instance of RelayHubV2, bound to a specific deployed contract.
func NewRelayHubV2Transactor(address common.Address, transactor bind.ContractTransactor) (*RelayHubV2Transactor, error) {
	contract, err := bindRelayHubV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RelayHubV2Transactor{contract: contract}, nil
}

// NewRelayHubV2Filterer creates a new log filterer instance of RelayHubV2, bound to a specific deployed contract.
func NewRelayHubV2Filterer(address common.Address, filterer bind.ContractFilterer) (*RelayHubV2Filterer, error) {
	contract, err := bindRelayHubV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RelayHubV2Filterer{contract: contract}, nil
}

// bindRelayHubV2 binds a generic wrapper to an already deployed contract.
func bindRelayHubV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RelayHubV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelayHubV2 *RelayHubV2Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RelayHubV2.Contract.RelayHubV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelayHubV2 *RelayHubV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayHubV2.Contract.RelayHubV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelayHubV2 *RelayHubV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelayHubV2.Contract.RelayHubV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelayHubV2 *RelayHubV2CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RelayHubV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelayHubV2 *RelayHubV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayHubV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelayHubV2 *RelayHubV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelayHubV2.Contract.contract.Transact(opts, method, params...)
}

// MAXWORKERCOUNT is a free data retrieval call binding the contract method 0x633ed2c3.
//
// Solidity: function MAX_WORKER_COUNT() view returns(uint256)
func (_RelayHubV2 *RelayHubV2Caller) MAXWORKERCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RelayHubV2.contract.Call(opts, out, "MAX_WORKER_COUNT")
	return *ret0, err
}

// MAXWORKERCOUNT is a free data retrieval call binding the contract method 0x633ed2c3.
//
// Solidity: function MAX_WORKER_COUNT() view returns(uint256)
func (_RelayHubV2 *RelayHubV2Session) MAXWORKERCOUNT() (*big.Int, error) {
	return _RelayHubV2.Contract.MAXWORKERCOUNT(&_RelayHubV2.CallOpts)
}

// MAXWORKERCOUNT is a free data retrieval call binding the contract method 0x633ed2c3.
//
// Solidity: function MAX_WORKER_COUNT() view returns(uint256)
func (_RelayHubV2 *RelayHubV2CallerSession) MAXWORKERCOUNT() (*big.Int, error) {
	return _RelayHubV2.Contract.MAXWORKERCOUNT(&_RelayHubV2.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address target) view returns(uint256)
func (_RelayHubV2 *RelayHubV2Caller) BalanceOf(opts *bind.CallOpts, target common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RelayHubV2.contract.Call(opts, out, "balanceOf", target)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address target) view returns(uint256)
func (_RelayHubV2 *RelayHubV2Session) BalanceOf(target common.Address) (*big.Int, error) {
	return _RelayHubV2.Contract.BalanceOf(&_RelayHubV2.CallOpts, target)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address target) view returns(uint256)
func (_RelayHubV2 *RelayHubV2CallerSession) BalanceOf(target common.Address) (*big.Int, error) {
	return _RelayHubV2.Contract.BalanceOf(&_RelayHubV2.CallOpts, target)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_RelayHubV2 *RelayHubV2Caller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RelayHubV2.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_RelayHubV2 *RelayHubV2Session) Balances(arg0 common.Address) (*big.Int, error) {
	return _RelayHubV2.Contract.Balances(&_RelayHubV2.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_RelayHubV2 *RelayHubV2CallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _RelayHubV2.Contract.Balances(&_RelayHubV2.CallOpts, arg0)
}

// CalculateCharge is a free data retrieval call binding the contract method 0x11d77486.
//
// Solidity: function calculateCharge(uint256 gasUsed, ISignatureVerifierGasData gasData) view returns(uint256)
func (_RelayHubV2 *RelayHubV2Caller) CalculateCharge(opts *bind.CallOpts, gasUsed *big.Int, gasData ISignatureVerifierGasData) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RelayHubV2.contract.Call(opts, out, "calculateCharge", gasUsed, gasData)
	return *ret0, err
}

// CalculateCharge is a free data retrieval call binding the contract method 0x11d77486.
//
// Solidity: function calculateCharge(uint256 gasUsed, ISignatureVerifierGasData gasData) view returns(uint256)
func (_RelayHubV2 *RelayHubV2Session) CalculateCharge(gasUsed *big.Int, gasData ISignatureVerifierGasData) (*big.Int, error) {
	return _RelayHubV2.Contract.CalculateCharge(&_RelayHubV2.CallOpts, gasUsed, gasData)
}

// CalculateCharge is a free data retrieval call binding the contract method 0x11d77486.
//
// Solidity: function calculateCharge(uint256 gasUsed, ISignatureVerifierGasData gasData) view returns(uint256)
func (_RelayHubV2 *RelayHubV2CallerSession) CalculateCharge(gasUsed *big.Int, gasData ISignatureVerifierGasData) (*big.Int, error) {
	return _RelayHubV2.Contract.CalculateCharge(&_RelayHubV2.CallOpts, gasUsed, gasData)
}

// CanRelay is a free data retrieval call binding the contract method 0x514d5f54.
//
// Solidity: function canRelay(ISignatureVerifierRelayRequest relayRequest, uint256 initialGas, bytes signature, bytes approvalData) view returns(bool success, bytes returnValue, IPaymasterGasLimits gasLimits)
func (_RelayHubV2 *RelayHubV2Caller) CanRelay(opts *bind.CallOpts, relayRequest ISignatureVerifierRelayRequest, initialGas *big.Int, signature []byte, approvalData []byte) (struct {
	Success     bool
	ReturnValue []byte
	GasLimits   IPaymasterGasLimits
}, error) {
	ret := new(struct {
		Success     bool
		ReturnValue []byte
		GasLimits   IPaymasterGasLimits
	})
	out := ret
	err := _RelayHubV2.contract.Call(opts, out, "canRelay", relayRequest, initialGas, signature, approvalData)
	return *ret, err
}

// CanRelay is a free data retrieval call binding the contract method 0x514d5f54.
//
// Solidity: function canRelay(ISignatureVerifierRelayRequest relayRequest, uint256 initialGas, bytes signature, bytes approvalData) view returns(bool success, bytes returnValue, IPaymasterGasLimits gasLimits)
func (_RelayHubV2 *RelayHubV2Session) CanRelay(relayRequest ISignatureVerifierRelayRequest, initialGas *big.Int, signature []byte, approvalData []byte) (struct {
	Success     bool
	ReturnValue []byte
	GasLimits   IPaymasterGasLimits
}, error) {
	return _RelayHubV2.Contract.CanRelay(&_RelayHubV2.CallOpts, relayRequest, initialGas, signature, approvalData)
}

// CanRelay is a free data retrieval call binding the contract method 0x514d5f54.
//
// Solidity: function canRelay(ISignatureVerifierRelayRequest relayRequest, uint256 initialGas, bytes signature, bytes approvalData) view returns(bool success, bytes returnValue, IPaymasterGasLimits gasLimits)
func (_RelayHubV2 *RelayHubV2CallerSession) CanRelay(relayRequest ISignatureVerifierRelayRequest, initialGas *big.Int, signature []byte, approvalData []byte) (struct {
	Success     bool
	ReturnValue []byte
	GasLimits   IPaymasterGasLimits
}, error) {
	return _RelayHubV2.Contract.CanRelay(&_RelayHubV2.CallOpts, relayRequest, initialGas, signature, approvalData)
}

// GetHubOverhead is a free data retrieval call binding the contract method 0xe1decef3.
//
// Solidity: function getHubOverhead() view returns(uint256)
func (_RelayHubV2 *RelayHubV2Caller) GetHubOverhead(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RelayHubV2.contract.Call(opts, out, "getHubOverhead")
	return *ret0, err
}

// GetHubOverhead is a free data retrieval call binding the contract method 0xe1decef3.
//
// Solidity: function getHubOverhead() view returns(uint256)
func (_RelayHubV2 *RelayHubV2Session) GetHubOverhead() (*big.Int, error) {
	return _RelayHubV2.Contract.GetHubOverhead(&_RelayHubV2.CallOpts)
}

// GetHubOverhead is a free data retrieval call binding the contract method 0xe1decef3.
//
// Solidity: function getHubOverhead() view returns(uint256)
func (_RelayHubV2 *RelayHubV2CallerSession) GetHubOverhead() (*big.Int, error) {
	return _RelayHubV2.Contract.GetHubOverhead(&_RelayHubV2.CallOpts)
}

// GetStakeManager is a free data retrieval call binding the contract method 0xca64f9e7.
//
// Solidity: function getStakeManager() view returns(address)
func (_RelayHubV2 *RelayHubV2Caller) GetStakeManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RelayHubV2.contract.Call(opts, out, "getStakeManager")
	return *ret0, err
}

// GetStakeManager is a free data retrieval call binding the contract method 0xca64f9e7.
//
// Solidity: function getStakeManager() view returns(address)
func (_RelayHubV2 *RelayHubV2Session) GetStakeManager() (common.Address, error) {
	return _RelayHubV2.Contract.GetStakeManager(&_RelayHubV2.CallOpts)
}

// GetStakeManager is a free data retrieval call binding the contract method 0xca64f9e7.
//
// Solidity: function getStakeManager() view returns(address)
func (_RelayHubV2 *RelayHubV2CallerSession) GetStakeManager() (common.Address, error) {
	return _RelayHubV2.Contract.GetStakeManager(&_RelayHubV2.CallOpts)
}

// Penalizer is a free data retrieval call binding the contract method 0xc4775a68.
//
// Solidity: function penalizer() view returns(address)
func (_RelayHubV2 *RelayHubV2Caller) Penalizer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RelayHubV2.contract.Call(opts, out, "penalizer")
	return *ret0, err
}

// Penalizer is a free data retrieval call binding the contract method 0xc4775a68.
//
// Solidity: function penalizer() view returns(address)
func (_RelayHubV2 *RelayHubV2Session) Penalizer() (common.Address, error) {
	return _RelayHubV2.Contract.Penalizer(&_RelayHubV2.CallOpts)
}

// Penalizer is a free data retrieval call binding the contract method 0xc4775a68.
//
// Solidity: function penalizer() view returns(address)
func (_RelayHubV2 *RelayHubV2CallerSession) Penalizer() (common.Address, error) {
	return _RelayHubV2.Contract.Penalizer(&_RelayHubV2.CallOpts)
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_RelayHubV2 *RelayHubV2Caller) StakeManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RelayHubV2.contract.Call(opts, out, "stakeManager")
	return *ret0, err
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_RelayHubV2 *RelayHubV2Session) StakeManager() (common.Address, error) {
	return _RelayHubV2.Contract.StakeManager(&_RelayHubV2.CallOpts)
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_RelayHubV2 *RelayHubV2CallerSession) StakeManager() (common.Address, error) {
	return _RelayHubV2.Contract.StakeManager(&_RelayHubV2.CallOpts)
}

// VersionHub is a free data retrieval call binding the contract method 0xd904c732.
//
// Solidity: function versionHub() view returns(string)
func (_RelayHubV2 *RelayHubV2Caller) VersionHub(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RelayHubV2.contract.Call(opts, out, "versionHub")
	return *ret0, err
}

// VersionHub is a free data retrieval call binding the contract method 0xd904c732.
//
// Solidity: function versionHub() view returns(string)
func (_RelayHubV2 *RelayHubV2Session) VersionHub() (string, error) {
	return _RelayHubV2.Contract.VersionHub(&_RelayHubV2.CallOpts)
}

// VersionHub is a free data retrieval call binding the contract method 0xd904c732.
//
// Solidity: function versionHub() view returns(string)
func (_RelayHubV2 *RelayHubV2CallerSession) VersionHub() (string, error) {
	return _RelayHubV2.Contract.VersionHub(&_RelayHubV2.CallOpts)
}

// WorkerCount is a free data retrieval call binding the contract method 0x194ac307.
//
// Solidity: function workerCount(address ) view returns(uint256)
func (_RelayHubV2 *RelayHubV2Caller) WorkerCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RelayHubV2.contract.Call(opts, out, "workerCount", arg0)
	return *ret0, err
}

// WorkerCount is a free data retrieval call binding the contract method 0x194ac307.
//
// Solidity: function workerCount(address ) view returns(uint256)
func (_RelayHubV2 *RelayHubV2Session) WorkerCount(arg0 common.Address) (*big.Int, error) {
	return _RelayHubV2.Contract.WorkerCount(&_RelayHubV2.CallOpts, arg0)
}

// WorkerCount is a free data retrieval call binding the contract method 0x194ac307.
//
// Solidity: function workerCount(address ) view returns(uint256)
func (_RelayHubV2 *RelayHubV2CallerSession) WorkerCount(arg0 common.Address) (*big.Int, error) {
	return _RelayHubV2.Contract.WorkerCount(&_RelayHubV2.CallOpts, arg0)
}

// WorkerToManager is a free data retrieval call binding the contract method 0xca998f56.
//
// Solidity: function workerToManager(address ) view returns(address)
func (_RelayHubV2 *RelayHubV2Caller) WorkerToManager(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RelayHubV2.contract.Call(opts, out, "workerToManager", arg0)
	return *ret0, err
}

// WorkerToManager is a free data retrieval call binding the contract method 0xca998f56.
//
// Solidity: function workerToManager(address ) view returns(address)
func (_RelayHubV2 *RelayHubV2Session) WorkerToManager(arg0 common.Address) (common.Address, error) {
	return _RelayHubV2.Contract.WorkerToManager(&_RelayHubV2.CallOpts, arg0)
}

// WorkerToManager is a free data retrieval call binding the contract method 0xca998f56.
//
// Solidity: function workerToManager(address ) view returns(address)
func (_RelayHubV2 *RelayHubV2CallerSession) WorkerToManager(arg0 common.Address) (common.Address, error) {
	return _RelayHubV2.Contract.WorkerToManager(&_RelayHubV2.CallOpts, arg0)
}

// AddRelayWorkers is a paid mutator transaction binding the contract method 0xc2da0786.
//
// Solidity: function addRelayWorkers(address[] newRelayWorkers) returns()
func (_RelayHubV2 *RelayHubV2Transactor) AddRelayWorkers(opts *bind.TransactOpts, newRelayWorkers []common.Address) (*types.Transaction, error) {
	return _RelayHubV2.contract.Transact(opts, "addRelayWorkers", newRelayWorkers)
}

// AddRelayWorkers is a paid mutator transaction binding the contract method 0xc2da0786.
//
// Solidity: function addRelayWorkers(address[] newRelayWorkers) returns()
func (_RelayHubV2 *RelayHubV2Session) AddRelayWorkers(newRelayWorkers []common.Address) (*types.Transaction, error) {
	return _RelayHubV2.Contract.AddRelayWorkers(&_RelayHubV2.TransactOpts, newRelayWorkers)
}

// AddRelayWorkers is a paid mutator transaction binding the contract method 0xc2da0786.
//
// Solidity: function addRelayWorkers(address[] newRelayWorkers) returns()
func (_RelayHubV2 *RelayHubV2TransactorSession) AddRelayWorkers(newRelayWorkers []common.Address) (*types.Transaction, error) {
	return _RelayHubV2.Contract.AddRelayWorkers(&_RelayHubV2.TransactOpts, newRelayWorkers)
}

// DepositFor is a paid mutator transaction binding the contract method 0xaa67c919.
//
// Solidity: function depositFor(address target) payable returns()
func (_RelayHubV2 *RelayHubV2Transactor) DepositFor(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _RelayHubV2.contract.Transact(opts, "depositFor", target)
}

// DepositFor is a paid mutator transaction binding the contract method 0xaa67c919.
//
// Solidity: function depositFor(address target) payable returns()
func (_RelayHubV2 *RelayHubV2Session) DepositFor(target common.Address) (*types.Transaction, error) {
	return _RelayHubV2.Contract.DepositFor(&_RelayHubV2.TransactOpts, target)
}

// DepositFor is a paid mutator transaction binding the contract method 0xaa67c919.
//
// Solidity: function depositFor(address target) payable returns()
func (_RelayHubV2 *RelayHubV2TransactorSession) DepositFor(target common.Address) (*types.Transaction, error) {
	return _RelayHubV2.Contract.DepositFor(&_RelayHubV2.TransactOpts, target)
}

// InnerRelayCall is a paid mutator transaction binding the contract method 0x3df1a408.
//
// Solidity: function innerRelayCall(ISignatureVerifierRelayRequest relayRequest, bytes signature, IPaymasterGasLimits gasLimits, uint256 totalInitialGas, bytes recipientContext) returns(uint8)
func (_RelayHubV2 *RelayHubV2Transactor) InnerRelayCall(opts *bind.TransactOpts, relayRequest ISignatureVerifierRelayRequest, signature []byte, gasLimits IPaymasterGasLimits, totalInitialGas *big.Int, recipientContext []byte) (*types.Transaction, error) {
	return _RelayHubV2.contract.Transact(opts, "innerRelayCall", relayRequest, signature, gasLimits, totalInitialGas, recipientContext)
}

// InnerRelayCall is a paid mutator transaction binding the contract method 0x3df1a408.
//
// Solidity: function innerRelayCall(ISignatureVerifierRelayRequest relayRequest, bytes signature, IPaymasterGasLimits gasLimits, uint256 totalInitialGas, bytes recipientContext) returns(uint8)
func (_RelayHubV2 *RelayHubV2Session) InnerRelayCall(relayRequest ISignatureVerifierRelayRequest, signature []byte, gasLimits IPaymasterGasLimits, totalInitialGas *big.Int, recipientContext []byte) (*types.Transaction, error) {
	return _RelayHubV2.Contract.InnerRelayCall(&_RelayHubV2.TransactOpts, relayRequest, signature, gasLimits, totalInitialGas, recipientContext)
}

// InnerRelayCall is a paid mutator transaction binding the contract method 0x3df1a408.
//
// Solidity: function innerRelayCall(ISignatureVerifierRelayRequest relayRequest, bytes signature, IPaymasterGasLimits gasLimits, uint256 totalInitialGas, bytes recipientContext) returns(uint8)
func (_RelayHubV2 *RelayHubV2TransactorSession) InnerRelayCall(relayRequest ISignatureVerifierRelayRequest, signature []byte, gasLimits IPaymasterGasLimits, totalInitialGas *big.Int, recipientContext []byte) (*types.Transaction, error) {
	return _RelayHubV2.Contract.InnerRelayCall(&_RelayHubV2.TransactOpts, relayRequest, signature, gasLimits, totalInitialGas, recipientContext)
}

// Penalize is a paid mutator transaction binding the contract method 0xebcd31ac.
//
// Solidity: function penalize(address relayWorker, address beneficiary) returns()
func (_RelayHubV2 *RelayHubV2Transactor) Penalize(opts *bind.TransactOpts, relayWorker common.Address, beneficiary common.Address) (*types.Transaction, error) {
	return _RelayHubV2.contract.Transact(opts, "penalize", relayWorker, beneficiary)
}

// Penalize is a paid mutator transaction binding the contract method 0xebcd31ac.
//
// Solidity: function penalize(address relayWorker, address beneficiary) returns()
func (_RelayHubV2 *RelayHubV2Session) Penalize(relayWorker common.Address, beneficiary common.Address) (*types.Transaction, error) {
	return _RelayHubV2.Contract.Penalize(&_RelayHubV2.TransactOpts, relayWorker, beneficiary)
}

// Penalize is a paid mutator transaction binding the contract method 0xebcd31ac.
//
// Solidity: function penalize(address relayWorker, address beneficiary) returns()
func (_RelayHubV2 *RelayHubV2TransactorSession) Penalize(relayWorker common.Address, beneficiary common.Address) (*types.Transaction, error) {
	return _RelayHubV2.Contract.Penalize(&_RelayHubV2.TransactOpts, relayWorker, beneficiary)
}

// RegisterRelayServer is a paid mutator transaction binding the contract method 0x83b71871.
//
// Solidity: function registerRelayServer(uint256 baseRelayFee, uint256 pctRelayFee, string url) returns()
func (_RelayHubV2 *RelayHubV2Transactor) RegisterRelayServer(opts *bind.TransactOpts, baseRelayFee *big.Int, pctRelayFee *big.Int, url string) (*types.Transaction, error) {
	return _RelayHubV2.contract.Transact(opts, "registerRelayServer", baseRelayFee, pctRelayFee, url)
}

// RegisterRelayServer is a paid mutator transaction binding the contract method 0x83b71871.
//
// Solidity: function registerRelayServer(uint256 baseRelayFee, uint256 pctRelayFee, string url) returns()
func (_RelayHubV2 *RelayHubV2Session) RegisterRelayServer(baseRelayFee *big.Int, pctRelayFee *big.Int, url string) (*types.Transaction, error) {
	return _RelayHubV2.Contract.RegisterRelayServer(&_RelayHubV2.TransactOpts, baseRelayFee, pctRelayFee, url)
}

// RegisterRelayServer is a paid mutator transaction binding the contract method 0x83b71871.
//
// Solidity: function registerRelayServer(uint256 baseRelayFee, uint256 pctRelayFee, string url) returns()
func (_RelayHubV2 *RelayHubV2TransactorSession) RegisterRelayServer(baseRelayFee *big.Int, pctRelayFee *big.Int, url string) (*types.Transaction, error) {
	return _RelayHubV2.Contract.RegisterRelayServer(&_RelayHubV2.TransactOpts, baseRelayFee, pctRelayFee, url)
}

// RelayCall is a paid mutator transaction binding the contract method 0xf636417f.
//
// Solidity: function relayCall(ISignatureVerifierRelayRequest relayRequest, bytes signature, bytes approvalData, uint256 externalGasLimit) returns(bool paymasterAccepted, string revertReason)
func (_RelayHubV2 *RelayHubV2Transactor) RelayCall(opts *bind.TransactOpts, relayRequest ISignatureVerifierRelayRequest, signature []byte, approvalData []byte, externalGasLimit *big.Int) (*types.Transaction, error) {
	return _RelayHubV2.contract.Transact(opts, "relayCall", relayRequest, signature, approvalData, externalGasLimit)
}

// RelayCall is a paid mutator transaction binding the contract method 0xf636417f.
//
// Solidity: function relayCall(ISignatureVerifierRelayRequest relayRequest, bytes signature, bytes approvalData, uint256 externalGasLimit) returns(bool paymasterAccepted, string revertReason)
func (_RelayHubV2 *RelayHubV2Session) RelayCall(relayRequest ISignatureVerifierRelayRequest, signature []byte, approvalData []byte, externalGasLimit *big.Int) (*types.Transaction, error) {
	return _RelayHubV2.Contract.RelayCall(&_RelayHubV2.TransactOpts, relayRequest, signature, approvalData, externalGasLimit)
}

// RelayCall is a paid mutator transaction binding the contract method 0xf636417f.
//
// Solidity: function relayCall(ISignatureVerifierRelayRequest relayRequest, bytes signature, bytes approvalData, uint256 externalGasLimit) returns(bool paymasterAccepted, string revertReason)
func (_RelayHubV2 *RelayHubV2TransactorSession) RelayCall(relayRequest ISignatureVerifierRelayRequest, signature []byte, approvalData []byte, externalGasLimit *big.Int) (*types.Transaction, error) {
	return _RelayHubV2.Contract.RelayCall(&_RelayHubV2.TransactOpts, relayRequest, signature, approvalData, externalGasLimit)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address dest) returns()
func (_RelayHubV2 *RelayHubV2Transactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, dest common.Address) (*types.Transaction, error) {
	return _RelayHubV2.contract.Transact(opts, "withdraw", amount, dest)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address dest) returns()
func (_RelayHubV2 *RelayHubV2Session) Withdraw(amount *big.Int, dest common.Address) (*types.Transaction, error) {
	return _RelayHubV2.Contract.Withdraw(&_RelayHubV2.TransactOpts, amount, dest)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address dest) returns()
func (_RelayHubV2 *RelayHubV2TransactorSession) Withdraw(amount *big.Int, dest common.Address) (*types.Transaction, error) {
	return _RelayHubV2.Contract.Withdraw(&_RelayHubV2.TransactOpts, amount, dest)
}

// RelayHubV2DepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the RelayHubV2 contract.
type RelayHubV2DepositedIterator struct {
	Event *RelayHubV2Deposited // Event containing the contract specifics and raw log

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
func (it *RelayHubV2DepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayHubV2Deposited)
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
		it.Event = new(RelayHubV2Deposited)
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
func (it *RelayHubV2DepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayHubV2DepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayHubV2Deposited represents a Deposited event raised by the RelayHubV2 contract.
type RelayHubV2Deposited struct {
	Paymaster common.Address
	From      common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed paymaster, address indexed from, uint256 amount)
func (_RelayHubV2 *RelayHubV2Filterer) FilterDeposited(opts *bind.FilterOpts, paymaster []common.Address, from []common.Address) (*RelayHubV2DepositedIterator, error) {

	var paymasterRule []interface{}
	for _, paymasterItem := range paymaster {
		paymasterRule = append(paymasterRule, paymasterItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _RelayHubV2.contract.FilterLogs(opts, "Deposited", paymasterRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &RelayHubV2DepositedIterator{contract: _RelayHubV2.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed paymaster, address indexed from, uint256 amount)
func (_RelayHubV2 *RelayHubV2Filterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *RelayHubV2Deposited, paymaster []common.Address, from []common.Address) (event.Subscription, error) {

	var paymasterRule []interface{}
	for _, paymasterItem := range paymaster {
		paymasterRule = append(paymasterRule, paymasterItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _RelayHubV2.contract.WatchLogs(opts, "Deposited", paymasterRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayHubV2Deposited)
				if err := _RelayHubV2.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed paymaster, address indexed from, uint256 amount)
func (_RelayHubV2 *RelayHubV2Filterer) ParseDeposited(log types.Log) (*RelayHubV2Deposited, error) {
	event := new(RelayHubV2Deposited)
	if err := _RelayHubV2.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayHubV2PenalizedIterator is returned from FilterPenalized and is used to iterate over the raw logs and unpacked data for Penalized events raised by the RelayHubV2 contract.
type RelayHubV2PenalizedIterator struct {
	Event *RelayHubV2Penalized // Event containing the contract specifics and raw log

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
func (it *RelayHubV2PenalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayHubV2Penalized)
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
		it.Event = new(RelayHubV2Penalized)
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
func (it *RelayHubV2PenalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayHubV2PenalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayHubV2Penalized represents a Penalized event raised by the RelayHubV2 contract.
type RelayHubV2Penalized struct {
	RelayWorker common.Address
	Sender      common.Address
	Reward      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPenalized is a free log retrieval operation binding the contract event 0xb0595266ccec357806b2691f348b128209f1060a0bda4f5c95f7090730351ff8.
//
// Solidity: event Penalized(address indexed relayWorker, address sender, uint256 reward)
func (_RelayHubV2 *RelayHubV2Filterer) FilterPenalized(opts *bind.FilterOpts, relayWorker []common.Address) (*RelayHubV2PenalizedIterator, error) {

	var relayWorkerRule []interface{}
	for _, relayWorkerItem := range relayWorker {
		relayWorkerRule = append(relayWorkerRule, relayWorkerItem)
	}

	logs, sub, err := _RelayHubV2.contract.FilterLogs(opts, "Penalized", relayWorkerRule)
	if err != nil {
		return nil, err
	}
	return &RelayHubV2PenalizedIterator{contract: _RelayHubV2.contract, event: "Penalized", logs: logs, sub: sub}, nil
}

// WatchPenalized is a free log subscription operation binding the contract event 0xb0595266ccec357806b2691f348b128209f1060a0bda4f5c95f7090730351ff8.
//
// Solidity: event Penalized(address indexed relayWorker, address sender, uint256 reward)
func (_RelayHubV2 *RelayHubV2Filterer) WatchPenalized(opts *bind.WatchOpts, sink chan<- *RelayHubV2Penalized, relayWorker []common.Address) (event.Subscription, error) {

	var relayWorkerRule []interface{}
	for _, relayWorkerItem := range relayWorker {
		relayWorkerRule = append(relayWorkerRule, relayWorkerItem)
	}

	logs, sub, err := _RelayHubV2.contract.WatchLogs(opts, "Penalized", relayWorkerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayHubV2Penalized)
				if err := _RelayHubV2.contract.UnpackLog(event, "Penalized", log); err != nil {
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

// ParsePenalized is a log parse operation binding the contract event 0xb0595266ccec357806b2691f348b128209f1060a0bda4f5c95f7090730351ff8.
//
// Solidity: event Penalized(address indexed relayWorker, address sender, uint256 reward)
func (_RelayHubV2 *RelayHubV2Filterer) ParsePenalized(log types.Log) (*RelayHubV2Penalized, error) {
	event := new(RelayHubV2Penalized)
	if err := _RelayHubV2.contract.UnpackLog(event, "Penalized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayHubV2RelayServerRegisteredIterator is returned from FilterRelayServerRegistered and is used to iterate over the raw logs and unpacked data for RelayServerRegistered events raised by the RelayHubV2 contract.
type RelayHubV2RelayServerRegisteredIterator struct {
	Event *RelayHubV2RelayServerRegistered // Event containing the contract specifics and raw log

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
func (it *RelayHubV2RelayServerRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayHubV2RelayServerRegistered)
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
		it.Event = new(RelayHubV2RelayServerRegistered)
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
func (it *RelayHubV2RelayServerRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayHubV2RelayServerRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayHubV2RelayServerRegistered represents a RelayServerRegistered event raised by the RelayHubV2 contract.
type RelayHubV2RelayServerRegistered struct {
	RelayManager common.Address
	BaseRelayFee *big.Int
	PctRelayFee  *big.Int
	RelayUrl     string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRelayServerRegistered is a free log retrieval operation binding the contract event 0x77f2d8afec4b9d82ffa0dea525320620292bd1067f575964994d5c4501479aed.
//
// Solidity: event RelayServerRegistered(address indexed relayManager, uint256 baseRelayFee, uint256 pctRelayFee, string relayUrl)
func (_RelayHubV2 *RelayHubV2Filterer) FilterRelayServerRegistered(opts *bind.FilterOpts, relayManager []common.Address) (*RelayHubV2RelayServerRegisteredIterator, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}

	logs, sub, err := _RelayHubV2.contract.FilterLogs(opts, "RelayServerRegistered", relayManagerRule)
	if err != nil {
		return nil, err
	}
	return &RelayHubV2RelayServerRegisteredIterator{contract: _RelayHubV2.contract, event: "RelayServerRegistered", logs: logs, sub: sub}, nil
}

// WatchRelayServerRegistered is a free log subscription operation binding the contract event 0x77f2d8afec4b9d82ffa0dea525320620292bd1067f575964994d5c4501479aed.
//
// Solidity: event RelayServerRegistered(address indexed relayManager, uint256 baseRelayFee, uint256 pctRelayFee, string relayUrl)
func (_RelayHubV2 *RelayHubV2Filterer) WatchRelayServerRegistered(opts *bind.WatchOpts, sink chan<- *RelayHubV2RelayServerRegistered, relayManager []common.Address) (event.Subscription, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}

	logs, sub, err := _RelayHubV2.contract.WatchLogs(opts, "RelayServerRegistered", relayManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayHubV2RelayServerRegistered)
				if err := _RelayHubV2.contract.UnpackLog(event, "RelayServerRegistered", log); err != nil {
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

// ParseRelayServerRegistered is a log parse operation binding the contract event 0x77f2d8afec4b9d82ffa0dea525320620292bd1067f575964994d5c4501479aed.
//
// Solidity: event RelayServerRegistered(address indexed relayManager, uint256 baseRelayFee, uint256 pctRelayFee, string relayUrl)
func (_RelayHubV2 *RelayHubV2Filterer) ParseRelayServerRegistered(log types.Log) (*RelayHubV2RelayServerRegistered, error) {
	event := new(RelayHubV2RelayServerRegistered)
	if err := _RelayHubV2.contract.UnpackLog(event, "RelayServerRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayHubV2RelayWorkersAddedIterator is returned from FilterRelayWorkersAdded and is used to iterate over the raw logs and unpacked data for RelayWorkersAdded events raised by the RelayHubV2 contract.
type RelayHubV2RelayWorkersAddedIterator struct {
	Event *RelayHubV2RelayWorkersAdded // Event containing the contract specifics and raw log

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
func (it *RelayHubV2RelayWorkersAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayHubV2RelayWorkersAdded)
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
		it.Event = new(RelayHubV2RelayWorkersAdded)
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
func (it *RelayHubV2RelayWorkersAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayHubV2RelayWorkersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayHubV2RelayWorkersAdded represents a RelayWorkersAdded event raised by the RelayHubV2 contract.
type RelayHubV2RelayWorkersAdded struct {
	RelayManager    common.Address
	NewRelayWorkers []common.Address
	WorkersCount    *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRelayWorkersAdded is a free log retrieval operation binding the contract event 0xebf4a9bffb39f7c5dbf3f65540183b9381ae226ac3d0a45b4cad484713bd4a28.
//
// Solidity: event RelayWorkersAdded(address indexed relayManager, address[] newRelayWorkers, uint256 workersCount)
func (_RelayHubV2 *RelayHubV2Filterer) FilterRelayWorkersAdded(opts *bind.FilterOpts, relayManager []common.Address) (*RelayHubV2RelayWorkersAddedIterator, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}

	logs, sub, err := _RelayHubV2.contract.FilterLogs(opts, "RelayWorkersAdded", relayManagerRule)
	if err != nil {
		return nil, err
	}
	return &RelayHubV2RelayWorkersAddedIterator{contract: _RelayHubV2.contract, event: "RelayWorkersAdded", logs: logs, sub: sub}, nil
}

// WatchRelayWorkersAdded is a free log subscription operation binding the contract event 0xebf4a9bffb39f7c5dbf3f65540183b9381ae226ac3d0a45b4cad484713bd4a28.
//
// Solidity: event RelayWorkersAdded(address indexed relayManager, address[] newRelayWorkers, uint256 workersCount)
func (_RelayHubV2 *RelayHubV2Filterer) WatchRelayWorkersAdded(opts *bind.WatchOpts, sink chan<- *RelayHubV2RelayWorkersAdded, relayManager []common.Address) (event.Subscription, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}

	logs, sub, err := _RelayHubV2.contract.WatchLogs(opts, "RelayWorkersAdded", relayManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayHubV2RelayWorkersAdded)
				if err := _RelayHubV2.contract.UnpackLog(event, "RelayWorkersAdded", log); err != nil {
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

// ParseRelayWorkersAdded is a log parse operation binding the contract event 0xebf4a9bffb39f7c5dbf3f65540183b9381ae226ac3d0a45b4cad484713bd4a28.
//
// Solidity: event RelayWorkersAdded(address indexed relayManager, address[] newRelayWorkers, uint256 workersCount)
func (_RelayHubV2 *RelayHubV2Filterer) ParseRelayWorkersAdded(log types.Log) (*RelayHubV2RelayWorkersAdded, error) {
	event := new(RelayHubV2RelayWorkersAdded)
	if err := _RelayHubV2.contract.UnpackLog(event, "RelayWorkersAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayHubV2TransactionRejectedByPaymasterIterator is returned from FilterTransactionRejectedByPaymaster and is used to iterate over the raw logs and unpacked data for TransactionRejectedByPaymaster events raised by the RelayHubV2 contract.
type RelayHubV2TransactionRejectedByPaymasterIterator struct {
	Event *RelayHubV2TransactionRejectedByPaymaster // Event containing the contract specifics and raw log

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
func (it *RelayHubV2TransactionRejectedByPaymasterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayHubV2TransactionRejectedByPaymaster)
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
		it.Event = new(RelayHubV2TransactionRejectedByPaymaster)
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
func (it *RelayHubV2TransactionRejectedByPaymasterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayHubV2TransactionRejectedByPaymasterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayHubV2TransactionRejectedByPaymaster represents a TransactionRejectedByPaymaster event raised by the RelayHubV2 contract.
type RelayHubV2TransactionRejectedByPaymaster struct {
	RelayManager common.Address
	Paymaster    common.Address
	From         common.Address
	To           common.Address
	RelayWorker  common.Address
	Selector     [4]byte
	Reason       string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransactionRejectedByPaymaster is a free log retrieval operation binding the contract event 0xe9d1ac88a261e4e2083100c35586e8fd4999357ea18741846147b31defddd9ef.
//
// Solidity: event TransactionRejectedByPaymaster(address indexed relayManager, address indexed paymaster, address indexed from, address to, address relayWorker, bytes4 selector, string reason)
func (_RelayHubV2 *RelayHubV2Filterer) FilterTransactionRejectedByPaymaster(opts *bind.FilterOpts, relayManager []common.Address, paymaster []common.Address, from []common.Address) (*RelayHubV2TransactionRejectedByPaymasterIterator, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}
	var paymasterRule []interface{}
	for _, paymasterItem := range paymaster {
		paymasterRule = append(paymasterRule, paymasterItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _RelayHubV2.contract.FilterLogs(opts, "TransactionRejectedByPaymaster", relayManagerRule, paymasterRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &RelayHubV2TransactionRejectedByPaymasterIterator{contract: _RelayHubV2.contract, event: "TransactionRejectedByPaymaster", logs: logs, sub: sub}, nil
}

// WatchTransactionRejectedByPaymaster is a free log subscription operation binding the contract event 0xe9d1ac88a261e4e2083100c35586e8fd4999357ea18741846147b31defddd9ef.
//
// Solidity: event TransactionRejectedByPaymaster(address indexed relayManager, address indexed paymaster, address indexed from, address to, address relayWorker, bytes4 selector, string reason)
func (_RelayHubV2 *RelayHubV2Filterer) WatchTransactionRejectedByPaymaster(opts *bind.WatchOpts, sink chan<- *RelayHubV2TransactionRejectedByPaymaster, relayManager []common.Address, paymaster []common.Address, from []common.Address) (event.Subscription, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}
	var paymasterRule []interface{}
	for _, paymasterItem := range paymaster {
		paymasterRule = append(paymasterRule, paymasterItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _RelayHubV2.contract.WatchLogs(opts, "TransactionRejectedByPaymaster", relayManagerRule, paymasterRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayHubV2TransactionRejectedByPaymaster)
				if err := _RelayHubV2.contract.UnpackLog(event, "TransactionRejectedByPaymaster", log); err != nil {
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

// ParseTransactionRejectedByPaymaster is a log parse operation binding the contract event 0xe9d1ac88a261e4e2083100c35586e8fd4999357ea18741846147b31defddd9ef.
//
// Solidity: event TransactionRejectedByPaymaster(address indexed relayManager, address indexed paymaster, address indexed from, address to, address relayWorker, bytes4 selector, string reason)
func (_RelayHubV2 *RelayHubV2Filterer) ParseTransactionRejectedByPaymaster(log types.Log) (*RelayHubV2TransactionRejectedByPaymaster, error) {
	event := new(RelayHubV2TransactionRejectedByPaymaster)
	if err := _RelayHubV2.contract.UnpackLog(event, "TransactionRejectedByPaymaster", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayHubV2TransactionRelayedIterator is returned from FilterTransactionRelayed and is used to iterate over the raw logs and unpacked data for TransactionRelayed events raised by the RelayHubV2 contract.
type RelayHubV2TransactionRelayedIterator struct {
	Event *RelayHubV2TransactionRelayed // Event containing the contract specifics and raw log

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
func (it *RelayHubV2TransactionRelayedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayHubV2TransactionRelayed)
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
		it.Event = new(RelayHubV2TransactionRelayed)
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
func (it *RelayHubV2TransactionRelayedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayHubV2TransactionRelayedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayHubV2TransactionRelayed represents a TransactionRelayed event raised by the RelayHubV2 contract.
type RelayHubV2TransactionRelayed struct {
	RelayManager common.Address
	RelayWorker  common.Address
	From         common.Address
	To           common.Address
	Paymaster    common.Address
	Selector     [4]byte
	Status       uint8
	Charge       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransactionRelayed is a free log retrieval operation binding the contract event 0xc9aa709786a3d5fe2cc947abc1ba8cbb0f6decb57aa74b84eb7f558125fee454.
//
// Solidity: event TransactionRelayed(address indexed relayManager, address indexed relayWorker, address indexed from, address to, address paymaster, bytes4 selector, uint8 status, uint256 charge)
func (_RelayHubV2 *RelayHubV2Filterer) FilterTransactionRelayed(opts *bind.FilterOpts, relayManager []common.Address, relayWorker []common.Address, from []common.Address) (*RelayHubV2TransactionRelayedIterator, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}
	var relayWorkerRule []interface{}
	for _, relayWorkerItem := range relayWorker {
		relayWorkerRule = append(relayWorkerRule, relayWorkerItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _RelayHubV2.contract.FilterLogs(opts, "TransactionRelayed", relayManagerRule, relayWorkerRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &RelayHubV2TransactionRelayedIterator{contract: _RelayHubV2.contract, event: "TransactionRelayed", logs: logs, sub: sub}, nil
}

// WatchTransactionRelayed is a free log subscription operation binding the contract event 0xc9aa709786a3d5fe2cc947abc1ba8cbb0f6decb57aa74b84eb7f558125fee454.
//
// Solidity: event TransactionRelayed(address indexed relayManager, address indexed relayWorker, address indexed from, address to, address paymaster, bytes4 selector, uint8 status, uint256 charge)
func (_RelayHubV2 *RelayHubV2Filterer) WatchTransactionRelayed(opts *bind.WatchOpts, sink chan<- *RelayHubV2TransactionRelayed, relayManager []common.Address, relayWorker []common.Address, from []common.Address) (event.Subscription, error) {

	var relayManagerRule []interface{}
	for _, relayManagerItem := range relayManager {
		relayManagerRule = append(relayManagerRule, relayManagerItem)
	}
	var relayWorkerRule []interface{}
	for _, relayWorkerItem := range relayWorker {
		relayWorkerRule = append(relayWorkerRule, relayWorkerItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _RelayHubV2.contract.WatchLogs(opts, "TransactionRelayed", relayManagerRule, relayWorkerRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayHubV2TransactionRelayed)
				if err := _RelayHubV2.contract.UnpackLog(event, "TransactionRelayed", log); err != nil {
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

// ParseTransactionRelayed is a log parse operation binding the contract event 0xc9aa709786a3d5fe2cc947abc1ba8cbb0f6decb57aa74b84eb7f558125fee454.
//
// Solidity: event TransactionRelayed(address indexed relayManager, address indexed relayWorker, address indexed from, address to, address paymaster, bytes4 selector, uint8 status, uint256 charge)
func (_RelayHubV2 *RelayHubV2Filterer) ParseTransactionRelayed(log types.Log) (*RelayHubV2TransactionRelayed, error) {
	event := new(RelayHubV2TransactionRelayed)
	if err := _RelayHubV2.contract.UnpackLog(event, "TransactionRelayed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayHubV2WithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the RelayHubV2 contract.
type RelayHubV2WithdrawnIterator struct {
	Event *RelayHubV2Withdrawn // Event containing the contract specifics and raw log

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
func (it *RelayHubV2WithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayHubV2Withdrawn)
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
		it.Event = new(RelayHubV2Withdrawn)
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
func (it *RelayHubV2WithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayHubV2WithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayHubV2Withdrawn represents a Withdrawn event raised by the RelayHubV2 contract.
type RelayHubV2Withdrawn struct {
	Account common.Address
	Dest    common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address indexed dest, uint256 amount)
func (_RelayHubV2 *RelayHubV2Filterer) FilterWithdrawn(opts *bind.FilterOpts, account []common.Address, dest []common.Address) (*RelayHubV2WithdrawnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var destRule []interface{}
	for _, destItem := range dest {
		destRule = append(destRule, destItem)
	}

	logs, sub, err := _RelayHubV2.contract.FilterLogs(opts, "Withdrawn", accountRule, destRule)
	if err != nil {
		return nil, err
	}
	return &RelayHubV2WithdrawnIterator{contract: _RelayHubV2.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address indexed dest, uint256 amount)
func (_RelayHubV2 *RelayHubV2Filterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *RelayHubV2Withdrawn, account []common.Address, dest []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var destRule []interface{}
	for _, destItem := range dest {
		destRule = append(destRule, destItem)
	}

	logs, sub, err := _RelayHubV2.contract.WatchLogs(opts, "Withdrawn", accountRule, destRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayHubV2Withdrawn)
				if err := _RelayHubV2.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address indexed dest, uint256 amount)
func (_RelayHubV2 *RelayHubV2Filterer) ParseWithdrawn(log types.Log) (*RelayHubV2Withdrawn, error) {
	event := new(RelayHubV2Withdrawn)
	if err := _RelayHubV2.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}
