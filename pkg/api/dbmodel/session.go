package dbmodel

type EventSession struct {
	SessionID   string `db:"session_id"`
	EventID     string `db:"event_id"`
	SessionName string `db:"session_name"`
	StartTime   string `db:"start_time"`
	EndTime     string `db:"end_time"`
}
