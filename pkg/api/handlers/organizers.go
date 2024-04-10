package handler

import (
	"context"
	"errors"
	"kraneapi/graph/model"
	service "kraneapi/pkg/api/services"
	"kraneapi/pkg/utils"
)

func CreateEventOrganizer(ctx context.Context, input model.CreateEventOrganizerInput) (*model.EventOrganizer, error) {

	// Check if the user has permission to update the event
	canUpdate, err := utils.CheckPermission(ctx, input.EventID, "event_organizers", "add")
	if err != nil {
		return nil, err
	}
	if !canUpdate {
		return nil, errors.New("permission denied")
	}

	// Check if the current user has permission to add the specified role
	canAdd, err := utils.CanAddRole(ctx, input.EventID, string(input.Role))
	if err != nil {
		return nil, err
	}
	if !canAdd {
		return nil, errors.New("permission denied")
	}

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
	// Check if the user has permission to update the event
	canUpdate, err := utils.CheckPermission(ctx, input.EventID, "event_organizers", "update")
	if err != nil {
		return nil, err
	}
	if !canUpdate {
		return nil, errors.New("permission denied")
	}

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
