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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractCosmicSignature\",\"name\":\"nft_\",\"type\":\"address\"},{\"internalType\":\"contractCosmicGame\",\"name\":\"game_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"charity_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newCharityAddress\",\"type\":\"address\"}],\"name\":\"CharityUpdatedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"ClaimRewardEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"depositTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNFTs\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"modulo\",\"type\":\"uint256\"}],\"name\":\"EthDepositEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalNFTs\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unstakeTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"StakeActionEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalNFTs\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"}],\"name\":\"UnstakeActionEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ETHDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"depositTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numStaked\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ETHDepositId\",\"type\":\"uint256\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"game\",\"outputs\":[{\"internalType\":\"contractCosmicGame\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nft\",\"outputs\":[{\"internalType\":\"contractCosmicSignature\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numETHDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakeActions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCharityAddress\",\"type\":\"address\"}],\"name\":\"setCharity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakedNFTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeEligibleTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620022eb380380620022eb8339818101604052810190620000379190620002f7565b620000576200004b6200012360201b60201c565b6200012b60201b60201c565b82600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505062000353565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200022182620001f4565b9050919050565b6000620002358262000214565b9050919050565b620002478162000228565b81146200025357600080fd5b50565b60008151905062000267816200023c565b92915050565b60006200027a82620001f4565b9050919050565b60006200028e826200026d565b9050919050565b620002a08162000281565b8114620002ac57600080fd5b50565b600081519050620002c08162000295565b92915050565b620002d18162000214565b8114620002dd57600080fd5b50565b600081519050620002f181620002c6565b92915050565b600080600060608486031215620003135762000312620001ef565b5b6000620003238682870162000256565b93505060206200033686828701620002af565b92505060406200034986828701620002e0565b9150509250925092565b611f8880620003636000396000f3fe6080604052600436106100dd5760003560e01c8063a531aa861161007f578063b865749d11610059578063b865749d1461027b578063c3fe3e28146102bc578063f2fde38b146102e7578063fb6f71a314610310576100dd565b8063a531aa861461020b578063a694fc3a14610236578063b6b55f251461025f576100dd565b806355279fdb116100bb57806355279fdb14610175578063715018a6146101a057806386bb8f37146101b75780638da5cb5b146101e0576100dd565b80632e17de78146100e2578063451f1adf1461010b57806347ccca021461014a575b600080fd5b3480156100ee57600080fd5b50610109600480360381019061010491906113e7565b610339565b005b34801561011757600080fd5b50610132600480360381019061012d91906113e7565b6105c2565b60405161014193929190611423565b60405180910390f35b34801561015657600080fd5b5061015f6105ec565b60405161016c91906114d9565b60405180910390f35b34801561018157600080fd5b5061018a610612565b60405161019791906114f4565b60405180910390f35b3480156101ac57600080fd5b506101b5610618565b005b3480156101c357600080fd5b506101de60048036038101906101d9919061150f565b6106a0565b005b3480156101ec57600080fd5b506101f5610a8c565b6040516102029190611570565b60405180910390f35b34801561021757600080fd5b50610220610ab5565b60405161022d91906114f4565b60405180910390f35b34801561024257600080fd5b5061025d600480360381019061025891906113e7565b610abb565b005b610279600480360381019061027491906113e7565b610d7e565b005b34801561028757600080fd5b506102a2600480360381019061029d91906113e7565b610fd9565b6040516102b395949392919061158b565b60405180910390f35b3480156102c857600080fd5b506102d161102f565b6040516102de9190611611565b60405180910390f35b3480156102f357600080fd5b5061030e60048036038101906103099190611658565b611055565b005b34801561031c57600080fd5b5061033760048036038101906103329190611658565b61114c565b005b6000600160008381526020019081526020016000206003015414610392576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610389906116e2565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff166001600083815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610436576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161042d9061174e565b60405180910390fd5b4260016000838152602001908152602001600020600401541061048e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610485906117ba565b60405180910390fd5b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166342842e0e3033600160006002548152602001908152602001600020600001546040518463ffffffff1660e01b8152600401610505939291906117da565b600060405180830381600087803b15801561051f57600080fd5b505af1158015610533573d6000803e3d6000fd5b505050504260016000838152602001908152602001600020600301819055506001600560008282546105659190611840565b925050819055506001600082815260200190815260200160002060000154817f33940a9fc3ee3c9150b29b34ef29ca397b5e1e340425a4f0da0346b5b90766c8600554336040516105b7929190611874565b60405180910390a350565b60036020528060005260406000206000915090508060000154908060010154908060020154905083565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60045481565b6106206112e0565b73ffffffffffffffffffffffffffffffffffffffff1661063e610a8c565b73ffffffffffffffffffffffffffffffffffffffff1614610694576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161068b906118e9565b60405180910390fd5b61069e60006112e8565b565b60006001600084815260200190815260200160002060030154116106f9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106f090611955565b60405180910390fd5b60016000838152602001908152602001600020600501600082815260200190815260200160002060009054906101000a900460ff161561076e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076590611955565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff166001600084815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610812576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610809906119c1565b60405180910390fd5b6003600082815260200190815260200160002060000154600160008481526020019081526020016000206002015410610880576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161087790611a2d565b60405180910390fd5b60036000828152602001908152602001600020600001546001600084815260200190815260200160002060030154116108ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108e590611a99565b60405180910390fd5b6001806000848152602001908152602001600020600501600083815260200190815260200160002060006101000a81548160ff0219169083151502179055506000600360008381526020019081526020016000206002015460036000848152602001908152602001600020600101546109679190611ae8565b905060006001600085815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16826040516109c590611b4a565b60006040518083038185875af1925050503d8060008114610a02576040519150601f19603f3d011682016040523d82523d6000602084013e610a07565b606091505b5050905080610a4b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a4290611bab565b60405180910390fd5b82847fdde81df5caa033e783e2a39d93e3a8718a7dc27ba95a4757f5433a01f794ec368433604051610a7e929190611874565b60405180910390a350505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60025481565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166342842e0e3330846040518463ffffffff1660e01b8152600401610b1a939291906117da565b600060405180830381600087803b158015610b3457600080fd5b505af1158015610b48573d6000803e3d6000fd5b5050505080600160006002548152602001908152602001600020600001819055503360016000600254815260200190815260200160002060010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555042600160006002548152602001908152602001600020600201819055506000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c4c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c709190611be0565b600554620f4240610c819190611c0d565b610c8b9190611ae8565b905060006227ac406105dc8384610ca29190611c0d565b610cac9190611ae8565b610cb69190611c4f565b90508042610cc49190611c4f565b60016000600254815260200190815260200160002060040181905550600160026000828254610cf39190611c4f565b92505081905550600160056000828254610d0d9190611c4f565b92505081905550826001600254610d249190611840565b7f057eba8c4bba00f858e4f586f9c02794abb0df789ef316c741f9073fe2c435db6005546001600060025481526020019081526020016000206004015433604051610d7193929190611c83565b60405180910390a3505050565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610e0e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e0590611d2c565b60405180910390fd5b600060055403610eec576000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1634604051610e6090611b4a565b60006040518083038185875af1925050503d8060008114610e9d576040519150601f19603f3d011682016040523d82523d6000602084013e610ea2565b606091505b5050905080610ee6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610edd90611dbe565b60405180910390fd5b50610fd6565b8060036000600454815260200190815260200160002060000181905550346003600060045481526020019081526020016000206001018190555060055460036000600454815260200190815260200160002060020181905550600160046000828254610f589190611c4f565b9250508190555060055434610f6d9190611dde565b60076000828254610f7e9190611c4f565b92505081905550807fdc0eacba8b1f88284dca5eec8be23173aefa7206298fe22de43e064b6ccd84186001600454610fb69190611840565b60055434600754604051610fcd9493929190611e0f565b60405180910390a25b50565b60016020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154908060030154908060040154905085565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61105d6112e0565b73ffffffffffffffffffffffffffffffffffffffff1661107b610a8c565b73ffffffffffffffffffffffffffffffffffffffff16146110d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110c8906118e9565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611140576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161113790611ec6565b60405180910390fd5b611149816112e8565b50565b6111546112e0565b73ffffffffffffffffffffffffffffffffffffffff16611172610a8c565b73ffffffffffffffffffffffffffffffffffffffff16146111c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111bf906118e9565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611237576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161122e90611f32565b60405180910390fd5b80600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fa0bd6b2fdbf082ae2356710c23fc8d76d56d418cecb4514d119c77a8617b4ffe60405160405180910390a250565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b6000819050919050565b6113c4816113b1565b81146113cf57600080fd5b50565b6000813590506113e1816113bb565b92915050565b6000602082840312156113fd576113fc6113ac565b5b600061140b848285016113d2565b91505092915050565b61141d816113b1565b82525050565b60006060820190506114386000830186611414565b6114456020830185611414565b6114526040830184611414565b949350505050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061149f61149a6114958461145a565b61147a565b61145a565b9050919050565b60006114b182611484565b9050919050565b60006114c3826114a6565b9050919050565b6114d3816114b8565b82525050565b60006020820190506114ee60008301846114ca565b92915050565b60006020820190506115096000830184611414565b92915050565b60008060408385031215611526576115256113ac565b5b6000611534858286016113d2565b9250506020611545858286016113d2565b9150509250929050565b600061155a8261145a565b9050919050565b61156a8161154f565b82525050565b60006020820190506115856000830184611561565b92915050565b600060a0820190506115a06000830188611414565b6115ad6020830187611561565b6115ba6040830186611414565b6115c76060830185611414565b6115d46080830184611414565b9695505050505050565b60006115e982611484565b9050919050565b60006115fb826115de565b9050919050565b61160b816115f0565b82525050565b60006020820190506116266000830184611602565b92915050565b6116358161154f565b811461164057600080fd5b50565b6000813590506116528161162c565b92915050565b60006020828403121561166e5761166d6113ac565b5b600061167c84828501611643565b91505092915050565b600082825260208201905092915050565b7f546f6b656e2068617320616c7265616479206265656e20756e7374616b656400600082015250565b60006116cc601f83611685565b91506116d782611696565b602082019050919050565b600060208201905081810360008301526116fb816116bf565b9050919050565b7f4f6e6c7920746865206f776e65722063616e20756e7374616b65000000000000600082015250565b6000611738601a83611685565b915061174382611702565b602082019050919050565b600060208201905081810360008301526117678161172b565b9050919050565b7f4e6f7420616c6c6f77656420746f20756e7374616b6520796574000000000000600082015250565b60006117a4601a83611685565b91506117af8261176e565b602082019050919050565b600060208201905081810360008301526117d381611797565b9050919050565b60006060820190506117ef6000830186611561565b6117fc6020830185611561565b6118096040830184611414565b949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061184b826113b1565b9150611856836113b1565b925082820390508181111561186e5761186d611811565b5b92915050565b60006040820190506118896000830185611414565b6118966020830184611561565b9392505050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b60006118d3602083611685565b91506118de8261189d565b602082019050919050565b60006020820190508181036000830152611902816118c6565b9050919050565b7f546f6b656e20686173206e6f74206265656e20756e7374616b65640000000000600082015250565b600061193f601b83611685565b915061194a82611909565b602082019050919050565b6000602082019050818103600083015261196e81611932565b9050919050565b7f4f6e6c7920746865206f776e65722063616e20636c61696d2072657761726400600082015250565b60006119ab601f83611685565b91506119b682611975565b602082019050919050565b600060208201905081810360008301526119da8161199e565b9050919050565b7f596f752077657265206e6f74207374616b6564207965742e0000000000000000600082015250565b6000611a17601883611685565b9150611a22826119e1565b602082019050919050565b60006020820190508181036000830152611a4681611a0a565b9050919050565b7f596f75207765726520616c726561647920756e7374616b65642e000000000000600082015250565b6000611a83601a83611685565b9150611a8e82611a4d565b602082019050919050565b60006020820190508181036000830152611ab281611a76565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611af3826113b1565b9150611afe836113b1565b925082611b0e57611b0d611ab9565b5b828204905092915050565b600081905092915050565b50565b6000611b34600083611b19565b9150611b3f82611b24565b600082019050919050565b6000611b5582611b27565b9150819050919050565b7f526577617264207472616e73666572206661696c65642e000000000000000000600082015250565b6000611b95601783611685565b9150611ba082611b5f565b602082019050919050565b60006020820190508181036000830152611bc481611b88565b9050919050565b600081519050611bda816113bb565b92915050565b600060208284031215611bf657611bf56113ac565b5b6000611c0484828501611bcb565b91505092915050565b6000611c18826113b1565b9150611c23836113b1565b9250828202611c31816113b1565b91508282048414831517611c4857611c47611811565b5b5092915050565b6000611c5a826113b1565b9150611c65836113b1565b9250828201905080821115611c7d57611c7c611811565b5b92915050565b6000606082019050611c986000830186611414565b611ca56020830185611414565b611cb26040830184611561565b949350505050565b7f4f6e6c792074686520436f736d696347616d6520636f6e74726163742063616e60008201527f206465706f7369742e0000000000000000000000000000000000000000000000602082015250565b6000611d16602983611685565b9150611d2182611cba565b604082019050919050565b60006020820190508181036000830152611d4581611d09565b9050919050565b7f5472616e7366657220746f206368617269747920636f6e74726163742066616960008201527f6c65642e00000000000000000000000000000000000000000000000000000000602082015250565b6000611da8602483611685565b9150611db382611d4c565b604082019050919050565b60006020820190508181036000830152611dd781611d9b565b9050919050565b6000611de9826113b1565b9150611df4836113b1565b925082611e0457611e03611ab9565b5b828206905092915050565b6000608082019050611e246000830187611414565b611e316020830186611414565b611e3e6040830185611414565b611e4b6060830184611414565b95945050505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000611eb0602683611685565b9150611ebb82611e54565b604082019050919050565b60006020820190508181036000830152611edf81611ea3565b9050919050565b7f5a65726f2d616464726573732077617320676976656e2e000000000000000000600082015250565b6000611f1c601783611685565b9150611f2782611ee6565b602082019050919050565b60006020820190508181036000830152611f4b81611f0f565b905091905056fea2646970667358221220525fa622c268b493fe33e4d91ebca3f2ccd3913be8aa968ecb3b03ee5d0b8c3064736f6c63430008130033",
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
// Solidity: event ClaimRewardEvent(uint256 indexed actionId, uint256 indexed depositId, uint256 reward, address staker)
func (_StakingWallet *StakingWalletFilterer) FilterClaimRewardEvent(opts *bind.FilterOpts, actionId []*big.Int, depositId []*big.Int) (*StakingWalletClaimRewardEventIterator, error) {

	var actionIdRule []interface{}
	for _, actionIdItem := range actionId {
		actionIdRule = append(actionIdRule, actionIdItem)
	}
	var depositIdRule []interface{}
	for _, depositIdItem := range depositId {
		depositIdRule = append(depositIdRule, depositIdItem)
	}

	logs, sub, err := _StakingWallet.contract.FilterLogs(opts, "ClaimRewardEvent", actionIdRule, depositIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletClaimRewardEventIterator{contract: _StakingWallet.contract, event: "ClaimRewardEvent", logs: logs, sub: sub}, nil
}

// WatchClaimRewardEvent is a free log subscription operation binding the contract event 0xdde81df5caa033e783e2a39d93e3a8718a7dc27ba95a4757f5433a01f794ec36.
//
// Solidity: event ClaimRewardEvent(uint256 indexed actionId, uint256 indexed depositId, uint256 reward, address staker)
func (_StakingWallet *StakingWalletFilterer) WatchClaimRewardEvent(opts *bind.WatchOpts, sink chan<- *StakingWalletClaimRewardEvent, actionId []*big.Int, depositId []*big.Int) (event.Subscription, error) {

	var actionIdRule []interface{}
	for _, actionIdItem := range actionId {
		actionIdRule = append(actionIdRule, actionIdItem)
	}
	var depositIdRule []interface{}
	for _, depositIdItem := range depositId {
		depositIdRule = append(depositIdRule, depositIdItem)
	}

	logs, sub, err := _StakingWallet.contract.WatchLogs(opts, "ClaimRewardEvent", actionIdRule, depositIdRule)
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
// Solidity: event ClaimRewardEvent(uint256 indexed actionId, uint256 indexed depositId, uint256 reward, address staker)
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
// Solidity: event StakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, uint256 unstakeTime, address staker)
func (_StakingWallet *StakingWalletFilterer) FilterStakeActionEvent(opts *bind.FilterOpts, actionId []*big.Int, tokenId []*big.Int) (*StakingWalletStakeActionEventIterator, error) {

	var actionIdRule []interface{}
	for _, actionIdItem := range actionId {
		actionIdRule = append(actionIdRule, actionIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _StakingWallet.contract.FilterLogs(opts, "StakeActionEvent", actionIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletStakeActionEventIterator{contract: _StakingWallet.contract, event: "StakeActionEvent", logs: logs, sub: sub}, nil
}

// WatchStakeActionEvent is a free log subscription operation binding the contract event 0x057eba8c4bba00f858e4f586f9c02794abb0df789ef316c741f9073fe2c435db.
//
// Solidity: event StakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, uint256 unstakeTime, address staker)
func (_StakingWallet *StakingWalletFilterer) WatchStakeActionEvent(opts *bind.WatchOpts, sink chan<- *StakingWalletStakeActionEvent, actionId []*big.Int, tokenId []*big.Int) (event.Subscription, error) {

	var actionIdRule []interface{}
	for _, actionIdItem := range actionId {
		actionIdRule = append(actionIdRule, actionIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _StakingWallet.contract.WatchLogs(opts, "StakeActionEvent", actionIdRule, tokenIdRule)
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
// Solidity: event StakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, uint256 unstakeTime, address staker)
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
	Taker     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnstakeActionEvent is a free log retrieval operation binding the contract event 0x33940a9fc3ee3c9150b29b34ef29ca397b5e1e340425a4f0da0346b5b90766c8.
//
// Solidity: event UnstakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address taker)
func (_StakingWallet *StakingWalletFilterer) FilterUnstakeActionEvent(opts *bind.FilterOpts, actionId []*big.Int, tokenId []*big.Int) (*StakingWalletUnstakeActionEventIterator, error) {

	var actionIdRule []interface{}
	for _, actionIdItem := range actionId {
		actionIdRule = append(actionIdRule, actionIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _StakingWallet.contract.FilterLogs(opts, "UnstakeActionEvent", actionIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletUnstakeActionEventIterator{contract: _StakingWallet.contract, event: "UnstakeActionEvent", logs: logs, sub: sub}, nil
}

// WatchUnstakeActionEvent is a free log subscription operation binding the contract event 0x33940a9fc3ee3c9150b29b34ef29ca397b5e1e340425a4f0da0346b5b90766c8.
//
// Solidity: event UnstakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address taker)
func (_StakingWallet *StakingWalletFilterer) WatchUnstakeActionEvent(opts *bind.WatchOpts, sink chan<- *StakingWalletUnstakeActionEvent, actionId []*big.Int, tokenId []*big.Int) (event.Subscription, error) {

	var actionIdRule []interface{}
	for _, actionIdItem := range actionId {
		actionIdRule = append(actionIdRule, actionIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _StakingWallet.contract.WatchLogs(opts, "UnstakeActionEvent", actionIdRule, tokenIdRule)
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
// Solidity: event UnstakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address taker)
func (_StakingWallet *StakingWalletFilterer) ParseUnstakeActionEvent(log types.Log) (*StakingWalletUnstakeActionEvent, error) {
	event := new(StakingWalletUnstakeActionEvent)
	if err := _StakingWallet.contract.UnpackLog(event, "UnstakeActionEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
