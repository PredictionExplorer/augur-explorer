"""Labels and formatting for Cosmic Game admin / config-change events."""
from __future__ import annotations

from datetime import datetime, timezone
from typing import Any

# Mirrors rwcg/websrv/templates/cosmicgame/cg_system_admin_events_in_range.html
ADMIN_EVENT_TYPE_NAMES: dict[int, str] = {
    0: "Undefined",
    1: "CharityEthDonationAmountPercentChanged",
    2: "MainEthPrizeAmountPercentChanged",
    3: "RaffleTotalEthPrizeAmountPercentChanged",
    4: "StakingTotalEthRewardAmountPercentChanged",
    5: "NumRaffleEthPrizesForBiddersChanged",
    6: "NumRaffleCosmicSignatureNftsForBiddersChanged",
    7: "DelayDurationBeforeRoundActivationChanged",
    8: "NumRaffleNftsForRandomWalkStakersChanged",
    9: "CharityAddressChanged",
    10: "RandomWalkAddressChanged",
    11: "PrizesWalletAddressChanged",
    12: "StakingWalletCosmicSignatureNftAddressChanged",
    13: "StakingWalletRandomWalkNftAddressChanged",
    14: "MarketingWalletAddressChanged",
    15: "CosmicTokenAddressChanged",
    16: "CosmicSignatureNftAddressChanged",
    17: "Upgraded",
    18: "TimeIncreaseChanged",
    19: "TimeoutDurationToClaimMainPrizeChanged",
    20: "PriceIncreaseChanged",
    21: "MainPrizeTimeIncrementMicroSecondsChanged",
    22: "InitialSecondsUntilPrizeChanged",
    23: "TreasurerAddressChanged",
    24: "ActivationTimeChanged",
    25: "CstDutchAuctionDurationDivisorChanged",
    26: "CstPrizeAmountChanged",
    27: "StartingBidPriceCstMinLimitChanged",
    28: "MarketingWalletCstContributionChanged",
    29: "CstRewardAmountForBiddingChanged",
    30: "MaxMessageLengthChanged",
    31: "TokenGenerationScriptURL",
    32: "BaseURI",
    33: "Initialized",
    34: "OwnershipTransferred",
    35: "TimeoutDurationToWithdrawPrizesChanged",
    36: "EthDutchAuctionDurationDivisorChanged",
    37: "EthDutchAuctionEndingBidPriceDivisorChanged",
    38: "ChronoWarriorEthPrizeAmountPercentChanged",
}


def format_event_value(event: dict[str, Any]) -> str:
    parts: list[str] = []
    if event.get("AddressValue"):
        parts.append(str(event["AddressValue"]))
    if event.get("IntegerValue") not in (None, "", 0):
        parts.append(str(event["IntegerValue"]))
    if event.get("FloatValue") not in (None, "", 0):
        parts.append(str(event["FloatValue"]))
    if event.get("StringValue"):
        parts.append(str(event["StringValue"]))
    return ", ".join(parts) if parts else "(no value)"


def summarize_admin_event(event: dict[str, Any]) -> str:
    record_type = int(event.get("RecordType") or 0)
    name = ADMIN_EVENT_TYPE_NAMES.get(record_type, f"RecordType{record_type}")
    dt = event.get("DateTime") or ""
    value = format_event_value(event)
    return f"- {name} at {dt}: {value}"


def parse_iso_timestamp(value: Any) -> int | None:
    if value is None:
        return None
    if isinstance(value, (int, float)):
        ts = int(value)
        return ts if ts > 0 else None
    text = str(value).strip()
    if not text:
        return None
    try:
        if text.isdigit():
            ts = int(text)
            return ts if ts > 0 else None
        normalized = text.replace(" ", "T") if "T" not in text and "+" in text else text
        dt = datetime.fromisoformat(normalized.replace("Z", "+00:00"))
        if dt.tzinfo is None:
            dt = dt.replace(tzinfo=timezone.utc)
        return int(dt.timestamp())
    except ValueError:
        return None


def filter_events_to_activation_window(
    events: list[dict[str, Any]],
    *,
    param_window_start_ts: int | None,
    activation_ts: int | None,
) -> list[dict[str, Any]]:
    """Keep admin events that fall in the param/activation window when timestamps are known."""
    if not events:
        return events
    if param_window_start_ts is None and activation_ts is None:
        return events
    filtered: list[dict[str, Any]] = []
    for event in events:
        ts = int(event.get("TimeStamp") or 0)
        if ts <= 0:
            filtered.append(event)
            continue
        if param_window_start_ts is not None and ts < param_window_start_ts:
            continue
        if activation_ts is not None and ts > activation_ts:
            continue
        filtered.append(event)
    return filtered if filtered else events
