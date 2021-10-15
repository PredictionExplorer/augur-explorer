// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// RWMarketABI is the input ABI used to generate the binding from.
const RWMarketABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"offerId\",\"type\":\"uint256\"}],\"name\":\"ItemBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"offerId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"NewOffer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"offerId\",\"type\":\"uint256\"}],\"name\":\"OfferCanceled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offerId\",\"type\":\"uint256\"}],\"name\":\"acceptBuyOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offerId\",\"type\":\"uint256\"}],\"name\":\"acceptSellOffer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offerId\",\"type\":\"uint256\"}],\"name\":\"cancelBuyOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offerId\",\"type\":\"uint256\"}],\"name\":\"cancelSellOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getBuyOffers\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"getBuyOffersBy\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"getBuyTokensBy\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getSellOffers\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"getSellOffersBy\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"getSellTokenBy\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"makeBuyOffer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"makeSellOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numOffers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"offers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// RWMarket is an auto generated Go binding around an Ethereum contract.
type RWMarket struct {
	RWMarketCaller     // Read-only binding to the contract
	RWMarketTransactor // Write-only binding to the contract
	RWMarketFilterer   // Log filterer for contract events
}

// RWMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type RWMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RWMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RWMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RWMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RWMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RWMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RWMarketSession struct {
	Contract     *RWMarket         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RWMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RWMarketCallerSession struct {
	Contract *RWMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RWMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RWMarketTransactorSession struct {
	Contract     *RWMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RWMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type RWMarketRaw struct {
	Contract *RWMarket // Generic contract binding to access the raw methods on
}

// RWMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RWMarketCallerRaw struct {
	Contract *RWMarketCaller // Generic read-only contract binding to access the raw methods on
}

// RWMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RWMarketTransactorRaw struct {
	Contract *RWMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRWMarket creates a new instance of RWMarket, bound to a specific deployed contract.
func NewRWMarket(address common.Address, backend bind.ContractBackend) (*RWMarket, error) {
	contract, err := bindRWMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RWMarket{RWMarketCaller: RWMarketCaller{contract: contract}, RWMarketTransactor: RWMarketTransactor{contract: contract}, RWMarketFilterer: RWMarketFilterer{contract: contract}}, nil
}

// NewRWMarketCaller creates a new read-only instance of RWMarket, bound to a specific deployed contract.
func NewRWMarketCaller(address common.Address, caller bind.ContractCaller) (*RWMarketCaller, error) {
	contract, err := bindRWMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RWMarketCaller{contract: contract}, nil
}

// NewRWMarketTransactor creates a new write-only instance of RWMarket, bound to a specific deployed contract.
func NewRWMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*RWMarketTransactor, error) {
	contract, err := bindRWMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RWMarketTransactor{contract: contract}, nil
}

// NewRWMarketFilterer creates a new log filterer instance of RWMarket, bound to a specific deployed contract.
func NewRWMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*RWMarketFilterer, error) {
	contract, err := bindRWMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RWMarketFilterer{contract: contract}, nil
}

// bindRWMarket binds a generic wrapper to an already deployed contract.
func bindRWMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RWMarketABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RWMarket *RWMarketRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RWMarket.Contract.RWMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RWMarket *RWMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RWMarket.Contract.RWMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RWMarket *RWMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RWMarket.Contract.RWMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RWMarket *RWMarketCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RWMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RWMarket *RWMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RWMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RWMarket *RWMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RWMarket.Contract.contract.Transact(opts, method, params...)
}

// GetBuyOffers is a free data retrieval call binding the contract method 0xdbea17f2.
//
// Solidity: function getBuyOffers(uint256 tokenId) view returns(uint256[])
func (_RWMarket *RWMarketCaller) GetBuyOffers(opts *bind.CallOpts, tokenId *big.Int) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _RWMarket.contract.Call(opts, out, "getBuyOffers", tokenId)
	return *ret0, err
}

// GetBuyOffers is a free data retrieval call binding the contract method 0xdbea17f2.
//
// Solidity: function getBuyOffers(uint256 tokenId) view returns(uint256[])
func (_RWMarket *RWMarketSession) GetBuyOffers(tokenId *big.Int) ([]*big.Int, error) {
	return _RWMarket.Contract.GetBuyOffers(&_RWMarket.CallOpts, tokenId)
}

