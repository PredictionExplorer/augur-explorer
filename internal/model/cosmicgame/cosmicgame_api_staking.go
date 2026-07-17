package cosmicgame

// CGStakingCSTHistoryRec is one CST staking-wallet action (stake or unstake,
// per ActionType) in the global staking history listing.
type CGStakingCSTHistoryRec struct {
	ActionType         int64
	RecordId           int64
	Tx                 Transaction
	UnstakeDate        string
	UnstakeTimeStamp   int64
	ActionId           int64
	TokenId            int64
	RoundNum           int64
	NumStakedNFTs      int64
	AccumNumStakedNFTs int64
	Modulo             string
	ModuloF64          float64
	Claimed            bool
	StakerAid          int64
	StakerAddr         string
}

// CGStakingRWalkHistoryRec is one RandomWalk staking-wallet action (stake or
// unstake, per ActionType) in the global staking history listing.
type CGStakingRWalkHistoryRec struct {
	ActionType            int64
	RecordId              int64
	Tx                    Transaction
	UnstakeDate           string
	UnstakeTimeStamp      int64
	ActionId              int64
	TokenId               int64
	RoundNum              int64
	NumStakedNFTs         int64
	AccumNumStakedNFTs    int64
	StakerAid             int64
	StakerAddr            string
	LastBlockTS           int64
	UnstakeExpirationDiff int64
}

// CGStakeActionCSTRec is one CST staking action row in the per-user action
// listing (without staker identity, which the route scope implies).
type CGStakeActionCSTRec struct {
	ActionType       int64
	RecordId         int64
	Tx               Transaction
	UnstakeDate      string
	UnstakeTimeStamp int64
	ActionId         int64
	TokenId          int64
	NumStakedNFTs    int64
	Modulo           string
	ModuloF64        float64
	Claimed          bool
}

// CGStakeActionRWalkRec is one RandomWalk staking action row in the per-user
// action listing.
type CGStakeActionRWalkRec struct {
	ActionType       int64
	RecordId         int64
	Tx               Transaction
	UnstakeDate      string
	UnstakeTimeStamp int64
	ActionId         int64
	TokenId          int64
	NumStakedNFTs    int64
}

// CGStakeActionInfoRec is the stake half of a stake/unstake action pair.
type CGStakeActionInfoRec struct {
	ActionType    int64
	RecordId      int64
	Tx            Transaction
	ActionId      int64
	TokenId       int64
	RoundNum      int64
	NumStakedNFTs int64
	StakerAid     int64
	StakerAddr    string
}

// CGUnstakeActionInfoRec is the unstake half of a stake/unstake action pair,
// with the reward collected at unstake time.
type CGUnstakeActionInfoRec struct {
	ActionType        int64
	RecordId          int64
	Tx                Transaction
	ActionId          int64
	TokenId           int64
	RoundNum          int64
	NumStakedNFTs     int64
	RewardAmount      string
	RewardAmountEth   float64
	RewardPerToken    string
	RewardPerTokenEth float64
	StakerAid         int64
	StakerAddr        string
}

// CGStakeUnstakeCombined pairs a stake action with its matching unstake (the
// unstake is zero-valued while the token is still staked).
type CGStakeUnstakeCombined struct {
	Stake   CGStakeActionInfoRec
	Unstake CGUnstakeActionInfoRec
}

// CGStakedTokenCSTRec is one currently staked Cosmic Signature token with
// its mint record and the locking stake action.
type CGStakedTokenCSTRec struct {
	TokenInfo      CGCosmicSignatureMintRec
	StakeEvtLogId  int64
	StakeBlockNum  int64
	StakeActionId  int64
	StakeTimeStamp int64
	StakeDateTime  string
	UserAddr       string
	UserAid        int64
}

// CGStakedTokenRWalkRec is one currently staked RandomWalk token with the
// locking stake action.
type CGStakedTokenRWalkRec struct {
	StakeEvtLogId  int64
	StakeBlockNum  int64
	StakeActionId  int64
	StakeTimeStamp int64
	StakeDateTime  string
	StakedTokenId  int64
	UserAddr       string
	UserAid        int64
}

