// Deploys a sample ERC20 token for testing token donation mechanism
package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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

// SampTokenMetaData contains all meta data concerning the SampToken contract.
var SampTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b506040518060400160405280601381526020017f455243323020546f6b656e2053616d706c6531000000000000000000000000008152506040518060400160405280600881526020016753616d706c65203160c01b815250816003908161007691906102a3565b50600461008382826102a3565b5050506100a3336c01431e0fae6d7217caa00000006100a860201b60201c565b610382565b6001600160a01b0382166100d65760405163ec442f0560e01b81525f60048201526024015b60405180910390fd5b6100e15f83836100e5565b5050565b6001600160a01b03831661010f578060025f828254610104919061035d565b9091555061017f9050565b6001600160a01b0383165f90815260208190526040902054818110156101615760405163391434e360e21b81526001600160a01b038516600482015260248101829052604481018390526064016100cd565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b03821661019b576002805482900390556101b9565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516101fe91815260200190565b60405180910390a3505050565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061023357607f821691505b60208210810361025157634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561029e57805f5260205f20601f840160051c8101602085101561027c5750805b601f840160051c820191505b8181101561029b575f8155600101610288565b50505b505050565b81516001600160401b038111156102bc576102bc61020b565b6102d0816102ca845461021f565b84610257565b6020601f821160018114610302575f83156102eb5750848201515b5f19600385901b1c1916600184901b17845561029b565b5f84815260208120601f198516915b828110156103315787850151825560209485019460019092019101610311565b508482101561034e57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b8082018082111561037c57634e487b7160e01b5f52601160045260245ffd5b92915050565b6107928061038f5f395ff3fe608060405234801561000f575f5ffd5b50600436106100a6575f3560e01c806342966c681161006e57806342966c681461011f57806370a082311461013457806379cc67901461015c57806395d89b411461016f578063a9059cbb14610177578063dd62ed3e1461018a575f5ffd5b806306fdde03146100aa578063095ea7b3146100c857806318160ddd146100eb57806323b872dd146100fd578063313ce56714610110575b5f5ffd5b6100b26101c2565b6040516100bf91906105eb565b60405180910390f35b6100db6100d636600461063b565b610252565b60405190151581526020016100bf565b6002545b6040519081526020016100bf565b6100db61010b366004610663565b61026b565b604051601281526020016100bf565b61013261012d36600461069d565b61028e565b005b6100ef6101423660046106b4565b6001600160a01b03165f9081526020819052604090205490565b61013261016a36600461063b565b61029b565b6100b26102b4565b6100db61018536600461063b565b6102c3565b6100ef6101983660046106d4565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b6060600380546101d190610705565b80601f01602080910402602001604051908101604052809291908181526020018280546101fd90610705565b80156102485780601f1061021f57610100808354040283529160200191610248565b820191905f5260205f20905b81548152906001019060200180831161022b57829003601f168201915b5050505050905090565b5f3361025f8185856102d0565b60019150505b92915050565b5f336102788582856102e2565b610283858585610362565b506001949350505050565b61029833826103bf565b50565b6102a68233836102e2565b6102b082826103bf565b5050565b6060600480546101d190610705565b5f3361025f818585610362565b6102dd83838360016103f3565b505050565b6001600160a01b038381165f908152600160209081526040808320938616835292905220545f19811461035c578181101561034e57604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064015b60405180910390fd5b61035c84848484035f6103f3565b50505050565b6001600160a01b03831661038b57604051634b637e8f60e11b81525f6004820152602401610345565b6001600160a01b0382166103b45760405163ec442f0560e01b81525f6004820152602401610345565b6102dd8383836104c5565b6001600160a01b0382166103e857604051634b637e8f60e11b81525f6004820152602401610345565b6102b0825f836104c5565b6001600160a01b03841661041c5760405163e602df0560e01b81525f6004820152602401610345565b6001600160a01b03831661044557604051634a1406b160e11b81525f6004820152602401610345565b6001600160a01b038085165f908152600160209081526040808320938716835292905220829055801561035c57826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516104b791815260200190565b60405180910390a350505050565b6001600160a01b0383166104ef578060025f8282546104e4919061073d565b9091555061055f9050565b6001600160a01b0383165f90815260208190526040902054818110156105415760405163391434e360e21b81526001600160a01b03851660048201526024810182905260448101839052606401610345565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b03821661057b57600280548290039055610599565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516105de91815260200190565b60405180910390a3505050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80356001600160a01b0381168114610636575f5ffd5b919050565b5f5f6040838503121561064c575f5ffd5b61065583610620565b946020939093013593505050565b5f5f5f60608486031215610675575f5ffd5b61067e84610620565b925061068c60208501610620565b929592945050506040919091013590565b5f602082840312156106ad575f5ffd5b5035919050565b5f602082840312156106c4575f5ffd5b6106cd82610620565b9392505050565b5f5f604083850312156106e5575f5ffd5b6106ee83610620565b91506106fc60208401610620565b90509250929050565b600181811c9082168061071957607f821691505b60208210810361073757634e487b7160e01b5f52602260045260245ffd5b50919050565b8082018082111561026557634e487b7160e01b5f52601160045260245ffdfea264697066735822122080fb5507d50ef0406d8f88991a7e2d36126f4f49acc1055dd41ac1553dbc15c464736f6c634300081e0033",
}

