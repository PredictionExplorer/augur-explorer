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

// CosmicSignatureConstantsBalanceInfo is an auto generated low-level Go binding around an user-defined struct.
type CosmicSignatureConstantsBalanceInfo struct {
	RoundNum *big.Int
	Amount   *big.Int
}

// CosmicSignatureConstantsDonatedTokenToClaim is an auto generated low-level Go binding around an user-defined struct.
type CosmicSignatureConstantsDonatedTokenToClaim struct {
	RoundNum     *big.Int
	TokenAddress common.Address
}

// IPrizesWalletMetaData contains all meta data concerning the IPrizesWallet contract.
var IPrizesWalletMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DonatedNftClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DonatedTokenClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"roundPrizeWinnerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"roundPrizeWinnerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"NftDonated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"TimeoutDurationToWithdrawPrizesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenDonated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"}],\"name\":\"claimDonatedNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"}],\"name\":\"claimDonatedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"indices_\",\"type\":\"uint256[]\"}],\"name\":\"claimManyDonatedNfts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"internalType\":\"structCosmicSignatureConstants.DonatedTokenToClaim[]\",\"name\":\"donatedTokensToClaim_\",\"type\":\"tuple[]\"}],\"name\":\"claimManyDonatedTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"roundPrizeWinnerAddress_\",\"type\":\"address\"}],\"name\":\"depositEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"donorAddress_\",\"type\":\"address\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"donateNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"donorAddress_\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"donateToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"}],\"name\":\"getDonatedTokenAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEthBalanceInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCosmicSignatureConstants.BalanceInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"roundPrizeWinnerAddress_\",\"type\":\"address\"}],\"name\":\"getEthBalanceInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCosmicSignatureConstants.BalanceInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"roundMainPrizeWinnerAddress_\",\"type\":\"address\"}],\"name\":\"registerRoundEnd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setTimeoutDurationToWithdrawPrizes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"roundPrizeWinnerAddress_\",\"type\":\"address\"}],\"name\":\"withdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"withdrawEth_\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"internalType\":\"structCosmicSignatureConstants.DonatedTokenToClaim[]\",\"name\":\"donatedTokensToClaim_\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"donatedNftIndices_\",\"type\":\"uint256[]\"}],\"name\":\"withdrawEverything\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// GetDonatedTokenAmount is a free data retrieval call binding the contract method 0x854fdf1b.
//
// Solidity: function getDonatedTokenAmount(uint256 roundNum_, address tokenAddress_) view returns(uint256)
func (_IPrizesWallet *IPrizesWalletCaller) GetDonatedTokenAmount(opts *bind.CallOpts, roundNum_ *big.Int, tokenAddress_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IPrizesWallet.contract.Call(opts, &out, "getDonatedTokenAmount", roundNum_, tokenAddress_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDonatedTokenAmount is a free data retrieval call binding the contract method 0x854fdf1b.
//
// Solidity: function getDonatedTokenAmount(uint256 roundNum_, address tokenAddress_) view returns(uint256)
func (_IPrizesWallet *IPrizesWalletSession) GetDonatedTokenAmount(roundNum_ *big.Int, tokenAddress_ common.Address) (*big.Int, error) {
	return _IPrizesWallet.Contract.GetDonatedTokenAmount(&_IPrizesWallet.CallOpts, roundNum_, tokenAddress_)
}

// GetDonatedTokenAmount is a free data retrieval call binding the contract method 0x854fdf1b.
//
// Solidity: function getDonatedTokenAmount(uint256 roundNum_, address tokenAddress_) view returns(uint256)
func (_IPrizesWallet *IPrizesWalletCallerSession) GetDonatedTokenAmount(roundNum_ *big.Int, tokenAddress_ common.Address) (*big.Int, error) {
	return _IPrizesWallet.Contract.GetDonatedTokenAmount(&_IPrizesWallet.CallOpts, roundNum_, tokenAddress_)
}

// GetEthBalanceInfo is a free data retrieval call binding the contract method 0x1a18c889.
//
// Solidity: function getEthBalanceInfo() view returns((uint256,uint256))
func (_IPrizesWallet *IPrizesWalletCaller) GetEthBalanceInfo(opts *bind.CallOpts) (CosmicSignatureConstantsBalanceInfo, error) {
	var out []interface{}
	err := _IPrizesWallet.contract.Call(opts, &out, "getEthBalanceInfo")

	if err != nil {
		return *new(CosmicSignatureConstantsBalanceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(CosmicSignatureConstantsBalanceInfo)).(*CosmicSignatureConstantsBalanceInfo)

	return out0, err

}

// GetEthBalanceInfo is a free data retrieval call binding the contract method 0x1a18c889.
//
// Solidity: function getEthBalanceInfo() view returns((uint256,uint256))
func (_IPrizesWallet *IPrizesWalletSession) GetEthBalanceInfo() (CosmicSignatureConstantsBalanceInfo, error) {
	return _IPrizesWallet.Contract.GetEthBalanceInfo(&_IPrizesWallet.CallOpts)
}

// GetEthBalanceInfo is a free data retrieval call binding the contract method 0x1a18c889.
//
// Solidity: function getEthBalanceInfo() view returns((uint256,uint256))
func (_IPrizesWallet *IPrizesWalletCallerSession) GetEthBalanceInfo() (CosmicSignatureConstantsBalanceInfo, error) {
	return _IPrizesWallet.Contract.GetEthBalanceInfo(&_IPrizesWallet.CallOpts)
}

// GetEthBalanceInfo0 is a free data retrieval call binding the contract method 0xa089e0be.
//
// Solidity: function getEthBalanceInfo(address roundPrizeWinnerAddress_) view returns((uint256,uint256))
func (_IPrizesWallet *IPrizesWalletCaller) GetEthBalanceInfo0(opts *bind.CallOpts, roundPrizeWinnerAddress_ common.Address) (CosmicSignatureConstantsBalanceInfo, error) {
	var out []interface{}
	err := _IPrizesWallet.contract.Call(opts, &out, "getEthBalanceInfo0", roundPrizeWinnerAddress_)

	if err != nil {
		return *new(CosmicSignatureConstantsBalanceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(CosmicSignatureConstantsBalanceInfo)).(*CosmicSignatureConstantsBalanceInfo)

	return out0, err

}

// GetEthBalanceInfo0 is a free data retrieval call binding the contract method 0xa089e0be.
//
// Solidity: function getEthBalanceInfo(address roundPrizeWinnerAddress_) view returns((uint256,uint256))
func (_IPrizesWallet *IPrizesWalletSession) GetEthBalanceInfo0(roundPrizeWinnerAddress_ common.Address) (CosmicSignatureConstantsBalanceInfo, error) {
	return _IPrizesWallet.Contract.GetEthBalanceInfo0(&_IPrizesWallet.CallOpts, roundPrizeWinnerAddress_)
}

// GetEthBalanceInfo0 is a free data retrieval call binding the contract method 0xa089e0be.
//
// Solidity: function getEthBalanceInfo(address roundPrizeWinnerAddress_) view returns((uint256,uint256))
func (_IPrizesWallet *IPrizesWalletCallerSession) GetEthBalanceInfo0(roundPrizeWinnerAddress_ common.Address) (CosmicSignatureConstantsBalanceInfo, error) {
	return _IPrizesWallet.Contract.GetEthBalanceInfo0(&_IPrizesWallet.CallOpts, roundPrizeWinnerAddress_)
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

// ClaimDonatedToken is a paid mutator transaction binding the contract method 0x95de6c2c.
//
// Solidity: function claimDonatedToken(uint256 roundNum_, address tokenAddress_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) ClaimDonatedToken(opts *bind.TransactOpts, roundNum_ *big.Int, tokenAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "claimDonatedToken", roundNum_, tokenAddress_)
}

// ClaimDonatedToken is a paid mutator transaction binding the contract method 0x95de6c2c.
//
// Solidity: function claimDonatedToken(uint256 roundNum_, address tokenAddress_) returns()
func (_IPrizesWallet *IPrizesWalletSession) ClaimDonatedToken(roundNum_ *big.Int, tokenAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.ClaimDonatedToken(&_IPrizesWallet.TransactOpts, roundNum_, tokenAddress_)
}

// ClaimDonatedToken is a paid mutator transaction binding the contract method 0x95de6c2c.
//
// Solidity: function claimDonatedToken(uint256 roundNum_, address tokenAddress_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) ClaimDonatedToken(roundNum_ *big.Int, tokenAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.ClaimDonatedToken(&_IPrizesWallet.TransactOpts, roundNum_, tokenAddress_)
}

// ClaimManyDonatedNfts is a paid mutator transaction binding the contract method 0xe4a6c2a4.
//
// Solidity: function claimManyDonatedNfts(uint256[] indices_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) ClaimManyDonatedNfts(opts *bind.TransactOpts, indices_ []*big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "claimManyDonatedNfts", indices_)
}

// ClaimManyDonatedNfts is a paid mutator transaction binding the contract method 0xe4a6c2a4.
//
// Solidity: function claimManyDonatedNfts(uint256[] indices_) returns()
func (_IPrizesWallet *IPrizesWalletSession) ClaimManyDonatedNfts(indices_ []*big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.ClaimManyDonatedNfts(&_IPrizesWallet.TransactOpts, indices_)
}

// ClaimManyDonatedNfts is a paid mutator transaction binding the contract method 0xe4a6c2a4.
//
// Solidity: function claimManyDonatedNfts(uint256[] indices_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) ClaimManyDonatedNfts(indices_ []*big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.ClaimManyDonatedNfts(&_IPrizesWallet.TransactOpts, indices_)
}

// ClaimManyDonatedTokens is a paid mutator transaction binding the contract method 0x50a15deb.
//
// Solidity: function claimManyDonatedTokens((uint256,address)[] donatedTokensToClaim_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) ClaimManyDonatedTokens(opts *bind.TransactOpts, donatedTokensToClaim_ []CosmicSignatureConstantsDonatedTokenToClaim) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "claimManyDonatedTokens", donatedTokensToClaim_)
}

