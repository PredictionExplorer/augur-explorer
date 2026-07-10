// Package cosmicgame provides PostgreSQL persistence for CosmicGame events,
// aggregates and API read models.
package cosmicgame

const maxDonationPageLimit = 200

// DonationPageCursor identifies the last immutable event returned by a
// newest-first round donation page.
type DonationPageCursor struct {
	EventLogID int64
}
