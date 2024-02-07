package requests

import (
	"cbsr.io/golang-grpc-template/modules/users/exceptions"
	"cbsr.io/golang-grpc-template/proto/users"
	"github.com/go-playground/validator"
)

var validate = validator.New()

type CreateUserRequest struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
	Email    string `validate:"required,email"`
	Name     string `validate:"required"`
}

func ValidateCreateUserRequest(req *users.CreateUserRequest) error {
	validations := CreateUserRequest{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
		Email:    req.GetEmail(),
		Name:     req.GetName(),
	}

	if err := validate.Struct(validations); err != nil {
		return exceptions.NewInvalidArgument(err.Error())
	}

	return nil
}

type GetUserByUsernameRequest struct {
	Username string `validate:"required"`
}

func ValidateGetUserByUsernameRequest(req *users.GetUserByUsernameRequest) error {
	validations := GetUserByUsernameRequest{
		Username: req.GetUsername(),
	}

	if err := validate.Struct(validations); err != nil {
		return exceptions.NewInvalidArgument(err.Error())
	}

	return nil
}

type GetUserByEmailRequest struct {
	Email string `validate:"required,email"`
}

func ValidateGetUserByEmailRequest(req *users.GetUserByEmailRequest) error {
	validations := GetUserByEmailRequest{
		Email: req.GetEmail(),
	}

	if err := validate.Struct(validations); err != nil {
		return exceptions.NewInvalidArgument(err.Error())
	}

	return nil
}

type GetUserByIDRequest struct {
	ID string `validate:"required,number"`
}

func ValidateGetUserByIDRequest(req *users.GetUserRequest) error {
	validations := GetUserByIDRequest{
		ID: req.GetId(),
	}

	if err := validate.Struct(validations); err != nil {
		return exceptions.NewInvalidArgument(err.Error())
	}

	return nil
}

type UpdateUserRequest struct {
	ID       string `validate:"required,number"`
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Name     string `validate:"required"`
}

func ValidateUpdateUserRequest(req *users.UpdateUserRequest) error {
	validations := UpdateUserRequest{
		ID:       req.GetId(),
		Username: req.GetUsername(),
		Email:    req.GetEmail(),
		Name:     req.GetName(),
	}

	if err := validate.Struct(validations); err != nil {
		return exceptions.NewInvalidArgument(err.Error())
	}

	return nil
}

type DeleteUserRequest struct {
	ID string `validate:"required,number"`
}

func ValidateDeleteUserRequest(req *users.DeleteUserRequest) error {
	validations := DeleteUserRequest{
		ID: req.GetId(),
	}

	if err := validate.Struct(validations); err != nil {
		return exceptions.NewInvalidArgument(err.Error())
	}

	return nil
}
