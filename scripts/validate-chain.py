#!/usr/bin/env python3
"""validate-chain.py -- chain-vs-database audit for the Cosmic Signature backend.

The SQL validators (validate-db.sql, validate-stats.sql, validate-durations.sql)
prove the database is INTERNALLY consistent. This script closes the remaining
gap: it reconciles the database against live Arbitrum contract state, catching
indexer omissions that no internal check can see.

Checks:
  1. log parity      -- eth_getLogs count per tracked contract vs evt_log count.
                        Catches silently skipped logs or block ranges. Scans from
                        the first indexed log of the contract family, so a missed
                        prefix BEFORE the first indexed log is out of scope.
                        All platform contracts are counted in ONE pass over the
                        range (eth_getLogs takes an address array); RandomWalkNFT
                        history older than the platform is scanned separately.
                        Scan progress is checkpointed to a state file after every
                        chunk (VALIDATE_CHAIN_STATE, default
                        ~/.cache/cosmic-signature/validate-chain-state.json,
                        keyed by chain id), so the first run pays for the full
                        history and later runs only scan new blocks. Interrupting
                        with Ctrl-C loses at most one chunk. Set
                        SKIP_LOG_PARITY=1 to run only the fast checks.
                        ARCHIVE MODE: when ARCHIVE_DIR points at a tx-collector
                        backup tree (<dir>/<block>/<txhash>_receipt.rlp, written
                        by `opsctl tx-collector run` or the old standalone
                        transaction-collector), the parity check reads the
                        archived receipts instead of calling eth_getLogs — no
                        node work at all, and finer grained: per-transaction
                        per-contract log counts are compared in both directions,
                        so a single log missing from evt_log inside an otherwise
                        indexed transaction is caught. ARCHIVE_CONFIG must point
                        at the collector's JSON config; only contracts the
                        collector discovers (its "contracts" map) are fully
                        covered, others are reported as WARN/skipped.
  2. CST supply      -- CosmicSignatureToken.totalSupply() vs mints-burns in
                        cg_erc20_transfer (otype 1 minus otype 2).
  3. NFT supply+owners -- CosmicSignatureNft.totalSupply() vs cg_mint_event count,
                        and ownerOf(token) vs cg_mint_event.cur_owner_aid for
                        every token. (Staked NFTs are owned by the staking
                        wallet on chain AND in the DB, so they compare equal.)
  4. staked counts   -- numStakedNfts() on both staking wallets vs
                        stake-minus-unstake event counts.
  5. PrizesWallet    -- on-chain ETH balance vs deposits minus withdrawals.
  6. roundStats      -- the V3 roundStats(r) getter vs cg_bid recomputation:
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
import signal
import subprocess
import sys
import time
import urllib.request
from concurrent.futures import ThreadPoolExecutor, as_completed

# --- function selectors (keccak4, generated with go-ethereum crypto.Keccak256) --
SEL_TOTAL_SUPPLY = "0x18160ddd"  # totalSupply()
SEL_OWNER_OF = "0x6352211e"      # ownerOf(uint256)
SEL_NUM_STAKED = "0xca7c1f92"    # numStakedNfts()
SEL_ROUND_STATS = "0x259d33c1"   # roundStats(uint256) -> 9 uint256 words
SEL_ROUND_NUM = "0x119b22b3"     # roundNum()

LOG_CHUNK = int(os.environ.get("LOG_CHUNK", "2000000"))
                           # eth_getLogs span per request; bisected on node errors
LOG_TIMEOUT = 900          # seconds per getLogs request (bloom scans are slow);
                           # single attempt so timeouts never stack queries
LOG_WORKERS = int(os.environ.get("LOG_WORKERS", "4"))  # concurrent getLogs
                           # requests; raise if the node has idle cores, lower
                           # to 1 if it must stay responsive for other traffic
VOLATILE_RETRIES = 4       # chain/DB agreement attempts for live values
VOLATILE_DELAY = 2.0       # seconds between attempts
STATE_PATH = os.environ.get(
    "VALIDATE_CHAIN_STATE",
    os.path.expanduser("~/.cache/cosmic-signature/validate-chain-state.json"))

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


def rpc(method, params, timeout=90, attempts=3):
    global _rpc_id
    _rpc_id += 1
    body = json.dumps({"jsonrpc": "2.0", "id": _rpc_id,
                       "method": method, "params": params}).encode()
    req = urllib.request.Request(
        RPC_URL, data=body, headers={"Content-Type": "application/json"})
    last = None
    for attempt in range(attempts):
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


def get_logs_by_addr(addrs, frm, to, counts):
    """Add per-address log counts in [frm, to] into counts (one node pass for
    all addresses), bisecting ranges the node rejects."""
    if frm > to:
        return
    try:
        logs = rpc("eth_getLogs",
                   [{"address": addrs, "fromBlock": hex(frm), "toBlock": hex(to)}],
                   timeout=LOG_TIMEOUT, attempts=1)
    except RpcError:
        if frm == to:
            raise
        mid = (frm + to) // 2
        get_logs_by_addr(addrs, frm, mid, counts)
        get_logs_by_addr(addrs, mid + 1, to, counts)
        return
    for log in logs:
        a = log["address"].lower()
        counts[a] = counts.get(a, 0) + 1


def load_state():
    try:
        with open(STATE_PATH) as f:
            return json.load(f)
    except (OSError, ValueError):
        return {}


def save_state(state):
    os.makedirs(os.path.dirname(STATE_PATH), exist_ok=True)
    tmp = STATE_PATH + ".tmp"
    with open(tmp, "w") as f:
        json.dump(state, f)
    os.replace(tmp, STATE_PATH)


def scan_logs(state, scan_key, addrs, frm, to):
    """Checkpointed multi-address log count over [frm, to].

    Returns {addr: count} covering exactly [frm, to], resuming from and
    updating the persisted state. Returns None if the checkpoint is ahead of
    `to` (e.g. the DB was restored from an older backup), since counts up to
    `to` can then no longer be derived from the checkpoint.
    """
    entry = state.get(scan_key)
    if entry is None or entry.get("from") != frm:
        entry = {"from": frm, "scanned_to": frm - 1, "counts": {}}
    if entry["scanned_to"] > to:
        return None
    counts = dict(entry["counts"])
    chunks = []
    start = entry["scanned_to"] + 1
    while start <= to:
        end = min(start + LOG_CHUNK - 1, to)
        chunks.append((start, end))
        start = end + 1
    if not chunks:
        return counts

    def scan_chunk(chunk):
        delta = {}
        get_logs_by_addr(addrs, chunk[0], chunk[1], delta)
        return delta

    # LOG_WORKERS chunks run concurrently; the checkpoint only ever advances
    # over the CONTIGUOUS completed prefix, so an interrupt mid-flight never
    # records a gap (out-of-order completions are held back until their
    # predecessors finish, and at most the in-flight chunks are lost).
    t0 = time.time()
    folded = 0
    pending = {}  # start -> per-address delta, completed but not yet folded
    with ThreadPoolExecutor(max_workers=LOG_WORKERS) as pool:
        futures = {pool.submit(scan_chunk, c): c for c in chunks}
        next_idx = 0  # next chunk to fold into the checkpoint
        for fut in as_completed(futures):
            chunk = futures[fut]
            pending[chunk[0]] = fut.result()
            while next_idx < len(chunks) and chunks[next_idx][0] in pending:
                cstart, cend = chunks[next_idx]
                for a, n in pending.pop(cstart).items():
                    counts[a] = counts.get(a, 0) + n
                state[scan_key] = {"from": frm, "scanned_to": cend,
                                   "counts": counts}
                save_state(state)
                next_idx += 1
                folded += 1
                rate = (time.time() - t0) / folded
                eta = int(rate * (len(chunks) - next_idx))
                sys.stderr.write(
                    f"[scan {scan_key}] blocks {cstart}..{cend} of {to}"
                    f" ({(cend - frm + 1) * 100 // (to - frm + 1)}%,"
                    f" {sum(counts.values())} logs,"
                    f" ETA {eta // 60}m{eta % 60:02d}s)\n")
                sys.stderr.flush()
    return counts


# --------------------------------------------------------------------------
# Archive mode: read tx-collector receipt backups instead of eth_getLogs.
# --------------------------------------------------------------------------
def rlp_decode(data, pos=0):
    """Decode one RLP item at pos; return (item, next_pos). Items are bytes or
    lists of items."""
    prefix = data[pos]
    if prefix < 0x80:
        return data[pos:pos + 1], pos + 1
    if prefix < 0xb8:
        n = prefix - 0x80
        return data[pos + 1:pos + 1 + n], pos + 1 + n
    if prefix < 0xc0:
        ln = prefix - 0xb7
        n = int.from_bytes(data[pos + 1:pos + 1 + ln], "big")
        s = pos + 1 + ln
        return data[s:s + n], s + n
    if prefix < 0xf8:
        n = prefix - 0xc0
        end = pos + 1 + n
        pos += 1
    else:
        ln = prefix - 0xf7
        n = int.from_bytes(data[pos + 1:pos + 1 + ln], "big")
        pos += 1 + ln
        end = pos + n
    items = []
    while pos < end:
        item, pos = rlp_decode(data, pos)
        items.append(item)
    return items, pos


def receipt_log_addresses(blob):
    """Yield lowercase 0x-addresses, one per log, from a backed-up receipt.

    Handles the three on-disk encodings: BackupReceipt v1
    ([version, type, status, cumGas, bloom, logs] with 4-field logs),
    legacy consensus receipts ([status, cumGas, bloom, logs] with 3-field
    logs), and typed-envelope blobs (an RLP byte string whose payload is
    typeByte || consensus receipt).
    """
    item, _ = rlp_decode(blob)
    if isinstance(item, bytes):  # typed envelope: strip type byte, re-decode
        item, _ = rlp_decode(item[1:])
    if not isinstance(item, list):
        raise ValueError("receipt blob is not an RLP list")
    if len(item) == 6 and item[0] == b"\x01":
        logs = item[5]
    elif len(item) == 4:
        logs = item[3]
    else:
        raise ValueError(f"unrecognized receipt layout ({len(item)} fields)")
    for log in logs:
        yield "0x" + log[0].hex()


def archive_log_counts(archive_dir, tracked, db_head):
    """Per-(tx_hash, addr) log counts for tracked addrs from the backup tree,
    limited to blocks <= db_head. Returns (counts, n_receipts, decode_errors)."""
    counts = {}
    n_receipts = 0
    errors = []
    with os.scandir(archive_dir) as blocks:
        block_dirs = [d for d in blocks if d.is_dir() and d.name.isdigit()
                      and int(d.name) <= db_head]
    for d in sorted(block_dirs, key=lambda d: int(d.name)):
        for entry in os.scandir(d.path):
            if not entry.name.endswith("_receipt.rlp"):
                continue
            tx_hash = entry.name[:-len("_receipt.rlp")].lower()
            n_receipts += 1
            try:
                with open(entry.path, "rb") as f:
                    blob = f.read()
                for addr in receipt_log_addresses(blob):
                    if addr in tracked:
                        counts[(tx_hash, addr)] = counts.get((tx_hash, addr), 0) + 1
            except (ValueError, IndexError, OSError) as exc:
                errors.append(f"{entry.path}: {exc}")
    return counts, n_receipts, errors


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
    # Die quietly when piped into head/grep instead of raising BrokenPipeError.
    signal.signal(signal.SIGPIPE, signal.SIG_DFL)
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

    # -- DB head: the upper bound for log scans and DB-side counts. Not a
    # health check (monitor cg_proc_status directly for indexer lag); the
    # GREATEST covers databases where the checkpoint was never initialized,
    # falling back to the newest block that contained a tracked event.
    db_head = int(db1(
        "SELECT GREATEST((SELECT COALESCE(MAX(block_num), 0) FROM block),"
        " (SELECT COALESCE(MAX(last_block_num), 0) FROM cg_proc_status))"))

    # -- 1. log-count parity per contract ------------------------------------
    # Scan window: from the first indexed cosmic-platform log (= deployment-era
    # constructor events) up to the DB head, counting every tracked contract in
    # one node pass. RandomWalkNFT predates the platform, so its older history
    # is scanned separately from its own first indexed log. Both scans resume
    # from the checkpoint file, so only new blocks are scanned on reruns.
    if os.environ.get("SKIP_LOG_PARITY"):
        add("WARN", "log parity", "skipped (SKIP_LOG_PARITY set)")
    elif os.environ.get("ARCHIVE_DIR"):
        cfg_path = os.environ.get("ARCHIVE_CONFIG")
        if not cfg_path:
            sys.stderr.write("error: ARCHIVE_CONFIG (collector JSON) is"
                             " required with ARCHIVE_DIR\n")
            sys.exit(2)
        with open(cfg_path) as f:
            scope = {a.lower() for a in json.load(f)["contracts"].values()}
        tracked = {a for _, a in labeled}
        arch, n_receipts, errors = archive_log_counts(
            os.environ["ARCHIVE_DIR"], tracked, db_head)
        add("PASS" if not errors else "FAIL",
            f"archive readable ({n_receipts} receipts <= block {db_head})",
            "; ".join(errors[:3]))
        db_groups = {(tx, a): int(n) for tx, a, n in db(
            "SELECT LOWER(t.tx_hash), LOWER(a.addr), COUNT(*)"
            " FROM evt_log e"
            " JOIN transaction t ON t.id = e.tx_id"
            " JOIN address a ON a.address_id = e.contract_aid"
            " WHERE LOWER(a.addr) = ANY(ARRAY["
            + ",".join(f"'{a}'" for a in sorted(tracked)) + "])"
            f" AND e.block_num <= {db_head} GROUP BY 1, 2")}
        for name, addr in labeled:
            if addr not in scope:
                add("WARN", f"log parity (archive): {name}",
                    "not in collector config; not covered by the archive")
                continue
            diffs = []
            a_total = d_total = 0
            txs = ({tx for tx, a in arch if a == addr}
                   | {tx for tx, a in db_groups if a == addr})
            for tx in sorted(txs):
                an = arch.get((tx, addr), 0)
                dn = db_groups.get((tx, addr), 0)
                a_total += an
                d_total += dn
                if an != dn:
                    diffs.append(f"tx {tx}: archive={an} db={dn}")
            add("PASS" if not diffs else "FAIL",
                f"log parity (archive): {name}",
                f"archive={a_total} db={d_total} (..{db_head})"
                + ("; " + "; ".join(diffs[:5]) if diffs else ""))
    else:
        chain_id = int(rpc("eth_chainId", []), 16)
        cosmic_from = int(db1(
            "SELECT MIN(e.block_num) FROM evt_log e"
            " JOIN address a ON a.address_id = e.contract_aid"
            " WHERE LOWER(a.addr) = ANY(ARRAY["
            + ",".join(f"'{a}'" for _, a in labeled if a != rwalk) + "])"))
        all_addrs = sorted(a for _, a in labeled)
        state = load_state()
        counts = scan_logs(state, f"{chain_id}:platform", all_addrs,
                           cosmic_from, db_head)
        rw_from = db1("SELECT MIN(e.block_num) FROM evt_log e"
                      " JOIN address a ON a.address_id = e.contract_aid"
                      f" WHERE LOWER(a.addr) = '{rwalk}'")
        rw_prefix = None
        if rw_from != "" and int(rw_from) < cosmic_from:
            rw_prefix = scan_logs(state, f"{chain_id}:randomwalk-prefix",
                                  [rwalk], int(rw_from), cosmic_from - 1)
        for name, addr in labeled:
            if counts is None or (addr == rwalk and rw_from != ""
                                  and int(rw_from) < cosmic_from
                                  and rw_prefix is None):
                add("WARN", f"log parity: {name}",
                    "checkpoint ahead of DB head; delete state file to rescan")
                continue
            if addr == rwalk and rw_from == "":
                add("WARN", f"log parity: {name}", "no indexed logs; skipped")
                continue
            chain_cnt = counts.get(addr, 0)
            if addr == rwalk and rw_prefix is not None:
                chain_cnt += rw_prefix.get(addr, 0)
            db_cnt = int(db1(
                "SELECT COUNT(*) FROM evt_log e"
                " JOIN address a ON a.address_id = e.contract_aid"
                f" WHERE LOWER(a.addr) = '{addr}' AND e.block_num <= {db_head}"))
            add("PASS" if chain_cnt == db_cnt else "FAIL",
                f"log parity: {name}",
                f"chain={chain_cnt} db={db_cnt} (..{db_head})")

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