// CGActionIdsForDepositWithClaimInfo links one stake action to one reward
// deposit with the claim state of that (action, deposit) cell.
type CGActionIdsForDepositWithClaimInfo struct {
	RecordId             int64
	DepositId            int64
	UserAid              int64
	StakeActionId        int64
	TokenId              int64
	Claimed              bool
	ClaimBlockNum        int64
	ClaimTimeStamp       int64
	ClaimDateTime        string
	ClaimTxId            int64
	ClaimTxHash          string
	ClaimRewardAmount    string
	ClaimRewardAmountEth float64
}

// CGEthDepositAsReward is one CST staking-wallet ETH deposit viewed from one
// staker's perspective: the staker's share and its collected/pending split.
type CGEthDepositAsReward struct {
	RecordId                int64
	Tx                      Transaction
	DepositDate             string
	DepositTimeStamp        int64
	NumStakedNFTsTotal      int64
	Amount                  string
	AmountEth               float64
	AmountPerToken          string
	AmountPerTokenEth       float64
	StakerAid               int64
	StakerAddr              string
	StakerNumStakedNFTs     int64
	StakerAmount            string
	StakerAmountEth         float64
	AmountCollected         string
	AmountCollectedEth      float64
	AmountPendingToClaim    string
	AmountPendingToClaimEth float64
}

// CGRewardToClaim is one reward deposit a staker still has an uncollected
// share in (the v1 to_claim listing).
type CGRewardToClaim struct {
	RecordId               int64
	Tx                     Transaction
	DepositDate            string
	DepositTimeStamp       int64
	DepositId              int64
	NumStakedNFTs          int64
	DepositAmount          string
	DepositAmountEth       float64
	YourTokensStaked       int64
	YourRewardAmount       string
	YourRewardAmountEth    float64
	YourCollectedAmount    string
	YourCollectedAmountEth float64
	PendingToClaim         string
	PendingToClaimEth      float64
	NumUnclaimedTokens     int64
	AmountPerToken         string
	AmountPerTokenEth      float64
	Modulo                 string
	ModuloF64              float64
}

// CGCollectedReward is one reward deposit a staker has already collected
// from (the v1 collected listing), with the per-deposit accounting.
type CGCollectedReward struct {
	RecordId                 int64
	Tx                       Transaction
	DepositDate              string
	DepositTimeStamp         int64
	DepositId                int64
	NumStakedNFTs            int64
	TotalDepositAmount       string
	TotalDepositAmountEth    float64
	YourTokensStaked         int64
	YourAmountToClaim        string
	YourAmountToClaimEth     float64
	DepositAmountPerToken    string
	DepositAmountPerTokenEth float64
	NumTokensCollected       int64
	YourCollectedAmount      string
	YourCollectedAmountEth   float64
	Modulo                   string
	ModuloF64                float64
	RoundNum                 int64
	StakerAid                int64
	StakerAddr               string
	FullyClaimed             bool
}

// CGStakingRewardGlobal is one CST staking reward deposit in the global
// (staker-independent) view with collected/pending totals.
type CGStakingRewardGlobal struct {
	RecordId                 int64
	Tx                       Transaction
	DepositDate              string
	DepositTimeStamp         int64
	DepositId                int64
	NumStakedNFTs            int64
	TotalDepositAmount       string
	TotalDepositAmountEth    float64
	DepositAmountPerToken    string
	DepositAmountPerTokenEth float64
	Modulo                   string
	ModuloF64                float64
	RoundNum                 int64
	FullyClaimed             bool
	AlreadyCollected         string
	AlreadyCollectedEth      float64
	PendingToCollect         string
	PendingToCollectEth      float64
}

