#!/usr/bin/env python3
"""validate-chain.py -- chain-vs-database audit for the Cosmic Signature backend.

The SQL validators (validate-db.sql, validate-stats.sql, validate-durations.sql)
prove the database is INTERNALLY consistent. This script closes the remaining
gap: it reconciles the database against live Arbitrum contract state, catching
indexer omissions that no internal check can see.

Checks:
  1. head lag        -- chain head vs highest indexed block (informational/WARN).
  2. log parity      -- eth_getLogs count per tracked contract vs evt_log count.
                        Catches silently skipped logs or block ranges. Scans from
                        the first indexed log of the contract family, so a missed
                        prefix BEFORE the first indexed log is out of scope.
  3. CST supply      -- CosmicSignatureToken.totalSupply() vs mints-burns in
                        cg_erc20_transfer (otype 1 minus otype 2).
  4. NFT supply+owners -- CosmicSignatureNft.totalSupply() vs cg_mint_event count,
                        and ownerOf(token) vs cg_mint_event.cur_owner_aid for
                        every token. (Staked NFTs are owned by the staking
                        wallet on chain AND in the DB, so they compare equal.)
  5. staked counts   -- numStakedNfts() on both staking wallets vs
                        stake-minus-unstake event counts.
  6. PrizesWallet    -- on-chain ETH balance vs deposits minus withdrawals.
  7. roundStats      -- the V3 roundStats(r) getter vs cg_bid recomputation:
                        numBids for every round (the V2->V3 storage
                        reinterpretation makes it valid for legacy rounds too);
                        numCstBids / totalSpentEthAmount / totalSpentCstAmount /
                        max bid prices only for rounds that started under V3
                        (the chain fields read 0 for earlier rounds).

Live-chain races: bids can land between a chain read and a DB read, and the ETL
tails the chain, so volatile checks are retried a few times and pass if chain
and DB agree in any attempt.

Environment (same DB conventions as validate-db.sh / cg-etl):
  RPC_URL         Arbitrum RPC endpoint (required). Full nodes are fine: the
                  script uses only 'latest' state and historical logs/receipts,
                  never historical state.
  DATABASE_URL    complete postgres:// URL; when set it wins over PGSQL_*.
  PGSQL_HOST      "host" or "host:port"; empty selects the local Unix socket.
  PGSQL_USERNAME  database user.
  PGSQL_PASSWORD  database password (optional).
  PGSQL_DATABASE  database name.

Read-only on both sides: SELECTs on the database, eth_call / eth_getLogs /
eth_getBalance / eth_blockNumber on the node.

Exit codes: 0 all checks pass (warnings allowed), 2 usage/config error,
3 at least one FAIL (matching the SQL validators' convention).
"""

import json
import os
import subprocess
import sys
import time
import urllib.request

# --- function selectors (keccak4, generated with go-ethereum crypto.Keccak256) --
SEL_TOTAL_SUPPLY = "0x18160ddd"  # totalSupply()
SEL_OWNER_OF = "0x6352211e"      # ownerOf(uint256)
SEL_NUM_STAKED = "0xca7c1f92"    # numStakedNfts()
SEL_ROUND_STATS = "0x259d33c1"   # roundStats(uint256) -> 9 uint256 words
SEL_ROUND_NUM = "0x119b22b3"     # roundNum()

LOG_CHUNK = 8_000_000      # initial eth_getLogs span; bisected on node errors
VOLATILE_RETRIES = 4       # chain/DB agreement attempts for live values
VOLATILE_DELAY = 2.0       # seconds between attempts
HEAD_LAG_WARN = 1800       # blocks (~7.5 min on Arbitrum) before WARN

results = []  # (status, check_name, details)


def add(status, name, details=""):
    results.append((status, name, details))
    print(f" {status:<5}| {name}" + (f" | {details}" if details else ""), flush=True)


# --------------------------------------------------------------------------
# Database access: shell out to psql with the same env handling as
# validate-db.sh so the script works wherever the SQL validators work.
# --------------------------------------------------------------------------
def build_psql_argv():
    psql = os.environ.get("PSQL", "psql")
    url = os.environ.get("DATABASE_URL")
    if url:
        return [psql, url]
    db = os.environ.get("PGSQL_DATABASE")
    if not db:
        sys.stderr.write("error: DATABASE_URL, or PGSQL_HOST, PGSQL_USERNAME,"
                         " and PGSQL_DATABASE must be set\n")
        sys.exit(2)
    argv = [psql, "-d", db]
    host = os.environ.get("PGSQL_HOST", "")
    if host:
        if ":" in host:
            host, port = host.rsplit(":", 1)
            argv += ["-p", port]
        argv += ["-h", host]
        os.environ["PGSSLMODE"] = "disable"
    user = os.environ.get("PGSQL_USERNAME")
    if user:
        argv += ["-U", user]
    pw = os.environ.get("PGSQL_PASSWORD")
    if pw:
        os.environ["PGPASSWORD"] = pw
    return argv


