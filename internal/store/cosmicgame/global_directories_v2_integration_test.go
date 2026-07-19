//go:build integration

package cosmicgame

import (
	"cmp"
	"context"
	"math/big"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"testing"
	"time"
)

// walkGlobalTokenPages exhausts the descending global token directory with
// the given filter and page size, verifying page bounds and strict
// descending order on the way.
func walkGlobalTokenPages(
	t *testing.T,
	r *Repo,
	filter GlobalTokenFilter,
	pageSize int,
) []GlobalTokenRecord {
	t.Helper()
	ctx := context.Background()
	var collected []GlobalTokenRecord
	var after *GlobalTokenPageCursor
	for {
		page, hasMore, err := r.CosmicSignatureTokensGlobalPage(ctx, filter, after, pageSize)
		if err != nil {
			t.Fatalf("page walk: %v", err)
		}
		if len(page) > pageSize {
			t.Fatalf("page length = %d, limit %d", len(page), pageSize)
		}
		for i := range page {
			if len(collected) > 0 &&
				page[i].TokenID >= collected[len(collected)-1].TokenID {
				t.Fatalf("unordered token IDs at %d", page[i].TokenID)
			}
			collected = append(collected, page[i])
		}
		if !hasMore {
			return collected
		}
		if len(page) == 0 {
			t.Fatal("hasMore without a cursor row")
		}
		after = &GlobalTokenPageCursor{TokenID: page[len(page)-1].TokenID}
	}
}

func TestCosmicSignatureTokensGlobalPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	// The directory must agree with the legacy global list. The legacy
	// query joins every stake cycle, so its rows are de-duplicated before
	// the comparison (the very defect the v2 directory fixes).
	legacy, err := r.CosmicSignatureTokens(ctx, 0, 1000)
	if err != nil {
		t.Fatal(err)
	}
	seen := map[int64]bool{}
	var legacyIDs []int64
	for i := range legacy {
		if !seen[legacy[i].TokenId] {
			seen[legacy[i].TokenId] = true
			legacyIDs = append(legacyIDs, legacy[i].TokenId)
		}
	}
	slices.SortFunc(legacyIDs, func(a, b int64) int { return cmp.Compare(b, a) })

	all := walkGlobalTokenPages(t, r, GlobalTokenFilter{}, 1)
	gotIDs := make([]int64, len(all))
	for i := range all {
		gotIDs[i] = all[i].TokenID
	}
	if !reflect.DeepEqual(gotIDs, legacyIDs) {
		t.Fatalf("paged tokens = %v, legacy = %v", gotIDs, legacyIDs)
	}

	// Every row resolves exactly one mint source and both owner columns.
	for i := range all {
		if all[i].Seed == "" || all[i].WinnerAddr == "" || all[i].CurOwnerAddr == "" {
			t.Fatalf("incomplete row %+v", all[i])
		}
	}

	// Token 1 belongs to alice, is named Genesis, and its stake/unstake
	// cycle has ended so membership is gone; token 5 sits in the staking
	// wallet; token 6 is the chrono-warrior NFT v1 mislabels.
	byID := map[int64]GlobalTokenRecord{}
	for i := range all {
		byID[all[i].TokenID] = all[i]
	}
	if token := byID[1]; token.TokenName != "Genesis" || token.Staked ||
		token.MintSource != MintSourceMainPrize {
		t.Fatalf("token 1 = %+v", token)
	}
	if token := byID[5]; !token.Staked || token.MintSource != MintSourceEnduranceChampion {
		t.Fatalf("token 5 = %+v", token)
	}
	if token := byID[6]; token.MintSource != MintSourceChronoWarriorPrize {
		t.Fatalf("token 6 = %+v", token)
	}
	// Token 2 was minted to dave and transferred to bob: the directory
	// keeps the mint provenance while the current owner follows.
	if token := byID[2]; token.MintSource != MintSourceBidderRaffle ||
		strings.EqualFold(token.WinnerAddr, token.CurOwnerAddr) {
		t.Fatalf("token 2 = %+v", token)
	}
}

