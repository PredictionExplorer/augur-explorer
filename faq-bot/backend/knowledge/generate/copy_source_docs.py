"""Copy and normalize existing markdown documentation."""
from __future__ import annotations

from pathlib import Path

from knowledge.config import (
    CURSOR_VREF_PATH,
    DOCS_BEGINNER_DIR,
    DOCS_EXPERT_DIR,
    DOCS_SOURCES_DIR,
    REPO_NAMES,
)
from knowledge.generate.utils import read_text, relative, write_text


SOURCE_MANIFEST = [
    # cursor-vref local docs (high value)
    (CURSOR_VREF_PATH / "README.md", "cursor-vref/README.md"),
    (CURSOR_VREF_PATH / "rwcg/README.md", "cursor-vref/rwcg/README.md"),
    (CURSOR_VREF_PATH / "rwcg/docs/BACKEND.md", "cursor-vref/rwcg/docs/BACKEND.md"),
    (CURSOR_VREF_PATH / "docs/BIDDING_FLOW_AUDIT.md", "cursor-vref/docs/BIDDING_FLOW_AUDIT.md"),
    (CURSOR_VREF_PATH / "docs/CLAIM_MAIN_PRIZE_AUDIT.md", "cursor-vref/docs/CLAIM_MAIN_PRIZE_AUDIT.md"),
    # smart contracts
    (REPO_NAMES["smart_contracts"] / "README.md", "smart_contracts/README.md"),
    (REPO_NAMES["smart_contracts"] / "docs/cosmic-signature-game-prizes.md", "smart_contracts/docs/cosmic-signature-game-prizes.md"),
    (REPO_NAMES["smart_contracts"] / "docs/cosmic-signature-contracts-functional-requirements.md", "smart_contracts/docs/cosmic-signature-contracts-functional-requirements.md"),
    (REPO_NAMES["smart_contracts"] / "docs/QUICKSTART.md", "smart_contracts/docs/QUICKSTART.md"),
    (REPO_NAMES["smart_contracts"] / "tasks/docs/Cosmic-Signature-Contracts-Deployment-And-Registration.md", "smart_contracts/tasks/docs/deployment-and-registration.md"),
    # backend
    (REPO_NAMES["backend_api"] / "rwcg/docs/BACKEND.md", "backend_api/rwcg/docs/BACKEND.md"),
    (REPO_NAMES["backend_api"] / "rwcg/README.md", "backend_api/rwcg/README.md"),
    # frontend
    (REPO_NAMES["frontend"] / "README.md", "frontend/README.md"),
    (REPO_NAMES["frontend"] / "QUICK_START.md", "frontend/QUICK_START.md"),
]


def run() -> None:
    copied = []
    missing = []
    for src, dest_name in SOURCE_MANIFEST:
        dest = DOCS_SOURCES_DIR / dest_name
        if not src.exists():
            missing.append(str(src))
            continue
        content = read_text(src)
        header = f"<!-- source: {relative(src, src.anchor)} -->\n\n" if False else f"<!-- source: {src} -->\n\n"
        write_text(dest, header + content)
        copied.append({"source": str(src), "dest": dest_name})

    write_text(
        DOCS_SOURCES_DIR / "_manifest.txt",
        "\n".join(f"{item['dest']} <= {item['source']}" for item in copied)
        + ("\n\nMissing:\n" + "\n".join(missing) if missing else ""),
    )


def generate_beginner_overview() -> None:
    overview = """# Cosmic Signature — Project Overview (Beginner)

Cosmic Signature is a blockchain-based bidding game on Arbitrum. Players place bids with ETH or CST tokens
during timed rounds. The last bidder before the timer expires can win the main prize.

## What you need
- A Web3 wallet (e.g. MetaMask)
- ETH on Arbitrum for gas and bidding, and/or CST tokens
- Visit the website and connect your wallet

## Core activities
1. **Bid** — go to `/game/play` during an active round
2. **Learn** — read `/game/how-it-works` and `/game/prizes`
3. **Stake NFTs** — visit `/stake` to earn rewards
4. **Claim winnings** — check `/account/winnings`
5. **Browse NFTs** — visit `/gallery` and `/account/nfts`

## Tokens and NFTs
- **ETH** — used for bidding and gas
- **CST (Cosmic Signature Token)** — ERC20 used for CST bids and rewards
- **Cosmic Signature NFT** — project NFT collection
- **RandomWalk NFT** — optional discount when bidding

## Where smart contracts live
Production contracts are in the Cosmic Signature GitHub repository under `contracts/production/`.
Live addresses are served by the backend dashboard API and shown on `/contracts`.

## More help
- How to bid: see `02-how-to-bid.md`
- Prizes: see `03-prizes-and-winning.md`
- Wallet setup: see `04-wallet-and-network.md`
"""
    write_text(DOCS_BEGINNER_DIR / "00-project-overview.md", overview)


