// Pure decode tests (no Docker) for the inherited OpenZeppelin / Governor
// events and for the event fields the v3.1 audit found dropped
// (FundTransferFailed.errStr, RandomWalk NftUnstaked.actionCounter,
// EthDepositReceived.rewardAmountPerStakedNft).
package cosmicgame

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

var (
	unitAlice = ethcommon.HexToAddress("0x2100000000000000000000000000000000000021")
	unitBob   = ethcommon.HexToAddress("0x2200000000000000000000000000000000000022")
	unitCarol = ethcommon.HexToAddress("0x2300000000000000000000000000000000000023")
	unitMeta  = &store.EthereumEventLog{EvtID: 9, BlockNum: 8, TimeStamp: 7, TxID: 6}
)

// unitLog packs an event the way the EVM emits it: topic0 plus the indexed
// values as topics, the non-indexed values ABI-encoded in data.
func unitLog(t *testing.T, contractABI *abi.ABI, event string, from ethcommon.Address, indexed []ethcommon.Hash, nonIndexed ...any) *types.Log {
	t.Helper()
	ev, ok := contractABI.Events[event]
	if !ok {
		t.Fatalf("event %s not in ABI", event)
	}
	data, err := ev.Inputs.NonIndexed().Pack(nonIndexed...)
	if err != nil {
		t.Fatalf("packing %s: %v", event, err)
	}
	return &types.Log{Address: from, Topics: append([]ethcommon.Hash{ev.ID}, indexed...), Data: data}
}

func addrTopic(a ethcommon.Address) ethcommon.Hash { return ethcommon.BytesToHash(a.Bytes()) }
func bigTopic(v int64) ethcommon.Hash              { return ethcommon.BigToHash(big.NewInt(v)) }

func TestDecodeFundTransferFailedKeepsErrStr(t *testing.T) {
	h := newUnitHandlers(t)
	// The staking wallet's event (CosmicSignatureEvents library) has the
	// same shape the game ABI declares; build it from there and stamp the
	// wallet as the emitter.
	lg := unitLog(t, h.gameABI, "FundTransferFailed", h.c.StakingCST,
		[]ethcommon.Hash{addrTopic(unitCarol)}, "ETH transfer to charity failed.", big.NewInt(1234))
	got, err := h.decodeFundTransferFailed(lg, unitMeta)
	if err != nil {
		t.Fatal(err)
	}
	if got.Contract != h.c.StakingCST.String() || got.Destination != unitCarol.String() ||
		got.Amount != "1234" || got.ErrStr != "ETH transfer to charity failed." {
		t.Fatalf("decoded = %+v", got)
	}
}

func TestDecodeNftUnstakedRWalkKeepsActionCounter(t *testing.T) {
	h := newUnitHandlers(t)
	lg := unitLog(t, h.stakingRWalkABI, "NftUnstaked", h.c.StakingRWalk,
		[]ethcommon.Hash{bigTopic(5), bigTopic(10), addrTopic(unitCarol)}, big.NewInt(77), big.NewInt(3))
	got, err := h.decodeNftUnstakedRWalk(lg, unitMeta)
	if err != nil {
		t.Fatal(err)
	}
	if got.ActionId != 5 || got.NftId != 10 || got.StakerAddress != unitCarol.String() ||
		got.ActionCounter != 77 || got.NumStakedNfts != 3 {
		t.Fatalf("decoded = %+v", got)
	}
}

func TestDecodeStakingEthDepositKeepsRewardPerStakedNft(t *testing.T) {
	h := newUnitHandlers(t)
	lg := unitLog(t, h.stakingCSTABI, "EthDepositReceived", h.c.StakingCST,
		[]ethcommon.Hash{bigTopic(4)}, big.NewInt(12), big.NewInt(1001), big.NewInt(555), big.NewInt(3))
	got, err := h.decodeStakingEthDeposit(lg, unitMeta)
	if err != nil {
		t.Fatal(err)
	}
	if got.RoundNum != 4 || got.DepositId != 12 || got.Amount != "1001" || got.NumStakedNfts != 3 ||
		got.AmountPerStaker != "333" || got.Modulo != "2" || got.RewardPerStakedNft != "555" {
		t.Fatalf("decoded = %+v", got)
	}
}

