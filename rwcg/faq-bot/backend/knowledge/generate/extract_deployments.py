"""Extract deployed contract addresses from JSON reports and deployment text files."""
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
    REPO_NAMES,
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


def _deployments_source_dir() -> Path:
    return DOCS_SOURCES_DIR / "deployments"


def _external_deployment_dir() -> Path | None:
    raw = os.getenv("DEPLOYMENT_FACTS_DIR", "").strip()
    if raw:
        path = Path(raw)
        if path.is_dir():
            return path
    return None


def _find_json_reports() -> list[Path]:
    repo = REPO_NAMES["smart_contracts"]
    reports: list[Path] = []
    for base in (repo / "tasks" / "output", repo / "output"):
        if base.exists():
            reports.extend(sorted(base.glob("deploy-cosmic-signature-contracts-report-*.json")))
    return reports


def _find_text_deployments() -> list[Path]:
    paths: list[Path] = []
    kb_dir = _deployments_source_dir()
    if kb_dir.exists():
        paths.extend(sorted(kb_dir.glob("*.txt")))
    ext = _external_deployment_dir()
    if ext:
        paths.extend(sorted(ext.glob("*.txt")))
    # de-dupe by resolved path
    seen: set[str] = set()
    unique: list[Path] = []
    for p in paths:
        key = str(p.resolve())
        if key not in seen:
            seen.add(key)
            unique.append(p)
    return unique


def _parse_deployment_text(content: str) -> dict[str, str]:
    return {name.strip(): addr for name, addr in _ADDR_LINE.findall(content)}


def _parse_environment_text(content: str) -> dict[str, str]:
    return {key: value.strip().strip('"').strip("'") for key, value in _ENV_LINE.findall(content)}


def _sync_deployment_seeds() -> None:
    """Copy bundled deployment seed files into the knowledge base."""
    if not DEPLOYMENT_SEEDS_DIR.exists():
        return
    dest = _deployments_source_dir()
    dest.mkdir(parents=True, exist_ok=True)
    for path in sorted(DEPLOYMENT_SEEDS_DIR.glob("*.txt")):
        shutil.copy2(path, dest / path.name)


def _find_environment_files() -> list[Path]:
    paths: list[Path] = []
    kb_dir = _deployments_source_dir()
    if kb_dir.exists():
        paths.extend(sorted(kb_dir.glob("*-environment.txt")))
    ext = _external_deployment_dir()
    if ext:
        paths.extend(sorted(ext.glob("*-environment.txt")))
    seen: set[str] = set()
    unique: list[Path] = []
    for p in paths:
        key = str(p.resolve())
        if key not in seen:
            seen.add(key)
            unique.append(p)
    return unique


def _network_key_from_stem(stem: str) -> str:
    stem = stem.replace("deployment-", "").removesuffix("-environment")
    mapping = {
        "ArbitrumMainnet": "arbitrum_one_mainnet",
        "arbitrum-mainnet": "arbitrum_one_mainnet",
        "arbitrum_mainnet": "arbitrum_one_mainnet",
    }
    return mapping.get(stem, stem.replace("-", "_").lower())


def _network_name_from_file(path: Path) -> str:
    return _network_key_from_stem(path.stem)


def run() -> None:
    _deployments_source_dir().mkdir(parents=True, exist_ok=True)
    _sync_deployment_seeds()

    networks: dict[str, dict] = {}
    environments: dict[str, dict] = {}

    for path in _find_json_reports():
        network = path.stem.replace("deploy-cosmic-signature-contracts-report-", "")
        try:
            data = json.loads(path.read_text(encoding="utf-8"))
        except Exception:
            continue
        if isinstance(data, dict):
            networks[network] = {
                "source_file": str(path.name),
                "source_type": "json",
                "addresses": data,
            }

    for path in _find_text_deployments():
        content = read_text(path)
        addresses = _parse_deployment_text(content)
        if not addresses:
            continue
        network = _network_name_from_file(path)
        rel = path
        try:
            rel = path.relative_to(KNOWLEDGE_BASE)
        except ValueError:
            pass
        networks[network] = {
            "source_file": str(rel),
            "source_type": "text",
            "network_label": "Arbitrum One (mainnet)" if "mainnet" in network else network,
            "address_count": len(addresses),
            "addresses": addresses,
        }

    for path in _find_environment_files():
        content = read_text(path)
        variables = _parse_environment_text(content)
        if not variables:
            continue
        network = _network_key_from_stem(path.stem)
        rel = path
        try:
            rel = path.relative_to(KNOWLEDGE_BASE)
        except ValueError:
            pass
        chain_id_match = re.search(r"^Chain ID:\s*(\d+)\s*$", content, re.MULTILINE)
        environments[network] = {
            "source_file": str(rel),
            "source_type": "text",
            "network_label": "Arbitrum One (mainnet)" if "mainnet" in network else network,
            "chain_id": int(chain_id_match.group(1)) if chain_id_match else None,
            "frontend_env": variables,
        }

    payload = {
        "networks": networks,
        "json_report_files_found": len(_find_json_reports()),
        "text_deployment_files_found": len(_find_text_deployments()),
        "knowledge_base_deployments_dir": str(_deployments_source_dir()),
        "external_deployment_dir": str(_external_deployment_dir() or ""),
    }
    write_json(FACTS_DIR / "deployed-addresses.json", payload)
    write_json(
        FACTS_DIR / "network-environment.json",
        {
            "networks": environments,
            "environment_files_found": len(_find_environment_files()),
            "knowledge_base_deployments_dir": str(_deployments_source_dir()),
            "external_deployment_dir": str(_external_deployment_dir() or ""),
        },
    )

    lines = [
        "# Deployed Contract Addresses (Expert)",
        "",
        "Canonical deployment records indexed for the FAQ bot.",
        "",
    ]
    if not networks:
        lines += [
            "No deployment files found.",
            "",
            f"Add `*.txt` files to `{_deployments_source_dir()}` or set `DEPLOYMENT_FACTS_DIR`.",
            "Format: `ContractName address: 0x...`",
        ]
    else:
        for network, info in sorted(networks.items()):
            lines.append(f"## {info.get('network_label', network)}")
            lines.append(f"- Source: `{info['source_file']}` ({info['source_type']})")
            lines.append(f"- Address count: **{len(info['addresses'])}**")
            lines.append("")
            for name, addr in info["addresses"].items():
                lines.append(f"- **{name}**: `{addr}`")
            lines.append("")
    write_text(DOCS_EXPERT_DIR / "07-deployed-addresses.md", "\n".join(lines))

    env_lines = [
        "# Network Environment (Expert)",
        "",
        "Production frontend environment variables for live Cosmic Signature deployments.",
        "",
    ]
    if not environments:
        env_lines += [
            "No environment files found.",
            "",
            f"Add `*-environment.txt` files to `{_deployments_source_dir()}`.",
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
            "Set `RPC_URL` or `FAQ_BOT_RPC_URL` in the backend environment. "
            "If unset, the bot uses `NEXT_PUBLIC_RPC_URL` from this file.",
            "",
        ]
    write_text(DOCS_EXPERT_DIR / "09-network-environment.md", "\n".join(env_lines))


if __name__ == "__main__":
    run()