// GetBuyOffers is a free data retrieval call binding the contract method 0xdbea17f2.
//
// Solidity: function getBuyOffers(uint256 tokenId) view returns(uint256[])
func (_RWMarket *RWMarketCallerSession) GetBuyOffers(tokenId *big.Int) ([]*big.Int, error) {
	return _RWMarket.Contract.GetBuyOffers(&_RWMarket.CallOpts, tokenId)
}

// GetBuyOffersBy is a free data retrieval call binding the contract method 0xfd5820eb.
//
// Solidity: function getBuyOffersBy(address buyer) view returns(uint256[])
func (_RWMarket *RWMarketCaller) GetBuyOffersBy(opts *bind.CallOpts, buyer common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _RWMarket.contract.Call(opts, out, "getBuyOffersBy", buyer)
	return *ret0, err
}

// GetBuyOffersBy is a free data retrieval call binding the contract method 0xfd5820eb.
//
// Solidity: function getBuyOffersBy(address buyer) view returns(uint256[])
func (_RWMarket *RWMarketSession) GetBuyOffersBy(buyer common.Address) ([]*big.Int, error) {
	return _RWMarket.Contract.GetBuyOffersBy(&_RWMarket.CallOpts, buyer)
}

// GetBuyOffersBy is a free data retrieval call binding the contract method 0xfd5820eb.
//
// Solidity: function getBuyOffersBy(address buyer) view returns(uint256[])
func (_RWMarket *RWMarketCallerSession) GetBuyOffersBy(buyer common.Address) ([]*big.Int, error) {
	return _RWMarket.Contract.GetBuyOffersBy(&_RWMarket.CallOpts, buyer)
}

// GetBuyTokensBy is a free data retrieval call binding the contract method 0x78d2c9f1.
//
// Solidity: function getBuyTokensBy(address buyer) view returns(uint256[])
func (_RWMarket *RWMarketCaller) GetBuyTokensBy(opts *bind.CallOpts, buyer common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _RWMarket.contract.Call(opts, out, "getBuyTokensBy", buyer)
	return *ret0, err
}

// GetBuyTokensBy is a free data retrieval call binding the contract method 0x78d2c9f1.
//
// Solidity: function getBuyTokensBy(address buyer) view returns(uint256[])
func (_RWMarket *RWMarketSession) GetBuyTokensBy(buyer common.Address) ([]*big.Int, error) {
	return _RWMarket.Contract.GetBuyTokensBy(&_RWMarket.CallOpts, buyer)
}

// GetBuyTokensBy is a free data retrieval call binding the contract method 0x78d2c9f1.
//
// Solidity: function getBuyTokensBy(address buyer) view returns(uint256[])
func (_RWMarket *RWMarketCallerSession) GetBuyTokensBy(buyer common.Address) ([]*big.Int, error) {
	return _RWMarket.Contract.GetBuyTokensBy(&_RWMarket.CallOpts, buyer)
}

// GetSellOffers is a free data retrieval call binding the contract method 0x5507f91c.
//
// Solidity: function getSellOffers(uint256 tokenId) view returns(uint256[])
func (_RWMarket *RWMarketCaller) GetSellOffers(opts *bind.CallOpts, tokenId *big.Int) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _RWMarket.contract.Call(opts, out, "getSellOffers", tokenId)
	return *ret0, err
}

// GetSellOffers is a free data retrieval call binding the contract method 0x5507f91c.
//
// Solidity: function getSellOffers(uint256 tokenId) view returns(uint256[])
func (_RWMarket *RWMarketSession) GetSellOffers(tokenId *big.Int) ([]*big.Int, error) {
	return _RWMarket.Contract.GetSellOffers(&_RWMarket.CallOpts, tokenId)
}

// GetSellOffers is a free data retrieval call binding the contract method 0x5507f91c.
//
// Solidity: function getSellOffers(uint256 tokenId) view returns(uint256[])
func (_RWMarket *RWMarketCallerSession) GetSellOffers(tokenId *big.Int) ([]*big.Int, error) {
	return _RWMarket.Contract.GetSellOffers(&_RWMarket.CallOpts, tokenId)
}

