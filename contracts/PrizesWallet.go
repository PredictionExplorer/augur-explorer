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

// IPrizesWalletDonatedTokenToClaim is an auto generated low-level Go binding around an user-defined struct.
type IPrizesWalletDonatedTokenToClaim struct {
	RoundNum     *big.Int
	TokenAddress common.Address
	Amount       *big.Int
}

// IPrizesWalletEthBalanceInfo is an auto generated low-level Go binding around an user-defined struct.
type IPrizesWalletEthBalanceInfo struct {
	RoundNum *big.Int
	Amount   *big.Int
}

// IPrizesWalletEthDeposit is an auto generated low-level Go binding around an user-defined struct.
type IPrizesWalletEthDeposit struct {
	PrizeWinnerAddress common.Address
	Amount             *big.Int
}

// IPrizesWalletMetaData contains all meta data concerning the IPrizesWallet contract.
var IPrizesWalletMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DonatedNftClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DonatedTokenClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prizeWinnerIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prizeWinnerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prizeWinnerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"NftDonated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"TimeoutDurationToWithdrawPrizesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenDonated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"}],\"name\":\"claimDonatedNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"claimDonatedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"indexes_\",\"type\":\"uint256[]\"}],\"name\":\"claimManyDonatedNfts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPrizesWallet.DonatedTokenToClaim[]\",\"name\":\"donatedTokensToClaim_\",\"type\":\"tuple[]\"}],\"name\":\"claimManyDonatedTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prizeWinnerIndex_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"prizeWinnerAddress_\",\"type\":\"address\"}],\"name\":\"depositEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"donorAddress_\",\"type\":\"address\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"donateNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"donorAddress_\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"donateToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"}],\"name\":\"getDonatedTokenBalanceAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEthBalanceInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPrizesWallet.EthBalanceInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prizeWinnerAddress_\",\"type\":\"address\"}],\"name\":\"getEthBalanceInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPrizesWallet.EthBalanceInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mainPrizeBeneficiaryAddress_\",\"type\":\"address\"}],\"name\":\"registerRoundEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mainPrizeBeneficiaryAddress_\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"prizeWinnerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPrizesWallet.EthDeposit[]\",\"name\":\"ethDeposits_\",\"type\":\"tuple[]\"}],\"name\":\"registerRoundEndAndDepositEthMany\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setTimeoutDurationToWithdrawPrizes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prizeWinnerAddress_\",\"type\":\"address\"}],\"name\":\"withdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"withdrawEth_\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPrizesWallet.DonatedTokenToClaim[]\",\"name\":\"donatedTokensToClaim_\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"donatedNftIndexes_\",\"type\":\"uint256[]\"}],\"name\":\"withdrawEverything\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IPrizesWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use IPrizesWalletMetaData.ABI instead.
var IPrizesWalletABI = IPrizesWalletMetaData.ABI

// IPrizesWallet is an auto generated Go binding around an Ethereum contract.
type IPrizesWallet struct {
	IPrizesWalletCaller     // Read-only binding to the contract
	IPrizesWalletTransactor // Write-only binding to the contract
	IPrizesWalletFilterer   // Log filterer for contract events
}

// IPrizesWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPrizesWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPrizesWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPrizesWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPrizesWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPrizesWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPrizesWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPrizesWalletSession struct {
	Contract     *IPrizesWallet    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPrizesWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPrizesWalletCallerSession struct {
	Contract *IPrizesWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IPrizesWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPrizesWalletTransactorSession struct {
	Contract     *IPrizesWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IPrizesWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPrizesWalletRaw struct {
	Contract *IPrizesWallet // Generic contract binding to access the raw methods on
}

// IPrizesWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPrizesWalletCallerRaw struct {
	Contract *IPrizesWalletCaller // Generic read-only contract binding to access the raw methods on
}

// IPrizesWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPrizesWalletTransactorRaw struct {
	Contract *IPrizesWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPrizesWallet creates a new instance of IPrizesWallet, bound to a specific deployed contract.
func NewIPrizesWallet(address common.Address, backend bind.ContractBackend) (*IPrizesWallet, error) {
	contract, err := bindIPrizesWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPrizesWallet{IPrizesWalletCaller: IPrizesWalletCaller{contract: contract}, IPrizesWalletTransactor: IPrizesWalletTransactor{contract: contract}, IPrizesWalletFilterer: IPrizesWalletFilterer{contract: contract}}, nil
}

// NewIPrizesWalletCaller creates a new read-only instance of IPrizesWallet, bound to a specific deployed contract.
func NewIPrizesWalletCaller(address common.Address, caller bind.ContractCaller) (*IPrizesWalletCaller, error) {
	contract, err := bindIPrizesWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletCaller{contract: contract}, nil
}

// NewIPrizesWalletTransactor creates a new write-only instance of IPrizesWallet, bound to a specific deployed contract.
func NewIPrizesWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*IPrizesWalletTransactor, error) {
	contract, err := bindIPrizesWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletTransactor{contract: contract}, nil
}

// NewIPrizesWalletFilterer creates a new log filterer instance of IPrizesWallet, bound to a specific deployed contract.
func NewIPrizesWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*IPrizesWalletFilterer, error) {
	contract, err := bindIPrizesWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletFilterer{contract: contract}, nil
}

// bindIPrizesWallet binds a generic wrapper to an already deployed contract.
func bindIPrizesWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPrizesWalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPrizesWallet *IPrizesWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPrizesWallet.Contract.IPrizesWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPrizesWallet *IPrizesWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.IPrizesWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPrizesWallet *IPrizesWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.IPrizesWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPrizesWallet *IPrizesWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPrizesWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPrizesWallet *IPrizesWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPrizesWallet *IPrizesWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.contract.Transact(opts, method, params...)
}

// GetDonatedTokenBalanceAmount is a free data retrieval call binding the contract method 0xe86a49d7.
//
// Solidity: function getDonatedTokenBalanceAmount(uint256 roundNum_, address tokenAddress_) view returns(uint256)
func (_IPrizesWallet *IPrizesWalletCaller) GetDonatedTokenBalanceAmount(opts *bind.CallOpts, roundNum_ *big.Int, tokenAddress_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IPrizesWallet.contract.Call(opts, &out, "getDonatedTokenBalanceAmount", roundNum_, tokenAddress_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDonatedTokenBalanceAmount is a free data retrieval call binding the contract method 0xe86a49d7.
//
// Solidity: function getDonatedTokenBalanceAmount(uint256 roundNum_, address tokenAddress_) view returns(uint256)
func (_IPrizesWallet *IPrizesWalletSession) GetDonatedTokenBalanceAmount(roundNum_ *big.Int, tokenAddress_ common.Address) (*big.Int, error) {
	return _IPrizesWallet.Contract.GetDonatedTokenBalanceAmount(&_IPrizesWallet.CallOpts, roundNum_, tokenAddress_)
}

// GetDonatedTokenBalanceAmount is a free data retrieval call binding the contract method 0xe86a49d7.
//
// Solidity: function getDonatedTokenBalanceAmount(uint256 roundNum_, address tokenAddress_) view returns(uint256)
func (_IPrizesWallet *IPrizesWalletCallerSession) GetDonatedTokenBalanceAmount(roundNum_ *big.Int, tokenAddress_ common.Address) (*big.Int, error) {
	return _IPrizesWallet.Contract.GetDonatedTokenBalanceAmount(&_IPrizesWallet.CallOpts, roundNum_, tokenAddress_)
}

// GetEthBalanceInfo is a free data retrieval call binding the contract method 0x1a18c889.
//
// Solidity: function getEthBalanceInfo() view returns((uint256,uint256))
func (_IPrizesWallet *IPrizesWalletCaller) GetEthBalanceInfo(opts *bind.CallOpts) (IPrizesWalletEthBalanceInfo, error) {
	var out []interface{}
	err := _IPrizesWallet.contract.Call(opts, &out, "getEthBalanceInfo")

	if err != nil {
		return *new(IPrizesWalletEthBalanceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IPrizesWalletEthBalanceInfo)).(*IPrizesWalletEthBalanceInfo)

	return out0, err

}

// GetEthBalanceInfo is a free data retrieval call binding the contract method 0x1a18c889.
//
// Solidity: function getEthBalanceInfo() view returns((uint256,uint256))
func (_IPrizesWallet *IPrizesWalletSession) GetEthBalanceInfo() (IPrizesWalletEthBalanceInfo, error) {
	return _IPrizesWallet.Contract.GetEthBalanceInfo(&_IPrizesWallet.CallOpts)
}

// GetEthBalanceInfo is a free data retrieval call binding the contract method 0x1a18c889.
//
// Solidity: function getEthBalanceInfo() view returns((uint256,uint256))
func (_IPrizesWallet *IPrizesWalletCallerSession) GetEthBalanceInfo() (IPrizesWalletEthBalanceInfo, error) {
	return _IPrizesWallet.Contract.GetEthBalanceInfo(&_IPrizesWallet.CallOpts)
}

// GetEthBalanceInfo0 is a free data retrieval call binding the contract method 0xa089e0be.
//
// Solidity: function getEthBalanceInfo(address prizeWinnerAddress_) view returns((uint256,uint256))
func (_IPrizesWallet *IPrizesWalletCaller) GetEthBalanceInfo0(opts *bind.CallOpts, prizeWinnerAddress_ common.Address) (IPrizesWalletEthBalanceInfo, error) {
	var out []interface{}
	err := _IPrizesWallet.contract.Call(opts, &out, "getEthBalanceInfo0", prizeWinnerAddress_)

	if err != nil {
		return *new(IPrizesWalletEthBalanceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IPrizesWalletEthBalanceInfo)).(*IPrizesWalletEthBalanceInfo)

	return out0, err

}

// GetEthBalanceInfo0 is a free data retrieval call binding the contract method 0xa089e0be.
//
// Solidity: function getEthBalanceInfo(address prizeWinnerAddress_) view returns((uint256,uint256))
func (_IPrizesWallet *IPrizesWalletSession) GetEthBalanceInfo0(prizeWinnerAddress_ common.Address) (IPrizesWalletEthBalanceInfo, error) {
	return _IPrizesWallet.Contract.GetEthBalanceInfo0(&_IPrizesWallet.CallOpts, prizeWinnerAddress_)
}

// GetEthBalanceInfo0 is a free data retrieval call binding the contract method 0xa089e0be.
//
// Solidity: function getEthBalanceInfo(address prizeWinnerAddress_) view returns((uint256,uint256))
func (_IPrizesWallet *IPrizesWalletCallerSession) GetEthBalanceInfo0(prizeWinnerAddress_ common.Address) (IPrizesWalletEthBalanceInfo, error) {
	return _IPrizesWallet.Contract.GetEthBalanceInfo0(&_IPrizesWallet.CallOpts, prizeWinnerAddress_)
}

// ClaimDonatedNft is a paid mutator transaction binding the contract method 0x94d907fc.
//
// Solidity: function claimDonatedNft(uint256 index_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) ClaimDonatedNft(opts *bind.TransactOpts, index_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "claimDonatedNft", index_)
}

// ClaimDonatedNft is a paid mutator transaction binding the contract method 0x94d907fc.
//
// Solidity: function claimDonatedNft(uint256 index_) returns()
func (_IPrizesWallet *IPrizesWalletSession) ClaimDonatedNft(index_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.ClaimDonatedNft(&_IPrizesWallet.TransactOpts, index_)
}

// ClaimDonatedNft is a paid mutator transaction binding the contract method 0x94d907fc.
//
// Solidity: function claimDonatedNft(uint256 index_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) ClaimDonatedNft(index_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.ClaimDonatedNft(&_IPrizesWallet.TransactOpts, index_)
}

// ClaimDonatedToken is a paid mutator transaction binding the contract method 0x9523b237.
//
// Solidity: function claimDonatedToken(uint256 roundNum_, address tokenAddress_, uint256 amount_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) ClaimDonatedToken(opts *bind.TransactOpts, roundNum_ *big.Int, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "claimDonatedToken", roundNum_, tokenAddress_, amount_)
}

// ClaimDonatedToken is a paid mutator transaction binding the contract method 0x9523b237.
//
// Solidity: function claimDonatedToken(uint256 roundNum_, address tokenAddress_, uint256 amount_) returns()
func (_IPrizesWallet *IPrizesWalletSession) ClaimDonatedToken(roundNum_ *big.Int, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.ClaimDonatedToken(&_IPrizesWallet.TransactOpts, roundNum_, tokenAddress_, amount_)
}

// ClaimDonatedToken is a paid mutator transaction binding the contract method 0x9523b237.
//
// Solidity: function claimDonatedToken(uint256 roundNum_, address tokenAddress_, uint256 amount_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) ClaimDonatedToken(roundNum_ *big.Int, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.ClaimDonatedToken(&_IPrizesWallet.TransactOpts, roundNum_, tokenAddress_, amount_)
}

// ClaimManyDonatedNfts is a paid mutator transaction binding the contract method 0xe4a6c2a4.
//
// Solidity: function claimManyDonatedNfts(uint256[] indexes_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) ClaimManyDonatedNfts(opts *bind.TransactOpts, indexes_ []*big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "claimManyDonatedNfts", indexes_)
}

// ClaimManyDonatedNfts is a paid mutator transaction binding the contract method 0xe4a6c2a4.
//
// Solidity: function claimManyDonatedNfts(uint256[] indexes_) returns()
func (_IPrizesWallet *IPrizesWalletSession) ClaimManyDonatedNfts(indexes_ []*big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.ClaimManyDonatedNfts(&_IPrizesWallet.TransactOpts, indexes_)
}

// ClaimManyDonatedNfts is a paid mutator transaction binding the contract method 0xe4a6c2a4.
//
// Solidity: function claimManyDonatedNfts(uint256[] indexes_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) ClaimManyDonatedNfts(indexes_ []*big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.ClaimManyDonatedNfts(&_IPrizesWallet.TransactOpts, indexes_)
}

// ClaimManyDonatedTokens is a paid mutator transaction binding the contract method 0x5e00aaba.
//
// Solidity: function claimManyDonatedTokens((uint256,address,uint256)[] donatedTokensToClaim_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) ClaimManyDonatedTokens(opts *bind.TransactOpts, donatedTokensToClaim_ []IPrizesWalletDonatedTokenToClaim) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "claimManyDonatedTokens", donatedTokensToClaim_)
}

// ClaimManyDonatedTokens is a paid mutator transaction binding the contract method 0x5e00aaba.
//
// Solidity: function claimManyDonatedTokens((uint256,address,uint256)[] donatedTokensToClaim_) returns()
func (_IPrizesWallet *IPrizesWalletSession) ClaimManyDonatedTokens(donatedTokensToClaim_ []IPrizesWalletDonatedTokenToClaim) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.ClaimManyDonatedTokens(&_IPrizesWallet.TransactOpts, donatedTokensToClaim_)
}

// ClaimManyDonatedTokens is a paid mutator transaction binding the contract method 0x5e00aaba.
//
// Solidity: function claimManyDonatedTokens((uint256,address,uint256)[] donatedTokensToClaim_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) ClaimManyDonatedTokens(donatedTokensToClaim_ []IPrizesWalletDonatedTokenToClaim) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.ClaimManyDonatedTokens(&_IPrizesWallet.TransactOpts, donatedTokensToClaim_)
}

// DepositEth is a paid mutator transaction binding the contract method 0xf11b35fa.
//
// Solidity: function depositEth(uint256 roundNum_, uint256 prizeWinnerIndex_, address prizeWinnerAddress_) payable returns()
func (_IPrizesWallet *IPrizesWalletTransactor) DepositEth(opts *bind.TransactOpts, roundNum_ *big.Int, prizeWinnerIndex_ *big.Int, prizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "depositEth", roundNum_, prizeWinnerIndex_, prizeWinnerAddress_)
}

// DepositEth is a paid mutator transaction binding the contract method 0xf11b35fa.
//
// Solidity: function depositEth(uint256 roundNum_, uint256 prizeWinnerIndex_, address prizeWinnerAddress_) payable returns()
func (_IPrizesWallet *IPrizesWalletSession) DepositEth(roundNum_ *big.Int, prizeWinnerIndex_ *big.Int, prizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.DepositEth(&_IPrizesWallet.TransactOpts, roundNum_, prizeWinnerIndex_, prizeWinnerAddress_)
}

// DepositEth is a paid mutator transaction binding the contract method 0xf11b35fa.
//
// Solidity: function depositEth(uint256 roundNum_, uint256 prizeWinnerIndex_, address prizeWinnerAddress_) payable returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) DepositEth(roundNum_ *big.Int, prizeWinnerIndex_ *big.Int, prizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.DepositEth(&_IPrizesWallet.TransactOpts, roundNum_, prizeWinnerIndex_, prizeWinnerAddress_)
}

// DonateNft is a paid mutator transaction binding the contract method 0xfe673fd3.
//
// Solidity: function donateNft(uint256 roundNum_, address donorAddress_, address nftAddress_, uint256 nftId_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) DonateNft(opts *bind.TransactOpts, roundNum_ *big.Int, donorAddress_ common.Address, nftAddress_ common.Address, nftId_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "donateNft", roundNum_, donorAddress_, nftAddress_, nftId_)
}

// DonateNft is a paid mutator transaction binding the contract method 0xfe673fd3.
//
// Solidity: function donateNft(uint256 roundNum_, address donorAddress_, address nftAddress_, uint256 nftId_) returns()
func (_IPrizesWallet *IPrizesWalletSession) DonateNft(roundNum_ *big.Int, donorAddress_ common.Address, nftAddress_ common.Address, nftId_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.DonateNft(&_IPrizesWallet.TransactOpts, roundNum_, donorAddress_, nftAddress_, nftId_)
}

// DonateNft is a paid mutator transaction binding the contract method 0xfe673fd3.
//
// Solidity: function donateNft(uint256 roundNum_, address donorAddress_, address nftAddress_, uint256 nftId_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) DonateNft(roundNum_ *big.Int, donorAddress_ common.Address, nftAddress_ common.Address, nftId_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.DonateNft(&_IPrizesWallet.TransactOpts, roundNum_, donorAddress_, nftAddress_, nftId_)
}

// DonateToken is a paid mutator transaction binding the contract method 0xe2051c7e.
//
// Solidity: function donateToken(uint256 roundNum_, address donorAddress_, address tokenAddress_, uint256 amount_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) DonateToken(opts *bind.TransactOpts, roundNum_ *big.Int, donorAddress_ common.Address, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "donateToken", roundNum_, donorAddress_, tokenAddress_, amount_)
}

