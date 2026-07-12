//go:build integration

package toolutil

import (
	"reflect"
	"sort"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
	"github.com/PredictionExplorer/augur-explorer/internal/testfixtures"
)

func TestContractAddressLookups(t *testing.T) {
	db := testdb.New(t)
	if err := testfixtures.Apply(t.Context(), db.SQL); err != nil {
		t.Fatal(err)
	}

	randomWalkIDs, err := GetContractAddressIds(db.SQL, ProjectRandomWalk)
	if err != nil {
		t.Fatal(err)
	}
	sort.Slice(randomWalkIDs, func(i, j int) bool { return randomWalkIDs[i] < randomWalkIDs[j] })
	if !reflect.DeepEqual(randomWalkIDs, []int64{8, 12}) {
		t.Fatalf("RandomWalk contract IDs = %v", randomWalkIDs)
	}
	cosmicGameIDs, err := GetContractAddressIds(db.SQL, ProjectCosmicGame)
	if err != nil {
		t.Fatal(err)
	}
	if len(cosmicGameIDs) != 11 {
		t.Fatalf("CosmicGame contract IDs = %v", cosmicGameIDs)
	}
	addresses, err := GetContractAddrsByAids(db.SQL, []int64{12, 8})
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(addresses, []string{
		"0x8000000000000000000000000000000000000008",
		"0x1200000000000000000000000000000000000012",
	}) {
		t.Fatalf("contract addresses = %v", addresses)
	}
	empty, err := GetContractAddrsByAids(db.SQL, nil)
	if err != nil || len(empty) != 0 {
		t.Fatalf("empty lookup = %v, %v", empty, err)
	}
	if _, err := GetContractAddressIds(db.SQL, "other"); err == nil {
		t.Fatal("unknown project was accepted")
	}
}

func TestContractAddressLookupsPropagateClosedDatabase(t *testing.T) {
	db := testdb.New(t)
	if err := db.SQL.Close(); err != nil {
		t.Fatal(err)
	}
	if _, err := GetContractAddressIds(db.SQL, ProjectCosmicGame); err == nil {
		t.Fatal("contract ID lookup succeeded on closed database")
	}
	if _, err := GetContractAddrsByAids(db.SQL, []int64{1}); err == nil {
		t.Fatal("address lookup succeeded on closed database")
	}
}