// GetSellOffersBy is a free data retrieval call binding the contract method 0x883332e2.
//
// Solidity: function getSellOffersBy(address seller) view returns(uint256[])
func (_RWMarket *RWMarketCaller) GetSellOffersBy(opts *bind.CallOpts, seller common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _RWMarket.contract.Call(opts, out, "getSellOffersBy", seller)
	return *ret0, err
}

// GetSellOffersBy is a free data retrieval call binding the contract method 0x883332e2.
//
// Solidity: function getSellOffersBy(address seller) view returns(uint256[])
func (_RWMarket *RWMarketSession) GetSellOffersBy(seller common.Address) ([]*big.Int, error) {
	return _RWMarket.Contract.GetSellOffersBy(&_RWMarket.CallOpts, seller)
}

// GetSellOffersBy is a free data retrieval call binding the contract method 0x883332e2.
//
// Solidity: function getSellOffersBy(address seller) view returns(uint256[])
func (_RWMarket *RWMarketCallerSession) GetSellOffersBy(seller common.Address) ([]*big.Int, error) {
	return _RWMarket.Contract.GetSellOffersBy(&_RWMarket.CallOpts, seller)
}

// GetSellTokenBy is a free data retrieval call binding the contract method 0x55479649.
//
// Solidity: function getSellTokenBy(address seller) view returns(uint256[])
func (_RWMarket *RWMarketCaller) GetSellTokenBy(opts *bind.CallOpts, seller common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _RWMarket.contract.Call(opts, out, "getSellTokenBy", seller)
	return *ret0, err
}

// GetSellTokenBy is a free data retrieval call binding the contract method 0x55479649.
//
// Solidity: function getSellTokenBy(address seller) view returns(uint256[])
func (_RWMarket *RWMarketSession) GetSellTokenBy(seller common.Address) ([]*big.Int, error) {
	return _RWMarket.Contract.GetSellTokenBy(&_RWMarket.CallOpts, seller)
}

// GetSellTokenBy is a free data retrieval call binding the contract method 0x55479649.
//
// Solidity: function getSellTokenBy(address seller) view returns(uint256[])
func (_RWMarket *RWMarketCallerSession) GetSellTokenBy(seller common.Address) ([]*big.Int, error) {
	return _RWMarket.Contract.GetSellTokenBy(&_RWMarket.CallOpts, seller)
}

// NumOffers is a free data retrieval call binding the contract method 0xcc6bee54.
//
// Solidity: function numOffers() view returns(uint256)
func (_RWMarket *RWMarketCaller) NumOffers(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RWMarket.contract.Call(opts, out, "numOffers")
	return *ret0, err
}

// NumOffers is a free data retrieval call binding the contract method 0xcc6bee54.
//
// Solidity: function numOffers() view returns(uint256)
func (_RWMarket *RWMarketSession) NumOffers() (*big.Int, error) {
	return _RWMarket.Contract.NumOffers(&_RWMarket.CallOpts)
}

// NumOffers is a free data retrieval call binding the contract method 0xcc6bee54.
//
// Solidity: function numOffers() view returns(uint256)
func (_RWMarket *RWMarketCallerSession) NumOffers() (*big.Int, error) {
	return _RWMarket.Contract.NumOffers(&_RWMarket.CallOpts)
}

// Offers is a free data retrieval call binding the contract method 0x8a72ea6a.
//
// Solidity: function offers(uint256 ) view returns(uint256 tokenId, uint256 price, address seller, address buyer, bool active)
func (_RWMarket *RWMarketCaller) Offers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TokenId *big.Int
	Price   *big.Int
	Seller  common.Address
	Buyer   common.Address
	Active  bool
}, error) {
	ret := new(struct {
		TokenId *big.Int
		Price   *big.Int
		Seller  common.Address
		Buyer   common.Address
		Active  bool
	})
	out := ret
	err := _RWMarket.contract.Call(opts, out, "offers", arg0)
	return *ret, err
}

