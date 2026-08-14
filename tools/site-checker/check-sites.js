#!/usr/bin/env node
/**
 * Site health checker.
 *
 * Loads each site in a real headless Chromium (via Playwright) so the React
 * app actually runs, and collects:
 *   - uncaught JS exceptions ("pageerror")
 *   - console messages of type "error"
 *   - network requests that failed at the transport level (DNS, TLS, timeout)
 *   - HTTP responses with status >= 400 (a non-200 on the main document is
 *     treated as CRITICAL)
 *
 * Also probes API endpoints and JSON-RPC servers directly, and checks SSL
 * certificate expiry for every host involved.
 *
 * Findings matching the ignore regexes in config.json (global "ignoreGlobal"
 * plus per-site "ignore") are reported separately as benign and do NOT
 * trigger the alarm.
 *
 * Exit codes: 0 = all healthy, 1 = authentic problems found, 2 = checker error.
 *
 * Usage: node check-sites.js [--config config.json] [--json report.json]
 *        [--only <site-name>] [--headed]
 *
 * With "--json -" the JSON report is written to stdout and all human-readable
 * output moves to stderr, so a parent process (e.g. srvmonitor) can read the
 * machine-readable report over a pipe.
 */

'use strict';

const fs = require('fs');
const path = require('path');
const tls = require('tls');
const { chromium } = require('playwright');

// ---------------------------------------------------------------- CLI args

function parseArgs(argv) {
  const args = { config: path.join(__dirname, 'config.json'), json: null, only: null, headed: false };
  for (let i = 2; i < argv.length; i++) {
    const a = argv[i];
    if (a === '--config') args.config = argv[++i];
    else if (a === '--json') args.json = argv[++i];
    else if (a === '--only') args.only = argv[++i];
    else if (a === '--headed') args.headed = true;
    else {
      console.error(`Unknown argument: ${a}`);
      process.exit(2);
    }
  }
  return args;
}

// ------------------------------------------------------------- findings

/**
 * A finding: { severity, kind, site, message }
 *   severity: "critical" | "error" | "warning"
 *   kind: "js-error" | "console-error" | "http-status" | "request-failed"
 *         | "endpoint" | "ssl" | "navigation"
 */
function makeCollector(ignoreRegexes) {
  const authentic = [];
  const ignored = [];
  const add = (finding) => {
    const target = ignoreRegexes.some((re) => re.test(finding.message)) ? ignored : authentic;
    // De-duplicate identical messages (React can spam the same error).
    if (!target.some((f) => f.message === finding.message && f.kind === finding.kind)) {
      target.push(finding);
    }
  };
  return { authentic, ignored, add };
}

function compileIgnores(patterns) {
  return (patterns || []).map((p) => new RegExp(p, 'i'));
}

// --------------------------------------------------------- browser check

