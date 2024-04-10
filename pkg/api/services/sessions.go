package service

import (
	"context"
	"kraneapi/graph/model"
	"kraneapi/pkg/api/dbmodel"
	"kraneapi/pkg/db"
	"time"

	"github.com/doug-martin/goqu/v9"
)

// CreateEventSession creates a new event session in the database.
func CreateEventSession(ctx context.Context, input model.CreateEventSessionInput) (string, error) {
	startTime, err := time.Parse(time.RFC3339, input.StartTime)
	if err != nil {
		return "", err
	}

	endTime, err := time.Parse(time.RFC3339, input.EndTime)
	if err != nil {
		return "", err
	}

	insert := db.GetDB().Insert("event_sessions").
		Cols("event_id", "session_name", "start_time", "end_time").
		Vals(goqu.Vals{input.EventID, input.SessionName, startTime, endTime}).
		Returning(goqu.C("session_id")).
		Executor()

	var id string
	if _, err := insert.ScanVal(&id); err != nil {
		return "", err
	}

	return id, nil
}

// GetAllEventSessions retrieves all event sessions from the database.
func GetAllEventSessions(ctx context.Context, eventID string) ([]*dbmodel.EventSession, error) {
	var eventSessions []*dbmodel.EventSession
	err := db.GetDB().From("event_sessions").Where(goqu.C("event_id").Eq(eventID)).Select("*").ScanStructs(&eventSessions)
	if err != nil {
		return nil, err
	}
	return eventSessions, nil
}

// UpdateEventSession updates an existing event session in the database.
func UpdateEventSession(ctx context.Context, input model.UpdateEventSessionInput) (string, error) {
	update := db.GetDB().Update("event_sessions").
		Set(map[string]interface{}{
			"session_name": input.SessionName,
			"start_time":   input.StartTime,
			"end_time":     input.EndTime,
		}).
		Where(goqu.C("session_id").Eq(input.SessionID), goqu.C("event_id").Eq(input.EventID)).
		Returning(goqu.C("session_id")).
		Executor()

	var id string
	if _, err := update.ScanVal(&id); err != nil {
		return "", err
	}

	return id, nil
}
