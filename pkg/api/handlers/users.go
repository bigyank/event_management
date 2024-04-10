package handler

import (
	"context"
	"kraneapi/graph/model"
	"kraneapi/pkg/api/dbmodel"
	"kraneapi/pkg/db"

	"github.com/doug-martin/goqu/v9"
)

type UserResolver struct{}

func AddUser(ctx context.Context, input model.AddUserInput) (*model.User, error) {
	userModel := dbmodel.User{
		Name:        input.Name,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
	}

	// Insert the user into the database using goqu and return the inserted row's data
	userinsert := db.GetDB().Insert("users").
		Cols("name", "email", "phone_number").
		Vals(goqu.Vals{userModel.Name, userModel.Email, userModel.PhoneNumber}).
		Returning(goqu.C("user_id")).Executor()

	var id string

	if _, err := userinsert.ScanVal(&id); err != nil {
		return nil, err
	}

	// Convert UserModel to User for the response
	user := model.User{
		UserID:      id,
		Name:        userModel.Name,
		Email:       userModel.Email,
		PhoneNumber: userModel.PhoneNumber,
	}

	return &user, nil
}

func GetUsers(ctx context.Context) ([]*model.User, error) {
	var userModels []*dbmodel.User
	err := db.GetDB().From("users").Select("*").ScanStructs(&userModels)
	if err != nil {
		return nil, err
	}

	// Convert UserModel to User for the response
	var users []*model.User
	for _, userModel := range userModels {
		user := model.User{
			UserID:      userModel.UserID,
			Name:        userModel.Name,
			Email:       userModel.Email,
			PhoneNumber: userModel.PhoneNumber,
		}
		users = append(users, &user)
	}

	return users, nil
}

func GetUserByID(ctx context.Context, input model.UserIDInput) (*model.User, error) {
	var userModel dbmodel.User
	_, err := db.GetDB().From("users").Where(goqu.C("user_id").Eq(input.ID)).Select("*").ScanStruct(&userModel)

	if err != nil {
		return nil, err
	}

	// Convert UserModel to User for the response
	user := model.User{
		UserID:      userModel.UserID,
		Name:        userModel.Name,
		Email:       userModel.Email,
		PhoneNumber: userModel.PhoneNumber,
	}

	return &user, nil
}
