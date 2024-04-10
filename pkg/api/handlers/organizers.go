package handler

import (
	"context"
	"kraneapi/graph/model"
	service "kraneapi/pkg/api/services"
)

func CreateEventOrganizer(ctx context.Context, input model.CreateEventOrganizerInput) (*model.EventOrganizer, error) {

	id, err := service.CreateEventOrganizer(ctx, input)
	if err != nil {
		return nil, err
	}

	eventOrganizer := model.EventOrganizer{
		EventOrganizerID: id,
		EventID:          input.EventID,
		UserID:           input.UserID,
		Role:             input.Role,
	}

	return &eventOrganizer, nil
}

func UpdateEventOrganizer(ctx context.Context, input model.UpdateEventOrganizerInput) (*model.EventOrganizer, error) {
	id, err := service.UpdateEventOrganizer(ctx, input)
	if err != nil {
		return nil, err
	}

	eventOrganizer := model.EventOrganizer{
		EventOrganizerID: id,
		EventID:          input.EventID,
		UserID:           input.UserID,
		Role:             input.Role,
	}

	return &eventOrganizer, nil
}
