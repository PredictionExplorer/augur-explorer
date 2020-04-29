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

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"fundingSuccess\",\"type\":\"bool\"}],\"name\":\"ExecuteTransactionStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldRelayHub\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRelayHub\",\"type\":\"address\"}],\"name\":\"RelayHubChanged\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_encodedFunction\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_maxPossibleCharge\",\"type\":\"uint256\"}],\"name\":\"acceptRelayedCall\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_reason\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_context\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"affiliates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"augur\",\"outputs\":[{\"internalType\":\"contractIAugur\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"augurTrading\",\"outputs\":[{\"internalType\":\"contractIAugurTrading\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cash\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"createOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ethExchange\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Exchange\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_payment\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_referralAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_fingerprint\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_desiredSignerBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxExchangeRateInDai\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_revertOnFailure\",\"type\":\"bool\"}],\"name\":\"executeWalletTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fillOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getCreate2WalletAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getHubAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"relay\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"encodedFunction\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"transactionFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"getRelayMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getWallet\",\"outputs\":[{\"internalType\":\"contractIAugurWallet\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIAugur\",\"name\":\"_augur\",\"type\":\"address\"},{\"internalType\":\"contractIAugurTrading\",\"name\":\"_augurTrading\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"actualCharge\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"preRetVal\",\"type\":\"bytes32\"}],\"name\":\"postRelayedCall\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"}],\"name\":\"preRelayedCall\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"relayHubVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"shareToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token0IsCash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wallets\",\"outputs\":[{\"internalType\":\"contractIAugurWallet\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"zeroXTrade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Token *TokenCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "WETH")
	return *ret0, err
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Token *TokenSession) WETH() (common.Address, error) {
	return _Token.Contract.WETH(&_Token.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Token *TokenCallerSession) WETH() (common.Address, error) {
	return _Token.Contract.WETH(&_Token.CallOpts)
}

// AcceptRelayedCall is a free data retrieval call binding the contract method 0x83947ea0.
//
// Solidity: function acceptRelayedCall(address , address _from, bytes _encodedFunction, uint256 , uint256 , uint256 , uint256 , bytes , uint256 _maxPossibleCharge) view returns(uint256 _reason, bytes _context)
func (_Token *TokenCaller) AcceptRelayedCall(opts *bind.CallOpts, arg0 common.Address, _from common.Address, _encodedFunction []byte, arg3 *big.Int, arg4 *big.Int, arg5 *big.Int, arg6 *big.Int, arg7 []byte, _maxPossibleCharge *big.Int) (struct {
	Reason  *big.Int
	Context []byte
}, error) {
	ret := new(struct {
		Reason  *big.Int
		Context []byte
	})
	out := ret
	err := _Token.contract.Call(opts, out, "acceptRelayedCall", arg0, _from, _encodedFunction, arg3, arg4, arg5, arg6, arg7, _maxPossibleCharge)
	return *ret, err
}

// AcceptRelayedCall is a free data retrieval call binding the contract method 0x83947ea0.
//
// Solidity: function acceptRelayedCall(address , address _from, bytes _encodedFunction, uint256 , uint256 , uint256 , uint256 , bytes , uint256 _maxPossibleCharge) view returns(uint256 _reason, bytes _context)
func (_Token *TokenSession) AcceptRelayedCall(arg0 common.Address, _from common.Address, _encodedFunction []byte, arg3 *big.Int, arg4 *big.Int, arg5 *big.Int, arg6 *big.Int, arg7 []byte, _maxPossibleCharge *big.Int) (struct {
	Reason  *big.Int
	Context []byte
}, error) {
	return _Token.Contract.AcceptRelayedCall(&_Token.CallOpts, arg0, _from, _encodedFunction, arg3, arg4, arg5, arg6, arg7, _maxPossibleCharge)
}

// AcceptRelayedCall is a free data retrieval call binding the contract method 0x83947ea0.
//
// Solidity: function acceptRelayedCall(address , address _from, bytes _encodedFunction, uint256 , uint256 , uint256 , uint256 , bytes , uint256 _maxPossibleCharge) view returns(uint256 _reason, bytes _context)
func (_Token *TokenCallerSession) AcceptRelayedCall(arg0 common.Address, _from common.Address, _encodedFunction []byte, arg3 *big.Int, arg4 *big.Int, arg5 *big.Int, arg6 *big.Int, arg7 []byte, _maxPossibleCharge *big.Int) (struct {
	Reason  *big.Int
	Context []byte
}, error) {
	return _Token.Contract.AcceptRelayedCall(&_Token.CallOpts, arg0, _from, _encodedFunction, arg3, arg4, arg5, arg6, arg7, _maxPossibleCharge)
}

// Affiliates is a free data retrieval call binding the contract method 0x3f03842a.
//
// Solidity: function affiliates() view returns(address)
func (_Token *TokenCaller) Affiliates(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "affiliates")
	return *ret0, err
}