async function checkSite(browser, site, cfg, ignoreRegexes) {
  const collector = makeCollector(ignoreRegexes);
  const { add } = collector;
  const context = await browser.newContext({
    userAgent:
      'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) ' +
      'Chrome/131.0.0.0 Safari/537.36 SiteChecker/1.0',
    viewport: { width: 1440, height: 900 },
  });
  const page = await context.newPage();
  page.setDefaultTimeout(cfg.navTimeoutMs);

  page.on('pageerror', (err) => {
    add({
      severity: 'error',
      kind: 'js-error',
      site: site.name,
      message: `Uncaught JS exception: ${err.message}`,
    });
  });

  page.on('console', (msg) => {
    if (msg.type() !== 'error') return;
    // Include the source URL: Chromium's "Failed to load resource" text has no
    // URL of its own, and the ignore patterns need it to filter benign noise.
    const loc = msg.location();
    const where = loc && loc.url ? ` [${loc.url}]` : '';
    add({
      severity: 'error',
      kind: 'console-error',
      site: site.name,
      message: `Console error: ${msg.text()}${where}`,
    });
  });

  page.on('requestfailed', (req) => {
    const failure = req.failure() ? req.failure().errorText : 'unknown failure';
    add({
      severity: 'error',
      kind: 'request-failed',
      site: site.name,
      message: `Request failed (${failure}): ${req.method()} ${req.url()}`,
    });
  });

  let mainStatus = null;
  page.on('response', (resp) => {
    const status = resp.status();
    if (status < 400) return;
    const isMainDoc = resp.url().replace(/\/$/, '') === site.url.replace(/\/$/, '');
    add({
      severity: isMainDoc ? 'critical' : 'error',
      kind: 'http-status',
      site: site.name,
      message: `HTTP ${status} on ${resp.request().method()} ${resp.url()}`,
    });
  });

  try {
    const resp = await page.goto(site.url, { waitUntil: 'domcontentloaded', timeout: cfg.navTimeoutMs });
    mainStatus = resp ? resp.status() : null;
    if (mainStatus !== null && mainStatus !== 200) {
      add({
        severity: 'critical',
        kind: 'http-status',
        site: site.name,
        message: `Main document returned HTTP ${mainStatus} (expected 200) for ${site.url}`,
      });
    }
    // Let React hydrate, fire its API calls, load images, etc.
    await page.waitForLoadState('networkidle', { timeout: cfg.navTimeoutMs }).catch(() => {});
    await page.waitForTimeout(cfg.settleMs);
    // Scroll to the bottom to trigger lazy-loaded content, then settle again.
    await page.evaluate(() => window.scrollTo(0, document.body.scrollHeight)).catch(() => {});
    await page.waitForTimeout(2000);

    // A blank <body> despite HTTP 200 usually means React crashed on mount.
    const bodyLength = await page.evaluate(() => (document.body ? document.body.innerText.trim().length : 0)).catch(() => -1);
    if (bodyLength === 0) {
      add({
        severity: 'critical',
        kind: 'navigation',
        site: site.name,
        message: `Page rendered an empty body (React likely failed to mount) at ${site.url}`,
      });
    }
  } catch (err) {
    add({
      severity: 'critical',
      kind: 'navigation',
      site: site.name,
      message: `Failed to load ${site.url}: ${err.message.split('\n')[0]}`,
    });
  } finally {
    await context.close();
  }

  return { name: site.name, url: site.url, mainStatus, ...collector };
}

// --------------------------------------------------------- direct probes

async function fetchWithTimeout(url, options, timeoutMs) {
  const controller = new AbortController();
  const timer = setTimeout(() => controller.abort(), timeoutMs);
  try {
    return await fetch(url, { ...options, signal: controller.signal });
  } finally {
    clearTimeout(timer);
  }
}

async function checkEndpoint(ep, cfg, ignoreRegexes) {
  const collector = makeCollector(ignoreRegexes);
  const { add } = collector;
  try {
    let resp;
    if (ep.type === 'jsonrpc') {
      resp = await fetchWithTimeout(
        ep.url,
        {
          method: 'POST',
          headers: { 'content-type': 'application/json' },
          body: JSON.stringify({ jsonrpc: '2.0', id: 1, method: 'eth_blockNumber', params: [] }),
        },
        cfg.navTimeoutMs
      );
      if (resp.status !== 200) {
        add({
          severity: 'critical',
          kind: 'endpoint',
          site: ep.name,
          message: `RPC returned HTTP ${resp.status} for ${ep.url}`,
        });
      } else {
        const body = await resp.json().catch(() => null);
        if (!body || typeof body.result !== 'string') {
          add({
            severity: 'critical',
            kind: 'endpoint',
            site: ep.name,
            message: `RPC eth_blockNumber returned unexpected body from ${ep.url}: ${JSON.stringify(body).slice(0, 200)}`,
          });
        }
      }
    } else {
      resp = await fetchWithTimeout(ep.url, { method: 'GET' }, cfg.navTimeoutMs);
      if (resp.status !== 200) {
        add({
          severity: 'critical',
          kind: 'endpoint',
          site: ep.name,
          message: `Endpoint returned HTTP ${resp.status} for ${ep.url}`,
        });
      }
    }
  } catch (err) {
    add({
      severity: 'critical',
      kind: 'endpoint',
      site: ep.name,
      message: `Endpoint unreachable: ${ep.url} (${err.cause ? err.cause.message || err.cause.code : err.message})`,
    });
  }
  return { name: ep.name, url: ep.url, ...collector };
}

