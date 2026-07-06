"""
Smart Code Understander
Parses frontend TSX/JSX files to extract UI elements, routes, and user interactions
Then generates user-friendly answers without needing an LLM
"""
import re
import logging
from typing import List, Dict, Any, Optional
from haystack import Document

logger = logging.getLogger(__name__)


class CodeUnderstander:
    """
    Understands code structure and generates user-friendly step-by-step answers
    by parsing frontend pages, extracting UI elements, and building instructions.
    Works without any LLM API key.
    """

    def __init__(self):
        self.page_route_map = {
            "game/play": {
                "name": "Play / Bidding Page",
                "url": "/game/play",
                "description": "The main game page where you can place bids",
            },
            "game/prizes": {
                "name": "Prizes Page",
                "url": "/game/prizes",
                "description": "Shows prize distribution and how prizes are awarded",
            },
            "game/leaderboard": {
                "name": "Leaderboard Page",
                "url": "/game/leaderboard",
                "description": "Shows top players and rankings",
            },
            "game/history": {
                "name": "History Page",
                "url": "/game/history",
                "description": "Shows past rounds and bid history",
            },
            "game/how-it-works": {
                "name": "How It Works Page",
                "url": "/game/how-it-works",
                "description": "Explains how the game works",
            },
            "game/statistics": {
                "name": "Statistics Page",
                "url": "/game/statistics",
                "description": "Shows game statistics and analytics",
            },
            "stake": {
                "name": "Staking Page",
                "url": "/stake",
                "description": "Where you can stake your NFTs to earn rewards",
            },
            "account/nfts": {
                "name": "My NFTs Page",
                "url": "/account/nfts",
                "description": "View and manage your Cosmic Signature NFTs",
            },
            "account/winnings": {
                "name": "Winnings Page",
                "url": "/account/winnings",
                "description": "View your prize winnings and claim rewards",
            },
            "account/activity": {
                "name": "Activity Page",
                "url": "/account/activity",
                "description": "View your bidding and activity history",
            },
            "account": {
                "name": "Account Page",
                "url": "/account",
                "description": "Your account overview",
            },
            "gallery": {
                "name": "Gallery Page",
                "url": "/gallery",
                "description": "Browse Cosmic Signature NFTs",
            },
            "contracts": {
                "name": "Contracts Page",
                "url": "/contracts",
                "description": "View smart contract addresses and information",
            },
        }

    def get_page_from_file(self, file_path: str) -> Optional[Dict]:
        fp = file_path.replace("\\", "/")
        for route_key, page_info in self.page_route_map.items():
            if route_key in fp:
                return page_info
        return None

    def extract_bid_info(self, content: str) -> Dict[str, Any]:
        info = {}
        bid_types = re.findall(r'"(ETH|CST)"', content)
        if bid_types:
            info["bid_types"] = list(set(bid_types))
        if "RandomWalk" in content or "useRandomWalkNft" in content:
            info["supports_nft"] = True
        if "bidMessage" in content or "BidMessage" in content:
            info["supports_message"] = True
        if "donation" in content.lower():
            info["supports_donation"] = True
        if "showAdvancedOptions" in content:
            info["has_advanced_options"] = True
        if "isConnected" in content or "useAccount" in content:
            info["requires_wallet"] = True
        if "claimPrize" in content:
            info["has_prize_claim"] = True
        return info

    def extract_prize_info(self, content: str) -> Dict[str, Any]:
        info = {"prize_types": []}
        prize_sections = re.findall(
            r'name:\s*["\']([^"\']+)["\'],\s*'
            r'(?:percentage:\s*[^,]+,\s*)?'
            r'description:\s*["\']([^"\']+)["\'],\s*'
            r'rewards:\s*\[([^\]]+)\]',
            content,
            re.DOTALL,
        )
        for name, description, rewards_str in prize_sections:
            reward_items = re.findall(r'`([^`]+)`|["\']([^"\']+)["\']', rewards_str)
            rewards = [r[0] or r[1] for r in reward_items if r[0] or r[1]]
            info["prize_types"].append(
                {"name": name, "description": description, "rewards": rewards}
            )
        return info

    def generate_user_answer(self, question: str, documents: List[Document]) -> str:
        """Generate a user-friendly answer by understanding the code"""
        question_lower = question.lower()
        bid_info: Dict = {}
        prize_info: Dict = {}
        all_page_info: List = []

        for doc in documents:
            file_path = doc.meta.get("file_path", "")
            content = doc.content
            page = self.get_page_from_file(file_path)
            if page:
                entry: Dict[str, Any] = {"page": page, "file": file_path}
                if "play" in file_path:
                    bid_info = self.extract_bid_info(content)
                    entry["bid_info"] = bid_info
                elif "prizes" in file_path or "prize" in file_path:
                    prize_info = self.extract_prize_info(content)
                    entry["prize_info"] = prize_info
                all_page_info.append(entry)

        if any(kw in question_lower for kw in ["bid", "place bid", "how to bid", "make a bid"]):
            return self._answer_bidding(bid_info)
        elif any(kw in question_lower for kw in ["prize", "win", "reward", "distribution", "payout"]):
            return self._answer_prizes(prize_info)
        elif any(kw in question_lower for kw in ["stake", "staking"]):
            return self._answer_staking()
        elif any(kw in question_lower for kw in ["wallet", "connect", "metamask"]):
            return self._answer_wallet()
        elif any(kw in question_lower for kw in ["contract", "smart contract", "how many contract"]):
            return self._answer_contracts(documents)
        elif any(kw in question_lower for kw in ["nft", "cosmic signature nft"]):
            return self._answer_nfts()
        elif any(kw in question_lower for kw in ["leaderboard", "ranking", "top player"]):
            return self._answer_leaderboard()
        elif any(kw in question_lower for kw in ["history", "past round", "previous round"]):
            return self._answer_history()
        else:
            return self._answer_general(question, documents, all_page_info)

    def _answer_bidding(self, bid_info: Dict) -> str:
        lines = [
            "**How to Place a Bid on Cosmic Signature**\n",
            "**Step 1 – Connect Your Wallet**",
            "- Install MetaMask (https://metamask.io) or another Web3 wallet",
            "- Click the **Connect Wallet** button in the top-right corner of the website",
            "- Approve the connection in your wallet popup\n",
            "**Step 2 – Go to the Play Page**",
            "- Navigate to **/game/play** from the main menu",
            "- This is the main bidding interface\n",
            "**Step 3 – Choose Your Bid Type**",
        ]
        bid_types = bid_info.get("bid_types", ["ETH", "CST"])
        lines.append(f"- Select your token: **{' or '.join(bid_types)}**")
        lines.append("  - **ETH** – Bid with Ether (standard option)")
        lines.append("  - **CST** – Bid with Cosmic Signature Tokens\n")

        if bid_info.get("supports_nft"):
            lines += [
                "**Step 4 – (Optional) Use a RandomWalk NFT**",
                "- If you own a RandomWalk NFT, toggle **Use RandomWalk NFT** for a bid discount",
                "- Select your NFT from the dropdown\n",
            ]
        else:
            lines += ["**Step 4 – Review the Current Bid Price**",
                      "- The current price is shown on the page and updates in real time\n"]

        if bid_info.get("supports_message"):
            lines += [
                "**Step 5 – (Optional) Add a Message**",
                "- Type a short public message to attach to your bid\n",
            ]

        lines += [
            "**Step 6 – Place Your Bid**",
            "- Click the **Place Bid** button",
            "- Approve and sign the transaction in your wallet",
            "- Wait for blockchain confirmation\n",
            "**Step 7 – You're the Last Bidder!**",
            "- A countdown timer resets after your bid",
            "- If no one outbids you before the timer expires, **you win the prize!**",
        ]

        if bid_info.get("has_advanced_options"):
            lines += [
                "\n**Advanced Options**",
                "- Expand **Advanced Options** to donate NFTs or tokens alongside your bid",
            ]
        return "\n".join(lines)

    def _answer_prizes(self, prize_info: Dict) -> str:
        lines = [
            "**Prize Distribution in Cosmic Signature**\n",
            "View the full breakdown on the **Prizes page**: `/game/prizes`\n",
            "When a round ends, the ETH prize pool is split as follows:\n",
        ]
        prize_types = prize_info.get("prize_types", [])
        if prize_types:
            for p in prize_types:
                lines.append(f"**{p['name']}**")
                lines.append(f"- *Who:* {p['description']}")
                for r in p.get("rewards", []):
                    lines.append(f"- *Reward:* {r}")
                lines.append("")
        else:
            lines += [
                "**Main Prize** – Last bidder when the timer runs out",
                "- % of ETH pool + 1 Cosmic Signature NFT\n",
                "**Endurance Champion** – Longest single stint as last bidder",
                "- 10x CST per bid + 1 Cosmic Signature NFT\n",
                "**Chrono-Warrior** – Longest total time as Endurance Champion",
                "- % of ETH pool\n",
                "**Raffle Winners** – Random selection among all bidders",
                "- Share of ETH pool + Cosmic Signature NFTs\n",
                "**NFT Stakers** – Players who stake their NFTs",
                "- Proportional share of ETH pool\n",
                "**Charity** – A portion goes to charitable causes\n",
            ]
        lines.append("📌 Visit `/game/prizes` to see exact percentages for the current round.")
        return "\n".join(lines)

    def _answer_staking(self) -> str:
        return "\n".join([
            "**How to Stake NFTs on Cosmic Signature**\n",
            "**Step 1 – Get a Cosmic Signature or RandomWalk NFT**",
            "- Win one by placing bids or winning a raffle\n",
            "**Step 2 – Go to the Staking Page**",
            "- Navigate to **/stake** from the main navigation\n",
            "**Step 3 – Stake Your NFT**",
            "- Connect your wallet",
            "- Select the NFT(s) you want to stake",
            "- Click **Stake** and approve the transaction\n",
            "**Step 4 – Earn Rewards**",
            "- Every round, a portion of the ETH prize pool goes to stakers",
            "- Rewards are proportional to your staked NFTs\n",
            "**Step 5 – Unstake When Ready**",
            "- Unstake at any time from `/stake`",
            "- Claim pending rewards before unstaking",
        ])

    def _answer_wallet(self) -> str:
        return "\n".join([
            "**How to Connect Your Wallet**\n",
            "**Step 1 – Install a Web3 Wallet**",
            "- Download MetaMask from https://metamask.io",
            "- Create an account and save your seed phrase securely\n",
            "**Step 2 – Click Connect Wallet**",
            "- Click the **Connect Wallet** button in the top-right corner of the site\n",
            "**Step 3 – Approve the Connection**",
            "- Select your wallet type (MetaMask, WalletConnect, etc.)",
            "- Click **Approve** in the wallet popup\n",
            "**Step 4 – Done!**",
            "- Your address appears in the top-right corner",
            "- You can now bid, stake NFTs, and claim prizes",
            "- Make sure you have ETH for bids and gas fees",
        ])

    def _answer_contracts(self, documents: List[Document]) -> str:
        sol_files = sorted({
            doc.meta.get("file_path", "")
            for doc in documents
            if doc.meta.get("file_path", "").endswith(".sol")
        })
        lines = ["**Smart Contracts in Cosmic Signature**\n"]
        if sol_files:
            lines.append(f"Found {len(sol_files)} Solidity contract file(s) in the indexed documents:\n")
            for f in sol_files[:15]:
                lines.append(f"- `{f}`")
            if len(sol_files) > 15:
                lines.append(f"- ... and {len(sol_files) - 15} more")
        else:
            lines += [
                "The project includes several key smart contracts:\n",
                "- **CosmicGame** – Main game logic (bids, prizes, rounds)",
                "- **CosmicToken (CST)** – ERC20 token used for bidding",
                "- **CosmicSignature NFT** – ERC721 NFT awarded to winners",
                "- **RandomWalkNFT** – NFT that gives bid discounts",
                "- **StakingWallet** – Manages NFT staking and rewards",
                "- **CharityWallet** – Handles charitable contributions",
            ]
        lines.append("\n📌 View all contract addresses on the **Contracts page**: `/contracts`")
        return "\n".join(lines)

    def _answer_nfts(self) -> str:
        return "\n".join([
            "**NFTs in Cosmic Signature**\n",
            "**Cosmic Signature NFT**",
            "- Awarded to round winners (main prize, raffle, endurance champion)",
            "- Can be staked to earn passive ETH rewards",
            "- Browse all NFTs in the **Gallery**: `/gallery`",
            "- View your NFTs at `/account/nfts`\n",
            "**RandomWalk NFT**",
            "- Gives a **discount on bid prices** when used",
            "- Each NFT can only be used once per game",
            "- Can also be staked for rewards\n",
            "**How to Get NFTs**",
            "- Win the main prize or raffle in a round",
            "- Be the Endurance Champion",
            "- Purchase on NFT marketplaces\n",
            "📌 View your NFTs at `/account/nfts`",
        ])

    def _answer_leaderboard(self) -> str:
        return "\n".join([
            "**Leaderboard & Rankings**\n",
            "1. Go to **/game/leaderboard** to see top players",
            "2. Go to **/game/statistics** for detailed analytics\n",
            "Your personal stats:",
            "- `/account/statistics` – Your game statistics",
            "- `/account/activity` – Your bidding history",
        ])

    def _answer_history(self) -> str:
        return "\n".join([
            "**Game History & Past Rounds**\n",
            "1. Go to **/game/history/rounds** to browse past rounds",
            "   - See who won each round and the prize amounts",
            "2. Click any round to see details: `/game/history/rounds/[id]`",
            "3. View individual bid details: `/game/history/bids/[id]`\n",
            "📌 Your personal history is at `/account/activity`",
        ])

    def _answer_general(self, question: str, documents: List[Document], all_page_info: List) -> str:
        lines = ["Based on the Cosmic Signature codebase, here's what I found:\n"]
        pages_found = [p["page"] for p in all_page_info if p.get("page")]
        if pages_found:
            lines.append("**Relevant pages:**")
            for page in pages_found[:3]:
                lines.append(f"- **{page['name']}** (`{page['url']}`) – {page['description']}")
            lines.append("")
        for doc in documents[:2]:
            repo = doc.meta.get("repository", "unknown")
            fp = doc.meta.get("file_path", "unknown")
            snippet = doc.content[:200].replace("\n", " ").strip()
            lines.append(f"**From `{repo}/{fp}`:**")
            lines.append(f"_{snippet}..._\n")
        lines += [
            "\n💡 Try asking more specifically:",
            "- \"How do I place a bid?\"",
            "- \"How are prizes distributed?\"",
            "- \"How do I stake NFTs?\"",
            "- \"How many smart contracts are there?\"",
        ]
        return "\n".join(lines)