def generate_beginner_bidding_doc() -> None:
    bidding = """# How to Place a Bid (Beginner)

## Before you start
1. Install and set up MetaMask (or another Web3 wallet)
2. Connect the wallet using **Connect Wallet** in the site header
3. Switch to the correct Arbitrum network shown by the app
4. Ensure you have enough ETH for gas and your bid

## Steps
1. Open **`/game/play`**
2. Review the current bid price (updates in real time)
3. Choose bid type:
   - **ETH** — pay with Ether
   - **CST** — pay with Cosmic Signature Token
4. Optional: enable **RandomWalk NFT** discount if you own an eligible NFT
5. Optional: add a short bid message or attach a donation
6. Click **Place Bid**
7. Confirm the transaction in your wallet
8. Wait for on-chain confirmation

## What happens next
- You become the **last bidder** and a countdown timer resets
- If nobody outbids you before the timer ends, you can win the main prize
- You may receive CST rewards for bidding (minted by the game contract)

## Important notes
- The **first bid in a round must be ETH**
- Bids happen **on-chain only** — the backend does not submit bids for you
- Donations require approving **PrizesWallet** as the spender for tokens/NFTs

## Troubleshooting
- **Insufficient funds** — add ETH or CST
- **Wrong network** — switch network in the wallet when prompted
- **Round inactive** — wait until the round opens
- **Transaction reverted** — check wallet error; often allowance, timing, or bid amount
"""
    write_text(DOCS_BEGINNER_DIR / "02-how-to-bid.md", bidding)


def generate_beginner_prizes_doc() -> None:
    prizes = """# Prizes and Winning (Beginner)

## Main prize
The main prize goes to the **last bidder** before the round timer expires.

## Other prizes
Cosmic Signature also distributes secondary prizes and rewards. See `/game/prizes` for:
- Prize categories
- Reward types (ETH, CST, NFTs)
- How winners are selected

## Claiming
1. Go to **`/account/winnings`**
2. Review available prizes
3. Click claim buttons for eligible rewards
4. Confirm wallet transactions

## Donations with bids
You can optionally donate ERC20 tokens or NFTs alongside a bid. These go to **PrizesWallet**.

## Expert reference
For contract-level prize logic, see source doc `smart_contracts/docs/cosmic-signature-game-prizes.md`
and the claim flow audit in `cursor-vref/docs/CLAIM_MAIN_PRIZE_AUDIT.md`.
"""
    write_text(DOCS_BEGINNER_DIR / "03-prizes-and-winning.md", prizes)


def generate_beginner_wallet_doc() -> None:
    wallet = """# Wallet and Network Setup (Beginner)

## Supported wallet
MetaMask or any injected EIP-1193 wallet supported by the site.

## Steps
1. Install MetaMask from https://metamask.io
2. Create or import a wallet
3. Add Arbitrum network if not present (the site may prompt you to switch)
4. Fund the wallet with ETH on Arbitrum
5. Connect on the website via **Connect Wallet**

## Contract addresses
The frontend loads live contract addresses from the backend dashboard API (`ContractAddrs`).
Visit **`/contracts`** to inspect addresses for the current deployment.

## Security tips
- Never share your seed phrase
- Verify you are on the official website
- Review transaction details before signing
"""
    write_text(DOCS_BEGINNER_DIR / "04-wallet-and-network.md", wallet)


def generate_expert_architecture_doc() -> None:
    arch = """# System Architecture (Expert)

## Components
1. **Smart contracts** — Cosmic Signature production contracts on Arbitrum
2. **Backend (`rwcg`)** — Go API + PostgreSQL + ETL indexers
3. **Frontend** — Next.js app consuming backend APIs and wallets for on-chain txs

## Data flow
- Users submit bids **directly on-chain** via the game contract
- Backend ETL indexes chain events into PostgreSQL
- Frontend reads game state, prices, and addresses from backend REST APIs
- Contract addresses are exposed via dashboard `ContractAddrs`

## Repositories
- Smart contracts: PredictionExplorer/Cosmic-Signature
- Backend: PredictionExplorer/augur-explorer (`rwcg/`)
- Frontend: PredictionExplorer/cosmic-front-alternate

## Key backend areas
- `rwcg/websrv/api/cosmicgame/` — REST handlers
- `rwcg/dbs/cosmicgame/` — DB access
- `rwcg/etl/cosmicgame/` — chain event ingestion
- `rwcg/contracts/cosmicgame/` — Go contract bindings

## Local monorepo note
Some operational audit docs live in `cursor-vref/docs/` and may be ahead of the public GitHub mirrors.
"""
    write_text(DOCS_EXPERT_DIR / "00-architecture.md", arch)


def generate_expert_deployment_doc() -> None:
    deployment = """# Deployment and Networks (Expert)

## Supported networks
- Hardhat local
- Arbitrum Sepolia (testnet)
- Arbitrum One (mainnet)

## Deployment process
1. Configure `tasks/config/deploy-cosmic-signature-contracts-config-<network>.json`
2. Set Hardhat vars for private key and Etherscan/Arbiscan API key (live networks)
3. Run `tasks/runners/run-deploy-cosmic-signature-contracts-<network>.bash`
4. Read `tasks/output/deploy-cosmic-signature-contracts-report-<network>.json`
5. Register/verify on Arbiscan for testnet/mainnet

## Instantiated contract roles
See generated facts file `deployed-contract-roles.json` for the 10 logical roles
(game, token, NFTs, wallets, DAO) expected per deployment.

## Live addresses in production UI
Frontend resolves addresses from backend dashboard statistics, not from static frontend env vars.

## OpenZeppelin upgrades
Preserve `.openzeppelin/` deployment metadata when upgrading proxied contracts.
"""
    write_text(DOCS_EXPERT_DIR / "02-deployment-and-networks.md", deployment)