func TestCosmicSignatureTokensGlobalPageFilters(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	// named=true must agree with the legacy named-token view.
	legacy, err := r.NamedTokens(ctx)
	if err != nil {
		t.Fatal(err)
	}
	wantNamed := make([]int64, len(legacy))
	for i := range legacy {
		wantNamed[i] = legacy[i].TokenId
	}
	slices.SortFunc(wantNamed, func(a, b int64) int { return cmp.Compare(b, a) })
	named := walkGlobalTokenPages(t, r, GlobalTokenFilter{NamedOnly: true}, 1)
	gotNamed := make([]int64, len(named))
	for i := range named {
		gotNamed[i] = named[i].TokenID
		if named[i].TokenName == "" {
			t.Fatalf("named filter returned unnamed token %d", named[i].TokenID)
		}
	}
	if !reflect.DeepEqual(gotNamed, wantNamed) {
		t.Fatalf("named tokens = %v, legacy = %v", gotNamed, wantNamed)
	}
	if len(gotNamed) == 0 {
		t.Fatal("fixture has no named tokens; the filter test is vacuous")
	}

	// The search filter must agree with the legacy substring search,
	// case-insensitively.
	searched, err := r.SearchTokensByName(ctx, "gene")
	if err != nil {
		t.Fatal(err)
	}
	wantFound := make([]int64, len(searched))
	for i := range searched {
		wantFound[i] = searched[i].TokenId
	}
	slices.SortFunc(wantFound, func(a, b int64) int { return cmp.Compare(b, a) })
	found := walkGlobalTokenPages(t, r, GlobalTokenFilter{NameContains: "gene"}, 1)
	gotFound := make([]int64, len(found))
	for i := range found {
		gotFound[i] = found[i].TokenID
	}
	if !reflect.DeepEqual(gotFound, wantFound) || len(gotFound) == 0 {
		t.Fatalf("searched tokens = %v, legacy = %v", gotFound, wantFound)
	}

	// ILIKE wildcards in the term are escaped: a bare % matches only
	// names containing a literal percent sign, so nothing here — the
	// legacy query would have returned every named token.
	wild, hasMore, err := r.CosmicSignatureTokensGlobalPage(
		ctx, GlobalTokenFilter{NameContains: "%"}, nil, 50)
	if err != nil || hasMore || len(wild) != 0 {
		t.Fatalf("wildcard search = %d rows, more=%v, err=%v", len(wild), hasMore, err)
	}
	underscore, hasMore, err := r.CosmicSignatureTokensGlobalPage(
		ctx, GlobalTokenFilter{NameContains: "_"}, nil, 50)
	if err != nil || hasMore || len(underscore) != 0 {
		t.Fatalf("underscore search = %d rows, more=%v, err=%v", len(underscore), hasMore, err)
	}

	// Contradictory filters and invalid limits fail before the query.
	if _, _, err := r.CosmicSignatureTokensGlobalPage(
		ctx, GlobalTokenFilter{NamedOnly: true, NameContains: "x"}, nil, 1); err == nil {
		t.Fatal("contradictory filter accepted")
	}
	if _, _, err := r.CosmicSignatureTokensGlobalPage(ctx, GlobalTokenFilter{}, nil, 0); err == nil {
		t.Fatal("zero limit accepted")
	}
	if _, _, err := r.CosmicSignatureTokensGlobalPage(
		ctx, GlobalTokenFilter{}, &GlobalTokenPageCursor{TokenID: -1}, 1); err == nil {
		t.Fatal("negative cursor accepted")
	}

	// The terminal cursor yields an empty page.
	exhausted, hasMore, err := r.CosmicSignatureTokensGlobalPage(
		ctx, GlobalTokenFilter{}, &GlobalTokenPageCursor{TokenID: 0}, 50)
	if err != nil || hasMore || len(exhausted) != 0 {
		t.Fatalf("exhausted page = %d rows, more=%v, err=%v", len(exhausted), hasMore, err)
	}
}

func TestCosmicSignatureTokenDetailV2(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	// Token 5 sits in the staking wallet: the detail carries the locking
	// stake action and agrees with the legacy info view.
	staked, err := r.CosmicSignatureTokenDetailV2(ctx, 5)
	if err != nil {
		t.Fatal(err)
	}
	if !staked.Staked || staked.CurrentStake == nil {
		t.Fatalf("token 5 detail = %+v", staked)
	}
	legacy, err := r.CosmicSignatureTokenInfo(ctx, 5)
	if err != nil {
		t.Fatal(err)
	}
	if staked.CurrentStake.StakeActionID != legacy.StakeActionId ||
		!strings.EqualFold(staked.CurrentStake.StakerAddr, legacy.StakedOwnerAddr) {
		t.Fatalf("stake detail = %+v, legacy action %d staker %s",
			staked.CurrentStake, legacy.StakeActionId, legacy.StakedOwnerAddr)
	}
	if _, err := time.Parse(time.RFC3339Nano, staked.CurrentStake.StakedAtText); err != nil {
		t.Fatalf("stake timestamp %q: %v", staked.CurrentStake.StakedAtText, err)
	}

	// Token 1 finished its stake cycle: no live stake, name preserved.
	unstaked, err := r.CosmicSignatureTokenDetailV2(ctx, 1)
	if err != nil || unstaked.Staked || unstaked.CurrentStake != nil ||
		unstaked.TokenName != "Genesis" || unstaked.MintSource != MintSourceMainPrize {
		t.Fatalf("token 1 detail = %+v, err=%v", unstaked, err)
	}

	// Token 6 is the chrono-warrior mint the legacy view mislabeled.
	chrono, err := r.CosmicSignatureTokenDetailV2(ctx, 6)
	if err != nil || chrono.MintSource != MintSourceChronoWarriorPrize {
		t.Fatalf("token 6 detail = %+v, err=%v", chrono, err)
	}

	if _, err := r.CosmicSignatureTokenDetailV2(ctx, 424242); err == nil {
		t.Fatal("unknown token returned a detail")
	}
	if _, err := r.CosmicSignatureTokenDetailV2(ctx, -1); err == nil {
		t.Fatal("negative token id accepted")
	}
}

func TestTokenNameHistoryPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	// Token 1 carries the fixture renames; the page walk must agree with
	// the legacy per-token history.
	legacy, err := r.TokenNameHistory(ctx, 1)
	if err != nil {
		t.Fatal(err)
	}
	if len(legacy) == 0 {
		t.Fatal("fixture has no renames for token 1; the test is vacuous")
	}
	wantIDs := make([]int64, len(legacy))
	for i := range legacy {
		wantIDs[i] = legacy[i].Tx.EvtLogId
	}

	var gotIDs []int64
	var after *TokenEventPageCursor
	for {
		page, hasMore, err := r.TokenNameHistoryPage(ctx, 1, after, 1)
		if err != nil {
			t.Fatalf("page walk: %v", err)
		}
		for i := range page {
			if page[i].TokenID != 1 {
				t.Fatalf("foreign token in page: %+v", page[i])
			}
			if len(gotIDs) > 0 && page[i].Tx.EvtLogId >= gotIDs[len(gotIDs)-1] {
				t.Fatalf("unordered event ids: %v then %d", gotIDs, page[i].Tx.EvtLogId)
			}
			if page[i].ChangedBy == "" {
				t.Fatalf("rename without an author: %+v", page[i])
			}
			gotIDs = append(gotIDs, page[i].Tx.EvtLogId)
		}
		if !hasMore {
			break
		}
		after = &TokenEventPageCursor{EventLogID: page[len(page)-1].Tx.EvtLogId}
	}
	if !reflect.DeepEqual(gotIDs, wantIDs) {
		t.Fatalf("paged renames = %v, legacy = %v", gotIDs, wantIDs)
	}

	// A minted but never renamed token gets an empty page.
	if legacy, err := r.TokenNameHistory(ctx, 2); err != nil || len(legacy) != 0 {
		t.Fatalf("fixture assumption changed: token 2 has %d renames, err=%v", len(legacy), err)
	}
	empty, hasMore, err := r.TokenNameHistoryPage(ctx, 2, nil, 50)
	if err != nil || hasMore || len(empty) != 0 {
		t.Fatalf("unnamed token page = %d rows, more=%v, err=%v", len(empty), hasMore, err)
	}

	if _, _, err := r.TokenNameHistoryPage(ctx, -1, nil, 1); err == nil {
		t.Fatal("negative token id accepted")
	}
}

func TestTokenTransfersPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	// Token 2 was minted to dave and transferred to bob; the page walk
	// must agree with the legacy ownership history.
	legacy, err := r.TokenOwnershipTransfers(ctx, 2, 0, 1000)
	if err != nil {
		t.Fatal(err)
	}
	if len(legacy) < 2 {
		t.Fatalf("fixture assumption changed: token 2 has %d transfers", len(legacy))
	}
	wantIDs := make([]int64, len(legacy))
	for i := range legacy {
		wantIDs[i] = legacy[i].Tx.EvtLogId
	}

	var gotIDs []int64
	var sawMint bool
	var after *TokenEventPageCursor
	for {
		page, hasMore, err := r.TokenTransfersPage(ctx, 2, after, 1)
		if err != nil {
			t.Fatalf("page walk: %v", err)
		}
		for i := range page {
			if page[i].TokenID != 2 {
				t.Fatalf("foreign token in page: %+v", page[i])
			}
			if page[i].TransferType == 1 {
				sawMint = true
			}
			gotIDs = append(gotIDs, page[i].Tx.EvtLogId)
		}
		if !hasMore {
			break
		}
		after = &TokenEventPageCursor{EventLogID: page[len(page)-1].Tx.EvtLogId}
	}
	if !reflect.DeepEqual(gotIDs, wantIDs) {
		t.Fatalf("paged transfers = %v, legacy = %v", gotIDs, wantIDs)
	}
	if !sawMint {
		t.Fatal("the mint row is missing from the ownership history")
	}

	if _, _, err := r.TokenTransfersPage(ctx, 2, &TokenEventPageCursor{EventLogID: 0}, 1); err == nil {
		t.Fatal("invalid cursor accepted")
	}
}

func TestCosmicSignatureHoldersPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	// The paged directory must carry the same multiset as the legacy
	// distribution; v2 adds the deterministic address-ID tie-break.
	legacy, err := r.CosmicSignatureTokenDistribution(ctx)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]int64{}
	for i := range legacy {
		want[strings.ToLower(legacy[i].OwnerAddr)] = legacy[i].NumTokens
	}

	var collected []CosmicSignatureHolderRecord
	var after *ParticipantPageCursor
	for {
		page, hasMore, err := r.CosmicSignatureHoldersPage(ctx, after, 1)
		if err != nil {
			t.Fatalf("page walk: %v", err)
		}
		for i := range page {
			if len(collected) > 0 {
				previous := collected[len(collected)-1]
				if page[i].TokenCount > previous.TokenCount ||
					(page[i].TokenCount == previous.TokenCount &&
						page[i].OwnerAid <= previous.OwnerAid) {
					t.Fatalf("unordered holders: %+v then %+v", previous, page[i])
				}
			}
			collected = append(collected, page[i])
		}
		if !hasMore {
			break
		}
		last := page[len(page)-1]
		after = &ParticipantPageCursor{
			Kind:      ParticipantCsTokenHolders,
			SortValue: strconv.FormatInt(last.TokenCount, 10),
			AddressID: last.OwnerAid,
		}
	}
	got := map[string]int64{}
	for i := range collected {
		got[strings.ToLower(collected[i].Address)] = collected[i].TokenCount
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("holders = %v, legacy = %v", got, want)
	}
	if len(collected) < 2 {
		t.Fatalf("fixture has %d holders; the ordering test is vacuous", len(collected))
	}

	if _, _, err := r.CosmicSignatureHoldersPage(ctx, &ParticipantPageCursor{
		Kind: ParticipantBidders, SortValue: "1", AddressID: 1,
	}, 1); err == nil {
		t.Fatal("cross-directory cursor accepted")
	}
}

func TestCosmicTokenHoldersPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	legacy, err := r.CosmicTokenHolders(ctx)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{}
	for i := range legacy {
		balance, ok := new(big.Int).SetString(legacy[i].Balance, 10)
		if !ok {
			t.Fatalf("legacy balance %q", legacy[i].Balance)
		}
		if balance.Sign() > 0 {
			want[strings.ToLower(legacy[i].OwnerAddr)] = balance.String()
		}
	}
	if len(want) < 2 {
		t.Fatalf("fixture has %d positive balances; the test is vacuous", len(want))
	}

	var collected []CosmicTokenHolderRecord
	var after *ParticipantPageCursor
	for {
		page, hasMore, err := r.CosmicTokenHoldersPage(ctx, after, 1)
		if err != nil {
			t.Fatalf("page walk: %v", err)
		}
		for i := range page {
			if len(collected) > 0 {
				previous := collected[len(collected)-1]
				previousBalance, _ := new(big.Int).SetString(previous.BalanceWei, 10)
				currentBalance, _ := new(big.Int).SetString(page[i].BalanceWei, 10)
				comparison := currentBalance.Cmp(previousBalance)
				if comparison > 0 || (comparison == 0 && page[i].OwnerAid <= previous.OwnerAid) {
					t.Fatalf("unordered balances: %+v then %+v", previous, page[i])
				}
			}
			collected = append(collected, page[i])
		}
		if !hasMore {
			break
		}
		last := page[len(page)-1]
		after = &ParticipantPageCursor{
			Kind:      ParticipantCosmicTokenHolders,
			SortValue: last.BalanceWei,
			AddressID: last.OwnerAid,
		}
	}
	got := map[string]string{}
	for i := range collected {
		got[strings.ToLower(collected[i].Address)] = collected[i].BalanceWei
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("holders = %v, legacy = %v", got, want)
	}
}