// DonateToken is a paid mutator transaction binding the contract method 0xe2051c7e.
//
// Solidity: function donateToken(uint256 roundNum_, address donorAddress_, address tokenAddress_, uint256 amount_) returns()
func (_IPrizesWallet *IPrizesWalletSession) DonateToken(roundNum_ *big.Int, donorAddress_ common.Address, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.DonateToken(&_IPrizesWallet.TransactOpts, roundNum_, donorAddress_, tokenAddress_, amount_)
}

// DonateToken is a paid mutator transaction binding the contract method 0xe2051c7e.
//
// Solidity: function donateToken(uint256 roundNum_, address donorAddress_, address tokenAddress_, uint256 amount_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) DonateToken(roundNum_ *big.Int, donorAddress_ common.Address, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.DonateToken(&_IPrizesWallet.TransactOpts, roundNum_, donorAddress_, tokenAddress_, amount_)
}

// RegisterRoundEnd is a paid mutator transaction binding the contract method 0xcc5810d8.
//
// Solidity: function registerRoundEnd(uint256 roundNum_, address mainPrizeBeneficiaryAddress_) returns(uint256)
func (_IPrizesWallet *IPrizesWalletTransactor) RegisterRoundEnd(opts *bind.TransactOpts, roundNum_ *big.Int, mainPrizeBeneficiaryAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "registerRoundEnd", roundNum_, mainPrizeBeneficiaryAddress_)
}

// RegisterRoundEnd is a paid mutator transaction binding the contract method 0xcc5810d8.
//
// Solidity: function registerRoundEnd(uint256 roundNum_, address mainPrizeBeneficiaryAddress_) returns(uint256)
func (_IPrizesWallet *IPrizesWalletSession) RegisterRoundEnd(roundNum_ *big.Int, mainPrizeBeneficiaryAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.RegisterRoundEnd(&_IPrizesWallet.TransactOpts, roundNum_, mainPrizeBeneficiaryAddress_)
}

// RegisterRoundEnd is a paid mutator transaction binding the contract method 0xcc5810d8.
//
// Solidity: function registerRoundEnd(uint256 roundNum_, address mainPrizeBeneficiaryAddress_) returns(uint256)
func (_IPrizesWallet *IPrizesWalletTransactorSession) RegisterRoundEnd(roundNum_ *big.Int, mainPrizeBeneficiaryAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.RegisterRoundEnd(&_IPrizesWallet.TransactOpts, roundNum_, mainPrizeBeneficiaryAddress_)
}

// RegisterRoundEndAndDepositEthMany is a paid mutator transaction binding the contract method 0x87565d14.
//
// Solidity: function registerRoundEndAndDepositEthMany(uint256 roundNum_, address mainPrizeBeneficiaryAddress_, (address,uint256)[] ethDeposits_) payable returns(uint256)
func (_IPrizesWallet *IPrizesWalletTransactor) RegisterRoundEndAndDepositEthMany(opts *bind.TransactOpts, roundNum_ *big.Int, mainPrizeBeneficiaryAddress_ common.Address, ethDeposits_ []IPrizesWalletEthDeposit) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "registerRoundEndAndDepositEthMany", roundNum_, mainPrizeBeneficiaryAddress_, ethDeposits_)
}

// RegisterRoundEndAndDepositEthMany is a paid mutator transaction binding the contract method 0x87565d14.
//
// Solidity: function registerRoundEndAndDepositEthMany(uint256 roundNum_, address mainPrizeBeneficiaryAddress_, (address,uint256)[] ethDeposits_) payable returns(uint256)
func (_IPrizesWallet *IPrizesWalletSession) RegisterRoundEndAndDepositEthMany(roundNum_ *big.Int, mainPrizeBeneficiaryAddress_ common.Address, ethDeposits_ []IPrizesWalletEthDeposit) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.RegisterRoundEndAndDepositEthMany(&_IPrizesWallet.TransactOpts, roundNum_, mainPrizeBeneficiaryAddress_, ethDeposits_)
}

// RegisterRoundEndAndDepositEthMany is a paid mutator transaction binding the contract method 0x87565d14.
//
// Solidity: function registerRoundEndAndDepositEthMany(uint256 roundNum_, address mainPrizeBeneficiaryAddress_, (address,uint256)[] ethDeposits_) payable returns(uint256)
func (_IPrizesWallet *IPrizesWalletTransactorSession) RegisterRoundEndAndDepositEthMany(roundNum_ *big.Int, mainPrizeBeneficiaryAddress_ common.Address, ethDeposits_ []IPrizesWalletEthDeposit) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.RegisterRoundEndAndDepositEthMany(&_IPrizesWallet.TransactOpts, roundNum_, mainPrizeBeneficiaryAddress_, ethDeposits_)
}

// SetTimeoutDurationToWithdrawPrizes is a paid mutator transaction binding the contract method 0x9e2842a8.
//
// Solidity: function setTimeoutDurationToWithdrawPrizes(uint256 newValue_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) SetTimeoutDurationToWithdrawPrizes(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "setTimeoutDurationToWithdrawPrizes", newValue_)
}

// SetTimeoutDurationToWithdrawPrizes is a paid mutator transaction binding the contract method 0x9e2842a8.
//
// Solidity: function setTimeoutDurationToWithdrawPrizes(uint256 newValue_) returns()
func (_IPrizesWallet *IPrizesWalletSession) SetTimeoutDurationToWithdrawPrizes(newValue_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.SetTimeoutDurationToWithdrawPrizes(&_IPrizesWallet.TransactOpts, newValue_)
}

// SetTimeoutDurationToWithdrawPrizes is a paid mutator transaction binding the contract method 0x9e2842a8.
//
// Solidity: function setTimeoutDurationToWithdrawPrizes(uint256 newValue_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) SetTimeoutDurationToWithdrawPrizes(newValue_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.SetTimeoutDurationToWithdrawPrizes(&_IPrizesWallet.TransactOpts, newValue_)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x25e16063.
//
// Solidity: function withdrawEth(address prizeWinnerAddress_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) WithdrawEth(opts *bind.TransactOpts, prizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "withdrawEth", prizeWinnerAddress_)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x25e16063.
//
// Solidity: function withdrawEth(address prizeWinnerAddress_) returns()
func (_IPrizesWallet *IPrizesWalletSession) WithdrawEth(prizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.WithdrawEth(&_IPrizesWallet.TransactOpts, prizeWinnerAddress_)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x25e16063.
//
// Solidity: function withdrawEth(address prizeWinnerAddress_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) WithdrawEth(prizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.WithdrawEth(&_IPrizesWallet.TransactOpts, prizeWinnerAddress_)
}

// WithdrawEth0 is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_IPrizesWallet *IPrizesWalletTransactor) WithdrawEth0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "withdrawEth0")
}

// WithdrawEth0 is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_IPrizesWallet *IPrizesWalletSession) WithdrawEth0() (*types.Transaction, error) {
	return _IPrizesWallet.Contract.WithdrawEth0(&_IPrizesWallet.TransactOpts)
}

// WithdrawEth0 is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) WithdrawEth0() (*types.Transaction, error) {
	return _IPrizesWallet.Contract.WithdrawEth0(&_IPrizesWallet.TransactOpts)
}

// WithdrawEverything is a paid mutator transaction binding the contract method 0xb4e6f803.
//
// Solidity: function withdrawEverything(bool withdrawEth_, (uint256,address,uint256)[] donatedTokensToClaim_, uint256[] donatedNftIndexes_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) WithdrawEverything(opts *bind.TransactOpts, withdrawEth_ bool, donatedTokensToClaim_ []IPrizesWalletDonatedTokenToClaim, donatedNftIndexes_ []*big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "withdrawEverything", withdrawEth_, donatedTokensToClaim_, donatedNftIndexes_)
}

// WithdrawEverything is a paid mutator transaction binding the contract method 0xb4e6f803.
//
// Solidity: function withdrawEverything(bool withdrawEth_, (uint256,address,uint256)[] donatedTokensToClaim_, uint256[] donatedNftIndexes_) returns()
func (_IPrizesWallet *IPrizesWalletSession) WithdrawEverything(withdrawEth_ bool, donatedTokensToClaim_ []IPrizesWalletDonatedTokenToClaim, donatedNftIndexes_ []*big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.WithdrawEverything(&_IPrizesWallet.TransactOpts, withdrawEth_, donatedTokensToClaim_, donatedNftIndexes_)
}

// WithdrawEverything is a paid mutator transaction binding the contract method 0xb4e6f803.
//
// Solidity: function withdrawEverything(bool withdrawEth_, (uint256,address,uint256)[] donatedTokensToClaim_, uint256[] donatedNftIndexes_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) WithdrawEverything(withdrawEth_ bool, donatedTokensToClaim_ []IPrizesWalletDonatedTokenToClaim, donatedNftIndexes_ []*big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.WithdrawEverything(&_IPrizesWallet.TransactOpts, withdrawEth_, donatedTokensToClaim_, donatedNftIndexes_)
}

