package handler

import (
	"context"
	"errors"
	"kraneapi/graph/model"
	service "kraneapi/pkg/api/services"
	"kraneapi/pkg/utils"
)

func CreateEvent(ctx context.Context, input model.AddEventInput) (*model.Event, error) {
	user, err := utils.GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	id, err := service.CreateEvent(ctx, input)
	if err != nil {
		return nil, err
	}

	// Add the user as Admin for the newly created event
	organizerinput := model.CreateEventOrganizerInput{
		EventID: id,
		UserID:  user.UserID,
		Role:    model.RoleAdmin,
	}
	_, err = service.CreateEventOrganizer(ctx, organizerinput)
	if err != nil {
		return nil, err
	}

	event := model.Event{
		ID:          id,
		EventName:   input.EventName,
		StartDate:   input.StartDate,
		EndDate:     input.EndDate,
		Location:    input.Location,
		Description: input.Description,
	}

	return &event, nil
}

func UpdateEvent(ctx context.Context, input model.UpdateEventInput) (*model.Event, error) {
	// Check if the user has permission to update the event
	canUpdate, err := utils.CheckPermission(ctx, input.ID, "events", "update")
	if err != nil {
		return nil, err
	}
	if !canUpdate {
		return nil, errors.New("permission denied")
	}

	id, err := service.UpdateEvent(ctx, input)
	if err != nil {
		return nil, err
	}

	event := model.Event{
		ID:          id,
		EventName:   input.EventName,
		StartDate:   input.StartDate,
		EndDate:     input.EndDate,
		Location:    input.Location,
		Description: input.Description,
	}

	return &event, nil
}
