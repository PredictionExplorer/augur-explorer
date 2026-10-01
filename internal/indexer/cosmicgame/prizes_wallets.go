// Prizes-wallet tracking. The PrizesWallet is not upgradeable: a fix means
// deploying a new contract and pointing the game at it with setPrizesWallet
// (PrizesWalletAddressChanged). Prizes deposited into the old wallet remain
// there until their winners withdraw them, so the ETL must keep watching
// every wallet the game has ever used. This file keeps that set current:
// seeded at startup from cg_contracts plus the indexed
// PrizesWalletAddressChanged history (Contracts.PrizesWallets), grown at
// runtime by storePrizesWalletAddressChanged. Unlike game implementations,
// no backfill is needed: a prizes wallet only emits events when the game
// calls into it, which cannot happen before the PrizesWalletAddressChanged
// event that announces it.

package cosmicgame

import (
	"context"
	"slices"

	ethcommon "github.com/ethereum/go-ethereum/common"
)

// PrizesWallets returns the live set of prizes-wallet contracts.
func (h *Handlers) PrizesWallets() []ethcommon.Address {
	h.pwMu.RLock()
	defer h.pwMu.RUnlock()
	return slices.Clone(h.prizesWallets)
}

// addPrizesWallet records wallet; it reports whether the address was new.
func (h *Handlers) addPrizesWallet(wallet ethcommon.Address) bool {
	if wallet == (ethcommon.Address{}) {
		return false
	}
	h.pwMu.Lock()
	defer h.pwMu.Unlock()
	if slices.Contains(h.prizesWallets, wallet) {
		return false
	}
	h.prizesWallets = append(h.prizesWallets, wallet)
	return true
}

// SetOnNewPrizesWallet installs the callback storePrizesWalletAddressChanged
// invokes for a wallet address it has not seen before. cmd/cg-etl uses it to
// add the address to the engine's FilterLogs set. The callback runs inside
// the block's database transaction and must not write to the store itself.
func (h *Handlers) SetOnNewPrizesWallet(fn func(ctx context.Context, wallet ethcommon.Address, blockNum int64)) {
	h.pwMu.Lock()
	defer h.pwMu.Unlock()
	h.onNewPrizesWallet = fn
}

// notePrizesWallet is called by storePrizesWalletAddressChanged with the
// newly configured wallet.
func (h *Handlers) notePrizesWallet(ctx context.Context, wallet ethcommon.Address, blockNum int64) {
	if !h.addPrizesWallet(wallet) {
		return
	}
	h.log.Info("new prizes wallet", "prizes_wallet", wallet.Hex(), "block", blockNum)
	h.pwMu.RLock()
	fn := h.onNewPrizesWallet
	h.pwMu.RUnlock()
	if fn != nil {
		fn(ctx, wallet, blockNum)
	}
}

// prizesWalletSources lists every prizes wallet known so far; the dynamic
// source set of the PrizesWallet-emitted event handlers.
func (h *Handlers) prizesWalletSources() []ethcommon.Address {
	return h.PrizesWallets()
}
