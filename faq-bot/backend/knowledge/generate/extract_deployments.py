"""Extract deployed contract addresses from bundled repo seed files (deterministic)."""
from __future__ import annotations

import json
import os
import re
import shutil
from pathlib import Path

from knowledge.config import (
    BACKEND_ROOT,
    CURSOR_VREF_PATH,
    DOCS_EXPERT_DIR,
    DOCS_SOURCES_DIR,
    FACTS_DIR,
    KNOWLEDGE_BASE,
)
from knowledge.generate.utils import read_text, write_json, write_text

_ADDR_LINE = re.compile(
    r"^(.+?)\s+address:\s+(0x[a-fA-F0-9]{40})\s*$",
    re.MULTILINE,
)
_ENV_LINE = re.compile(
    r"^(NEXT_PUBLIC_[A-Z0-9_]+):\s*(.+?)\s*$",
    re.MULTILINE,
)

DEPLOYMENT_SEEDS_DIR = BACKEND_ROOT / "knowledge" / "deployments"
MAINNET_NETWORK_KEY = "arbitrum_one_mainnet"
GAME_PROXY_KEY = "CosmicSignatureGame proxy"


def _deployments_source_dir() -> Path:
    return DOCS_SOURCES_DIR / "deployments"


def _allow_external_deployment_dir() -> bool:
    raw = os.getenv("FAQ_BOT_ALLOW_EXTERNAL_DEPLOYMENT", "").strip().lower()
    return raw in {"1", "true", "yes", "on"}


def _external_deployment_dir() -> Path | None:
    if not _allow_external_deployment_dir():
        return None
    raw = os.getenv("DEPLOYMENT_FACTS_DIR", "").strip()
    if raw:
        path = Path(raw)
        if path.is_dir():
            return path
    return None


def _is_environment_file(path: Path) -> bool:
    return path.name.endswith("-environment.txt")


def _find_canonical_address_files() -> list[Path]:
    """Address seeds shipped in the faq-bot repo — same on laptop and production."""
    if not DEPLOYMENT_SEEDS_DIR.is_dir():
        return []
    return sorted(
        p for p in DEPLOYMENT_SEEDS_DIR.glob("*.txt") if not _is_environment_file(p)
    )


def _find_canonical_environment_files() -> list[Path]:
    if not DEPLOYMENT_SEEDS_DIR.is_dir():
        return []
    return sorted(DEPLOYMENT_SEEDS_DIR.glob("*-environment.txt"))


def _find_optional_external_address_files() -> list[Path]:
    ext = _external_deployment_dir()
    if not ext:
        return []
    return sorted(p for p in ext.glob("*.txt") if not _is_environment_file(p))


def _parse_deployment_text(content: str) -> dict[str, str]:
    return {name.strip(): addr for name, addr in _ADDR_LINE.findall(content)}


def _parse_environment_text(content: str) -> dict[str, str]:
    return {key: value.strip().strip('"').strip("'") for key, value in _ENV_LINE.findall(content)}


def _sync_deployment_seeds() -> None:
    """Copy bundled deployment seed files into the knowledge base (audit trail)."""
    if not DEPLOYMENT_SEEDS_DIR.exists():
        return
    dest = _deployments_source_dir()
    dest.mkdir(parents=True, exist_ok=True)
    for path in sorted(DEPLOYMENT_SEEDS_DIR.glob("*.txt")):
        shutil.copy2(path, dest / path.name)


def _network_key_from_stem(stem: str) -> str:
    stem = stem.replace("deployment-", "").removesuffix("-environment")
    mapping = {
        "ArbitrumMainnet": MAINNET_NETWORK_KEY,
        "arbitrum-mainnet": MAINNET_NETWORK_KEY,
        "arbitrum_mainnet": MAINNET_NETWORK_KEY,
    }
    return mapping.get(stem, stem.replace("-", "_").lower())


def _network_name_from_file(path: Path) -> str:
    return _network_key_from_stem(path.stem)


def _load_address_networks() -> tuple[dict[str, dict], list[str]]:
    networks: dict[str, dict] = {}
    sources_used: list[str] = []

    for path in _find_canonical_address_files():
        addresses = _parse_deployment_text(read_text(path))
        if not addresses:
            continue
        network = _network_name_from_file(path)
        rel = f"knowledge/deployments/{path.name}"
        networks[network] = {
            "source_file": rel,
            "source_type": "repo_seed",
            "network_label": "Arbitrum One (mainnet)" if "mainnet" in network else network,
            "address_count": len(addresses),
            "addresses": addresses,
        }
        sources_used.append(rel)

    if networks:
        return networks, sources_used

    for path in _find_optional_external_address_files():
        addresses = _parse_deployment_text(read_text(path))
        if not addresses:
            continue
        network = _network_name_from_file(path)
        networks[network] = {
            "source_file": str(path),
            "source_type": "external",
            "network_label": "Arbitrum One (mainnet)" if "mainnet" in network else network,
            "address_count": len(addresses),
            "addresses": addresses,
        }
        sources_used.append(str(path))

    return networks, sources_used


