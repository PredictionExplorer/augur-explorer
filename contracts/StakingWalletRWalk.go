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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractRandomWalkNFT\",\"name\":\"rwalk_\",\"type\":\"address\"},{\"internalType\":\"contractCosmicGame\",\"name\":\"game_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"}],\"name\":\"AccessError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"NoTokensStaked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"OneTimeStaking\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"TokenAlreadyDeleted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"}],\"name\":\"TokenAlreadyInserted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"}],\"name\":\"TokenAlreadyUnstaked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalNFTs\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"StakeActionEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalNFTs\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"UnstakeActionEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ETHDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"depositTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numStaked\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"game\",\"outputs\":[{\"internalType\":\"contractCosmicGame\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isTokenStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"lastActionIdByTokenId\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastActionIds\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numETHDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakeActions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakedNFTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numTokensStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"entropy\",\"type\":\"bytes32\"}],\"name\":\"pickRandomStaker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomWalk\",\"outputs\":[{\"internalType\":\"contractRandomWalkNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeActions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"stakerByTokenId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenIndex\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIndices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"unstakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"usedTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"wasTokenUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052346100305761001a61001461014b565b906103cc565b610022610035565b61211261052e823961211290f35b61003b565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906100679061003f565b810190811060018060401b0382111761007f57604052565b610049565b90610097610090610035565b928361005d565b565b5f80fd5b60018060a01b031690565b6100b19061009d565b90565b6100bd906100a8565b90565b6100c9816100b4565b036100d057565b5f80fd5b905051906100e1826100c0565b565b6100ec9061009d565b90565b6100f8906100e3565b90565b610104816100ef565b0361010b57565b5f80fd5b9050519061011c826100fb565b565b9190604083820312610146578061013a610143925f86016100d4565b9360200161010f565b90565b610099565b6101696126408038038061015e81610084565b92833981019061011e565b9091565b90565b61018461017f6101899261009d565b61016d565b61009d565b90565b61019590610170565b90565b6101a19061018c565b90565b90565b6101bb6101b66101c0926101a4565b61016d565b61009d565b90565b6101cc906101a7565b90565b60209181520190565b60207f616e646f6d57616c6b20746f6b656e2e00000000000000000000000000000000917f5a65726f2d616464726573732077617320676976656e20666f722074686520525f8201520152565b61023260306040926101cf565b61023b816101d8565b0190565b6102549060208101905f818303910152610225565b90565b1561025e57565b610266610035565b63eac0d38960e01b81528061027d6004820161023f565b0390fd5b61028a9061018c565b90565b60207f616d652e00000000000000000000000000000000000000000000000000000000917f5a65726f2d616464726573732077617320676976656e20666f722074686520675f8201520152565b6102e760246040926101cf565b6102f08161028d565b0190565b6103099060208101905f8183039101526102da565b90565b1561031357565b61031b610035565b63eac0d38960e01b815280610332600482016102f4565b0390fd5b5f1b90565b9061034c60018060a01b0391610336565b9181191691161790565b61035f90610170565b90565b61036b90610356565b90565b90565b9061038661038161038d92610362565b61036e565b825461033b565b9055565b61039a90610170565b90565b6103a690610391565b90565b90565b906103c16103bc6103c89261039d565b6103a9565b825461033b565b9055565b9061043d610444926103dc610446565b6104096103e882610198565b6104026103fc6103f75f6101c3565b6100a8565b916100a8565b1415610257565b61043661041584610281565b61042f6104296104245f6101c3565b6100a8565b916100a8565b141561030c565b600a610371565b600b6103ac565b565b61045661045161045c565b6104ce565b565b5f90565b610464610458565b503390565b5f1c90565b60018060a01b031690565b61048561048a91610469565b61046e565b90565b6104979054610479565b90565b6104a39061018c565b90565b90565b906104be6104b96104c59261049a565b6104a6565b825461033b565b9055565b5f0190565b6104d75f61048d565b6104e1825f6104a9565b9061051561050f7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09361049a565b9161049a565b9161051e610035565b80610528816104c9565b0390a356fe60806040526004361015610013575b610de0565b61001d5f356101ac565b80630d50c189146101a75780630f7ee879146101a257806317db62131461019d5780632a3247aa146101985780632e17de7814610193578063418104251461018e57806344d110b914610189578063451f1adf146101845780634f6ccce71461017f5780635111a2d61461017a57806355279fdb146101755780635fda0acc146101705780636427d9a91461016b578063715018a614610166578063889d1e1a146101615780638da5cb5b1461015c578063a2b136fb14610157578063a531aa8614610152578063a694fc3a1461014d578063c065894e14610148578063c078855514610143578063c3fe3e281461013e578063f0a5242414610139578063f2fde38b146101345763fe939afc0361000e57610dad565b610d7a565b610d04565b610ccf565b610c20565b610b76565b610b43565b610b0e565b610ac6565b6109d1565b61099c565b610969565b610934565b610882565b61084d565b610809565b61073e565b610706565b610631565b6105cd565b61051d565b6104e8565b6104b3565b610415565b61031f565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906101f0906101c8565b810190811067ffffffffffffffff82111761020a57604052565b6101d2565b9061022261021b6101b2565b92836101e6565b565b67ffffffffffffffff811161023c5760208091020190565b6101d2565b5f80fd5b90565b61025181610245565b0361025857565b5f80fd5b9050359061026982610248565b565b9092919261028061027b82610224565b61020f565b93818552602080860192028301928184116102bd57915b8383106102a45750505050565b602080916102b2848661025c565b815201920191610297565b610241565b9080601f830112156102e0578160206102dd9335910161026b565b90565b6101c4565b90602082820312610315575f82013567ffffffffffffffff81116103105761030d92016102c2565b90565b6101c0565b6101bc565b5f0190565b3461034d576103376103323660046102e5565b610e43565b61033f6101b2565b806103498161031a565b0390f35b6101b8565b9060208282031261036b57610368915f0161025c565b90565b6101bc565b90565b61038761038261038c92610245565b610370565b610245565b90565b9061039990610373565b5f5260205260405f2090565b1c90565b60ff1690565b6103bf9060086103c493026103a5565b6103a9565b90565b906103d291546103af565b90565b6103eb906103e66003915f9261038f565b6103c7565b90565b151590565b6103fc906103ee565b9052565b9190610413905f602085019401906103f3565b565b346104455761044161043061042b366004610352565b6103d5565b6104386101b2565b91829182610400565b0390f35b6101b8565b5f91031261045457565b6101bc565b90565b61046c90600861047193026103a5565b610459565b90565b9061047f915461045c565b90565b61048e60095f90610474565b90565b61049a90610245565b9052565b91906104b1905f60208501940190610491565b565b346104e3576104c336600461044a565b6104df6104ce610482565b6104d66101b2565b9182918261049e565b0390f35b6101b8565b34610518576105146105036104fe366004610352565b610ec2565b61050b6101b2565b91829182610400565b0390f35b6101b8565b3461054b57610535610530366004610352565b61115d565b61053d6101b2565b806105478161031a565b0390f35b6101b8565b90565b61055c81610550565b0361056357565b5f80fd5b9050359061057482610553565b565b9060208282031261058f5761058c915f01610567565b90565b6101bc565b60018060a01b031690565b6105a890610594565b90565b6105b49061059f565b9052565b91906105cb905f602085019401906105ab565b565b346105fd576105f96105e86105e3366004610576565b611469565b6105f06101b2565b918291826105b8565b0390f35b6101b8565b9061060c90610373565b5f5260205260405f2090565b61062e906106296005915f92610602565b610474565b90565b346106615761065d61064c610647366004610352565b610618565b6106546101b2565b9182918261049e565b0390f35b6101b8565b9061067090610373565b5f5260205260405f2090565b5f1c90565b61068d6106929161067c565b610459565b90565b61069f9054610681565b90565b6106ad906007610666565b6106b85f8201610695565b916106d160026106ca60018501610695565b9301610695565b90565b6040906106fd61070494969593966106f360608401985f850190610491565b6020830190610491565b0190610491565b565b346107395761073561072161071c366004610352565b6106a2565b61072c9391936101b2565b938493846106d4565b0390f35b6101b8565b3461076e5761076a610759610754366004610352565b6114ff565b6107616101b2565b9182918261049e565b0390f35b6101b8565b60018060a01b031690565b61078e90600861079393026103a5565b610773565b90565b906107a1915461077e565b90565b6107b0600a5f90610796565b90565b6107c76107c26107cc92610594565b610370565b610594565b90565b6107d8906107b3565b90565b6107e4906107cf565b90565b6107f0906107db565b9052565b9190610807905f602085019401906107e7565b565b346108395761081936600461044a565b6108356108246107a4565b61082c6101b2565b918291826107f4565b0390f35b6101b8565b61084a60085f90610474565b90565b3461087d5761085d36600461044a565b61087961086861083e565b6108706101b2565b9182918261049e565b0390f35b6101b8565b346108b25761089236600461044a565b6108ae61089d61151f565b6108a56101b2565b9182918261049e565b0390f35b6101b8565b906108c190610373565b5f5260205260405f2090565b90565b6108e09060086108e593026103a5565b6108cd565b90565b906108f391546108d0565b90565b61090c906109076006915f926108b7565b6108e8565b90565b90565b61091b9061090f565b9052565b9190610932905f60208501940190610912565b565b346109645761096061094f61094a366004610352565b6108f6565b6109576101b2565b9182918261091f565b0390f35b6101b8565b346109975761097936600461044a565b610981611625565b6109896101b2565b806109938161031a565b0390f35b6101b8565b346109cc576109c86109b76109b2366004610352565b611652565b6109bf6101b2565b9182918261091f565b0390f35b6101b8565b34610a01576109e136600461044a565b6109fd6109ec6116ab565b6109f46101b2565b918291826105b8565b0390f35b6101b8565b90610a1090610373565b5f5260205260405f2090565b60018060a01b031690565b610a33610a389161067c565b610a1c565b90565b610a459054610a27565b90565b610a53906001610a06565b90610a5f5f8301610695565b91610a6c60018201610a3b565b91610a856003610a7e60028501610695565b9301610695565b90565b610abd610ac494610ab3606094989795610aa9608086019a5f870190610491565b60208501906105ab565b6040830190610491565b0190610491565b565b34610afa57610af6610ae1610adc366004610352565b610a48565b90610aed9492946101b2565b94859485610a88565b0390f35b6101b8565b610b0b60025f90610474565b90565b34610b3e57610b1e36600461044a565b610b3a610b29610aff565b610b316101b2565b9182918261049e565b0390f35b6101b8565b34610b7157610b5b610b56366004610352565b611822565b610b636101b2565b80610b6d8161031a565b0390f35b6101b8565b34610ba657610ba2610b91610b8c366004610352565b611a61565b610b996101b2565b918291826105b8565b0390f35b6101b8565b634e487b7160e01b5f52603260045260245ffd5b5490565b5f5260205f2090565b610bd581610bbf565b821015610bef57610be7600191610bc3565b910201905f90565b610bab565b6004610bff81610bbf565b821015610c1c57610c1991610c1391610bcc565b90610474565b90565b5f80fd5b34610c5057610c4c610c3b610c36366004610352565b610bf4565b610c436101b2565b9182918261049e565b0390f35b6101b8565b60018060a01b031690565b610c70906008610c7593026103a5565b610c55565b90565b90610c839154610c60565b90565b610c92600b5f90610c78565b90565b610c9e906107b3565b90565b610caa90610c95565b90565b610cb690610ca1565b9052565b9190610ccd905f60208501940190610cad565b565b34610cff57610cdf36600461044a565b610cfb610cea610c86565b610cf26101b2565b91829182610cba565b0390f35b6101b8565b34610d3457610d30610d1f610d1a366004610352565b611ac3565b610d276101b2565b91829182610400565b0390f35b6101b8565b610d428161059f565b03610d4957565b5f80fd5b90503590610d5a82610d39565b565b90602082820312610d7557610d72915f01610d4d565b90565b6101bc565b34610da857610d92610d8d366004610d5c565b611c03565b610d9a6101b2565b80610da48161031a565b0390f35b6101b8565b34610ddb57610dc5610dc03660046102e5565b611c0e565b610dcd6101b2565b80610dd78161031a565b0390f35b6101b8565b5f80fd5b90565b610dfb610df6610e0092610de4565b610370565b610245565b90565b6001610e0f9101610245565b90565b5190565b90610e2082610e12565b811015610e31576020809102010190565b610bab565b610e409051610245565b90565b90610e4d5f610de7565b5b80610e69610e63610e5e86610e12565b610245565b91610245565b1015610e9857610e9390610e8e610e89610e84868490610e16565b610e36565b61115d565b610e03565b610e4e565b509050565b5f90565b610ead610eb29161067c565b6103a9565b90565b610ebf9054610ea1565b90565b610ed9610ede91610ed1610e9d565b50600361038f565b610eb5565b610ef1610eeb60016103ee565b916103ee565b1490565b60209181520190565b5f7f546f6b656e2068617320616c7265616479206265656e20756e7374616b65642e910152565b610f3160208092610ef5565b610f3a81610efe565b0190565b9190610f61906020610f59604086018681035f880152610f25565b940190610491565b565b15610f6b5750565b610f8d90610f776101b2565b91829163aed59e4f60e01b835260048301610f3e565b0390fd5b5f7f4f6e6c7920746865206f776e65722063616e20756e7374616b652e0000000000910152565b610fc5601b602092610ef5565b610fce81610f91565b0190565b916040611003929493610ffc610ff1606083018381035f850152610fb8565b966020830190610491565b01906105ab565b565b1561100e575050565b61102f6110196101b2565b9283926345c2e43b60e01b845260048401610fd2565b0390fd5b61103f6110449161067c565b610773565b90565b6110519054611033565b90565b61105d906107cf565b90565b5f80fd5b60e01b90565b5f91031261107457565b6101bc565b6040906110a26110a9949695939661109860608401985f8501906105ab565b60208301906105ab565b0190610491565b565b6110b36101b2565b3d5f823e3d90fd5b5f1b90565b906110cc5f19916110bb565b9181191691161790565b90565b906110ee6110e96110f592610373565b6110d6565b82546110c0565b9055565b90565b61111061110b611115926110f9565b610370565b610245565b90565b634e487b7160e01b5f52601160045260245ffd5b61113b61114191939293610245565b92610245565b820391821161114c57565b611118565b61115a906107cf565b90565b611191611177600361117160018590610a06565b01610695565b6111896111835f610de7565b91610245565b148290610f63565b6111c56111aa60016111a4818590610a06565b01610a3b565b6111bc6111b63361059f565b9161059f565b14823391611005565b6111db5f6111d560018490610a06565b01610695565b906111e582611e34565b6111f76111f2600a611047565b6107db565b6323b872dd61120530611054565b339261121d5f61121760018890610a06565b01610695565b92813b15611330575f6112439161124e82966112376101b2565b98899788968795611064565b855260048501611079565b03925af1801561132b576112ff575b5061127642600361127060018590610a06565b016110d9565b61129c61129561128660016110fc565b6112906009610695565b61112c565b60096110d9565b6112a66009610695565b9133916112fa6112e86112e26112dc7f33940a9fc3ee3c9150b29b34ef29ca397b5e1e340425a4f0da0346b5b90766c894610373565b94610373565b94611151565b946112f16101b2565b9182918261049e565b0390a4565b61131e905f3d8111611324575b61131681836101e6565b81019061106a565b5f61125d565b503d61130c565b6110ab565b611060565b5f90565b60207f74616b65642e0000000000000000000000000000000000000000000000000000917f546865726520617265206e6f2052616e646f6d57616c6b20746f6b656e7320735f8201520152565b6113936026604092610ef5565b61139c81611339565b0190565b6113b59060208101905f818303910152611386565b90565b156113bf57565b6113c76101b2565b63bc8b155960e01b8152806113de600482016113a0565b0390fd5b6113ee6113f39161067c565b610373565b90565b634e487b7160e01b5f52601260045260245ffd5b61141661141c91610245565b91610245565b908115611427570690565b6113f6565b61143861143d9161067c565b6108cd565b90565b61144a905461142c565b90565b61146161145c6114669261090f565b610370565b610245565b90565b60016114f26114e26114dd6114d66114d06114f896611486611335565b506114ac6114946004610bbf565b6114a66114a05f610de7565b91610245565b116113b8565b6114ca6114ba6004926113e2565b6114c46004610bbf565b9061140a565b90610bcc565b90610474565b60066108b7565b611440565b6114ec839161144d565b90610a06565b01610a3b565b90565b5f90565b61151661151c9161150e6114fb565b506004610bcc565b90610474565b90565b6115276114fb565b506115326004610bbf565b90565b5f7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572910152565b61156860208092610ef5565b61157181611535565b0190565b61158a9060208101905f81830391015261155c565b90565b1561159457565b61159c6101b2565b62461bcd60e51b8152806115b260048201611575565b0390fd5b6115e06115c16116ab565b6115da6115d46115cf611f10565b61059f565b9161059f565b1461158d565b6115e8611612565b565b6115fe6115f961160392610de4565b610370565b610594565b90565b61160f906115ea565b90565b61162361161e5f611606565b611f1d565b565b61162d6115b6565b565b5f90565b90565b61164a61164561164f92611633565b610370565b61090f565b90565b61165a61162f565b5061166f61166a60058390610602565b610695565b61168161167b5f610de7565b91610245565b1461169c576116946116999160066108b7565b611440565b90565b506116a8600119611636565b90565b6116b3611335565b506116bd5f610a3b565b90565b60207f6564206f6e6c79206f6e63650000000000000000000000000000000000000000917f5374616b696e672f756e7374616b696e6720746f6b656e20697320616c6c6f775f8201520152565b61171a602c604092610ef5565b611723816116c0565b0190565b919061174a906020611742604086018681035f88015261170d565b940190610491565b565b156117545750565b611776906117606101b2565b918291632290948760e21b835260048301611727565b0390fd5b9061178660ff916110bb565b9181191691161790565b611799906103ee565b90565b90565b906117b46117af6117bb92611790565b61179c565b825461177a565b9055565b906117d060018060a01b03916110bb565b9181191691161790565b90565b906117f26117ed6117f992611151565b6117da565b82546117bf565b9055565b61180c61181291939293610245565b92610245565b820180921161181d57565b611118565b6118556118396118346003849061038f565b610eb5565b61184c61184660016103ee565b916103ee565b1415829061174c565b61186b60016118666003849061038f565b61179f565b61187d611878600a611047565b6107db565b6323b872dd3361188c30611054565b928492813b15611a40575f6118b4916118bf82966118a86101b2565b98899788968795611064565b855260048501611079565b03925af18015611a3b57611a0f575b506118e3816118dd6002610695565b9061206f565b611903815f6118fd60016118f76002610695565b90610a06565b016110d9565b61192333600161191d816119176002610695565b90610a06565b016117dd565b61194342600261193d600161193783610695565b90610a06565b016110d9565b61196961196261195360016110fc565b61195d6002610695565b6117fd565b60026110d9565b61198f61198861197960016110fc565b6119836009610695565b6117fd565b60096110d9565b6119ac61199c6002610695565b6119a660016110fc565b9061112c565b6119b66009610695565b913391611a0a6119f86119f26119ec7fde0c27ec227b5d0c0a194ea6f25d7031639f9c10d8bf6e2f9414ff8088a6e20d94610373565b94610373565b94611151565b94611a016101b2565b9182918261049e565b0390a4565b611a2e905f3d8111611a34575b611a2681836101e6565b81019061106a565b5f6118ce565b503d611a1c565b6110ab565b611060565b611a59611a54611a5e92610de4565b610370565b61090f565b90565b611a7c90611a6d611335565b50611a7661162f565b50611652565b80611a8f611a895f611a45565b9161090f565b12611ab6576001611aad611ab392611aa7839161144d565b90610a06565b01610a3b565b90565b50611ac05f611606565b90565b611ada611adf91611ad2610e9d565b506005610602565b610695565b611af1611aeb5f610de7565b91610245565b141590565b611b2990611b24611b056116ab565b611b1e611b18611b13611f10565b61059f565b9161059f565b1461158d565b611bd3565b565b60207f6464726573730000000000000000000000000000000000000000000000000000917f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201520152565b611b856026604092610ef5565b611b8e81611b2b565b0190565b611ba79060208101905f818303910152611b78565b90565b15611bb157565b611bb96101b2565b62461bcd60e51b815280611bcf60048201611b92565b0390fd5b611c0190611bfc81611bf5611bef611bea5f611606565b61059f565b9161059f565b1415611baa565b611f1d565b565b611c0c90611af6565b565b90611c185f610de7565b5b80611c34611c2e611c2986610e12565b610245565b91610245565b1015611c6357611c5e90611c59611c54611c4f868490610e16565b610e36565b611822565b610e03565b611c19565b509050565b5f7f546f6b656e206973206e6f7420696e20746865206c6973742e00000000000000910152565b611c9c6019602092610ef5565b611ca581611c68565b0190565b9190611ccc906020611cc4604086018681035f880152611c8f565b940190610491565b565b15611cd65750565b611cf890611ce26101b2565b918291639aa6fa6560e01b835260048301611ca9565b0390fd5b1b90565b91906008611d1b910291611d155f1984611cfc565b92611cfc565b9181191691161790565b9190611d3b611d36611d4393610373565b6110d6565b908354611d00565b9055565b611d5991611d536114fb565b91611d25565b565b90565b634e487b7160e01b5f52603160045260245ffd5b5490565b5f5260205f2090565b611d8881611d72565b821015611da257611d9a600191611d76565b910201905f90565b610bab565b611db081611d72565b8015611dd1576001900390611dce611dc88383611d7f565b90611d47565b55565b611d5e565b90565b611ded611de8611df292611dd6565b610370565b61090f565b90565b611e09611e04611e0e9261090f565b610370565b61090f565b90565b90565b90611e29611e24611e3092611df5565b611e11565b82546110c0565b9055565b611f0e90611e4b611e4482611ac3565b8290611cce565b611ecd611e62611e5d60058490610602565b610695565b611ec8611e96611e906004611e8a611e7a6004610bbf565b611e8460016110fc565b9061112c565b90610bcc565b90610474565b91611ec083611eba6004611eb485611eae60016110fc565b9061112c565b90610bcc565b90611d25565b916005610602565b6110d9565b611ee25f611edd60058490610602565b611d47565b611ef4611eef6004611d5b565b611da7565b611f09611f015f19611dd9565b9160066108b7565b611e14565b565b611f18611335565b503390565b611f265f610a3b565b611f30825f6117dd565b90611f64611f5e7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e093611151565b91611151565b91611f6d6101b2565b80611f778161031a565b0390a3565b5f7f546f6b656e20616c726561647920696e20746865206c6973742e000000000000910152565b611fb0601a602092610ef5565b611fb981611f7c565b0190565b916040611fee929493611fe7611fdc606083018381035f850152611fa3565b966020830190610491565b0190610491565b565b15611ff9575050565b61201a6120046101b2565b92839263597558c560e11b845260048401611fbd565b0390fd5b908154916801000000000000000083101561204e578261204691600161204c95018155611d7f565b90611d25565b565b6101d2565b61206761206261206c92610245565b610370565b61090f565b90565b6120d56120cd6120da9361209661208e61208886611ac3565b156103ee565b858391611ff0565b6120aa6120a36004611d5b565b859061201e565b6120c86120b76004610bbf565b6120c360058790610602565b6110d9565b612053565b9160066108b7565b611e14565b56fea26469706673582212206129591c942d32791ad2131786c06ad37792f42eb1f769fec3f4ae08d02b1d0364736f6c634300081a0033",
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
// Solidity: function ETHDeposits(uint256 ) view returns(uint256 depositTime, uint256 depositAmount, uint256 numStaked)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) ETHDeposits(opts *bind.CallOpts, arg0 *big.Int) (struct {
	DepositTime   *big.Int
	DepositAmount *big.Int
	NumStaked     *big.Int
}, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "ETHDeposits", arg0)

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
func (_StakingWalletRWalk *StakingWalletRWalkSession) ETHDeposits(arg0 *big.Int) (struct {
	DepositTime   *big.Int
	DepositAmount *big.Int
	NumStaked     *big.Int
}, error) {
	return _StakingWalletRWalk.Contract.ETHDeposits(&_StakingWalletRWalk.CallOpts, arg0)
}