// Offers is a free data retrieval call binding the contract method 0x8a72ea6a.
//
// Solidity: function offers(uint256 ) view returns(uint256 tokenId, uint256 price, address seller, address buyer, bool active)
func (_RWMarket *RWMarketSession) Offers(arg0 *big.Int) (struct {
	TokenId *big.Int
	Price   *big.Int
	Seller  common.Address
	Buyer   common.Address
	Active  bool
}, error) {
	return _RWMarket.Contract.Offers(&_RWMarket.CallOpts, arg0)
}

// Offers is a free data retrieval call binding the contract method 0x8a72ea6a.
//
// Solidity: function offers(uint256 ) view returns(uint256 tokenId, uint256 price, address seller, address buyer, bool active)
func (_RWMarket *RWMarketCallerSession) Offers(arg0 *big.Int) (struct {
	TokenId *big.Int
	Price   *big.Int
	Seller  common.Address
	Buyer   common.Address
	Active  bool
}, error) {
	return _RWMarket.Contract.Offers(&_RWMarket.CallOpts, arg0)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) pure returns(bytes4)
func (_RWMarket *RWMarketCaller) OnERC721Received(opts *bind.CallOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _RWMarket.contract.Call(opts, out, "onERC721Received", operator, from, tokenId, data)
	return *ret0, err
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) pure returns(bytes4)
func (_RWMarket *RWMarketSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) ([4]byte, error) {
	return _RWMarket.Contract.OnERC721Received(&_RWMarket.CallOpts, operator, from, tokenId, data)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) pure returns(bytes4)
func (_RWMarket *RWMarketCallerSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) ([4]byte, error) {
	return _RWMarket.Contract.OnERC721Received(&_RWMarket.CallOpts, operator, from, tokenId, data)
}

// AcceptBuyOffer is a paid mutator transaction binding the contract method 0x45502349.
//
// Solidity: function acceptBuyOffer(uint256 offerId) returns()
func (_RWMarket *RWMarketTransactor) AcceptBuyOffer(opts *bind.TransactOpts, offerId *big.Int) (*types.Transaction, error) {
	return _RWMarket.contract.Transact(opts, "acceptBuyOffer", offerId)
}

// AcceptBuyOffer is a paid mutator transaction binding the contract method 0x45502349.
//
// Solidity: function acceptBuyOffer(uint256 offerId) returns()
func (_RWMarket *RWMarketSession) AcceptBuyOffer(offerId *big.Int) (*types.Transaction, error) {
	return _RWMarket.Contract.AcceptBuyOffer(&_RWMarket.TransactOpts, offerId)
}

// AcceptBuyOffer is a paid mutator transaction binding the contract method 0x45502349.
//
// Solidity: function acceptBuyOffer(uint256 offerId) returns()
func (_RWMarket *RWMarketTransactorSession) AcceptBuyOffer(offerId *big.Int) (*types.Transaction, error) {
	return _RWMarket.Contract.AcceptBuyOffer(&_RWMarket.TransactOpts, offerId)
}

// AcceptSellOffer is a paid mutator transaction binding the contract method 0x09b06721.
//
// Solidity: function acceptSellOffer(uint256 offerId) payable returns()
func (_RWMarket *RWMarketTransactor) AcceptSellOffer(opts *bind.TransactOpts, offerId *big.Int) (*types.Transaction, error) {
	return _RWMarket.contract.Transact(opts, "acceptSellOffer", offerId)
}

// AcceptSellOffer is a paid mutator transaction binding the contract method 0x09b06721.
//
// Solidity: function acceptSellOffer(uint256 offerId) payable returns()
func (_RWMarket *RWMarketSession) AcceptSellOffer(offerId *big.Int) (*types.Transaction, error) {
	return _RWMarket.Contract.AcceptSellOffer(&_RWMarket.TransactOpts, offerId)
}

// AcceptSellOffer is a paid mutator transaction binding the contract method 0x09b06721.
//
// Solidity: function acceptSellOffer(uint256 offerId) payable returns()
func (_RWMarket *RWMarketTransactorSession) AcceptSellOffer(offerId *big.Int) (*types.Transaction, error) {
	return _RWMarket.Contract.AcceptSellOffer(&_RWMarket.TransactOpts, offerId)
}