// ClaimManyDonatedTokens is a paid mutator transaction binding the contract method 0x50a15deb.
//
// Solidity: function claimManyDonatedTokens((uint256,address)[] donatedTokensToClaim_) returns()
func (_IPrizesWallet *IPrizesWalletSession) ClaimManyDonatedTokens(donatedTokensToClaim_ []CosmicSignatureConstantsDonatedTokenToClaim) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.ClaimManyDonatedTokens(&_IPrizesWallet.TransactOpts, donatedTokensToClaim_)
}

// ClaimManyDonatedTokens is a paid mutator transaction binding the contract method 0x50a15deb.
//
// Solidity: function claimManyDonatedTokens((uint256,address)[] donatedTokensToClaim_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) ClaimManyDonatedTokens(donatedTokensToClaim_ []CosmicSignatureConstantsDonatedTokenToClaim) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.ClaimManyDonatedTokens(&_IPrizesWallet.TransactOpts, donatedTokensToClaim_)
}

// DepositEth is a paid mutator transaction binding the contract method 0x46f6b4e1.
//
// Solidity: function depositEth(uint256 roundNum_, address roundPrizeWinnerAddress_) payable returns()
func (_IPrizesWallet *IPrizesWalletTransactor) DepositEth(opts *bind.TransactOpts, roundNum_ *big.Int, roundPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "depositEth", roundNum_, roundPrizeWinnerAddress_)
}

// DepositEth is a paid mutator transaction binding the contract method 0x46f6b4e1.
//
// Solidity: function depositEth(uint256 roundNum_, address roundPrizeWinnerAddress_) payable returns()
func (_IPrizesWallet *IPrizesWalletSession) DepositEth(roundNum_ *big.Int, roundPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.DepositEth(&_IPrizesWallet.TransactOpts, roundNum_, roundPrizeWinnerAddress_)
}

// DepositEth is a paid mutator transaction binding the contract method 0x46f6b4e1.
//
// Solidity: function depositEth(uint256 roundNum_, address roundPrizeWinnerAddress_) payable returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) DepositEth(roundNum_ *big.Int, roundPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.DepositEth(&_IPrizesWallet.TransactOpts, roundNum_, roundPrizeWinnerAddress_)
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
// Solidity: function registerRoundEnd(uint256 roundNum_, address roundMainPrizeWinnerAddress_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) RegisterRoundEnd(opts *bind.TransactOpts, roundNum_ *big.Int, roundMainPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "registerRoundEnd", roundNum_, roundMainPrizeWinnerAddress_)
}

// RegisterRoundEnd is a paid mutator transaction binding the contract method 0xcc5810d8.
//
// Solidity: function registerRoundEnd(uint256 roundNum_, address roundMainPrizeWinnerAddress_) returns()
func (_IPrizesWallet *IPrizesWalletSession) RegisterRoundEnd(roundNum_ *big.Int, roundMainPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.RegisterRoundEnd(&_IPrizesWallet.TransactOpts, roundNum_, roundMainPrizeWinnerAddress_)
}

// RegisterRoundEnd is a paid mutator transaction binding the contract method 0xcc5810d8.
//
// Solidity: function registerRoundEnd(uint256 roundNum_, address roundMainPrizeWinnerAddress_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) RegisterRoundEnd(roundNum_ *big.Int, roundMainPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.RegisterRoundEnd(&_IPrizesWallet.TransactOpts, roundNum_, roundMainPrizeWinnerAddress_)
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
// Solidity: function withdrawEth(address roundPrizeWinnerAddress_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) WithdrawEth(opts *bind.TransactOpts, roundPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "withdrawEth", roundPrizeWinnerAddress_)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x25e16063.
//
// Solidity: function withdrawEth(address roundPrizeWinnerAddress_) returns()
func (_IPrizesWallet *IPrizesWalletSession) WithdrawEth(roundPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.WithdrawEth(&_IPrizesWallet.TransactOpts, roundPrizeWinnerAddress_)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x25e16063.
//
// Solidity: function withdrawEth(address roundPrizeWinnerAddress_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) WithdrawEth(roundPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.WithdrawEth(&_IPrizesWallet.TransactOpts, roundPrizeWinnerAddress_)
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

// WithdrawEverything is a paid mutator transaction binding the contract method 0xa72be1b2.
//
// Solidity: function withdrawEverything(bool withdrawEth_, (uint256,address)[] donatedTokensToClaim_, uint256[] donatedNftIndices_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) WithdrawEverything(opts *bind.TransactOpts, withdrawEth_ bool, donatedTokensToClaim_ []CosmicSignatureConstantsDonatedTokenToClaim, donatedNftIndices_ []*big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "withdrawEverything", withdrawEth_, donatedTokensToClaim_, donatedNftIndices_)
}

// WithdrawEverything is a paid mutator transaction binding the contract method 0xa72be1b2.
//
// Solidity: function withdrawEverything(bool withdrawEth_, (uint256,address)[] donatedTokensToClaim_, uint256[] donatedNftIndices_) returns()
func (_IPrizesWallet *IPrizesWalletSession) WithdrawEverything(withdrawEth_ bool, donatedTokensToClaim_ []CosmicSignatureConstantsDonatedTokenToClaim, donatedNftIndices_ []*big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.WithdrawEverything(&_IPrizesWallet.TransactOpts, withdrawEth_, donatedTokensToClaim_, donatedNftIndices_)
}

// WithdrawEverything is a paid mutator transaction binding the contract method 0xa72be1b2.
//
// Solidity: function withdrawEverything(bool withdrawEth_, (uint256,address)[] donatedTokensToClaim_, uint256[] donatedNftIndices_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) WithdrawEverything(withdrawEth_ bool, donatedTokensToClaim_ []CosmicSignatureConstantsDonatedTokenToClaim, donatedNftIndices_ []*big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.WithdrawEverything(&_IPrizesWallet.TransactOpts, withdrawEth_, donatedTokensToClaim_, donatedNftIndices_)
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
// Solidity: event DonatedNftClaimed(uint256 indexed roundNum, address beneficiaryAddress, address nftAddress, uint256 nftId, uint256 index)
func (_IPrizesWallet *IPrizesWalletFilterer) FilterDonatedNftClaimed(opts *bind.FilterOpts, roundNum []*big.Int) (*IPrizesWalletDonatedNftClaimedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _IPrizesWallet.contract.FilterLogs(opts, "DonatedNftClaimed", roundNumRule)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletDonatedNftClaimedIterator{contract: _IPrizesWallet.contract, event: "DonatedNftClaimed", logs: logs, sub: sub}, nil
}

// WatchDonatedNftClaimed is a free log subscription operation binding the contract event 0x03c2b6e01c9ca39e4073132f279b61b987d41a82b08cea3dd4a2fabc37067ac3.
//
// Solidity: event DonatedNftClaimed(uint256 indexed roundNum, address beneficiaryAddress, address nftAddress, uint256 nftId, uint256 index)
func (_IPrizesWallet *IPrizesWalletFilterer) WatchDonatedNftClaimed(opts *bind.WatchOpts, sink chan<- *IPrizesWalletDonatedNftClaimed, roundNum []*big.Int) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _IPrizesWallet.contract.WatchLogs(opts, "DonatedNftClaimed", roundNumRule)
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
// Solidity: event DonatedNftClaimed(uint256 indexed roundNum, address beneficiaryAddress, address nftAddress, uint256 nftId, uint256 index)
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
// Solidity: event DonatedTokenClaimed(uint256 indexed roundNum, address beneficiaryAddress, address tokenAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) FilterDonatedTokenClaimed(opts *bind.FilterOpts, roundNum []*big.Int) (*IPrizesWalletDonatedTokenClaimedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _IPrizesWallet.contract.FilterLogs(opts, "DonatedTokenClaimed", roundNumRule)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletDonatedTokenClaimedIterator{contract: _IPrizesWallet.contract, event: "DonatedTokenClaimed", logs: logs, sub: sub}, nil
}

// WatchDonatedTokenClaimed is a free log subscription operation binding the contract event 0xaf1adae2e1e983ca738335dc2e37194114142793f394ac934a45ea632a8a5bd0.
//
// Solidity: event DonatedTokenClaimed(uint256 indexed roundNum, address beneficiaryAddress, address tokenAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) WatchDonatedTokenClaimed(opts *bind.WatchOpts, sink chan<- *IPrizesWalletDonatedTokenClaimed, roundNum []*big.Int) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _IPrizesWallet.contract.WatchLogs(opts, "DonatedTokenClaimed", roundNumRule)
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
// Solidity: event DonatedTokenClaimed(uint256 indexed roundNum, address beneficiaryAddress, address tokenAddress, uint256 amount)
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
	RoundNum                *big.Int
	RoundPrizeWinnerAddress common.Address
	Amount                  *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterEthReceived is a free log retrieval operation binding the contract event 0x999946acc98c7b7dacc26921697d55abbcb1637484b0a73040f0b06287e43326.
