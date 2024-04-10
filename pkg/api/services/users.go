package service

import (
	"context"
	"kraneapi/pkg/api/dbmodel"
	"kraneapi/pkg/db"

	"github.com/doug-martin/goqu/v9"
)

func AddUser(ctx context.Context, userModel dbmodel.User) (string, error) {
	userInsert := db.GetDB().Insert("users").
		Cols("name", "email", "phone_number").
		Vals(goqu.Vals{userModel.Name, userModel.Email, userModel.PhoneNumber}).
		Returning(goqu.C("user_id")).Executor()

	var id string
	if _, err := userInsert.ScanVal(&id); err != nil {
		return "", err
	}

	return id, nil
}

// GetUsers retrieves all users from the database.
func GetUsers(ctx context.Context) ([]*dbmodel.User, error) {
	var userModels []*dbmodel.User
	err := db.GetDB().From("users").Select("*").ScanStructs(&userModels)
	if err != nil {
		return nil, err
	}
	return userModels, nil
}

// GetUserByID retrieves a user by their ID.
func GetUserByID(ctx context.Context, id string) (*dbmodel.User, error) {
	var userModel dbmodel.User
	_, err := db.GetDB().From("users").Where(goqu.C("user_id").Eq(id)).Select("*").ScanStruct(&userModel)
	if err != nil {
		return nil, err
	}
	return &userModel, nil
}