// CancelBuyOffer is a paid mutator transaction binding the contract method 0x1a46829c.
//
// Solidity: function cancelBuyOffer(uint256 offerId) returns()
func (_RWMarket *RWMarketTransactor) CancelBuyOffer(opts *bind.TransactOpts, offerId *big.Int) (*types.Transaction, error) {
	return _RWMarket.contract.Transact(opts, "cancelBuyOffer", offerId)
}

// CancelBuyOffer is a paid mutator transaction binding the contract method 0x1a46829c.
//
// Solidity: function cancelBuyOffer(uint256 offerId) returns()
func (_RWMarket *RWMarketSession) CancelBuyOffer(offerId *big.Int) (*types.Transaction, error) {
	return _RWMarket.Contract.CancelBuyOffer(&_RWMarket.TransactOpts, offerId)
}

// CancelBuyOffer is a paid mutator transaction binding the contract method 0x1a46829c.
//
// Solidity: function cancelBuyOffer(uint256 offerId) returns()
func (_RWMarket *RWMarketTransactorSession) CancelBuyOffer(offerId *big.Int) (*types.Transaction, error) {
	return _RWMarket.Contract.CancelBuyOffer(&_RWMarket.TransactOpts, offerId)
}

// CancelSellOffer is a paid mutator transaction binding the contract method 0x98d91a26.
//
// Solidity: function cancelSellOffer(uint256 offerId) returns()
func (_RWMarket *RWMarketTransactor) CancelSellOffer(opts *bind.TransactOpts, offerId *big.Int) (*types.Transaction, error) {
	return _RWMarket.contract.Transact(opts, "cancelSellOffer", offerId)
}

// CancelSellOffer is a paid mutator transaction binding the contract method 0x98d91a26.
//
// Solidity: function cancelSellOffer(uint256 offerId) returns()
func (_RWMarket *RWMarketSession) CancelSellOffer(offerId *big.Int) (*types.Transaction, error) {
	return _RWMarket.Contract.CancelSellOffer(&_RWMarket.TransactOpts, offerId)
}

// CancelSellOffer is a paid mutator transaction binding the contract method 0x98d91a26.
//
// Solidity: function cancelSellOffer(uint256 offerId) returns()
func (_RWMarket *RWMarketTransactorSession) CancelSellOffer(offerId *big.Int) (*types.Transaction, error) {
	return _RWMarket.Contract.CancelSellOffer(&_RWMarket.TransactOpts, offerId)
}

// MakeBuyOffer is a paid mutator transaction binding the contract method 0x48bf7c6a.
//
// Solidity: function makeBuyOffer(uint256 tokenId) payable returns()
func (_RWMarket *RWMarketTransactor) MakeBuyOffer(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _RWMarket.contract.Transact(opts, "makeBuyOffer", tokenId)
}

// MakeBuyOffer is a paid mutator transaction binding the contract method 0x48bf7c6a.
//
// Solidity: function makeBuyOffer(uint256 tokenId) payable returns()
func (_RWMarket *RWMarketSession) MakeBuyOffer(tokenId *big.Int) (*types.Transaction, error) {
	return _RWMarket.Contract.MakeBuyOffer(&_RWMarket.TransactOpts, tokenId)
}

// MakeBuyOffer is a paid mutator transaction binding the contract method 0x48bf7c6a.
//
// Solidity: function makeBuyOffer(uint256 tokenId) payable returns()
func (_RWMarket *RWMarketTransactorSession) MakeBuyOffer(tokenId *big.Int) (*types.Transaction, error) {
	return _RWMarket.Contract.MakeBuyOffer(&_RWMarket.TransactOpts, tokenId)
}

// MakeSellOffer is a paid mutator transaction binding the contract method 0xc728c75e.
//
// Solidity: function makeSellOffer(uint256 tokenId, uint256 price) returns()
func (_RWMarket *RWMarketTransactor) MakeSellOffer(opts *bind.TransactOpts, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _RWMarket.contract.Transact(opts, "makeSellOffer", tokenId, price)
}

// MakeSellOffer is a paid mutator transaction binding the contract method 0xc728c75e.
//
// Solidity: function makeSellOffer(uint256 tokenId, uint256 price) returns()
func (_RWMarket *RWMarketSession) MakeSellOffer(tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _RWMarket.Contract.MakeSellOffer(&_RWMarket.TransactOpts, tokenId, price)
}

