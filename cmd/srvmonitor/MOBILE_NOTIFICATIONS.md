# Mobile Notifications for Server Monitor

## Overview

The server monitoring app now supports Android notifications via Termux API. When errors are detected (servers down, databases unreachable, etc.), the app can send notifications to your Android device.

## Features

- **Threshold-based notifications**: Only sends notifications after an alarm has been triggered 5 times in the last 30 minutes
- **Cooldown period**: Prevents spam by waiting 60 minutes between notifications for the same alarm
- **Automatic cleanup**: Removes old alarm data to prevent memory buildup
- **Configurable**: Can be easily enabled or disabled via environment variable

## Requirements

1. **Termux** - The terminal emulator for Android
2. **Termux:API** - Install from F-Droid or Google Play Store
3. **termux-api package** - Install in Termux with:
   ```bash
   pkg install termux-api
   ```

## Configuration

Enable mobile notifications by setting the environment variable:

```bash
export MOBILE_NOTIF=yes
```

Or disable it:

```bash
export MOBILE_NOTIF=no
# or simply leave it unset
```

Valid values for enabling: `yes`, `true`, `1` (case-insensitive)

## How It Works

1. **Error Detection**: All monitors (RPC, Database, Web API, etc.) report errors to a central error channel
2. **Alarm Tracking**: Each error is recorded with a timestamp
3. **Threshold Check**: When the same error occurs 5 times within 30 minutes, a notification is triggered
4. **Cooldown**: After sending a notification, the system waits 60 minutes before sending another notification for the same alarm
5. **Cleanup**: Old alarm data (older than 30 minutes for occurrences, 24 hours for notification history) is automatically cleaned up

## Notification Details

When a notification is sent, it includes:
- **Title**: "Server Monitor Alert"
- **Content**: The error message (e.g., "Block difference is zero (last block = 214872827)")
- **Priority**: High
- **Sound**: Enabled

## Testing

To test if notifications work:

1. Make sure Termux:API is installed
2. Run in terminal:
   ```bash
   termux-notification --title "Test" --content "Hello from Termux" --sound
   ```
3. You should see a notification on your Android device

## Troubleshooting

**Notifications not appearing:**
- Ensure Termux:API app is installed
- Check that `termux-api` package is installed: `pkg list-installed | grep termux-api`
- Verify notification permissions are granted to Termux
- Check the log file (`/tmp/srvmonitor.log`) for error messages

**Too many notifications:**
- The default threshold is 5 occurrences in 30 minutes
- The cooldown period is 60 minutes between notifications for the same alarm
- These can be adjusted in `monitor/alarm_tracker.go` if needed

## Constants (Adjustable)

In `monitor/alarm_tracker.go`:
- `AlarmThreshold`: Number of occurrences to trigger notification (default: 5)
- `AlarmTimeWindow`: Time window to track alarms (default: 30 minutes)
- `NotificationCooldown`: Minimum time between notifications for same alarm (default: 60 minutes)

## Example

1. RPC node goes down
2. Error "Block difference is zero" is logged
3. Over the next 30 minutes, this error occurs 5 times
4. On the 5th occurrence, you receive an Android notification
5. Even if the error continues, you won't get another notification for this specific error for 60 minutes
6. Once the cooldown expires, if the alarm is still occurring, you'll get another notification