PSQL_ARGV = None


def db(sql):
    """Run one read-only query; return rows as lists of strings ('' for NULL)."""
    # -X skips ~/.psqlrc, whose chatter (e.g. "Pager usage is off.") would
    # otherwise be parsed as a result row.
    out = subprocess.run(
        PSQL_ARGV + ["-X", "-tA", "-F", "\t", "-v", "ON_ERROR_STOP=1", "-c", sql],
        capture_output=True, text=True)
    if out.returncode != 0:
        sys.stderr.write(f"psql failed:\n{sql}\n{out.stderr}\n")
        sys.exit(2)
    return [line.split("\t") for line in out.stdout.splitlines() if line != ""]


def db1(sql):
    """Single scalar; '' when NULL / no row."""
    rows = db(sql)
    return rows[0][0] if rows and rows[0] else ""


# --------------------------------------------------------------------------
# RPC access
# --------------------------------------------------------------------------
RPC_URL = None
_rpc_id = 0


def rpc(method, params, timeout=90):
    global _rpc_id
    _rpc_id += 1
    body = json.dumps({"jsonrpc": "2.0", "id": _rpc_id,
                       "method": method, "params": params}).encode()
    req = urllib.request.Request(
        RPC_URL, data=body, headers={"Content-Type": "application/json"})
    last = None
    for attempt in range(3):
        try:
            with urllib.request.urlopen(req, timeout=timeout) as resp:
                reply = json.loads(resp.read())
            if "error" in reply:
                raise RpcError(reply["error"].get("message", str(reply["error"])))
            return reply["result"]
        except RpcError:
            raise
        except Exception as exc:  # transport hiccup: retry
            last = exc
            time.sleep(1.5 * (attempt + 1))
    raise RuntimeError(f"RPC {method} failed after retries: {last}")


class RpcError(Exception):
    """Node-reported JSON-RPC error (e.g. too many getLogs results)."""


def eth_call(to, data):
    return rpc("eth_call", [{"to": to, "data": data}, "latest"])


def call_uint(to, data):
    return int(eth_call(to, data), 16)


def encode_uint(n):
    return format(n, "064x")


def get_logs_count(addr, frm, to):
    """Count logs for addr in [frm, to], bisecting ranges the node rejects."""
    if frm > to:
        return 0
    try:
        logs = rpc("eth_getLogs", [{
            "address": addr, "fromBlock": hex(frm), "toBlock": hex(to)}])
        return len(logs)
    except RpcError:
        if frm == to:
            raise
        mid = (frm + to) // 2
        return get_logs_count(addr, frm, mid) + get_logs_count(addr, mid + 1, to)


def get_logs_count_chunked(addr, frm, to):
    total = 0
    start = frm
    while start <= to:
        end = min(start + LOG_CHUNK - 1, to)
        total += get_logs_count(addr, start, end)
        start = end + 1
    return total


def volatile_match(chain_fn, db_fn, label):
    """Compare a live chain value with a DB value, retrying through ETL lag.

    Returns (ok, chain_value, db_value) using the last attempt's values.
    """
    chain_v = db_v = None
    for attempt in range(VOLATILE_RETRIES):
        if attempt:
            time.sleep(VOLATILE_DELAY)
        chain_v = chain_fn()
        db_v = db_fn()
        if chain_v == db_v:
            return True, chain_v, db_v
    return False, chain_v, db_v