//
// Solidity: event EthReceived(uint256 indexed roundNum, address indexed roundPrizeWinnerAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) FilterEthReceived(opts *bind.FilterOpts, roundNum []*big.Int, roundPrizeWinnerAddress []common.Address) (*IPrizesWalletEthReceivedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var roundPrizeWinnerAddressRule []interface{}
	for _, roundPrizeWinnerAddressItem := range roundPrizeWinnerAddress {
		roundPrizeWinnerAddressRule = append(roundPrizeWinnerAddressRule, roundPrizeWinnerAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.FilterLogs(opts, "EthReceived", roundNumRule, roundPrizeWinnerAddressRule)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletEthReceivedIterator{contract: _IPrizesWallet.contract, event: "EthReceived", logs: logs, sub: sub}, nil
}

// WatchEthReceived is a free log subscription operation binding the contract event 0x999946acc98c7b7dacc26921697d55abbcb1637484b0a73040f0b06287e43326.
//
// Solidity: event EthReceived(uint256 indexed roundNum, address indexed roundPrizeWinnerAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) WatchEthReceived(opts *bind.WatchOpts, sink chan<- *IPrizesWalletEthReceived, roundNum []*big.Int, roundPrizeWinnerAddress []common.Address) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var roundPrizeWinnerAddressRule []interface{}
	for _, roundPrizeWinnerAddressItem := range roundPrizeWinnerAddress {
		roundPrizeWinnerAddressRule = append(roundPrizeWinnerAddressRule, roundPrizeWinnerAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.WatchLogs(opts, "EthReceived", roundNumRule, roundPrizeWinnerAddressRule)
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

// ParseEthReceived is a log parse operation binding the contract event 0x999946acc98c7b7dacc26921697d55abbcb1637484b0a73040f0b06287e43326.
//
// Solidity: event EthReceived(uint256 indexed roundNum, address indexed roundPrizeWinnerAddress, uint256 amount)
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
	RoundPrizeWinnerAddress common.Address
	BeneficiaryAddress      common.Address
	Amount                  *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterEthWithdrawn is a free log retrieval operation binding the contract event 0x4f43b861ba36494acfe938f3815fba7fac6981bdc611b6ccdc14c08f59292383.
//
// Solidity: event EthWithdrawn(address indexed roundPrizeWinnerAddress, address indexed beneficiaryAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) FilterEthWithdrawn(opts *bind.FilterOpts, roundPrizeWinnerAddress []common.Address, beneficiaryAddress []common.Address) (*IPrizesWalletEthWithdrawnIterator, error) {

	var roundPrizeWinnerAddressRule []interface{}
	for _, roundPrizeWinnerAddressItem := range roundPrizeWinnerAddress {
		roundPrizeWinnerAddressRule = append(roundPrizeWinnerAddressRule, roundPrizeWinnerAddressItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.FilterLogs(opts, "EthWithdrawn", roundPrizeWinnerAddressRule, beneficiaryAddressRule)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletEthWithdrawnIterator{contract: _IPrizesWallet.contract, event: "EthWithdrawn", logs: logs, sub: sub}, nil
}

// WatchEthWithdrawn is a free log subscription operation binding the contract event 0x4f43b861ba36494acfe938f3815fba7fac6981bdc611b6ccdc14c08f59292383.
//
// Solidity: event EthWithdrawn(address indexed roundPrizeWinnerAddress, address indexed beneficiaryAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) WatchEthWithdrawn(opts *bind.WatchOpts, sink chan<- *IPrizesWalletEthWithdrawn, roundPrizeWinnerAddress []common.Address, beneficiaryAddress []common.Address) (event.Subscription, error) {

	var roundPrizeWinnerAddressRule []interface{}
	for _, roundPrizeWinnerAddressItem := range roundPrizeWinnerAddress {
		roundPrizeWinnerAddressRule = append(roundPrizeWinnerAddressRule, roundPrizeWinnerAddressItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.WatchLogs(opts, "EthWithdrawn", roundPrizeWinnerAddressRule, beneficiaryAddressRule)
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
// Solidity: event EthWithdrawn(address indexed roundPrizeWinnerAddress, address indexed beneficiaryAddress, uint256 amount)
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"game_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"callerAddress\",\"type\":\"address\"}],\"name\":\"CallDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DonatedNftAlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DonatedNftClaimDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"DonatedTokenClaimDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"operationPermittedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimeStamp\",\"type\":\"uint256\"}],\"name\":\"EarlyWithdrawal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundTransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"InvalidDonatedNftIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DonatedNftClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DonatedTokenClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"roundPrizeWinnerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"roundPrizeWinnerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"NftDonated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"TimeoutDurationToWithdrawPrizesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenDonated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"}],\"name\":\"claimDonatedNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"}],\"name\":\"claimDonatedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"indices_\",\"type\":\"uint256[]\"}],\"name\":\"claimManyDonatedNfts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"internalType\":\"structCosmicSignatureConstants.DonatedTokenToClaim[]\",\"name\":\"donatedTokensToClaim_\",\"type\":\"tuple[]\"}],\"name\":\"claimManyDonatedTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"roundPrizeWinnerAddress_\",\"type\":\"address\"}],\"name\":\"depositEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"donorAddress_\",\"type\":\"address\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"donateNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"donorAddress_\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"donateToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"donatedNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"donatedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"game\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"}],\"name\":\"getDonatedTokenAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEthBalanceInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCosmicSignatureConstants.BalanceInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"roundPrizeWinnerAddress_\",\"type\":\"address\"}],\"name\":\"getEthBalanceInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCosmicSignatureConstants.BalanceInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mainPrizeWinnerAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numDonatedNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"roundMainPrizeWinnerAddress_\",\"type\":\"address\"}],\"name\":\"registerRoundEnd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"roundTimeoutTimesToWithdrawPrizes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setTimeoutDurationToWithdrawPrizes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeoutDurationToWithdrawPrizes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"roundPrizeWinnerAddress_\",\"type\":\"address\"}],\"name\":\"withdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"withdrawEth_\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"internalType\":\"structCosmicSignatureConstants.DonatedTokenToClaim[]\",\"name\":\"donatedTokensToClaim_\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"donatedNftIndices_\",\"type\":\"uint256[]\"}],\"name\":\"withdrawEverything\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523461002f576100196100146100f4565b610339565b610021610034565b612d076104b98239612d0790f35b61003a565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906100669061003e565b810190811060018060401b0382111761007e57604052565b610048565b9061009661008f610034565b928361005c565b565b5f80fd5b60018060a01b031690565b6100b09061009c565b90565b6100bc816100a7565b036100c357565b5f80fd5b905051906100d4826100b3565b565b906020828203126100ef576100ec915f016100c7565b90565b610098565b6101126131c08038038061010781610083565b9283398101906100d6565b90565b90565b90565b90565b61013261012d61013792610115565b61011b565b610118565b90565b610146620d2f0061011e565b90565b5f1b90565b9061015a5f1991610149565b9181191691161790565b61017861017361017d92610118565b61011b565b610118565b90565b90565b9061019861019361019f92610164565b610180565b825461014e565b9055565b90565b6101ba6101b56101bf926101a3565b61011b565b610118565b90565b906101d76101d26101de926101a6565b610180565b825461014e565b9055565b6101f66101f16101fb926101a3565b61011b565b61009c565b90565b610207906101e2565b90565b60209181520190565b5f7f5468652070726f76696465642061646472657373206973207a65726f2e000000910152565b610247601d60209261020a565b61025081610213565b0190565b6102699060208101905f81830391015261023a565b90565b8061028761028161027c5f6101fe565b6100a7565b916100a7565b14610297576102959061032c565b565b61029f610034565b63eac0d38960e01b8152806102b660048201610254565b0390fd5b906102cb60018060a01b0391610149565b9181191691161790565b6102e96102e46102ee9261009c565b61011b565b61009c565b90565b6102fa906102d5565b90565b610306906102f1565b90565b90565b9061032161031c610328926102fd565b610309565b82546102ba565b9055565b61033790600161030c565b565b61038c906103463361038e565b61036061035161013a565b68010000000000000002610183565b6103875f7c01000000000000000200000000000000000000000200000000000000036101c2565b61026c565b565b61039790610399565b565b6103a2906103a4565b565b6103ad906103d1565b565b6103b8906100a7565b9052565b91906103cf905f602085019401906103af565b565b806103ec6103e66103e15f6101fe565b6100a7565b916100a7565b146103fc576103fa90610459565b565b61041f6104085f6101fe565b5f918291631e4fbdf760e01b8352600483016103bc565b0390fd5b5f1c90565b60018060a01b031690565b61043f61044491610423565b610428565b90565b6104519054610433565b90565b5f0190565b6104625f610447565b61046c825f61030c565b906104a061049a7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0936102fd565b916102fd565b916104a9610034565b806104b381610454565b0390a356fe60806040526004361015610013575b610e58565b61001d5f356101ac565b80631a18c889146101a757806325e16063146101a257806346f6b4e11461019d5780634b5e1b191461019857806350a15deb146101935780636224dd3f1461018e578063715018a614610189578063854fdf1b146101845780638da5cb5b1461017f578063909466221461017a57806394d907fc1461017557806395de6c2c146101705780639cf10d321461016b5780639e2842a814610166578063a089e0be14610161578063a0ef91df1461015c578063a72be1b214610157578063c3fe3e2814610152578063cc5810d81461014d578063d7f4f8be14610148578063e2051c7e14610143578063e4a6c2a41461013e578063f2fde38b14610139578063f9eba390146101345763fe673fd30361000e57610e21565b610d77565b610cdf565b610cab565b610c3e565b610bc3565b610b18565b610ae3565b610a6c565b610967565b610932565b6108ff565b6108c7565b61072e565b6106fb565b6106c6565b610666565b61060e565b61057f565b61054a565b6104ff565b610449565b610331565b6102ae565b610218565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f9103126101ca57565b6101bc565b90565b6101db906101cf565b9052565b90602080610201936101f75f8201515f8601906101d2565b01519101906101d2565b565b9190610216905f604085019401906101df565b565b34610248576102283660046101c0565b610244610233610f8d565b61023b6101b2565b91829182610203565b0390f35b6101b8565b5f80fd5b60018060a01b031690565b61026590610251565b90565b6102718161025c565b0361027857565b5f80fd5b9050359061028982610268565b565b906020828203126102a4576102a1915f0161027c565b90565b6101bc565b5f0190565b346102dc576102c66102c136600461028b565b611243565b6102ce6101b2565b806102d8816102a9565b0390f35b6101b8565b6102ea816101cf565b036102f157565b5f80fd5b90503590610302826102e1565b565b919060408382031261032c5780610320610329925f86016102f5565b9360200161027c565b90565b6101bc565b61034561033f366004610304565b906115c1565b61034d6101b2565b80610357816102a9565b0390f35b9060208282031261037457610371915f016102f5565b90565b6101bc565b634e487b7160e01b5f52603260045260245ffd5b506801000000000000000090565b90565b6103a78161038d565b8210156103c1576103b960019161039b565b910201905f90565b610379565b1c90565b90565b6103dd9060086103e293026103c6565b6103ca565b90565b906103f091546103cd565b90565b680100000000000000036104068161038d565b821015610423576104209161041a9161039e565b906103e5565b90565b5f80fd5b610430906101cf565b9052565b9190610447905f60208501940190610427565b565b346104795761047561046461045f36600461035b565b6103f3565b61046c6101b2565b91829182610434565b0390f35b6101b8565b5f80fd5b5f80fd5b5f80fd5b909182601f830112156104c45781359167ffffffffffffffff83116104bf5760200192604083028401116104ba57565b610486565b610482565b61047e565b906020828203126104fa575f82013567ffffffffffffffff81116104f5576104f1920161048a565b9091565b61024d565b6101bc565b3461052e576105186105123660046104c9565b9061160f565b6105206101b2565b8061052a816102a9565b0390f35b6101b8565b610547680100000000000000025f906103e5565b90565b3461057a5761055a3660046101c0565b610576610565610533565b61056d6101b2565b91829182610434565b0390f35b6101b8565b346105ad5761058f3660046101c0565b6105976116d5565b61059f6101b2565b806105a9816102a9565b0390f35b6101b8565b6105bb9061025c565b90565b6105c7816105b2565b036105ce57565b5f80fd5b905035906105df826105be565b565b919060408382031261060957806105fd610606925f86016102f5565b936020016105d2565b90565b6101bc565b3461063f5761063b61062a6106243660046105e1565b906116df565b6106326101b2565b91829182610434565b0390f35b6101b8565b61064d9061025c565b9052565b9190610664905f60208501940190610644565b565b34610696576106763660046101c0565b610692610681611722565b6106896101b2565b91829182610651565b0390f35b6101b8565b6106c37c01000000000000000200000000000000000000000200000000000000035f906103e5565b90565b346106f6576106d63660046101c0565b6106f26106e161069b565b6106e96101b2565b91829182610434565b0390f35b6101b8565b346107295761071361070e36600461035b565b611aa9565b61071b6101b2565b80610725816102a9565b0390f35b6101b8565b3461075d576107476107413660046105e1565b90611f39565b61074f6101b2565b80610759816102a9565b0390f35b6101b8565b506801000000000000000090565b90565b61077c81610762565b8210156107965761078e600391610770565b910201905f90565b610379565b5f1c90565b6107ac6107b19161079b565b6103ca565b90565b6107be90546107a0565b90565b60018060a01b031690565b6107d86107dd9161079b565b6107c1565b90565b6107ea90546107cc565b90565b7c01000000000000000200000000000000000000000200000000000000049061081582610762565b81101561084d5761082591610773565b506108315f82016107b4565b9161084a6002610843600185016107e0565b93016107b4565b90565b5f80fd5b90565b61086861086361086d92610251565b610851565b610251565b90565b61087990610854565b90565b61088590610870565b90565b6108919061087c565b9052565b6040906108be6108c594969593966108b460608401985f850190610427565b6020830190610888565b0190610427565b565b346108fa576108f66108e26108dd36600461035b565b6107ed565b6108ed9391936101b2565b93849384610895565b0390f35b6101b8565b3461092d5761091761091236600461035b565b6120e0565b61091f6101b2565b80610929816102a9565b0390f35b6101b8565b346109625761095e61094d61094836600461028b565b6120eb565b6109556101b2565b91829182610203565b0390f35b6101b8565b34610995576109773660046101c0565b61097f61211d565b6109876101b2565b80610991816102a9565b0390f35b6101b8565b151590565b6109a88161099a565b036109af57565b5f80fd5b905035906109c08261099f565b565b909182601f830112156109fc5781359167ffffffffffffffff83116109f75760200192602083028401116109f257565b610486565b610482565b61047e565b606081830312610a6757610a17825f83016109b3565b92602082013567ffffffffffffffff8111610a625783610a3891840161048a565b929093604082013567ffffffffffffffff8111610a5d57610a5992016109c2565b9091565b61024d565b61024d565b6101bc565b34610a9e57610a88610a7f366004610a01565b939290926121e8565b610a906101b2565b80610a9a816102a9565b0390f35b6101b8565b60018060a01b031690565b610abe906008610ac393026103c6565b610aa3565b90565b90610ad19154610aae565b90565b610ae060015f90610ac6565b90565b34610b1357610af33660046101c0565b610b0f610afe610ad4565b610b066101b2565b91829182610651565b0390f35b6101b8565b34610b4757610b31610b2b366004610304565b906122b9565b610b396101b2565b80610b43816102a9565b0390f35b6101b8565b50600160e01b90565b90565b610b6181610b4c565b821015610b7b57610b73600191610b55565b910201905f90565b610379565b7402000000000000000000000002000000000000000390610ba082610b4c565b811015610bbf57610bb55f91610bbc93610b58565b50016107b4565b90565b5f80fd5b34610bf357610bef610bde610bd936600461035b565b610b80565b610be66101b2565b91829182610434565b0390f35b6101b8565b608081830312610c3957610c0e825f83016102f5565b92610c36610c1f846020850161027c565b93610c2d81604086016105d2565b936060016102f5565b90565b6101bc565b34610c7057610c5a610c51366004610bf8565b92919091612488565b610c626101b2565b80610c6c816102a9565b0390f35b6101b8565b90602082820312610ca6575f82013567ffffffffffffffff8111610ca157610c9d92016109c2565b9091565b61024d565b6101bc565b34610cda57610cc4610cbe366004610c75565b906124af565b610ccc6101b2565b80610cd6816102a9565b0390f35b6101b8565b34610d0d57610cf7610cf236600461028b565b612574565b610cff6101b2565b80610d09816102a9565b0390f35b6101b8565b506801000000000000000090565b90565b610d2c81610d12565b821015610d4657610d3e600191610d20565b910201905f90565b610379565b6002610d5681610d12565b821015610d7357610d7091610d6a91610d23565b90610ac6565b90565b5f80fd5b34610da757610da3610d92610d8d36600461035b565b610d4b565b610d9a6101b2565b91829182610651565b0390f35b6101b8565b610db59061025c565b90565b610dc181610dac565b03610dc857565b5f80fd5b90503590610dd982610db8565b565b608081830312610e1c57610df1825f83016102f5565b92610e19610e02846020850161027c565b93610e108160408601610dcc565b936060016102f5565b90565b6101bc565b34610e5357610e3d610e34366004610ddb565b92919091612822565b610e456101b2565b80610e4f816102a9565b0390f35b6101b8565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b90610e8490610e5c565b810190811067ffffffffffffffff821117610e9e57604052565b610e66565b90610eb6610eaf6101b2565b9283610e7a565b565b610ec26040610ea3565b90565b5f90565b610ed1610eb8565b9060208083610ede610ec5565b815201610ee9610ec5565b81525050565b610ef7610ec9565b90565b610f0390610854565b90565b50600160a01b90565b90565b610f1b81610f06565b821015610f3557610f2d600291610f0f565b910201905f90565b610379565b90610f44906101cf565b9052565b90610f7f610f766001610f59610eb8565b94610f70610f685f83016107b4565b5f8801610f3a565b016107b4565b60208401610f3a565b565b610f8a90610f48565b90565b610f95610eef565b50610fbb610fb568020000000000000003610faf33610efa565b90610f12565b50610f81565b90565b90565b90565b610fd8610fd3610fdd92610fc1565b610851565b6101cf565b90565b60209181520190565b5f7f4e6f7420656e6f7567682074696d652068617320656c61707365642e00000000910152565b61101d601c602092610fe0565b61102681610fe9565b0190565b91604061105b929493611054611049606083018381035f850152611010565b966020830190610427565b0190610427565b565b15611066575050565b6110876110716101b2565b92839263496fca8f60e11b84526004840161102a565b0390fd5b1b90565b919060086110aa9102916110a45f198461108b565b9261108b565b9181191691161790565b6110c86110c36110cd926101cf565b610851565b6101cf565b90565b90565b91906110e96110e46110f1936110b4565b6110d0565b90835461108f565b9055565b5f90565b61110b916111056110f5565b916110d3565b565b61111690610870565b90565b905090565b6111295f8092611119565b0190565b6111369061111e565b90565b67ffffffffffffffff811161115757611153602091610e5c565b0190565b610e66565b9061116e61116983611139565b610ea3565b918252565b606090565b3d5f14611193576111883d61115c565b903d5f602084013e5b565b61119b611173565b90611191565b5f7f455448207769746864726177616c206661696c65642e00000000000000000000910152565b6111d56016602092610fe0565b6111de816111a1565b0190565b91604061121392949361120c611201606083018381035f8501526111c8565b966020830190610644565b0190610427565b565b1561121e575050565b61123f6112296101b2565b928392630aa7db6360e11b8452600484016111e2565b0390fd5b61135a9061126c6112666802000000000000000361126084610efa565b90610f12565b50610fbe565b906112ba6112976112916801000000000000000361128b5f87016107b4565b9061039e565b906103e5565b426112aa6112a4836101cf565b916101cf565b1015908161135c575b429161105d565b6112df5f806112cb600186016107b4565b946112d982600183016110f9565b016110f9565b33829161132a6113186113127f4f43b861ba36494acfe938f3815fba7fac6981bdc611b6ccdc14c08f592923839361110d565b9361110d565b936113216101b2565b91829182610434565b0390a35f8033836113396101b2565b90816113448161112d565b03925af1611350611178565b5090339091611215565b565b80915061137161136b5f610fc4565b916101cf565b11906112b3565b6113846113899161079b565b610aa3565b90565b6113969054611378565b90565b60407f6d6574686f642e00000000000000000000000000000000000000000000000000917f4f6e6c792074686520436f736d69635369676e617475726547616d6520636f6e5f8201527f7472616374206973207065726d697474656420746f2063616c6c20746869732060208201520152565b6114196047606092610fe0565b61142281611399565b0190565b9190611449906020611441604086018681035f88015261140c565b940190610644565b565b156114535750565b6114759061145f6101b2565b91829163336e7de360e11b835260048301611426565b0390fd5b906114aa916114a53361149d611497611492600161138c565b61025c565b9161025c565b14339061144b565b611520565b565b5f1b90565b906114bd5f19916114ac565b9181191691161790565b906114dc6114d76114e3926110b4565b6110d0565b82546114b1565b9055565b634e487b7160e01b5f52601160045260245ffd5b61150a611510919392936101cf565b926101cf565b820180921161151b57565b6114e7565b6115716115486115426802000000000000000361153c86610efa565b90610f12565b50610fbe565b611554835f83016114c7565b61156b600134920191611566836107b4565b6114fb565b906114c7565b9034916115bc6115aa6115a47f999946acc98c7b7dacc26921697d55abbcb1637484b0a73040f0b06287e43326936110b4565b9361110d565b936115b36101b2565b91829182610434565b0390a3565b906115cb91611479565b565b60016115d991016101cf565b90565b5090565b91908110156115f0576040020190565b610379565b356115ff816102e1565b90565b3561160c816105be565b90565b919061161a5f610fc4565b5b8061163861163261162d8786906115dc565b6101cf565b916101cf565b10156116825761167d9061167861165b5f611655888786916115e0565b016115f5565b611672602061166c898887916115e0565b01611602565b90611f39565b6115cd565b61161b565b50509050565b611690612830565b6116986116c2565b565b6116ae6116a96116b392610fc1565b610851565b610251565b90565b6116bf9061169a565b90565b6116d36116ce5f6116b6565b61289e565b565b6116dd611688565b565b6117146116f961171b935f936116f36110f5565b50612965565b74020000000000000000000000020000000000000003610b58565b50016107b4565b90565b5f90565b61172a61171e565b506117345f61138c565b90565b90565b9061174490610dac565b9052565b6117526060610ea3565b90565b906117a461179b6002611766611748565b9461177d6117755f83016107b4565b5f8801610f3a565b61179561178c600183016107e0565b6020880161173a565b016107b4565b60408401610f3a565b565b6117af90611755565b90565b6117bc9051610dac565b90565b5f7f446f6e61746564204e465420616c726561647920636c61696d65642e00000000910152565b6117f3601c602092610fe0565b6117fc816117bf565b0190565b919061182390602061181b604086018681035f8801526117e6565b940190610427565b565b5f7f496e76616c696420646f6e61746564204e465420696e6465782e000000000000910152565b611859601a602092610fe0565b61186281611825565b0190565b9190611889906020611881604086018681035f88015261184c565b940190610427565b565b61189590516101cf565b90565b60407f686973204e465420756e74696c20612074696d656f757420657870697265732e917f4f6e6c79207468652062696464696e6720726f756e64206d61696e207072697a5f8201527f652077696e6e6572206973207065726d697474656420746f20636c61696d207460208201520152565b61191760608092610fe0565b61192081611898565b0190565b91604061195592949361194e611943606083018381035f85015261190b565b966020830190610644565b0190610427565b565b15611960575050565b61198161196b6101b2565b92839263d1b9b13b60e01b845260048401611924565b0390fd5b919060086119a591029161199f60018060a01b038461108b565b9261108b565b9181191691161790565b6119b890610854565b90565b6119c4906119af565b90565b90565b91906119e06119db6119e8936119bb565b6119c7565b908354611985565b9055565b5f90565b611a02916119fc6119ec565b916119ca565b565b611a39611a4094611a2f606094989795611a25608086019a5f870190610644565b6020850190610888565b6040830190610427565b0190610427565b565b611a4b90610870565b90565b5f80fd5b60e01b90565b5f910312611a6257565b6101bc565b604090611a90611a979496959396611a8660608401985f850190610644565b6020830190610644565b0190610427565b565b611aa16101b2565b3d5f823e3d90fd5b90611adb611ad57c01000000000000000200000000000000000000000200000000000000048490610773565b50611737565b611ae4816117a6565b90611af9611af4602084016117b2565b61087c565b611b13611b0d611b085f6116b6565b61025c565b9161025c565b145f14611ba7578380611b53611b4d611b487c01000000000000000200000000000000000000000200000000000000036107b4565b6101cf565b916101cf565b10155f14611b8157611b7d90611b676101b2565b9182916306420ad160e41b835260048301611866565b0390fd5b611ba390611b8d6101b2565b91829163aa576f4960e01b835260048301611800565b0390fd5b5f6002611c07929593949533611be2611bdc611bd7611bd186611bcb898d0161188b565b90610d23565b90610ac6565b61025c565b9161025c565b03611d19575b611bf4838083016110f9565b611c0183600183016119f0565b016110f9565b611c125f830161188b565b33611c1f602085016117b2565b90611c6c611c2f6040870161188b565b94611c5a7f03c2b6e01c9ca39e4073132f279b61b987d41a82b08cea3dd4a2fabc37067ac3956110b4565b95611c636101b2565b94859485611a04565b0390a2611c83611c7e602083016117b2565b61087c565b6323b872dd90611c9230611a42565b90611ca0604033950161188b565b92813b15611d14575f611cc691611cd18296611cba6101b2565b98899788968795611a52565b855260048501611a67565b03925af18015611d0f57611ce3575b50565b611d02905f3d8111611d08575b611cfa8183610e7a565b810190611a58565b5f611ce0565b503d611cf0565b611a99565b611a4e565b611d68611d43611d3d68010000000000000003611d37878b0161188b565b9061039e565b906103e5565b42611d56611d50836101cf565b916101cf565b10159081611d6d575b50338791611957565b611be8565b9050611d81611d7b86610fc4565b916101cf565b115f611d5f565b60607f612074696d656f757420657870697265732e0000000000000000000000000000917f4f6e6c79207468652062696464696e6720726f756e64206d61696e207072697a5f8201527f652077696e6e6572206973207065726d697474656420746f20636c61696d207460208201527f686973204552432d323020746f6b656e20646f6e6174696f6e20756e74696c2060408201520152565b611e2e6072608092610fe0565b611e3781611d88565b0190565b611e4490610870565b90565b611e5090611e3b565b9052565b606090611e8d611e949496959396611e83611e78608085018581035f870152611e21565b986020850190610427565b6040830190610644565b0190611e47565b565b9290919215611ea457505050565b611ec690611eb06101b2565b93849363b039722960e01b855260048501611e54565b0390fd5b90565b611ed76020610ea3565b90565b90611ef9611ef15f611eea611ecd565b94016107b4565b5f8401610f3a565b565b611f0490611eda565b90565b604090611f30611f379496959396611f2660608401985f850190610644565b6020830190611e47565b0190610427565b565b6120169133611f64611f5e611f59611f5360028790610d23565b90610ac6565b61025c565b9161025c565b03612018575b611f9c611f96611f7b848490612965565b74020000000000000000000000020000000000000003610b58565b50611eca565b91611fb25f80611fab86611efb565b95016110f9565b339082611fc05f860161188b565b92612000611fee7faf1adae2e1e983ca738335dc2e37194114142793f394ac934a45ea632a8a5bd0946110b4565b94611ff76101b2565b93849384611f07565b0390a2906120105f33920161188b565b916129f3565b565b61205f61203861203268010000000000000003859061039e565b906103e5565b4261204b612045836101cf565b916101cf565b10159081612064575b508333908492611e96565b611f6a565b90506120786120725f610fc4565b916101cf565b115f612054565b6120909061208b612830565b612092565b565b6120a581680100000000000000026114c7565b6120db7f8717bb199c6bc4a5dadb21547205f9ef8ec037dda246a5526d6a6471306ea52e916120d26101b2565b91829182610434565b0390a1565b6120e99061207f565b565b61211461211a916120fa610eef565b5061210e6802000000000000000391610efa565b90610f12565b50610f81565b90565b6121e661214561213f6802000000000000000361213933610efa565b90610f12565b50610fbe565b61216a5f80612156600185016107b4565b9361216482600183016110f9565b016110f9565b333382916121b66121a461219e7f4f43b861ba36494acfe938f3815fba7fac6981bdc611b6ccdc14c08f592923839361110d565b9361110d565b936121ad6101b2565b91829182610434565b0390a35f8033836121c56101b2565b90816121d08161112d565b03925af16121dc611178565b5090339091611215565b565b916121fd91612202959493612204575b61160f565b6124af565b565b61220c61211d565b6121f8565b906122429161223d3361223561222f61222a600161138c565b61025c565b9161025c565b14339061144b565b612269565b565b90565b919061225d6122586122659361110d565b612244565b908354611985565b9055565b6122826122b79261227c60028490610d23565b90612247565b6122b16122a14261229b680100000000000000026107b4565b906114fb565b916801000000000000000361039e565b906110d3565b565b906122c391612211565b565b906122f89392916122f3336122eb6122e56122e0600161138c565b61025c565b9161025c565b14339061144b565b612353565b565b5f7f5468652070726f76696465642061646472657373206973207a65726f2e000000910152565b61232e601d602092610fe0565b612337816122fa565b0190565b6123509060208101905f818303910152612321565b90565b92919061235f82611e3b565b61237961237361236e5f6116b6565b61025c565b9161025c565b1461238957612387936123c4565b565b6123916101b2565b63eac0d38960e01b8152806123a86004820161233b565b0390fd5b6123b590610854565b90565b6123c1906123ac565b90565b9161248693919261241c6124006123fa6123df848690612965565b74020000000000000000000000020000000000000003610b58565b50611eca565b6124165f86920191612411836107b4565b6114fb565b906114c7565b838290849261247261246061245a6124547f3f94f6171c96ab491d8c268f80da5e460b148878065711a0bb099bbc297615af946110b4565b9461110d565b946123b8565b946124696101b2565b91829182610434565b0390a49161247f30611a42565b9192612a41565b565b906124949392916122c5565b565b5090565b91908110156124aa576020020190565b610379565b9190916124bb5f610fc4565b5b806124d96124d36124ce858890612496565b6101cf565b916101cf565b101561250957612504906124ff6124fa6124f58588859161249a565b6115f5565b611aa9565b6115cd565b6124bc565b50509050565b6125209061251b612830565b612522565b565b8061253d6125376125325f6116b6565b61025c565b9161025c565b1461254d5761254b9061289e565b565b6125706125595f6116b6565b5f918291631e4fbdf760e01b835260048301610651565b0390fd5b61257d9061250f565b565b906125b29392916125ad336125a561259f61259a600161138c565b61025c565b9161025c565b14339061144b565b6125b4565b565b9291906125c08261087c565b6125da6125d46125cf5f6116b6565b61025c565b9161025c565b146125ea576125e893612687565b565b6125f26101b2565b63eac0d38960e01b8152806126096004820161233b565b0390fd5b9061261e60018060a01b03916114ac565b9181191691161790565b9061263d612638612644926119bb565b6119c7565b825461260d565b9055565b91602061266992949361266260408201965f830190610427565b0190610427565b565b612674906101cf565b5f1981146126825760010190565b6114e7565b9161279361277061279893946126b97c01000000000000000200000000000000000000000200000000000000036107b4565b906127116126ee6126e87c01000000000000000200000000000000000000000200000000000000048590610773565b50611737565b6126fa835f83016114c7565b6127078760018301612628565b60028a91016114c7565b86859089928461275361274d6127477fb12e72bab0c2a8fe98175a3c02792645ebdf65d37cf3983517e6d1c5ab842e23946110b4565b9461110d565b946119bb565b9461276861275f6101b2565b92839283612648565b0390a461266b565b7c01000000000000000200000000000000000000000200000000000000036114c7565b61087c565b6323b872dd91906127a830611a42565b9392813b1561281d575f6127cf916127da82966127c36101b2565b98899788968795611a52565b855260048501611a67565b03925af18015612818576127ec575b50565b61280b905f3d8111612811575b6128038183610e7a565b810190611a58565b5f6127e9565b503d6127f9565b611a99565b611a4e565b9061282e93929161257f565b565b612838611722565b61285161284b612846612a91565b61025c565b9161025c565b0361285857565b61287a612863612a91565b5f91829163118cdaa760e01b835260048301610651565b0390fd5b9061289361288e61289a9261110d565b612244565b825461260d565b9055565b6128a75f61138c565b6128b1825f61287e565b906128e56128df7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09361110d565b9161110d565b916128ee6101b2565b806128f8816102a9565b0390a3565b61291161290c61291692610251565b610851565b6101cf565b90565b90565b60ff1690565b61293661293161293b92612919565b610851565b61291c565b90565b61295d906129576129516129629461291c565b916101cf565b9061108b565b6101cf565b90565b61298b61298661298161299b9361297a6110f5565b5094611e3b565b610efa565b6128fd565b6129956040612922565b9061293e565b1790565b63ffffffff1690565b63ffffffff60e01b1690565b6129c86129c36129cd9261299f565b611a52565b6129a8565b90565b9160206129f19294936129ea60408201965f830190610644565b0190610427565b565b90612a3a612a3f93612a2b60049493612a1263a9059cbb9193916129b4565b92612a1b6101b2565b96879460208601908152016129d0565b60208201810382520383610e7a565b612acf565b565b600492612a7b612a8f9593612a8a9394612a626323b872dd929491926129b4565b93612a6b6101b2565b9788956020870190815201611a67565b60208201810382520383610e7a565b612acf565b565b612a9961171e565b503390565b5190565b90505190612aaf8261099f565b565b90602082820312612aca57612ac7915f01612aa2565b90565b6101bc565b90612ae290612add83611e3b565b612b60565b612aeb81612a9e565b612afd612af75f610fc4565b916101cf565b14159081612b35575b50612b0e5750565b612b1a612b3191611e3b565b5f918291635274afe760e01b835260048301610651565b0390fd5b612b5a9150612b54906020612b4982612a9e565b818301019101612ab1565b1561099a565b5f612b06565b90612b7e91612b6d611173565b5090612b785f610fc4565b91612b8d565b90565b612b8a90610870565b90565b9091612b97611173565b50612ba130612b81565b31612bb4612bae836101cf565b916101cf565b10612be0575f8091612bdd948491602082019151925af190612bd4611178565b90919091612c07565b90565b612c03612bec30612b81565b5f91829163cd78605960e01b835260048301610651565b0390fd5b90612c1b90612c14611173565b501561099a565b5f14612c275750612c8b565b612c3082612a9e565b612c42612c3c5f610fc4565b916101cf565b1480612c70575b612c51575090565b612c6c905f918291639996b31560e01b835260048301610651565b0390fd5b50803b612c85612c7f5f610fc4565b916101cf565b14612c49565b612c9481612a9e565b612ca6612ca05f610fc4565b916101cf565b115f14612cb557805190602001fd5b5f630a12f52160e11b815280612ccd600482016102a9565b0390fdfea2646970667358221220a6e7be27aa259afe9b86ac20b1478af1632b3724c30998abe29cfa999ca1b88b64736f6c634300081c0033",
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
// Solidity: function donatedTokens(uint256 ) view returns(uint256 amount)
func (_PrizesWallet *PrizesWalletCaller) DonatedTokens(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "donatedTokens", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DonatedTokens is a free data retrieval call binding the contract method 0xd7f4f8be.
//
// Solidity: function donatedTokens(uint256 ) view returns(uint256 amount)
func (_PrizesWallet *PrizesWalletSession) DonatedTokens(arg0 *big.Int) (*big.Int, error) {
	return _PrizesWallet.Contract.DonatedTokens(&_PrizesWallet.CallOpts, arg0)
}

// DonatedTokens is a free data retrieval call binding the contract method 0xd7f4f8be.
//
// Solidity: function donatedTokens(uint256 ) view returns(uint256 amount)
func (_PrizesWallet *PrizesWalletCallerSession) DonatedTokens(arg0 *big.Int) (*big.Int, error) {
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

// GetDonatedTokenAmount is a free data retrieval call binding the contract method 0x854fdf1b.
//
// Solidity: function getDonatedTokenAmount(uint256 roundNum_, address tokenAddress_) view returns(uint256)
func (_PrizesWallet *PrizesWalletCaller) GetDonatedTokenAmount(opts *bind.CallOpts, roundNum_ *big.Int, tokenAddress_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "getDonatedTokenAmount", roundNum_, tokenAddress_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDonatedTokenAmount is a free data retrieval call binding the contract method 0x854fdf1b.
//
// Solidity: function getDonatedTokenAmount(uint256 roundNum_, address tokenAddress_) view returns(uint256)
func (_PrizesWallet *PrizesWalletSession) GetDonatedTokenAmount(roundNum_ *big.Int, tokenAddress_ common.Address) (*big.Int, error) {
	return _PrizesWallet.Contract.GetDonatedTokenAmount(&_PrizesWallet.CallOpts, roundNum_, tokenAddress_)
}

// GetDonatedTokenAmount is a free data retrieval call binding the contract method 0x854fdf1b.
//
// Solidity: function getDonatedTokenAmount(uint256 roundNum_, address tokenAddress_) view returns(uint256)
func (_PrizesWallet *PrizesWalletCallerSession) GetDonatedTokenAmount(roundNum_ *big.Int, tokenAddress_ common.Address) (*big.Int, error) {
	return _PrizesWallet.Contract.GetDonatedTokenAmount(&_PrizesWallet.CallOpts, roundNum_, tokenAddress_)
}

// GetEthBalanceInfo is a free data retrieval call binding the contract method 0x1a18c889.
//
// Solidity: function getEthBalanceInfo() view returns((uint256,uint256))
func (_PrizesWallet *PrizesWalletCaller) GetEthBalanceInfo(opts *bind.CallOpts) (CosmicSignatureConstantsBalanceInfo, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "getEthBalanceInfo")

	if err != nil {
		return *new(CosmicSignatureConstantsBalanceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(CosmicSignatureConstantsBalanceInfo)).(*CosmicSignatureConstantsBalanceInfo)

	return out0, err

}

// GetEthBalanceInfo is a free data retrieval call binding the contract method 0x1a18c889.
//
// Solidity: function getEthBalanceInfo() view returns((uint256,uint256))
func (_PrizesWallet *PrizesWalletSession) GetEthBalanceInfo() (CosmicSignatureConstantsBalanceInfo, error) {
	return _PrizesWallet.Contract.GetEthBalanceInfo(&_PrizesWallet.CallOpts)
}

// GetEthBalanceInfo is a free data retrieval call binding the contract method 0x1a18c889.
//
// Solidity: function getEthBalanceInfo() view returns((uint256,uint256))
func (_PrizesWallet *PrizesWalletCallerSession) GetEthBalanceInfo() (CosmicSignatureConstantsBalanceInfo, error) {
	return _PrizesWallet.Contract.GetEthBalanceInfo(&_PrizesWallet.CallOpts)
}

// GetEthBalanceInfo0 is a free data retrieval call binding the contract method 0xa089e0be.
//
// Solidity: function getEthBalanceInfo(address roundPrizeWinnerAddress_) view returns((uint256,uint256))
func (_PrizesWallet *PrizesWalletCaller) GetEthBalanceInfo0(opts *bind.CallOpts, roundPrizeWinnerAddress_ common.Address) (CosmicSignatureConstantsBalanceInfo, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "getEthBalanceInfo0", roundPrizeWinnerAddress_)

	if err != nil {
		return *new(CosmicSignatureConstantsBalanceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(CosmicSignatureConstantsBalanceInfo)).(*CosmicSignatureConstantsBalanceInfo)

	return out0, err

}

// GetEthBalanceInfo0 is a free data retrieval call binding the contract method 0xa089e0be.
//
// Solidity: function getEthBalanceInfo(address roundPrizeWinnerAddress_) view returns((uint256,uint256))
func (_PrizesWallet *PrizesWalletSession) GetEthBalanceInfo0(roundPrizeWinnerAddress_ common.Address) (CosmicSignatureConstantsBalanceInfo, error) {
	return _PrizesWallet.Contract.GetEthBalanceInfo0(&_PrizesWallet.CallOpts, roundPrizeWinnerAddress_)
}

// GetEthBalanceInfo0 is a free data retrieval call binding the contract method 0xa089e0be.
//
// Solidity: function getEthBalanceInfo(address roundPrizeWinnerAddress_) view returns((uint256,uint256))
func (_PrizesWallet *PrizesWalletCallerSession) GetEthBalanceInfo0(roundPrizeWinnerAddress_ common.Address) (CosmicSignatureConstantsBalanceInfo, error) {
	return _PrizesWallet.Contract.GetEthBalanceInfo0(&_PrizesWallet.CallOpts, roundPrizeWinnerAddress_)
}

// MainPrizeWinnerAddresses is a free data retrieval call binding the contract method 0xf9eba390.
//
// Solidity: function mainPrizeWinnerAddresses(uint256 ) view returns(address)
func (_PrizesWallet *PrizesWalletCaller) MainPrizeWinnerAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "mainPrizeWinnerAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MainPrizeWinnerAddresses is a free data retrieval call binding the contract method 0xf9eba390.
//
// Solidity: function mainPrizeWinnerAddresses(uint256 ) view returns(address)
func (_PrizesWallet *PrizesWalletSession) MainPrizeWinnerAddresses(arg0 *big.Int) (common.Address, error) {
	return _PrizesWallet.Contract.MainPrizeWinnerAddresses(&_PrizesWallet.CallOpts, arg0)
}

// MainPrizeWinnerAddresses is a free data retrieval call binding the contract method 0xf9eba390.
//
// Solidity: function mainPrizeWinnerAddresses(uint256 ) view returns(address)
func (_PrizesWallet *PrizesWalletCallerSession) MainPrizeWinnerAddresses(arg0 *big.Int) (common.Address, error) {
	return _PrizesWallet.Contract.MainPrizeWinnerAddresses(&_PrizesWallet.CallOpts, arg0)
}

// NumDonatedNfts is a free data retrieval call binding the contract method 0x90946622.
//
// Solidity: function numDonatedNfts() view returns(uint256)
func (_PrizesWallet *PrizesWalletCaller) NumDonatedNfts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "numDonatedNfts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumDonatedNfts is a free data retrieval call binding the contract method 0x90946622.
//
// Solidity: function numDonatedNfts() view returns(uint256)
func (_PrizesWallet *PrizesWalletSession) NumDonatedNfts() (*big.Int, error) {
	return _PrizesWallet.Contract.NumDonatedNfts(&_PrizesWallet.CallOpts)
}

// NumDonatedNfts is a free data retrieval call binding the contract method 0x90946622.
//
// Solidity: function numDonatedNfts() view returns(uint256)
func (_PrizesWallet *PrizesWalletCallerSession) NumDonatedNfts() (*big.Int, error) {
	return _PrizesWallet.Contract.NumDonatedNfts(&_PrizesWallet.CallOpts)
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

// ClaimDonatedToken is a paid mutator transaction binding the contract method 0x95de6c2c.
//
// Solidity: function claimDonatedToken(uint256 roundNum_, address tokenAddress_) returns()
func (_PrizesWallet *PrizesWalletTransactor) ClaimDonatedToken(opts *bind.TransactOpts, roundNum_ *big.Int, tokenAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "claimDonatedToken", roundNum_, tokenAddress_)
}

// ClaimDonatedToken is a paid mutator transaction binding the contract method 0x95de6c2c.
//
// Solidity: function claimDonatedToken(uint256 roundNum_, address tokenAddress_) returns()
func (_PrizesWallet *PrizesWalletSession) ClaimDonatedToken(roundNum_ *big.Int, tokenAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.ClaimDonatedToken(&_PrizesWallet.TransactOpts, roundNum_, tokenAddress_)
}

// ClaimDonatedToken is a paid mutator transaction binding the contract method 0x95de6c2c.
//
// Solidity: function claimDonatedToken(uint256 roundNum_, address tokenAddress_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) ClaimDonatedToken(roundNum_ *big.Int, tokenAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.ClaimDonatedToken(&_PrizesWallet.TransactOpts, roundNum_, tokenAddress_)
}

// ClaimManyDonatedNfts is a paid mutator transaction binding the contract method 0xe4a6c2a4.
//
// Solidity: function claimManyDonatedNfts(uint256[] indices_) returns()
func (_PrizesWallet *PrizesWalletTransactor) ClaimManyDonatedNfts(opts *bind.TransactOpts, indices_ []*big.Int) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "claimManyDonatedNfts", indices_)
}

// ClaimManyDonatedNfts is a paid mutator transaction binding the contract method 0xe4a6c2a4.
//
// Solidity: function claimManyDonatedNfts(uint256[] indices_) returns()
func (_PrizesWallet *PrizesWalletSession) ClaimManyDonatedNfts(indices_ []*big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.ClaimManyDonatedNfts(&_PrizesWallet.TransactOpts, indices_)
}

// ClaimManyDonatedNfts is a paid mutator transaction binding the contract method 0xe4a6c2a4.
//
// Solidity: function claimManyDonatedNfts(uint256[] indices_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) ClaimManyDonatedNfts(indices_ []*big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.ClaimManyDonatedNfts(&_PrizesWallet.TransactOpts, indices_)
}

// ClaimManyDonatedTokens is a paid mutator transaction binding the contract method 0x50a15deb.
//
// Solidity: function claimManyDonatedTokens((uint256,address)[] donatedTokensToClaim_) returns()
func (_PrizesWallet *PrizesWalletTransactor) ClaimManyDonatedTokens(opts *bind.TransactOpts, donatedTokensToClaim_ []CosmicSignatureConstantsDonatedTokenToClaim) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "claimManyDonatedTokens", donatedTokensToClaim_)
}

