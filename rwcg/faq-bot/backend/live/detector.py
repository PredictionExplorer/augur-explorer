"""Detect questions that need live on-chain contract state."""
from __future__ import annotations

import re

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

_CHAMPION_ROLES = (
    "chrono warrior",
    "endurance champion",
    "last eth bidder",
    "last cst bidder",
    "special winner",
    "special winners",
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
    "who is the chrono",
    "who is chrono",
    "who is the endurance",
    "who is endurance",
)

_WHO_IS = re.compile(r"\bwho(?:'s|\s+is)\b", re.IGNORECASE)
_WHEN_BECAME = re.compile(
    r"\bwhen\s+did\b|\bwhen\s+.*\s+became\b|\bbecame\s+(?:the\s+)?(?:chrono|endurance)",
    re.IGNORECASE,
)

_ROUND_END_PHRASES = (
    "when will the round end",
    "when does the round end",
    "when will this round end",
    "when will the cycle end",
    "when does the cycle end",
    "when is the round over",
    "when does the round finish",
    "round end date",
    "round end time",
    "cycle end date",
    "cycle end time",
    "time until round end",
    "time until the round ends",
    "how long until the round",
    "how long until round",
    "how long left in the round",
    "countdown to round end",
    "main prize time",
)

_WHEN_ROUND_END = re.compile(
    r"\bwhen\b.*\b(?:round|cycle)\b.*\b(?:end|finish|over|expire)\b|"
    r"\b(?:round|cycle)\b.*\b(?:end|finish)\b.*\bwhen\b|"
    r"\bhow\s+long\b.*\b(?:round|cycle)\b.*\b(?:end|left|remain)\b",
    re.IGNORECASE,
)


def normalize_question(question: str) -> str:
    """Unify compound protocol role names for keyword matching."""
    q = question.lower()
    q = q.replace("chrono-warrior", "chrono warrior")
    q = q.replace("chronowarrior", "chrono warrior")
    q = q.replace("endurance-champion", "endurance champion")
    q = q.replace("endurancechampion", "endurance champion")
    return q


def needs_round_end_time(question: str) -> bool:
    """True when the user asks when the current round/cycle ends."""
    q = normalize_question(question)
    if any(p in q for p in _ROUND_END_PHRASES):
        return True
    if _WHEN_ROUND_END.search(q):
        return True
    if "end" in q and "round" in q and any(w in q for w in ("countdown", "timer", "when", "how long")):
        return True
    if "end" in q and "cycle" in q and any(w in q for w in ("countdown", "timer", "when", "how long")):
        return True
    return False


def needs_backend_url_info(question: str) -> bool:
    """True when the user asks for the Cosmic Game REST API / backend base URL."""
    q = normalize_question(question)
    if any(t in q for t in ("next_public_api", "faq_bot_api", "cosmic_game_api")):
        return True
    if any(p in q for p in ("backend url", "api url", "api base", "rest api url", "backend endpoint")):
        return True
    if ("url" in q or "endpoint" in q or "base url" in q) and any(
        w in q for w in ("backend", "api", "rest", "server", "rwcg", "cosmicgame", "cosmic game")
    ):
        return True
    if any(p in q for p in ("what is the api", "where is the api", "url for the backend", "url of the backend")):
        return True
    return False


def needs_staking_stats(question: str) -> bool:
    """True when the user asks for current/historical NFT staker counts."""
    q = normalize_question(question)
    if not any(w in q for w in ("staker", "stakers", "staking", "staked")):
        return False
    if any(
        w in q
        for w in (
            "how many",
            "number of",
            "count",
            "total",
            "currently",
            "current",
            "right now",
            "live",
            "today",
            "at the moment",
        )
    ):
        return True
    if any(p in q for p in ("unique stakers", "active stakers", "staking stats")):
        return True
    return False


def needs_champions_state(question: str) -> bool:
    """True when the user asks about Chrono Warrior, Endurance Champion, or special winners."""
    q = normalize_question(question)
    if any(role in q for role in _CHAMPION_ROLES):
        return True
    if _WHO_IS.search(q) and any(w in q for w in ("warrior", "champion", "bidder")):
        return True
    if _WHEN_BECAME.search(q) and any(w in q for w in ("warrior", "champion")):
        return True
    return False


def needs_live_state(question: str) -> bool:
    q = normalize_question(question)
    if needs_round_end_time(question):
        return True
    if needs_champions_state(question):
        return True
    if any(p in q for p in _PHRASES):
        return True
    has_live = any(w in q for w in _LIVE_WORDS)
    has_state = any(w in q for w in _STATE_WORDS)
    return has_live and has_state
