// V4 (in development) events-only ABI fragments.
//
// The Cosmic Signature V4 upgrade adds four events to the game proxy and a
// new non-upgradeable CosmicSignatureArtAuction contract. V4 is not deployed
// yet and its interfaces may still change, so — like the ChaosZero package —
// these hand-written, events-only fragments carry exactly what the ETL
// dispatches, extracted from the Cosmic-Signature repo's `V4` branch
// artifacts (solc 0.8.38). When V4 stabilizes and ships, fold the contracts
// into buildjson/combined.json and bindings.gen.go, and retire this file.
//
// The indexer package's topic tests pin every V4 topic-hash constant against
// these fragments, so an upstream event-signature change surfaces as a test
// failure here rather than as silently dropped events.
package cosmicgame

// GameV4GettersABI carries the V4-only view functions the backend probes:
// gamever detects a V4 deployment by artAuctionOpeningBid() answering, and
// the contractsync drift audit compares the three getters against indexed
// history.
const GameV4GettersABI = `[
{"type":"function","name":"artAuctionOpeningBid","stateMutability":"view","inputs":[],"outputs":[{"name":"","type":"uint256"}]},
{"type":"function","name":"artAuction","stateMutability":"view","inputs":[],"outputs":[{"name":"","type":"address"}]},
{"type":"function","name":"chaosZeroMarket","stateMutability":"view","inputs":[],"outputs":[{"name":"","type":"address"}]}
]`

// GameV4ABI describes the events CosmicSignatureGameV4 adds on top of the
// V3 game ABI: the art-auction wiring announcements and the per-cycle
// ChaosZero market funding record.
const GameV4ABI = `[
{"type":"event","name":"ArtAuctionConfigured","anonymous":false,"inputs":[
	{"name":"artAuctionAddress","type":"address","indexed":true}]},
{"type":"event","name":"ArtAuctionNftsMinted","anonymous":false,"inputs":[
	{"name":"roundNum","type":"uint256","indexed":true},
	{"name":"artAuctionAddress","type":"address","indexed":true},
	{"name":"firstNftId","type":"uint256","indexed":true},
	{"name":"numNfts","type":"uint256","indexed":false}]},
{"type":"event","name":"ArtAuctionOpeningBidChanged","anonymous":false,"inputs":[
	{"name":"openingBid","type":"uint256","indexed":false}]},
{"type":"event","name":"ChaosZeroCycleFunded","anonymous":false,"inputs":[
	{"name":"market","type":"address","indexed":true},
	{"name":"cycle","type":"uint256","indexed":true},
	{"name":"amount","type":"uint256","indexed":false},
	{"name":"cumulativeAmount","type":"uint256","indexed":false}]}
]`

// ArtAuctionABI describes every event of the non-upgradeable
// CosmicSignatureArtAuction contract (fully funded English auctions over the
// eight extra Cosmic Signature NFTs each V4 round mints into its custody).
const ArtAuctionABI = `[
{"type":"event","name":"AuctionBatchRegistered","anonymous":false,"inputs":[
	{"name":"roundNum","type":"uint256","indexed":true},
	{"name":"firstNftId","type":"uint256","indexed":true},
	{"name":"count","type":"uint256","indexed":false},
	{"name":"eligibleAt","type":"uint256","indexed":false},
	{"name":"openingBid","type":"uint256","indexed":false}]},
{"type":"event","name":"AuctionBidPlaced","anonymous":false,"inputs":[
	{"name":"nftId","type":"uint256","indexed":true},
	{"name":"roundNum","type":"uint256","indexed":true},
	{"name":"bidder","type":"address","indexed":true},
	{"name":"amount","type":"uint256","indexed":false},
	{"name":"ethPaid","type":"uint256","indexed":false},
	{"name":"endAt","type":"uint256","indexed":false}]},
{"type":"event","name":"AuctionNftDelivered","anonymous":false,"inputs":[
	{"name":"nftId","type":"uint256","indexed":true},
	{"name":"winner","type":"address","indexed":true},
	{"name":"recipient","type":"address","indexed":true}]},
{"type":"event","name":"AuctionProceedsForwarded","anonymous":false,"inputs":[
	{"name":"nftId","type":"uint256","indexed":true},
	{"name":"roundNum","type":"uint256","indexed":true},
	{"name":"amount","type":"uint256","indexed":false}]},
{"type":"event","name":"AuctionRefundCredited","anonymous":false,"inputs":[
	{"name":"bidder","type":"address","indexed":true},
	{"name":"nftId","type":"uint256","indexed":true},
	{"name":"amount","type":"uint256","indexed":false}]},
{"type":"event","name":"AuctionRefundWithdrawn","anonymous":false,"inputs":[
	{"name":"bidder","type":"address","indexed":true},
	{"name":"recipient","type":"address","indexed":true},
	{"name":"amount","type":"uint256","indexed":false}]},
{"type":"event","name":"AuctionSettled","anonymous":false,"inputs":[
	{"name":"nftId","type":"uint256","indexed":true},
	{"name":"roundNum","type":"uint256","indexed":true},
	{"name":"winner","type":"address","indexed":true},
	{"name":"amount","type":"uint256","indexed":false}]},
{"type":"event","name":"AuctionSurplusForwarded","anonymous":false,"inputs":[
	{"name":"amount","type":"uint256","indexed":false}]}
]`