// Affiliates is a free data retrieval call binding the contract method 0x3f03842a.
//
// Solidity: function affiliates() view returns(address)
func (_Token *TokenSession) Affiliates() (common.Address, error) {
	return _Token.Contract.Affiliates(&_Token.CallOpts)
}

// Affiliates is a free data retrieval call binding the contract method 0x3f03842a.
//
// Solidity: function affiliates() view returns(address)
func (_Token *TokenCallerSession) Affiliates() (common.Address, error) {
	return _Token.Contract.Affiliates(&_Token.CallOpts)
}

// Augur is a free data retrieval call binding the contract method 0x7a0d8f8a.
//
// Solidity: function augur() view returns(address)
func (_Token *TokenCaller) Augur(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "augur")
	return *ret0, err
}

// Augur is a free data retrieval call binding the contract method 0x7a0d8f8a.
//
// Solidity: function augur() view returns(address)
func (_Token *TokenSession) Augur() (common.Address, error) {
	return _Token.Contract.Augur(&_Token.CallOpts)
}

// Augur is a free data retrieval call binding the contract method 0x7a0d8f8a.
//
// Solidity: function augur() view returns(address)
func (_Token *TokenCallerSession) Augur() (common.Address, error) {
	return _Token.Contract.Augur(&_Token.CallOpts)
}

// AugurTrading is a free data retrieval call binding the contract method 0x8e1bfa73.
//
// Solidity: function augurTrading() view returns(address)
func (_Token *TokenCaller) AugurTrading(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "augurTrading")
	return *ret0, err
}

// AugurTrading is a free data retrieval call binding the contract method 0x8e1bfa73.
//
// Solidity: function augurTrading() view returns(address)
func (_Token *TokenSession) AugurTrading() (common.Address, error) {
	return _Token.Contract.AugurTrading(&_Token.CallOpts)
}

// AugurTrading is a free data retrieval call binding the contract method 0x8e1bfa73.
//
// Solidity: function augurTrading() view returns(address)
func (_Token *TokenCallerSession) AugurTrading() (common.Address, error) {
	return _Token.Contract.AugurTrading(&_Token.CallOpts)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_Token *TokenCaller) Cash(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "cash")
	return *ret0, err
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_Token *TokenSession) Cash() (common.Address, error) {
	return _Token.Contract.Cash(&_Token.CallOpts)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_Token *TokenCallerSession) Cash() (common.Address, error) {
	return _Token.Contract.Cash(&_Token.CallOpts)
}

// CreateOrder is a free data retrieval call binding the contract method 0x6512e6ec.
//
// Solidity: function createOrder() view returns(address)
func (_Token *TokenCaller) CreateOrder(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "createOrder")
	return *ret0, err
}

// CreateOrder is a free data retrieval call binding the contract method 0x6512e6ec.
//
// Solidity: function createOrder() view returns(address)
func (_Token *TokenSession) CreateOrder() (common.Address, error) {
	return _Token.Contract.CreateOrder(&_Token.CallOpts)
}

// CreateOrder is a free data retrieval call binding the contract method 0x6512e6ec.
//
// Solidity: function createOrder() view returns(address)
func (_Token *TokenCallerSession) CreateOrder() (common.Address, error) {
	return _Token.Contract.CreateOrder(&_Token.CallOpts)
}

