package srvmonitor

import (
	"context"
	"log"
	"os/exec"
	"sync"
	"time"
)

// Alarm notification policy.
const (
	AlarmThreshold       = 5                // Number of occurrences to trigger notification
	AlarmTimeWindow      = 30 * time.Minute // Time window to track alarms
	NotificationCooldown = 60 * time.Minute // Minimum time between notifications for same alarm
	alarmCleanupInterval = 10 * time.Minute // How often stale alarm state is pruned
)

// CommandRunner executes an external command and returns its output.
// Monitors use it for ssh/scp/termux invocations; tests inject fakes.
type CommandRunner func(ctx context.Context, name string, args ...string) ([]byte, error)

// runCombinedOutput is the production CommandRunner: it returns the
// command's combined stdout+stderr, so failures carry their error text.
func runCombinedOutput(ctx context.Context, name string, args ...string) ([]byte, error) {
	return exec.CommandContext(ctx, name, args...).CombinedOutput() //nolint:gosec // G204: command and args come from operator-provided monitor config
}

// runStdout is a CommandRunner returning only stdout, for commands whose
// output is parsed (stderr noise like ssh banners would corrupt it).
func runStdout(ctx context.Context, name string, args ...string) ([]byte, error) {
	return exec.CommandContext(ctx, name, args...).Output() //nolint:gosec // G204: command and args come from operator-provided monitor config
}

// AlarmTracker tracks alarms and sends notifications once an alarm repeats
// past AlarmThreshold within AlarmTimeWindow, at most once per
// NotificationCooldown.
type AlarmTracker struct {
	enabled          bool
	occurrences      map[string][]time.Time // Maps alarm message to timestamps
	lastNotification map[string]time.Time   // Last notification time for each alarm
	mutex            sync.Mutex
	logger           *log.Logger

	now             func() time.Time // nil means the wall clock; tests inject
	run             CommandRunner
	cleanupInterval time.Duration
}

// NewAlarmTracker creates a new alarm tracker. When enabled is false every
// operation is a no-op apart from logging.
func NewAlarmTracker(enabled bool, logger *log.Logger) *AlarmTracker {
	return &AlarmTracker{
		enabled:          enabled,
		occurrences:      make(map[string][]time.Time),
		lastNotification: make(map[string]time.Time),
		logger:           logger,
		run:              runCombinedOutput,
		cleanupInterval:  alarmCleanupInterval,
	}
}

// clock returns the tracker's notion of "now".
func (a *AlarmTracker) clock() time.Time {
	if a.now != nil {
		return a.now()
	}
	return time.Now()
}

// RecordAlarm records an alarm occurrence and sends a notification when the
// threshold is reached outside the cooldown window.
func (a *AlarmTracker) RecordAlarm(ctx context.Context, message string) {
	if !a.enabled || message == "" {
		return
	}

	a.mutex.Lock()
	defer a.mutex.Unlock()

	now := a.clock()

	// Clean up old occurrences outside the time window
	a.cleanOldOccurrences(message, now)

	// Add new occurrence
	a.occurrences[message] = append(a.occurrences[message], now)

	// Check if we should send notification
	occCount := len(a.occurrences[message])

	a.logger.Printf("Alarm recorded: '%s' (count in last 30min: %d)", message, occCount)

	// Only send notification if:
	// 1. We've reached the threshold (5 occurrences)
	// 2. We haven't sent a notification for this alarm recently (cooldown period)
	if occCount >= AlarmThreshold {
		lastNotif, exists := a.lastNotification[message]

		if !exists || now.Sub(lastNotif) >= NotificationCooldown {
			a.sendNotification(ctx, message, occCount)
			a.lastNotification[message] = now
		} else {
			a.logger.Printf("Notification suppressed (cooldown): %v remaining",
				NotificationCooldown-now.Sub(lastNotif))
		}
	}
}

// cleanOldOccurrences removes occurrences outside the time window.
func (a *AlarmTracker) cleanOldOccurrences(message string, now time.Time) {
	timestamps, exists := a.occurrences[message]
	if !exists {
		return
	}

	cutoff := now.Add(-AlarmTimeWindow)
	filtered := make([]time.Time, 0, len(timestamps))
	for _, ts := range timestamps {
		if ts.After(cutoff) {
			filtered = append(filtered, ts)
		}
	}
	a.occurrences[message] = filtered
}

// SendTestNotification sends a test notification (bypasses threshold checks).
func (a *AlarmTracker) SendTestNotification(ctx context.Context, message string) {
	if !a.enabled {
		a.logger.Printf("Cannot send test notification: mobile notifications disabled")
		return
	}
	a.sendNotification(ctx, message, 0)
}

// sendNotification sends an Android notification or alert.
func (a *AlarmTracker) sendNotification(ctx context.Context, message string, count int) {
	a.logger.Printf("Sending notification: '%s' (triggered %d times)", message, count)

	// Try termux-notification first (F-Droid version)
	_, err := a.run(ctx, "termux-notification",
		"--title", "Server Monitor Alert",
		"--content", message,
		"--priority", "high",
		"--sound")
	if err == nil {
		a.logger.Printf("Notification sent successfully")
		return
	}

	// If termux-notification fails, try termux-vibrate (works on Google Play if installed)
	a.logger.Printf("termux-notification failed: %v, trying termux-vibrate fallback", err)

	if _, vibrateErr := a.run(ctx, "termux-vibrate", "-d", "1000"); vibrateErr != nil {
		a.logger.Printf("termux-vibrate also failed: %v", vibrateErr)
		// Last resort: terminal beep (BEL character)
		a.logger.Printf("\a\a\a ALERT (no notification sent): %s", message)
		return
	}
	a.logger.Printf("Vibration alert sent successfully")
}

// RunCleanup prunes stale alarm state every cleanup interval until ctx is
// cancelled. The legacy cleanup goroutine had no stop condition.
func (a *AlarmTracker) RunCleanup(ctx context.Context) {
	ticker := time.NewTicker(a.cleanupInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			a.CleanupOldData()
		}
	}
}

// CleanupOldData removes alarm occurrences outside the tracking window and
// notification timestamps older than 24 hours.
func (a *AlarmTracker) CleanupOldData() {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	now := a.clock()

	// Clean up occurrences
	for message := range a.occurrences {
		a.cleanOldOccurrences(message, now)

		// Remove empty entries
		if len(a.occurrences[message]) == 0 {
			delete(a.occurrences, message)
		}
	}

	// Clean up old notification timestamps (keep only last 24 hours)
	cutoff := now.Add(-24 * time.Hour)
	for message, lastNotif := range a.lastNotification {
		if lastNotif.Before(cutoff) {
			delete(a.lastNotification, message)
		}
	}
}
