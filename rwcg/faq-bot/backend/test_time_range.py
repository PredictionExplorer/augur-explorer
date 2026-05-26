"""Tests for time-range parsing and API detection."""
from __future__ import annotations

from datetime import datetime, timezone

from live.api_detector import needs_backend_api, needs_time_range_bids
from live.detector import needs_champions_state, needs_live_state, needs_round_end_time, needs_staking_stats, needs_backend_url_info, normalize_question
from live.time_range import parse_time_range


def _dt(y: int, m: int, d: int, h: int = 12) -> datetime:
    return datetime(y, m, d, h, 0, 0, tzinfo=timezone.utc)


def test_ambiguous_weekend_asks_clarification():
    result = parse_time_range(
        "how many bids were made during the weekend?",
        now=_dt(2026, 5, 25, 22),
    )
    assert result is not None
    assert result.needs_clarification
    assert "weekend" in (result.clarification or "").lower()


def test_last_weekend_resolves_on_sunday():
    result = parse_time_range("how many bids last weekend?", now=_dt(2026, 5, 25, 22))
    assert result is not None
    assert result.resolved
    assert result.init_ts is not None
    assert result.fin_ts is not None
    assert "last completed weekend" in (result.period_label or "").lower()


def test_this_weekend_resolves_on_sunday():
    result = parse_time_range("bids this weekend", now=_dt(2026, 5, 25, 18))
    assert result is not None
    assert result.resolved
    assert "this weekend" in (result.period_label or "").lower()


def test_this_weekend_monday_needs_clarification():
    result = parse_time_range("bids this weekend", now=_dt(2026, 5, 26, 10))
    assert result is not None
    assert result.needs_clarification


def test_away_without_dates_needs_clarification():
    result = parse_time_range(
        "I was not around, update me — how many bids?",
        now=_dt(2026, 5, 25, 10),
    )
    assert result is not None
    assert result.needs_clarification


def test_user_weekend_question_triggers_backend_api():
    q = "how many bids were made during the weekend? I was not around , want you to update me on the status"
    assert needs_time_range_bids(q)
    assert needs_backend_api(q)


def test_chrono_warrior_question_triggers_live_and_api():
    q = "Who is ChronoWarrior right now and when did this user became ChronoWarrior?"
    assert needs_champions_state(q)
    assert needs_live_state(q)
    assert needs_backend_api(q)


def test_round_end_question_triggers_live_and_api():
    q = "when will the round end?"
    assert needs_round_end_time(q)
    assert needs_live_state(q)
    assert needs_backend_api(q)


def test_backend_url_question_triggers_config_and_retrieval():
    q = "What is the URL for the backend?"
    assert needs_backend_url_info(q)


def test_stakers_question_triggers_backend_api():
    q = "how many stakers are there currently ?"
    assert needs_staking_stats(q)
    assert needs_backend_api(q)


def test_normalize_chronowarrior_compound():
    q = normalize_question("ChronoWarrior")
    assert "chrono warrior" in q


def test_iso_date_range_resolves():
    result = parse_time_range(
        "how many bids between 2026-05-24 and 2026-05-25?",
        now=_dt(2026, 5, 26),
    )
    assert result is not None
    assert result.resolved