// SampTokenABI is the input ABI used to generate the binding from.
var SampTokenABI = SampTokenMetaData.ABI

// SampTokenBin is the compiled bytecode used for deploying new contracts.
var SampTokenBin = SampTokenMetaData.Bin

// DeploySampToken deploys a new Ethereum contract, binding an instance of SampToken to it.
func DeploySampToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SampToken, error) {
	parsed, err := SampTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SampTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SampToken{SampTokenCaller: SampTokenCaller{contract: contract}, SampTokenTransactor: SampTokenTransactor{contract: contract}, SampTokenFilterer: SampTokenFilterer{contract: contract}}, nil
}

// SampToken is an auto generated Go binding around an Ethereum contract.
type SampToken struct {
	SampTokenCaller     // Read-only binding to the contract
	SampTokenTransactor // Write-only binding to the contract
	SampTokenFilterer   // Log filterer for contract events
}

// SampTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type SampTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SampTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SampTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// bindSampToken binds a generic wrapper to an already deployed contract.
func bindSampToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SampTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

var (
	RPC_URL string
)

func main() {
	RPC_URL = os.Getenv("RPC_URL")
	if RPC_URL == "" {
		fmt.Printf("Error: RPC_URL environment variable is not set\n")
		os.Exit(1)
	}

	eclient, err := ethclient.Dial(RPC_URL)
	if err != nil {
		fmt.Printf("Can't connect to ETH RPC: %v\n", err)
		os.Exit(1)
	}

	big_chain_id, err := eclient.NetworkID(context.Background())
	if err != nil {
		fmt.Printf("Error getting network ID: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Using chain_id=%v\n", big_chain_id.String())

	from_pkey_str := os.Getenv("PKEY_HEX")
	if from_pkey_str == "" {
		fmt.Printf("Usage: %v\n\nDeploys a sample ERC20 token (100 billion tokens with 18 decimals).\nToken Name: ERC20 Token Sample1, Symbol: Sample 1\n\nEnvironment: RPC_URL and PKEY_HEX (64-char hex private key, no 0x prefix) must be set.\n", os.Args[0])
		os.Exit(1)
	}
	if len(from_pkey_str) != 64 {
		fmt.Printf("PKEY_HEX must be 64 hex characters (without 0x prefix)\n")
		os.Exit(1)
	}

	from_PrivateKey, err := crypto.HexToECDSA(from_pkey_str)
	if err != nil {
		fmt.Printf("Error parsing private key: %v\n", err)
		os.Exit(1)
	}

	from_publicKey := from_PrivateKey.Public()
	from_publicKeyECDSA, ok := from_publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Printf("Couldn't derive public key from private key\n")
		os.Exit(1)
	}
	from_address := crypto.PubkeyToAddress(*from_publicKeyECDSA)
	fmt.Printf("Deployer address: %v\n", from_address.Hex())

	from_nonce, err := eclient.PendingNonceAt(context.Background(), from_address)
	if err != nil {
		fmt.Printf("Error getting account's nonce: %v\n", err)
		os.Exit(1)
	}

	gasPrice, err := eclient.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested gas price: %v\n", err)
		os.Exit(1)
	}

	txopts := bind.NewKeyedTransactor(from_PrivateKey)
	txopts.Nonce = big.NewInt(int64(from_nonce))
	txopts.Value = big.NewInt(0)
	txopts.GasLimit = uint64(3000000) // Enough for ERC20 deployment
	txopts.GasPrice = gasPrice.Add(gasPrice, big.NewInt(20000))

	fmt.Printf("Gas price: %v\n", gasPrice.String())

	signfunc := func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		signer := types.NewEIP155Signer(big_chain_id)
		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), from_PrivateKey)
		if err != nil {
			fmt.Printf("Error signing: %v\n", err)
			return nil, err
		}
		return tx.WithSignature(signer, signature)
	}
	txopts.Signer = signfunc

	fmt.Printf("Deploying SampToken contract...\n")
	contractAddr, tx, _, err := DeploySampToken(txopts, eclient)
	if err != nil {
		fmt.Printf("Error deploying contract: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Deployment tx hash: %v\n", tx.Hash().String())
	fmt.Printf("Contract will be deployed at: %v\n", contractAddr.Hex())
	fmt.Printf("\nWaiting for transaction to be mined...\n")

	receipt, err := bind.WaitMined(context.Background(), eclient, tx)
	if err != nil {
		fmt.Printf("Error waiting for mining: %v\n", err)
		os.Exit(1)
	}

	if receipt.Status == types.ReceiptStatusSuccessful {
		fmt.Printf("\n=== DEPLOYMENT SUCCESSFUL ===\n")
		fmt.Printf("Contract Address: %v\n", receipt.ContractAddress.Hex())
		fmt.Printf("Block Number: %v\n", receipt.BlockNumber)
		fmt.Printf("Gas Used: %v\n", receipt.GasUsed)
		fmt.Printf("Token Name: ERC20 Token Sample1\n")
		fmt.Printf("Token Symbol: Sample 1\n")
		fmt.Printf("Total Supply: 100,000,000,000 tokens (with 18 decimals)\n")
	} else {
		fmt.Printf("Transaction failed with status: %v\n", receipt.Status)
		os.Exit(1)
	}
}
