# RWCG

This repository contains the RandomWalk and CosmicGame projects (blockchain applications on Arbitrum)
(all code previous to CosmicGame/RandomWalk was moved to `previous-code` directory)
## rwcg Overview

The `rwcg/` directory provides the complete backend infrastructure:

| Component | Description |
|-----------|-------------|
| **websrv** | HTTP/HTTPS web server with JSON APIs and HTML pages |
| **etl** | Event extraction pipelines for CosmicGame and RandomWalk smart contracts |
| **dbs** | PostgreSQL database layer and query operations |
| **contracts** | Go bindings for smart contracts (generated via abigen) |
| **primitives** | Shared types and utilities |
| **notibot** | Notification bot for Twitter and Discord |
| **freezer-scanner** | Geth freezer database reader for historical data |
| **tweets/wanotif** | Twitter and WhatsApp notification integrations |
| **tools** | CLI utilities for verification and data export |

## Build

```bash
cd rwcg
make all      # builds etl, websrv, notibot
make tools    # builds CLI tools
```

## See Also

- [rwcg/README.md](rwcg/README.md) — detailed documentation on the web server and API endpoints
