#!/usr/bin/env python3
"""Generate the full Cosmic Signature FAQ knowledge base."""
from __future__ import annotations

import argparse
import sys
from pathlib import Path

BACKEND_ROOT = Path(__file__).resolve().parents[2]
if str(BACKEND_ROOT) not in sys.path:
    sys.path.insert(0, str(BACKEND_ROOT))

from knowledge.config import ensure_output_dirs, FACTS_DIR, KNOWLEDGE_BASE
from knowledge.generate import (
    copy_source_docs,
    extract_abis,
    extract_api,
    extract_contracts,
    extract_deployments,
    extract_routes,
    extract_ui_pages,
    sync_repos,
)
from knowledge.generate.utils import write_text


def main() -> int:
    parser = argparse.ArgumentParser(description="Generate FAQ knowledge base")
    parser.add_argument("--skip-source-copy", action="store_true")
    parser.add_argument(
        "--skip-repo-sync",
        action="store_true",
        help="Do not clone/update GitHub repos into backend/data/repos (offline or pre-populated cache)",
    )
    args = parser.parse_args()

    ensure_output_dirs()
    if not args.skip_repo_sync:
        print("Syncing upstream GitHub repos → backend/data/repos/ ...")
        sync_repos.sync_repositories()
        print("  ✓ repository cache")
    print("Generating facts and docs...")

    extract_contracts.run()
    print("  ✓ contracts inventory")

    extract_routes.run()
    print("  ✓ frontend routes")

    extract_api.run()
    print("  ✓ backend API")

    extract_deployments.run()
    print("  ✓ deployment addresses (repo seeds)")

    deployed = FACTS_DIR / "deployed-addresses.json"
    if deployed.is_file():
        import json

        data = json.loads(deployed.read_text(encoding="utf-8"))
        mainnet = (data.get("networks") or {}).get("arbitrum_one_mainnet") or {}
        proxy = (mainnet.get("addresses") or {}).get("CosmicSignatureGame proxy")
        if proxy:
            print(f"  ✓ mainnet CosmicSignatureGame proxy: {proxy}")

    extract_abis.run()
    print("  ✓ contract ABIs")

    if not args.skip_source_copy:
        copy_source_docs.run()
        copy_source_docs.generate_beginner_overview()
        copy_source_docs.generate_beginner_bidding_doc()
        copy_source_docs.generate_beginner_prizes_doc()
        copy_source_docs.generate_beginner_wallet_doc()
        copy_source_docs.generate_expert_architecture_doc()
        copy_source_docs.generate_expert_deployment_doc()
        print("  ✓ source docs + curated summaries")

    extract_ui_pages.run()
    print("  ✓ UI page extracts")

    write_text(
        KNOWLEDGE_BASE / "README.md",
        "# Cosmic Signature FAQ Knowledge Base\n\n"
        f"Path: `{KNOWLEDGE_BASE}`\n\n"
        "Re-run generators: `python -m knowledge.generate.run_all` from `backend/`.\n\n"
        "See `knowledge/GUIDELINES.md` in the faq-bot repo for LLM-assisted regeneration.\n",
    )
    print(f"Done. Output: {KNOWLEDGE_BASE}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
