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

// CheckPermission checks if the user has the permission to perform the specified action.
func CheckPermission(ctx context.Context, eventID string, table string, action string) (bool, error) {
	// Retrieve the current user's details
	user, err := GetCurrentUser(ctx)
	if err != nil {
		return false, err
	}

	// Use the user's ID to determine their role for the event
	role, err := GetUserRoleForEvent(ctx, user.UserID, eventID)
	if err != nil {
		return false, err
	}

	// Check if the user has permission to perform the action
	canPerform, err := CanPerformAction(ctx, table, action, role)
	if err != nil {
		return false, err
	}

	return canPerform, nil
}

// canAddRole checks if the current user has permission to add the specified role.
func CanAddRole(ctx context.Context, eventID string, newUserRole string) (bool, error) {

	// Retrieve the current user's details
	user, err := GetCurrentUser(ctx)
	if err != nil {
		return false, err
	}

	// Use the user's ID to determine their role for the event
	role, err := GetUserRoleForEvent(ctx, user.UserID, eventID)
	if err != nil {
		return false, err
	}

	// Define the permissions for adding roles
	permissions := map[string][]string{
		"ADMIN":       {"ADMIN", "CONTRIBUTOR", "ATTENDEE"},
		"CONTRIBUTOR": {"ATTENDEE"},
	}

	// Check if the current user's role is in the permissions map
	roles, ok := permissions[role]
	if !ok {
		return false, errors.New("invalid current user role")
	}

	// Check if the new user's role is allowed
	for _, role := range roles {
		if role == newUserRole {
			return true, nil
		}
	}

	return false, nil
}
