"""Fetch live indexed data from the Cosmic Game backend REST API."""
from __future__ import annotations

import json
import logging
import os
from datetime import datetime, timezone
from typing import Any

import aiohttp

from knowledge.config import KNOWLEDGE_BASE
from live.api_detector import (
    bid_count_fetch_path,
    bidding_frequency_path,
    donation_fetch_paths,
    extract_round_num,
    needs_backend_api,
    needs_status_recap,
    needs_time_range_bids,
)
from live.api_fetch_result import ApiFetchResult
from live.detector import needs_champions_state, needs_round_end_time, needs_staking_stats
from live.time_range import parse_time_range

logger = logging.getLogger(__name__)


class CosmicGameApiClient:
    """Read-only HTTP client for /api/cosmicgame/* endpoints."""

    def __init__(self, base_url: str | None = None, timeout_seconds: int | None = None):
        self.base_url = (base_url or os.getenv("FAQ_BOT_API_URL") or os.getenv("COSMIC_GAME_API_URL") or "").strip()
        self.timeout_seconds = timeout_seconds or int(os.getenv("FAQ_BOT_API_TIMEOUT_SECONDS", "20"))
        self._apply_kb_defaults()

    def _apply_kb_defaults(self) -> None:
        if self.base_url:
            return
        env_facts = KNOWLEDGE_BASE / "facts" / "network-environment.json"
        try:
            if env_facts.exists():
                data = json.loads(env_facts.read_text(encoding="utf-8"))
                mainnet = data.get("networks", {}).get("arbitrum_one_mainnet", {})
                self.base_url = (mainnet.get("frontend_env") or {}).get("NEXT_PUBLIC_API_URL", "").strip()
        except Exception as exc:
            logger.warning("Failed to load KB API URL default: %s", exc)

    @property
    def is_configured(self) -> bool:
        return bool(self.base_url)

    def config_status(self) -> dict[str, str | bool]:
        return {
            "configured": self.is_configured,
            "base_url": self.base_url or "",
        }

    def _url(self, path: str) -> str:
        base = self.base_url.rstrip("/") + "/"
        return base + path.lstrip("/")

    async def _get_json(self, path: str) -> tuple[dict[str, Any] | None, str | None]:
        url = self._url(path)
        timeout = aiohttp.ClientTimeout(total=self.timeout_seconds)
        try:
            async with aiohttp.ClientSession(timeout=timeout) as session:
                async with session.get(url) as resp:
                    text = await resp.text()
                    if resp.status != 200:
                        return None, f"HTTP {resp.status} for {path}: {text[:300]}"
                    try:
                        return json.loads(text), None
                    except json.JSONDecodeError as exc:
                        return None, f"Invalid JSON from {path}: {exc}"
        except aiohttp.ClientError as exc:
            return None, f"Request failed for {path}: {exc}"
        except TimeoutError:
            return None, f"Timeout after {self.timeout_seconds}s for {path}"

    @staticmethod
    def _format_utc_timestamp(ts: int | None) -> str:
        if ts is None or ts <= 0:
            return "unknown"
        return datetime.fromtimestamp(int(ts), tz=timezone.utc).strftime("%Y-%m-%d %H:%M:%S UTC")

    @staticmethod
    def _summarize_special_winners(data: dict[str, Any]) -> list[str]:
        lines = [
            "SPECIAL WINNERS / CHAMPIONS (live contract read via bid/current_special_winners)",
        ]

        chrono_addr = data.get("ChronoWarriorAddress") or ""
        endurance_addr = data.get("EnduranceChampionAddress") or ""
        chrono_dur = data.get("ChronoWarriorDuration")
        endurance_dur = data.get("EnduranceChampionDuration")
        chrono_live = data.get("ChronoWarriorIsLive")
        ec_start = int(data.get("EnduranceChampionStartTimeStamp") or 0)
        prev_ec_dur = int(data.get("PrevEnduranceChampionDuration") or 0)
        chrono_segment_start = ec_start + prev_ec_dur if ec_start else 0
        source_ts = int(data.get("SourceBlockTimeStamp") or 0)

        lines.append(f"ChronoWarriorAddress: {chrono_addr}")
        lines.append(f"ChronoWarriorDurationSeconds: {chrono_dur}")
        lines.append(f"ChronoWarriorIsLive: {chrono_live}")
        lines.append(f"EnduranceChampionAddress: {endurance_addr}")
        lines.append(f"EnduranceChampionDurationSeconds: {endurance_dur}")
        lines.append(f"EnduranceChampionStartTimeStamp: {ec_start}")
        lines.append(f"EnduranceChampionStart (UTC): {CosmicGameApiClient._format_utc_timestamp(ec_start)}")
        lines.append(f"PrevEnduranceChampionDurationSeconds: {prev_ec_dur}")
        lines.append(f"ChronoSegmentStartTimeStamp: {chrono_segment_start}")
        lines.append(
            f"ChronoSegmentStart (UTC): {CosmicGameApiClient._format_utc_timestamp(chrono_segment_start)}"
        )
        lines.append(f"SourceBlockTimeStamp (UTC): {CosmicGameApiClient._format_utc_timestamp(source_ts)}")
        lines.append(f"LastBidderAddress: {data.get('LastBidderAddress', '')}")
        lines.append(f"LastBidderLastBidTime (UTC): {CosmicGameApiClient._format_utc_timestamp(data.get('LastBidderLastBidTime'))}")

        if chrono_live and chrono_segment_start:
            lines.append(
                "ChronoWarriorStatus: active (ChronoWarriorIsLive=true). "
                f"The current chrono segment started at {CosmicGameApiClient._format_utc_timestamp(chrono_segment_start)}. "
                "They became Chrono Warrior when this segment's hold exceeded the stored Chrono Warrior "
                "duration threshold (after that segment start)."
            )
        elif chrono_addr and not chrono_live:
            lines.append(
                "ChronoWarriorStatus: address is set but ChronoWarriorIsLive=false "
                "(not currently accumulating Chrono Warrior time)."
            )

        lines.append(
            "Use ChronoWarriorAddress for 'who is Chrono Warrior'. "
            "For 'when did they become Chrono Warrior': if ChronoWarriorIsLive is true, use "
            "ChronoSegmentStart (UTC) and explain they crossed the threshold after that time; "
            "do not invent exact timestamps beyond what is provided."
        )
        return lines

    @staticmethod
    def _parse_unix_timestamp(value: Any) -> int | None:
        if value is None:
            return None
        if isinstance(value, bool):
            return None
        if isinstance(value, (int, float)):
            ts = int(value)
            return ts if ts > 0 else None
        if isinstance(value, str):
            text = value.strip()
            if not text:
                return None
            try:
                ts = int(text, 0)
            except ValueError:
                return None
            return ts if ts > 0 else None
        return None

    @staticmethod
    def _format_duration(seconds: int | None) -> str:
        if seconds is None or seconds < 0:
            return "unknown"
        days, rem = divmod(int(seconds), 86400)
        hours, rem = divmod(rem, 3600)
        minutes, secs = divmod(rem, 60)
        parts: list[str] = []
        if days:
            parts.append(f"{days}d")
        if hours:
            parts.append(f"{hours}h")
        if minutes:
            parts.append(f"{minutes}m")
        if secs or not parts:
            parts.append(f"{secs}s")
        return " ".join(parts)

    def _summarize_round_end(
        self,
        *,
        prize_time_data: dict[str, Any] | None,
        until_data: dict[str, Any] | None,
        current_data: dict[str, Any] | None,
        dashboard: dict[str, Any] | None,
    ) -> list[str]:
        lines = [
            "ROUND END / COUNTDOWN (live contract read — authoritative for when the cycle can finalize)",
        ]

        now_ts = self._parse_unix_timestamp((current_data or {}).get("CurrentTimeStamp"))
        prize_ts = self._parse_unix_timestamp((prize_time_data or {}).get("CurRoundPrizeTime"))
        seconds_until = self._parse_unix_timestamp((until_data or {}).get("TimeUntilPrize"))

        cur_round = (dashboard or {}).get("CurRoundNum")
        if cur_round is not None:
            lines.append(f"Current round (CurRoundNum): {cur_round}")

        stats = (dashboard or {}).get("CurRoundStats") or {}
        activation_ts = self._parse_unix_timestamp(stats.get("ActivationTime"))
        last_bidder = (dashboard or {}).get("LastBidderAddr") or ""

        if activation_ts and now_ts and activation_ts > now_ts:
            lines.append("ROUND_STATUS: pre-activation (cycle not open yet)")
            lines.append(f"ROUND_OPENS_TIMESTAMP: {activation_ts}")
            lines.append(f"ROUND_OPENS_UTC: {self._format_utc_timestamp(activation_ts)}")
            if now_ts:
                lines.append(
                    f"SECONDS_UNTIL_ROUND_OPENS: {max(0, activation_ts - now_ts)} "
                    f"({self._format_duration(max(0, activation_ts - now_ts))})"
                )
            lines.append(
                "Answer pre-activation questions with ROUND_OPENS_UTC. "
                "The main-prize countdown starts after the first gesture."
            )
            return lines

        zero_addr = "0x0000000000000000000000000000000000000000"
        if last_bidder.lower() == zero_addr or not prize_ts:
            lines.append("ROUND_STATUS: active, waiting for first gesture")
            lines.append(
                "The Signature Allocation countdown has not started yet — it begins after the first gesture. "
                "There is no fixed round end time until then."
            )
            return lines

        if prize_ts:
            lines.append(f"ROUND_END_TIMESTAMP (MainPrizeTime): {prize_ts}")
            lines.append(f"ROUND_END_UTC: {self._format_utc_timestamp(prize_ts)}")
        if seconds_until is not None:
            lines.append(f"SECONDS_UNTIL_ROUND_END: {seconds_until}")
            lines.append(f"TIME_UNTIL_ROUND_END: {self._format_duration(seconds_until)}")
        if now_ts:
            lines.append(f"CURRENT_TIME_UTC: {self._format_utc_timestamp(now_ts)}")

        lines.append(
            "Use ROUND_END_UTC as the projected cycle end. Each new gesture extends the timer, "
            "so this is the current deadline if no one gestures again."
        )
        return lines

    async def fetch_round_end_state(self) -> ApiFetchResult:
        if not self.is_configured:
            return ApiFetchResult(error="FAQ_BOT_API_URL is not set and no KB API default found")

        sources: list[str] = []
        errors: list[str] = []
        prize_time_data: dict[str, Any] | None = None
        until_data: dict[str, Any] | None = None
        current_data: dict[str, Any] | None = None
        dashboard: dict[str, Any] | None = None

        async def fetch(path: str) -> dict[str, Any] | None:
            sources.append(f"live:api:{path}")
            data, err = await self._get_json(path)
            if err:
                errors.append(f"{path}: {err}")
                return None
            return data

        prize_time_data = await fetch("rounds/current/time")
        until_data = await fetch("time/until_prize")
        current_data = await fetch("time/current")
        dashboard = await fetch("statistics/dashboard")

        if not any((prize_time_data, until_data, current_data, dashboard)):
            return ApiFetchResult(error="; ".join(errors) or "round end API fetch failed", sources=sources)

        parts = [
            f"Fetched at (UTC): {datetime.now(timezone.utc).isoformat()}",
            f"API base: {self.base_url}",
            *self._summarize_round_end(
                prize_time_data=prize_time_data,
                until_data=until_data,
                current_data=current_data,
                dashboard=dashboard,
            ),
        ]
        if errors:
            parts.append("Partial fetch errors: " + "; ".join(errors))

        block = (
            "LIVE ROUND END STATE (indexed on-chain read — authoritative for cycle end time/countdown)\n"
            + "\n".join(parts)
        )
        return ApiFetchResult(block=block, sources=sources)

    @staticmethod
    def _summarize_staking_stats(dashboard: dict[str, Any]) -> list[str]:
        main = dashboard.get("MainStats") or {}
        cst = main.get("StakeStatisticsCST") or {}
        rwalk = main.get("StakeStatisticsRWalk") or {}

        active_cst = cst.get("NumActiveStakers")
        active_rwalk = rwalk.get("NumActiveStakers")
        lines = [
            "LIVE STAKING STATS (statistics/dashboard MainStats — authoritative for staker counts)",
            f"NumActiveStakersCST (currently staking Cosmic Signature NFTs): {active_cst}",
            f"NumActiveStakersRandomWalk (currently staking RandomWalk NFTs): {active_rwalk}",
            f"NumUniqueStakersCST (addresses with CST stake history): {main.get('NumUniqueStakersCST')}",
            f"NumUniqueStakersRandomWalk (addresses with RandomWalk stake history): {main.get('NumUniqueStakersRWalk')}",
            f"NumUniqueStakersBoth (addresses that staked both collections): {main.get('NumUniqueStakersBoth')}",
            f"TotalTokensStakedCST: {cst.get('TotalTokensStaked')}",
            f"TotalTokensStakedRandomWalk: {rwalk.get('TotalTokensStaked')}",
        ]
        lines.append(
            "For 'how many stakers currently', report NumActiveStakersCST and "
            "NumActiveStakersRandomWalk (one address may appear in both). "
            "Use NumUniqueStakers* only when the user asks about all-time/historical unique stakers."
        )
        return lines

    async def fetch_staking_stats(self) -> ApiFetchResult:
        if not self.is_configured:
            return ApiFetchResult(error="FAQ_BOT_API_URL is not set and no KB API default found")

        path = "statistics/dashboard"
        data, err = await self._get_json(path)
        if err or not data:
            return ApiFetchResult(error=err or "empty response", sources=[f"live:api:{path}"])

        parts = [
            f"Fetched at (UTC): {datetime.now(timezone.utc).isoformat()}",
            f"API base: {self.base_url}",
            *self._summarize_staking_stats(data),
        ]
        block = (
            "LIVE STAKING STATE (indexed dashboard read — authoritative for staker counts)\n"
            + "\n".join(parts)
        )
        return ApiFetchResult(block=block, sources=[f"live:api:{path}"])

    async def fetch_champions_state(self) -> ApiFetchResult:
        if not self.is_configured:
            return ApiFetchResult(error="FAQ_BOT_API_URL is not set and no KB API default found")

        path = "bid/current_special_winners"
        data, err = await self._get_json(path)
        if err or not data:
            return ApiFetchResult(error=err or "empty response", sources=[f"live:api:{path}"])

        if data.get("status") == 0:
            return ApiFetchResult(
                error=str(data.get("error") or "API error"),
                sources=[f"live:api:{path}"],
            )

        parts = [
            f"Fetched at (UTC): {datetime.now(timezone.utc).isoformat()}",
            f"API base: {self.base_url}",
            *self._summarize_special_winners(data),
        ]
        block = (
            "LIVE CHAMPIONS STATE (indexed on-chain read — authoritative for Chrono Warrior, "
            "Endurance Champion, and special winner addresses/timers)\n"
            + "\n".join(parts)
        )
        return ApiFetchResult(block=block, sources=[f"live:api:{path}"])

    @staticmethod
    def _summarize_dashboard(data: dict[str, Any], round_num: int | None) -> list[str]:
        lines: list[str] = []
        cur_round = data.get("CurRoundNum")
        lines.append(f"Current round (CurRoundNum): {cur_round}")
        stats = data.get("CurRoundStats") or {}
        if stats and (round_num is None or stats.get("RoundNum") == round_num):
            lines.append(f"Round {stats.get('RoundNum')} TotalBids (round total): {stats.get('TotalBids')}")
            lines.append(f"Round {stats.get('RoundNum')} NumERC20Donations: {stats.get('NumERC20Donations')}")
            lines.append(f"Round {stats.get('RoundNum')} TotalDonatedNFTs: {stats.get('TotalDonatedNFTs')}")
            lines.append(
                f"Round {stats.get('RoundNum')} TotalDonatedCount (ETH donations): "
                f"{stats.get('TotalDonatedCount')} "
                f"({stats.get('TotalDonatedAmountEth')} ETH total)"
            )
        return lines

    @staticmethod
    def _sum_frequency_bids(data: dict[str, Any]) -> int:
        buckets = data.get("FrequencyHistory") or []
        total = 0
        for bucket in buckets:
            try:
                total += int(bucket.get("NumBids", 0))
            except (TypeError, ValueError):
                continue
        return total

    @staticmethod
    def _frequency_interval_secs(init_ts: int, fin_ts: int) -> int:
        span = max(0, fin_ts - init_ts)
        if span <= 3 * 86400:
            return 3600
        return 86400

    @staticmethod
    def _summarize_payload(label: str, path: str, data: dict[str, Any]) -> list[str]:
        lines = [f"[{label}] GET /api/cosmicgame/{path}"]

        if "DonationsERC20ByRoundSummarized" in data:
            items = data["DonationsERC20ByRoundSummarized"] or []
            lines.append(f"ERC20 donation summary rows: {len(items)}")
        if "DonationsERC20ByRoundAll" in data:
            items = data["DonationsERC20ByRoundAll"] or []
            lines.append(f"ERC20 donation events (all rows): {len(items)}")
        if "NFTDonations" in data:
            items = data["NFTDonations"] or []
            lines.append(f"NFT donations: {len(items)}")
        if "DirectCGDonations" in data:
            items = data["DirectCGDonations"] or []
            lines.append(f"Direct ETH donations to game: {len(items)}")
        if "CosmicGameDonations" in data:
            items = data["CosmicGameDonations"] or []
            lines.append(f"ETH donations with info: {len(items)}")
        if "TotalRows" in data:
            lines.append(f"Total bid rows for round {data.get('RoundNum')}: {data.get('TotalRows')}")
        if "RoundInfo" in data:
            info = data["RoundInfo"] or {}
            round_stats = info.get("RoundStats") or info
            lines.append(f"RoundInfo RoundNum: {round_stats.get('RoundNum')}")
            lines.append(f"RoundInfo TotalBids: {round_stats.get('TotalBids')}")
            lines.append(f"RoundInfo NumERC20Donations: {round_stats.get('NumERC20Donations')}")
            lines.append(f"RoundInfo TotalDonatedNFTs: {round_stats.get('TotalDonatedNFTs')}")

        if data.get("status") == 0 and data.get("error"):
            lines.append(f"API error: {data.get('error')}")

        compact = json.dumps(data, indent=2)
        if len(compact) > 6000:
            compact = compact[:6000] + "\n...(truncated)"
        lines.append(compact)
        return lines

    async def fetch_for_question(self, question: str) -> ApiFetchResult:
        if not needs_backend_api(question):
            return ApiFetchResult()

        if not self.is_configured:
            return ApiFetchResult(error="FAQ_BOT_API_URL is not set and no KB API default found")

        q_lower = question.lower()
        wants_round_end = needs_round_end_time(question)
        wants_staking = needs_staking_stats(question)
        wants_champions = needs_champions_state(question)
        wants_time_bids = needs_time_range_bids(question)
        wants_recap = needs_status_recap(question)
        wants_bids = any(
            p in q_lower for p in ("how many bid", "number of bid", "total bid", "count bid", "num bid")
        ) or wants_time_bids
        wants_donations = any(w in q_lower for w in ("donat", "donor")) or (
            "how many" in q_lower and any(w in q_lower for w in ("token", "nft"))
        )

        time_parse = parse_time_range(question) if wants_time_bids else None
        if time_parse and time_parse.needs_clarification:
            return ApiFetchResult(
                clarification_answer=time_parse.clarification,
                sources=["live:api:time-range-clarification"],
            )

        parts: list[str] = [
            f"Fetched at (UTC): {datetime.now(timezone.utc).isoformat()}",
            f"API base: {self.base_url}",
        ]
        sources: list[str] = []
        errors: list[str] = []

        async def fetch_one(label: str, path: str) -> dict[str, Any] | None:
            sources.append(f"live:api:{path}")
            data, err = await self._get_json(path)
            if err:
                errors.append(f"{path}: {err}")
                parts.append(f"[{label}] ERROR: {err}")
                return None
            return data

        if wants_staking:
            staking = await self.fetch_staking_stats()
            sources.extend(staking.sources)
            if staking.block:
                parts.append(staking.block)
            elif staking.error:
                errors.append(staking.error)
                parts.append(f"[staking] ERROR: {staking.error}")

        if wants_round_end:
            round_end = await self.fetch_round_end_state()
            sources.extend(round_end.sources)
            if round_end.block:
                parts.append(round_end.block)
            elif round_end.error:
                errors.append(round_end.error)
                parts.append(f"[round_end] ERROR: {round_end.error}")

        if wants_champions:
            champs = await self.fetch_champions_state()
            sources.extend(champs.sources)
            if champs.block:
                parts.append(champs.block)
            elif champs.error:
                errors.append(champs.error)
                parts.append(f"[champions] ERROR: {champs.error}")

        round_num = extract_round_num(question)
        dashboard = await fetch_one("dashboard", "statistics/dashboard")
        if dashboard and round_num is None:
            round_num = dashboard.get("CurRoundNum")

        if time_parse and time_parse.resolved:
            assert time_parse.init_ts is not None and time_parse.fin_ts is not None
            interval = self._frequency_interval_secs(time_parse.init_ts, time_parse.fin_ts)
            label, path = bidding_frequency_path(time_parse.init_ts, time_parse.fin_ts, interval)
            freq_data = await fetch_one(label, path)
            if freq_data:
                period_bids = self._sum_frequency_bids(freq_data)
                parts.append(f"PERIOD_BID_COUNT ({time_parse.period_label}): {period_bids}")
                parts.append(
                    "Use PERIOD_BID_COUNT for time-window bid questions. "
                    "Do not substitute the current-round TotalBids for a calendar period."
                )
                bucket_count = len(freq_data.get("FrequencyHistory") or [])
                parts.append(f"Frequency buckets returned: {bucket_count} (interval {interval}s)")

        if wants_recap or dashboard:
            if round_num is not None:
                parts.append(f"Resolved round: {round_num}")
            if dashboard:
                parts.extend(self._summarize_dashboard(dashboard, round_num))
                parts.append(
                    "CURRENT_ROUND_STATUS: use dashboard fields above for catch-up / status questions."
                )

        if round_num is None and not (time_parse and time_parse.resolved) and not wants_champions and not wants_round_end and not wants_staking:
            if errors and dashboard is None:
                return ApiFetchResult(error="; ".join(errors), sources=sources)
            if not wants_time_bids:
                return ApiFetchResult(error="Could not determine round number", sources=sources)

        if round_num is not None and (wants_donations or wants_recap):
            for label, path in donation_fetch_paths(round_num, question):
                data = await fetch_one(label, path)
                if data:
                    parts.extend(self._summarize_payload(label, path, data))

        if round_num is not None and wants_bids and not (time_parse and time_parse.resolved):
            label, path = bid_count_fetch_path(round_num)
            data = await fetch_one(label, path)
            if data:
                parts.extend(self._summarize_payload(label, path, data))

        if (
            round_num is not None
            and not wants_donations
            and not wants_bids
            and not wants_recap
            and not (time_parse and time_parse.resolved)
        ):
            label, path = f"round_info_{round_num}", f"rounds/info/{round_num}"
            data = await fetch_one(label, path)
            if data:
                parts.extend(self._summarize_payload(label, path, data))

        if errors and dashboard is None and not (time_parse and time_parse.resolved) and not wants_champions:
            return ApiFetchResult(error="; ".join(errors), sources=sources)

        if len(parts) <= 2 and not wants_champions and not wants_round_end and not wants_staking:
            return ApiFetchResult(error="No applicable API data fetched", sources=sources)

        block = (
            "LIVE BACKEND API DATA (indexed Cosmic Game REST API — authoritative for "
            "donation counts, bid totals by time period, and round statistics)\n"
            + "\n".join(parts)
        )
        return ApiFetchResult(block=block, sources=sources)
