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

// ERC20VotesCheckpoint is an auto generated low-level Go binding around an user-defined struct.
type ERC20VotesCheckpoint struct {
	FromBlock uint32
	Votes     *big.Int
}

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

// CosmicSignatureMetaData contains all meta data concerning the CosmicSignature contract.
var CosmicSignatureMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cosmicGameContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newURI\",\"type\":\"string\"}],\"name\":\"BaseURIEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"seed\",\"type\":\"bytes32\"}],\"name\":\"MintEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newURL\",\"type\":\"string\"}],\"name\":\"TokenGenerationScriptURLEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newName\",\"type\":\"string\"}],\"name\":\"TokenNameEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cosmicGameContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entropy\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"seeds\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTokenGenerationScriptURL\",\"type\":\"string\"}],\"name\":\"setTokenGenerationScriptURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setTokenName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenGenerationScriptURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenNames\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040526000600e556040518060400160405280600a81526020017f697066733a2f2f54424400000000000000000000000000000000000000000000815250601090816200004f91906200056a565b503480156200005d57600080fd5b506040516200491c3803806200491c8339818101604052810190620000839190620006bb565b6040518060400160405280600f81526020017f436f736d69635369676e617475726500000000000000000000000000000000008152506040518060400160405280600381526020017f435353000000000000000000000000000000000000000000000000000000000081525081600090816200010091906200056a565b5080600190816200011291906200056a565b50505062000135620001296200022260201b60201c565b6200022a60201b60201c565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603620001a7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200019e906200074e565b60405180910390fd5b42600143620001b791906200079f565b40604051602001620001cb92919062000856565b60405160208183030381529060405280519060200120600d819055508073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250505062000898565b600033905090565b6000600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200037257607f821691505b6020821081036200038857620003876200032a565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620003f27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620003b3565b620003fe8683620003b3565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006200044b620004456200043f8462000416565b62000420565b62000416565b9050919050565b6000819050919050565b62000467836200042a565b6200047f620004768262000452565b848454620003c0565b825550505050565b600090565b6200049662000487565b620004a38184846200045c565b505050565b5b81811015620004cb57620004bf6000826200048c565b600181019050620004a9565b5050565b601f8211156200051a57620004e4816200038e565b620004ef84620003a3565b81016020851015620004ff578190505b620005176200050e85620003a3565b830182620004a8565b50505b505050565b600082821c905092915050565b60006200053f600019846008026200051f565b1980831691505092915050565b60006200055a83836200052c565b9150826002028217905092915050565b6200057582620002f0565b67ffffffffffffffff811115620005915762000590620002fb565b5b6200059d825462000359565b620005aa828285620004cf565b600060209050601f831160018114620005e25760008415620005cd578287015190505b620005d985826200054c565b86555062000649565b601f198416620005f2866200038e565b60005b828110156200061c57848901518255600182019150602085019450602081019050620005f5565b868310156200063c578489015162000638601f8916826200052c565b8355505b6001600288020188555050505b505050505050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620006838262000656565b9050919050565b620006958162000676565b8114620006a157600080fd5b50565b600081519050620006b5816200068a565b92915050565b600060208284031215620006d457620006d362000651565b5b6000620006e484828501620006a4565b91505092915050565b600082825260208201905092915050565b7f5a65726f2d616464726573732077617320676976656e2e000000000000000000600082015250565b600062000736601783620006ed565b91506200074382620006fe565b602082019050919050565b60006020820190508181036000830152620007698162000727565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000620007ac8262000416565b9150620007b98362000416565b9250828203905081811115620007d457620007d362000770565b5b92915050565b7f6e65774e46540000000000000000000000000000000000000000000000000000600082015250565b600062000812600683620006ed565b91506200081f82620007da565b602082019050919050565b620008358162000416565b82525050565b6000819050919050565b62000850816200083b565b82525050565b60006060820190508181036000830152620008718162000803565b90506200088260208301856200082a565b62000891604083018462000845565b9392505050565b608051614061620008bb60003960008181610aa6015261105d01526140616000f3fe608060405234801561001057600080fd5b50600436106101cf5760003560e01c80636352211e11610104578063a22cb465116100a2578063de9d90bc11610071578063de9d90bc14610540578063e985e9c51461055e578063f0503e801461058e578063f2fde38b146105be576101cf565b8063a22cb465146104bc578063b88d4fde146104d8578063c87b56dd146104f4578063cdb0e89e14610524576101cf565b80638da5cb5b116100de5780638da5cb5b146104445780638e0354ec146104625780638e499bcf1461048057806395d89b411461049e576101cf565b80636352211e146103da57806370a082311461040a578063715018a61461043a576101cf565b8063310495ab1161017157806346519a181161014b57806346519a181461035457806347ce07cc146103705780634f6ccce71461038e57806355f804b3146103be576101cf565b8063310495ab146102d857806340c10f191461030857806342842e0e14610338576101cf565b8063095ea7b3116101ad578063095ea7b31461025257806318160ddd1461026e57806323b872dd1461028c5780632f745c59146102a8576101cf565b806301ffc9a7146101d457806306fdde0314610204578063081812fc14610222575b600080fd5b6101ee60048036038101906101e99190612779565b6105da565b6040516101fb91906127c1565b60405180910390f35b61020c610654565b604051610219919061286c565b60405180910390f35b61023c600480360381019061023791906128c4565b6106e6565b6040516102499190612932565b60405180910390f35b61026c60048036038101906102679190612979565b61076b565b005b610276610882565b60405161028391906129c8565b60405180910390f35b6102a660048036038101906102a191906129e3565b61088f565b005b6102c260048036038101906102bd9190612979565b6108ef565b6040516102cf91906129c8565b60405180910390f35b6102f260048036038101906102ed91906128c4565b610994565b6040516102ff919061286c565b60405180910390f35b610322600480360381019061031d9190612979565b610a34565b60405161032f91906129c8565b60405180910390f35b610352600480360381019061034d91906129e3565b610c1e565b005b61036e60048036038101906103699190612b6b565b610c3e565b005b610378610d04565b6040516103859190612bcd565b60405180910390f35b6103a860048036038101906103a391906128c4565b610d0a565b6040516103b591906129c8565b60405180910390f35b6103d860048036038101906103d39190612b6b565b610d7b565b005b6103f460048036038101906103ef91906128c4565b610e41565b6040516104019190612932565b60405180910390f35b610424600480360381019061041f9190612be8565b610ef2565b60405161043191906129c8565b60405180910390f35b610442610fa9565b005b61044c611031565b6040516104599190612932565b60405180910390f35b61046a61105b565b6040516104779190612932565b60405180910390f35b61048861107f565b60405161049591906129c8565b60405180910390f35b6104a6611085565b6040516104b3919061286c565b60405180910390f35b6104d660048036038101906104d19190612c41565b611117565b005b6104f260048036038101906104ed9190612d22565b611297565b005b61050e600480360381019061050991906128c4565b6112f9565b60405161051b919061286c565b60405180910390f35b61053e60048036038101906105399190612da5565b6113a0565b005b610548611492565b604051610555919061286c565b60405180910390f35b61057860048036038101906105739190612e01565b611520565b60405161058591906127c1565b60405180910390f35b6105a860048036038101906105a391906128c4565b6115b4565b6040516105b59190612bcd565b60405180910390f35b6105d860048036038101906105d39190612be8565b6115cc565b005b60007f780e9d63000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061064d575061064c826116c3565b5b9050919050565b60606000805461066390612e70565b80601f016020809104026020016040519081016040528092919081815260200182805461068f90612e70565b80156106dc5780601f106106b1576101008083540402835291602001916106dc565b820191906000526020600020905b8154815290600101906020018083116106bf57829003601f168201915b5050505050905090565b60006106f1826117a5565b610730576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161072790612f13565b60405180910390fd5b6004600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b600061077682610e41565b90508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036107e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107dd90612fa5565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16610805611811565b73ffffffffffffffffffffffffffffffffffffffff16148061083457506108338161082e611811565b611520565b5b610873576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161086a90613037565b60405180910390fd5b61087d8383611819565b505050565b6000600880549050905090565b6108a061089a611811565b826118d2565b6108df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d6906130c9565b60405180910390fd5b6108ea8383836119b0565b505050565b60006108fa83610ef2565b821061093b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109329061315b565b60405180910390fd5b600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600083815260200190815260200160002054905092915050565b600c60205280600052604060002060009150905080546109b390612e70565b80601f01602080910402602001604051908101604052809291908181526020018280546109df90612e70565b8015610a2c5780601f10610a0157610100808354040283529160200191610a2c565b820191906000526020600020905b815481529060010190602001808311610a0f57829003601f168201915b505050505081565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610aa4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a9b906131c7565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610ae3611811565b73ffffffffffffffffffffffffffffffffffffffff1614610b39576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b3090613259565b60405180910390fd5b6000600e5490506001600e6000828254610b5391906132a8565b92505081905550600d5442600143610b6b91906132dc565b408387604051602001610b82959493929190613310565b60405160208183030381529060405280519060200120600d81905550600d54600b600083815260200190815260200160002081905550610bc28482611c0b565b828473ffffffffffffffffffffffffffffffffffffffff16827fc646da88dc2b2526461a0ebb4326e2418ec0bea89496b632b7c9ee42fbfe1d4d600d54604051610c0c9190612bcd565b60405180910390a48091505092915050565b610c3983838360405180602001604052806000815250611297565b505050565b610c46611811565b73ffffffffffffffffffffffffffffffffffffffff16610c64611031565b73ffffffffffffffffffffffffffffffffffffffff1614610cba576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cb1906133af565b60405180910390fd5b8060109081610cc9919061357b565b507f0119741ee0f95fab26124262a82c3c0e9e1c7ff4bb33c6fba5f3b11c9b6d0bad81604051610cf9919061286c565b60405180910390a150565b600d5481565b6000610d14610882565b8210610d55576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d4c906136bf565b60405180910390fd5b60088281548110610d6957610d686136df565b5b90600052602060002001549050919050565b610d83611811565b73ffffffffffffffffffffffffffffffffffffffff16610da1611031565b73ffffffffffffffffffffffffffffffffffffffff1614610df7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dee906133af565b60405180910390fd5b80600f9081610e06919061357b565b507f2fc013f885e8a815b9d697da28bc143b4dced47528c41b46e2b35fd0f4be718c81604051610e36919061286c565b60405180910390a150565b6000806002600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610ee9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ee090613780565b60405180910390fd5b80915050919050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610f62576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f5990613812565b60405180910390fd5b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b610fb1611811565b73ffffffffffffffffffffffffffffffffffffffff16610fcf611031565b73ffffffffffffffffffffffffffffffffffffffff1614611025576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161101c906133af565b60405180910390fd5b61102f6000611dd8565b565b6000600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b7f000000000000000000000000000000000000000000000000000000000000000081565b600e5481565b60606001805461109490612e70565b80601f01602080910402602001604051908101604052809291908181526020018280546110c090612e70565b801561110d5780601f106110e25761010080835404028352916020019161110d565b820191906000526020600020905b8154815290600101906020018083116110f057829003601f168201915b5050505050905090565b61111f611811565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361118c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111839061387e565b60405180910390fd5b8060056000611199611811565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff16611246611811565b73ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c318360405161128b91906127c1565b60405180910390a35050565b6112a86112a2611811565b836118d2565b6112e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112de906130c9565b60405180910390fd5b6112f384848484611e9e565b50505050565b6060611304826117a5565b611343576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161133a90613910565b60405180910390fd5b600061134d611efa565b9050600081511161136d5760405180602001604052806000815250611398565b8061137784611f8c565b60405160200161138892919061396c565b6040516020818303038152906040525b915050919050565b6113b16113ab611811565b836118d2565b6113f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113e790613a02565b60405180910390fd5b602081511115611435576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161142c90613a6e565b60405180910390fd5b80600c60008481526020019081526020016000209081611455919061357b565b50817f8ad5e159ff95649c8a9f323ac5a457e741897cf44ce07dfce0e98b84ef9d5f1282604051611486919061286c565b60405180910390a25050565b6010805461149f90612e70565b80601f01602080910402602001604051908101604052809291908181526020018280546114cb90612e70565b80156115185780601f106114ed57610100808354040283529160200191611518565b820191906000526020600020905b8154815290600101906020018083116114fb57829003601f168201915b505050505081565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b600b6020528060005260406000206000915090505481565b6115d4611811565b73ffffffffffffffffffffffffffffffffffffffff166115f2611031565b73ffffffffffffffffffffffffffffffffffffffff1614611648576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161163f906133af565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036116b7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116ae90613b00565b60405180910390fd5b6116c081611dd8565b50565b60007f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061178e57507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b8061179e575061179d826120ec565b5b9050919050565b60008073ffffffffffffffffffffffffffffffffffffffff166002600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614159050919050565b600033905090565b816004600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff1661188c83610e41565b73ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b60006118dd826117a5565b61191c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161191390613b92565b60405180910390fd5b600061192783610e41565b90508073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16148061199657508373ffffffffffffffffffffffffffffffffffffffff1661197e846106e6565b73ffffffffffffffffffffffffffffffffffffffff16145b806119a757506119a68185611520565b5b91505092915050565b8273ffffffffffffffffffffffffffffffffffffffff166119d082610e41565b73ffffffffffffffffffffffffffffffffffffffff1614611a26576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a1d90613c24565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611a95576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a8c90613cb6565b60405180910390fd5b611aa0838383612156565b611aab600082611819565b6001600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611afb91906132dc565b925050819055506001600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611b5291906132a8565b92505081905550816002600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611c7a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c7190613d22565b60405180910390fd5b611c83816117a5565b15611cc3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611cba90613d8e565b60405180910390fd5b611ccf60008383612156565b6001600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611d1f91906132a8565b92505081905550816002600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a45050565b6000600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b611ea98484846119b0565b611eb584848484612268565b611ef4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611eeb90613e20565b60405180910390fd5b50505050565b6060600f8054611f0990612e70565b80601f0160208091040260200160405190810160405280929190818152602001828054611f3590612e70565b8015611f825780601f10611f5757610100808354040283529160200191611f82565b820191906000526020600020905b815481529060010190602001808311611f6557829003601f168201915b5050505050905090565b606060008203611fd3576040518060400160405280600181526020017f300000000000000000000000000000000000000000000000000000000000000081525090506120e7565b600082905060005b60008214612005578080611fee90613e40565b915050600a82611ffe9190613eb7565b9150611fdb565b60008167ffffffffffffffff81111561202157612020612a40565b5b6040519080825280601f01601f1916602001820160405280156120535781602001600182028036833780820191505090505b5090505b600085146120e05760018261206c91906132dc565b9150600a8561207b9190613ee8565b603061208791906132a8565b60f81b81838151811061209d5761209c6136df565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a856120d99190613eb7565b9450612057565b8093505050505b919050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b6121618383836123ef565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036121a35761219e816123f4565b6121e2565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16146121e1576121e0838261243d565b5b5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036122245761221f816125aa565b612263565b8273ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161461226257612261828261267b565b5b5b505050565b60006122898473ffffffffffffffffffffffffffffffffffffffff166126fa565b156123e2578373ffffffffffffffffffffffffffffffffffffffff1663150b7a026122b2611811565b8786866040518563ffffffff1660e01b81526004016122d49493929190613f6e565b6020604051808303816000875af192505050801561231057506040513d601f19601f8201168201806040525081019061230d9190613fcf565b60015b612392573d8060008114612340576040519150601f19603f3d011682016040523d82523d6000602084013e612345565b606091505b50600081510361238a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161238190613e20565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149150506123e7565b600190505b949350505050565b505050565b6008805490506009600083815260200190815260200160002081905550600881908060018154018082558091505060019003906000526020600020016000909190919091505550565b6000600161244a84610ef2565b61245491906132dc565b9050600060076000848152602001908152602001600020549050818114612539576000600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002054905080600660008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002081905550816007600083815260200190815260200160002081905550505b6007600084815260200190815260200160002060009055600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008381526020019081526020016000206000905550505050565b600060016008805490506125be91906132dc565b90506000600960008481526020019081526020016000205490506000600883815481106125ee576125ed6136df565b5b9060005260206000200154905080600883815481106126105761260f6136df565b5b90600052602060002001819055508160096000838152602001908152602001600020819055506009600085815260200190815260200160002060009055600880548061265f5761265e613ffc565b5b6001900381819060005260206000200160009055905550505050565b600061268683610ef2565b905081600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600083815260200190815260200160002081905550806007600084815260200190815260200160002081905550505050565b600080823b905060008111915050919050565b6000604051905090565b600080fd5b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b61275681612721565b811461276157600080fd5b50565b6000813590506127738161274d565b92915050565b60006020828403121561278f5761278e612717565b5b600061279d84828501612764565b91505092915050565b60008115159050919050565b6127bb816127a6565b82525050565b60006020820190506127d660008301846127b2565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156128165780820151818401526020810190506127fb565b60008484015250505050565b6000601f19601f8301169050919050565b600061283e826127dc565b61284881856127e7565b93506128588185602086016127f8565b61286181612822565b840191505092915050565b600060208201905081810360008301526128868184612833565b905092915050565b6000819050919050565b6128a18161288e565b81146128ac57600080fd5b50565b6000813590506128be81612898565b92915050565b6000602082840312156128da576128d9612717565b5b60006128e8848285016128af565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061291c826128f1565b9050919050565b61292c81612911565b82525050565b60006020820190506129476000830184612923565b92915050565b61295681612911565b811461296157600080fd5b50565b6000813590506129738161294d565b92915050565b600080604083850312156129905761298f612717565b5b600061299e85828601612964565b92505060206129af858286016128af565b9150509250929050565b6129c28161288e565b82525050565b60006020820190506129dd60008301846129b9565b92915050565b6000806000606084860312156129fc576129fb612717565b5b6000612a0a86828701612964565b9350506020612a1b86828701612964565b9250506040612a2c868287016128af565b9150509250925092565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b612a7882612822565b810181811067ffffffffffffffff82111715612a9757612a96612a40565b5b80604052505050565b6000612aaa61270d565b9050612ab68282612a6f565b919050565b600067ffffffffffffffff821115612ad657612ad5612a40565b5b612adf82612822565b9050602081019050919050565b82818337600083830152505050565b6000612b0e612b0984612abb565b612aa0565b905082815260208101848484011115612b2a57612b29612a3b565b5b612b35848285612aec565b509392505050565b600082601f830112612b5257612b51612a36565b5b8135612b62848260208601612afb565b91505092915050565b600060208284031215612b8157612b80612717565b5b600082013567ffffffffffffffff811115612b9f57612b9e61271c565b5b612bab84828501612b3d565b91505092915050565b6000819050919050565b612bc781612bb4565b82525050565b6000602082019050612be26000830184612bbe565b92915050565b600060208284031215612bfe57612bfd612717565b5b6000612c0c84828501612964565b91505092915050565b612c1e816127a6565b8114612c2957600080fd5b50565b600081359050612c3b81612c15565b92915050565b60008060408385031215612c5857612c57612717565b5b6000612c6685828601612964565b9250506020612c7785828601612c2c565b9150509250929050565b600067ffffffffffffffff821115612c9c57612c9b612a40565b5b612ca582612822565b9050602081019050919050565b6000612cc5612cc084612c81565b612aa0565b905082815260208101848484011115612ce157612ce0612a3b565b5b612cec848285612aec565b509392505050565b600082601f830112612d0957612d08612a36565b5b8135612d19848260208601612cb2565b91505092915050565b60008060008060808587031215612d3c57612d3b612717565b5b6000612d4a87828801612964565b9450506020612d5b87828801612964565b9350506040612d6c878288016128af565b925050606085013567ffffffffffffffff811115612d8d57612d8c61271c565b5b612d9987828801612cf4565b91505092959194509250565b60008060408385031215612dbc57612dbb612717565b5b6000612dca858286016128af565b925050602083013567ffffffffffffffff811115612deb57612dea61271c565b5b612df785828601612b3d565b9150509250929050565b60008060408385031215612e1857612e17612717565b5b6000612e2685828601612964565b9250506020612e3785828601612964565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680612e8857607f821691505b602082108103612e9b57612e9a612e41565b5b50919050565b7f4552433732313a20617070726f76656420717565727920666f72206e6f6e657860008201527f697374656e7420746f6b656e0000000000000000000000000000000000000000602082015250565b6000612efd602c836127e7565b9150612f0882612ea1565b604082019050919050565b60006020820190508181036000830152612f2c81612ef0565b9050919050565b7f4552433732313a20617070726f76616c20746f2063757272656e74206f776e6560008201527f7200000000000000000000000000000000000000000000000000000000000000602082015250565b6000612f8f6021836127e7565b9150612f9a82612f33565b604082019050919050565b60006020820190508181036000830152612fbe81612f82565b9050919050565b7f4552433732313a20617070726f76652063616c6c6572206973206e6f74206f7760008201527f6e6572206e6f7220617070726f76656420666f7220616c6c0000000000000000602082015250565b60006130216038836127e7565b915061302c82612fc5565b604082019050919050565b6000602082019050818103600083015261305081613014565b9050919050565b7f4552433732313a207472616e736665722063616c6c6572206973206e6f74206f60008201527f776e6572206e6f7220617070726f766564000000000000000000000000000000602082015250565b60006130b36031836127e7565b91506130be82613057565b604082019050919050565b600060208201905081810360008301526130e2816130a6565b9050919050565b7f455243373231456e756d657261626c653a206f776e657220696e646578206f7560008201527f74206f6620626f756e6473000000000000000000000000000000000000000000602082015250565b6000613145602b836127e7565b9150613150826130e9565b604082019050919050565b6000602082019050818103600083015261317481613138565b9050919050565b7f5a65726f2d616464726573732077617320676976656e2e000000000000000000600082015250565b60006131b16017836127e7565b91506131bc8261317b565b602082019050919050565b600060208201905081810360008301526131e0816131a4565b9050919050565b7f4f6e6c792074686520436f736d696347616d6520636f6e74726163742063616e60008201527f206d696e742e0000000000000000000000000000000000000000000000000000602082015250565b60006132436026836127e7565b915061324e826131e7565b604082019050919050565b6000602082019050818103600083015261327281613236565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006132b38261288e565b91506132be8361288e565b92508282019050808211156132d6576132d5613279565b5b92915050565b60006132e78261288e565b91506132f28361288e565b925082820390508181111561330a57613309613279565b5b92915050565b600060a0820190506133256000830188612bbe565b61333260208301876129b9565b61333f6040830186612bbe565b61334c60608301856129b9565b6133596080830184612923565b9695505050505050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b60006133996020836127e7565b91506133a482613363565b602082019050919050565b600060208201905081810360008301526133c88161338c565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026134317fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826133f4565b61343b86836133f4565b95508019841693508086168417925050509392505050565b6000819050919050565b600061347861347361346e8461288e565b613453565b61288e565b9050919050565b6000819050919050565b6134928361345d565b6134a661349e8261347f565b848454613401565b825550505050565b600090565b6134bb6134ae565b6134c6818484613489565b505050565b5b818110156134ea576134df6000826134b3565b6001810190506134cc565b5050565b601f82111561352f57613500816133cf565b613509846133e4565b81016020851015613518578190505b61352c613524856133e4565b8301826134cb565b50505b505050565b600082821c905092915050565b600061355260001984600802613534565b1980831691505092915050565b600061356b8383613541565b9150826002028217905092915050565b613584826127dc565b67ffffffffffffffff81111561359d5761359c612a40565b5b6135a78254612e70565b6135b28282856134ee565b600060209050601f8311600181146135e557600084156135d3578287015190505b6135dd858261355f565b865550613645565b601f1984166135f3866133cf565b60005b8281101561361b578489015182556001820191506020850194506020810190506135f6565b868310156136385784890151613634601f891682613541565b8355505b6001600288020188555050505b505050505050565b7f455243373231456e756d657261626c653a20676c6f62616c20696e646578206f60008201527f7574206f6620626f756e64730000000000000000000000000000000000000000602082015250565b60006136a9602c836127e7565b91506136b48261364d565b604082019050919050565b600060208201905081810360008301526136d88161369c565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4552433732313a206f776e657220717565727920666f72206e6f6e657869737460008201527f656e7420746f6b656e0000000000000000000000000000000000000000000000602082015250565b600061376a6029836127e7565b91506137758261370e565b604082019050919050565b600060208201905081810360008301526137998161375d565b9050919050565b7f4552433732313a2062616c616e636520717565727920666f7220746865207a6560008201527f726f206164647265737300000000000000000000000000000000000000000000602082015250565b60006137fc602a836127e7565b9150613807826137a0565b604082019050919050565b6000602082019050818103600083015261382b816137ef565b9050919050565b7f4552433732313a20617070726f766520746f2063616c6c657200000000000000600082015250565b60006138686019836127e7565b915061387382613832565b602082019050919050565b600060208201905081810360008301526138978161385b565b9050919050565b7f4552433732314d657461646174613a2055524920717565727920666f72206e6f60008201527f6e6578697374656e7420746f6b656e0000000000000000000000000000000000602082015250565b60006138fa602f836127e7565b91506139058261389e565b604082019050919050565b60006020820190508181036000830152613929816138ed565b9050919050565b600081905092915050565b6000613946826127dc565b6139508185613930565b93506139608185602086016127f8565b80840191505092915050565b6000613978828561393b565b9150613984828461393b565b91508190509392505050565b7f736574546f6b656e4e616d652063616c6c6572206973206e6f74206f776e657260008201527f206e6f7220617070726f7665642e000000000000000000000000000000000000602082015250565b60006139ec602e836127e7565b91506139f782613990565b604082019050919050565b60006020820190508181036000830152613a1b816139df565b9050919050565b7f546f6b656e206e616d6520697320746f6f206c6f6e672e000000000000000000600082015250565b6000613a586017836127e7565b9150613a6382613a22565b602082019050919050565b60006020820190508181036000830152613a8781613a4b565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000613aea6026836127e7565b9150613af582613a8e565b604082019050919050565b60006020820190508181036000830152613b1981613add565b9050919050565b7f4552433732313a206f70657261746f7220717565727920666f72206e6f6e657860008201527f697374656e7420746f6b656e0000000000000000000000000000000000000000602082015250565b6000613b7c602c836127e7565b9150613b8782613b20565b604082019050919050565b60006020820190508181036000830152613bab81613b6f565b9050919050565b7f4552433732313a207472616e73666572206f6620746f6b656e2074686174206960008201527f73206e6f74206f776e0000000000000000000000000000000000000000000000602082015250565b6000613c0e6029836127e7565b9150613c1982613bb2565b604082019050919050565b60006020820190508181036000830152613c3d81613c01565b9050919050565b7f4552433732313a207472616e7366657220746f20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b6000613ca06024836127e7565b9150613cab82613c44565b604082019050919050565b60006020820190508181036000830152613ccf81613c93565b9050919050565b7f4552433732313a206d696e7420746f20746865207a65726f2061646472657373600082015250565b6000613d0c6020836127e7565b9150613d1782613cd6565b602082019050919050565b60006020820190508181036000830152613d3b81613cff565b9050919050565b7f4552433732313a20746f6b656e20616c7265616479206d696e74656400000000600082015250565b6000613d78601c836127e7565b9150613d8382613d42565b602082019050919050565b60006020820190508181036000830152613da781613d6b565b9050919050565b7f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560008201527f63656976657220696d706c656d656e7465720000000000000000000000000000602082015250565b6000613e0a6032836127e7565b9150613e1582613dae565b604082019050919050565b60006020820190508181036000830152613e3981613dfd565b9050919050565b6000613e4b8261288e565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203613e7d57613e7c613279565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000613ec28261288e565b9150613ecd8361288e565b925082613edd57613edc613e88565b5b828204905092915050565b6000613ef38261288e565b9150613efe8361288e565b925082613f0e57613f0d613e88565b5b828206905092915050565b600081519050919050565b600082825260208201905092915050565b6000613f4082613f19565b613f4a8185613f24565b9350613f5a8185602086016127f8565b613f6381612822565b840191505092915050565b6000608082019050613f836000830187612923565b613f906020830186612923565b613f9d60408301856129b9565b8181036060830152613faf8184613f35565b905095945050505050565b600081519050613fc98161274d565b92915050565b600060208284031215613fe557613fe4612717565b5b6000613ff384828501613fba565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea2646970667358221220f43fee36f14a57721ad7fcf57597308b422a43f901324eb3f60f82418c79e78964736f6c63430008130033",
}

