#!/usr/bin/env python3
"""Generate the cg_error_lookup seed rows from the Solidity sources.

cg_error_lookup maps a signature (32-byte topic0 for events, 4-byte selector
for custom errors, hex without 0x) to the human-readable message the contracts
attach to it. V3 stripped most `errStr` parameters from the custom errors to
save gas, leaving the wording only as inline comments at the revert sites
(`revert X(/* "..." */ ...)`); this script harvests both the live literals and
those commented literals so a revert selector or a failure event can still be
related to its description.

Usage:
    scripts/gen-cg-error-lookup.py [CONTRACTS_PRODUCTION_DIR] > seed.sql

The output is pasted into the migration that owns the table (00034) and is
re-run whenever the contracts add or reword an error.
"""
import re
import sys
from collections import defaultdict
from pathlib import Path

try:
    from Crypto.Hash import keccak
except ImportError:  # pragma: no cover
    sys.exit("pycryptodome is required: pip install pycryptodome")

DEFAULT_ROOT = "/home/niko/eth/dev/b/cg-v3/Cosmic-Signature-v3.1-2026-08-19/contracts/production"

# Events that report a failure. Everything else emitted by the contracts is
# ordinary telemetry and has no message to look up.
FAILURE_EVENTS = {
    "FundTransferFailed",
    "EthTransferToCharityFailed",
    "ArbSysArbBlockNumberCallFailed",
    "ArbSysArbBlockHashCallFailed",
    "ArbGasInfoGetGasBacklogCallFailed",
    "ArbGasInfoGetL1PricingUnitsSinceUpdateCallFailed",
}

# Wording of the parameterless V3.1 Arbitrum failure events (the message the
# pre-V3.1 ArbitrumError(string) carried for the same failure).
FIXED_EVENT_MESSAGES = {
    "ArbSysArbBlockNumberCallFailed": "ArbSys.arbBlockNumber call failed.",
    "ArbSysArbBlockHashCallFailed": "ArbSys.arbBlockHash call failed.",
    "ArbGasInfoGetGasBacklogCallFailed": "ArbGasInfo.getGasBacklog call failed.",
    "ArbGasInfoGetL1PricingUnitsSinceUpdateCallFailed": "ArbGasInfo.getL1PricingUnitsSinceUpdate call failed.",
    "EthTransferToCharityFailed": "ETH transfer to charity failed.",
}

# Signatures that no longer exist in the sources but are present in already
# indexed logs of the deployed V1/V2 contracts.
LEGACY_ROWS = [
    ("event", "ArbitrumError", "ArbitrumError(string)", "CosmicSignatureEvents.sol (pre-V3.1)",
     ["ArbSys.arbBlockNumber call failed.", "ArbSys.arbBlockHash call failed.",
      "ArbGasInfo.getGasBacklog call failed.", "ArbGasInfo.getL1PricingUnitsSinceUpdate call failed."]),
    # The pre-V3 contracts passed the wording at each emit site; it is no
    # longer in the sources, so the row documents the signature only.
    ("event", "ERC20TransferFailed", "ERC20TransferFailed(string,address,uint256)", "CosmicSignatureErrors.sol (pre-V3)", []),
]


def kec(s: str) -> str:
    k = keccak.new(digest_bits=256)
    k.update(s.encode())
    return k.hexdigest()


def canon(t: str) -> str:
    t = re.sub(r"\b(indexed|memory|calldata|storage|payable)\b", "", t).strip()
    t = re.sub(r"\s+", "", t)
    m = re.fullmatch(r"(.+?)(\[.*\])", t)
    suffix = ""
    if m:
        t, suffix = m.group(1), m.group(2)
    if t == "uint":
        t = "uint256"
    elif t == "int":
        t = "int256"
    elif re.fullmatch(r"u?int\d+|bytes\d*|string|bool|address", t):
        pass
    else:
        t = "address"  # contract / interface typed parameter
    return t + suffix


def split_params(s: str):
    out, depth, cur = [], 0, []
    for ch in s:
        if ch == "(":
            depth += 1
        elif ch == ")":
            depth -= 1
        if ch == "," and depth == 0:
            out.append("".join(cur))
            cur = []
        else:
            cur.append(ch)
    out.append("".join(cur))
    return [p.strip() for p in out if p.strip()]