// ClaimManyDonatedTokens is a paid mutator transaction binding the contract method 0x50a15deb.
//
// Solidity: function claimManyDonatedTokens((uint256,address)[] donatedTokensToClaim_) returns()
func (_PrizesWallet *PrizesWalletSession) ClaimManyDonatedTokens(donatedTokensToClaim_ []CosmicSignatureConstantsDonatedTokenToClaim) (*types.Transaction, error) {
	return _PrizesWallet.Contract.ClaimManyDonatedTokens(&_PrizesWallet.TransactOpts, donatedTokensToClaim_)
}

// ClaimManyDonatedTokens is a paid mutator transaction binding the contract method 0x50a15deb.
//
// Solidity: function claimManyDonatedTokens((uint256,address)[] donatedTokensToClaim_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) ClaimManyDonatedTokens(donatedTokensToClaim_ []CosmicSignatureConstantsDonatedTokenToClaim) (*types.Transaction, error) {
	return _PrizesWallet.Contract.ClaimManyDonatedTokens(&_PrizesWallet.TransactOpts, donatedTokensToClaim_)
}

// DepositEth is a paid mutator transaction binding the contract method 0x46f6b4e1.
//
// Solidity: function depositEth(uint256 roundNum_, address roundPrizeWinnerAddress_) payable returns()
func (_PrizesWallet *PrizesWalletTransactor) DepositEth(opts *bind.TransactOpts, roundNum_ *big.Int, roundPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "depositEth", roundNum_, roundPrizeWinnerAddress_)
}

