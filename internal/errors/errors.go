package errors

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

var (
	// ErrNotFound is returned when a resource is not found
	ErrNotFound = status.Error(codes.NotFound, "resource not found")
	// ErrAlreadyExists is returned when a resource already exists
	ErrAlreadyExists = status.Error(codes.AlreadyExists, "resource already exists")
	// ErrInvalidArgument is returned when an argument is invalid
	ErrInvalidArgument = status.Error(codes.InvalidArgument, "invalid argument")
	// ErrInternal is returned when an internal error occurs
	ErrInternal = status.Error(codes.Internal, "internal error")
	// ErrUnauthenticated is returned when authentication fails
	ErrUnauthenticated = status.Error(codes.Unauthenticated, "unauthenticated")
	// ErrInvalidClientName is returned when a client name is invalid
	ErrInvalidClientName = errors.New("invalid client name")
)

// WrapDBError converts gorm errors to gRPC status errors
func WrapDBError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrNotFound
	}

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return ErrAlreadyExists
	}

	return status.Errorf(codes.Internal, "database error: %v", err)
}

// NotFound returns a not found error with a custom message
func NotFound(resource string) error {
	return status.Errorf(codes.NotFound, "%s not found", resource)
}

// AlreadyExists returns an already exists error with a custom message
func AlreadyExists(resource string) error {
	return status.Errorf(codes.AlreadyExists, "%s already exists", resource)
}

// InvalidArgument returns an invalid argument error with a custom message
func InvalidArgument(msg string) error {
	return status.Errorf(codes.InvalidArgument, "%s", msg)
}

// Internal returns an internal error with a custom message
func Internal(msg string) error {
	return status.Errorf(codes.Internal, "%s", msg)
}