def strip_doc_and_line_comments(src: str) -> str:
    # Keep inline /* "..." */ (that is where V3 parks the message); drop
    # NatSpec blocks and // lines so commented-out code is not harvested.
    src = re.sub(r"/\*\*.*?\*/", "", src, flags=re.S)
    return re.sub(r"//[^\n]*", "", src)


def balanced_args(src: str, start: int) -> str:
    depth, i = 1, start
    while i < len(src) and depth:
        depth += src[i] == "("
        depth -= src[i] == ")"
        i += 1
    return src[start : i - 1]


def main() -> None:
    root = Path(sys.argv[1] if len(sys.argv) > 1 else DEFAULT_ROOT)
    decls = {}  # (kind, name) -> (abi_signature, file)
    messages = defaultdict(set)  # (kind, name) -> {message}

    for path in sorted(root.rglob("*.sol")):
        src = strip_doc_and_line_comments(path.read_text(errors="replace"))
        for m in re.finditer(r"\b(event|error)\s+([A-Za-z_]\w*)\s*\(([^;]*?)\)\s*;", src, flags=re.S):
            kind, name, params = m.groups()
            if kind == "event" and name not in FAILURE_EVENTS:
                continue  # ordinary telemetry (versioned events like BidPlaced legitimately differ)
            params = re.sub(r"/\*.*?\*/", "", params, flags=re.S)
            types = []
            for prm in split_params(params):
                prm = re.sub(r"\s+", " ", prm)
                typ = re.sub(r"\s+[A-Za-z_]\w*$", "", prm) if " " in prm else prm
                types.append(canon(typ))
            sig = f"{name}({','.join(types)})"
            prev = decls.get((kind, name))
            if prev and prev[0] != sig:
                sys.exit(f"{kind} {name}: conflicting declarations {prev[0]} vs {sig}")
            decls[(kind, name)] = (sig, path.name)

    # Use sites: `emit X(...)` feeds the event row; `revert X(...)` and
    # `require(..., X(...))` feed the error row. The same name can be both
    # (FundTransferFailed), so the verb decides.
    site_re = re.compile(r"(?<![\w.])(emit\s+)?(?:CosmicSignatureErrors\.|CosmicSignatureEvents\.)?([A-Z]\w*)\s*\(")
    for path in sorted(root.rglob("*.sol")):
        src = strip_doc_and_line_comments(path.read_text(errors="replace"))
        for m in site_re.finditer(src):
            kind = "event" if m.group(1) else "error"
            name = m.group(2)
            if (kind, name) not in decls:
                continue
            args = balanced_args(src, m.end())
            lit = re.search(r'"((?:[^"\\]|\\.)*)"', args)
            if lit:
                messages[(kind, name)].add(lit.group(1))

    rows = []
    for (kind, name), (sig, fname) in decls.items():
        msgs = sorted(messages.get((kind, name), set()))
        if name in FIXED_EVENT_MESSAGES and not msgs:
            msgs = [FIXED_EVENT_MESSAGES[name]]
        h = kec(sig)
        rows.append((kind, name, sig, h if kind == "event" else h[:8], fname, msgs))
    for kind, name, sig, fname, msgs in LEGACY_ROWS:
        h = kec(sig)
        rows.append((kind, name, sig, h if kind == "event" else h[:8], fname, msgs))
    rows.sort(key=lambda r: (r[0], r[1]))

    def q(s: str) -> str:
        return "'" + s.replace("'", "''") + "'"

    print("INSERT INTO cg_error_lookup(signature, kind, name, abi_signature, source, err_str, err_strs) VALUES")
    lines = []
    for kind, name, sig, key, fname, msgs in rows:
        err_str = q(msgs[0]) if msgs else "NULL"
        arr = "ARRAY[" + ",".join(q(x) for x in msgs) + "]::TEXT[]" if msgs else "'{}'::TEXT[]"
        lines.append(f"\t({q(key)}, {q(kind)}, {q(name)}, {q(sig)}, {q(fname)}, {err_str}, {arr})")
    print(",\n".join(lines) + ";")


if __name__ == "__main__":
    main()
