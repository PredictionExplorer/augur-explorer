// Token events: CosmicSignature NFT mint/name/transfer, CosmicToken (ERC20)
// transfers and MarketingWallet reward payments.

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

func (h *Handlers) decodeNftNameChanged(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGTokenNameEvent, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	var ethEvt cgc.ICosmicSignatureNftNftNameChanged
	if err := h.signatureABI.UnpackIntoInterface(&ethEvt, "NftNameChanged", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGTokenNameEvent{}
	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.TokenId = lg.Topics[1].Big().Int64()
	evt.TokenName = ethEvt.NftName
	return evt, nil
}

func (h *Handlers) storeNftNameChanged(ctx context.Context, evt *cgmodel.CGTokenNameEvent) error {
	h.log.Info("NftNameChanged", "evt_id", evt.EvtId, "token_id", evt.TokenId, "name", evt.TokenName)

	if err := h.repo.DeleteTokenName(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertTokenName(ctx, evt)
}

func (h *Handlers) decodeNftMinted(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGMintEvent, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	var ethEvt cgc.CosmicSignatureNftNftMinted
	if err := h.signatureABI.UnpackIntoInterface(&ethEvt, "NftMinted", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGMintEvent{}
	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.TokenId = lg.Topics[3].Big().Int64()
	evt.RoundNum = lg.Topics[1].Big().Int64()
	evt.OwnerAddr = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	// NftSeed is uint256; big.Int.Bytes() drops leading zero bytes, so format
	// with a fixed 64-hex-char width (32 bytes) to preserve leading zeros.
	evt.Seed = fmt.Sprintf("%064x", ethEvt.NftSeed)
	return evt, nil
}

func (h *Handlers) storeNftMinted(ctx context.Context, evt *cgmodel.CGMintEvent) error {
	h.log.Info("NftMinted",
		"evt_id", evt.EvtId, "round", evt.RoundNum, "token_id", evt.TokenId,
		"owner", evt.OwnerAddr, "seed", evt.Seed)

	if err := h.repo.DeleteMint(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertMint(ctx, evt)
}

func (h *Handlers) decodeMarketingRewardPaid(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGMarketingRewardPaid, error) {
	if err := requireTopics(lg, 2); err != nil {
		return nil, err
	}
	var ethEvt cgc.MarketingWalletRewardPaid
	if err := h.marketingWalletABI.UnpackIntoInterface(&ethEvt, "RewardPaid", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGMarketingRewardPaid{}
	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.ContractAddr = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.Marketer = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	evt.Amount = ethEvt.Amount.String()
	return evt, nil
}

func (h *Handlers) storeMarketingRewardPaid(ctx context.Context, evt *cgmodel.CGMarketingRewardPaid) error {
	h.log.Info("Marketing RewardPaid", "evt_id", evt.EvtId, "marketer", evt.Marketer, "amount", evt.Amount)

	if err := h.repo.DeleteMarketingRewardPaid(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertMarketingRewardPaid(ctx, evt)
}

// decodeCosmicSignatureTransfer handles the ERC721 Transfer of the
// CosmicSignature NFT (the Transfer topic0 is shared with the ERC20
// CosmicToken; the registry's source filter separates them).
func (h *Handlers) decodeCosmicSignatureTransfer(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGERC721Transfer, error) {
	if err := requireTopics(lg, 4); err != nil {
		return nil, err
	}
	evt := &cgmodel.CGERC721Transfer{}
	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.From = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	evt.To = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.TokenId = lg.Topics[3].Big().Int64()
	return evt, nil
}

func (h *Handlers) storeCosmicSignatureTransfer(ctx context.Context, evt *cgmodel.CGERC721Transfer) error {
	h.log.Info("CosmicSignature Transfer",
		"evt_id", evt.EvtId, "from", evt.From, "to", evt.To, "token_id", evt.TokenId)

	if err := h.repo.DeleteCosmicSignatureTransfer(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertCosmicSignatureTransfer(ctx, evt)
}

// decodeCosmicTokenTransfer handles the ERC20 Transfer of the CosmicToken.
func (h *Handlers) decodeCosmicTokenTransfer(lg *types.Log, elog *store.EthereumEventLog) (*cgmodel.CGERC20Transfer, error) {
	if err := requireTopics(lg, 3); err != nil {
		return nil, err
	}
	var ethEvt cgc.ERC20Transfer
	if err := h.erc20ABI.UnpackIntoInterface(&ethEvt, "Transfer", lg.Data); err != nil {
		return nil, err
	}

	evt := &cgmodel.CGERC20Transfer{}
	evt.EvtId = elog.EvtId
	evt.BlockNum = elog.BlockNum
	evt.TxId = elog.TxId
	evt.Contract = lg.Address.String()
	evt.TimeStamp = elog.TimeStamp
	evt.From = ethcommon.BytesToAddress(lg.Topics[1][12:]).String()
	evt.To = ethcommon.BytesToAddress(lg.Topics[2][12:]).String()
	evt.Value = ethEvt.Value.String()
	return evt, nil
}

func (h *Handlers) storeCosmicTokenTransfer(ctx context.Context, evt *cgmodel.CGERC20Transfer) error {
	h.log.Info("CosmicToken Transfer",
		"evt_id", evt.EvtId, "from", evt.From, "to", evt.To, "value", evt.Value)

	if err := h.repo.DeleteCosmicTokenTransfer(ctx, evt.EvtId); err != nil {
		return err
	}
	return h.repo.InsertCosmicTokenTransfer(ctx, evt)
}
