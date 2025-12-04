// Package responses provides response mapping for user service.
package responses

import (
	"fmt"

	"github.com/leonardocbsr/golang-grpc-template/pkg/users/models"
	users "github.com/leonardocbsr/golang-grpc-template/pkg/users/proto"
)

// ToGetResponse converts a User model to a GetUserResponse.
func ToGetResponse(user *models.User) *users.GetUserResponse {
	return &users.GetUserResponse{
		Id:        fmt.Sprint(user.ID),
		Username:  user.Username,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}
}

// ToCreateResponse converts a User model to a CreateUserResponse.
func ToCreateResponse(user *models.User) *users.CreateUserResponse {
	return &users.CreateUserResponse{
		Id:        fmt.Sprint(user.ID),
		Username:  user.Username,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}
}

// ToUpdateResponse converts a User model to an UpdateUserResponse.
func ToUpdateResponse(user *models.User) *users.UpdateUserResponse {
	return &users.UpdateUserResponse{
		Id:        fmt.Sprint(user.ID),
		Username:  user.Username,
		Email:     user.Email,
		Name:      user.Name,
		UpdatedAt: user.UpdatedAt.String(),
	}
}

// ToDeleteResponse converts a User model to a DeleteUserResponse.
func ToDeleteResponse(user *models.User) *users.DeleteUserResponse {
	return &users.DeleteUserResponse{
		Id:        fmt.Sprint(user.ID),
		DeletedAt: user.DeletedAt.Time.String(),
	}
}
