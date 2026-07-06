"""Clone or update upstream GitHub repos used by knowledge generators."""
from __future__ import annotations

import logging
import os

import git

from knowledge.config import GITHUB_REPOS, REPOS_DIR

logger = logging.getLogger(__name__)


def _skip_sync() -> bool:
    raw = os.getenv("FAQ_BOT_SKIP_REPO_SYNC", "").strip().lower()
    return raw in {"1", "true", "yes", "on"}


def sync_repositories() -> dict[str, str]:
    """
    Ensure REPOS_DIR contains shallow clones of smart_contracts, backend_api, frontend.

    Fresh git clones of faq-bot do not include data/repos/ — generators need these
    sources (especially frontend/src/contracts for ABIs).
    """
    if _skip_sync():
        logger.info("Skipping repository sync (FAQ_BOT_SKIP_REPO_SYNC is set)")
        return {"status": "skipped"}

    REPOS_DIR.mkdir(parents=True, exist_ok=True)
    results: dict[str, str] = {}

    for name, url in GITHUB_REPOS.items():
        repo_path = REPOS_DIR / name
        try:
            if repo_path.exists() and (repo_path / ".git").is_dir():
                logger.info("Updating repository %s", name)
                repo = git.Repo(repo_path)
                repo.remotes.origin.pull()
                results[name] = "updated"
            else:
                logger.info("Cloning repository %s from %s", name, url)
                git.Repo.clone_from(url, repo_path, depth=1)
                results[name] = "cloned"
        except Exception as exc:
            raise RuntimeError(
                f"Failed to sync repository '{name}' into {repo_path}: {exc}"
            ) from exc

    return results


def run() -> None:
    sync_repositories()
