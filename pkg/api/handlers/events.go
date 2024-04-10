package handler

import (
	"context"
	"kraneapi/graph/model"
	"kraneapi/pkg/db"
	"kraneapi/utils"
	"time"

	"github.com/doug-martin/goqu/v9"
)

func CreateEvent(ctx context.Context, input model.AddEventInput) (*model.Event, error) {
	user, err := utils.GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	start, err := time.Parse(time.RFC3339, input.StartDate)
	if err != nil {
		return nil, err
	}

	end, err := time.Parse(time.RFC3339, input.EndDate)
	if err != nil {
		return nil, err
	}

	insert := db.GetDB().Insert("events").
		Cols("event_name", "start_date", "end_date", "location", "description").
		Vals(goqu.Vals{input.EventName, start, end, input.Location, input.Description}).
		Returning(goqu.C("event_id")).
		Executor()

	var id string

	if _, err := insert.ScanVal(&id); err != nil {
		return nil, err
	}

	// Add the user as Admin for the newly created event

	organizerinput := model.CreateEventOrganizerInput{
		EventID: id,
		UserID:  user.UserID,
		Role:    model.RoleAdmin,
	}
	_, err = CreateEventOrganizer(ctx, organizerinput)
	if err != nil {
		return nil, err
	}

	// Convert UserModel to User for the response
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
	var start, end time.Time
	var err error

	start, err = time.Parse(time.RFC3339, input.StartDate)
	if err != nil {
		return nil, err
	}

	end, err = time.Parse(time.RFC3339, input.EndDate)
	if err != nil {
		return nil, err
	}

	// Prepare the update statement
	update := db.GetDB().Update("events").
		Set(map[string]interface{}{
			"event_name":  input.EventName,
			"start_date":  start,
			"end_date":    end,
			"location":    input.Location,
			"description": input.Description,
		}).
		Where(goqu.C("event_id").Eq(input.ID)).
		Returning(goqu.C("event_id")).
		Executor()

	var id string

	if _, err := update.ScanVal(&id); err != nil {
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