// EthExchange is a free data retrieval call binding the contract method 0x53e569a1.
//
// Solidity: function ethExchange() view returns(address)
func (_Token *TokenCaller) EthExchange(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "ethExchange")
	return *ret0, err
}

// EthExchange is a free data retrieval call binding the contract method 0x53e569a1.
//
// Solidity: function ethExchange() view returns(address)
func (_Token *TokenSession) EthExchange() (common.Address, error) {
	return _Token.Contract.EthExchange(&_Token.CallOpts)
}

// EthExchange is a free data retrieval call binding the contract method 0x53e569a1.
//
// Solidity: function ethExchange() view returns(address)
func (_Token *TokenCallerSession) EthExchange() (common.Address, error) {
	return _Token.Contract.EthExchange(&_Token.CallOpts)
}

// FillOrder is a free data retrieval call binding the contract method 0x5c1ad844.
//
// Solidity: function fillOrder() view returns(address)
func (_Token *TokenCaller) FillOrder(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "fillOrder")
	return *ret0, err
}

// FillOrder is a free data retrieval call binding the contract method 0x5c1ad844.
//
// Solidity: function fillOrder() view returns(address)
func (_Token *TokenSession) FillOrder() (common.Address, error) {
	return _Token.Contract.FillOrder(&_Token.CallOpts)
}

