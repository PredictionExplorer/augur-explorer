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

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220adf975bab74bae8a1b2ee567b0874a9ed87daca30bd4fdf4940cf168032ee1d864736f6c63430008130033",
}

// AddressABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressMetaData.ABI instead.
var AddressABI = AddressMetaData.ABI

// AddressBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressMetaData.Bin instead.
var AddressBin = AddressMetaData.Bin

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// CosmicGameMetaData contains all meta data concerning the CosmicGame contract.
var CosmicGameMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newActivationTime\",\"type\":\"uint256\"}],\"name\":\"ActivationTimeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lastBidder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"randomWalkNFTId\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prizeTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"BidEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCharity\",\"type\":\"address\"}],\"name\":\"CharityAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCharityPercentage\",\"type\":\"uint256\"}],\"name\":\"CharityPercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCosmicSignature\",\"type\":\"address\"}],\"name\":\"CosmicSignatureAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCosmicToken\",\"type\":\"address\"}],\"name\":\"CosmicTokenAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddressdonatedNFTs\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"DonatedNFTClaimedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DonationEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newInitialBidAmountFraction\",\"type\":\"uint256\"}],\"name\":\"InitialBidAmountFractionChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newInitialSecondsUntilPrize\",\"type\":\"uint256\"}],\"name\":\"InitialSecondsUntilPrizeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"NFTDonationEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNanoSecondsExtra\",\"type\":\"uint256\"}],\"name\":\"NanoSecondsExtraChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNumHolderNFTWinnersPerRound\",\"type\":\"uint256\"}],\"name\":\"NumHolderNFTWinnersPerRoundChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNumRaffleNFTWinnersPerRound\",\"type\":\"uint256\"}],\"name\":\"NumRaffleNFTWinnersPerRoundChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNumRaffleWinnersPerRound\",\"type\":\"uint256\"}],\"name\":\"NumRaffleWinnersPerRoundChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPriceIncrease\",\"type\":\"uint256\"}],\"name\":\"PriceIncreaseChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prizeNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PrizeClaimEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrizePercentage\",\"type\":\"uint256\"}],\"name\":\"PrizePercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"RaffleNFTClaimedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winnerIndex\",\"type\":\"uint256\"}],\"name\":\"RaffleNFTWinnerEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRafflePercentage\",\"type\":\"uint256\"}],\"name\":\"RafflePercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRaffleWallet\",\"type\":\"address\"}],\"name\":\"RaffleWalletAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRandomWalk\",\"type\":\"address\"}],\"name\":\"RandomWalkAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTimeIncrease\",\"type\":\"uint256\"}],\"name\":\"TimeIncreaseChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_MESSAGE_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MILLION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_REWARD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"bidAndDonateNFT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"randomWalkNFTId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"bidWithRWLK\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"randomWalkNFTId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"bidWithRWLKAndDonateNFT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"charity\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"charityAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"charityPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"claimDonatedNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimPrize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimRaffleNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"donatedNFTs\",\"outputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialBidAmountFraction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialSecondsUntilPrize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBidder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nanoSecondsExtra\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nft\",\"outputs\":[{\"internalType\":\"contractCosmicSignature\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numDonatedNFTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numHolderNFTWinnersPerRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRaffleNFTWinnersPerRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRaffleParticipants\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRaffleWinnersPerRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceIncrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prizeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prizePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prizeTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"raffleAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"raffleEntropy\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"raffleNFTWinners\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"raffleParticipants\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rafflePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"raffleWallet\",\"outputs\":[{\"internalType\":\"contractRaffleWallet\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomWalk\",\"outputs\":[{\"internalType\":\"contractRandomWalkNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newActivationTime\",\"type\":\"uint256\"}],\"name\":\"setActivationTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setCharity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newCharityPercentage\",\"type\":\"uint256\"}],\"name\":\"setCharityPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newInitialSecondsUntilPrize\",\"type\":\"uint256\"}],\"name\":\"setInitialSecondsUntilPrize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newNanoSecondsExtra\",\"type\":\"uint256\"}],\"name\":\"setNanoSecondsExtra\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setNftContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newNumHolderNFTWinnersPerRound\",\"type\":\"uint256\"}],\"name\":\"setNumHolderNFTWinnersPerRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newNumRaffleNFTWinnersPerRound\",\"type\":\"uint256\"}],\"name\":\"setNumRaffleNFTWinnersPerRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newNumRaffleWinnersPerRound\",\"type\":\"uint256\"}],\"name\":\"setNumRaffleWinnersPerRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPriceIncrease\",\"type\":\"uint256\"}],\"name\":\"setPriceIncrease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newRafflePercentage\",\"type\":\"uint256\"}],\"name\":\"setRafflePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setRaffleWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setRandomWalk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimeIncrease\",\"type\":\"uint256\"}],\"name\":\"setTimeIncrease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setTokenContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeIncrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeUntilActivation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeUntilPrize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractCosmicToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newInitialBidAmountFraction\",\"type\":\"uint256\"}],\"name\":\"updateInitialBidAmountFraction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPrizePercentage\",\"type\":\"uint256\"}],\"name\":\"updatePrizePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"usedRandomWalkNFTs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"winners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052620f695060015565034630b8a000600255620f42a460035566038d7ea4c680006004556000600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600a600755620151806008556103e86009556019600a556005600b556003600c556005600d556002600e55600060105563644709f0601355348015620000ae57600080fd5b50620000cf620000c36200015760201b60201c565b6200015f60201b60201c565b424340604051602001620000e5929190620002ba565b60405160208183030381529060405280519060200120601481905550620001116200015760201b60201c565b600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620002fc565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600082825260208201905092915050565b7f436f736d6963205369676e617475726520323032330000000000000000000000600082015250565b60006200026c60158362000223565b9150620002798262000234565b602082019050919050565b6000819050919050565b620002998162000284565b82525050565b6000819050919050565b620002b4816200029f565b82525050565b60006060820190508181036000830152620002d5816200025d565b9050620002e660208301856200028e565b620002f56040830184620002a9565b9392505050565b615bfe806200030c6000396000f3fe6080604052600436106103e85760003560e01c80638da5cb5b11610208578063d59d747811610118578063ec34866d116100ab578063f2fde38b1161007a578063f2fde38b14610e37578063f717882214610e60578063f8c3405014610e8b578063fb6f71a314610eb6578063fc0c546a14610edf57610407565b8063ec34866d14610dc0578063ed08339814610deb578063ed88c68e14610e16578063efeacc5614610e2057610407565b8063dbc945c0116100e7578063dbc945c014610d04578063dec08d8e14610d2f578063e1381d7e14610d5a578063ebc3f1d114610d8357610407565b8063d59d747814610c45578063d6e1741714610c6e578063d94d031614610cae578063da4493f614610cd957610407565b8063a6f9cc151161019b578063c01c5de21161016a578063c01c5de214610b72578063c7c8378d14610b9b578063c94028c214610bc4578063cb819dc014610bef578063d4ba890214610c1a57610407565b8063a6f9cc1514610adb578063ae6247f214610b04578063bb0bc58e14610b2d578063bbcd5bbe14610b4957610407565b8063a2fb1175116101d7578063a2fb117514610a2e578063a325b7d214610a6b578063a672f6e114610a87578063a6ceac2c14610ab057610407565b80638da5cb5b146109825780639136d6d9146109ad5780639250c33c146109d8578063934aa02314610a0357610407565b80634ac3a39511610303578063785fa6271161029657806380de163d1161026557806380de163d1461089b5780638547af30146108c657806386e378c9146108f15780638b1222741461092e5780638b1329e01461095757610407565b8063785fa62714610800578063799d431d1461082b5780637aef951c146108565780637c5486a21461087257610407565b80635e6e47aa116102d25780635e6e47aa14610780578063647b3e7f146107a957806370740ac9146107d2578063715018a6146107e957610407565b80634ac3a395146106da5780635111a2d614610703578063519645881461072e57806352f5ad771461075757610407565b80632aab32231161037b5780633c83adc41161034a5780633c83adc41461063057806340e023221461065957806347ccca02146106845780634a773f33146106af57610407565b80632aab32231461058457806332bc934c146105af57806332d382cd146105da5780633bec7b691461060557610407565b806311dc7335116103b757806311dc7335146104c6578063150b7a02146104f157806319afe4731461052e5780632a0272111461055957610407565b806304a57c091461040c57806305ba9b6714610449578063062fb10014610472578063119b22b31461049b57610407565b366104075761040560405180602001604052806000815250610f0a565b005b600080fd5b34801561041857600080fd5b50610433600480360381019061042e919061435e565b6110f1565b60405161044091906143cc565b60405180910390f35b34801561045557600080fd5b50610470600480360381019061046b919061435e565b611124565b005b34801561047e57600080fd5b5061049960048036038101906104949190614413565b6111e3565b005b3480156104a757600080fd5b506104b0611349565b6040516104bd919061444f565b60405180910390f35b3480156104d257600080fd5b506104db61134f565b6040516104e8919061444f565b60405180910390f35b3480156104fd57600080fd5b50610518600480360381019061051391906144cf565b611355565b6040516105259190614592565b60405180910390f35b34801561053a57600080fd5b5061054361136a565b604051610550919061444f565b60405180910390f35b34801561056557600080fd5b5061056e611370565b60405161057b919061444f565b60405180910390f35b34801561059057600080fd5b50610599611376565b6040516105a6919061444f565b60405180910390f35b3480156105bb57600080fd5b506105c461137c565b6040516105d1919061444f565b60405180910390f35b3480156105e657600080fd5b506105ef611383565b6040516105fc919061460c565b60405180910390f35b34801561061157600080fd5b5061061a6113a9565b604051610627919061444f565b60405180910390f35b34801561063c57600080fd5b506106576004803603810190610652919061435e565b6113af565b005b34801561066557600080fd5b5061066e6114b1565b60405161067b919061444f565b60405180910390f35b34801561069057600080fd5b506106996114b7565b6040516106a69190614648565b60405180910390f35b3480156106bb57600080fd5b506106c46114dd565b6040516106d1919061444f565b60405180910390f35b3480156106e657600080fd5b5061070160048036038101906106fc919061435e565b6114e3565b005b34801561070f57600080fd5b506107186115a2565b6040516107259190614684565b60405180910390f35b34801561073a57600080fd5b506107556004803603810190610750919061435e565b6115c8565b005b34801561076357600080fd5b5061077e60048036038101906107799190614413565b611687565b005b34801561078c57600080fd5b506107a760048036038101906107a2919061435e565b6117ed565b005b3480156107b557600080fd5b506107d060048036038101906107cb919061435e565b6118ef565b005b3480156107de57600080fd5b506107e76119ae565b005b3480156107f557600080fd5b506107fe612619565b005b34801561080c57600080fd5b506108156126a1565b604051610822919061444f565b60405180910390f35b34801561083757600080fd5b506108406126c2565b60405161084d919061444f565b60405180910390f35b610870600480360381019061086b91906147e0565b610f0a565b005b34801561087e57600080fd5b506108996004803603810190610894919061435e565b6126c8565b005b3480156108a757600080fd5b506108b0612787565b6040516108bd9190614842565b60405180910390f35b3480156108d257600080fd5b506108db61278d565b6040516108e891906143cc565b60405180910390f35b3480156108fd57600080fd5b506109186004803603810190610913919061435e565b6127b3565b6040516109259190614878565b60405180910390f35b34801561093a57600080fd5b506109556004803603810190610950919061435e565b6127d3565b005b34801561096357600080fd5b5061096c612892565b604051610979919061444f565b60405180910390f35b34801561098e57600080fd5b506109976128bb565b6040516109a491906143cc565b60405180910390f35b3480156109b957600080fd5b506109c26128e4565b6040516109cf919061444f565b60405180910390f35b3480156109e457600080fd5b506109ed6128ea565b6040516109fa919061444f565b60405180910390f35b348015610a0f57600080fd5b50610a186128f0565b604051610a2591906143cc565b60405180910390f35b348015610a3a57600080fd5b50610a556004803603810190610a50919061435e565b612916565b604051610a6291906143cc565b60405180910390f35b610a856004803603810190610a8091906148d1565b612949565b005b348015610a9357600080fd5b50610aae6004803603810190610aa9919061435e565b612963565b005b348015610abc57600080fd5b50610ac5612a22565b604051610ad2919061444f565b60405180910390f35b348015610ae757600080fd5b50610b026004803603810190610afd9190614413565b612a28565b005b348015610b1057600080fd5b50610b2b6004803603810190610b269190614954565b612b8e565b005b610b476004803603810190610b4291906149b0565b612db5565b005b348015610b5557600080fd5b50610b706004803603810190610b6b9190614413565b612dcd565b005b348015610b7e57600080fd5b50610b996004803603810190610b94919061435e565b612f33565b005b348015610ba757600080fd5b50610bc26004803603810190610bbd919061435e565b612ff2565b005b348015610bd057600080fd5b50610bd96130f4565b604051610be6919061444f565b60405180910390f35b348015610bfb57600080fd5b50610c04613115565b604051610c11919061444f565b60405180910390f35b348015610c2657600080fd5b50610c2f61311b565b604051610c3c919061444f565b60405180910390f35b348015610c5157600080fd5b50610c6c6004803603810190610c67919061435e565b613128565b005b348015610c7a57600080fd5b50610c956004803603810190610c90919061435e565b613421565b604051610ca59493929190614a40565b60405180910390f35b348015610cba57600080fd5b50610cc361347e565b604051610cd0919061444f565b60405180910390f35b348015610ce557600080fd5b50610cee613484565b604051610cfb919061444f565b60405180910390f35b348015610d1057600080fd5b50610d1961348a565b604051610d26919061444f565b60405180910390f35b348015610d3b57600080fd5b50610d446134ab565b604051610d51919061444f565b60405180910390f35b348015610d6657600080fd5b50610d816004803603810190610d7c919061435e565b6134b1565b005b348015610d8f57600080fd5b50610daa6004803603810190610da59190614413565b613570565b604051610db7919061444f565b60405180910390f35b348015610dcc57600080fd5b50610dd5613588565b604051610de2919061444f565b60405180910390f35b348015610df757600080fd5b50610e006135ad565b604051610e0d919061444f565b60405180910390f35b610e1e6135b3565b005b348015610e2c57600080fd5b50610e356136ac565b005b348015610e4357600080fd5b50610e5e6004803603810190610e599190614413565b61392b565b005b348015610e6c57600080fd5b50610e75613a22565b604051610e82919061444f565b60405180910390f35b348015610e9757600080fd5b50610ea0613a4b565b604051610ead919061444f565b60405180910390f35b348015610ec257600080fd5b50610edd6004803603810190610ed89190614413565b613a51565b005b348015610eeb57600080fd5b50610ef4613bd9565b604051610f019190614aa6565b60405180910390f35b6000610f14613588565b905080341015610f59576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f5090614b44565b60405180910390fd5b80600481905550610f6982613bff565b600454341115611050576000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660045434610fbc9190614b93565b604051610fc890614bf8565b60006040518083038185875af1925050503d8060008114611005576040519150601f19603f3d011682016040523d82523d6000602084013e61100a565b606091505b505090508061104e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161104590614c59565b60405180910390fd5b505b601054600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fc7beef8f8d25068377c10e7b7f30ef5622f443eb674e07835df1daf5fe84acbf6004547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600f54876040516110e59493929190614d2c565b60405180910390a35050565b60166020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61112c613f4a565b73ffffffffffffffffffffffffffffffffffffffff1661114a6128bb565b73ffffffffffffffffffffffffffffffffffffffff16146111a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161119790614dc4565b60405180910390fd5b806002819055507f678d086a09e1be49401b1e3a6e14db1878e8d8b88e71d0cfe24a32726d0e38b96002546040516111d8919061444f565b60405180910390a150565b6111eb613f4a565b73ffffffffffffffffffffffffffffffffffffffff166112096128bb565b73ffffffffffffffffffffffffffffffffffffffff161461125f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161125690614dc4565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036112ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112c590614e30565b60405180910390fd5b80601860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f508d510537704d37ed64691a706392abc33e59ce8a327b4952e112820ceb88a68160405161133e91906143cc565b60405180910390a150565b60105481565b600a5481565b600063150b7a0260e01b905095945050505050565b60045481565b61011881565b601a5481565b620f424081565b601860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60175481565b6113b7613f4a565b73ffffffffffffffffffffffffffffffffffffffff166113d56128bb565b73ffffffffffffffffffffffffffffffffffffffff161461142b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161142290614dc4565b60405180910390fd5b6064811061146e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161146590614ec2565b60405180910390fd5b80600a819055507f595fa5ba64cd6f66df19b53b59bb4a275bae1ec5b362da95e1fa4c5feb62be1e600a546040516114a6919061444f565b60405180910390a150565b60095481565b601c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600c5481565b6114eb613f4a565b73ffffffffffffffffffffffffffffffffffffffff166115096128bb565b73ffffffffffffffffffffffffffffffffffffffff161461155f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161155690614dc4565b60405180910390fd5b806003819055507fed46e73b174ced51fb529cdc6c1b4d8abf49387e6d849b71648afb63c81d12cd600354604051611597919061444f565b60405180910390a150565b601d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6115d0613f4a565b73ffffffffffffffffffffffffffffffffffffffff166115ee6128bb565b73ffffffffffffffffffffffffffffffffffffffff1614611644576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161163b90614dc4565b60405180910390fd5b806008819055507f6da281754ba85ee0c5983a8e8f05a92910c2a0c5b80e68c126216d65f162a30560085460405161167c919061444f565b60405180910390a150565b61168f613f4a565b73ffffffffffffffffffffffffffffffffffffffff166116ad6128bb565b73ffffffffffffffffffffffffffffffffffffffff1614611703576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116fa90614dc4565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611772576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161176990614e30565b60405180910390fd5b80601c60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f7142a592d5404b5fc4a294c66f70e32b2a7776bd807f722d59268def87c765d1816040516117e291906143cc565b60405180910390a150565b6117f5613f4a565b73ffffffffffffffffffffffffffffffffffffffff166118136128bb565b73ffffffffffffffffffffffffffffffffffffffff1614611869576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161186090614dc4565b60405180910390fd5b606481106118ac576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118a390614ec2565b60405180910390fd5b806007819055507f0918b499c15fcb0e087d411d53664cde23577e1aa4a9cbfbdf735ddd1871e7d56007546040516118e4919061444f565b60405180910390a150565b6118f7613f4a565b73ffffffffffffffffffffffffffffffffffffffff166119156128bb565b73ffffffffffffffffffffffffffffffffffffffff161461196b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161196290614dc4565b60405180910390fd5b80600c819055507f5e3567ae359e61c47ad2f11d2828dea7e062b2ff63dea385cdb7900a50504c7c600c546040516119a3919061444f565b60405180910390a150565b42600f5411156119f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119ea90614f2e565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603611a84576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a7b90614f9a565b60405180910390fd5b62015180600f5442611a969190614b93565b1015611b3457600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16611add613f4a565b73ffffffffffffffffffffffffffffffffffffffff1614611b33576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b2a90615052565b60405180910390fd5b5b6000611b3e613f4a565b90506000600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060126000601054815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160106000828254611be99190615072565b925050819055506000601c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636a62784260e01b83604051602401611c4491906143cc565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051611cae91906150e2565b6000604051808303816000865af19150503d8060008114611ceb576040519150601f19603f3d011682016040523d82523d6000602084013e611cf0565b606091505b5050905080611d34576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d2b9061516b565b60405180910390fd5b6000611d3e6126a1565b90506000611d4a61348a565b90506000611d566130f4565b905060005b600d54811015611e39576000611d6f613f52565b90506001601560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611dc19190615072565b925050819055506001601054611dd79190614b93565b8173ffffffffffffffffffffffffffffffffffffffff167f80348bf864c08069d1368c42ed36b7a60560f73267f63d58e9be69f4b021bacc84604051611e1d919061444f565b60405180910390a3508080611e319061518b565b915050611d5b565b506000601d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611ea9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ecd91906151e8565b90506000601c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611f3e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f6291906151e8565b90506000600d54905060005b600e548110156122a357600084111561210457611f89613fa7565b60008460145460001c611f9c9190615244565b90506000601d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636352211e836040518263ffffffff1660e01b8152600401611ffb919061444f565b602060405180830381865afa158015612018573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061203c919061528a565b90506001601560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461208e9190615072565b9250508190555060016010546120a49190614b93565b8173ffffffffffffffffffffffffffffffffffffffff167f80348bf864c08069d1368c42ed36b7a60560f73267f63d58e9be69f4b021bacc866040516120ea919061444f565b60405180910390a36001846120ff9190615072565b935050505b600083111561229057612115613fa7565b60008360145460001c6121289190615244565b90506000601c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636352211e836040518263ffffffff1660e01b8152600401612187919061444f565b602060405180830381865afa1580156121a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121c8919061528a565b90506001601560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461221a9190615072565b9250508190555060016010546122309190614b93565b8173ffffffffffffffffffffffffffffffffffffffff167f80348bf864c08069d1368c42ed36b7a60560f73267f63d58e9be69f4b021bacc86604051612276919061444f565b60405180910390a360018461228b9190615072565b935050505b808061229b9061518b565b915050611f6e565b5060008873ffffffffffffffffffffffffffffffffffffffff16876040516122ca90614bf8565b60006040518083038185875af1925050503d8060008114612307576040519150601f19603f3d011682016040523d82523d6000602084013e61230c565b606091505b5050905080612350576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161234790615303565b60405180910390fd5b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168660405161239690614bf8565b60006040518083038185875af1925050503d80600081146123d3576040519150601f19603f3d011682016040523d82523d6000602084013e6123d8565b606091505b5050809150508061241e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161241590615395565b60405180910390fd5b60005b600c548110156125a0576000612435613f52565b9050601860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16876347e7ef2460e01b8360016010546124899190614b93565b60405160240161249a9291906153b5565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505060405161250491906150e2565b60006040518083038185875af1925050503d8060008114612541576040519150601f19603f3d011682016040523d82523d6000602084013e612546565b606091505b5050809350508261258c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125839061542a565b60405180910390fd5b5080806125989061518b565b915050612421565b5060006017819055506125b1613fdd565b8873ffffffffffffffffffffffffffffffffffffffff1660016010546125d79190614b93565b7f27bc828c399c2947fea27bca8a75ced2e94ff2651d607271f051e39db52286ce89604051612606919061444f565b60405180910390a3505050505050505050565b612621613f4a565b73ffffffffffffffffffffffffffffffffffffffff1661263f6128bb565b73ffffffffffffffffffffffffffffffffffffffff1614612695576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161268c90614dc4565b60405180910390fd5b61269f6000613ffa565b565b60006064600a54476126b3919061544a565b6126bd919061548c565b905090565b60075481565b6126d0613f4a565b73ffffffffffffffffffffffffffffffffffffffff166126ee6128bb565b73ffffffffffffffffffffffffffffffffffffffff1614612744576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161273b90614dc4565b60405180910390fd5b806013819055507f584ff9a7b3a86db5397052f2e440da0ed60a95f646a3a884863cd92262e682b660135460405161277c919061444f565b60405180910390a150565b60145481565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60116020528060005260406000206000915054906101000a900460ff1681565b6127db613f4a565b73ffffffffffffffffffffffffffffffffffffffff166127f96128bb565b73ffffffffffffffffffffffffffffffffffffffff161461284f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161284690614dc4565b60405180910390fd5b806001819055507fcbc1f49adfa29e2f2f0f5c9e057722496a4bc95a6a5446deaa423a02b30c64ac600154604051612887919061444f565b60405180910390a150565b600042600f5410156128a757600090506128b8565b42600f546128b59190614b93565b90505b90565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60025481565b600b5481565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60126020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6129538484612b8e565b61295d82826140be565b50505050565b61296b613f4a565b73ffffffffffffffffffffffffffffffffffffffff166129896128bb565b73ffffffffffffffffffffffffffffffffffffffff16146129df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016129d690614dc4565b60405180910390fd5b806009819055507f3b311f029da1b90c3de2e0b3168436c5ed7d8b0ae81b7d4894c12da03835c628600954604051612a17919061444f565b60405180910390a150565b600d5481565b612a30613f4a565b73ffffffffffffffffffffffffffffffffffffffff16612a4e6128bb565b73ffffffffffffffffffffffffffffffffffffffff1614612aa4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a9b90614dc4565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612b13576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b0a90614e30565b60405180910390fd5b80601d60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f9dc3707e4b88be65295b0083b44ffa94059c80372add9b0d24d6a3b371a03b9281604051612b8391906143cc565b60405180910390a150565b6011600083815260200190815260200160002060009054906101000a900460ff1615612bef576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612be69061552f565b60405180910390fd5b612bf7613f4a565b73ffffffffffffffffffffffffffffffffffffffff16601d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636352211e846040518263ffffffff1660e01b8152600401612c68919061444f565b602060405180830381865afa158015612c85573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ca9919061528a565b73ffffffffffffffffffffffffffffffffffffffff1614612cff576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612cf6906155c1565b60405180910390fd5b60016011600084815260200190815260200160002060006101000a81548160ff021916908315150217905550612d3481613bff565b601054600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fc7beef8f8d25068377c10e7b7f30ef5622f443eb674e07835df1daf5fe84acbf60045485600f5486604051612da994939291906155f0565b60405180910390a35050565b612dbe83610f0a565b612dc882826140be565b505050565b612dd5613f4a565b73ffffffffffffffffffffffffffffffffffffffff16612df36128bb565b73ffffffffffffffffffffffffffffffffffffffff1614612e49576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e4090614dc4565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612eb8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612eaf90614e30565b60405180910390fd5b80601b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f3ab2dcf539aa3c30314265602fc86dee8e343e3c3905730f870835f36c400deb81604051612f2891906143cc565b60405180910390a150565b612f3b613f4a565b73ffffffffffffffffffffffffffffffffffffffff16612f596128bb565b73ffffffffffffffffffffffffffffffffffffffff1614612faf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612fa690614dc4565b60405180910390fd5b80600e819055507f0cc7e37c68566f67d0fe13bf38246d7447cf99a0c481c2ef9963969bb4f5ebc8600e54604051612fe7919061444f565b60405180910390a150565b612ffa613f4a565b73ffffffffffffffffffffffffffffffffffffffff166130186128bb565b73ffffffffffffffffffffffffffffffffffffffff161461306e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161306590614dc4565b60405180910390fd5b606481106130b1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016130a890614ec2565b60405180910390fd5b80600b819055507fd2002a256ab2f8a5c1b865685754b20437c0a26e7249e40ae5df2993966f99f2600b546040516130e9919061444f565b60405180910390a150565b60006064600b5447613106919061544a565b613110919061548c565b905090565b600f5481565b68056bc75e2d6310000081565b601a54811061316c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161316390615688565b60405180910390fd5b6000601260006019600085815260200190815260200160002060020154815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff166131d9613f4a565b73ffffffffffffffffffffffffffffffffffffffff161461322f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016132269061571a565b60405180910390fd5b6019600083815260200190815260200160002060030160009054906101000a900460ff1615613293576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161328a906157ac565b60405180910390fd5b60016019600084815260200190815260200160002060030160006101000a81548160ff0219169083151502179055506019600083815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166342842e0e308360196000878152602001908152602001600020600101546040518463ffffffff1660e01b815260040161334b939291906157cc565b600060405180830381600087803b15801561336557600080fd5b505af1158015613379573d6000803e3d6000fd5b5050505060196000838152602001908152602001600020600201547f0d0e4b90a96d931964d5b0421a6b5b48dd73abf538cdb9ae015069d49e9a467983836019600087815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660196000888152602001908152602001600020600101546040516134159493929190615803565b60405180910390a25050565b60196020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154908060030160009054906101000a900460ff16905084565b60035481565b60135481565b600060646007544761349c919061544a565b6134a6919061548c565b905090565b600e5481565b6134b9613f4a565b73ffffffffffffffffffffffffffffffffffffffff166134d76128bb565b73ffffffffffffffffffffffffffffffffffffffff161461352d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161352490614dc4565b60405180910390fd5b80600d819055507f72e4278828b8a868e0ba5b4887f954797ea786f8bac991128636171f5eed471d600d54604051613565919061444f565b60405180910390a150565b60156020528060005260406000206000915090505481565b6000620f424060015460045461359e919061544a565b6135a8919061548c565b905090565b60085481565b600034116135f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016135ed906158ba565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361365557613654613fdd565b5b61365d613f4a565b73ffffffffffffffffffffffffffffffffffffffff167f8b7fe5be5699654fd637d2250cb0d47e88205730710745e78e9d8bcaf8aad8f1346040516136a2919061444f565b60405180910390a2565b6000601560006136ba613f4a565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205411613735576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161372c9061594c565b60405180910390fd5b600160156000613743613f4a565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461378c9190614b93565b925050819055506000601c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636a62784260e01b6137de613f4a565b6040516024016137ee91906143cc565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505060405161385891906150e2565b6000604051808303816000865af19150503d8060008114613895576040519150601f19603f3d011682016040523d82523d6000602084013e61389a565b606091505b50509050806138de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016138d59061516b565b60405180910390fd5b6138e6613f4a565b73ffffffffffffffffffffffffffffffffffffffff167fe05ba2c5fcd9a60f30b179cb0e775070cc8ce9667b0e663e984ee6a02f694cee60405160405180910390a250565b613933613f4a565b73ffffffffffffffffffffffffffffffffffffffff166139516128bb565b73ffffffffffffffffffffffffffffffffffffffff16146139a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161399e90614dc4565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603613a16576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613a0d906159de565b60405180910390fd5b613a1f81613ffa565b50565b6000426013541015613a375760009050613a48565b42601354613a459190614b93565b90505b90565b60015481565b613a59613f4a565b73ffffffffffffffffffffffffffffffffffffffff16613a776128bb565b73ffffffffffffffffffffffffffffffffffffffff1614613acd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613ac490614dc4565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603613b3c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613b3390614e30565b60405180910390fd5b80600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051613bce91906143cc565b60405180910390a150565b601b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b601354421015613c44576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613c3b90615a4a565b60405180910390fd5b61011881511115613c8a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613c8190615ab6565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603613cf55760085442613cee9190615072565b600f819055505b613cfd613f4a565b600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660166000601754815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160176000828254613dc69190615072565b925050819055506000601b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1960e01b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1668056bc75e2d63100000604051602401613e4e9291906153b5565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051613eb891906150e2565b6000604051808303816000865af19150503d8060008114613ef5576040519150601f19603f3d011682016040523d82523d6000602084013e613efa565b606091505b5050905080613f3e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613f3590615b48565b60405180910390fd5b613f466142a0565b5050565b600033905090565b6000613f5c613fa7565b6016600060175460145460001c613f739190615244565b815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b601454424340604051602001613fbf93929190615b68565b60405160208183030381529060405280519060200120601481905550565b600954613fe86126a1565b613ff2919061548c565b600481905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b8173ffffffffffffffffffffffffffffffffffffffff166342842e0e6140e2613f4a565b30846040518463ffffffff1660e01b8152600401614102939291906157cc565b600060405180830381600087803b15801561411c57600080fd5b505af1158015614130573d6000803e3d6000fd5b5050505060405180608001604052808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200160105481526020016000151581525060196000601a54815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101556040820151816002015560608201518160030160006101000a81548160ff0219169083151502179055509050506001601a60008282546142169190615072565b925050819055506010548273ffffffffffffffffffffffffffffffffffffffff1661423f613f4a565b73ffffffffffffffffffffffffffffffffffffffff167fc85be1734ed6a0f4d6adf924d4d41406e2729878c652110a96e2fdec64e118d1846001601a546142869190614b93565b604051614294929190615b9f565b60405180910390a45050565b6000633b9aca006002546142b4919061548c565b9050806142c3600f54426142fa565b6142cd9190615072565b600f81905550620f42406003546002546142e7919061544a565b6142f1919061548c565b60028190555050565b60008183101561430a578161430c565b825b905092915050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b61433b81614328565b811461434657600080fd5b50565b60008135905061435881614332565b92915050565b6000602082840312156143745761437361431e565b5b600061438284828501614349565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006143b68261438b565b9050919050565b6143c6816143ab565b82525050565b60006020820190506143e160008301846143bd565b92915050565b6143f0816143ab565b81146143fb57600080fd5b50565b60008135905061440d816143e7565b92915050565b6000602082840312156144295761442861431e565b5b6000614437848285016143fe565b91505092915050565b61444981614328565b82525050565b60006020820190506144646000830184614440565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f84011261448f5761448e61446a565b5b8235905067ffffffffffffffff8111156144ac576144ab61446f565b5b6020830191508360018202830111156144c8576144c7614474565b5b9250929050565b6000806000806000608086880312156144eb576144ea61431e565b5b60006144f9888289016143fe565b955050602061450a888289016143fe565b945050604061451b88828901614349565b935050606086013567ffffffffffffffff81111561453c5761453b614323565b5b61454888828901614479565b92509250509295509295909350565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b61458c81614557565b82525050565b60006020820190506145a76000830184614583565b92915050565b6000819050919050565b60006145d26145cd6145c88461438b565b6145ad565b61438b565b9050919050565b60006145e4826145b7565b9050919050565b60006145f6826145d9565b9050919050565b614606816145eb565b82525050565b600060208201905061462160008301846145fd565b92915050565b6000614632826145d9565b9050919050565b61464281614627565b82525050565b600060208201905061465d6000830184614639565b92915050565b600061466e826145d9565b9050919050565b61467e81614663565b82525050565b60006020820190506146996000830184614675565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6146ed826146a4565b810181811067ffffffffffffffff8211171561470c5761470b6146b5565b5b80604052505050565b600061471f614314565b905061472b82826146e4565b919050565b600067ffffffffffffffff82111561474b5761474a6146b5565b5b614754826146a4565b9050602081019050919050565b82818337600083830152505050565b600061478361477e84614730565b614715565b90508281526020810184848401111561479f5761479e61469f565b5b6147aa848285614761565b509392505050565b600082601f8301126147c7576147c661446a565b5b81356147d7848260208601614770565b91505092915050565b6000602082840312156147f6576147f561431e565b5b600082013567ffffffffffffffff81111561481457614813614323565b5b614820848285016147b2565b91505092915050565b6000819050919050565b61483c81614829565b82525050565b60006020820190506148576000830184614833565b92915050565b60008115159050919050565b6148728161485d565b82525050565b600060208201905061488d6000830184614869565b92915050565b600061489e826143ab565b9050919050565b6148ae81614893565b81146148b957600080fd5b50565b6000813590506148cb816148a5565b92915050565b600080600080608085870312156148eb576148ea61431e565b5b60006148f987828801614349565b945050602085013567ffffffffffffffff81111561491a57614919614323565b5b614926878288016147b2565b9350506040614937878288016148bc565b925050606061494887828801614349565b91505092959194509250565b6000806040838503121561496b5761496a61431e565b5b600061497985828601614349565b925050602083013567ffffffffffffffff81111561499a57614999614323565b5b6149a6858286016147b2565b9150509250929050565b6000806000606084860312156149c9576149c861431e565b5b600084013567ffffffffffffffff8111156149e7576149e6614323565b5b6149f3868287016147b2565b9350506020614a04868287016148bc565b9250506040614a1586828701614349565b9150509250925092565b6000614a2a826145d9565b9050919050565b614a3a81614a1f565b82525050565b6000608082019050614a556000830187614a31565b614a626020830186614440565b614a6f6040830185614440565b614a7c6060830184614869565b95945050505050565b6000614a90826145d9565b9050919050565b614aa081614a85565b82525050565b6000602082019050614abb6000830184614a97565b92915050565b600082825260208201905092915050565b7f5468652076616c7565207375626d69747465642077697468207468697320747260008201527f616e73616374696f6e20697320746f6f206c6f772e0000000000000000000000602082015250565b6000614b2e603583614ac1565b9150614b3982614ad2565b604082019050919050565b60006020820190508181036000830152614b5d81614b21565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000614b9e82614328565b9150614ba983614328565b9250828203905081811115614bc157614bc0614b64565b5b92915050565b600081905092915050565b50565b6000614be2600083614bc7565b9150614bed82614bd2565b600082019050919050565b6000614c0382614bd5565b9150819050919050565b7f526566756e64207472616e73666572206661696c65642e000000000000000000600082015250565b6000614c43601783614ac1565b9150614c4e82614c0d565b602082019050919050565b60006020820190508181036000830152614c7281614c36565b9050919050565b6000819050919050565b6000819050919050565b6000614ca8614ca3614c9e84614c79565b6145ad565b614c83565b9050919050565b614cb881614c8d565b82525050565b600081519050919050565b60005b83811015614ce7578082015181840152602081019050614ccc565b60008484015250505050565b6000614cfe82614cbe565b614d088185614ac1565b9350614d18818560208601614cc9565b614d21816146a4565b840191505092915050565b6000608082019050614d416000830187614440565b614d4e6020830186614caf565b614d5b6040830185614440565b8181036060830152614d6d8184614cf3565b905095945050505050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000614dae602083614ac1565b9150614db982614d78565b602082019050919050565b60006020820190508181036000830152614ddd81614da1565b9050919050565b7f5a65726f2d616464726573732077617320676976656e2e000000000000000000600082015250565b6000614e1a601783614ac1565b9150614e2582614de4565b602082019050919050565b60006020820190508181036000830152614e4981614e0d565b9050919050565b7f50657263656e746167652076616c7565206f766572666c6f772c206d7573742060008201527f6265206c6f776572207468616e203130302e0000000000000000000000000000602082015250565b6000614eac603283614ac1565b9150614eb782614e50565b604082019050919050565b60006020820190508181036000830152614edb81614e9f565b9050919050565b7f4e6f7420656e6f7567682074696d652068617320656c61707365642e00000000600082015250565b6000614f18601c83614ac1565b9150614f2382614ee2565b602082019050919050565b60006020820190508181036000830152614f4781614f0b565b9050919050565b7f5468657265206973206e6f206c617374206269646465722e0000000000000000600082015250565b6000614f84601883614ac1565b9150614f8f82614f4e565b602082019050919050565b60006020820190508181036000830152614fb381614f77565b9050919050565b7f4f6e6c7920746865206c617374206269646465722063616e20636c61696d207460008201527f6865207072697a6520647572696e672074686520666972737420323420686f7560208201527f72732e0000000000000000000000000000000000000000000000000000000000604082015250565b600061503c604383614ac1565b915061504782614fba565b606082019050919050565b6000602082019050818103600083015261506b8161502f565b9050919050565b600061507d82614328565b915061508883614328565b92508282019050808211156150a05761509f614b64565b5b92915050565b600081519050919050565b60006150bc826150a6565b6150c68185614bc7565b93506150d6818560208601614cc9565b80840191505092915050565b60006150ee82846150b1565b915081905092915050565b7f436f736d69635369676e6174757265206d696e742829206661696c656420746f60008201527f206d696e74204e46542e00000000000000000000000000000000000000000000602082015250565b6000615155602a83614ac1565b9150615160826150f9565b604082019050919050565b6000602082019050818103600083015261518481615148565b9050919050565b600061519682614328565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036151c8576151c7614b64565b5b600182019050919050565b6000815190506151e281614332565b92915050565b6000602082840312156151fe576151fd61431e565b5b600061520c848285016151d3565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061524f82614328565b915061525a83614328565b92508261526a57615269615215565b5b828206905092915050565b600081519050615284816143e7565b92915050565b6000602082840312156152a05761529f61431e565b5b60006152ae84828501615275565b91505092915050565b7f5472616e7366657220746f207468652077696e6e6572206661696c65642e0000600082015250565b60006152ed601e83614ac1565b91506152f8826152b7565b602082019050919050565b6000602082019050818103600083015261531c816152e0565b9050919050565b7f5472616e7366657220746f206368617269747920636f6e74726163742066616960008201527f6c65642e00000000000000000000000000000000000000000000000000000000602082015250565b600061537f602483614ac1565b915061538a82615323565b604082019050919050565b600060208201905081810360008301526153ae81615372565b9050919050565b60006040820190506153ca60008301856143bd565b6153d76020830184614440565b9392505050565b7f526166666c65206465706f736974206661696c65642e00000000000000000000600082015250565b6000615414601683614ac1565b915061541f826153de565b602082019050919050565b6000602082019050818103600083015261544381615407565b9050919050565b600061545582614328565b915061546083614328565b925082820261546e81614328565b9150828204841483151761548557615484614b64565b5b5092915050565b600061549782614328565b91506154a283614328565b9250826154b2576154b1615215565b5b828204905092915050565b7f546869732052616e646f6d57616c6b4e46542068617320616c7265616479206260008201527f65656e207573656420666f722062696464696e672e0000000000000000000000602082015250565b6000615519603583614ac1565b9150615524826154bd565b604082019050919050565b600060208201905081810360008301526155488161550c565b9050919050565b7f596f75206d75737420626520746865206f776e6572206f66207468652052616e60008201527f646f6d57616c6b4e46542e000000000000000000000000000000000000000000602082015250565b60006155ab602b83614ac1565b91506155b68261554f565b604082019050919050565b600060208201905081810360008301526155da8161559e565b9050919050565b6155ea81614c83565b82525050565b60006080820190506156056000830187614440565b61561260208301866155e1565b61561f6040830185614440565b81810360608301526156318184614cf3565b905095945050505050565b7f54686520646f6e61746564204e465420646f6573206e6f742065786973742e00600082015250565b6000615672601f83614ac1565b915061567d8261563c565b602082019050919050565b600060208201905081810360008301526156a181615665565b9050919050565b7f596f7520617265206e6f74207468652077696e6e6572206f662074686520726f60008201527f756e642e00000000000000000000000000000000000000000000000000000000602082015250565b6000615704602483614ac1565b915061570f826156a8565b604082019050919050565b60006020820190508181036000830152615733816156f7565b9050919050565b7f546865204e46542068617320616c7265616479206265656e20636c61696d656460008201527f2e00000000000000000000000000000000000000000000000000000000000000602082015250565b6000615796602183614ac1565b91506157a18261573a565b604082019050919050565b600060208201905081810360008301526157c581615789565b9050919050565b60006060820190506157e160008301866143bd565b6157ee60208301856143bd565b6157fb6040830184614440565b949350505050565b60006080820190506158186000830187614440565b61582560208301866143bd565b61583260408301856143bd565b61583f6060830184614440565b95945050505050565b7f446f6e6174696f6e20616d6f756e74206d75737420626520677265617465722060008201527f7468616e20302e00000000000000000000000000000000000000000000000000602082015250565b60006158a4602783614ac1565b91506158af82615848565b604082019050919050565b600060208201905081810360008301526158d381615897565b9050919050565b7f596f752068617665206e6f20756e636c61696d656420726166666c65204e465460008201527f732e000000000000000000000000000000000000000000000000000000000000602082015250565b6000615936602283614ac1565b9150615941826158da565b604082019050919050565b6000602082019050818103600083015261596581615929565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b60006159c8602683614ac1565b91506159d38261596c565b604082019050919050565b600060208201905081810360008301526159f7816159bb565b9050919050565b7f4e6f7420616374697665207965742e0000000000000000000000000000000000600082015250565b6000615a34600f83614ac1565b9150615a3f826159fe565b602082019050919050565b60006020820190508181036000830152615a6381615a27565b9050919050565b7f4d65737361676520697320746f6f206c6f6e672e000000000000000000000000600082015250565b6000615aa0601483614ac1565b9150615aab82615a6a565b602082019050919050565b60006020820190508181036000830152615acf81615a93565b9050919050565b7f436f736d6963546f6b656e206d696e742829206661696c656420746f206d696e60008201527f742072657761726420746f6b656e732e00000000000000000000000000000000602082015250565b6000615b32603083614ac1565b9150615b3d82615ad6565b604082019050919050565b60006020820190508181036000830152615b6181615b25565b9050919050565b6000606082019050615b7d6000830186614833565b615b8a6020830185614440565b615b976040830184614833565b949350505050565b6000604082019050615bb46000830185614440565b615bc16020830184614440565b939250505056fea2646970667358221220976ef97e4f2755e16750ab566311812c8d03b6431b021968d0aff257c0018d8e64736f6c63430008130033",
}

// CosmicGameABI is the input ABI used to generate the binding from.
// Deprecated: Use CosmicGameMetaData.ABI instead.
var CosmicGameABI = CosmicGameMetaData.ABI

// CosmicGameBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CosmicGameMetaData.Bin instead.
var CosmicGameBin = CosmicGameMetaData.Bin

// DeployCosmicGame deploys a new Ethereum contract, binding an instance of CosmicGame to it.
func DeployCosmicGame(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CosmicGame, error) {
	parsed, err := CosmicGameMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CosmicGameBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CosmicGame{CosmicGameCaller: CosmicGameCaller{contract: contract}, CosmicGameTransactor: CosmicGameTransactor{contract: contract}, CosmicGameFilterer: CosmicGameFilterer{contract: contract}}, nil
}

// CosmicGame is an auto generated Go binding around an Ethereum contract.
type CosmicGame struct {
	CosmicGameCaller     // Read-only binding to the contract
	CosmicGameTransactor // Write-only binding to the contract
	CosmicGameFilterer   // Log filterer for contract events
}

// CosmicGameCaller is an auto generated read-only Go binding around an Ethereum contract.
type CosmicGameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmicGameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CosmicGameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmicGameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CosmicGameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmicGameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CosmicGameSession struct {
	Contract     *CosmicGame       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CosmicGameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CosmicGameCallerSession struct {
	Contract *CosmicGameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CosmicGameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CosmicGameTransactorSession struct {
	Contract     *CosmicGameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CosmicGameRaw is an auto generated low-level Go binding around an Ethereum contract.
type CosmicGameRaw struct {
	Contract *CosmicGame // Generic contract binding to access the raw methods on
}

// CosmicGameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CosmicGameCallerRaw struct {
	Contract *CosmicGameCaller // Generic read-only contract binding to access the raw methods on
}

// CosmicGameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CosmicGameTransactorRaw struct {
	Contract *CosmicGameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCosmicGame creates a new instance of CosmicGame, bound to a specific deployed contract.
func NewCosmicGame(address common.Address, backend bind.ContractBackend) (*CosmicGame, error) {
	contract, err := bindCosmicGame(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CosmicGame{CosmicGameCaller: CosmicGameCaller{contract: contract}, CosmicGameTransactor: CosmicGameTransactor{contract: contract}, CosmicGameFilterer: CosmicGameFilterer{contract: contract}}, nil
}

// NewCosmicGameCaller creates a new read-only instance of CosmicGame, bound to a specific deployed contract.
func NewCosmicGameCaller(address common.Address, caller bind.ContractCaller) (*CosmicGameCaller, error) {
	contract, err := bindCosmicGame(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CosmicGameCaller{contract: contract}, nil
}

// NewCosmicGameTransactor creates a new write-only instance of CosmicGame, bound to a specific deployed contract.
func NewCosmicGameTransactor(address common.Address, transactor bind.ContractTransactor) (*CosmicGameTransactor, error) {
	contract, err := bindCosmicGame(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CosmicGameTransactor{contract: contract}, nil
}

// NewCosmicGameFilterer creates a new log filterer instance of CosmicGame, bound to a specific deployed contract.
func NewCosmicGameFilterer(address common.Address, filterer bind.ContractFilterer) (*CosmicGameFilterer, error) {
	contract, err := bindCosmicGame(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CosmicGameFilterer{contract: contract}, nil
}

// bindCosmicGame binds a generic wrapper to an already deployed contract.
func bindCosmicGame(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CosmicGameABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CosmicGame *CosmicGameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CosmicGame.Contract.CosmicGameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CosmicGame *CosmicGameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicGame.Contract.CosmicGameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CosmicGame *CosmicGameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CosmicGame.Contract.CosmicGameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CosmicGame *CosmicGameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CosmicGame.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CosmicGame *CosmicGameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicGame.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CosmicGame *CosmicGameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CosmicGame.Contract.contract.Transact(opts, method, params...)
}

// MAXMESSAGELENGTH is a free data retrieval call binding the contract method 0x2a027211.
//
// Solidity: function MAX_MESSAGE_LENGTH() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) MAXMESSAGELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "MAX_MESSAGE_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXMESSAGELENGTH is a free data retrieval call binding the contract method 0x2a027211.
//
// Solidity: function MAX_MESSAGE_LENGTH() view returns(uint256)
func (_CosmicGame *CosmicGameSession) MAXMESSAGELENGTH() (*big.Int, error) {
	return _CosmicGame.Contract.MAXMESSAGELENGTH(&_CosmicGame.CallOpts)
}

// MAXMESSAGELENGTH is a free data retrieval call binding the contract method 0x2a027211.
//
// Solidity: function MAX_MESSAGE_LENGTH() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) MAXMESSAGELENGTH() (*big.Int, error) {
	return _CosmicGame.Contract.MAXMESSAGELENGTH(&_CosmicGame.CallOpts)
}

// MILLION is a free data retrieval call binding the contract method 0x32bc934c.
//
// Solidity: function MILLION() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) MILLION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "MILLION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MILLION is a free data retrieval call binding the contract method 0x32bc934c.
//
// Solidity: function MILLION() view returns(uint256)
func (_CosmicGame *CosmicGameSession) MILLION() (*big.Int, error) {
	return _CosmicGame.Contract.MILLION(&_CosmicGame.CallOpts)
}

// MILLION is a free data retrieval call binding the contract method 0x32bc934c.
//
// Solidity: function MILLION() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) MILLION() (*big.Int, error) {
	return _CosmicGame.Contract.MILLION(&_CosmicGame.CallOpts)
}

// TOKENREWARD is a free data retrieval call binding the contract method 0xd4ba8902.
//
// Solidity: function TOKEN_REWARD() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) TOKENREWARD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "TOKEN_REWARD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOKENREWARD is a free data retrieval call binding the contract method 0xd4ba8902.
//
// Solidity: function TOKEN_REWARD() view returns(uint256)
func (_CosmicGame *CosmicGameSession) TOKENREWARD() (*big.Int, error) {
	return _CosmicGame.Contract.TOKENREWARD(&_CosmicGame.CallOpts)
}

// TOKENREWARD is a free data retrieval call binding the contract method 0xd4ba8902.
//
// Solidity: function TOKEN_REWARD() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) TOKENREWARD() (*big.Int, error) {
	return _CosmicGame.Contract.TOKENREWARD(&_CosmicGame.CallOpts)
}

// ActivationTime is a free data retrieval call binding the contract method 0xda4493f6.
//
// Solidity: function activationTime() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) ActivationTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "activationTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActivationTime is a free data retrieval call binding the contract method 0xda4493f6.
//
// Solidity: function activationTime() view returns(uint256)
func (_CosmicGame *CosmicGameSession) ActivationTime() (*big.Int, error) {
	return _CosmicGame.Contract.ActivationTime(&_CosmicGame.CallOpts)
}

// ActivationTime is a free data retrieval call binding the contract method 0xda4493f6.
//
// Solidity: function activationTime() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) ActivationTime() (*big.Int, error) {
	return _CosmicGame.Contract.ActivationTime(&_CosmicGame.CallOpts)
}

// BidPrice is a free data retrieval call binding the contract method 0x19afe473.
//
// Solidity: function bidPrice() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) BidPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "bidPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BidPrice is a free data retrieval call binding the contract method 0x19afe473.
//
// Solidity: function bidPrice() view returns(uint256)
func (_CosmicGame *CosmicGameSession) BidPrice() (*big.Int, error) {
	return _CosmicGame.Contract.BidPrice(&_CosmicGame.CallOpts)
}

// BidPrice is a free data retrieval call binding the contract method 0x19afe473.
//
// Solidity: function bidPrice() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) BidPrice() (*big.Int, error) {
	return _CosmicGame.Contract.BidPrice(&_CosmicGame.CallOpts)
}

// Charity is a free data retrieval call binding the contract method 0x934aa023.
//
// Solidity: function charity() view returns(address)
func (_CosmicGame *CosmicGameCaller) Charity(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "charity")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Charity is a free data retrieval call binding the contract method 0x934aa023.
//
// Solidity: function charity() view returns(address)
func (_CosmicGame *CosmicGameSession) Charity() (common.Address, error) {
	return _CosmicGame.Contract.Charity(&_CosmicGame.CallOpts)
}

// Charity is a free data retrieval call binding the contract method 0x934aa023.
//
// Solidity: function charity() view returns(address)
func (_CosmicGame *CosmicGameCallerSession) Charity() (common.Address, error) {
	return _CosmicGame.Contract.Charity(&_CosmicGame.CallOpts)
}

// CharityAmount is a free data retrieval call binding the contract method 0xdbc945c0.
//
// Solidity: function charityAmount() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) CharityAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "charityAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CharityAmount is a free data retrieval call binding the contract method 0xdbc945c0.
//
// Solidity: function charityAmount() view returns(uint256)
func (_CosmicGame *CosmicGameSession) CharityAmount() (*big.Int, error) {
	return _CosmicGame.Contract.CharityAmount(&_CosmicGame.CallOpts)
}

// CharityAmount is a free data retrieval call binding the contract method 0xdbc945c0.
//
// Solidity: function charityAmount() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) CharityAmount() (*big.Int, error) {
	return _CosmicGame.Contract.CharityAmount(&_CosmicGame.CallOpts)
}

// CharityPercentage is a free data retrieval call binding the contract method 0x799d431d.
//
// Solidity: function charityPercentage() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) CharityPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "charityPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CharityPercentage is a free data retrieval call binding the contract method 0x799d431d.
//
// Solidity: function charityPercentage() view returns(uint256)
func (_CosmicGame *CosmicGameSession) CharityPercentage() (*big.Int, error) {
	return _CosmicGame.Contract.CharityPercentage(&_CosmicGame.CallOpts)
}

// CharityPercentage is a free data retrieval call binding the contract method 0x799d431d.
//
// Solidity: function charityPercentage() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) CharityPercentage() (*big.Int, error) {
	return _CosmicGame.Contract.CharityPercentage(&_CosmicGame.CallOpts)
}

// DonatedNFTs is a free data retrieval call binding the contract method 0xd6e17417.
//
// Solidity: function donatedNFTs(uint256 ) view returns(address nftAddress, uint256 tokenId, uint256 round, bool claimed)
func (_CosmicGame *CosmicGameCaller) DonatedNFTs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	NftAddress common.Address
	TokenId    *big.Int
	Round      *big.Int
	Claimed    bool
}, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "donatedNFTs", arg0)

	outstruct := new(struct {
		NftAddress common.Address
		TokenId    *big.Int
		Round      *big.Int
		Claimed    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NftAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Round = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Claimed = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// DonatedNFTs is a free data retrieval call binding the contract method 0xd6e17417.
//
// Solidity: function donatedNFTs(uint256 ) view returns(address nftAddress, uint256 tokenId, uint256 round, bool claimed)
func (_CosmicGame *CosmicGameSession) DonatedNFTs(arg0 *big.Int) (struct {
	NftAddress common.Address
	TokenId    *big.Int
	Round      *big.Int
	Claimed    bool
}, error) {
	return _CosmicGame.Contract.DonatedNFTs(&_CosmicGame.CallOpts, arg0)
}

// DonatedNFTs is a free data retrieval call binding the contract method 0xd6e17417.
//
// Solidity: function donatedNFTs(uint256 ) view returns(address nftAddress, uint256 tokenId, uint256 round, bool claimed)
func (_CosmicGame *CosmicGameCallerSession) DonatedNFTs(arg0 *big.Int) (struct {
	NftAddress common.Address
	TokenId    *big.Int
	Round      *big.Int
	Claimed    bool
}, error) {
	return _CosmicGame.Contract.DonatedNFTs(&_CosmicGame.CallOpts, arg0)
}

// GetBidPrice is a free data retrieval call binding the contract method 0xec34866d.
//
// Solidity: function getBidPrice() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) GetBidPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "getBidPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBidPrice is a free data retrieval call binding the contract method 0xec34866d.
//
// Solidity: function getBidPrice() view returns(uint256)
func (_CosmicGame *CosmicGameSession) GetBidPrice() (*big.Int, error) {
	return _CosmicGame.Contract.GetBidPrice(&_CosmicGame.CallOpts)
}

// GetBidPrice is a free data retrieval call binding the contract method 0xec34866d.
//
// Solidity: function getBidPrice() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) GetBidPrice() (*big.Int, error) {
	return _CosmicGame.Contract.GetBidPrice(&_CosmicGame.CallOpts)
}

// InitialBidAmountFraction is a free data retrieval call binding the contract method 0x40e02322.
//
// Solidity: function initialBidAmountFraction() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) InitialBidAmountFraction(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "initialBidAmountFraction")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialBidAmountFraction is a free data retrieval call binding the contract method 0x40e02322.
//
// Solidity: function initialBidAmountFraction() view returns(uint256)
func (_CosmicGame *CosmicGameSession) InitialBidAmountFraction() (*big.Int, error) {
	return _CosmicGame.Contract.InitialBidAmountFraction(&_CosmicGame.CallOpts)
}

// InitialBidAmountFraction is a free data retrieval call binding the contract method 0x40e02322.
//
// Solidity: function initialBidAmountFraction() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) InitialBidAmountFraction() (*big.Int, error) {
	return _CosmicGame.Contract.InitialBidAmountFraction(&_CosmicGame.CallOpts)
}

// InitialSecondsUntilPrize is a free data retrieval call binding the contract method 0xed083398.
//
// Solidity: function initialSecondsUntilPrize() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) InitialSecondsUntilPrize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "initialSecondsUntilPrize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialSecondsUntilPrize is a free data retrieval call binding the contract method 0xed083398.
//
// Solidity: function initialSecondsUntilPrize() view returns(uint256)
func (_CosmicGame *CosmicGameSession) InitialSecondsUntilPrize() (*big.Int, error) {
	return _CosmicGame.Contract.InitialSecondsUntilPrize(&_CosmicGame.CallOpts)
}

// InitialSecondsUntilPrize is a free data retrieval call binding the contract method 0xed083398.
//
// Solidity: function initialSecondsUntilPrize() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) InitialSecondsUntilPrize() (*big.Int, error) {
	return _CosmicGame.Contract.InitialSecondsUntilPrize(&_CosmicGame.CallOpts)
}

// LastBidder is a free data retrieval call binding the contract method 0x8547af30.
//
// Solidity: function lastBidder() view returns(address)
func (_CosmicGame *CosmicGameCaller) LastBidder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "lastBidder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LastBidder is a free data retrieval call binding the contract method 0x8547af30.
//
// Solidity: function lastBidder() view returns(address)
func (_CosmicGame *CosmicGameSession) LastBidder() (common.Address, error) {
	return _CosmicGame.Contract.LastBidder(&_CosmicGame.CallOpts)
}

// LastBidder is a free data retrieval call binding the contract method 0x8547af30.
//
// Solidity: function lastBidder() view returns(address)
func (_CosmicGame *CosmicGameCallerSession) LastBidder() (common.Address, error) {
	return _CosmicGame.Contract.LastBidder(&_CosmicGame.CallOpts)
}

// NanoSecondsExtra is a free data retrieval call binding the contract method 0x9136d6d9.
//
// Solidity: function nanoSecondsExtra() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) NanoSecondsExtra(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "nanoSecondsExtra")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NanoSecondsExtra is a free data retrieval call binding the contract method 0x9136d6d9.
//
// Solidity: function nanoSecondsExtra() view returns(uint256)
func (_CosmicGame *CosmicGameSession) NanoSecondsExtra() (*big.Int, error) {
	return _CosmicGame.Contract.NanoSecondsExtra(&_CosmicGame.CallOpts)
}

// NanoSecondsExtra is a free data retrieval call binding the contract method 0x9136d6d9.
//
// Solidity: function nanoSecondsExtra() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) NanoSecondsExtra() (*big.Int, error) {
	return _CosmicGame.Contract.NanoSecondsExtra(&_CosmicGame.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_CosmicGame *CosmicGameCaller) Nft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "nft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_CosmicGame *CosmicGameSession) Nft() (common.Address, error) {
	return _CosmicGame.Contract.Nft(&_CosmicGame.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_CosmicGame *CosmicGameCallerSession) Nft() (common.Address, error) {
	return _CosmicGame.Contract.Nft(&_CosmicGame.CallOpts)
}

// NumDonatedNFTs is a free data retrieval call binding the contract method 0x2aab3223.
//
// Solidity: function numDonatedNFTs() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) NumDonatedNFTs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "numDonatedNFTs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumDonatedNFTs is a free data retrieval call binding the contract method 0x2aab3223.
//
// Solidity: function numDonatedNFTs() view returns(uint256)
func (_CosmicGame *CosmicGameSession) NumDonatedNFTs() (*big.Int, error) {
	return _CosmicGame.Contract.NumDonatedNFTs(&_CosmicGame.CallOpts)
}

// NumDonatedNFTs is a free data retrieval call binding the contract method 0x2aab3223.
//
// Solidity: function numDonatedNFTs() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) NumDonatedNFTs() (*big.Int, error) {
	return _CosmicGame.Contract.NumDonatedNFTs(&_CosmicGame.CallOpts)
}

// NumHolderNFTWinnersPerRound is a free data retrieval call binding the contract method 0xdec08d8e.
//
// Solidity: function numHolderNFTWinnersPerRound() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) NumHolderNFTWinnersPerRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "numHolderNFTWinnersPerRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumHolderNFTWinnersPerRound is a free data retrieval call binding the contract method 0xdec08d8e.
//
// Solidity: function numHolderNFTWinnersPerRound() view returns(uint256)
func (_CosmicGame *CosmicGameSession) NumHolderNFTWinnersPerRound() (*big.Int, error) {
	return _CosmicGame.Contract.NumHolderNFTWinnersPerRound(&_CosmicGame.CallOpts)
}

// NumHolderNFTWinnersPerRound is a free data retrieval call binding the contract method 0xdec08d8e.
//
// Solidity: function numHolderNFTWinnersPerRound() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) NumHolderNFTWinnersPerRound() (*big.Int, error) {
	return _CosmicGame.Contract.NumHolderNFTWinnersPerRound(&_CosmicGame.CallOpts)
}

// NumRaffleNFTWinnersPerRound is a free data retrieval call binding the contract method 0xa6ceac2c.
//
// Solidity: function numRaffleNFTWinnersPerRound() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) NumRaffleNFTWinnersPerRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "numRaffleNFTWinnersPerRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumRaffleNFTWinnersPerRound is a free data retrieval call binding the contract method 0xa6ceac2c.
//
// Solidity: function numRaffleNFTWinnersPerRound() view returns(uint256)
func (_CosmicGame *CosmicGameSession) NumRaffleNFTWinnersPerRound() (*big.Int, error) {
	return _CosmicGame.Contract.NumRaffleNFTWinnersPerRound(&_CosmicGame.CallOpts)
}

// NumRaffleNFTWinnersPerRound is a free data retrieval call binding the contract method 0xa6ceac2c.
//
// Solidity: function numRaffleNFTWinnersPerRound() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) NumRaffleNFTWinnersPerRound() (*big.Int, error) {
	return _CosmicGame.Contract.NumRaffleNFTWinnersPerRound(&_CosmicGame.CallOpts)
}

// NumRaffleParticipants is a free data retrieval call binding the contract method 0x3bec7b69.
//
// Solidity: function numRaffleParticipants() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) NumRaffleParticipants(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "numRaffleParticipants")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumRaffleParticipants is a free data retrieval call binding the contract method 0x3bec7b69.
//
// Solidity: function numRaffleParticipants() view returns(uint256)
func (_CosmicGame *CosmicGameSession) NumRaffleParticipants() (*big.Int, error) {
	return _CosmicGame.Contract.NumRaffleParticipants(&_CosmicGame.CallOpts)
}

// NumRaffleParticipants is a free data retrieval call binding the contract method 0x3bec7b69.
//
// Solidity: function numRaffleParticipants() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) NumRaffleParticipants() (*big.Int, error) {
	return _CosmicGame.Contract.NumRaffleParticipants(&_CosmicGame.CallOpts)
}

// NumRaffleWinnersPerRound is a free data retrieval call binding the contract method 0x4a773f33.
//
// Solidity: function numRaffleWinnersPerRound() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) NumRaffleWinnersPerRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "numRaffleWinnersPerRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumRaffleWinnersPerRound is a free data retrieval call binding the contract method 0x4a773f33.
//
// Solidity: function numRaffleWinnersPerRound() view returns(uint256)
func (_CosmicGame *CosmicGameSession) NumRaffleWinnersPerRound() (*big.Int, error) {
	return _CosmicGame.Contract.NumRaffleWinnersPerRound(&_CosmicGame.CallOpts)
}

// NumRaffleWinnersPerRound is a free data retrieval call binding the contract method 0x4a773f33.
//
// Solidity: function numRaffleWinnersPerRound() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) NumRaffleWinnersPerRound() (*big.Int, error) {
	return _CosmicGame.Contract.NumRaffleWinnersPerRound(&_CosmicGame.CallOpts)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_CosmicGame *CosmicGameCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_CosmicGame *CosmicGameSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _CosmicGame.Contract.OnERC721Received(&_CosmicGame.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_CosmicGame *CosmicGameCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _CosmicGame.Contract.OnERC721Received(&_CosmicGame.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CosmicGame *CosmicGameCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CosmicGame *CosmicGameSession) Owner() (common.Address, error) {
	return _CosmicGame.Contract.Owner(&_CosmicGame.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CosmicGame *CosmicGameCallerSession) Owner() (common.Address, error) {
	return _CosmicGame.Contract.Owner(&_CosmicGame.CallOpts)
}

// PriceIncrease is a free data retrieval call binding the contract method 0xf8c34050.
//
// Solidity: function priceIncrease() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) PriceIncrease(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "priceIncrease")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceIncrease is a free data retrieval call binding the contract method 0xf8c34050.
//
// Solidity: function priceIncrease() view returns(uint256)
func (_CosmicGame *CosmicGameSession) PriceIncrease() (*big.Int, error) {
	return _CosmicGame.Contract.PriceIncrease(&_CosmicGame.CallOpts)
}

// PriceIncrease is a free data retrieval call binding the contract method 0xf8c34050.
//
// Solidity: function priceIncrease() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) PriceIncrease() (*big.Int, error) {
	return _CosmicGame.Contract.PriceIncrease(&_CosmicGame.CallOpts)
}

// PrizeAmount is a free data retrieval call binding the contract method 0x785fa627.
//
// Solidity: function prizeAmount() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) PrizeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "prizeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrizeAmount is a free data retrieval call binding the contract method 0x785fa627.
//
// Solidity: function prizeAmount() view returns(uint256)
func (_CosmicGame *CosmicGameSession) PrizeAmount() (*big.Int, error) {
	return _CosmicGame.Contract.PrizeAmount(&_CosmicGame.CallOpts)
}

// PrizeAmount is a free data retrieval call binding the contract method 0x785fa627.
//
// Solidity: function prizeAmount() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) PrizeAmount() (*big.Int, error) {
	return _CosmicGame.Contract.PrizeAmount(&_CosmicGame.CallOpts)
}

// PrizePercentage is a free data retrieval call binding the contract method 0x11dc7335.
//
// Solidity: function prizePercentage() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) PrizePercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "prizePercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrizePercentage is a free data retrieval call binding the contract method 0x11dc7335.
//
// Solidity: function prizePercentage() view returns(uint256)
func (_CosmicGame *CosmicGameSession) PrizePercentage() (*big.Int, error) {
	return _CosmicGame.Contract.PrizePercentage(&_CosmicGame.CallOpts)
}

// PrizePercentage is a free data retrieval call binding the contract method 0x11dc7335.
//
// Solidity: function prizePercentage() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) PrizePercentage() (*big.Int, error) {
	return _CosmicGame.Contract.PrizePercentage(&_CosmicGame.CallOpts)
}

// PrizeTime is a free data retrieval call binding the contract method 0xcb819dc0.
//
// Solidity: function prizeTime() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) PrizeTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "prizeTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrizeTime is a free data retrieval call binding the contract method 0xcb819dc0.
//
// Solidity: function prizeTime() view returns(uint256)
func (_CosmicGame *CosmicGameSession) PrizeTime() (*big.Int, error) {
	return _CosmicGame.Contract.PrizeTime(&_CosmicGame.CallOpts)
}

// PrizeTime is a free data retrieval call binding the contract method 0xcb819dc0.
//
// Solidity: function prizeTime() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) PrizeTime() (*big.Int, error) {
	return _CosmicGame.Contract.PrizeTime(&_CosmicGame.CallOpts)
}

// RaffleAmount is a free data retrieval call binding the contract method 0xc94028c2.
//
// Solidity: function raffleAmount() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) RaffleAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "raffleAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RaffleAmount is a free data retrieval call binding the contract method 0xc94028c2.
//
// Solidity: function raffleAmount() view returns(uint256)
func (_CosmicGame *CosmicGameSession) RaffleAmount() (*big.Int, error) {
	return _CosmicGame.Contract.RaffleAmount(&_CosmicGame.CallOpts)
}

// RaffleAmount is a free data retrieval call binding the contract method 0xc94028c2.
//
// Solidity: function raffleAmount() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) RaffleAmount() (*big.Int, error) {
	return _CosmicGame.Contract.RaffleAmount(&_CosmicGame.CallOpts)
}

// RaffleEntropy is a free data retrieval call binding the contract method 0x80de163d.
//
// Solidity: function raffleEntropy() view returns(bytes32)
func (_CosmicGame *CosmicGameCaller) RaffleEntropy(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "raffleEntropy")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RaffleEntropy is a free data retrieval call binding the contract method 0x80de163d.
//
// Solidity: function raffleEntropy() view returns(bytes32)
func (_CosmicGame *CosmicGameSession) RaffleEntropy() ([32]byte, error) {
	return _CosmicGame.Contract.RaffleEntropy(&_CosmicGame.CallOpts)
}

// RaffleEntropy is a free data retrieval call binding the contract method 0x80de163d.
//
// Solidity: function raffleEntropy() view returns(bytes32)
func (_CosmicGame *CosmicGameCallerSession) RaffleEntropy() ([32]byte, error) {
	return _CosmicGame.Contract.RaffleEntropy(&_CosmicGame.CallOpts)
}

// RaffleNFTWinners is a free data retrieval call binding the contract method 0xebc3f1d1.
//
// Solidity: function raffleNFTWinners(address ) view returns(uint256)
func (_CosmicGame *CosmicGameCaller) RaffleNFTWinners(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "raffleNFTWinners", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RaffleNFTWinners is a free data retrieval call binding the contract method 0xebc3f1d1.
//
// Solidity: function raffleNFTWinners(address ) view returns(uint256)
func (_CosmicGame *CosmicGameSession) RaffleNFTWinners(arg0 common.Address) (*big.Int, error) {
	return _CosmicGame.Contract.RaffleNFTWinners(&_CosmicGame.CallOpts, arg0)
}

// RaffleNFTWinners is a free data retrieval call binding the contract method 0xebc3f1d1.
//
// Solidity: function raffleNFTWinners(address ) view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) RaffleNFTWinners(arg0 common.Address) (*big.Int, error) {
	return _CosmicGame.Contract.RaffleNFTWinners(&_CosmicGame.CallOpts, arg0)
}

// RaffleParticipants is a free data retrieval call binding the contract method 0x04a57c09.
//
// Solidity: function raffleParticipants(uint256 ) view returns(address)
func (_CosmicGame *CosmicGameCaller) RaffleParticipants(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "raffleParticipants", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RaffleParticipants is a free data retrieval call binding the contract method 0x04a57c09.
//
// Solidity: function raffleParticipants(uint256 ) view returns(address)
func (_CosmicGame *CosmicGameSession) RaffleParticipants(arg0 *big.Int) (common.Address, error) {
	return _CosmicGame.Contract.RaffleParticipants(&_CosmicGame.CallOpts, arg0)
}

// RaffleParticipants is a free data retrieval call binding the contract method 0x04a57c09.
//
// Solidity: function raffleParticipants(uint256 ) view returns(address)
func (_CosmicGame *CosmicGameCallerSession) RaffleParticipants(arg0 *big.Int) (common.Address, error) {
	return _CosmicGame.Contract.RaffleParticipants(&_CosmicGame.CallOpts, arg0)
}

// RafflePercentage is a free data retrieval call binding the contract method 0x9250c33c.
//
// Solidity: function rafflePercentage() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) RafflePercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "rafflePercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RafflePercentage is a free data retrieval call binding the contract method 0x9250c33c.
//
// Solidity: function rafflePercentage() view returns(uint256)
func (_CosmicGame *CosmicGameSession) RafflePercentage() (*big.Int, error) {
	return _CosmicGame.Contract.RafflePercentage(&_CosmicGame.CallOpts)
}

// RafflePercentage is a free data retrieval call binding the contract method 0x9250c33c.
//
// Solidity: function rafflePercentage() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) RafflePercentage() (*big.Int, error) {
	return _CosmicGame.Contract.RafflePercentage(&_CosmicGame.CallOpts)
}

// RaffleWallet is a free data retrieval call binding the contract method 0x32d382cd.
//
// Solidity: function raffleWallet() view returns(address)
func (_CosmicGame *CosmicGameCaller) RaffleWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "raffleWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RaffleWallet is a free data retrieval call binding the contract method 0x32d382cd.
//
// Solidity: function raffleWallet() view returns(address)
func (_CosmicGame *CosmicGameSession) RaffleWallet() (common.Address, error) {
	return _CosmicGame.Contract.RaffleWallet(&_CosmicGame.CallOpts)
}

// RaffleWallet is a free data retrieval call binding the contract method 0x32d382cd.
//
// Solidity: function raffleWallet() view returns(address)
func (_CosmicGame *CosmicGameCallerSession) RaffleWallet() (common.Address, error) {
	return _CosmicGame.Contract.RaffleWallet(&_CosmicGame.CallOpts)
}

// RandomWalk is a free data retrieval call binding the contract method 0x5111a2d6.
//
// Solidity: function randomWalk() view returns(address)
func (_CosmicGame *CosmicGameCaller) RandomWalk(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "randomWalk")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RandomWalk is a free data retrieval call binding the contract method 0x5111a2d6.
//
// Solidity: function randomWalk() view returns(address)
func (_CosmicGame *CosmicGameSession) RandomWalk() (common.Address, error) {
	return _CosmicGame.Contract.RandomWalk(&_CosmicGame.CallOpts)
}

// RandomWalk is a free data retrieval call binding the contract method 0x5111a2d6.
//
// Solidity: function randomWalk() view returns(address)
func (_CosmicGame *CosmicGameCallerSession) RandomWalk() (common.Address, error) {
	return _CosmicGame.Contract.RandomWalk(&_CosmicGame.CallOpts)
}

// RoundNum is a free data retrieval call binding the contract method 0x119b22b3.
//
// Solidity: function roundNum() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) RoundNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "roundNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundNum is a free data retrieval call binding the contract method 0x119b22b3.
//
// Solidity: function roundNum() view returns(uint256)
func (_CosmicGame *CosmicGameSession) RoundNum() (*big.Int, error) {
	return _CosmicGame.Contract.RoundNum(&_CosmicGame.CallOpts)
}

// RoundNum is a free data retrieval call binding the contract method 0x119b22b3.
//
// Solidity: function roundNum() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) RoundNum() (*big.Int, error) {
	return _CosmicGame.Contract.RoundNum(&_CosmicGame.CallOpts)
}

// TimeIncrease is a free data retrieval call binding the contract method 0xd94d0316.
//
// Solidity: function timeIncrease() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) TimeIncrease(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "timeIncrease")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeIncrease is a free data retrieval call binding the contract method 0xd94d0316.
//
// Solidity: function timeIncrease() view returns(uint256)
func (_CosmicGame *CosmicGameSession) TimeIncrease() (*big.Int, error) {
	return _CosmicGame.Contract.TimeIncrease(&_CosmicGame.CallOpts)
}

// TimeIncrease is a free data retrieval call binding the contract method 0xd94d0316.
//
// Solidity: function timeIncrease() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) TimeIncrease() (*big.Int, error) {
	return _CosmicGame.Contract.TimeIncrease(&_CosmicGame.CallOpts)
}

// TimeUntilActivation is a free data retrieval call binding the contract method 0xf7178822.
//
// Solidity: function timeUntilActivation() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) TimeUntilActivation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "timeUntilActivation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeUntilActivation is a free data retrieval call binding the contract method 0xf7178822.
//
// Solidity: function timeUntilActivation() view returns(uint256)
func (_CosmicGame *CosmicGameSession) TimeUntilActivation() (*big.Int, error) {
	return _CosmicGame.Contract.TimeUntilActivation(&_CosmicGame.CallOpts)
}

// TimeUntilActivation is a free data retrieval call binding the contract method 0xf7178822.
//
// Solidity: function timeUntilActivation() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) TimeUntilActivation() (*big.Int, error) {
	return _CosmicGame.Contract.TimeUntilActivation(&_CosmicGame.CallOpts)
}

// TimeUntilPrize is a free data retrieval call binding the contract method 0x8b1329e0.
//
// Solidity: function timeUntilPrize() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) TimeUntilPrize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "timeUntilPrize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeUntilPrize is a free data retrieval call binding the contract method 0x8b1329e0.
//
// Solidity: function timeUntilPrize() view returns(uint256)
func (_CosmicGame *CosmicGameSession) TimeUntilPrize() (*big.Int, error) {
	return _CosmicGame.Contract.TimeUntilPrize(&_CosmicGame.CallOpts)
}

// TimeUntilPrize is a free data retrieval call binding the contract method 0x8b1329e0.
//
// Solidity: function timeUntilPrize() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) TimeUntilPrize() (*big.Int, error) {
	return _CosmicGame.Contract.TimeUntilPrize(&_CosmicGame.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CosmicGame *CosmicGameCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CosmicGame *CosmicGameSession) Token() (common.Address, error) {
	return _CosmicGame.Contract.Token(&_CosmicGame.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CosmicGame *CosmicGameCallerSession) Token() (common.Address, error) {
	return _CosmicGame.Contract.Token(&_CosmicGame.CallOpts)
}

// UsedRandomWalkNFTs is a free data retrieval call binding the contract method 0x86e378c9.
//
// Solidity: function usedRandomWalkNFTs(uint256 ) view returns(bool)
func (_CosmicGame *CosmicGameCaller) UsedRandomWalkNFTs(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "usedRandomWalkNFTs", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedRandomWalkNFTs is a free data retrieval call binding the contract method 0x86e378c9.
//
// Solidity: function usedRandomWalkNFTs(uint256 ) view returns(bool)
func (_CosmicGame *CosmicGameSession) UsedRandomWalkNFTs(arg0 *big.Int) (bool, error) {
	return _CosmicGame.Contract.UsedRandomWalkNFTs(&_CosmicGame.CallOpts, arg0)
}

// UsedRandomWalkNFTs is a free data retrieval call binding the contract method 0x86e378c9.
//
// Solidity: function usedRandomWalkNFTs(uint256 ) view returns(bool)
func (_CosmicGame *CosmicGameCallerSession) UsedRandomWalkNFTs(arg0 *big.Int) (bool, error) {
	return _CosmicGame.Contract.UsedRandomWalkNFTs(&_CosmicGame.CallOpts, arg0)
}

// Winners is a free data retrieval call binding the contract method 0xa2fb1175.
//
// Solidity: function winners(uint256 ) view returns(address)
func (_CosmicGame *CosmicGameCaller) Winners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "winners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Winners is a free data retrieval call binding the contract method 0xa2fb1175.
//
// Solidity: function winners(uint256 ) view returns(address)
func (_CosmicGame *CosmicGameSession) Winners(arg0 *big.Int) (common.Address, error) {
	return _CosmicGame.Contract.Winners(&_CosmicGame.CallOpts, arg0)
}

// Winners is a free data retrieval call binding the contract method 0xa2fb1175.
//
// Solidity: function winners(uint256 ) view returns(address)
func (_CosmicGame *CosmicGameCallerSession) Winners(arg0 *big.Int) (common.Address, error) {
	return _CosmicGame.Contract.Winners(&_CosmicGame.CallOpts, arg0)
}

// Bid is a paid mutator transaction binding the contract method 0x7aef951c.
//
// Solidity: function bid(string message) payable returns()
func (_CosmicGame *CosmicGameTransactor) Bid(opts *bind.TransactOpts, message string) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "bid", message)
}

// Bid is a paid mutator transaction binding the contract method 0x7aef951c.
//
// Solidity: function bid(string message) payable returns()
func (_CosmicGame *CosmicGameSession) Bid(message string) (*types.Transaction, error) {
	return _CosmicGame.Contract.Bid(&_CosmicGame.TransactOpts, message)
}

// Bid is a paid mutator transaction binding the contract method 0x7aef951c.
//
// Solidity: function bid(string message) payable returns()
func (_CosmicGame *CosmicGameTransactorSession) Bid(message string) (*types.Transaction, error) {
	return _CosmicGame.Contract.Bid(&_CosmicGame.TransactOpts, message)
}

// BidAndDonateNFT is a paid mutator transaction binding the contract method 0xbb0bc58e.
//
// Solidity: function bidAndDonateNFT(string message, address nftAddress, uint256 tokenId) payable returns()
func (_CosmicGame *CosmicGameTransactor) BidAndDonateNFT(opts *bind.TransactOpts, message string, nftAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "bidAndDonateNFT", message, nftAddress, tokenId)
}

// BidAndDonateNFT is a paid mutator transaction binding the contract method 0xbb0bc58e.
//
// Solidity: function bidAndDonateNFT(string message, address nftAddress, uint256 tokenId) payable returns()
func (_CosmicGame *CosmicGameSession) BidAndDonateNFT(message string, nftAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.BidAndDonateNFT(&_CosmicGame.TransactOpts, message, nftAddress, tokenId)
}

// BidAndDonateNFT is a paid mutator transaction binding the contract method 0xbb0bc58e.
//
// Solidity: function bidAndDonateNFT(string message, address nftAddress, uint256 tokenId) payable returns()
func (_CosmicGame *CosmicGameTransactorSession) BidAndDonateNFT(message string, nftAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.BidAndDonateNFT(&_CosmicGame.TransactOpts, message, nftAddress, tokenId)
}

// BidWithRWLK is a paid mutator transaction binding the contract method 0xae6247f2.
//
// Solidity: function bidWithRWLK(uint256 randomWalkNFTId, string message) returns()
func (_CosmicGame *CosmicGameTransactor) BidWithRWLK(opts *bind.TransactOpts, randomWalkNFTId *big.Int, message string) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "bidWithRWLK", randomWalkNFTId, message)
}

// BidWithRWLK is a paid mutator transaction binding the contract method 0xae6247f2.
//
// Solidity: function bidWithRWLK(uint256 randomWalkNFTId, string message) returns()
func (_CosmicGame *CosmicGameSession) BidWithRWLK(randomWalkNFTId *big.Int, message string) (*types.Transaction, error) {
	return _CosmicGame.Contract.BidWithRWLK(&_CosmicGame.TransactOpts, randomWalkNFTId, message)
}

// BidWithRWLK is a paid mutator transaction binding the contract method 0xae6247f2.
//
// Solidity: function bidWithRWLK(uint256 randomWalkNFTId, string message) returns()
func (_CosmicGame *CosmicGameTransactorSession) BidWithRWLK(randomWalkNFTId *big.Int, message string) (*types.Transaction, error) {
	return _CosmicGame.Contract.BidWithRWLK(&_CosmicGame.TransactOpts, randomWalkNFTId, message)
}

// BidWithRWLKAndDonateNFT is a paid mutator transaction binding the contract method 0xa325b7d2.
//
// Solidity: function bidWithRWLKAndDonateNFT(uint256 randomWalkNFTId, string message, address nftAddress, uint256 tokenId) payable returns()
func (_CosmicGame *CosmicGameTransactor) BidWithRWLKAndDonateNFT(opts *bind.TransactOpts, randomWalkNFTId *big.Int, message string, nftAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "bidWithRWLKAndDonateNFT", randomWalkNFTId, message, nftAddress, tokenId)
}

// BidWithRWLKAndDonateNFT is a paid mutator transaction binding the contract method 0xa325b7d2.
//
// Solidity: function bidWithRWLKAndDonateNFT(uint256 randomWalkNFTId, string message, address nftAddress, uint256 tokenId) payable returns()
func (_CosmicGame *CosmicGameSession) BidWithRWLKAndDonateNFT(randomWalkNFTId *big.Int, message string, nftAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.BidWithRWLKAndDonateNFT(&_CosmicGame.TransactOpts, randomWalkNFTId, message, nftAddress, tokenId)
}

// BidWithRWLKAndDonateNFT is a paid mutator transaction binding the contract method 0xa325b7d2.
//
// Solidity: function bidWithRWLKAndDonateNFT(uint256 randomWalkNFTId, string message, address nftAddress, uint256 tokenId) payable returns()
func (_CosmicGame *CosmicGameTransactorSession) BidWithRWLKAndDonateNFT(randomWalkNFTId *big.Int, message string, nftAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.BidWithRWLKAndDonateNFT(&_CosmicGame.TransactOpts, randomWalkNFTId, message, nftAddress, tokenId)
}

// ClaimDonatedNFT is a paid mutator transaction binding the contract method 0xd59d7478.
//
// Solidity: function claimDonatedNFT(uint256 num) returns()
func (_CosmicGame *CosmicGameTransactor) ClaimDonatedNFT(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "claimDonatedNFT", num)
}

// ClaimDonatedNFT is a paid mutator transaction binding the contract method 0xd59d7478.
//
// Solidity: function claimDonatedNFT(uint256 num) returns()
func (_CosmicGame *CosmicGameSession) ClaimDonatedNFT(num *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.ClaimDonatedNFT(&_CosmicGame.TransactOpts, num)
}

// ClaimDonatedNFT is a paid mutator transaction binding the contract method 0xd59d7478.
//
// Solidity: function claimDonatedNFT(uint256 num) returns()
func (_CosmicGame *CosmicGameTransactorSession) ClaimDonatedNFT(num *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.ClaimDonatedNFT(&_CosmicGame.TransactOpts, num)
}

// ClaimPrize is a paid mutator transaction binding the contract method 0x70740ac9.
//
// Solidity: function claimPrize() returns()
func (_CosmicGame *CosmicGameTransactor) ClaimPrize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "claimPrize")
}

// ClaimPrize is a paid mutator transaction binding the contract method 0x70740ac9.
//
// Solidity: function claimPrize() returns()
func (_CosmicGame *CosmicGameSession) ClaimPrize() (*types.Transaction, error) {
	return _CosmicGame.Contract.ClaimPrize(&_CosmicGame.TransactOpts)
}

// ClaimPrize is a paid mutator transaction binding the contract method 0x70740ac9.
//
// Solidity: function claimPrize() returns()
func (_CosmicGame *CosmicGameTransactorSession) ClaimPrize() (*types.Transaction, error) {
	return _CosmicGame.Contract.ClaimPrize(&_CosmicGame.TransactOpts)
}

// ClaimRaffleNFT is a paid mutator transaction binding the contract method 0xefeacc56.
//
// Solidity: function claimRaffleNFT() returns()
func (_CosmicGame *CosmicGameTransactor) ClaimRaffleNFT(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "claimRaffleNFT")
}

// ClaimRaffleNFT is a paid mutator transaction binding the contract method 0xefeacc56.
//
// Solidity: function claimRaffleNFT() returns()
func (_CosmicGame *CosmicGameSession) ClaimRaffleNFT() (*types.Transaction, error) {
	return _CosmicGame.Contract.ClaimRaffleNFT(&_CosmicGame.TransactOpts)
}

// ClaimRaffleNFT is a paid mutator transaction binding the contract method 0xefeacc56.
//
// Solidity: function claimRaffleNFT() returns()
func (_CosmicGame *CosmicGameTransactorSession) ClaimRaffleNFT() (*types.Transaction, error) {
	return _CosmicGame.Contract.ClaimRaffleNFT(&_CosmicGame.TransactOpts)
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_CosmicGame *CosmicGameTransactor) Donate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "donate")
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_CosmicGame *CosmicGameSession) Donate() (*types.Transaction, error) {
	return _CosmicGame.Contract.Donate(&_CosmicGame.TransactOpts)
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_CosmicGame *CosmicGameTransactorSession) Donate() (*types.Transaction, error) {
	return _CosmicGame.Contract.Donate(&_CosmicGame.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CosmicGame *CosmicGameTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CosmicGame *CosmicGameSession) RenounceOwnership() (*types.Transaction, error) {
	return _CosmicGame.Contract.RenounceOwnership(&_CosmicGame.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CosmicGame *CosmicGameTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CosmicGame.Contract.RenounceOwnership(&_CosmicGame.TransactOpts)
}

// SetActivationTime is a paid mutator transaction binding the contract method 0x7c5486a2.
//
// Solidity: function setActivationTime(uint256 newActivationTime) returns()
func (_CosmicGame *CosmicGameTransactor) SetActivationTime(opts *bind.TransactOpts, newActivationTime *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setActivationTime", newActivationTime)
}

// SetActivationTime is a paid mutator transaction binding the contract method 0x7c5486a2.
//
// Solidity: function setActivationTime(uint256 newActivationTime) returns()
func (_CosmicGame *CosmicGameSession) SetActivationTime(newActivationTime *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetActivationTime(&_CosmicGame.TransactOpts, newActivationTime)
}

// SetActivationTime is a paid mutator transaction binding the contract method 0x7c5486a2.
//
// Solidity: function setActivationTime(uint256 newActivationTime) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetActivationTime(newActivationTime *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetActivationTime(&_CosmicGame.TransactOpts, newActivationTime)
}

// SetCharity is a paid mutator transaction binding the contract method 0xfb6f71a3.
//
// Solidity: function setCharity(address addr) returns()
func (_CosmicGame *CosmicGameTransactor) SetCharity(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setCharity", addr)
}

// SetCharity is a paid mutator transaction binding the contract method 0xfb6f71a3.
//
// Solidity: function setCharity(address addr) returns()
func (_CosmicGame *CosmicGameSession) SetCharity(addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetCharity(&_CosmicGame.TransactOpts, addr)
}

// SetCharity is a paid mutator transaction binding the contract method 0xfb6f71a3.
//
// Solidity: function setCharity(address addr) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetCharity(addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetCharity(&_CosmicGame.TransactOpts, addr)
}

// SetCharityPercentage is a paid mutator transaction binding the contract method 0x5e6e47aa.
//
// Solidity: function setCharityPercentage(uint256 newCharityPercentage) returns()
func (_CosmicGame *CosmicGameTransactor) SetCharityPercentage(opts *bind.TransactOpts, newCharityPercentage *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setCharityPercentage", newCharityPercentage)
}

// SetCharityPercentage is a paid mutator transaction binding the contract method 0x5e6e47aa.
//
// Solidity: function setCharityPercentage(uint256 newCharityPercentage) returns()
func (_CosmicGame *CosmicGameSession) SetCharityPercentage(newCharityPercentage *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetCharityPercentage(&_CosmicGame.TransactOpts, newCharityPercentage)
}

// SetCharityPercentage is a paid mutator transaction binding the contract method 0x5e6e47aa.
//
// Solidity: function setCharityPercentage(uint256 newCharityPercentage) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetCharityPercentage(newCharityPercentage *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetCharityPercentage(&_CosmicGame.TransactOpts, newCharityPercentage)
}

// SetInitialSecondsUntilPrize is a paid mutator transaction binding the contract method 0x51964588.
//
// Solidity: function setInitialSecondsUntilPrize(uint256 newInitialSecondsUntilPrize) returns()
func (_CosmicGame *CosmicGameTransactor) SetInitialSecondsUntilPrize(opts *bind.TransactOpts, newInitialSecondsUntilPrize *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setInitialSecondsUntilPrize", newInitialSecondsUntilPrize)
}

// SetInitialSecondsUntilPrize is a paid mutator transaction binding the contract method 0x51964588.
//
// Solidity: function setInitialSecondsUntilPrize(uint256 newInitialSecondsUntilPrize) returns()
func (_CosmicGame *CosmicGameSession) SetInitialSecondsUntilPrize(newInitialSecondsUntilPrize *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetInitialSecondsUntilPrize(&_CosmicGame.TransactOpts, newInitialSecondsUntilPrize)
}

// SetInitialSecondsUntilPrize is a paid mutator transaction binding the contract method 0x51964588.
//
// Solidity: function setInitialSecondsUntilPrize(uint256 newInitialSecondsUntilPrize) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetInitialSecondsUntilPrize(newInitialSecondsUntilPrize *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetInitialSecondsUntilPrize(&_CosmicGame.TransactOpts, newInitialSecondsUntilPrize)
}

// SetNanoSecondsExtra is a paid mutator transaction binding the contract method 0x05ba9b67.
//
// Solidity: function setNanoSecondsExtra(uint256 newNanoSecondsExtra) returns()
func (_CosmicGame *CosmicGameTransactor) SetNanoSecondsExtra(opts *bind.TransactOpts, newNanoSecondsExtra *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setNanoSecondsExtra", newNanoSecondsExtra)
}

// SetNanoSecondsExtra is a paid mutator transaction binding the contract method 0x05ba9b67.
//
// Solidity: function setNanoSecondsExtra(uint256 newNanoSecondsExtra) returns()
func (_CosmicGame *CosmicGameSession) SetNanoSecondsExtra(newNanoSecondsExtra *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetNanoSecondsExtra(&_CosmicGame.TransactOpts, newNanoSecondsExtra)
}

// SetNanoSecondsExtra is a paid mutator transaction binding the contract method 0x05ba9b67.
//
// Solidity: function setNanoSecondsExtra(uint256 newNanoSecondsExtra) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetNanoSecondsExtra(newNanoSecondsExtra *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetNanoSecondsExtra(&_CosmicGame.TransactOpts, newNanoSecondsExtra)
}

// SetNftContract is a paid mutator transaction binding the contract method 0x52f5ad77.
//
// Solidity: function setNftContract(address addr) returns()
func (_CosmicGame *CosmicGameTransactor) SetNftContract(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setNftContract", addr)
}

// SetNftContract is a paid mutator transaction binding the contract method 0x52f5ad77.
//
// Solidity: function setNftContract(address addr) returns()
func (_CosmicGame *CosmicGameSession) SetNftContract(addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetNftContract(&_CosmicGame.TransactOpts, addr)
}

// SetNftContract is a paid mutator transaction binding the contract method 0x52f5ad77.
//
// Solidity: function setNftContract(address addr) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetNftContract(addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetNftContract(&_CosmicGame.TransactOpts, addr)
}

// SetNumHolderNFTWinnersPerRound is a paid mutator transaction binding the contract method 0xc01c5de2.
//
// Solidity: function setNumHolderNFTWinnersPerRound(uint256 newNumHolderNFTWinnersPerRound) returns()
func (_CosmicGame *CosmicGameTransactor) SetNumHolderNFTWinnersPerRound(opts *bind.TransactOpts, newNumHolderNFTWinnersPerRound *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setNumHolderNFTWinnersPerRound", newNumHolderNFTWinnersPerRound)
}

// SetNumHolderNFTWinnersPerRound is a paid mutator transaction binding the contract method 0xc01c5de2.
//
// Solidity: function setNumHolderNFTWinnersPerRound(uint256 newNumHolderNFTWinnersPerRound) returns()
func (_CosmicGame *CosmicGameSession) SetNumHolderNFTWinnersPerRound(newNumHolderNFTWinnersPerRound *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetNumHolderNFTWinnersPerRound(&_CosmicGame.TransactOpts, newNumHolderNFTWinnersPerRound)
}

// SetNumHolderNFTWinnersPerRound is a paid mutator transaction binding the contract method 0xc01c5de2.
//
// Solidity: function setNumHolderNFTWinnersPerRound(uint256 newNumHolderNFTWinnersPerRound) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetNumHolderNFTWinnersPerRound(newNumHolderNFTWinnersPerRound *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetNumHolderNFTWinnersPerRound(&_CosmicGame.TransactOpts, newNumHolderNFTWinnersPerRound)
}

// SetNumRaffleNFTWinnersPerRound is a paid mutator transaction binding the contract method 0xe1381d7e.
//
// Solidity: function setNumRaffleNFTWinnersPerRound(uint256 newNumRaffleNFTWinnersPerRound) returns()
func (_CosmicGame *CosmicGameTransactor) SetNumRaffleNFTWinnersPerRound(opts *bind.TransactOpts, newNumRaffleNFTWinnersPerRound *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setNumRaffleNFTWinnersPerRound", newNumRaffleNFTWinnersPerRound)
}

// SetNumRaffleNFTWinnersPerRound is a paid mutator transaction binding the contract method 0xe1381d7e.
//
// Solidity: function setNumRaffleNFTWinnersPerRound(uint256 newNumRaffleNFTWinnersPerRound) returns()
func (_CosmicGame *CosmicGameSession) SetNumRaffleNFTWinnersPerRound(newNumRaffleNFTWinnersPerRound *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetNumRaffleNFTWinnersPerRound(&_CosmicGame.TransactOpts, newNumRaffleNFTWinnersPerRound)
}

// SetNumRaffleNFTWinnersPerRound is a paid mutator transaction binding the contract method 0xe1381d7e.
//
// Solidity: function setNumRaffleNFTWinnersPerRound(uint256 newNumRaffleNFTWinnersPerRound) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetNumRaffleNFTWinnersPerRound(newNumRaffleNFTWinnersPerRound *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetNumRaffleNFTWinnersPerRound(&_CosmicGame.TransactOpts, newNumRaffleNFTWinnersPerRound)
}

// SetNumRaffleWinnersPerRound is a paid mutator transaction binding the contract method 0x647b3e7f.
//
// Solidity: function setNumRaffleWinnersPerRound(uint256 newNumRaffleWinnersPerRound) returns()
func (_CosmicGame *CosmicGameTransactor) SetNumRaffleWinnersPerRound(opts *bind.TransactOpts, newNumRaffleWinnersPerRound *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setNumRaffleWinnersPerRound", newNumRaffleWinnersPerRound)
}

// SetNumRaffleWinnersPerRound is a paid mutator transaction binding the contract method 0x647b3e7f.
//
// Solidity: function setNumRaffleWinnersPerRound(uint256 newNumRaffleWinnersPerRound) returns()
func (_CosmicGame *CosmicGameSession) SetNumRaffleWinnersPerRound(newNumRaffleWinnersPerRound *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetNumRaffleWinnersPerRound(&_CosmicGame.TransactOpts, newNumRaffleWinnersPerRound)
}

// SetNumRaffleWinnersPerRound is a paid mutator transaction binding the contract method 0x647b3e7f.
//
// Solidity: function setNumRaffleWinnersPerRound(uint256 newNumRaffleWinnersPerRound) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetNumRaffleWinnersPerRound(newNumRaffleWinnersPerRound *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetNumRaffleWinnersPerRound(&_CosmicGame.TransactOpts, newNumRaffleWinnersPerRound)
}

// SetPriceIncrease is a paid mutator transaction binding the contract method 0x8b122274.
//
// Solidity: function setPriceIncrease(uint256 newPriceIncrease) returns()
func (_CosmicGame *CosmicGameTransactor) SetPriceIncrease(opts *bind.TransactOpts, newPriceIncrease *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setPriceIncrease", newPriceIncrease)
}

// SetPriceIncrease is a paid mutator transaction binding the contract method 0x8b122274.
//
// Solidity: function setPriceIncrease(uint256 newPriceIncrease) returns()
func (_CosmicGame *CosmicGameSession) SetPriceIncrease(newPriceIncrease *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetPriceIncrease(&_CosmicGame.TransactOpts, newPriceIncrease)
}

// SetPriceIncrease is a paid mutator transaction binding the contract method 0x8b122274.
//
// Solidity: function setPriceIncrease(uint256 newPriceIncrease) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetPriceIncrease(newPriceIncrease *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetPriceIncrease(&_CosmicGame.TransactOpts, newPriceIncrease)
}

// SetRafflePercentage is a paid mutator transaction binding the contract method 0xc7c8378d.
//
// Solidity: function setRafflePercentage(uint256 newRafflePercentage) returns()
func (_CosmicGame *CosmicGameTransactor) SetRafflePercentage(opts *bind.TransactOpts, newRafflePercentage *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setRafflePercentage", newRafflePercentage)
}

// SetRafflePercentage is a paid mutator transaction binding the contract method 0xc7c8378d.
//
// Solidity: function setRafflePercentage(uint256 newRafflePercentage) returns()
func (_CosmicGame *CosmicGameSession) SetRafflePercentage(newRafflePercentage *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetRafflePercentage(&_CosmicGame.TransactOpts, newRafflePercentage)
}

// SetRafflePercentage is a paid mutator transaction binding the contract method 0xc7c8378d.
//
// Solidity: function setRafflePercentage(uint256 newRafflePercentage) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetRafflePercentage(newRafflePercentage *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetRafflePercentage(&_CosmicGame.TransactOpts, newRafflePercentage)
}

// SetRaffleWallet is a paid mutator transaction binding the contract method 0x062fb100.
//
// Solidity: function setRaffleWallet(address addr) returns()
func (_CosmicGame *CosmicGameTransactor) SetRaffleWallet(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setRaffleWallet", addr)
}

// SetRaffleWallet is a paid mutator transaction binding the contract method 0x062fb100.
//
// Solidity: function setRaffleWallet(address addr) returns()
func (_CosmicGame *CosmicGameSession) SetRaffleWallet(addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetRaffleWallet(&_CosmicGame.TransactOpts, addr)
}

// SetRaffleWallet is a paid mutator transaction binding the contract method 0x062fb100.
//
// Solidity: function setRaffleWallet(address addr) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetRaffleWallet(addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetRaffleWallet(&_CosmicGame.TransactOpts, addr)
}

// SetRandomWalk is a paid mutator transaction binding the contract method 0xa6f9cc15.
//
// Solidity: function setRandomWalk(address addr) returns()
func (_CosmicGame *CosmicGameTransactor) SetRandomWalk(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setRandomWalk", addr)
}

// SetRandomWalk is a paid mutator transaction binding the contract method 0xa6f9cc15.
//
// Solidity: function setRandomWalk(address addr) returns()
func (_CosmicGame *CosmicGameSession) SetRandomWalk(addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetRandomWalk(&_CosmicGame.TransactOpts, addr)
}

// SetRandomWalk is a paid mutator transaction binding the contract method 0xa6f9cc15.
//
// Solidity: function setRandomWalk(address addr) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetRandomWalk(addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetRandomWalk(&_CosmicGame.TransactOpts, addr)
}

// SetTimeIncrease is a paid mutator transaction binding the contract method 0x4ac3a395.
//
// Solidity: function setTimeIncrease(uint256 newTimeIncrease) returns()
func (_CosmicGame *CosmicGameTransactor) SetTimeIncrease(opts *bind.TransactOpts, newTimeIncrease *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setTimeIncrease", newTimeIncrease)
}

// SetTimeIncrease is a paid mutator transaction binding the contract method 0x4ac3a395.
//
// Solidity: function setTimeIncrease(uint256 newTimeIncrease) returns()
func (_CosmicGame *CosmicGameSession) SetTimeIncrease(newTimeIncrease *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetTimeIncrease(&_CosmicGame.TransactOpts, newTimeIncrease)
}

// SetTimeIncrease is a paid mutator transaction binding the contract method 0x4ac3a395.
//
// Solidity: function setTimeIncrease(uint256 newTimeIncrease) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetTimeIncrease(newTimeIncrease *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetTimeIncrease(&_CosmicGame.TransactOpts, newTimeIncrease)
}

// SetTokenContract is a paid mutator transaction binding the contract method 0xbbcd5bbe.
//
// Solidity: function setTokenContract(address addr) returns()
func (_CosmicGame *CosmicGameTransactor) SetTokenContract(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setTokenContract", addr)
}

// SetTokenContract is a paid mutator transaction binding the contract method 0xbbcd5bbe.
//
// Solidity: function setTokenContract(address addr) returns()
func (_CosmicGame *CosmicGameSession) SetTokenContract(addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetTokenContract(&_CosmicGame.TransactOpts, addr)
}

// SetTokenContract is a paid mutator transaction binding the contract method 0xbbcd5bbe.
//
// Solidity: function setTokenContract(address addr) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetTokenContract(addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetTokenContract(&_CosmicGame.TransactOpts, addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CosmicGame *CosmicGameTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CosmicGame *CosmicGameSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CosmicGame.Contract.TransferOwnership(&_CosmicGame.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CosmicGame *CosmicGameTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CosmicGame.Contract.TransferOwnership(&_CosmicGame.TransactOpts, newOwner)
}

// UpdateInitialBidAmountFraction is a paid mutator transaction binding the contract method 0xa672f6e1.
//
// Solidity: function updateInitialBidAmountFraction(uint256 newInitialBidAmountFraction) returns()
func (_CosmicGame *CosmicGameTransactor) UpdateInitialBidAmountFraction(opts *bind.TransactOpts, newInitialBidAmountFraction *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "updateInitialBidAmountFraction", newInitialBidAmountFraction)
}

// UpdateInitialBidAmountFraction is a paid mutator transaction binding the contract method 0xa672f6e1.
//
// Solidity: function updateInitialBidAmountFraction(uint256 newInitialBidAmountFraction) returns()
func (_CosmicGame *CosmicGameSession) UpdateInitialBidAmountFraction(newInitialBidAmountFraction *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.UpdateInitialBidAmountFraction(&_CosmicGame.TransactOpts, newInitialBidAmountFraction)
}

// UpdateInitialBidAmountFraction is a paid mutator transaction binding the contract method 0xa672f6e1.
//
// Solidity: function updateInitialBidAmountFraction(uint256 newInitialBidAmountFraction) returns()
func (_CosmicGame *CosmicGameTransactorSession) UpdateInitialBidAmountFraction(newInitialBidAmountFraction *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.UpdateInitialBidAmountFraction(&_CosmicGame.TransactOpts, newInitialBidAmountFraction)
}

// UpdatePrizePercentage is a paid mutator transaction binding the contract method 0x3c83adc4.
//
// Solidity: function updatePrizePercentage(uint256 newPrizePercentage) returns()
func (_CosmicGame *CosmicGameTransactor) UpdatePrizePercentage(opts *bind.TransactOpts, newPrizePercentage *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "updatePrizePercentage", newPrizePercentage)
}

// UpdatePrizePercentage is a paid mutator transaction binding the contract method 0x3c83adc4.
//
// Solidity: function updatePrizePercentage(uint256 newPrizePercentage) returns()
func (_CosmicGame *CosmicGameSession) UpdatePrizePercentage(newPrizePercentage *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.UpdatePrizePercentage(&_CosmicGame.TransactOpts, newPrizePercentage)
}

// UpdatePrizePercentage is a paid mutator transaction binding the contract method 0x3c83adc4.
//
// Solidity: function updatePrizePercentage(uint256 newPrizePercentage) returns()
func (_CosmicGame *CosmicGameTransactorSession) UpdatePrizePercentage(newPrizePercentage *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.UpdatePrizePercentage(&_CosmicGame.TransactOpts, newPrizePercentage)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CosmicGame *CosmicGameTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicGame.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CosmicGame *CosmicGameSession) Receive() (*types.Transaction, error) {
	return _CosmicGame.Contract.Receive(&_CosmicGame.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CosmicGame *CosmicGameTransactorSession) Receive() (*types.Transaction, error) {
	return _CosmicGame.Contract.Receive(&_CosmicGame.TransactOpts)
}

// CosmicGameActivationTimeChangedIterator is returned from FilterActivationTimeChanged and is used to iterate over the raw logs and unpacked data for ActivationTimeChanged events raised by the CosmicGame contract.
type CosmicGameActivationTimeChangedIterator struct {
	Event *CosmicGameActivationTimeChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameActivationTimeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameActivationTimeChanged)
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
		it.Event = new(CosmicGameActivationTimeChanged)
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
func (it *CosmicGameActivationTimeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameActivationTimeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameActivationTimeChanged represents a ActivationTimeChanged event raised by the CosmicGame contract.
type CosmicGameActivationTimeChanged struct {
	NewActivationTime *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterActivationTimeChanged is a free log retrieval operation binding the contract event 0x584ff9a7b3a86db5397052f2e440da0ed60a95f646a3a884863cd92262e682b6.
//
// Solidity: event ActivationTimeChanged(uint256 newActivationTime)
func (_CosmicGame *CosmicGameFilterer) FilterActivationTimeChanged(opts *bind.FilterOpts) (*CosmicGameActivationTimeChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "ActivationTimeChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameActivationTimeChangedIterator{contract: _CosmicGame.contract, event: "ActivationTimeChanged", logs: logs, sub: sub}, nil
}

// WatchActivationTimeChanged is a free log subscription operation binding the contract event 0x584ff9a7b3a86db5397052f2e440da0ed60a95f646a3a884863cd92262e682b6.
//
// Solidity: event ActivationTimeChanged(uint256 newActivationTime)
func (_CosmicGame *CosmicGameFilterer) WatchActivationTimeChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameActivationTimeChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "ActivationTimeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameActivationTimeChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "ActivationTimeChanged", log); err != nil {
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

// ParseActivationTimeChanged is a log parse operation binding the contract event 0x584ff9a7b3a86db5397052f2e440da0ed60a95f646a3a884863cd92262e682b6.
//
// Solidity: event ActivationTimeChanged(uint256 newActivationTime)
func (_CosmicGame *CosmicGameFilterer) ParseActivationTimeChanged(log types.Log) (*CosmicGameActivationTimeChanged, error) {
	event := new(CosmicGameActivationTimeChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "ActivationTimeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameBidEventIterator is returned from FilterBidEvent and is used to iterate over the raw logs and unpacked data for BidEvent events raised by the CosmicGame contract.
type CosmicGameBidEventIterator struct {
	Event *CosmicGameBidEvent // Event containing the contract specifics and raw log

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
func (it *CosmicGameBidEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameBidEvent)
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
		it.Event = new(CosmicGameBidEvent)
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
func (it *CosmicGameBidEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameBidEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameBidEvent represents a BidEvent event raised by the CosmicGame contract.
type CosmicGameBidEvent struct {
	LastBidder      common.Address
	Round           *big.Int
	BidPrice        *big.Int
	RandomWalkNFTId *big.Int
	PrizeTime       *big.Int
	Message         string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBidEvent is a free log retrieval operation binding the contract event 0xc7beef8f8d25068377c10e7b7f30ef5622f443eb674e07835df1daf5fe84acbf.
//
// Solidity: event BidEvent(address indexed lastBidder, uint256 indexed round, uint256 bidPrice, int256 randomWalkNFTId, uint256 prizeTime, string message)
func (_CosmicGame *CosmicGameFilterer) FilterBidEvent(opts *bind.FilterOpts, lastBidder []common.Address, round []*big.Int) (*CosmicGameBidEventIterator, error) {

	var lastBidderRule []interface{}
	for _, lastBidderItem := range lastBidder {
		lastBidderRule = append(lastBidderRule, lastBidderItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "BidEvent", lastBidderRule, roundRule)
	if err != nil {
		return nil, err
	}
	return &CosmicGameBidEventIterator{contract: _CosmicGame.contract, event: "BidEvent", logs: logs, sub: sub}, nil
}

// WatchBidEvent is a free log subscription operation binding the contract event 0xc7beef8f8d25068377c10e7b7f30ef5622f443eb674e07835df1daf5fe84acbf.
//
// Solidity: event BidEvent(address indexed lastBidder, uint256 indexed round, uint256 bidPrice, int256 randomWalkNFTId, uint256 prizeTime, string message)
func (_CosmicGame *CosmicGameFilterer) WatchBidEvent(opts *bind.WatchOpts, sink chan<- *CosmicGameBidEvent, lastBidder []common.Address, round []*big.Int) (event.Subscription, error) {

	var lastBidderRule []interface{}
	for _, lastBidderItem := range lastBidder {
		lastBidderRule = append(lastBidderRule, lastBidderItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "BidEvent", lastBidderRule, roundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameBidEvent)
				if err := _CosmicGame.contract.UnpackLog(event, "BidEvent", log); err != nil {
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

// ParseBidEvent is a log parse operation binding the contract event 0xc7beef8f8d25068377c10e7b7f30ef5622f443eb674e07835df1daf5fe84acbf.
//
// Solidity: event BidEvent(address indexed lastBidder, uint256 indexed round, uint256 bidPrice, int256 randomWalkNFTId, uint256 prizeTime, string message)
func (_CosmicGame *CosmicGameFilterer) ParseBidEvent(log types.Log) (*CosmicGameBidEvent, error) {
	event := new(CosmicGameBidEvent)
	if err := _CosmicGame.contract.UnpackLog(event, "BidEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameCharityAddressChangedIterator is returned from FilterCharityAddressChanged and is used to iterate over the raw logs and unpacked data for CharityAddressChanged events raised by the CosmicGame contract.
type CosmicGameCharityAddressChangedIterator struct {
	Event *CosmicGameCharityAddressChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameCharityAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameCharityAddressChanged)
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
		it.Event = new(CosmicGameCharityAddressChanged)
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
func (it *CosmicGameCharityAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameCharityAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameCharityAddressChanged represents a CharityAddressChanged event raised by the CosmicGame contract.
type CosmicGameCharityAddressChanged struct {
	NewCharity common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCharityAddressChanged is a free log retrieval operation binding the contract event 0x1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c.
//
// Solidity: event CharityAddressChanged(address newCharity)
func (_CosmicGame *CosmicGameFilterer) FilterCharityAddressChanged(opts *bind.FilterOpts) (*CosmicGameCharityAddressChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "CharityAddressChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameCharityAddressChangedIterator{contract: _CosmicGame.contract, event: "CharityAddressChanged", logs: logs, sub: sub}, nil
}

// WatchCharityAddressChanged is a free log subscription operation binding the contract event 0x1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c.
//
// Solidity: event CharityAddressChanged(address newCharity)
func (_CosmicGame *CosmicGameFilterer) WatchCharityAddressChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameCharityAddressChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "CharityAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameCharityAddressChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "CharityAddressChanged", log); err != nil {
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

// ParseCharityAddressChanged is a log parse operation binding the contract event 0x1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c.
//
// Solidity: event CharityAddressChanged(address newCharity)
func (_CosmicGame *CosmicGameFilterer) ParseCharityAddressChanged(log types.Log) (*CosmicGameCharityAddressChanged, error) {
	event := new(CosmicGameCharityAddressChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "CharityAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameCharityPercentageChangedIterator is returned from FilterCharityPercentageChanged and is used to iterate over the raw logs and unpacked data for CharityPercentageChanged events raised by the CosmicGame contract.
type CosmicGameCharityPercentageChangedIterator struct {
	Event *CosmicGameCharityPercentageChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameCharityPercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameCharityPercentageChanged)
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
		it.Event = new(CosmicGameCharityPercentageChanged)
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
func (it *CosmicGameCharityPercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameCharityPercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameCharityPercentageChanged represents a CharityPercentageChanged event raised by the CosmicGame contract.
type CosmicGameCharityPercentageChanged struct {
	NewCharityPercentage *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterCharityPercentageChanged is a free log retrieval operation binding the contract event 0x0918b499c15fcb0e087d411d53664cde23577e1aa4a9cbfbdf735ddd1871e7d5.
//
// Solidity: event CharityPercentageChanged(uint256 newCharityPercentage)
func (_CosmicGame *CosmicGameFilterer) FilterCharityPercentageChanged(opts *bind.FilterOpts) (*CosmicGameCharityPercentageChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "CharityPercentageChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameCharityPercentageChangedIterator{contract: _CosmicGame.contract, event: "CharityPercentageChanged", logs: logs, sub: sub}, nil
}

// WatchCharityPercentageChanged is a free log subscription operation binding the contract event 0x0918b499c15fcb0e087d411d53664cde23577e1aa4a9cbfbdf735ddd1871e7d5.
//
// Solidity: event CharityPercentageChanged(uint256 newCharityPercentage)
func (_CosmicGame *CosmicGameFilterer) WatchCharityPercentageChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameCharityPercentageChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "CharityPercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameCharityPercentageChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "CharityPercentageChanged", log); err != nil {
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

// ParseCharityPercentageChanged is a log parse operation binding the contract event 0x0918b499c15fcb0e087d411d53664cde23577e1aa4a9cbfbdf735ddd1871e7d5.
//
// Solidity: event CharityPercentageChanged(uint256 newCharityPercentage)
func (_CosmicGame *CosmicGameFilterer) ParseCharityPercentageChanged(log types.Log) (*CosmicGameCharityPercentageChanged, error) {
	event := new(CosmicGameCharityPercentageChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "CharityPercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameCosmicSignatureAddressChangedIterator is returned from FilterCosmicSignatureAddressChanged and is used to iterate over the raw logs and unpacked data for CosmicSignatureAddressChanged events raised by the CosmicGame contract.
type CosmicGameCosmicSignatureAddressChangedIterator struct {
	Event *CosmicGameCosmicSignatureAddressChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameCosmicSignatureAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameCosmicSignatureAddressChanged)
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
		it.Event = new(CosmicGameCosmicSignatureAddressChanged)
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
func (it *CosmicGameCosmicSignatureAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameCosmicSignatureAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameCosmicSignatureAddressChanged represents a CosmicSignatureAddressChanged event raised by the CosmicGame contract.
type CosmicGameCosmicSignatureAddressChanged struct {
	NewCosmicSignature common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCosmicSignatureAddressChanged is a free log retrieval operation binding the contract event 0x7142a592d5404b5fc4a294c66f70e32b2a7776bd807f722d59268def87c765d1.
//
// Solidity: event CosmicSignatureAddressChanged(address newCosmicSignature)
func (_CosmicGame *CosmicGameFilterer) FilterCosmicSignatureAddressChanged(opts *bind.FilterOpts) (*CosmicGameCosmicSignatureAddressChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "CosmicSignatureAddressChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameCosmicSignatureAddressChangedIterator{contract: _CosmicGame.contract, event: "CosmicSignatureAddressChanged", logs: logs, sub: sub}, nil
}

// WatchCosmicSignatureAddressChanged is a free log subscription operation binding the contract event 0x7142a592d5404b5fc4a294c66f70e32b2a7776bd807f722d59268def87c765d1.
//
// Solidity: event CosmicSignatureAddressChanged(address newCosmicSignature)
func (_CosmicGame *CosmicGameFilterer) WatchCosmicSignatureAddressChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameCosmicSignatureAddressChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "CosmicSignatureAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameCosmicSignatureAddressChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "CosmicSignatureAddressChanged", log); err != nil {
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

// ParseCosmicSignatureAddressChanged is a log parse operation binding the contract event 0x7142a592d5404b5fc4a294c66f70e32b2a7776bd807f722d59268def87c765d1.
//
// Solidity: event CosmicSignatureAddressChanged(address newCosmicSignature)
func (_CosmicGame *CosmicGameFilterer) ParseCosmicSignatureAddressChanged(log types.Log) (*CosmicGameCosmicSignatureAddressChanged, error) {
	event := new(CosmicGameCosmicSignatureAddressChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "CosmicSignatureAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameCosmicTokenAddressChangedIterator is returned from FilterCosmicTokenAddressChanged and is used to iterate over the raw logs and unpacked data for CosmicTokenAddressChanged events raised by the CosmicGame contract.
type CosmicGameCosmicTokenAddressChangedIterator struct {
	Event *CosmicGameCosmicTokenAddressChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameCosmicTokenAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameCosmicTokenAddressChanged)
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
		it.Event = new(CosmicGameCosmicTokenAddressChanged)
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
func (it *CosmicGameCosmicTokenAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameCosmicTokenAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameCosmicTokenAddressChanged represents a CosmicTokenAddressChanged event raised by the CosmicGame contract.
type CosmicGameCosmicTokenAddressChanged struct {
	NewCosmicToken common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCosmicTokenAddressChanged is a free log retrieval operation binding the contract event 0x3ab2dcf539aa3c30314265602fc86dee8e343e3c3905730f870835f36c400deb.
//
// Solidity: event CosmicTokenAddressChanged(address newCosmicToken)
func (_CosmicGame *CosmicGameFilterer) FilterCosmicTokenAddressChanged(opts *bind.FilterOpts) (*CosmicGameCosmicTokenAddressChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "CosmicTokenAddressChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameCosmicTokenAddressChangedIterator{contract: _CosmicGame.contract, event: "CosmicTokenAddressChanged", logs: logs, sub: sub}, nil
}

// WatchCosmicTokenAddressChanged is a free log subscription operation binding the contract event 0x3ab2dcf539aa3c30314265602fc86dee8e343e3c3905730f870835f36c400deb.
//
// Solidity: event CosmicTokenAddressChanged(address newCosmicToken)
func (_CosmicGame *CosmicGameFilterer) WatchCosmicTokenAddressChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameCosmicTokenAddressChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "CosmicTokenAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameCosmicTokenAddressChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "CosmicTokenAddressChanged", log); err != nil {
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

// ParseCosmicTokenAddressChanged is a log parse operation binding the contract event 0x3ab2dcf539aa3c30314265602fc86dee8e343e3c3905730f870835f36c400deb.
//
// Solidity: event CosmicTokenAddressChanged(address newCosmicToken)
func (_CosmicGame *CosmicGameFilterer) ParseCosmicTokenAddressChanged(log types.Log) (*CosmicGameCosmicTokenAddressChanged, error) {
	event := new(CosmicGameCosmicTokenAddressChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "CosmicTokenAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameDonatedNFTClaimedEventIterator is returned from FilterDonatedNFTClaimedEvent and is used to iterate over the raw logs and unpacked data for DonatedNFTClaimedEvent events raised by the CosmicGame contract.
type CosmicGameDonatedNFTClaimedEventIterator struct {
	Event *CosmicGameDonatedNFTClaimedEvent // Event containing the contract specifics and raw log

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
func (it *CosmicGameDonatedNFTClaimedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameDonatedNFTClaimedEvent)
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
		it.Event = new(CosmicGameDonatedNFTClaimedEvent)
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
func (it *CosmicGameDonatedNFTClaimedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameDonatedNFTClaimedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameDonatedNFTClaimedEvent represents a DonatedNFTClaimedEvent event raised by the CosmicGame contract.
type CosmicGameDonatedNFTClaimedEvent struct {
	Round                 *big.Int
	Index                 *big.Int
	Winner                common.Address
	NftAddressdonatedNFTs common.Address
	TokenId               *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterDonatedNFTClaimedEvent is a free log retrieval operation binding the contract event 0x0d0e4b90a96d931964d5b0421a6b5b48dd73abf538cdb9ae015069d49e9a4679.
//
// Solidity: event DonatedNFTClaimedEvent(uint256 indexed round, uint256 index, address winner, address nftAddressdonatedNFTs, uint256 tokenId)
func (_CosmicGame *CosmicGameFilterer) FilterDonatedNFTClaimedEvent(opts *bind.FilterOpts, round []*big.Int) (*CosmicGameDonatedNFTClaimedEventIterator, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "DonatedNFTClaimedEvent", roundRule)
	if err != nil {
		return nil, err
	}
	return &CosmicGameDonatedNFTClaimedEventIterator{contract: _CosmicGame.contract, event: "DonatedNFTClaimedEvent", logs: logs, sub: sub}, nil
}

// WatchDonatedNFTClaimedEvent is a free log subscription operation binding the contract event 0x0d0e4b90a96d931964d5b0421a6b5b48dd73abf538cdb9ae015069d49e9a4679.
//
// Solidity: event DonatedNFTClaimedEvent(uint256 indexed round, uint256 index, address winner, address nftAddressdonatedNFTs, uint256 tokenId)
func (_CosmicGame *CosmicGameFilterer) WatchDonatedNFTClaimedEvent(opts *bind.WatchOpts, sink chan<- *CosmicGameDonatedNFTClaimedEvent, round []*big.Int) (event.Subscription, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "DonatedNFTClaimedEvent", roundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameDonatedNFTClaimedEvent)
				if err := _CosmicGame.contract.UnpackLog(event, "DonatedNFTClaimedEvent", log); err != nil {
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

// ParseDonatedNFTClaimedEvent is a log parse operation binding the contract event 0x0d0e4b90a96d931964d5b0421a6b5b48dd73abf538cdb9ae015069d49e9a4679.
//
// Solidity: event DonatedNFTClaimedEvent(uint256 indexed round, uint256 index, address winner, address nftAddressdonatedNFTs, uint256 tokenId)
func (_CosmicGame *CosmicGameFilterer) ParseDonatedNFTClaimedEvent(log types.Log) (*CosmicGameDonatedNFTClaimedEvent, error) {
	event := new(CosmicGameDonatedNFTClaimedEvent)
	if err := _CosmicGame.contract.UnpackLog(event, "DonatedNFTClaimedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameDonationEventIterator is returned from FilterDonationEvent and is used to iterate over the raw logs and unpacked data for DonationEvent events raised by the CosmicGame contract.
type CosmicGameDonationEventIterator struct {
	Event *CosmicGameDonationEvent // Event containing the contract specifics and raw log

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
func (it *CosmicGameDonationEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameDonationEvent)
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
		it.Event = new(CosmicGameDonationEvent)
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
func (it *CosmicGameDonationEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameDonationEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameDonationEvent represents a DonationEvent event raised by the CosmicGame contract.
type CosmicGameDonationEvent struct {
	Donor  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDonationEvent is a free log retrieval operation binding the contract event 0x8b7fe5be5699654fd637d2250cb0d47e88205730710745e78e9d8bcaf8aad8f1.
//
// Solidity: event DonationEvent(address indexed donor, uint256 amount)
func (_CosmicGame *CosmicGameFilterer) FilterDonationEvent(opts *bind.FilterOpts, donor []common.Address) (*CosmicGameDonationEventIterator, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "DonationEvent", donorRule)
	if err != nil {
		return nil, err
	}
	return &CosmicGameDonationEventIterator{contract: _CosmicGame.contract, event: "DonationEvent", logs: logs, sub: sub}, nil
}

// WatchDonationEvent is a free log subscription operation binding the contract event 0x8b7fe5be5699654fd637d2250cb0d47e88205730710745e78e9d8bcaf8aad8f1.
//
// Solidity: event DonationEvent(address indexed donor, uint256 amount)
func (_CosmicGame *CosmicGameFilterer) WatchDonationEvent(opts *bind.WatchOpts, sink chan<- *CosmicGameDonationEvent, donor []common.Address) (event.Subscription, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "DonationEvent", donorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameDonationEvent)
				if err := _CosmicGame.contract.UnpackLog(event, "DonationEvent", log); err != nil {
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

// ParseDonationEvent is a log parse operation binding the contract event 0x8b7fe5be5699654fd637d2250cb0d47e88205730710745e78e9d8bcaf8aad8f1.
//
// Solidity: event DonationEvent(address indexed donor, uint256 amount)
func (_CosmicGame *CosmicGameFilterer) ParseDonationEvent(log types.Log) (*CosmicGameDonationEvent, error) {
	event := new(CosmicGameDonationEvent)
	if err := _CosmicGame.contract.UnpackLog(event, "DonationEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameInitialBidAmountFractionChangedIterator is returned from FilterInitialBidAmountFractionChanged and is used to iterate over the raw logs and unpacked data for InitialBidAmountFractionChanged events raised by the CosmicGame contract.
type CosmicGameInitialBidAmountFractionChangedIterator struct {
	Event *CosmicGameInitialBidAmountFractionChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameInitialBidAmountFractionChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameInitialBidAmountFractionChanged)
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
		it.Event = new(CosmicGameInitialBidAmountFractionChanged)
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
func (it *CosmicGameInitialBidAmountFractionChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameInitialBidAmountFractionChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameInitialBidAmountFractionChanged represents a InitialBidAmountFractionChanged event raised by the CosmicGame contract.
type CosmicGameInitialBidAmountFractionChanged struct {
	NewInitialBidAmountFraction *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterInitialBidAmountFractionChanged is a free log retrieval operation binding the contract event 0x3b311f029da1b90c3de2e0b3168436c5ed7d8b0ae81b7d4894c12da03835c628.
//
// Solidity: event InitialBidAmountFractionChanged(uint256 newInitialBidAmountFraction)
func (_CosmicGame *CosmicGameFilterer) FilterInitialBidAmountFractionChanged(opts *bind.FilterOpts) (*CosmicGameInitialBidAmountFractionChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "InitialBidAmountFractionChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameInitialBidAmountFractionChangedIterator{contract: _CosmicGame.contract, event: "InitialBidAmountFractionChanged", logs: logs, sub: sub}, nil
}

// WatchInitialBidAmountFractionChanged is a free log subscription operation binding the contract event 0x3b311f029da1b90c3de2e0b3168436c5ed7d8b0ae81b7d4894c12da03835c628.
//
// Solidity: event InitialBidAmountFractionChanged(uint256 newInitialBidAmountFraction)
func (_CosmicGame *CosmicGameFilterer) WatchInitialBidAmountFractionChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameInitialBidAmountFractionChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "InitialBidAmountFractionChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameInitialBidAmountFractionChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "InitialBidAmountFractionChanged", log); err != nil {
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

// ParseInitialBidAmountFractionChanged is a log parse operation binding the contract event 0x3b311f029da1b90c3de2e0b3168436c5ed7d8b0ae81b7d4894c12da03835c628.
//
// Solidity: event InitialBidAmountFractionChanged(uint256 newInitialBidAmountFraction)
func (_CosmicGame *CosmicGameFilterer) ParseInitialBidAmountFractionChanged(log types.Log) (*CosmicGameInitialBidAmountFractionChanged, error) {
	event := new(CosmicGameInitialBidAmountFractionChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "InitialBidAmountFractionChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameInitialSecondsUntilPrizeChangedIterator is returned from FilterInitialSecondsUntilPrizeChanged and is used to iterate over the raw logs and unpacked data for InitialSecondsUntilPrizeChanged events raised by the CosmicGame contract.
type CosmicGameInitialSecondsUntilPrizeChangedIterator struct {
	Event *CosmicGameInitialSecondsUntilPrizeChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameInitialSecondsUntilPrizeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameInitialSecondsUntilPrizeChanged)
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
		it.Event = new(CosmicGameInitialSecondsUntilPrizeChanged)
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
func (it *CosmicGameInitialSecondsUntilPrizeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameInitialSecondsUntilPrizeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameInitialSecondsUntilPrizeChanged represents a InitialSecondsUntilPrizeChanged event raised by the CosmicGame contract.
type CosmicGameInitialSecondsUntilPrizeChanged struct {
	NewInitialSecondsUntilPrize *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterInitialSecondsUntilPrizeChanged is a free log retrieval operation binding the contract event 0x6da281754ba85ee0c5983a8e8f05a92910c2a0c5b80e68c126216d65f162a305.
//
// Solidity: event InitialSecondsUntilPrizeChanged(uint256 newInitialSecondsUntilPrize)
func (_CosmicGame *CosmicGameFilterer) FilterInitialSecondsUntilPrizeChanged(opts *bind.FilterOpts) (*CosmicGameInitialSecondsUntilPrizeChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "InitialSecondsUntilPrizeChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameInitialSecondsUntilPrizeChangedIterator{contract: _CosmicGame.contract, event: "InitialSecondsUntilPrizeChanged", logs: logs, sub: sub}, nil
}

// WatchInitialSecondsUntilPrizeChanged is a free log subscription operation binding the contract event 0x6da281754ba85ee0c5983a8e8f05a92910c2a0c5b80e68c126216d65f162a305.
//
// Solidity: event InitialSecondsUntilPrizeChanged(uint256 newInitialSecondsUntilPrize)
func (_CosmicGame *CosmicGameFilterer) WatchInitialSecondsUntilPrizeChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameInitialSecondsUntilPrizeChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "InitialSecondsUntilPrizeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameInitialSecondsUntilPrizeChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "InitialSecondsUntilPrizeChanged", log); err != nil {
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

// ParseInitialSecondsUntilPrizeChanged is a log parse operation binding the contract event 0x6da281754ba85ee0c5983a8e8f05a92910c2a0c5b80e68c126216d65f162a305.
//
// Solidity: event InitialSecondsUntilPrizeChanged(uint256 newInitialSecondsUntilPrize)
func (_CosmicGame *CosmicGameFilterer) ParseInitialSecondsUntilPrizeChanged(log types.Log) (*CosmicGameInitialSecondsUntilPrizeChanged, error) {
	event := new(CosmicGameInitialSecondsUntilPrizeChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "InitialSecondsUntilPrizeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameNFTDonationEventIterator is returned from FilterNFTDonationEvent and is used to iterate over the raw logs and unpacked data for NFTDonationEvent events raised by the CosmicGame contract.
type CosmicGameNFTDonationEventIterator struct {
	Event *CosmicGameNFTDonationEvent // Event containing the contract specifics and raw log

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
func (it *CosmicGameNFTDonationEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameNFTDonationEvent)
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
		it.Event = new(CosmicGameNFTDonationEvent)
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
func (it *CosmicGameNFTDonationEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameNFTDonationEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameNFTDonationEvent represents a NFTDonationEvent event raised by the CosmicGame contract.
type CosmicGameNFTDonationEvent struct {
	Donor      common.Address
	NftAddress common.Address
	Round      *big.Int
	TokenId    *big.Int
	Index      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNFTDonationEvent is a free log retrieval operation binding the contract event 0xc85be1734ed6a0f4d6adf924d4d41406e2729878c652110a96e2fdec64e118d1.
//
// Solidity: event NFTDonationEvent(address indexed donor, address indexed nftAddress, uint256 indexed round, uint256 tokenId, uint256 index)
func (_CosmicGame *CosmicGameFilterer) FilterNFTDonationEvent(opts *bind.FilterOpts, donor []common.Address, nftAddress []common.Address, round []*big.Int) (*CosmicGameNFTDonationEventIterator, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "NFTDonationEvent", donorRule, nftAddressRule, roundRule)
	if err != nil {
		return nil, err
	}
	return &CosmicGameNFTDonationEventIterator{contract: _CosmicGame.contract, event: "NFTDonationEvent", logs: logs, sub: sub}, nil
}

// WatchNFTDonationEvent is a free log subscription operation binding the contract event 0xc85be1734ed6a0f4d6adf924d4d41406e2729878c652110a96e2fdec64e118d1.
//
// Solidity: event NFTDonationEvent(address indexed donor, address indexed nftAddress, uint256 indexed round, uint256 tokenId, uint256 index)
func (_CosmicGame *CosmicGameFilterer) WatchNFTDonationEvent(opts *bind.WatchOpts, sink chan<- *CosmicGameNFTDonationEvent, donor []common.Address, nftAddress []common.Address, round []*big.Int) (event.Subscription, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "NFTDonationEvent", donorRule, nftAddressRule, roundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameNFTDonationEvent)
				if err := _CosmicGame.contract.UnpackLog(event, "NFTDonationEvent", log); err != nil {
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

// ParseNFTDonationEvent is a log parse operation binding the contract event 0xc85be1734ed6a0f4d6adf924d4d41406e2729878c652110a96e2fdec64e118d1.
//
// Solidity: event NFTDonationEvent(address indexed donor, address indexed nftAddress, uint256 indexed round, uint256 tokenId, uint256 index)
func (_CosmicGame *CosmicGameFilterer) ParseNFTDonationEvent(log types.Log) (*CosmicGameNFTDonationEvent, error) {
	event := new(CosmicGameNFTDonationEvent)
	if err := _CosmicGame.contract.UnpackLog(event, "NFTDonationEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameNanoSecondsExtraChangedIterator is returned from FilterNanoSecondsExtraChanged and is used to iterate over the raw logs and unpacked data for NanoSecondsExtraChanged events raised by the CosmicGame contract.
type CosmicGameNanoSecondsExtraChangedIterator struct {
	Event *CosmicGameNanoSecondsExtraChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameNanoSecondsExtraChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameNanoSecondsExtraChanged)
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
		it.Event = new(CosmicGameNanoSecondsExtraChanged)
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
func (it *CosmicGameNanoSecondsExtraChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameNanoSecondsExtraChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameNanoSecondsExtraChanged represents a NanoSecondsExtraChanged event raised by the CosmicGame contract.
type CosmicGameNanoSecondsExtraChanged struct {
	NewNanoSecondsExtra *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNanoSecondsExtraChanged is a free log retrieval operation binding the contract event 0x678d086a09e1be49401b1e3a6e14db1878e8d8b88e71d0cfe24a32726d0e38b9.
//
// Solidity: event NanoSecondsExtraChanged(uint256 newNanoSecondsExtra)
func (_CosmicGame *CosmicGameFilterer) FilterNanoSecondsExtraChanged(opts *bind.FilterOpts) (*CosmicGameNanoSecondsExtraChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "NanoSecondsExtraChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameNanoSecondsExtraChangedIterator{contract: _CosmicGame.contract, event: "NanoSecondsExtraChanged", logs: logs, sub: sub}, nil
}

// WatchNanoSecondsExtraChanged is a free log subscription operation binding the contract event 0x678d086a09e1be49401b1e3a6e14db1878e8d8b88e71d0cfe24a32726d0e38b9.
//
// Solidity: event NanoSecondsExtraChanged(uint256 newNanoSecondsExtra)
func (_CosmicGame *CosmicGameFilterer) WatchNanoSecondsExtraChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameNanoSecondsExtraChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "NanoSecondsExtraChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameNanoSecondsExtraChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "NanoSecondsExtraChanged", log); err != nil {
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

// ParseNanoSecondsExtraChanged is a log parse operation binding the contract event 0x678d086a09e1be49401b1e3a6e14db1878e8d8b88e71d0cfe24a32726d0e38b9.
//
// Solidity: event NanoSecondsExtraChanged(uint256 newNanoSecondsExtra)
func (_CosmicGame *CosmicGameFilterer) ParseNanoSecondsExtraChanged(log types.Log) (*CosmicGameNanoSecondsExtraChanged, error) {
	event := new(CosmicGameNanoSecondsExtraChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "NanoSecondsExtraChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameNumHolderNFTWinnersPerRoundChangedIterator is returned from FilterNumHolderNFTWinnersPerRoundChanged and is used to iterate over the raw logs and unpacked data for NumHolderNFTWinnersPerRoundChanged events raised by the CosmicGame contract.
type CosmicGameNumHolderNFTWinnersPerRoundChangedIterator struct {
	Event *CosmicGameNumHolderNFTWinnersPerRoundChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameNumHolderNFTWinnersPerRoundChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameNumHolderNFTWinnersPerRoundChanged)
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
		it.Event = new(CosmicGameNumHolderNFTWinnersPerRoundChanged)
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
func (it *CosmicGameNumHolderNFTWinnersPerRoundChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameNumHolderNFTWinnersPerRoundChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameNumHolderNFTWinnersPerRoundChanged represents a NumHolderNFTWinnersPerRoundChanged event raised by the CosmicGame contract.
type CosmicGameNumHolderNFTWinnersPerRoundChanged struct {
	NewNumHolderNFTWinnersPerRound *big.Int
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterNumHolderNFTWinnersPerRoundChanged is a free log retrieval operation binding the contract event 0x0cc7e37c68566f67d0fe13bf38246d7447cf99a0c481c2ef9963969bb4f5ebc8.
//
// Solidity: event NumHolderNFTWinnersPerRoundChanged(uint256 newNumHolderNFTWinnersPerRound)
func (_CosmicGame *CosmicGameFilterer) FilterNumHolderNFTWinnersPerRoundChanged(opts *bind.FilterOpts) (*CosmicGameNumHolderNFTWinnersPerRoundChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "NumHolderNFTWinnersPerRoundChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameNumHolderNFTWinnersPerRoundChangedIterator{contract: _CosmicGame.contract, event: "NumHolderNFTWinnersPerRoundChanged", logs: logs, sub: sub}, nil
}

// WatchNumHolderNFTWinnersPerRoundChanged is a free log subscription operation binding the contract event 0x0cc7e37c68566f67d0fe13bf38246d7447cf99a0c481c2ef9963969bb4f5ebc8.
//
// Solidity: event NumHolderNFTWinnersPerRoundChanged(uint256 newNumHolderNFTWinnersPerRound)
func (_CosmicGame *CosmicGameFilterer) WatchNumHolderNFTWinnersPerRoundChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameNumHolderNFTWinnersPerRoundChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "NumHolderNFTWinnersPerRoundChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameNumHolderNFTWinnersPerRoundChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "NumHolderNFTWinnersPerRoundChanged", log); err != nil {
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

// ParseNumHolderNFTWinnersPerRoundChanged is a log parse operation binding the contract event 0x0cc7e37c68566f67d0fe13bf38246d7447cf99a0c481c2ef9963969bb4f5ebc8.
//
// Solidity: event NumHolderNFTWinnersPerRoundChanged(uint256 newNumHolderNFTWinnersPerRound)
func (_CosmicGame *CosmicGameFilterer) ParseNumHolderNFTWinnersPerRoundChanged(log types.Log) (*CosmicGameNumHolderNFTWinnersPerRoundChanged, error) {
	event := new(CosmicGameNumHolderNFTWinnersPerRoundChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "NumHolderNFTWinnersPerRoundChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameNumRaffleNFTWinnersPerRoundChangedIterator is returned from FilterNumRaffleNFTWinnersPerRoundChanged and is used to iterate over the raw logs and unpacked data for NumRaffleNFTWinnersPerRoundChanged events raised by the CosmicGame contract.
type CosmicGameNumRaffleNFTWinnersPerRoundChangedIterator struct {
	Event *CosmicGameNumRaffleNFTWinnersPerRoundChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameNumRaffleNFTWinnersPerRoundChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameNumRaffleNFTWinnersPerRoundChanged)
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
		it.Event = new(CosmicGameNumRaffleNFTWinnersPerRoundChanged)
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
func (it *CosmicGameNumRaffleNFTWinnersPerRoundChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameNumRaffleNFTWinnersPerRoundChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameNumRaffleNFTWinnersPerRoundChanged represents a NumRaffleNFTWinnersPerRoundChanged event raised by the CosmicGame contract.
type CosmicGameNumRaffleNFTWinnersPerRoundChanged struct {
	NewNumRaffleNFTWinnersPerRound *big.Int
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterNumRaffleNFTWinnersPerRoundChanged is a free log retrieval operation binding the contract event 0x72e4278828b8a868e0ba5b4887f954797ea786f8bac991128636171f5eed471d.
//
// Solidity: event NumRaffleNFTWinnersPerRoundChanged(uint256 newNumRaffleNFTWinnersPerRound)
func (_CosmicGame *CosmicGameFilterer) FilterNumRaffleNFTWinnersPerRoundChanged(opts *bind.FilterOpts) (*CosmicGameNumRaffleNFTWinnersPerRoundChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "NumRaffleNFTWinnersPerRoundChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameNumRaffleNFTWinnersPerRoundChangedIterator{contract: _CosmicGame.contract, event: "NumRaffleNFTWinnersPerRoundChanged", logs: logs, sub: sub}, nil
}

// WatchNumRaffleNFTWinnersPerRoundChanged is a free log subscription operation binding the contract event 0x72e4278828b8a868e0ba5b4887f954797ea786f8bac991128636171f5eed471d.
//
// Solidity: event NumRaffleNFTWinnersPerRoundChanged(uint256 newNumRaffleNFTWinnersPerRound)
func (_CosmicGame *CosmicGameFilterer) WatchNumRaffleNFTWinnersPerRoundChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameNumRaffleNFTWinnersPerRoundChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "NumRaffleNFTWinnersPerRoundChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameNumRaffleNFTWinnersPerRoundChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "NumRaffleNFTWinnersPerRoundChanged", log); err != nil {
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

// ParseNumRaffleNFTWinnersPerRoundChanged is a log parse operation binding the contract event 0x72e4278828b8a868e0ba5b4887f954797ea786f8bac991128636171f5eed471d.
//
// Solidity: event NumRaffleNFTWinnersPerRoundChanged(uint256 newNumRaffleNFTWinnersPerRound)
func (_CosmicGame *CosmicGameFilterer) ParseNumRaffleNFTWinnersPerRoundChanged(log types.Log) (*CosmicGameNumRaffleNFTWinnersPerRoundChanged, error) {
	event := new(CosmicGameNumRaffleNFTWinnersPerRoundChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "NumRaffleNFTWinnersPerRoundChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameNumRaffleWinnersPerRoundChangedIterator is returned from FilterNumRaffleWinnersPerRoundChanged and is used to iterate over the raw logs and unpacked data for NumRaffleWinnersPerRoundChanged events raised by the CosmicGame contract.
type CosmicGameNumRaffleWinnersPerRoundChangedIterator struct {
	Event *CosmicGameNumRaffleWinnersPerRoundChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameNumRaffleWinnersPerRoundChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameNumRaffleWinnersPerRoundChanged)
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
		it.Event = new(CosmicGameNumRaffleWinnersPerRoundChanged)
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
func (it *CosmicGameNumRaffleWinnersPerRoundChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameNumRaffleWinnersPerRoundChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameNumRaffleWinnersPerRoundChanged represents a NumRaffleWinnersPerRoundChanged event raised by the CosmicGame contract.
type CosmicGameNumRaffleWinnersPerRoundChanged struct {
	NewNumRaffleWinnersPerRound *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterNumRaffleWinnersPerRoundChanged is a free log retrieval operation binding the contract event 0x5e3567ae359e61c47ad2f11d2828dea7e062b2ff63dea385cdb7900a50504c7c.
//
// Solidity: event NumRaffleWinnersPerRoundChanged(uint256 newNumRaffleWinnersPerRound)
func (_CosmicGame *CosmicGameFilterer) FilterNumRaffleWinnersPerRoundChanged(opts *bind.FilterOpts) (*CosmicGameNumRaffleWinnersPerRoundChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "NumRaffleWinnersPerRoundChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameNumRaffleWinnersPerRoundChangedIterator{contract: _CosmicGame.contract, event: "NumRaffleWinnersPerRoundChanged", logs: logs, sub: sub}, nil
}

// WatchNumRaffleWinnersPerRoundChanged is a free log subscription operation binding the contract event 0x5e3567ae359e61c47ad2f11d2828dea7e062b2ff63dea385cdb7900a50504c7c.
//
// Solidity: event NumRaffleWinnersPerRoundChanged(uint256 newNumRaffleWinnersPerRound)
func (_CosmicGame *CosmicGameFilterer) WatchNumRaffleWinnersPerRoundChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameNumRaffleWinnersPerRoundChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "NumRaffleWinnersPerRoundChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameNumRaffleWinnersPerRoundChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "NumRaffleWinnersPerRoundChanged", log); err != nil {
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

// ParseNumRaffleWinnersPerRoundChanged is a log parse operation binding the contract event 0x5e3567ae359e61c47ad2f11d2828dea7e062b2ff63dea385cdb7900a50504c7c.
//
// Solidity: event NumRaffleWinnersPerRoundChanged(uint256 newNumRaffleWinnersPerRound)
func (_CosmicGame *CosmicGameFilterer) ParseNumRaffleWinnersPerRoundChanged(log types.Log) (*CosmicGameNumRaffleWinnersPerRoundChanged, error) {
	event := new(CosmicGameNumRaffleWinnersPerRoundChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "NumRaffleWinnersPerRoundChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CosmicGame contract.
type CosmicGameOwnershipTransferredIterator struct {
	Event *CosmicGameOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CosmicGameOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameOwnershipTransferred)
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
		it.Event = new(CosmicGameOwnershipTransferred)
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
func (it *CosmicGameOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameOwnershipTransferred represents a OwnershipTransferred event raised by the CosmicGame contract.
type CosmicGameOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CosmicGame *CosmicGameFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CosmicGameOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CosmicGameOwnershipTransferredIterator{contract: _CosmicGame.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CosmicGame *CosmicGameFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CosmicGameOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameOwnershipTransferred)
				if err := _CosmicGame.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CosmicGame *CosmicGameFilterer) ParseOwnershipTransferred(log types.Log) (*CosmicGameOwnershipTransferred, error) {
	event := new(CosmicGameOwnershipTransferred)
	if err := _CosmicGame.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGamePriceIncreaseChangedIterator is returned from FilterPriceIncreaseChanged and is used to iterate over the raw logs and unpacked data for PriceIncreaseChanged events raised by the CosmicGame contract.
type CosmicGamePriceIncreaseChangedIterator struct {
	Event *CosmicGamePriceIncreaseChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGamePriceIncreaseChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGamePriceIncreaseChanged)
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
		it.Event = new(CosmicGamePriceIncreaseChanged)
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
func (it *CosmicGamePriceIncreaseChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGamePriceIncreaseChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGamePriceIncreaseChanged represents a PriceIncreaseChanged event raised by the CosmicGame contract.
type CosmicGamePriceIncreaseChanged struct {
	NewPriceIncrease *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPriceIncreaseChanged is a free log retrieval operation binding the contract event 0xcbc1f49adfa29e2f2f0f5c9e057722496a4bc95a6a5446deaa423a02b30c64ac.
//
// Solidity: event PriceIncreaseChanged(uint256 newPriceIncrease)
func (_CosmicGame *CosmicGameFilterer) FilterPriceIncreaseChanged(opts *bind.FilterOpts) (*CosmicGamePriceIncreaseChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "PriceIncreaseChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGamePriceIncreaseChangedIterator{contract: _CosmicGame.contract, event: "PriceIncreaseChanged", logs: logs, sub: sub}, nil
}

// WatchPriceIncreaseChanged is a free log subscription operation binding the contract event 0xcbc1f49adfa29e2f2f0f5c9e057722496a4bc95a6a5446deaa423a02b30c64ac.
//
// Solidity: event PriceIncreaseChanged(uint256 newPriceIncrease)
func (_CosmicGame *CosmicGameFilterer) WatchPriceIncreaseChanged(opts *bind.WatchOpts, sink chan<- *CosmicGamePriceIncreaseChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "PriceIncreaseChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGamePriceIncreaseChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "PriceIncreaseChanged", log); err != nil {
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

// ParsePriceIncreaseChanged is a log parse operation binding the contract event 0xcbc1f49adfa29e2f2f0f5c9e057722496a4bc95a6a5446deaa423a02b30c64ac.
//
// Solidity: event PriceIncreaseChanged(uint256 newPriceIncrease)
func (_CosmicGame *CosmicGameFilterer) ParsePriceIncreaseChanged(log types.Log) (*CosmicGamePriceIncreaseChanged, error) {
	event := new(CosmicGamePriceIncreaseChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "PriceIncreaseChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGamePrizeClaimEventIterator is returned from FilterPrizeClaimEvent and is used to iterate over the raw logs and unpacked data for PrizeClaimEvent events raised by the CosmicGame contract.
type CosmicGamePrizeClaimEventIterator struct {
	Event *CosmicGamePrizeClaimEvent // Event containing the contract specifics and raw log

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
func (it *CosmicGamePrizeClaimEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGamePrizeClaimEvent)
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
		it.Event = new(CosmicGamePrizeClaimEvent)
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
func (it *CosmicGamePrizeClaimEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGamePrizeClaimEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGamePrizeClaimEvent represents a PrizeClaimEvent event raised by the CosmicGame contract.
type CosmicGamePrizeClaimEvent struct {
	PrizeNum    *big.Int
	Destination common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPrizeClaimEvent is a free log retrieval operation binding the contract event 0x27bc828c399c2947fea27bca8a75ced2e94ff2651d607271f051e39db52286ce.
//
// Solidity: event PrizeClaimEvent(uint256 indexed prizeNum, address indexed destination, uint256 amount)
func (_CosmicGame *CosmicGameFilterer) FilterPrizeClaimEvent(opts *bind.FilterOpts, prizeNum []*big.Int, destination []common.Address) (*CosmicGamePrizeClaimEventIterator, error) {

	var prizeNumRule []interface{}
	for _, prizeNumItem := range prizeNum {
		prizeNumRule = append(prizeNumRule, prizeNumItem)
	}
	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "PrizeClaimEvent", prizeNumRule, destinationRule)
	if err != nil {
		return nil, err
	}
	return &CosmicGamePrizeClaimEventIterator{contract: _CosmicGame.contract, event: "PrizeClaimEvent", logs: logs, sub: sub}, nil
}

// WatchPrizeClaimEvent is a free log subscription operation binding the contract event 0x27bc828c399c2947fea27bca8a75ced2e94ff2651d607271f051e39db52286ce.
//
// Solidity: event PrizeClaimEvent(uint256 indexed prizeNum, address indexed destination, uint256 amount)
func (_CosmicGame *CosmicGameFilterer) WatchPrizeClaimEvent(opts *bind.WatchOpts, sink chan<- *CosmicGamePrizeClaimEvent, prizeNum []*big.Int, destination []common.Address) (event.Subscription, error) {

	var prizeNumRule []interface{}
	for _, prizeNumItem := range prizeNum {
		prizeNumRule = append(prizeNumRule, prizeNumItem)
	}
	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "PrizeClaimEvent", prizeNumRule, destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGamePrizeClaimEvent)
				if err := _CosmicGame.contract.UnpackLog(event, "PrizeClaimEvent", log); err != nil {
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

// ParsePrizeClaimEvent is a log parse operation binding the contract event 0x27bc828c399c2947fea27bca8a75ced2e94ff2651d607271f051e39db52286ce.
//
// Solidity: event PrizeClaimEvent(uint256 indexed prizeNum, address indexed destination, uint256 amount)
func (_CosmicGame *CosmicGameFilterer) ParsePrizeClaimEvent(log types.Log) (*CosmicGamePrizeClaimEvent, error) {
	event := new(CosmicGamePrizeClaimEvent)
	if err := _CosmicGame.contract.UnpackLog(event, "PrizeClaimEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGamePrizePercentageChangedIterator is returned from FilterPrizePercentageChanged and is used to iterate over the raw logs and unpacked data for PrizePercentageChanged events raised by the CosmicGame contract.
type CosmicGamePrizePercentageChangedIterator struct {
	Event *CosmicGamePrizePercentageChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGamePrizePercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGamePrizePercentageChanged)
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
		it.Event = new(CosmicGamePrizePercentageChanged)
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
func (it *CosmicGamePrizePercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGamePrizePercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGamePrizePercentageChanged represents a PrizePercentageChanged event raised by the CosmicGame contract.
type CosmicGamePrizePercentageChanged struct {
	NewPrizePercentage *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPrizePercentageChanged is a free log retrieval operation binding the contract event 0x595fa5ba64cd6f66df19b53b59bb4a275bae1ec5b362da95e1fa4c5feb62be1e.
//
// Solidity: event PrizePercentageChanged(uint256 newPrizePercentage)
func (_CosmicGame *CosmicGameFilterer) FilterPrizePercentageChanged(opts *bind.FilterOpts) (*CosmicGamePrizePercentageChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "PrizePercentageChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGamePrizePercentageChangedIterator{contract: _CosmicGame.contract, event: "PrizePercentageChanged", logs: logs, sub: sub}, nil
}

// WatchPrizePercentageChanged is a free log subscription operation binding the contract event 0x595fa5ba64cd6f66df19b53b59bb4a275bae1ec5b362da95e1fa4c5feb62be1e.
//
// Solidity: event PrizePercentageChanged(uint256 newPrizePercentage)
func (_CosmicGame *CosmicGameFilterer) WatchPrizePercentageChanged(opts *bind.WatchOpts, sink chan<- *CosmicGamePrizePercentageChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "PrizePercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGamePrizePercentageChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "PrizePercentageChanged", log); err != nil {
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

// ParsePrizePercentageChanged is a log parse operation binding the contract event 0x595fa5ba64cd6f66df19b53b59bb4a275bae1ec5b362da95e1fa4c5feb62be1e.
//
// Solidity: event PrizePercentageChanged(uint256 newPrizePercentage)
func (_CosmicGame *CosmicGameFilterer) ParsePrizePercentageChanged(log types.Log) (*CosmicGamePrizePercentageChanged, error) {
	event := new(CosmicGamePrizePercentageChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "PrizePercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameRaffleNFTClaimedEventIterator is returned from FilterRaffleNFTClaimedEvent and is used to iterate over the raw logs and unpacked data for RaffleNFTClaimedEvent events raised by the CosmicGame contract.
type CosmicGameRaffleNFTClaimedEventIterator struct {
	Event *CosmicGameRaffleNFTClaimedEvent // Event containing the contract specifics and raw log

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
func (it *CosmicGameRaffleNFTClaimedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameRaffleNFTClaimedEvent)
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
		it.Event = new(CosmicGameRaffleNFTClaimedEvent)
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
func (it *CosmicGameRaffleNFTClaimedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameRaffleNFTClaimedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameRaffleNFTClaimedEvent represents a RaffleNFTClaimedEvent event raised by the CosmicGame contract.
type CosmicGameRaffleNFTClaimedEvent struct {
	Winner common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRaffleNFTClaimedEvent is a free log retrieval operation binding the contract event 0xe05ba2c5fcd9a60f30b179cb0e775070cc8ce9667b0e663e984ee6a02f694cee.
//
// Solidity: event RaffleNFTClaimedEvent(address indexed winner)
func (_CosmicGame *CosmicGameFilterer) FilterRaffleNFTClaimedEvent(opts *bind.FilterOpts, winner []common.Address) (*CosmicGameRaffleNFTClaimedEventIterator, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "RaffleNFTClaimedEvent", winnerRule)
	if err != nil {
		return nil, err
	}
	return &CosmicGameRaffleNFTClaimedEventIterator{contract: _CosmicGame.contract, event: "RaffleNFTClaimedEvent", logs: logs, sub: sub}, nil
}

// WatchRaffleNFTClaimedEvent is a free log subscription operation binding the contract event 0xe05ba2c5fcd9a60f30b179cb0e775070cc8ce9667b0e663e984ee6a02f694cee.
//
// Solidity: event RaffleNFTClaimedEvent(address indexed winner)
func (_CosmicGame *CosmicGameFilterer) WatchRaffleNFTClaimedEvent(opts *bind.WatchOpts, sink chan<- *CosmicGameRaffleNFTClaimedEvent, winner []common.Address) (event.Subscription, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "RaffleNFTClaimedEvent", winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameRaffleNFTClaimedEvent)
				if err := _CosmicGame.contract.UnpackLog(event, "RaffleNFTClaimedEvent", log); err != nil {
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

// ParseRaffleNFTClaimedEvent is a log parse operation binding the contract event 0xe05ba2c5fcd9a60f30b179cb0e775070cc8ce9667b0e663e984ee6a02f694cee.
//
// Solidity: event RaffleNFTClaimedEvent(address indexed winner)
func (_CosmicGame *CosmicGameFilterer) ParseRaffleNFTClaimedEvent(log types.Log) (*CosmicGameRaffleNFTClaimedEvent, error) {
	event := new(CosmicGameRaffleNFTClaimedEvent)
	if err := _CosmicGame.contract.UnpackLog(event, "RaffleNFTClaimedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameRaffleNFTWinnerEventIterator is returned from FilterRaffleNFTWinnerEvent and is used to iterate over the raw logs and unpacked data for RaffleNFTWinnerEvent events raised by the CosmicGame contract.
type CosmicGameRaffleNFTWinnerEventIterator struct {
	Event *CosmicGameRaffleNFTWinnerEvent // Event containing the contract specifics and raw log

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
func (it *CosmicGameRaffleNFTWinnerEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameRaffleNFTWinnerEvent)
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
		it.Event = new(CosmicGameRaffleNFTWinnerEvent)
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
func (it *CosmicGameRaffleNFTWinnerEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameRaffleNFTWinnerEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameRaffleNFTWinnerEvent represents a RaffleNFTWinnerEvent event raised by the CosmicGame contract.
type CosmicGameRaffleNFTWinnerEvent struct {
	Winner      common.Address
	Round       *big.Int
	WinnerIndex *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRaffleNFTWinnerEvent is a free log retrieval operation binding the contract event 0x80348bf864c08069d1368c42ed36b7a60560f73267f63d58e9be69f4b021bacc.
//
// Solidity: event RaffleNFTWinnerEvent(address indexed winner, uint256 indexed round, uint256 winnerIndex)
func (_CosmicGame *CosmicGameFilterer) FilterRaffleNFTWinnerEvent(opts *bind.FilterOpts, winner []common.Address, round []*big.Int) (*CosmicGameRaffleNFTWinnerEventIterator, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "RaffleNFTWinnerEvent", winnerRule, roundRule)
	if err != nil {
		return nil, err
	}
	return &CosmicGameRaffleNFTWinnerEventIterator{contract: _CosmicGame.contract, event: "RaffleNFTWinnerEvent", logs: logs, sub: sub}, nil
}

// WatchRaffleNFTWinnerEvent is a free log subscription operation binding the contract event 0x80348bf864c08069d1368c42ed36b7a60560f73267f63d58e9be69f4b021bacc.
//
// Solidity: event RaffleNFTWinnerEvent(address indexed winner, uint256 indexed round, uint256 winnerIndex)
func (_CosmicGame *CosmicGameFilterer) WatchRaffleNFTWinnerEvent(opts *bind.WatchOpts, sink chan<- *CosmicGameRaffleNFTWinnerEvent, winner []common.Address, round []*big.Int) (event.Subscription, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "RaffleNFTWinnerEvent", winnerRule, roundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameRaffleNFTWinnerEvent)
				if err := _CosmicGame.contract.UnpackLog(event, "RaffleNFTWinnerEvent", log); err != nil {
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

// ParseRaffleNFTWinnerEvent is a log parse operation binding the contract event 0x80348bf864c08069d1368c42ed36b7a60560f73267f63d58e9be69f4b021bacc.
//
// Solidity: event RaffleNFTWinnerEvent(address indexed winner, uint256 indexed round, uint256 winnerIndex)
func (_CosmicGame *CosmicGameFilterer) ParseRaffleNFTWinnerEvent(log types.Log) (*CosmicGameRaffleNFTWinnerEvent, error) {
	event := new(CosmicGameRaffleNFTWinnerEvent)
	if err := _CosmicGame.contract.UnpackLog(event, "RaffleNFTWinnerEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameRafflePercentageChangedIterator is returned from FilterRafflePercentageChanged and is used to iterate over the raw logs and unpacked data for RafflePercentageChanged events raised by the CosmicGame contract.
type CosmicGameRafflePercentageChangedIterator struct {
	Event *CosmicGameRafflePercentageChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameRafflePercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameRafflePercentageChanged)
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
		it.Event = new(CosmicGameRafflePercentageChanged)
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
func (it *CosmicGameRafflePercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameRafflePercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameRafflePercentageChanged represents a RafflePercentageChanged event raised by the CosmicGame contract.
type CosmicGameRafflePercentageChanged struct {
	NewRafflePercentage *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRafflePercentageChanged is a free log retrieval operation binding the contract event 0xd2002a256ab2f8a5c1b865685754b20437c0a26e7249e40ae5df2993966f99f2.
//
// Solidity: event RafflePercentageChanged(uint256 newRafflePercentage)
func (_CosmicGame *CosmicGameFilterer) FilterRafflePercentageChanged(opts *bind.FilterOpts) (*CosmicGameRafflePercentageChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "RafflePercentageChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameRafflePercentageChangedIterator{contract: _CosmicGame.contract, event: "RafflePercentageChanged", logs: logs, sub: sub}, nil
}

// WatchRafflePercentageChanged is a free log subscription operation binding the contract event 0xd2002a256ab2f8a5c1b865685754b20437c0a26e7249e40ae5df2993966f99f2.
//
// Solidity: event RafflePercentageChanged(uint256 newRafflePercentage)
func (_CosmicGame *CosmicGameFilterer) WatchRafflePercentageChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameRafflePercentageChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "RafflePercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameRafflePercentageChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "RafflePercentageChanged", log); err != nil {
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

// ParseRafflePercentageChanged is a log parse operation binding the contract event 0xd2002a256ab2f8a5c1b865685754b20437c0a26e7249e40ae5df2993966f99f2.
//
// Solidity: event RafflePercentageChanged(uint256 newRafflePercentage)
func (_CosmicGame *CosmicGameFilterer) ParseRafflePercentageChanged(log types.Log) (*CosmicGameRafflePercentageChanged, error) {
	event := new(CosmicGameRafflePercentageChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "RafflePercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameRaffleWalletAddressChangedIterator is returned from FilterRaffleWalletAddressChanged and is used to iterate over the raw logs and unpacked data for RaffleWalletAddressChanged events raised by the CosmicGame contract.
type CosmicGameRaffleWalletAddressChangedIterator struct {
	Event *CosmicGameRaffleWalletAddressChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameRaffleWalletAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameRaffleWalletAddressChanged)
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
		it.Event = new(CosmicGameRaffleWalletAddressChanged)
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
func (it *CosmicGameRaffleWalletAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameRaffleWalletAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameRaffleWalletAddressChanged represents a RaffleWalletAddressChanged event raised by the CosmicGame contract.
type CosmicGameRaffleWalletAddressChanged struct {
	NewRaffleWallet common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRaffleWalletAddressChanged is a free log retrieval operation binding the contract event 0x508d510537704d37ed64691a706392abc33e59ce8a327b4952e112820ceb88a6.
//
// Solidity: event RaffleWalletAddressChanged(address newRaffleWallet)
func (_CosmicGame *CosmicGameFilterer) FilterRaffleWalletAddressChanged(opts *bind.FilterOpts) (*CosmicGameRaffleWalletAddressChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "RaffleWalletAddressChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameRaffleWalletAddressChangedIterator{contract: _CosmicGame.contract, event: "RaffleWalletAddressChanged", logs: logs, sub: sub}, nil
}

// WatchRaffleWalletAddressChanged is a free log subscription operation binding the contract event 0x508d510537704d37ed64691a706392abc33e59ce8a327b4952e112820ceb88a6.
//
// Solidity: event RaffleWalletAddressChanged(address newRaffleWallet)
func (_CosmicGame *CosmicGameFilterer) WatchRaffleWalletAddressChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameRaffleWalletAddressChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "RaffleWalletAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameRaffleWalletAddressChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "RaffleWalletAddressChanged", log); err != nil {
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

// ParseRaffleWalletAddressChanged is a log parse operation binding the contract event 0x508d510537704d37ed64691a706392abc33e59ce8a327b4952e112820ceb88a6.
//
// Solidity: event RaffleWalletAddressChanged(address newRaffleWallet)
func (_CosmicGame *CosmicGameFilterer) ParseRaffleWalletAddressChanged(log types.Log) (*CosmicGameRaffleWalletAddressChanged, error) {
	event := new(CosmicGameRaffleWalletAddressChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "RaffleWalletAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameRandomWalkAddressChangedIterator is returned from FilterRandomWalkAddressChanged and is used to iterate over the raw logs and unpacked data for RandomWalkAddressChanged events raised by the CosmicGame contract.
type CosmicGameRandomWalkAddressChangedIterator struct {
	Event *CosmicGameRandomWalkAddressChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameRandomWalkAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameRandomWalkAddressChanged)
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
		it.Event = new(CosmicGameRandomWalkAddressChanged)
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
func (it *CosmicGameRandomWalkAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameRandomWalkAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameRandomWalkAddressChanged represents a RandomWalkAddressChanged event raised by the CosmicGame contract.
type CosmicGameRandomWalkAddressChanged struct {
	NewRandomWalk common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRandomWalkAddressChanged is a free log retrieval operation binding the contract event 0x9dc3707e4b88be65295b0083b44ffa94059c80372add9b0d24d6a3b371a03b92.
//
// Solidity: event RandomWalkAddressChanged(address newRandomWalk)
func (_CosmicGame *CosmicGameFilterer) FilterRandomWalkAddressChanged(opts *bind.FilterOpts) (*CosmicGameRandomWalkAddressChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "RandomWalkAddressChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameRandomWalkAddressChangedIterator{contract: _CosmicGame.contract, event: "RandomWalkAddressChanged", logs: logs, sub: sub}, nil
}

// WatchRandomWalkAddressChanged is a free log subscription operation binding the contract event 0x9dc3707e4b88be65295b0083b44ffa94059c80372add9b0d24d6a3b371a03b92.
//
// Solidity: event RandomWalkAddressChanged(address newRandomWalk)
func (_CosmicGame *CosmicGameFilterer) WatchRandomWalkAddressChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameRandomWalkAddressChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "RandomWalkAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameRandomWalkAddressChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "RandomWalkAddressChanged", log); err != nil {
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

// ParseRandomWalkAddressChanged is a log parse operation binding the contract event 0x9dc3707e4b88be65295b0083b44ffa94059c80372add9b0d24d6a3b371a03b92.
//
// Solidity: event RandomWalkAddressChanged(address newRandomWalk)
func (_CosmicGame *CosmicGameFilterer) ParseRandomWalkAddressChanged(log types.Log) (*CosmicGameRandomWalkAddressChanged, error) {
	event := new(CosmicGameRandomWalkAddressChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "RandomWalkAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameTimeIncreaseChangedIterator is returned from FilterTimeIncreaseChanged and is used to iterate over the raw logs and unpacked data for TimeIncreaseChanged events raised by the CosmicGame contract.
type CosmicGameTimeIncreaseChangedIterator struct {
	Event *CosmicGameTimeIncreaseChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameTimeIncreaseChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameTimeIncreaseChanged)
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
		it.Event = new(CosmicGameTimeIncreaseChanged)
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
func (it *CosmicGameTimeIncreaseChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameTimeIncreaseChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameTimeIncreaseChanged represents a TimeIncreaseChanged event raised by the CosmicGame contract.
type CosmicGameTimeIncreaseChanged struct {
	NewTimeIncrease *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimeIncreaseChanged is a free log retrieval operation binding the contract event 0xed46e73b174ced51fb529cdc6c1b4d8abf49387e6d849b71648afb63c81d12cd.
//
// Solidity: event TimeIncreaseChanged(uint256 newTimeIncrease)
func (_CosmicGame *CosmicGameFilterer) FilterTimeIncreaseChanged(opts *bind.FilterOpts) (*CosmicGameTimeIncreaseChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "TimeIncreaseChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameTimeIncreaseChangedIterator{contract: _CosmicGame.contract, event: "TimeIncreaseChanged", logs: logs, sub: sub}, nil
}

// WatchTimeIncreaseChanged is a free log subscription operation binding the contract event 0xed46e73b174ced51fb529cdc6c1b4d8abf49387e6d849b71648afb63c81d12cd.
//
// Solidity: event TimeIncreaseChanged(uint256 newTimeIncrease)
func (_CosmicGame *CosmicGameFilterer) WatchTimeIncreaseChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameTimeIncreaseChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "TimeIncreaseChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameTimeIncreaseChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "TimeIncreaseChanged", log); err != nil {
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

// ParseTimeIncreaseChanged is a log parse operation binding the contract event 0xed46e73b174ced51fb529cdc6c1b4d8abf49387e6d849b71648afb63c81d12cd.
//
// Solidity: event TimeIncreaseChanged(uint256 newTimeIncrease)
func (_CosmicGame *CosmicGameFilterer) ParseTimeIncreaseChanged(log types.Log) (*CosmicGameTimeIncreaseChanged, error) {
	event := new(CosmicGameTimeIncreaseChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "TimeIncreaseChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CountersMetaData contains all meta data concerning the Counters contract.
var CountersMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212202b30207ad4c731f97f8f5c63a174f7a19eb2f35182448202f31f4ff8d52b1a2f64736f6c63430008130033",
}

// CountersABI is the input ABI used to generate the binding from.
// Deprecated: Use CountersMetaData.ABI instead.
var CountersABI = CountersMetaData.ABI

// CountersBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CountersMetaData.Bin instead.
var CountersBin = CountersMetaData.Bin

// DeployCounters deploys a new Ethereum contract, binding an instance of Counters to it.
func DeployCounters(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Counters, error) {
	parsed, err := CountersMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CountersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Counters{CountersCaller: CountersCaller{contract: contract}, CountersTransactor: CountersTransactor{contract: contract}, CountersFilterer: CountersFilterer{contract: contract}}, nil
}

// Counters is an auto generated Go binding around an Ethereum contract.
type Counters struct {
	CountersCaller     // Read-only binding to the contract
	CountersTransactor // Write-only binding to the contract
	CountersFilterer   // Log filterer for contract events
}

// CountersCaller is an auto generated read-only Go binding around an Ethereum contract.
type CountersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CountersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CountersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CountersSession struct {
	Contract     *Counters         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CountersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CountersCallerSession struct {
	Contract *CountersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CountersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CountersTransactorSession struct {
	Contract     *CountersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CountersRaw is an auto generated low-level Go binding around an Ethereum contract.
type CountersRaw struct {
	Contract *Counters // Generic contract binding to access the raw methods on
}

// CountersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CountersCallerRaw struct {
	Contract *CountersCaller // Generic read-only contract binding to access the raw methods on
}

// CountersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CountersTransactorRaw struct {
	Contract *CountersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCounters creates a new instance of Counters, bound to a specific deployed contract.
func NewCounters(address common.Address, backend bind.ContractBackend) (*Counters, error) {
	contract, err := bindCounters(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Counters{CountersCaller: CountersCaller{contract: contract}, CountersTransactor: CountersTransactor{contract: contract}, CountersFilterer: CountersFilterer{contract: contract}}, nil
}

// NewCountersCaller creates a new read-only instance of Counters, bound to a specific deployed contract.
func NewCountersCaller(address common.Address, caller bind.ContractCaller) (*CountersCaller, error) {
	contract, err := bindCounters(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CountersCaller{contract: contract}, nil
}

// NewCountersTransactor creates a new write-only instance of Counters, bound to a specific deployed contract.
func NewCountersTransactor(address common.Address, transactor bind.ContractTransactor) (*CountersTransactor, error) {
	contract, err := bindCounters(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CountersTransactor{contract: contract}, nil
}

// NewCountersFilterer creates a new log filterer instance of Counters, bound to a specific deployed contract.
func NewCountersFilterer(address common.Address, filterer bind.ContractFilterer) (*CountersFilterer, error) {
	contract, err := bindCounters(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CountersFilterer{contract: contract}, nil
}

// bindCounters binds a generic wrapper to an already deployed contract.
func bindCounters(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CountersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counters *CountersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counters.Contract.CountersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counters *CountersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counters.Contract.CountersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counters *CountersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counters.Contract.CountersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counters *CountersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counters.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counters *CountersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counters.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counters *CountersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counters.Contract.contract.Transact(opts, method, params...)
}

// ECDSAMetaData contains all meta data concerning the ECDSA contract.
var ECDSAMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220fddca1388b4cfc79615214a0c4649257a2326f1b0fbd4f2a3cfbd72f6a47c3ca64736f6c63430008130033",
}

// ECDSAABI is the input ABI used to generate the binding from.
// Deprecated: Use ECDSAMetaData.ABI instead.
var ECDSAABI = ECDSAMetaData.ABI

// ECDSABin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ECDSAMetaData.Bin instead.
var ECDSABin = ECDSAMetaData.Bin

// DeployECDSA deploys a new Ethereum contract, binding an instance of ECDSA to it.
func DeployECDSA(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECDSA, error) {
	parsed, err := ECDSAMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ECDSABin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// ECDSA is an auto generated Go binding around an Ethereum contract.
type ECDSA struct {
	ECDSACaller     // Read-only binding to the contract
	ECDSATransactor // Write-only binding to the contract
	ECDSAFilterer   // Log filterer for contract events
}

// ECDSACaller is an auto generated read-only Go binding around an Ethereum contract.
type ECDSACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSATransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECDSATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECDSAFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECDSASession struct {
	Contract     *ECDSA            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECDSACallerSession struct {
	Contract *ECDSACaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ECDSATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECDSATransactorSession struct {
	Contract     *ECDSATransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSARaw is an auto generated low-level Go binding around an Ethereum contract.
type ECDSARaw struct {
	Contract *ECDSA // Generic contract binding to access the raw methods on
}

// ECDSACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECDSACallerRaw struct {
	Contract *ECDSACaller // Generic read-only contract binding to access the raw methods on
}

// ECDSATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECDSATransactorRaw struct {
	Contract *ECDSATransactor // Generic write-only contract binding to access the raw methods on
}

// NewECDSA creates a new instance of ECDSA, bound to a specific deployed contract.
func NewECDSA(address common.Address, backend bind.ContractBackend) (*ECDSA, error) {
	contract, err := bindECDSA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// NewECDSACaller creates a new read-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSACaller(address common.Address, caller bind.ContractCaller) (*ECDSACaller, error) {
	contract, err := bindECDSA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSACaller{contract: contract}, nil
}

// NewECDSATransactor creates a new write-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSATransactor(address common.Address, transactor bind.ContractTransactor) (*ECDSATransactor, error) {
	contract, err := bindECDSA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSATransactor{contract: contract}, nil
}

// NewECDSAFilterer creates a new log filterer instance of ECDSA, bound to a specific deployed contract.
func NewECDSAFilterer(address common.Address, filterer bind.ContractFilterer) (*ECDSAFilterer, error) {
	contract, err := bindECDSA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECDSAFilterer{contract: contract}, nil
}

// bindECDSA binds a generic wrapper to an already deployed contract.
func bindECDSA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECDSAABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSARaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.ECDSACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSACallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transact(opts, method, params...)
}

// EIP712MetaData contains all meta data concerning the EIP712 contract.
var EIP712MetaData = &bind.MetaData{
	ABI: "[]",
}

// EIP712ABI is the input ABI used to generate the binding from.
// Deprecated: Use EIP712MetaData.ABI instead.
var EIP712ABI = EIP712MetaData.ABI

// EIP712 is an auto generated Go binding around an Ethereum contract.
type EIP712 struct {
	EIP712Caller     // Read-only binding to the contract
	EIP712Transactor // Write-only binding to the contract
	EIP712Filterer   // Log filterer for contract events
}

// EIP712Caller is an auto generated read-only Go binding around an Ethereum contract.
type EIP712Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP712Transactor is an auto generated write-only Go binding around an Ethereum contract.
type EIP712Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP712Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EIP712Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP712Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EIP712Session struct {
	Contract     *EIP712           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EIP712CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EIP712CallerSession struct {
	Contract *EIP712Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EIP712TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EIP712TransactorSession struct {
	Contract     *EIP712Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EIP712Raw is an auto generated low-level Go binding around an Ethereum contract.
type EIP712Raw struct {
	Contract *EIP712 // Generic contract binding to access the raw methods on
}

// EIP712CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EIP712CallerRaw struct {
	Contract *EIP712Caller // Generic read-only contract binding to access the raw methods on
}

// EIP712TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EIP712TransactorRaw struct {
	Contract *EIP712Transactor // Generic write-only contract binding to access the raw methods on
}

// NewEIP712 creates a new instance of EIP712, bound to a specific deployed contract.
func NewEIP712(address common.Address, backend bind.ContractBackend) (*EIP712, error) {
	contract, err := bindEIP712(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EIP712{EIP712Caller: EIP712Caller{contract: contract}, EIP712Transactor: EIP712Transactor{contract: contract}, EIP712Filterer: EIP712Filterer{contract: contract}}, nil
}

// NewEIP712Caller creates a new read-only instance of EIP712, bound to a specific deployed contract.
func NewEIP712Caller(address common.Address, caller bind.ContractCaller) (*EIP712Caller, error) {
	contract, err := bindEIP712(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EIP712Caller{contract: contract}, nil
}

// NewEIP712Transactor creates a new write-only instance of EIP712, bound to a specific deployed contract.
func NewEIP712Transactor(address common.Address, transactor bind.ContractTransactor) (*EIP712Transactor, error) {
	contract, err := bindEIP712(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EIP712Transactor{contract: contract}, nil
}

// NewEIP712Filterer creates a new log filterer instance of EIP712, bound to a specific deployed contract.
func NewEIP712Filterer(address common.Address, filterer bind.ContractFilterer) (*EIP712Filterer, error) {
	contract, err := bindEIP712(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EIP712Filterer{contract: contract}, nil
}

// bindEIP712 binds a generic wrapper to an already deployed contract.
func bindEIP712(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EIP712ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EIP712 *EIP712Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EIP712.Contract.EIP712Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EIP712 *EIP712Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIP712.Contract.EIP712Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EIP712 *EIP712Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EIP712.Contract.EIP712Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EIP712 *EIP712CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EIP712.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EIP712 *EIP712TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIP712.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EIP712 *EIP712TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EIP712.Contract.contract.Transact(opts, method, params...)
}

// ERC165MetaData contains all meta data concerning the ERC165 contract.
var ERC165MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ERC165ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC165MetaData.ABI instead.
var ERC165ABI = ERC165MetaData.ABI

// ERC165 is an auto generated Go binding around an Ethereum contract.
type ERC165 struct {
	ERC165Caller     // Read-only binding to the contract
	ERC165Transactor // Write-only binding to the contract
	ERC165Filterer   // Log filterer for contract events
}

// ERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC165Session struct {
	Contract     *ERC165           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC165CallerSession struct {
	Contract *ERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC165TransactorSession struct {
	Contract     *ERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC165Raw struct {
	Contract *ERC165 // Generic contract binding to access the raw methods on
}

// ERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC165CallerRaw struct {
	Contract *ERC165Caller // Generic read-only contract binding to access the raw methods on
}

// ERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC165TransactorRaw struct {
	Contract *ERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC165 creates a new instance of ERC165, bound to a specific deployed contract.
func NewERC165(address common.Address, backend bind.ContractBackend) (*ERC165, error) {
	contract, err := bindERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC165{ERC165Caller: ERC165Caller{contract: contract}, ERC165Transactor: ERC165Transactor{contract: contract}, ERC165Filterer: ERC165Filterer{contract: contract}}, nil
}

// NewERC165Caller creates a new read-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Caller(address common.Address, caller bind.ContractCaller) (*ERC165Caller, error) {
	contract, err := bindERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Caller{contract: contract}, nil
}

// NewERC165Transactor creates a new write-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC165Transactor, error) {
	contract, err := bindERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Transactor{contract: contract}, nil
}

// NewERC165Filterer creates a new log filterer instance of ERC165, bound to a specific deployed contract.
func NewERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC165Filterer, error) {
	contract, err := bindERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC165Filterer{contract: contract}, nil
}

// bindERC165 binds a generic wrapper to an already deployed contract.
func bindERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.ERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// ERC20BurnableMetaData contains all meta data concerning the ERC20Burnable contract.
var ERC20BurnableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERC20BurnableABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20BurnableMetaData.ABI instead.
var ERC20BurnableABI = ERC20BurnableMetaData.ABI

// ERC20Burnable is an auto generated Go binding around an Ethereum contract.
type ERC20Burnable struct {
	ERC20BurnableCaller     // Read-only binding to the contract
	ERC20BurnableTransactor // Write-only binding to the contract
	ERC20BurnableFilterer   // Log filterer for contract events
}

// ERC20BurnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BurnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BurnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BurnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BurnableSession struct {
	Contract     *ERC20Burnable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BurnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BurnableCallerSession struct {
	Contract *ERC20BurnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ERC20BurnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BurnableTransactorSession struct {
	Contract     *ERC20BurnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC20BurnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BurnableRaw struct {
	Contract *ERC20Burnable // Generic contract binding to access the raw methods on
}

// ERC20BurnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BurnableCallerRaw struct {
	Contract *ERC20BurnableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BurnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BurnableTransactorRaw struct {
	Contract *ERC20BurnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Burnable creates a new instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20Burnable(address common.Address, backend bind.ContractBackend) (*ERC20Burnable, error) {
	contract, err := bindERC20Burnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Burnable{ERC20BurnableCaller: ERC20BurnableCaller{contract: contract}, ERC20BurnableTransactor: ERC20BurnableTransactor{contract: contract}, ERC20BurnableFilterer: ERC20BurnableFilterer{contract: contract}}, nil
}

// NewERC20BurnableCaller creates a new read-only instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20BurnableCaller(address common.Address, caller bind.ContractCaller) (*ERC20BurnableCaller, error) {
	contract, err := bindERC20Burnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableCaller{contract: contract}, nil
}

// NewERC20BurnableTransactor creates a new write-only instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20BurnableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BurnableTransactor, error) {
	contract, err := bindERC20Burnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableTransactor{contract: contract}, nil
}

// NewERC20BurnableFilterer creates a new log filterer instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20BurnableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BurnableFilterer, error) {
	contract, err := bindERC20Burnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableFilterer{contract: contract}, nil
}

// bindERC20Burnable binds a generic wrapper to an already deployed contract.
func bindERC20Burnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BurnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Burnable *ERC20BurnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Burnable.Contract.ERC20BurnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Burnable *ERC20BurnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.ERC20BurnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Burnable *ERC20BurnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.ERC20BurnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Burnable *ERC20BurnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Burnable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Burnable *ERC20BurnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Burnable *ERC20BurnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.Allowance(&_ERC20Burnable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.Allowance(&_ERC20Burnable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.BalanceOf(&_ERC20Burnable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.BalanceOf(&_ERC20Burnable.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Burnable *ERC20BurnableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Burnable *ERC20BurnableSession) Decimals() (uint8, error) {
	return _ERC20Burnable.Contract.Decimals(&_ERC20Burnable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Burnable *ERC20BurnableCallerSession) Decimals() (uint8, error) {
	return _ERC20Burnable.Contract.Decimals(&_ERC20Burnable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Burnable *ERC20BurnableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Burnable *ERC20BurnableSession) Name() (string, error) {
	return _ERC20Burnable.Contract.Name(&_ERC20Burnable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Burnable *ERC20BurnableCallerSession) Name() (string, error) {
	return _ERC20Burnable.Contract.Name(&_ERC20Burnable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Burnable *ERC20BurnableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Burnable *ERC20BurnableSession) Symbol() (string, error) {
	return _ERC20Burnable.Contract.Symbol(&_ERC20Burnable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Burnable *ERC20BurnableCallerSession) Symbol() (string, error) {
	return _ERC20Burnable.Contract.Symbol(&_ERC20Burnable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Burnable *ERC20BurnableSession) TotalSupply() (*big.Int, error) {
	return _ERC20Burnable.Contract.TotalSupply(&_ERC20Burnable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Burnable.Contract.TotalSupply(&_ERC20Burnable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Approve(&_ERC20Burnable.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Approve(&_ERC20Burnable.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Burn(&_ERC20Burnable.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Burn(&_ERC20Burnable.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.BurnFrom(&_ERC20Burnable.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.BurnFrom(&_ERC20Burnable.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.DecreaseAllowance(&_ERC20Burnable.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.DecreaseAllowance(&_ERC20Burnable.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.IncreaseAllowance(&_ERC20Burnable.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.IncreaseAllowance(&_ERC20Burnable.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Transfer(&_ERC20Burnable.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Transfer(&_ERC20Burnable.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.TransferFrom(&_ERC20Burnable.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.TransferFrom(&_ERC20Burnable.TransactOpts, sender, recipient, amount)
}

// ERC20BurnableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Burnable contract.
type ERC20BurnableApprovalIterator struct {
	Event *ERC20BurnableApproval // Event containing the contract specifics and raw log

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
func (it *ERC20BurnableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BurnableApproval)
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
		it.Event = new(ERC20BurnableApproval)
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
func (it *ERC20BurnableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BurnableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BurnableApproval represents a Approval event raised by the ERC20Burnable contract.
type ERC20BurnableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20BurnableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Burnable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableApprovalIterator{contract: _ERC20Burnable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20BurnableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Burnable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BurnableApproval)
				if err := _ERC20Burnable.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) ParseApproval(log types.Log) (*ERC20BurnableApproval, error) {
	event := new(ERC20BurnableApproval)
	if err := _ERC20Burnable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BurnableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Burnable contract.
type ERC20BurnableTransferIterator struct {
	Event *ERC20BurnableTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20BurnableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BurnableTransfer)
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
		it.Event = new(ERC20BurnableTransfer)
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
func (it *ERC20BurnableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BurnableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BurnableTransfer represents a Transfer event raised by the ERC20Burnable contract.
type ERC20BurnableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20BurnableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Burnable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableTransferIterator{contract: _ERC20Burnable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20BurnableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Burnable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BurnableTransfer)
				if err := _ERC20Burnable.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) ParseTransfer(log types.Log) (*ERC20BurnableTransfer, error) {
	event := new(ERC20BurnableTransfer)
	if err := _ERC20Burnable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PermitMetaData contains all meta data concerning the ERC20Permit contract.
var ERC20PermitMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERC20PermitABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20PermitMetaData.ABI instead.
var ERC20PermitABI = ERC20PermitMetaData.ABI

// ERC20Permit is an auto generated Go binding around an Ethereum contract.
type ERC20Permit struct {
	ERC20PermitCaller     // Read-only binding to the contract
	ERC20PermitTransactor // Write-only binding to the contract
	ERC20PermitFilterer   // Log filterer for contract events
}

// ERC20PermitCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20PermitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PermitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20PermitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PermitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20PermitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PermitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20PermitSession struct {
	Contract     *ERC20Permit      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20PermitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20PermitCallerSession struct {
	Contract *ERC20PermitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ERC20PermitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20PermitTransactorSession struct {
	Contract     *ERC20PermitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ERC20PermitRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20PermitRaw struct {
	Contract *ERC20Permit // Generic contract binding to access the raw methods on
}

// ERC20PermitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20PermitCallerRaw struct {
	Contract *ERC20PermitCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20PermitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20PermitTransactorRaw struct {
	Contract *ERC20PermitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Permit creates a new instance of ERC20Permit, bound to a specific deployed contract.
func NewERC20Permit(address common.Address, backend bind.ContractBackend) (*ERC20Permit, error) {
	contract, err := bindERC20Permit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Permit{ERC20PermitCaller: ERC20PermitCaller{contract: contract}, ERC20PermitTransactor: ERC20PermitTransactor{contract: contract}, ERC20PermitFilterer: ERC20PermitFilterer{contract: contract}}, nil
}

// NewERC20PermitCaller creates a new read-only instance of ERC20Permit, bound to a specific deployed contract.
func NewERC20PermitCaller(address common.Address, caller bind.ContractCaller) (*ERC20PermitCaller, error) {
	contract, err := bindERC20Permit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20PermitCaller{contract: contract}, nil
}

// NewERC20PermitTransactor creates a new write-only instance of ERC20Permit, bound to a specific deployed contract.
func NewERC20PermitTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20PermitTransactor, error) {
	contract, err := bindERC20Permit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20PermitTransactor{contract: contract}, nil
}

// NewERC20PermitFilterer creates a new log filterer instance of ERC20Permit, bound to a specific deployed contract.
func NewERC20PermitFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20PermitFilterer, error) {
	contract, err := bindERC20Permit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20PermitFilterer{contract: contract}, nil
}

// bindERC20Permit binds a generic wrapper to an already deployed contract.
func bindERC20Permit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20PermitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Permit *ERC20PermitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Permit.Contract.ERC20PermitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Permit *ERC20PermitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Permit.Contract.ERC20PermitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Permit *ERC20PermitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Permit.Contract.ERC20PermitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Permit *ERC20PermitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Permit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Permit *ERC20PermitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Permit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Permit *ERC20PermitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Permit.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ERC20Permit *ERC20PermitCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20Permit.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ERC20Permit *ERC20PermitSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _ERC20Permit.Contract.DOMAINSEPARATOR(&_ERC20Permit.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ERC20Permit *ERC20PermitCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _ERC20Permit.Contract.DOMAINSEPARATOR(&_ERC20Permit.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Permit *ERC20PermitCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Permit.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Permit *ERC20PermitSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Permit.Contract.Allowance(&_ERC20Permit.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Permit *ERC20PermitCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Permit.Contract.Allowance(&_ERC20Permit.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Permit *ERC20PermitCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Permit.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Permit *ERC20PermitSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Permit.Contract.BalanceOf(&_ERC20Permit.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Permit *ERC20PermitCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Permit.Contract.BalanceOf(&_ERC20Permit.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Permit *ERC20PermitCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20Permit.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Permit *ERC20PermitSession) Decimals() (uint8, error) {
	return _ERC20Permit.Contract.Decimals(&_ERC20Permit.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Permit *ERC20PermitCallerSession) Decimals() (uint8, error) {
	return _ERC20Permit.Contract.Decimals(&_ERC20Permit.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Permit *ERC20PermitCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Permit.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Permit *ERC20PermitSession) Name() (string, error) {
	return _ERC20Permit.Contract.Name(&_ERC20Permit.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Permit *ERC20PermitCallerSession) Name() (string, error) {
	return _ERC20Permit.Contract.Name(&_ERC20Permit.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_ERC20Permit *ERC20PermitCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Permit.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_ERC20Permit *ERC20PermitSession) Nonces(owner common.Address) (*big.Int, error) {
	return _ERC20Permit.Contract.Nonces(&_ERC20Permit.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_ERC20Permit *ERC20PermitCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _ERC20Permit.Contract.Nonces(&_ERC20Permit.CallOpts, owner)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Permit *ERC20PermitCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Permit.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Permit *ERC20PermitSession) Symbol() (string, error) {
	return _ERC20Permit.Contract.Symbol(&_ERC20Permit.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Permit *ERC20PermitCallerSession) Symbol() (string, error) {
	return _ERC20Permit.Contract.Symbol(&_ERC20Permit.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Permit *ERC20PermitCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Permit.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Permit *ERC20PermitSession) TotalSupply() (*big.Int, error) {
	return _ERC20Permit.Contract.TotalSupply(&_ERC20Permit.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Permit *ERC20PermitCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Permit.Contract.TotalSupply(&_ERC20Permit.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Permit *ERC20PermitTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Permit *ERC20PermitSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.Contract.Approve(&_ERC20Permit.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Permit *ERC20PermitTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.Contract.Approve(&_ERC20Permit.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Permit *ERC20PermitTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Permit *ERC20PermitSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.Contract.DecreaseAllowance(&_ERC20Permit.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Permit *ERC20PermitTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.Contract.DecreaseAllowance(&_ERC20Permit.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Permit *ERC20PermitTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Permit *ERC20PermitSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.Contract.IncreaseAllowance(&_ERC20Permit.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Permit *ERC20PermitTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.Contract.IncreaseAllowance(&_ERC20Permit.TransactOpts, spender, addedValue)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ERC20Permit *ERC20PermitTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ERC20Permit.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ERC20Permit *ERC20PermitSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ERC20Permit.Contract.Permit(&_ERC20Permit.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ERC20Permit *ERC20PermitTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ERC20Permit.Contract.Permit(&_ERC20Permit.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Permit *ERC20PermitTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Permit *ERC20PermitSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.Contract.Transfer(&_ERC20Permit.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Permit *ERC20PermitTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.Contract.Transfer(&_ERC20Permit.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Permit *ERC20PermitTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Permit *ERC20PermitSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.Contract.TransferFrom(&_ERC20Permit.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Permit *ERC20PermitTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.Contract.TransferFrom(&_ERC20Permit.TransactOpts, sender, recipient, amount)
}

// ERC20PermitApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Permit contract.
type ERC20PermitApprovalIterator struct {
	Event *ERC20PermitApproval // Event containing the contract specifics and raw log

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
func (it *ERC20PermitApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PermitApproval)
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
		it.Event = new(ERC20PermitApproval)
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
func (it *ERC20PermitApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PermitApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PermitApproval represents a Approval event raised by the ERC20Permit contract.
type ERC20PermitApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Permit *ERC20PermitFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20PermitApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Permit.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PermitApprovalIterator{contract: _ERC20Permit.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Permit *ERC20PermitFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20PermitApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Permit.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PermitApproval)
				if err := _ERC20Permit.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Permit *ERC20PermitFilterer) ParseApproval(log types.Log) (*ERC20PermitApproval, error) {
	event := new(ERC20PermitApproval)
	if err := _ERC20Permit.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PermitTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Permit contract.
type ERC20PermitTransferIterator struct {
	Event *ERC20PermitTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20PermitTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PermitTransfer)
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
		it.Event = new(ERC20PermitTransfer)
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
func (it *ERC20PermitTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PermitTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PermitTransfer represents a Transfer event raised by the ERC20Permit contract.
type ERC20PermitTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Permit *ERC20PermitFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20PermitTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Permit.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PermitTransferIterator{contract: _ERC20Permit.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Permit *ERC20PermitFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20PermitTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Permit.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PermitTransfer)
				if err := _ERC20Permit.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Permit *ERC20PermitFilterer) ParseTransfer(log types.Log) (*ERC20PermitTransfer, error) {
	event := new(ERC20PermitTransfer)
	if err := _ERC20Permit.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20VotesMetaData contains all meta data concerning the ERC20Votes contract.
var ERC20VotesMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromDelegate\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toDelegate\",\"type\":\"address\"}],\"name\":\"DelegateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"DelegateVotesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"pos\",\"type\":\"uint32\"}],\"name\":\"checkpoints\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"fromBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint224\",\"name\":\"votes\",\"type\":\"uint224\"}],\"internalType\":\"structERC20Votes.Checkpoint\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"delegateBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"delegates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getPastTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getPastVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"numCheckpoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERC20VotesABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20VotesMetaData.ABI instead.
var ERC20VotesABI = ERC20VotesMetaData.ABI

// ERC20Votes is an auto generated Go binding around an Ethereum contract.
type ERC20Votes struct {
	ERC20VotesCaller     // Read-only binding to the contract
	ERC20VotesTransactor // Write-only binding to the contract
	ERC20VotesFilterer   // Log filterer for contract events
}

// ERC20VotesCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20VotesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20VotesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20VotesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20VotesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20VotesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20VotesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20VotesSession struct {
	Contract     *ERC20Votes       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20VotesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20VotesCallerSession struct {
	Contract *ERC20VotesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC20VotesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20VotesTransactorSession struct {
	Contract     *ERC20VotesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC20VotesRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20VotesRaw struct {
	Contract *ERC20Votes // Generic contract binding to access the raw methods on
}

// ERC20VotesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20VotesCallerRaw struct {
	Contract *ERC20VotesCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20VotesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20VotesTransactorRaw struct {
	Contract *ERC20VotesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Votes creates a new instance of ERC20Votes, bound to a specific deployed contract.
func NewERC20Votes(address common.Address, backend bind.ContractBackend) (*ERC20Votes, error) {
	contract, err := bindERC20Votes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Votes{ERC20VotesCaller: ERC20VotesCaller{contract: contract}, ERC20VotesTransactor: ERC20VotesTransactor{contract: contract}, ERC20VotesFilterer: ERC20VotesFilterer{contract: contract}}, nil
}

// NewERC20VotesCaller creates a new read-only instance of ERC20Votes, bound to a specific deployed contract.
func NewERC20VotesCaller(address common.Address, caller bind.ContractCaller) (*ERC20VotesCaller, error) {
	contract, err := bindERC20Votes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20VotesCaller{contract: contract}, nil
}

// NewERC20VotesTransactor creates a new write-only instance of ERC20Votes, bound to a specific deployed contract.
func NewERC20VotesTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20VotesTransactor, error) {
	contract, err := bindERC20Votes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20VotesTransactor{contract: contract}, nil
}

// NewERC20VotesFilterer creates a new log filterer instance of ERC20Votes, bound to a specific deployed contract.
func NewERC20VotesFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20VotesFilterer, error) {
	contract, err := bindERC20Votes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20VotesFilterer{contract: contract}, nil
}

// bindERC20Votes binds a generic wrapper to an already deployed contract.
func bindERC20Votes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20VotesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Votes *ERC20VotesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Votes.Contract.ERC20VotesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Votes *ERC20VotesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Votes.Contract.ERC20VotesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Votes *ERC20VotesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Votes.Contract.ERC20VotesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Votes *ERC20VotesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Votes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Votes *ERC20VotesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Votes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Votes *ERC20VotesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Votes.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ERC20Votes *ERC20VotesCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20Votes.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ERC20Votes *ERC20VotesSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _ERC20Votes.Contract.DOMAINSEPARATOR(&_ERC20Votes.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ERC20Votes *ERC20VotesCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _ERC20Votes.Contract.DOMAINSEPARATOR(&_ERC20Votes.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Votes *ERC20VotesCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Votes.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Votes *ERC20VotesSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Votes.Contract.Allowance(&_ERC20Votes.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Votes *ERC20VotesCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Votes.Contract.Allowance(&_ERC20Votes.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Votes *ERC20VotesCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Votes.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Votes *ERC20VotesSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Votes.Contract.BalanceOf(&_ERC20Votes.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Votes *ERC20VotesCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Votes.Contract.BalanceOf(&_ERC20Votes.CallOpts, account)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address account, uint32 pos) view returns((uint32,uint224))
func (_ERC20Votes *ERC20VotesCaller) Checkpoints(opts *bind.CallOpts, account common.Address, pos uint32) (ERC20VotesCheckpoint, error) {
	var out []interface{}
	err := _ERC20Votes.contract.Call(opts, &out, "checkpoints", account, pos)

	if err != nil {
		return *new(ERC20VotesCheckpoint), err
	}

	out0 := *abi.ConvertType(out[0], new(ERC20VotesCheckpoint)).(*ERC20VotesCheckpoint)

	return out0, err

}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address account, uint32 pos) view returns((uint32,uint224))
func (_ERC20Votes *ERC20VotesSession) Checkpoints(account common.Address, pos uint32) (ERC20VotesCheckpoint, error) {
	return _ERC20Votes.Contract.Checkpoints(&_ERC20Votes.CallOpts, account, pos)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address account, uint32 pos) view returns((uint32,uint224))
func (_ERC20Votes *ERC20VotesCallerSession) Checkpoints(account common.Address, pos uint32) (ERC20VotesCheckpoint, error) {
	return _ERC20Votes.Contract.Checkpoints(&_ERC20Votes.CallOpts, account, pos)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Votes *ERC20VotesCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20Votes.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Votes *ERC20VotesSession) Decimals() (uint8, error) {
	return _ERC20Votes.Contract.Decimals(&_ERC20Votes.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Votes *ERC20VotesCallerSession) Decimals() (uint8, error) {
	return _ERC20Votes.Contract.Decimals(&_ERC20Votes.CallOpts)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address account) view returns(address)
func (_ERC20Votes *ERC20VotesCaller) Delegates(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var out []interface{}
	err := _ERC20Votes.contract.Call(opts, &out, "delegates", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address account) view returns(address)
func (_ERC20Votes *ERC20VotesSession) Delegates(account common.Address) (common.Address, error) {
	return _ERC20Votes.Contract.Delegates(&_ERC20Votes.CallOpts, account)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address account) view returns(address)
func (_ERC20Votes *ERC20VotesCallerSession) Delegates(account common.Address) (common.Address, error) {
	return _ERC20Votes.Contract.Delegates(&_ERC20Votes.CallOpts, account)
}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 blockNumber) view returns(uint256)
func (_ERC20Votes *ERC20VotesCaller) GetPastTotalSupply(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Votes.contract.Call(opts, &out, "getPastTotalSupply", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 blockNumber) view returns(uint256)
func (_ERC20Votes *ERC20VotesSession) GetPastTotalSupply(blockNumber *big.Int) (*big.Int, error) {
	return _ERC20Votes.Contract.GetPastTotalSupply(&_ERC20Votes.CallOpts, blockNumber)
}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 blockNumber) view returns(uint256)
func (_ERC20Votes *ERC20VotesCallerSession) GetPastTotalSupply(blockNumber *big.Int) (*big.Int, error) {
	return _ERC20Votes.Contract.GetPastTotalSupply(&_ERC20Votes.CallOpts, blockNumber)
}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address account, uint256 blockNumber) view returns(uint256)
func (_ERC20Votes *ERC20VotesCaller) GetPastVotes(opts *bind.CallOpts, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Votes.contract.Call(opts, &out, "getPastVotes", account, blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address account, uint256 blockNumber) view returns(uint256)
func (_ERC20Votes *ERC20VotesSession) GetPastVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _ERC20Votes.Contract.GetPastVotes(&_ERC20Votes.CallOpts, account, blockNumber)
}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address account, uint256 blockNumber) view returns(uint256)
func (_ERC20Votes *ERC20VotesCallerSession) GetPastVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _ERC20Votes.Contract.GetPastVotes(&_ERC20Votes.CallOpts, account, blockNumber)
}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address account) view returns(uint256)
func (_ERC20Votes *ERC20VotesCaller) GetVotes(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Votes.contract.Call(opts, &out, "getVotes", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address account) view returns(uint256)
func (_ERC20Votes *ERC20VotesSession) GetVotes(account common.Address) (*big.Int, error) {
	return _ERC20Votes.Contract.GetVotes(&_ERC20Votes.CallOpts, account)
}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address account) view returns(uint256)
func (_ERC20Votes *ERC20VotesCallerSession) GetVotes(account common.Address) (*big.Int, error) {
	return _ERC20Votes.Contract.GetVotes(&_ERC20Votes.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Votes *ERC20VotesCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Votes.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Votes *ERC20VotesSession) Name() (string, error) {
	return _ERC20Votes.Contract.Name(&_ERC20Votes.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Votes *ERC20VotesCallerSession) Name() (string, error) {
	return _ERC20Votes.Contract.Name(&_ERC20Votes.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_ERC20Votes *ERC20VotesCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Votes.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_ERC20Votes *ERC20VotesSession) Nonces(owner common.Address) (*big.Int, error) {
	return _ERC20Votes.Contract.Nonces(&_ERC20Votes.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_ERC20Votes *ERC20VotesCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _ERC20Votes.Contract.Nonces(&_ERC20Votes.CallOpts, owner)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address account) view returns(uint32)
func (_ERC20Votes *ERC20VotesCaller) NumCheckpoints(opts *bind.CallOpts, account common.Address) (uint32, error) {
	var out []interface{}
	err := _ERC20Votes.contract.Call(opts, &out, "numCheckpoints", account)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address account) view returns(uint32)
func (_ERC20Votes *ERC20VotesSession) NumCheckpoints(account common.Address) (uint32, error) {
	return _ERC20Votes.Contract.NumCheckpoints(&_ERC20Votes.CallOpts, account)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address account) view returns(uint32)
func (_ERC20Votes *ERC20VotesCallerSession) NumCheckpoints(account common.Address) (uint32, error) {
	return _ERC20Votes.Contract.NumCheckpoints(&_ERC20Votes.CallOpts, account)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Votes *ERC20VotesCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Votes.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Votes *ERC20VotesSession) Symbol() (string, error) {
	return _ERC20Votes.Contract.Symbol(&_ERC20Votes.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Votes *ERC20VotesCallerSession) Symbol() (string, error) {
	return _ERC20Votes.Contract.Symbol(&_ERC20Votes.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Votes *ERC20VotesCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Votes.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Votes *ERC20VotesSession) TotalSupply() (*big.Int, error) {
	return _ERC20Votes.Contract.TotalSupply(&_ERC20Votes.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Votes *ERC20VotesCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Votes.Contract.TotalSupply(&_ERC20Votes.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Votes *ERC20VotesTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Votes.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Votes *ERC20VotesSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Votes.Contract.Approve(&_ERC20Votes.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Votes *ERC20VotesTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Votes.Contract.Approve(&_ERC20Votes.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Votes *ERC20VotesTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Votes.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Votes *ERC20VotesSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Votes.Contract.DecreaseAllowance(&_ERC20Votes.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Votes *ERC20VotesTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Votes.Contract.DecreaseAllowance(&_ERC20Votes.TransactOpts, spender, subtractedValue)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_ERC20Votes *ERC20VotesTransactor) Delegate(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _ERC20Votes.contract.Transact(opts, "delegate", delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_ERC20Votes *ERC20VotesSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _ERC20Votes.Contract.Delegate(&_ERC20Votes.TransactOpts, delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_ERC20Votes *ERC20VotesTransactorSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _ERC20Votes.Contract.Delegate(&_ERC20Votes.TransactOpts, delegatee)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_ERC20Votes *ERC20VotesTransactor) DelegateBySig(opts *bind.TransactOpts, delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ERC20Votes.contract.Transact(opts, "delegateBySig", delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_ERC20Votes *ERC20VotesSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ERC20Votes.Contract.DelegateBySig(&_ERC20Votes.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_ERC20Votes *ERC20VotesTransactorSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ERC20Votes.Contract.DelegateBySig(&_ERC20Votes.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Votes *ERC20VotesTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Votes.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Votes *ERC20VotesSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Votes.Contract.IncreaseAllowance(&_ERC20Votes.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Votes *ERC20VotesTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Votes.Contract.IncreaseAllowance(&_ERC20Votes.TransactOpts, spender, addedValue)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ERC20Votes *ERC20VotesTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ERC20Votes.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ERC20Votes *ERC20VotesSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ERC20Votes.Contract.Permit(&_ERC20Votes.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ERC20Votes *ERC20VotesTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ERC20Votes.Contract.Permit(&_ERC20Votes.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Votes *ERC20VotesTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Votes.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Votes *ERC20VotesSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Votes.Contract.Transfer(&_ERC20Votes.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Votes *ERC20VotesTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Votes.Contract.Transfer(&_ERC20Votes.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Votes *ERC20VotesTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Votes.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Votes *ERC20VotesSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Votes.Contract.TransferFrom(&_ERC20Votes.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Votes *ERC20VotesTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Votes.Contract.TransferFrom(&_ERC20Votes.TransactOpts, sender, recipient, amount)
}

// ERC20VotesApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Votes contract.
type ERC20VotesApprovalIterator struct {
	Event *ERC20VotesApproval // Event containing the contract specifics and raw log

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
func (it *ERC20VotesApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20VotesApproval)
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
		it.Event = new(ERC20VotesApproval)
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
func (it *ERC20VotesApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20VotesApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20VotesApproval represents a Approval event raised by the ERC20Votes contract.
type ERC20VotesApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Votes *ERC20VotesFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20VotesApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Votes.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20VotesApprovalIterator{contract: _ERC20Votes.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Votes *ERC20VotesFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20VotesApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Votes.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20VotesApproval)
				if err := _ERC20Votes.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Votes *ERC20VotesFilterer) ParseApproval(log types.Log) (*ERC20VotesApproval, error) {
	event := new(ERC20VotesApproval)
	if err := _ERC20Votes.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20VotesDelegateChangedIterator is returned from FilterDelegateChanged and is used to iterate over the raw logs and unpacked data for DelegateChanged events raised by the ERC20Votes contract.
type ERC20VotesDelegateChangedIterator struct {
	Event *ERC20VotesDelegateChanged // Event containing the contract specifics and raw log

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
func (it *ERC20VotesDelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20VotesDelegateChanged)
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
		it.Event = new(ERC20VotesDelegateChanged)
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
func (it *ERC20VotesDelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20VotesDelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20VotesDelegateChanged represents a DelegateChanged event raised by the ERC20Votes contract.
type ERC20VotesDelegateChanged struct {
	Delegator    common.Address
	FromDelegate common.Address
	ToDelegate   common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegateChanged is a free log retrieval operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_ERC20Votes *ERC20VotesFilterer) FilterDelegateChanged(opts *bind.FilterOpts, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (*ERC20VotesDelegateChangedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _ERC20Votes.contract.FilterLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return &ERC20VotesDelegateChangedIterator{contract: _ERC20Votes.contract, event: "DelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateChanged is a free log subscription operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_ERC20Votes *ERC20VotesFilterer) WatchDelegateChanged(opts *bind.WatchOpts, sink chan<- *ERC20VotesDelegateChanged, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _ERC20Votes.contract.WatchLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20VotesDelegateChanged)
				if err := _ERC20Votes.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
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

// ParseDelegateChanged is a log parse operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_ERC20Votes *ERC20VotesFilterer) ParseDelegateChanged(log types.Log) (*ERC20VotesDelegateChanged, error) {
	event := new(ERC20VotesDelegateChanged)
	if err := _ERC20Votes.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20VotesDelegateVotesChangedIterator is returned from FilterDelegateVotesChanged and is used to iterate over the raw logs and unpacked data for DelegateVotesChanged events raised by the ERC20Votes contract.
type ERC20VotesDelegateVotesChangedIterator struct {
	Event *ERC20VotesDelegateVotesChanged // Event containing the contract specifics and raw log

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
func (it *ERC20VotesDelegateVotesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20VotesDelegateVotesChanged)
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
		it.Event = new(ERC20VotesDelegateVotesChanged)
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
func (it *ERC20VotesDelegateVotesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20VotesDelegateVotesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20VotesDelegateVotesChanged represents a DelegateVotesChanged event raised by the ERC20Votes contract.
type ERC20VotesDelegateVotesChanged struct {
	Delegate        common.Address
	PreviousBalance *big.Int
	NewBalance      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelegateVotesChanged is a free log retrieval operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_ERC20Votes *ERC20VotesFilterer) FilterDelegateVotesChanged(opts *bind.FilterOpts, delegate []common.Address) (*ERC20VotesDelegateVotesChangedIterator, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _ERC20Votes.contract.FilterLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return &ERC20VotesDelegateVotesChangedIterator{contract: _ERC20Votes.contract, event: "DelegateVotesChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateVotesChanged is a free log subscription operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_ERC20Votes *ERC20VotesFilterer) WatchDelegateVotesChanged(opts *bind.WatchOpts, sink chan<- *ERC20VotesDelegateVotesChanged, delegate []common.Address) (event.Subscription, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _ERC20Votes.contract.WatchLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20VotesDelegateVotesChanged)
				if err := _ERC20Votes.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
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

// ParseDelegateVotesChanged is a log parse operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_ERC20Votes *ERC20VotesFilterer) ParseDelegateVotesChanged(log types.Log) (*ERC20VotesDelegateVotesChanged, error) {
	event := new(ERC20VotesDelegateVotesChanged)
	if err := _ERC20Votes.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20VotesTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Votes contract.
type ERC20VotesTransferIterator struct {
	Event *ERC20VotesTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20VotesTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20VotesTransfer)
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
		it.Event = new(ERC20VotesTransfer)
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
func (it *ERC20VotesTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20VotesTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20VotesTransfer represents a Transfer event raised by the ERC20Votes contract.
type ERC20VotesTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Votes *ERC20VotesFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20VotesTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Votes.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20VotesTransferIterator{contract: _ERC20Votes.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Votes *ERC20VotesFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20VotesTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Votes.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20VotesTransfer)
				if err := _ERC20Votes.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Votes *ERC20VotesFilterer) ParseTransfer(log types.Log) (*ERC20VotesTransfer, error) {
	event := new(ERC20VotesTransfer)
	if err := _ERC20Votes.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721MetaData contains all meta data concerning the ERC721 contract.
var ERC721MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162002946380380620029468339818101604052810190620000379190620001f6565b8160009081620000489190620004c6565b5080600190816200005a9190620004c6565b505050620005ad565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620000cc8262000081565b810181811067ffffffffffffffff82111715620000ee57620000ed62000092565b5b80604052505050565b60006200010362000063565b9050620001118282620000c1565b919050565b600067ffffffffffffffff82111562000134576200013362000092565b5b6200013f8262000081565b9050602081019050919050565b60005b838110156200016c5780820151818401526020810190506200014f565b60008484015250505050565b60006200018f620001898462000116565b620000f7565b905082815260208101848484011115620001ae57620001ad6200007c565b5b620001bb8482856200014c565b509392505050565b600082601f830112620001db57620001da62000077565b5b8151620001ed84826020860162000178565b91505092915050565b6000806040838503121562000210576200020f6200006d565b5b600083015167ffffffffffffffff81111562000231576200023062000072565b5b6200023f85828601620001c3565b925050602083015167ffffffffffffffff81111562000263576200026262000072565b5b6200027185828601620001c3565b9150509250929050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620002ce57607f821691505b602082108103620002e457620002e362000286565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026200034e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200030f565b6200035a86836200030f565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620003a7620003a16200039b8462000372565b6200037c565b62000372565b9050919050565b6000819050919050565b620003c38362000386565b620003db620003d282620003ae565b8484546200031c565b825550505050565b600090565b620003f2620003e3565b620003ff818484620003b8565b505050565b5b8181101562000427576200041b600082620003e8565b60018101905062000405565b5050565b601f82111562000476576200044081620002ea565b6200044b84620002ff565b810160208510156200045b578190505b620004736200046a85620002ff565b83018262000404565b50505b505050565b600082821c905092915050565b60006200049b600019846008026200047b565b1980831691505092915050565b6000620004b6838362000488565b9150826002028217905092915050565b620004d1826200027b565b67ffffffffffffffff811115620004ed57620004ec62000092565b5b620004f98254620002b5565b620005068282856200042b565b600060209050601f8311600181146200053e576000841562000529578287015190505b620005358582620004a8565b865550620005a5565b601f1984166200054e86620002ea565b60005b82811015620005785784890151825560018201915060208501945060208101905062000551565b8683101562000598578489015162000594601f89168262000488565b8355505b6001600288020188555050505b505050505050565b61238980620005bd6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80636352211e1161008c578063a22cb46511610066578063a22cb46514610224578063b88d4fde14610240578063c87b56dd1461025c578063e985e9c51461028c576100cf565b80636352211e146101a657806370a08231146101d657806395d89b4114610206576100cf565b806301ffc9a7146100d457806306fdde0314610104578063081812fc14610122578063095ea7b31461015257806323b872dd1461016e57806342842e0e1461018a575b600080fd5b6100ee60048036038101906100e99190611411565b6102bc565b6040516100fb9190611459565b60405180910390f35b61010c61039e565b6040516101199190611504565b60405180910390f35b61013c6004803603810190610137919061155c565b610430565b60405161014991906115ca565b60405180910390f35b61016c60048036038101906101679190611611565b6104b5565b005b61018860048036038101906101839190611651565b6105cc565b005b6101a4600480360381019061019f9190611651565b61062c565b005b6101c060048036038101906101bb919061155c565b61064c565b6040516101cd91906115ca565b60405180910390f35b6101f060048036038101906101eb91906116a4565b6106fd565b6040516101fd91906116e0565b60405180910390f35b61020e6107b4565b60405161021b9190611504565b60405180910390f35b61023e60048036038101906102399190611727565b610846565b005b61025a6004803603810190610255919061189c565b6109c6565b005b6102766004803603810190610271919061155c565b610a28565b6040516102839190611504565b60405180910390f35b6102a660048036038101906102a1919061191f565b610acf565b6040516102b39190611459565b60405180910390f35b60007f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061038757507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b80610397575061039682610b63565b5b9050919050565b6060600080546103ad9061198e565b80601f01602080910402602001604051908101604052809291908181526020018280546103d99061198e565b80156104265780601f106103fb57610100808354040283529160200191610426565b820191906000526020600020905b81548152906001019060200180831161040957829003601f168201915b5050505050905090565b600061043b82610bcd565b61047a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047190611a31565b60405180910390fd5b6004600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b60006104c08261064c565b90508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610530576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161052790611ac3565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1661054f610c39565b73ffffffffffffffffffffffffffffffffffffffff16148061057e575061057d81610578610c39565b610acf565b5b6105bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105b490611b55565b60405180910390fd5b6105c78383610c41565b505050565b6105dd6105d7610c39565b82610cfa565b61061c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161061390611be7565b60405180910390fd5b610627838383610dd8565b505050565b610647838383604051806020016040528060008152506109c6565b505050565b6000806002600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036106f4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106eb90611c79565b60405180910390fd5b80915050919050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361076d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076490611d0b565b60405180910390fd5b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6060600180546107c39061198e565b80601f01602080910402602001604051908101604052809291908181526020018280546107ef9061198e565b801561083c5780601f106108115761010080835404028352916020019161083c565b820191906000526020600020905b81548152906001019060200180831161081f57829003601f168201915b5050505050905090565b61084e610c39565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036108bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108b290611d77565b60405180910390fd5b80600560006108c8610c39565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff16610975610c39565b73ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31836040516109ba9190611459565b60405180910390a35050565b6109d76109d1610c39565b83610cfa565b610a16576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a0d90611be7565b60405180910390fd5b610a2284848484611033565b50505050565b6060610a3382610bcd565b610a72576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a6990611e09565b60405180910390fd5b6000610a7c61108f565b90506000815111610a9c5760405180602001604052806000815250610ac7565b80610aa6846110a6565b604051602001610ab7929190611e65565b6040516020818303038152906040525b915050919050565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b60008073ffffffffffffffffffffffffffffffffffffffff166002600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614159050919050565b600033905090565b816004600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff16610cb48361064c565b73ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6000610d0582610bcd565b610d44576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d3b90611efb565b60405180910390fd5b6000610d4f8361064c565b90508073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480610dbe57508373ffffffffffffffffffffffffffffffffffffffff16610da684610430565b73ffffffffffffffffffffffffffffffffffffffff16145b80610dcf5750610dce8185610acf565b5b91505092915050565b8273ffffffffffffffffffffffffffffffffffffffff16610df88261064c565b73ffffffffffffffffffffffffffffffffffffffff1614610e4e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e4590611f8d565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610ebd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610eb49061201f565b60405180910390fd5b610ec8838383611206565b610ed3600082610c41565b6001600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610f23919061206e565b925050819055506001600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610f7a91906120a2565b92505081905550816002600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b61103e848484610dd8565b61104a8484848461120b565b611089576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161108090612148565b60405180910390fd5b50505050565b606060405180602001604052806000815250905090565b6060600082036110ed576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050611201565b600082905060005b6000821461111f57808061110890612168565b915050600a8261111891906121df565b91506110f5565b60008167ffffffffffffffff81111561113b5761113a611771565b5b6040519080825280601f01601f19166020018201604052801561116d5781602001600182028036833780820191505090505b5090505b600085146111fa57600182611186919061206e565b9150600a856111959190612210565b60306111a191906120a2565b60f81b8183815181106111b7576111b6612241565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a856111f391906121df565b9450611171565b8093505050505b919050565b505050565b600061122c8473ffffffffffffffffffffffffffffffffffffffff16611392565b15611385578373ffffffffffffffffffffffffffffffffffffffff1663150b7a02611255610c39565b8786866040518563ffffffff1660e01b815260040161127794939291906122c5565b6020604051808303816000875af19250505080156112b357506040513d601f19601f820116820180604052508101906112b09190612326565b60015b611335573d80600081146112e3576040519150601f19603f3d011682016040523d82523d6000602084013e6112e8565b606091505b50600081510361132d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161132490612148565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161491505061138a565b600190505b949350505050565b600080823b905060008111915050919050565b6000604051905090565b600080fd5b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6113ee816113b9565b81146113f957600080fd5b50565b60008135905061140b816113e5565b92915050565b600060208284031215611427576114266113af565b5b6000611435848285016113fc565b91505092915050565b60008115159050919050565b6114538161143e565b82525050565b600060208201905061146e600083018461144a565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156114ae578082015181840152602081019050611493565b60008484015250505050565b6000601f19601f8301169050919050565b60006114d682611474565b6114e0818561147f565b93506114f0818560208601611490565b6114f9816114ba565b840191505092915050565b6000602082019050818103600083015261151e81846114cb565b905092915050565b6000819050919050565b61153981611526565b811461154457600080fd5b50565b60008135905061155681611530565b92915050565b600060208284031215611572576115716113af565b5b600061158084828501611547565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006115b482611589565b9050919050565b6115c4816115a9565b82525050565b60006020820190506115df60008301846115bb565b92915050565b6115ee816115a9565b81146115f957600080fd5b50565b60008135905061160b816115e5565b92915050565b60008060408385031215611628576116276113af565b5b6000611636858286016115fc565b925050602061164785828601611547565b9150509250929050565b60008060006060848603121561166a576116696113af565b5b6000611678868287016115fc565b9350506020611689868287016115fc565b925050604061169a86828701611547565b9150509250925092565b6000602082840312156116ba576116b96113af565b5b60006116c8848285016115fc565b91505092915050565b6116da81611526565b82525050565b60006020820190506116f560008301846116d1565b92915050565b6117048161143e565b811461170f57600080fd5b50565b600081359050611721816116fb565b92915050565b6000806040838503121561173e5761173d6113af565b5b600061174c858286016115fc565b925050602061175d85828601611712565b9150509250929050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6117a9826114ba565b810181811067ffffffffffffffff821117156117c8576117c7611771565b5b80604052505050565b60006117db6113a5565b90506117e782826117a0565b919050565b600067ffffffffffffffff82111561180757611806611771565b5b611810826114ba565b9050602081019050919050565b82818337600083830152505050565b600061183f61183a846117ec565b6117d1565b90508281526020810184848401111561185b5761185a61176c565b5b61186684828561181d565b509392505050565b600082601f83011261188357611882611767565b5b813561189384826020860161182c565b91505092915050565b600080600080608085870312156118b6576118b56113af565b5b60006118c4878288016115fc565b94505060206118d5878288016115fc565b93505060406118e687828801611547565b925050606085013567ffffffffffffffff811115611907576119066113b4565b5b6119138782880161186e565b91505092959194509250565b60008060408385031215611936576119356113af565b5b6000611944858286016115fc565b9250506020611955858286016115fc565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806119a657607f821691505b6020821081036119b9576119b861195f565b5b50919050565b7f4552433732313a20617070726f76656420717565727920666f72206e6f6e657860008201527f697374656e7420746f6b656e0000000000000000000000000000000000000000602082015250565b6000611a1b602c8361147f565b9150611a26826119bf565b604082019050919050565b60006020820190508181036000830152611a4a81611a0e565b9050919050565b7f4552433732313a20617070726f76616c20746f2063757272656e74206f776e6560008201527f7200000000000000000000000000000000000000000000000000000000000000602082015250565b6000611aad60218361147f565b9150611ab882611a51565b604082019050919050565b60006020820190508181036000830152611adc81611aa0565b9050919050565b7f4552433732313a20617070726f76652063616c6c6572206973206e6f74206f7760008201527f6e6572206e6f7220617070726f76656420666f7220616c6c0000000000000000602082015250565b6000611b3f60388361147f565b9150611b4a82611ae3565b604082019050919050565b60006020820190508181036000830152611b6e81611b32565b9050919050565b7f4552433732313a207472616e736665722063616c6c6572206973206e6f74206f60008201527f776e6572206e6f7220617070726f766564000000000000000000000000000000602082015250565b6000611bd160318361147f565b9150611bdc82611b75565b604082019050919050565b60006020820190508181036000830152611c0081611bc4565b9050919050565b7f4552433732313a206f776e657220717565727920666f72206e6f6e657869737460008201527f656e7420746f6b656e0000000000000000000000000000000000000000000000602082015250565b6000611c6360298361147f565b9150611c6e82611c07565b604082019050919050565b60006020820190508181036000830152611c9281611c56565b9050919050565b7f4552433732313a2062616c616e636520717565727920666f7220746865207a6560008201527f726f206164647265737300000000000000000000000000000000000000000000602082015250565b6000611cf5602a8361147f565b9150611d0082611c99565b604082019050919050565b60006020820190508181036000830152611d2481611ce8565b9050919050565b7f4552433732313a20617070726f766520746f2063616c6c657200000000000000600082015250565b6000611d6160198361147f565b9150611d6c82611d2b565b602082019050919050565b60006020820190508181036000830152611d9081611d54565b9050919050565b7f4552433732314d657461646174613a2055524920717565727920666f72206e6f60008201527f6e6578697374656e7420746f6b656e0000000000000000000000000000000000602082015250565b6000611df3602f8361147f565b9150611dfe82611d97565b604082019050919050565b60006020820190508181036000830152611e2281611de6565b9050919050565b600081905092915050565b6000611e3f82611474565b611e498185611e29565b9350611e59818560208601611490565b80840191505092915050565b6000611e718285611e34565b9150611e7d8284611e34565b91508190509392505050565b7f4552433732313a206f70657261746f7220717565727920666f72206e6f6e657860008201527f697374656e7420746f6b656e0000000000000000000000000000000000000000602082015250565b6000611ee5602c8361147f565b9150611ef082611e89565b604082019050919050565b60006020820190508181036000830152611f1481611ed8565b9050919050565b7f4552433732313a207472616e73666572206f6620746f6b656e2074686174206960008201527f73206e6f74206f776e0000000000000000000000000000000000000000000000602082015250565b6000611f7760298361147f565b9150611f8282611f1b565b604082019050919050565b60006020820190508181036000830152611fa681611f6a565b9050919050565b7f4552433732313a207472616e7366657220746f20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b600061200960248361147f565b915061201482611fad565b604082019050919050565b6000602082019050818103600083015261203881611ffc565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061207982611526565b915061208483611526565b925082820390508181111561209c5761209b61203f565b5b92915050565b60006120ad82611526565b91506120b883611526565b92508282019050808211156120d0576120cf61203f565b5b92915050565b7f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560008201527f63656976657220696d706c656d656e7465720000000000000000000000000000602082015250565b600061213260328361147f565b915061213d826120d6565b604082019050919050565b6000602082019050818103600083015261216181612125565b9050919050565b600061217382611526565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036121a5576121a461203f565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006121ea82611526565b91506121f583611526565b925082612205576122046121b0565b5b828204905092915050565b600061221b82611526565b915061222683611526565b925082612236576122356121b0565b5b828206905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600081519050919050565b600082825260208201905092915050565b600061229782612270565b6122a1818561227b565b93506122b1818560208601611490565b6122ba816114ba565b840191505092915050565b60006080820190506122da60008301876115bb565b6122e760208301866115bb565b6122f460408301856116d1565b8181036060830152612306818461228c565b905095945050505050565b600081519050612320816113e5565b92915050565b60006020828403121561233c5761233b6113af565b5b600061234a84828501612311565b9150509291505056fea2646970667358221220fb1be3510c119434556b065ac9555a3e4ddcae4e3b6ee61680e47c45ff5b744e64736f6c63430008130033",
}

// ERC721ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC721MetaData.ABI instead.
var ERC721ABI = ERC721MetaData.ABI

// ERC721Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC721MetaData.Bin instead.
var ERC721Bin = ERC721MetaData.Bin

// DeployERC721 deploys a new Ethereum contract, binding an instance of ERC721 to it.
func DeployERC721(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string) (common.Address, *types.Transaction, *ERC721, error) {
	parsed, err := ERC721MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC721Bin), backend, name_, symbol_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721{ERC721Caller: ERC721Caller{contract: contract}, ERC721Transactor: ERC721Transactor{contract: contract}, ERC721Filterer: ERC721Filterer{contract: contract}}, nil
}

// ERC721 is an auto generated Go binding around an Ethereum contract.
type ERC721 struct {
	ERC721Caller     // Read-only binding to the contract
	ERC721Transactor // Write-only binding to the contract
	ERC721Filterer   // Log filterer for contract events
}

// ERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721Session struct {
	Contract     *ERC721           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721CallerSession struct {
	Contract *ERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721TransactorSession struct {
	Contract     *ERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721Raw struct {
	Contract *ERC721 // Generic contract binding to access the raw methods on
}

// ERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721CallerRaw struct {
	Contract *ERC721Caller // Generic read-only contract binding to access the raw methods on
}

// ERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721TransactorRaw struct {
	Contract *ERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721 creates a new instance of ERC721, bound to a specific deployed contract.
func NewERC721(address common.Address, backend bind.ContractBackend) (*ERC721, error) {
	contract, err := bindERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721{ERC721Caller: ERC721Caller{contract: contract}, ERC721Transactor: ERC721Transactor{contract: contract}, ERC721Filterer: ERC721Filterer{contract: contract}}, nil
}

// NewERC721Caller creates a new read-only instance of ERC721, bound to a specific deployed contract.
func NewERC721Caller(address common.Address, caller bind.ContractCaller) (*ERC721Caller, error) {
	contract, err := bindERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721Caller{contract: contract}, nil
}

// NewERC721Transactor creates a new write-only instance of ERC721, bound to a specific deployed contract.
func NewERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC721Transactor, error) {
	contract, err := bindERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721Transactor{contract: contract}, nil
}

// NewERC721Filterer creates a new log filterer instance of ERC721, bound to a specific deployed contract.
func NewERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC721Filterer, error) {
	contract, err := bindERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721Filterer{contract: contract}, nil
}

// bindERC721 binds a generic wrapper to an already deployed contract.
func bindERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721 *ERC721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721.Contract.ERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721 *ERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721.Contract.ERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721 *ERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721.Contract.ERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721 *ERC721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721 *ERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721 *ERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721 *ERC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721 *ERC721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721.Contract.BalanceOf(&_ERC721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721 *ERC721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721.Contract.BalanceOf(&_ERC721.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.GetApproved(&_ERC721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.GetApproved(&_ERC721.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721 *ERC721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721 *ERC721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721.Contract.IsApprovedForAll(&_ERC721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721 *ERC721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721.Contract.IsApprovedForAll(&_ERC721.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721 *ERC721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721 *ERC721Session) Name() (string, error) {
	return _ERC721.Contract.Name(&_ERC721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721 *ERC721CallerSession) Name() (string, error) {
	return _ERC721.Contract.Name(&_ERC721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.OwnerOf(&_ERC721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.OwnerOf(&_ERC721.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721 *ERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721 *ERC721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721.Contract.SupportsInterface(&_ERC721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721 *ERC721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721.Contract.SupportsInterface(&_ERC721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721 *ERC721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721 *ERC721Session) Symbol() (string, error) {
	return _ERC721.Contract.Symbol(&_ERC721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721 *ERC721CallerSession) Symbol() (string, error) {
	return _ERC721.Contract.Symbol(&_ERC721.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721 *ERC721Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721 *ERC721Session) TokenURI(tokenId *big.Int) (string, error) {
	return _ERC721.Contract.TokenURI(&_ERC721.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721 *ERC721CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ERC721.Contract.TokenURI(&_ERC721.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Approve(&_ERC721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721 *ERC721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Approve(&_ERC721.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721 *ERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721 *ERC721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom0(&_ERC721.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721 *ERC721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom0(&_ERC721.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721 *ERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721 *ERC721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721.Contract.SetApprovalForAll(&_ERC721.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721 *ERC721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721.Contract.SetApprovalForAll(&_ERC721.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.TransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.TransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

// ERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721 contract.
type ERC721ApprovalIterator struct {
	Event *ERC721Approval // Event containing the contract specifics and raw log

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
func (it *ERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721Approval)
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
		it.Event = new(ERC721Approval)
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
func (it *ERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721Approval represents a Approval event raised by the ERC721 contract.
type ERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ERC721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721ApprovalIterator{contract: _ERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721Approval)
				if err := _ERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) ParseApproval(log types.Log) (*ERC721Approval, error) {
	event := new(ERC721Approval)
	if err := _ERC721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721 contract.
type ERC721ApprovalForAllIterator struct {
	Event *ERC721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721ApprovalForAll)
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
		it.Event = new(ERC721ApprovalForAll)
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
func (it *ERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721ApprovalForAll represents a ApprovalForAll event raised by the ERC721 contract.
type ERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721 *ERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ERC721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721ApprovalForAllIterator{contract: _ERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721 *ERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721ApprovalForAll)
				if err := _ERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_ERC721 *ERC721Filterer) ParseApprovalForAll(log types.Log) (*ERC721ApprovalForAll, error) {
	event := new(ERC721ApprovalForAll)
	if err := _ERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721 contract.
type ERC721TransferIterator struct {
	Event *ERC721Transfer // Event containing the contract specifics and raw log

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
func (it *ERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721Transfer)
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
		it.Event = new(ERC721Transfer)
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
func (it *ERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721Transfer represents a Transfer event raised by the ERC721 contract.
type ERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ERC721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721TransferIterator{contract: _ERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721Transfer)
				if err := _ERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) ParseTransfer(log types.Log) (*ERC721Transfer, error) {
	event := new(ERC721Transfer)
	if err := _ERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721EnumerableMetaData contains all meta data concerning the ERC721Enumerable contract.
var ERC721EnumerableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERC721EnumerableABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC721EnumerableMetaData.ABI instead.
var ERC721EnumerableABI = ERC721EnumerableMetaData.ABI

// ERC721Enumerable is an auto generated Go binding around an Ethereum contract.
type ERC721Enumerable struct {
	ERC721EnumerableCaller     // Read-only binding to the contract
	ERC721EnumerableTransactor // Write-only binding to the contract
	ERC721EnumerableFilterer   // Log filterer for contract events
}

// ERC721EnumerableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721EnumerableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721EnumerableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721EnumerableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721EnumerableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721EnumerableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721EnumerableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721EnumerableSession struct {
	Contract     *ERC721Enumerable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721EnumerableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721EnumerableCallerSession struct {
	Contract *ERC721EnumerableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ERC721EnumerableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721EnumerableTransactorSession struct {
	Contract     *ERC721EnumerableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ERC721EnumerableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721EnumerableRaw struct {
	Contract *ERC721Enumerable // Generic contract binding to access the raw methods on
}

// ERC721EnumerableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721EnumerableCallerRaw struct {
	Contract *ERC721EnumerableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721EnumerableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721EnumerableTransactorRaw struct {
	Contract *ERC721EnumerableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Enumerable creates a new instance of ERC721Enumerable, bound to a specific deployed contract.
func NewERC721Enumerable(address common.Address, backend bind.ContractBackend) (*ERC721Enumerable, error) {
	contract, err := bindERC721Enumerable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Enumerable{ERC721EnumerableCaller: ERC721EnumerableCaller{contract: contract}, ERC721EnumerableTransactor: ERC721EnumerableTransactor{contract: contract}, ERC721EnumerableFilterer: ERC721EnumerableFilterer{contract: contract}}, nil
}

// NewERC721EnumerableCaller creates a new read-only instance of ERC721Enumerable, bound to a specific deployed contract.
func NewERC721EnumerableCaller(address common.Address, caller bind.ContractCaller) (*ERC721EnumerableCaller, error) {
	contract, err := bindERC721Enumerable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableCaller{contract: contract}, nil
}

// NewERC721EnumerableTransactor creates a new write-only instance of ERC721Enumerable, bound to a specific deployed contract.
func NewERC721EnumerableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721EnumerableTransactor, error) {
	contract, err := bindERC721Enumerable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableTransactor{contract: contract}, nil
}

// NewERC721EnumerableFilterer creates a new log filterer instance of ERC721Enumerable, bound to a specific deployed contract.
func NewERC721EnumerableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721EnumerableFilterer, error) {
	contract, err := bindERC721Enumerable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableFilterer{contract: contract}, nil
}

// bindERC721Enumerable binds a generic wrapper to an already deployed contract.
func bindERC721Enumerable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721EnumerableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Enumerable *ERC721EnumerableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721Enumerable.Contract.ERC721EnumerableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Enumerable *ERC721EnumerableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.ERC721EnumerableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Enumerable *ERC721EnumerableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.ERC721EnumerableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Enumerable *ERC721EnumerableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721Enumerable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Enumerable *ERC721EnumerableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Enumerable *ERC721EnumerableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC721Enumerable.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721Enumerable.Contract.BalanceOf(&_ERC721Enumerable.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721Enumerable.Contract.BalanceOf(&_ERC721Enumerable.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721Enumerable *ERC721EnumerableCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC721Enumerable.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721Enumerable *ERC721EnumerableSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721Enumerable.Contract.GetApproved(&_ERC721Enumerable.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721Enumerable.Contract.GetApproved(&_ERC721Enumerable.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721Enumerable *ERC721EnumerableCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ERC721Enumerable.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721Enumerable *ERC721EnumerableSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721Enumerable.Contract.IsApprovedForAll(&_ERC721Enumerable.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721Enumerable.Contract.IsApprovedForAll(&_ERC721Enumerable.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721Enumerable *ERC721EnumerableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC721Enumerable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721Enumerable *ERC721EnumerableSession) Name() (string, error) {
	return _ERC721Enumerable.Contract.Name(&_ERC721Enumerable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) Name() (string, error) {
	return _ERC721Enumerable.Contract.Name(&_ERC721Enumerable.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721Enumerable *ERC721EnumerableCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC721Enumerable.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721Enumerable *ERC721EnumerableSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721Enumerable.Contract.OwnerOf(&_ERC721Enumerable.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721Enumerable.Contract.OwnerOf(&_ERC721Enumerable.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721Enumerable *ERC721EnumerableCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC721Enumerable.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721Enumerable *ERC721EnumerableSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721Enumerable.Contract.SupportsInterface(&_ERC721Enumerable.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721Enumerable.Contract.SupportsInterface(&_ERC721Enumerable.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721Enumerable *ERC721EnumerableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC721Enumerable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721Enumerable *ERC721EnumerableSession) Symbol() (string, error) {
	return _ERC721Enumerable.Contract.Symbol(&_ERC721Enumerable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) Symbol() (string, error) {
	return _ERC721Enumerable.Contract.Symbol(&_ERC721Enumerable.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC721Enumerable.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ERC721Enumerable.Contract.TokenByIndex(&_ERC721Enumerable.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ERC721Enumerable.Contract.TokenByIndex(&_ERC721Enumerable.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC721Enumerable.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ERC721Enumerable.Contract.TokenOfOwnerByIndex(&_ERC721Enumerable.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ERC721Enumerable.Contract.TokenOfOwnerByIndex(&_ERC721Enumerable.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721Enumerable *ERC721EnumerableCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ERC721Enumerable.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721Enumerable *ERC721EnumerableSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ERC721Enumerable.Contract.TokenURI(&_ERC721Enumerable.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ERC721Enumerable.Contract.TokenURI(&_ERC721Enumerable.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC721Enumerable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableSession) TotalSupply() (*big.Int, error) {
	return _ERC721Enumerable.Contract.TotalSupply(&_ERC721Enumerable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC721Enumerable.Contract.TotalSupply(&_ERC721Enumerable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.Approve(&_ERC721Enumerable.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.Approve(&_ERC721Enumerable.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.SafeTransferFrom(&_ERC721Enumerable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.SafeTransferFrom(&_ERC721Enumerable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Enumerable.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721Enumerable *ERC721EnumerableSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.SafeTransferFrom0(&_ERC721Enumerable.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.SafeTransferFrom0(&_ERC721Enumerable.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721Enumerable.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721Enumerable *ERC721EnumerableSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.SetApprovalForAll(&_ERC721Enumerable.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.SetApprovalForAll(&_ERC721Enumerable.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.TransferFrom(&_ERC721Enumerable.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.TransferFrom(&_ERC721Enumerable.TransactOpts, from, to, tokenId)
}

// ERC721EnumerableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721Enumerable contract.
type ERC721EnumerableApprovalIterator struct {
	Event *ERC721EnumerableApproval // Event containing the contract specifics and raw log

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
func (it *ERC721EnumerableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721EnumerableApproval)
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
		it.Event = new(ERC721EnumerableApproval)
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
func (it *ERC721EnumerableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721EnumerableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721EnumerableApproval represents a Approval event raised by the ERC721Enumerable contract.
type ERC721EnumerableApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721Enumerable *ERC721EnumerableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ERC721EnumerableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableApprovalIterator{contract: _ERC721Enumerable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721Enumerable *ERC721EnumerableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721EnumerableApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721EnumerableApproval)
				if err := _ERC721Enumerable.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721Enumerable *ERC721EnumerableFilterer) ParseApproval(log types.Log) (*ERC721EnumerableApproval, error) {
	event := new(ERC721EnumerableApproval)
	if err := _ERC721Enumerable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721EnumerableApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721Enumerable contract.
type ERC721EnumerableApprovalForAllIterator struct {
	Event *ERC721EnumerableApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721EnumerableApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721EnumerableApprovalForAll)
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
		it.Event = new(ERC721EnumerableApprovalForAll)
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
func (it *ERC721EnumerableApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721EnumerableApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721EnumerableApprovalForAll represents a ApprovalForAll event raised by the ERC721Enumerable contract.
type ERC721EnumerableApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721Enumerable *ERC721EnumerableFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ERC721EnumerableApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableApprovalForAllIterator{contract: _ERC721Enumerable.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721Enumerable *ERC721EnumerableFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721EnumerableApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721EnumerableApprovalForAll)
				if err := _ERC721Enumerable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_ERC721Enumerable *ERC721EnumerableFilterer) ParseApprovalForAll(log types.Log) (*ERC721EnumerableApprovalForAll, error) {
	event := new(ERC721EnumerableApprovalForAll)
	if err := _ERC721Enumerable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721EnumerableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721Enumerable contract.
type ERC721EnumerableTransferIterator struct {
	Event *ERC721EnumerableTransfer // Event containing the contract specifics and raw log

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
func (it *ERC721EnumerableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721EnumerableTransfer)
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
		it.Event = new(ERC721EnumerableTransfer)
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
func (it *ERC721EnumerableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721EnumerableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721EnumerableTransfer represents a Transfer event raised by the ERC721Enumerable contract.
type ERC721EnumerableTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721Enumerable *ERC721EnumerableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ERC721EnumerableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableTransferIterator{contract: _ERC721Enumerable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721Enumerable *ERC721EnumerableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721EnumerableTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721EnumerableTransfer)
				if err := _ERC721Enumerable.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721Enumerable *ERC721EnumerableFilterer) ParseTransfer(log types.Log) (*ERC721EnumerableTransfer, error) {
	event := new(ERC721EnumerableTransfer)
	if err := _ERC721Enumerable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC165MetaData contains all meta data concerning the IERC165 contract.
var IERC165MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IERC165ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC165MetaData.ABI instead.
var IERC165ABI = IERC165MetaData.ABI

// IERC165 is an auto generated Go binding around an Ethereum contract.
type IERC165 struct {
	IERC165Caller     // Read-only binding to the contract
	IERC165Transactor // Write-only binding to the contract
	IERC165Filterer   // Log filterer for contract events
}

// IERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC165Session struct {
	Contract     *IERC165          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC165CallerSession struct {
	Contract *IERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC165TransactorSession struct {
	Contract     *IERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC165Raw struct {
	Contract *IERC165 // Generic contract binding to access the raw methods on
}

// IERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC165CallerRaw struct {
	Contract *IERC165Caller // Generic read-only contract binding to access the raw methods on
}

// IERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC165TransactorRaw struct {
	Contract *IERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC165 creates a new instance of IERC165, bound to a specific deployed contract.
func NewIERC165(address common.Address, backend bind.ContractBackend) (*IERC165, error) {
	contract, err := bindIERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC165{IERC165Caller: IERC165Caller{contract: contract}, IERC165Transactor: IERC165Transactor{contract: contract}, IERC165Filterer: IERC165Filterer{contract: contract}}, nil
}

// NewIERC165Caller creates a new read-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Caller(address common.Address, caller bind.ContractCaller) (*IERC165Caller, error) {
	contract, err := bindIERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Caller{contract: contract}, nil
}

// NewIERC165Transactor creates a new write-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC165Transactor, error) {
	contract, err := bindIERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Transactor{contract: contract}, nil
}

// NewIERC165Filterer creates a new log filterer instance of IERC165, bound to a specific deployed contract.
func NewIERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC165Filterer, error) {
	contract, err := bindIERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC165Filterer{contract: contract}, nil
}

// bindIERC165 binds a generic wrapper to an already deployed contract.
func bindIERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.IERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20MetadataMetaData contains all meta data concerning the IERC20Metadata contract.
var IERC20MetadataMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20MetadataABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetadataMetaData.ABI instead.
var IERC20MetadataABI = IERC20MetadataMetaData.ABI

// IERC20Metadata is an auto generated Go binding around an Ethereum contract.
type IERC20Metadata struct {
	IERC20MetadataCaller     // Read-only binding to the contract
	IERC20MetadataTransactor // Write-only binding to the contract
	IERC20MetadataFilterer   // Log filterer for contract events
}

// IERC20MetadataCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20MetadataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MetadataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20MetadataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MetadataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20MetadataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MetadataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20MetadataSession struct {
	Contract     *IERC20Metadata   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20MetadataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20MetadataCallerSession struct {
	Contract *IERC20MetadataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IERC20MetadataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20MetadataTransactorSession struct {
	Contract     *IERC20MetadataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IERC20MetadataRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20MetadataRaw struct {
	Contract *IERC20Metadata // Generic contract binding to access the raw methods on
}

// IERC20MetadataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20MetadataCallerRaw struct {
	Contract *IERC20MetadataCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20MetadataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20MetadataTransactorRaw struct {
	Contract *IERC20MetadataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Metadata creates a new instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20Metadata(address common.Address, backend bind.ContractBackend) (*IERC20Metadata, error) {
	contract, err := bindIERC20Metadata(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Metadata{IERC20MetadataCaller: IERC20MetadataCaller{contract: contract}, IERC20MetadataTransactor: IERC20MetadataTransactor{contract: contract}, IERC20MetadataFilterer: IERC20MetadataFilterer{contract: contract}}, nil
}

// NewIERC20MetadataCaller creates a new read-only instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20MetadataCaller(address common.Address, caller bind.ContractCaller) (*IERC20MetadataCaller, error) {
	contract, err := bindIERC20Metadata(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataCaller{contract: contract}, nil
}

// NewIERC20MetadataTransactor creates a new write-only instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20MetadataTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20MetadataTransactor, error) {
	contract, err := bindIERC20Metadata(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataTransactor{contract: contract}, nil
}

// NewIERC20MetadataFilterer creates a new log filterer instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20MetadataFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20MetadataFilterer, error) {
	contract, err := bindIERC20Metadata(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataFilterer{contract: contract}, nil
}

// bindIERC20Metadata binds a generic wrapper to an already deployed contract.
func bindIERC20Metadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20MetadataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Metadata *IERC20MetadataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Metadata.Contract.IERC20MetadataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Metadata *IERC20MetadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.IERC20MetadataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Metadata *IERC20MetadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.IERC20MetadataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Metadata *IERC20MetadataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Metadata.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Metadata *IERC20MetadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Metadata *IERC20MetadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.Allowance(&_IERC20Metadata.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.Allowance(&_IERC20Metadata.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.BalanceOf(&_IERC20Metadata.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.BalanceOf(&_IERC20Metadata.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Metadata *IERC20MetadataCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Metadata *IERC20MetadataSession) Decimals() (uint8, error) {
	return _IERC20Metadata.Contract.Decimals(&_IERC20Metadata.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Metadata *IERC20MetadataCallerSession) Decimals() (uint8, error) {
	return _IERC20Metadata.Contract.Decimals(&_IERC20Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Metadata *IERC20MetadataCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Metadata *IERC20MetadataSession) Name() (string, error) {
	return _IERC20Metadata.Contract.Name(&_IERC20Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Metadata *IERC20MetadataCallerSession) Name() (string, error) {
	return _IERC20Metadata.Contract.Name(&_IERC20Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Metadata *IERC20MetadataCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Metadata *IERC20MetadataSession) Symbol() (string, error) {
	return _IERC20Metadata.Contract.Symbol(&_IERC20Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Metadata *IERC20MetadataCallerSession) Symbol() (string, error) {
	return _IERC20Metadata.Contract.Symbol(&_IERC20Metadata.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Metadata *IERC20MetadataSession) TotalSupply() (*big.Int, error) {
	return _IERC20Metadata.Contract.TotalSupply(&_IERC20Metadata.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20Metadata.Contract.TotalSupply(&_IERC20Metadata.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Approve(&_IERC20Metadata.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Approve(&_IERC20Metadata.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Transfer(&_IERC20Metadata.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Transfer(&_IERC20Metadata.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.TransferFrom(&_IERC20Metadata.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.TransferFrom(&_IERC20Metadata.TransactOpts, sender, recipient, amount)
}

// IERC20MetadataApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20Metadata contract.
type IERC20MetadataApprovalIterator struct {
	Event *IERC20MetadataApproval // Event containing the contract specifics and raw log

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
func (it *IERC20MetadataApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20MetadataApproval)
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
		it.Event = new(IERC20MetadataApproval)
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
func (it *IERC20MetadataApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20MetadataApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20MetadataApproval represents a Approval event raised by the IERC20Metadata contract.
type IERC20MetadataApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20MetadataApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Metadata.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataApprovalIterator{contract: _IERC20Metadata.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20MetadataApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Metadata.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20MetadataApproval)
				if err := _IERC20Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) ParseApproval(log types.Log) (*IERC20MetadataApproval, error) {
	event := new(IERC20MetadataApproval)
	if err := _IERC20Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20MetadataTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20Metadata contract.
type IERC20MetadataTransferIterator struct {
	Event *IERC20MetadataTransfer // Event containing the contract specifics and raw log

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
func (it *IERC20MetadataTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20MetadataTransfer)
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
		it.Event = new(IERC20MetadataTransfer)
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
func (it *IERC20MetadataTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20MetadataTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20MetadataTransfer represents a Transfer event raised by the IERC20Metadata contract.
type IERC20MetadataTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20MetadataTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Metadata.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataTransferIterator{contract: _IERC20Metadata.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20MetadataTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Metadata.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20MetadataTransfer)
				if err := _IERC20Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) ParseTransfer(log types.Log) (*IERC20MetadataTransfer, error) {
	event := new(IERC20MetadataTransfer)
	if err := _IERC20Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20PermitMetaData contains all meta data concerning the IERC20Permit contract.
var IERC20PermitMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20PermitABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20PermitMetaData.ABI instead.
var IERC20PermitABI = IERC20PermitMetaData.ABI

// IERC20Permit is an auto generated Go binding around an Ethereum contract.
type IERC20Permit struct {
	IERC20PermitCaller     // Read-only binding to the contract
	IERC20PermitTransactor // Write-only binding to the contract
	IERC20PermitFilterer   // Log filterer for contract events
}

// IERC20PermitCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20PermitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20PermitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20PermitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20PermitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20PermitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20PermitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20PermitSession struct {
	Contract     *IERC20Permit     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20PermitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20PermitCallerSession struct {
	Contract *IERC20PermitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IERC20PermitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20PermitTransactorSession struct {
	Contract     *IERC20PermitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IERC20PermitRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20PermitRaw struct {
	Contract *IERC20Permit // Generic contract binding to access the raw methods on
}

// IERC20PermitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20PermitCallerRaw struct {
	Contract *IERC20PermitCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20PermitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20PermitTransactorRaw struct {
	Contract *IERC20PermitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Permit creates a new instance of IERC20Permit, bound to a specific deployed contract.
func NewIERC20Permit(address common.Address, backend bind.ContractBackend) (*IERC20Permit, error) {
	contract, err := bindIERC20Permit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Permit{IERC20PermitCaller: IERC20PermitCaller{contract: contract}, IERC20PermitTransactor: IERC20PermitTransactor{contract: contract}, IERC20PermitFilterer: IERC20PermitFilterer{contract: contract}}, nil
}

// NewIERC20PermitCaller creates a new read-only instance of IERC20Permit, bound to a specific deployed contract.
func NewIERC20PermitCaller(address common.Address, caller bind.ContractCaller) (*IERC20PermitCaller, error) {
	contract, err := bindIERC20Permit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20PermitCaller{contract: contract}, nil
}

// NewIERC20PermitTransactor creates a new write-only instance of IERC20Permit, bound to a specific deployed contract.
func NewIERC20PermitTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20PermitTransactor, error) {
	contract, err := bindIERC20Permit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20PermitTransactor{contract: contract}, nil
}

// NewIERC20PermitFilterer creates a new log filterer instance of IERC20Permit, bound to a specific deployed contract.
func NewIERC20PermitFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20PermitFilterer, error) {
	contract, err := bindIERC20Permit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20PermitFilterer{contract: contract}, nil
}

// bindIERC20Permit binds a generic wrapper to an already deployed contract.
func bindIERC20Permit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20PermitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Permit *IERC20PermitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Permit.Contract.IERC20PermitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Permit *IERC20PermitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Permit.Contract.IERC20PermitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Permit *IERC20PermitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Permit.Contract.IERC20PermitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Permit *IERC20PermitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Permit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Permit *IERC20PermitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Permit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Permit *IERC20PermitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Permit.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC20Permit *IERC20PermitCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IERC20Permit.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC20Permit *IERC20PermitSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IERC20Permit.Contract.DOMAINSEPARATOR(&_IERC20Permit.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC20Permit *IERC20PermitCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IERC20Permit.Contract.DOMAINSEPARATOR(&_IERC20Permit.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IERC20Permit *IERC20PermitCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Permit.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IERC20Permit *IERC20PermitSession) Nonces(owner common.Address) (*big.Int, error) {
	return _IERC20Permit.Contract.Nonces(&_IERC20Permit.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IERC20Permit *IERC20PermitCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _IERC20Permit.Contract.Nonces(&_IERC20Permit.CallOpts, owner)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20Permit *IERC20PermitTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20Permit.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20Permit *IERC20PermitSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20Permit.Contract.Permit(&_IERC20Permit.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20Permit *IERC20PermitTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20Permit.Contract.Permit(&_IERC20Permit.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// MathMetaData contains all meta data concerning the Math contract.
var MathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ce99fef1524ff99fb01e1f85f06090ad10a2f76c97737945f99ae93c007563ad64736f6c63430008130033",
}

// MathABI is the input ABI used to generate the binding from.
// Deprecated: Use MathMetaData.ABI instead.
var MathABI = MathMetaData.ABI

// MathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MathMetaData.Bin instead.
var MathBin = MathMetaData.Bin

// DeployMath deploys a new Ethereum contract, binding an instance of Math to it.
func DeployMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Math, error) {
	parsed, err := MathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	MathCaller     // Read-only binding to the contract
	MathTransactor // Write-only binding to the contract
	MathFilterer   // Log filterer for contract events
}

// MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathSession struct {
	Contract     *Math             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathCallerSession struct {
	Contract *MathCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathTransactorSession struct {
	Contract     *MathTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathRaw struct {
	Contract *Math // Generic contract binding to access the raw methods on
}

// MathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathCallerRaw struct {
	Contract *MathCaller // Generic read-only contract binding to access the raw methods on
}

// MathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathTransactorRaw struct {
	Contract *MathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMath creates a new instance of Math, bound to a specific deployed contract.
func NewMath(address common.Address, backend bind.ContractBackend) (*Math, error) {
	contract, err := bindMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// NewMathCaller creates a new read-only instance of Math, bound to a specific deployed contract.
func NewMathCaller(address common.Address, caller bind.ContractCaller) (*MathCaller, error) {
	contract, err := bindMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathCaller{contract: contract}, nil
}

// NewMathTransactor creates a new write-only instance of Math, bound to a specific deployed contract.
func NewMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTransactor, error) {
	contract, err := bindMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathTransactor{contract: contract}, nil
}

// NewMathFilterer creates a new log filterer instance of Math, bound to a specific deployed contract.
func NewMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MathFilterer, error) {
	contract, err := bindMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathFilterer{contract: contract}, nil
}

// bindMath binds a generic wrapper to an already deployed contract.
func bindMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Math.Contract.MathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Math.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.contract.Transact(opts, method, params...)
}

// OwnableMetaData contains all meta data concerning the Ownable contract.
var OwnableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableMetaData.ABI instead.
var OwnableABI = OwnableMetaData.ABI

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomWalkNFTMetaData contains all meta data concerning the RandomWalkNFT contract.
var RandomWalkNFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"seed\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"MintEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newName\",\"type\":\"string\"}],\"name\":\"TokenNameEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entropy\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastMintTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastMinter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numWithdrawals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"saleTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"seeds\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"seedsOfOwner\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setTokenName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeUntilSale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeUntilWithdrawal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenGenerationScript\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenNames\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"walletOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawalAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawalNums\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalWaitSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405263618dae80600b5566038d7ea4c68000600c5562278d00600d5560006010556000601460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600b5460155560006016556040518060600160405280603581526020016200560660359139601890816200009a9190620004db565b50348015620000a857600080fd5b506040518060400160405280600d81526020017f52616e646f6d57616c6b4e4654000000000000000000000000000000000000008152506040518060400160405280600481526020017f52574c4b000000000000000000000000000000000000000000000000000000008152508160009081620001269190620004db565b508060019081620001389190620004db565b5050506200015b6200014f6200019360201b60201c565b6200019b60201b60201c565b42434060405160200162000171929190620006c1565b6040516020818303038152906040528051906020012060138190555062000703565b600033905090565b6000600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620002e357607f821691505b602082108103620002f957620002f86200029b565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620003637fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000324565b6200036f868362000324565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620003bc620003b6620003b08462000387565b62000391565b62000387565b9050919050565b6000819050919050565b620003d8836200039b565b620003f0620003e782620003c3565b84845462000331565b825550505050565b600090565b62000407620003f8565b62000414818484620003cd565b505050565b5b818110156200043c5762000430600082620003fd565b6001810190506200041a565b5050565b601f8211156200048b576200045581620002ff565b620004608462000314565b8101602085101562000470578190505b620004886200047f8562000314565b83018262000419565b50505b505050565b600082821c905092915050565b6000620004b06000198460080262000490565b1980831691505092915050565b6000620004cb83836200049d565b9150826002028217905092915050565b620004e68262000261565b67ffffffffffffffff8111156200050257620005016200026c565b5b6200050e8254620002ca565b6200051b82828562000440565b600060209050601f8311600181146200055357600084156200053e578287015190505b6200054a8582620004bd565b865550620005ba565b601f1984166200056386620002ff565b60005b828110156200058d5784890151825560018201915060208501945060208101905062000566565b86831015620005ad5784890151620005a9601f8916826200049d565b8355505b6001600288020188555050505b505050505050565b600082825260208201905092915050565b7f412074776f2d64696d656e73696f6e616c2072616e646f6d2077616c6b20776960008201527f6c6c2072657475726e20746f2074686520706f696e742077686572652069742060208201527f737461727465642c2062757420612074687265652d64696d656e73696f6e616c60408201527f206f6e65206d6179206e6f742e00000000000000000000000000000000000000606082015250565b60006200067d606d83620005c2565b91506200068a82620005d3565b608082019050919050565b620006a08162000387565b82525050565b6000819050919050565b620006bb81620006a6565b82525050565b60006060820190508181036000830152620006dc816200066e565b9050620006ed602083018562000695565b620006fc6040830184620006b0565b9392505050565b614ef380620007136000396000f3fe6080604052600436106102515760003560e01c8063661a1cfe11610139578063a7f93ebd116100b6578063e6e268f41161007a578063e6e268f4146108d1578063e985e9c5146108fc578063ee67d6a014610939578063f0503e8014610964578063f2fde38b146109a1578063f454aae1146109ca57610251565b8063a7f93ebd146107da578063b88d4fde14610805578063c87b56dd1461082e578063cb8efe951461086b578063cdb0e89e146108a857610251565b80638da5cb5b116100fd5780638da5cb5b1461070557806395d89b41146107305780639d4635201461075b578063a035b1fe14610786578063a22cb465146107b157610251565b8063661a1cfe1461061e5780636e56f6f91461064957806370a0823114610686578063715018a6146106c357806375794a3c146106da57610251565b80632f745c59116101d2578063438b630011610196578063438b6300146104e857806347ce07cc146105255780634cd609bc146105505780634f6ccce71461057b57806355f804b3146105b85780636352211e146105e157610251565b80632f745c5914610403578063310495ab146104405780633ccfd60b1461047d5780633e8af4da1461049457806342842e0e146104bf57610251565b8063157e394511610219578063157e39451461032e5780631596facb1461035957806317d209ab1461038457806318160ddd146103af57806323b872dd146103da57610251565b806301ffc9a71461025657806306fdde0314610293578063081812fc146102be578063095ea7b3146102fb5780631249c58b14610324575b600080fd5b34801561026257600080fd5b5061027d600480360381019061027891906131fd565b610a07565b60405161028a9190613245565b60405180910390f35b34801561029f57600080fd5b506102a8610a81565b6040516102b591906132f0565b60405180910390f35b3480156102ca57600080fd5b506102e560048036038101906102e09190613348565b610b13565b6040516102f291906133b6565b60405180910390f35b34801561030757600080fd5b50610322600480360381019061031d91906133fd565b610b98565b005b61032c610caf565b005b34801561033a57600080fd5b50610343610fbc565b604051610350919061344c565b60405180910390f35b34801561036557600080fd5b5061036e610fc2565b60405161037b919061344c565b60405180910390f35b34801561039057600080fd5b50610399610fc8565b6040516103a6919061344c565b60405180910390f35b3480156103bb57600080fd5b506103c4610fce565b6040516103d1919061344c565b60405180910390f35b3480156103e657600080fd5b5061040160048036038101906103fc9190613467565b610fdb565b005b34801561040f57600080fd5b5061042a600480360381019061042591906133fd565b61103b565b604051610437919061344c565b60405180910390f35b34801561044c57600080fd5b5061046760048036038101906104629190613348565b6110e0565b60405161047491906132f0565b60405180910390f35b34801561048957600080fd5b50610492611180565b005b3480156104a057600080fd5b506104a9611421565b6040516104b6919061344c565b60405180910390f35b3480156104cb57600080fd5b506104e660048036038101906104e19190613467565b61145b565b005b3480156104f457600080fd5b5061050f600480360381019061050a91906134ba565b61147b565b60405161051c91906135a5565b60405180910390f35b34801561053157600080fd5b5061053a611584565b60405161054791906135e0565b60405180910390f35b34801561055c57600080fd5b5061056561158a565b60405161057291906133b6565b60405180910390f35b34801561058757600080fd5b506105a2600480360381019061059d9190613348565b6115b0565b6040516105af919061344c565b60405180910390f35b3480156105c457600080fd5b506105df60048036038101906105da9190613730565b611621565b005b3480156105ed57600080fd5b5061060860048036038101906106039190613348565b6116b0565b60405161061591906133b6565b60405180910390f35b34801561062a57600080fd5b50610633611761565b604051610640919061344c565b60405180910390f35b34801561065557600080fd5b50610670600480360381019061066b9190613348565b61178a565b60405161067d919061344c565b60405180910390f35b34801561069257600080fd5b506106ad60048036038101906106a891906134ba565b6117a2565b6040516106ba919061344c565b60405180910390f35b3480156106cf57600080fd5b506106d8611859565b005b3480156106e657600080fd5b506106ef6118e1565b6040516106fc919061344c565b60405180910390f35b34801561071157600080fd5b5061071a6118e7565b60405161072791906133b6565b60405180910390f35b34801561073c57600080fd5b50610745611911565b60405161075291906132f0565b60405180910390f35b34801561076757600080fd5b506107706119a3565b60405161077d919061344c565b60405180910390f35b34801561079257600080fd5b5061079b6119a9565b6040516107a8919061344c565b60405180910390f35b3480156107bd57600080fd5b506107d860048036038101906107d391906137a5565b6119af565b005b3480156107e657600080fd5b506107ef611b2f565b6040516107fc919061344c565b60405180910390f35b34801561081157600080fd5b5061082c60048036038101906108279190613886565b611b53565b005b34801561083a57600080fd5b5061085560048036038101906108509190613348565b611bb5565b60405161086291906132f0565b60405180910390f35b34801561087757600080fd5b50610892600480360381019061088d91906134ba565b611c5c565b60405161089f91906139c7565b60405180910390f35b3480156108b457600080fd5b506108cf60048036038101906108ca91906139e9565b611d7e565b005b3480156108dd57600080fd5b506108e6611e71565b6040516108f3919061344c565b60405180910390f35b34801561090857600080fd5b50610923600480360381019061091e9190613a45565b611e85565b6040516109309190613245565b60405180910390f35b34801561094557600080fd5b5061094e611f19565b60405161095b91906132f0565b60405180910390f35b34801561097057600080fd5b5061098b60048036038101906109869190613348565b611fa7565b60405161099891906135e0565b60405180910390f35b3480156109ad57600080fd5b506109c860048036038101906109c391906134ba565b611fbf565b005b3480156109d657600080fd5b506109f160048036038101906109ec9190613348565b6120b6565b6040516109fe919061344c565b60405180910390f35b60007f780e9d63000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610a7a5750610a79826120ce565b5b9050919050565b606060008054610a9090613ab4565b80601f0160208091040260200160405190810160405280929190818152602001828054610abc90613ab4565b8015610b095780601f10610ade57610100808354040283529160200191610b09565b820191906000526020600020905b815481529060010190602001808311610aec57829003601f168201915b5050505050905090565b6000610b1e826121b0565b610b5d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b5490613b57565b60405180910390fd5b6004600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6000610ba3826116b0565b90508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610c13576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c0a90613be9565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16610c3261221c565b73ffffffffffffffffffffffffffffffffffffffff161480610c615750610c6081610c5b61221c565b611e85565b5b610ca0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c9790613c7b565b60405180910390fd5b610caa8383612224565b505050565b6000610cb9611b2f565b905080341015610cfe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cf590613d0d565b60405180910390fd5b600b54421015610d43576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d3a90613d79565b60405180910390fd5b610d4b61221c565b601460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055504260158190555080600c8190555060006016549050600160166000828254610db39190613dc8565b9250508190555060135442434083601460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051602001610df8959493929190613dfc565b60405160208183030381529060405280519060200120601381905550601354600e600083815260200190815260200160002081905550610e5a601460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16826122dd565b600c54341115610f41576000601460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600c5434610ead9190613e4f565b604051610eb990613eb4565b60006040518083038185875af1925050503d8060008114610ef6576040519150601f19603f3d011682016040523d82523d6000602084013e610efb565b606091505b5050905080610f3f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f3690613f15565b60405180910390fd5b505b601460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16817fad2bc79f659de022c64ef55c71f16d0cf125452ed5fc5757b2edc331f58565ec601354600c54604051610fb0929190613f35565b60405180910390a35050565b60105481565b600b5481565b600d5481565b6000600880549050905090565b610fec610fe661221c565b826122fb565b61102b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161102290613fd0565b60405180910390fd5b6110368383836123d9565b505050565b6000611046836117a2565b8210611087576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161107e90614062565b60405180910390fd5b600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600083815260200190815260200160002054905092915050565b600f60205280600052604060002060009150905080546110ff90613ab4565b80601f016020809104026020016040519081016040528092919081815260200182805461112b90613ab4565b80156111785780601f1061114d57610100808354040283529160200191611178565b820191906000526020600020905b81548152906001019060200180831161115b57829003601f168201915b505050505081565b601460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166111c161221c565b73ffffffffffffffffffffffffffffffffffffffff1614611217576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161120e906140ce565b60405180910390fd5b6000611221611421565b14611261576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112589061413a565b60405180910390fd5b6000601460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000601460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600060016016546112db9190613e4f565b905060006112e7611e71565b90506001601060008282546112fc9190613dc8565b92505081905550601054601160008481526020019081526020016000208190555080601260008481526020019081526020016000208190555060008373ffffffffffffffffffffffffffffffffffffffff168260405161135b90613eb4565b60006040518083038185875af1925050503d8060008114611398576040519150601f19603f3d011682016040523d82523d6000602084013e61139d565b606091505b50509050806113e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113d890613f15565b60405180910390fd5b827fa11b556ace4b11a5cae8675a293b51e8cde3a06387d34010861789dfd9e9abc7858460405161141392919061415a565b60405180910390a250505050565b600080600d546015546114349190613dc8565b905042811015611448576000915050611458565b42816114549190613e4f565b9150505b90565b61147683838360405180602001604052806000815250611b53565b505050565b60606000611488836117a2565b9050600081036114e457600067ffffffffffffffff8111156114ad576114ac613605565b5b6040519080825280602002602001820160405280156114db5781602001602082028036833780820191505090505b5091505061157f565b60008167ffffffffffffffff811115611500576114ff613605565b5b60405190808252806020026020018201604052801561152e5781602001602082028036833780820191505090505b50905060005b8281101561157857611546858261103b565b82828151811061155957611558614183565b5b6020026020010181815250508080611570906141b2565b915050611534565b5080925050505b919050565b60135481565b601460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006115ba610fce565b82106115fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115f29061426c565b60405180910390fd5b6008828154811061160f5761160e614183565b5b90600052602060002001549050919050565b61162961221c565b73ffffffffffffffffffffffffffffffffffffffff166116476118e7565b73ffffffffffffffffffffffffffffffffffffffff161461169d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611694906142d8565b60405180910390fd5b80601790816116ac91906144a4565b5050565b6000806002600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611758576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161174f906145e8565b60405180910390fd5b80915050919050565b600042600b5410156117765760009050611787565b42600b546117849190613e4f565b90505b90565b60116020528060005260406000206000915090505481565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611812576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118099061467a565b60405180910390fd5b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b61186161221c565b73ffffffffffffffffffffffffffffffffffffffff1661187f6118e7565b73ffffffffffffffffffffffffffffffffffffffff16146118d5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118cc906142d8565b60405180910390fd5b6118df6000612634565b565b60165481565b6000600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60606001805461192090613ab4565b80601f016020809104026020016040519081016040528092919081815260200182805461194c90613ab4565b80156119995780601f1061196e57610100808354040283529160200191611999565b820191906000526020600020905b81548152906001019060200180831161197c57829003601f168201915b5050505050905090565b60155481565b600c5481565b6119b761221c565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611a24576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a1b906146e6565b60405180910390fd5b8060056000611a3161221c565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff16611ade61221c565b73ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3183604051611b239190613245565b60405180910390a35050565b600061271061271b600c54611b449190614706565b611b4e9190614777565b905090565b611b64611b5e61221c565b836122fb565b611ba3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b9a90613fd0565b60405180910390fd5b611baf848484846126fa565b50505050565b6060611bc0826121b0565b611bff576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611bf69061481a565b60405180910390fd5b6000611c09612756565b90506000815111611c295760405180602001604052806000815250611c54565b80611c33846127e8565b604051602001611c44929190614876565b6040516020818303038152906040525b915050919050565b60606000611c69836117a2565b905060008103611cc557600067ffffffffffffffff811115611c8e57611c8d613605565b5b604051908082528060200260200182016040528015611cbc5781602001602082028036833780820191505090505b50915050611d79565b60008167ffffffffffffffff811115611ce157611ce0613605565b5b604051908082528060200260200182016040528015611d0f5781602001602082028036833780820191505090505b50905060005b82811015611d72576000611d29868361103b565b9050600e600082815260200190815260200160002054838381518110611d5257611d51614183565b5b602002602001018181525050508080611d6a906141b2565b915050611d15565b5080925050505b919050565b611d8f611d8961221c565b836122fb565b611dce576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611dc59061490c565b60405180910390fd5b602081511115611e13576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e0a90614978565b60405180910390fd5b80600f60008481526020019081526020016000209081611e3391906144a4565b507f8ad5e159ff95649c8a9f323ac5a457e741897cf44ce07dfce0e98b84ef9d5f128282604051611e65929190614998565b60405180910390a15050565b6000600247611e809190614777565b905090565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b60188054611f2690613ab4565b80601f0160208091040260200160405190810160405280929190818152602001828054611f5290613ab4565b8015611f9f5780601f10611f7457610100808354040283529160200191611f9f565b820191906000526020600020905b815481529060010190602001808311611f8257829003601f168201915b505050505081565b600e6020528060005260406000206000915090505481565b611fc761221c565b73ffffffffffffffffffffffffffffffffffffffff16611fe56118e7565b73ffffffffffffffffffffffffffffffffffffffff161461203b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612032906142d8565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036120aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120a190614a3a565b60405180910390fd5b6120b381612634565b50565b60126020528060005260406000206000915090505481565b60007f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061219957507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b806121a957506121a882612948565b5b9050919050565b60008073ffffffffffffffffffffffffffffffffffffffff166002600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614159050919050565b600033905090565b816004600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff16612297836116b0565b73ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6122f78282604051806020016040528060008152506129b2565b5050565b6000612306826121b0565b612345576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161233c90614acc565b60405180910390fd5b6000612350836116b0565b90508073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614806123bf57508373ffffffffffffffffffffffffffffffffffffffff166123a784610b13565b73ffffffffffffffffffffffffffffffffffffffff16145b806123d057506123cf8185611e85565b5b91505092915050565b8273ffffffffffffffffffffffffffffffffffffffff166123f9826116b0565b73ffffffffffffffffffffffffffffffffffffffff161461244f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161244690614b5e565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036124be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016124b590614bf0565b60405180910390fd5b6124c9838383612a0d565b6124d4600082612224565b6001600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546125249190613e4f565b925050819055506001600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461257b9190613dc8565b92505081905550816002600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b6000600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6127058484846123d9565b61271184848484612b1f565b612750576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161274790614c82565b60405180910390fd5b50505050565b60606017805461276590613ab4565b80601f016020809104026020016040519081016040528092919081815260200182805461279190613ab4565b80156127de5780601f106127b3576101008083540402835291602001916127de565b820191906000526020600020905b8154815290600101906020018083116127c157829003601f168201915b5050505050905090565b60606000820361282f576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050612943565b600082905060005b6000821461286157808061284a906141b2565b915050600a8261285a9190614777565b9150612837565b60008167ffffffffffffffff81111561287d5761287c613605565b5b6040519080825280601f01601f1916602001820160405280156128af5781602001600182028036833780820191505090505b5090505b6000851461293c576001826128c89190613e4f565b9150600a856128d79190614ca2565b60306128e39190613dc8565b60f81b8183815181106128f9576128f8614183565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a856129359190614777565b94506128b3565b8093505050505b919050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b6129bc8383612ca6565b6129c96000848484612b1f565b612a08576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016129ff90614c82565b60405180910390fd5b505050565b612a18838383612e73565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603612a5a57612a5581612e78565b612a99565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614612a9857612a978382612ec1565b5b5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603612adb57612ad68161302e565b612b1a565b8273ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614612b1957612b1882826130ff565b5b5b505050565b6000612b408473ffffffffffffffffffffffffffffffffffffffff1661317e565b15612c99578373ffffffffffffffffffffffffffffffffffffffff1663150b7a02612b6961221c565b8786866040518563ffffffff1660e01b8152600401612b8b9493929190614d28565b6020604051808303816000875af1925050508015612bc757506040513d601f19601f82011682018060405250810190612bc49190614d89565b60015b612c49573d8060008114612bf7576040519150601f19603f3d011682016040523d82523d6000602084013e612bfc565b606091505b506000815103612c41576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c3890614c82565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614915050612c9e565b600190505b949350505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603612d15576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d0c90614e02565b60405180910390fd5b612d1e816121b0565b15612d5e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d5590614e6e565b60405180910390fd5b612d6a60008383612a0d565b6001600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254612dba9190613dc8565b92505081905550816002600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a45050565b505050565b6008805490506009600083815260200190815260200160002081905550600881908060018154018082558091505060019003906000526020600020016000909190919091505550565b60006001612ece846117a2565b612ed89190613e4f565b9050600060076000848152602001908152602001600020549050818114612fbd576000600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002054905080600660008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002081905550816007600083815260200190815260200160002081905550505b6007600084815260200190815260200160002060009055600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008381526020019081526020016000206000905550505050565b600060016008805490506130429190613e4f565b905060006009600084815260200190815260200160002054905060006008838154811061307257613071614183565b5b90600052602060002001549050806008838154811061309457613093614183565b5b9060005260206000200181905550816009600083815260200190815260200160002081905550600960008581526020019081526020016000206000905560088054806130e3576130e2614e8e565b5b6001900381819060005260206000200160009055905550505050565b600061310a836117a2565b905081600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600083815260200190815260200160002081905550806007600084815260200190815260200160002081905550505050565b600080823b905060008111915050919050565b6000604051905090565b600080fd5b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6131da816131a5565b81146131e557600080fd5b50565b6000813590506131f7816131d1565b92915050565b6000602082840312156132135761321261319b565b5b6000613221848285016131e8565b91505092915050565b60008115159050919050565b61323f8161322a565b82525050565b600060208201905061325a6000830184613236565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561329a57808201518184015260208101905061327f565b60008484015250505050565b6000601f19601f8301169050919050565b60006132c282613260565b6132cc818561326b565b93506132dc81856020860161327c565b6132e5816132a6565b840191505092915050565b6000602082019050818103600083015261330a81846132b7565b905092915050565b6000819050919050565b61332581613312565b811461333057600080fd5b50565b6000813590506133428161331c565b92915050565b60006020828403121561335e5761335d61319b565b5b600061336c84828501613333565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006133a082613375565b9050919050565b6133b081613395565b82525050565b60006020820190506133cb60008301846133a7565b92915050565b6133da81613395565b81146133e557600080fd5b50565b6000813590506133f7816133d1565b92915050565b600080604083850312156134145761341361319b565b5b6000613422858286016133e8565b925050602061343385828601613333565b9150509250929050565b61344681613312565b82525050565b6000602082019050613461600083018461343d565b92915050565b6000806000606084860312156134805761347f61319b565b5b600061348e868287016133e8565b935050602061349f868287016133e8565b92505060406134b086828701613333565b9150509250925092565b6000602082840312156134d0576134cf61319b565b5b60006134de848285016133e8565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61351c81613312565b82525050565b600061352e8383613513565b60208301905092915050565b6000602082019050919050565b6000613552826134e7565b61355c81856134f2565b935061356783613503565b8060005b8381101561359857815161357f8882613522565b975061358a8361353a565b92505060018101905061356b565b5085935050505092915050565b600060208201905081810360008301526135bf8184613547565b905092915050565b6000819050919050565b6135da816135c7565b82525050565b60006020820190506135f560008301846135d1565b92915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61363d826132a6565b810181811067ffffffffffffffff8211171561365c5761365b613605565b5b80604052505050565b600061366f613191565b905061367b8282613634565b919050565b600067ffffffffffffffff82111561369b5761369a613605565b5b6136a4826132a6565b9050602081019050919050565b82818337600083830152505050565b60006136d36136ce84613680565b613665565b9050828152602081018484840111156136ef576136ee613600565b5b6136fa8482856136b1565b509392505050565b600082601f830112613717576137166135fb565b5b81356137278482602086016136c0565b91505092915050565b6000602082840312156137465761374561319b565b5b600082013567ffffffffffffffff811115613764576137636131a0565b5b61377084828501613702565b91505092915050565b6137828161322a565b811461378d57600080fd5b50565b60008135905061379f81613779565b92915050565b600080604083850312156137bc576137bb61319b565b5b60006137ca858286016133e8565b92505060206137db85828601613790565b9150509250929050565b600067ffffffffffffffff821115613800576137ff613605565b5b613809826132a6565b9050602081019050919050565b6000613829613824846137e5565b613665565b90508281526020810184848401111561384557613844613600565b5b6138508482856136b1565b509392505050565b600082601f83011261386d5761386c6135fb565b5b813561387d848260208601613816565b91505092915050565b600080600080608085870312156138a05761389f61319b565b5b60006138ae878288016133e8565b94505060206138bf878288016133e8565b93505060406138d087828801613333565b925050606085013567ffffffffffffffff8111156138f1576138f06131a0565b5b6138fd87828801613858565b91505092959194509250565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61393e816135c7565b82525050565b60006139508383613935565b60208301905092915050565b6000602082019050919050565b600061397482613909565b61397e8185613914565b935061398983613925565b8060005b838110156139ba5781516139a18882613944565b97506139ac8361395c565b92505060018101905061398d565b5085935050505092915050565b600060208201905081810360008301526139e18184613969565b905092915050565b60008060408385031215613a00576139ff61319b565b5b6000613a0e85828601613333565b925050602083013567ffffffffffffffff811115613a2f57613a2e6131a0565b5b613a3b85828601613702565b9150509250929050565b60008060408385031215613a5c57613a5b61319b565b5b6000613a6a858286016133e8565b9250506020613a7b858286016133e8565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680613acc57607f821691505b602082108103613adf57613ade613a85565b5b50919050565b7f4552433732313a20617070726f76656420717565727920666f72206e6f6e657860008201527f697374656e7420746f6b656e0000000000000000000000000000000000000000602082015250565b6000613b41602c8361326b565b9150613b4c82613ae5565b604082019050919050565b60006020820190508181036000830152613b7081613b34565b9050919050565b7f4552433732313a20617070726f76616c20746f2063757272656e74206f776e6560008201527f7200000000000000000000000000000000000000000000000000000000000000602082015250565b6000613bd360218361326b565b9150613bde82613b77565b604082019050919050565b60006020820190508181036000830152613c0281613bc6565b9050919050565b7f4552433732313a20617070726f76652063616c6c6572206973206e6f74206f7760008201527f6e6572206e6f7220617070726f76656420666f7220616c6c0000000000000000602082015250565b6000613c6560388361326b565b9150613c7082613c09565b604082019050919050565b60006020820190508181036000830152613c9481613c58565b9050919050565b7f5468652076616c7565207375626d69747465642077697468207468697320747260008201527f616e73616374696f6e20697320746f6f206c6f772e0000000000000000000000602082015250565b6000613cf760358361326b565b9150613d0282613c9b565b604082019050919050565b60006020820190508181036000830152613d2681613cea565b9050919050565b7f5468652073616c65206973206e6f74206f70656e207965742e00000000000000600082015250565b6000613d6360198361326b565b9150613d6e82613d2d565b602082019050919050565b60006020820190508181036000830152613d9281613d56565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000613dd382613312565b9150613dde83613312565b9250828201905080821115613df657613df5613d99565b5b92915050565b600060a082019050613e1160008301886135d1565b613e1e602083018761343d565b613e2b60408301866135d1565b613e38606083018561343d565b613e4560808301846133a7565b9695505050505050565b6000613e5a82613312565b9150613e6583613312565b9250828203905081811115613e7d57613e7c613d99565b5b92915050565b600081905092915050565b50565b6000613e9e600083613e83565b9150613ea982613e8e565b600082019050919050565b6000613ebf82613e91565b9150819050919050565b7f5472616e73666572206661696c65642e00000000000000000000000000000000600082015250565b6000613eff60108361326b565b9150613f0a82613ec9565b602082019050919050565b60006020820190508181036000830152613f2e81613ef2565b9050919050565b6000604082019050613f4a60008301856135d1565b613f57602083018461343d565b9392505050565b7f4552433732313a207472616e736665722063616c6c6572206973206e6f74206f60008201527f776e6572206e6f7220617070726f766564000000000000000000000000000000602082015250565b6000613fba60318361326b565b9150613fc582613f5e565b604082019050919050565b60006020820190508181036000830152613fe981613fad565b9050919050565b7f455243373231456e756d657261626c653a206f776e657220696e646578206f7560008201527f74206f6620626f756e6473000000000000000000000000000000000000000000602082015250565b600061404c602b8361326b565b915061405782613ff0565b604082019050919050565b6000602082019050818103600083015261407b8161403f565b9050919050565b7f4f6e6c79206c617374206d696e7465722063616e2077697468647261772e0000600082015250565b60006140b8601e8361326b565b91506140c382614082565b602082019050919050565b600060208201905081810360008301526140e7816140ab565b9050919050565b7f4e6f7420656e6f7567682074696d652068617320656c61707365642e00000000600082015250565b6000614124601c8361326b565b915061412f826140ee565b602082019050919050565b6000602082019050818103600083015261415381614117565b9050919050565b600060408201905061416f60008301856133a7565b61417c602083018461343d565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006141bd82613312565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036141ef576141ee613d99565b5b600182019050919050565b7f455243373231456e756d657261626c653a20676c6f62616c20696e646578206f60008201527f7574206f6620626f756e64730000000000000000000000000000000000000000602082015250565b6000614256602c8361326b565b9150614261826141fa565b604082019050919050565b6000602082019050818103600083015261428581614249565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b60006142c260208361326b565b91506142cd8261428c565b602082019050919050565b600060208201905081810360008301526142f1816142b5565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b60006008830261435a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261431d565b614364868361431d565b95508019841693508086168417925050509392505050565b6000819050919050565b60006143a161439c61439784613312565b61437c565b613312565b9050919050565b6000819050919050565b6143bb83614386565b6143cf6143c7826143a8565b84845461432a565b825550505050565b600090565b6143e46143d7565b6143ef8184846143b2565b505050565b5b81811015614413576144086000826143dc565b6001810190506143f5565b5050565b601f82111561445857614429816142f8565b6144328461430d565b81016020851015614441578190505b61445561444d8561430d565b8301826143f4565b50505b505050565b600082821c905092915050565b600061447b6000198460080261445d565b1980831691505092915050565b6000614494838361446a565b9150826002028217905092915050565b6144ad82613260565b67ffffffffffffffff8111156144c6576144c5613605565b5b6144d08254613ab4565b6144db828285614417565b600060209050601f83116001811461450e57600084156144fc578287015190505b6145068582614488565b86555061456e565b601f19841661451c866142f8565b60005b828110156145445784890151825560018201915060208501945060208101905061451f565b86831015614561578489015161455d601f89168261446a565b8355505b6001600288020188555050505b505050505050565b7f4552433732313a206f776e657220717565727920666f72206e6f6e657869737460008201527f656e7420746f6b656e0000000000000000000000000000000000000000000000602082015250565b60006145d260298361326b565b91506145dd82614576565b604082019050919050565b60006020820190508181036000830152614601816145c5565b9050919050565b7f4552433732313a2062616c616e636520717565727920666f7220746865207a6560008201527f726f206164647265737300000000000000000000000000000000000000000000602082015250565b6000614664602a8361326b565b915061466f82614608565b604082019050919050565b6000602082019050818103600083015261469381614657565b9050919050565b7f4552433732313a20617070726f766520746f2063616c6c657200000000000000600082015250565b60006146d060198361326b565b91506146db8261469a565b602082019050919050565b600060208201905081810360008301526146ff816146c3565b9050919050565b600061471182613312565b915061471c83613312565b925082820261472a81613312565b9150828204841483151761474157614740613d99565b5b5092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061478282613312565b915061478d83613312565b92508261479d5761479c614748565b5b828204905092915050565b7f4552433732314d657461646174613a2055524920717565727920666f72206e6f60008201527f6e6578697374656e7420746f6b656e0000000000000000000000000000000000602082015250565b6000614804602f8361326b565b915061480f826147a8565b604082019050919050565b60006020820190508181036000830152614833816147f7565b9050919050565b600081905092915050565b600061485082613260565b61485a818561483a565b935061486a81856020860161327c565b80840191505092915050565b60006148828285614845565b915061488e8284614845565b91508190509392505050565b7f736574546f6b656e4e616d652063616c6c6572206973206e6f74206f776e657260008201527f206e6f7220617070726f76656400000000000000000000000000000000000000602082015250565b60006148f6602d8361326b565b91506149018261489a565b604082019050919050565b60006020820190508181036000830152614925816148e9565b9050919050565b7f546f6b656e206e616d6520697320746f6f206c6f6e672e000000000000000000600082015250565b600061496260178361326b565b915061496d8261492c565b602082019050919050565b6000602082019050818103600083015261499181614955565b9050919050565b60006040820190506149ad600083018561343d565b81810360208301526149bf81846132b7565b90509392505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000614a2460268361326b565b9150614a2f826149c8565b604082019050919050565b60006020820190508181036000830152614a5381614a17565b9050919050565b7f4552433732313a206f70657261746f7220717565727920666f72206e6f6e657860008201527f697374656e7420746f6b656e0000000000000000000000000000000000000000602082015250565b6000614ab6602c8361326b565b9150614ac182614a5a565b604082019050919050565b60006020820190508181036000830152614ae581614aa9565b9050919050565b7f4552433732313a207472616e73666572206f6620746f6b656e2074686174206960008201527f73206e6f74206f776e0000000000000000000000000000000000000000000000602082015250565b6000614b4860298361326b565b9150614b5382614aec565b604082019050919050565b60006020820190508181036000830152614b7781614b3b565b9050919050565b7f4552433732313a207472616e7366657220746f20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b6000614bda60248361326b565b9150614be582614b7e565b604082019050919050565b60006020820190508181036000830152614c0981614bcd565b9050919050565b7f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560008201527f63656976657220696d706c656d656e7465720000000000000000000000000000602082015250565b6000614c6c60328361326b565b9150614c7782614c10565b604082019050919050565b60006020820190508181036000830152614c9b81614c5f565b9050919050565b6000614cad82613312565b9150614cb883613312565b925082614cc857614cc7614748565b5b828206905092915050565b600081519050919050565b600082825260208201905092915050565b6000614cfa82614cd3565b614d048185614cde565b9350614d1481856020860161327c565b614d1d816132a6565b840191505092915050565b6000608082019050614d3d60008301876133a7565b614d4a60208301866133a7565b614d57604083018561343d565b8181036060830152614d698184614cef565b905095945050505050565b600081519050614d83816131d1565b92915050565b600060208284031215614d9f57614d9e61319b565b5b6000614dad84828501614d74565b91505092915050565b7f4552433732313a206d696e7420746f20746865207a65726f2061646472657373600082015250565b6000614dec60208361326b565b9150614df782614db6565b602082019050919050565b60006020820190508181036000830152614e1b81614ddf565b9050919050565b7f4552433732313a20746f6b656e20616c7265616479206d696e74656400000000600082015250565b6000614e58601c8361326b565b9150614e6382614e22565b602082019050919050565b60006020820190508181036000830152614e8781614e4b565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea264697066735822122047852f6ed623c78af33027a5946e09618e4c4c3862daee4d6a23ca6c550b736264736f6c63430008130033697066733a2f2f516d50375a385662514c7079747a586e6365654141633444357458333958567a6f4565555a77454b3861506b3857",
}

// RandomWalkNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use RandomWalkNFTMetaData.ABI instead.
var RandomWalkNFTABI = RandomWalkNFTMetaData.ABI

// RandomWalkNFTBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RandomWalkNFTMetaData.Bin instead.
var RandomWalkNFTBin = RandomWalkNFTMetaData.Bin

// DeployRandomWalkNFT deploys a new Ethereum contract, binding an instance of RandomWalkNFT to it.
func DeployRandomWalkNFT(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RandomWalkNFT, error) {
	parsed, err := RandomWalkNFTMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RandomWalkNFTBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RandomWalkNFT{RandomWalkNFTCaller: RandomWalkNFTCaller{contract: contract}, RandomWalkNFTTransactor: RandomWalkNFTTransactor{contract: contract}, RandomWalkNFTFilterer: RandomWalkNFTFilterer{contract: contract}}, nil
}

// RandomWalkNFT is an auto generated Go binding around an Ethereum contract.
type RandomWalkNFT struct {
	RandomWalkNFTCaller     // Read-only binding to the contract
	RandomWalkNFTTransactor // Write-only binding to the contract
	RandomWalkNFTFilterer   // Log filterer for contract events
}

// RandomWalkNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type RandomWalkNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomWalkNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RandomWalkNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomWalkNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RandomWalkNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomWalkNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RandomWalkNFTSession struct {
	Contract     *RandomWalkNFT    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RandomWalkNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RandomWalkNFTCallerSession struct {
	Contract *RandomWalkNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// RandomWalkNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RandomWalkNFTTransactorSession struct {
	Contract     *RandomWalkNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RandomWalkNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type RandomWalkNFTRaw struct {
	Contract *RandomWalkNFT // Generic contract binding to access the raw methods on
}

// RandomWalkNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RandomWalkNFTCallerRaw struct {
	Contract *RandomWalkNFTCaller // Generic read-only contract binding to access the raw methods on
}

// RandomWalkNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RandomWalkNFTTransactorRaw struct {
	Contract *RandomWalkNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRandomWalkNFT creates a new instance of RandomWalkNFT, bound to a specific deployed contract.
func NewRandomWalkNFT(address common.Address, backend bind.ContractBackend) (*RandomWalkNFT, error) {
	contract, err := bindRandomWalkNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RandomWalkNFT{RandomWalkNFTCaller: RandomWalkNFTCaller{contract: contract}, RandomWalkNFTTransactor: RandomWalkNFTTransactor{contract: contract}, RandomWalkNFTFilterer: RandomWalkNFTFilterer{contract: contract}}, nil
}

// NewRandomWalkNFTCaller creates a new read-only instance of RandomWalkNFT, bound to a specific deployed contract.
func NewRandomWalkNFTCaller(address common.Address, caller bind.ContractCaller) (*RandomWalkNFTCaller, error) {
	contract, err := bindRandomWalkNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RandomWalkNFTCaller{contract: contract}, nil
}

// NewRandomWalkNFTTransactor creates a new write-only instance of RandomWalkNFT, bound to a specific deployed contract.
func NewRandomWalkNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*RandomWalkNFTTransactor, error) {
	contract, err := bindRandomWalkNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RandomWalkNFTTransactor{contract: contract}, nil
}

// NewRandomWalkNFTFilterer creates a new log filterer instance of RandomWalkNFT, bound to a specific deployed contract.
func NewRandomWalkNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*RandomWalkNFTFilterer, error) {
	contract, err := bindRandomWalkNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RandomWalkNFTFilterer{contract: contract}, nil
}

// bindRandomWalkNFT binds a generic wrapper to an already deployed contract.
func bindRandomWalkNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RandomWalkNFTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RandomWalkNFT *RandomWalkNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RandomWalkNFT.Contract.RandomWalkNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RandomWalkNFT *RandomWalkNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.RandomWalkNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RandomWalkNFT *RandomWalkNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.RandomWalkNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RandomWalkNFT *RandomWalkNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RandomWalkNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RandomWalkNFT *RandomWalkNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RandomWalkNFT *RandomWalkNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _RandomWalkNFT.Contract.BalanceOf(&_RandomWalkNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _RandomWalkNFT.Contract.BalanceOf(&_RandomWalkNFT.CallOpts, owner)
}

// Entropy is a free data retrieval call binding the contract method 0x47ce07cc.
//
// Solidity: function entropy() view returns(bytes32)
func (_RandomWalkNFT *RandomWalkNFTCaller) Entropy(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "entropy")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Entropy is a free data retrieval call binding the contract method 0x47ce07cc.
//
// Solidity: function entropy() view returns(bytes32)
func (_RandomWalkNFT *RandomWalkNFTSession) Entropy() ([32]byte, error) {
	return _RandomWalkNFT.Contract.Entropy(&_RandomWalkNFT.CallOpts)
}

// Entropy is a free data retrieval call binding the contract method 0x47ce07cc.
//
// Solidity: function entropy() view returns(bytes32)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) Entropy() ([32]byte, error) {
	return _RandomWalkNFT.Contract.Entropy(&_RandomWalkNFT.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_RandomWalkNFT *RandomWalkNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_RandomWalkNFT *RandomWalkNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _RandomWalkNFT.Contract.GetApproved(&_RandomWalkNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _RandomWalkNFT.Contract.GetApproved(&_RandomWalkNFT.CallOpts, tokenId)
}

// GetMintPrice is a free data retrieval call binding the contract method 0xa7f93ebd.
//
// Solidity: function getMintPrice() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) GetMintPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "getMintPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMintPrice is a free data retrieval call binding the contract method 0xa7f93ebd.
//
// Solidity: function getMintPrice() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) GetMintPrice() (*big.Int, error) {
	return _RandomWalkNFT.Contract.GetMintPrice(&_RandomWalkNFT.CallOpts)
}

// GetMintPrice is a free data retrieval call binding the contract method 0xa7f93ebd.
//
// Solidity: function getMintPrice() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) GetMintPrice() (*big.Int, error) {
	return _RandomWalkNFT.Contract.GetMintPrice(&_RandomWalkNFT.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_RandomWalkNFT *RandomWalkNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_RandomWalkNFT *RandomWalkNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _RandomWalkNFT.Contract.IsApprovedForAll(&_RandomWalkNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _RandomWalkNFT.Contract.IsApprovedForAll(&_RandomWalkNFT.CallOpts, owner, operator)
}

// LastMintTime is a free data retrieval call binding the contract method 0x9d463520.
//
// Solidity: function lastMintTime() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) LastMintTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "lastMintTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastMintTime is a free data retrieval call binding the contract method 0x9d463520.
//
// Solidity: function lastMintTime() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) LastMintTime() (*big.Int, error) {
	return _RandomWalkNFT.Contract.LastMintTime(&_RandomWalkNFT.CallOpts)
}

// LastMintTime is a free data retrieval call binding the contract method 0x9d463520.
//
// Solidity: function lastMintTime() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) LastMintTime() (*big.Int, error) {
	return _RandomWalkNFT.Contract.LastMintTime(&_RandomWalkNFT.CallOpts)
}

// LastMinter is a free data retrieval call binding the contract method 0x4cd609bc.
//
// Solidity: function lastMinter() view returns(address)
func (_RandomWalkNFT *RandomWalkNFTCaller) LastMinter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "lastMinter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LastMinter is a free data retrieval call binding the contract method 0x4cd609bc.
//
// Solidity: function lastMinter() view returns(address)
func (_RandomWalkNFT *RandomWalkNFTSession) LastMinter() (common.Address, error) {
	return _RandomWalkNFT.Contract.LastMinter(&_RandomWalkNFT.CallOpts)
}

// LastMinter is a free data retrieval call binding the contract method 0x4cd609bc.
//
// Solidity: function lastMinter() view returns(address)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) LastMinter() (common.Address, error) {
	return _RandomWalkNFT.Contract.LastMinter(&_RandomWalkNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RandomWalkNFT *RandomWalkNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RandomWalkNFT *RandomWalkNFTSession) Name() (string, error) {
	return _RandomWalkNFT.Contract.Name(&_RandomWalkNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) Name() (string, error) {
	return _RandomWalkNFT.Contract.Name(&_RandomWalkNFT.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) NextTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "nextTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) NextTokenId() (*big.Int, error) {
	return _RandomWalkNFT.Contract.NextTokenId(&_RandomWalkNFT.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) NextTokenId() (*big.Int, error) {
	return _RandomWalkNFT.Contract.NextTokenId(&_RandomWalkNFT.CallOpts)
}

// NumWithdrawals is a free data retrieval call binding the contract method 0x157e3945.
//
// Solidity: function numWithdrawals() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) NumWithdrawals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "numWithdrawals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumWithdrawals is a free data retrieval call binding the contract method 0x157e3945.
//
// Solidity: function numWithdrawals() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) NumWithdrawals() (*big.Int, error) {
	return _RandomWalkNFT.Contract.NumWithdrawals(&_RandomWalkNFT.CallOpts)
}

// NumWithdrawals is a free data retrieval call binding the contract method 0x157e3945.
//
// Solidity: function numWithdrawals() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) NumWithdrawals() (*big.Int, error) {
	return _RandomWalkNFT.Contract.NumWithdrawals(&_RandomWalkNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RandomWalkNFT *RandomWalkNFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RandomWalkNFT *RandomWalkNFTSession) Owner() (common.Address, error) {
	return _RandomWalkNFT.Contract.Owner(&_RandomWalkNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) Owner() (common.Address, error) {
	return _RandomWalkNFT.Contract.Owner(&_RandomWalkNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_RandomWalkNFT *RandomWalkNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_RandomWalkNFT *RandomWalkNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _RandomWalkNFT.Contract.OwnerOf(&_RandomWalkNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _RandomWalkNFT.Contract.OwnerOf(&_RandomWalkNFT.CallOpts, tokenId)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) Price() (*big.Int, error) {
	return _RandomWalkNFT.Contract.Price(&_RandomWalkNFT.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) Price() (*big.Int, error) {
	return _RandomWalkNFT.Contract.Price(&_RandomWalkNFT.CallOpts)
}

// SaleTime is a free data retrieval call binding the contract method 0x1596facb.
//
// Solidity: function saleTime() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) SaleTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "saleTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SaleTime is a free data retrieval call binding the contract method 0x1596facb.
//
// Solidity: function saleTime() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) SaleTime() (*big.Int, error) {
	return _RandomWalkNFT.Contract.SaleTime(&_RandomWalkNFT.CallOpts)
}

// SaleTime is a free data retrieval call binding the contract method 0x1596facb.
//
// Solidity: function saleTime() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) SaleTime() (*big.Int, error) {
	return _RandomWalkNFT.Contract.SaleTime(&_RandomWalkNFT.CallOpts)
}

// Seeds is a free data retrieval call binding the contract method 0xf0503e80.
//
// Solidity: function seeds(uint256 ) view returns(bytes32)
func (_RandomWalkNFT *RandomWalkNFTCaller) Seeds(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "seeds", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Seeds is a free data retrieval call binding the contract method 0xf0503e80.
//
// Solidity: function seeds(uint256 ) view returns(bytes32)
func (_RandomWalkNFT *RandomWalkNFTSession) Seeds(arg0 *big.Int) ([32]byte, error) {
	return _RandomWalkNFT.Contract.Seeds(&_RandomWalkNFT.CallOpts, arg0)
}

// Seeds is a free data retrieval call binding the contract method 0xf0503e80.
//
// Solidity: function seeds(uint256 ) view returns(bytes32)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) Seeds(arg0 *big.Int) ([32]byte, error) {
	return _RandomWalkNFT.Contract.Seeds(&_RandomWalkNFT.CallOpts, arg0)
}

// SeedsOfOwner is a free data retrieval call binding the contract method 0xcb8efe95.
//
// Solidity: function seedsOfOwner(address _owner) view returns(bytes32[])
func (_RandomWalkNFT *RandomWalkNFTCaller) SeedsOfOwner(opts *bind.CallOpts, _owner common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "seedsOfOwner", _owner)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// SeedsOfOwner is a free data retrieval call binding the contract method 0xcb8efe95.
//
// Solidity: function seedsOfOwner(address _owner) view returns(bytes32[])
func (_RandomWalkNFT *RandomWalkNFTSession) SeedsOfOwner(_owner common.Address) ([][32]byte, error) {
	return _RandomWalkNFT.Contract.SeedsOfOwner(&_RandomWalkNFT.CallOpts, _owner)
}

// SeedsOfOwner is a free data retrieval call binding the contract method 0xcb8efe95.
//
// Solidity: function seedsOfOwner(address _owner) view returns(bytes32[])
func (_RandomWalkNFT *RandomWalkNFTCallerSession) SeedsOfOwner(_owner common.Address) ([][32]byte, error) {
	return _RandomWalkNFT.Contract.SeedsOfOwner(&_RandomWalkNFT.CallOpts, _owner)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RandomWalkNFT *RandomWalkNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RandomWalkNFT *RandomWalkNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RandomWalkNFT.Contract.SupportsInterface(&_RandomWalkNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RandomWalkNFT.Contract.SupportsInterface(&_RandomWalkNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RandomWalkNFT *RandomWalkNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RandomWalkNFT *RandomWalkNFTSession) Symbol() (string, error) {
	return _RandomWalkNFT.Contract.Symbol(&_RandomWalkNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) Symbol() (string, error) {
	return _RandomWalkNFT.Contract.Symbol(&_RandomWalkNFT.CallOpts)
}

// TimeUntilSale is a free data retrieval call binding the contract method 0x661a1cfe.
//
// Solidity: function timeUntilSale() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) TimeUntilSale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "timeUntilSale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeUntilSale is a free data retrieval call binding the contract method 0x661a1cfe.
//
// Solidity: function timeUntilSale() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) TimeUntilSale() (*big.Int, error) {
	return _RandomWalkNFT.Contract.TimeUntilSale(&_RandomWalkNFT.CallOpts)
}

// TimeUntilSale is a free data retrieval call binding the contract method 0x661a1cfe.
//
// Solidity: function timeUntilSale() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) TimeUntilSale() (*big.Int, error) {
	return _RandomWalkNFT.Contract.TimeUntilSale(&_RandomWalkNFT.CallOpts)
}

// TimeUntilWithdrawal is a free data retrieval call binding the contract method 0x3e8af4da.
//
// Solidity: function timeUntilWithdrawal() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) TimeUntilWithdrawal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "timeUntilWithdrawal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeUntilWithdrawal is a free data retrieval call binding the contract method 0x3e8af4da.
//
// Solidity: function timeUntilWithdrawal() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) TimeUntilWithdrawal() (*big.Int, error) {
	return _RandomWalkNFT.Contract.TimeUntilWithdrawal(&_RandomWalkNFT.CallOpts)
}

// TimeUntilWithdrawal is a free data retrieval call binding the contract method 0x3e8af4da.
//
// Solidity: function timeUntilWithdrawal() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) TimeUntilWithdrawal() (*big.Int, error) {
	return _RandomWalkNFT.Contract.TimeUntilWithdrawal(&_RandomWalkNFT.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _RandomWalkNFT.Contract.TokenByIndex(&_RandomWalkNFT.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _RandomWalkNFT.Contract.TokenByIndex(&_RandomWalkNFT.CallOpts, index)
}

// TokenGenerationScript is a free data retrieval call binding the contract method 0xee67d6a0.
//
// Solidity: function tokenGenerationScript() view returns(string)
func (_RandomWalkNFT *RandomWalkNFTCaller) TokenGenerationScript(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "tokenGenerationScript")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenGenerationScript is a free data retrieval call binding the contract method 0xee67d6a0.
//
// Solidity: function tokenGenerationScript() view returns(string)
func (_RandomWalkNFT *RandomWalkNFTSession) TokenGenerationScript() (string, error) {
	return _RandomWalkNFT.Contract.TokenGenerationScript(&_RandomWalkNFT.CallOpts)
}

// TokenGenerationScript is a free data retrieval call binding the contract method 0xee67d6a0.
//
// Solidity: function tokenGenerationScript() view returns(string)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) TokenGenerationScript() (string, error) {
	return _RandomWalkNFT.Contract.TokenGenerationScript(&_RandomWalkNFT.CallOpts)
}

// TokenNames is a free data retrieval call binding the contract method 0x310495ab.
//
// Solidity: function tokenNames(uint256 ) view returns(string)
func (_RandomWalkNFT *RandomWalkNFTCaller) TokenNames(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "tokenNames", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenNames is a free data retrieval call binding the contract method 0x310495ab.
//
// Solidity: function tokenNames(uint256 ) view returns(string)
func (_RandomWalkNFT *RandomWalkNFTSession) TokenNames(arg0 *big.Int) (string, error) {
	return _RandomWalkNFT.Contract.TokenNames(&_RandomWalkNFT.CallOpts, arg0)
}

// TokenNames is a free data retrieval call binding the contract method 0x310495ab.
//
// Solidity: function tokenNames(uint256 ) view returns(string)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) TokenNames(arg0 *big.Int) (string, error) {
	return _RandomWalkNFT.Contract.TokenNames(&_RandomWalkNFT.CallOpts, arg0)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _RandomWalkNFT.Contract.TokenOfOwnerByIndex(&_RandomWalkNFT.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _RandomWalkNFT.Contract.TokenOfOwnerByIndex(&_RandomWalkNFT.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_RandomWalkNFT *RandomWalkNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_RandomWalkNFT *RandomWalkNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _RandomWalkNFT.Contract.TokenURI(&_RandomWalkNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _RandomWalkNFT.Contract.TokenURI(&_RandomWalkNFT.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) TotalSupply() (*big.Int, error) {
	return _RandomWalkNFT.Contract.TotalSupply(&_RandomWalkNFT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) TotalSupply() (*big.Int, error) {
	return _RandomWalkNFT.Contract.TotalSupply(&_RandomWalkNFT.CallOpts)
}

// WalletOfOwner is a free data retrieval call binding the contract method 0x438b6300.
//
// Solidity: function walletOfOwner(address _owner) view returns(uint256[])
func (_RandomWalkNFT *RandomWalkNFTCaller) WalletOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "walletOfOwner", _owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// WalletOfOwner is a free data retrieval call binding the contract method 0x438b6300.
//
// Solidity: function walletOfOwner(address _owner) view returns(uint256[])
func (_RandomWalkNFT *RandomWalkNFTSession) WalletOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _RandomWalkNFT.Contract.WalletOfOwner(&_RandomWalkNFT.CallOpts, _owner)
}

// WalletOfOwner is a free data retrieval call binding the contract method 0x438b6300.
//
// Solidity: function walletOfOwner(address _owner) view returns(uint256[])
func (_RandomWalkNFT *RandomWalkNFTCallerSession) WalletOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _RandomWalkNFT.Contract.WalletOfOwner(&_RandomWalkNFT.CallOpts, _owner)
}

// WithdrawalAmount is a free data retrieval call binding the contract method 0xe6e268f4.
//
// Solidity: function withdrawalAmount() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) WithdrawalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "withdrawalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalAmount is a free data retrieval call binding the contract method 0xe6e268f4.
//
// Solidity: function withdrawalAmount() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) WithdrawalAmount() (*big.Int, error) {
	return _RandomWalkNFT.Contract.WithdrawalAmount(&_RandomWalkNFT.CallOpts)
}

// WithdrawalAmount is a free data retrieval call binding the contract method 0xe6e268f4.
//
// Solidity: function withdrawalAmount() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) WithdrawalAmount() (*big.Int, error) {
	return _RandomWalkNFT.Contract.WithdrawalAmount(&_RandomWalkNFT.CallOpts)
}

// WithdrawalAmounts is a free data retrieval call binding the contract method 0xf454aae1.
//
// Solidity: function withdrawalAmounts(uint256 ) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) WithdrawalAmounts(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "withdrawalAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalAmounts is a free data retrieval call binding the contract method 0xf454aae1.
//
// Solidity: function withdrawalAmounts(uint256 ) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) WithdrawalAmounts(arg0 *big.Int) (*big.Int, error) {
	return _RandomWalkNFT.Contract.WithdrawalAmounts(&_RandomWalkNFT.CallOpts, arg0)
}

// WithdrawalAmounts is a free data retrieval call binding the contract method 0xf454aae1.
//
// Solidity: function withdrawalAmounts(uint256 ) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) WithdrawalAmounts(arg0 *big.Int) (*big.Int, error) {
	return _RandomWalkNFT.Contract.WithdrawalAmounts(&_RandomWalkNFT.CallOpts, arg0)
}

// WithdrawalNums is a free data retrieval call binding the contract method 0x6e56f6f9.
//
// Solidity: function withdrawalNums(uint256 ) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) WithdrawalNums(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "withdrawalNums", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalNums is a free data retrieval call binding the contract method 0x6e56f6f9.
//
// Solidity: function withdrawalNums(uint256 ) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) WithdrawalNums(arg0 *big.Int) (*big.Int, error) {
	return _RandomWalkNFT.Contract.WithdrawalNums(&_RandomWalkNFT.CallOpts, arg0)
}

// WithdrawalNums is a free data retrieval call binding the contract method 0x6e56f6f9.
//
// Solidity: function withdrawalNums(uint256 ) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) WithdrawalNums(arg0 *big.Int) (*big.Int, error) {
	return _RandomWalkNFT.Contract.WithdrawalNums(&_RandomWalkNFT.CallOpts, arg0)
}

// WithdrawalWaitSeconds is a free data retrieval call binding the contract method 0x17d209ab.
//
// Solidity: function withdrawalWaitSeconds() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) WithdrawalWaitSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "withdrawalWaitSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalWaitSeconds is a free data retrieval call binding the contract method 0x17d209ab.
//
// Solidity: function withdrawalWaitSeconds() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) WithdrawalWaitSeconds() (*big.Int, error) {
	return _RandomWalkNFT.Contract.WithdrawalWaitSeconds(&_RandomWalkNFT.CallOpts)
}

// WithdrawalWaitSeconds is a free data retrieval call binding the contract method 0x17d209ab.
//
// Solidity: function withdrawalWaitSeconds() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) WithdrawalWaitSeconds() (*big.Int, error) {
	return _RandomWalkNFT.Contract.WithdrawalWaitSeconds(&_RandomWalkNFT.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RandomWalkNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_RandomWalkNFT *RandomWalkNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.Approve(&_RandomWalkNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.Approve(&_RandomWalkNFT.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_RandomWalkNFT *RandomWalkNFTTransactor) Mint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomWalkNFT.contract.Transact(opts, "mint")
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_RandomWalkNFT *RandomWalkNFTSession) Mint() (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.Mint(&_RandomWalkNFT.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_RandomWalkNFT *RandomWalkNFTTransactorSession) Mint() (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.Mint(&_RandomWalkNFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RandomWalkNFT *RandomWalkNFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomWalkNFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RandomWalkNFT *RandomWalkNFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.RenounceOwnership(&_RandomWalkNFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RandomWalkNFT *RandomWalkNFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.RenounceOwnership(&_RandomWalkNFT.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RandomWalkNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_RandomWalkNFT *RandomWalkNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.SafeTransferFrom(&_RandomWalkNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.SafeTransferFrom(&_RandomWalkNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _RandomWalkNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_RandomWalkNFT *RandomWalkNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.SafeTransferFrom0(&_RandomWalkNFT.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.SafeTransferFrom0(&_RandomWalkNFT.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _RandomWalkNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_RandomWalkNFT *RandomWalkNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.SetApprovalForAll(&_RandomWalkNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.SetApprovalForAll(&_RandomWalkNFT.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactor) SetBaseURI(opts *bind.TransactOpts, baseURI string) (*types.Transaction, error) {
	return _RandomWalkNFT.contract.Transact(opts, "setBaseURI", baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_RandomWalkNFT *RandomWalkNFTSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.SetBaseURI(&_RandomWalkNFT.TransactOpts, baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactorSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.SetBaseURI(&_RandomWalkNFT.TransactOpts, baseURI)
}

// SetTokenName is a paid mutator transaction binding the contract method 0xcdb0e89e.
//
// Solidity: function setTokenName(uint256 tokenId, string name) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactor) SetTokenName(opts *bind.TransactOpts, tokenId *big.Int, name string) (*types.Transaction, error) {
	return _RandomWalkNFT.contract.Transact(opts, "setTokenName", tokenId, name)
}

// SetTokenName is a paid mutator transaction binding the contract method 0xcdb0e89e.
//
// Solidity: function setTokenName(uint256 tokenId, string name) returns()
func (_RandomWalkNFT *RandomWalkNFTSession) SetTokenName(tokenId *big.Int, name string) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.SetTokenName(&_RandomWalkNFT.TransactOpts, tokenId, name)
}

// SetTokenName is a paid mutator transaction binding the contract method 0xcdb0e89e.
//
// Solidity: function setTokenName(uint256 tokenId, string name) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactorSession) SetTokenName(tokenId *big.Int, name string) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.SetTokenName(&_RandomWalkNFT.TransactOpts, tokenId, name)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RandomWalkNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_RandomWalkNFT *RandomWalkNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.TransferFrom(&_RandomWalkNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.TransferFrom(&_RandomWalkNFT.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RandomWalkNFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RandomWalkNFT *RandomWalkNFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.TransferOwnership(&_RandomWalkNFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.TransferOwnership(&_RandomWalkNFT.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_RandomWalkNFT *RandomWalkNFTTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomWalkNFT.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_RandomWalkNFT *RandomWalkNFTSession) Withdraw() (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.Withdraw(&_RandomWalkNFT.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_RandomWalkNFT *RandomWalkNFTTransactorSession) Withdraw() (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.Withdraw(&_RandomWalkNFT.TransactOpts)
}

// RandomWalkNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the RandomWalkNFT contract.
type RandomWalkNFTApprovalIterator struct {
	Event *RandomWalkNFTApproval // Event containing the contract specifics and raw log

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
func (it *RandomWalkNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomWalkNFTApproval)
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
		it.Event = new(RandomWalkNFTApproval)
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
func (it *RandomWalkNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomWalkNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomWalkNFTApproval represents a Approval event raised by the RandomWalkNFT contract.
type RandomWalkNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_RandomWalkNFT *RandomWalkNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*RandomWalkNFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _RandomWalkNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &RandomWalkNFTApprovalIterator{contract: _RandomWalkNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_RandomWalkNFT *RandomWalkNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *RandomWalkNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _RandomWalkNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomWalkNFTApproval)
				if err := _RandomWalkNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_RandomWalkNFT *RandomWalkNFTFilterer) ParseApproval(log types.Log) (*RandomWalkNFTApproval, error) {
	event := new(RandomWalkNFTApproval)
	if err := _RandomWalkNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomWalkNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the RandomWalkNFT contract.
type RandomWalkNFTApprovalForAllIterator struct {
	Event *RandomWalkNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *RandomWalkNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomWalkNFTApprovalForAll)
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
		it.Event = new(RandomWalkNFTApprovalForAll)
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
func (it *RandomWalkNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomWalkNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomWalkNFTApprovalForAll represents a ApprovalForAll event raised by the RandomWalkNFT contract.
type RandomWalkNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_RandomWalkNFT *RandomWalkNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*RandomWalkNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _RandomWalkNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &RandomWalkNFTApprovalForAllIterator{contract: _RandomWalkNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_RandomWalkNFT *RandomWalkNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *RandomWalkNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _RandomWalkNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomWalkNFTApprovalForAll)
				if err := _RandomWalkNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_RandomWalkNFT *RandomWalkNFTFilterer) ParseApprovalForAll(log types.Log) (*RandomWalkNFTApprovalForAll, error) {
	event := new(RandomWalkNFTApprovalForAll)
	if err := _RandomWalkNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomWalkNFTMintEventIterator is returned from FilterMintEvent and is used to iterate over the raw logs and unpacked data for MintEvent events raised by the RandomWalkNFT contract.
type RandomWalkNFTMintEventIterator struct {
	Event *RandomWalkNFTMintEvent // Event containing the contract specifics and raw log

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
func (it *RandomWalkNFTMintEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomWalkNFTMintEvent)
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
		it.Event = new(RandomWalkNFTMintEvent)
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
func (it *RandomWalkNFTMintEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomWalkNFTMintEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomWalkNFTMintEvent represents a MintEvent event raised by the RandomWalkNFT contract.
type RandomWalkNFTMintEvent struct {
	TokenId *big.Int
	Owner   common.Address
	Seed    [32]byte
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMintEvent is a free log retrieval operation binding the contract event 0xad2bc79f659de022c64ef55c71f16d0cf125452ed5fc5757b2edc331f58565ec.
//
// Solidity: event MintEvent(uint256 indexed tokenId, address indexed owner, bytes32 seed, uint256 price)
func (_RandomWalkNFT *RandomWalkNFTFilterer) FilterMintEvent(opts *bind.FilterOpts, tokenId []*big.Int, owner []common.Address) (*RandomWalkNFTMintEventIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _RandomWalkNFT.contract.FilterLogs(opts, "MintEvent", tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &RandomWalkNFTMintEventIterator{contract: _RandomWalkNFT.contract, event: "MintEvent", logs: logs, sub: sub}, nil
}

// WatchMintEvent is a free log subscription operation binding the contract event 0xad2bc79f659de022c64ef55c71f16d0cf125452ed5fc5757b2edc331f58565ec.
//
// Solidity: event MintEvent(uint256 indexed tokenId, address indexed owner, bytes32 seed, uint256 price)
func (_RandomWalkNFT *RandomWalkNFTFilterer) WatchMintEvent(opts *bind.WatchOpts, sink chan<- *RandomWalkNFTMintEvent, tokenId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _RandomWalkNFT.contract.WatchLogs(opts, "MintEvent", tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomWalkNFTMintEvent)
				if err := _RandomWalkNFT.contract.UnpackLog(event, "MintEvent", log); err != nil {
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

// ParseMintEvent is a log parse operation binding the contract event 0xad2bc79f659de022c64ef55c71f16d0cf125452ed5fc5757b2edc331f58565ec.
//
// Solidity: event MintEvent(uint256 indexed tokenId, address indexed owner, bytes32 seed, uint256 price)
func (_RandomWalkNFT *RandomWalkNFTFilterer) ParseMintEvent(log types.Log) (*RandomWalkNFTMintEvent, error) {
	event := new(RandomWalkNFTMintEvent)
	if err := _RandomWalkNFT.contract.UnpackLog(event, "MintEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomWalkNFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RandomWalkNFT contract.
type RandomWalkNFTOwnershipTransferredIterator struct {
	Event *RandomWalkNFTOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RandomWalkNFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomWalkNFTOwnershipTransferred)
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
		it.Event = new(RandomWalkNFTOwnershipTransferred)
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
func (it *RandomWalkNFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomWalkNFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomWalkNFTOwnershipTransferred represents a OwnershipTransferred event raised by the RandomWalkNFT contract.
type RandomWalkNFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RandomWalkNFT *RandomWalkNFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RandomWalkNFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RandomWalkNFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RandomWalkNFTOwnershipTransferredIterator{contract: _RandomWalkNFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RandomWalkNFT *RandomWalkNFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RandomWalkNFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RandomWalkNFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomWalkNFTOwnershipTransferred)
				if err := _RandomWalkNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RandomWalkNFT *RandomWalkNFTFilterer) ParseOwnershipTransferred(log types.Log) (*RandomWalkNFTOwnershipTransferred, error) {
	event := new(RandomWalkNFTOwnershipTransferred)
	if err := _RandomWalkNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomWalkNFTTokenNameEventIterator is returned from FilterTokenNameEvent and is used to iterate over the raw logs and unpacked data for TokenNameEvent events raised by the RandomWalkNFT contract.
type RandomWalkNFTTokenNameEventIterator struct {
	Event *RandomWalkNFTTokenNameEvent // Event containing the contract specifics and raw log

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
func (it *RandomWalkNFTTokenNameEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomWalkNFTTokenNameEvent)
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
		it.Event = new(RandomWalkNFTTokenNameEvent)
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
func (it *RandomWalkNFTTokenNameEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomWalkNFTTokenNameEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomWalkNFTTokenNameEvent represents a TokenNameEvent event raised by the RandomWalkNFT contract.
type RandomWalkNFTTokenNameEvent struct {
	TokenId *big.Int
	NewName string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenNameEvent is a free log retrieval operation binding the contract event 0x8ad5e159ff95649c8a9f323ac5a457e741897cf44ce07dfce0e98b84ef9d5f12.
//
// Solidity: event TokenNameEvent(uint256 tokenId, string newName)
func (_RandomWalkNFT *RandomWalkNFTFilterer) FilterTokenNameEvent(opts *bind.FilterOpts) (*RandomWalkNFTTokenNameEventIterator, error) {

	logs, sub, err := _RandomWalkNFT.contract.FilterLogs(opts, "TokenNameEvent")
	if err != nil {
		return nil, err
	}
	return &RandomWalkNFTTokenNameEventIterator{contract: _RandomWalkNFT.contract, event: "TokenNameEvent", logs: logs, sub: sub}, nil
}

// WatchTokenNameEvent is a free log subscription operation binding the contract event 0x8ad5e159ff95649c8a9f323ac5a457e741897cf44ce07dfce0e98b84ef9d5f12.
//
// Solidity: event TokenNameEvent(uint256 tokenId, string newName)
func (_RandomWalkNFT *RandomWalkNFTFilterer) WatchTokenNameEvent(opts *bind.WatchOpts, sink chan<- *RandomWalkNFTTokenNameEvent) (event.Subscription, error) {

	logs, sub, err := _RandomWalkNFT.contract.WatchLogs(opts, "TokenNameEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomWalkNFTTokenNameEvent)
				if err := _RandomWalkNFT.contract.UnpackLog(event, "TokenNameEvent", log); err != nil {
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

// ParseTokenNameEvent is a log parse operation binding the contract event 0x8ad5e159ff95649c8a9f323ac5a457e741897cf44ce07dfce0e98b84ef9d5f12.
//
// Solidity: event TokenNameEvent(uint256 tokenId, string newName)
func (_RandomWalkNFT *RandomWalkNFTFilterer) ParseTokenNameEvent(log types.Log) (*RandomWalkNFTTokenNameEvent, error) {
	event := new(RandomWalkNFTTokenNameEvent)
	if err := _RandomWalkNFT.contract.UnpackLog(event, "TokenNameEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomWalkNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the RandomWalkNFT contract.
type RandomWalkNFTTransferIterator struct {
	Event *RandomWalkNFTTransfer // Event containing the contract specifics and raw log

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
func (it *RandomWalkNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomWalkNFTTransfer)
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
		it.Event = new(RandomWalkNFTTransfer)
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
func (it *RandomWalkNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomWalkNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomWalkNFTTransfer represents a Transfer event raised by the RandomWalkNFT contract.
type RandomWalkNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_RandomWalkNFT *RandomWalkNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*RandomWalkNFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _RandomWalkNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &RandomWalkNFTTransferIterator{contract: _RandomWalkNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_RandomWalkNFT *RandomWalkNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RandomWalkNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _RandomWalkNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomWalkNFTTransfer)
				if err := _RandomWalkNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_RandomWalkNFT *RandomWalkNFTFilterer) ParseTransfer(log types.Log) (*RandomWalkNFTTransfer, error) {
	event := new(RandomWalkNFTTransfer)
	if err := _RandomWalkNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomWalkNFTWithdrawalEventIterator is returned from FilterWithdrawalEvent and is used to iterate over the raw logs and unpacked data for WithdrawalEvent events raised by the RandomWalkNFT contract.
type RandomWalkNFTWithdrawalEventIterator struct {
	Event *RandomWalkNFTWithdrawalEvent // Event containing the contract specifics and raw log

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
func (it *RandomWalkNFTWithdrawalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomWalkNFTWithdrawalEvent)
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
		it.Event = new(RandomWalkNFTWithdrawalEvent)
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
func (it *RandomWalkNFTWithdrawalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomWalkNFTWithdrawalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomWalkNFTWithdrawalEvent represents a WithdrawalEvent event raised by the RandomWalkNFT contract.
type RandomWalkNFTWithdrawalEvent struct {
	TokenId     *big.Int
	Destination common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalEvent is a free log retrieval operation binding the contract event 0xa11b556ace4b11a5cae8675a293b51e8cde3a06387d34010861789dfd9e9abc7.
//
// Solidity: event WithdrawalEvent(uint256 indexed tokenId, address destination, uint256 amount)
func (_RandomWalkNFT *RandomWalkNFTFilterer) FilterWithdrawalEvent(opts *bind.FilterOpts, tokenId []*big.Int) (*RandomWalkNFTWithdrawalEventIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _RandomWalkNFT.contract.FilterLogs(opts, "WithdrawalEvent", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &RandomWalkNFTWithdrawalEventIterator{contract: _RandomWalkNFT.contract, event: "WithdrawalEvent", logs: logs, sub: sub}, nil
}

// WatchWithdrawalEvent is a free log subscription operation binding the contract event 0xa11b556ace4b11a5cae8675a293b51e8cde3a06387d34010861789dfd9e9abc7.
//
// Solidity: event WithdrawalEvent(uint256 indexed tokenId, address destination, uint256 amount)
func (_RandomWalkNFT *RandomWalkNFTFilterer) WatchWithdrawalEvent(opts *bind.WatchOpts, sink chan<- *RandomWalkNFTWithdrawalEvent, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _RandomWalkNFT.contract.WatchLogs(opts, "WithdrawalEvent", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomWalkNFTWithdrawalEvent)
				if err := _RandomWalkNFT.contract.UnpackLog(event, "WithdrawalEvent", log); err != nil {
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

// ParseWithdrawalEvent is a log parse operation binding the contract event 0xa11b556ace4b11a5cae8675a293b51e8cde3a06387d34010861789dfd9e9abc7.
//
// Solidity: event WithdrawalEvent(uint256 indexed tokenId, address destination, uint256 amount)
func (_RandomWalkNFT *RandomWalkNFTFilterer) ParseWithdrawalEvent(log types.Log) (*RandomWalkNFTWithdrawalEvent, error) {
	event := new(RandomWalkNFTWithdrawalEvent)
	if err := _RandomWalkNFT.contract.UnpackLog(event, "WithdrawalEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeCastMetaData contains all meta data concerning the SafeCast contract.
var SafeCastMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208845b917a644c445e474788dafb2895e18348825c919f69032b30088a13f102b64736f6c63430008130033",
}

// SafeCastABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeCastMetaData.ABI instead.
var SafeCastABI = SafeCastMetaData.ABI

// SafeCastBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeCastMetaData.Bin instead.
var SafeCastBin = SafeCastMetaData.Bin

// DeploySafeCast deploys a new Ethereum contract, binding an instance of SafeCast to it.
func DeploySafeCast(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeCast, error) {
	parsed, err := SafeCastMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeCastBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeCast{SafeCastCaller: SafeCastCaller{contract: contract}, SafeCastTransactor: SafeCastTransactor{contract: contract}, SafeCastFilterer: SafeCastFilterer{contract: contract}}, nil
}

// SafeCast is an auto generated Go binding around an Ethereum contract.
type SafeCast struct {
	SafeCastCaller     // Read-only binding to the contract
	SafeCastTransactor // Write-only binding to the contract
	SafeCastFilterer   // Log filterer for contract events
}

// SafeCastCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeCastCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeCastTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeCastFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeCastSession struct {
	Contract     *SafeCast         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeCastCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeCastCallerSession struct {
	Contract *SafeCastCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeCastTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeCastTransactorSession struct {
	Contract     *SafeCastTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeCastRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeCastRaw struct {
	Contract *SafeCast // Generic contract binding to access the raw methods on
}

// SafeCastCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeCastCallerRaw struct {
	Contract *SafeCastCaller // Generic read-only contract binding to access the raw methods on
}

// SafeCastTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeCastTransactorRaw struct {
	Contract *SafeCastTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeCast creates a new instance of SafeCast, bound to a specific deployed contract.
func NewSafeCast(address common.Address, backend bind.ContractBackend) (*SafeCast, error) {
	contract, err := bindSafeCast(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeCast{SafeCastCaller: SafeCastCaller{contract: contract}, SafeCastTransactor: SafeCastTransactor{contract: contract}, SafeCastFilterer: SafeCastFilterer{contract: contract}}, nil
}

// NewSafeCastCaller creates a new read-only instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastCaller(address common.Address, caller bind.ContractCaller) (*SafeCastCaller, error) {
	contract, err := bindSafeCast(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeCastCaller{contract: contract}, nil
}

// NewSafeCastTransactor creates a new write-only instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeCastTransactor, error) {
	contract, err := bindSafeCast(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeCastTransactor{contract: contract}, nil
}

// NewSafeCastFilterer creates a new log filterer instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeCastFilterer, error) {
	contract, err := bindSafeCast(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeCastFilterer{contract: contract}, nil
}

// bindSafeCast binds a generic wrapper to an already deployed contract.
func bindSafeCast(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeCastABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeCast *SafeCastRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeCast.Contract.SafeCastCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeCast *SafeCastRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeCast.Contract.SafeCastTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeCast *SafeCastRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeCast.Contract.SafeCastTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeCast *SafeCastCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeCast.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeCast *SafeCastTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeCast.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeCast *SafeCastTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeCast.Contract.contract.Transact(opts, method, params...)
}

// StringsMetaData contains all meta data concerning the Strings contract.
var StringsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212206b5ef8b0c3688c5123425dff80caa6213642336c7d6195a5211aa3f0ec616ac164736f6c63430008130033",
}

// StringsABI is the input ABI used to generate the binding from.
// Deprecated: Use StringsMetaData.ABI instead.
var StringsABI = StringsMetaData.ABI

// StringsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StringsMetaData.Bin instead.
var StringsBin = StringsMetaData.Bin

// DeployStrings deploys a new Ethereum contract, binding an instance of Strings to it.
func DeployStrings(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Strings, error) {
	parsed, err := StringsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StringsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Strings{StringsCaller: StringsCaller{contract: contract}, StringsTransactor: StringsTransactor{contract: contract}, StringsFilterer: StringsFilterer{contract: contract}}, nil
}

// Strings is an auto generated Go binding around an Ethereum contract.
type Strings struct {
	StringsCaller     // Read-only binding to the contract
	StringsTransactor // Write-only binding to the contract
	StringsFilterer   // Log filterer for contract events
}

// StringsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StringsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StringsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StringsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StringsSession struct {
	Contract     *Strings          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StringsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StringsCallerSession struct {
	Contract *StringsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StringsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StringsTransactorSession struct {
	Contract     *StringsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StringsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StringsRaw struct {
	Contract *Strings // Generic contract binding to access the raw methods on
}

// StringsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StringsCallerRaw struct {
	Contract *StringsCaller // Generic read-only contract binding to access the raw methods on
}

// StringsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StringsTransactorRaw struct {
	Contract *StringsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrings creates a new instance of Strings, bound to a specific deployed contract.
func NewStrings(address common.Address, backend bind.ContractBackend) (*Strings, error) {
	contract, err := bindStrings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Strings{StringsCaller: StringsCaller{contract: contract}, StringsTransactor: StringsTransactor{contract: contract}, StringsFilterer: StringsFilterer{contract: contract}}, nil
}

// NewStringsCaller creates a new read-only instance of Strings, bound to a specific deployed contract.
func NewStringsCaller(address common.Address, caller bind.ContractCaller) (*StringsCaller, error) {
	contract, err := bindStrings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StringsCaller{contract: contract}, nil
}

// NewStringsTransactor creates a new write-only instance of Strings, bound to a specific deployed contract.
func NewStringsTransactor(address common.Address, transactor bind.ContractTransactor) (*StringsTransactor, error) {
	contract, err := bindStrings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StringsTransactor{contract: contract}, nil
}

// NewStringsFilterer creates a new log filterer instance of Strings, bound to a specific deployed contract.
func NewStringsFilterer(address common.Address, filterer bind.ContractFilterer) (*StringsFilterer, error) {
	contract, err := bindStrings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StringsFilterer{contract: contract}, nil
}

// bindStrings binds a generic wrapper to an already deployed contract.
func bindStrings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StringsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strings *StringsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Strings.Contract.StringsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strings *StringsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strings.Contract.StringsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strings *StringsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strings.Contract.StringsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strings *StringsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Strings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strings *StringsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strings *StringsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strings.Contract.contract.Transact(opts, method, params...)
}
