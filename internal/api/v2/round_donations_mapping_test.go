package v2

import (
	"strings"
	"testing"
	"time"

	cgprimitives "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func TestMapRoundEthDonationVariants(t *testing.T) {
	t.Parallel()

	t.Run("plain", func(t *testing.T) {
		t.Parallel()
		record := validRoundEthDonationRecord(cgstore.RoundEthDonationPlain)
		record.EthAmountWei = "00042"
		got, err := mapRoundEthDonation(record)
		if err != nil {
			t.Fatalf("mapRoundEthDonation: %v", err)
		}
		if got.Kind != Plain || got.EthAmountWei != "42" ||
			got.ContractRecordId != nil || got.Data != nil {
			t.Fatalf("plain donation = %+v", got)
		}
		assertDonationEventFields(t, got.EventLogId, got.Round, got.DonorAddress, got.TransactionHash, got.OccurredAt)
	})

	t.Run("with info", func(t *testing.T) {
		t.Parallel()
		record := validRoundEthDonationRecord(cgstore.RoundEthDonationWithInfo)
		recordID := int64(0)
		data := ""
		record.ContractRecordID = &recordID
		record.Data = &data
		got, err := mapRoundEthDonation(record)
		if err != nil {
			t.Fatalf("mapRoundEthDonation: %v", err)
		}
		if got.Kind != WithInfo || got.ContractRecordId == nil || *got.ContractRecordId != 0 ||
			got.Data == nil || *got.Data != "" {
			t.Fatalf("with-info donation = %+v", got)
		}
	})
}

func TestMapRoundERC20Donation(t *testing.T) {
	t.Parallel()

	record := validRoundERC20DonationRecord()
	record.AmountBaseUnits = "123456789012345678901234567890"
	got, err := mapRoundERC20Donation(record)
	if err != nil {
		t.Fatalf("mapRoundERC20Donation: %v", err)
	}
	if got.AmountBaseUnits != record.AmountBaseUnits ||
		got.TokenAddress != "0x3333333333333333333333333333333333333333" {
		t.Fatalf("ERC-20 donation = %+v", got)
	}
	assertDonationEventFields(t, got.EventLogId, got.Round, got.DonorAddress, got.TransactionHash, got.OccurredAt)
}

func TestMapRoundNFTDonation(t *testing.T) {
	t.Parallel()

	record := validRoundNFTDonationRecord()
	record.TokenURI = ""
	got, err := mapRoundNFTDonation(record)
	if err != nil {
		t.Fatalf("mapRoundNFTDonation: %v", err)
	}
	if got.TokenId != 777 || got.DonationIndex != 12 || got.TokenUri != "" ||
		got.TokenAddress != "0x3333333333333333333333333333333333333333" {
		t.Fatalf("NFT donation = %+v", got)
	}
	assertDonationEventFields(t, got.EventLogId, got.Round, got.DonorAddress, got.TransactionHash, got.OccurredAt)
}

