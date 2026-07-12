# notibot

Notification bot: announces RandomWalk NFT events (mints, marketplace offers
and purchases, floor-price changes) on Twitter and Discord, and keeps the
Discord statistics channels (mint count, current price, last mint, last
reward) up to date.

The monitoring engine lives in `internal/notify/rwbot` and is shared with
`rwctl notify-bot` (the Twitter-only variant); this binary adds the Discord
sink and the ffmpeg video-resampling adapter, and posts the token video as a
reply to every mint tweet.

## Running

```bash
notibot --twitter --discord
```

At least one of `--twitter` / `--discord` is required.

## Configuration

| Variable | Description |
|----------|-------------|
| `RPC_URL` | Ethereum/Arbitrum JSON-RPC endpoint (required) |
| `PGSQL_HOST`, `PGSQL_USERNAME`, `PGSQL_PASSWORD`, `PGSQL_DATABASE` | PostgreSQL connection (required) |
| `TWITTER_KEYS_FILE` | JSON credentials file under `$HOME/configs/` (`--twitter`) |
| `DISCORD_KEYS_FILE` | JSON credentials file under `$HOME/configs/` (`--discord`) |

The Discord credentials file carries the bot token plus the notification and
statistics channel ids:

```json
{
  "TokenKey": "...",
  "ChannelId": 0,
  "MainChannelId": 0,
  "MintStatsChanId": 0,
  "PriceStatsChanId": 0,
  "DateStatsChanId": 0,
  "RewardStatsChanId": 0
}
```

Statistics channels with id 0 are skipped. Discord bot permission required to
update the statistics channels: Manage Channel, Connect. Recommended
permissions for other users so they stay out of the statistical channels:
View Channel yes, Manage Channel no, Connect no.

Logs go to `$HOME/ae_logs/notibot_{info,error,db}.log`. The bot resumes from
the `rw_messaging_status` watermark persisted after every processed event, so
restarts never re-announce history.
