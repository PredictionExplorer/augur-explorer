"""Extract contract ABIs from frontend JSON artifacts into the knowledge base."""
from __future__ import annotations

import json
import os
import shutil
from pathlib import Path
from typing import Any

from knowledge.config import (
    DOCS_EXPERT_DIR,
    DOCS_SOURCES_DIR,
    FACTS_DIR,
    GITHUB_REPOS,
    REPOS_DIR,
    REPO_NAMES,
)
from knowledge.generate.utils import read_text, write_json, write_text

KB_CONTRACTS_DIR = DOCS_SOURCES_DIR / "frontend-contracts"


def resolve_frontend_contracts_dir() -> Path:
    """Locate frontend ABI JSONs (env override, then data/repos clone)."""
    override = os.getenv("FAQ_BOT_FRONTEND_CONTRACTS_DIR", "").strip()
    if override:
        path = Path(override)
        if path.is_dir():
            return path
        raise FileNotFoundError(
            f"FAQ_BOT_FRONTEND_CONTRACTS_DIR is set but not a directory: {path}"
        )

    default = REPO_NAMES["frontend"] / "src" / "contracts"
    if default.is_dir():
        return default

    raise FileNotFoundError(
        f"Frontend contract ABIs not found at {default}. "
        f"Run knowledge generation without --skip-repo-sync to clone "
        f"{GITHUB_REPOS['frontend']} into {REPOS_DIR / 'frontend'}, "
        "or set FAQ_BOT_FRONTEND_CONTRACTS_DIR to a directory containing *.json ABI files."
    )

# Frontend ABI filename -> deployed role / frontend env key
CONTRACT_META: dict[str, dict[str, str]] = {
    "CosmicGame.json": {
        "role": "CosmicSignatureGame",
        "frontend_key": "COSMIC_GAME",
        "description": "Main game proxy (bidding, rounds, prizes)",
    },
    "CosmicToken.json": {
        "role": "CosmicSignatureToken",
        "frontend_key": "COSMIC_SIGNATURE_TOKEN",
        "description": "CST ERC-20 token",
    },
    "CosmicSignature.json": {
        "role": "CosmicSignatureNft",
        "frontend_key": "COSMIC_SIGNATURE_NFT",
        "description": "Cosmic Signature NFT collection",
    },
    "RandomWalkNFT.json": {
        "role": "RandomWalkNFT",
        "frontend_key": "RANDOM_WALK_NFT",
        "description": "RandomWalk NFT (bid discount)",
    },
    "PrizesWallet.json": {
        "role": "PrizesWallet",
        "frontend_key": "PRIZES_WALLET",
        "description": "Donation approval target for bid donations",
    },
    "CharityWallet.json": {
        "role": "CharityWallet",
        "frontend_key": "CHARITY_WALLET",
        "description": "Charity wallet",
    },
    "EthDonations.json": {
        "role": "EthDonations",
        "frontend_key": "ETH_DONATIONS",
        "description": "ETH donations helper contract",
    },
    "StakingWalletCosmicSignatureNft.json": {
        "role": "StakingWalletCosmicSignatureNft",
        "frontend_key": "STAKING_WALLET_CST",
        "description": "CST NFT staking wallet",
    },
    "StakingWalletRandomWalkNft.json": {
        "role": "StakingWalletRandomWalkNft",
        "frontend_key": "STAKING_WALLET_RWLK",
        "description": "RandomWalk NFT staking wallet",
    },
}

BID_FUNCTION_NAMES = (
    "bidWithEth",
    "bidWithEthAndDonateNft",
    "bidWithEthAndDonateToken",
    "bidWithCst",
    "bidWithCstAndDonateNft",
    "bidWithCstAndDonateToken",
)

BID_VIEW_FUNCTION_NAMES = (
    "getNextEthBidPrice",
    "getNextEthBidPriceAdvanced",
    "nextEthBidPrice",
    "ethDutchAuctionBeginningBidPrice",
    "ethBidPriceIncreaseDivisor",
    "bidMessageLengthMaxLimit",
)