func TestRoundDonationMappersRejectInvalidRecords(t *testing.T) {
	t.Parallel()

	tests := map[string]func() error{
		"eth negative round": func() error {
			record := validRoundEthDonationRecord(cgstore.RoundEthDonationPlain)
			record.RoundNum = -1
			_, err := mapRoundEthDonation(record)
			return err
		},
		"eth invalid transaction": func() error {
			record := validRoundEthDonationRecord(cgstore.RoundEthDonationPlain)
			record.Tx.EvtLogId = 0
			_, err := mapRoundEthDonation(record)
			return err
		},
		"eth invalid donor": func() error {
			record := validRoundEthDonationRecord(cgstore.RoundEthDonationPlain)
			record.DonorAddr = "not-an-address"
			_, err := mapRoundEthDonation(record)
			return err
		},
		"eth invalid amount": func() error {
			record := validRoundEthDonationRecord(cgstore.RoundEthDonationPlain)
			record.EthAmountWei = "-1"
			_, err := mapRoundEthDonation(record)
			return err
		},
		"plain carries info": func() error {
			record := validRoundEthDonationRecord(cgstore.RoundEthDonationPlain)
			value := int64(1)
			record.ContractRecordID = &value
			_, err := mapRoundEthDonation(record)
			return err
		},
		"with info missing fields": func() error {
			record := validRoundEthDonationRecord(cgstore.RoundEthDonationWithInfo)
			_, err := mapRoundEthDonation(record)
			return err
		},
		"with info negative record": func() error {
			record := validRoundEthDonationRecord(cgstore.RoundEthDonationWithInfo)
			recordID := int64(-1)
			data := "{}"
			record.ContractRecordID = &recordID
			record.Data = &data
			_, err := mapRoundEthDonation(record)
			return err
		},
		"unknown eth kind": func() error {
			record := validRoundEthDonationRecord(cgstore.RoundEthDonationKind("other"))
			_, err := mapRoundEthDonation(record)
			return err
		},
		"erc20 negative round": func() error {
			record := validRoundERC20DonationRecord()
			record.RoundNum = -1
			_, err := mapRoundERC20Donation(record)
			return err
		},
		"erc20 invalid donor": func() error {
			record := validRoundERC20DonationRecord()
			record.DonorAddr = "bad"
			_, err := mapRoundERC20Donation(record)
			return err
		},
		"erc20 invalid token": func() error {
			record := validRoundERC20DonationRecord()
			record.TokenAddr = "bad"
			_, err := mapRoundERC20Donation(record)
			return err
		},
		"erc20 invalid amount": func() error {
			record := validRoundERC20DonationRecord()
			record.AmountBaseUnits = "1.5"
			_, err := mapRoundERC20Donation(record)
			return err
		},
		"nft negative token": func() error {
			record := validRoundNFTDonationRecord()
			record.TokenID = -1
			_, err := mapRoundNFTDonation(record)
			return err
		},
		"nft negative index": func() error {
			record := validRoundNFTDonationRecord()
			record.DonationIndex = -1
			_, err := mapRoundNFTDonation(record)
			return err
		},
		"nft invalid donor": func() error {
			record := validRoundNFTDonationRecord()
			record.DonorAddr = "bad"
			_, err := mapRoundNFTDonation(record)
			return err
		},
		"nft invalid token": func() error {
			record := validRoundNFTDonationRecord()
			record.TokenAddr = "bad"
			_, err := mapRoundNFTDonation(record)
			return err
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if err := test(); err == nil {
				t.Fatal("mapper accepted invalid record")
			}
		})
	}
}

func validRoundEthDonationRecord(kind cgstore.RoundEthDonationKind) cgstore.RoundEthDonationRecord {
	return cgstore.RoundEthDonationRecord{
		Tx:           validDonationTransaction(),
		RoundNum:     4,
		DonorAddr:    "0x2222222222222222222222222222222222222222",
		EthAmountWei: "42",
		Kind:         kind,
	}
}

func validRoundERC20DonationRecord() cgstore.RoundERC20DonationRecord {
	return cgstore.RoundERC20DonationRecord{
		Tx:              validDonationTransaction(),
		RoundNum:        4,
		DonorAddr:       "0x2222222222222222222222222222222222222222",
		TokenAddr:       "0x3333333333333333333333333333333333333333",
		AmountBaseUnits: "42",
	}
}

func validRoundNFTDonationRecord() cgstore.RoundNFTDonationRecord {
	// #nosec G101 -- deterministic chain fixture values, not credentials.
	return cgstore.RoundNFTDonationRecord{
		Tx:            validDonationTransaction(),
		RoundNum:      4,
		DonorAddr:     "0x2222222222222222222222222222222222222222",
		TokenAddr:     "0x3333333333333333333333333333333333333333",
		TokenID:       777,
		DonationIndex: 12,
		TokenURI:      "ipfs://donation/777",
	}
}

func validDonationTransaction() cgprimitives.Transaction {
	return cgprimitives.Transaction{
		EvtLogId:  100,
		BlockNum:  200,
		TxId:      300,
		TxHash:    "0x" + strings.Repeat("AB", 32),
		TimeStamp: 1_767_226_100,
		DateTime:  "2026-01-01T01:01:40+01:00",
	}
}

func assertDonationEventFields(
	t *testing.T,
	eventLogID,
	round int64,
	donorAddress,
	transactionHash string,
	occurredAt time.Time,
) {
	t.Helper()
	if eventLogID != 100 || round != 4 ||
		donorAddress != "0x2222222222222222222222222222222222222222" ||
		transactionHash != "0x"+strings.Repeat("ab", 32) ||
		!occurredAt.Equal(time.Date(2026, 1, 1, 0, 1, 40, 0, time.UTC)) ||
		occurredAt.Location() != time.UTC {
		t.Fatalf("event fields = id:%d round:%d donor:%s tx:%s time:%s",
			eventLogID, round, donorAddress, transactionHash, occurredAt)
	}
}