func TestDecodeNftApprovals(t *testing.T) {
	h := newUnitHandlers(t)

	approval := unitLog(t, h.signatureABI, "Approval", h.c.Signature,
		[]ethcommon.Hash{addrTopic(unitAlice), addrTopic(unitBob), bigTopic(42)})
	got, err := h.decodeNftApproval(approval, unitMeta)
	if err != nil {
		t.Fatal(err)
	}
	if got.Owner != unitAlice.String() || got.Approved != unitBob.String() || got.TokenId != 42 || got.EvtId != 9 {
		t.Fatalf("Approval decoded = %+v", got)
	}
	if _, err := h.decodeNftApproval(&types.Log{Topics: approval.Topics[:3]}, unitMeta); err == nil {
		t.Fatal("Approval with three topics must fail, not panic")
	}

	forAll := unitLog(t, h.signatureABI, "ApprovalForAll", h.c.Signature,
		[]ethcommon.Hash{addrTopic(unitAlice), addrTopic(unitBob)}, true)
	all, err := h.decodeNftApprovalForAll(forAll, unitMeta)
	if err != nil {
		t.Fatal(err)
	}
	if all.Owner != unitAlice.String() || all.Operator != unitBob.String() || !all.Approved {
		t.Fatalf("ApprovalForAll decoded = %+v", all)
	}
}

func TestDecodeTokenEvents(t *testing.T) {
	h := newUnitHandlers(t)

	approval := unitLog(t, h.tokenABI, "Approval", h.c.Token,
		[]ethcommon.Hash{addrTopic(unitAlice), addrTopic(unitBob)}, big.NewInt(1e9))
	got, err := h.decodeTokenApproval(approval, unitMeta)
	if err != nil {
		t.Fatal(err)
	}
	if got.Owner != unitAlice.String() || got.Spender != unitBob.String() || got.Value != "1000000000" {
		t.Fatalf("Approval decoded = %+v", got)
	}

	delegated := unitLog(t, h.tokenABI, "DelegateChanged", h.c.Token,
		[]ethcommon.Hash{addrTopic(unitAlice), addrTopic(unitBob), addrTopic(unitCarol)})
	d, err := h.decodeDelegateChanged(delegated, unitMeta)
	if err != nil {
		t.Fatal(err)
	}
	if d.Delegator != unitAlice.String() || d.FromDelegate != unitBob.String() || d.ToDelegate != unitCarol.String() {
		t.Fatalf("DelegateChanged decoded = %+v", d)
	}

	votes := unitLog(t, h.tokenABI, "DelegateVotesChanged", h.c.Token,
		[]ethcommon.Hash{addrTopic(unitCarol)}, big.NewInt(10), big.NewInt(25))
	v, err := h.decodeDelegateVotesChanged(votes, unitMeta)
	if err != nil {
		t.Fatal(err)
	}
	if v.Delegate != unitCarol.String() || v.PreviousVotes != "10" || v.NewVotes != "25" {
		t.Fatalf("DelegateVotesChanged decoded = %+v", v)
	}

	domain := unitLog(t, h.tokenABI, "EIP712DomainChanged", h.c.Token, nil)
	e, err := h.decodeEIP712DomainChanged(domain, unitMeta)
	if err != nil {
		t.Fatal(err)
	}
	if e.Contract != h.c.Token.String() || e.EvtId != 9 {
		t.Fatalf("EIP712DomainChanged decoded = %+v", e)
	}
}

