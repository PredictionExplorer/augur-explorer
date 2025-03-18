package cosmicgame

type CGStakingCSTHistoryRec struct {
	ActionType					int64
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	UnstakeEvtLogId				int64
	UnstakeBlockNum				int64
	UnstakeDate					string
	UnstakeTimeStamp			int64
	ActionId					int64
	TokenId						int64
	RoundNum					int64
	NumStakedNFTs				int64
	AccumNumStakedNFTs			int64
	AmountPerHolder				string
	AmountPerHolderEth			float64
	Modulo						string
	ModuloF64					float64
	Claimed						bool
	StakerAid					int64
	StakerAddr					string
//	LastBlockTS					int64 DISCONTINUED, removal pending
//	UnstakeExpirationDiff		int64 DISCONTINUED, removal pending
}
type CGStakingRWalkHistoryRec struct {
	ActionType					int64
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	UnstakeEvtLogId				int64
	UnstakeBlockNum				int64
	UnstakeDate					string
	UnstakeTimeStamp			int64
	ActionId					int64
	TokenId						int64
	RoundNum					int64
	NumStakedNFTs				int64
	AccumNumStakedNFTs			int64
	StakerAid					int64
	StakerAddr					string
	LastBlockTS					int64
	UnstakeExpirationDiff		int64
}
type CGStakeActionCSTRec struct {
	ActionType					int64
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	UnstakeDate					string
	UnstakeTimeStamp			int64
	ActionId					int64
	TokenId						int64
	NumStakedNFTs				int64
	AmountPerHolder				string
	AmountPerHolderEth			float64
	Modulo						string
	ModuloF64					float64
	Claimed						bool
}
type CGStakeActionRWalkRec struct {
	ActionType					int64
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	UnstakeDate					string
	UnstakeTimeStamp			int64
	ActionId					int64
	TokenId						int64
	NumStakedNFTs				int64
}
type CGStakeActionInfoRec struct {
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	ActionId					int64
	TokenId						int64
	RoundNum					int64
	NumStakedNFTs				int64
	UnstakeTimeStamp			int64
	UnstakeDate					string
	StakerAid					int64
	StakerAddr					string
}
type CGUnstakeActionInfoRec struct {
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	ActionId					int64
	TokenId						int64
	RoundNum					int64
	NumStakedNFTs				int64
	StakerAid					int64
	StakerAddr					string
	RewardAmount				string
	RewardAmountEth				float64
}
type CGStakeUnstakeCombined struct {
	Stake						CGStakeActionInfoRec
	Unstake						CGUnstakeActionInfoRec
}
type CGStakedTokenCSTRec struct {
	TokenInfo					CGCosmicSignatureMintRec
	StakeEvtLogId				int64
	StakeBlockNum				int64
	StakeActionId				int64
	StakeTimeStamp				int64
	StakeDateTime				string
	UserAddr					string
	UserAid						int64
}
type CGStakedTokenRWalkRec struct {
	StakeEvtLogId				int64
	StakeBlockNum				int64
	StakeActionId				int64
	StakeTimeStamp				int64
	StakeDateTime				string
	StakedTokenId				int64
	UserAddr					string
	UserAid						int64
}
type CGActionIdsForDeposit struct {
	RecordId					int64
	DepositId					int64
	UserAid						int64
	StakeActionId				int64
	TokenId						int64
	Claimed						bool
	StakeActionTimeStamp		int64
	CurChainTimeStamp			int64
	TimeStampDiff				int64	// subtraction of UnstakeEligibleTimestamp from CurChainTimestamp
	Amount						string
	AmountEth					float64
}
type CGActionIdsForDepositWithClaimInfo struct {
	RecordId					int64
	DepositId					int64
	UserAid						int64
	StakeActionId				int64
	TokenId						int64
	Claimed						bool
	ClaimBlockNum				int64
	ClaimTimeStamp				int64
	ClaimDateTime				string
	ClaimTxId					int64
	ClaimTxHash					string
	ClaimRewardAmount			string
	ClaimRewardAmountEth		float64
}
type CGStakingRewardRec struct {
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	RoundNum					int64
	Amount						string
	AmountEth					float64
	StakerAid					int64
	StakerAddr					string
}
type CGEthDepositAsReward struct {
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	DepositDate					string
	DepositTimeStamp			int64
	NumStakedNFTsTotal			int64
	Amount						string
	AmountEth					float64
	AmountPerToken				string
	AmountPerTokenEth			float64
	StakerAid					int64
	StakerAddr					string
	StakerAmount				string
	StakerAmountEth				float64
	StakerNumStakedNFTs			int64
}
type CGEthDepositRec struct {
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	DepositDate					string
	DepositTimeStamp			int64
	NumStakedNFTs				int64
	Amount						string
	AmountEth					float64
	AmountPerHolder				string
	AmountPerHolderEth			float64
	Modulo						string
	ModuloF64					float64
}
type CGRewardToClaim struct {
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	DepositDate					string
	DepositTimeStamp			int64
	DepositId					int64
	NumStakedNFTs				int64
	DepositAmount				string
	DepositAmountEth			float64
	YourTokensStaked			int64
	YourRewardAmount			string
	YourRewardAmountEth			float64
	YourCollectedAmount			string
	YourCollectedAmountEth		float64
	PendingToClaim				string
	PendingToClaimEth			float64
	NumUnclaimedTokens			int64
	AmountPerToken				string
	AmountPerTokenEth			float64
	Modulo						string
	ModuloF64					float64
}
type CGCollectedReward struct {
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	DepositDate					string
	DepositTimeStamp			int64
	DepositId					int64
	NumStakedNFTs				int64
	TotalDepositAmount			string
	TotalDepositAmountEth		float64
	YourTokensStaked			int64
	YourAmountToClaim			string
	YourAmountToClaimEth		float64
	DepositAmountPerToken		string
	DepositAmountPerTokenEth	float64
	NumTokensCollected			int64
	YourCollectedAmount			string
	YourCollectedAmountEth		float64
	Modulo						string
	ModuloF64					float64
	RoundNum					int64
	StakerAid					int64
	StakerAddr					string
	FullyClaimed				bool
}
type CGStakingRewardGlobal	struct {
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	DepositDate					string
	DepositTimeStamp			int64
	DepositId					int64
	NumStakedNFTs				int64
	TotalDepositAmount			string
	TotalDepositAmountEth		float64
	DepositAmountPerToken		string
	DepositAmountPerTokenEth	float64
	Modulo						string
	ModuloF64					float64
	RoundNum					int64
	FullyClaimed				bool
	AlreadyCollected			string
	AlreadyCollectedEth			float64
	PendingToCollect			string
	PendingToCollectEth			float64
}
type CGUniqueStakerCST struct {
	StakerAid					int64
	StakerAddr					string
	TotalTokensStaked			int64
	NumStakeActions				int64
	NumUnstakeActions			int64
	TotalReward					string
	TotalRewardEth				float64
	UnclaimedReward				string
	UnclaimedRewardEth			float64
	TotalTokensMinted			int64
}
type CGUniqueStakerRWalk struct {
	StakerAid					int64
	StakerAddr					string
	TotalTokensStaked			int64
	NumStakeActions				int64
	NumUnstakeActions			int64
	TotalTokensMinted			int64
}
type CGUniqueStakersBoth struct {
	StakerAid					int64
	StakerAddr					string
	TotalStakedTokensBoth		int64
	CSTStats					CGUniqueStakerCST
	RWalkStats					CGUniqueStakerRWalk
}
type CGStakeStatsCST struct {
	TotalNumStakeActions		int64
	TotalNumUnstakeActions		int64
	TotalTokensStaked			int64
	TotalReward					string
	TotalRewardEth				float64
	UnclaimedReward				string
	UnclaimedRewardEth			float64
	NumActiveStakers			int64
	NumDeposits					int64
	TotalTokensMinted			int64		// if CosmicGame is configured to mint NFTs for CST stakers, this counts tokens minted
}
type CGStakeStatsRWalk struct {
	TotalNumStakeActions		int64
	TotalNumUnstakeActions		int64
	TotalTokensStaked			int64
	TotalTokensMinted			int64
	NumActiveStakers			int64
}
type CGNftStakedInfoRec struct {
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	ActionId					int64
	TokenId						int64
	RoundNum					int64
	NumStakedNFTs				int64
	UnstakeTimeStamp			int64
	UnstakeDate					string
	StakerAid					int64
	StakerAddr					string
}
type CGNftUnstakedInfoRec struct {
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	ActionId					int64
	TokenId						int64
	RoundNum					int64
	NumStakedNFTs				int64
	StakerAid					int64
	StakerAddr					string
	RewardAmount				string
	RewardAmountEth				float64
}
type CGNftStakeUnstakeCombined struct {
	Stake						CGNftStakedInfoRec
	Unstake						CGNftUnstakedInfoRec
	DepositId					int64
	DepositTimeStamp			int64
	DepositDateTime				string
	RoundNum					int64
	Reward						string
	RewardEth					float64
	Claimed						bool
}
type CGCombinedDepositRewardRec struct {	// for showing tree-like structure of deposits-rewards
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	DepositId					int64
	DepositRoundNum				int64
	NumStakedNFTs				int64
	DepositAmount				string
	DepositAmountEth			float64
	YourTokensStaked			int64
	YourClaimableAmount			string
	YourClaimableAmountEth		float64
	NumTokensCollected			int64
	AmountPerToken				string
	AmountPerTokenEth			float64
	Modulo						string
	ModuloF64					float64
	FullyClaimed				bool
	ClaimedAmount				string
	ClaimedAmountEth			float64
	Actions						[]CGNftStakeUnstakeCombined
}
type CGRewardPaidRec struct {
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	ActionId					int64
	TokenId						int64
	StakerAid					int64
	StakerAddr					string
	RewardAmount				string
	RewardAmountEth				float64
}
type CGStakingCstRewardPerTokenRec struct {
	TokenId						int64
	RewardCollectedEth			float64
	RewardToCollectEth			float64
	UserAid						int64
	UserAddr					string
}
type UserStakingInfo struct {
	CSTStakingInfo				CGStakeStatsCST
	RWalkStakingInfo			CGStakeStatsRWalk
}
