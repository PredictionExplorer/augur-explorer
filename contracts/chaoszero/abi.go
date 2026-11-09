// Package chaoszero carries the event ABI of the deployed ChaosZero
// contracts: the eight-market IndependentOutcomeMarket (an ERC-1155) and the
// OutcomeSettlementVault its constructor deploys. The contracts are
// immutable singletons with no owner and no upgrade path, so these
// hand-written, events-only ABI fragments are pinned to the Arbitrum One
// deployment of 21 October 2026 (market
// 0x2BC4759A2b44e1197EAdc4E035Eb8c52299658d2). The indexer package's
// registry test asserts every topic-hash constant against these fragments.
package chaoszero

// MarketABI describes the IndependentOutcomeMarket events the ETL
// dispatches: the eight market events plus the ERC-1155 transfer/approval
// events it inherits from OpenZeppelin.
const MarketABI = `[
{"type":"event","name":"RoundInitialized","anonymous":false,"inputs":[
	{"name":"roundId","type":"uint256","indexed":true}]},
{"type":"event","name":"RoundFunded","anonymous":false,"inputs":[
	{"name":"roundId","type":"uint256","indexed":true},
	{"name":"caller","type":"address","indexed":true},
	{"name":"donated","type":"uint256","indexed":false},
	{"name":"transfersUsed","type":"uint256","indexed":false},
	{"name":"carryUsed","type":"uint256","indexed":false},
	{"name":"allocations","type":"uint256[8]","indexed":false}]},
{"type":"event","name":"Trade","anonymous":false,"inputs":[
	{"name":"roundId","type":"uint256","indexed":true},
	{"name":"user","type":"address","indexed":true},
	{"name":"question","type":"uint8","indexed":true},
	{"name":"yes","type":"bool","indexed":false},
	{"name":"buy","type":"bool","indexed":false},
	{"name":"cash","type":"uint256","indexed":false},
	{"name":"units","type":"uint256","indexed":false}]},
{"type":"event","name":"PositionSplit","anonymous":false,"inputs":[
	{"name":"roundId","type":"uint256","indexed":true},
	{"name":"user","type":"address","indexed":true},
	{"name":"question","type":"uint8","indexed":true},
	{"name":"amount","type":"uint256","indexed":false}]},
{"type":"event","name":"PositionsMerged","anonymous":false,"inputs":[
	{"name":"roundId","type":"uint256","indexed":true},
	{"name":"user","type":"address","indexed":true},
	{"name":"amounts","type":"uint256[8]","indexed":false},
	{"name":"collateralOut","type":"uint256","indexed":false}]},
{"type":"event","name":"RoundSettled","anonymous":false,"inputs":[
	{"name":"roundId","type":"uint256","indexed":true},
	{"name":"values","type":"uint8","indexed":false},
	{"name":"escrowed","type":"uint256[8]","indexed":false},
	{"name":"carried","type":"uint256","indexed":false}]},
{"type":"event","name":"Claimed","anonymous":false,"inputs":[
	{"name":"roundId","type":"uint256","indexed":true},
	{"name":"user","type":"address","indexed":true},
	{"name":"amounts","type":"uint256[8]","indexed":false},
	{"name":"total","type":"uint256","indexed":false}]},
{"type":"event","name":"AnswersObserved","anonymous":false,"inputs":[
	{"name":"roundId","type":"uint256","indexed":true},
	{"name":"knownMask","type":"uint8","indexed":false},
	{"name":"values","type":"uint8","indexed":false}]},
{"type":"event","name":"TransferSingle","anonymous":false,"inputs":[
	{"name":"operator","type":"address","indexed":true},
	{"name":"from","type":"address","indexed":true},
	{"name":"to","type":"address","indexed":true},
	{"name":"id","type":"uint256","indexed":false},
	{"name":"value","type":"uint256","indexed":false}]},
{"type":"event","name":"TransferBatch","anonymous":false,"inputs":[
	{"name":"operator","type":"address","indexed":true},
	{"name":"from","type":"address","indexed":true},
	{"name":"to","type":"address","indexed":true},
	{"name":"ids","type":"uint256[]","indexed":false},
	{"name":"values","type":"uint256[]","indexed":false}]},
{"type":"event","name":"ApprovalForAll","anonymous":false,"inputs":[
	{"name":"account","type":"address","indexed":true},
	{"name":"operator","type":"address","indexed":true},
	{"name":"approved","type":"bool","indexed":false}]}
]`

// VaultABI describes the OutcomeSettlementVault events.
const VaultABI = `[
{"type":"event","name":"CycleEscrowed","anonymous":false,"inputs":[
	{"name":"roundId","type":"uint256","indexed":true},
	{"name":"amounts","type":"uint256[8]","indexed":false}]},
{"type":"event","name":"Paid","anonymous":false,"inputs":[
	{"name":"roundId","type":"uint256","indexed":true},
	{"name":"question","type":"uint8","indexed":true},
	{"name":"recipient","type":"address","indexed":true},
	{"name":"amount","type":"uint256","indexed":false}]}
]`
