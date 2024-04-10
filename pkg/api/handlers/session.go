package handler

import (
	"context"
	"kraneapi/graph/model"
	"kraneapi/pkg/api/dbmodel"
	"kraneapi/pkg/db"
	"time"

	"github.com/doug-martin/goqu/v9"
)

func CreateEventSession(ctx context.Context, input model.CreateEventSessionInput) (*model.EventSession, error) {
	startTime, err := time.Parse(time.RFC3339, input.StartTime)
	if err != nil {
		return nil, err
	}

	endTime, err := time.Parse(time.RFC3339, input.EndTime)
	if err != nil {
		return nil, err
	}

	// Prepare the insert statement
	insert := db.GetDB().Insert("event_sessions").
		Cols("event_id", "session_name", "start_time", "end_time").
		Vals(goqu.Vals{input.EventID, input.SessionName, startTime, endTime}).
		Returning(goqu.C("session_id")).
		Executor()

	var id string
	if _, err := insert.ScanVal(&id); err != nil {
		return nil, err
	}

	// Convert the inserted row's data to an EventSession model
	eventSession := model.EventSession{
		SessionID:   id,
		EventID:     input.EventID,
		SessionName: input.SessionName,
		StartTime:   input.StartTime,
		EndTime:     input.EndTime,
	}

	return &eventSession, nil
}

func GetAllEventSessions(ctx context.Context) ([]*model.EventSession, error) {
	var eventSessions []*dbmodel.EventSession

	err := db.GetDB().From("event_sessions").Select("*").ScanStructs(&eventSessions)

	if err != nil {
		return nil, err
	}

	var sessions []*model.EventSession
	for _, sessionModel := range eventSessions {
		session := model.EventSession{
			SessionID:   sessionModel.SessionID,
			EventID:     sessionModel.EventID,
			SessionName: sessionModel.SessionName,
			StartTime:   sessionModel.StartTime,
			EndTime:     sessionModel.EndTime,
		}
		sessions = append(sessions, &session)
	}

	return sessions, nil
}

func UpdateEventSession(ctx context.Context, input model.UpdateEventSessionInput) (*model.EventSession, error) {

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
		return nil, err
	}

	return &model.EventSession{
		SessionID:   id,
		EventID:     input.EventID,
		SessionName: input.SessionName,
		StartTime:   input.StartTime,
		EndTime:     input.EndTime,
	}, nil
}