// ETHDeposits is a free data retrieval call binding the contract method 0x451f1adf.
//
// Solidity: function ETHDeposits(uint256 ) view returns(uint256 depositTime, uint256 depositAmount, uint256 numStaked)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) ETHDeposits(arg0 *big.Int) (struct {
	DepositTime   *big.Int
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
// Solidity: function stakeActions(uint256 ) view returns(uint256 tokenId, address owner, uint256 stakeTime, uint256 unstakeTime)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) StakeActions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TokenId     *big.Int
	Owner       common.Address
	StakeTime   *big.Int
	UnstakeTime *big.Int
}, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "stakeActions", arg0)

	outstruct := new(struct {
		TokenId     *big.Int
		Owner       common.Address
		StakeTime   *big.Int
		UnstakeTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.StakeTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UnstakeTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// StakeActions is a free data retrieval call binding the contract method 0xa2b136fb.
//
// Solidity: function stakeActions(uint256 ) view returns(uint256 tokenId, address owner, uint256 stakeTime, uint256 unstakeTime)
func (_StakingWalletRWalk *StakingWalletRWalkSession) StakeActions(arg0 *big.Int) (struct {
	TokenId     *big.Int
	Owner       common.Address
	StakeTime   *big.Int
	UnstakeTime *big.Int
}, error) {
	return _StakingWalletRWalk.Contract.StakeActions(&_StakingWalletRWalk.CallOpts, arg0)
}

// StakeActions is a free data retrieval call binding the contract method 0xa2b136fb.
//
// Solidity: function stakeActions(uint256 ) view returns(uint256 tokenId, address owner, uint256 stakeTime, uint256 unstakeTime)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) StakeActions(arg0 *big.Int) (struct {
	TokenId     *big.Int
	Owner       common.Address
	StakeTime   *big.Int
	UnstakeTime *big.Int
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

// UsedTokens is a free data retrieval call binding the contract method 0x0f7ee879.
//
// Solidity: function usedTokens(uint256 ) view returns(bool)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) UsedTokens(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "usedTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedTokens is a free data retrieval call binding the contract method 0x0f7ee879.
//
// Solidity: function usedTokens(uint256 ) view returns(bool)
func (_StakingWalletRWalk *StakingWalletRWalkSession) UsedTokens(arg0 *big.Int) (bool, error) {
	return _StakingWalletRWalk.Contract.UsedTokens(&_StakingWalletRWalk.CallOpts, arg0)
}

// UsedTokens is a free data retrieval call binding the contract method 0x0f7ee879.
//
// Solidity: function usedTokens(uint256 ) view returns(bool)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) UsedTokens(arg0 *big.Int) (bool, error) {
	return _StakingWalletRWalk.Contract.UsedTokens(&_StakingWalletRWalk.CallOpts, arg0)
}

// WasTokenUsed is a free data retrieval call binding the contract method 0x2a3247aa.
//
// Solidity: function wasTokenUsed(uint256 _tokenId) view returns(bool)
func (_StakingWalletRWalk *StakingWalletRWalkCaller) WasTokenUsed(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _StakingWalletRWalk.contract.Call(opts, &out, "wasTokenUsed", _tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WasTokenUsed is a free data retrieval call binding the contract method 0x2a3247aa.
//
// Solidity: function wasTokenUsed(uint256 _tokenId) view returns(bool)
func (_StakingWalletRWalk *StakingWalletRWalkSession) WasTokenUsed(_tokenId *big.Int) (bool, error) {
	return _StakingWalletRWalk.Contract.WasTokenUsed(&_StakingWalletRWalk.CallOpts, _tokenId)
}

// WasTokenUsed is a free data retrieval call binding the contract method 0x2a3247aa.
//
// Solidity: function wasTokenUsed(uint256 _tokenId) view returns(bool)
func (_StakingWalletRWalk *StakingWalletRWalkCallerSession) WasTokenUsed(_tokenId *big.Int) (bool, error) {
	return _StakingWalletRWalk.Contract.WasTokenUsed(&_StakingWalletRWalk.CallOpts, _tokenId)
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
	ActionId  *big.Int
	TokenId   *big.Int
	TotalNFTs *big.Int
	Staker    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeActionEvent is a free log retrieval operation binding the contract event 0xde0c27ec227b5d0c0a194ea6f25d7031639f9c10d8bf6e2f9414ff8088a6e20d.
//
// Solidity: event StakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address indexed staker)
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

// WatchStakeActionEvent is a free log subscription operation binding the contract event 0xde0c27ec227b5d0c0a194ea6f25d7031639f9c10d8bf6e2f9414ff8088a6e20d.
//
// Solidity: event StakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address indexed staker)
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

// ParseStakeActionEvent is a log parse operation binding the contract event 0xde0c27ec227b5d0c0a194ea6f25d7031639f9c10d8bf6e2f9414ff8088a6e20d.
//
// Solidity: event StakeActionEvent(uint256 indexed actionId, uint256 indexed tokenId, uint256 totalNFTs, address indexed staker)
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