function checkSslCert(hostPort, warnDays, timeoutMs) {
  const [host, portStr] = hostPort.split(':');
  const port = Number(portStr || 443);
  return new Promise((resolve) => {
    const findings = [];
    const socket = tls.connect(
      { host, port, servername: host, timeout: timeoutMs },
      () => {
        const cert = socket.getPeerCertificate();
        if (!socket.authorized) {
          findings.push({
            severity: 'critical',
            kind: 'ssl',
            site: hostPort,
            message: `SSL certificate NOT trusted for ${hostPort}: ${socket.authorizationError}`,
          });
        }
        if (cert && cert.valid_to) {
          const expires = new Date(cert.valid_to);
          const daysLeft = Math.floor((expires - Date.now()) / 86400000);
          if (daysLeft < 0) {
            findings.push({
              severity: 'critical',
              kind: 'ssl',
              site: hostPort,
              message: `SSL certificate EXPIRED ${-daysLeft} day(s) ago for ${hostPort} (valid_to: ${cert.valid_to})`,
            });
          } else if (daysLeft < warnDays) {
            findings.push({
              severity: 'warning',
              kind: 'ssl',
              site: hostPort,
              message: `SSL certificate expires in ${daysLeft} day(s) for ${hostPort} (valid_to: ${cert.valid_to})`,
            });
          }
          resolve({ host: hostPort, daysLeft, findings });
        } else {
          findings.push({
            severity: 'critical',
            kind: 'ssl',
            site: hostPort,
            message: `Could not read SSL certificate for ${hostPort}`,
          });
          resolve({ host: hostPort, daysLeft: null, findings });
        }
        socket.end();
      }
    );
    socket.on('timeout', () => {
      socket.destroy();
      findings.push({
        severity: 'critical',
        kind: 'ssl',
        site: hostPort,
        message: `TLS connection timed out to ${hostPort}`,
      });
      resolve({ host: hostPort, daysLeft: null, findings });
    });
    socket.on('error', (err) => {
      findings.push({
        severity: 'critical',
        kind: 'ssl',
        site: hostPort,
        message: `TLS connection failed to ${hostPort}: ${err.message}`,
      });
      resolve({ host: hostPort, daysLeft: null, findings });
    });
  });
}

// ---------------------------------------------------------------- report

const COLORS = process.stdout.isTTY
  ? { red: '\x1b[31m', yellow: '\x1b[33m', green: '\x1b[32m', dim: '\x1b[2m', bold: '\x1b[1m', reset: '\x1b[0m' }
  : { red: '', yellow: '', green: '', dim: '', bold: '', reset: '' };

function severityTag(sev) {
  if (sev === 'critical') return `${COLORS.red}${COLORS.bold}[CRITICAL]${COLORS.reset}`;
  if (sev === 'error') return `${COLORS.red}[ERROR]${COLORS.reset}`;
  return `${COLORS.yellow}[WARN]${COLORS.reset}`;
}

function printSection(title, authentic, ignored) {
  console.log(`${COLORS.bold}== ${title} ==${COLORS.reset}`);
  if (authentic.length === 0) {
    console.log(`  ${COLORS.green}OK${COLORS.reset} — no authentic errors`);
  } else {
    for (const f of authentic) {
      console.log(`  ${severityTag(f.severity)} ${f.message}`);
    }
  }
  if (ignored.length > 0) {
    console.log(`  ${COLORS.dim}ignored as benign (${ignored.length}):${COLORS.reset}`);
    for (const f of ignored) {
      console.log(`  ${COLORS.dim}  - ${f.message.slice(0, 200)}${COLORS.reset}`);
    }
  }
  console.log('');
}

// ------------------------------------------------------------------ main