func TestCosmicTokenStatisticsV2(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	record, err := r.CosmicTokenStatisticsV2(ctx)
	if err != nil {
		t.Fatal(err)
	}
	legacy, err := r.CosmicTokenStatistics(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// The exact aggregates must agree with the legacy statistics where
	// both exist; the sums drop the float twins.
	equalDecimal := func(name, got, want string) {
		gotInt, ok1 := new(big.Int).SetString(got, 10)
		wantInt, ok2 := new(big.Int).SetString(want, 10)
		if !ok1 || !ok2 || gotInt.Cmp(wantInt) != 0 {
			t.Errorf("%s = %s, legacy %s", name, got, want)
		}
	}
	equalDecimal("total supply", record.TotalSupplyWei, legacy.TotalSupply)
	equalDecimal("bidding rewards", record.BiddingRewardsWei, legacy.EarnedFromBidding)
	equalDecimal("marketing rewards", record.MarketingRewardsWei, legacy.DistributedToMarketers)
	equalDecimal("main prizes", record.MainPrizesWei, legacy.GivenAsMainPrizes)
	equalDecimal("raffle prizes", record.RafflePrizesWei, legacy.GivenAsRafflePrizes)
	equalDecimal("chrono prizes", record.ChronoWarriorPrizesWei, legacy.GivenAsChronoWarriorPrizes)
	if record.HolderCount != legacy.TotalHolders ||
		record.MintCount != legacy.TotalMints ||
		record.BurnCount != legacy.TotalBurns {
		t.Errorf("counters = %d/%d/%d, legacy %d/%d/%d",
			record.HolderCount, record.MintCount, record.BurnCount,
			legacy.TotalHolders, legacy.TotalMints, legacy.TotalBurns)
	}
	// v1 counts otype=0 rows as "transfers"; v2 counts all rows.
	if record.TransferCount != legacy.TotalTransfers+legacy.TotalMints+legacy.TotalBurns {
		t.Errorf("transfer count = %d, legacy total %d",
			record.TransferCount, legacy.TotalTransfers+legacy.TotalMints+legacy.TotalBurns)
	}

	// The seven sources sum to the total and the net closes the loop —
	// this is the invariant the mapper enforces on every response.
	sum := new(big.Int)
	for _, source := range []string{
		record.BiddingRewardsWei,
		record.MainPrizesWei,
		record.RafflePrizesWei,
		record.ChronoWarriorPrizesWei,
		record.EnduranceChampionPrizesWei,
		record.LastCstBidderPrizesWei,
		record.MarketingRewardsWei,
	} {
		part, ok := new(big.Int).SetString(source, 10)
		if !ok {
			t.Fatalf("source %q", source)
		}
		sum.Add(sum, part)
	}
	if sum.String() != record.TotalEarnedWei {
		t.Fatalf("sources sum to %s, total %s", sum.String(), record.TotalEarnedWei)
	}
	consumed, _ := new(big.Int).SetString(record.ConsumedInBidsWei, 10)
	if new(big.Int).Sub(sum, consumed).String() != record.NetWei {
		t.Fatalf("net %s does not close earned %s - consumed %s",
			record.NetWei, sum.String(), record.ConsumedInBidsWei)
	}

	// The endurance and last-CST sources v1 omitted are present and
	// nonzero in the fixture set.
	for name, source := range map[string]string{
		"endurance": record.EnduranceChampionPrizesWei,
		"last CST":  record.LastCstBidderPrizesWei,
	} {
		value, ok := new(big.Int).SetString(source, 10)
		if !ok || value.Sign() <= 0 {
			t.Errorf("%s source = %q, want a positive amount", name, source)
		}
	}

	// The embedded top-holder list agrees with the legacy one.
	if len(record.TopHolders) != len(legacy.TopHolders) {
		t.Fatalf("top holders = %d, legacy %d", len(record.TopHolders), len(legacy.TopHolders))
	}
	for i := range record.TopHolders {
		if !strings.EqualFold(record.TopHolders[i].Address, legacy.TopHolders[i].OwnerAddr) {
			t.Errorf("top holder %d = %s, legacy %s",
				i, record.TopHolders[i].Address, legacy.TopHolders[i].OwnerAddr)
		}
		equalDecimal("top holder balance", record.TopHolders[i].BalanceWei, legacy.TopHolders[i].Balance)
	}
}

func TestCosmicTokenSupplyByBidPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	legacy, err := r.CosmicTokenSupplyHistoryByBid(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(legacy) < 4 {
		t.Fatalf("fixture has %d bids; the paging test is vacuous", len(legacy))
	}

	// Walking with a small page size must reproduce the legacy ledger
	// exactly — order, per-bid amounts and running totals included. This
	// proves the base-aggregate-plus-window pagination equals the naive
	// full-window form.
	var collected []SupplyChangeRecord
	var after *SupplyChangePageCursor
	for {
		page, hasMore, err := r.CosmicTokenSupplyByBidPage(ctx, after, 3)
		if err != nil {
			t.Fatalf("page walk: %v", err)
		}
		if len(page) > 3 {
			t.Fatalf("page length = %d", len(page))
		}
		collected = append(collected, page...)
		if !hasMore {
			break
		}
		if len(page) == 0 {
			t.Fatal("hasMore without a cursor row")
		}
		after = &SupplyChangePageCursor{EventLogID: page[len(page)-1].Tx.EvtLogId}
	}
	if len(collected) != len(legacy) {
		t.Fatalf("collected %d rows, legacy %d", len(collected), len(legacy))
	}
	equalDecimal := func(name, got, want string, row int) {
		gotInt, ok1 := new(big.Int).SetString(got, 10)
		wantInt, ok2 := new(big.Int).SetString(want, 10)
		if !ok1 || !ok2 || gotInt.Cmp(wantInt) != 0 {
			t.Errorf("row %d %s = %s, legacy %s", row, name, got, want)
		}
	}
	for i := range collected {
		if collected[i].Tx.EvtLogId != legacy[i].Tx.EvtLogId {
			t.Fatalf("row %d event = %d, legacy %d",
				i, collected[i].Tx.EvtLogId, legacy[i].Tx.EvtLogId)
		}
		equalDecimal("minted", collected[i].MintedWei, legacy[i].MintAmount, i)
		equalDecimal("burned", collected[i].BurnedWei, legacy[i].BurnAmount, i)
		equalDecimal("net", collected[i].NetWei, legacy[i].Amount, i)
		equalDecimal("running total", collected[i].TotalSupplyWei, legacy[i].TotalSupply, i)
		if collected[i].BidType != legacy[i].BidType ||
			!strings.EqualFold(collected[i].BidderAddr, legacy[i].BidderAddr) {
			t.Errorf("row %d identity = %d/%s, legacy %d/%s", i,
				collected[i].BidType, collected[i].BidderAddr,
				legacy[i].BidType, legacy[i].BidderAddr)
		}
	}

	// The fixture must include at least one CST bid so the burn branch is
	// exercised.
	var sawBurn bool
	for i := range collected {
		burned, _ := new(big.Int).SetString(collected[i].BurnedWei, 10)
		if burned.Sign() > 0 {
			sawBurn = true
		}
	}
	if !sawBurn {
		t.Fatal("no CST bid burned supply; the burn branch is untested")
	}

	if _, _, err := r.CosmicTokenSupplyByBidPage(ctx, nil, 0); err == nil {
		t.Fatal("zero limit accepted")
	}
	if _, _, err := r.CosmicTokenSupplyByBidPage(ctx, &SupplyChangePageCursor{}, 1); err == nil {
		t.Fatal("zero cursor accepted")
	}
}

func TestCosmicTokenSupplyDaily(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	// The fixture bids all land in one era; compare a generous window
	// against the legacy inclusive-day query (the v2 window is half-open,
	// so its exclusive end is one day past the legacy inclusive end).
	from := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC)
	legacy, err := r.CosmicTokenSupplyHistoryByDate(ctx, "20250101", "20261230")
	if err != nil {
		t.Fatal(err)
	}
	if len(legacy) == 0 {
		t.Fatal("fixture has no daily rows; the test is vacuous")
	}

	got, err := r.CosmicTokenSupplyDaily(ctx, from, to)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != len(legacy) {
		t.Fatalf("daily rows = %d, legacy %d", len(got), len(legacy))
	}
	for i := range got {
		gotDay := strings.ReplaceAll(got[i].Date, "-", "")
		if gotDay != legacy[i].Date {
			t.Fatalf("row %d day = %s, legacy %s", i, got[i].Date, legacy[i].Date)
		}
		if got[i].BidCount != legacy[i].NumBids {
			t.Errorf("row %d bids = %d, legacy %d", i, got[i].BidCount, legacy[i].NumBids)
		}
		for name, pair := range map[string][2]string{
			"minted":  {got[i].MintedWei, legacy[i].MintAmount},
			"burned":  {got[i].BurnedWei, legacy[i].BurnAmount},
			"net":     {got[i].NetWei, legacy[i].Amount},
			"running": {got[i].TotalSupplyWei, legacy[i].TotalSupply},
		} {
			gotInt, ok1 := new(big.Int).SetString(pair[0], 10)
			wantInt, ok2 := new(big.Int).SetString(pair[1], 10)
			if !ok1 || !ok2 || gotInt.Cmp(wantInt) != 0 {
				t.Errorf("row %d %s = %s, legacy %s", i, name, pair[0], pair[1])
			}
		}
	}

	// The upper bound is exclusive: ending exactly on the first fixture
	// day returns nothing before it.
	firstDay, err := time.Parse("2006-01-02", got[0].Date)
	if err != nil {
		t.Fatal(err)
	}
	before, err := r.CosmicTokenSupplyDaily(ctx, from, firstDay)
	if err != nil || len(before) != 0 {
		t.Fatalf("pre-window rows = %d, err=%v", len(before), err)
	}
	including, err := r.CosmicTokenSupplyDaily(ctx, from, firstDay.AddDate(0, 0, 1))
	if err != nil || len(including) != 1 {
		t.Fatalf("first-day window rows = %d, err=%v", len(including), err)
	}

	// Window validation.
	if _, err := r.CosmicTokenSupplyDaily(ctx, to, from); err == nil {
		t.Fatal("reversed window accepted")
	}
	if _, err := r.CosmicTokenSupplyDaily(ctx, from, from); err == nil {
		t.Fatal("empty window accepted")
	}
	if _, err := r.CosmicTokenSupplyDaily(ctx, from,
		from.AddDate(0, 0, MaxSupplyDailyWindowDays+1)); err == nil {
		t.Fatal("oversized window accepted")
	}
	edge, err := r.CosmicTokenSupplyDaily(ctx,
		time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(1990, 1, 2, 0, 0, 0, 0, time.UTC))
	if err != nil || len(edge) != 0 {
		t.Fatalf("empty era rows = %d, err=%v", len(edge), err)
	}
}

func TestMarketingRewardsGlobalPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	legacy, err := r.MarketingRewardHistoryGlobal(ctx, 0, 1000)
	if err != nil {
		t.Fatal(err)
	}
	if len(legacy) == 0 {
		t.Fatal("fixture has no marketing rewards; the test is vacuous")
	}

	var collected []MarketingRewardRecord
	var after *UserEventPageCursor
	for {
		page, hasMore, err := r.MarketingRewardsGlobalPage(ctx, after, 1)
		if err != nil {
			t.Fatalf("page walk: %v", err)
		}
		for i := range page {
			if len(collected) > 0 &&
				page[i].Tx.EvtLogId >= collected[len(collected)-1].Tx.EvtLogId {
				t.Fatalf("unordered rewards at %d", page[i].Tx.EvtLogId)
			}
			collected = append(collected, page[i])
		}
		if !hasMore {
			break
		}
		after = &UserEventPageCursor{EventLogID: page[len(page)-1].Tx.EvtLogId}
	}
	if len(collected) != len(legacy) {
		t.Fatalf("collected %d rewards, legacy %d", len(collected), len(legacy))
	}
	for i := range collected {
		if collected[i].Tx.EvtLogId != legacy[i].Tx.EvtLogId ||
			!strings.EqualFold(collected[i].MarketerAddr, legacy[i].MarketerAddr) {
			t.Errorf("row %d = %+v, legacy %+v", i, collected[i], legacy[i])
		}
		gotAmount, ok1 := new(big.Int).SetString(collected[i].AmountWei, 10)
		wantAmount, ok2 := new(big.Int).SetString(legacy[i].Amount, 10)
		if !ok1 || !ok2 || gotAmount.Cmp(wantAmount) != 0 {
			t.Errorf("row %d amount = %s, legacy %s", i, collected[i].AmountWei, legacy[i].Amount)
		}
	}
}

