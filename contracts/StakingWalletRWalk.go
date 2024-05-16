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

// StakingWalletRWalkMetaData contains all meta data concerning the StakingWalletRWalk contract.
var StakingWalletRWalkMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractRandomWalkNFT\",\"name\":\"rwalk_\",\"type\":\"address\"},{\"internalType\":\"contractCosmicGame\",\"name\":\"game_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPeriod\",\"type\":\"uint256\"}],\"name\":\"MinStakePeriodChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ModuloSentEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalNFTs\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unstakeTime\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"StakeActionEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalNFTs\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"UnstakeActionEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ETHDeposits\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"depositTime\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numStaked\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"game\",\"outputs\":[{\"internalType\":\"contractCosmicGame\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isTokenStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"lastActionIdByTokenId\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastActionIds\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minStakePeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numETHDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakeActions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakedNFTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numTokensStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"entropy\",\"type\":\"bytes32\"}],\"name\":\"pickRandomStaker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomWalk\",\"outputs\":[{\"internalType\":\"contractRandomWalkNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newStakePeriod\",\"type\":\"uint32\"}],\"name\":\"setMinStakePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeActions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stakeTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"unstakeTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"unstakeEligibleTime\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"stakerByTokenId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenIndex\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIndices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"unstakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405262278d00600960006101000a81548163ffffffff021916908363ffffffff1602179055503480156200003557600080fd5b50604051620026a5380380620026a583398181016040528101906200005b91906200038c565b6200007b6200006f620001e960201b60201c565b620001f160201b60201c565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603620000ed576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000e4906200045a565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036200015f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200015690620004f2565b60405180910390fd5b81600960046101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505062000514565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620002e782620002ba565b9050919050565b6000620002fb82620002da565b9050919050565b6200030d81620002ee565b81146200031957600080fd5b50565b6000815190506200032d8162000302565b92915050565b60006200034082620002ba565b9050919050565b6000620003548262000333565b9050919050565b620003668162000347565b81146200037257600080fd5b50565b60008151905062000386816200035b565b92915050565b60008060408385031215620003a657620003a5620002b5565b5b6000620003b6858286016200031c565b9250506020620003c98582860162000375565b9150509250929050565b600082825260208201905092915050565b7f5a65726f2d616464726573732077617320676976656e20666f7220746865205260008201527f616e646f6d57616c6b20746f6b656e2e00000000000000000000000000000000602082015250565b600062000442603083620003d3565b91506200044f82620003e4565b604082019050919050565b60006020820190508181036000830152620004758162000433565b9050919050565b7f5a65726f2d616464726573732077617320676976656e20666f7220746865206760008201527f616d652e00000000000000000000000000000000000000000000000000000000602082015250565b6000620004da602483620003d3565b9150620004e7826200047c565b604082019050919050565b600060208201905081810360008301526200050d81620004cb565b9050919050565b61218180620005246000396000f3fe608060405234801561001057600080fd5b50600436106101735760003560e01c8063715018a6116100de578063c065894e11610097578063d84f2bc611610071578063d84f2bc61461047a578063f0a5242414610498578063f2fde38b146104c8578063fe939afc146104e457610173565b8063c065894e146103fc578063c07885551461042c578063c3fe3e281461045c57610173565b8063715018a614610336578063889d1e1a146103405780638da5cb5b14610370578063a2b136fb1461038e578063a531aa86146103c2578063a694fc3a146103e057610173565b8063451f1adf11610130578063451f1adf1461024a5780634f6ccce71461027c5780635111a2d6146102ac57806355279fdb146102ca5780635fda0acc146102e85780636427d9a91461030657610173565b80630d50c1891461017857806317db62131461019457806325646e1f146101b25780632e17de78146101ce57806341810425146101ea57806344d110b91461021a575b600080fd5b610192600480360381019061018d91906115e2565b610500565b005b61019c610546565b6040516101a9919061163a565b60405180910390f35b6101cc60048036038101906101c79190611691565b61054c565b005b6101e860048036038101906101e391906116be565b610623565b005b61020460048036038101906101ff9190611721565b610914565b604051610211919061178f565b60405180910390f35b610234600480360381019061022f91906116be565b6109ea565b604051610241919061163a565b60405180910390f35b610264600480360381019061025f91906116be565b610a02565b604051610273939291906117b9565b60405180910390f35b610296600480360381019061029191906116be565b610a3c565b6040516102a3919061163a565b60405180910390f35b6102b4610a9b565b6040516102c1919061184f565b60405180910390f35b6102d2610ac1565b6040516102df919061163a565b60405180910390f35b6102f0610ac7565b6040516102fd919061163a565b60405180910390f35b610320600480360381019061031b91906116be565b610ad4565b60405161032d9190611883565b60405180910390f35b61033e610aec565b005b61035a600480360381019061035591906116be565b610b74565b6040516103679190611883565b60405180910390f35b610378610be1565b604051610385919061178f565b60405180910390f35b6103a860048036038101906103a391906116be565b610c0a565b6040516103b995949392919061189e565b60405180910390f35b6103ca610c90565b6040516103d7919061163a565b60405180910390f35b6103fa60048036038101906103f591906116be565b610c96565b005b610416600480360381019061041191906116be565b610f4d565b604051610423919061178f565b60405180910390f35b610446600480360381019061044191906116be565b610fb4565b604051610453919061163a565b60405180910390f35b610464610fd8565b6040516104719190611924565b60405180910390f35b610482610ffe565b60405161048f919061193f565b60405180910390f35b6104b260048036038101906104ad91906116be565b611014565b6040516104bf9190611975565b60405180910390f35b6104e260048036038101906104dd91906119bc565b611034565b005b6104fe60048036038101906104f991906115e2565b61112b565b005b60005b81518110156105425761052f828281518110610522576105216119e9565b5b6020026020010151610623565b808061053a90611a47565b915050610503565b5050565b60085481565b610554611171565b73ffffffffffffffffffffffffffffffffffffffff16610572610be1565b73ffffffffffffffffffffffffffffffffffffffff16146105c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105bf90611aec565b60405180910390fd5b80600960006101000a81548163ffffffff021916908363ffffffff1602179055507f3bc8e083bbd3346984c352a42f59129e9299433a69e464ed246798ba091831a5816040516106189190611b3d565b60405180910390a150565b60006001600083815260200190815260200160002060010160189054906101000a900463ffffffff1663ffffffff1614610692576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161068990611ba4565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff166001600083815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610736576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161072d90611c10565b60405180910390fd5b4260016000838152602001908152602001600020600101601c9054906101000a900463ffffffff1663ffffffff16106107a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079b90611c7c565b60405180910390fd5b6000600160008381526020019081526020016000206000015490506107c881611179565b600960049054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd303360016000878152602001908152602001600020600001546040518463ffffffff1660e01b815260040161083d93929190611c9c565b600060405180830381600087803b15801561085757600080fd5b505af115801561086b573d6000803e3d6000fd5b50505050426001600084815260200190815260200160002060010160186101000a81548163ffffffff021916908363ffffffff1602179055506001600860008282546108b79190611cd3565b925050819055503373ffffffffffffffffffffffffffffffffffffffff1681837f33940a9fc3ee3c9150b29b34ef29ca397b5e1e340425a4f0da0346b5b90766c8600854604051610908919061163a565b60405180910390a45050565b6000806003805490501161095d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161095490611d79565b60405180910390fd5b6000600380805490508460001c6109749190611dc8565b81548110610985576109846119e9565b5b906000526020600020015490506000600560008381526020019081526020016000205490506001600082815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1692505050919050565b60046020528060005260406000206000915090505481565b60066020528060005260406000206000915090508060000160009054906101000a900463ffffffff16908060010154908060020154905083565b6000808211610a80576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a7790611e6b565b60405180910390fd5b60046000838152602001908152602001600020549050919050565b600960049054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60075481565b6000600380549050905090565b60056020528060005260406000206000915090505481565b610af4611171565b73ffffffffffffffffffffffffffffffffffffffff16610b12610be1565b73ffffffffffffffffffffffffffffffffffffffff1614610b68576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b5f90611aec565b60405180910390fd5b610b7260006112d0565b565b6000806004600084815260200190815260200160002054905060008103610bbe577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe915050610bdc565b60006005600085815260200190815260200160002054905080925050505b919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60016020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160149054906101000a900463ffffffff16908060010160189054906101000a900463ffffffff169080600101601c9054906101000a900463ffffffff16905085565b60025481565b600960049054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b8152600401610cf593929190611c9c565b600060405180830381600087803b158015610d0f57600080fd5b505af1158015610d23573d6000803e3d6000fd5b50505050610d3381600254611394565b80600160006002548152602001908152602001600020600001819055503360016000600254815260200190815260200160002060010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055504260016000600254815260200190815260200160002060010160146101000a81548163ffffffff021916908363ffffffff1602179055506000600960009054906101000a900463ffffffff1642610dfe9190611e8b565b9050428163ffffffff1611610e48576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e3f90611f35565b60405180910390fd5b80600160006002548152602001908152602001600020600101601c6101000a81548163ffffffff021916908363ffffffff160217905550600160026000828254610e929190611f55565b92505081905550600160086000828254610eac9190611f55565b925050819055503373ffffffffffffffffffffffffffffffffffffffff16826001600254610eda9190611cd3565b7f057eba8c4bba00f858e4f586f9c02794abb0df789ef316c741f9073fe2c435db600854600160006001600254610f119190611cd3565b8152602001908152602001600020600101601c9054906101000a900463ffffffff16604051610f41929190611f89565b60405180910390a45050565b600080610f5983610b74565b90506000811215610f6e576000915050610faf565b60006001600083815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905080925050505b919050565b60038181548110610fc457600080fd5b906000526020600020016000915090505481565b600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600960009054906101000a900463ffffffff1681565b600080600460008481526020019081526020016000205414159050919050565b61103c611171565b73ffffffffffffffffffffffffffffffffffffffff1661105a610be1565b73ffffffffffffffffffffffffffffffffffffffff16146110b0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110a790611aec565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361111f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161111690612024565b60405180910390fd5b611128816112d0565b50565b60005b815181101561116d5761115a82828151811061114d5761114c6119e9565b5b6020026020010151610c96565b808061116590611a47565b91505061112e565b5050565b600033905090565b61118281611014565b6111c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111b890612090565b60405180910390fd5b6000600460008381526020019081526020016000205490506000600360016003805490506111ef9190611cd3565b81548110611200576111ff6119e9565b5b9060005260206000200154905080600360018461121d9190611cd3565b8154811061122e5761122d6119e9565b5b90600052602060002001819055508160046000838152602001908152602001600020819055506004600084815260200190815260200160002060009055600380548061127d5761127c6120b0565b5b600190038181906000526020600020016000905590557fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6005600085815260200190815260200160002081905550505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b61139d82611014565b156113dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113d49061212b565b60405180910390fd5b600382908060018154018082558091505060019003906000526020600020016000909190919091505560038054905060046000848152602001908152602001600020819055508060056000848152602001908152602001600020819055505050565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6114a182611458565b810181811067ffffffffffffffff821117156114c0576114bf611469565b5b80604052505050565b60006114d361143f565b90506114df8282611498565b919050565b600067ffffffffffffffff8211156114ff576114fe611469565b5b602082029050602081019050919050565b600080fd5b6000819050919050565b61152881611515565b811461153357600080fd5b50565b6000813590506115458161151f565b92915050565b600061155e611559846114e4565b6114c9565b9050808382526020820190506020840283018581111561158157611580611510565b5b835b818110156115aa57806115968882611536565b845260208401935050602081019050611583565b5050509392505050565b600082601f8301126115c9576115c8611453565b5b81356115d984826020860161154b565b91505092915050565b6000602082840312156115f8576115f7611449565b5b600082013567ffffffffffffffff8111156116165761161561144e565b5b611622848285016115b4565b91505092915050565b61163481611515565b82525050565b600060208201905061164f600083018461162b565b92915050565b600063ffffffff82169050919050565b61166e81611655565b811461167957600080fd5b50565b60008135905061168b81611665565b92915050565b6000602082840312156116a7576116a6611449565b5b60006116b58482850161167c565b91505092915050565b6000602082840312156116d4576116d3611449565b5b60006116e284828501611536565b91505092915050565b6000819050919050565b6116fe816116eb565b811461170957600080fd5b50565b60008135905061171b816116f5565b92915050565b60006020828403121561173757611736611449565b5b60006117458482850161170c565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006117798261174e565b9050919050565b6117898161176e565b82525050565b60006020820190506117a46000830184611780565b92915050565b6117b381611655565b82525050565b60006060820190506117ce60008301866117aa565b6117db602083018561162b565b6117e8604083018461162b565b949350505050565b6000819050919050565b600061181561181061180b8461174e565b6117f0565b61174e565b9050919050565b6000611827826117fa565b9050919050565b60006118398261181c565b9050919050565b6118498161182e565b82525050565b60006020820190506118646000830184611840565b92915050565b6000819050919050565b61187d8161186a565b82525050565b60006020820190506118986000830184611874565b92915050565b600060a0820190506118b3600083018861162b565b6118c06020830187611780565b6118cd60408301866117aa565b6118da60608301856117aa565b6118e760808301846117aa565b9695505050505050565b60006118fc826117fa565b9050919050565b600061190e826118f1565b9050919050565b61191e81611903565b82525050565b60006020820190506119396000830184611915565b92915050565b600060208201905061195460008301846117aa565b92915050565b60008115159050919050565b61196f8161195a565b82525050565b600060208201905061198a6000830184611966565b92915050565b6119998161176e565b81146119a457600080fd5b50565b6000813590506119b681611990565b92915050565b6000602082840312156119d2576119d1611449565b5b60006119e0848285016119a7565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611a5282611515565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611a8457611a83611a18565b5b600182019050919050565b600082825260208201905092915050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000611ad6602083611a8f565b9150611ae182611aa0565b602082019050919050565b60006020820190508181036000830152611b0581611ac9565b9050919050565b6000611b27611b22611b1d84611655565b6117f0565b611515565b9050919050565b611b3781611b0c565b82525050565b6000602082019050611b526000830184611b2e565b92915050565b7f546f6b656e2068617320616c7265616479206265656e20756e7374616b65642e600082015250565b6000611b8e602083611a8f565b9150611b9982611b58565b602082019050919050565b60006020820190508181036000830152611bbd81611b81565b9050919050565b7f4f6e6c7920746865206f776e65722063616e20756e7374616b652e0000000000600082015250565b6000611bfa601b83611a8f565b9150611c0582611bc4565b602082019050919050565b60006020820190508181036000830152611c2981611bed565b9050919050565b7f4e6f7420616c6c6f77656420746f20756e7374616b65207965742e0000000000600082015250565b6000611c66601b83611a8f565b9150611c7182611c30565b602082019050919050565b60006020820190508181036000830152611c9581611c59565b9050919050565b6000606082019050611cb16000830186611780565b611cbe6020830185611780565b611ccb604083018461162b565b949350505050565b6000611cde82611515565b9150611ce983611515565b9250828203905081811115611d0157611d00611a18565b5b92915050565b7f546865726520617265206e6f2052616e646f6d57616c6b20746f6b656e73207360008201527f74616b65642e0000000000000000000000000000000000000000000000000000602082015250565b6000611d63602683611a8f565b9150611d6e82611d07565b604082019050919050565b60006020820190508181036000830152611d9281611d56565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611dd382611515565b9150611dde83611515565b925082611dee57611ded611d99565b5b828206905092915050565b7f5a65726f2077617320676976656e2c20746f6b656e20696e646963657320737460008201527f6172742066726f6d203100000000000000000000000000000000000000000000602082015250565b6000611e55602a83611a8f565b9150611e6082611df9565b604082019050919050565b60006020820190508181036000830152611e8481611e48565b9050919050565b6000611e9682611655565b9150611ea183611655565b9250828201905063ffffffff811115611ebd57611ebc611a18565b5b92915050565b7f556e7374616b652074696d652073686f756c642062652062696767657220746860008201527f616e20626c6f636b2074696d657374616d700000000000000000000000000000602082015250565b6000611f1f603283611a8f565b9150611f2a82611ec3565b604082019050919050565b60006020820190508181036000830152611f4e81611f12565b9050919050565b6000611f6082611515565b9150611f6b83611515565b9250828201905080821115611f8357611f82611a18565b5b92915050565b6000604082019050611f9e600083018561162b565b611fab6020830184611b2e565b9392505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b600061200e602683611a8f565b915061201982611fb2565b604082019050919050565b6000602082019050818103600083015261203d81612001565b9050919050565b7f546f6b656e206973206e6f7420696e20746865206c6973742e00000000000000600082015250565b600061207a601983611a8f565b915061208582612044565b602082019050919050565b600060208201905081810360008301526120a98161206d565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b7f546f6b656e20616c726561647920696e20746865206c6973742e000000000000600082015250565b6000612115601a83611a8f565b9150612120826120df565b602082019050919050565b6000602082019050818103600083015261214481612108565b905091905056fea264697066735822122040cdbbc045ecf16b7e738e3c599214b87174a0f24e2f28529ca8e46e98b303f164736f6c63430008130033",
}

// StakingWalletRWalkABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingWalletRWalkMetaData.ABI instead.
var StakingWalletRWalkABI = StakingWalletRWalkMetaData.ABI

// StakingWalletRWalkBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingWalletRWalkMetaData.Bin instead.
var StakingWalletRWalkBin = StakingWalletRWalkMetaData.Bin

// DeployStakingWalletRWalk deploys a new Ethereum contract, binding an instance of StakingWalletRWalk to it.
func DeployStakingWalletRWalk(auth *bind.TransactOpts, backend bind.ContractBackend, rwalk_ common.Address, game_ common.Address) (common.Address, *types.Transaction, *StakingWalletRWalk, error) {
	parsed, err := StakingWalletRWalkMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingWalletRWalkBin), backend, rwalk_, game_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StakingWalletRWalk{StakingWalletRWalkCaller: StakingWalletRWalkCaller{contract: contract}, StakingWalletRWalkTransactor: StakingWalletRWalkTransactor{contract: contract}, StakingWalletRWalkFilterer: StakingWalletRWalkFilterer{contract: contract}}, nil
}

// StakingWalletRWalk is an auto generated Go binding around an Ethereum contract.
type StakingWalletRWalk struct {
	StakingWalletRWalkCaller     // Read-only binding to the contract
	StakingWalletRWalkTransactor // Write-only binding to the contract
	StakingWalletRWalkFilterer   // Log filterer for contract events
}

// StakingWalletRWalkCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingWalletRWalkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletRWalkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingWalletRWalkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletRWalkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingWalletRWalkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletRWalkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingWalletRWalkSession struct {
	Contract     *StakingWalletRWalk // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StakingWalletRWalkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingWalletRWalkCallerSession struct {
	Contract *StakingWalletRWalkCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// StakingWalletRWalkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingWalletRWalkTransactorSession struct {
	Contract     *StakingWalletRWalkTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// StakingWalletRWalkRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingWalletRWalkRaw struct {
	Contract *StakingWalletRWalk // Generic contract binding to access the raw methods on
}

// StakingWalletRWalkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingWalletRWalkCallerRaw struct {
	Contract *StakingWalletRWalkCaller // Generic read-only contract binding to access the raw methods on
}

// StakingWalletRWalkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingWalletRWalkTransactorRaw struct {
	Contract *StakingWalletRWalkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingWalletRWalk creates a new instance of StakingWalletRWalk, bound to a specific deployed contract.
func NewStakingWalletRWalk(address common.Address, backend bind.ContractBackend) (*StakingWalletRWalk, error) {
	contract, err := bindStakingWalletRWalk(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRWalk{StakingWalletRWalkCaller: StakingWalletRWalkCaller{contract: contract}, StakingWalletRWalkTransactor: StakingWalletRWalkTransactor{contract: contract}, StakingWalletRWalkFilterer: StakingWalletRWalkFilterer{contract: contract}}, nil
}

// NewStakingWalletRWalkCaller creates a new read-only instance of StakingWalletRWalk, bound to a specific deployed contract.
func NewStakingWalletRWalkCaller(address common.Address, caller bind.ContractCaller) (*StakingWalletRWalkCaller, error) {
	contract, err := bindStakingWalletRWalk(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRWalkCaller{contract: contract}, nil
}

// NewStakingWalletRWalkTransactor creates a new write-only instance of StakingWalletRWalk, bound to a specific deployed contract.
func NewStakingWalletRWalkTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingWalletRWalkTransactor, error) {
	contract, err := bindStakingWalletRWalk(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRWalkTransactor{contract: contract}, nil
}

// NewStakingWalletRWalkFilterer creates a new log filterer instance of StakingWalletRWalk, bound to a specific deployed contract.
func NewStakingWalletRWalkFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingWalletRWalkFilterer, error) {
	contract, err := bindStakingWalletRWalk(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRWalkFilterer{contract: contract}, nil
}

// bindStakingWalletRWalk binds a generic wrapper to an already deployed contract.
func bindStakingWalletRWalk(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingWalletRWalkMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingWalletRWalk *StakingWalletRWalkRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingWalletRWalk.Contract.StakingWalletRWalkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingWalletRWalk *StakingWalletRWalkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.StakingWalletRWalkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingWalletRWalk *StakingWalletRWalkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.StakingWalletRWalkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingWalletRWalk *StakingWalletRWalkCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingWalletRWalk.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingWalletRWalk *StakingWalletRWalkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingWalletRWalk *StakingWalletRWalkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.contract.Transact(opts, method, params...)
}

// ETHDeposits is a free data retrieval call binding the contract method 0x451f1adf.
//
// Solidity: function ETHDeposits(uint256 ) view returns(uint32 depositTime, uint256 depositAmount, uint256 numStaked)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) ETHDeposits(opts *bind.CallOpts, arg0 *big.Int) (struct {
	DepositTime   uint32
	DepositAmount *big.Int
	NumStaked     *big.Int
}, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "ETHDeposits", arg0)

	outstruct := new(struct {
		DepositTime   uint32
		DepositAmount *big.Int
		NumStaked     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DepositTime = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.DepositAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NumStaked = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ETHDeposits is a free data retrieval call binding the contract method 0x451f1adf.
//
// Solidity: function ETHDeposits(uint256 ) view returns(uint32 depositTime, uint256 depositAmount, uint256 numStaked)
func (_StakingWalletRWalk *StakingWalletRWalkSession) ETHDeposits(arg0 *big.Int) (struct {
	DepositTime   uint32
	DepositAmount *big.Int
	NumStaked     *big.Int
}, error) {
	return _StakingWalletRWalk.Contract.ETHDeposits(&_StakingWalletRWalk.CallOpts, arg0)
}

// ETHDeposits is a free data retrieval call binding the contract method 0x451f1adf.
//
// Solidity: function ETHDeposits(uint256 ) view returns(uint32 depositTime, uint256 depositAmount, uint256 numStaked)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) ETHDeposits(arg0 *big.Int) (struct {
	DepositTime   uint32
	DepositAmount *big.Int
	NumStaked     *big.Int
}, error) {
	return _StakingWalletRWalk.Contract.ETHDeposits(&_StakingWalletRWalk.CallOpts, arg0)
}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) Game(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "game")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_StakingWalletRWalk *StakingWalletRWalkSession) Game() (common.Address, error) {
	return _StakingWalletRWalk.Contract.Game(&_StakingWalletRWalk.CallOpts)
}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) Game() (common.Address, error) {
	return _StakingWalletRWalk.Contract.Game(&_StakingWalletRWalk.CallOpts)
}

