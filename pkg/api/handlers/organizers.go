package handler

import (
	"context"
	"errors"
	"kraneapi/graph/model"
	"kraneapi/pkg/db"

	"github.com/doug-martin/goqu/v9"
)

func CreateEventOrganizer(ctx context.Context, input model.CreateEventOrganizerInput) (*model.EventOrganizer, error) {

	eventOrganizerInsert := db.GetDB().Insert("event_organizers").
		Cols("event_id", "user_id", "role").
		Vals(goqu.Vals{input.EventID, input.UserID, input.Role}).
		Returning(goqu.C("event_organizer_id"))

	var id string

	if _, err := eventOrganizerInsert.Executor().ScanVal(&id); err != nil {
		return nil, err
	}

	// Convert the input to EventOrganizer for the response
	eventOrganizer := model.EventOrganizer{
		EventOrganizerID: id,
		EventID:          input.EventID,
		UserID:           input.UserID,
		Role:             input.Role,
	}

	return &eventOrganizer, nil
}

func UpdateEventOrganizer(ctx context.Context, input model.UpdateEventOrganizerInput) (*model.EventOrganizer, error) {
	eventOrganizerUpdate := db.GetDB().Update("event_organizers").
		Set(goqu.Record{"role": input.Role}).
		Where(goqu.C("event_id").Eq(input.EventID), goqu.C("user_id").Eq(input.UserID)).
		Returning(goqu.C("event_organizer_id"))

	var id string

	if applied, err := eventOrganizerUpdate.Executor().ScanVal(&id); err != nil || !applied {
		if !applied {
			return nil, errors.New("no event organizer found with the provided event ID and user ID")
		}
		return nil, err
	}

	// Convert the input to EventOrganizer for the response
	eventOrganizer := model.EventOrganizer{
		EventOrganizerID: id,
		EventID:          input.EventID,
		UserID:           input.UserID,
		Role:             input.Role,
	}

	return &eventOrganizer, nil
}
