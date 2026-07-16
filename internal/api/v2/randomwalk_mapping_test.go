package v2

import (
	"strings"
	"testing"
	"time"

	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

const (
	rwTestHash  = "0xF000000000000000000000000000000000000000000000000000000000001030"
	rwTestAlice = "0x2100000000000000000000000000000000000021"
	rwTestCarol = "0x2300000000000000000000000000000000000023"
	rwTestDave  = "0x2400000000000000000000000000000000000024"
)

func validRandomWalkEventTx() rwstore.EventTx {
	return rwstore.EventTx{
		EvtLogID:  5080,
		BlockNum:  130,
		TxID:      1030,
		TxHash:    rwTestHash,
		TimeStamp: 1767228600,
		DateTime:  "2026-01-01T00:50:00Z",
	}
}

func validRandomWalkTokenRecord() rwstore.TokenRecord {
	return rwstore.TokenRecord{
		MintTx:           validRandomWalkEventTx(),
		TokenID:          10,
		MinterAid:        23,
		MinterAddr:       rwTestCarol,
		Seed:             "aa00000000000000000000000000000000000000000000000000000000000010",
		SeedNum:          "16",
		MintPriceWei:     "50000000000000000",
		CurOwnerAid:      24,
		CurOwnerAddr:     rwTestDave,
		TokenName:        "Wanderer",
		LastPriceWei:     "1000000000000000000",
		TradeCount:       1,
		TradingVolumeWei: "1000000000000000000",
	}
}

func TestMapRandomWalkTx(t *testing.T) {
	t.Parallel()
	mapped, err := mapRandomWalkTx(validRandomWalkEventTx())
	if err != nil {
		t.Fatal(err)
	}
	if mapped.EventLogID != 5080 || mapped.BlockNumber != 130 {
		t.Fatalf("tx = %+v", mapped)
	}
	if mapped.TransactionHash != strings.ToLower(rwTestHash) {
		t.Fatalf("hash not normalized: %q", mapped.TransactionHash)
	}
	if mapped.OccurredAt.Location() != time.UTC {
		t.Fatalf("timestamp not UTC: %v", mapped.OccurredAt)
	}

	broken := validRandomWalkEventTx()
	broken.EvtLogID = 0
	if _, err := mapRandomWalkTx(broken); err == nil {
		t.Fatal("zero event-log id accepted")
	}
	broken = validRandomWalkEventTx()
	broken.TxHash = "0x1234"
	if _, err := mapRandomWalkTx(broken); err == nil {
		t.Fatal("short hash accepted")
	}
	broken = validRandomWalkEventTx()
	broken.DateTime = "not-a-time"
	if _, err := mapRandomWalkTx(broken); err == nil {
		t.Fatal("unparseable timestamp accepted")
	}
}

func TestMapRandomWalkToken(t *testing.T) {
	t.Parallel()
	token, err := mapRandomWalkToken(validRandomWalkTokenRecord())
	if err != nil {
		t.Fatal(err)
	}
	if token.NftTokenId != 10 || token.MintPriceWei != "50000000000000000" ||
		token.LastPriceWei != "1000000000000000000" || token.SeedNumber != "16" ||
		token.MinterAddress != rwTestCarol || token.CurrentOwnerAddress != rwTestDave ||
		token.TokenName == nil || *token.TokenName != "Wanderer" {
		t.Fatalf("token = %+v", token)
	}

	unnamed := validRandomWalkTokenRecord()
	unnamed.TokenName = ""
	mapped, err := mapRandomWalkToken(unnamed)
	if err != nil || mapped.TokenName != nil {
		t.Fatalf("unnamed token = %+v, %v", mapped, err)
	}

	failures := map[string]func(*rwstore.TokenRecord){
		"negative token":  func(r *rwstore.TokenRecord) { r.TokenID = -1 },
		"empty seed":      func(r *rwstore.TokenRecord) { r.Seed = "" },
		"bad seed number": func(r *rwstore.TokenRecord) { r.SeedNum = "0x10" },
		"bad mint price":  func(r *rwstore.TokenRecord) { r.MintPriceWei = "-5" },
		"bad last price":  func(r *rwstore.TokenRecord) { r.LastPriceWei = "" },
		"bad volume":      func(r *rwstore.TokenRecord) { r.TradingVolumeWei = "1.5" },
		"negative trades": func(r *rwstore.TokenRecord) { r.TradeCount = -1 },
		"zero minter":     func(r *rwstore.TokenRecord) { r.MinterAddr = "0x0000000000000000000000000000000000000000" },
		"bad owner":       func(r *rwstore.TokenRecord) { r.CurOwnerAddr = "nope" },
		"bad tx":          func(r *rwstore.TokenRecord) { r.MintTx.TxHash = "" },
	}
	for name, mutate := range failures {
		record := validRandomWalkTokenRecord()
		mutate(&record)
		if _, err := mapRandomWalkToken(record); err == nil {
			t.Errorf("%s accepted", name)
		}
	}
}

func TestMapRandomWalkTokenDetail(t *testing.T) {
	t.Parallel()
	renamed := rwstore.TokenDetailRecord{
		TokenRecord:    validRandomWalkTokenRecord(),
		NameChangeTs:   1767229000,
		NameChangeText: "2026-01-01T00:56:40Z",
	}
	detail, err := mapRandomWalkTokenDetail(renamed)
	if err != nil || detail.NameChangedAt == nil ||
		!detail.NameChangedAt.Equal(time.Unix(1767229000, 0)) {
		t.Fatalf("detail = %+v, %v", detail, err)
	}

	never := rwstore.TokenDetailRecord{TokenRecord: validRandomWalkTokenRecord()}
	detail, err = mapRandomWalkTokenDetail(never)
	if err != nil || detail.NameChangedAt != nil {
		t.Fatalf("never-renamed detail = %+v, %v", detail, err)
	}

	broken := renamed
	broken.NameChangeText = "bogus"
	if _, err := mapRandomWalkTokenDetail(broken); err == nil {
		t.Fatal("unparseable rename timestamp accepted")
	}
}

func TestMapRandomWalkTokenNameChange(t *testing.T) {
	t.Parallel()
	record := rwstore.TokenNameChangeRecord{
		Tx:       validRandomWalkEventTx(),
		TokenID:  10,
		NewName:  "Wanderer",
		OwnerAid: 23,
		Owner:    rwTestCarol,
	}
	mapped, err := mapRandomWalkTokenNameChange(record)
	if err != nil || mapped.NewName != "Wanderer" || mapped.OwnerAddress != rwTestCarol {
		t.Fatalf("rename = %+v, %v", mapped, err)
	}

	// An empty name is a legitimate cleared-name event.
	record.NewName = ""
	mapped, err = mapRandomWalkTokenNameChange(record)
	if err != nil || mapped.NewName != "" {
		t.Fatalf("cleared rename = %+v, %v", mapped, err)
	}

	record.Owner = ""
	if _, err := mapRandomWalkTokenNameChange(record); err == nil {
		t.Fatal("missing renamer accepted")
	}
}

func validRandomWalkTokenEventRecord(kind rwstore.TokenEventKind) rwstore.TokenEventRecord {
	record := rwstore.TokenEventRecord{
		Tx:      validRandomWalkEventTx(),
		Kind:    kind,
		TokenID: 10,
	}
	switch kind {
	case rwstore.TokenEventMint:
		record.MinterAddr = rwTestCarol
		record.Seed = "aa10"
		record.SeedNum = "16"
		record.PriceWei = "50000000000000000"
	case rwstore.TokenEventTransfer:
		record.FromAddr = rwTestCarol
		record.ToAddr = rwTestDave
	case rwstore.TokenEventNameChange:
		record.NewName = "Wanderer"
		record.HasNewName = true
	case rwstore.TokenEventListed, rwstore.TokenEventOfferCanceled:
		record.OfferID = 4
		record.HasOffer = true
		record.OfferType = 1
		record.PriceWei = "3000000000000000000"
		record.MakerAddr = rwTestDave
	case rwstore.TokenEventPurchase:
		record.OfferID = 1
		record.HasOffer = true
		record.OfferType = 1
		record.PriceWei = "1000000000000000000"
		record.BuyerAddr = rwTestDave
		record.SellerAddr = rwTestCarol
	}
	return record
}

func TestMapRandomWalkTokenEventKinds(t *testing.T) {
	t.Parallel()
	expected := map[rwstore.TokenEventKind]RandomWalkTokenEventType{
		rwstore.TokenEventMint:          RandomWalkTokenEventTypeMint,
		rwstore.TokenEventTransfer:      RandomWalkTokenEventTypeTransfer,
		rwstore.TokenEventNameChange:    RandomWalkTokenEventTypeNameChange,
		rwstore.TokenEventListed:        RandomWalkTokenEventTypeListed,
		rwstore.TokenEventOfferCanceled: RandomWalkTokenEventTypeOfferCanceled,
		rwstore.TokenEventPurchase:      RandomWalkTokenEventTypePurchase,
	}
	for kind, eventType := range expected {
		mapped, err := mapRandomWalkTokenEvent(validRandomWalkTokenEventRecord(kind))
		if err != nil {
			t.Fatalf("%s: %v", kind, err)
		}
		if mapped.EventType != eventType {
			t.Fatalf("%s type = %q", kind, mapped.EventType)
		}
		switch kind {
		case rwstore.TokenEventMint:
			if mapped.MinterAddress == nil || mapped.Seed == nil ||
				mapped.SeedNumber == nil || mapped.PriceWei == nil {
				t.Fatalf("mint fields = %+v", mapped)
			}
		case rwstore.TokenEventTransfer:
			if mapped.FromAddress == nil || mapped.ToAddress == nil {
				t.Fatalf("transfer fields = %+v", mapped)
			}
		case rwstore.TokenEventNameChange:
			if mapped.NewName == nil {
				t.Fatalf("rename fields = %+v", mapped)
			}
		case rwstore.TokenEventListed, rwstore.TokenEventOfferCanceled:
			if mapped.OfferId == nil || mapped.OfferSide == nil ||
				mapped.PriceWei == nil || mapped.MakerAddress == nil {
				t.Fatalf("%s fields = %+v", kind, mapped)
			}
		case rwstore.TokenEventPurchase:
			if mapped.OfferId == nil || mapped.OfferSide == nil || mapped.PriceWei == nil ||
				mapped.BuyerAddress == nil || mapped.SellerAddress == nil {
				t.Fatalf("purchase fields = %+v", mapped)
			}
		}
	}

	// A burn keeps the zero recipient.
	burn := validRandomWalkTokenEventRecord(rwstore.TokenEventTransfer)
	burn.ToAddr = "0x0000000000000000000000000000000000000000"
	mapped, err := mapRandomWalkTokenEvent(burn)
	if err != nil || *mapped.ToAddress != "0x0000000000000000000000000000000000000000" {
		t.Fatalf("burn = %+v, %v", mapped, err)
	}

	// A cleared name maps to an empty newName, distinct from absence.
	cleared := validRandomWalkTokenEventRecord(rwstore.TokenEventNameChange)
	cleared.NewName = ""
	mapped, err = mapRandomWalkTokenEvent(cleared)
	if err != nil || mapped.NewName == nil || *mapped.NewName != "" {
		t.Fatalf("cleared rename = %+v, %v", mapped, err)
	}

	mutations := map[string]struct {
		kind   rwstore.TokenEventKind
		mutate func(*rwstore.TokenEventRecord)
	}{
		"negative token":        {rwstore.TokenEventMint, func(r *rwstore.TokenEventRecord) { r.TokenID = -1 }},
		"broken transaction":    {rwstore.TokenEventMint, func(r *rwstore.TokenEventRecord) { r.Tx.TxHash = "" }},
		"zero minter":           {rwstore.TokenEventMint, func(r *rwstore.TokenEventRecord) { r.MinterAddr = "0x0000000000000000000000000000000000000000" }},
		"mint without seed":     {rwstore.TokenEventMint, func(r *rwstore.TokenEventRecord) { r.Seed = "" }},
		"bad seed number":       {rwstore.TokenEventMint, func(r *rwstore.TokenEventRecord) { r.SeedNum = "0x10" }},
		"bad mint price":        {rwstore.TokenEventMint, func(r *rwstore.TokenEventRecord) { r.PriceWei = "-1" }},
		"bad transfer sender":   {rwstore.TokenEventTransfer, func(r *rwstore.TokenEventRecord) { r.FromAddr = "zzz" }},
		"bad transfer receiver": {rwstore.TokenEventTransfer, func(r *rwstore.TokenEventRecord) { r.ToAddr = "" }},
		"rename without name": {rwstore.TokenEventNameChange, func(r *rwstore.TokenEventRecord) {
			r.HasNewName = false
			r.NewName = ""
		}},
		"listing without offer": {rwstore.TokenEventListed, func(r *rwstore.TokenEventRecord) { r.HasOffer = false }},
		"negative offer id":     {rwstore.TokenEventListed, func(r *rwstore.TokenEventRecord) { r.OfferID = -4 }},
		"unknown side":          {rwstore.TokenEventPurchase, func(r *rwstore.TokenEventRecord) { r.OfferType = 9 }},
		"listing without price": {rwstore.TokenEventListed, func(r *rwstore.TokenEventRecord) { r.PriceWei = "" }},
		"zero maker":            {rwstore.TokenEventListed, func(r *rwstore.TokenEventRecord) { r.MakerAddr = "0x0000000000000000000000000000000000000000" }},
		"zero buyer":            {rwstore.TokenEventPurchase, func(r *rwstore.TokenEventRecord) { r.BuyerAddr = "0x0000000000000000000000000000000000000000" }},
		"zero seller":           {rwstore.TokenEventPurchase, func(r *rwstore.TokenEventRecord) { r.SellerAddr = "0x0000000000000000000000000000000000000000" }},
		"unknown kind":          {rwstore.TokenEventMint, func(r *rwstore.TokenEventRecord) { r.Kind = "bogus" }},
	}
	for name, tc := range mutations {
		record := validRandomWalkTokenEventRecord(tc.kind)
		tc.mutate(&record)
		if _, err := mapRandomWalkTokenEvent(record); err == nil {
			t.Errorf("%s accepted", name)
		}
	}
}

func TestMapRandomWalkTokenNameChangeRejectsInvalidInput(t *testing.T) {
	t.Parallel()
	valid := rwstore.TokenNameChangeRecord{
		Tx:       validRandomWalkEventTx(),
		TokenID:  10,
		NewName:  "Wanderer",
		OwnerAid: 23,
		Owner:    rwTestCarol,
	}
	negative := valid
	negative.TokenID = -1
	if _, err := mapRandomWalkTokenNameChange(negative); err == nil {
		t.Fatal("negative token id accepted")
	}
	brokenTx := valid
	brokenTx.Tx.TxHash = "bogus"
	if _, err := mapRandomWalkTokenNameChange(brokenTx); err == nil {
		t.Fatal("broken transaction accepted")
	}
}

func TestMapRandomWalkTokenDetailPropagatesBaseFailure(t *testing.T) {
	t.Parallel()
	broken := rwstore.TokenDetailRecord{TokenRecord: validRandomWalkTokenRecord()}
	broken.Seed = ""
	if _, err := mapRandomWalkTokenDetail(broken); err == nil {
		t.Fatal("broken base record accepted")
	}
}

func validRandomWalkOfferHistoryRecord() rwstore.OfferHistoryRecord {
	return rwstore.OfferHistoryRecord{
		ListTx:    validRandomWalkEventTx(),
		OfferID:   2,
		OfferType: 1,
		TokenID:   11,
		PriceWei:  "2000000000000000000",
		MakerAid:  24,
		MakerAddr: rwTestDave,
		Active:    true,
	}
}

func TestMapRandomWalkOfferHistoryEntry(t *testing.T) {
	t.Parallel()
	active, err := mapRandomWalkOfferHistoryEntry(validRandomWalkOfferHistoryRecord())
	if err != nil || active.Status != Active || active.Purchase != nil ||
		active.Cancellation != nil || active.SellerProfitWei != nil {
		t.Fatalf("active = %+v, %v", active, err)
	}
	if active.Side != Sell || active.MakerAddress != rwTestDave {
		t.Fatalf("active = %+v", active)
	}

	boughtTx := validRandomWalkEventTx()
	boughtTx.EvtLogID = 5089
	bought := validRandomWalkOfferHistoryRecord()
	bought.Active = false
	bought.ProfitWei = "-50000000000000000"
	bought.Purchase = &rwstore.OfferOutcomePurchase{
		Tx:         boughtTx,
		BuyerAid:   24,
		BuyerAddr:  rwTestDave,
		SellerAid:  23,
		SellerAddr: rwTestCarol,
	}
	mapped, err := mapRandomWalkOfferHistoryEntry(bought)
	if err != nil || mapped.Status != Bought || mapped.Purchase == nil ||
		mapped.Purchase.EventLogId != 5089 ||
		mapped.SellerProfitWei == nil || *mapped.SellerProfitWei != "-50000000000000000" {
		t.Fatalf("bought = %+v, %v", mapped, err)
	}

	cancelTx := validRandomWalkEventTx()
	cancelTx.EvtLogID = 5094
	canceled := validRandomWalkOfferHistoryRecord()
	canceled.Active = false
	canceled.Cancellation = &cancelTx
	mapped, err = mapRandomWalkOfferHistoryEntry(canceled)
	if err != nil || mapped.Status != Canceled || mapped.Cancellation == nil ||
		mapped.Cancellation.EventLogId != 5094 {
		t.Fatalf("canceled = %+v, %v", mapped, err)
	}

	both := bought
	both.Cancellation = &cancelTx
	if _, err := mapRandomWalkOfferHistoryEntry(both); err == nil {
		t.Fatal("bought+canceled offer accepted")
	}
	staleActive := bought
	staleActive.Cancellation = nil
	staleActive.Active = true
	if _, err := mapRandomWalkOfferHistoryEntry(staleActive); err == nil {
		t.Fatal("bought offer still active accepted")
	}
	closedNoOutcome := validRandomWalkOfferHistoryRecord()
	closedNoOutcome.Active = false
	if _, err := mapRandomWalkOfferHistoryEntry(closedNoOutcome); err == nil {
		t.Fatal("closed offer without an outcome accepted")
	}
	badProfit := bought
	badProfit.ProfitWei = "1.5"
	if _, err := mapRandomWalkOfferHistoryEntry(badProfit); err == nil {
		t.Fatal("fractional profit accepted")
	}

	failures := map[string]func(*rwstore.OfferHistoryRecord){
		"negative offer id":  func(r *rwstore.OfferHistoryRecord) { r.OfferID = -1 },
		"broken transaction": func(r *rwstore.OfferHistoryRecord) { r.ListTx.TxHash = "short" },
		"unknown side":       func(r *rwstore.OfferHistoryRecord) { r.OfferType = 9 },
		"bad price":          func(r *rwstore.OfferHistoryRecord) { r.PriceWei = "" },
		"zero maker":         func(r *rwstore.OfferHistoryRecord) { r.MakerAddr = "0x0000000000000000000000000000000000000000" },
	}
	for name, mutate := range failures {
		record := validRandomWalkOfferHistoryRecord()
		mutate(&record)
		if _, err := mapRandomWalkOfferHistoryEntry(record); err == nil {
			t.Errorf("%s accepted", name)
		}
	}
	outcomeFailures := map[string]func(*rwstore.OfferHistoryRecord){
		"broken purchase transaction": func(r *rwstore.OfferHistoryRecord) { r.Purchase.Tx.TxHash = "" },
		"zero purchase buyer": func(r *rwstore.OfferHistoryRecord) {
			r.Purchase.BuyerAddr = "0x0000000000000000000000000000000000000000"
		},
		"zero purchase seller": func(r *rwstore.OfferHistoryRecord) {
			r.Purchase.SellerAddr = "0x0000000000000000000000000000000000000000"
		},
	}
	for name, mutate := range outcomeFailures {
		record := bought
		purchase := *bought.Purchase
		record.Purchase = &purchase
		mutate(&record)
		if _, err := mapRandomWalkOfferHistoryEntry(record); err == nil {
			t.Errorf("%s accepted", name)
		}
	}
	activeCancel := canceled
	activeCancel.Active = true
	if _, err := mapRandomWalkOfferHistoryEntry(activeCancel); err == nil {
		t.Fatal("canceled offer still active accepted")
	}
	brokenCancel := canceled
	brokenCancelTx := *canceled.Cancellation
	brokenCancelTx.TxHash = "nope"
	brokenCancel.Cancellation = &brokenCancelTx
	if _, err := mapRandomWalkOfferHistoryEntry(brokenCancel); err == nil {
		t.Fatal("broken cancellation transaction accepted")
	}
}

func TestMapRandomWalkTradeAndOffer(t *testing.T) {
	t.Parallel()
	trade := rwstore.TradeRecord{
		Tx:         validRandomWalkEventTx(),
		OfferID:    1,
		OfferType:  1,
		TokenID:    10,
		PriceWei:   "1000000000000000000",
		BuyerAid:   24,
		BuyerAddr:  rwTestDave,
		SellerAid:  23,
		SellerAddr: rwTestCarol,
		ProfitWei:  "950000000000000000",
	}
	mapped, err := mapRandomWalkTrade(trade)
	if err != nil || mapped.Side != Sell || mapped.BuyerAddress != rwTestDave ||
		mapped.SellerProfitWei == nil || *mapped.SellerProfitWei != "950000000000000000" {
		t.Fatalf("trade = %+v, %v", mapped, err)
	}
	trade.ProfitWei = ""
	mapped, err = mapRandomWalkTrade(trade)
	if err != nil || mapped.SellerProfitWei != nil {
		t.Fatalf("untracked trade = %+v, %v", mapped, err)
	}
	trade.OfferType = 3
	if _, err := mapRandomWalkTrade(trade); err == nil {
		t.Fatal("unknown side accepted")
	}

	offer := rwstore.OfferRecord{
		ListTx:    validRandomWalkEventTx(),
		OfferID:   6,
		OfferType: 0,
		TokenID:   12,
		PriceWei:  "500000000000000000",
		MakerAid:  21,
		MakerAddr: rwTestAlice,
	}
	book, err := mapRandomWalkMarketplaceOffer(offer)
	if err != nil || book.Side != Buy || book.MakerAddress != rwTestAlice {
		t.Fatalf("offer = %+v, %v", book, err)
	}
	offer.PriceWei = "-1"
	if _, err := mapRandomWalkMarketplaceOffer(offer); err == nil {
		t.Fatal("negative price accepted")
	}

	tradeFailures := map[string]func(*rwstore.TradeRecord){
		"negative token":     func(r *rwstore.TradeRecord) { r.TokenID = -1 },
		"broken transaction": func(r *rwstore.TradeRecord) { r.Tx.EvtLogID = 0 },
		"bad price":          func(r *rwstore.TradeRecord) { r.PriceWei = "x" },
		"zero buyer":         func(r *rwstore.TradeRecord) { r.BuyerAddr = "0x0000000000000000000000000000000000000000" },
		"zero seller":        func(r *rwstore.TradeRecord) { r.SellerAddr = "0x0000000000000000000000000000000000000000" },
	}
	validTrade := rwstore.TradeRecord{
		Tx:         validRandomWalkEventTx(),
		OfferID:    1,
		OfferType:  1,
		TokenID:    10,
		PriceWei:   "1",
		BuyerAid:   24,
		BuyerAddr:  rwTestDave,
		SellerAid:  23,
		SellerAddr: rwTestCarol,
	}
	for name, mutate := range tradeFailures {
		record := validTrade
		mutate(&record)
		if _, err := mapRandomWalkTrade(record); err == nil {
			t.Errorf("trade %s accepted", name)
		}
	}
	badTradeProfit := validTrade
	badTradeProfit.ProfitWei = "x"
	if _, err := mapRandomWalkTrade(badTradeProfit); err == nil {
		t.Fatal("invalid trade profit accepted")
	}

	offerFailures := map[string]func(*rwstore.OfferRecord){
		"negative offer":     func(r *rwstore.OfferRecord) { r.OfferID = -1 },
		"broken transaction": func(r *rwstore.OfferRecord) { r.ListTx.DateTime = "bogus" },
		"unknown side":       func(r *rwstore.OfferRecord) { r.OfferType = 7 },
		"zero maker":         func(r *rwstore.OfferRecord) { r.MakerAddr = "0x0000000000000000000000000000000000000000" },
	}
	validOffer := rwstore.OfferRecord{
		ListTx:    validRandomWalkEventTx(),
		OfferID:   6,
		OfferType: 0,
		TokenID:   12,
		PriceWei:  "5",
		MakerAid:  21,
		MakerAddr: rwTestAlice,
	}
	for name, mutate := range offerFailures {
		record := validOffer
		mutate(&record)
		if _, err := mapRandomWalkMarketplaceOffer(record); err == nil {
			t.Errorf("offer %s accepted", name)
		}
	}
}

func TestMapRandomWalkFloorPrice(t *testing.T) {
	t.Parallel()
	empty, err := mapRandomWalkFloorPrice(rwstore.FloorPriceRecord{})
	if err != nil || empty.ActiveSellOfferCount != 0 || empty.Floor != nil {
		t.Fatalf("empty book = %+v, %v", empty, err)
	}

	populated := rwstore.FloorPriceRecord{
		ActiveSellOfferCount: 2,
		Floor: &rwstore.FloorListingRecord{
			OfferID:      2,
			TokenID:      11,
			PriceWei:     "2000000000000000000",
			ListedAtTs:   1767229300,
			ListedAtText: "2026-01-01T01:01:40Z",
		},
	}
	mapped, err := mapRandomWalkFloorPrice(populated)
	if err != nil || mapped.Floor == nil || mapped.Floor.PriceWei != "2000000000000000000" {
		t.Fatalf("floor = %+v, %v", mapped, err)
	}

	if _, err := mapRandomWalkFloorPrice(rwstore.FloorPriceRecord{
		ActiveSellOfferCount: 3,
	}); err == nil {
		t.Fatal("non-empty book without floor accepted")
	}
	orphan := populated
	orphan.ActiveSellOfferCount = 0
	if _, err := mapRandomWalkFloorPrice(orphan); err == nil {
		t.Fatal("floor listing on empty book accepted")
	}
	if _, err := mapRandomWalkFloorPrice(rwstore.FloorPriceRecord{
		ActiveSellOfferCount: -1,
	}); err == nil {
		t.Fatal("negative offer count accepted")
	}
	failures := map[string]func(*rwstore.FloorListingRecord){
		"negative floor token": func(r *rwstore.FloorListingRecord) { r.TokenID = -1 },
		"bad floor price":      func(r *rwstore.FloorListingRecord) { r.PriceWei = "x" },
		"bad floor timestamp":  func(r *rwstore.FloorListingRecord) { r.ListedAtText = "bogus" },
	}
	for name, mutate := range failures {
		record := populated
		floor := *populated.Floor
		record.Floor = &floor
		mutate(record.Floor)
		if _, err := mapRandomWalkFloorPrice(record); err == nil {
			t.Errorf("%s accepted", name)
		}
	}
}

func TestMapRandomWalkUserProfile(t *testing.T) {
	t.Parallel()
	record := rwstore.UserProfileRecord{
		Aid:              23,
		Address:          rwTestCarol,
		MintedTokenCount: 1,
		OwnedTokenCount:  0,
		TradeCount:       1,
		TradingVolumeWei: "1000000000000000000",
		ProfitWei:        "-950000000000000000",
		WithdrawalCount:  1,
	}
	profile, err := mapRandomWalkUserProfile(record)
	if err != nil || profile.Address != rwTestCarol ||
		profile.ProfitWei != "-950000000000000000" {
		t.Fatalf("profile = %+v, %v", profile, err)
	}

	record.TradingVolumeWei = "wat"
	if _, err := mapRandomWalkUserProfile(record); err == nil {
		t.Fatal("invalid volume accepted")
	}
	record.TradingVolumeWei = "1"
	record.Address = "nope"
	if _, err := mapRandomWalkUserProfile(record); err == nil {
		t.Fatal("invalid address accepted")
	}
	record.Address = rwTestCarol
	record.WithdrawalCount = -1
	if _, err := mapRandomWalkUserProfile(record); err == nil {
		t.Fatal("negative counter accepted")
	}
	record.WithdrawalCount = 0
	record.ProfitWei = "x"
	if _, err := mapRandomWalkUserProfile(record); err == nil {
		t.Fatal("invalid profit accepted")
	}

	zero := zeroRandomWalkUserProfile(rwTestAlice)
	if zero.Address != rwTestAlice || zero.ProfitWei != "0" ||
		zero.TradingVolumeWei != "0" || zero.OwnedTokenCount != 0 {
		t.Fatalf("zero profile = %+v", zero)
	}
}

func TestMapRandomWalkOwnedToken(t *testing.T) {
	t.Parallel()
	record := rwstore.OwnedTokenRecord{
		TokenID:          10,
		Seed:             "aa10",
		SeedNum:          "16",
		TokenName:        "Wanderer",
		LastPriceWei:     "1000000000000000000",
		TradeCount:       1,
		TradingVolumeWei: "1000000000000000000",
		HasMint:          true,
		MintTs:           1767228600,
		MintText:         "2026-01-01T00:50:00Z",
		MintPriceWei:     "50000000000000000",
	}
	token, err := mapRandomWalkOwnedToken(record)
	if err != nil || token.Seed == nil || token.MintedAt == nil ||
		token.MintPriceWei == nil || *token.MintPriceWei != "50000000000000000" {
		t.Fatalf("owned token = %+v, %v", token, err)
	}

	// A trigger-created placeholder row has no seed and no mint provenance.
	placeholder := rwstore.OwnedTokenRecord{
		TokenID:          99,
		LastPriceWei:     "0",
		TradingVolumeWei: "0",
	}
	token, err = mapRandomWalkOwnedToken(placeholder)
	if err != nil || token.Seed != nil || token.SeedNumber != nil ||
		token.MintedAt != nil || token.MintPriceWei != nil || token.TokenName != nil {
		t.Fatalf("placeholder token = %+v, %v", token, err)
	}

	broken := record
	broken.MintText = "bogus"
	if _, err := mapRandomWalkOwnedToken(broken); err == nil {
		t.Fatal("unparseable mint timestamp accepted")
	}

	failures := map[string]func(*rwstore.OwnedTokenRecord){
		"negative token":  func(r *rwstore.OwnedTokenRecord) { r.TokenID = -1 },
		"bad last price":  func(r *rwstore.OwnedTokenRecord) { r.LastPriceWei = "x" },
		"bad volume":      func(r *rwstore.OwnedTokenRecord) { r.TradingVolumeWei = "" },
		"bad seed number": func(r *rwstore.OwnedTokenRecord) { r.SeedNum = "seed" },
		"bad mint price":  func(r *rwstore.OwnedTokenRecord) { r.MintPriceWei = "-1" },
	}
	for name, mutate := range failures {
		mutated := record
		mutate(&mutated)
		if _, err := mapRandomWalkOwnedToken(mutated); err == nil {
			t.Errorf("%s accepted", name)
		}
	}
}

func validRandomWalkStatisticsRecord() rwstore.StatisticsRecord {
	return rwstore.StatisticsRecord{
		MintedCount:            4,
		UniqueOwnerCount:       3,
		TokenTradeCount:        1,
		TokenTradingVolumeWei:  "1000000000000000000",
		MintFundsWei:           "230000000000000000",
		MarketTradeCount:       1,
		MarketTradingVolumeWei: "1000000000000000000",
		ActiveSellOfferCount:   2,
		ActiveBuyOfferCount:    0,
		WithdrawalCount:        1,
		LastMint: &rwstore.LastMintRecord{
			TokenID:    13,
			PriceWei:   "65000000000000000",
			MintTs:     1767228900,
			MintText:   "2026-01-01T00:55:00Z",
			MinterAid:  22,
			MinterAddr: "0x2200000000000000000000000000000000000022",
		},
		LatestWithdrawal: &rwstore.LatestWithdrawalRecord{
			AmountWei:      "30000000000000000",
			OccurredTs:     1767229700,
			OccurredText:   "2026-01-01T01:08:20Z",
			WithdrawerAid:  23,
			WithdrawerAddr: rwTestCarol,
			TokenID:        10,
		},
	}
}

func TestMapRandomWalkStatistics(t *testing.T) {
	t.Parallel()
	statistics, err := mapRandomWalkStatistics(validRandomWalkStatisticsRecord())
	if err != nil {
		t.Fatal(err)
	}
	if statistics.Tokens.MintedCount != 4 || statistics.Tokens.LastMint == nil ||
		statistics.Tokens.LastMint.NftTokenId != 13 ||
		statistics.Marketplace.ActiveSellOfferCount != 2 ||
		statistics.Withdrawals.Latest == nil ||
		statistics.Withdrawals.Latest.AmountWei != "30000000000000000" {
		t.Fatalf("statistics = %+v", statistics)
	}

	fresh := validRandomWalkStatisticsRecord()
	fresh.LastMint = nil
	fresh.LatestWithdrawal = nil
	statistics, err = mapRandomWalkStatistics(fresh)
	if err != nil || statistics.Tokens.LastMint != nil || statistics.Withdrawals.Latest != nil {
		t.Fatalf("fresh statistics = %+v, %v", statistics, err)
	}

	failures := map[string]func(*rwstore.StatisticsRecord){
		"negative count":           func(r *rwstore.StatisticsRecord) { r.MintedCount = -1 },
		"bad volume":               func(r *rwstore.StatisticsRecord) { r.TokenTradingVolumeWei = "x" },
		"bad mint funds":           func(r *rwstore.StatisticsRecord) { r.MintFundsWei = "" },
		"bad market volume":        func(r *rwstore.StatisticsRecord) { r.MarketTradingVolumeWei = "x" },
		"bad last mint":            func(r *rwstore.StatisticsRecord) { r.LastMint.PriceWei = "-1" },
		"negative last mint token": func(r *rwstore.StatisticsRecord) { r.LastMint.TokenID = -1 },
		"zero last minter": func(r *rwstore.StatisticsRecord) {
			r.LastMint.MinterAddr = "0x0000000000000000000000000000000000000000"
		},
		"bad last mint timestamp":   func(r *rwstore.StatisticsRecord) { r.LastMint.MintText = "bogus" },
		"bad withdrawal":            func(r *rwstore.StatisticsRecord) { r.LatestWithdrawal.AmountWei = "" },
		"negative withdrawal token": func(r *rwstore.StatisticsRecord) { r.LatestWithdrawal.TokenID = -1 },
		"zero withdrawer": func(r *rwstore.StatisticsRecord) {
			r.LatestWithdrawal.WithdrawerAddr = "0x0000000000000000000000000000000000000000"
		},
		"bad withdrawal timestamp":     func(r *rwstore.StatisticsRecord) { r.LatestWithdrawal.OccurredText = "bogus" },
		"negative active offer counts": func(r *rwstore.StatisticsRecord) { r.ActiveBuyOfferCount = -1 },
	}
	for name, mutate := range failures {
		record := validRandomWalkStatisticsRecord()
		mutate(&record)
		if _, err := mapRandomWalkStatistics(record); err == nil {
			t.Errorf("%s accepted", name)
		}
	}
}

func TestMapRandomWalkTradingVolume(t *testing.T) {
	t.Parallel()
	window := analyticsWindow{
		analyticsRange:  analyticsRange{from: 100, to: 400},
		intervalSeconds: 100,
	}
	buckets := []rwstore.VolumeBucketRecord{
		{BucketStart: 100, TradeCount: 1, VolumeWei: "7"},
		{BucketStart: 200, TradeCount: 0, VolumeWei: "0"},
		{BucketStart: 300, TradeCount: 2, VolumeWei: "3"},
	}
	mapped, err := mapRandomWalkTradingVolume(window, "5", buckets)
	if err != nil {
		t.Fatal(err)
	}
	if mapped.BaseVolumeWei != "5" || len(mapped.Buckets) != 3 ||
		mapped.Buckets[0].CumulativeVolumeWei != "12" ||
		mapped.Buckets[1].CumulativeVolumeWei != "12" ||
		mapped.Buckets[2].CumulativeVolumeWei != "15" {
		t.Fatalf("volume = %+v", mapped)
	}

	outside := []rwstore.VolumeBucketRecord{{BucketStart: 400, TradeCount: 0, VolumeWei: "0"}}
	if _, err := mapRandomWalkTradingVolume(window, "0", outside); err == nil {
		t.Fatal("bucket at exclusive end accepted")
	}
	unordered := []rwstore.VolumeBucketRecord{
		{BucketStart: 200, TradeCount: 0, VolumeWei: "0"},
		{BucketStart: 100, TradeCount: 0, VolumeWei: "0"},
	}
	if _, err := mapRandomWalkTradingVolume(window, "0", unordered); err == nil {
		t.Fatal("unordered buckets accepted")
	}
	ghost := []rwstore.VolumeBucketRecord{{BucketStart: 100, TradeCount: 0, VolumeWei: "9"}}
	if _, err := mapRandomWalkTradingVolume(window, "0", ghost); err == nil {
		t.Fatal("volume without trades accepted")
	}
	if _, err := mapRandomWalkTradingVolume(window, "x", nil); err == nil {
		t.Fatal("invalid base volume accepted")
	}
	negative := []rwstore.VolumeBucketRecord{{BucketStart: 100, TradeCount: -1, VolumeWei: "0"}}
	if _, err := mapRandomWalkTradingVolume(window, "0", negative); err == nil {
		t.Fatal("negative trade count accepted")
	}
	badVolume := []rwstore.VolumeBucketRecord{{BucketStart: 100, TradeCount: 1, VolumeWei: "x"}}
	if _, err := mapRandomWalkTradingVolume(window, "0", badVolume); err == nil {
		t.Fatal("invalid bucket volume accepted")
	}
}

func TestMapRandomWalkListingFloorHistory(t *testing.T) {
	t.Parallel()
	window := analyticsWindow{
		analyticsRange:  analyticsRange{from: 100, to: 400},
		intervalSeconds: 100,
	}
	points := []rwstore.FloorPointRecord{
		{BucketStart: 100, FloorWei: "5"},
		{BucketStart: 300, FloorWei: "2"},
	}
	mapped, err := mapRandomWalkListingFloorHistory(window, points)
	if err != nil || len(mapped.Points) != 2 || mapped.Points[1].FloorPriceWei != "2" {
		t.Fatalf("floor history = %+v, %v", mapped, err)
	}
	if _, err := mapRandomWalkListingFloorHistory(window, []rwstore.FloorPointRecord{
		{BucketStart: 300, FloorWei: "2"},
		{BucketStart: 100, FloorWei: "5"},
	}); err == nil {
		t.Fatal("unordered points accepted")
	}
	if _, err := mapRandomWalkListingFloorHistory(window, []rwstore.FloorPointRecord{
		{BucketStart: 100, FloorWei: "-1"},
	}); err == nil {
		t.Fatal("negative floor accepted")
	}
}

func TestMapRandomWalkMintReport(t *testing.T) {
	t.Parallel()
	months := []rwstore.MonthlyMintRecord{
		{Year: 2021, Month: 11, MintCount: 2, MintedWei: "10"},
		{Year: 2022, Month: 1, MintCount: 1, MintedWei: "5"},
	}
	mapped, err := mapRandomWalkMintReport(months)
	if err != nil || len(mapped.Months) != 2 ||
		mapped.Months[0].CumulativeMintedWei != "10" ||
		mapped.Months[1].CumulativeMintedWei != "15" {
		t.Fatalf("mint report = %+v, %v", mapped, err)
	}

	if _, err := mapRandomWalkMintReport([]rwstore.MonthlyMintRecord{
		{Year: 2022, Month: 13, MintCount: 1, MintedWei: "1"},
	}); err == nil {
		t.Fatal("month 13 accepted")
	}
	if _, err := mapRandomWalkMintReport([]rwstore.MonthlyMintRecord{
		{Year: 2022, Month: 5, MintCount: 1, MintedWei: "1"},
		{Year: 2022, Month: 5, MintCount: 1, MintedWei: "1"},
	}); err == nil {
		t.Fatal("duplicate month accepted")
	}
	if _, err := mapRandomWalkMintReport([]rwstore.MonthlyMintRecord{
		{Year: 2022, Month: 5, MintCount: 0, MintedWei: "0"},
	}); err == nil {
		t.Fatal("zero-mint month accepted")
	}
	if _, err := mapRandomWalkMintReport([]rwstore.MonthlyMintRecord{
		{Year: 2022, Month: 5, MintCount: 1, MintedWei: "x"},
	}); err == nil {
		t.Fatal("invalid mint amount accepted")
	}
}

func TestMapRandomWalkWithdrawal(t *testing.T) {
	t.Parallel()
	record := rwstore.WithdrawalRecord{
		Tx:             validRandomWalkEventTx(),
		WithdrawerAid:  23,
		WithdrawerAddr: rwTestCarol,
		TokenID:        10,
		AmountWei:      "30000000000000000",
	}
	mapped, err := mapRandomWalkWithdrawal(record)
	if err != nil || mapped.WithdrawerAddress != rwTestCarol ||
		mapped.AmountWei != "30000000000000000" || mapped.NftTokenId != 10 {
		t.Fatalf("withdrawal = %+v, %v", mapped, err)
	}
	record.AmountWei = ""
	if _, err := mapRandomWalkWithdrawal(record); err == nil {
		t.Fatal("empty amount accepted")
	}
	record.AmountWei = "1"
	record.TokenID = -1
	if _, err := mapRandomWalkWithdrawal(record); err == nil {
		t.Fatal("negative token id accepted")
	}
	record.TokenID = 10
	record.Tx.BlockNum = -1
	if _, err := mapRandomWalkWithdrawal(record); err == nil {
		t.Fatal("broken transaction accepted")
	}
	record.Tx.BlockNum = 130
	record.WithdrawerAddr = "0x0000000000000000000000000000000000000000"
	if _, err := mapRandomWalkWithdrawal(record); err == nil {
		t.Fatal("zero withdrawer accepted")
	}
}

func TestCanonicalAnyAddress(t *testing.T) {
	t.Parallel()
	zero, err := canonicalAnyAddress("burn recipient", "0x0000000000000000000000000000000000000000")
	if err != nil || zero != "0x0000000000000000000000000000000000000000" {
		t.Fatalf("zero address = %q, %v", zero, err)
	}
	if _, err := canonicalAnyAddress("burn recipient", "not-hex"); err == nil {
		t.Fatal("invalid address accepted")
	}
}
