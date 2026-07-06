"""Tests for deterministic deployment address extraction."""
from __future__ import annotations

import json
import tempfile
import unittest
from pathlib import Path
from unittest import mock

from knowledge.generate import extract_deployments as ed


class ExtractDeploymentsTest(unittest.TestCase):
    def test_canonical_mainnet_seed_parses_game_proxy(self):
        seed = ed.DEPLOYMENT_SEEDS_DIR / "arbitrum-mainnet.txt"
        self.assertTrue(seed.is_file(), f"missing canonical seed {seed}")
        addresses = ed._parse_deployment_text(seed.read_text(encoding="utf-8"))
        self.assertIn(ed.GAME_PROXY_KEY, addresses)
        self.assertRegex(addresses[ed.GAME_PROXY_KEY], r"^0x[a-fA-F0-9]{40}$")

    def test_load_address_networks_from_repo_seeds(self):
        networks, sources = ed._load_address_networks()
        self.assertIn(ed.MAINNET_NETWORK_KEY, networks)
        self.assertTrue(sources)
        addrs = networks[ed.MAINNET_NETWORK_KEY]["addresses"]
        self.assertIn(ed.GAME_PROXY_KEY, addrs)

    def test_run_writes_deployed_addresses_with_mainnet(self):
        with tempfile.TemporaryDirectory() as tmp:
            kb = Path(tmp) / "kb"
            kb.mkdir()
            with mock.patch.object(ed, "KNOWLEDGE_BASE", kb), mock.patch.object(
                ed, "FACTS_DIR", kb / "facts"
            ), mock.patch.object(ed, "DOCS_EXPERT_DIR", kb / "docs" / "expert"), mock.patch.object(
                ed, "DOCS_SOURCES_DIR", kb / "docs" / "sources"
            ):
                ed.FACTS_DIR.mkdir(parents=True)
                ed.DOCS_EXPERT_DIR.mkdir(parents=True)
                ed.run()
                data = json.loads((ed.FACTS_DIR / "deployed-addresses.json").read_text())
                mainnet = data["networks"][ed.MAINNET_NETWORK_KEY]["addresses"]
                self.assertIn(ed.GAME_PROXY_KEY, mainnet)


if __name__ == "__main__":
    unittest.main()
