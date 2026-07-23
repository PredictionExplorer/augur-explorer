# srvmonitor

Interactive terminal dashboard monitoring the operational estate: RPC nodes,
PostgreSQL databases, ETL progress, web APIs, disk usage, SSL certificate
expiry, the RandomWalk image server and web-server anomalies.

All monitoring logic lives in [`internal/srvmonitor`](../../internal/srvmonitor)
(monitors, manager, alarm tracker, layout, configuration) with the termbox UI
isolated in `internal/srvmonitor/termboxui`. This binary wires environment
configuration, the terminal and signal handling together; every monitor is
tested against fakes, httptest servers, a fake chain and testcontainers
PostgreSQL — see the package's test suite.

## Running

```bash
# Set environment variables (see below), then:
./srvmonitor

# Exit: press 'q' or Ctrl+C. SIGUSR1 sends a test mobile notification.
```

## Configuration (environment)

| Variables | Monitor |
|---|---|
| `RPC0_NAME`/`RPC0_URL`/`RPC0_CHAINID` … `RPC9_*` (≥1 required) | RPC nodes: head must advance between two reads |
| `OFFICIAL_RPC_MAINNET` / `_ARBITRUM` / `_SEPOLIA` / `_SEPOLIA_ARB` | Names of the official reference nodes for lag calculation |
| `DB_L1_NAME_SRV1`/`DB_L1_HOST_SRV1`/`DB_L1_DBNAME_SRV1`/`DB_L1_USER_SRV1`/`DB_L1_PASS_SRV1` … `SRV4` (≥1 required) | Layer 1 databases: `block` table must grow |
| `APP_STATUS_SRV1_TITLE`/`_HOST`/`_DBNAME`/`_USER`/`_PASS` … `SRV4` | Application DBs: last indexed block + lag vs official node |
| `DB_L1EVT1_NAME`/`_HOST`/`_DBNAME`/`_USER`/`_PASS`/`_TABLE`/`_COLUMN` … `DB_L1EVT6_*` | Up to six event tables: a status column must increase |
| `SRV1_WEB_API_NAME`/`_HOST`/`_PORT`/`_URI`/`_PUBLIC_URL` … `SRV6` | Web APIs: internal probe (lenient) + optional public TLS probe (strict 200) |
| `SSH_CMD_DF_SRV1_NAME`/`_USER`/`_IP`/`_DEVICES` … `SRV3` | Disk usage via `ssh … df` |
| `SSL_CERT1_HOST`/`_PORT`/`_NAME`/`_SERVERNAME` … `SSL_CERT12` | Certificate expiry (warn at ≤10 days) |
| `ANOMALY_SSH_USER`/`ANOMALY_SSH_HOST`/`ANOMALY_REMOTE_FILE`/`ANOMALY_TITLE`; `ANOMALY_STALE_SECS` (default `1800`) | Web-server anomaly file fetched via scp; heartbeat age detects a stopped `loganomaly` producer |
| `DB_RWLK_*_SRV`, `RWALK_CONTRACT_ADDR` (+ `RPC1_URL`) | RandomWalk thumbnails: latest mints and a random spot-check must exist; DB and contract token ids must match |
| `MOBILE_NOTIF` (`yes`/`true`/`1`) | Android alerts via termux-notification (see `MOBILE_NOTIFICATIONS.md`) |
| `TMPDIR` | Log and anomaly-file directory (default `/tmp`) |

Web API probes are part of the v1 sunset measurement. Configuration rejects
deprecated `/api/cosmicgame/*` and `/api/randomwalk/*` probe paths so a
monitor cannot silently keep `deprecated="true"` traffic alive. Use
readiness for the internal process/DB path and a stable DB-backed v2 resource
for the public TLS/proxy path:

```sh
SRV1_WEB_API_NAME=cosmic-api
SRV1_WEB_API_HOST=127.0.0.1
SRV1_WEB_API_PORT=8080
SRV1_WEB_API_URI=/readyz
SRV1_WEB_API_PUBLIC_URL=https://api.example/api/v2/cosmicgame/rounds?limit=1
```

Do not use contract-state resources such as
`/api/v2/cosmicgame/rounds/current` as uptime probes: they deliberately
return 503 while the RPC-backed cache is unavailable. `opsctl smoketest`
checks those resources as part of the full v2 degradation suite.

`loganomaly` prepends `#TS=<unix-seconds>` on every successful generation,
including runs with no anomalies. The marker is metadata and never appears
as an anomaly row. A valid marker older than `ANOMALY_STALE_SECS` paints a
red `STALE` age and emits the same alarm key on every poll; markerless or
malformed-marker files remain supported as legacy feeds with no staleness
claim. SCP/read failures clear prior freshness state and are reported as
fetch/read errors, not as stale data.

## Alerting

Monitor failures flow into a central manager: they are logged, painted on
the shared error area, and — with `MOBILE_NOTIF` — turned into Android
notifications once an alarm repeats 5 times within 30 minutes (one
notification per hour per alarm). `NOTIFICATION_TESTING.md` describes the
SIGUSR1 test flow.

## Logging

Logs are written to `$TMPDIR/srvmonitor.log` (rotated to
`srvmonitor-old.log` on exit): startup configuration summary, monitor
registration, layout positions and every monitor error.
