# site-checker

Health checker for our production sites:

- https://app.cosmicsignature.com/ (cosmicgame frontend)
- https://randomwalknft.com/
- https://axiomzero.market/

Plain `curl` cannot catch most real-world failures because the sites are React
apps: the HTML document returns 200 even when the app is broken. This checker
loads each site in a **real headless Chromium** (Playwright), lets React
hydrate and fire its API/RPC calls, and collects everything that goes wrong.

## What it checks

For each site, rendered in the browser:

1. **Uncaught JS exceptions** (`pageerror`) — a React crash, a broken bundle,
   an API response the code can't handle.
2. **Console errors** — anything the app logs at `error` level, tagged with
   the source URL.
3. **HTTP status >= 400 on any request** the page makes — missing NFT images
   (404), failing API calls (5xx), etc. A **non-200 on the main document
   itself is CRITICAL**.
4. **Transport-level request failures** — DNS, TLS, connection refused,
   timeouts.
5. **Empty body detection** — HTTP 200 but a blank page means React failed to
   mount; flagged CRITICAL.

Independently of the browser:

6. **Direct API probes** — `GET https://a{1,2}.cosmicsignature.com/api/cosmicgame/statistics/dashboard`
   must return 200.
7. **Direct RPC probes** — JSON-RPC `eth_blockNumber` POSTed to
   `https://rpc{1,2}.cosmicsignature.com:38546` must return 200 with a valid
   result.
8. **SSL certificates** — every host (sites, API servers, RPC servers) is
   checked for chain validity and expiry (warning under 14 days, critical if
   expired or untrusted).

## False-positive filtering

Findings whose message matches any regex in `ignoreGlobal` (in `config.json`)
or in a site's own `ignore` list are reported separately as "benign" and do
**not** trigger the alarm. The shipped list covers the known noise:

- WalletConnect / Reown / web3modal "projectId not on Allowlist" 403s
  (`pulse.walletconnect.org`, `api.web3modal.org`, etc.)
- wagmi / MetaMask "no injected provider" complaints (no wallet extension
  exists in a headless browser)
- `net::ERR_ABORTED` — Next.js route prefetches and media preloads that the
  browser cancels on purpose; not a server failure. Genuine failures of those
  same URLs surface as HTTP-status or connection errors instead, which are
  still caught.

When a new benign error appears, add a regex for it to `ignoreGlobal`
(applies everywhere) or to the site's `ignore` array (that site only).
Matching is case-insensitive against the full finding text, which includes
the URL.

## Usage

```bash
cd site-checker
npm install                      # once
npx playwright install chromium  # once, downloads the headless browser

node check-sites.js                          # check everything
node check-sites.js --only axiomzero         # single site, skips endpoint/SSL probes
node check-sites.js --json report.json       # also write machine-readable report
node check-sites.js --json -                 # JSON report to stdout, human text to stderr
                                             # (used by srvmonitor over a pipe)
node check-sites.js --headed                 # watch the browser (debugging)
node check-sites.js --config other.json      # alternate config
```

## Exit codes (for cron / alerting)

| Code | Meaning                                             |
|------|-----------------------------------------------------|
| 0    | All checks passed (benign findings may be ignored)  |
| 1    | ALARM — authentic problems found                    |
| 2    | The checker itself failed to run                    |

Example cron entry that emails on failure (cron mails output when a command
prints and exits non-zero, if MAILTO is set):

```cron
MAILTO=you@example.com
*/15 * * * * cd /home/niko/eth/dev/b/backend/tools/site-checker && node check-sites.js > /tmp/site-check.log 2>&1 || cat /tmp/site-check.log
```

Or hook any alerting command to the exit code:

```bash
node check-sites.js || ./send-alarm.sh "$(cat /tmp/site-check.log)"
```

## Config reference (`config.json`)

- `sites[]` — `name`, `url`, optional `ignore` (per-site regex list)
- `endpoints[]` — `name`, `url`, `type`: `"http"` (GET, expect 200) or
  `"jsonrpc"` (POST `eth_blockNumber`, expect 200 + result)
- `sslHosts[]` — `host:port` list for certificate checks
- `ignoreGlobal[]` — case-insensitive regexes applied to every finding
- `navTimeoutMs` — navigation / probe timeout (default 45000)
- `settleMs` — how long to let React settle after load (default 8000)
- `sslWarnDays` — warn when a certificate expires within this many days
