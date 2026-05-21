"""Extract deployed contract addresses from deployment report JSON files."""
from __future__ import annotations

import json
from pathlib import Path

from knowledge.config import DOCS_EXPERT_DIR, FACTS_DIR, REPO_NAMES
from knowledge.generate.utils import write_json, write_text


def _find_reports() -> list[Path]:
    repo = REPO_NAMES["smart_contracts"]
    patterns = [
        repo / "tasks" / "output",
        repo / "output",
    ]
    reports: list[Path] = []
    for base in patterns:
        if not base.exists():
            continue
        reports.extend(sorted(base.glob("deploy-cosmic-signature-contracts-report-*.json")))
    return reports


def run() -> None:
    reports = _find_reports()
    networks: dict[str, dict] = {}
    for path in reports:
        network = path.stem.replace("deploy-cosmic-signature-contracts-report-", "")
        try:
            data = json.loads(path.read_text(encoding="utf-8"))
        except Exception:
            continue
        if isinstance(data, dict):
            networks[network] = {"source_file": str(path.name), "addresses": data}

    payload = {
        "networks": networks,
        "report_files_found": len(reports),
        "note": (
            "No deployment report JSON was found in the repo cache."
            if not networks
            else "Live addresses extracted from deployment reports."
        ),
    }
    write_json(FACTS_DIR / "deployed-addresses.json", payload)

    lines = [
        "# Deployed Contract Addresses (Expert)",
        "",
        f"Deployment report files found: **{len(reports)}**",
        "",
    ]
    if not networks:
        lines += [
            "No `deploy-cosmic-signature-contracts-report-*.json` files are present in the cached repo.",
            "After deployment, copy the report JSON into the smart contracts repo or set `DEPLOYMENT_REPORTS_DIR`.",
            "The frontend/backend still expose live addresses via dashboard `ContractAddrs`.",
        ]
    else:
        for network, info in networks.items():
            lines.append(f"## {network}")
            lines.append(f"Source: `{info['source_file']}`")
            for key, val in info["addresses"].items():
                lines.append(f"- `{key}`: `{val}`")
            lines.append("")
    write_text(DOCS_EXPERT_DIR / "07-deployed-addresses.md", "\n".join(lines))
