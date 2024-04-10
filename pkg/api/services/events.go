package service

import (
	"context"
	"kraneapi/graph/model"
	"kraneapi/pkg/db"
	"time"

	"github.com/doug-martin/goqu/v9"
)

// CreateEvent creates a new event in the database.
func CreateEvent(ctx context.Context, input model.AddEventInput) (string, error) {
	start, err := time.Parse(time.RFC3339, input.StartDate)
	if err != nil {
		return "", err
	}

	end, err := time.Parse(time.RFC3339, input.EndDate)
	if err != nil {
		return "", err
	}

	insert := db.GetDB().Insert("events").
		Cols("event_name", "start_date", "end_date", "location", "description").
		Vals(goqu.Vals{input.EventName, start, end, input.Location, input.Description}).
		Returning(goqu.C("event_id")).
		Executor()

	var id string
	if _, err := insert.ScanVal(&id); err != nil {
		return "", err
	}

	return id, nil
}

// UpdateEvent updates an existing event in the database.
func UpdateEvent(ctx context.Context, input model.UpdateEventInput) (string, error) {
	start, err := time.Parse(time.RFC3339, input.StartDate)
	if err != nil {
		return "", err
	}

	end, err := time.Parse(time.RFC3339, input.EndDate)
	if err != nil {
		return "", err
	}

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
		return "", err
	}

	return id, nil
}
