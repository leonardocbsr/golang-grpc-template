package errors

import (
	"errors"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func TestWrapDBError(t *testing.T) {
	tests := []struct {
		name         string
		err          error
		expectedCode codes.Code
		expectedNil  bool
	}{
		{
			name:        "nil error returns nil",
			err:         nil,
			expectedNil: true,
		},
		{
			name:         "record not found returns NotFound",
			err:          gorm.ErrRecordNotFound,
			expectedCode: codes.NotFound,
		},
		{
			name:         "duplicated key returns AlreadyExists",
			err:          gorm.ErrDuplicatedKey,
			expectedCode: codes.AlreadyExists,
		},
		{
			name:         "other errors return Internal",
			err:          errors.New("some db error"),
			expectedCode: codes.Internal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WrapDBError(tt.err)

			if tt.expectedNil {
				if result != nil {
					t.Errorf("expected nil, got %v", result)
				}
				return
			}

			st, ok := status.FromError(result)
			if !ok {
				t.Fatalf("expected gRPC status error, got %T", result)
			}

			if st.Code() != tt.expectedCode {
				t.Errorf("expected code %v, got %v", tt.expectedCode, st.Code())
			}
		})
	}
}

func TestNotFound(t *testing.T) {
	err := NotFound("user")

	st, ok := status.FromError(err)
	if !ok {
		t.Fatalf("expected gRPC status error, got %T", err)
	}

	if st.Code() != codes.NotFound {
		t.Errorf("expected NotFound code, got %v", st.Code())
	}

	if st.Message() != "user not found" {
		t.Errorf("expected 'user not found', got %q", st.Message())
	}
}

func TestAlreadyExists(t *testing.T) {
	err := AlreadyExists("email")

	st, ok := status.FromError(err)
	if !ok {
		t.Fatalf("expected gRPC status error, got %T", err)
	}

	if st.Code() != codes.AlreadyExists {
		t.Errorf("expected AlreadyExists code, got %v", st.Code())
	}

	if st.Message() != "email already exists" {
		t.Errorf("expected 'email already exists', got %q", st.Message())
	}
}

func TestInvalidArgument(t *testing.T) {
	err := InvalidArgument("invalid email format")

	st, ok := status.FromError(err)
	if !ok {
		t.Fatalf("expected gRPC status error, got %T", err)
	}

	if st.Code() != codes.InvalidArgument {
		t.Errorf("expected InvalidArgument code, got %v", st.Code())
	}

	if st.Message() != "invalid email format" {
		t.Errorf("expected 'invalid email format', got %q", st.Message())
	}
}

func TestInternal(t *testing.T) {
	err := Internal("database connection failed")

	st, ok := status.FromError(err)
	if !ok {
		t.Fatalf("expected gRPC status error, got %T", err)
	}

	if st.Code() != codes.Internal {
		t.Errorf("expected Internal code, got %v", st.Code())
	}

	if st.Message() != "database connection failed" {
		t.Errorf("expected 'database connection failed', got %q", st.Message())
	}
}

func TestPredefinedErrors(t *testing.T) {
	tests := []struct {
		name         string
		err          error
		expectedCode codes.Code
	}{
		{"ErrNotFound", ErrNotFound, codes.NotFound},
		{"ErrAlreadyExists", ErrAlreadyExists, codes.AlreadyExists},
		{"ErrInvalidArgument", ErrInvalidArgument, codes.InvalidArgument},
		{"ErrInternal", ErrInternal, codes.Internal},
		{"ErrUnauthenticated", ErrUnauthenticated, codes.Unauthenticated},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st, ok := status.FromError(tt.err)
			if !ok {
				t.Fatalf("expected gRPC status error, got %T", tt.err)
			}

			if st.Code() != tt.expectedCode {
				t.Errorf("expected code %v, got %v", tt.expectedCode, st.Code())
			}
		})
	}
}