func TestDecodeDaoEvents(t *testing.T) {
	h := newUnitHandlers(t)
	pid, _ := new(big.Int).SetString("98765432109876543210987654321098765432109876543210", 10)

	created := unitLog(t, h.daoABI, "ProposalCreated", h.c.Dao, nil,
		pid, unitAlice,
		[]ethcommon.Address{unitBob, unitCarol},
		[]*big.Int{big.NewInt(0), big.NewInt(5)},
		[]string{"", "setFoo(uint256)"},
		[][]byte{{0x01, 0x02}, {}},
		big.NewInt(1000), big.NewInt(2000), "Proposal #1")
	c, err := h.decodeDaoProposalCreated(created, unitMeta)
	if err != nil {
		t.Fatal(err)
	}
	if c.ProposalId != pid.String() || c.Proposer != unitAlice.String() ||
		len(c.Targets) != 2 || c.Targets[1] != unitCarol.String() ||
		len(c.Values) != 2 || c.Values[1] != "5" ||
		len(c.Signatures) != 2 || c.Signatures[1] != "setFoo(uint256)" ||
		len(c.Calldatas) != 2 || string(c.Calldatas[0]) != "\x01\x02" ||
		c.VoteStart != 1000 || c.VoteEnd != 2000 || c.Description != "Proposal #1" {
		t.Fatalf("ProposalCreated decoded = %+v", c)
	}

	queued := unitLog(t, h.daoABI, "ProposalQueued", h.c.Dao, nil, pid, big.NewInt(1700000000))
	q, err := h.decodeDaoProposalQueued(queued, unitMeta)
	if err != nil {
		t.Fatal(err)
	}
	if q.State != cgmodel.DaoProposalQueued || q.ProposalId != pid.String() || q.EtaSeconds != 1700000000 {
		t.Fatalf("ProposalQueued decoded = %+v", q)
	}
	executed := unitLog(t, h.daoABI, "ProposalExecuted", h.c.Dao, nil, pid)
	x, err := h.decodeDaoProposalExecuted(executed, unitMeta)
	if err != nil {
		t.Fatal(err)
	}
	if x.State != cgmodel.DaoProposalExecuted || x.EtaSeconds != 0 {
		t.Fatalf("ProposalExecuted decoded = %+v", x)
	}
	canceled := unitLog(t, h.daoABI, "ProposalCanceled", h.c.Dao, nil, pid)
	k, err := h.decodeDaoProposalCanceled(canceled, unitMeta)
	if err != nil {
		t.Fatal(err)
	}
	if k.State != cgmodel.DaoProposalCanceled {
		t.Fatalf("ProposalCanceled decoded = %+v", k)
	}

	vote := unitLog(t, h.daoABI, "VoteCast", h.c.Dao,
		[]ethcommon.Hash{addrTopic(unitBob)}, pid, uint8(1), big.NewInt(300), "yes")
	v, err := h.decodeDaoVoteCast(vote, unitMeta)
	if err != nil {
		t.Fatal(err)
	}
	if v.Voter != unitBob.String() || v.ProposalId != pid.String() || v.Support != 1 ||
		v.Weight != "300" || v.Reason != "yes" || v.Params != nil {
		t.Fatalf("VoteCast decoded = %+v", v)
	}
	voteParams := unitLog(t, h.daoABI, "VoteCastWithParams", h.c.Dao,
		[]ethcommon.Hash{addrTopic(unitBob)}, pid, uint8(2), big.NewInt(1), "", []byte{0xaa})
	vp, err := h.decodeDaoVoteCastWithParams(voteParams, unitMeta)
	if err != nil {
		t.Fatal(err)
	}
	if vp.Support != 2 || string(vp.Params) != "\xaa" {
		t.Fatalf("VoteCastWithParams decoded = %+v", vp)
	}

	for _, tc := range []struct {
		event   string
		setting int64
	}{
		{"ProposalThresholdSet", cgmodel.DaoSettingProposalThreshold},
		{"VotingDelaySet", cgmodel.DaoSettingVotingDelay},
		{"VotingPeriodSet", cgmodel.DaoSettingVotingPeriod},
		{"QuorumNumeratorUpdated", cgmodel.DaoSettingQuorumNumerator},
	} {
		lg := unitLog(t, h.daoABI, tc.event, h.c.Dao, nil, big.NewInt(3), big.NewInt(4))
		s, err := h.decodeDaoSettingChanged(tc.event, tc.setting)(lg, unitMeta)
		if err != nil {
			t.Fatalf("%s: %v", tc.event, err)
		}
		if s.Setting != tc.setting || s.OldValue != "3" || s.NewValue != "4" {
			t.Fatalf("%s decoded = %+v", tc.event, s)
		}
	}
}

func TestDaoDecodersRejectOutOfRange(t *testing.T) {
	h := newUnitHandlers(t)
	huge := new(big.Int).Lsh(big.NewInt(1), 70)
	created := unitLog(t, h.daoABI, "ProposalCreated", h.c.Dao, nil,
		big.NewInt(1), unitAlice, []ethcommon.Address{}, []*big.Int{}, []string{}, [][]byte{},
		huge, big.NewInt(2), "")
	if _, err := h.decodeDaoProposalCreated(created, unitMeta); err == nil {
		t.Fatal("voteStart beyond int64 must fail the decode")
	}
	queued := unitLog(t, h.daoABI, "ProposalQueued", h.c.Dao, nil, big.NewInt(1), huge)
	if _, err := h.decodeDaoProposalQueued(queued, unitMeta); err == nil {
		t.Fatal("etaSeconds beyond int64 must fail the decode")
	}
}
