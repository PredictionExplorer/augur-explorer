#!/usr/bin/env python3
"""
Regenerate Random Walk PNG / JPG thumbs / MP4s from the indexer DB.

Uses rw_token (seed_hex, rwalk_aid) when present; otherwise rw_mint_evt (seed, contract_aid).
Note: rw_token.rwalk_aid is the NFT contract's address id — the same value as rw_mint_evt.contract_aid,
not necessarily 1. If you omit --rwalk-aid and RWALK_AID, the script picks the only contract_aid in the DB
or errors if there are several.

Environment:
  DATABASE_URL — postgresql://... (optional if PGSQL_* are set)
  PGSQL_HOST, PGSQL_USERNAME, PGSQL_DATABASE, PGSQL_PASSWORD — same as rwalk-local-hh.env
  RWALK_IMG_OUTPUT_DIR — output dir (defaults to NFT_ASSETS_ROOT if set)
  RWALK_AID — optional; same as contract_aid / rwalk_aid

Examples:
  source ~/configs/rwalk-local-hh.env
  ./regenerate_from_db.py --skip-existing
  ./regenerate_from_db.py --rwalk-aid 11 --token-id 0
"""

from __future__ import annotations

import argparse
import logging
import os
import sys
from urllib.parse import quote_plus

import psycopg2
from psycopg2.extras import RealDictCursor

from image_generator import create_media

ZERO_SEED = "0x" + "0" * 64


def resolve_database_url() -> str:
    direct = os.environ.get("DATABASE_URL", "").strip()
    if direct:
        return direct
    user = os.environ.get("PGSQL_USERNAME", "").strip()
    password = os.environ.get("PGSQL_PASSWORD", "").strip()
    database = os.environ.get("PGSQL_DATABASE", "").strip()
    hostport = os.environ.get("PGSQL_HOST", "").strip()
    if not user or not database or not hostport:
        return ""
    if ":" in hostport:
        host, port = hostport.rsplit(":", 1)
    else:
        host, port = hostport, "5432"
    return "postgresql://{}:{}@{}:{}/{}".format(
        quote_plus(user),
        quote_plus(password),
        host,
        port,
        quote_plus(database),
    )


def discover_contract_aid(cur) -> tuple[int | None, list[int]]:
    cur.execute(
        """
        SELECT DISTINCT aid FROM (
            SELECT rwalk_aid AS aid FROM rw_token WHERE rwalk_aid IS NOT NULL
            UNION
            SELECT contract_aid AS aid FROM rw_mint_evt WHERE contract_aid IS NOT NULL
        ) s ORDER BY aid
        """
    )
    aids = [int(r["aid"]) for r in cur.fetchall()]
    if len(aids) == 1:
        return aids[0], aids
    return None, aids


def resolve_rwalk_aid_arg(explicit: int | None) -> int | None:
    if explicit is not None:
        return explicit
    env = os.environ.get("RWALK_AID", "").strip()
    if env:
        return int(env)
    return None


def fetch_rows(cur, rwalk_aid: int, token_id: int | None):
    """Prefer rw_token; if no rows, use rw_mint_evt (same contract_aid / rwalk_aid)."""
    if token_id is not None:
        q_tok = (
            "SELECT token_id, seed_hex AS seed_hex FROM rw_token "
            "WHERE rwalk_aid = %s AND token_id = %s AND seed_hex IS NOT NULL AND trim(seed_hex) <> ''"
        )
        cur.execute(q_tok, (rwalk_aid, token_id))
        rows = cur.fetchall()
        if rows:
            return rows, "rw_token"
        q_mint = (
            "SELECT token_id, seed AS seed_hex FROM rw_mint_evt "
            "WHERE contract_aid = %s AND token_id = %s AND seed IS NOT NULL AND trim(seed::text) <> ''"
        )
        cur.execute(q_mint, (rwalk_aid, token_id))
        return cur.fetchall(), "rw_mint_evt"

    q_tok = (
        "SELECT token_id, seed_hex AS seed_hex FROM rw_token "
        "WHERE rwalk_aid = %s AND seed_hex IS NOT NULL AND trim(seed_hex) <> '' "
        "ORDER BY token_id"
    )
    cur.execute(q_tok, (rwalk_aid,))
    rows = cur.fetchall()
    if rows:
        return rows, "rw_token"
    q_mint = (
        "SELECT token_id, seed AS seed_hex FROM rw_mint_evt "
        "WHERE contract_aid = %s AND seed IS NOT NULL AND trim(seed::text) <> '' "
        "ORDER BY token_id"
    )
    cur.execute(q_mint, (rwalk_aid,))
    return cur.fetchall(), "rw_mint_evt"


