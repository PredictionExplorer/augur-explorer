"""Detect questions that need live Cosmic Game backend API data."""
from __future__ import annotations

import re

_ROUND_NUM = re.compile(r"\bround\s*(?:#|number\s*)?(\d+)\b", re.IGNORECASE)
_PRIZE_NUM = re.compile(r"\bprize\s*(?:#|number\s*)?(\d+)\b", re.IGNORECASE)

_DONATION_WORDS = ("donat", "donor")
_BID_COUNT_WORDS = ("how many bid", "number of bid", "total bid", "count bid", "num bid")
_ROUND_WORDS = ("round", "prize round")


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


def needs_backend_api(question: str) -> bool:
    """True when indexed backend REST data is needed (donations, bid counts, round stats)."""
    q = question.lower().strip()
    if any(w in q for w in _DONATION_WORDS):
        return True
    if any(p in q for p in _BID_COUNT_WORDS):
        return True
    if "round" in q and any(w in q for w in ("stats", "statistics", "summary", "info")):
        return True
    if "how many" in q and any(w in q for w in ("token", "nft", "donation", "bid")):
        return True
    if "number of" in q and any(w in q for w in ("token", "nft", "donation", "bid")):
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
