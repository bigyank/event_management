package service

import (
	"context"
	"errors"
	"kraneapi/graph/model"
	"kraneapi/pkg/db"

	"github.com/doug-martin/goqu/v9"
)

// CreateEventOrganizer creates a new event organizer in the database.
func CreateEventOrganizer(ctx context.Context, input model.CreateEventOrganizerInput) (string, error) {
	eventOrganizerInsert := db.GetDB().Insert("event_organizers").
		Cols("event_id", "user_id", "role").
		Vals(goqu.Vals{input.EventID, input.UserID, input.Role}).
		Returning(goqu.C("event_organizer_id"))

	var id string
	if _, err := eventOrganizerInsert.Executor().ScanVal(&id); err != nil {
		return "", err
	}

	return id, nil
}

// UpdateEventOrganizer updates an existing event organizer in the database.
func UpdateEventOrganizer(ctx context.Context, input model.UpdateEventOrganizerInput) (string, error) {
	eventOrganizerUpdate := db.GetDB().Update("event_organizers").
		Set(goqu.Record{"role": input.Role}).
		Where(goqu.C("event_id").Eq(input.EventID), goqu.C("user_id").Eq(input.UserID)).
		Returning(goqu.C("event_organizer_id"))

	var id string
	if applied, err := eventOrganizerUpdate.Executor().ScanVal(&id); err != nil || !applied {
		if !applied {
			return "", errors.New("no event organizer found with the provided event ID and user ID")
		}
		return "", err
	}

	return id, nil
}
