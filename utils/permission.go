package utils

import (
	"context"
	"errors"
)

var permissions = map[string]map[string]map[string]bool{
	"event_organizers": {
		"ADMIN": {
			"add":    true,
			"update": true,
			"delete": true,
			"view":   true,
		},
		"CONTRIBUTOR": {
			"add":  true,
			"view": true,
		},
		"ATTENDEE": {
			"view": true,
		},
	},
	"events": {
		"ADMIN": {
			"update": true,
			"delete": true,
			"view":   true,
		},
		"CONTRIBUTOR": {
			"add":  true,
			"view": true,
		},
		"ATTENDEE": {
			"view": true,
		},
	},
	"event_sessions": {
		"ADMIN": {
			"add":    true,
			"update": true,
			"delete": true,
			"view":   true,
		},
		"CONTRIBUTOR": {
			"add":    true,
			"update": true,
			"view":   true,
		},
		"ATTENDEE": {
			"view": true,
		}},
	"expenses": {
		"ADMIN": {
			"add":    true,
			"update": true,
			"delete": true,
			"view":   true,
		},
		"CONTRIBUTOR": {
			"view": true,
		},
		"ATTENDEE": {
			"view": false,
		}},
}

func CanPerformAction(ctx context.Context, table string, action string, role string) (bool, error) {

	// Directly access the permissions for the specific table
	permissionsForTable, ok := permissions[table]
	if !ok {
		return false, errors.New("invalid table")
	}
	// Access the permissions for the specific role
	permissionsForRole, ok := permissionsForTable[role]
	if !ok {
		return false, errors.New("invalid role")
	}

	// Check if the action is allowed for the user's role on the specified table
	if permissionsForRole[action] {
		return true, nil
	}

	return false, nil
}
