package primitives

import (


)
type EConditionPreparation struct {// Eevent of ConditionalToken (Gnosis)
	//Signature: ab3760c3bd2bb38b5bcf54dc79802ed67338b4cf29f3054ded67ed24661e4177
	ConditionId      [32]byte
	Oracle           common.Address
	QuestionId       [32]byte
	OutcomeSlotCount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}
type EConditionResolution struct {// Event of ConditionalToken (Gnosis)
	//Signature: b44d84d3289691f71497564b85d4233648d9dbae8cbdbb4329f301c3a0185894
	ConditionId      [32]byte
	Oracle           common.Address
	QuestionId       [32]byte
	OutcomeSlotCount *big.Int
	PayoutNumerators []*big.Int
	Raw              types.Log // Blockchain specific contextual infos
}
type EPayoutRedemption struct {// Event of ConditionalToken (Gnosis)
	//Signature: 2682012a4a4f1973119f1c9b90745d1bd91fa2bab387344f044cb3586864d18d
	Redeemer           common.Address
	CollateralToken    common.Address
	ParentCollectionId [32]byte
	ConditionId        [32]byte
	IndexSets          []*big.Int
	Payout             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}
type EPositionSplit struct {// Event of ConditionalToken (Gnosis)
	//Signature: 2e6bb91f8cbcda0c93623c54d0403a43514fabc40084ec96b6d5379a74786298
	Stakeholder        common.Address
	CollateralToken    common.Address
	ParentCollectionId [32]byte
	ConditionId        [32]byte
	Partition          []*big.Int
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}
type EPositionsMerge struct { //Event of ConditionalToken (Gnosis)
	//Signature: f13ca62553fcc2bcd2372180a43949c1e4cebba603901ede2f4e14f36b282ca
	Stakeholder        common.Address
	CollateralToken    common.Address
	ParentCollectionId [32]byte
	ConditionId        [32]byte
	Partition          []*big.Int
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}
type EURI struct {// Event of ConditionalToken (Gnosis)
	//Signature: 6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

type PolymarketProcStatus struct {
	LastIdProcessed			int64
}