// DepositEth is a paid mutator transaction binding the contract method 0x46f6b4e1.
//
// Solidity: function depositEth(uint256 roundNum_, address roundPrizeWinnerAddress_) payable returns()
func (_PrizesWallet *PrizesWalletSession) DepositEth(roundNum_ *big.Int, roundPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.DepositEth(&_PrizesWallet.TransactOpts, roundNum_, roundPrizeWinnerAddress_)
}

// DepositEth is a paid mutator transaction binding the contract method 0x46f6b4e1.
//
// Solidity: function depositEth(uint256 roundNum_, address roundPrizeWinnerAddress_) payable returns()
func (_PrizesWallet *PrizesWalletTransactorSession) DepositEth(roundNum_ *big.Int, roundPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.DepositEth(&_PrizesWallet.TransactOpts, roundNum_, roundPrizeWinnerAddress_)
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
// Solidity: function registerRoundEnd(uint256 roundNum_, address roundMainPrizeWinnerAddress_) returns()
func (_PrizesWallet *PrizesWalletTransactor) RegisterRoundEnd(opts *bind.TransactOpts, roundNum_ *big.Int, roundMainPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "registerRoundEnd", roundNum_, roundMainPrizeWinnerAddress_)
}

// RegisterRoundEnd is a paid mutator transaction binding the contract method 0xcc5810d8.
//
// Solidity: function registerRoundEnd(uint256 roundNum_, address roundMainPrizeWinnerAddress_) returns()
func (_PrizesWallet *PrizesWalletSession) RegisterRoundEnd(roundNum_ *big.Int, roundMainPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.RegisterRoundEnd(&_PrizesWallet.TransactOpts, roundNum_, roundMainPrizeWinnerAddress_)
}

