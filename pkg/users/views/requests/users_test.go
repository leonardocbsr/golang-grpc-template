package requests

import (
	"strings"
	"testing"

	users "github.com/leonardocbsr/golang-grpc-template/pkg/users/proto"
)

func TestValidateCreateUserRequest(t *testing.T) {
	tests := []struct {
		name          string
		req           *users.CreateUserRequest
		wantError     bool
		errorContains string
	}{
		{
			name: "valid request",
			req: &users.CreateUserRequest{
				Username: "testuser",
				Password: "password123",
				Email:    "test@example.com",
				Name:     "Test User",
			},
			wantError: false,
		},
		{
			name: "missing username",
			req: &users.CreateUserRequest{
				Username: "",
				Password: "password123",
				Email:    "test@example.com",
				Name:     "Test User",
			},
			wantError:     true,
			errorContains: "Username",
		},
		{
			name: "username too short",
			req: &users.CreateUserRequest{
				Username: "ab",
				Password: "password123",
				Email:    "test@example.com",
				Name:     "Test User",
			},
			wantError:     true,
			errorContains: "Username",
		},
		{
			name: "password too short",
			req: &users.CreateUserRequest{
				Username: "testuser",
				Password: "short",
				Email:    "test@example.com",
				Name:     "Test User",
			},
			wantError:     true,
			errorContains: "Password",
		},
		{
			name: "password exactly 8 chars (minimum)",
			req: &users.CreateUserRequest{
				Username: "testuser",
				Password: "exactly8",
				Email:    "test@example.com",
				Name:     "Test User",
			},
			wantError: false,
		},
		{
			name: "invalid email",
			req: &users.CreateUserRequest{
				Username: "testuser",
				Password: "password123",
				Email:    "invalid-email",
				Name:     "Test User",
			},
			wantError:     true,
			errorContains: "Email",
		},
		{
			name: "missing name",
			req: &users.CreateUserRequest{
				Username: "testuser",
				Password: "password123",
				Email:    "test@example.com",
				Name:     "",
			},
			wantError:     true,
			errorContains: "Name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCreateUserRequest(tt.req)

			if tt.wantError {
				if err == nil {
					t.Error("expected error, got nil")
					return
				}
				if tt.errorContains != "" && !strings.Contains(err.Error(), tt.errorContains) {
					t.Errorf("expected error to contain %q, got %q", tt.errorContains, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
			}
		})
	}
}

func TestValidateGetUserByUsernameRequest(t *testing.T) {
	tests := []struct {
		name      string
		req       *users.GetUserByUsernameRequest
		wantError bool
	}{
		{
			name:      "valid request",
			req:       &users.GetUserByUsernameRequest{Username: "testuser"},
			wantError: false,
		},
		{
			name:      "missing username",
			req:       &users.GetUserByUsernameRequest{Username: ""},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateGetUserByUsernameRequest(tt.req)

			if tt.wantError && err == nil {
				t.Error("expected error, got nil")
			}
			if !tt.wantError && err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	}
}

func TestValidateGetUserByEmailRequest(t *testing.T) {
	tests := []struct {
		name      string
		req       *users.GetUserByEmailRequest
		wantError bool
	}{
		{
			name:      "valid request",
			req:       &users.GetUserByEmailRequest{Email: "test@example.com"},
			wantError: false,
		},
		{
			name:      "missing email",
			req:       &users.GetUserByEmailRequest{Email: ""},
			wantError: true,
		},
		{
			name:      "invalid email format",
			req:       &users.GetUserByEmailRequest{Email: "not-an-email"},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateGetUserByEmailRequest(tt.req)

			if tt.wantError && err == nil {
				t.Error("expected error, got nil")
			}
			if !tt.wantError && err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	}
}

func TestValidateGetUserByIDRequest(t *testing.T) {
	tests := []struct {
		name      string
		req       *users.GetUserRequest
		wantError bool
	}{
		{
			name:      "valid request",
			req:       &users.GetUserRequest{Id: "123"},
			wantError: false,
		},
		{
			name:      "missing id",
			req:       &users.GetUserRequest{Id: ""},
			wantError: true,
		},
		{
			name:      "non-numeric id",
			req:       &users.GetUserRequest{Id: "abc"},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateGetUserByIDRequest(tt.req)

			if tt.wantError && err == nil {
				t.Error("expected error, got nil")
			}
			if !tt.wantError && err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	}
}

func TestValidateUpdateUserRequest(t *testing.T) {
	tests := []struct {
		name      string
		req       *users.UpdateUserRequest
		wantError bool
	}{
		{
			name: "valid request",
			req: &users.UpdateUserRequest{
				Id:       "123",
				Username: "testuser",
				Email:    "test@example.com",
				Name:     "Test User",
			},
			wantError: false,
		},
		{
			name: "missing id",
			req: &users.UpdateUserRequest{
				Id:       "",
				Username: "testuser",
				Email:    "test@example.com",
				Name:     "Test User",
			},
			wantError: true,
		},
		{
			name: "username too short",
			req: &users.UpdateUserRequest{
				Id:       "123",
				Username: "ab",
				Email:    "test@example.com",
				Name:     "Test User",
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUpdateUserRequest(tt.req)

			if tt.wantError && err == nil {
				t.Error("expected error, got nil")
			}
			if !tt.wantError && err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	}
}

func TestValidateDeleteUserRequest(t *testing.T) {
	tests := []struct {
		name      string
		req       *users.DeleteUserRequest
		wantError bool
	}{
		{
			name:      "valid request",
			req:       &users.DeleteUserRequest{Id: "123"},
			wantError: false,
		},
		{
			name:      "missing id",
			req:       &users.DeleteUserRequest{Id: ""},
			wantError: true,
		},
		{
			name:      "non-numeric id",
			req:       &users.DeleteUserRequest{Id: "abc"},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateDeleteUserRequest(tt.req)

			if tt.wantError && err == nil {
				t.Error("expected error, got nil")
			}
			if !tt.wantError && err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	}
}