func TestGlobalTokensRejectAmbiguousMintSource(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	// A synthetic round-1 endurance prize pointing at dave's round-1
	// main-prize token makes token 7 match two mint sources; the global
	// directory and the detail must fail loudly instead of picking one.
	if _, err := r.pool().Exec(ctx, `INSERT INTO evt_log(id, block_num, tx_id, contract_aid, topic0_sig, log_index, log_rlp)
		VALUES (8031, 120, 1021, 2, 'ep000002', 97, '\x00')`); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := r.pool().Exec(context.Background(),
			"DELETE FROM evt_log WHERE id=8031"); err != nil {
			t.Errorf("cleaning synthetic endurance prize: %v", err)
		}
	})
	if _, err := r.pool().Exec(ctx, `INSERT INTO cg_endurance_prize(evtlog_id, block_num, tx_id, time_stamp, contract_aid, winner_aid, round_num, erc721_token_id, erc20_amount)
		VALUES (8031, 120, 1021, TO_TIMESTAMP(1767227600), 2, 23, 1, 7, 1)`); err != nil {
		t.Fatal(err)
	}

	if _, _, err := r.CosmicSignatureTokensGlobalPage(ctx, GlobalTokenFilter{}, nil, 50); err == nil {
		t.Fatal("global directory accepted a token with two mint sources")
	} else if !strings.Contains(err.Error(), "mint sources") {
		t.Fatalf("unexpected ambiguity error: %v", err)
	}
	if _, err := r.CosmicSignatureTokenDetailV2(ctx, 7); err == nil {
		t.Fatal("token detail accepted a token with two mint sources")
	}
}