// MakeSellOffer is a paid mutator transaction binding the contract method 0xc728c75e.
//
// Solidity: function makeSellOffer(uint256 tokenId, uint256 price) returns()
func (_RWMarket *RWMarketTransactorSession) MakeSellOffer(tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _RWMarket.Contract.MakeSellOffer(&_RWMarket.TransactOpts, tokenId, price)
}

// RWMarketItemBoughtIterator is returned from FilterItemBought and is used to iterate over the raw logs and unpacked data for ItemBought events raised by the RWMarket contract.
type RWMarketItemBoughtIterator struct {
	Event *RWMarketItemBought // Event containing the contract specifics and raw log

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
func (it *RWMarketItemBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWMarketItemBought)
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
		it.Event = new(RWMarketItemBought)
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
func (it *RWMarketItemBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWMarketItemBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWMarketItemBought represents a ItemBought event raised by the RWMarket contract.
type RWMarketItemBought struct {
	OfferId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterItemBought is a free log retrieval operation binding the contract event 0x7f7e375fbeaef0d6ebfc53af15b7aeed1d704e3424f34ef67e914b1f68f8c8ef.
//
// Solidity: event ItemBought(uint256 indexed offerId)
func (_RWMarket *RWMarketFilterer) FilterItemBought(opts *bind.FilterOpts, offerId []*big.Int) (*RWMarketItemBoughtIterator, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}

	logs, sub, err := _RWMarket.contract.FilterLogs(opts, "ItemBought", offerIdRule)
	if err != nil {
		return nil, err
	}
	return &RWMarketItemBoughtIterator{contract: _RWMarket.contract, event: "ItemBought", logs: logs, sub: sub}, nil
}

// WatchItemBought is a free log subscription operation binding the contract event 0x7f7e375fbeaef0d6ebfc53af15b7aeed1d704e3424f34ef67e914b1f68f8c8ef.
//
// Solidity: event ItemBought(uint256 indexed offerId)
func (_RWMarket *RWMarketFilterer) WatchItemBought(opts *bind.WatchOpts, sink chan<- *RWMarketItemBought, offerId []*big.Int) (event.Subscription, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}

	logs, sub, err := _RWMarket.contract.WatchLogs(opts, "ItemBought", offerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWMarketItemBought)
				if err := _RWMarket.contract.UnpackLog(event, "ItemBought", log); err != nil {
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

// ParseItemBought is a log parse operation binding the contract event 0x7f7e375fbeaef0d6ebfc53af15b7aeed1d704e3424f34ef67e914b1f68f8c8ef.
//
// Solidity: event ItemBought(uint256 indexed offerId)
func (_RWMarket *RWMarketFilterer) ParseItemBought(log types.Log) (*RWMarketItemBought, error) {
	event := new(RWMarketItemBought)
	if err := _RWMarket.contract.UnpackLog(event, "ItemBought", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RWMarketNewOfferIterator is returned from FilterNewOffer and is used to iterate over the raw logs and unpacked data for NewOffer events raised by the RWMarket contract.
type RWMarketNewOfferIterator struct {
	Event *RWMarketNewOffer // Event containing the contract specifics and raw log

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
func (it *RWMarketNewOfferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWMarketNewOffer)
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
		it.Event = new(RWMarketNewOffer)
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
func (it *RWMarketNewOfferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWMarketNewOfferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWMarketNewOffer represents a NewOffer event raised by the RWMarket contract.
type RWMarketNewOffer struct {
	OfferId *big.Int
	TokenId *big.Int
	Seller  common.Address
	Buyer   common.Address
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewOffer is a free log retrieval operation binding the contract event 0x8b4d06c200b17b9c1150172953ceb6fa3e7ace7623f6f933707badfa52c354cf.
//
// Solidity: event NewOffer(uint256 indexed offerId, uint256 indexed tokenId, address seller, address buyer, uint256 price)
func (_RWMarket *RWMarketFilterer) FilterNewOffer(opts *bind.FilterOpts, offerId []*big.Int, tokenId []*big.Int) (*RWMarketNewOfferIterator, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _RWMarket.contract.FilterLogs(opts, "NewOffer", offerIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &RWMarketNewOfferIterator{contract: _RWMarket.contract, event: "NewOffer", logs: logs, sub: sub}, nil
}

// WatchNewOffer is a free log subscription operation binding the contract event 0x8b4d06c200b17b9c1150172953ceb6fa3e7ace7623f6f933707badfa52c354cf.
//
// Solidity: event NewOffer(uint256 indexed offerId, uint256 indexed tokenId, address seller, address buyer, uint256 price)
func (_RWMarket *RWMarketFilterer) WatchNewOffer(opts *bind.WatchOpts, sink chan<- *RWMarketNewOffer, offerId []*big.Int, tokenId []*big.Int) (event.Subscription, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _RWMarket.contract.WatchLogs(opts, "NewOffer", offerIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWMarketNewOffer)
				if err := _RWMarket.contract.UnpackLog(event, "NewOffer", log); err != nil {
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

// ParseNewOffer is a log parse operation binding the contract event 0x8b4d06c200b17b9c1150172953ceb6fa3e7ace7623f6f933707badfa52c354cf.
//
// Solidity: event NewOffer(uint256 indexed offerId, uint256 indexed tokenId, address seller, address buyer, uint256 price)
func (_RWMarket *RWMarketFilterer) ParseNewOffer(log types.Log) (*RWMarketNewOffer, error) {
	event := new(RWMarketNewOffer)
	if err := _RWMarket.contract.UnpackLog(event, "NewOffer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RWMarketOfferCanceledIterator is returned from FilterOfferCanceled and is used to iterate over the raw logs and unpacked data for OfferCanceled events raised by the RWMarket contract.
type RWMarketOfferCanceledIterator struct {
	Event *RWMarketOfferCanceled // Event containing the contract specifics and raw log

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
func (it *RWMarketOfferCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWMarketOfferCanceled)
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
		it.Event = new(RWMarketOfferCanceled)
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
func (it *RWMarketOfferCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWMarketOfferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWMarketOfferCanceled represents a OfferCanceled event raised by the RWMarket contract.
type RWMarketOfferCanceled struct {
	OfferId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOfferCanceled is a free log retrieval operation binding the contract event 0x0ff09947dd7d2583091e8cbfb427fecacb697bf895187b243fd0072c0ee9b951.
//
// Solidity: event OfferCanceled(uint256 indexed offerId)
func (_RWMarket *RWMarketFilterer) FilterOfferCanceled(opts *bind.FilterOpts, offerId []*big.Int) (*RWMarketOfferCanceledIterator, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}

	logs, sub, err := _RWMarket.contract.FilterLogs(opts, "OfferCanceled", offerIdRule)
	if err != nil {
		return nil, err
	}
	return &RWMarketOfferCanceledIterator{contract: _RWMarket.contract, event: "OfferCanceled", logs: logs, sub: sub}, nil
}

// WatchOfferCanceled is a free log subscription operation binding the contract event 0x0ff09947dd7d2583091e8cbfb427fecacb697bf895187b243fd0072c0ee9b951.
//
// Solidity: event OfferCanceled(uint256 indexed offerId)
func (_RWMarket *RWMarketFilterer) WatchOfferCanceled(opts *bind.WatchOpts, sink chan<- *RWMarketOfferCanceled, offerId []*big.Int) (event.Subscription, error) {

	var offerIdRule []interface{}
	for _, offerIdItem := range offerId {
		offerIdRule = append(offerIdRule, offerIdItem)
	}

	logs, sub, err := _RWMarket.contract.WatchLogs(opts, "OfferCanceled", offerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWMarketOfferCanceled)
				if err := _RWMarket.contract.UnpackLog(event, "OfferCanceled", log); err != nil {
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

// ParseOfferCanceled is a log parse operation binding the contract event 0x0ff09947dd7d2583091e8cbfb427fecacb697bf895187b243fd0072c0ee9b951.
//
// Solidity: event OfferCanceled(uint256 indexed offerId)
func (_RWMarket *RWMarketFilterer) ParseOfferCanceled(log types.Log) (*RWMarketOfferCanceled, error) {
	event := new(RWMarketOfferCanceled)
	if err := _RWMarket.contract.UnpackLog(event, "OfferCanceled", log); err != nil {
		return nil, err
	}
	return event, nil
}