# --------------------------------------------------------------------------
# Checks
# --------------------------------------------------------------------------
def main():
    global PSQL_ARGV, RPC_URL
    RPC_URL = os.environ.get("RPC_URL")
    if not RPC_URL:
        sys.stderr.write("error: RPC_URL must be set (Arbitrum RPC endpoint)\n")
        sys.exit(2)
    PSQL_ARGV = build_psql_argv()

    print("=== Cosmic Signature chain-vs-DB validator ===\n")

    # -- contract addresses (registry + every proxy implementation) ----------
    row = db("SELECT cosmic_game_addr, cosmic_signature_addr, cosmic_token_addr,"
             " cosmic_dao_addr, charity_wallet_addr, prizes_wallet_addr,"
             " random_walk_addr, staking_wallet_cst_addr,"
             " staking_wallet_rwalk_addr, marketing_wallet_addr,"
             " implementation_addr FROM cg_contracts")[0]
    (game, nft, cst, dao, charity, prizes, rwalk,
     staking_cst, staking_rwalk, marketing, impl_v1) = [a.lower() for a in row]
    impls = [r[0].lower() for r in db(
        "SELECT DISTINCT LOWER(a.addr) FROM cg_adm_upgraded u"
        " JOIN address a ON a.address_id = u.implementation_aid")]

    labeled = [("game proxy", game), ("CosmicSignatureNft", nft),
               ("CosmicSignatureToken", cst), ("CosmicSignatureDao", dao),
               ("CharityWallet", charity), ("PrizesWallet", prizes),
               ("RandomWalkNFT", rwalk),
               ("StakingWallet CST", staking_cst),
               ("StakingWallet RWalk", staking_rwalk),
               ("MarketingWallet", marketing)]
    seen = {a for _, a in labeled}
    for i, a in enumerate(sorted(set(impls + [impl_v1]))):
        if a not in seen:
            labeled.append((f"game impl {a[:10]}", a))
            seen.add(a)

    # -- 1. head lag ---------------------------------------------------------
    chain_head = int(rpc("eth_blockNumber", []), 16)
    db_head = int(db1("SELECT MAX(block_num) FROM block"))
    lag = chain_head - db_head
    add("WARN" if lag > HEAD_LAG_WARN else "PASS",
        "indexer head lag",
        f"chain={chain_head} db={db_head} lag={lag} blocks")

    # -- 2. log-count parity per contract ------------------------------------
    # Scan window: from the first indexed cosmic-platform log (= deployment-era
    # constructor events) up to the DB head. RandomWalkNFT predates the
    # platform, so it uses its own first indexed log.
    cosmic_from = int(db1(
        "SELECT MIN(e.block_num) FROM evt_log e"
        " JOIN address a ON a.address_id = e.contract_aid"
        " WHERE LOWER(a.addr) = ANY(ARRAY["
        + ",".join(f"'{a}'" for _, a in labeled if a != rwalk) + "])"))
    for name, addr in labeled:
        frm = cosmic_from
        if addr == rwalk:
            v = db1("SELECT MIN(e.block_num) FROM evt_log e"
                    " JOIN address a ON a.address_id = e.contract_aid"
                    f" WHERE LOWER(a.addr) = '{addr}'")
            if v == "":
                add("WARN", f"log parity: {name}", "no indexed logs; skipped")
                continue
            frm = int(v)
        db_cnt = int(db1(
            "SELECT COUNT(*) FROM evt_log e"
            " JOIN address a ON a.address_id = e.contract_aid"
            f" WHERE LOWER(a.addr) = '{addr}' AND e.block_num <= {db_head}"))
        chain_cnt = get_logs_count_chunked(addr, frm, db_head)
        add("PASS" if chain_cnt == db_cnt else "FAIL",
            f"log parity: {name}",
            f"chain={chain_cnt} db={db_cnt} (blocks {frm}..{db_head})")

    # -- 3. CST total supply --------------------------------------------------
    ok, cv, dv = volatile_match(
        lambda: call_uint(cst, SEL_TOTAL_SUPPLY),
        lambda: int(db1(
            "SELECT COALESCE(SUM(CASE otype WHEN 1 THEN value"
            " WHEN 2 THEN -value ELSE 0 END), 0)::NUMERIC(80,0)"
            " FROM cg_erc20_transfer e JOIN address a"
            " ON a.address_id = e.contract_aid"
            f" WHERE LOWER(a.addr) = '{cst}'")),
        "CST supply")
    add("PASS" if ok else "FAIL", "CST totalSupply vs mints-burns",
        f"chain={cv} db={dv}")

    # -- 4. NFT supply and per-token ownership --------------------------------
    ok, cv, dv = volatile_match(
        lambda: call_uint(nft, SEL_TOTAL_SUPPLY),
        lambda: int(db1("SELECT COUNT(*) FROM cg_mint_event")),
        "NFT supply")
    add("PASS" if ok else "FAIL", "CS NFT totalSupply vs cg_mint_event count",
        f"chain={cv} db={dv}")

    owners = db("SELECT m.token_id, LOWER(a.addr) FROM cg_mint_event m"
                " JOIN address a ON a.address_id = m.cur_owner_aid"
                " ORDER BY m.token_id")
    mismatches = []
    for token_id, db_owner in owners:
        chain_owner = "0x" + eth_call(
            nft, SEL_OWNER_OF + encode_uint(int(token_id)))[-40:]
        if chain_owner != db_owner:
            mismatches.append(f"token {token_id}: chain={chain_owner} db={db_owner}")
    add("PASS" if not mismatches else "FAIL",
        f"ownerOf() vs cur_owner for {len(owners)} NFTs",
        "; ".join(mismatches[:5]) + ("..." if len(mismatches) > 5 else ""))

    # -- 5. staked-NFT counts --------------------------------------------------
    for name, wallet, staked_tbl, unstaked_tbl in [
            ("CST", staking_cst, "cg_nft_staked_cst", "cg_nft_unstaked_cst"),
            ("RandomWalk", staking_rwalk, "cg_nft_staked_rwalk",
             "cg_nft_unstaked_rwalk")]:
        ok, cv, dv = volatile_match(
            lambda w=wallet: call_uint(w, SEL_NUM_STAKED),
            lambda s=staked_tbl, u=unstaked_tbl: int(db1(
                f"SELECT (SELECT COUNT(*) FROM {s})"
                f" - (SELECT COUNT(*) FROM {u})")),
            f"staked {name}")
        add("PASS" if ok else "FAIL",
            f"numStakedNfts ({name}) vs stake-unstake events",
            f"chain={cv} db={dv}")

    # -- 6. PrizesWallet balance ----------------------------------------------
    ok, cv, dv = volatile_match(
        lambda: int(rpc("eth_getBalance", [prizes, "latest"]), 16),
        lambda: int(db1(
            "SELECT (SELECT COALESCE(SUM(amount),0) FROM cg_prize_deposit)"
            " - (SELECT COALESCE(SUM(amount),0) FROM cg_prize_withdrawal)")),
        "PrizesWallet balance")
    add("PASS" if ok else "FAIL",
        "PrizesWallet ETH balance vs deposits-withdrawals",
        f"chain={cv} db={dv} wei")

    # -- 7. roundStats(r) vs cg_bid recomputation ------------------------------
    cur_round = call_uint(game, SEL_ROUND_NUM)
    v3_block = db1("SELECT MIN(block_num) FROM cg_adm_initialized WHERE version = 3")
    for r in range(cur_round + 1):
        def stats():
            w = eth_call(game, SEL_ROUND_STATS + encode_uint(r))[2:]
            return [int(w[i * 64:(i + 1) * 64], 16) for i in range(len(w) // 64)]
        def db_stats():
            row = db(
                "SELECT COUNT(*),"
                " COUNT(*) FILTER (WHERE bid_type = 2),"
                " COALESCE(SUM(eth_price) FILTER (WHERE bid_type <> 2),0)::NUMERIC(80,0),"
                " COALESCE(SUM(cst_price) FILTER (WHERE bid_type = 2),0)::NUMERIC(80,0),"
                " COALESCE(MAX(eth_price) FILTER (WHERE bid_type <> 2),0)::NUMERIC(80,0),"
                " COALESCE(MAX(cst_price) FILTER (WHERE bid_type = 2),0)::NUMERIC(80,0)"
                f" FROM cg_bid WHERE round_num = {r}")[0]
            return [int(x) for x in row]
        first_bid_block = db1(f"SELECT MIN(block_num) FROM cg_bid WHERE round_num = {r}")
        v3_round = (v3_block != "" and first_bid_block != ""
                    and int(first_bid_block) >= int(v3_block))
        n_fields = 6 if v3_round else 1
        field_names = ["numBids", "numCstBids", "totalSpentEth", "totalSpentCst",
                       "maxEthBidPrice", "maxCstBidPrice"][:n_fields]
        ok, cv, dv = volatile_match(
            lambda: stats()[:n_fields], lambda: db_stats()[:n_fields],
            f"roundStats({r})")
        diffs = [f"{n}: chain={c} db={d}"
                 for n, c, d in zip(field_names, cv, dv) if c != d]
        scope = "V3 full" if v3_round else "legacy numBids-only"
        add("PASS" if ok else "FAIL",
            f"roundStats({r}) vs cg_bid ({scope})",
            "; ".join(diffs) if diffs else f"numBids={cv[0]}")

    # -- summary ---------------------------------------------------------------
    fails = sum(1 for s, _, _ in results if s == "FAIL")
    warns = sum(1 for s, _, _ in results if s == "WARN")
    print(f"\nchecks: {len(results)}  passed: {len(results) - fails - warns}"
          f"  failed: {fails}  warnings: {warns}")
    if fails:
        print("CHAIN VALIDATION FAILED: at least one check found a mismatch.")
        sys.exit(3)
    print("CHAIN VALIDATION OK: database matches on-chain state.")


if __name__ == "__main__":
    main()
