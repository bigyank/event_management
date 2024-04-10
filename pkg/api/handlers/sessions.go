package handler

import (
	"context"
	"errors"
	"kraneapi/graph/model"
	service "kraneapi/pkg/api/services"
	"kraneapi/utils"
)

func CreateEventSession(ctx context.Context, input model.CreateEventSessionInput) (*model.EventSession, error) {

	// Check if the user has permission to add event sessions
	canAdd, err := utils.CheckPermission(ctx, input.EventID, "event_sessions", "add")
	if err != nil {
		return nil, err
	}
	if !canAdd {
		return nil, errors.New("permission denied")
	}
	id, err := service.CreateEventSession(ctx, input)
	if err != nil {
		return nil, err
	}

	eventSession := model.EventSession{
		SessionID:   id,
		EventID:     input.EventID,
		SessionName: input.SessionName,
		StartTime:   input.StartTime,
		EndTime:     input.EndTime,
	}

	return &eventSession, nil
}

func GetAllEventSessions(ctx context.Context, eventID string) ([]*model.EventSession, error) {
	// Check if the user has permission to update the event
	canView, err := utils.CheckPermission(ctx, eventID, "event_sessions", "view")
	if err != nil {
		return nil, err
	}
	if !canView {
		return nil, errors.New("permission denied")
	}

	eventSessions, err := service.GetAllEventSessions(ctx, eventID)
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

	// Check if the user has permission to update the event
	canUpdate, err := utils.CheckPermission(ctx, input.EventID, "event_sessions", "update")
	if err != nil {
		return nil, err
	}
	if !canUpdate {
		return nil, errors.New("permission denied")
	}

	id, err := service.UpdateEventSession(ctx, input)
	if err != nil {
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