def _format_abi_type(item: dict[str, Any]) -> str:
    typ = item.get("type", "")
    if typ == "tuple":
        inner = ",".join(_format_abi_type(c) for c in item.get("components", []))
        if item.get("internalType", "").endswith("[]"):
            return f"({inner})[]"
        return f"({inner})"
    if typ.endswith("[]"):
        base = {**item, "type": typ[:-2]}
        return f"{_format_abi_type(base)}[]"
    return typ


def _function_signature(entry: dict[str, Any]) -> str:
    inputs = entry.get("inputs", [])
    types = ",".join(_format_abi_type(i) for i in inputs)
    mut = entry.get("stateMutability", "")
    suffix = f" {mut}" if mut and mut not in ("nonpayable",) else ""
    if mut == "payable":
        suffix = " payable"
    return f"{entry['name']}({types}){suffix}"


def _summarize_entry(entry: dict[str, Any]) -> dict[str, Any]:
    etype = entry.get("type")
    summary: dict[str, Any] = {
        "type": etype,
        "name": entry.get("name"),
    }
    if etype == "function":
        summary["signature"] = _function_signature(entry)
        summary["stateMutability"] = entry.get("stateMutability")
        summary["inputs"] = [
            {
                "name": i.get("name"),
                "type": _format_abi_type(i),
                "internalType": i.get("internalType"),
            }
            for i in entry.get("inputs", [])
        ]
        outputs = entry.get("outputs", [])
        if outputs:
            summary["outputs"] = [
                {
                    "name": o.get("name"),
                    "type": _format_abi_type(o),
                    "internalType": o.get("internalType"),
                }
                for o in outputs
            ]
    elif etype == "error":
        summary["signature"] = _function_signature(entry)
        summary["inputs"] = [
            {"name": i.get("name"), "type": _format_abi_type(i)}
            for i in entry.get("inputs", [])
        ]
    elif etype == "event":
        summary["signature"] = _function_signature(entry)
    return summary


def _load_abi(path: Path) -> list[dict[str, Any]]:
    raw = json.loads(read_text(path))
    if isinstance(raw, list):
        return raw
    if isinstance(raw, dict) and isinstance(raw.get("abi"), list):
        return raw["abi"]
    raise ValueError(f"Unsupported ABI format in {path}")


def _summarize_contract(filename: str, abi: list[dict[str, Any]]) -> dict[str, Any]:
    meta = CONTRACT_META.get(filename, {"role": filename.replace(".json", ""), "frontend_key": "", "description": ""})
    functions = [e for e in abi if e.get("type") == "function"]
    events = [e for e in abi if e.get("type") == "event"]
    errors = [e for e in abi if e.get("type") == "error"]

    bid_functions = [_summarize_entry(e) for e in functions if e.get("name") in BID_FUNCTION_NAMES]
    bid_views = [_summarize_entry(e) for e in functions if e.get("name") in BID_VIEW_FUNCTION_NAMES]

    return {
        **meta,
        "source_file": f"frontend/src/contracts/{filename}",
        "kb_abi_copy": f"docs/sources/frontend-contracts/{filename}",
        "counts": {
            "functions": len(functions),
            "events": len(events),
            "errors": len(errors),
        },
        "bid_functions": bid_functions,
        "bid_view_functions": bid_views,
        "custom_errors": [_summarize_entry(e) for e in errors],
        "function_names": sorted(e.get("name", "") for e in functions if e.get("name")),
    }


