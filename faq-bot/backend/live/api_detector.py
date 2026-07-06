"""Detect questions that need live Cosmic Game backend API data."""
from __future__ import annotations

import re

from live.detector import needs_champions_state, needs_round_end_time, needs_staking_stats

_ROUND_NUM = re.compile(r"\bround\s*(?:#|number\s*)?(\d+)\b", re.IGNORECASE)
_PRIZE_NUM = re.compile(r"\bprize\s*(?:#|number\s*)?(\d+)\b", re.IGNORECASE)

_DONATION_WORDS = ("donat", "donor")
_BID_COUNT_WORDS = ("how many bid", "number of bid", "total bid", "count bid", "num bid")
_BID_ACTIVITY_WORDS = (
    "bid",
    "bids",
    "bidding",
    "gesture",
    "gestures",
    "were made",
    "happened",
    "activity",
)
_TIME_RANGE_WORDS = (
    "weekend",
    "yesterday",
    "today",
    "last week",
    "last day",
    "last ",
    " days",
    "during",
    "since",
    "between",
    "while i was",
    "when i was",
    "not around",
    "away",
)
_STATUS_RECAP_WORDS = (
    "update me",
    "catch me up",
    "catch up",
    "what happened",
    "what's happening",
    "whats happening",
    "status",
    "brief me",
    "fill me in",
    "while i was",
    "when i was",
    "not around",
    "was away",
    "missed",
)

_ACTIVATION_WINDOW_PHRASES = (
    "activation window",
    "param window",
    "parameter window",
    "before activation",
    "before the round activated",
    "before the round opens",
    "during activation",
    "config changed during",
    "configuration changed during",
    "configuration of the round",
    "config of the round",
    "admin events",
    "admin config",
    "round configuration",
    "round config",
    "delay before round activation",
    "delaydurationbeforeroundactivation",
)

_CONFIG_CHANGE_WORDS = (
    "config",
    "configuration",
    "configured",
    "admin",
    "parameter",
    "param ",
    "setter",
    "changed during",
    "changes during",
)


def needs_activation_window_config(question: str) -> bool:
    """True when the user asks about admin config changes during a round activation window."""
    q = question.lower().strip()
    if any(p in q for p in _ACTIVATION_WINDOW_PHRASES):
        return True
    if "activation" in q and any(w in q for w in _CONFIG_CHANGE_WORDS):
        return True
    if "round" in q and "changed" in q and any(w in q for w in ("config", "configuration", "admin", "parameter")):
        return True
    return False


def extract_round_num(question: str) -> int | None:
    for pattern in (_ROUND_NUM, _PRIZE_NUM):
        match = pattern.search(question)
        if match:
            return int(match.group(1))
    return None


def wants_current_round(question: str) -> bool:
    q = question.lower()
    return any(
        phrase in q
        for phrase in (
            "this round",
            "current round",
            "the round",
            "right now",
            "currently",
            "at the moment",
            "today",
        )
    )


def mentions_erc20_tokens(question: str) -> bool:
    q = question.lower()
    return any(w in q for w in ("token", "erc20", "erc-20", "cst"))


def mentions_nft(question: str) -> bool:
    q = question.lower()
    return "nft" in q


def mentions_eth_donation(question: str) -> bool:
    q = question.lower()
    return "eth" in q and any(w in q for w in _DONATION_WORDS)


def mentions_bid_activity(question: str) -> bool:
    q = question.lower()
    return any(w in q for w in _BID_ACTIVITY_WORDS)


def mentions_time_range(question: str) -> bool:
    q = question.lower()
    return any(w in q for w in _TIME_RANGE_WORDS)


def needs_time_range_bids(question: str) -> bool:
    """True when the user asks for bid counts over a calendar/time window."""
    q = question.lower().strip()
    if not mentions_bid_activity(q):
        return False
    if any(p in q for p in _BID_COUNT_WORDS):
        return mentions_time_range(q) or "weekend" in q or "yesterday" in q or "today" in q
    if mentions_time_range(q) and any(w in q for w in ("bid", "bids", "gesture", "gestures")):
        return True
    if "how many" in q and mentions_bid_activity(q) and mentions_time_range(q):
        return True
    return False


def needs_status_recap(question: str) -> bool:
    """True when the user wants a catch-up / current-state summary."""
    q = question.lower().strip()
    return any(phrase in q for phrase in _STATUS_RECAP_WORDS)


def needs_backend_api(question: str) -> bool:
    """True when indexed backend REST data is needed (donations, bid counts, round stats)."""
    q = question.lower().strip()
    if needs_activation_window_config(question):
        return True
    if needs_round_end_time(question):
        return True
    if needs_staking_stats(question):
        return True
    if needs_champions_state(question):
        return True
    if needs_time_range_bids(q):
        return True
    if needs_status_recap(q):
        return True
    if any(w in q for w in _DONATION_WORDS):
        return True
    if any(p in q for p in _BID_COUNT_WORDS):
        return True
    if "round" in q and any(w in q for w in ("stats", "statistics", "summary", "info")):
        return True
    if "how many" in q and any(w in q for w in ("token", "nft", "donation", "bid", "staker", "staking")):
        return True
    if "number of" in q and any(w in q for w in ("token", "nft", "donation", "bid", "staker", "staking")):
        return True
    if mentions_time_range(q) and mentions_bid_activity(q):
        return True
    return False


def donation_fetch_paths(round_num: int, question: str) -> list[tuple[str, str]]:
    """Return (label, path_suffix) pairs relative to API base."""
    paths: list[tuple[str, str]] = []
    generic = not mentions_nft(question) and not mentions_erc20_tokens(question)

    if mentions_erc20_tokens(question) or generic:
        paths.append(
            (f"erc20_summarized_round_{round_num}", f"donations/erc20/by_round/summarized/{round_num}")
        )
        paths.append((f"erc20_all_round_{round_num}", f"donations/erc20/by_round/all/{round_num}"))

    if mentions_nft(question) or generic:
        paths.append((f"nft_round_{round_num}", f"donations/nft/by_round/{round_num}"))

    if mentions_eth_donation(question) or generic:
        paths.append((f"eth_simple_round_{round_num}", f"donations/eth/simple/by_round/{round_num}"))
        paths.append(
            (f"eth_with_info_round_{round_num}", f"donations/eth/with_info/by_round/{round_num}")
        )

    return paths


def bid_count_fetch_path(round_num: int) -> tuple[str, str]:
    return (f"bids_round_{round_num}", f"bid/list/by_round/{round_num}/0/0/1000")


def bidding_frequency_path(init_ts: int, fin_ts: int, interval_secs: int) -> tuple[str, str]:
    path = f"statistics/bidding/frequency/{init_ts}/{fin_ts}/{interval_secs}"
    return ("bidding_frequency", path)
