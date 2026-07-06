# Testing Notifications with SIGUSR1

## Overview

You can test if Android notifications are working by sending a `SIGUSR1` signal to the running srvmonitor process. This will immediately send a test notification without waiting for real errors.

## How to Test

### 1. Start srvmonitor with notifications enabled

```bash
export MOBILE_NOTIF=yes
./srvmonitor
```

### 2. Find the process ID

In another terminal:

```bash
# Find the PID
ps aux | grep srvmonitor
# or
pgrep srvmonitor
```

### 3. Send SIGUSR1 signal

```bash
kill -SIGUSR1 <PID>

# Example:
kill -SIGUSR1 12345
```

### 4. Check the result

You should:
- **See a notification** on your Android device with the title "Server Monitor Alert"
- **See logs** in `/tmp/srvmonitor.log` (or `$TMPDIR/srvmonitor.log`):
  ```
  SIGUSR1 signal received, sending test notification
  SIGUSR1 received: Sending test notification
  Sending notification: 'TEST: Server Monitor notification test (triggered by SIGUSR1)' (triggered 0 times)
  Notification sent successfully
  ```

## One-Liner Test

```bash
# Start srvmonitor in background, wait a bit, send test signal
./srvmonitor & sleep 5 && kill -SIGUSR1 $! && fg
```

## If Notifications Don't Appear

1. **Check if MOBILE_NOTIF is enabled:**
   ```bash
   echo $MOBILE_NOTIF
   ```
   Should show `yes`

2. **Check the logs:**
   ```bash
   tail -f /tmp/srvmonitor.log
   # or
   tail -f $TMPDIR/srvmonitor.log
   ```

3. **Verify termux-notification works directly:**
   ```bash
   termux-notification --title "Test" --content "Hello" --sound
   ```

4. **Make sure you have:**
   - F-Droid version of Termux (not Google Play)
   - Termux:API app installed from F-Droid
   - Termux API package: `pkg install termux-api`

## Production Use

In production, real notifications are sent when:
- The **same error** occurs **5 times** within **30 minutes**
- At least **60 minutes** have passed since the last notification for that specific error

The SIGUSR1 test signal bypasses these thresholds and sends an immediate notification.

## Quick Test Script

Save this as `test-notif.sh`:

```bash
#!/bin/bash
echo "Starting srvmonitor with notifications enabled..."
export MOBILE_NOTIF=yes
./srvmonitor &
PID=$!
echo "srvmonitor started with PID: $PID"
sleep 5
echo "Sending test notification signal..."
kill -SIGUSR1 $PID
echo "Check your phone for notification!"
echo "Press Ctrl+C to stop srvmonitor"
wait $PID
```

Make it executable and run:
```bash
chmod +x test-notif.sh
./test-notif.sh
```