// RegisterRoundEnd is a paid mutator transaction binding the contract method 0xcc5810d8.
//
// Solidity: function registerRoundEnd(uint256 roundNum_, address roundMainPrizeWinnerAddress_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) RegisterRoundEnd(roundNum_ *big.Int, roundMainPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.RegisterRoundEnd(&_PrizesWallet.TransactOpts, roundNum_, roundMainPrizeWinnerAddress_)
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
// Solidity: function withdrawEth(address roundPrizeWinnerAddress_) returns()
func (_PrizesWallet *PrizesWalletTransactor) WithdrawEth(opts *bind.TransactOpts, roundPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "withdrawEth", roundPrizeWinnerAddress_)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x25e16063.
//
// Solidity: function withdrawEth(address roundPrizeWinnerAddress_) returns()
func (_PrizesWallet *PrizesWalletSession) WithdrawEth(roundPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.WithdrawEth(&_PrizesWallet.TransactOpts, roundPrizeWinnerAddress_)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x25e16063.
//
// Solidity: function withdrawEth(address roundPrizeWinnerAddress_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) WithdrawEth(roundPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.WithdrawEth(&_PrizesWallet.TransactOpts, roundPrizeWinnerAddress_)
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

// WithdrawEverything is a paid mutator transaction binding the contract method 0xa72be1b2.
//
// Solidity: function withdrawEverything(bool withdrawEth_, (uint256,address)[] donatedTokensToClaim_, uint256[] donatedNftIndices_) returns()
func (_PrizesWallet *PrizesWalletTransactor) WithdrawEverything(opts *bind.TransactOpts, withdrawEth_ bool, donatedTokensToClaim_ []CosmicSignatureConstantsDonatedTokenToClaim, donatedNftIndices_ []*big.Int) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "withdrawEverything", withdrawEth_, donatedTokensToClaim_, donatedNftIndices_)
}

