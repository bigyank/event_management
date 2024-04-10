// utils/auth_utils.go

package utils

import (
	"context"
	"errors"
	"fmt"
	auth "kraneapi/middleware"
	"kraneapi/pkg/api/dbmodel"
	"kraneapi/pkg/db"

	"github.com/doug-martin/goqu/v9"
)

func GetCurrentUserIdHeader(ctx context.Context) (string, error) {
	userID := auth.ForContext(ctx)
	if userID == "" {
		return "", fmt.Errorf("unauthorized: user ID not provided")
	}
	return userID, nil
}

func GetCurrentUser(ctx context.Context) (*dbmodel.User, error) {
	userID, err := GetCurrentUserIdHeader(ctx)
	if err != nil {
		return nil, err
	}

	var user dbmodel.User
	found, err := db.GetDB().From("users").
		Where(goqu.C("user_id").Eq(userID)).
		ScanStruct(&user)

	if err != nil {
		return nil, err
	}

	if !found {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func GetUserRoleForEvent(ctx context.Context, userID string, eventID string) (string, error) {

	var eventOrganizer dbmodel.EventOrganizer
	found, err := db.GetDB().From("event_organizers").
		Where(goqu.C("user_id").Eq(userID), goqu.C("event_id").Eq(eventID)).
		ScanStruct(&eventOrganizer)

	if err != nil {
		return "", err
	}

	if !found {
		return "", errors.New("user not found")
	}

	return eventOrganizer.Role, nil
}