// FillOrder is a free data retrieval call binding the contract method 0x5c1ad844.
//
// Solidity: function fillOrder() view returns(address)
func (_Token *TokenCallerSession) FillOrder() (common.Address, error) {
	return _Token.Contract.FillOrder(&_Token.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_Token *TokenCaller) GetAmountIn(opts *bind.CallOpts, amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getAmountIn", amountOut, reserveIn, reserveOut)
	return *ret0, err
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_Token *TokenSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Token.Contract.GetAmountIn(&_Token.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_Token *TokenCallerSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Token.Contract.GetAmountIn(&_Token.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Token *TokenCaller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getAmountOut", amountIn, reserveIn, reserveOut)
	return *ret0, err
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Token *TokenSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Token.Contract.GetAmountOut(&_Token.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Token *TokenCallerSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Token.Contract.GetAmountOut(&_Token.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetCreate2WalletAddress is a free data retrieval call binding the contract method 0x80a35930.
//
// Solidity: function getCreate2WalletAddress(address _owner) view returns(address)
func (_Token *TokenCaller) GetCreate2WalletAddress(opts *bind.CallOpts, _owner common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getCreate2WalletAddress", _owner)
	return *ret0, err
}

// GetCreate2WalletAddress is a free data retrieval call binding the contract method 0x80a35930.
//
// Solidity: function getCreate2WalletAddress(address _owner) view returns(address)
func (_Token *TokenSession) GetCreate2WalletAddress(_owner common.Address) (common.Address, error) {
	return _Token.Contract.GetCreate2WalletAddress(&_Token.CallOpts, _owner)
}

// GetCreate2WalletAddress is a free data retrieval call binding the contract method 0x80a35930.
//
// Solidity: function getCreate2WalletAddress(address _owner) view returns(address)
func (_Token *TokenCallerSession) GetCreate2WalletAddress(_owner common.Address) (common.Address, error) {
	return _Token.Contract.GetCreate2WalletAddress(&_Token.CallOpts, _owner)
}

// GetHubAddr is a free data retrieval call binding the contract method 0x74e861d6.
//
// Solidity: function getHubAddr() view returns(address)
func (_Token *TokenCaller) GetHubAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getHubAddr")
	return *ret0, err
}

// GetHubAddr is a free data retrieval call binding the contract method 0x74e861d6.
//
// Solidity: function getHubAddr() view returns(address)
func (_Token *TokenSession) GetHubAddr() (common.Address, error) {
	return _Token.Contract.GetHubAddr(&_Token.CallOpts)
}

// GetHubAddr is a free data retrieval call binding the contract method 0x74e861d6.
//
// Solidity: function getHubAddr() view returns(address)
func (_Token *TokenCallerSession) GetHubAddr() (common.Address, error) {
	return _Token.Contract.GetHubAddr(&_Token.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() view returns(bool)
func (_Token *TokenCaller) GetInitialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getInitialized")
	return *ret0, err
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() view returns(bool)
func (_Token *TokenSession) GetInitialized() (bool, error) {
	return _Token.Contract.GetInitialized(&_Token.CallOpts)
}

// GetInitialized is a free data retrieval call binding the contract method 0xee89dab4.
//
// Solidity: function getInitialized() view returns(bool)
func (_Token *TokenCallerSession) GetInitialized() (bool, error) {
	return _Token.Contract.GetInitialized(&_Token.CallOpts)
}

// GetRelayMessageHash is a free data retrieval call binding the contract method 0xd5b9562b.
//
// Solidity: function getRelayMessageHash(address relay, address from, address to, bytes encodedFunction, uint256 transactionFee, uint256 gasPrice, uint256 gasLimit, uint256 nonce) view returns(bytes32)
func (_Token *TokenCaller) GetRelayMessageHash(opts *bind.CallOpts, relay common.Address, from common.Address, to common.Address, encodedFunction []byte, transactionFee *big.Int, gasPrice *big.Int, gasLimit *big.Int, nonce *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getRelayMessageHash", relay, from, to, encodedFunction, transactionFee, gasPrice, gasLimit, nonce)
	return *ret0, err
}

// GetRelayMessageHash is a free data retrieval call binding the contract method 0xd5b9562b.
//
// Solidity: function getRelayMessageHash(address relay, address from, address to, bytes encodedFunction, uint256 transactionFee, uint256 gasPrice, uint256 gasLimit, uint256 nonce) view returns(bytes32)
func (_Token *TokenSession) GetRelayMessageHash(relay common.Address, from common.Address, to common.Address, encodedFunction []byte, transactionFee *big.Int, gasPrice *big.Int, gasLimit *big.Int, nonce *big.Int) ([32]byte, error) {
	return _Token.Contract.GetRelayMessageHash(&_Token.CallOpts, relay, from, to, encodedFunction, transactionFee, gasPrice, gasLimit, nonce)
}

// GetRelayMessageHash is a free data retrieval call binding the contract method 0xd5b9562b.
//
// Solidity: function getRelayMessageHash(address relay, address from, address to, bytes encodedFunction, uint256 transactionFee, uint256 gasPrice, uint256 gasLimit, uint256 nonce) view returns(bytes32)
func (_Token *TokenCallerSession) GetRelayMessageHash(relay common.Address, from common.Address, to common.Address, encodedFunction []byte, transactionFee *big.Int, gasPrice *big.Int, gasLimit *big.Int, nonce *big.Int) ([32]byte, error) {
	return _Token.Contract.GetRelayMessageHash(&_Token.CallOpts, relay, from, to, encodedFunction, transactionFee, gasPrice, gasLimit, nonce)
}

// GetWallet is a free data retrieval call binding the contract method 0x04d0a647.
//
// Solidity: function getWallet(address _account) view returns(address)
func (_Token *TokenCaller) GetWallet(opts *bind.CallOpts, _account common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getWallet", _account)
	return *ret0, err
}

// GetWallet is a free data retrieval call binding the contract method 0x04d0a647.
//
// Solidity: function getWallet(address _account) view returns(address)
func (_Token *TokenSession) GetWallet(_account common.Address) (common.Address, error) {
	return _Token.Contract.GetWallet(&_Token.CallOpts, _account)
}

// GetWallet is a free data retrieval call binding the contract method 0x04d0a647.
//
// Solidity: function getWallet(address _account) view returns(address)
func (_Token *TokenCallerSession) GetWallet(_account common.Address) (common.Address, error) {
	return _Token.Contract.GetWallet(&_Token.CallOpts, _account)
}

// RelayHubVersion is a free data retrieval call binding the contract method 0xad61ccd5.
//
// Solidity: function relayHubVersion() view returns(string)
func (_Token *TokenCaller) RelayHubVersion(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "relayHubVersion")
	return *ret0, err
}

// RelayHubVersion is a free data retrieval call binding the contract method 0xad61ccd5.
//
// Solidity: function relayHubVersion() view returns(string)
func (_Token *TokenSession) RelayHubVersion() (string, error) {
	return _Token.Contract.RelayHubVersion(&_Token.CallOpts)
}

// RelayHubVersion is a free data retrieval call binding the contract method 0xad61ccd5.
//
// Solidity: function relayHubVersion() view returns(string)
func (_Token *TokenCallerSession) RelayHubVersion() (string, error) {
	return _Token.Contract.RelayHubVersion(&_Token.CallOpts)
}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() view returns(address)
func (_Token *TokenCaller) ShareToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "shareToken")
	return *ret0, err
}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() view returns(address)
func (_Token *TokenSession) ShareToken() (common.Address, error) {
	return _Token.Contract.ShareToken(&_Token.CallOpts)
}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() view returns(address)
func (_Token *TokenCallerSession) ShareToken() (common.Address, error) {
	return _Token.Contract.ShareToken(&_Token.CallOpts)
}

// Token0IsCash is a free data retrieval call binding the contract method 0x7b66c82a.
//
// Solidity: function token0IsCash() view returns(bool)
func (_Token *TokenCaller) Token0IsCash(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "token0IsCash")
	return *ret0, err
}

// Token0IsCash is a free data retrieval call binding the contract method 0x7b66c82a.
//
// Solidity: function token0IsCash() view returns(bool)
func (_Token *TokenSession) Token0IsCash() (bool, error) {
	return _Token.Contract.Token0IsCash(&_Token.CallOpts)
}

// Token0IsCash is a free data retrieval call binding the contract method 0x7b66c82a.
//
// Solidity: function token0IsCash() view returns(bool)
func (_Token *TokenCallerSession) Token0IsCash() (bool, error) {
	return _Token.Contract.Token0IsCash(&_Token.CallOpts)
}

// Wallets is a free data retrieval call binding the contract method 0x89b08f11.
//
// Solidity: function wallets(address ) view returns(address)
func (_Token *TokenCaller) Wallets(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "wallets", arg0)
	return *ret0, err
}

// Wallets is a free data retrieval call binding the contract method 0x89b08f11.
//
// Solidity: function wallets(address ) view returns(address)
func (_Token *TokenSession) Wallets(arg0 common.Address) (common.Address, error) {
	return _Token.Contract.Wallets(&_Token.CallOpts, arg0)
}

// Wallets is a free data retrieval call binding the contract method 0x89b08f11.
//
// Solidity: function wallets(address ) view returns(address)
func (_Token *TokenCallerSession) Wallets(arg0 common.Address) (common.Address, error) {
	return _Token.Contract.Wallets(&_Token.CallOpts, arg0)
}

// ZeroXTrade is a free data retrieval call binding the contract method 0x1a505b76.
//
// Solidity: function zeroXTrade() view returns(address)
func (_Token *TokenCaller) ZeroXTrade(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "zeroXTrade")
	return *ret0, err
}

// ZeroXTrade is a free data retrieval call binding the contract method 0x1a505b76.
//
// Solidity: function zeroXTrade() view returns(address)
func (_Token *TokenSession) ZeroXTrade() (common.Address, error) {
	return _Token.Contract.ZeroXTrade(&_Token.CallOpts)
}

// ZeroXTrade is a free data retrieval call binding the contract method 0x1a505b76.
//
// Solidity: function zeroXTrade() view returns(address)
func (_Token *TokenCallerSession) ZeroXTrade() (common.Address, error) {
	return _Token.Contract.ZeroXTrade(&_Token.CallOpts)
}

// ExecuteWalletTransaction is a paid mutator transaction binding the contract method 0x78dc0eed.
//
// Solidity: function executeWalletTransaction(address _to, bytes _data, uint256 _value, uint256 _payment, address _referralAddress, bytes32 _fingerprint, uint256 _desiredSignerBalance, uint256 _maxExchangeRateInDai, bool _revertOnFailure) returns()
func (_Token *TokenTransactor) ExecuteWalletTransaction(opts *bind.TransactOpts, _to common.Address, _data []byte, _value *big.Int, _payment *big.Int, _referralAddress common.Address, _fingerprint [32]byte, _desiredSignerBalance *big.Int, _maxExchangeRateInDai *big.Int, _revertOnFailure bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "executeWalletTransaction", _to, _data, _value, _payment, _referralAddress, _fingerprint, _desiredSignerBalance, _maxExchangeRateInDai, _revertOnFailure)
}

