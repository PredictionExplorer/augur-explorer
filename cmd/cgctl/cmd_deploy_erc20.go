package main

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"
)

// SampTokenMetaData contains all meta data concerning the SampToken contract.
var SampTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b506040518060400160405280601381526020017f455243323020546f6b656e2053616d706c6531000000000000000000000000008152506040518060400160405280600881526020016753616d706c65203160c01b815250816003908161007691906102a3565b50600461008382826102a3565b5050506100a3336c01431e0fae6d7217caa00000006100a860201b60201c565b610382565b6001600160a01b0382166100d65760405163ec442f0560e01b81525f60048201526024015b60405180910390fd5b6100e15f83836100e5565b5050565b6001600160a01b03831661010f578060025f828254610104919061035d565b9091555061017f9050565b6001600160a01b0383165f90815260208190526040902054818110156101615760405163391434e360e21b81526001600160a01b038516600482015260248101829052604481018390526064016100cd565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b03821661019b576002805482900390556101b9565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516101fe91815260200190565b60405180910390a3505050565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061023357607f821691505b60208210810361025157634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561029e57805f5260205f20601f840160051c8101602085101561027c5750805b601f840160051c820191505b8181101561029b575f8155600101610288565b50505b505050565b81516001600160401b038111156102bc576102bc61020b565b6102d0816102ca845461021f565b84610257565b6020601f821160018114610302575f83156102eb5750848201515b5f19600385901b1c1916600184901b17845561029b565b5f84815260208120601f198516915b828110156103315787850151825560209485019460019092019101610311565b508482101561034e57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b8082018082111561037c57634e487b7160e01b5f52601160045260245ffd5b92915050565b6107928061038f5f395ff3fe608060405234801561000f575f5ffd5b50600436106100a6575f3560e01c806342966c681161006e57806342966c681461011f57806370a082311461013457806379cc67901461015c57806395d89b411461016f578063a9059cbb14610177578063dd62ed3e1461018a575f5ffd5b806306fdde03146100aa578063095ea7b3146100c857806318160ddd146100eb57806323b872dd146100fd578063313ce56714610110575b5f5ffd5b6100b26101c2565b6040516100bf91906105eb565b60405180910390f35b6100db6100d636600461063b565b610252565b60405190151581526020016100bf565b6002545b6040519081526020016100bf565b6100db61010b366004610663565b61026b565b604051601281526020016100bf565b61013261012d36600461069d565b61028e565b005b6100ef6101423660046106b4565b6001600160a01b03165f9081526020819052604090205490565b61013261016a36600461063b565b61029b565b6100b26102b4565b6100db61018536600461063b565b6102c3565b6100ef6101983660046106d4565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b6060600380546101d190610705565b80601f01602080910402602001604051908101604052809291908181526020018280546101fd90610705565b80156102485780601f1061021f57610100808354040283529160200191610248565b820191905f5260205f20905b81548152906001019060200180831161022b57829003601f168201915b5050505050905090565b5f3361025f8185856102d0565b60019150505b92915050565b5f336102788582856102e2565b610283858585610362565b506001949350505050565b61029833826103bf565b50565b6102a68233836102e2565b6102b082826103bf565b5050565b6060600480546101d190610705565b5f3361025f818585610362565b6102dd83838360016103f3565b505050565b6001600160a01b038381165f908152600160209081526040808320938616835292905220545f19811461035c578181101561034e57604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064015b60405180910390fd5b61035c84848484035f6103f3565b50505050565b6001600160a01b03831661038b57604051634b637e8f60e11b81525f6004820152602401610345565b6001600160a01b0382166103b45760405163ec442f0560e01b81525f6004820152602401610345565b6102dd8383836104c5565b6001600160a01b0382166103e857604051634b637e8f60e11b81525f6004820152602401610345565b6102b0825f836104c5565b6001600160a01b03841661041c5760405163e602df0560e01b81525f6004820152602401610345565b6001600160a01b03831661044557604051634a1406b160e11b81525f6004820152602401610345565b6001600160a01b038085165f908152600160209081526040808320938716835292905220829055801561035c57826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516104b791815260200190565b60405180910390a350505050565b6001600160a01b0383166104ef578060025f8282546104e4919061073d565b9091555061055f9050565b6001600160a01b0383165f90815260208190526040902054818110156105415760405163391434e360e21b81526001600160a01b03851660048201526024810182905260448101839052606401610345565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b03821661057b57600280548290039055610599565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516105de91815260200190565b60405180910390a3505050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80356001600160a01b0381168114610636575f5ffd5b919050565b5f5f6040838503121561064c575f5ffd5b61065583610620565b946020939093013593505050565b5f5f5f60608486031215610675575f5ffd5b61067e84610620565b925061068c60208501610620565b929592945050506040919091013590565b5f602082840312156106ad575f5ffd5b5035919050565b5f602082840312156106c4575f5ffd5b6106cd82610620565b9392505050565b5f5f604083850312156106e5575f5ffd5b6106ee83610620565b91506106fc60208401610620565b90509250929050565b600181811c9082168061071957607f821691505b60208210810361073757634e487b7160e01b5f52602260045260245ffd5b50919050565b8082018082111561026557634e487b7160e01b5f52601160045260245ffdfea264697066735822122080fb5507d50ef0406d8f88991a7e2d36126f4f49acc1055dd41ac1553dbc15c464736f6c634300081e0033",
}