// IsTokenStaked is a free data retrieval call binding the contract method 0xf0a52424.
//
// Solidity: function isTokenStaked(uint256 tokenId) view returns(bool)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) IsTokenStaked(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "isTokenStaked", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenStaked is a free data retrieval call binding the contract method 0xf0a52424.
//
// Solidity: function isTokenStaked(uint256 tokenId) view returns(bool)
func (_StakingWalletRWalk *StakingWalletRWalkSession) IsTokenStaked(tokenId *big.Int) (bool, error) {
	return _StakingWalletRWalk.Contract.IsTokenStaked(&_StakingWalletRWalk.CallOpts, tokenId)
}

// IsTokenStaked is a free data retrieval call binding the contract method 0xf0a52424.
//
// Solidity: function isTokenStaked(uint256 tokenId) view returns(bool)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) IsTokenStaked(tokenId *big.Int) (bool, error) {
	return _StakingWalletRWalk.Contract.IsTokenStaked(&_StakingWalletRWalk.CallOpts, tokenId)
}

// LastActionIdByTokenId is a free data retrieval call binding the contract method 0x889d1e1a.
//
// Solidity: function lastActionIdByTokenId(uint256 tokenId) view returns(int256)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) LastActionIdByTokenId(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "lastActionIdByTokenId", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastActionIdByTokenId is a free data retrieval call binding the contract method 0x889d1e1a.
//
// Solidity: function lastActionIdByTokenId(uint256 tokenId) view returns(int256)
func (_StakingWalletRWalk *StakingWalletRWalkSession) LastActionIdByTokenId(tokenId *big.Int) (*big.Int, error) {
	return _StakingWalletRWalk.Contract.LastActionIdByTokenId(&_StakingWalletRWalk.CallOpts, tokenId)
}

// LastActionIdByTokenId is a free data retrieval call binding the contract method 0x889d1e1a.
//
// Solidity: function lastActionIdByTokenId(uint256 tokenId) view returns(int256)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) LastActionIdByTokenId(tokenId *big.Int) (*big.Int, error) {
	return _StakingWalletRWalk.Contract.LastActionIdByTokenId(&_StakingWalletRWalk.CallOpts, tokenId)
}

// LastActionIds is a free data retrieval call binding the contract method 0x6427d9a9.
//
// Solidity: function lastActionIds(uint256 ) view returns(int256)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) LastActionIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "lastActionIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastActionIds is a free data retrieval call binding the contract method 0x6427d9a9.
//
// Solidity: function lastActionIds(uint256 ) view returns(int256)
func (_StakingWalletRWalk *StakingWalletRWalkSession) LastActionIds(arg0 *big.Int) (*big.Int, error) {
	return _StakingWalletRWalk.Contract.LastActionIds(&_StakingWalletRWalk.CallOpts, arg0)
}

// LastActionIds is a free data retrieval call binding the contract method 0x6427d9a9.
//
// Solidity: function lastActionIds(uint256 ) view returns(int256)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) LastActionIds(arg0 *big.Int) (*big.Int, error) {
	return _StakingWalletRWalk.Contract.LastActionIds(&_StakingWalletRWalk.CallOpts, arg0)
}

// MinStakePeriod is a free data retrieval call binding the contract method 0xd84f2bc6.
//
// Solidity: function minStakePeriod() view returns(uint32)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) MinStakePeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "minStakePeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MinStakePeriod is a free data retrieval call binding the contract method 0xd84f2bc6.
//
// Solidity: function minStakePeriod() view returns(uint32)
func (_StakingWalletRWalk *StakingWalletRWalkSession) MinStakePeriod() (uint32, error) {
	return _StakingWalletRWalk.Contract.MinStakePeriod(&_StakingWalletRWalk.CallOpts)
}

// MinStakePeriod is a free data retrieval call binding the contract method 0xd84f2bc6.
//
// Solidity: function minStakePeriod() view returns(uint32)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) MinStakePeriod() (uint32, error) {
	return _StakingWalletRWalk.Contract.MinStakePeriod(&_StakingWalletRWalk.CallOpts)
}

