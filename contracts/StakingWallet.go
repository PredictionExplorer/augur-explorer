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

// StakingWalletMetaData contains all meta data concerning the StakingWallet contract.
var StakingWalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractCosmicSignature\",\"name\":\"nft_\",\"type\":\"address\"},{\"internalType\":\"contractCosmicGame\",\"name\":\"game_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"charity_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newCharityAddress\",\"type\":\"address\"}],\"name\":\"CharityUpdatedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"ClaimRewardEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"depositTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNFTs\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"modulo\",\"type\":\"uint256\"}],\"name\":\"EthDepositEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalNFTs\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unstakeTime\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"StakeActionEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalNFTs\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"UnstakeActionEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ETHDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"depositTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numStaked\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"charity\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"actions\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"deposits\",\"type\":\"uint256[]\"}],\"name\":\"claimManyRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ETHDepositId\",\"type\":\"uint256\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"game\",\"outputs\":[{\"internalType\":\"contractCosmicGame\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"modulo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nft\",\"outputs\":[{\"internalType\":\"contractCosmicSignature\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numETHDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakeActions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakedNFTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCharityAddress\",\"type\":\"address\"}],\"name\":\"setCharity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakedNFTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeEligibleTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"unstakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620028c7380380620028c78339818101604052810190620000379190620002f7565b620000576200004b6200012360201b60201c565b6200012b60201b60201c565b82600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505062000353565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200022182620001f4565b9050919050565b6000620002358262000214565b9050919050565b620002478162000228565b81146200025357600080fd5b50565b60008151905062000267816200023c565b92915050565b60006200027a82620001f4565b9050919050565b60006200028e826200026d565b9050919050565b620002a08162000281565b8114620002ac57600080fd5b50565b600081519050620002c08162000295565b92915050565b620002d18162000214565b8114620002dd57600080fd5b50565b600081519050620002f181620002c6565b92915050565b600080600060608486031215620003135762000312620001ef565b5b6000620003238682870162000256565b93505060206200033686828701620002af565b92505060406200034986828701620002e0565b9150509250925092565b61256480620003636000396000f3fe60806040526004361061011f5760003560e01c80638da5cb5b116100a0578063b865749d11610064578063b865749d14610390578063c3fe3e28146103d1578063f2fde38b146103fc578063fb6f71a314610425578063fe939afc1461044e5761011f565b80638da5cb5b146102ca578063934aa023146102f5578063a531aa8614610320578063a694fc3a1461034b578063b6b55f25146103745761011f565b806347ccca02116100e757806347ccca021461020b57806355279fdb146102365780636034eb5b14610261578063715018a61461028a57806386bb8f37146102a15761011f565b80630d50c1891461012457806317db62131461014d57806329745262146101785780632e17de78146101a3578063451f1adf146101cc575b600080fd5b34801561013057600080fd5b5061014b60048036038101906101469190611830565b610477565b005b34801561015957600080fd5b506101626104bd565b60405161016f9190611888565b60405180910390f35b34801561018457600080fd5b5061018d6104c3565b60405161019a9190611888565b60405180910390f35b3480156101af57600080fd5b506101ca60048036038101906101c591906118a3565b6104c9565b005b3480156101d857600080fd5b506101f360048036038101906101ee91906118a3565b610767565b604051610202939291906118d0565b60405180910390f35b34801561021757600080fd5b50610220610791565b60405161022d9190611986565b60405180910390f35b34801561024257600080fd5b5061024b6107b7565b6040516102589190611888565b60405180910390f35b34801561026d57600080fd5b50610288600480360381019061028391906119a1565b6107bd565b005b34801561029657600080fd5b5061029f610863565b005b3480156102ad57600080fd5b506102c860048036038101906102c39190611a19565b6108eb565b005b3480156102d657600080fd5b506102df610cec565b6040516102ec9190611a7a565b60405180910390f35b34801561030157600080fd5b5061030a610d15565b6040516103179190611a7a565b60405180910390f35b34801561032c57600080fd5b50610335610d3b565b6040516103429190611888565b60405180910390f35b34801561035757600080fd5b50610372600480360381019061036d91906118a3565b610d41565b005b61038e600480360381019061038991906118a3565b611019565b005b34801561039c57600080fd5b506103b760048036038101906103b291906118a3565b611274565b6040516103c8959493929190611a95565b60405180910390f35b3480156103dd57600080fd5b506103e66112ca565b6040516103f39190611b1b565b60405180910390f35b34801561040857600080fd5b50610423600480360381019061041e9190611b62565b6112f0565b005b34801561043157600080fd5b5061044c60048036038101906104479190611b62565b6113e7565b005b34801561045a57600080fd5b5061047560048036038101906104709190611830565b61157b565b005b60005b81518110156104b9576104a682828151811061049957610498611b8f565b5b60200260200101516104c9565b80806104b190611bed565b91505061047a565b5050565b60055481565b60075481565b6000600160008381526020019081526020016000206003015414610522576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161051990611c92565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff166001600083815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146105c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105bd90611cfe565b60405180910390fd5b4260016000838152602001908152602001600020600401541061061e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161061590611d6a565b60405180910390fd5b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3033600160006002548152602001908152602001600020600001546040518463ffffffff1660e01b815260040161069593929190611d8a565b600060405180830381600087803b1580156106af57600080fd5b505af11580156106c3573d6000803e3d6000fd5b505050504260016000838152602001908152602001600020600301819055506001600560008282546106f59190611dc1565b925050819055503373ffffffffffffffffffffffffffffffffffffffff166001600083815260200190815260200160002060000154827f33940a9fc3ee3c9150b29b34ef29ca397b5e1e340425a4f0da0346b5b90766c860055460405161075c9190611888565b60405180910390a450565b60036020528060005260406000206000915090508060000154908060010154908060020154905083565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60045481565b8051825114610801576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107f890611e67565b60405180910390fd5b60005b825181101561085e5761084b83828151811061082357610822611b8f565b5b602002602001015183838151811061083e5761083d611b8f565b5b60200260200101516108eb565b808061085690611bed565b915050610804565b505050565b61086b6115c1565b73ffffffffffffffffffffffffffffffffffffffff16610889610cec565b73ffffffffffffffffffffffffffffffffffffffff16146108df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d690611ed3565b60405180910390fd5b6108e960006115c9565b565b6000600160008481526020019081526020016000206003015411610944576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161093b90611f3f565b60405180910390fd5b60016000838152602001908152602001600020600501600082815260200190815260200160002060009054906101000a900460ff16156109b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109b090611f3f565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff166001600084815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610a5d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a5490611fab565b60405180910390fd5b6003600082815260200190815260200160002060000154600160008481526020019081526020016000206002015410610acb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ac290612017565b60405180910390fd5b6003600082815260200190815260200160002060000154600160008481526020019081526020016000206003015411610b39576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b3090612083565b60405180910390fd5b6001806000848152602001908152602001600020600501600083815260200190815260200160002060006101000a81548160ff021916908315150217905550600060036000838152602001908152602001600020600201546003600084815260200190815260200160002060010154610bb291906120d2565b905060006001600085815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1682604051610c1090612134565b60006040518083038185875af1925050503d8060008114610c4d576040519150601f19603f3d011682016040523d82523d6000602084013e610c52565b606091505b5050905080610c96576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c8d90612195565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff1683857fdde81df5caa033e783e2a39d93e3a8718a7dc27ba95a4757f5433a01f794ec3685604051610cde9190611888565b60405180910390a450505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60025481565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b8152600401610da093929190611d8a565b600060405180830381600087803b158015610dba57600080fd5b505af1158015610dce573d6000803e3d6000fd5b5050505080600160006002548152602001908152602001600020600001819055503360016000600254815260200190815260200160002060010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555042600160006002548152602001908152602001600020600201819055506000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ed2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ef691906121ca565b600554620f4240610f0791906121f7565b610f1191906120d2565b905060006227ac406105dc8384610f2891906121f7565b610f3291906120d2565b610f3c9190612239565b90508042610f4a9190612239565b60016000600254815260200190815260200160002060040181905550600160026000828254610f799190612239565b92505081905550600160056000828254610f939190612239565b925050819055503373ffffffffffffffffffffffffffffffffffffffff16836001600254610fc19190611dc1565b7f057eba8c4bba00f858e4f586f9c02794abb0df789ef316c741f9073fe2c435db6005546001600060025481526020019081526020016000206004015460405161100c92919061226d565b60405180910390a4505050565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146110a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110a090612308565b60405180910390fd5b600060055403611187576000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16346040516110fb90612134565b60006040518083038185875af1925050503d8060008114611138576040519150601f19603f3d011682016040523d82523d6000602084013e61113d565b606091505b5050905080611181576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111789061239a565b60405180910390fd5b50611271565b80600360006004548152602001908152602001600020600001819055503460036000600454815260200190815260200160002060010181905550600554600360006004548152602001908152602001600020600201819055506001600460008282546111f39190612239565b925050819055506005543461120891906123ba565b600760008282546112199190612239565b92505081905550807fdc0eacba8b1f88284dca5eec8be23173aefa7206298fe22de43e064b6ccd841860016004546112519190611dc1565b6005543460075460405161126894939291906123eb565b60405180910390a25b50565b60016020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154908060030154908060040154905085565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6112f86115c1565b73ffffffffffffffffffffffffffffffffffffffff16611316610cec565b73ffffffffffffffffffffffffffffffffffffffff161461136c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161136390611ed3565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036113db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113d2906124a2565b60405180910390fd5b6113e4816115c9565b50565b6113ef6115c1565b73ffffffffffffffffffffffffffffffffffffffff1661140d610cec565b73ffffffffffffffffffffffffffffffffffffffff1614611463576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161145a90611ed3565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036114d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114c99061250e565b60405180910390fd5b80600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fa0bd6b2fdbf082ae2356710c23fc8d76d56d418cecb4514d119c77a8617b4ffe60405160405180910390a250565b60005b81518110156115bd576115aa82828151811061159d5761159c611b8f565b5b6020026020010151610d41565b80806115b590611bed565b91505061157e565b5050565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6116ef826116a6565b810181811067ffffffffffffffff8211171561170e5761170d6116b7565b5b80604052505050565b600061172161168d565b905061172d82826116e6565b919050565b600067ffffffffffffffff82111561174d5761174c6116b7565b5b602082029050602081019050919050565b600080fd5b6000819050919050565b61177681611763565b811461178157600080fd5b50565b6000813590506117938161176d565b92915050565b60006117ac6117a784611732565b611717565b905080838252602082019050602084028301858111156117cf576117ce61175e565b5b835b818110156117f857806117e48882611784565b8452602084019350506020810190506117d1565b5050509392505050565b600082601f830112611817576118166116a1565b5b8135611827848260208601611799565b91505092915050565b60006020828403121561184657611845611697565b5b600082013567ffffffffffffffff8111156118645761186361169c565b5b61187084828501611802565b91505092915050565b61188281611763565b82525050565b600060208201905061189d6000830184611879565b92915050565b6000602082840312156118b9576118b8611697565b5b60006118c784828501611784565b91505092915050565b60006060820190506118e56000830186611879565b6118f26020830185611879565b6118ff6040830184611879565b949350505050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061194c61194761194284611907565b611927565b611907565b9050919050565b600061195e82611931565b9050919050565b600061197082611953565b9050919050565b61198081611965565b82525050565b600060208201905061199b6000830184611977565b92915050565b600080604083850312156119b8576119b7611697565b5b600083013567ffffffffffffffff8111156119d6576119d561169c565b5b6119e285828601611802565b925050602083013567ffffffffffffffff811115611a0357611a0261169c565b5b611a0f85828601611802565b9150509250929050565b60008060408385031215611a3057611a2f611697565b5b6000611a3e85828601611784565b9250506020611a4f85828601611784565b9150509250929050565b6000611a6482611907565b9050919050565b611a7481611a59565b82525050565b6000602082019050611a8f6000830184611a6b565b92915050565b600060a082019050611aaa6000830188611879565b611ab76020830187611a6b565b611ac46040830186611879565b611ad16060830185611879565b611ade6080830184611879565b9695505050505050565b6000611af382611931565b9050919050565b6000611b0582611ae8565b9050919050565b611b1581611afa565b82525050565b6000602082019050611b306000830184611b0c565b92915050565b611b3f81611a59565b8114611b4a57600080fd5b50565b600081359050611b5c81611b36565b92915050565b600060208284031215611b7857611b77611697565b5b6000611b8684828501611b4d565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611bf882611763565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611c2a57611c29611bbe565b5b600182019050919050565b600082825260208201905092915050565b7f546f6b656e2068617320616c7265616479206265656e20756e7374616b656400600082015250565b6000611c7c601f83611c35565b9150611c8782611c46565b602082019050919050565b60006020820190508181036000830152611cab81611c6f565b9050919050565b7f4f6e6c7920746865206f776e65722063616e20756e7374616b65000000000000600082015250565b6000611ce8601a83611c35565b9150611cf382611cb2565b602082019050919050565b60006020820190508181036000830152611d1781611cdb565b9050919050565b7f4e6f7420616c6c6f77656420746f20756e7374616b6520796574000000000000600082015250565b6000611d54601a83611c35565b9150611d5f82611d1e565b602082019050919050565b60006020820190508181036000830152611d8381611d47565b9050919050565b6000606082019050611d9f6000830186611a6b565b611dac6020830185611a6b565b611db96040830184611879565b949350505050565b6000611dcc82611763565b9150611dd783611763565b9250828203905081811115611def57611dee611bbe565b5b92915050565b7f417272617920617267756d656e7473206d757374206265206f6620746865207360008201527f616d65206c656e6774682e000000000000000000000000000000000000000000602082015250565b6000611e51602b83611c35565b9150611e5c82611df5565b604082019050919050565b60006020820190508181036000830152611e8081611e44565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000611ebd602083611c35565b9150611ec882611e87565b602082019050919050565b60006020820190508181036000830152611eec81611eb0565b9050919050565b7f546f6b656e20686173206e6f74206265656e20756e7374616b65640000000000600082015250565b6000611f29601b83611c35565b9150611f3482611ef3565b602082019050919050565b60006020820190508181036000830152611f5881611f1c565b9050919050565b7f4f6e6c7920746865206f776e65722063616e20636c61696d2072657761726400600082015250565b6000611f95601f83611c35565b9150611fa082611f5f565b602082019050919050565b60006020820190508181036000830152611fc481611f88565b9050919050565b7f596f752077657265206e6f74207374616b6564207965742e0000000000000000600082015250565b6000612001601883611c35565b915061200c82611fcb565b602082019050919050565b6000602082019050818103600083015261203081611ff4565b9050919050565b7f596f75207765726520616c726561647920756e7374616b65642e000000000000600082015250565b600061206d601a83611c35565b915061207882612037565b602082019050919050565b6000602082019050818103600083015261209c81612060565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006120dd82611763565b91506120e883611763565b9250826120f8576120f76120a3565b5b828204905092915050565b600081905092915050565b50565b600061211e600083612103565b91506121298261210e565b600082019050919050565b600061213f82612111565b9150819050919050565b7f526577617264207472616e73666572206661696c65642e000000000000000000600082015250565b600061217f601783611c35565b915061218a82612149565b602082019050919050565b600060208201905081810360008301526121ae81612172565b9050919050565b6000815190506121c48161176d565b92915050565b6000602082840312156121e0576121df611697565b5b60006121ee848285016121b5565b91505092915050565b600061220282611763565b915061220d83611763565b925082820261221b81611763565b9150828204841483151761223257612231611bbe565b5b5092915050565b600061224482611763565b915061224f83611763565b925082820190508082111561226757612266611bbe565b5b92915050565b60006040820190506122826000830185611879565b61228f6020830184611879565b9392505050565b7f4f6e6c792074686520436f736d696347616d6520636f6e74726163742063616e60008201527f206465706f7369742e0000000000000000000000000000000000000000000000602082015250565b60006122f2602983611c35565b91506122fd82612296565b604082019050919050565b60006020820190508181036000830152612321816122e5565b9050919050565b7f5472616e7366657220746f206368617269747920636f6e74726163742066616960008201527f6c65642e00000000000000000000000000000000000000000000000000000000602082015250565b6000612384602483611c35565b915061238f82612328565b604082019050919050565b600060208201905081810360008301526123b381612377565b9050919050565b60006123c582611763565b91506123d083611763565b9250826123e0576123df6120a3565b5b828206905092915050565b60006080820190506124006000830187611879565b61240d6020830186611879565b61241a6040830185611879565b6124276060830184611879565b95945050505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b600061248c602683611c35565b915061249782612430565b604082019050919050565b600060208201905081810360008301526124bb8161247f565b9050919050565b7f5a65726f2d616464726573732077617320676976656e2e000000000000000000600082015250565b60006124f8601783611c35565b9150612503826124c2565b602082019050919050565b60006020820190508181036000830152612527816124eb565b905091905056fea26469706673582212208f22a3cbb37e44349ac53714753444d541645c57e3921513456cf75d71a4723d64736f6c63430008130033",
}

// StakingWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingWalletMetaData.ABI instead.
var StakingWalletABI = StakingWalletMetaData.ABI

// StakingWalletBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingWalletMetaData.Bin instead.
var StakingWalletBin = StakingWalletMetaData.Bin

// DeployStakingWallet deploys a new Ethereum contract, binding an instance of StakingWallet to it.
func DeployStakingWallet(auth *bind.TransactOpts, backend bind.ContractBackend, nft_ common.Address, game_ common.Address, charity_ common.Address) (common.Address, *types.Transaction, *StakingWallet, error) {
	parsed, err := StakingWalletMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingWalletBin), backend, nft_, game_, charity_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StakingWallet{StakingWalletCaller: StakingWalletCaller{contract: contract}, StakingWalletTransactor: StakingWalletTransactor{contract: contract}, StakingWalletFilterer: StakingWalletFilterer{contract: contract}}, nil
}

// StakingWallet is an auto generated Go binding around an Ethereum contract.
type StakingWallet struct {
	StakingWalletCaller     // Read-only binding to the contract
	StakingWalletTransactor // Write-only binding to the contract
	StakingWalletFilterer   // Log filterer for contract events
}

// StakingWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingWalletSession struct {
	Contract     *StakingWallet    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingWalletCallerSession struct {
	Contract *StakingWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StakingWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingWalletTransactorSession struct {
	Contract     *StakingWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StakingWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingWalletRaw struct {
	Contract *StakingWallet // Generic contract binding to access the raw methods on
}

// StakingWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingWalletCallerRaw struct {
	Contract *StakingWalletCaller // Generic read-only contract binding to access the raw methods on
}

// StakingWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingWalletTransactorRaw struct {
	Contract *StakingWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingWallet creates a new instance of StakingWallet, bound to a specific deployed contract.
func NewStakingWallet(address common.Address, backend bind.ContractBackend) (*StakingWallet, error) {
	contract, err := bindStakingWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingWallet{StakingWalletCaller: StakingWalletCaller{contract: contract}, StakingWalletTransactor: StakingWalletTransactor{contract: contract}, StakingWalletFilterer: StakingWalletFilterer{contract: contract}}, nil
}

// NewStakingWalletCaller creates a new read-only instance of StakingWallet, bound to a specific deployed contract.
func NewStakingWalletCaller(address common.Address, caller bind.ContractCaller) (*StakingWalletCaller, error) {
	contract, err := bindStakingWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingWalletCaller{contract: contract}, nil
}

// NewStakingWalletTransactor creates a new write-only instance of StakingWallet, bound to a specific deployed contract.
func NewStakingWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingWalletTransactor, error) {
	contract, err := bindStakingWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingWalletTransactor{contract: contract}, nil
}

