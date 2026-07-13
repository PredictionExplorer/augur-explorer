// Donation events: ETH/ERC20/NFT donations, donated token/NFT claims and
// charity wallet traffic (donations received/sent, funds transferred to charity).

package cosmicgame

import (
	"context"
	"fmt"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

func (h *Handlers) decodeEthDonated(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGDonationEvent, error) {
	if err := requireTopics(lg, 3); err != nil {
		return nil, err
	}
	var ethEvt cgc.CosmicSignatureGameEthDonated
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "EthDonated", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGDonationEvent{}
	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.DonorAddr = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.Amount = ethEvt.Amount.String()
	evt.RoundNum = lg.Topics[1].Big().Int64()
	return evt, nil
}

func (h *Handlers) storeEthDonated(ctx context.Context, evt *cgmodel.CGDonationEvent) error {
	h.log.Info("EthDonated",
		"evt_id", evt.EvtId, "round", evt.RoundNum, "donor", evt.DonorAddr, "amount", evt.Amount)

	if err := h.repo.DeleteEthDonation(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertEthDonation(ctx, evt)
}

func (h *Handlers) decodeEthDonatedWithInfo(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGDonationWithInfoEvent, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	var ethEvt cgc.CosmicSignatureGameEthDonatedWithInfo
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "EthDonatedWithInfo", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGDonationWithInfoEvent{}
	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.DonorAddr = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.RecordId = lg.Topics[3].Big().Int64()
	evt.Amount = ethEvt.Amount.String()
	evt.RoundNum = lg.Topics[1].Big().Int64()
	return evt, nil
}

func (h *Handlers) storeEthDonatedWithInfo(ctx context.Context, evt *cgmodel.CGDonationWithInfoEvent) error {
	dataJSON, err := h.fetchDonationInfo(ctx, evt.RecordId)
	if err != nil {
		return fmt.Errorf("EthDonatedWithInfo (evt id %v): fetching donation info record: %w", evt.EvtId, err)
	}

	h.log.Info("EthDonatedWithInfo",
		"evt_id", evt.EvtId, "round", evt.RoundNum, "donor", evt.DonorAddr,
		"record_id", evt.RecordId, "amount", evt.Amount, "data_json", dataJSON)

	if err := h.repo.DeleteEthDonationWithInfo(ctx, evt.EvtId); err != nil {
		return err
	}
	if err := h.repo.InsertEthDonationWithInfo(ctx, evt); err != nil {
		return err
	}
	return h.repo.InsertDonationJSON(ctx, evt.RecordId, dataJSON)
}

func (h *Handlers) decodeDonationReceived(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGDonationReceivedEvent, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	var ethEvt cgc.CharityWalletDonationReceived
	if err := h.charityWalletABI.UnpackIntoInterface(&ethEvt, "DonationReceived", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGDonationReceivedEvent{}
	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.DonorAddr = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	evt.Amount = ethEvt.Amount.String()
	return evt, nil
}

func (h *Handlers) storeDonationReceived(ctx context.Context, evt *cgmodel.CGDonationReceivedEvent) error {
	// The round is resolved from the MainPrizeClaimed event of the same
	// transaction; a standalone donation reports -1.
	var err error
	evt.RoundNum, err = h.prizeRoundInTx(ctx, evt.TxId)
	if err != nil {
		return fmt.Errorf("DonationReceived (evt id %v): %w", evt.EvtId, err)
	}

	h.log.Info("DonationReceived",
		"evt_id", evt.EvtId, "round", evt.RoundNum, "donor", evt.DonorAddr, "amount", evt.Amount)

	if err := h.repo.DeleteDonationReceived(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertDonationReceived(ctx, evt)
}

// decodeDonationSent handles the CharityWallet's FundsTransferredToCharity
// (the topic0 is shared with the game/marketing-wallet event of the same
// name; the registry's source filter separates them).
func (h *Handlers) decodeDonationSent(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGDonationSentEvent, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	var ethEvt cgc.CharityWalletFundsTransferredToCharity
	if err := h.charityWalletABI.UnpackIntoInterface(&ethEvt, "FundsTransferredToCharity", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGDonationSentEvent{}
	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.CharityAddr = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	evt.Amount = ethEvt.Amount.String()
	return evt, nil
}

func (h *Handlers) storeDonationSent(ctx context.Context, evt *cgmodel.CGDonationSentEvent) error {
	h.log.Info("CharityWallet FundsTransferredToCharity",
		"evt_id", evt.EvtId, "charity", evt.CharityAddr, "amount", evt.Amount)

	if err := h.repo.DeleteDonationSent(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertDonationSent(ctx, evt)
}

func (h *Handlers) decodeTokenDonated(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGERC20DonationEvent, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	var ethEvt cgc.IPrizesWalletTokenDonated
	if err := h.prizesWalletABI.UnpackIntoInterface(&ethEvt, "TokenDonated", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGERC20DonationEvent{}
	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.RoundNum = lg.Topics[1].Big().Int64()
	evt.DonorAddr = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.TokenAddr = ethcommon.BytesToAddress(lg.Topics[3][12:]).String()
	evt.Amount = ethEvt.Amount.String()
	return evt, nil
}

func (h *Handlers) storeTokenDonated(ctx context.Context, evt *cgmodel.CGERC20DonationEvent) error {
	// A bidAndDonateToken() call places the bid event directly before this
	// one; when EvtId-1 is an Approval event instead, the bid is at EvtId-2.
	var err error
	evt.BidId, err = h.bidIDByEvtlog(ctx, evt.EvtId-1)
	if err != nil {
		return err
	}
	if evt.BidId == -1 {
		evt.BidId, err = h.bidIDByEvtlog(ctx, evt.EvtId-2)
		if err != nil {
			return err
		}
	}

	h.log.Info("TokenDonated",
		"evt_id", evt.EvtId, "round", evt.RoundNum, "donor", evt.DonorAddr,
		"token", evt.TokenAddr, "amount", evt.Amount, "bid_id", evt.BidId)

	if err := h.repo.DeleteERC20Donation(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertERC20Donation(ctx, evt)
}

func (h *Handlers) decodeNftDonated(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGNFTDonationEvent, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	var ethEvt cgc.IPrizesWalletNftDonated
	if err := h.prizesWalletABI.UnpackIntoInterface(&ethEvt, "NftDonated", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGNFTDonationEvent{}
	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.DonorAddr = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.TokenAddr = ethcommon.BytesToAddress(lg.Topics[3][12:]).String()
	evt.RoundNum = lg.Topics[1].Big().Int64()
	evt.TokenId = ethEvt.NftId.Int64()
	evt.Index = ethEvt.Index.Int64()
	return evt, nil
}

func (h *Handlers) storeNftDonated(ctx context.Context, evt *cgmodel.CGNFTDonationEvent) error {
	// BidRowIDByEvtlogID reports 0 when the previous event carries no bid
	// (e.g. a pure Donate() call); only real DB failures error.
	var err error
	evt.BidId, err = h.repo.BidRowIDByEvtlogID(ctx, evt.EvtId-1)
	if err != nil {
		return fmt.Errorf("NftDonated (evt id %v): %w", evt.EvtId, err)
	}
	evt.NFTTokenURI, err = h.fetchTokenURI(ctx, evt.TokenId, ethcommon.HexToAddress(evt.TokenAddr))
	if err != nil {
		return fmt.Errorf("NftDonated (evt id %v): %w", evt.EvtId, err)
	}

	h.log.Info("NftDonated",
		"evt_id", evt.EvtId, "round", evt.RoundNum, "donor", evt.DonorAddr,
		"nft", evt.TokenAddr, "token_id", evt.TokenId, "token_uri", evt.NFTTokenURI,
		"bid_id", evt.BidId)

	if err := h.repo.DeleteNFTDonation(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertNFTDonation(ctx, evt)
}

func (h *Handlers) decodeDonatedTokenClaimed(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGDonatedTokenClaimed, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	var ethEvt cgc.PrizesWalletDonatedTokenClaimed
	if err := h.prizesWalletABI.UnpackIntoInterface(&ethEvt, "DonatedTokenClaimed", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGDonatedTokenClaimed{}
	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.RoundNum = lg.Topics[1].Big().Int64()
	evt.Amount = ethEvt.Amount.String()
	evt.BeneficiaryAddr = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.TokenAddr = ethcommon.BytesToAddress(lg.Topics[3][12:]).String()
	return evt, nil
}

func (h *Handlers) storeDonatedTokenClaimed(ctx context.Context, evt *cgmodel.CGDonatedTokenClaimed) error {
	h.log.Info("DonatedTokenClaimed",
		"evt_id", evt.EvtId, "round", evt.RoundNum, "beneficiary", evt.BeneficiaryAddr,
		"token", evt.TokenAddr, "amount", evt.Amount)

	if err := h.repo.DeleteDonatedTokenClaim(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertDonatedTokenClaim(ctx, evt)
}

func (h *Handlers) decodeDonatedNftClaimed(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGDonatedNFTClaimed, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	var ethEvt cgc.PrizesWalletDonatedNftClaimed
	if err := h.prizesWalletABI.UnpackIntoInterface(&ethEvt, "DonatedNftClaimed", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGDonatedNFTClaimed{}
	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.RoundNum = lg.Topics[1].Big().Int64()
	evt.BeneficiaryAddr = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.TokenAddr = ethcommon.BytesToAddress(lg.Topics[3][12:]).String()
	evt.TokenId = ethEvt.NftId.String()
	evt.Index = ethEvt.Index.Int64()
	return evt, nil
}

func (h *Handlers) storeDonatedNftClaimed(ctx context.Context, evt *cgmodel.CGDonatedNFTClaimed) error {
	h.log.Info("DonatedNftClaimed",
		"evt_id", evt.EvtId, "round", evt.RoundNum, "index", evt.Index,
		"beneficiary", evt.BeneficiaryAddr, "nft", evt.TokenAddr, "token_id", evt.TokenId)

	if err := h.repo.DeleteDonatedNFTClaim(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertDonatedNFTClaim(ctx, evt)
}

// decodeFundsToCharity handles the game's FundsTransferredToCharity emitted
// through the marketing wallet (same topic0 as the CharityWallet event;
// separated by source).
func (h *Handlers) decodeFundsToCharity(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGFundsToCharity, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	var ethEvt cgc.CosmicSignatureGameFundsTransferredToCharity
	if err := h.gameABI.UnpackIntoInterface(&ethEvt, "FundsTransferredToCharity", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGFundsToCharity{}
	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.CharityAddr = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	evt.Amount = ethEvt.Amount.String()
	return evt, nil
}

func (h *Handlers) storeFundsToCharity(ctx context.Context, evt *cgmodel.CGFundsToCharity) error {
	h.log.Info("FundsTransferredToCharity",
		"evt_id", evt.EvtId, "charity", evt.CharityAddr, "amount", evt.Amount)

	if err := h.repo.DeleteFundsToCharity(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertFundsToCharity(ctx, evt)
}