// CGUniqueStakerCST is one row of the unique CST stakers directory with
// lifetime action counts and reward accounting.
type CGUniqueStakerCST struct {
	StakerAid          int64
	StakerAddr         string
	TotalTokensStaked  int64
	NumStakeActions    int64
	NumUnstakeActions  int64
	TotalReward        string
	TotalRewardEth     float64
	UnclaimedReward    string
	UnclaimedRewardEth float64
}

// CGUniqueStakerRWalk is one row of the unique RandomWalk stakers directory
// with lifetime action counts and raffle mints earned by staking.
type CGUniqueStakerRWalk struct {
	StakerAid         int64
	StakerAddr        string
	TotalTokensStaked int64
	NumStakeActions   int64
	NumUnstakeActions int64
	TotalTokensMinted int64
}

// CGUniqueStakersBoth is one row of the dual-stakers directory: a wallet
// that staked both CST and RandomWalk tokens, with both stat blocks.
type CGUniqueStakersBoth struct {
	StakerAid             int64
	StakerAddr            string
	TotalStakedTokensBoth int64
	CSTStats              CGUniqueStakerCST
	RWalkStats            CGUniqueStakerRWalk
}

// CGStakeStatsCST is the global (or per-user) CST staking statistics block.
type CGStakeStatsCST struct {
	TotalNumStakeActions   int64
	TotalNumUnstakeActions int64
	TotalTokensStaked      int64
	TotalReward            string
	TotalRewardEth         float64
	UnclaimedReward        string
	UnclaimedRewardEth     float64
	NumActiveStakers       int64
	NumDeposits            int64
}

// CGStakeStatsRWalk is the global (or per-user) RandomWalk staking
// statistics block.
type CGStakeStatsRWalk struct {
	TotalNumStakeActions   int64
	TotalNumUnstakeActions int64
	TotalTokensStaked      int64
	TotalTokensMinted      int64
	NumActiveStakers       int64
}

// CGNftStakedInfoRec is the stake half of a deposit-scoped action pair.
type CGNftStakedInfoRec struct {
	RecordId      int64
	Tx            Transaction
	ActionId      int64
	TokenId       int64
	NumStakedNFTs int64
	StakerAid     int64
}

// CGNftUnstakedInfoRec is the unstake half of a deposit-scoped action pair,
// with the reward collected at unstake time.
type CGNftUnstakedInfoRec struct {
	RecordId        int64
	Tx              Transaction
	ActionId        int64
	TokenId         int64
	NumStakedNFTs   int64
	RewardAmount    string
	RewardAmountEth float64
	StakerAid       int64
}

// CGNftStakeUnstakeCombined pairs stake and unstake actions under one reward
// deposit with the per-pair reward and claim state.
type CGNftStakeUnstakeCombined struct {
	Stake            CGNftStakedInfoRec
	Unstake          CGNftUnstakedInfoRec
	DepositId        int64
	DepositTimeStamp int64
	DepositDateTime  string
	RoundNum         int64
	Reward           string
	RewardEth        float64
	Claimed          bool
}

// CGCombinedDepositRewardRec is one reward deposit with the staker's action
// pairs nested under it (the deposits-rewards tree view).
type CGCombinedDepositRewardRec struct { // for showing tree-like structure of deposits-rewards
	RecordId               int64
	Tx                     Transaction
	DepositId              int64
	DepositRoundNum        int64
	NumStakedNFTs          int64
	DepositAmount          string
	DepositAmountEth       float64
	YourTokensStaked       int64
	YourClaimableAmountEth float64
	NumTokensCollected     int64
	AmountPerToken         string
	AmountPerTokenEth      float64
	Modulo                 string
	ModuloF64              float64
	FullyClaimed           bool
	ClaimedAmountEth       float64
	Actions                []CGNftStakeUnstakeCombined
}

// CGStakingCstRewardPerTokenRec is one token's collected/pending CST staking
// reward totals for its current owner.
type CGStakingCstRewardPerTokenRec struct {
	TokenId            int64
	RewardCollectedEth float64
	RewardToCollectEth float64
	UserAid            int64
	UserAddr           string
}