// WithdrawEverything is a paid mutator transaction binding the contract method 0xa72be1b2.
//
// Solidity: function withdrawEverything(bool withdrawEth_, (uint256,address)[] donatedTokensToClaim_, uint256[] donatedNftIndices_) returns()
func (_PrizesWallet *PrizesWalletSession) WithdrawEverything(withdrawEth_ bool, donatedTokensToClaim_ []CosmicSignatureConstantsDonatedTokenToClaim, donatedNftIndices_ []*big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.WithdrawEverything(&_PrizesWallet.TransactOpts, withdrawEth_, donatedTokensToClaim_, donatedNftIndices_)
}

// WithdrawEverything is a paid mutator transaction binding the contract method 0xa72be1b2.
//
// Solidity: function withdrawEverything(bool withdrawEth_, (uint256,address)[] donatedTokensToClaim_, uint256[] donatedNftIndices_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) WithdrawEverything(withdrawEth_ bool, donatedTokensToClaim_ []CosmicSignatureConstantsDonatedTokenToClaim, donatedNftIndices_ []*big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.WithdrawEverything(&_PrizesWallet.TransactOpts, withdrawEth_, donatedTokensToClaim_, donatedNftIndices_)
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
// Solidity: event DonatedNftClaimed(uint256 indexed roundNum, address beneficiaryAddress, address nftAddress, uint256 nftId, uint256 index)
func (_PrizesWallet *PrizesWalletFilterer) FilterDonatedNftClaimed(opts *bind.FilterOpts, roundNum []*big.Int) (*PrizesWalletDonatedNftClaimedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _PrizesWallet.contract.FilterLogs(opts, "DonatedNftClaimed", roundNumRule)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletDonatedNftClaimedIterator{contract: _PrizesWallet.contract, event: "DonatedNftClaimed", logs: logs, sub: sub}, nil
}