// NumETHDeposits is a free data retrieval call binding the contract method 0x55279fdb.
//
// Solidity: function numETHDeposits() view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) NumETHDeposits(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "numETHDeposits")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumETHDeposits is a free data retrieval call binding the contract method 0x55279fdb.
//
// Solidity: function numETHDeposits() view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkSession) NumETHDeposits() (*big.Int, error) {
	return _StakingWalletRWalk.Contract.NumETHDeposits(&_StakingWalletRWalk.CallOpts)
}

// NumETHDeposits is a free data retrieval call binding the contract method 0x55279fdb.
//
// Solidity: function numETHDeposits() view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) NumETHDeposits() (*big.Int, error) {
	return _StakingWalletRWalk.Contract.NumETHDeposits(&_StakingWalletRWalk.CallOpts)
}

// NumStakeActions is a free data retrieval call binding the contract method 0xa531aa86.
//
// Solidity: function numStakeActions() view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) NumStakeActions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "numStakeActions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumStakeActions is a free data retrieval call binding the contract method 0xa531aa86.
//
// Solidity: function numStakeActions() view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkSession) NumStakeActions() (*big.Int, error) {
	return _StakingWalletRWalk.Contract.NumStakeActions(&_StakingWalletRWalk.CallOpts)
}

// NumStakeActions is a free data retrieval call binding the contract method 0xa531aa86.
//
// Solidity: function numStakeActions() view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) NumStakeActions() (*big.Int, error) {
	return _StakingWalletRWalk.Contract.NumStakeActions(&_StakingWalletRWalk.CallOpts)
}

// NumStakedNFTs is a free data retrieval call binding the contract method 0x17db6213.
//
// Solidity: function numStakedNFTs() view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) NumStakedNFTs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "numStakedNFTs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumStakedNFTs is a free data retrieval call binding the contract method 0x17db6213.
//
// Solidity: function numStakedNFTs() view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkSession) NumStakedNFTs() (*big.Int, error) {
	return _StakingWalletRWalk.Contract.NumStakedNFTs(&_StakingWalletRWalk.CallOpts)
}

// NumStakedNFTs is a free data retrieval call binding the contract method 0x17db6213.
//
// Solidity: function numStakedNFTs() view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) NumStakedNFTs() (*big.Int, error) {
	return _StakingWalletRWalk.Contract.NumStakedNFTs(&_StakingWalletRWalk.CallOpts)
}

// NumTokensStaked is a free data retrieval call binding the contract method 0x5fda0acc.
//
// Solidity: function numTokensStaked() view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) NumTokensStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "numTokensStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumTokensStaked is a free data retrieval call binding the contract method 0x5fda0acc.
//
// Solidity: function numTokensStaked() view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkSession) NumTokensStaked() (*big.Int, error) {
	return _StakingWalletRWalk.Contract.NumTokensStaked(&_StakingWalletRWalk.CallOpts)
}

// NumTokensStaked is a free data retrieval call binding the contract method 0x5fda0acc.
//
// Solidity: function numTokensStaked() view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) NumTokensStaked() (*big.Int, error) {
	return _StakingWalletRWalk.Contract.NumTokensStaked(&_StakingWalletRWalk.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingWalletRWalk *StakingWalletRWalkSession) Owner() (common.Address, error) {
	return _StakingWalletRWalk.Contract.Owner(&_StakingWalletRWalk.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) Owner() (common.Address, error) {
	return _StakingWalletRWalk.Contract.Owner(&_StakingWalletRWalk.CallOpts)
}

// PickRandomStaker is a free data retrieval call binding the contract method 0x41810425.
//
// Solidity: function pickRandomStaker(bytes32 entropy) view returns(address)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) PickRandomStaker(opts *bind.CallOpts, entropy [32]byte) (common.Address, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "pickRandomStaker", entropy)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PickRandomStaker is a free data retrieval call binding the contract method 0x41810425.
//
// Solidity: function pickRandomStaker(bytes32 entropy) view returns(address)
func (_StakingWalletRWalk *StakingWalletRWalkSession) PickRandomStaker(entropy [32]byte) (common.Address, error) {
	return _StakingWalletRWalk.Contract.PickRandomStaker(&_StakingWalletRWalk.CallOpts, entropy)
}

// PickRandomStaker is a free data retrieval call binding the contract method 0x41810425.
//
// Solidity: function pickRandomStaker(bytes32 entropy) view returns(address)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) PickRandomStaker(entropy [32]byte) (common.Address, error) {
	return _StakingWalletRWalk.Contract.PickRandomStaker(&_StakingWalletRWalk.CallOpts, entropy)
}

// RandomWalk is a free data retrieval call binding the contract method 0x5111a2d6.
//
// Solidity: function randomWalk() view returns(address)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) RandomWalk(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "randomWalk")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RandomWalk is a free data retrieval call binding the contract method 0x5111a2d6.
//
// Solidity: function randomWalk() view returns(address)
func (_StakingWalletRWalk *StakingWalletRWalkSession) RandomWalk() (common.Address, error) {
	return _StakingWalletRWalk.Contract.RandomWalk(&_StakingWalletRWalk.CallOpts)
}

// RandomWalk is a free data retrieval call binding the contract method 0x5111a2d6.
//
// Solidity: function randomWalk() view returns(address)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) RandomWalk() (common.Address, error) {
	return _StakingWalletRWalk.Contract.RandomWalk(&_StakingWalletRWalk.CallOpts)
}

// StakeActions is a free data retrieval call binding the contract method 0xa2b136fb.
//
// Solidity: function stakeActions(uint256 ) view returns(uint256 tokenId, address owner, uint32 stakeTime, uint32 unstakeTime, uint32 unstakeEligibleTime)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) StakeActions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TokenId             *big.Int
	Owner               common.Address
	StakeTime           uint32
	UnstakeTime         uint32
	UnstakeEligibleTime uint32
}, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "stakeActions", arg0)

	outstruct := new(struct {
		TokenId             *big.Int
		Owner               common.Address
		StakeTime           uint32
		UnstakeTime         uint32
		UnstakeEligibleTime uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.StakeTime = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.UnstakeTime = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.UnstakeEligibleTime = *abi.ConvertType(out[4], new(uint32)).(*uint32)

	return *outstruct, err

}

// StakeActions is a free data retrieval call binding the contract method 0xa2b136fb.
//
// Solidity: function stakeActions(uint256 ) view returns(uint256 tokenId, address owner, uint32 stakeTime, uint32 unstakeTime, uint32 unstakeEligibleTime)
func (_StakingWalletRWalk *StakingWalletRWalkSession) StakeActions(arg0 *big.Int) (struct {
	TokenId             *big.Int
	Owner               common.Address
	StakeTime           uint32
	UnstakeTime         uint32
	UnstakeEligibleTime uint32
}, error) {
	return _StakingWalletRWalk.Contract.StakeActions(&_StakingWalletRWalk.CallOpts, arg0)
}

// StakeActions is a free data retrieval call binding the contract method 0xa2b136fb.
//
// Solidity: function stakeActions(uint256 ) view returns(uint256 tokenId, address owner, uint32 stakeTime, uint32 unstakeTime, uint32 unstakeEligibleTime)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) StakeActions(arg0 *big.Int) (struct {
	TokenId             *big.Int
	Owner               common.Address
	StakeTime           uint32
	UnstakeTime         uint32
	UnstakeEligibleTime uint32
}, error) {
	return _StakingWalletRWalk.Contract.StakeActions(&_StakingWalletRWalk.CallOpts, arg0)
}

