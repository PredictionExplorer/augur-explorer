package primitives

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

)
//------------ ConditionalTokens contract events
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
//------------------------------ FixedPriceMarketMaker contract events
type EFundingAdded struct {
	//Signature: ec2dc3e5a3bb9aa0a1deb905d2bd23640d07f107e6ceb484024501aad964a951.
	Funder       common.Address
	AmountsAdded []*big.Int
	SharesMinted *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}
type EFundingRemoved struct {
	//Signature: 8b4b2c8ebd04c47fc8bce136a85df9b93fcb1f47c8aa296457d4391519d190e7
	Funder                       common.Address
	AmountsRemoved               []*big.Int
	CollateralRemovedFromFeePool *big.Int
	SharesBurnt                  *big.Int
	Raw                          types.Log // Blockchain specific contextual infos
}
type EBuy struct {
	//Signature: 4f62630f51608fc8a7603a9391a5101e58bd7c276139366fc107dc3b67c3dcf8
	Buyer               common.Address
	InvestmentAmount    *big.Int
	FeeAmount           *big.Int
	OutcomeIndex        *big.Int
	OutcomeTokensBought *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}
type ESell struct {
	//Signature: adcf2a240ed9300d681d9a3f5382b6c1beed1b7e46643e0c7b42cbe6e2d766b4
	Seller            common.Address
	ReturnAmount      *big.Int
	FeeAmount         *big.Int
	OutcomeIndex      *big.Int
	OutcomeTokensSold *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}



//-------------- other types
type PolymarketProcStatus struct {
	LastIdProcessed			int64
	LastBlockNum			int64
}
type Pol_ConditionPreparation struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	SlotCount				int64
	Contract                string
	ConditionId				string
	OracleAddr				string
	QuestionId				string
	OutcomeSlotCount		string
}
type Pol_ConditionResolution struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	ConditionId				string
	OracleAddr				string
	QuestionId				string
	OutcomeSlotCount		string
	PayoutNumerators		string
}
type Pol_PositionSplit struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	ConditionId				string
	StakeHolderAddr			string
	CollateralToken			string
	ParentCollectionId		string
	Partition				string
	Amount					string
}
type Pol_PositionMerge struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	ConditionId				string
	StakeHolderAddr			string
	CollateralToken			string
	ParentCollectionId		string
	Partition				string
	Amount					string
}
type Pol_PayoutRedemption struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	Redeemer				string
	CollateralToken			string
	ParentCollectionId		string
	ConditionId				string
	IndexSets				string
	Payout					string
}
type Pol_URI struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	Value					string
	Id						string
}
type Pol_Market_API_Response_JSON struct {
	records					[]Pol_Market_API_Record_JSON

}
type Pol_Market_API_UseCase_JSON struct {
	Id				int64		`json:"id"`
	Title			string		`json:"title"`
	Content			string		`json:"content"`
}
type Pol_Market_API_SEO_MetaImage_Obj_JSON struct {
	Ext				string		`json:"ext"`
	Url				string		`json:"url"`
	Hash			string		`json:"hash"`
	Mime			string		`json:"mime"`
	Name			string		`json:"name"`
	Path			string		`json:"path"`
	Size			float64		`json:"size"`
	Width			int64		`json:"width"`
	Height			int64		`json:"height"`
}
type Pol_Market_API_SEO_MetaImage_JSON struct {
	Large				Pol_Market_API_SEO_MetaImage_Obj_JSON
	Small				Pol_Market_API_SEO_MetaImage_Obj_JSON
}
type Pol_Market_API_SEO_JSON struct {
	Id				int64		`json:"id"`
	MetaTitle		string		`json:"meta_title"`
	MetaDescription	string		`json:"meta_description"`
	MetaImage		Pol_Market_API_SEO_MetaImage_JSON `json:"meta_image"`

}
type Pol_Market_API_Record_JSON struct {
	// source: https://strapi-matic.poly.market/markets
	MarketId				int64		`json:"id"`
	Question				string		`json:"question"`
	ConditionId				string		`json:"conditionId"`
	Slug					string		`json:"slug"`
	TwitterCardImage		string		`json:"twitter_card_image"`
	ResolutionSource		string		`json:"resolution_source"`
	EndDate					string		`json:"end_date"`
	Category				string		`json:"category"`
	AmmType					string		`json:"amm_type"`
	Liquidity				string		`json:"liquidity"`
	SponsorName				string		`json:"sponsor_name"`
	SponsorImage			string		`json:"sponsor_image"`
	StartDate				string		`json:"start_date"`
	XAxisValue				int64		`json:"x_axis_value"`
	YAxisValue				int64		`json:"y_axis_value"`
	DenominationToken		string		`json:"denomination_token"`
	Fee						string		`json:"fee"`
	Image					string		`json:"image"`
	Icon					string		`json:"icon"`
	LowerBound				string		`json:"lower_bound"`
	UpperBound				string		`json:"upper_bound"`
	Description				string		`json:"description"`
	Tags					[]string	`json:"tags"`
	Outcomes				[]string	`json:"outcomes"`
	OutcomePrices			[]string	`json:"outcomePrices"`
	Volume					string		`json:"volume"`
	Active					bool		`json:"active"`
	MarketType				string		`json:"market_type"`
	FormatType				string		`json:"format_type"`
	LowerBoundDate			string		`json:"lower_bound_date"`
	UpperBoundDate			string		`json:"upper_bound_date"`
	Closed					bool		`json:"closed"`
	MarketMakerAddr			string		`json:"marketMakerAddress"`
	CreatedAtDate			string		`json:"created_at"`
	UpdatedAt				string		`json:"updated_at"`
	ClosedTimeDate			string		`json:"closed_time"`
	WideFormat				bool		`json:"wide_format"`
	New						bool		`json:"new"`
	SentDiscord				bool		`json:"sent_discord"`
	MailChimpTag			string		`json:"mailchimp_tag"`
	Featured				bool		`json:"featured"`
	SubmittedBy				string		`json:"submitted_by"`
	Subcategory				string		`json:"subcategory"`
	CategoryMailChimpTag	string		`json:"category_mailchimp_tag"`
	UseCases				[]Pol_Market_API_UseCase_JSON	`json:"use_cases"`
	SEO						Pol_Market_API_SEO_JSON 	`json:"seo"`
}
type Pol_Market_API_Record_Complementary struct {
	EndDateTs				int64
	StartDateTs				int64
	LowerBoundDateTs		int64
	UpperBoundDateTs		int64
	CreatedAtDateTs			int64
	UpdatedAtTs				int64
	ClosedTimeDateTs		int64
	MarketTypeCode			int64
}
type Pol_FundingAdded struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	Funder					string
	AmountsAdded			string
	SharesMinted			string
}
type Pol_FundingRemoved struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	Funder					string
	AmountsRemoved			string
	SharesBurnt				string
	CollateralRemoved		string
}
type Pol_Buy struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	Buyer					string
	InvestmentAmount		string
	FeeAmount				string
	OutcomeIdx				int64
	TokensBought			string
}
type Pol_Sell struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	Seller					string
	ReturnAmount			string
	FeeAmount				string
	OutcomeIdx				int64
	TokensSold				string
}
