package monitor

import (
	"log"
	"os/exec"
	"sync"
	"time"
)

const (
	AlarmThreshold      = 5              // Number of occurrences to trigger notification
	AlarmTimeWindow     = 30 * time.Minute // Time window to track alarms
	NotificationCooldown = 60 * time.Minute // Minimum time between notifications for same alarm
)

// AlarmOccurrence tracks when an alarm occurred
type AlarmOccurrence struct {
	Timestamp time.Time
	Message   string
}

// AlarmTracker tracks alarms and sends notifications
type AlarmTracker struct {
	enabled           bool
	occurrences       map[string][]time.Time // Maps alarm message to timestamps
	lastNotification  map[string]time.Time   // Last notification time for each alarm
	mutex             sync.Mutex
	logger            *log.Logger
}

// NewAlarmTracker creates a new alarm tracker
func NewAlarmTracker(enabled bool, logger *log.Logger) *AlarmTracker {
	return &AlarmTracker{
		enabled:          enabled,
		occurrences:      make(map[string][]time.Time),
		lastNotification: make(map[string]time.Time),
		logger:           logger,
	}
}

// RecordAlarm records an alarm occurrence and checks if notification should be sent
func (a *AlarmTracker) RecordAlarm(message string) {
	if !a.enabled || message == "" {
		return
	}
	
	a.mutex.Lock()
	defer a.mutex.Unlock()
	
	now := time.Now()
	
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
			a.sendNotification(message, occCount)
			a.lastNotification[message] = now
		} else {
			a.logger.Printf("Notification suppressed (cooldown): %v remaining",
				NotificationCooldown-now.Sub(lastNotif))
		}
	}
}

// cleanOldOccurrences removes occurrences outside the time window
func (a *AlarmTracker) cleanOldOccurrences(message string, now time.Time) {
	timestamps, exists := a.occurrences[message]
	if !exists {
		return
	}
	
	// Filter out timestamps older than the time window
	cutoff := now.Add(-AlarmTimeWindow)
	filtered := make([]time.Time, 0)
	
	for _, ts := range timestamps {
		if ts.After(cutoff) {
			filtered = append(filtered, ts)
		}
	}
	
	a.occurrences[message] = filtered
}

// SendTestNotification sends a test notification (bypasses threshold checks)
func (a *AlarmTracker) SendTestNotification(message string) {
	if !a.enabled {
		a.logger.Printf("Cannot send test notification: mobile notifications disabled")
		return
	}
	a.sendNotification(message, 0)
}

// sendNotification sends an Android notification or alert
func (a *AlarmTracker) sendNotification(message string, count int) {
	a.logger.Printf("Sending notification: '%s' (triggered %d times)", message, count)
	
	// Try termux-notification first (F-Droid version)
	cmd := exec.Command("termux-notification",
		"--title", "Server Monitor Alert",
		"--content", message,
		"--priority", "high",
		"--sound")
	
	err := cmd.Run()
	if err != nil {
		// If termux-notification fails, try termux-vibrate (works on Google Play if installed)
		a.logger.Printf("termux-notification failed: %v, trying termux-vibrate fallback", err)
		
		// Vibrate pattern: 1 second on, 0.5 second off, 1 second on
		vibrateCmd := exec.Command("termux-vibrate", "-d", "1000")
		vibrateErr := vibrateCmd.Run()
		if vibrateErr != nil {
			a.logger.Printf("termux-vibrate also failed: %v", vibrateErr)
			// Last resort: terminal beep (BEL character)
			a.logger.Printf("\a\a\a ALERT (no notification sent): %s", message)
		} else {
			a.logger.Printf("Vibration alert sent successfully")
		}
	} else {
		a.logger.Printf("Notification sent successfully")
	}
}

// CleanupOldData periodically cleans up old data (call this in a goroutine)
func (a *AlarmTracker) CleanupOldData() {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	
	now := time.Now()
	
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

// StartCleanupRoutine starts a background routine to clean up old data
func (a *AlarmTracker) StartCleanupRoutine() {
	go func() {
		ticker := time.NewTicker(10 * time.Minute)
		defer ticker.Stop()
		
		for range ticker.C {
			a.CleanupOldData()
		}
	}()
}