// StakedTokens is a free data retrieval call binding the contract method 0xc0788555.
//
// Solidity: function stakedTokens(uint256 ) view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) StakedTokens(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "stakedTokens", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedTokens is a free data retrieval call binding the contract method 0xc0788555.
//
// Solidity: function stakedTokens(uint256 ) view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkSession) StakedTokens(arg0 *big.Int) (*big.Int, error) {
	return _StakingWalletRWalk.Contract.StakedTokens(&_StakingWalletRWalk.CallOpts, arg0)
}

// StakedTokens is a free data retrieval call binding the contract method 0xc0788555.
//
// Solidity: function stakedTokens(uint256 ) view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) StakedTokens(arg0 *big.Int) (*big.Int, error) {
	return _StakingWalletRWalk.Contract.StakedTokens(&_StakingWalletRWalk.CallOpts, arg0)
}

// StakerByTokenId is a free data retrieval call binding the contract method 0xc065894e.
//
// Solidity: function stakerByTokenId(uint256 tokenId) view returns(address)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) StakerByTokenId(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "stakerByTokenId", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakerByTokenId is a free data retrieval call binding the contract method 0xc065894e.
//
// Solidity: function stakerByTokenId(uint256 tokenId) view returns(address)
func (_StakingWalletRWalk *StakingWalletRWalkSession) StakerByTokenId(tokenId *big.Int) (common.Address, error) {
	return _StakingWalletRWalk.Contract.StakerByTokenId(&_StakingWalletRWalk.CallOpts, tokenId)
}

// StakerByTokenId is a free data retrieval call binding the contract method 0xc065894e.
//
// Solidity: function stakerByTokenId(uint256 tokenId) view returns(address)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) StakerByTokenId(tokenId *big.Int) (common.Address, error) {
	return _StakingWalletRWalk.Contract.StakerByTokenId(&_StakingWalletRWalk.CallOpts, tokenId)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 tokenIndex) view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) TokenByIndex(opts *bind.CallOpts, tokenIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "tokenByIndex", tokenIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 tokenIndex) view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkSession) TokenByIndex(tokenIndex *big.Int) (*big.Int, error) {
	return _StakingWalletRWalk.Contract.TokenByIndex(&_StakingWalletRWalk.CallOpts, tokenIndex)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 tokenIndex) view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) TokenByIndex(tokenIndex *big.Int) (*big.Int, error) {
	return _StakingWalletRWalk.Contract.TokenByIndex(&_StakingWalletRWalk.CallOpts, tokenIndex)
}

// TokenIndices is a free data retrieval call binding the contract method 0x44d110b9.
//
// Solidity: function tokenIndices(uint256 ) view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) TokenIndices(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "tokenIndices", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIndices is a free data retrieval call binding the contract method 0x44d110b9.
//
// Solidity: function tokenIndices(uint256 ) view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkSession) TokenIndices(arg0 *big.Int) (*big.Int, error) {
	return _StakingWalletRWalk.Contract.TokenIndices(&_StakingWalletRWalk.CallOpts, arg0)
}

// TokenIndices is a free data retrieval call binding the contract method 0x44d110b9.
//
// Solidity: function tokenIndices(uint256 ) view returns(uint256)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) TokenIndices(arg0 *big.Int) (*big.Int, error) {
	return _StakingWalletRWalk.Contract.TokenIndices(&_StakingWalletRWalk.CallOpts, arg0)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingWalletRWalk *StakingWalletRWalkTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingWalletRWalk.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingWalletRWalk *StakingWalletRWalkSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.RenounceOwnership(&_StakingWalletRWalk.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingWalletRWalk *StakingWalletRWalkTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.RenounceOwnership(&_StakingWalletRWalk.TransactOpts)
}

// SetMinStakePeriod is a paid mutator transaction binding the contract method 0x25646e1f.
//
// Solidity: function setMinStakePeriod(uint32 newStakePeriod) returns()
func (_StakingWalletRWalk *StakingWalletRWalkTransactor) SetMinStakePeriod(opts *bind.TransactOpts, newStakePeriod uint32) (*types.Transaction, error) {
	return _StakingWalletRWalk.contract.Transact(opts, "setMinStakePeriod", newStakePeriod)
}

// SetMinStakePeriod is a paid mutator transaction binding the contract method 0x25646e1f.
//
// Solidity: function setMinStakePeriod(uint32 newStakePeriod) returns()
func (_StakingWalletRWalk *StakingWalletRWalkSession) SetMinStakePeriod(newStakePeriod uint32) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.SetMinStakePeriod(&_StakingWalletRWalk.TransactOpts, newStakePeriod)
}

// SetMinStakePeriod is a paid mutator transaction binding the contract method 0x25646e1f.
//
// Solidity: function setMinStakePeriod(uint32 newStakePeriod) returns()
func (_StakingWalletRWalk *StakingWalletRWalkTransactorSession) SetMinStakePeriod(newStakePeriod uint32) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.SetMinStakePeriod(&_StakingWalletRWalk.TransactOpts, newStakePeriod)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _tokenId) returns()
func (_StakingWalletRWalk *StakingWalletRWalkTransactor) Stake(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _StakingWalletRWalk.contract.Transact(opts, "stake", _tokenId)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _tokenId) returns()
func (_StakingWalletRWalk *StakingWalletRWalkSession) Stake(_tokenId *big.Int) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.Stake(&_StakingWalletRWalk.TransactOpts, _tokenId)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _tokenId) returns()
func (_StakingWalletRWalk *StakingWalletRWalkTransactorSession) Stake(_tokenId *big.Int) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.Stake(&_StakingWalletRWalk.TransactOpts, _tokenId)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] ids) returns()
func (_StakingWalletRWalk *StakingWalletRWalkTransactor) StakeMany(opts *bind.TransactOpts, ids []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRWalk.contract.Transact(opts, "stakeMany", ids)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] ids) returns()
func (_StakingWalletRWalk *StakingWalletRWalkSession) StakeMany(ids []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.StakeMany(&_StakingWalletRWalk.TransactOpts, ids)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] ids) returns()
func (_StakingWalletRWalk *StakingWalletRWalkTransactorSession) StakeMany(ids []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.StakeMany(&_StakingWalletRWalk.TransactOpts, ids)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingWalletRWalk *StakingWalletRWalkTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakingWalletRWalk.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingWalletRWalk *StakingWalletRWalkSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.TransferOwnership(&_StakingWalletRWalk.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingWalletRWalk *StakingWalletRWalkTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.TransferOwnership(&_StakingWalletRWalk.TransactOpts, newOwner)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId) returns()
func (_StakingWalletRWalk *StakingWalletRWalkTransactor) Unstake(opts *bind.TransactOpts, stakeActionId *big.Int) (*types.Transaction, error) {
	return _StakingWalletRWalk.contract.Transact(opts, "unstake", stakeActionId)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId) returns()
func (_StakingWalletRWalk *StakingWalletRWalkSession) Unstake(stakeActionId *big.Int) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.Unstake(&_StakingWalletRWalk.TransactOpts, stakeActionId)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId) returns()
func (_StakingWalletRWalk *StakingWalletRWalkTransactorSession) Unstake(stakeActionId *big.Int) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.Unstake(&_StakingWalletRWalk.TransactOpts, stakeActionId)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] ids) returns()
func (_StakingWalletRWalk *StakingWalletRWalkTransactor) UnstakeMany(opts *bind.TransactOpts, ids []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRWalk.contract.Transact(opts, "unstakeMany", ids)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] ids) returns()
func (_StakingWalletRWalk *StakingWalletRWalkSession) UnstakeMany(ids []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.UnstakeMany(&_StakingWalletRWalk.TransactOpts, ids)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] ids) returns()
func (_StakingWalletRWalk *StakingWalletRWalkTransactorSession) UnstakeMany(ids []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRWalk.Contract.UnstakeMany(&_StakingWalletRWalk.TransactOpts, ids)
}

