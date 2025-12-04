// Package requests provides validation for user service requests.
package requests

import (
	"github.com/go-playground/validator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	users "github.com/leonardocbsr/golang-grpc-template/pkg/users/proto"
)

var validate = validator.New()

// validateStruct validates a struct and returns a gRPC InvalidArgument error on failure.
func validateStruct(v any) error {
	if err := validate.Struct(v); err != nil {
		return status.Errorf(codes.InvalidArgument, "validation failed: %s", err.Error())
	}
	return nil
}

// CreateUserRequest holds validation rules for creating a user.
type CreateUserRequest struct {
	Username string `validate:"required,min=3,max=50"`
	Password string `validate:"required,min=8,max=72"`
	Email    string `validate:"required,email"`
	Name     string `validate:"required,min=1,max=100"`
}

// ValidateCreateUserRequest validates a create user request.
func ValidateCreateUserRequest(req *users.CreateUserRequest) error {
	return validateStruct(CreateUserRequest{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
		Email:    req.GetEmail(),
		Name:     req.GetName(),
	})
}

// GetUserByUsernameRequest holds validation rules for getting a user by username.
type GetUserByUsernameRequest struct {
	Username string `validate:"required"`
}

// ValidateGetUserByUsernameRequest validates a get user by username request.
func ValidateGetUserByUsernameRequest(req *users.GetUserByUsernameRequest) error {
	return validateStruct(GetUserByUsernameRequest{
		Username: req.GetUsername(),
	})
}

// GetUserByEmailRequest holds validation rules for getting a user by email.
type GetUserByEmailRequest struct {
	Email string `validate:"required,email"`
}

// ValidateGetUserByEmailRequest validates a get user by email request.
func ValidateGetUserByEmailRequest(req *users.GetUserByEmailRequest) error {
	return validateStruct(GetUserByEmailRequest{
		Email: req.GetEmail(),
	})
}

// GetUserByIDRequest holds validation rules for getting a user by ID.
type GetUserByIDRequest struct {
	ID string `validate:"required,number"`
}

// ValidateGetUserByIDRequest validates a get user by ID request.
func ValidateGetUserByIDRequest(req *users.GetUserRequest) error {
	return validateStruct(GetUserByIDRequest{
		ID: req.GetId(),
	})
}

// UpdateUserRequest holds validation rules for updating a user.
type UpdateUserRequest struct {
	ID       string `validate:"required,number"`
	Username string `validate:"required,min=3,max=50"`
	Email    string `validate:"required,email"`
	Name     string `validate:"required,min=1,max=100"`
}

// ValidateUpdateUserRequest validates an update user request.
func ValidateUpdateUserRequest(req *users.UpdateUserRequest) error {
	return validateStruct(UpdateUserRequest{
		ID:       req.GetId(),
		Username: req.GetUsername(),
		Email:    req.GetEmail(),
		Name:     req.GetName(),
	})
}

// DeleteUserRequest holds validation rules for deleting a user.
type DeleteUserRequest struct {
	ID string `validate:"required,number"`
}

// ValidateDeleteUserRequest validates a delete user request.
func ValidateDeleteUserRequest(req *users.DeleteUserRequest) error {
	return validateStruct(DeleteUserRequest{
		ID: req.GetId(),
	})
}