// WatchDonatedNftClaimed is a free log subscription operation binding the contract event 0x03c2b6e01c9ca39e4073132f279b61b987d41a82b08cea3dd4a2fabc37067ac3.
//
// Solidity: event DonatedNftClaimed(uint256 indexed roundNum, address beneficiaryAddress, address nftAddress, uint256 nftId, uint256 index)
func (_PrizesWallet *PrizesWalletFilterer) WatchDonatedNftClaimed(opts *bind.WatchOpts, sink chan<- *PrizesWalletDonatedNftClaimed, roundNum []*big.Int) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _PrizesWallet.contract.WatchLogs(opts, "DonatedNftClaimed", roundNumRule)
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
// Solidity: event DonatedNftClaimed(uint256 indexed roundNum, address beneficiaryAddress, address nftAddress, uint256 nftId, uint256 index)
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
// Solidity: event DonatedTokenClaimed(uint256 indexed roundNum, address beneficiaryAddress, address tokenAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) FilterDonatedTokenClaimed(opts *bind.FilterOpts, roundNum []*big.Int) (*PrizesWalletDonatedTokenClaimedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _PrizesWallet.contract.FilterLogs(opts, "DonatedTokenClaimed", roundNumRule)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletDonatedTokenClaimedIterator{contract: _PrizesWallet.contract, event: "DonatedTokenClaimed", logs: logs, sub: sub}, nil
}

// WatchDonatedTokenClaimed is a free log subscription operation binding the contract event 0xaf1adae2e1e983ca738335dc2e37194114142793f394ac934a45ea632a8a5bd0.
//
// Solidity: event DonatedTokenClaimed(uint256 indexed roundNum, address beneficiaryAddress, address tokenAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) WatchDonatedTokenClaimed(opts *bind.WatchOpts, sink chan<- *PrizesWalletDonatedTokenClaimed, roundNum []*big.Int) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _PrizesWallet.contract.WatchLogs(opts, "DonatedTokenClaimed", roundNumRule)
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
// Solidity: event DonatedTokenClaimed(uint256 indexed roundNum, address beneficiaryAddress, address tokenAddress, uint256 amount)
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
	RoundNum                *big.Int
	RoundPrizeWinnerAddress common.Address
	Amount                  *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterEthReceived is a free log retrieval operation binding the contract event 0x999946acc98c7b7dacc26921697d55abbcb1637484b0a73040f0b06287e43326.
//
// Solidity: event EthReceived(uint256 indexed roundNum, address indexed roundPrizeWinnerAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) FilterEthReceived(opts *bind.FilterOpts, roundNum []*big.Int, roundPrizeWinnerAddress []common.Address) (*PrizesWalletEthReceivedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var roundPrizeWinnerAddressRule []interface{}
	for _, roundPrizeWinnerAddressItem := range roundPrizeWinnerAddress {
		roundPrizeWinnerAddressRule = append(roundPrizeWinnerAddressRule, roundPrizeWinnerAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.FilterLogs(opts, "EthReceived", roundNumRule, roundPrizeWinnerAddressRule)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletEthReceivedIterator{contract: _PrizesWallet.contract, event: "EthReceived", logs: logs, sub: sub}, nil
}

// WatchEthReceived is a free log subscription operation binding the contract event 0x999946acc98c7b7dacc26921697d55abbcb1637484b0a73040f0b06287e43326.
//
// Solidity: event EthReceived(uint256 indexed roundNum, address indexed roundPrizeWinnerAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) WatchEthReceived(opts *bind.WatchOpts, sink chan<- *PrizesWalletEthReceived, roundNum []*big.Int, roundPrizeWinnerAddress []common.Address) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var roundPrizeWinnerAddressRule []interface{}
	for _, roundPrizeWinnerAddressItem := range roundPrizeWinnerAddress {
		roundPrizeWinnerAddressRule = append(roundPrizeWinnerAddressRule, roundPrizeWinnerAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.WatchLogs(opts, "EthReceived", roundNumRule, roundPrizeWinnerAddressRule)
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

// ParseEthReceived is a log parse operation binding the contract event 0x999946acc98c7b7dacc26921697d55abbcb1637484b0a73040f0b06287e43326.
//
// Solidity: event EthReceived(uint256 indexed roundNum, address indexed roundPrizeWinnerAddress, uint256 amount)
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
	RoundPrizeWinnerAddress common.Address
	BeneficiaryAddress      common.Address
	Amount                  *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterEthWithdrawn is a free log retrieval operation binding the contract event 0x4f43b861ba36494acfe938f3815fba7fac6981bdc611b6ccdc14c08f59292383.
//
// Solidity: event EthWithdrawn(address indexed roundPrizeWinnerAddress, address indexed beneficiaryAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) FilterEthWithdrawn(opts *bind.FilterOpts, roundPrizeWinnerAddress []common.Address, beneficiaryAddress []common.Address) (*PrizesWalletEthWithdrawnIterator, error) {

	var roundPrizeWinnerAddressRule []interface{}
	for _, roundPrizeWinnerAddressItem := range roundPrizeWinnerAddress {
		roundPrizeWinnerAddressRule = append(roundPrizeWinnerAddressRule, roundPrizeWinnerAddressItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.FilterLogs(opts, "EthWithdrawn", roundPrizeWinnerAddressRule, beneficiaryAddressRule)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletEthWithdrawnIterator{contract: _PrizesWallet.contract, event: "EthWithdrawn", logs: logs, sub: sub}, nil
}

// WatchEthWithdrawn is a free log subscription operation binding the contract event 0x4f43b861ba36494acfe938f3815fba7fac6981bdc611b6ccdc14c08f59292383.
//
// Solidity: event EthWithdrawn(address indexed roundPrizeWinnerAddress, address indexed beneficiaryAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) WatchEthWithdrawn(opts *bind.WatchOpts, sink chan<- *PrizesWalletEthWithdrawn, roundPrizeWinnerAddress []common.Address, beneficiaryAddress []common.Address) (event.Subscription, error) {

	var roundPrizeWinnerAddressRule []interface{}
	for _, roundPrizeWinnerAddressItem := range roundPrizeWinnerAddress {
		roundPrizeWinnerAddressRule = append(roundPrizeWinnerAddressRule, roundPrizeWinnerAddressItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.WatchLogs(opts, "EthWithdrawn", roundPrizeWinnerAddressRule, beneficiaryAddressRule)
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
// Solidity: event EthWithdrawn(address indexed roundPrizeWinnerAddress, address indexed beneficiaryAddress, uint256 amount)
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