// ExecuteWalletTransaction is a paid mutator transaction binding the contract method 0x78dc0eed.
//
// Solidity: function executeWalletTransaction(address _to, bytes _data, uint256 _value, uint256 _payment, address _referralAddress, bytes32 _fingerprint, uint256 _desiredSignerBalance, uint256 _maxExchangeRateInDai, bool _revertOnFailure) returns()
func (_Token *TokenSession) ExecuteWalletTransaction(_to common.Address, _data []byte, _value *big.Int, _payment *big.Int, _referralAddress common.Address, _fingerprint [32]byte, _desiredSignerBalance *big.Int, _maxExchangeRateInDai *big.Int, _revertOnFailure bool) (*types.Transaction, error) {
	return _Token.Contract.ExecuteWalletTransaction(&_Token.TransactOpts, _to, _data, _value, _payment, _referralAddress, _fingerprint, _desiredSignerBalance, _maxExchangeRateInDai, _revertOnFailure)
}

// ExecuteWalletTransaction is a paid mutator transaction binding the contract method 0x78dc0eed.
//
// Solidity: function executeWalletTransaction(address _to, bytes _data, uint256 _value, uint256 _payment, address _referralAddress, bytes32 _fingerprint, uint256 _desiredSignerBalance, uint256 _maxExchangeRateInDai, bool _revertOnFailure) returns()
func (_Token *TokenTransactorSession) ExecuteWalletTransaction(_to common.Address, _data []byte, _value *big.Int, _payment *big.Int, _referralAddress common.Address, _fingerprint [32]byte, _desiredSignerBalance *big.Int, _maxExchangeRateInDai *big.Int, _revertOnFailure bool) (*types.Transaction, error) {
	return _Token.Contract.ExecuteWalletTransaction(&_Token.TransactOpts, _to, _data, _value, _payment, _referralAddress, _fingerprint, _desiredSignerBalance, _maxExchangeRateInDai, _revertOnFailure)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _augur, address _augurTrading) payable returns(bool)
