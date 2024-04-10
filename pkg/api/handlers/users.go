package handler

import (
	"context"
	"kraneapi/graph/model"
	"kraneapi/pkg/api/dbmodel"
	service "kraneapi/pkg/api/services"
)

// Implement the AddUser method of the mutationResolver interface
func AddUser(ctx context.Context, input model.AddUserInput) (*model.User, error) {
	userModel := dbmodel.User{
		Name:        input.Name,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
	}

	id, err := service.AddUser(ctx, userModel)
	if err != nil {
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
	userModels, err := service.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

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
	userModel, err := service.GetUserByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	user := model.User{
		UserID:      userModel.UserID,
		Name:        userModel.Name,
		Email:       userModel.Email,
		PhoneNumber: userModel.PhoneNumber,
	}

	return &user, nil
}
