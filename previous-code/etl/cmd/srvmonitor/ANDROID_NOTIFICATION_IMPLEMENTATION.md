# Android Notification Implementation Summary

## Files Modified

### 1. `config/config.go`
- Added `MobileNotif bool` field to Config struct
- Added environment variable loading for `MOBILE_NOTIF`
- Accepts: `yes`, `true`, `1` (case-insensitive) to enable
- Logs the mobile notification status on startup

### 2. `monitor/manager.go`
- Updated `Manager` struct to include `alarmTracker *AlarmTracker`
- Modified `NewManager()` to accept `mobileNotifEnabled bool` parameter
- Creates and initializes `AlarmTracker` when notifications are enabled
- Starts cleanup routine for alarm data
- Modified `handleErrors()` to record alarms via `alarmTracker.RecordAlarm(err)`

### 3. `main.go`
- Updated manager creation to pass `cfg.MobileNotif` to `NewManager()`
- Added logging of mobile notification status during startup

## Files Created

### 4. `monitor/alarm_tracker.go` (NEW)
Complete alarm tracking and notification system with:

**Key Features:**
- Tracks alarm occurrences with timestamps
- Counts occurrences in a 30-minute sliding window
- Sends notification after 5 occurrences of the same alarm
- Implements 60-minute cooldown between notifications for the same alarm
- Automatic cleanup of old data every 10 minutes
- Thread-safe with mutex protection

**Main Components:**
- `AlarmTracker` struct: Manages all alarm state
- `RecordAlarm()`: Records alarm and checks if notification should be sent
- `sendNotification()`: Uses termux-notification to send Android notification
- `cleanOldOccurrences()`: Removes outdated alarm records
- `CleanupOldData()`: Periodic cleanup of old data
- `StartCleanupRoutine()`: Background goroutine for cleanup

**Constants:**
- `AlarmThreshold = 5`: Occurrences needed to trigger notification
- `AlarmTimeWindow = 30 minutes`: Time window for counting occurrences
- `NotificationCooldown = 60 minutes`: Time between notifications for same alarm

**Notification Command:**
```bash
termux-notification --title "Server Monitor Alert" --content "<error message>" --priority high --sound
```

## Documentation Created

### 5. `MOBILE_NOTIFICATIONS.md` (NEW)
Complete user documentation including:
- Overview and features
- Requirements (Termux, Termux:API)
- Configuration instructions
- How it works
- Testing instructions
- Troubleshooting guide
- Adjustable constants

## How It Works (Flow)

1. **Configuration Loading**:
   - App reads `MOBILE_NOTIF` environment variable
   - Creates AlarmTracker with enabled/disabled state

2. **Monitor Startup**:
   - Manager creates AlarmTracker
   - If enabled, starts cleanup routine
   - All monitors report errors to error channel

3. **Error Detection**:
   - Monitor detects error (e.g., "Block difference is zero")
   - Sends error message to error channel
   - Manager receives error in `handleErrors()`

4. **Alarm Recording**:
   - Manager calls `alarmTracker.RecordAlarm(errorMsg)`
   - AlarmTracker records timestamp for this error
   - Cleans up old occurrences (> 30 minutes)
   - Counts recent occurrences

5. **Notification Decision**:
   - If count >= 5 AND cooldown expired: Send notification
   - If count >= 5 BUT cooldown active: Suppress (log only)
   - If count < 5: No action (just track)

6. **Notification Sending**:
   - Executes `termux-notification` command
   - Records notification timestamp
   - Prevents duplicates via cooldown

7. **Cleanup**:
   - Every 10 minutes: Remove old occurrences
   - Remove empty alarm entries
   - Remove old notification timestamps (> 24 hours)

## Benefits

1. **No False Alarms**: Requires 5 occurrences before notifying
2. **No Spam**: 60-minute cooldown prevents notification flood
3. **Memory Efficient**: Automatic cleanup of old data
4. **Thread-Safe**: Proper mutex protection for concurrent access
5. **Configurable**: Easy to enable/disable
6. **Informative**: Full error message in notification
7. **Auditable**: All actions logged

## Testing Recommendations

1. **Test notification system**:
   ```bash
   termux-notification --title "Test" --content "Hello" --sound
   ```

2. **Test with app** (enable notifications):
   ```bash
   export MOBILE_NOTIF=yes
   ./srvmonitor
   ```

3. **Simulate errors**: Stop a monitored service and wait for 5 error occurrences

4. **Check logs**: Review `/tmp/srvmonitor.log` for:
   - "Mobile notifications enabled"
   - "Alarm recorded: '<message>' (count in last 30min: X)"
   - "Sending notification: '<message>' (triggered X times)"
   - "Notification sent successfully"

## Potential Enhancements (Future)

- Make threshold and cooldown configurable via environment variables
- Add notification categories (critical, warning, info)
- Support multiple notification channels
- Add notification history display in UI
- Allow manual notification testing from UI
- Support notification actions (e.g., "Acknowledge", "Ignore for 1 hour")






