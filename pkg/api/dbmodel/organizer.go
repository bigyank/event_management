package dbmodel

type EventOrganizer struct {
	EventOrganizerID int    `db:"event_organizer_id"`
	EventID          int    `db:"event_id"`
	UserID           int    `db:"user_id"`
	Role             string `db:"role"`
}