// NewStakingWalletFilterer creates a new log filterer instance of StakingWallet, bound to a specific deployed contract.
func NewStakingWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingWalletFilterer, error) {
	contract, err := bindStakingWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingWalletFilterer{contract: contract}, nil
}

// bindStakingWallet binds a generic wrapper to an already deployed contract.
func bindStakingWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingWalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingWallet *StakingWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingWallet.Contract.StakingWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingWallet *StakingWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingWallet.Contract.StakingWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingWallet *StakingWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingWallet.Contract.StakingWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingWallet *StakingWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingWallet *StakingWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingWallet *StakingWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingWallet.Contract.contract.Transact(opts, method, params...)
}

// ETHDeposits is a free data retrieval call binding the contract method 0x451f1adf.
//
// Solidity: function ETHDeposits(uint256 ) view returns(uint256 depositTime, uint256 depositAmount, uint256 numStaked)
func (_StakingWallet *StakingWalletCaller) ETHDeposits(opts *bind.CallOpts, arg0 *big.Int) (struct {
	DepositTime   *big.Int
	DepositAmount *big.Int
	NumStaked     *big.Int
}, error) {
	var out []interface{}
	err := _StakingWallet.contract.Call(opts, &out, "ETHDeposits", arg0)

	outstruct := new(struct {
		DepositTime   *big.Int
		DepositAmount *big.Int
		NumStaked     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DepositTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DepositAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NumStaked = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ETHDeposits is a free data retrieval call binding the contract method 0x451f1adf.
//
// Solidity: function ETHDeposits(uint256 ) view returns(uint256 depositTime, uint256 depositAmount, uint256 numStaked)
func (_StakingWallet *StakingWalletSession) ETHDeposits(arg0 *big.Int) (struct {
	DepositTime   *big.Int
	DepositAmount *big.Int
	NumStaked     *big.Int
}, error) {
	return _StakingWallet.Contract.ETHDeposits(&_StakingWallet.CallOpts, arg0)
}

// ETHDeposits is a free data retrieval call binding the contract method 0x451f1adf.
//
// Solidity: function ETHDeposits(uint256 ) view returns(uint256 depositTime, uint256 depositAmount, uint256 numStaked)
func (_StakingWallet *StakingWalletCallerSession) ETHDeposits(arg0 *big.Int) (struct {
	DepositTime   *big.Int
	DepositAmount *big.Int
	NumStaked     *big.Int
}, error) {
	return _StakingWallet.Contract.ETHDeposits(&_StakingWallet.CallOpts, arg0)
}

// Charity is a free data retrieval call binding the contract method 0x934aa023.
//
// Solidity: function charity() view returns(address)
func (_StakingWallet *StakingWalletCaller) Charity(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingWallet.contract.Call(opts, &out, "charity")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Charity is a free data retrieval call binding the contract method 0x934aa023.
//
// Solidity: function charity() view returns(address)
func (_StakingWallet *StakingWalletSession) Charity() (common.Address, error) {
	return _StakingWallet.Contract.Charity(&_StakingWallet.CallOpts)
}

// Charity is a free data retrieval call binding the contract method 0x934aa023.
//
// Solidity: function charity() view returns(address)
func (_StakingWallet *StakingWalletCallerSession) Charity() (common.Address, error) {
	return _StakingWallet.Contract.Charity(&_StakingWallet.CallOpts)
}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_StakingWallet *StakingWalletCaller) Game(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingWallet.contract.Call(opts, &out, "game")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_StakingWallet *StakingWalletSession) Game() (common.Address, error) {
	return _StakingWallet.Contract.Game(&_StakingWallet.CallOpts)
}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_StakingWallet *StakingWalletCallerSession) Game() (common.Address, error) {
	return _StakingWallet.Contract.Game(&_StakingWallet.CallOpts)
}

// Modulo is a free data retrieval call binding the contract method 0x29745262.
//
// Solidity: function modulo() view returns(uint256)
func (_StakingWallet *StakingWalletCaller) Modulo(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWallet.contract.Call(opts, &out, "modulo")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Modulo is a free data retrieval call binding the contract method 0x29745262.
//
// Solidity: function modulo() view returns(uint256)
func (_StakingWallet *StakingWalletSession) Modulo() (*big.Int, error) {
	return _StakingWallet.Contract.Modulo(&_StakingWallet.CallOpts)
}

// Modulo is a free data retrieval call binding the contract method 0x29745262.
//
// Solidity: function modulo() view returns(uint256)
func (_StakingWallet *StakingWalletCallerSession) Modulo() (*big.Int, error) {
	return _StakingWallet.Contract.Modulo(&_StakingWallet.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_StakingWallet *StakingWalletCaller) Nft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingWallet.contract.Call(opts, &out, "nft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_StakingWallet *StakingWalletSession) Nft() (common.Address, error) {
	return _StakingWallet.Contract.Nft(&_StakingWallet.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_StakingWallet *StakingWalletCallerSession) Nft() (common.Address, error) {
	return _StakingWallet.Contract.Nft(&_StakingWallet.CallOpts)
}

// NumETHDeposits is a free data retrieval call binding the contract method 0x55279fdb.
//
// Solidity: function numETHDeposits() view returns(uint256)
func (_StakingWallet *StakingWalletCaller) NumETHDeposits(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWallet.contract.Call(opts, &out, "numETHDeposits")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumETHDeposits is a free data retrieval call binding the contract method 0x55279fdb.
//
// Solidity: function numETHDeposits() view returns(uint256)
func (_StakingWallet *StakingWalletSession) NumETHDeposits() (*big.Int, error) {
	return _StakingWallet.Contract.NumETHDeposits(&_StakingWallet.CallOpts)
}

// NumETHDeposits is a free data retrieval call binding the contract method 0x55279fdb.
//
// Solidity: function numETHDeposits() view returns(uint256)
func (_StakingWallet *StakingWalletCallerSession) NumETHDeposits() (*big.Int, error) {
	return _StakingWallet.Contract.NumETHDeposits(&_StakingWallet.CallOpts)
}

// NumStakeActions is a free data retrieval call binding the contract method 0xa531aa86.
//
// Solidity: function numStakeActions() view returns(uint256)
func (_StakingWallet *StakingWalletCaller) NumStakeActions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWallet.contract.Call(opts, &out, "numStakeActions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumStakeActions is a free data retrieval call binding the contract method 0xa531aa86.
//
// Solidity: function numStakeActions() view returns(uint256)
func (_StakingWallet *StakingWalletSession) NumStakeActions() (*big.Int, error) {
	return _StakingWallet.Contract.NumStakeActions(&_StakingWallet.CallOpts)
}

// NumStakeActions is a free data retrieval call binding the contract method 0xa531aa86.
//
// Solidity: function numStakeActions() view returns(uint256)
func (_StakingWallet *StakingWalletCallerSession) NumStakeActions() (*big.Int, error) {
	return _StakingWallet.Contract.NumStakeActions(&_StakingWallet.CallOpts)
}

// NumStakedNFTs is a free data retrieval call binding the contract method 0x17db6213.
//
// Solidity: function numStakedNFTs() view returns(uint256)
func (_StakingWallet *StakingWalletCaller) NumStakedNFTs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWallet.contract.Call(opts, &out, "numStakedNFTs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumStakedNFTs is a free data retrieval call binding the contract method 0x17db6213.
//
// Solidity: function numStakedNFTs() view returns(uint256)
func (_StakingWallet *StakingWalletSession) NumStakedNFTs() (*big.Int, error) {
	return _StakingWallet.Contract.NumStakedNFTs(&_StakingWallet.CallOpts)
}

// NumStakedNFTs is a free data retrieval call binding the contract method 0x17db6213.
//
// Solidity: function numStakedNFTs() view returns(uint256)
func (_StakingWallet *StakingWalletCallerSession) NumStakedNFTs() (*big.Int, error) {
	return _StakingWallet.Contract.NumStakedNFTs(&_StakingWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingWallet *StakingWalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingWallet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingWallet *StakingWalletSession) Owner() (common.Address, error) {
	return _StakingWallet.Contract.Owner(&_StakingWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingWallet *StakingWalletCallerSession) Owner() (common.Address, error) {
	return _StakingWallet.Contract.Owner(&_StakingWallet.CallOpts)
}

// StakedNFTs is a free data retrieval call binding the contract method 0xb865749d.
//
// Solidity: function stakedNFTs(uint256 ) view returns(uint256 tokenId, address owner, uint256 stakeTime, uint256 unstakeTime, uint256 unstakeEligibleTime)
func (_StakingWallet *StakingWalletCaller) StakedNFTs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TokenId             *big.Int
	Owner               common.Address
	StakeTime           *big.Int
	UnstakeTime         *big.Int
	UnstakeEligibleTime *big.Int
}, error) {
	var out []interface{}
	err := _StakingWallet.contract.Call(opts, &out, "stakedNFTs", arg0)

	outstruct := new(struct {
		TokenId             *big.Int
		Owner               common.Address
		StakeTime           *big.Int
		UnstakeTime         *big.Int
		UnstakeEligibleTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.StakeTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UnstakeTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.UnstakeEligibleTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// StakedNFTs is a free data retrieval call binding the contract method 0xb865749d.
//
// Solidity: function stakedNFTs(uint256 ) view returns(uint256 tokenId, address owner, uint256 stakeTime, uint256 unstakeTime, uint256 unstakeEligibleTime)
func (_StakingWallet *StakingWalletSession) StakedNFTs(arg0 *big.Int) (struct {
	TokenId             *big.Int
	Owner               common.Address
	StakeTime           *big.Int
	UnstakeTime         *big.Int
	UnstakeEligibleTime *big.Int
}, error) {
	return _StakingWallet.Contract.StakedNFTs(&_StakingWallet.CallOpts, arg0)
}

// StakedNFTs is a free data retrieval call binding the contract method 0xb865749d.
//
// Solidity: function stakedNFTs(uint256 ) view returns(uint256 tokenId, address owner, uint256 stakeTime, uint256 unstakeTime, uint256 unstakeEligibleTime)
func (_StakingWallet *StakingWalletCallerSession) StakedNFTs(arg0 *big.Int) (struct {
	TokenId             *big.Int
	Owner               common.Address
	StakeTime           *big.Int
	UnstakeTime         *big.Int
	UnstakeEligibleTime *big.Int
}, error) {
	return _StakingWallet.Contract.StakedNFTs(&_StakingWallet.CallOpts, arg0)
}

// ClaimManyRewards is a paid mutator transaction binding the contract method 0x6034eb5b.
//
// Solidity: function claimManyRewards(uint256[] actions, uint256[] deposits) returns()
func (_StakingWallet *StakingWalletTransactor) ClaimManyRewards(opts *bind.TransactOpts, actions []*big.Int, deposits []*big.Int) (*types.Transaction, error) {
	return _StakingWallet.contract.Transact(opts, "claimManyRewards", actions, deposits)
}

// ClaimManyRewards is a paid mutator transaction binding the contract method 0x6034eb5b.
//
// Solidity: function claimManyRewards(uint256[] actions, uint256[] deposits) returns()
func (_StakingWallet *StakingWalletSession) ClaimManyRewards(actions []*big.Int, deposits []*big.Int) (*types.Transaction, error) {
	return _StakingWallet.Contract.ClaimManyRewards(&_StakingWallet.TransactOpts, actions, deposits)
}

// ClaimManyRewards is a paid mutator transaction binding the contract method 0x6034eb5b.
//
// Solidity: function claimManyRewards(uint256[] actions, uint256[] deposits) returns()
func (_StakingWallet *StakingWalletTransactorSession) ClaimManyRewards(actions []*big.Int, deposits []*big.Int) (*types.Transaction, error) {
	return _StakingWallet.Contract.ClaimManyRewards(&_StakingWallet.TransactOpts, actions, deposits)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x86bb8f37.
//
// Solidity: function claimReward(uint256 stakeActionId, uint256 ETHDepositId) returns()
func (_StakingWallet *StakingWalletTransactor) ClaimReward(opts *bind.TransactOpts, stakeActionId *big.Int, ETHDepositId *big.Int) (*types.Transaction, error) {
	return _StakingWallet.contract.Transact(opts, "claimReward", stakeActionId, ETHDepositId)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x86bb8f37.
//
// Solidity: function claimReward(uint256 stakeActionId, uint256 ETHDepositId) returns()
func (_StakingWallet *StakingWalletSession) ClaimReward(stakeActionId *big.Int, ETHDepositId *big.Int) (*types.Transaction, error) {
	return _StakingWallet.Contract.ClaimReward(&_StakingWallet.TransactOpts, stakeActionId, ETHDepositId)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x86bb8f37.
//
// Solidity: function claimReward(uint256 stakeActionId, uint256 ETHDepositId) returns()
func (_StakingWallet *StakingWalletTransactorSession) ClaimReward(stakeActionId *big.Int, ETHDepositId *big.Int) (*types.Transaction, error) {
	return _StakingWallet.Contract.ClaimReward(&_StakingWallet.TransactOpts, stakeActionId, ETHDepositId)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 timestamp) payable returns()
func (_StakingWallet *StakingWalletTransactor) Deposit(opts *bind.TransactOpts, timestamp *big.Int) (*types.Transaction, error) {
	return _StakingWallet.contract.Transact(opts, "deposit", timestamp)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 timestamp) payable returns()
func (_StakingWallet *StakingWalletSession) Deposit(timestamp *big.Int) (*types.Transaction, error) {
	return _StakingWallet.Contract.Deposit(&_StakingWallet.TransactOpts, timestamp)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 timestamp) payable returns()
func (_StakingWallet *StakingWalletTransactorSession) Deposit(timestamp *big.Int) (*types.Transaction, error) {
	return _StakingWallet.Contract.Deposit(&_StakingWallet.TransactOpts, timestamp)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingWallet *StakingWalletTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingWallet.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingWallet *StakingWalletSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingWallet.Contract.RenounceOwnership(&_StakingWallet.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingWallet *StakingWalletTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingWallet.Contract.RenounceOwnership(&_StakingWallet.TransactOpts)
}

// SetCharity is a paid mutator transaction binding the contract method 0xfb6f71a3.
//
// Solidity: function setCharity(address newCharityAddress) returns()
func (_StakingWallet *StakingWalletTransactor) SetCharity(opts *bind.TransactOpts, newCharityAddress common.Address) (*types.Transaction, error) {
	return _StakingWallet.contract.Transact(opts, "setCharity", newCharityAddress)
}

// SetCharity is a paid mutator transaction binding the contract method 0xfb6f71a3.
//
// Solidity: function setCharity(address newCharityAddress) returns()
func (_StakingWallet *StakingWalletSession) SetCharity(newCharityAddress common.Address) (*types.Transaction, error) {
	return _StakingWallet.Contract.SetCharity(&_StakingWallet.TransactOpts, newCharityAddress)
}

// SetCharity is a paid mutator transaction binding the contract method 0xfb6f71a3.
//
// Solidity: function setCharity(address newCharityAddress) returns()
func (_StakingWallet *StakingWalletTransactorSession) SetCharity(newCharityAddress common.Address) (*types.Transaction, error) {
	return _StakingWallet.Contract.SetCharity(&_StakingWallet.TransactOpts, newCharityAddress)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _tokenId) returns()
func (_StakingWallet *StakingWalletTransactor) Stake(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _StakingWallet.contract.Transact(opts, "stake", _tokenId)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _tokenId) returns()
func (_StakingWallet *StakingWalletSession) Stake(_tokenId *big.Int) (*types.Transaction, error) {
	return _StakingWallet.Contract.Stake(&_StakingWallet.TransactOpts, _tokenId)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _tokenId) returns()
func (_StakingWallet *StakingWalletTransactorSession) Stake(_tokenId *big.Int) (*types.Transaction, error) {
	return _StakingWallet.Contract.Stake(&_StakingWallet.TransactOpts, _tokenId)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] ids) returns()
func (_StakingWallet *StakingWalletTransactor) StakeMany(opts *bind.TransactOpts, ids []*big.Int) (*types.Transaction, error) {
	return _StakingWallet.contract.Transact(opts, "stakeMany", ids)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] ids) returns()
func (_StakingWallet *StakingWalletSession) StakeMany(ids []*big.Int) (*types.Transaction, error) {
	return _StakingWallet.Contract.StakeMany(&_StakingWallet.TransactOpts, ids)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] ids) returns()
func (_StakingWallet *StakingWalletTransactorSession) StakeMany(ids []*big.Int) (*types.Transaction, error) {
	return _StakingWallet.Contract.StakeMany(&_StakingWallet.TransactOpts, ids)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingWallet *StakingWalletTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakingWallet.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingWallet *StakingWalletSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingWallet.Contract.TransferOwnership(&_StakingWallet.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingWallet *StakingWalletTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingWallet.Contract.TransferOwnership(&_StakingWallet.TransactOpts, newOwner)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId) returns()
func (_StakingWallet *StakingWalletTransactor) Unstake(opts *bind.TransactOpts, stakeActionId *big.Int) (*types.Transaction, error) {
	return _StakingWallet.contract.Transact(opts, "unstake", stakeActionId)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId) returns()
func (_StakingWallet *StakingWalletSession) Unstake(stakeActionId *big.Int) (*types.Transaction, error) {
	return _StakingWallet.Contract.Unstake(&_StakingWallet.TransactOpts, stakeActionId)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId) returns()
func (_StakingWallet *StakingWalletTransactorSession) Unstake(stakeActionId *big.Int) (*types.Transaction, error) {
	return _StakingWallet.Contract.Unstake(&_StakingWallet.TransactOpts, stakeActionId)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] ids) returns()
func (_StakingWallet *StakingWalletTransactor) UnstakeMany(opts *bind.TransactOpts, ids []*big.Int) (*types.Transaction, error) {
	return _StakingWallet.contract.Transact(opts, "unstakeMany", ids)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] ids) returns()
func (_StakingWallet *StakingWalletSession) UnstakeMany(ids []*big.Int) (*types.Transaction, error) {
	return _StakingWallet.Contract.UnstakeMany(&_StakingWallet.TransactOpts, ids)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] ids) returns()
func (_StakingWallet *StakingWalletTransactorSession) UnstakeMany(ids []*big.Int) (*types.Transaction, error) {
	return _StakingWallet.Contract.UnstakeMany(&_StakingWallet.TransactOpts, ids)
}

// StakingWalletCharityUpdatedEventIterator is returned from FilterCharityUpdatedEvent and is used to iterate over the raw logs and unpacked data for CharityUpdatedEvent events raised by the StakingWallet contract.
type StakingWalletCharityUpdatedEventIterator struct {
	Event *StakingWalletCharityUpdatedEvent // Event containing the contract specifics and raw log

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
func (it *StakingWalletCharityUpdatedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletCharityUpdatedEvent)
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
		it.Event = new(StakingWalletCharityUpdatedEvent)
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
func (it *StakingWalletCharityUpdatedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletCharityUpdatedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletCharityUpdatedEvent represents a CharityUpdatedEvent event raised by the StakingWallet contract.
type StakingWalletCharityUpdatedEvent struct {
	NewCharityAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCharityUpdatedEvent is a free log retrieval operation binding the contract event 0xa0bd6b2fdbf082ae2356710c23fc8d76d56d418cecb4514d119c77a8617b4ffe.
//
// Solidity: event CharityUpdatedEvent(address indexed newCharityAddress)
func (_StakingWallet *StakingWalletFilterer) FilterCharityUpdatedEvent(opts *bind.FilterOpts, newCharityAddress []common.Address) (*StakingWalletCharityUpdatedEventIterator, error) {

	var newCharityAddressRule []interface{}
	for _, newCharityAddressItem := range newCharityAddress {
		newCharityAddressRule = append(newCharityAddressRule, newCharityAddressItem)
	}

	logs, sub, err := _StakingWallet.contract.FilterLogs(opts, "CharityUpdatedEvent", newCharityAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletCharityUpdatedEventIterator{contract: _StakingWallet.contract, event: "CharityUpdatedEvent", logs: logs, sub: sub}, nil
}

// WatchCharityUpdatedEvent is a free log subscription operation binding the contract event 0xa0bd6b2fdbf082ae2356710c23fc8d76d56d418cecb4514d119c77a8617b4ffe.
//
// Solidity: event CharityUpdatedEvent(address indexed newCharityAddress)
func (_StakingWallet *StakingWalletFilterer) WatchCharityUpdatedEvent(opts *bind.WatchOpts, sink chan<- *StakingWalletCharityUpdatedEvent, newCharityAddress []common.Address) (event.Subscription, error) {

	var newCharityAddressRule []interface{}
	for _, newCharityAddressItem := range newCharityAddress {
		newCharityAddressRule = append(newCharityAddressRule, newCharityAddressItem)
	}

	logs, sub, err := _StakingWallet.contract.WatchLogs(opts, "CharityUpdatedEvent", newCharityAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletCharityUpdatedEvent)
				if err := _StakingWallet.contract.UnpackLog(event, "CharityUpdatedEvent", log); err != nil {
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

// ParseCharityUpdatedEvent is a log parse operation binding the contract event 0xa0bd6b2fdbf082ae2356710c23fc8d76d56d418cecb4514d119c77a8617b4ffe.
//
// Solidity: event CharityUpdatedEvent(address indexed newCharityAddress)
func (_StakingWallet *StakingWalletFilterer) ParseCharityUpdatedEvent(log types.Log) (*StakingWalletCharityUpdatedEvent, error) {
	event := new(StakingWalletCharityUpdatedEvent)
	if err := _StakingWallet.contract.UnpackLog(event, "CharityUpdatedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletClaimRewardEventIterator is returned from FilterClaimRewardEvent and is used to iterate over the raw logs and unpacked data for ClaimRewardEvent events raised by the StakingWallet contract.
type StakingWalletClaimRewardEventIterator struct {
	Event *StakingWalletClaimRewardEvent // Event containing the contract specifics and raw log

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
func (it *StakingWalletClaimRewardEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletClaimRewardEvent)
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
		it.Event = new(StakingWalletClaimRewardEvent)
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
func (it *StakingWalletClaimRewardEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletClaimRewardEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletClaimRewardEvent represents a ClaimRewardEvent event raised by the StakingWallet contract.
type StakingWalletClaimRewardEvent struct {
	ActionId  *big.Int
	DepositId *big.Int
	Reward    *big.Int
	Staker    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimRewardEvent is a free log retrieval operation binding the contract event 0xdde81df5caa033e783e2a39d93e3a8718a7dc27ba95a4757f5433a01f794ec36.
//
// Solidity: event ClaimRewardEvent(uint256 indexed actionId, uint256 indexed depositId, uint256 reward, address indexed staker)
func (_StakingWallet *StakingWalletFilterer) FilterClaimRewardEvent(opts *bind.FilterOpts, actionId []*big.Int, depositId []*big.Int, staker []common.Address) (*StakingWalletClaimRewardEventIterator, error) {

	var actionIdRule []interface{}
	for _, actionIdItem := range actionId {
		actionIdRule = append(actionIdRule, actionIdItem)
	}
	var depositIdRule []interface{}
	for _, depositIdItem := range depositId {
		depositIdRule = append(depositIdRule, depositIdItem)
	}

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakingWallet.contract.FilterLogs(opts, "ClaimRewardEvent", actionIdRule, depositIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletClaimRewardEventIterator{contract: _StakingWallet.contract, event: "ClaimRewardEvent", logs: logs, sub: sub}, nil
}

// WatchClaimRewardEvent is a free log subscription operation binding the contract event 0xdde81df5caa033e783e2a39d93e3a8718a7dc27ba95a4757f5433a01f794ec36.
//
// Solidity: event ClaimRewardEvent(uint256 indexed actionId, uint256 indexed depositId, uint256 reward, address indexed staker)
func (_StakingWallet *StakingWalletFilterer) WatchClaimRewardEvent(opts *bind.WatchOpts, sink chan<- *StakingWalletClaimRewardEvent, actionId []*big.Int, depositId []*big.Int, staker []common.Address) (event.Subscription, error) {

	var actionIdRule []interface{}
	for _, actionIdItem := range actionId {
		actionIdRule = append(actionIdRule, actionIdItem)
	}
	var depositIdRule []interface{}
	for _, depositIdItem := range depositId {
		depositIdRule = append(depositIdRule, depositIdItem)
	}

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakingWallet.contract.WatchLogs(opts, "ClaimRewardEvent", actionIdRule, depositIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletClaimRewardEvent)
				if err := _StakingWallet.contract.UnpackLog(event, "ClaimRewardEvent", log); err != nil {
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

// ParseClaimRewardEvent is a log parse operation binding the contract event 0xdde81df5caa033e783e2a39d93e3a8718a7dc27ba95a4757f5433a01f794ec36.
//
// Solidity: event ClaimRewardEvent(uint256 indexed actionId, uint256 indexed depositId, uint256 reward, address indexed staker)
func (_StakingWallet *StakingWalletFilterer) ParseClaimRewardEvent(log types.Log) (*StakingWalletClaimRewardEvent, error) {
	event := new(StakingWalletClaimRewardEvent)
	if err := _StakingWallet.contract.UnpackLog(event, "ClaimRewardEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletEthDepositEventIterator is returned from FilterEthDepositEvent and is used to iterate over the raw logs and unpacked data for EthDepositEvent events raised by the StakingWallet contract.
type StakingWalletEthDepositEventIterator struct {
	Event *StakingWalletEthDepositEvent // Event containing the contract specifics and raw log

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
func (it *StakingWalletEthDepositEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletEthDepositEvent)
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
		it.Event = new(StakingWalletEthDepositEvent)
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
func (it *StakingWalletEthDepositEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletEthDepositEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletEthDepositEvent represents a EthDepositEvent event raised by the StakingWallet contract.
type StakingWalletEthDepositEvent struct {
	DepositTime   *big.Int
	DepositNum    *big.Int
	NumStakedNFTs *big.Int
	Amount        *big.Int
	Modulo        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEthDepositEvent is a free log retrieval operation binding the contract event 0xdc0eacba8b1f88284dca5eec8be23173aefa7206298fe22de43e064b6ccd8418.
//
// Solidity: event EthDepositEvent(uint256 indexed depositTime, uint256 depositNum, uint256 numStakedNFTs, uint256 amount, uint256 modulo)
func (_StakingWallet *StakingWalletFilterer) FilterEthDepositEvent(opts *bind.FilterOpts, depositTime []*big.Int) (*StakingWalletEthDepositEventIterator, error) {

	var depositTimeRule []interface{}
	for _, depositTimeItem := range depositTime {
		depositTimeRule = append(depositTimeRule, depositTimeItem)
	}

	logs, sub, err := _StakingWallet.contract.FilterLogs(opts, "EthDepositEvent", depositTimeRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletEthDepositEventIterator{contract: _StakingWallet.contract, event: "EthDepositEvent", logs: logs, sub: sub}, nil
}

// WatchEthDepositEvent is a free log subscription operation binding the contract event 0xdc0eacba8b1f88284dca5eec8be23173aefa7206298fe22de43e064b6ccd8418.
//
// Solidity: event EthDepositEvent(uint256 indexed depositTime, uint256 depositNum, uint256 numStakedNFTs, uint256 amount, uint256 modulo)
func (_StakingWallet *StakingWalletFilterer) WatchEthDepositEvent(opts *bind.WatchOpts, sink chan<- *StakingWalletEthDepositEvent, depositTime []*big.Int) (event.Subscription, error) {

	var depositTimeRule []interface{}
	for _, depositTimeItem := range depositTime {
		depositTimeRule = append(depositTimeRule, depositTimeItem)
	}

	logs, sub, err := _StakingWallet.contract.WatchLogs(opts, "EthDepositEvent", depositTimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletEthDepositEvent)
				if err := _StakingWallet.contract.UnpackLog(event, "EthDepositEvent", log); err != nil {
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

// ParseEthDepositEvent is a log parse operation binding the contract event 0xdc0eacba8b1f88284dca5eec8be23173aefa7206298fe22de43e064b6ccd8418.
//
// Solidity: event EthDepositEvent(uint256 indexed depositTime, uint256 depositNum, uint256 numStakedNFTs, uint256 amount, uint256 modulo)
func (_StakingWallet *StakingWalletFilterer) ParseEthDepositEvent(log types.Log) (*StakingWalletEthDepositEvent, error) {
	event := new(StakingWalletEthDepositEvent)
	if err := _StakingWallet.contract.UnpackLog(event, "EthDepositEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakingWallet contract.
type StakingWalletOwnershipTransferredIterator struct {
	Event *StakingWalletOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingWalletOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletOwnershipTransferred)
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
		it.Event = new(StakingWalletOwnershipTransferred)
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
func (it *StakingWalletOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletOwnershipTransferred represents a OwnershipTransferred event raised by the StakingWallet contract.
type StakingWalletOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingWallet *StakingWalletFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingWalletOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingWallet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletOwnershipTransferredIterator{contract: _StakingWallet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingWallet *StakingWalletFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingWalletOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingWallet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletOwnershipTransferred)
				if err := _StakingWallet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StakingWallet *StakingWalletFilterer) ParseOwnershipTransferred(log types.Log) (*StakingWalletOwnershipTransferred, error) {
	event := new(StakingWalletOwnershipTransferred)
	if err := _StakingWallet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletStakeActionEventIterator is returned from FilterStakeActionEvent and is used to iterate over the raw logs and unpacked data for StakeActionEvent events raised by the StakingWallet contract.
type StakingWalletStakeActionEventIterator struct {
	Event *StakingWalletStakeActionEvent // Event containing the contract specifics and raw log

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
func (it *StakingWalletStakeActionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletStakeActionEvent)
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
		it.Event = new(StakingWalletStakeActionEvent)
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
func (it *StakingWalletStakeActionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletStakeActionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletStakeActionEvent represents a StakeActionEvent event raised by the StakingWallet contract.
type StakingWalletStakeActionEvent struct {
	ActionId    *big.Int
	TokenId     *big.Int
	TotalNFTs   *big.Int
	UnstakeTime *big.Int
	Staker      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStakeActionEvent is a free log retrieval operation binding the contract event 0x057eba8c4bba00f858e4f586f9c02794abb0df789ef316c741f9073fe2c435db.
//
// Solidity: event StakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, uint256 unstakeTime, address indexed staker)
func (_StakingWallet *StakingWalletFilterer) FilterStakeActionEvent(opts *bind.FilterOpts, actionId []*big.Int, tokenId []*big.Int, staker []common.Address) (*StakingWalletStakeActionEventIterator, error) {

	var actionIdRule []interface{}
	for _, actionIdItem := range actionId {
		actionIdRule = append(actionIdRule, actionIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakingWallet.contract.FilterLogs(opts, "StakeActionEvent", actionIdRule, tokenIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletStakeActionEventIterator{contract: _StakingWallet.contract, event: "StakeActionEvent", logs: logs, sub: sub}, nil
}

// WatchStakeActionEvent is a free log subscription operation binding the contract event 0x057eba8c4bba00f858e4f586f9c02794abb0df789ef316c741f9073fe2c435db.
//
// Solidity: event StakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, uint256 unstakeTime, address indexed staker)
func (_StakingWallet *StakingWalletFilterer) WatchStakeActionEvent(opts *bind.WatchOpts, sink chan<- *StakingWalletStakeActionEvent, actionId []*big.Int, tokenId []*big.Int, staker []common.Address) (event.Subscription, error) {

	var actionIdRule []interface{}
	for _, actionIdItem := range actionId {
		actionIdRule = append(actionIdRule, actionIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakingWallet.contract.WatchLogs(opts, "StakeActionEvent", actionIdRule, tokenIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletStakeActionEvent)
				if err := _StakingWallet.contract.UnpackLog(event, "StakeActionEvent", log); err != nil {
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

// ParseStakeActionEvent is a log parse operation binding the contract event 0x057eba8c4bba00f858e4f586f9c02794abb0df789ef316c741f9073fe2c435db.
//
// Solidity: event StakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, uint256 unstakeTime, address indexed staker)
func (_StakingWallet *StakingWalletFilterer) ParseStakeActionEvent(log types.Log) (*StakingWalletStakeActionEvent, error) {
	event := new(StakingWalletStakeActionEvent)
	if err := _StakingWallet.contract.UnpackLog(event, "StakeActionEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletUnstakeActionEventIterator is returned from FilterUnstakeActionEvent and is used to iterate over the raw logs and unpacked data for UnstakeActionEvent events raised by the StakingWallet contract.
type StakingWalletUnstakeActionEventIterator struct {
	Event *StakingWalletUnstakeActionEvent // Event containing the contract specifics and raw log

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
func (it *StakingWalletUnstakeActionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletUnstakeActionEvent)
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
		it.Event = new(StakingWalletUnstakeActionEvent)
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
func (it *StakingWalletUnstakeActionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletUnstakeActionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletUnstakeActionEvent represents a UnstakeActionEvent event raised by the StakingWallet contract.
type StakingWalletUnstakeActionEvent struct {
	ActionId  *big.Int
	TokenId   *big.Int
	TotalNFTs *big.Int
	Staker    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnstakeActionEvent is a free log retrieval operation binding the contract event 0x33940a9fc3ee3c9150b29b34ef29ca397b5e1e340425a4f0da0346b5b90766c8.
//
// Solidity: event UnstakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address indexed staker)
func (_StakingWallet *StakingWalletFilterer) FilterUnstakeActionEvent(opts *bind.FilterOpts, actionId []*big.Int, tokenId []*big.Int, staker []common.Address) (*StakingWalletUnstakeActionEventIterator, error) {

	var actionIdRule []interface{}
	for _, actionIdItem := range actionId {
		actionIdRule = append(actionIdRule, actionIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakingWallet.contract.FilterLogs(opts, "UnstakeActionEvent", actionIdRule, tokenIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletUnstakeActionEventIterator{contract: _StakingWallet.contract, event: "UnstakeActionEvent", logs: logs, sub: sub}, nil
}

// WatchUnstakeActionEvent is a free log subscription operation binding the contract event 0x33940a9fc3ee3c9150b29b34ef29ca397b5e1e340425a4f0da0346b5b90766c8.
//
// Solidity: event UnstakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address indexed staker)
func (_StakingWallet *StakingWalletFilterer) WatchUnstakeActionEvent(opts *bind.WatchOpts, sink chan<- *StakingWalletUnstakeActionEvent, actionId []*big.Int, tokenId []*big.Int, staker []common.Address) (event.Subscription, error) {

	var actionIdRule []interface{}
	for _, actionIdItem := range actionId {
		actionIdRule = append(actionIdRule, actionIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakingWallet.contract.WatchLogs(opts, "UnstakeActionEvent", actionIdRule, tokenIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletUnstakeActionEvent)
				if err := _StakingWallet.contract.UnpackLog(event, "UnstakeActionEvent", log); err != nil {
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

// ParseUnstakeActionEvent is a log parse operation binding the contract event 0x33940a9fc3ee3c9150b29b34ef29ca397b5e1e340425a4f0da0346b5b90766c8.
//
// Solidity: event UnstakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address indexed staker)
func (_StakingWallet *StakingWalletFilterer) ParseUnstakeActionEvent(log types.Log) (*StakingWalletUnstakeActionEvent, error) {
	event := new(StakingWalletUnstakeActionEvent)
	if err := _StakingWallet.contract.UnpackLog(event, "UnstakeActionEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