// deploySampToken deploys a new SampToken contract and returns its address
// and the deployment transaction.
func deploySampToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, error) {
	parsed, err := SampTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, errors.New("GetABI returned nil")
	}

	address, tx, _, err := bind.DeployContract(auth, *parsed, common.FromHex(SampTokenMetaData.Bin), backend)
	if err != nil {
		return common.Address{}, nil, err
	}
	return address, tx, nil
}

// newDeployERC20Cmd builds the deploy-erc20 subcommand.
func newDeployERC20Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "deploy-erc20",
		Short: "Deploy a sample ERC-20 token for donation testing",
		Long: `Deploy a sample ERC-20 token (100 billion tokens with 18 decimals) for
testing the token donation mechanism.
Token Name: ERC20 Token Sample1, Symbol: Sample 1.

` + txEnvHelp,
		Args: cobra.NoArgs,
		RunE: runDeployERC20,
	}
}

func init() { register(newDeployERC20Cmd()) }

func runDeployERC20(cmd *cobra.Command, _ []string) error {
	w := cmd.OutOrStdout()
	s, err := newTxSession(cmd, false)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "Using chain_id=%v\n", s.Net.ChainID.String())
	fmt.Fprintf(w, "Deployer address: %v\n", s.Acc.Address.Hex())

	// Gas price headroom above the block base fee comes from the session's
	// GAS_PRICE_MULTIPLIER policy (default 2x, same as the old script).
	txopts := s.TransactOpts(big.NewInt(0), uint64(3000000))
	fmt.Fprintf(w, "Gas price: %v\n", txopts.GasPrice.String())

	fmt.Fprintf(w, "Deploying SampToken contract...\n")
	contractAddr, tx, err := deploySampToken(txopts, s.Net.Client)
	if err != nil {
		return fmt.Errorf("deploying contract: %w", err)
	}

	fmt.Fprintf(w, "Deployment tx hash: %v\n", tx.Hash().String())
	fmt.Fprintf(w, "Contract will be deployed at: %v\n", contractAddr.Hex())
	fmt.Fprintf(w, "\nWaiting for transaction to be mined...\n")

	receipt, err := s.WaitForReceipt(cmd.Context(), tx)
	if err != nil {
		return fmt.Errorf("waiting for mining: %w", err)
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("transaction failed with status: %v", receipt.Status)
	}
	fmt.Fprintf(w, "\n=== DEPLOYMENT SUCCESSFUL ===\n")
	fmt.Fprintf(w, "Contract Address: %v\n", receipt.ContractAddress.Hex())
	fmt.Fprintf(w, "Block Number: %v\n", receipt.BlockNumber)
	fmt.Fprintf(w, "Gas Used: %v\n", receipt.GasUsed)
	fmt.Fprintf(w, "Token Name: ERC20 Token Sample1\n")
	fmt.Fprintf(w, "Token Symbol: Sample 1\n")
	fmt.Fprintf(w, "Total Supply: 100,000,000,000 tokens (with 18 decimals)\n")
	return nil
}
