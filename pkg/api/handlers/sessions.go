package handler

import (
	"context"
	"kraneapi/graph/model"
	service "kraneapi/pkg/api/services"
)

func CreateEventSession(ctx context.Context, input model.CreateEventSessionInput) (*model.EventSession, error) {
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