def _integration_notes() -> dict[str, Any]:
    return {
        "arbitrum_one_chain_id": 42161,
        "arbitrum_sepolia_chain_id": 421614,
        "cosmic_game_proxy_mainnet": "0x6a714Ae7B5b6eA520F6BCA23d2E609C4Fd5863F2",
        "simple_eth_bid": {
            "contract": "CosmicSignatureGame proxy",
            "abi_file": "CosmicGame.json",
            "function": "bidWithEth",
            "signature": "bidWithEth(int256,string) payable",
            "inputs": [
                {"name": "randomWalkNftId_", "type": "int256", "notes": "Use -1 when not using a RandomWalk NFT bid discount."},
                {"name": "message_", "type": "string", "notes": "Bid message; length limited by bidMessageLengthMaxLimit()."},
            ],
            "msg_value": "Payable. Must be >= getNextEthBidPrice() or transaction reverts with InsufficientReceivedBidAmount.",
            "randomwalk_discount": "When randomWalkNftId_ is a valid owned token id, frontend sends ~50% of the computed ETH value.",
            "price_discovery": "Call getNextEthBidPrice() on the proxy before bidding. Frontend adds a small price buffer percentage.",
            "first_bid_rule": "The first bid in a round must be ETH (not CST).",
        },
        "abi_source": "Copied from frontend/src/contracts/*.json (used by the live dApp via viem/wagmi).",
        "go_binding_source": "Monorepo Go bindings: rwcg/contracts/cosmicgame/CosmicSignatureGame.go (generated via abigen).",
    }


