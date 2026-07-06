"""Extract contract inventory and deployment facts."""
from __future__ import annotations

import re
from pathlib import Path

from knowledge.config import FACTS_DIR, REPO_NAMES
from knowledge.generate.utils import iter_files, relative, write_json, write_text


PRODUCTION_CONTRACTS_DIR = REPO_NAMES["smart_contracts"] / "contracts" / "production"


def _contract_name(path: Path) -> str:
    return path.stem


def extract_contract_inventory() -> dict:
    repo = REPO_NAMES["smart_contracts"]
    all_sol = iter_files(repo, [".sol"])
    production = iter_files(PRODUCTION_CONTRACTS_DIR, [".sol"]) if PRODUCTION_CONTRACTS_DIR.exists() else []
    tests = [p for p in all_sol if "/tests/" in str(p).replace("\\", "/")]
    interfaces = [p for p in all_sol if "/interfaces/" in str(p).replace("\\", "/")]
    libraries = [p for p in all_sol if "/libraries/" in str(p).replace("\\", "/")]

    production_names = sorted(_contract_name(p) for p in production)
    interface_names = sorted(_contract_name(p) for p in interfaces)

    return {
        "repository": "smart_contracts",
        "totals": {
            "all_solidity_files": len(all_sol),
            "production_contract_files": len(production),
            "test_contract_files": len(tests),
            "interface_files": len(interfaces),
            "library_files": len(libraries),
        },
        "production_contracts": [
            {
                "name": _contract_name(p),
                "path": relative(p, repo),
                "category": (
                    "interface"
                    if "/interfaces/" in str(p)
                    else "library"
                    if "/libraries/" in str(p)
                    else "contract"
                ),
            }
            for p in production
        ],
        "core_contract_names": production_names,
        "interface_names": interface_names,
        "supported_networks": ["hardhat_local", "arbitrum_sepolia", "arbitrum_one"],
        "deployment_report_pattern": "tasks/output/deploy-cosmic-signature-contracts-report-<network>.json",
        "deployment_guide": "tasks/docs/Cosmic-Signature-Contracts-Deployment-And-Registration.md",
        "notes": [
            "Production Cosmic Signature functionality is implemented by contracts under contracts/production/.",
            "Deployed addresses are written to deploy-cosmic-signature-contracts-report-<network>.json after deployment.",
            "The frontend reads live addresses from the backend dashboard API ContractAddrs field.",
            "Arbitrum One and Arbitrum Sepolia are the supported live/test networks per deployment docs.",
        ],
    }


def extract_deployed_contract_roles() -> dict:
    """Map logical contract roles to frontend/backend names."""
    return {
        "instantiated_roles_on_network": [
            {"role": "CosmicSignatureGame", "frontend_key": "COSMIC_GAME", "backend_key": "CosmicGameAddr"},
            {"role": "CosmicSignatureToken", "frontend_key": "COSMIC_SIGNATURE_TOKEN", "backend_key": "CosmicTokenAddr"},
            {"role": "CosmicSignatureNft", "frontend_key": "COSMIC_SIGNATURE_NFT", "backend_key": "CosmicSignatureAddr"},
            {"role": "RandomWalkNFT", "frontend_key": "RANDOM_WALK_NFT", "backend_key": "RandomWalkAddr"},
            {"role": "PrizesWallet", "frontend_key": "PRIZES_WALLET", "backend_key": "PrizesWalletAddr"},
            {"role": "StakingWalletCosmicSignatureNft", "frontend_key": "STAKING_WALLET_CST", "backend_key": "StakingWalletCSTAddr"},
            {"role": "StakingWalletRandomWalkNft", "frontend_key": "STAKING_WALLET_RWLK", "backend_key": "StakingWalletRWalkAddr"},
            {"role": "CharityWallet", "frontend_key": "CHARITY_WALLET", "backend_key": "CharityWalletAddr"},
            {"role": "MarketingWallet", "frontend_key": "MARKETING_WALLET", "backend_key": "MarketingWalletAddr"},
            {"role": "CosmicSignatureDao", "frontend_key": "COSMIC_DAO", "backend_key": "CosmicDaoAddr"},
        ],
        "expected_instantiated_count": 10,
        "address_source": "Backend dashboard API statistics/dashboard -> ContractAddrs",
        "deployment_artifacts": "tasks/output/deploy-cosmic-signature-contracts-report-arbitrum_one.json (when deployed)",
    }


def generate_expert_contracts_doc(inventory: dict, roles: dict) -> str:
    totals = inventory["totals"]
    lines = [
        "# Smart Contracts Inventory (Expert)",
        "",
        "## Summary",
        f"- Total Solidity files in repo: **{totals['all_solidity_files']}**",
        f"- Production contract files: **{totals['production_contract_files']}**",
        f"- Interface files: **{totals['interface_files']}**",
        f"- Library files: **{totals['library_files']}**",
        f"- Test/mock contract files: **{totals['test_contract_files']}**",
        "",
        "## Production contracts",
    ]
    for item in inventory["production_contracts"]:
        lines.append(f"- `{item['name']}` ({item['category']}) — `{item['path']}`")
    lines += [
        "",
        "## Instantiated on-chain roles",
        f"The live system expects **{roles['expected_instantiated_count']}** deployed contract roles per network:",
    ]
    for role in roles["instantiated_roles_on_network"]:
        lines.append(
            f"- **{role['role']}** (frontend `{role['frontend_key']}`, backend `{role['backend_key']}`)"
        )
    lines += [
        "",
        "## Networks",
        "- Hardhat local development network",
        "- Arbitrum Sepolia (testnet)",
        "- Arbitrum One (mainnet)",
        "",
        "## Deployment",
        "- Guide: `tasks/docs/Cosmic-Signature-Contracts-Deployment-And-Registration.md`",
        "- Report JSON: `tasks/output/deploy-cosmic-signature-contracts-report-<network>.json`",
        "- Frontend resolves addresses from backend dashboard, not hardcoded env vars.",
    ]
    return "\n".join(lines)


def run() -> None:
    from knowledge.config import DOCS_EXPERT_DIR

    inventory = extract_contract_inventory()
    roles = extract_deployed_contract_roles()
    write_json(FACTS_DIR / "contracts-inventory.json", inventory)
    write_json(FACTS_DIR / "deployed-contract-roles.json", roles)
    write_text(
        DOCS_EXPERT_DIR / "01-smart-contracts-inventory.md",
        generate_expert_contracts_doc(inventory, roles),
    )