// IPrizesWalletDonatedNftClaimedIterator is returned from FilterDonatedNftClaimed and is used to iterate over the raw logs and unpacked data for DonatedNftClaimed events raised by the IPrizesWallet contract.
type IPrizesWalletDonatedNftClaimedIterator struct {
	Event *IPrizesWalletDonatedNftClaimed // Event containing the contract specifics and raw log

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
func (it *IPrizesWalletDonatedNftClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPrizesWalletDonatedNftClaimed)
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
		it.Event = new(IPrizesWalletDonatedNftClaimed)
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
func (it *IPrizesWalletDonatedNftClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPrizesWalletDonatedNftClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPrizesWalletDonatedNftClaimed represents a DonatedNftClaimed event raised by the IPrizesWallet contract.
type IPrizesWalletDonatedNftClaimed struct {
	RoundNum           *big.Int
	BeneficiaryAddress common.Address
	NftAddress         common.Address
	NftId              *big.Int
	Index              *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDonatedNftClaimed is a free log retrieval operation binding the contract event 0x03c2b6e01c9ca39e4073132f279b61b987d41a82b08cea3dd4a2fabc37067ac3.
//
// Solidity: event DonatedNftClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, address indexed nftAddress, uint256 nftId, uint256 index)
func (_IPrizesWallet *IPrizesWalletFilterer) FilterDonatedNftClaimed(opts *bind.FilterOpts, roundNum []*big.Int, beneficiaryAddress []common.Address, nftAddress []common.Address) (*IPrizesWalletDonatedNftClaimedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.FilterLogs(opts, "DonatedNftClaimed", roundNumRule, beneficiaryAddressRule, nftAddressRule)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletDonatedNftClaimedIterator{contract: _IPrizesWallet.contract, event: "DonatedNftClaimed", logs: logs, sub: sub}, nil
}

// WatchDonatedNftClaimed is a free log subscription operation binding the contract event 0x03c2b6e01c9ca39e4073132f279b61b987d41a82b08cea3dd4a2fabc37067ac3.
//
// Solidity: event DonatedNftClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, address indexed nftAddress, uint256 nftId, uint256 index)
func (_IPrizesWallet *IPrizesWalletFilterer) WatchDonatedNftClaimed(opts *bind.WatchOpts, sink chan<- *IPrizesWalletDonatedNftClaimed, roundNum []*big.Int, beneficiaryAddress []common.Address, nftAddress []common.Address) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.WatchLogs(opts, "DonatedNftClaimed", roundNumRule, beneficiaryAddressRule, nftAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPrizesWalletDonatedNftClaimed)
				if err := _IPrizesWallet.contract.UnpackLog(event, "DonatedNftClaimed", log); err != nil {
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

// ParseDonatedNftClaimed is a log parse operation binding the contract event 0x03c2b6e01c9ca39e4073132f279b61b987d41a82b08cea3dd4a2fabc37067ac3.
//
// Solidity: event DonatedNftClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, address indexed nftAddress, uint256 nftId, uint256 index)
func (_IPrizesWallet *IPrizesWalletFilterer) ParseDonatedNftClaimed(log types.Log) (*IPrizesWalletDonatedNftClaimed, error) {
	event := new(IPrizesWalletDonatedNftClaimed)
	if err := _IPrizesWallet.contract.UnpackLog(event, "DonatedNftClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPrizesWalletDonatedTokenClaimedIterator is returned from FilterDonatedTokenClaimed and is used to iterate over the raw logs and unpacked data for DonatedTokenClaimed events raised by the IPrizesWallet contract.
type IPrizesWalletDonatedTokenClaimedIterator struct {
	Event *IPrizesWalletDonatedTokenClaimed // Event containing the contract specifics and raw log

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
func (it *IPrizesWalletDonatedTokenClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPrizesWalletDonatedTokenClaimed)
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
		it.Event = new(IPrizesWalletDonatedTokenClaimed)
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
func (it *IPrizesWalletDonatedTokenClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPrizesWalletDonatedTokenClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPrizesWalletDonatedTokenClaimed represents a DonatedTokenClaimed event raised by the IPrizesWallet contract.
type IPrizesWalletDonatedTokenClaimed struct {
	RoundNum           *big.Int
	BeneficiaryAddress common.Address
	TokenAddress       common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDonatedTokenClaimed is a free log retrieval operation binding the contract event 0xaf1adae2e1e983ca738335dc2e37194114142793f394ac934a45ea632a8a5bd0.
//
// Solidity: event DonatedTokenClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, address indexed tokenAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) FilterDonatedTokenClaimed(opts *bind.FilterOpts, roundNum []*big.Int, beneficiaryAddress []common.Address, tokenAddress []common.Address) (*IPrizesWalletDonatedTokenClaimedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.FilterLogs(opts, "DonatedTokenClaimed", roundNumRule, beneficiaryAddressRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletDonatedTokenClaimedIterator{contract: _IPrizesWallet.contract, event: "DonatedTokenClaimed", logs: logs, sub: sub}, nil
}

// WatchDonatedTokenClaimed is a free log subscription operation binding the contract event 0xaf1adae2e1e983ca738335dc2e37194114142793f394ac934a45ea632a8a5bd0.
//
// Solidity: event DonatedTokenClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, address indexed tokenAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) WatchDonatedTokenClaimed(opts *bind.WatchOpts, sink chan<- *IPrizesWalletDonatedTokenClaimed, roundNum []*big.Int, beneficiaryAddress []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.WatchLogs(opts, "DonatedTokenClaimed", roundNumRule, beneficiaryAddressRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPrizesWalletDonatedTokenClaimed)
				if err := _IPrizesWallet.contract.UnpackLog(event, "DonatedTokenClaimed", log); err != nil {
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

// ParseDonatedTokenClaimed is a log parse operation binding the contract event 0xaf1adae2e1e983ca738335dc2e37194114142793f394ac934a45ea632a8a5bd0.
//
// Solidity: event DonatedTokenClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, address indexed tokenAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) ParseDonatedTokenClaimed(log types.Log) (*IPrizesWalletDonatedTokenClaimed, error) {
	event := new(IPrizesWalletDonatedTokenClaimed)
	if err := _IPrizesWallet.contract.UnpackLog(event, "DonatedTokenClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPrizesWalletEthReceivedIterator is returned from FilterEthReceived and is used to iterate over the raw logs and unpacked data for EthReceived events raised by the IPrizesWallet contract.
type IPrizesWalletEthReceivedIterator struct {
	Event *IPrizesWalletEthReceived // Event containing the contract specifics and raw log

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
func (it *IPrizesWalletEthReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPrizesWalletEthReceived)
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
		it.Event = new(IPrizesWalletEthReceived)
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
func (it *IPrizesWalletEthReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPrizesWalletEthReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPrizesWalletEthReceived represents a EthReceived event raised by the IPrizesWallet contract.
type IPrizesWalletEthReceived struct {
	RoundNum           *big.Int
	PrizeWinnerIndex   *big.Int
	PrizeWinnerAddress common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterEthReceived is a free log retrieval operation binding the contract event 0x8e369548f9ef214c7c1823c098a23763321fd761ac5cf78958e1db1b3648e7a2.
//
// Solidity: event EthReceived(uint256 indexed roundNum, uint256 prizeWinnerIndex, address indexed prizeWinnerAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) FilterEthReceived(opts *bind.FilterOpts, roundNum []*big.Int, prizeWinnerAddress []common.Address) (*IPrizesWalletEthReceivedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	var prizeWinnerAddressRule []interface{}
	for _, prizeWinnerAddressItem := range prizeWinnerAddress {
		prizeWinnerAddressRule = append(prizeWinnerAddressRule, prizeWinnerAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.FilterLogs(opts, "EthReceived", roundNumRule, prizeWinnerAddressRule)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletEthReceivedIterator{contract: _IPrizesWallet.contract, event: "EthReceived", logs: logs, sub: sub}, nil
}

// WatchEthReceived is a free log subscription operation binding the contract event 0x8e369548f9ef214c7c1823c098a23763321fd761ac5cf78958e1db1b3648e7a2.
//
// Solidity: event EthReceived(uint256 indexed roundNum, uint256 prizeWinnerIndex, address indexed prizeWinnerAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) WatchEthReceived(opts *bind.WatchOpts, sink chan<- *IPrizesWalletEthReceived, roundNum []*big.Int, prizeWinnerAddress []common.Address) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	var prizeWinnerAddressRule []interface{}
	for _, prizeWinnerAddressItem := range prizeWinnerAddress {
		prizeWinnerAddressRule = append(prizeWinnerAddressRule, prizeWinnerAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.WatchLogs(opts, "EthReceived", roundNumRule, prizeWinnerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPrizesWalletEthReceived)
				if err := _IPrizesWallet.contract.UnpackLog(event, "EthReceived", log); err != nil {
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

// ParseEthReceived is a log parse operation binding the contract event 0x8e369548f9ef214c7c1823c098a23763321fd761ac5cf78958e1db1b3648e7a2.
//
// Solidity: event EthReceived(uint256 indexed roundNum, uint256 prizeWinnerIndex, address indexed prizeWinnerAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) ParseEthReceived(log types.Log) (*IPrizesWalletEthReceived, error) {
	event := new(IPrizesWalletEthReceived)
	if err := _IPrizesWallet.contract.UnpackLog(event, "EthReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPrizesWalletEthWithdrawnIterator is returned from FilterEthWithdrawn and is used to iterate over the raw logs and unpacked data for EthWithdrawn events raised by the IPrizesWallet contract.
type IPrizesWalletEthWithdrawnIterator struct {
	Event *IPrizesWalletEthWithdrawn // Event containing the contract specifics and raw log

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
func (it *IPrizesWalletEthWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPrizesWalletEthWithdrawn)
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
		it.Event = new(IPrizesWalletEthWithdrawn)
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
func (it *IPrizesWalletEthWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPrizesWalletEthWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPrizesWalletEthWithdrawn represents a EthWithdrawn event raised by the IPrizesWallet contract.
type IPrizesWalletEthWithdrawn struct {
	PrizeWinnerAddress common.Address
	BeneficiaryAddress common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterEthWithdrawn is a free log retrieval operation binding the contract event 0x4f43b861ba36494acfe938f3815fba7fac6981bdc611b6ccdc14c08f59292383.
//
// Solidity: event EthWithdrawn(address indexed prizeWinnerAddress, address indexed beneficiaryAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) FilterEthWithdrawn(opts *bind.FilterOpts, prizeWinnerAddress []common.Address, beneficiaryAddress []common.Address) (*IPrizesWalletEthWithdrawnIterator, error) {

	var prizeWinnerAddressRule []interface{}
	for _, prizeWinnerAddressItem := range prizeWinnerAddress {
		prizeWinnerAddressRule = append(prizeWinnerAddressRule, prizeWinnerAddressItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.FilterLogs(opts, "EthWithdrawn", prizeWinnerAddressRule, beneficiaryAddressRule)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletEthWithdrawnIterator{contract: _IPrizesWallet.contract, event: "EthWithdrawn", logs: logs, sub: sub}, nil
}

// WatchEthWithdrawn is a free log subscription operation binding the contract event 0x4f43b861ba36494acfe938f3815fba7fac6981bdc611b6ccdc14c08f59292383.
//
// Solidity: event EthWithdrawn(address indexed prizeWinnerAddress, address indexed beneficiaryAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) WatchEthWithdrawn(opts *bind.WatchOpts, sink chan<- *IPrizesWalletEthWithdrawn, prizeWinnerAddress []common.Address, beneficiaryAddress []common.Address) (event.Subscription, error) {

	var prizeWinnerAddressRule []interface{}
	for _, prizeWinnerAddressItem := range prizeWinnerAddress {
		prizeWinnerAddressRule = append(prizeWinnerAddressRule, prizeWinnerAddressItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.WatchLogs(opts, "EthWithdrawn", prizeWinnerAddressRule, beneficiaryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPrizesWalletEthWithdrawn)
				if err := _IPrizesWallet.contract.UnpackLog(event, "EthWithdrawn", log); err != nil {
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

// ParseEthWithdrawn is a log parse operation binding the contract event 0x4f43b861ba36494acfe938f3815fba7fac6981bdc611b6ccdc14c08f59292383.
//
// Solidity: event EthWithdrawn(address indexed prizeWinnerAddress, address indexed beneficiaryAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) ParseEthWithdrawn(log types.Log) (*IPrizesWalletEthWithdrawn, error) {
	event := new(IPrizesWalletEthWithdrawn)
	if err := _IPrizesWallet.contract.UnpackLog(event, "EthWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPrizesWalletNftDonatedIterator is returned from FilterNftDonated and is used to iterate over the raw logs and unpacked data for NftDonated events raised by the IPrizesWallet contract.
type IPrizesWalletNftDonatedIterator struct {
	Event *IPrizesWalletNftDonated // Event containing the contract specifics and raw log

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
func (it *IPrizesWalletNftDonatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPrizesWalletNftDonated)
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
		it.Event = new(IPrizesWalletNftDonated)
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
func (it *IPrizesWalletNftDonatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPrizesWalletNftDonatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPrizesWalletNftDonated represents a NftDonated event raised by the IPrizesWallet contract.
type IPrizesWalletNftDonated struct {
	RoundNum     *big.Int
	DonorAddress common.Address
	NftAddress   common.Address
	NftId        *big.Int
	Index        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNftDonated is a free log retrieval operation binding the contract event 0xb12e72bab0c2a8fe98175a3c02792645ebdf65d37cf3983517e6d1c5ab842e23.
//
// Solidity: event NftDonated(uint256 indexed roundNum, address indexed donorAddress, address indexed nftAddress, uint256 nftId, uint256 index)
func (_IPrizesWallet *IPrizesWalletFilterer) FilterNftDonated(opts *bind.FilterOpts, roundNum []*big.Int, donorAddress []common.Address, nftAddress []common.Address) (*IPrizesWalletNftDonatedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.FilterLogs(opts, "NftDonated", roundNumRule, donorAddressRule, nftAddressRule)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletNftDonatedIterator{contract: _IPrizesWallet.contract, event: "NftDonated", logs: logs, sub: sub}, nil
}

// WatchNftDonated is a free log subscription operation binding the contract event 0xb12e72bab0c2a8fe98175a3c02792645ebdf65d37cf3983517e6d1c5ab842e23.
//
// Solidity: event NftDonated(uint256 indexed roundNum, address indexed donorAddress, address indexed nftAddress, uint256 nftId, uint256 index)
func (_IPrizesWallet *IPrizesWalletFilterer) WatchNftDonated(opts *bind.WatchOpts, sink chan<- *IPrizesWalletNftDonated, roundNum []*big.Int, donorAddress []common.Address, nftAddress []common.Address) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.WatchLogs(opts, "NftDonated", roundNumRule, donorAddressRule, nftAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPrizesWalletNftDonated)
				if err := _IPrizesWallet.contract.UnpackLog(event, "NftDonated", log); err != nil {
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

// ParseNftDonated is a log parse operation binding the contract event 0xb12e72bab0c2a8fe98175a3c02792645ebdf65d37cf3983517e6d1c5ab842e23.
//
// Solidity: event NftDonated(uint256 indexed roundNum, address indexed donorAddress, address indexed nftAddress, uint256 nftId, uint256 index)
func (_IPrizesWallet *IPrizesWalletFilterer) ParseNftDonated(log types.Log) (*IPrizesWalletNftDonated, error) {
	event := new(IPrizesWalletNftDonated)
	if err := _IPrizesWallet.contract.UnpackLog(event, "NftDonated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPrizesWalletTimeoutDurationToWithdrawPrizesChangedIterator is returned from FilterTimeoutDurationToWithdrawPrizesChanged and is used to iterate over the raw logs and unpacked data for TimeoutDurationToWithdrawPrizesChanged events raised by the IPrizesWallet contract.
type IPrizesWalletTimeoutDurationToWithdrawPrizesChangedIterator struct {
	Event *IPrizesWalletTimeoutDurationToWithdrawPrizesChanged // Event containing the contract specifics and raw log

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
func (it *IPrizesWalletTimeoutDurationToWithdrawPrizesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPrizesWalletTimeoutDurationToWithdrawPrizesChanged)
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
		it.Event = new(IPrizesWalletTimeoutDurationToWithdrawPrizesChanged)
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
func (it *IPrizesWalletTimeoutDurationToWithdrawPrizesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPrizesWalletTimeoutDurationToWithdrawPrizesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPrizesWalletTimeoutDurationToWithdrawPrizesChanged represents a TimeoutDurationToWithdrawPrizesChanged event raised by the IPrizesWallet contract.
type IPrizesWalletTimeoutDurationToWithdrawPrizesChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTimeoutDurationToWithdrawPrizesChanged is a free log retrieval operation binding the contract event 0x8717bb199c6bc4a5dadb21547205f9ef8ec037dda246a5526d6a6471306ea52e.
//
// Solidity: event TimeoutDurationToWithdrawPrizesChanged(uint256 newValue)
func (_IPrizesWallet *IPrizesWalletFilterer) FilterTimeoutDurationToWithdrawPrizesChanged(opts *bind.FilterOpts) (*IPrizesWalletTimeoutDurationToWithdrawPrizesChangedIterator, error) {

	logs, sub, err := _IPrizesWallet.contract.FilterLogs(opts, "TimeoutDurationToWithdrawPrizesChanged")
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletTimeoutDurationToWithdrawPrizesChangedIterator{contract: _IPrizesWallet.contract, event: "TimeoutDurationToWithdrawPrizesChanged", logs: logs, sub: sub}, nil
}

// WatchTimeoutDurationToWithdrawPrizesChanged is a free log subscription operation binding the contract event 0x8717bb199c6bc4a5dadb21547205f9ef8ec037dda246a5526d6a6471306ea52e.
//
// Solidity: event TimeoutDurationToWithdrawPrizesChanged(uint256 newValue)
func (_IPrizesWallet *IPrizesWalletFilterer) WatchTimeoutDurationToWithdrawPrizesChanged(opts *bind.WatchOpts, sink chan<- *IPrizesWalletTimeoutDurationToWithdrawPrizesChanged) (event.Subscription, error) {

	logs, sub, err := _IPrizesWallet.contract.WatchLogs(opts, "TimeoutDurationToWithdrawPrizesChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPrizesWalletTimeoutDurationToWithdrawPrizesChanged)
				if err := _IPrizesWallet.contract.UnpackLog(event, "TimeoutDurationToWithdrawPrizesChanged", log); err != nil {
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

// ParseTimeoutDurationToWithdrawPrizesChanged is a log parse operation binding the contract event 0x8717bb199c6bc4a5dadb21547205f9ef8ec037dda246a5526d6a6471306ea52e.
//
// Solidity: event TimeoutDurationToWithdrawPrizesChanged(uint256 newValue)
func (_IPrizesWallet *IPrizesWalletFilterer) ParseTimeoutDurationToWithdrawPrizesChanged(log types.Log) (*IPrizesWalletTimeoutDurationToWithdrawPrizesChanged, error) {
	event := new(IPrizesWalletTimeoutDurationToWithdrawPrizesChanged)
	if err := _IPrizesWallet.contract.UnpackLog(event, "TimeoutDurationToWithdrawPrizesChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPrizesWalletTokenDonatedIterator is returned from FilterTokenDonated and is used to iterate over the raw logs and unpacked data for TokenDonated events raised by the IPrizesWallet contract.
type IPrizesWalletTokenDonatedIterator struct {
	Event *IPrizesWalletTokenDonated // Event containing the contract specifics and raw log

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
func (it *IPrizesWalletTokenDonatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPrizesWalletTokenDonated)
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
		it.Event = new(IPrizesWalletTokenDonated)
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
func (it *IPrizesWalletTokenDonatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPrizesWalletTokenDonatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPrizesWalletTokenDonated represents a TokenDonated event raised by the IPrizesWallet contract.
type IPrizesWalletTokenDonated struct {
	RoundNum     *big.Int
	DonorAddress common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenDonated is a free log retrieval operation binding the contract event 0x3f94f6171c96ab491d8c268f80da5e460b148878065711a0bb099bbc297615af.
//
// Solidity: event TokenDonated(uint256 indexed roundNum, address indexed donorAddress, address indexed tokenAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) FilterTokenDonated(opts *bind.FilterOpts, roundNum []*big.Int, donorAddress []common.Address, tokenAddress []common.Address) (*IPrizesWalletTokenDonatedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.FilterLogs(opts, "TokenDonated", roundNumRule, donorAddressRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletTokenDonatedIterator{contract: _IPrizesWallet.contract, event: "TokenDonated", logs: logs, sub: sub}, nil
}

// WatchTokenDonated is a free log subscription operation binding the contract event 0x3f94f6171c96ab491d8c268f80da5e460b148878065711a0bb099bbc297615af.
//
// Solidity: event TokenDonated(uint256 indexed roundNum, address indexed donorAddress, address indexed tokenAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) WatchTokenDonated(opts *bind.WatchOpts, sink chan<- *IPrizesWalletTokenDonated, roundNum []*big.Int, donorAddress []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.WatchLogs(opts, "TokenDonated", roundNumRule, donorAddressRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPrizesWalletTokenDonated)
				if err := _IPrizesWallet.contract.UnpackLog(event, "TokenDonated", log); err != nil {
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

// ParseTokenDonated is a log parse operation binding the contract event 0x3f94f6171c96ab491d8c268f80da5e460b148878065711a0bb099bbc297615af.
//
// Solidity: event TokenDonated(uint256 indexed roundNum, address indexed donorAddress, address indexed tokenAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) ParseTokenDonated(log types.Log) (*IPrizesWalletTokenDonated, error) {
	event := new(IPrizesWalletTokenDonated)
	if err := _IPrizesWallet.contract.UnpackLog(event, "TokenDonated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizesWalletMetaData contains all meta data concerning the PrizesWallet contract.
var PrizesWalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"game_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DonatedNftAlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"operationPermittedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimeStamp\",\"type\":\"uint256\"}],\"name\":\"DonatedNftClaimDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"operationPermittedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimeStamp\",\"type\":\"uint256\"}],\"name\":\"DonatedTokenClaimDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"prizeWinnerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"operationPermittedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimeStamp\",\"type\":\"uint256\"}],\"name\":\"EthWithdrawalDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundTransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"InvalidDonatedNftIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"callerAddress\",\"type\":\"address\"}],\"name\":\"UnauthorizedCaller\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DonatedNftClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DonatedTokenClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prizeWinnerIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prizeWinnerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prizeWinnerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"NftDonated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"TimeoutDurationToWithdrawPrizesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenDonated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"}],\"name\":\"claimDonatedNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"claimDonatedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"indexes_\",\"type\":\"uint256[]\"}],\"name\":\"claimManyDonatedNfts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPrizesWallet.DonatedTokenToClaim[]\",\"name\":\"donatedTokensToClaim_\",\"type\":\"tuple[]\"}],\"name\":\"claimManyDonatedTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prizeWinnerIndex_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"prizeWinnerAddress_\",\"type\":\"address\"}],\"name\":\"depositEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"donorAddress_\",\"type\":\"address\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"donateNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"donorAddress_\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"donateToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"donatedNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"donatedTokens\",\"outputs\":[{\"internalType\":\"contractDonatedTokenHolder\",\"name\":\"holder\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"game\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"}],\"name\":\"getDonatedTokenBalanceAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEthBalanceInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPrizesWallet.EthBalanceInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prizeWinnerAddress_\",\"type\":\"address\"}],\"name\":\"getEthBalanceInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPrizesWallet.EthBalanceInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mainPrizeBeneficiaryAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextDonatedNftIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mainPrizeBeneficiaryAddress_\",\"type\":\"address\"}],\"name\":\"registerRoundEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mainPrizeBeneficiaryAddress_\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"prizeWinnerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPrizesWallet.EthDeposit[]\",\"name\":\"ethDeposits_\",\"type\":\"tuple[]\"}],\"name\":\"registerRoundEndAndDepositEthMany\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"roundTimeoutTimesToWithdrawPrizes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setTimeoutDurationToWithdrawPrizes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeoutDurationToWithdrawPrizes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prizeWinnerAddress_\",\"type\":\"address\"}],\"name\":\"withdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"withdrawEth_\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPrizesWallet.DonatedTokenToClaim[]\",\"name\":\"donatedTokensToClaim_\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"donatedNftIndexes_\",\"type\":\"uint256[]\"}],\"name\":\"withdrawEverything\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523461003f57610019610014610104565b61020b565b610021610044565b613d4e6104db8239608051818181610bb001526124f80152613d4e90f35b61004a565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906100769061004e565b810190811060018060401b0382111761008e57604052565b610058565b906100a661009f610044565b928361006c565b565b5f80fd5b60018060a01b031690565b6100c0906100ac565b90565b6100cc816100b7565b036100d357565b5f80fd5b905051906100e4826100c3565b565b906020828203126100ff576100fc915f016100d7565b90565b6100a8565b6101226142298038038061011781610093565b9283398101906100e6565b90565b90565b90565b90565b61014261013d61014792610125565b61012b565b610128565b90565b610156622e248061012e565b90565b5f1b90565b9061016a5f1991610159565b9181191691161790565b61018861018361018d92610128565b61012b565b610128565b90565b90565b906101a86101a36101af92610174565b610190565b825461015e565b9055565b90565b6101ca6101c56101cf926101b3565b61012b565b610128565b90565b906101e76101e26101ee926101b6565b610190565b825461015e565b9055565b610204906101ff8161038f565b610206565b565b608052565b61025d9061021f61021a610320565b61025f565b61023961022a61014a565b68010000000000000001610193565b6102585f740200000000000000000000000300000000000000026101d2565b6101f2565b565b6102689061026a565b565b61027390610275565b565b61027e906102ca565b565b61029461028f610299926101b3565b61012b565b6100ac565b90565b6102a590610280565b90565b6102b1906100b7565b9052565b91906102c8905f602085019401906102a8565b565b806102e56102df6102da5f61029c565b6100b7565b916100b7565b146102f5576102f39061047b565b565b6103186103015f61029c565b5f918291631e4fbdf760e01b8352600483016102b5565b0390fd5b5f90565b61032861031c565b503390565b60209181520190565b5f7f5468652070726f76696465642061646472657373206973207a65726f2e000000910152565b61036a601d60209261032d565b61037381610336565b0190565b61038c9060208101905f81830391015261035d565b90565b6103a96103a361039e5f61029c565b6100b7565b916100b7565b146103b057565b6103b8610044565b63eac0d38960e01b8152806103cf60048201610377565b0390fd5b5f1c90565b60018060a01b031690565b6103ef6103f4916103d3565b6103d8565b90565b61040190546103e3565b90565b9061041560018060a01b0391610159565b9181191691161790565b61043361042e610438926100ac565b61012b565b6100ac565b90565b6104449061041f565b90565b6104509061043b565b90565b90565b9061046b61046661047292610447565b610453565b8254610404565b9055565b5f0190565b6104845f6103f7565b61048e825f610456565b906104c26104bc7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e093610447565b91610447565b916104cb610044565b806104d581610476565b0390a356fe60806040526004361015610013575b611002565b61001d5f356101bc565b80631a18c889146101b757806325e16063146101b25780634b5e1b19146101ad5780635e00aaba146101a85780636224dd3f146101a35780636ff1facd1461019e578063715018a61461019957806376e42c6a1461019457806387565d141461018f5780638da5cb5b1461018a57806394d907fc146101855780639523b237146101805780639cf10d321461017b5780639e2842a814610176578063a089e0be14610171578063a0ef91df1461016c578063b4e6f80314610167578063c3fe3e2814610162578063cc5810d81461015d578063d7f4f8be14610158578063e2051c7e14610153578063e4a6c2a41461014e578063e86a49d714610149578063f11b35fa14610144578063f2fde38b1461013f5763fe673fd30361000e57610fcb565b610f23565b610ef9565b610e89565b610e28565b610dbb565b610d40565b610c34565b610bd2565b610b77565b610a72565b610a3d565b610a0a565b6109d2565b610841565b6107a5565b610770565b610741565b61067b565b610590565b61055b565b610503565b6104b8565b610402565b6102be565b610228565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f9103126101da57565b6101cc565b90565b6101eb906101df565b9052565b90602080610211936102075f8201515f8601906101e2565b01519101906101e2565b565b9190610226905f604085019401906101ef565b565b34610258576102383660046101d0565b610254610243611137565b61024b6101c2565b91829182610213565b0390f35b6101c8565b5f80fd5b60018060a01b031690565b61027590610261565b90565b6102818161026c565b0361028857565b5f80fd5b9050359061029982610278565b565b906020828203126102b4576102b1915f0161028c565b90565b6101cc565b5f0190565b346102ec576102d66102d136600461029b565b61136e565b6102de6101c2565b806102e8816102b9565b0390f35b6101c8565b6102fa816101df565b0361030157565b5f80fd5b90503590610312826102f1565b565b9060208282031261032d5761032a915f01610305565b90565b6101cc565b634e487b7160e01b5f52603260045260245ffd5b506801000000000000000090565b90565b61036081610346565b82101561037a57610372600191610354565b910201905f90565b610332565b1c90565b90565b61039690600861039b930261037f565b610383565b90565b906103a99154610386565b90565b680100000000000000026103bf81610346565b8210156103dc576103d9916103d391610357565b9061039e565b90565b5f80fd5b6103e9906101df565b9052565b9190610400905f602085019401906103e0565b565b346104325761042e61041d610418366004610314565b6103ac565b6104256101c2565b918291826103ed565b0390f35b6101c8565b5f80fd5b5f80fd5b5f80fd5b909182601f8301121561047d5781359167ffffffffffffffff831161047857602001926060830284011161047357565b61043f565b61043b565b610437565b906020828203126104b3575f82013567ffffffffffffffff81116104ae576104aa9201610443565b9091565b61025d565b6101cc565b346104e7576104d16104cb366004610482565b906113a1565b6104d96101c2565b806104e3816102b9565b0390f35b6101c8565b610500680100000000000000015f9061039e565b90565b34610533576105133660046101d0565b61052f61051e6104ec565b6105266101c2565b918291826103ed565b0390f35b6101c8565b610558740200000000000000000000000300000000000000025f9061039e565b90565b3461058b5761056b3660046101d0565b610587610576610538565b61057e6101c2565b918291826103ed565b0390f35b6101c8565b346105be576105a03660046101d0565b6105a86113fa565b6105b06101c2565b806105ba816102b9565b0390f35b6101c8565b506801000000000000000090565b90565b6105dd816105c3565b8210156105f7576105ef6001916105d1565b910201905f90565b610332565b60018060a01b031690565b61061790600861061c930261037f565b6105fc565b90565b9061062a9154610607565b90565b6001610638816105c3565b821015610655576106529161064c916105d4565b9061061f565b90565b5f80fd5b6106629061026c565b9052565b9190610679905f60208501940190610659565b565b346106ab576106a7610696610691366004610314565b61062d565b61069e6101c2565b91829182610666565b0390f35b6101c8565b909182601f830112156106ea5781359167ffffffffffffffff83116106e55760200192604083028401116106e057565b61043f565b61043b565b610437565b9160608383031261073c57610706825f8501610305565b92610714836020830161028c565b92604082013567ffffffffffffffff81116107375761073392016106b0565b9091565b61025d565b6101cc565b61076c61075b6107523660046106ef565b92919091611521565b6107636101c2565b918291826103ed565b0390f35b346107a0576107803660046101d0565b61079c61078b61155d565b6107936101c2565b91829182610666565b0390f35b6101c8565b346107d3576107bd6107b8366004610314565b611598565b6107c56101c2565b806107cf816102b9565b0390f35b6101c8565b6107e19061026c565b90565b6107ed816107d8565b036107f457565b5f80fd5b90503590610805826107e4565b565b909160608284031261083c57610839610822845f8501610305565b9361083081602086016107f8565b93604001610305565b90565b6101cc565b346108705761085a610854366004610807565b916115cf565b6108626101c2565b8061086c816102b9565b0390f35b6101c8565b506801000000000000000090565b90565b61088f81610875565b8210156108a9576108a1600391610883565b910201905f90565b610332565b5f1c90565b6108bf6108c4916108ae565b610383565b90565b6108d190546108b3565b90565b60018060a01b031690565b6108eb6108f0916108ae565b6108d4565b90565b6108fd90546108df565b90565b740200000000000000000000000300000000000000039061092082610875565b8110156109585761093091610886565b5061093c5f82016108c7565b91610955600261094e600185016108f3565b93016108c7565b90565b5f80fd5b90565b61097361096e61097892610261565b61095c565b610261565b90565b6109849061095f565b90565b6109909061097b565b90565b61099c90610987565b9052565b6040906109c96109d094969593966109bf60608401985f8501906103e0565b6020830190610993565b01906103e0565b565b34610a0557610a016109ed6109e8366004610314565b610900565b6109f89391936101c2565b938493846109a0565b0390f35b6101c8565b34610a3857610a22610a1d366004610314565b611697565b610a2a6101c2565b80610a34816102b9565b0390f35b6101c8565b34610a6d57610a69610a58610a5336600461029b565b6116a2565b610a606101c2565b91829182610213565b0390f35b6101c8565b34610aa057610a823660046101d0565b610a8a6116f8565b610a926101c2565b80610a9c816102b9565b0390f35b6101c8565b151590565b610ab381610aa5565b03610aba57565b5f80fd5b90503590610acb82610aaa565b565b909182601f83011215610b075781359167ffffffffffffffff8311610b02576020019260208302840111610afd57565b61043f565b61043b565b610437565b606081830312610b7257610b22825f8301610abe565b92602082013567ffffffffffffffff8111610b6d5783610b43918401610443565b929093604082013567ffffffffffffffff8111610b6857610b649201610acd565b9091565b61025d565b61025d565b6101cc565b34610ba957610b93610b8a366004610b0c565b9392909261174a565b610b9b6101c2565b80610ba5816102b9565b0390f35b6101c8565b7f000000000000000000000000000000000000000000000000000000000000000090565b34610c0257610be23660046101d0565b610bfe610bed610bae565b610bf56101c2565b91829182610666565b0390f35b6101c8565b9190604083820312610c2f5780610c23610c2c925f8601610305565b9360200161028c565b90565b6101cc565b34610c6557610c61610c50610c4a366004610c07565b9061179c565b610c586101c2565b918291826103ed565b0390f35b6101c8565b506801000000000000000090565b90565b610c8481610c6a565b821015610c9e57610c96600191610c78565b910201905f90565b610332565b60018060a01b031690565b610cba610cbf916108ae565b610ca3565b90565b610ccc9054610cae565b90565b7402000000000000000000000002000000000000000290610cef82610c6a565b811015610d0e57610d045f91610d0b93610c7b565b5001610cc2565b90565b5f80fd5b610d1b9061097b565b90565b610d2790610d12565b9052565b9190610d3e905f60208501940190610d1e565b565b34610d7057610d6c610d5b610d56366004610314565b610ccf565b610d636101c2565b91829182610d2b565b0390f35b6101c8565b608081830312610db657610d8b825f8301610305565b92610db3610d9c846020850161028c565b93610daa81604086016107f8565b93606001610305565b90565b6101cc565b34610ded57610dd7610dce366004610d75565b92919091611b12565b610ddf6101c2565b80610de9816102b9565b0390f35b6101c8565b90602082820312610e23575f82013567ffffffffffffffff8111610e1e57610e1a9201610acd565b9091565b61025d565b6101cc565b34610e5757610e41610e3b366004610df2565b90611b48565b610e496101c2565b80610e53816102b9565b0390f35b6101c8565b9190604083820312610e845780610e78610e81925f8601610305565b936020016107f8565b90565b6101cc565b34610eba57610eb6610ea5610e9f366004610e5c565b90611b81565b610ead6101c2565b918291826103ed565b0390f35b6101c8565b9091606082840312610ef457610ef1610eda845f8501610305565b93610ee88160208601610305565b9360400161028c565b90565b6101cc565b610f0d610f07366004610ebf565b91611cd4565b610f156101c2565b80610f1f816102b9565b0390f35b34610f5157610f3b610f3636600461029b565b611d46565b610f436101c2565b80610f4d816102b9565b0390f35b6101c8565b610f5f9061026c565b90565b610f6b81610f56565b03610f7257565b5f80fd5b90503590610f8382610f62565b565b608081830312610fc657610f9b825f8301610305565b92610fc3610fac846020850161028c565b93610fba8160408601610f76565b93606001610305565b90565b6101cc565b34610ffd57610fe7610fde366004610f85565b92919091611fc3565b610fef6101c2565b80610ff9816102b9565b0390f35b6101c8565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b9061102e90611006565b810190811067ffffffffffffffff82111761104857604052565b611010565b906110606110596101c2565b9283611024565b565b61106c604061104d565b90565b5f90565b61107b611062565b906020808361108861106f565b81520161109361106f565b81525050565b6110a1611073565b90565b6110ad9061095f565b90565b50600160a01b90565b90565b6110c5816110b0565b8210156110df576110d76002916110b9565b910201905f90565b610332565b906110ee906101df565b9052565b906111296111206001611103611062565b9461111a6111125f83016108c7565b5f88016110e4565b016108c7565b602084016110e4565b565b611134906110f2565b90565b61113f611099565b5061116c6111666802000000000000000261116061115b611fd1565b6110a4565b906110bc565b5061112b565b90565b6111809061117b61202c565b6112c8565b611188612071565b565b90565b90565b6111a461119f6111a99261118d565b61095c565b6101df565b90565b60209181520190565b60407f206265666f726520612074696d656f757420657870697265732e000000000000917f4f6e6c792074686520455448207072697a652077696e6e6572206973207065725f8201527f6d697474656420746f2077697468647261772074686569722062616c616e636560208201520152565b611235605a6060926111ac565b61123e816111b5565b0190565b61128761128e9461127d60809498979561127361126860a087018781035f890152611228565b9a6020870190610659565b6040850190610659565b60608301906103e0565b01906103e0565b565b939193929092156112a15750505050565b906112c4916112ae6101c2565b94859463380b9dd360e21b865260048601611242565b0390fd5b611350906112f16112eb680200000000000000026112e5846110a4565b906110bc565b5061118a565b9061134b61131c611316680100000000000000026113105f87016108c7565b90610357565b9061039e565b4261132f611329836101df565b916101df565b10159081611352575b8390611342611fd1565b90914293611290565b6121e8565b565b8091506113676113615f611190565b916101df565b1190611338565b6113779061116f565b565b9061138b9161138661202c565b611395565b611393612071565b565b9061139f916122f0565b565b906113ab91611379565b565b6113b5612366565b6113bd6113e7565b565b6113d36113ce6113d89261118d565b61095c565b610261565b90565b6113e4906113bf565b90565b6113f86113f35f6113db565b6123d7565b565b6114026113ad565b565b5f90565b9061141d9493929161141861202c565b611428565b90611426612071565b565b9061143d949392916114386124e8565b6114a3565b90565b5090565b634e487b7160e01b5f52601160045260245ffd5b611461906101df565b5f811461146f576001900390565b611444565b9190811015611484576040020190565b610332565b3561149381610278565b90565b356114a0816102f1565b90565b916114b3919394959250836125c7565b906114bf848690611440565b5b806114d36114cd5f611190565b916101df565b1115611518576114e290611458565b916115126114f286888691611474565b85908561150c60206115055f8501611489565b9301611496565b92612624565b916114c0565b50935091505090565b90611535939291611530611404565b611408565b90565b5f90565b61154861154d916108ae565b6105fc565b90565b61155a905461153c565b90565b611565611538565b5061156f5f611550565b90565b6115839061157e61202c565b61158d565b61158b612071565b565b611596906129b5565b565b6115a190611572565b565b906115b692916115b161202c565b6115c0565b6115be612071565b565b916115cd92919091612e07565b565b906115da92916115a3565b565b6115ed906115e8612366565b611649565b565b5f1b90565b906116005f19916115ef565b9181191691161790565b61161e611619611623926101df565b61095c565b6101df565b90565b90565b9061163e6116396116459261160a565b611626565b82546115f4565b9055565b61165c8168010000000000000001611629565b6116927f8717bb199c6bc4a5dadb21547205f9ef8ec037dda246a5526d6a6471306ea52e916116896101c2565b918291826103ed565b0390a1565b6116a0906115dc565b565b6116cb6116d1916116b1611099565b506116c568020000000000000002916110a4565b906110bc565b5061112b565b90565b6116dc61202c565b6116e46116ee565b6116ec612071565b565b6116f661302d565b565b6117006116d4565b565b906117179493929161171261202c565b611721565b61171f612071565b565b916117369161173b95949361173d575b6122f0565b613084565b565b61174561302d565b611731565b9061175794939291611702565b565b9061176c929161176761202c565b611777565b90611775612071565b565b9061178a92916117856124e8565b61178d565b90565b906117999291506125c7565b90565b906117ae916117a9611404565b611759565b90565b906117c59392916117c061202c565b6117cf565b6117cd612071565b565b906117e39392916117de6124e8565b61191a565b565b90565b6117f19061026c565b90565b906117fe906117e8565b9052565b61180c602061104d565b90565b9061182e6118265f61181f611802565b9401610cc2565b5f84016117f4565b565b6118399061180f565b90565b61184690516117e8565b90565b5f80fd5b60e01b90565b5f91031261185d57565b6101cc565b61186b9061097b565b90565b61187790611862565b9052565b919061188e905f6020850194019061186e565b565b6118986101c2565b3d5f823e3d90fd5b906118b160018060a01b03916115ef565b9181191691161790565b6118c49061095f565b90565b6118d0906118bb565b90565b90565b906118eb6118e66118f2926118c7565b6118d3565b82546118a0565b9055565b6118ff9061097b565b90565b61190b9061095f565b90565b61191790611902565b90565b92909192611947611941740200000000000000000000000200000000000000028390610c7b565b506117e5565b9061195182611830565b916119656119605f850161183c565b610d12565b61197f6119796119745f6113db565b61026c565b9161026c565b145f14611a6f578561198f6101c2565b90610a8a820182811067ffffffffffffffff821117611a6a5782916119bb91610a8a61328f853961187b565b03905ff0908115611a6557611a6396611a5c936119f55f936119e2611a5796868a016117f4565b846119ee818a0161183c565b91016118d6565b5b8782908892611a4c611a3a611a34611a2e7f3f94f6171c96ab491d8c268f80da5e460b148878065711a0bb099bbc297615af9461160a565b946118f6565b9461190e565b94611a436101c2565b918291826103ed565b0390a495930161183c565b610d12565b919261310e565b565b611890565b611010565b50611a83611a7e5f840161183c565b610d12565b630b820b5486823b15611b0d57611ab992611aae5f8094611aa26101c2565b9687958694859361184d565b83526004830161187b565b03925af18015611b0857611a6396611a5c93611a57935f93611adc575b506119f6565b611afb90843d8111611b01575b611af38183611024565b810190611853565b5f611ad6565b503d611ae9565b611890565b611849565b90611b1e9392916117b1565b565b90611b3291611b2d61202c565b611b3c565b611b3a612071565b565b90611b4691613084565b565b90611b5291611b20565b565b90505190611b61826102f1565b565b90602082820312611b7c57611b79915f01611b54565b90565b6101cc565b611bb5611baf611bba92611b93611404565b5074020000000000000000000000020000000000000002610c7b565b506117e5565b611830565b90611bce611bc95f840161183c565b610d12565b611be8611be2611bdd5f6113db565b61026c565b9161026c565b145f14611bfe575050611bfa5f611190565b5b90565b611c4891611c0d602092611862565b611c3d611c28611c235f6370a08231950161183c565b610d12565b92611c316101c2565b9586948593849361184d565b835260048301610666565b03915afa908115611c8e575f91611c60575b50611bfb565b611c81915060203d8111611c87575b611c798183611024565b810190611b63565b5f611c5a565b503d611c6f565b611890565b90611ca69291611ca161202c565b611cb0565b611cae612071565b565b90611cc39291611cbe6124e8565b611cc5565b565b91611cd292913492612624565b565b90611cdf9291611c93565b565b611cf290611ced612366565b611cf4565b565b80611d0f611d09611d045f6113db565b61026c565b9161026c565b14611d1f57611d1d906123d7565b565b611d42611d2b5f6113db565b5f918291631e4fbdf760e01b835260048301610666565b0390fd5b611d4f90611ce1565b565b90611d65939291611d6061202c565b611d6f565b611d6d612071565b565b90611d83939291611d7e6124e8565b611e40565b565b90565b611d919061095f565b90565b611d9d90611d88565b90565b90565b90611db8611db3611dbf92611d94565b611da0565b82546118a0565b9055565b916020611de4929493611ddd60408201965f8301906103e0565b01906103e0565b565b611def906101df565b5f198114611dfd5760010190565b611444565b611e0b9061097b565b90565b604090611e37611e3e9496959396611e2d60608401985f850190610659565b6020830190610659565b01906103e0565b565b91611f34611f19611f399394611e6a740200000000000000000000000300000000000000026108c7565b90611eba611e97611e91740200000000000000000000000300000000000000038590610886565b50611d85565b611ea3835f8301611629565b611eb08760018301611da3565b60028a9101611629565b868590899284611efc611ef6611ef07fb12e72bab0c2a8fe98175a3c02792645ebdf65d37cf3983517e6d1c5ab842e239461160a565b946118f6565b94611d94565b94611f11611f086101c2565b92839283611dc3565b0390a4611de6565b74020000000000000000000000030000000000000002611629565b610987565b6323b872dd9190611f4930611e02565b9392813b15611fbe575f611f7091611f7b8296611f646101c2565b9889978896879561184d565b855260048501611e0e565b03925af18015611fb957611f8d575b50565b611fac905f3d8111611fb2575b611fa48183611024565b810190611853565b5f611f8a565b503d611f9a565b611890565b611849565b90611fcf939291611d51565b565b611fd9611538565b503390565b90565b90565b611ff8611ff3611ffd92611fde565b6115ef565b611fe1565b90565b6120297f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00611fe4565b90565b612034613162565b6120555761205361204b612046612000565b613196565b6001906131ab565b565b5f633ee5aeb560e01b81528061206d600482016102b9565b0390fd5b61208b61208461207f612000565b613196565b5f906131ab565b565b1b90565b919060086120ac9102916120a65f198461208d565b9261208d565b9181191691161790565b91906120cc6120c76120d49361160a565b611626565b908354612091565b9055565b6120ea916120e4611404565b916120b6565b565b905090565b6120fc5f80926120ec565b0190565b612109906120f1565b90565b67ffffffffffffffff811161212a57612126602091611006565b0190565b611010565b9061214161213c8361210c565b61104d565b918252565b606090565b3d5f146121665761215b3d61212f565b903d5f602084013e5b565b61216e612146565b90612164565b5f7f455448207769746864726177616c206661696c65642e00000000000000000000910152565b6121a860166020926111ac565b6121b181612174565b0190565b9160406121e69294936121df6121d4606083018381035f85015261219b565b966020830190610659565b01906103e0565b565b61220d5f806121f9600186016108c7565b9461220782600183016120d8565b016120d8565b612215611fd1565b829161225f61224d6122477f4f43b861ba36494acfe938f3815fba7fac6981bdc611b6ccdc14c08f59292383936118f6565b936118f6565b936122566101c2565b918291826103ed565b0390a36122965f8061226f611fd1565b846122786101c2565b908161228381612100565b03925af161228f61214b565b5015610aa5565b61229d5750565b6122a5611fd1565b6122c66122b06101c2565b928392630aa7db6360e11b8452600484016121b5565b0390fd5b5090565b91908110156122de576060020190565b610332565b356122ed816107e4565b90565b90916122fd8284906122ca565b5b8061231161230b5f611190565b916101df565b11156123605761232090611458565b9061235a612330848685916122ce565b61233b5f8201611496565b90612354604061234d602084016122e3565b9201611496565b91612e07565b906122fe565b50915050565b61236e61155d565b61238761238161237c611fd1565b61026c565b9161026c565b0361238e57565b6123b0612399611fd1565b5f91829163118cdaa760e01b835260048301610666565b0390fd5b90565b906123cc6123c76123d3926118f6565b6123b4565b82546118a0565b9055565b6123e05f611550565b6123ea825f6123b7565b9061241e6124187f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0936118f6565b916118f6565b916124276101c2565b80612431816102b9565b0390a3565b60407f6d6574686f642e00000000000000000000000000000000000000000000000000917f4f6e6c792074686520436f736d69635369676e617475726547616d6520636f6e5f8201527f7472616374206973207065726d697474656420746f2063616c6c20746869732060208201520152565b6124b660476060926111ac565b6124bf81612436565b0190565b91906124e69060206124de604086018681035f8801526124a9565b940190610659565b565b6124f0611fd1565b61252261251c7f000000000000000000000000000000000000000000000000000000000000000061026c565b9161026c565b0361252957565b612552612534611fd1565b61253c6101c2565b91829163ced50f6760e01b8352600483016124c3565b0390fd5b9190600861257691029161257060018060a01b038461208d565b9261208d565b9181191691161790565b919061259661259161259e936118f6565b6123b4565b908354612556565b9055565b6125b16125b7919392936101df565b926101df565b82018092116125c257565b611444565b906125e7906125d4611404565b506125e1600184906105d4565b90612580565b61262161260642612600680100000000000000016108c7565b906125a2565b9161261b839168010000000000000002610357565b906120b6565b90565b92919261267861264f61264968020000000000000002612643886110a4565b906110bc565b5061118a565b61265b835f8301611629565b61267260018692019161266d836108c7565b6125a2565b90611629565b9092916126ae6126a87f8e369548f9ef214c7c1823c098a23763321fd761ac5cf78958e1db1b3648e7a29361160a565b936118f6565b936126c36126ba6101c2565b92839283611dc3565b0390a3565b906126d290610f56565b9052565b6126e0606061104d565b90565b9061273261272960026126f46126d6565b9461270b6127035f83016108c7565b5f88016110e4565b61272361271a600183016108f3565b602088016126c8565b016108c7565b604084016110e4565b565b61273d906126e3565b90565b61274a9051610f56565b90565b5f7f496e76616c696420646f6e61746564204e465420696e6465782e000000000000910152565b612781601a6020926111ac565b61278a8161274d565b0190565b9160406127bf9294936127b86127ad606083018381035f850152612774565b966020830190610659565b01906103e0565b565b5f7f446f6e61746564204e465420616c726561647920636c61696d65642e00000000910152565b6127f5601c6020926111ac565b6127fe816127c1565b0190565b91604061283392949361282c612821606083018381035f8501526127e8565b966020830190610659565b01906103e0565b565b61283f90516101df565b90565b60607f70697265732e0000000000000000000000000000000000000000000000000000917f4f6e6c79207468652062696464696e6720726f756e64206d61696e207072697a5f8201527f652062656e6566696369617279206973207065726d697474656420746f20636c60208201527f61696d2074686973204e4654206265666f726520612074696d656f757420657860408201520152565b6128e860666080926111ac565b6128f181612842565b0190565b61293a6129419461293060809498979561292661291b60a087018781035f8901526128db565b9a6020870190610659565b60408501906103e0565b60608301906103e0565b01906103e0565b565b939193929092156129545750505050565b90612977916129616101c2565b94859463b97dcd1760e01b8652600486016128f5565b0390fd5b919061299161298c61299993611d94565b611da0565b908354612556565b9055565b5f90565b6129b3916129ad61299d565b9161297b565b565b6129de6129d8740200000000000000000000000300000000000000038390610886565b50611d85565b906129e882612734565b916129fd6129f860208501612740565b610987565b612a17612a11612a0c5f6113db565b61026c565b9161026c565b14612c2c575f6002612a8092612a2b611fd1565b612a5b612a55612a50612a4a6001612a44898d01612835565b906105d4565b9061061f565b61026c565b9161026c565b03612bb2575b612a6d838083016120d8565b612a7a83600183016129a1565b016120d8565b612a8b5f8301612835565b612a93611fd1565b612a9f60208501612740565b91612aac60408601612835565b93612ae9612ae3612add7f03c2b6e01c9ca39e4073132f279b61b987d41a82b08cea3dd4a2fabc37067ac39461160a565b946118f6565b94611d94565b94612afe612af56101c2565b92839283611dc3565b0390a4612b15612b1060208301612740565b610987565b6323b872dd90612b2430611e02565b90612b396040612b32611fd1565b9501612835565b92813b15612bad575f612b5f91612b6a8296612b536101c2565b9889978896879561184d565b855260048501611e0e565b03925af18015612ba857612b7c575b50565b612b9b905f3d8111612ba1575b612b938183611024565b810190611853565b5f612b79565b503d612b89565b611890565b611849565b612c0b612bdc612bd668010000000000000002612bd0878b01612835565b90610357565b9061039e565b42612bef612be9836101df565b916101df565b10159081612c10575b612c00611fd1565b908890914293612943565b612a61565b809150612c25612c1f87611190565b916101df565b1190612bf8565b5080612c5d612c57612c52740200000000000000000000000300000000000000026108c7565b6101df565b916101df565b1015612c9057612c6b611fd1565b612c8c612c766101c2565b92839263581c778d60e01b845260048401612802565b0390fd5b612c98611fd1565b612cb9612ca36101c2565b9283926373b047ef60e11b84526004840161278e565b0390fd5b60607f65666f726520612074696d656f757420657870697265732e0000000000000000917f4f6e6c79207468652062696464696e6720726f756e64206d61696e207072697a5f8201527f652062656e6566696369617279206973207065726d697474656420746f20636c60208201527f61696d2074686973204552432d323020746f6b656e20646f6e6174696f6e206260408201520152565b612d6360786080926111ac565b612d6c81612cbd565b0190565b90959492612dcb94612dba612dc492612db060a096612da6612d9b60c089018981035f8b0152612d56565b9c60208901906103e0565b6040870190610659565b606085019061186e565b60808301906103e0565b01906103e0565b565b94929093919415612ddf575050505050565b90612e039291612ded6101c2565b95869563c8568d6360e01b875260048701612d70565b0390fd5b9091612e11611fd1565b612e37612e31612e2c612e26600187906105d4565b9061061f565b61026c565b9161026c565b03612fbc575b612e6e612e69612e63740200000000000000000000000200000000000000028590610c7b565b506117e5565b611830565b9181612e82612e7c5f611190565b916101df565b14612f11575b612efb5f612f0f95612f0093612e9c611fd1565b82908792612ef1612edf612ed9612ed37faf1adae2e1e983ca738335dc2e37194114142793f394ac934a45ea632a8a5bd09461160a565b946118f6565b9461190e565b94612ee86101c2565b918291826103ed565b0390a4940161183c565b610d12565b612f08611fd1565b919261310e565b565b9050826020612f22612f5e95611862565b6370a0823190612f53612f3e612f395f890161183c565b610d12565b92612f476101c2565b9889948593849361184d565b835260048301610666565b03915afa8015612fb7575f612f0092612efb92612f0f978391612f89575b5094935095505050612e88565b612faa915060203d8111612fb0575b612fa28183611024565b810190611b63565b5f612f7c565b503d612f98565b611890565b61300c612fdc612fd6680100000000000000028590610357565b9061039e565b42612fef612fe9836101df565b916101df565b10159081613011575b8490613002611fd1565b8791924294612dcd565b612e3d565b8091506130266130205f611190565b916101df565b1190612ff8565b61306961305c6130566802000000000000000261305061304b611fd1565b6110a4565b906110bc565b5061118a565b613064611fd1565b6121e8565b565b5090565b919081101561307f576020020190565b610332565b909161309182849061306b565b5b806130a561309f5f611190565b916101df565b11156130d7576130b490611458565b906130d16130cc6130c78587869161306f565b611496565b6129b5565b90613092565b50915050565b63ffffffff1690565b63ffffffff60e01b1690565b61310661310161310b926130dd565b61184d565b6130e6565b90565b60049261314861315c9593613157939461312f6323b872dd929491926130f2565b936131386101c2565b9788956020870190815201611e0e565b60208201810382520383611024565b6131cd565b565b5f90565b61316a61315e565b5061318361317e613179612000565b613196565b613281565b90565b5f90565b61319390611fe1565b90565b6131a8906131a2613186565b5061318a565b90565b5d565b90565b6131c56131c06131ca926131ae565b61095c565b6101df565b90565b905f6020916131da611404565b506131e3611404565b50828151910182855af115613276573d5f51906132086132025f611190565b916101df565b145f1461325c575061321981611862565b3b61322c6132265f611190565b916101df565b145b6132355750565b61324161325891611862565b5f918291635274afe760e01b835260048301610666565b0390fd5b61326f61326960016131b1565b916101df565b141561322e565b6040513d5f823e3d90fd5b61328961315e565b505c9056fe60a06040523461003f57610019610014610110565b610131565b610021610044565b6105d46104b6823960805181818161018e01526101eb01526105d490f35b61004a565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906100769061004e565b810190811060018060401b0382111761008e57604052565b610058565b906100a661009f610044565b928361006c565b565b5f80fd5b60018060a01b031690565b6100c0906100ac565b90565b6100cc906100b7565b90565b6100d8816100c3565b036100df57565b5f80fd5b905051906100f0826100cf565b565b9060208282031261010b57610108915f016100e3565b90565b6100a8565b61012e610a8a8038038061012381610093565b9283398101906100f2565b90565b61013e903360805261014d565b565b61014a90516100b7565b90565b6101639061015b6080610140565b5f1991610267565b565b90565b61017c610177610181926100ac565b610165565b6100ac565b90565b61018d90610168565b90565b61019990610184565b90565b63ffffffff1690565b63ffffffff60e01b1690565b60e01b90565b6101cb6101c66101d09261019c565b6101b1565b6101a5565b90565b6101dc906100b7565b9052565b90565b6101ec906101e0565b9052565b91602061021192949361020a60408201965f8301906101d3565b01906101e3565b565b151590565b90565b61022f61022a61023492610218565b610165565b6101e0565b90565b6102409061021b565b9052565b91602061026592949361025e60408201965f8301906101d3565b0190610237565b565b90916102aa60049161029b63095ea7b36102828793916101b7565b9261028b610044565b95869460208601908152016101f0565b6020820181038252038261006c565b906102bf6102b982849061034a565b15610213565b6102c9575b505050565b610316600461031b94610311849161030263095ea7b36102e95f916101b7565b926102f2610044565b9687946020860190815201610244565b6020820181038252038361006c565b610401565b610401565b5f80806102c4565b5f90565b5f90565b90565b61034261033d6103479261032b565b610165565b6101e0565b90565b905f602091610357610323565b50610360610323565b50610369610327565b50610372610327565b50828151910182855af13d915f5191928361038e575b50505090565b909192506103a461039e5f61021b565b916101e0565b145f146103d2576103b59150610190565b3b6103c86103c25f61021b565b916101e0565b115b5f8080610388565b506103e66103e0600161032e565b916101e0565b146103ca565b91906103ff905f602085019401906101d3565b565b905f60209161040e610327565b50610417610327565b50828151910182855af1156104aa573d5f519061043c6104365f61021b565b916101e0565b145f14610490575061044d81610190565b3b61046061045a5f61021b565b916101e0565b145b6104695750565b61047561048c91610190565b5f918291635274afe760e01b8352600483016103ec565b0390fd5b6104a361049d600161032e565b916101e0565b1415610462565b6040513d5f823e3d90fdfe60806040526004361015610013575b6100dc565b61001d5f3561002c565b630b820b540361000e576100a9565b60e01c90565b60405190565b5f80fd5b5f80fd5b60018060a01b031690565b61005490610040565b90565b6100609061004b565b90565b61006c81610057565b0361007357565b5f80fd5b9050359061008482610063565b565b9060208282031261009f5761009c915f01610077565b90565b61003c565b5f0190565b346100d7576100c16100bc366004610086565b6100fe565b6100c9610032565b806100d3816100a4565b0390f35b610038565b5f80fd5b6100f1906100ec610185565b6100f3565b565b6100fc906101e5565b565b610107906100e0565b565b60209181520190565b5f7f4465706c6f796572206f6e6c792e000000000000000000000000000000000000910152565b610146600e602092610109565b61014f81610112565b0190565b61015c9061004b565b9052565b919061018390602061017b604086018681035f880152610139565b940190610153565b565b336101b86101b27f000000000000000000000000000000000000000000000000000000000000000061004b565b9161004b565b036101bf57565b6101e1336101cb610032565b91829163ced50f6760e01b835260048301610160565b0390fd5b610212907f00000000000000000000000000000000000000000000000000000000000000005f1991610350565b565b90565b61022b61022661023092610040565b610214565b610040565b90565b61023c90610217565b90565b61024890610233565b90565b63ffffffff1690565b63ffffffff60e01b1690565b60e01b90565b61027a61027561027f9261024b565b610260565b610254565b90565b90565b61028e90610282565b9052565b9160206102b39294936102ac60408201965f830190610153565b0190610285565b565b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906102dd906102b5565b810190811067ffffffffffffffff8211176102f757604052565b6102bf565b151590565b90565b61031861031361031d92610301565b610214565b610282565b90565b61032990610304565b9052565b91602061034e92949361034760408201965f830190610153565b0190610320565b565b909161039360049161038463095ea7b361036b879391610266565b92610374610032565b9586946020860190815201610292565b602082018103825203826102d3565b906103a86103a2828490610433565b156102fc565b6103b2575b505050565b6103ff6004610404946103fa84916103eb63095ea7b36103d25f91610266565b926103db610032565b968794602086019081520161032d565b602082018103825203836102d3565b6104ea565b6104ea565b5f80806103ad565b5f90565b5f90565b90565b61042b61042661043092610414565b610214565b610282565b90565b905f60209161044061040c565b5061044961040c565b50610452610410565b5061045b610410565b50828151910182855af13d915f51919283610477575b50505090565b9091925061048d6104875f610304565b91610282565b145f146104bb5761049e915061023f565b3b6104b16104ab5f610304565b91610282565b115b5f8080610471565b506104cf6104c96001610417565b91610282565b146104b3565b91906104e8905f60208501940190610153565b565b905f6020916104f7610410565b50610500610410565b50828151910182855af115610593573d5f519061052561051f5f610304565b91610282565b145f1461057957506105368161023f565b3b6105496105435f610304565b91610282565b145b6105525750565b61055e6105759161023f565b5f918291635274afe760e01b8352600483016104d5565b0390fd5b61058c6105866001610417565b91610282565b141561054b565b6040513d5f823e3d90fdfea264697066735822122087d61c6822975843ffd4abcd10e31147c9dead87998f5d391fdd21132f72cc9564736f6c634300081e0033a2646970667358221220916e0b73b2dce002935d19e7fc7e4c33d6f6aa602b59f589855f6444c7d4351364736f6c634300081e0033",
}

// PrizesWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use PrizesWalletMetaData.ABI instead.
var PrizesWalletABI = PrizesWalletMetaData.ABI

// PrizesWalletBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PrizesWalletMetaData.Bin instead.
var PrizesWalletBin = PrizesWalletMetaData.Bin

// DeployPrizesWallet deploys a new Ethereum contract, binding an instance of PrizesWallet to it.
func DeployPrizesWallet(auth *bind.TransactOpts, backend bind.ContractBackend, game_ common.Address) (common.Address, *types.Transaction, *PrizesWallet, error) {
	parsed, err := PrizesWalletMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PrizesWalletBin), backend, game_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PrizesWallet{PrizesWalletCaller: PrizesWalletCaller{contract: contract}, PrizesWalletTransactor: PrizesWalletTransactor{contract: contract}, PrizesWalletFilterer: PrizesWalletFilterer{contract: contract}}, nil
}

// PrizesWallet is an auto generated Go binding around an Ethereum contract.
type PrizesWallet struct {
	PrizesWalletCaller     // Read-only binding to the contract
	PrizesWalletTransactor // Write-only binding to the contract
	PrizesWalletFilterer   // Log filterer for contract events
}

// PrizesWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type PrizesWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrizesWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PrizesWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrizesWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PrizesWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrizesWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PrizesWalletSession struct {
	Contract     *PrizesWallet     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PrizesWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PrizesWalletCallerSession struct {
	Contract *PrizesWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PrizesWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PrizesWalletTransactorSession struct {
	Contract     *PrizesWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PrizesWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type PrizesWalletRaw struct {
	Contract *PrizesWallet // Generic contract binding to access the raw methods on
}

// PrizesWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PrizesWalletCallerRaw struct {
	Contract *PrizesWalletCaller // Generic read-only contract binding to access the raw methods on
}

// PrizesWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PrizesWalletTransactorRaw struct {
	Contract *PrizesWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrizesWallet creates a new instance of PrizesWallet, bound to a specific deployed contract.
func NewPrizesWallet(address common.Address, backend bind.ContractBackend) (*PrizesWallet, error) {
	contract, err := bindPrizesWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PrizesWallet{PrizesWalletCaller: PrizesWalletCaller{contract: contract}, PrizesWalletTransactor: PrizesWalletTransactor{contract: contract}, PrizesWalletFilterer: PrizesWalletFilterer{contract: contract}}, nil
}

// NewPrizesWalletCaller creates a new read-only instance of PrizesWallet, bound to a specific deployed contract.
func NewPrizesWalletCaller(address common.Address, caller bind.ContractCaller) (*PrizesWalletCaller, error) {
	contract, err := bindPrizesWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletCaller{contract: contract}, nil
}

// NewPrizesWalletTransactor creates a new write-only instance of PrizesWallet, bound to a specific deployed contract.
func NewPrizesWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*PrizesWalletTransactor, error) {
	contract, err := bindPrizesWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletTransactor{contract: contract}, nil
}

// NewPrizesWalletFilterer creates a new log filterer instance of PrizesWallet, bound to a specific deployed contract.
func NewPrizesWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*PrizesWalletFilterer, error) {
	contract, err := bindPrizesWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletFilterer{contract: contract}, nil
}

// bindPrizesWallet binds a generic wrapper to an already deployed contract.
func bindPrizesWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PrizesWalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrizesWallet *PrizesWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PrizesWallet.Contract.PrizesWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrizesWallet *PrizesWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrizesWallet.Contract.PrizesWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrizesWallet *PrizesWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrizesWallet.Contract.PrizesWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrizesWallet *PrizesWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PrizesWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrizesWallet *PrizesWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrizesWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrizesWallet *PrizesWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrizesWallet.Contract.contract.Transact(opts, method, params...)
}

// DonatedNfts is a free data retrieval call binding the contract method 0x9cf10d32.
//
// Solidity: function donatedNfts(uint256 ) view returns(uint256 roundNum, address nftAddress, uint256 nftId)
func (_PrizesWallet *PrizesWalletCaller) DonatedNfts(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RoundNum   *big.Int
	NftAddress common.Address
	NftId      *big.Int
}, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "donatedNfts", arg0)

	outstruct := new(struct {
		RoundNum   *big.Int
		NftAddress common.Address
		NftId      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundNum = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NftAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.NftId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DonatedNfts is a free data retrieval call binding the contract method 0x9cf10d32.
//
// Solidity: function donatedNfts(uint256 ) view returns(uint256 roundNum, address nftAddress, uint256 nftId)
func (_PrizesWallet *PrizesWalletSession) DonatedNfts(arg0 *big.Int) (struct {
	RoundNum   *big.Int
	NftAddress common.Address
	NftId      *big.Int
}, error) {
	return _PrizesWallet.Contract.DonatedNfts(&_PrizesWallet.CallOpts, arg0)
}

// DonatedNfts is a free data retrieval call binding the contract method 0x9cf10d32.
//
// Solidity: function donatedNfts(uint256 ) view returns(uint256 roundNum, address nftAddress, uint256 nftId)
func (_PrizesWallet *PrizesWalletCallerSession) DonatedNfts(arg0 *big.Int) (struct {
	RoundNum   *big.Int
	NftAddress common.Address
	NftId      *big.Int
}, error) {
	return _PrizesWallet.Contract.DonatedNfts(&_PrizesWallet.CallOpts, arg0)
}

// DonatedTokens is a free data retrieval call binding the contract method 0xd7f4f8be.
//
// Solidity: function donatedTokens(uint256 ) view returns(address holder)
func (_PrizesWallet *PrizesWalletCaller) DonatedTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "donatedTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DonatedTokens is a free data retrieval call binding the contract method 0xd7f4f8be.
//
// Solidity: function donatedTokens(uint256 ) view returns(address holder)
func (_PrizesWallet *PrizesWalletSession) DonatedTokens(arg0 *big.Int) (common.Address, error) {
	return _PrizesWallet.Contract.DonatedTokens(&_PrizesWallet.CallOpts, arg0)
}

// DonatedTokens is a free data retrieval call binding the contract method 0xd7f4f8be.
//
// Solidity: function donatedTokens(uint256 ) view returns(address holder)
func (_PrizesWallet *PrizesWalletCallerSession) DonatedTokens(arg0 *big.Int) (common.Address, error) {
	return _PrizesWallet.Contract.DonatedTokens(&_PrizesWallet.CallOpts, arg0)
}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_PrizesWallet *PrizesWalletCaller) Game(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "game")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_PrizesWallet *PrizesWalletSession) Game() (common.Address, error) {
	return _PrizesWallet.Contract.Game(&_PrizesWallet.CallOpts)
}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_PrizesWallet *PrizesWalletCallerSession) Game() (common.Address, error) {
	return _PrizesWallet.Contract.Game(&_PrizesWallet.CallOpts)
}

// GetDonatedTokenBalanceAmount is a free data retrieval call binding the contract method 0xe86a49d7.
//
// Solidity: function getDonatedTokenBalanceAmount(uint256 roundNum_, address tokenAddress_) view returns(uint256)
func (_PrizesWallet *PrizesWalletCaller) GetDonatedTokenBalanceAmount(opts *bind.CallOpts, roundNum_ *big.Int, tokenAddress_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "getDonatedTokenBalanceAmount", roundNum_, tokenAddress_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDonatedTokenBalanceAmount is a free data retrieval call binding the contract method 0xe86a49d7.
//
// Solidity: function getDonatedTokenBalanceAmount(uint256 roundNum_, address tokenAddress_) view returns(uint256)
func (_PrizesWallet *PrizesWalletSession) GetDonatedTokenBalanceAmount(roundNum_ *big.Int, tokenAddress_ common.Address) (*big.Int, error) {
	return _PrizesWallet.Contract.GetDonatedTokenBalanceAmount(&_PrizesWallet.CallOpts, roundNum_, tokenAddress_)
}

// GetDonatedTokenBalanceAmount is a free data retrieval call binding the contract method 0xe86a49d7.
//
// Solidity: function getDonatedTokenBalanceAmount(uint256 roundNum_, address tokenAddress_) view returns(uint256)
func (_PrizesWallet *PrizesWalletCallerSession) GetDonatedTokenBalanceAmount(roundNum_ *big.Int, tokenAddress_ common.Address) (*big.Int, error) {
	return _PrizesWallet.Contract.GetDonatedTokenBalanceAmount(&_PrizesWallet.CallOpts, roundNum_, tokenAddress_)
}

// GetEthBalanceInfo is a free data retrieval call binding the contract method 0x1a18c889.
//
// Solidity: function getEthBalanceInfo() view returns((uint256,uint256))
func (_PrizesWallet *PrizesWalletCaller) GetEthBalanceInfo(opts *bind.CallOpts) (IPrizesWalletEthBalanceInfo, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "getEthBalanceInfo")

	if err != nil {
		return *new(IPrizesWalletEthBalanceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IPrizesWalletEthBalanceInfo)).(*IPrizesWalletEthBalanceInfo)

	return out0, err

}

// GetEthBalanceInfo is a free data retrieval call binding the contract method 0x1a18c889.
//
// Solidity: function getEthBalanceInfo() view returns((uint256,uint256))
func (_PrizesWallet *PrizesWalletSession) GetEthBalanceInfo() (IPrizesWalletEthBalanceInfo, error) {
	return _PrizesWallet.Contract.GetEthBalanceInfo(&_PrizesWallet.CallOpts)
}

// GetEthBalanceInfo is a free data retrieval call binding the contract method 0x1a18c889.
//
// Solidity: function getEthBalanceInfo() view returns((uint256,uint256))
func (_PrizesWallet *PrizesWalletCallerSession) GetEthBalanceInfo() (IPrizesWalletEthBalanceInfo, error) {
	return _PrizesWallet.Contract.GetEthBalanceInfo(&_PrizesWallet.CallOpts)
}

// GetEthBalanceInfo0 is a free data retrieval call binding the contract method 0xa089e0be.
//
// Solidity: function getEthBalanceInfo(address prizeWinnerAddress_) view returns((uint256,uint256))
func (_PrizesWallet *PrizesWalletCaller) GetEthBalanceInfo0(opts *bind.CallOpts, prizeWinnerAddress_ common.Address) (IPrizesWalletEthBalanceInfo, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "getEthBalanceInfo0", prizeWinnerAddress_)

	if err != nil {
		return *new(IPrizesWalletEthBalanceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IPrizesWalletEthBalanceInfo)).(*IPrizesWalletEthBalanceInfo)

	return out0, err

}

// GetEthBalanceInfo0 is a free data retrieval call binding the contract method 0xa089e0be.
//
// Solidity: function getEthBalanceInfo(address prizeWinnerAddress_) view returns((uint256,uint256))
func (_PrizesWallet *PrizesWalletSession) GetEthBalanceInfo0(prizeWinnerAddress_ common.Address) (IPrizesWalletEthBalanceInfo, error) {
	return _PrizesWallet.Contract.GetEthBalanceInfo0(&_PrizesWallet.CallOpts, prizeWinnerAddress_)
}

// GetEthBalanceInfo0 is a free data retrieval call binding the contract method 0xa089e0be.
//
// Solidity: function getEthBalanceInfo(address prizeWinnerAddress_) view returns((uint256,uint256))
func (_PrizesWallet *PrizesWalletCallerSession) GetEthBalanceInfo0(prizeWinnerAddress_ common.Address) (IPrizesWalletEthBalanceInfo, error) {
	return _PrizesWallet.Contract.GetEthBalanceInfo0(&_PrizesWallet.CallOpts, prizeWinnerAddress_)
}

// MainPrizeBeneficiaryAddresses is a free data retrieval call binding the contract method 0x76e42c6a.
//
// Solidity: function mainPrizeBeneficiaryAddresses(uint256 ) view returns(address)
func (_PrizesWallet *PrizesWalletCaller) MainPrizeBeneficiaryAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "mainPrizeBeneficiaryAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MainPrizeBeneficiaryAddresses is a free data retrieval call binding the contract method 0x76e42c6a.
//
// Solidity: function mainPrizeBeneficiaryAddresses(uint256 ) view returns(address)
func (_PrizesWallet *PrizesWalletSession) MainPrizeBeneficiaryAddresses(arg0 *big.Int) (common.Address, error) {
	return _PrizesWallet.Contract.MainPrizeBeneficiaryAddresses(&_PrizesWallet.CallOpts, arg0)
}

// MainPrizeBeneficiaryAddresses is a free data retrieval call binding the contract method 0x76e42c6a.
//
// Solidity: function mainPrizeBeneficiaryAddresses(uint256 ) view returns(address)
func (_PrizesWallet *PrizesWalletCallerSession) MainPrizeBeneficiaryAddresses(arg0 *big.Int) (common.Address, error) {
	return _PrizesWallet.Contract.MainPrizeBeneficiaryAddresses(&_PrizesWallet.CallOpts, arg0)
}

// NextDonatedNftIndex is a free data retrieval call binding the contract method 0x6ff1facd.
//
// Solidity: function nextDonatedNftIndex() view returns(uint256)
func (_PrizesWallet *PrizesWalletCaller) NextDonatedNftIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "nextDonatedNftIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextDonatedNftIndex is a free data retrieval call binding the contract method 0x6ff1facd.
//
// Solidity: function nextDonatedNftIndex() view returns(uint256)
func (_PrizesWallet *PrizesWalletSession) NextDonatedNftIndex() (*big.Int, error) {
	return _PrizesWallet.Contract.NextDonatedNftIndex(&_PrizesWallet.CallOpts)
}

// NextDonatedNftIndex is a free data retrieval call binding the contract method 0x6ff1facd.
//
// Solidity: function nextDonatedNftIndex() view returns(uint256)
func (_PrizesWallet *PrizesWalletCallerSession) NextDonatedNftIndex() (*big.Int, error) {
	return _PrizesWallet.Contract.NextDonatedNftIndex(&_PrizesWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PrizesWallet *PrizesWalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PrizesWallet *PrizesWalletSession) Owner() (common.Address, error) {
	return _PrizesWallet.Contract.Owner(&_PrizesWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PrizesWallet *PrizesWalletCallerSession) Owner() (common.Address, error) {
	return _PrizesWallet.Contract.Owner(&_PrizesWallet.CallOpts)
}

// RoundTimeoutTimesToWithdrawPrizes is a free data retrieval call binding the contract method 0x4b5e1b19.
//
// Solidity: function roundTimeoutTimesToWithdrawPrizes(uint256 ) view returns(uint256)
func (_PrizesWallet *PrizesWalletCaller) RoundTimeoutTimesToWithdrawPrizes(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "roundTimeoutTimesToWithdrawPrizes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundTimeoutTimesToWithdrawPrizes is a free data retrieval call binding the contract method 0x4b5e1b19.
//
// Solidity: function roundTimeoutTimesToWithdrawPrizes(uint256 ) view returns(uint256)
func (_PrizesWallet *PrizesWalletSession) RoundTimeoutTimesToWithdrawPrizes(arg0 *big.Int) (*big.Int, error) {
	return _PrizesWallet.Contract.RoundTimeoutTimesToWithdrawPrizes(&_PrizesWallet.CallOpts, arg0)
}

// RoundTimeoutTimesToWithdrawPrizes is a free data retrieval call binding the contract method 0x4b5e1b19.
//
// Solidity: function roundTimeoutTimesToWithdrawPrizes(uint256 ) view returns(uint256)
func (_PrizesWallet *PrizesWalletCallerSession) RoundTimeoutTimesToWithdrawPrizes(arg0 *big.Int) (*big.Int, error) {
	return _PrizesWallet.Contract.RoundTimeoutTimesToWithdrawPrizes(&_PrizesWallet.CallOpts, arg0)
}

// TimeoutDurationToWithdrawPrizes is a free data retrieval call binding the contract method 0x6224dd3f.
//
// Solidity: function timeoutDurationToWithdrawPrizes() view returns(uint256)
func (_PrizesWallet *PrizesWalletCaller) TimeoutDurationToWithdrawPrizes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "timeoutDurationToWithdrawPrizes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeoutDurationToWithdrawPrizes is a free data retrieval call binding the contract method 0x6224dd3f.
//
// Solidity: function timeoutDurationToWithdrawPrizes() view returns(uint256)
func (_PrizesWallet *PrizesWalletSession) TimeoutDurationToWithdrawPrizes() (*big.Int, error) {
	return _PrizesWallet.Contract.TimeoutDurationToWithdrawPrizes(&_PrizesWallet.CallOpts)
}

// TimeoutDurationToWithdrawPrizes is a free data retrieval call binding the contract method 0x6224dd3f.
//
// Solidity: function timeoutDurationToWithdrawPrizes() view returns(uint256)
func (_PrizesWallet *PrizesWalletCallerSession) TimeoutDurationToWithdrawPrizes() (*big.Int, error) {
	return _PrizesWallet.Contract.TimeoutDurationToWithdrawPrizes(&_PrizesWallet.CallOpts)
}

// ClaimDonatedNft is a paid mutator transaction binding the contract method 0x94d907fc.
//
// Solidity: function claimDonatedNft(uint256 index_) returns()
func (_PrizesWallet *PrizesWalletTransactor) ClaimDonatedNft(opts *bind.TransactOpts, index_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "claimDonatedNft", index_)
}

// ClaimDonatedNft is a paid mutator transaction binding the contract method 0x94d907fc.
//
// Solidity: function claimDonatedNft(uint256 index_) returns()
func (_PrizesWallet *PrizesWalletSession) ClaimDonatedNft(index_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.ClaimDonatedNft(&_PrizesWallet.TransactOpts, index_)
}

// ClaimDonatedNft is a paid mutator transaction binding the contract method 0x94d907fc.
//
// Solidity: function claimDonatedNft(uint256 index_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) ClaimDonatedNft(index_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.ClaimDonatedNft(&_PrizesWallet.TransactOpts, index_)
}

// ClaimDonatedToken is a paid mutator transaction binding the contract method 0x9523b237.
//
// Solidity: function claimDonatedToken(uint256 roundNum_, address tokenAddress_, uint256 amount_) returns()
func (_PrizesWallet *PrizesWalletTransactor) ClaimDonatedToken(opts *bind.TransactOpts, roundNum_ *big.Int, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "claimDonatedToken", roundNum_, tokenAddress_, amount_)
}

// ClaimDonatedToken is a paid mutator transaction binding the contract method 0x9523b237.
//
// Solidity: function claimDonatedToken(uint256 roundNum_, address tokenAddress_, uint256 amount_) returns()
func (_PrizesWallet *PrizesWalletSession) ClaimDonatedToken(roundNum_ *big.Int, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.ClaimDonatedToken(&_PrizesWallet.TransactOpts, roundNum_, tokenAddress_, amount_)
}

// ClaimDonatedToken is a paid mutator transaction binding the contract method 0x9523b237.
//
// Solidity: function claimDonatedToken(uint256 roundNum_, address tokenAddress_, uint256 amount_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) ClaimDonatedToken(roundNum_ *big.Int, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.ClaimDonatedToken(&_PrizesWallet.TransactOpts, roundNum_, tokenAddress_, amount_)
}

// ClaimManyDonatedNfts is a paid mutator transaction binding the contract method 0xe4a6c2a4.
//
// Solidity: function claimManyDonatedNfts(uint256[] indexes_) returns()
func (_PrizesWallet *PrizesWalletTransactor) ClaimManyDonatedNfts(opts *bind.TransactOpts, indexes_ []*big.Int) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "claimManyDonatedNfts", indexes_)
}

// ClaimManyDonatedNfts is a paid mutator transaction binding the contract method 0xe4a6c2a4.
//
// Solidity: function claimManyDonatedNfts(uint256[] indexes_) returns()
func (_PrizesWallet *PrizesWalletSession) ClaimManyDonatedNfts(indexes_ []*big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.ClaimManyDonatedNfts(&_PrizesWallet.TransactOpts, indexes_)
}

// ClaimManyDonatedNfts is a paid mutator transaction binding the contract method 0xe4a6c2a4.
//
// Solidity: function claimManyDonatedNfts(uint256[] indexes_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) ClaimManyDonatedNfts(indexes_ []*big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.ClaimManyDonatedNfts(&_PrizesWallet.TransactOpts, indexes_)
}

// ClaimManyDonatedTokens is a paid mutator transaction binding the contract method 0x5e00aaba.
//
// Solidity: function claimManyDonatedTokens((uint256,address,uint256)[] donatedTokensToClaim_) returns()
func (_PrizesWallet *PrizesWalletTransactor) ClaimManyDonatedTokens(opts *bind.TransactOpts, donatedTokensToClaim_ []IPrizesWalletDonatedTokenToClaim) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "claimManyDonatedTokens", donatedTokensToClaim_)
}

// ClaimManyDonatedTokens is a paid mutator transaction binding the contract method 0x5e00aaba.
//
// Solidity: function claimManyDonatedTokens((uint256,address,uint256)[] donatedTokensToClaim_) returns()
func (_PrizesWallet *PrizesWalletSession) ClaimManyDonatedTokens(donatedTokensToClaim_ []IPrizesWalletDonatedTokenToClaim) (*types.Transaction, error) {
	return _PrizesWallet.Contract.ClaimManyDonatedTokens(&_PrizesWallet.TransactOpts, donatedTokensToClaim_)
}

// ClaimManyDonatedTokens is a paid mutator transaction binding the contract method 0x5e00aaba.
//
// Solidity: function claimManyDonatedTokens((uint256,address,uint256)[] donatedTokensToClaim_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) ClaimManyDonatedTokens(donatedTokensToClaim_ []IPrizesWalletDonatedTokenToClaim) (*types.Transaction, error) {
	return _PrizesWallet.Contract.ClaimManyDonatedTokens(&_PrizesWallet.TransactOpts, donatedTokensToClaim_)
}

// DepositEth is a paid mutator transaction binding the contract method 0xf11b35fa.
//
// Solidity: function depositEth(uint256 roundNum_, uint256 prizeWinnerIndex_, address prizeWinnerAddress_) payable returns()
func (_PrizesWallet *PrizesWalletTransactor) DepositEth(opts *bind.TransactOpts, roundNum_ *big.Int, prizeWinnerIndex_ *big.Int, prizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "depositEth", roundNum_, prizeWinnerIndex_, prizeWinnerAddress_)
}

// DepositEth is a paid mutator transaction binding the contract method 0xf11b35fa.
//
// Solidity: function depositEth(uint256 roundNum_, uint256 prizeWinnerIndex_, address prizeWinnerAddress_) payable returns()
func (_PrizesWallet *PrizesWalletSession) DepositEth(roundNum_ *big.Int, prizeWinnerIndex_ *big.Int, prizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.DepositEth(&_PrizesWallet.TransactOpts, roundNum_, prizeWinnerIndex_, prizeWinnerAddress_)
}

// DepositEth is a paid mutator transaction binding the contract method 0xf11b35fa.
//
// Solidity: function depositEth(uint256 roundNum_, uint256 prizeWinnerIndex_, address prizeWinnerAddress_) payable returns()
func (_PrizesWallet *PrizesWalletTransactorSession) DepositEth(roundNum_ *big.Int, prizeWinnerIndex_ *big.Int, prizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.DepositEth(&_PrizesWallet.TransactOpts, roundNum_, prizeWinnerIndex_, prizeWinnerAddress_)
}

// DonateNft is a paid mutator transaction binding the contract method 0xfe673fd3.
//
// Solidity: function donateNft(uint256 roundNum_, address donorAddress_, address nftAddress_, uint256 nftId_) returns()
func (_PrizesWallet *PrizesWalletTransactor) DonateNft(opts *bind.TransactOpts, roundNum_ *big.Int, donorAddress_ common.Address, nftAddress_ common.Address, nftId_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "donateNft", roundNum_, donorAddress_, nftAddress_, nftId_)
}

// DonateNft is a paid mutator transaction binding the contract method 0xfe673fd3.
//
// Solidity: function donateNft(uint256 roundNum_, address donorAddress_, address nftAddress_, uint256 nftId_) returns()
func (_PrizesWallet *PrizesWalletSession) DonateNft(roundNum_ *big.Int, donorAddress_ common.Address, nftAddress_ common.Address, nftId_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.DonateNft(&_PrizesWallet.TransactOpts, roundNum_, donorAddress_, nftAddress_, nftId_)
}

// DonateNft is a paid mutator transaction binding the contract method 0xfe673fd3.
//
// Solidity: function donateNft(uint256 roundNum_, address donorAddress_, address nftAddress_, uint256 nftId_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) DonateNft(roundNum_ *big.Int, donorAddress_ common.Address, nftAddress_ common.Address, nftId_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.DonateNft(&_PrizesWallet.TransactOpts, roundNum_, donorAddress_, nftAddress_, nftId_)
}

// DonateToken is a paid mutator transaction binding the contract method 0xe2051c7e.
//
// Solidity: function donateToken(uint256 roundNum_, address donorAddress_, address tokenAddress_, uint256 amount_) returns()
func (_PrizesWallet *PrizesWalletTransactor) DonateToken(opts *bind.TransactOpts, roundNum_ *big.Int, donorAddress_ common.Address, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "donateToken", roundNum_, donorAddress_, tokenAddress_, amount_)
}

// DonateToken is a paid mutator transaction binding the contract method 0xe2051c7e.
//
// Solidity: function donateToken(uint256 roundNum_, address donorAddress_, address tokenAddress_, uint256 amount_) returns()
func (_PrizesWallet *PrizesWalletSession) DonateToken(roundNum_ *big.Int, donorAddress_ common.Address, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.DonateToken(&_PrizesWallet.TransactOpts, roundNum_, donorAddress_, tokenAddress_, amount_)
}

// DonateToken is a paid mutator transaction binding the contract method 0xe2051c7e.
//
// Solidity: function donateToken(uint256 roundNum_, address donorAddress_, address tokenAddress_, uint256 amount_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) DonateToken(roundNum_ *big.Int, donorAddress_ common.Address, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.DonateToken(&_PrizesWallet.TransactOpts, roundNum_, donorAddress_, tokenAddress_, amount_)
}

// RegisterRoundEnd is a paid mutator transaction binding the contract method 0xcc5810d8.
//
// Solidity: function registerRoundEnd(uint256 roundNum_, address mainPrizeBeneficiaryAddress_) returns(uint256)
func (_PrizesWallet *PrizesWalletTransactor) RegisterRoundEnd(opts *bind.TransactOpts, roundNum_ *big.Int, mainPrizeBeneficiaryAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "registerRoundEnd", roundNum_, mainPrizeBeneficiaryAddress_)
}

// RegisterRoundEnd is a paid mutator transaction binding the contract method 0xcc5810d8.
//
// Solidity: function registerRoundEnd(uint256 roundNum_, address mainPrizeBeneficiaryAddress_) returns(uint256)
func (_PrizesWallet *PrizesWalletSession) RegisterRoundEnd(roundNum_ *big.Int, mainPrizeBeneficiaryAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.RegisterRoundEnd(&_PrizesWallet.TransactOpts, roundNum_, mainPrizeBeneficiaryAddress_)
}

// RegisterRoundEnd is a paid mutator transaction binding the contract method 0xcc5810d8.
//
// Solidity: function registerRoundEnd(uint256 roundNum_, address mainPrizeBeneficiaryAddress_) returns(uint256)
func (_PrizesWallet *PrizesWalletTransactorSession) RegisterRoundEnd(roundNum_ *big.Int, mainPrizeBeneficiaryAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.RegisterRoundEnd(&_PrizesWallet.TransactOpts, roundNum_, mainPrizeBeneficiaryAddress_)
}

// RegisterRoundEndAndDepositEthMany is a paid mutator transaction binding the contract method 0x87565d14.
//
// Solidity: function registerRoundEndAndDepositEthMany(uint256 roundNum_, address mainPrizeBeneficiaryAddress_, (address,uint256)[] ethDeposits_) payable returns(uint256)
func (_PrizesWallet *PrizesWalletTransactor) RegisterRoundEndAndDepositEthMany(opts *bind.TransactOpts, roundNum_ *big.Int, mainPrizeBeneficiaryAddress_ common.Address, ethDeposits_ []IPrizesWalletEthDeposit) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "registerRoundEndAndDepositEthMany", roundNum_, mainPrizeBeneficiaryAddress_, ethDeposits_)
}

// RegisterRoundEndAndDepositEthMany is a paid mutator transaction binding the contract method 0x87565d14.
//
// Solidity: function registerRoundEndAndDepositEthMany(uint256 roundNum_, address mainPrizeBeneficiaryAddress_, (address,uint256)[] ethDeposits_) payable returns(uint256)
func (_PrizesWallet *PrizesWalletSession) RegisterRoundEndAndDepositEthMany(roundNum_ *big.Int, mainPrizeBeneficiaryAddress_ common.Address, ethDeposits_ []IPrizesWalletEthDeposit) (*types.Transaction, error) {
	return _PrizesWallet.Contract.RegisterRoundEndAndDepositEthMany(&_PrizesWallet.TransactOpts, roundNum_, mainPrizeBeneficiaryAddress_, ethDeposits_)
}

// RegisterRoundEndAndDepositEthMany is a paid mutator transaction binding the contract method 0x87565d14.
//
// Solidity: function registerRoundEndAndDepositEthMany(uint256 roundNum_, address mainPrizeBeneficiaryAddress_, (address,uint256)[] ethDeposits_) payable returns(uint256)
func (_PrizesWallet *PrizesWalletTransactorSession) RegisterRoundEndAndDepositEthMany(roundNum_ *big.Int, mainPrizeBeneficiaryAddress_ common.Address, ethDeposits_ []IPrizesWalletEthDeposit) (*types.Transaction, error) {
	return _PrizesWallet.Contract.RegisterRoundEndAndDepositEthMany(&_PrizesWallet.TransactOpts, roundNum_, mainPrizeBeneficiaryAddress_, ethDeposits_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PrizesWallet *PrizesWalletTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PrizesWallet *PrizesWalletSession) RenounceOwnership() (*types.Transaction, error) {
	return _PrizesWallet.Contract.RenounceOwnership(&_PrizesWallet.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PrizesWallet *PrizesWalletTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PrizesWallet.Contract.RenounceOwnership(&_PrizesWallet.TransactOpts)
}

// SetTimeoutDurationToWithdrawPrizes is a paid mutator transaction binding the contract method 0x9e2842a8.
//
// Solidity: function setTimeoutDurationToWithdrawPrizes(uint256 newValue_) returns()
func (_PrizesWallet *PrizesWalletTransactor) SetTimeoutDurationToWithdrawPrizes(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "setTimeoutDurationToWithdrawPrizes", newValue_)
}

// SetTimeoutDurationToWithdrawPrizes is a paid mutator transaction binding the contract method 0x9e2842a8.
//
// Solidity: function setTimeoutDurationToWithdrawPrizes(uint256 newValue_) returns()
func (_PrizesWallet *PrizesWalletSession) SetTimeoutDurationToWithdrawPrizes(newValue_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.SetTimeoutDurationToWithdrawPrizes(&_PrizesWallet.TransactOpts, newValue_)
}

// SetTimeoutDurationToWithdrawPrizes is a paid mutator transaction binding the contract method 0x9e2842a8.
//
// Solidity: function setTimeoutDurationToWithdrawPrizes(uint256 newValue_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) SetTimeoutDurationToWithdrawPrizes(newValue_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.SetTimeoutDurationToWithdrawPrizes(&_PrizesWallet.TransactOpts, newValue_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PrizesWallet *PrizesWalletTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PrizesWallet *PrizesWalletSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.TransferOwnership(&_PrizesWallet.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.TransferOwnership(&_PrizesWallet.TransactOpts, newOwner)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x25e16063.
//
// Solidity: function withdrawEth(address prizeWinnerAddress_) returns()
func (_PrizesWallet *PrizesWalletTransactor) WithdrawEth(opts *bind.TransactOpts, prizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "withdrawEth", prizeWinnerAddress_)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x25e16063.
//
// Solidity: function withdrawEth(address prizeWinnerAddress_) returns()
func (_PrizesWallet *PrizesWalletSession) WithdrawEth(prizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.WithdrawEth(&_PrizesWallet.TransactOpts, prizeWinnerAddress_)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x25e16063.
//
// Solidity: function withdrawEth(address prizeWinnerAddress_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) WithdrawEth(prizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.WithdrawEth(&_PrizesWallet.TransactOpts, prizeWinnerAddress_)
}

// WithdrawEth0 is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_PrizesWallet *PrizesWalletTransactor) WithdrawEth0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "withdrawEth0")
}

// WithdrawEth0 is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_PrizesWallet *PrizesWalletSession) WithdrawEth0() (*types.Transaction, error) {
	return _PrizesWallet.Contract.WithdrawEth0(&_PrizesWallet.TransactOpts)
}

// WithdrawEth0 is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_PrizesWallet *PrizesWalletTransactorSession) WithdrawEth0() (*types.Transaction, error) {
	return _PrizesWallet.Contract.WithdrawEth0(&_PrizesWallet.TransactOpts)
}

// WithdrawEverything is a paid mutator transaction binding the contract method 0xb4e6f803.
//
// Solidity: function withdrawEverything(bool withdrawEth_, (uint256,address,uint256)[] donatedTokensToClaim_, uint256[] donatedNftIndexes_) returns()
func (_PrizesWallet *PrizesWalletTransactor) WithdrawEverything(opts *bind.TransactOpts, withdrawEth_ bool, donatedTokensToClaim_ []IPrizesWalletDonatedTokenToClaim, donatedNftIndexes_ []*big.Int) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "withdrawEverything", withdrawEth_, donatedTokensToClaim_, donatedNftIndexes_)
}

// WithdrawEverything is a paid mutator transaction binding the contract method 0xb4e6f803.
//
// Solidity: function withdrawEverything(bool withdrawEth_, (uint256,address,uint256)[] donatedTokensToClaim_, uint256[] donatedNftIndexes_) returns()
func (_PrizesWallet *PrizesWalletSession) WithdrawEverything(withdrawEth_ bool, donatedTokensToClaim_ []IPrizesWalletDonatedTokenToClaim, donatedNftIndexes_ []*big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.WithdrawEverything(&_PrizesWallet.TransactOpts, withdrawEth_, donatedTokensToClaim_, donatedNftIndexes_)
}

// WithdrawEverything is a paid mutator transaction binding the contract method 0xb4e6f803.
//
// Solidity: function withdrawEverything(bool withdrawEth_, (uint256,address,uint256)[] donatedTokensToClaim_, uint256[] donatedNftIndexes_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) WithdrawEverything(withdrawEth_ bool, donatedTokensToClaim_ []IPrizesWalletDonatedTokenToClaim, donatedNftIndexes_ []*big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.WithdrawEverything(&_PrizesWallet.TransactOpts, withdrawEth_, donatedTokensToClaim_, donatedNftIndexes_)
}

// PrizesWalletDonatedNftClaimedIterator is returned from FilterDonatedNftClaimed and is used to iterate over the raw logs and unpacked data for DonatedNftClaimed events raised by the PrizesWallet contract.
type PrizesWalletDonatedNftClaimedIterator struct {
	Event *PrizesWalletDonatedNftClaimed // Event containing the contract specifics and raw log

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
func (it *PrizesWalletDonatedNftClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizesWalletDonatedNftClaimed)
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
		it.Event = new(PrizesWalletDonatedNftClaimed)
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
func (it *PrizesWalletDonatedNftClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizesWalletDonatedNftClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizesWalletDonatedNftClaimed represents a DonatedNftClaimed event raised by the PrizesWallet contract.
type PrizesWalletDonatedNftClaimed struct {
	RoundNum           *big.Int
	BeneficiaryAddress common.Address
	NftAddress         common.Address
	NftId              *big.Int
	Index              *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDonatedNftClaimed is a free log retrieval operation binding the contract event 0x03c2b6e01c9ca39e4073132f279b61b987d41a82b08cea3dd4a2fabc37067ac3.
//
// Solidity: event DonatedNftClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, address indexed nftAddress, uint256 nftId, uint256 index)
func (_PrizesWallet *PrizesWalletFilterer) FilterDonatedNftClaimed(opts *bind.FilterOpts, roundNum []*big.Int, beneficiaryAddress []common.Address, nftAddress []common.Address) (*PrizesWalletDonatedNftClaimedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.FilterLogs(opts, "DonatedNftClaimed", roundNumRule, beneficiaryAddressRule, nftAddressRule)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletDonatedNftClaimedIterator{contract: _PrizesWallet.contract, event: "DonatedNftClaimed", logs: logs, sub: sub}, nil
}

// WatchDonatedNftClaimed is a free log subscription operation binding the contract event 0x03c2b6e01c9ca39e4073132f279b61b987d41a82b08cea3dd4a2fabc37067ac3.
//
// Solidity: event DonatedNftClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, address indexed nftAddress, uint256 nftId, uint256 index)
func (_PrizesWallet *PrizesWalletFilterer) WatchDonatedNftClaimed(opts *bind.WatchOpts, sink chan<- *PrizesWalletDonatedNftClaimed, roundNum []*big.Int, beneficiaryAddress []common.Address, nftAddress []common.Address) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.WatchLogs(opts, "DonatedNftClaimed", roundNumRule, beneficiaryAddressRule, nftAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizesWalletDonatedNftClaimed)
				if err := _PrizesWallet.contract.UnpackLog(event, "DonatedNftClaimed", log); err != nil {
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

// ParseDonatedNftClaimed is a log parse operation binding the contract event 0x03c2b6e01c9ca39e4073132f279b61b987d41a82b08cea3dd4a2fabc37067ac3.
//
// Solidity: event DonatedNftClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, address indexed nftAddress, uint256 nftId, uint256 index)
func (_PrizesWallet *PrizesWalletFilterer) ParseDonatedNftClaimed(log types.Log) (*PrizesWalletDonatedNftClaimed, error) {
	event := new(PrizesWalletDonatedNftClaimed)
	if err := _PrizesWallet.contract.UnpackLog(event, "DonatedNftClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizesWalletDonatedTokenClaimedIterator is returned from FilterDonatedTokenClaimed and is used to iterate over the raw logs and unpacked data for DonatedTokenClaimed events raised by the PrizesWallet contract.
type PrizesWalletDonatedTokenClaimedIterator struct {
	Event *PrizesWalletDonatedTokenClaimed // Event containing the contract specifics and raw log

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
func (it *PrizesWalletDonatedTokenClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizesWalletDonatedTokenClaimed)
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
		it.Event = new(PrizesWalletDonatedTokenClaimed)
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
func (it *PrizesWalletDonatedTokenClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizesWalletDonatedTokenClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizesWalletDonatedTokenClaimed represents a DonatedTokenClaimed event raised by the PrizesWallet contract.
type PrizesWalletDonatedTokenClaimed struct {
	RoundNum           *big.Int
	BeneficiaryAddress common.Address
	TokenAddress       common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDonatedTokenClaimed is a free log retrieval operation binding the contract event 0xaf1adae2e1e983ca738335dc2e37194114142793f394ac934a45ea632a8a5bd0.
//
// Solidity: event DonatedTokenClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, address indexed tokenAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) FilterDonatedTokenClaimed(opts *bind.FilterOpts, roundNum []*big.Int, beneficiaryAddress []common.Address, tokenAddress []common.Address) (*PrizesWalletDonatedTokenClaimedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.FilterLogs(opts, "DonatedTokenClaimed", roundNumRule, beneficiaryAddressRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletDonatedTokenClaimedIterator{contract: _PrizesWallet.contract, event: "DonatedTokenClaimed", logs: logs, sub: sub}, nil
}

// WatchDonatedTokenClaimed is a free log subscription operation binding the contract event 0xaf1adae2e1e983ca738335dc2e37194114142793f394ac934a45ea632a8a5bd0.
//
// Solidity: event DonatedTokenClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, address indexed tokenAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) WatchDonatedTokenClaimed(opts *bind.WatchOpts, sink chan<- *PrizesWalletDonatedTokenClaimed, roundNum []*big.Int, beneficiaryAddress []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.WatchLogs(opts, "DonatedTokenClaimed", roundNumRule, beneficiaryAddressRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizesWalletDonatedTokenClaimed)
				if err := _PrizesWallet.contract.UnpackLog(event, "DonatedTokenClaimed", log); err != nil {
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

// ParseDonatedTokenClaimed is a log parse operation binding the contract event 0xaf1adae2e1e983ca738335dc2e37194114142793f394ac934a45ea632a8a5bd0.
//
// Solidity: event DonatedTokenClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, address indexed tokenAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) ParseDonatedTokenClaimed(log types.Log) (*PrizesWalletDonatedTokenClaimed, error) {
	event := new(PrizesWalletDonatedTokenClaimed)
	if err := _PrizesWallet.contract.UnpackLog(event, "DonatedTokenClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizesWalletEthReceivedIterator is returned from FilterEthReceived and is used to iterate over the raw logs and unpacked data for EthReceived events raised by the PrizesWallet contract.
type PrizesWalletEthReceivedIterator struct {
	Event *PrizesWalletEthReceived // Event containing the contract specifics and raw log

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
func (it *PrizesWalletEthReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizesWalletEthReceived)
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
		it.Event = new(PrizesWalletEthReceived)
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
func (it *PrizesWalletEthReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizesWalletEthReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizesWalletEthReceived represents a EthReceived event raised by the PrizesWallet contract.
type PrizesWalletEthReceived struct {
	RoundNum           *big.Int
	PrizeWinnerIndex   *big.Int
	PrizeWinnerAddress common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterEthReceived is a free log retrieval operation binding the contract event 0x8e369548f9ef214c7c1823c098a23763321fd761ac5cf78958e1db1b3648e7a2.
//
// Solidity: event EthReceived(uint256 indexed roundNum, uint256 prizeWinnerIndex, address indexed prizeWinnerAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) FilterEthReceived(opts *bind.FilterOpts, roundNum []*big.Int, prizeWinnerAddress []common.Address) (*PrizesWalletEthReceivedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	var prizeWinnerAddressRule []interface{}
	for _, prizeWinnerAddressItem := range prizeWinnerAddress {
		prizeWinnerAddressRule = append(prizeWinnerAddressRule, prizeWinnerAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.FilterLogs(opts, "EthReceived", roundNumRule, prizeWinnerAddressRule)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletEthReceivedIterator{contract: _PrizesWallet.contract, event: "EthReceived", logs: logs, sub: sub}, nil
}

// WatchEthReceived is a free log subscription operation binding the contract event 0x8e369548f9ef214c7c1823c098a23763321fd761ac5cf78958e1db1b3648e7a2.
//
// Solidity: event EthReceived(uint256 indexed roundNum, uint256 prizeWinnerIndex, address indexed prizeWinnerAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) WatchEthReceived(opts *bind.WatchOpts, sink chan<- *PrizesWalletEthReceived, roundNum []*big.Int, prizeWinnerAddress []common.Address) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	var prizeWinnerAddressRule []interface{}
	for _, prizeWinnerAddressItem := range prizeWinnerAddress {
		prizeWinnerAddressRule = append(prizeWinnerAddressRule, prizeWinnerAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.WatchLogs(opts, "EthReceived", roundNumRule, prizeWinnerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizesWalletEthReceived)
				if err := _PrizesWallet.contract.UnpackLog(event, "EthReceived", log); err != nil {
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

// ParseEthReceived is a log parse operation binding the contract event 0x8e369548f9ef214c7c1823c098a23763321fd761ac5cf78958e1db1b3648e7a2.
//
// Solidity: event EthReceived(uint256 indexed roundNum, uint256 prizeWinnerIndex, address indexed prizeWinnerAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) ParseEthReceived(log types.Log) (*PrizesWalletEthReceived, error) {
	event := new(PrizesWalletEthReceived)
	if err := _PrizesWallet.contract.UnpackLog(event, "EthReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizesWalletEthWithdrawnIterator is returned from FilterEthWithdrawn and is used to iterate over the raw logs and unpacked data for EthWithdrawn events raised by the PrizesWallet contract.
type PrizesWalletEthWithdrawnIterator struct {
	Event *PrizesWalletEthWithdrawn // Event containing the contract specifics and raw log

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
func (it *PrizesWalletEthWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizesWalletEthWithdrawn)
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
		it.Event = new(PrizesWalletEthWithdrawn)
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
func (it *PrizesWalletEthWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizesWalletEthWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizesWalletEthWithdrawn represents a EthWithdrawn event raised by the PrizesWallet contract.
type PrizesWalletEthWithdrawn struct {
	PrizeWinnerAddress common.Address
	BeneficiaryAddress common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterEthWithdrawn is a free log retrieval operation binding the contract event 0x4f43b861ba36494acfe938f3815fba7fac6981bdc611b6ccdc14c08f59292383.
//
// Solidity: event EthWithdrawn(address indexed prizeWinnerAddress, address indexed beneficiaryAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) FilterEthWithdrawn(opts *bind.FilterOpts, prizeWinnerAddress []common.Address, beneficiaryAddress []common.Address) (*PrizesWalletEthWithdrawnIterator, error) {

	var prizeWinnerAddressRule []interface{}
	for _, prizeWinnerAddressItem := range prizeWinnerAddress {
		prizeWinnerAddressRule = append(prizeWinnerAddressRule, prizeWinnerAddressItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.FilterLogs(opts, "EthWithdrawn", prizeWinnerAddressRule, beneficiaryAddressRule)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletEthWithdrawnIterator{contract: _PrizesWallet.contract, event: "EthWithdrawn", logs: logs, sub: sub}, nil
}

// WatchEthWithdrawn is a free log subscription operation binding the contract event 0x4f43b861ba36494acfe938f3815fba7fac6981bdc611b6ccdc14c08f59292383.
//
// Solidity: event EthWithdrawn(address indexed prizeWinnerAddress, address indexed beneficiaryAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) WatchEthWithdrawn(opts *bind.WatchOpts, sink chan<- *PrizesWalletEthWithdrawn, prizeWinnerAddress []common.Address, beneficiaryAddress []common.Address) (event.Subscription, error) {

	var prizeWinnerAddressRule []interface{}
	for _, prizeWinnerAddressItem := range prizeWinnerAddress {
		prizeWinnerAddressRule = append(prizeWinnerAddressRule, prizeWinnerAddressItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.WatchLogs(opts, "EthWithdrawn", prizeWinnerAddressRule, beneficiaryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizesWalletEthWithdrawn)
				if err := _PrizesWallet.contract.UnpackLog(event, "EthWithdrawn", log); err != nil {
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

// ParseEthWithdrawn is a log parse operation binding the contract event 0x4f43b861ba36494acfe938f3815fba7fac6981bdc611b6ccdc14c08f59292383.
//
// Solidity: event EthWithdrawn(address indexed prizeWinnerAddress, address indexed beneficiaryAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) ParseEthWithdrawn(log types.Log) (*PrizesWalletEthWithdrawn, error) {
	event := new(PrizesWalletEthWithdrawn)
	if err := _PrizesWallet.contract.UnpackLog(event, "EthWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizesWalletNftDonatedIterator is returned from FilterNftDonated and is used to iterate over the raw logs and unpacked data for NftDonated events raised by the PrizesWallet contract.
type PrizesWalletNftDonatedIterator struct {
	Event *PrizesWalletNftDonated // Event containing the contract specifics and raw log

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
func (it *PrizesWalletNftDonatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizesWalletNftDonated)
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
		it.Event = new(PrizesWalletNftDonated)
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
func (it *PrizesWalletNftDonatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizesWalletNftDonatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizesWalletNftDonated represents a NftDonated event raised by the PrizesWallet contract.
type PrizesWalletNftDonated struct {
	RoundNum     *big.Int
	DonorAddress common.Address
	NftAddress   common.Address
	NftId        *big.Int
	Index        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNftDonated is a free log retrieval operation binding the contract event 0xb12e72bab0c2a8fe98175a3c02792645ebdf65d37cf3983517e6d1c5ab842e23.
//
// Solidity: event NftDonated(uint256 indexed roundNum, address indexed donorAddress, address indexed nftAddress, uint256 nftId, uint256 index)
func (_PrizesWallet *PrizesWalletFilterer) FilterNftDonated(opts *bind.FilterOpts, roundNum []*big.Int, donorAddress []common.Address, nftAddress []common.Address) (*PrizesWalletNftDonatedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.FilterLogs(opts, "NftDonated", roundNumRule, donorAddressRule, nftAddressRule)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletNftDonatedIterator{contract: _PrizesWallet.contract, event: "NftDonated", logs: logs, sub: sub}, nil
}

// WatchNftDonated is a free log subscription operation binding the contract event 0xb12e72bab0c2a8fe98175a3c02792645ebdf65d37cf3983517e6d1c5ab842e23.
//
// Solidity: event NftDonated(uint256 indexed roundNum, address indexed donorAddress, address indexed nftAddress, uint256 nftId, uint256 index)
func (_PrizesWallet *PrizesWalletFilterer) WatchNftDonated(opts *bind.WatchOpts, sink chan<- *PrizesWalletNftDonated, roundNum []*big.Int, donorAddress []common.Address, nftAddress []common.Address) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.WatchLogs(opts, "NftDonated", roundNumRule, donorAddressRule, nftAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizesWalletNftDonated)
				if err := _PrizesWallet.contract.UnpackLog(event, "NftDonated", log); err != nil {
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

// ParseNftDonated is a log parse operation binding the contract event 0xb12e72bab0c2a8fe98175a3c02792645ebdf65d37cf3983517e6d1c5ab842e23.
//
// Solidity: event NftDonated(uint256 indexed roundNum, address indexed donorAddress, address indexed nftAddress, uint256 nftId, uint256 index)
func (_PrizesWallet *PrizesWalletFilterer) ParseNftDonated(log types.Log) (*PrizesWalletNftDonated, error) {
	event := new(PrizesWalletNftDonated)
	if err := _PrizesWallet.contract.UnpackLog(event, "NftDonated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizesWalletOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PrizesWallet contract.
type PrizesWalletOwnershipTransferredIterator struct {
	Event *PrizesWalletOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PrizesWalletOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizesWalletOwnershipTransferred)
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
		it.Event = new(PrizesWalletOwnershipTransferred)
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
func (it *PrizesWalletOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizesWalletOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizesWalletOwnershipTransferred represents a OwnershipTransferred event raised by the PrizesWallet contract.
type PrizesWalletOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PrizesWallet *PrizesWalletFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PrizesWalletOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PrizesWallet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletOwnershipTransferredIterator{contract: _PrizesWallet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PrizesWallet *PrizesWalletFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PrizesWalletOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PrizesWallet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizesWalletOwnershipTransferred)
				if err := _PrizesWallet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PrizesWallet *PrizesWalletFilterer) ParseOwnershipTransferred(log types.Log) (*PrizesWalletOwnershipTransferred, error) {
	event := new(PrizesWalletOwnershipTransferred)
	if err := _PrizesWallet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizesWalletTimeoutDurationToWithdrawPrizesChangedIterator is returned from FilterTimeoutDurationToWithdrawPrizesChanged and is used to iterate over the raw logs and unpacked data for TimeoutDurationToWithdrawPrizesChanged events raised by the PrizesWallet contract.
type PrizesWalletTimeoutDurationToWithdrawPrizesChangedIterator struct {
	Event *PrizesWalletTimeoutDurationToWithdrawPrizesChanged // Event containing the contract specifics and raw log

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
func (it *PrizesWalletTimeoutDurationToWithdrawPrizesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizesWalletTimeoutDurationToWithdrawPrizesChanged)
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
		it.Event = new(PrizesWalletTimeoutDurationToWithdrawPrizesChanged)
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
func (it *PrizesWalletTimeoutDurationToWithdrawPrizesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizesWalletTimeoutDurationToWithdrawPrizesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizesWalletTimeoutDurationToWithdrawPrizesChanged represents a TimeoutDurationToWithdrawPrizesChanged event raised by the PrizesWallet contract.
type PrizesWalletTimeoutDurationToWithdrawPrizesChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTimeoutDurationToWithdrawPrizesChanged is a free log retrieval operation binding the contract event 0x8717bb199c6bc4a5dadb21547205f9ef8ec037dda246a5526d6a6471306ea52e.
//
// Solidity: event TimeoutDurationToWithdrawPrizesChanged(uint256 newValue)
func (_PrizesWallet *PrizesWalletFilterer) FilterTimeoutDurationToWithdrawPrizesChanged(opts *bind.FilterOpts) (*PrizesWalletTimeoutDurationToWithdrawPrizesChangedIterator, error) {

	logs, sub, err := _PrizesWallet.contract.FilterLogs(opts, "TimeoutDurationToWithdrawPrizesChanged")
	if err != nil {
		return nil, err
	}
	return &PrizesWalletTimeoutDurationToWithdrawPrizesChangedIterator{contract: _PrizesWallet.contract, event: "TimeoutDurationToWithdrawPrizesChanged", logs: logs, sub: sub}, nil
}

// WatchTimeoutDurationToWithdrawPrizesChanged is a free log subscription operation binding the contract event 0x8717bb199c6bc4a5dadb21547205f9ef8ec037dda246a5526d6a6471306ea52e.
//
// Solidity: event TimeoutDurationToWithdrawPrizesChanged(uint256 newValue)
func (_PrizesWallet *PrizesWalletFilterer) WatchTimeoutDurationToWithdrawPrizesChanged(opts *bind.WatchOpts, sink chan<- *PrizesWalletTimeoutDurationToWithdrawPrizesChanged) (event.Subscription, error) {

	logs, sub, err := _PrizesWallet.contract.WatchLogs(opts, "TimeoutDurationToWithdrawPrizesChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizesWalletTimeoutDurationToWithdrawPrizesChanged)
				if err := _PrizesWallet.contract.UnpackLog(event, "TimeoutDurationToWithdrawPrizesChanged", log); err != nil {
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

// ParseTimeoutDurationToWithdrawPrizesChanged is a log parse operation binding the contract event 0x8717bb199c6bc4a5dadb21547205f9ef8ec037dda246a5526d6a6471306ea52e.
//
// Solidity: event TimeoutDurationToWithdrawPrizesChanged(uint256 newValue)
func (_PrizesWallet *PrizesWalletFilterer) ParseTimeoutDurationToWithdrawPrizesChanged(log types.Log) (*PrizesWalletTimeoutDurationToWithdrawPrizesChanged, error) {
	event := new(PrizesWalletTimeoutDurationToWithdrawPrizesChanged)
	if err := _PrizesWallet.contract.UnpackLog(event, "TimeoutDurationToWithdrawPrizesChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizesWalletTokenDonatedIterator is returned from FilterTokenDonated and is used to iterate over the raw logs and unpacked data for TokenDonated events raised by the PrizesWallet contract.
type PrizesWalletTokenDonatedIterator struct {
	Event *PrizesWalletTokenDonated // Event containing the contract specifics and raw log

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
func (it *PrizesWalletTokenDonatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizesWalletTokenDonated)
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
		it.Event = new(PrizesWalletTokenDonated)
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
func (it *PrizesWalletTokenDonatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizesWalletTokenDonatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizesWalletTokenDonated represents a TokenDonated event raised by the PrizesWallet contract.
type PrizesWalletTokenDonated struct {
	RoundNum     *big.Int
	DonorAddress common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenDonated is a free log retrieval operation binding the contract event 0x3f94f6171c96ab491d8c268f80da5e460b148878065711a0bb099bbc297615af.
//
// Solidity: event TokenDonated(uint256 indexed roundNum, address indexed donorAddress, address indexed tokenAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) FilterTokenDonated(opts *bind.FilterOpts, roundNum []*big.Int, donorAddress []common.Address, tokenAddress []common.Address) (*PrizesWalletTokenDonatedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.FilterLogs(opts, "TokenDonated", roundNumRule, donorAddressRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletTokenDonatedIterator{contract: _PrizesWallet.contract, event: "TokenDonated", logs: logs, sub: sub}, nil
}

// WatchTokenDonated is a free log subscription operation binding the contract event 0x3f94f6171c96ab491d8c268f80da5e460b148878065711a0bb099bbc297615af.
//
// Solidity: event TokenDonated(uint256 indexed roundNum, address indexed donorAddress, address indexed tokenAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) WatchTokenDonated(opts *bind.WatchOpts, sink chan<- *PrizesWalletTokenDonated, roundNum []*big.Int, donorAddress []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.WatchLogs(opts, "TokenDonated", roundNumRule, donorAddressRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizesWalletTokenDonated)
				if err := _PrizesWallet.contract.UnpackLog(event, "TokenDonated", log); err != nil {
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

// ParseTokenDonated is a log parse operation binding the contract event 0x3f94f6171c96ab491d8c268f80da5e460b148878065711a0bb099bbc297615af.
//
// Solidity: event TokenDonated(uint256 indexed roundNum, address indexed donorAddress, address indexed tokenAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) ParseTokenDonated(log types.Log) (*PrizesWalletTokenDonated, error) {
	event := new(PrizesWalletTokenDonated)
	if err := _PrizesWallet.contract.UnpackLog(event, "TokenDonated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
