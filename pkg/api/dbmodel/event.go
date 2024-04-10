package dbmodel

import "time"

// Event represents an event in the database.
type Event struct {
	ID          string    `db:"event_id"`
	EventName   string    `db:"event_name"`
	StartDate   time.Time `db:"start_date"`
	EndDate     time.Time `db:"end_date"`
	Location    string    `db:"location"`
	Description string    `db:"description"`
}
