"""Parse natural-language time ranges for bid-count questions (UTC).

When a phrase is ambiguous we return needs_clarification=True instead of guessing.
"""
from __future__ import annotations

import re
from dataclasses import dataclass
from datetime import date, datetime, time, timedelta, timezone

_AMBIGUOUS_WEEKEND = re.compile(
    r"\b(?:during|over|on)?\s*the\s+weekend\b|\bweekend\b",
    re.IGNORECASE,
)
_THIS_WEEKEND = re.compile(r"\bthis\s+weekend\b", re.IGNORECASE)
_LAST_WEEKEND = re.compile(r"\b(?:last|past|previous)\s+weekend\b", re.IGNORECASE)
_YESTERDAY = re.compile(r"\byesterday\b", re.IGNORECASE)
_TODAY = re.compile(r"\btoday\b", re.IGNORECASE)
_LAST_N_DAYS = re.compile(r"\blast\s+(\d+)\s+days?\b", re.IGNORECASE)
_AWAY_WITHOUT_DATES = re.compile(
    r"\b(?:while|when)\s+i\s+was\s+(?:away|not\s+around|offline)\b",
    re.IGNORECASE,
)

_ISO_DATE = re.compile(r"\b(20\d{2})-(\d{2})-(\d{2})\b")


@dataclass(frozen=True)
class TimeRangeParseResult:
    resolved: bool
    needs_clarification: bool
    clarification: str | None
    init_ts: int | None
    fin_ts: int | None
    period_label: str | None


def _utc_day_start(d: date) -> datetime:
    return datetime.combine(d, time.min, tzinfo=timezone.utc)


def _utc_day_end_inclusive(d: date) -> datetime:
    return datetime.combine(d, time.max, tzinfo=timezone.utc)


def _fmt_period(start: datetime, end: datetime) -> str:
    return f"{start.strftime('%Y-%m-%d %H:%M')} UTC – {end.strftime('%Y-%m-%d %H:%M')} UTC"


def _resolved(start: datetime, end: datetime, label: str) -> TimeRangeParseResult:
    if end < start:
        end = start
    return TimeRangeParseResult(
        resolved=True,
        needs_clarification=False,
        clarification=None,
        init_ts=int(start.timestamp()),
        fin_ts=int(end.timestamp()),
        period_label=label or _fmt_period(start, end),
    )


def _clarify(message: str) -> TimeRangeParseResult:
    return TimeRangeParseResult(
        resolved=False,
        needs_clarification=True,
        clarification=message,
        init_ts=None,
        fin_ts=None,
        period_label=None,
    )


def _last_completed_weekend(now: datetime) -> tuple[datetime, datetime]:
    """Most recent fully elapsed Sat 00:00 UTC – Sun 23:59:59 UTC."""
    today = now.date()
    weekday = now.weekday()  # Mon=0 … Sun=6
    # Days back to last Sunday (including today if Sunday)
    days_to_sunday = (weekday + 1) % 7
    if days_to_sunday == 0:
        # Today is Sunday — last *completed* weekend ended yesterday (Sat was day before)
        last_sunday = today - timedelta(days=7)
    else:
        last_sunday = today - timedelta(days=days_to_sunday)
    last_saturday = last_sunday - timedelta(days=1)
    start = _utc_day_start(last_saturday)
    end = _utc_day_end_inclusive(last_sunday)
    return start, end


def _this_weekend_so_far(now: datetime) -> tuple[datetime, datetime] | None:
    """Sat 00:00 UTC through now when today is Saturday or Sunday."""
    weekday = now.weekday()
    if weekday == 5:  # Saturday
        start = _utc_day_start(now.date())
        return start, now
    if weekday == 6:  # Sunday
        start = _utc_day_start(now.date() - timedelta(days=1))
        return start, now
    return None


def parse_time_range(question: str, now: datetime | None = None) -> TimeRangeParseResult | None:
    """Return a parse result when the question asks about bids in a time window, else None."""
    now = now or datetime.now(timezone.utc)
    q = question.strip()

    iso_matches = list(_ISO_DATE.finditer(q))
    if len(iso_matches) >= 2:
        d1 = date(int(iso_matches[0].group(1)), int(iso_matches[0].group(2)), int(iso_matches[0].group(3)))
        d2 = date(int(iso_matches[1].group(1)), int(iso_matches[1].group(2)), int(iso_matches[1].group(3)))
        start, end = _utc_day_start(min(d1, d2)), _utc_day_end_inclusive(max(d1, d2))
        return _resolved(start, end, _fmt_period(start, end))

    if _LAST_WEEKEND.search(q):
        start, end = _last_completed_weekend(now)
        return _resolved(start, end, f"last completed weekend ({_fmt_period(start, end)})")

    if _THIS_WEEKEND.search(q):
        window = _this_weekend_so_far(now)
        if window:
            start, end = window
            return _resolved(start, end, f"this weekend so far ({_fmt_period(start, end)})")
        return _clarify(
            "Which weekend do you mean — the one that just ended (last weekend), "
            "or the upcoming Saturday–Sunday? All times are UTC."
        )

    if _YESTERDAY.search(q):
        d = now.date() - timedelta(days=1)
        start, end = _utc_day_start(d), _utc_day_end_inclusive(d)
        return _resolved(start, end, f"yesterday ({_fmt_period(start, end)})")

    if _TODAY.search(q):
        start, end = _utc_day_start(now.date()), now
        return _resolved(start, end, f"today so far ({_fmt_period(start, end)})")

    last_days = _LAST_N_DAYS.search(q)
    if last_days:
        n = max(1, int(last_days.group(1)))
        start = now - timedelta(days=n)
        return _resolved(start, now, f"last {n} day(s) ({_fmt_period(start, now)})")

    if _AWAY_WITHOUT_DATES.search(q) and not iso_matches:
        return _clarify(
            "What dates should I use? For example: “last weekend”, “yesterday”, "
            "“May 24–25”, or “last 3 days” (UTC)."
        )

    if _AMBIGUOUS_WEEKEND.search(q) and not _LAST_WEEKEND.search(q) and not _THIS_WEEKEND.search(q):
        return _clarify(
            "Which weekend should I count bids for — last completed weekend (Sat–Sun UTC), "
            "this weekend so far, or specific dates?"
        )

    return None