func (_Token *TokenTransactor) Initialize(opts *bind.TransactOpts, _augur common.Address, _augurTrading common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "initialize", _augur, _augurTrading)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _augur, address _augurTrading) payable returns(bool)
func (_Token *TokenSession) Initialize(_augur common.Address, _augurTrading common.Address) (*types.Transaction, error) {
	return _Token.Contract.Initialize(&_Token.TransactOpts, _augur, _augurTrading)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _augur, address _augurTrading) payable returns(bool)
func (_Token *TokenTransactorSession) Initialize(_augur common.Address, _augurTrading common.Address) (*types.Transaction, error) {
	return _Token.Contract.Initialize(&_Token.TransactOpts, _augur, _augurTrading)
}

// PostRelayedCall is a paid mutator transaction binding the contract method 0xe06e0e22.
//
// Solidity: function postRelayedCall(bytes context, bool success, uint256 actualCharge, bytes32 preRetVal) returns()
func (_Token *TokenTransactor) PostRelayedCall(opts *bind.TransactOpts, context []byte, success bool, actualCharge *big.Int, preRetVal [32]byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "postRelayedCall", context, success, actualCharge, preRetVal)
}

// PostRelayedCall is a paid mutator transaction binding the contract method 0xe06e0e22.
//
// Solidity: function postRelayedCall(bytes context, bool success, uint256 actualCharge, bytes32 preRetVal) returns()
func (_Token *TokenSession) PostRelayedCall(context []byte, success bool, actualCharge *big.Int, preRetVal [32]byte) (*types.Transaction, error) {
	return _Token.Contract.PostRelayedCall(&_Token.TransactOpts, context, success, actualCharge, preRetVal)
}

