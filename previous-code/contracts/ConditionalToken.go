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

// ConditionalTokenMetaData contains all meta data concerning the ConditionalToken contract.
var ConditionalTokenMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"payoutNumerators\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"ids\",\"type\":\"uint256[]\"},{\"name\":\"values\",\"type\":\"uint256[]\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owners\",\"type\":\"address[]\"},{\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"payoutDenominator\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"questionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"outcomeSlotCount\",\"type\":\"uint256\"}],\"name\":\"ConditionPreparation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"questionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"outcomeSlotCount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"payoutNumerators\",\"type\":\"uint256[]\"}],\"name\":\"ConditionResolution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"stakeholder\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"collateralToken\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"parentCollectionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"partition\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PositionSplit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"stakeholder\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"collateralToken\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"parentCollectionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"partition\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PositionsMerge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"collateralToken\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"parentCollectionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"indexSets\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"payout\",\"type\":\"uint256\"}],\"name\":\"PayoutRedemption\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"oracle\",\"type\":\"address\"},{\"name\":\"questionId\",\"type\":\"bytes32\"},{\"name\":\"outcomeSlotCount\",\"type\":\"uint256\"}],\"name\":\"prepareCondition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"questionId\",\"type\":\"bytes32\"},{\"name\":\"payouts\",\"type\":\"uint256[]\"}],\"name\":\"reportPayouts\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"collateralToken\",\"type\":\"address\"},{\"name\":\"parentCollectionId\",\"type\":\"bytes32\"},{\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"name\":\"partition\",\"type\":\"uint256[]\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"splitPosition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"collateralToken\",\"type\":\"address\"},{\"name\":\"parentCollectionId\",\"type\":\"bytes32\"},{\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"name\":\"partition\",\"type\":\"uint256[]\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mergePositions\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"collateralToken\",\"type\":\"address\"},{\"name\":\"parentCollectionId\",\"type\":\"bytes32\"},{\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"name\":\"indexSets\",\"type\":\"uint256[]\"}],\"name\":\"redeemPositions\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"conditionId\",\"type\":\"bytes32\"}],\"name\":\"getOutcomeSlotCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"oracle\",\"type\":\"address\"},{\"name\":\"questionId\",\"type\":\"bytes32\"},{\"name\":\"outcomeSlotCount\",\"type\":\"uint256\"}],\"name\":\"getConditionId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"parentCollectionId\",\"type\":\"bytes32\"},{\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"name\":\"indexSet\",\"type\":\"uint256\"}],\"name\":\"getCollectionId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"collateralToken\",\"type\":\"address\"},{\"name\":\"collectionId\",\"type\":\"bytes32\"}],\"name\":\"getPositionId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// ConditionalTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use ConditionalTokenMetaData.ABI instead.
var ConditionalTokenABI = ConditionalTokenMetaData.ABI

// ConditionalToken is an auto generated Go binding around an Ethereum contract.
type ConditionalToken struct {
	ConditionalTokenCaller     // Read-only binding to the contract
	ConditionalTokenTransactor // Write-only binding to the contract
	ConditionalTokenFilterer   // Log filterer for contract events
}

// ConditionalTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConditionalTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConditionalTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConditionalTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConditionalTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConditionalTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConditionalTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConditionalTokenSession struct {
	Contract     *ConditionalToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConditionalTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConditionalTokenCallerSession struct {
	Contract *ConditionalTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ConditionalTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConditionalTokenTransactorSession struct {
	Contract     *ConditionalTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ConditionalTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConditionalTokenRaw struct {
	Contract *ConditionalToken // Generic contract binding to access the raw methods on
}

// ConditionalTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConditionalTokenCallerRaw struct {
	Contract *ConditionalTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ConditionalTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConditionalTokenTransactorRaw struct {
	Contract *ConditionalTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConditionalToken creates a new instance of ConditionalToken, bound to a specific deployed contract.
func NewConditionalToken(address common.Address, backend bind.ContractBackend) (*ConditionalToken, error) {
	contract, err := bindConditionalToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConditionalToken{ConditionalTokenCaller: ConditionalTokenCaller{contract: contract}, ConditionalTokenTransactor: ConditionalTokenTransactor{contract: contract}, ConditionalTokenFilterer: ConditionalTokenFilterer{contract: contract}}, nil
}

// NewConditionalTokenCaller creates a new read-only instance of ConditionalToken, bound to a specific deployed contract.
func NewConditionalTokenCaller(address common.Address, caller bind.ContractCaller) (*ConditionalTokenCaller, error) {
	contract, err := bindConditionalToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConditionalTokenCaller{contract: contract}, nil
}

// NewConditionalTokenTransactor creates a new write-only instance of ConditionalToken, bound to a specific deployed contract.
func NewConditionalTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ConditionalTokenTransactor, error) {
	contract, err := bindConditionalToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConditionalTokenTransactor{contract: contract}, nil
}

// NewConditionalTokenFilterer creates a new log filterer instance of ConditionalToken, bound to a specific deployed contract.
func NewConditionalTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ConditionalTokenFilterer, error) {
	contract, err := bindConditionalToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConditionalTokenFilterer{contract: contract}, nil
}

// bindConditionalToken binds a generic wrapper to an already deployed contract.
func bindConditionalToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConditionalTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConditionalToken *ConditionalTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConditionalToken.Contract.ConditionalTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConditionalToken *ConditionalTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConditionalToken.Contract.ConditionalTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConditionalToken *ConditionalTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConditionalToken.Contract.ConditionalTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConditionalToken *ConditionalTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConditionalToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConditionalToken *ConditionalTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConditionalToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConditionalToken *ConditionalTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConditionalToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 id) view returns(uint256)
func (_ConditionalToken *ConditionalTokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ConditionalToken.contract.Call(opts, &out, "balanceOf", owner, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 id) view returns(uint256)
func (_ConditionalToken *ConditionalTokenSession) BalanceOf(owner common.Address, id *big.Int) (*big.Int, error) {
	return _ConditionalToken.Contract.BalanceOf(&_ConditionalToken.CallOpts, owner, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 id) view returns(uint256)
func (_ConditionalToken *ConditionalTokenCallerSession) BalanceOf(owner common.Address, id *big.Int) (*big.Int, error) {
	return _ConditionalToken.Contract.BalanceOf(&_ConditionalToken.CallOpts, owner, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ids) view returns(uint256[])
func (_ConditionalToken *ConditionalTokenCaller) BalanceOfBatch(opts *bind.CallOpts, owners []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ConditionalToken.contract.Call(opts, &out, "balanceOfBatch", owners, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ids) view returns(uint256[])
func (_ConditionalToken *ConditionalTokenSession) BalanceOfBatch(owners []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ConditionalToken.Contract.BalanceOfBatch(&_ConditionalToken.CallOpts, owners, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ids) view returns(uint256[])
func (_ConditionalToken *ConditionalTokenCallerSession) BalanceOfBatch(owners []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ConditionalToken.Contract.BalanceOfBatch(&_ConditionalToken.CallOpts, owners, ids)
}

// GetCollectionId is a free data retrieval call binding the contract method 0x856296f7.
//
// Solidity: function getCollectionId(bytes32 parentCollectionId, bytes32 conditionId, uint256 indexSet) view returns(bytes32)
func (_ConditionalToken *ConditionalTokenCaller) GetCollectionId(opts *bind.CallOpts, parentCollectionId [32]byte, conditionId [32]byte, indexSet *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ConditionalToken.contract.Call(opts, &out, "getCollectionId", parentCollectionId, conditionId, indexSet)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCollectionId is a free data retrieval call binding the contract method 0x856296f7.
//
// Solidity: function getCollectionId(bytes32 parentCollectionId, bytes32 conditionId, uint256 indexSet) view returns(bytes32)
func (_ConditionalToken *ConditionalTokenSession) GetCollectionId(parentCollectionId [32]byte, conditionId [32]byte, indexSet *big.Int) ([32]byte, error) {
	return _ConditionalToken.Contract.GetCollectionId(&_ConditionalToken.CallOpts, parentCollectionId, conditionId, indexSet)
}

// GetCollectionId is a free data retrieval call binding the contract method 0x856296f7.
//
// Solidity: function getCollectionId(bytes32 parentCollectionId, bytes32 conditionId, uint256 indexSet) view returns(bytes32)
func (_ConditionalToken *ConditionalTokenCallerSession) GetCollectionId(parentCollectionId [32]byte, conditionId [32]byte, indexSet *big.Int) ([32]byte, error) {
	return _ConditionalToken.Contract.GetCollectionId(&_ConditionalToken.CallOpts, parentCollectionId, conditionId, indexSet)
}

// GetConditionId is a free data retrieval call binding the contract method 0x852c6ae2.
//
// Solidity: function getConditionId(address oracle, bytes32 questionId, uint256 outcomeSlotCount) pure returns(bytes32)
func (_ConditionalToken *ConditionalTokenCaller) GetConditionId(opts *bind.CallOpts, oracle common.Address, questionId [32]byte, outcomeSlotCount *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ConditionalToken.contract.Call(opts, &out, "getConditionId", oracle, questionId, outcomeSlotCount)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetConditionId is a free data retrieval call binding the contract method 0x852c6ae2.
//
// Solidity: function getConditionId(address oracle, bytes32 questionId, uint256 outcomeSlotCount) pure returns(bytes32)
func (_ConditionalToken *ConditionalTokenSession) GetConditionId(oracle common.Address, questionId [32]byte, outcomeSlotCount *big.Int) ([32]byte, error) {
	return _ConditionalToken.Contract.GetConditionId(&_ConditionalToken.CallOpts, oracle, questionId, outcomeSlotCount)
}

// GetConditionId is a free data retrieval call binding the contract method 0x852c6ae2.
//
// Solidity: function getConditionId(address oracle, bytes32 questionId, uint256 outcomeSlotCount) pure returns(bytes32)
func (_ConditionalToken *ConditionalTokenCallerSession) GetConditionId(oracle common.Address, questionId [32]byte, outcomeSlotCount *big.Int) ([32]byte, error) {
	return _ConditionalToken.Contract.GetConditionId(&_ConditionalToken.CallOpts, oracle, questionId, outcomeSlotCount)
}

// GetOutcomeSlotCount is a free data retrieval call binding the contract method 0xd42dc0c2.
//
// Solidity: function getOutcomeSlotCount(bytes32 conditionId) view returns(uint256)
func (_ConditionalToken *ConditionalTokenCaller) GetOutcomeSlotCount(opts *bind.CallOpts, conditionId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ConditionalToken.contract.Call(opts, &out, "getOutcomeSlotCount", conditionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOutcomeSlotCount is a free data retrieval call binding the contract method 0xd42dc0c2.
//
// Solidity: function getOutcomeSlotCount(bytes32 conditionId) view returns(uint256)
func (_ConditionalToken *ConditionalTokenSession) GetOutcomeSlotCount(conditionId [32]byte) (*big.Int, error) {
	return _ConditionalToken.Contract.GetOutcomeSlotCount(&_ConditionalToken.CallOpts, conditionId)
}

// GetOutcomeSlotCount is a free data retrieval call binding the contract method 0xd42dc0c2.
//
// Solidity: function getOutcomeSlotCount(bytes32 conditionId) view returns(uint256)
func (_ConditionalToken *ConditionalTokenCallerSession) GetOutcomeSlotCount(conditionId [32]byte) (*big.Int, error) {
	return _ConditionalToken.Contract.GetOutcomeSlotCount(&_ConditionalToken.CallOpts, conditionId)
}

// GetPositionId is a free data retrieval call binding the contract method 0x39dd7530.
//
// Solidity: function getPositionId(address collateralToken, bytes32 collectionId) pure returns(uint256)
func (_ConditionalToken *ConditionalTokenCaller) GetPositionId(opts *bind.CallOpts, collateralToken common.Address, collectionId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ConditionalToken.contract.Call(opts, &out, "getPositionId", collateralToken, collectionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPositionId is a free data retrieval call binding the contract method 0x39dd7530.
//
// Solidity: function getPositionId(address collateralToken, bytes32 collectionId) pure returns(uint256)
func (_ConditionalToken *ConditionalTokenSession) GetPositionId(collateralToken common.Address, collectionId [32]byte) (*big.Int, error) {
	return _ConditionalToken.Contract.GetPositionId(&_ConditionalToken.CallOpts, collateralToken, collectionId)
}

// GetPositionId is a free data retrieval call binding the contract method 0x39dd7530.
//
// Solidity: function getPositionId(address collateralToken, bytes32 collectionId) pure returns(uint256)
func (_ConditionalToken *ConditionalTokenCallerSession) GetPositionId(collateralToken common.Address, collectionId [32]byte) (*big.Int, error) {
	return _ConditionalToken.Contract.GetPositionId(&_ConditionalToken.CallOpts, collateralToken, collectionId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ConditionalToken *ConditionalTokenCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ConditionalToken.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ConditionalToken *ConditionalTokenSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ConditionalToken.Contract.IsApprovedForAll(&_ConditionalToken.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ConditionalToken *ConditionalTokenCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ConditionalToken.Contract.IsApprovedForAll(&_ConditionalToken.CallOpts, owner, operator)
}

// PayoutDenominator is a free data retrieval call binding the contract method 0xdd34de67.
//
// Solidity: function payoutDenominator(bytes32 ) view returns(uint256)
func (_ConditionalToken *ConditionalTokenCaller) PayoutDenominator(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ConditionalToken.contract.Call(opts, &out, "payoutDenominator", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PayoutDenominator is a free data retrieval call binding the contract method 0xdd34de67.
//
// Solidity: function payoutDenominator(bytes32 ) view returns(uint256)
func (_ConditionalToken *ConditionalTokenSession) PayoutDenominator(arg0 [32]byte) (*big.Int, error) {
	return _ConditionalToken.Contract.PayoutDenominator(&_ConditionalToken.CallOpts, arg0)
}

// PayoutDenominator is a free data retrieval call binding the contract method 0xdd34de67.
//
// Solidity: function payoutDenominator(bytes32 ) view returns(uint256)
func (_ConditionalToken *ConditionalTokenCallerSession) PayoutDenominator(arg0 [32]byte) (*big.Int, error) {
	return _ConditionalToken.Contract.PayoutDenominator(&_ConditionalToken.CallOpts, arg0)
}

// PayoutNumerators is a free data retrieval call binding the contract method 0x0504c814.
//
// Solidity: function payoutNumerators(bytes32 , uint256 ) view returns(uint256)
func (_ConditionalToken *ConditionalTokenCaller) PayoutNumerators(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ConditionalToken.contract.Call(opts, &out, "payoutNumerators", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PayoutNumerators is a free data retrieval call binding the contract method 0x0504c814.
//
// Solidity: function payoutNumerators(bytes32 , uint256 ) view returns(uint256)
func (_ConditionalToken *ConditionalTokenSession) PayoutNumerators(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _ConditionalToken.Contract.PayoutNumerators(&_ConditionalToken.CallOpts, arg0, arg1)
}

// PayoutNumerators is a free data retrieval call binding the contract method 0x0504c814.
//
// Solidity: function payoutNumerators(bytes32 , uint256 ) view returns(uint256)
func (_ConditionalToken *ConditionalTokenCallerSession) PayoutNumerators(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _ConditionalToken.Contract.PayoutNumerators(&_ConditionalToken.CallOpts, arg0, arg1)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ConditionalToken *ConditionalTokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ConditionalToken.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ConditionalToken *ConditionalTokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ConditionalToken.Contract.SupportsInterface(&_ConditionalToken.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ConditionalToken *ConditionalTokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ConditionalToken.Contract.SupportsInterface(&_ConditionalToken.CallOpts, interfaceId)
}

// MergePositions is a paid mutator transaction binding the contract method 0x9e7212ad.
//
// Solidity: function mergePositions(address collateralToken, bytes32 parentCollectionId, bytes32 conditionId, uint256[] partition, uint256 amount) returns()
func (_ConditionalToken *ConditionalTokenTransactor) MergePositions(opts *bind.TransactOpts, collateralToken common.Address, parentCollectionId [32]byte, conditionId [32]byte, partition []*big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ConditionalToken.contract.Transact(opts, "mergePositions", collateralToken, parentCollectionId, conditionId, partition, amount)
}

// MergePositions is a paid mutator transaction binding the contract method 0x9e7212ad.
//
// Solidity: function mergePositions(address collateralToken, bytes32 parentCollectionId, bytes32 conditionId, uint256[] partition, uint256 amount) returns()
func (_ConditionalToken *ConditionalTokenSession) MergePositions(collateralToken common.Address, parentCollectionId [32]byte, conditionId [32]byte, partition []*big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ConditionalToken.Contract.MergePositions(&_ConditionalToken.TransactOpts, collateralToken, parentCollectionId, conditionId, partition, amount)
}

// MergePositions is a paid mutator transaction binding the contract method 0x9e7212ad.
//
// Solidity: function mergePositions(address collateralToken, bytes32 parentCollectionId, bytes32 conditionId, uint256[] partition, uint256 amount) returns()
func (_ConditionalToken *ConditionalTokenTransactorSession) MergePositions(collateralToken common.Address, parentCollectionId [32]byte, conditionId [32]byte, partition []*big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ConditionalToken.Contract.MergePositions(&_ConditionalToken.TransactOpts, collateralToken, parentCollectionId, conditionId, partition, amount)
}

// PrepareCondition is a paid mutator transaction binding the contract method 0xd96ee754.
//
// Solidity: function prepareCondition(address oracle, bytes32 questionId, uint256 outcomeSlotCount) returns()
func (_ConditionalToken *ConditionalTokenTransactor) PrepareCondition(opts *bind.TransactOpts, oracle common.Address, questionId [32]byte, outcomeSlotCount *big.Int) (*types.Transaction, error) {
	return _ConditionalToken.contract.Transact(opts, "prepareCondition", oracle, questionId, outcomeSlotCount)
}

// PrepareCondition is a paid mutator transaction binding the contract method 0xd96ee754.
//
// Solidity: function prepareCondition(address oracle, bytes32 questionId, uint256 outcomeSlotCount) returns()
func (_ConditionalToken *ConditionalTokenSession) PrepareCondition(oracle common.Address, questionId [32]byte, outcomeSlotCount *big.Int) (*types.Transaction, error) {
	return _ConditionalToken.Contract.PrepareCondition(&_ConditionalToken.TransactOpts, oracle, questionId, outcomeSlotCount)
}

// PrepareCondition is a paid mutator transaction binding the contract method 0xd96ee754.
//
// Solidity: function prepareCondition(address oracle, bytes32 questionId, uint256 outcomeSlotCount) returns()
func (_ConditionalToken *ConditionalTokenTransactorSession) PrepareCondition(oracle common.Address, questionId [32]byte, outcomeSlotCount *big.Int) (*types.Transaction, error) {
	return _ConditionalToken.Contract.PrepareCondition(&_ConditionalToken.TransactOpts, oracle, questionId, outcomeSlotCount)
}

// RedeemPositions is a paid mutator transaction binding the contract method 0x01b7037c.
//
// Solidity: function redeemPositions(address collateralToken, bytes32 parentCollectionId, bytes32 conditionId, uint256[] indexSets) returns()
func (_ConditionalToken *ConditionalTokenTransactor) RedeemPositions(opts *bind.TransactOpts, collateralToken common.Address, parentCollectionId [32]byte, conditionId [32]byte, indexSets []*big.Int) (*types.Transaction, error) {
	return _ConditionalToken.contract.Transact(opts, "redeemPositions", collateralToken, parentCollectionId, conditionId, indexSets)
}

// RedeemPositions is a paid mutator transaction binding the contract method 0x01b7037c.
//
// Solidity: function redeemPositions(address collateralToken, bytes32 parentCollectionId, bytes32 conditionId, uint256[] indexSets) returns()
func (_ConditionalToken *ConditionalTokenSession) RedeemPositions(collateralToken common.Address, parentCollectionId [32]byte, conditionId [32]byte, indexSets []*big.Int) (*types.Transaction, error) {
	return _ConditionalToken.Contract.RedeemPositions(&_ConditionalToken.TransactOpts, collateralToken, parentCollectionId, conditionId, indexSets)
}

// RedeemPositions is a paid mutator transaction binding the contract method 0x01b7037c.
//
// Solidity: function redeemPositions(address collateralToken, bytes32 parentCollectionId, bytes32 conditionId, uint256[] indexSets) returns()
func (_ConditionalToken *ConditionalTokenTransactorSession) RedeemPositions(collateralToken common.Address, parentCollectionId [32]byte, conditionId [32]byte, indexSets []*big.Int) (*types.Transaction, error) {
	return _ConditionalToken.Contract.RedeemPositions(&_ConditionalToken.TransactOpts, collateralToken, parentCollectionId, conditionId, indexSets)
}

// ReportPayouts is a paid mutator transaction binding the contract method 0xc49298ac.
//
// Solidity: function reportPayouts(bytes32 questionId, uint256[] payouts) returns()
func (_ConditionalToken *ConditionalTokenTransactor) ReportPayouts(opts *bind.TransactOpts, questionId [32]byte, payouts []*big.Int) (*types.Transaction, error) {
	return _ConditionalToken.contract.Transact(opts, "reportPayouts", questionId, payouts)
}

// ReportPayouts is a paid mutator transaction binding the contract method 0xc49298ac.
//
// Solidity: function reportPayouts(bytes32 questionId, uint256[] payouts) returns()
func (_ConditionalToken *ConditionalTokenSession) ReportPayouts(questionId [32]byte, payouts []*big.Int) (*types.Transaction, error) {
	return _ConditionalToken.Contract.ReportPayouts(&_ConditionalToken.TransactOpts, questionId, payouts)
}

// ReportPayouts is a paid mutator transaction binding the contract method 0xc49298ac.
//
// Solidity: function reportPayouts(bytes32 questionId, uint256[] payouts) returns()
func (_ConditionalToken *ConditionalTokenTransactorSession) ReportPayouts(questionId [32]byte, payouts []*big.Int) (*types.Transaction, error) {
	return _ConditionalToken.Contract.ReportPayouts(&_ConditionalToken.TransactOpts, questionId, payouts)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_ConditionalToken *ConditionalTokenTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _ConditionalToken.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_ConditionalToken *ConditionalTokenSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _ConditionalToken.Contract.SafeBatchTransferFrom(&_ConditionalToken.TransactOpts, from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_ConditionalToken *ConditionalTokenTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _ConditionalToken.Contract.SafeBatchTransferFrom(&_ConditionalToken.TransactOpts, from, to, ids, values, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_ConditionalToken *ConditionalTokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ConditionalToken.contract.Transact(opts, "safeTransferFrom", from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_ConditionalToken *ConditionalTokenSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ConditionalToken.Contract.SafeTransferFrom(&_ConditionalToken.TransactOpts, from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_ConditionalToken *ConditionalTokenTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ConditionalToken.Contract.SafeTransferFrom(&_ConditionalToken.TransactOpts, from, to, id, value, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ConditionalToken *ConditionalTokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ConditionalToken.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ConditionalToken *ConditionalTokenSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ConditionalToken.Contract.SetApprovalForAll(&_ConditionalToken.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ConditionalToken *ConditionalTokenTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ConditionalToken.Contract.SetApprovalForAll(&_ConditionalToken.TransactOpts, operator, approved)
}

// SplitPosition is a paid mutator transaction binding the contract method 0x72ce4275.
//
// Solidity: function splitPosition(address collateralToken, bytes32 parentCollectionId, bytes32 conditionId, uint256[] partition, uint256 amount) returns()
func (_ConditionalToken *ConditionalTokenTransactor) SplitPosition(opts *bind.TransactOpts, collateralToken common.Address, parentCollectionId [32]byte, conditionId [32]byte, partition []*big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ConditionalToken.contract.Transact(opts, "splitPosition", collateralToken, parentCollectionId, conditionId, partition, amount)
}

// SplitPosition is a paid mutator transaction binding the contract method 0x72ce4275.
//
// Solidity: function splitPosition(address collateralToken, bytes32 parentCollectionId, bytes32 conditionId, uint256[] partition, uint256 amount) returns()
func (_ConditionalToken *ConditionalTokenSession) SplitPosition(collateralToken common.Address, parentCollectionId [32]byte, conditionId [32]byte, partition []*big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ConditionalToken.Contract.SplitPosition(&_ConditionalToken.TransactOpts, collateralToken, parentCollectionId, conditionId, partition, amount)
}

// SplitPosition is a paid mutator transaction binding the contract method 0x72ce4275.
//
// Solidity: function splitPosition(address collateralToken, bytes32 parentCollectionId, bytes32 conditionId, uint256[] partition, uint256 amount) returns()
func (_ConditionalToken *ConditionalTokenTransactorSession) SplitPosition(collateralToken common.Address, parentCollectionId [32]byte, conditionId [32]byte, partition []*big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ConditionalToken.Contract.SplitPosition(&_ConditionalToken.TransactOpts, collateralToken, parentCollectionId, conditionId, partition, amount)
}

// ConditionalTokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ConditionalToken contract.
type ConditionalTokenApprovalForAllIterator struct {
	Event *ConditionalTokenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ConditionalTokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConditionalTokenApprovalForAll)
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
		it.Event = new(ConditionalTokenApprovalForAll)
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
func (it *ConditionalTokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConditionalTokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConditionalTokenApprovalForAll represents a ApprovalForAll event raised by the ConditionalToken contract.
type ConditionalTokenApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ConditionalToken *ConditionalTokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ConditionalTokenApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ConditionalToken.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ConditionalTokenApprovalForAllIterator{contract: _ConditionalToken.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ConditionalToken *ConditionalTokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ConditionalTokenApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ConditionalToken.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConditionalTokenApprovalForAll)
				if err := _ConditionalToken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ConditionalToken *ConditionalTokenFilterer) ParseApprovalForAll(log types.Log) (*ConditionalTokenApprovalForAll, error) {
	event := new(ConditionalTokenApprovalForAll)
	if err := _ConditionalToken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConditionalTokenConditionPreparationIterator is returned from FilterConditionPreparation and is used to iterate over the raw logs and unpacked data for ConditionPreparation events raised by the ConditionalToken contract.
type ConditionalTokenConditionPreparationIterator struct {
	Event *ConditionalTokenConditionPreparation // Event containing the contract specifics and raw log

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
func (it *ConditionalTokenConditionPreparationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConditionalTokenConditionPreparation)
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
		it.Event = new(ConditionalTokenConditionPreparation)
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
func (it *ConditionalTokenConditionPreparationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConditionalTokenConditionPreparationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConditionalTokenConditionPreparation represents a ConditionPreparation event raised by the ConditionalToken contract.
type ConditionalTokenConditionPreparation struct {
	ConditionId      [32]byte
	Oracle           common.Address
	QuestionId       [32]byte
	OutcomeSlotCount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterConditionPreparation is a free log retrieval operation binding the contract event 0xab3760c3bd2bb38b5bcf54dc79802ed67338b4cf29f3054ded67ed24661e4177.
//
// Solidity: event ConditionPreparation(bytes32 indexed conditionId, address indexed oracle, bytes32 indexed questionId, uint256 outcomeSlotCount)
func (_ConditionalToken *ConditionalTokenFilterer) FilterConditionPreparation(opts *bind.FilterOpts, conditionId [][32]byte, oracle []common.Address, questionId [][32]byte) (*ConditionalTokenConditionPreparationIterator, error) {

	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}
	var questionIdRule []interface{}
	for _, questionIdItem := range questionId {
		questionIdRule = append(questionIdRule, questionIdItem)
	}

	logs, sub, err := _ConditionalToken.contract.FilterLogs(opts, "ConditionPreparation", conditionIdRule, oracleRule, questionIdRule)
	if err != nil {
		return nil, err
	}
	return &ConditionalTokenConditionPreparationIterator{contract: _ConditionalToken.contract, event: "ConditionPreparation", logs: logs, sub: sub}, nil
}

// WatchConditionPreparation is a free log subscription operation binding the contract event 0xab3760c3bd2bb38b5bcf54dc79802ed67338b4cf29f3054ded67ed24661e4177.
//
// Solidity: event ConditionPreparation(bytes32 indexed conditionId, address indexed oracle, bytes32 indexed questionId, uint256 outcomeSlotCount)
func (_ConditionalToken *ConditionalTokenFilterer) WatchConditionPreparation(opts *bind.WatchOpts, sink chan<- *ConditionalTokenConditionPreparation, conditionId [][32]byte, oracle []common.Address, questionId [][32]byte) (event.Subscription, error) {

	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}
	var questionIdRule []interface{}
	for _, questionIdItem := range questionId {
		questionIdRule = append(questionIdRule, questionIdItem)
	}

	logs, sub, err := _ConditionalToken.contract.WatchLogs(opts, "ConditionPreparation", conditionIdRule, oracleRule, questionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConditionalTokenConditionPreparation)
				if err := _ConditionalToken.contract.UnpackLog(event, "ConditionPreparation", log); err != nil {
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

// ParseConditionPreparation is a log parse operation binding the contract event 0xab3760c3bd2bb38b5bcf54dc79802ed67338b4cf29f3054ded67ed24661e4177.
//
// Solidity: event ConditionPreparation(bytes32 indexed conditionId, address indexed oracle, bytes32 indexed questionId, uint256 outcomeSlotCount)
func (_ConditionalToken *ConditionalTokenFilterer) ParseConditionPreparation(log types.Log) (*ConditionalTokenConditionPreparation, error) {
	event := new(ConditionalTokenConditionPreparation)
	if err := _ConditionalToken.contract.UnpackLog(event, "ConditionPreparation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConditionalTokenConditionResolutionIterator is returned from FilterConditionResolution and is used to iterate over the raw logs and unpacked data for ConditionResolution events raised by the ConditionalToken contract.
type ConditionalTokenConditionResolutionIterator struct {
	Event *ConditionalTokenConditionResolution // Event containing the contract specifics and raw log

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
func (it *ConditionalTokenConditionResolutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConditionalTokenConditionResolution)
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
		it.Event = new(ConditionalTokenConditionResolution)
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
func (it *ConditionalTokenConditionResolutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConditionalTokenConditionResolutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConditionalTokenConditionResolution represents a ConditionResolution event raised by the ConditionalToken contract.
type ConditionalTokenConditionResolution struct {
	ConditionId      [32]byte
	Oracle           common.Address
	QuestionId       [32]byte
	OutcomeSlotCount *big.Int
	PayoutNumerators []*big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterConditionResolution is a free log retrieval operation binding the contract event 0xb44d84d3289691f71497564b85d4233648d9dbae8cbdbb4329f301c3a0185894.
//
// Solidity: event ConditionResolution(bytes32 indexed conditionId, address indexed oracle, bytes32 indexed questionId, uint256 outcomeSlotCount, uint256[] payoutNumerators)
func (_ConditionalToken *ConditionalTokenFilterer) FilterConditionResolution(opts *bind.FilterOpts, conditionId [][32]byte, oracle []common.Address, questionId [][32]byte) (*ConditionalTokenConditionResolutionIterator, error) {

	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}
	var questionIdRule []interface{}
	for _, questionIdItem := range questionId {
		questionIdRule = append(questionIdRule, questionIdItem)
	}

	logs, sub, err := _ConditionalToken.contract.FilterLogs(opts, "ConditionResolution", conditionIdRule, oracleRule, questionIdRule)
	if err != nil {
		return nil, err
	}
	return &ConditionalTokenConditionResolutionIterator{contract: _ConditionalToken.contract, event: "ConditionResolution", logs: logs, sub: sub}, nil
}

// WatchConditionResolution is a free log subscription operation binding the contract event 0xb44d84d3289691f71497564b85d4233648d9dbae8cbdbb4329f301c3a0185894.
//
// Solidity: event ConditionResolution(bytes32 indexed conditionId, address indexed oracle, bytes32 indexed questionId, uint256 outcomeSlotCount, uint256[] payoutNumerators)
func (_ConditionalToken *ConditionalTokenFilterer) WatchConditionResolution(opts *bind.WatchOpts, sink chan<- *ConditionalTokenConditionResolution, conditionId [][32]byte, oracle []common.Address, questionId [][32]byte) (event.Subscription, error) {

	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}
	var questionIdRule []interface{}
	for _, questionIdItem := range questionId {
		questionIdRule = append(questionIdRule, questionIdItem)
	}

	logs, sub, err := _ConditionalToken.contract.WatchLogs(opts, "ConditionResolution", conditionIdRule, oracleRule, questionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConditionalTokenConditionResolution)
				if err := _ConditionalToken.contract.UnpackLog(event, "ConditionResolution", log); err != nil {
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

// ParseConditionResolution is a log parse operation binding the contract event 0xb44d84d3289691f71497564b85d4233648d9dbae8cbdbb4329f301c3a0185894.
//
// Solidity: event ConditionResolution(bytes32 indexed conditionId, address indexed oracle, bytes32 indexed questionId, uint256 outcomeSlotCount, uint256[] payoutNumerators)
func (_ConditionalToken *ConditionalTokenFilterer) ParseConditionResolution(log types.Log) (*ConditionalTokenConditionResolution, error) {
	event := new(ConditionalTokenConditionResolution)
	if err := _ConditionalToken.contract.UnpackLog(event, "ConditionResolution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConditionalTokenPayoutRedemptionIterator is returned from FilterPayoutRedemption and is used to iterate over the raw logs and unpacked data for PayoutRedemption events raised by the ConditionalToken contract.
type ConditionalTokenPayoutRedemptionIterator struct {
	Event *ConditionalTokenPayoutRedemption // Event containing the contract specifics and raw log

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
func (it *ConditionalTokenPayoutRedemptionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConditionalTokenPayoutRedemption)
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
		it.Event = new(ConditionalTokenPayoutRedemption)
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
func (it *ConditionalTokenPayoutRedemptionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConditionalTokenPayoutRedemptionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConditionalTokenPayoutRedemption represents a PayoutRedemption event raised by the ConditionalToken contract.
type ConditionalTokenPayoutRedemption struct {
	Redeemer           common.Address
	CollateralToken    common.Address
	ParentCollectionId [32]byte
	ConditionId        [32]byte
	IndexSets          []*big.Int
	Payout             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPayoutRedemption is a free log retrieval operation binding the contract event 0x2682012a4a4f1973119f1c9b90745d1bd91fa2bab387344f044cb3586864d18d.
//
// Solidity: event PayoutRedemption(address indexed redeemer, address indexed collateralToken, bytes32 indexed parentCollectionId, bytes32 conditionId, uint256[] indexSets, uint256 payout)
func (_ConditionalToken *ConditionalTokenFilterer) FilterPayoutRedemption(opts *bind.FilterOpts, redeemer []common.Address, collateralToken []common.Address, parentCollectionId [][32]byte) (*ConditionalTokenPayoutRedemptionIterator, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var collateralTokenRule []interface{}
	for _, collateralTokenItem := range collateralToken {
		collateralTokenRule = append(collateralTokenRule, collateralTokenItem)
	}
	var parentCollectionIdRule []interface{}
	for _, parentCollectionIdItem := range parentCollectionId {
		parentCollectionIdRule = append(parentCollectionIdRule, parentCollectionIdItem)
	}

	logs, sub, err := _ConditionalToken.contract.FilterLogs(opts, "PayoutRedemption", redeemerRule, collateralTokenRule, parentCollectionIdRule)
	if err != nil {
		return nil, err
	}
	return &ConditionalTokenPayoutRedemptionIterator{contract: _ConditionalToken.contract, event: "PayoutRedemption", logs: logs, sub: sub}, nil
}

// WatchPayoutRedemption is a free log subscription operation binding the contract event 0x2682012a4a4f1973119f1c9b90745d1bd91fa2bab387344f044cb3586864d18d.
//
// Solidity: event PayoutRedemption(address indexed redeemer, address indexed collateralToken, bytes32 indexed parentCollectionId, bytes32 conditionId, uint256[] indexSets, uint256 payout)
func (_ConditionalToken *ConditionalTokenFilterer) WatchPayoutRedemption(opts *bind.WatchOpts, sink chan<- *ConditionalTokenPayoutRedemption, redeemer []common.Address, collateralToken []common.Address, parentCollectionId [][32]byte) (event.Subscription, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var collateralTokenRule []interface{}
	for _, collateralTokenItem := range collateralToken {
		collateralTokenRule = append(collateralTokenRule, collateralTokenItem)
	}
	var parentCollectionIdRule []interface{}
	for _, parentCollectionIdItem := range parentCollectionId {
		parentCollectionIdRule = append(parentCollectionIdRule, parentCollectionIdItem)
	}

	logs, sub, err := _ConditionalToken.contract.WatchLogs(opts, "PayoutRedemption", redeemerRule, collateralTokenRule, parentCollectionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConditionalTokenPayoutRedemption)
				if err := _ConditionalToken.contract.UnpackLog(event, "PayoutRedemption", log); err != nil {
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

// ParsePayoutRedemption is a log parse operation binding the contract event 0x2682012a4a4f1973119f1c9b90745d1bd91fa2bab387344f044cb3586864d18d.
//
// Solidity: event PayoutRedemption(address indexed redeemer, address indexed collateralToken, bytes32 indexed parentCollectionId, bytes32 conditionId, uint256[] indexSets, uint256 payout)
func (_ConditionalToken *ConditionalTokenFilterer) ParsePayoutRedemption(log types.Log) (*ConditionalTokenPayoutRedemption, error) {
	event := new(ConditionalTokenPayoutRedemption)
	if err := _ConditionalToken.contract.UnpackLog(event, "PayoutRedemption", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConditionalTokenPositionSplitIterator is returned from FilterPositionSplit and is used to iterate over the raw logs and unpacked data for PositionSplit events raised by the ConditionalToken contract.
type ConditionalTokenPositionSplitIterator struct {
	Event *ConditionalTokenPositionSplit // Event containing the contract specifics and raw log

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
func (it *ConditionalTokenPositionSplitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConditionalTokenPositionSplit)
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
		it.Event = new(ConditionalTokenPositionSplit)
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
func (it *ConditionalTokenPositionSplitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConditionalTokenPositionSplitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConditionalTokenPositionSplit represents a PositionSplit event raised by the ConditionalToken contract.
type ConditionalTokenPositionSplit struct {
	Stakeholder        common.Address
	CollateralToken    common.Address
	ParentCollectionId [32]byte
	ConditionId        [32]byte
	Partition          []*big.Int
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPositionSplit is a free log retrieval operation binding the contract event 0x2e6bb91f8cbcda0c93623c54d0403a43514fabc40084ec96b6d5379a74786298.
//
// Solidity: event PositionSplit(address indexed stakeholder, address collateralToken, bytes32 indexed parentCollectionId, bytes32 indexed conditionId, uint256[] partition, uint256 amount)
func (_ConditionalToken *ConditionalTokenFilterer) FilterPositionSplit(opts *bind.FilterOpts, stakeholder []common.Address, parentCollectionId [][32]byte, conditionId [][32]byte) (*ConditionalTokenPositionSplitIterator, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}

	var parentCollectionIdRule []interface{}
	for _, parentCollectionIdItem := range parentCollectionId {
		parentCollectionIdRule = append(parentCollectionIdRule, parentCollectionIdItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _ConditionalToken.contract.FilterLogs(opts, "PositionSplit", stakeholderRule, parentCollectionIdRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return &ConditionalTokenPositionSplitIterator{contract: _ConditionalToken.contract, event: "PositionSplit", logs: logs, sub: sub}, nil
}

// WatchPositionSplit is a free log subscription operation binding the contract event 0x2e6bb91f8cbcda0c93623c54d0403a43514fabc40084ec96b6d5379a74786298.
//
// Solidity: event PositionSplit(address indexed stakeholder, address collateralToken, bytes32 indexed parentCollectionId, bytes32 indexed conditionId, uint256[] partition, uint256 amount)
func (_ConditionalToken *ConditionalTokenFilterer) WatchPositionSplit(opts *bind.WatchOpts, sink chan<- *ConditionalTokenPositionSplit, stakeholder []common.Address, parentCollectionId [][32]byte, conditionId [][32]byte) (event.Subscription, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}

	var parentCollectionIdRule []interface{}
	for _, parentCollectionIdItem := range parentCollectionId {
		parentCollectionIdRule = append(parentCollectionIdRule, parentCollectionIdItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _ConditionalToken.contract.WatchLogs(opts, "PositionSplit", stakeholderRule, parentCollectionIdRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConditionalTokenPositionSplit)
				if err := _ConditionalToken.contract.UnpackLog(event, "PositionSplit", log); err != nil {
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

// ParsePositionSplit is a log parse operation binding the contract event 0x2e6bb91f8cbcda0c93623c54d0403a43514fabc40084ec96b6d5379a74786298.
//
// Solidity: event PositionSplit(address indexed stakeholder, address collateralToken, bytes32 indexed parentCollectionId, bytes32 indexed conditionId, uint256[] partition, uint256 amount)
func (_ConditionalToken *ConditionalTokenFilterer) ParsePositionSplit(log types.Log) (*ConditionalTokenPositionSplit, error) {
	event := new(ConditionalTokenPositionSplit)
	if err := _ConditionalToken.contract.UnpackLog(event, "PositionSplit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConditionalTokenPositionsMergeIterator is returned from FilterPositionsMerge and is used to iterate over the raw logs and unpacked data for PositionsMerge events raised by the ConditionalToken contract.
type ConditionalTokenPositionsMergeIterator struct {
	Event *ConditionalTokenPositionsMerge // Event containing the contract specifics and raw log

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
func (it *ConditionalTokenPositionsMergeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConditionalTokenPositionsMerge)
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
		it.Event = new(ConditionalTokenPositionsMerge)
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
func (it *ConditionalTokenPositionsMergeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConditionalTokenPositionsMergeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConditionalTokenPositionsMerge represents a PositionsMerge event raised by the ConditionalToken contract.
type ConditionalTokenPositionsMerge struct {
	Stakeholder        common.Address
	CollateralToken    common.Address
	ParentCollectionId [32]byte
	ConditionId        [32]byte
	Partition          []*big.Int
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPositionsMerge is a free log retrieval operation binding the contract event 0x6f13ca62553fcc2bcd2372180a43949c1e4cebba603901ede2f4e14f36b282ca.
//
// Solidity: event PositionsMerge(address indexed stakeholder, address collateralToken, bytes32 indexed parentCollectionId, bytes32 indexed conditionId, uint256[] partition, uint256 amount)
func (_ConditionalToken *ConditionalTokenFilterer) FilterPositionsMerge(opts *bind.FilterOpts, stakeholder []common.Address, parentCollectionId [][32]byte, conditionId [][32]byte) (*ConditionalTokenPositionsMergeIterator, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}

	var parentCollectionIdRule []interface{}
	for _, parentCollectionIdItem := range parentCollectionId {
		parentCollectionIdRule = append(parentCollectionIdRule, parentCollectionIdItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _ConditionalToken.contract.FilterLogs(opts, "PositionsMerge", stakeholderRule, parentCollectionIdRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return &ConditionalTokenPositionsMergeIterator{contract: _ConditionalToken.contract, event: "PositionsMerge", logs: logs, sub: sub}, nil
}

// WatchPositionsMerge is a free log subscription operation binding the contract event 0x6f13ca62553fcc2bcd2372180a43949c1e4cebba603901ede2f4e14f36b282ca.
//
// Solidity: event PositionsMerge(address indexed stakeholder, address collateralToken, bytes32 indexed parentCollectionId, bytes32 indexed conditionId, uint256[] partition, uint256 amount)
func (_ConditionalToken *ConditionalTokenFilterer) WatchPositionsMerge(opts *bind.WatchOpts, sink chan<- *ConditionalTokenPositionsMerge, stakeholder []common.Address, parentCollectionId [][32]byte, conditionId [][32]byte) (event.Subscription, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}

	var parentCollectionIdRule []interface{}
	for _, parentCollectionIdItem := range parentCollectionId {
		parentCollectionIdRule = append(parentCollectionIdRule, parentCollectionIdItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _ConditionalToken.contract.WatchLogs(opts, "PositionsMerge", stakeholderRule, parentCollectionIdRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConditionalTokenPositionsMerge)
				if err := _ConditionalToken.contract.UnpackLog(event, "PositionsMerge", log); err != nil {
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

// ParsePositionsMerge is a log parse operation binding the contract event 0x6f13ca62553fcc2bcd2372180a43949c1e4cebba603901ede2f4e14f36b282ca.
//
// Solidity: event PositionsMerge(address indexed stakeholder, address collateralToken, bytes32 indexed parentCollectionId, bytes32 indexed conditionId, uint256[] partition, uint256 amount)
func (_ConditionalToken *ConditionalTokenFilterer) ParsePositionsMerge(log types.Log) (*ConditionalTokenPositionsMerge, error) {
	event := new(ConditionalTokenPositionsMerge)
	if err := _ConditionalToken.contract.UnpackLog(event, "PositionsMerge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConditionalTokenTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the ConditionalToken contract.
type ConditionalTokenTransferBatchIterator struct {
	Event *ConditionalTokenTransferBatch // Event containing the contract specifics and raw log

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
func (it *ConditionalTokenTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConditionalTokenTransferBatch)
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
		it.Event = new(ConditionalTokenTransferBatch)
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
func (it *ConditionalTokenTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConditionalTokenTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConditionalTokenTransferBatch represents a TransferBatch event raised by the ConditionalToken contract.
type ConditionalTokenTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ConditionalToken *ConditionalTokenFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ConditionalTokenTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ConditionalToken.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ConditionalTokenTransferBatchIterator{contract: _ConditionalToken.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ConditionalToken *ConditionalTokenFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *ConditionalTokenTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ConditionalToken.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConditionalTokenTransferBatch)
				if err := _ConditionalToken.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ConditionalToken *ConditionalTokenFilterer) ParseTransferBatch(log types.Log) (*ConditionalTokenTransferBatch, error) {
	event := new(ConditionalTokenTransferBatch)
	if err := _ConditionalToken.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConditionalTokenTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the ConditionalToken contract.
type ConditionalTokenTransferSingleIterator struct {
	Event *ConditionalTokenTransferSingle // Event containing the contract specifics and raw log

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
func (it *ConditionalTokenTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConditionalTokenTransferSingle)
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
		it.Event = new(ConditionalTokenTransferSingle)
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
func (it *ConditionalTokenTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConditionalTokenTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConditionalTokenTransferSingle represents a TransferSingle event raised by the ConditionalToken contract.
type ConditionalTokenTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ConditionalToken *ConditionalTokenFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ConditionalTokenTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ConditionalToken.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ConditionalTokenTransferSingleIterator{contract: _ConditionalToken.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ConditionalToken *ConditionalTokenFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *ConditionalTokenTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ConditionalToken.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConditionalTokenTransferSingle)
				if err := _ConditionalToken.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ConditionalToken *ConditionalTokenFilterer) ParseTransferSingle(log types.Log) (*ConditionalTokenTransferSingle, error) {
	event := new(ConditionalTokenTransferSingle)
	if err := _ConditionalToken.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConditionalTokenURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the ConditionalToken contract.
type ConditionalTokenURIIterator struct {
	Event *ConditionalTokenURI // Event containing the contract specifics and raw log

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
func (it *ConditionalTokenURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConditionalTokenURI)
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
		it.Event = new(ConditionalTokenURI)
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
func (it *ConditionalTokenURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConditionalTokenURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConditionalTokenURI represents a URI event raised by the ConditionalToken contract.
type ConditionalTokenURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ConditionalToken *ConditionalTokenFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*ConditionalTokenURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ConditionalToken.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &ConditionalTokenURIIterator{contract: _ConditionalToken.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ConditionalToken *ConditionalTokenFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *ConditionalTokenURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ConditionalToken.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConditionalTokenURI)
				if err := _ConditionalToken.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ConditionalToken *ConditionalTokenFilterer) ParseURI(log types.Log) (*ConditionalTokenURI, error) {
	event := new(ConditionalTokenURI)
	if err := _ConditionalToken.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