def _load_environment_networks() -> dict[str, dict]:
    environments: dict[str, dict] = {}
    seed_files = _find_canonical_environment_files()
    ext = _external_deployment_dir()
    if ext:
        seed_files.extend(sorted(ext.glob("*-environment.txt")))

    seen: set[str] = set()
    for path in seed_files:
        key = path.name
        if key in seen:
            continue
        seen.add(key)
        variables = _parse_environment_text(read_text(path))
        if not variables:
            continue
        network = _network_key_from_stem(path.stem)
        rel = f"knowledge/deployments/{path.name}"
        if path.parent != DEPLOYMENT_SEEDS_DIR:
            rel = str(path)
        chain_id_match = re.search(r"^Chain ID:\s*(\d+)\s*$", read_text(path), re.MULTILINE)
        environments[network] = {
            "source_file": rel,
            "source_type": "repo_seed" if path.parent == DEPLOYMENT_SEEDS_DIR else "external",
            "network_label": "Arbitrum One (mainnet)" if "mainnet" in network else network,
            "chain_id": int(chain_id_match.group(1)) if chain_id_match else None,
            "frontend_env": variables,
        }
    return environments


def _validate_address_networks(networks: dict[str, dict]) -> None:
    mainnet = networks.get(MAINNET_NETWORK_KEY) or {}
    addresses = mainnet.get("addresses") or {}
    proxy = addresses.get(GAME_PROXY_KEY) or addresses.get("CosmicSignatureGame")
    if proxy:
        return
    seed_dir = DEPLOYMENT_SEEDS_DIR
    raise RuntimeError(
        "deployed-addresses generation failed: missing canonical mainnet game proxy address. "
        f"Add `{GAME_PROXY_KEY} address: 0x...` to a *.txt file under {seed_dir} "
        f"(excluding *-environment.txt). Found networks: {sorted(networks.keys()) or '(none)'}"
    )


def run() -> None:
    _deployments_source_dir().mkdir(parents=True, exist_ok=True)
    _sync_deployment_seeds()

    networks, address_sources = _load_address_networks()
    environments = _load_environment_networks()
    _validate_address_networks(networks)

    payload = {
        "networks": networks,
        "canonical_seed_dir": str(DEPLOYMENT_SEEDS_DIR),
        "address_seed_files": [p.name for p in _find_canonical_address_files()],
        "environment_seed_files": [p.name for p in _find_canonical_environment_files()],
        "address_sources_used": address_sources,
        "external_deployment_enabled": _allow_external_deployment_dir(),
        "external_deployment_dir": str(_external_deployment_dir() or ""),
        "knowledge_base_deployments_dir": str(_deployments_source_dir()),
    }
    write_json(FACTS_DIR / "deployed-addresses.json", payload)
    write_json(
        FACTS_DIR / "network-environment.json",
        {
            "networks": environments,
            "canonical_seed_dir": str(DEPLOYMENT_SEEDS_DIR),
            "environment_seed_files": [p.name for p in _find_canonical_environment_files()],
            "external_deployment_enabled": _allow_external_deployment_dir(),
            "external_deployment_dir": str(_external_deployment_dir() or ""),
            "knowledge_base_deployments_dir": str(_deployments_source_dir()),
        },
    )

    lines = [
        "# Deployed Contract Addresses (Expert)",
        "",
        "Canonical deployment records indexed for the FAQ bot.",
        "",
        f"Source of truth: `{DEPLOYMENT_SEEDS_DIR}` (copied into the knowledge base on each generate run).",
        "",
    ]
    for network, info in sorted(networks.items()):
        lines.append(f"## {info.get('network_label', network)}")
        lines.append(f"- Source: `{info['source_file']}` ({info['source_type']})")
        lines.append(f"- Address count: **{len(info['addresses'])}**")
        lines.append("")
        for name, addr in sorted(info["addresses"].items()):
            lines.append(f"- **{name}**: `{addr}`")
        lines.append("")
    write_text(DOCS_EXPERT_DIR / "07-deployed-addresses.md", "\n".join(lines))

    env_lines = [
        "# Network Environment (Expert)",
        "",
        "Production frontend environment variables for live Cosmic Signature deployments.",
        "",
        f"Source of truth: `{DEPLOYMENT_SEEDS_DIR}` (*-environment.txt seeds).",
        "",
    ]
    if not environments:
        env_lines += [
            "No environment files found.",
            "",
            f"Add `*-environment.txt` under `{DEPLOYMENT_SEEDS_DIR}`.",
            "Format: `NEXT_PUBLIC_NETWORK: mainnet`",
        ]
    else:
        for network, info in sorted(environments.items()):
            env_lines.append(f"## {info.get('network_label', network)}")
            env_lines.append(f"- Source: `{info['source_file']}`")
            if info.get("chain_id"):
                env_lines.append(f"- Chain ID: **{info['chain_id']}**")
            env_lines.append("")
            env_lines.append("### Frontend environment")
            env_lines.append("```bash")
            for key, value in info["frontend_env"].items():
                env_lines.append(f'export {key}="{value}"')
            env_lines.append("```")
            env_lines.append("")
        env_lines += [
            "## Live on-chain state (cginfo)",
            "",
            "For **current** bid prices, round status, timers, and bidders, the FAQ bot runs:",
            "",
            "```bash",
            "RPC_URL=<rpc> ./cginfo <CosmicSignatureGame_proxy>",
            "```",
            "",
            f"Default binary: `{CURSOR_VREF_PATH / 'rwcg' / 'etl' / 'cosmicgame' / 'scripts' / 'cginfo'}`",
            "",
            "Game proxy address comes from `facts/deployed-addresses.json` (repo seed). "
            "Set `COSMIC_GAME_ADDR` only as an optional runtime override.",
            "",
        ]
    write_text(DOCS_EXPERT_DIR / "09-network-environment.md", "\n".join(env_lines))


if __name__ == "__main__":
    run()
