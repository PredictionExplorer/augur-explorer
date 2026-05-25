"""Detect questions that need live on-chain contract state."""
from __future__ import annotations


_LIVE_WORDS = (
    "current",
    "right now",
    "currently",
    "at the moment",
    "live",
    " now",
    "today",
)

_STATE_WORDS = (
    "price",
    "bid",
    "round",
    "bids",
    "bidder",
    "prize",
    "countdown",
    "timer",
    "eth price",
    "cst price",
    "next bid",
    "cost to bid",
    "main prize",
    "contract balance",
    "gas price",
    "chrono warrior",
    "endurance champion",
    "dutch auction",
    "time until",
    "claim",
)

_PHRASES = (
    "what is the current",
    "what's the current",
    "how much eth",
    "how much to bid",
    "how much cst",
    "next eth bid",
    "next cst bid",
    "getnextethbidprice",
    "round status",
    "total bids",
    "unique bidders",
    "last eth bidder",
    "last cst bidder",
    "price to bid",
)


def needs_live_state(question: str) -> bool:
    q = question.lower().strip()
    if any(p in q for p in _PHRASES):
        return True
    has_live = any(w in q for w in _LIVE_WORDS)
    has_state = any(w in q for w in _STATE_WORDS)
    return has_live and has_state