def required_files_exist(out_dir: str, token_id: int) -> bool:
    base = f"{token_id:06d}"
    for bg in ("white", "black"):
        fn = f"{base}_{bg}"
        paths = [
            os.path.join(out_dir, f"{fn}.png"),
            os.path.join(out_dir, f"{fn}_thumb.jpg"),
            os.path.join(out_dir, f"{fn}_single.mp4"),
            os.path.join(out_dir, f"{fn}_triple.mp4"),
        ]
        if any(not os.path.isfile(p) for p in paths):
            return False
    return True


def main() -> None:
    logging.basicConfig(level=logging.INFO, format="%(levelname)s %(message)s")
    p = argparse.ArgumentParser(description="Regenerate RandomWalk assets from rw_token / rw_mint_evt")
    p.add_argument(
        "--rwalk-aid",
        type=int,
        default=None,
        help="NFT contract address id (rw_token.rwalk_aid = rw_mint_evt.contract_aid); default: RWALK_AID env or auto-detect",
    )
    p.add_argument("--skip-existing", action="store_true", help="skip tokens whose 8 files already exist")
    p.add_argument("--token-id", type=int, default=None, help="only this token_id")
    p.add_argument("--limit", type=int, default=None, help="max rows (after ordering)")
    args = p.parse_args()

    db_url = resolve_database_url()
    if not db_url:
        logging.error(
            "Set DATABASE_URL or PGSQL_HOST, PGSQL_USERNAME, PGSQL_DATABASE, PGSQL_PASSWORD"
        )
        sys.exit(1)
    out_dir = (
        os.environ.get("RWALK_IMG_OUTPUT_DIR", "").strip()
        or os.environ.get("NFT_ASSETS_ROOT", "").strip()
    )
    if not out_dir:
        logging.error("Set RWALK_IMG_OUTPUT_DIR or NFT_ASSETS_ROOT")
        sys.exit(1)
    os.makedirs(out_dir, exist_ok=True)

    conn = psycopg2.connect(db_url)
    try:
        with conn.cursor(cursor_factory=RealDictCursor) as cur:
            rwalk_aid = resolve_rwalk_aid_arg(args.rwalk_aid)
            if rwalk_aid is None:
                discovered, all_aids = discover_contract_aid(cur)
                if discovered is not None:
                    rwalk_aid = discovered
                    logging.info("using contract_aid / rwalk_aid=%s (auto-detected)", rwalk_aid)
                elif not all_aids:
                    logging.error("No rows in rw_token or rw_mint_evt; nothing to regenerate")
                    sys.exit(1)
                else:
                    logging.error(
                        "Multiple contract ids %s — set RWALK_AID or --rwalk-aid (e.g. your rw_mint_evt.contract_aid)",
                        all_aids,
                    )
                    sys.exit(1)
            else:
                logging.info("using contract_aid / rwalk_aid=%s", rwalk_aid)

            rows, source = fetch_rows(cur, rwalk_aid, args.token_id)
            logging.info("loaded %s row(s) from %s", len(rows), source)
    finally:
        conn.close()

    if args.limit is not None:
        rows = rows[: args.limit]

    done = 0
    skipped = 0
    for row in rows:
        tid = int(row["token_id"])
        seed = (row["seed_hex"] or "").strip()
        if not seed or seed.lower() == ZERO_SEED.lower():
            logging.warning("skip token_id=%s invalid/zero seed", tid)
            skipped += 1
            continue
        if args.skip_existing and required_files_exist(out_dir, tid):
            logging.info("skip token_id=%s (complete set on disk)", tid)
            skipped += 1
            continue
        file_name = f"{tid:06d}"
        logging.info("generate token_id=%s", tid)
        try:
            create_media(file_name, seed)
        except Exception as e:
            logging.exception("token_id=%s failed: %s", tid, e)
            sys.exit(1)
        done += 1

    logging.info("finished: generated=%s skipped=%s total_rows=%s", done, skipped, len(rows))


if __name__ == "__main__":
    main()