def generate_expert_integration_doc(summary: dict[str, Any]) -> str:
    notes = summary["integration"]
    bid = notes["simple_eth_bid"]
    cosmic = summary["contracts"].get("CosmicGame.json", {})
    bid_fn = next((f for f in cosmic.get("bid_functions", []) if f["name"] == "bidWithEth"), None)

    lines = [
        "# Contract ABIs and Integration (Expert)",
        "",
        "## Overview",
        "- ABIs are exported from the frontend dApp at `frontend/src/contracts/*.json`.",
        "- Full JSON copies live in `docs/sources/frontend-contracts/`.",
        "- Structured summaries live in `facts/contract-abis-summary.json`.",
        "",
        "## Mainnet CosmicSignatureGame proxy",
        f"- **Address (Arbitrum One):** `{notes['cosmic_game_proxy_mainnet']}`",
        f"- **Chain ID:** `{notes['arbitrum_one_chain_id']}`",
        f"- **ABI file:** `{bid['abi_file']}`",
        "",
        "## Simple ETH bid",
        f"- **Function:** `{bid['signature']}`",
        f"- **msg.value:** {bid['msg_value']}",
        f"- **randomWalkNftId_:** {bid['inputs'][0]['notes']}",
        f"- **message_:** {bid['inputs'][1]['notes']}",
        f"- **Price:** {bid['price_discovery']}",
        f"- **RandomWalk discount:** {bid['randomwalk_discount']}",
        f"- **Round rule:** {bid['first_bid_rule']}",
        "",
        "## All bid entry points (CosmicGame)",
    ]
    for fn in cosmic.get("bid_functions", []):
        lines.append(f"- `{fn['signature']}`")

    lines += [
        "",
        "## Useful view calls before bidding",
    ]
    for fn in cosmic.get("bid_view_functions", []):
        lines.append(f"- `{fn['signature']}`")

    lines += [
        "",
        "## Go example (simple ETH bid)",
        "",
        "Requires `github.com/ethereum/go-ethereum`. Load the ABI from `docs/sources/frontend-contracts/CosmicGame.json`.",
        "",
        "```go",
        "package main",
        "",
        "import (",
        '    "context"',
        '    "log"',
        '    "math/big"',
        '    "os"',
        '    "strings"',
        "",
        '    "github.com/ethereum/go-ethereum/accounts/abi"',
        '    "github.com/ethereum/go-ethereum/accounts/abi/bind"',
        '    "github.com/ethereum/go-ethereum/common"',
        '    "github.com/ethereum/go-ethereum/crypto"',
        '    "github.com/ethereum/go-ethereum/ethclient"',
        ")",
        "",
        f'const gameProxy = "{notes["cosmic_game_proxy_mainnet"]}"',
        "",
        "func main() {",
        "    ctx := context.Background()",
        '    client, err := ethclient.Dial(os.Getenv("RPC_URL"))',
        "    if err != nil { log.Fatal(err) }",
        "",
        '    pk, err := crypto.HexToECDSA(strings.TrimPrefix(os.Getenv("PRIVATE_KEY"), "0x"))',
        "    if err != nil { log.Fatal(err) }",
        "    from := crypto.PubkeyToAddress(pk.PublicKey)",
        "    gameAddr := common.HexToAddress(gameProxy)",
        "",
        '    gameABI, err := abi.JSON(strings.NewReader(os.Getenv("GAME_ABI_JSON")))',
        "    if err != nil { log.Fatal(err) }",
        "",
        "    // Read required ETH bid amount",
        '    callData, _ := gameABI.Pack("getNextEthBidPrice")',
        "    out, err := client.CallContract(ctx, bind.CallOpts{From: from, Pending: true}, gameAddr, callData)",
        "    if err != nil { log.Fatal(err) }",
        "    val, err := gameABI.Unpack(\"getNextEthBidPrice\", out)",
        "    if err != nil { log.Fatal(err) }",
        "    bidWei := val[0].(*big.Int)",
        "",
        "    // bidWithEth(int256 randomWalkNftId_, string message_) payable",
        "    randomWalkNftId := big.NewInt(-1) // no RandomWalk discount",
        '    message := "hello from go"',
        '    txData, err := gameABI.Pack("bidWithEth", randomWalkNftId, message)',
        "    if err != nil { log.Fatal(err) }",
        "",
        "    auth, err := bind.NewKeyedTransactorWithChainID(pk, big.NewInt(42161))",
        "    if err != nil { log.Fatal(err) }",
        "    auth.Value = bidWei",
        "",
        "    tx := bind.NewBoundContract(gameAddr, gameABI, client, client, client)",
        "    txn, err := tx.RawTransact(auth, txData)",
        "    if err != nil { log.Fatal(err) }",
        '    log.Println("sent:", txn.Hash().Hex())',
        "}",
        "```",
        "",
        "## Indexed contract ABIs",
    ]
    for filename, info in sorted(summary["contracts"].items()):
        lines.append(
            f"- **{info.get('role', filename)}** — `{filename}` "
            f"({info['counts']['functions']} functions, {info['counts']['errors']} errors) — {info.get('description', '')}"
        )

    if bid_fn:
        lines += [
            "",
            "## bidWithEth ABI fragment",
            "",
            "```json",
            json.dumps(bid_fn, indent=2),
            "```",
        ]

    lines += [
        "",
        "## Sources",
        "- `docs/sources/frontend-contracts/CosmicGame.json`",
        "- `facts/contract-abis-summary.json`",
        "- `facts/deployed-addresses.json`",
        "- Frontend hook: `src/hooks/usePlaceBid.ts`",
    ]
    return "\n".join(lines)


def run() -> None:
    frontend_contracts_src = resolve_frontend_contracts_dir()

    KB_CONTRACTS_DIR.mkdir(parents=True, exist_ok=True)
    contracts_summary: dict[str, Any] = {}

    for path in sorted(frontend_contracts_src.glob("*.json")):
        abi = _load_abi(path)
        contracts_summary[path.name] = _summarize_contract(path.name, abi)
        shutil.copy2(path, KB_CONTRACTS_DIR / path.name)

    summary = {
        "source_directory": str(frontend_contracts_src),
        "kb_abi_directory": str(KB_CONTRACTS_DIR),
        "contract_count": len(contracts_summary),
        "contracts": contracts_summary,
        "integration": _integration_notes(),
    }

    write_json(FACTS_DIR / "contract-abis-summary.json", summary)
    write_text(
        DOCS_EXPERT_DIR / "08-contract-abis-and-integration.md",
        generate_expert_integration_doc(summary),
    )


if __name__ == "__main__":
    run()