// StakingWalletRWalkMinStakePeriodChangedIterator is returned from FilterMinStakePeriodChanged and is used to iterate over the raw logs and unpacked data for MinStakePeriodChanged events raised by the StakingWalletRWalk contract.
type StakingWalletRWalkMinStakePeriodChangedIterator struct {
	Event *StakingWalletRWalkMinStakePeriodChanged // Event containing the contract specifics and raw log

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
func (it *StakingWalletRWalkMinStakePeriodChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletRWalkMinStakePeriodChanged)
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
		it.Event = new(StakingWalletRWalkMinStakePeriodChanged)
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
func (it *StakingWalletRWalkMinStakePeriodChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletRWalkMinStakePeriodChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletRWalkMinStakePeriodChanged represents a MinStakePeriodChanged event raised by the StakingWalletRWalk contract.
type StakingWalletRWalkMinStakePeriodChanged struct {
	NewPeriod *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinStakePeriodChanged is a free log retrieval operation binding the contract event 0x3bc8e083bbd3346984c352a42f59129e9299433a69e464ed246798ba091831a5.
//
// Solidity: event MinStakePeriodChanged(uint256 newPeriod)
func (_StakingWalletRWalk *StakingWalletRWalkFilterer) FilterMinStakePeriodChanged(opts *bind.FilterOpts) (*StakingWalletRWalkMinStakePeriodChangedIterator, error) {

	logs, sub, err := _StakingWalletRWalk.contract.FilterLogs(opts, "MinStakePeriodChanged")
	if err != nil {
		return nil, err
	}
	return &StakingWalletRWalkMinStakePeriodChangedIterator{contract: _StakingWalletRWalk.contract, event: "MinStakePeriodChanged", logs: logs, sub: sub}, nil
}

// WatchMinStakePeriodChanged is a free log subscription operation binding the contract event 0x3bc8e083bbd3346984c352a42f59129e9299433a69e464ed246798ba091831a5.
//
// Solidity: event MinStakePeriodChanged(uint256 newPeriod)
func (_StakingWalletRWalk *StakingWalletRWalkFilterer) WatchMinStakePeriodChanged(opts *bind.WatchOpts, sink chan<- *StakingWalletRWalkMinStakePeriodChanged) (event.Subscription, error) {

	logs, sub, err := _StakingWalletRWalk.contract.WatchLogs(opts, "MinStakePeriodChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletRWalkMinStakePeriodChanged)
				if err := _StakingWalletRWalk.contract.UnpackLog(event, "MinStakePeriodChanged", log); err != nil {
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

// ParseMinStakePeriodChanged is a log parse operation binding the contract event 0x3bc8e083bbd3346984c352a42f59129e9299433a69e464ed246798ba091831a5.
//
// Solidity: event MinStakePeriodChanged(uint256 newPeriod)
func (_StakingWalletRWalk *StakingWalletRWalkFilterer) ParseMinStakePeriodChanged(log types.Log) (*StakingWalletRWalkMinStakePeriodChanged, error) {
	event := new(StakingWalletRWalkMinStakePeriodChanged)
	if err := _StakingWalletRWalk.contract.UnpackLog(event, "MinStakePeriodChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletRWalkModuloSentEventIterator is returned from FilterModuloSentEvent and is used to iterate over the raw logs and unpacked data for ModuloSentEvent events raised by the StakingWalletRWalk contract.
type StakingWalletRWalkModuloSentEventIterator struct {
	Event *StakingWalletRWalkModuloSentEvent // Event containing the contract specifics and raw log

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
func (it *StakingWalletRWalkModuloSentEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletRWalkModuloSentEvent)
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
		it.Event = new(StakingWalletRWalkModuloSentEvent)
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
func (it *StakingWalletRWalkModuloSentEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletRWalkModuloSentEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletRWalkModuloSentEvent represents a ModuloSentEvent event raised by the StakingWalletRWalk contract.
type StakingWalletRWalkModuloSentEvent struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterModuloSentEvent is a free log retrieval operation binding the contract event 0x6905286e1ecb9b47a50224e76e31cf1383f75212bc0b06c8684317782566a0a7.
//
// Solidity: event ModuloSentEvent(uint256 amount)
func (_StakingWalletRWalk *StakingWalletRWalkFilterer) FilterModuloSentEvent(opts *bind.FilterOpts) (*StakingWalletRWalkModuloSentEventIterator, error) {

	logs, sub, err := _StakingWalletRWalk.contract.FilterLogs(opts, "ModuloSentEvent")
	if err != nil {
		return nil, err
	}
	return &StakingWalletRWalkModuloSentEventIterator{contract: _StakingWalletRWalk.contract, event: "ModuloSentEvent", logs: logs, sub: sub}, nil
}

// WatchModuloSentEvent is a free log subscription operation binding the contract event 0x6905286e1ecb9b47a50224e76e31cf1383f75212bc0b06c8684317782566a0a7.
//
// Solidity: event ModuloSentEvent(uint256 amount)
func (_StakingWalletRWalk *StakingWalletRWalkFilterer) WatchModuloSentEvent(opts *bind.WatchOpts, sink chan<- *StakingWalletRWalkModuloSentEvent) (event.Subscription, error) {

	logs, sub, err := _StakingWalletRWalk.contract.WatchLogs(opts, "ModuloSentEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletRWalkModuloSentEvent)
				if err := _StakingWalletRWalk.contract.UnpackLog(event, "ModuloSentEvent", log); err != nil {
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

// ParseModuloSentEvent is a log parse operation binding the contract event 0x6905286e1ecb9b47a50224e76e31cf1383f75212bc0b06c8684317782566a0a7.
//
// Solidity: event ModuloSentEvent(uint256 amount)
func (_StakingWalletRWalk *StakingWalletRWalkFilterer) ParseModuloSentEvent(log types.Log) (*StakingWalletRWalkModuloSentEvent, error) {
	event := new(StakingWalletRWalkModuloSentEvent)
	if err := _StakingWalletRWalk.contract.UnpackLog(event, "ModuloSentEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletRWalkOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakingWalletRWalk contract.
type StakingWalletRWalkOwnershipTransferredIterator struct {
	Event *StakingWalletRWalkOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingWalletRWalkOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletRWalkOwnershipTransferred)
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
		it.Event = new(StakingWalletRWalkOwnershipTransferred)
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
func (it *StakingWalletRWalkOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletRWalkOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletRWalkOwnershipTransferred represents a OwnershipTransferred event raised by the StakingWalletRWalk contract.
type StakingWalletRWalkOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingWalletRWalk *StakingWalletRWalkFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingWalletRWalkOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingWalletRWalk.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRWalkOwnershipTransferredIterator{contract: _StakingWalletRWalk.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingWalletRWalk *StakingWalletRWalkFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingWalletRWalkOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingWalletRWalk.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletRWalkOwnershipTransferred)
				if err := _StakingWalletRWalk.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StakingWalletRWalk *StakingWalletRWalkFilterer) ParseOwnershipTransferred(log types.Log) (*StakingWalletRWalkOwnershipTransferred, error) {
	event := new(StakingWalletRWalkOwnershipTransferred)
	if err := _StakingWalletRWalk.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletRWalkStakeActionEventIterator is returned from FilterStakeActionEvent and is used to iterate over the raw logs and unpacked data for StakeActionEvent events raised by the StakingWalletRWalk contract.
type StakingWalletRWalkStakeActionEventIterator struct {
	Event *StakingWalletRWalkStakeActionEvent // Event containing the contract specifics and raw log

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
func (it *StakingWalletRWalkStakeActionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletRWalkStakeActionEvent)
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
		it.Event = new(StakingWalletRWalkStakeActionEvent)
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
func (it *StakingWalletRWalkStakeActionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletRWalkStakeActionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletRWalkStakeActionEvent represents a StakeActionEvent event raised by the StakingWalletRWalk contract.
type StakingWalletRWalkStakeActionEvent struct {
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
func (_StakingWalletRWalk *StakingWalletRWalkFilterer) FilterStakeActionEvent(opts *bind.FilterOpts, actionId []*big.Int, tokenId []*big.Int, staker []common.Address) (*StakingWalletRWalkStakeActionEventIterator, error) {

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

	logs, sub, err := _StakingWalletRWalk.contract.FilterLogs(opts, "StakeActionEvent", actionIdRule, tokenIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRWalkStakeActionEventIterator{contract: _StakingWalletRWalk.contract, event: "StakeActionEvent", logs: logs, sub: sub}, nil
}

// WatchStakeActionEvent is a free log subscription operation binding the contract event 0x057eba8c4bba00f858e4f586f9c02794abb0df789ef316c741f9073fe2c435db.
//
// Solidity: event StakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, uint256 unstakeTime, address indexed staker)
func (_StakingWalletRWalk *StakingWalletRWalkFilterer) WatchStakeActionEvent(opts *bind.WatchOpts, sink chan<- *StakingWalletRWalkStakeActionEvent, actionId []*big.Int, tokenId []*big.Int, staker []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakingWalletRWalk.contract.WatchLogs(opts, "StakeActionEvent", actionIdRule, tokenIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletRWalkStakeActionEvent)
				if err := _StakingWalletRWalk.contract.UnpackLog(event, "StakeActionEvent", log); err != nil {
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
func (_StakingWalletRWalk *StakingWalletRWalkFilterer) ParseStakeActionEvent(log types.Log) (*StakingWalletRWalkStakeActionEvent, error) {
	event := new(StakingWalletRWalkStakeActionEvent)
	if err := _StakingWalletRWalk.contract.UnpackLog(event, "StakeActionEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletRWalkUnstakeActionEventIterator is returned from FilterUnstakeActionEvent and is used to iterate over the raw logs and unpacked data for UnstakeActionEvent events raised by the StakingWalletRWalk contract.
type StakingWalletRWalkUnstakeActionEventIterator struct {
	Event *StakingWalletRWalkUnstakeActionEvent // Event containing the contract specifics and raw log

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
func (it *StakingWalletRWalkUnstakeActionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletRWalkUnstakeActionEvent)
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
		it.Event = new(StakingWalletRWalkUnstakeActionEvent)
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
func (it *StakingWalletRWalkUnstakeActionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletRWalkUnstakeActionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletRWalkUnstakeActionEvent represents a UnstakeActionEvent event raised by the StakingWalletRWalk contract.
type StakingWalletRWalkUnstakeActionEvent struct {
	ActionId  *big.Int
	TokenId   *big.Int
	TotalNFTs *big.Int
	Staker    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnstakeActionEvent is a free log retrieval operation binding the contract event 0x33940a9fc3ee3c9150b29b34ef29ca397b5e1e340425a4f0da0346b5b90766c8.
//
// Solidity: event UnstakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address indexed staker)
func (_StakingWalletRWalk *StakingWalletRWalkFilterer) FilterUnstakeActionEvent(opts *bind.FilterOpts, actionId []*big.Int, tokenId []*big.Int, staker []common.Address) (*StakingWalletRWalkUnstakeActionEventIterator, error) {

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

	logs, sub, err := _StakingWalletRWalk.contract.FilterLogs(opts, "UnstakeActionEvent", actionIdRule, tokenIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRWalkUnstakeActionEventIterator{contract: _StakingWalletRWalk.contract, event: "UnstakeActionEvent", logs: logs, sub: sub}, nil
}

// WatchUnstakeActionEvent is a free log subscription operation binding the contract event 0x33940a9fc3ee3c9150b29b34ef29ca397b5e1e340425a4f0da0346b5b90766c8.
//
// Solidity: event UnstakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address indexed staker)
func (_StakingWalletRWalk *StakingWalletRWalkFilterer) WatchUnstakeActionEvent(opts *bind.WatchOpts, sink chan<- *StakingWalletRWalkUnstakeActionEvent, actionId []*big.Int, tokenId []*big.Int, staker []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakingWalletRWalk.contract.WatchLogs(opts, "UnstakeActionEvent", actionIdRule, tokenIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletRWalkUnstakeActionEvent)
				if err := _StakingWalletRWalk.contract.UnpackLog(event, "UnstakeActionEvent", log); err != nil {
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
func (_StakingWalletRWalk *StakingWalletRWalkFilterer) ParseUnstakeActionEvent(log types.Log) (*StakingWalletRWalkUnstakeActionEvent, error) {
	event := new(StakingWalletRWalkUnstakeActionEvent)
	if err := _StakingWalletRWalk.contract.UnpackLog(event, "UnstakeActionEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