// CosmicSignatureABI is the input ABI used to generate the binding from.
// Deprecated: Use CosmicSignatureMetaData.ABI instead.
var CosmicSignatureABI = CosmicSignatureMetaData.ABI

// CosmicSignatureBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CosmicSignatureMetaData.Bin instead.
var CosmicSignatureBin = CosmicSignatureMetaData.Bin

// DeployCosmicSignature deploys a new Ethereum contract, binding an instance of CosmicSignature to it.
func DeployCosmicSignature(auth *bind.TransactOpts, backend bind.ContractBackend, _cosmicGameContract common.Address) (common.Address, *types.Transaction, *CosmicSignature, error) {
	parsed, err := CosmicSignatureMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CosmicSignatureBin), backend, _cosmicGameContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CosmicSignature{CosmicSignatureCaller: CosmicSignatureCaller{contract: contract}, CosmicSignatureTransactor: CosmicSignatureTransactor{contract: contract}, CosmicSignatureFilterer: CosmicSignatureFilterer{contract: contract}}, nil
}

// CosmicSignature is an auto generated Go binding around an Ethereum contract.
type CosmicSignature struct {
	CosmicSignatureCaller     // Read-only binding to the contract
	CosmicSignatureTransactor // Write-only binding to the contract
	CosmicSignatureFilterer   // Log filterer for contract events
}