func TestCosmicSignatureTokenDetailRejectsOrphanStake(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	// A membership row whose locking action does not exist would render a
	// staked token without stake data; the detail must fail loudly.
	if _, err := r.pool().Exec(ctx, `INSERT INTO cg_staked_token_cst(staker_aid, token_id, stake_action_id)
		VALUES (21, 1, 999999)`); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := r.pool().Exec(context.Background(),
			"DELETE FROM cg_staked_token_cst WHERE token_id=1 AND stake_action_id=999999"); err != nil {
			t.Errorf("cleaning orphan stake row: %v", err)
		}
	})

	if _, err := r.CosmicSignatureTokenDetailV2(ctx, 1); err == nil {
		t.Fatal("token detail accepted a stake without its action")
	} else if !strings.Contains(err.Error(), "stake action") {
		t.Fatalf("unexpected orphan-stake error: %v", err)
	}
}

func TestGlobalDirectoryErrorPaths(t *testing.T) {
	r := repo(t)
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()

	calls := map[string]func(ctx context.Context) error{
		"CosmicSignatureTokensGlobalPage": func(ctx context.Context) error {
			_, _, err := r.CosmicSignatureTokensGlobalPage(ctx, GlobalTokenFilter{}, nil, 1)
			return err
		},
		"CosmicSignatureTokenDetailV2": func(ctx context.Context) error {
			_, err := r.CosmicSignatureTokenDetailV2(ctx, 1)
			return err
		},
		"TokenNameHistoryPage": func(ctx context.Context) error {
			_, _, err := r.TokenNameHistoryPage(ctx, 1, nil, 1)
			return err
		},
		"TokenTransfersPage": func(ctx context.Context) error {
			_, _, err := r.TokenTransfersPage(ctx, 1, nil, 1)
			return err
		},
		"CosmicSignatureHoldersPage": func(ctx context.Context) error {
			_, _, err := r.CosmicSignatureHoldersPage(ctx, nil, 1)
			return err
		},
		"CosmicTokenHoldersPage": func(ctx context.Context) error {
			_, _, err := r.CosmicTokenHoldersPage(ctx, nil, 1)
			return err
		},
		"CosmicTokenStatisticsV2": func(ctx context.Context) error {
			_, err := r.CosmicTokenStatisticsV2(ctx)
			return err
		},
		"CosmicTokenSupplyByBidPage": func(ctx context.Context) error {
			_, _, err := r.CosmicTokenSupplyByBidPage(ctx, nil, 1)
			return err
		},
		"CosmicTokenSupplyDaily": func(ctx context.Context) error {
			_, err := r.CosmicTokenSupplyDaily(ctx,
				time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC))
			return err
		},
		"MarketingRewardsGlobalPage": func(ctx context.Context) error {
			_, _, err := r.MarketingRewardsGlobalPage(ctx, nil, 1)
			return err
		},
	}
	for name, call := range calls {
		if err := call(cancelled); err == nil {
			t.Errorf("%s succeeded on a cancelled context", name)
		}
	}
}

func TestGlobalDirectoryReadIndexesExist(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	// #nosec G101 -- index names and column lists, not credentials.
	wantIndexes := map[string]string{
		"cg_token_name_token_evt_idx":      "(token_id, evtlog_id desc)",
		"cg_erc721_transfer_token_evt_idx": "(token_id, evtlog_id desc)",
		"cg_mint_event_named_token_idx":    "(token_id)",
		"cg_costok_owner_balance_idx":      "(cur_balance desc, owner_aid)",
	}
	for indexName, wantColumns := range wantIndexes {
		var definition string
		err := r.pool().QueryRow(ctx, `SELECT indexdef FROM pg_indexes
			WHERE schemaname='public' AND indexname=$1`, indexName).Scan(&definition)
		if err != nil {
			t.Errorf("read index %s: %v", indexName, err)
			continue
		}
		normalized := strings.ToLower(strings.Join(strings.Fields(definition), " "))
		if !strings.Contains(normalized, wantColumns) {
			t.Errorf("index %s definition = %s, want columns %s", indexName, definition, wantColumns)
		}
	}
}