async function main() {
  const args = parseArgs(process.argv);
  const jsonToStdout = args.json === '-';
  if (jsonToStdout) {
    // Keep stdout clean for the JSON report; humans read stderr.
    console.log = console.error.bind(console);
  }
  const cfg = JSON.parse(fs.readFileSync(args.config, 'utf8'));
  const globalIgnores = compileIgnores(cfg.ignoreGlobal);

  const report = {
    startedAt: new Date().toISOString(),
    sites: [],
    endpoints: [],
    ssl: [],
    authenticCount: 0,
    ignoredCount: 0,
  };

  console.log(`${COLORS.bold}Site checker — ${report.startedAt}${COLORS.reset}\n`);

  // 1. Direct endpoint probes (API + RPC) — run in parallel.
  const endpoints = args.only ? [] : cfg.endpoints || [];
  const endpointResults = await Promise.all(endpoints.map((ep) => checkEndpoint(ep, cfg, globalIgnores)));

  // 2. SSL certificate checks — run in parallel.
  const sslHosts = args.only ? [] : cfg.sslHosts || [];
  const sslResults = await Promise.all(sslHosts.map((h) => checkSslCert(h, cfg.sslWarnDays || 14, 15000)));

  // 3. Browser checks — sequential, one context per site.
  const sites = (cfg.sites || []).filter((s) => !args.only || s.name === args.only);
  const browser = await chromium.launch({ headless: !args.headed });
  const siteResults = [];
  try {
    for (const site of sites) {
      const ignores = globalIgnores.concat(compileIgnores(site.ignore));
      console.log(`${COLORS.dim}loading ${site.url} ...${COLORS.reset}`);
      siteResults.push(await checkSite(browser, site, cfg, ignores));
    }
  } finally {
    await browser.close();
  }
  console.log('');

  // ------------------------------------------------------------- output

  for (const r of siteResults) {
    printSection(`Site: ${r.name} (${r.url}) — main document HTTP ${r.mainStatus ?? 'n/a'}`, r.authentic, r.ignored);
    report.sites.push({
      name: r.name,
      url: r.url,
      mainStatus: r.mainStatus,
      authentic: r.authentic,
      ignored: r.ignored,
    });
    report.authenticCount += r.authentic.length;
    report.ignoredCount += r.ignored.length;
  }

  if (endpointResults.length > 0) {
    const all = endpointResults.flatMap((r) => r.authentic);
    const allIgnored = endpointResults.flatMap((r) => r.ignored);
    printSection('Direct endpoint probes (API + RPC)', all, allIgnored);
    report.endpoints = endpointResults.map((r) => ({ name: r.name, url: r.url, authentic: r.authentic, ignored: r.ignored }));
    report.authenticCount += all.length;
  }

  if (sslResults.length > 0) {
    const all = sslResults.flatMap((r) => r.findings);
    console.log(`${COLORS.bold}== SSL certificates ==${COLORS.reset}`);
    for (const r of sslResults) {
      if (r.findings.length === 0) {
        console.log(`  ${COLORS.green}OK${COLORS.reset} ${r.host} — ${r.daysLeft} day(s) until expiry`);
      } else {
        for (const f of r.findings) console.log(`  ${severityTag(f.severity)} ${f.message}`);
      }
    }
    console.log('');
    report.ssl = sslResults;
    report.authenticCount += all.filter((f) => f.severity !== 'warning').length;
  }

  const warnings = sslResults.flatMap((r) => r.findings).filter((f) => f.severity === 'warning').length;
  const failed = report.authenticCount > 0;
  if (failed) {
    console.log(`${COLORS.red}${COLORS.bold}ALARM: ${report.authenticCount} authentic problem(s) found.${COLORS.reset}`);
  } else if (warnings > 0) {
    console.log(`${COLORS.yellow}${COLORS.bold}All checks passed, but ${warnings} warning(s) (see above).${COLORS.reset}`);
  } else {
    console.log(`${COLORS.green}${COLORS.bold}All checks passed. ${report.ignoredCount} benign finding(s) ignored.${COLORS.reset}`);
  }

  report.finishedAt = new Date().toISOString();
  if (jsonToStdout) {
    process.stdout.write(JSON.stringify(report) + '\n');
  } else if (args.json) {
    fs.writeFileSync(args.json, JSON.stringify(report, null, 2));
    console.log(`JSON report written to ${args.json}`);
  }

  process.exit(failed ? 1 : 0);
}

main().catch((err) => {
  console.error(`Checker itself failed: ${err.stack || err}`);
  process.exit(2);
});