// CosmicSignatureCaller is an auto generated read-only Go binding around an Ethereum contract.
type CosmicSignatureCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmicSignatureTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CosmicSignatureTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmicSignatureFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CosmicSignatureFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmicSignatureSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CosmicSignatureSession struct {
	Contract     *CosmicSignature  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CosmicSignatureCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CosmicSignatureCallerSession struct {
	Contract *CosmicSignatureCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CosmicSignatureTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CosmicSignatureTransactorSession struct {
	Contract     *CosmicSignatureTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CosmicSignatureRaw is an auto generated low-level Go binding around an Ethereum contract.
type CosmicSignatureRaw struct {
	Contract *CosmicSignature // Generic contract binding to access the raw methods on
}

// CosmicSignatureCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CosmicSignatureCallerRaw struct {
	Contract *CosmicSignatureCaller // Generic read-only contract binding to access the raw methods on
}

// CosmicSignatureTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CosmicSignatureTransactorRaw struct {
	Contract *CosmicSignatureTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCosmicSignature creates a new instance of CosmicSignature, bound to a specific deployed contract.
func NewCosmicSignature(address common.Address, backend bind.ContractBackend) (*CosmicSignature, error) {
	contract, err := bindCosmicSignature(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CosmicSignature{CosmicSignatureCaller: CosmicSignatureCaller{contract: contract}, CosmicSignatureTransactor: CosmicSignatureTransactor{contract: contract}, CosmicSignatureFilterer: CosmicSignatureFilterer{contract: contract}}, nil
}

// NewCosmicSignatureCaller creates a new read-only instance of CosmicSignature, bound to a specific deployed contract.
func NewCosmicSignatureCaller(address common.Address, caller bind.ContractCaller) (*CosmicSignatureCaller, error) {
	contract, err := bindCosmicSignature(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureCaller{contract: contract}, nil
}

// NewCosmicSignatureTransactor creates a new write-only instance of CosmicSignature, bound to a specific deployed contract.
func NewCosmicSignatureTransactor(address common.Address, transactor bind.ContractTransactor) (*CosmicSignatureTransactor, error) {
	contract, err := bindCosmicSignature(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureTransactor{contract: contract}, nil
}

// NewCosmicSignatureFilterer creates a new log filterer instance of CosmicSignature, bound to a specific deployed contract.
func NewCosmicSignatureFilterer(address common.Address, filterer bind.ContractFilterer) (*CosmicSignatureFilterer, error) {
	contract, err := bindCosmicSignature(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureFilterer{contract: contract}, nil
}

// bindCosmicSignature binds a generic wrapper to an already deployed contract.
func bindCosmicSignature(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CosmicSignatureABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CosmicSignature *CosmicSignatureRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CosmicSignature.Contract.CosmicSignatureCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CosmicSignature *CosmicSignatureRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicSignature.Contract.CosmicSignatureTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CosmicSignature *CosmicSignatureRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CosmicSignature.Contract.CosmicSignatureTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CosmicSignature *CosmicSignatureCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CosmicSignature.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CosmicSignature *CosmicSignatureTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicSignature.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CosmicSignature *CosmicSignatureTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CosmicSignature.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CosmicSignature *CosmicSignatureCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignature.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CosmicSignature *CosmicSignatureSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CosmicSignature.Contract.BalanceOf(&_CosmicSignature.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CosmicSignature *CosmicSignatureCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CosmicSignature.Contract.BalanceOf(&_CosmicSignature.CallOpts, owner)
}

// CosmicGameContract is a free data retrieval call binding the contract method 0x8e0354ec.
//
// Solidity: function cosmicGameContract() view returns(address)
func (_CosmicSignature *CosmicSignatureCaller) CosmicGameContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignature.contract.Call(opts, &out, "cosmicGameContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CosmicGameContract is a free data retrieval call binding the contract method 0x8e0354ec.
//
// Solidity: function cosmicGameContract() view returns(address)
func (_CosmicSignature *CosmicSignatureSession) CosmicGameContract() (common.Address, error) {
	return _CosmicSignature.Contract.CosmicGameContract(&_CosmicSignature.CallOpts)
}

// CosmicGameContract is a free data retrieval call binding the contract method 0x8e0354ec.
//
// Solidity: function cosmicGameContract() view returns(address)
func (_CosmicSignature *CosmicSignatureCallerSession) CosmicGameContract() (common.Address, error) {
	return _CosmicSignature.Contract.CosmicGameContract(&_CosmicSignature.CallOpts)
}

// Entropy is a free data retrieval call binding the contract method 0x47ce07cc.
//
// Solidity: function entropy() view returns(bytes32)
func (_CosmicSignature *CosmicSignatureCaller) Entropy(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CosmicSignature.contract.Call(opts, &out, "entropy")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Entropy is a free data retrieval call binding the contract method 0x47ce07cc.
//
// Solidity: function entropy() view returns(bytes32)
func (_CosmicSignature *CosmicSignatureSession) Entropy() ([32]byte, error) {
	return _CosmicSignature.Contract.Entropy(&_CosmicSignature.CallOpts)
}

// Entropy is a free data retrieval call binding the contract method 0x47ce07cc.
//
// Solidity: function entropy() view returns(bytes32)
func (_CosmicSignature *CosmicSignatureCallerSession) Entropy() ([32]byte, error) {
	return _CosmicSignature.Contract.Entropy(&_CosmicSignature.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CosmicSignature *CosmicSignatureCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignature.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CosmicSignature *CosmicSignatureSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _CosmicSignature.Contract.GetApproved(&_CosmicSignature.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CosmicSignature *CosmicSignatureCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _CosmicSignature.Contract.GetApproved(&_CosmicSignature.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CosmicSignature *CosmicSignatureCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _CosmicSignature.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CosmicSignature *CosmicSignatureSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _CosmicSignature.Contract.IsApprovedForAll(&_CosmicSignature.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CosmicSignature *CosmicSignatureCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _CosmicSignature.Contract.IsApprovedForAll(&_CosmicSignature.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CosmicSignature *CosmicSignatureCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CosmicSignature.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CosmicSignature *CosmicSignatureSession) Name() (string, error) {
	return _CosmicSignature.Contract.Name(&_CosmicSignature.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CosmicSignature *CosmicSignatureCallerSession) Name() (string, error) {
	return _CosmicSignature.Contract.Name(&_CosmicSignature.CallOpts)
}

// NumTokens is a free data retrieval call binding the contract method 0x8e499bcf.
//
// Solidity: function numTokens() view returns(uint256)
func (_CosmicSignature *CosmicSignatureCaller) NumTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignature.contract.Call(opts, &out, "numTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumTokens is a free data retrieval call binding the contract method 0x8e499bcf.
//
// Solidity: function numTokens() view returns(uint256)
func (_CosmicSignature *CosmicSignatureSession) NumTokens() (*big.Int, error) {
	return _CosmicSignature.Contract.NumTokens(&_CosmicSignature.CallOpts)
}

// NumTokens is a free data retrieval call binding the contract method 0x8e499bcf.
//
// Solidity: function numTokens() view returns(uint256)
func (_CosmicSignature *CosmicSignatureCallerSession) NumTokens() (*big.Int, error) {
	return _CosmicSignature.Contract.NumTokens(&_CosmicSignature.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CosmicSignature *CosmicSignatureCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignature.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CosmicSignature *CosmicSignatureSession) Owner() (common.Address, error) {
	return _CosmicSignature.Contract.Owner(&_CosmicSignature.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CosmicSignature *CosmicSignatureCallerSession) Owner() (common.Address, error) {
	return _CosmicSignature.Contract.Owner(&_CosmicSignature.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CosmicSignature *CosmicSignatureCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignature.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CosmicSignature *CosmicSignatureSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _CosmicSignature.Contract.OwnerOf(&_CosmicSignature.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CosmicSignature *CosmicSignatureCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _CosmicSignature.Contract.OwnerOf(&_CosmicSignature.CallOpts, tokenId)
}

// Seeds is a free data retrieval call binding the contract method 0xf0503e80.
//
// Solidity: function seeds(uint256 ) view returns(bytes32)
func (_CosmicSignature *CosmicSignatureCaller) Seeds(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _CosmicSignature.contract.Call(opts, &out, "seeds", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Seeds is a free data retrieval call binding the contract method 0xf0503e80.
//
// Solidity: function seeds(uint256 ) view returns(bytes32)
func (_CosmicSignature *CosmicSignatureSession) Seeds(arg0 *big.Int) ([32]byte, error) {
	return _CosmicSignature.Contract.Seeds(&_CosmicSignature.CallOpts, arg0)
}

// Seeds is a free data retrieval call binding the contract method 0xf0503e80.
//
// Solidity: function seeds(uint256 ) view returns(bytes32)
func (_CosmicSignature *CosmicSignatureCallerSession) Seeds(arg0 *big.Int) ([32]byte, error) {
	return _CosmicSignature.Contract.Seeds(&_CosmicSignature.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CosmicSignature *CosmicSignatureCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CosmicSignature.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CosmicSignature *CosmicSignatureSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CosmicSignature.Contract.SupportsInterface(&_CosmicSignature.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CosmicSignature *CosmicSignatureCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CosmicSignature.Contract.SupportsInterface(&_CosmicSignature.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CosmicSignature *CosmicSignatureCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CosmicSignature.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CosmicSignature *CosmicSignatureSession) Symbol() (string, error) {
	return _CosmicSignature.Contract.Symbol(&_CosmicSignature.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CosmicSignature *CosmicSignatureCallerSession) Symbol() (string, error) {
	return _CosmicSignature.Contract.Symbol(&_CosmicSignature.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_CosmicSignature *CosmicSignatureCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignature.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_CosmicSignature *CosmicSignatureSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _CosmicSignature.Contract.TokenByIndex(&_CosmicSignature.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_CosmicSignature *CosmicSignatureCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _CosmicSignature.Contract.TokenByIndex(&_CosmicSignature.CallOpts, index)
}

// TokenGenerationScriptURL is a free data retrieval call binding the contract method 0xde9d90bc.
//
// Solidity: function tokenGenerationScriptURL() view returns(string)
func (_CosmicSignature *CosmicSignatureCaller) TokenGenerationScriptURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CosmicSignature.contract.Call(opts, &out, "tokenGenerationScriptURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenGenerationScriptURL is a free data retrieval call binding the contract method 0xde9d90bc.
//
// Solidity: function tokenGenerationScriptURL() view returns(string)
func (_CosmicSignature *CosmicSignatureSession) TokenGenerationScriptURL() (string, error) {
	return _CosmicSignature.Contract.TokenGenerationScriptURL(&_CosmicSignature.CallOpts)
}

// TokenGenerationScriptURL is a free data retrieval call binding the contract method 0xde9d90bc.
//
// Solidity: function tokenGenerationScriptURL() view returns(string)
func (_CosmicSignature *CosmicSignatureCallerSession) TokenGenerationScriptURL() (string, error) {
	return _CosmicSignature.Contract.TokenGenerationScriptURL(&_CosmicSignature.CallOpts)
}

// TokenNames is a free data retrieval call binding the contract method 0x310495ab.
//
// Solidity: function tokenNames(uint256 ) view returns(string)
func (_CosmicSignature *CosmicSignatureCaller) TokenNames(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _CosmicSignature.contract.Call(opts, &out, "tokenNames", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenNames is a free data retrieval call binding the contract method 0x310495ab.
//
// Solidity: function tokenNames(uint256 ) view returns(string)
func (_CosmicSignature *CosmicSignatureSession) TokenNames(arg0 *big.Int) (string, error) {
	return _CosmicSignature.Contract.TokenNames(&_CosmicSignature.CallOpts, arg0)
}

// TokenNames is a free data retrieval call binding the contract method 0x310495ab.
//
// Solidity: function tokenNames(uint256 ) view returns(string)
func (_CosmicSignature *CosmicSignatureCallerSession) TokenNames(arg0 *big.Int) (string, error) {
	return _CosmicSignature.Contract.TokenNames(&_CosmicSignature.CallOpts, arg0)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_CosmicSignature *CosmicSignatureCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignature.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_CosmicSignature *CosmicSignatureSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _CosmicSignature.Contract.TokenOfOwnerByIndex(&_CosmicSignature.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_CosmicSignature *CosmicSignatureCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _CosmicSignature.Contract.TokenOfOwnerByIndex(&_CosmicSignature.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CosmicSignature *CosmicSignatureCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _CosmicSignature.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CosmicSignature *CosmicSignatureSession) TokenURI(tokenId *big.Int) (string, error) {
	return _CosmicSignature.Contract.TokenURI(&_CosmicSignature.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CosmicSignature *CosmicSignatureCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _CosmicSignature.Contract.TokenURI(&_CosmicSignature.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CosmicSignature *CosmicSignatureCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignature.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CosmicSignature *CosmicSignatureSession) TotalSupply() (*big.Int, error) {
	return _CosmicSignature.Contract.TotalSupply(&_CosmicSignature.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CosmicSignature *CosmicSignatureCallerSession) TotalSupply() (*big.Int, error) {
	return _CosmicSignature.Contract.TotalSupply(&_CosmicSignature.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CosmicSignature *CosmicSignatureTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CosmicSignature.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CosmicSignature *CosmicSignatureSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CosmicSignature.Contract.Approve(&_CosmicSignature.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CosmicSignature *CosmicSignatureTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CosmicSignature.Contract.Approve(&_CosmicSignature.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address owner, uint256 roundNum) returns(uint256)
func (_CosmicSignature *CosmicSignatureTransactor) Mint(opts *bind.TransactOpts, owner common.Address, roundNum *big.Int) (*types.Transaction, error) {
	return _CosmicSignature.contract.Transact(opts, "mint", owner, roundNum)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address owner, uint256 roundNum) returns(uint256)
func (_CosmicSignature *CosmicSignatureSession) Mint(owner common.Address, roundNum *big.Int) (*types.Transaction, error) {
	return _CosmicSignature.Contract.Mint(&_CosmicSignature.TransactOpts, owner, roundNum)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address owner, uint256 roundNum) returns(uint256)
func (_CosmicSignature *CosmicSignatureTransactorSession) Mint(owner common.Address, roundNum *big.Int) (*types.Transaction, error) {
	return _CosmicSignature.Contract.Mint(&_CosmicSignature.TransactOpts, owner, roundNum)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CosmicSignature *CosmicSignatureTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicSignature.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CosmicSignature *CosmicSignatureSession) RenounceOwnership() (*types.Transaction, error) {
	return _CosmicSignature.Contract.RenounceOwnership(&_CosmicSignature.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CosmicSignature *CosmicSignatureTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CosmicSignature.Contract.RenounceOwnership(&_CosmicSignature.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CosmicSignature *CosmicSignatureTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CosmicSignature.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CosmicSignature *CosmicSignatureSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CosmicSignature.Contract.SafeTransferFrom(&_CosmicSignature.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CosmicSignature *CosmicSignatureTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CosmicSignature.Contract.SafeTransferFrom(&_CosmicSignature.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_CosmicSignature *CosmicSignatureTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _CosmicSignature.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_CosmicSignature *CosmicSignatureSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _CosmicSignature.Contract.SafeTransferFrom0(&_CosmicSignature.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_CosmicSignature *CosmicSignatureTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _CosmicSignature.Contract.SafeTransferFrom0(&_CosmicSignature.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CosmicSignature *CosmicSignatureTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _CosmicSignature.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CosmicSignature *CosmicSignatureSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CosmicSignature.Contract.SetApprovalForAll(&_CosmicSignature.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CosmicSignature *CosmicSignatureTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CosmicSignature.Contract.SetApprovalForAll(&_CosmicSignature.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_CosmicSignature *CosmicSignatureTransactor) SetBaseURI(opts *bind.TransactOpts, baseURI string) (*types.Transaction, error) {
	return _CosmicSignature.contract.Transact(opts, "setBaseURI", baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_CosmicSignature *CosmicSignatureSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _CosmicSignature.Contract.SetBaseURI(&_CosmicSignature.TransactOpts, baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_CosmicSignature *CosmicSignatureTransactorSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _CosmicSignature.Contract.SetBaseURI(&_CosmicSignature.TransactOpts, baseURI)
}

// SetTokenGenerationScriptURL is a paid mutator transaction binding the contract method 0x46519a18.
//
// Solidity: function setTokenGenerationScriptURL(string newTokenGenerationScriptURL) returns()
func (_CosmicSignature *CosmicSignatureTransactor) SetTokenGenerationScriptURL(opts *bind.TransactOpts, newTokenGenerationScriptURL string) (*types.Transaction, error) {
	return _CosmicSignature.contract.Transact(opts, "setTokenGenerationScriptURL", newTokenGenerationScriptURL)
}

// SetTokenGenerationScriptURL is a paid mutator transaction binding the contract method 0x46519a18.
//
// Solidity: function setTokenGenerationScriptURL(string newTokenGenerationScriptURL) returns()
func (_CosmicSignature *CosmicSignatureSession) SetTokenGenerationScriptURL(newTokenGenerationScriptURL string) (*types.Transaction, error) {
	return _CosmicSignature.Contract.SetTokenGenerationScriptURL(&_CosmicSignature.TransactOpts, newTokenGenerationScriptURL)
}

// SetTokenGenerationScriptURL is a paid mutator transaction binding the contract method 0x46519a18.
//
// Solidity: function setTokenGenerationScriptURL(string newTokenGenerationScriptURL) returns()
func (_CosmicSignature *CosmicSignatureTransactorSession) SetTokenGenerationScriptURL(newTokenGenerationScriptURL string) (*types.Transaction, error) {
	return _CosmicSignature.Contract.SetTokenGenerationScriptURL(&_CosmicSignature.TransactOpts, newTokenGenerationScriptURL)
}

// SetTokenName is a paid mutator transaction binding the contract method 0xcdb0e89e.
//
// Solidity: function setTokenName(uint256 tokenId, string name) returns()
func (_CosmicSignature *CosmicSignatureTransactor) SetTokenName(opts *bind.TransactOpts, tokenId *big.Int, name string) (*types.Transaction, error) {
	return _CosmicSignature.contract.Transact(opts, "setTokenName", tokenId, name)
}

// SetTokenName is a paid mutator transaction binding the contract method 0xcdb0e89e.
//
// Solidity: function setTokenName(uint256 tokenId, string name) returns()
func (_CosmicSignature *CosmicSignatureSession) SetTokenName(tokenId *big.Int, name string) (*types.Transaction, error) {
	return _CosmicSignature.Contract.SetTokenName(&_CosmicSignature.TransactOpts, tokenId, name)
}

// SetTokenName is a paid mutator transaction binding the contract method 0xcdb0e89e.
//
// Solidity: function setTokenName(uint256 tokenId, string name) returns()
func (_CosmicSignature *CosmicSignatureTransactorSession) SetTokenName(tokenId *big.Int, name string) (*types.Transaction, error) {
	return _CosmicSignature.Contract.SetTokenName(&_CosmicSignature.TransactOpts, tokenId, name)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CosmicSignature *CosmicSignatureTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CosmicSignature.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CosmicSignature *CosmicSignatureSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CosmicSignature.Contract.TransferFrom(&_CosmicSignature.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CosmicSignature *CosmicSignatureTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CosmicSignature.Contract.TransferFrom(&_CosmicSignature.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CosmicSignature *CosmicSignatureTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CosmicSignature.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CosmicSignature *CosmicSignatureSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CosmicSignature.Contract.TransferOwnership(&_CosmicSignature.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CosmicSignature *CosmicSignatureTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CosmicSignature.Contract.TransferOwnership(&_CosmicSignature.TransactOpts, newOwner)
}

// CosmicSignatureApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CosmicSignature contract.
type CosmicSignatureApprovalIterator struct {
	Event *CosmicSignatureApproval // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureApproval)
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
		it.Event = new(CosmicSignatureApproval)
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
func (it *CosmicSignatureApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureApproval represents a Approval event raised by the CosmicSignature contract.
type CosmicSignatureApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CosmicSignature *CosmicSignatureFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*CosmicSignatureApprovalIterator, error) {

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

	logs, sub, err := _CosmicSignature.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureApprovalIterator{contract: _CosmicSignature.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CosmicSignature *CosmicSignatureFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CosmicSignatureApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _CosmicSignature.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureApproval)
				if err := _CosmicSignature.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_CosmicSignature *CosmicSignatureFilterer) ParseApproval(log types.Log) (*CosmicSignatureApproval, error) {
	event := new(CosmicSignatureApproval)
	if err := _CosmicSignature.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the CosmicSignature contract.
type CosmicSignatureApprovalForAllIterator struct {
	Event *CosmicSignatureApprovalForAll // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureApprovalForAll)
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
		it.Event = new(CosmicSignatureApprovalForAll)
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
func (it *CosmicSignatureApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureApprovalForAll represents a ApprovalForAll event raised by the CosmicSignature contract.
type CosmicSignatureApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CosmicSignature *CosmicSignatureFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*CosmicSignatureApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CosmicSignature.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureApprovalForAllIterator{contract: _CosmicSignature.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CosmicSignature *CosmicSignatureFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CosmicSignatureApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CosmicSignature.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureApprovalForAll)
				if err := _CosmicSignature.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_CosmicSignature *CosmicSignatureFilterer) ParseApprovalForAll(log types.Log) (*CosmicSignatureApprovalForAll, error) {
	event := new(CosmicSignatureApprovalForAll)
	if err := _CosmicSignature.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureBaseURIEventIterator is returned from FilterBaseURIEvent and is used to iterate over the raw logs and unpacked data for BaseURIEvent events raised by the CosmicSignature contract.
type CosmicSignatureBaseURIEventIterator struct {
	Event *CosmicSignatureBaseURIEvent // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureBaseURIEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureBaseURIEvent)
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
		it.Event = new(CosmicSignatureBaseURIEvent)
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
func (it *CosmicSignatureBaseURIEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureBaseURIEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureBaseURIEvent represents a BaseURIEvent event raised by the CosmicSignature contract.
type CosmicSignatureBaseURIEvent struct {
	NewURI string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBaseURIEvent is a free log retrieval operation binding the contract event 0x2fc013f885e8a815b9d697da28bc143b4dced47528c41b46e2b35fd0f4be718c.
//
// Solidity: event BaseURIEvent(string newURI)
func (_CosmicSignature *CosmicSignatureFilterer) FilterBaseURIEvent(opts *bind.FilterOpts) (*CosmicSignatureBaseURIEventIterator, error) {

	logs, sub, err := _CosmicSignature.contract.FilterLogs(opts, "BaseURIEvent")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureBaseURIEventIterator{contract: _CosmicSignature.contract, event: "BaseURIEvent", logs: logs, sub: sub}, nil
}

// WatchBaseURIEvent is a free log subscription operation binding the contract event 0x2fc013f885e8a815b9d697da28bc143b4dced47528c41b46e2b35fd0f4be718c.
//
// Solidity: event BaseURIEvent(string newURI)
func (_CosmicSignature *CosmicSignatureFilterer) WatchBaseURIEvent(opts *bind.WatchOpts, sink chan<- *CosmicSignatureBaseURIEvent) (event.Subscription, error) {

	logs, sub, err := _CosmicSignature.contract.WatchLogs(opts, "BaseURIEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureBaseURIEvent)
				if err := _CosmicSignature.contract.UnpackLog(event, "BaseURIEvent", log); err != nil {
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

// ParseBaseURIEvent is a log parse operation binding the contract event 0x2fc013f885e8a815b9d697da28bc143b4dced47528c41b46e2b35fd0f4be718c.
//
// Solidity: event BaseURIEvent(string newURI)
func (_CosmicSignature *CosmicSignatureFilterer) ParseBaseURIEvent(log types.Log) (*CosmicSignatureBaseURIEvent, error) {
	event := new(CosmicSignatureBaseURIEvent)
	if err := _CosmicSignature.contract.UnpackLog(event, "BaseURIEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureMintEventIterator is returned from FilterMintEvent and is used to iterate over the raw logs and unpacked data for MintEvent events raised by the CosmicSignature contract.
type CosmicSignatureMintEventIterator struct {
	Event *CosmicSignatureMintEvent // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureMintEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureMintEvent)
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
		it.Event = new(CosmicSignatureMintEvent)
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
func (it *CosmicSignatureMintEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureMintEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureMintEvent represents a MintEvent event raised by the CosmicSignature contract.
type CosmicSignatureMintEvent struct {
	TokenId  *big.Int
	Owner    common.Address
	RoundNum *big.Int
	Seed     [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMintEvent is a free log retrieval operation binding the contract event 0xc646da88dc2b2526461a0ebb4326e2418ec0bea89496b632b7c9ee42fbfe1d4d.
//
// Solidity: event MintEvent(uint256 indexed tokenId, address indexed owner, uint256 indexed roundNum, bytes32 seed)
func (_CosmicSignature *CosmicSignatureFilterer) FilterMintEvent(opts *bind.FilterOpts, tokenId []*big.Int, owner []common.Address, roundNum []*big.Int) (*CosmicSignatureMintEventIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _CosmicSignature.contract.FilterLogs(opts, "MintEvent", tokenIdRule, ownerRule, roundNumRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureMintEventIterator{contract: _CosmicSignature.contract, event: "MintEvent", logs: logs, sub: sub}, nil
}

// WatchMintEvent is a free log subscription operation binding the contract event 0xc646da88dc2b2526461a0ebb4326e2418ec0bea89496b632b7c9ee42fbfe1d4d.
//
// Solidity: event MintEvent(uint256 indexed tokenId, address indexed owner, uint256 indexed roundNum, bytes32 seed)
func (_CosmicSignature *CosmicSignatureFilterer) WatchMintEvent(opts *bind.WatchOpts, sink chan<- *CosmicSignatureMintEvent, tokenId []*big.Int, owner []common.Address, roundNum []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _CosmicSignature.contract.WatchLogs(opts, "MintEvent", tokenIdRule, ownerRule, roundNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureMintEvent)
				if err := _CosmicSignature.contract.UnpackLog(event, "MintEvent", log); err != nil {
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

// ParseMintEvent is a log parse operation binding the contract event 0xc646da88dc2b2526461a0ebb4326e2418ec0bea89496b632b7c9ee42fbfe1d4d.
//
// Solidity: event MintEvent(uint256 indexed tokenId, address indexed owner, uint256 indexed roundNum, bytes32 seed)
func (_CosmicSignature *CosmicSignatureFilterer) ParseMintEvent(log types.Log) (*CosmicSignatureMintEvent, error) {
	event := new(CosmicSignatureMintEvent)
	if err := _CosmicSignature.contract.UnpackLog(event, "MintEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CosmicSignature contract.
type CosmicSignatureOwnershipTransferredIterator struct {
	Event *CosmicSignatureOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureOwnershipTransferred)
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
		it.Event = new(CosmicSignatureOwnershipTransferred)
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
func (it *CosmicSignatureOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureOwnershipTransferred represents a OwnershipTransferred event raised by the CosmicSignature contract.
type CosmicSignatureOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CosmicSignature *CosmicSignatureFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CosmicSignatureOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CosmicSignature.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureOwnershipTransferredIterator{contract: _CosmicSignature.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CosmicSignature *CosmicSignatureFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CosmicSignatureOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CosmicSignature.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureOwnershipTransferred)
				if err := _CosmicSignature.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CosmicSignature *CosmicSignatureFilterer) ParseOwnershipTransferred(log types.Log) (*CosmicSignatureOwnershipTransferred, error) {
	event := new(CosmicSignatureOwnershipTransferred)
	if err := _CosmicSignature.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureTokenGenerationScriptURLEventIterator is returned from FilterTokenGenerationScriptURLEvent and is used to iterate over the raw logs and unpacked data for TokenGenerationScriptURLEvent events raised by the CosmicSignature contract.
type CosmicSignatureTokenGenerationScriptURLEventIterator struct {
	Event *CosmicSignatureTokenGenerationScriptURLEvent // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureTokenGenerationScriptURLEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureTokenGenerationScriptURLEvent)
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
		it.Event = new(CosmicSignatureTokenGenerationScriptURLEvent)
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
func (it *CosmicSignatureTokenGenerationScriptURLEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureTokenGenerationScriptURLEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureTokenGenerationScriptURLEvent represents a TokenGenerationScriptURLEvent event raised by the CosmicSignature contract.
type CosmicSignatureTokenGenerationScriptURLEvent struct {
	NewURL string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenGenerationScriptURLEvent is a free log retrieval operation binding the contract event 0x0119741ee0f95fab26124262a82c3c0e9e1c7ff4bb33c6fba5f3b11c9b6d0bad.
//
// Solidity: event TokenGenerationScriptURLEvent(string newURL)
func (_CosmicSignature *CosmicSignatureFilterer) FilterTokenGenerationScriptURLEvent(opts *bind.FilterOpts) (*CosmicSignatureTokenGenerationScriptURLEventIterator, error) {

	logs, sub, err := _CosmicSignature.contract.FilterLogs(opts, "TokenGenerationScriptURLEvent")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureTokenGenerationScriptURLEventIterator{contract: _CosmicSignature.contract, event: "TokenGenerationScriptURLEvent", logs: logs, sub: sub}, nil
}

// WatchTokenGenerationScriptURLEvent is a free log subscription operation binding the contract event 0x0119741ee0f95fab26124262a82c3c0e9e1c7ff4bb33c6fba5f3b11c9b6d0bad.
//
// Solidity: event TokenGenerationScriptURLEvent(string newURL)
func (_CosmicSignature *CosmicSignatureFilterer) WatchTokenGenerationScriptURLEvent(opts *bind.WatchOpts, sink chan<- *CosmicSignatureTokenGenerationScriptURLEvent) (event.Subscription, error) {

	logs, sub, err := _CosmicSignature.contract.WatchLogs(opts, "TokenGenerationScriptURLEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureTokenGenerationScriptURLEvent)
				if err := _CosmicSignature.contract.UnpackLog(event, "TokenGenerationScriptURLEvent", log); err != nil {
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

// ParseTokenGenerationScriptURLEvent is a log parse operation binding the contract event 0x0119741ee0f95fab26124262a82c3c0e9e1c7ff4bb33c6fba5f3b11c9b6d0bad.
//
// Solidity: event TokenGenerationScriptURLEvent(string newURL)
func (_CosmicSignature *CosmicSignatureFilterer) ParseTokenGenerationScriptURLEvent(log types.Log) (*CosmicSignatureTokenGenerationScriptURLEvent, error) {
	event := new(CosmicSignatureTokenGenerationScriptURLEvent)
	if err := _CosmicSignature.contract.UnpackLog(event, "TokenGenerationScriptURLEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureTokenNameEventIterator is returned from FilterTokenNameEvent and is used to iterate over the raw logs and unpacked data for TokenNameEvent events raised by the CosmicSignature contract.
type CosmicSignatureTokenNameEventIterator struct {
	Event *CosmicSignatureTokenNameEvent // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureTokenNameEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureTokenNameEvent)
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
		it.Event = new(CosmicSignatureTokenNameEvent)
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
func (it *CosmicSignatureTokenNameEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureTokenNameEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureTokenNameEvent represents a TokenNameEvent event raised by the CosmicSignature contract.
type CosmicSignatureTokenNameEvent struct {
	TokenId *big.Int
	NewName string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenNameEvent is a free log retrieval operation binding the contract event 0x8ad5e159ff95649c8a9f323ac5a457e741897cf44ce07dfce0e98b84ef9d5f12.
//
// Solidity: event TokenNameEvent(uint256 indexed tokenId, string newName)
func (_CosmicSignature *CosmicSignatureFilterer) FilterTokenNameEvent(opts *bind.FilterOpts, tokenId []*big.Int) (*CosmicSignatureTokenNameEventIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CosmicSignature.contract.FilterLogs(opts, "TokenNameEvent", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureTokenNameEventIterator{contract: _CosmicSignature.contract, event: "TokenNameEvent", logs: logs, sub: sub}, nil
}

// WatchTokenNameEvent is a free log subscription operation binding the contract event 0x8ad5e159ff95649c8a9f323ac5a457e741897cf44ce07dfce0e98b84ef9d5f12.
//
// Solidity: event TokenNameEvent(uint256 indexed tokenId, string newName)
func (_CosmicSignature *CosmicSignatureFilterer) WatchTokenNameEvent(opts *bind.WatchOpts, sink chan<- *CosmicSignatureTokenNameEvent, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CosmicSignature.contract.WatchLogs(opts, "TokenNameEvent", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureTokenNameEvent)
				if err := _CosmicSignature.contract.UnpackLog(event, "TokenNameEvent", log); err != nil {
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
// Solidity: event TokenNameEvent(uint256 indexed tokenId, string newName)
func (_CosmicSignature *CosmicSignatureFilterer) ParseTokenNameEvent(log types.Log) (*CosmicSignatureTokenNameEvent, error) {
	event := new(CosmicSignatureTokenNameEvent)
	if err := _CosmicSignature.contract.UnpackLog(event, "TokenNameEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CosmicSignature contract.
type CosmicSignatureTransferIterator struct {
	Event *CosmicSignatureTransfer // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureTransfer)
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
		it.Event = new(CosmicSignatureTransfer)
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
func (it *CosmicSignatureTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureTransfer represents a Transfer event raised by the CosmicSignature contract.
type CosmicSignatureTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CosmicSignature *CosmicSignatureFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*CosmicSignatureTransferIterator, error) {

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

	logs, sub, err := _CosmicSignature.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureTransferIterator{contract: _CosmicSignature.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CosmicSignature *CosmicSignatureFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CosmicSignatureTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _CosmicSignature.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureTransfer)
				if err := _CosmicSignature.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_CosmicSignature *CosmicSignatureFilterer) ParseTransfer(log types.Log) (*CosmicSignatureTransfer, error) {
	event := new(CosmicSignatureTransfer)
	if err := _CosmicSignature.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