// PostRelayedCall is a paid mutator transaction binding the contract method 0xe06e0e22.
//
// Solidity: function postRelayedCall(bytes context, bool success, uint256 actualCharge, bytes32 preRetVal) returns()
func (_Token *TokenTransactorSession) PostRelayedCall(context []byte, success bool, actualCharge *big.Int, preRetVal [32]byte) (*types.Transaction, error) {
	return _Token.Contract.PostRelayedCall(&_Token.TransactOpts, context, success, actualCharge, preRetVal)
}

// PreRelayedCall is a paid mutator transaction binding the contract method 0x80274db7.
//
// Solidity: function preRelayedCall(bytes context) returns(bytes32)
func (_Token *TokenTransactor) PreRelayedCall(opts *bind.TransactOpts, context []byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "preRelayedCall", context)
}

// PreRelayedCall is a paid mutator transaction binding the contract method 0x80274db7.
//
// Solidity: function preRelayedCall(bytes context) returns(bytes32)
func (_Token *TokenSession) PreRelayedCall(context []byte) (*types.Transaction, error) {
	return _Token.Contract.PreRelayedCall(&_Token.TransactOpts, context)
}

// PreRelayedCall is a paid mutator transaction binding the contract method 0x80274db7.
//
// Solidity: function preRelayedCall(bytes context) returns(bytes32)
func (_Token *TokenTransactorSession) PreRelayedCall(context []byte) (*types.Transaction, error) {
	return _Token.Contract.PreRelayedCall(&_Token.TransactOpts, context)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Token *TokenTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Token.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Token *TokenSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Token.Contract.Fallback(&_Token.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Token *TokenTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Token.Contract.Fallback(&_Token.TransactOpts, calldata)
}

// TokenExecuteTransactionStatusIterator is returned from FilterExecuteTransactionStatus and is used to iterate over the raw logs and unpacked data for ExecuteTransactionStatus events raised by the Token contract.
type TokenExecuteTransactionStatusIterator struct {
	Event *TokenExecuteTransactionStatus // Event containing the contract specifics and raw log

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
func (it *TokenExecuteTransactionStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenExecuteTransactionStatus)
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
		it.Event = new(TokenExecuteTransactionStatus)
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
func (it *TokenExecuteTransactionStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenExecuteTransactionStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenExecuteTransactionStatus represents a ExecuteTransactionStatus event raised by the Token contract.
type TokenExecuteTransactionStatus struct {
	Success        bool
	FundingSuccess bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterExecuteTransactionStatus is a free log retrieval operation binding the contract event 0xee9c28a7fe7177d351e891cb4ca5b7a4e4aba4974be67fb7665ba1ad0e703439.
//
// Solidity: event ExecuteTransactionStatus(bool success, bool fundingSuccess)
func (_Token *TokenFilterer) FilterExecuteTransactionStatus(opts *bind.FilterOpts) (*TokenExecuteTransactionStatusIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "ExecuteTransactionStatus")
	if err != nil {
		return nil, err
	}
	return &TokenExecuteTransactionStatusIterator{contract: _Token.contract, event: "ExecuteTransactionStatus", logs: logs, sub: sub}, nil
}

// WatchExecuteTransactionStatus is a free log subscription operation binding the contract event 0xee9c28a7fe7177d351e891cb4ca5b7a4e4aba4974be67fb7665ba1ad0e703439.
//
// Solidity: event ExecuteTransactionStatus(bool success, bool fundingSuccess)
func (_Token *TokenFilterer) WatchExecuteTransactionStatus(opts *bind.WatchOpts, sink chan<- *TokenExecuteTransactionStatus) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "ExecuteTransactionStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenExecuteTransactionStatus)
				if err := _Token.contract.UnpackLog(event, "ExecuteTransactionStatus", log); err != nil {
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

// ParseExecuteTransactionStatus is a log parse operation binding the contract event 0xee9c28a7fe7177d351e891cb4ca5b7a4e4aba4974be67fb7665ba1ad0e703439.
//
// Solidity: event ExecuteTransactionStatus(bool success, bool fundingSuccess)
func (_Token *TokenFilterer) ParseExecuteTransactionStatus(log types.Log) (*TokenExecuteTransactionStatus, error) {
	event := new(TokenExecuteTransactionStatus)
	if err := _Token.contract.UnpackLog(event, "ExecuteTransactionStatus", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenRelayHubChangedIterator is returned from FilterRelayHubChanged and is used to iterate over the raw logs and unpacked data for RelayHubChanged events raised by the Token contract.
type TokenRelayHubChangedIterator struct {
	Event *TokenRelayHubChanged // Event containing the contract specifics and raw log

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
func (it *TokenRelayHubChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRelayHubChanged)
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
		it.Event = new(TokenRelayHubChanged)
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
func (it *TokenRelayHubChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRelayHubChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRelayHubChanged represents a RelayHubChanged event raised by the Token contract.
type TokenRelayHubChanged struct {
	OldRelayHub common.Address
	NewRelayHub common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRelayHubChanged is a free log retrieval operation binding the contract event 0xb9f84b8e65164b14439ae3620df0a4d8786d896996c0282b683f9d8c08f046e8.
//
// Solidity: event RelayHubChanged(address indexed oldRelayHub, address indexed newRelayHub)
func (_Token *TokenFilterer) FilterRelayHubChanged(opts *bind.FilterOpts, oldRelayHub []common.Address, newRelayHub []common.Address) (*TokenRelayHubChangedIterator, error) {

	var oldRelayHubRule []interface{}
	for _, oldRelayHubItem := range oldRelayHub {
		oldRelayHubRule = append(oldRelayHubRule, oldRelayHubItem)
	}
	var newRelayHubRule []interface{}
	for _, newRelayHubItem := range newRelayHub {
		newRelayHubRule = append(newRelayHubRule, newRelayHubItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "RelayHubChanged", oldRelayHubRule, newRelayHubRule)
	if err != nil {
		return nil, err
	}
	return &TokenRelayHubChangedIterator{contract: _Token.contract, event: "RelayHubChanged", logs: logs, sub: sub}, nil
}

// WatchRelayHubChanged is a free log subscription operation binding the contract event 0xb9f84b8e65164b14439ae3620df0a4d8786d896996c0282b683f9d8c08f046e8.
//
// Solidity: event RelayHubChanged(address indexed oldRelayHub, address indexed newRelayHub)
func (_Token *TokenFilterer) WatchRelayHubChanged(opts *bind.WatchOpts, sink chan<- *TokenRelayHubChanged, oldRelayHub []common.Address, newRelayHub []common.Address) (event.Subscription, error) {

	var oldRelayHubRule []interface{}
	for _, oldRelayHubItem := range oldRelayHub {
		oldRelayHubRule = append(oldRelayHubRule, oldRelayHubItem)
	}
	var newRelayHubRule []interface{}
	for _, newRelayHubItem := range newRelayHub {
		newRelayHubRule = append(newRelayHubRule, newRelayHubItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "RelayHubChanged", oldRelayHubRule, newRelayHubRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRelayHubChanged)
				if err := _Token.contract.UnpackLog(event, "RelayHubChanged", log); err != nil {
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

// ParseRelayHubChanged is a log parse operation binding the contract event 0xb9f84b8e65164b14439ae3620df0a4d8786d896996c0282b683f9d8c08f046e8.
//
// Solidity: event RelayHubChanged(address indexed oldRelayHub, address indexed newRelayHub)
func (_Token *TokenFilterer) ParseRelayHubChanged(log types.Log) (*TokenRelayHubChanged, error) {
	event := new(TokenRelayHubChanged)
	if err := _Token.contract.UnpackLog(event, "RelayHubChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}
