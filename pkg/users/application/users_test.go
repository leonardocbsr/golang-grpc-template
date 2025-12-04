package application

import (
	"context"
	"testing"

	mock_users "github.com/leonardocbsr/golang-grpc-template/internal/mocks/users/repository"
	"github.com/leonardocbsr/golang-grpc-template/pkg/users/models"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreateUser_HashesPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_users.NewMockUserRepository(ctrl)
	svc := NewUserService(mockRepo)

	ctx := context.Background()
	password := "mysecretpassword"

	mockRepo.EXPECT().
		Create(ctx, gomock.Any()).
		DoAndReturn(func(_ context.Context, user *models.User) error {
			// Verify password is hashed
			if user.Password == password {
				t.Error("password should be hashed, not stored in plaintext")
			}
			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
				t.Errorf("password hash verification failed: %v", err)
			}
			return nil
		})

	user, err := svc.CreateUser(ctx, "testuser", password, "test@example.com", "Test User")
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}

	if user.Username != "testuser" {
		t.Errorf("expected username 'testuser', got %q", user.Username)
	}
}

func TestCreateUser_ReturnsUserData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_users.NewMockUserRepository(ctrl)
	svc := NewUserService(mockRepo)

	ctx := context.Background()

	mockRepo.EXPECT().
		Create(ctx, gomock.Any()).
		Return(nil)

	user, err := svc.CreateUser(ctx, "testuser", "password123", "test@example.com", "Test User")
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}

	if user.Username != "testuser" {
		t.Errorf("expected username 'testuser', got %q", user.Username)
	}
	if user.Email != "test@example.com" {
		t.Errorf("expected email 'test@example.com', got %q", user.Email)
	}
	if user.Name != "Test User" {
		t.Errorf("expected name 'Test User', got %q", user.Name)
	}
}

func TestCreateUser_RepositoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_users.NewMockUserRepository(ctrl)
	svc := NewUserService(mockRepo)

	ctx := context.Background()
	expectedErr := status.Error(codes.AlreadyExists, "user already exists")

	mockRepo.EXPECT().
		Create(ctx, gomock.Any()).
		Return(expectedErr)

	_, err := svc.CreateUser(ctx, "testuser", "password123", "test@example.com", "Test User")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	st, ok := status.FromError(err)
	if !ok {
		t.Fatalf("expected gRPC status error, got %T", err)
	}
	if st.Code() != codes.AlreadyExists {
		t.Errorf("expected AlreadyExists code, got %v", st.Code())
	}
}

func TestVerifyPassword_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_users.NewMockUserRepository(ctrl)
	svc := NewUserService(mockRepo)

	ctx := context.Background()
	password := "correctpassword"

	// Pre-hash the password as it would be stored
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	storedUser := &models.User{
		Username: "testuser",
		Password: string(hashedPassword),
		Email:    "test@example.com",
		Name:     "Test User",
	}

	mockRepo.EXPECT().
		GetByUsername(ctx, "testuser").
		Return(storedUser, nil)

	user, err := svc.VerifyPassword(ctx, "testuser", password)
	if err != nil {
		t.Fatalf("VerifyPassword failed: %v", err)
	}

	if user.Username != "testuser" {
		t.Errorf("expected username 'testuser', got %q", user.Username)
	}
}

func TestVerifyPassword_WrongPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_users.NewMockUserRepository(ctrl)
	svc := NewUserService(mockRepo)

	ctx := context.Background()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("correctpassword"), bcrypt.DefaultCost)

	storedUser := &models.User{
		Username: "testuser",
		Password: string(hashedPassword),
	}

	mockRepo.EXPECT().
		GetByUsername(ctx, "testuser").
		Return(storedUser, nil)

	_, err := svc.VerifyPassword(ctx, "testuser", "wrongpassword")
	if err == nil {
		t.Fatal("expected error for wrong password, got nil")
	}

	st, ok := status.FromError(err)
	if !ok {
		t.Fatalf("expected gRPC status error, got %T", err)
	}
	if st.Code() != codes.Unauthenticated {
		t.Errorf("expected Unauthenticated code, got %v", st.Code())
	}
}

func TestVerifyPassword_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_users.NewMockUserRepository(ctrl)
	svc := NewUserService(mockRepo)

	ctx := context.Background()

	mockRepo.EXPECT().
		GetByUsername(ctx, "nonexistent").
		Return(nil, status.Error(codes.NotFound, "user not found"))

	_, err := svc.VerifyPassword(ctx, "nonexistent", "password")
	if err == nil {
		t.Fatal("expected error for nonexistent user, got nil")
	}

	st, ok := status.FromError(err)
	if !ok {
		t.Fatalf("expected gRPC status error, got %T", err)
	}
	if st.Code() != codes.NotFound {
		t.Errorf("expected NotFound code, got %v", st.Code())
	}
}

func TestGetUserByUsername(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_users.NewMockUserRepository(ctrl)
	svc := NewUserService(mockRepo)

	ctx := context.Background()

	expectedUser := &models.User{
		Username: "testuser",
		Email:    "test@example.com",
		Name:     "Test User",
	}

	mockRepo.EXPECT().
		GetByUsername(ctx, "testuser").
		Return(expectedUser, nil)

	user, err := svc.GetUserByUsername(ctx, "testuser")
	if err != nil {
		t.Fatalf("GetUserByUsername failed: %v", err)
	}

	if user.Username != "testuser" {
		t.Errorf("expected username 'testuser', got %q", user.Username)
	}
}

func TestGetUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_users.NewMockUserRepository(ctrl)
	svc := NewUserService(mockRepo)

	ctx := context.Background()

	expectedUser := &models.User{
		Username: "testuser",
		Email:    "test@example.com",
		Name:     "Test User",
	}

	mockRepo.EXPECT().
		GetByEmail(ctx, "test@example.com").
		Return(expectedUser, nil)

	user, err := svc.GetUserByEmail(ctx, "test@example.com")
	if err != nil {
		t.Fatalf("GetUserByEmail failed: %v", err)
	}

	if user.Email != "test@example.com" {
		t.Errorf("expected email 'test@example.com', got %q", user.Email)
	}
}

func TestGetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_users.NewMockUserRepository(ctrl)
	svc := NewUserService(mockRepo)

	ctx := context.Background()

	expectedUser := &models.User{
		Username: "testuser",
		Email:    "test@example.com",
	}

	mockRepo.EXPECT().
		GetByID(ctx, "123").
		Return(expectedUser, nil)

	user, err := svc.GetUserByID(ctx, "123")
	if err != nil {
		t.Fatalf("GetUserByID failed: %v", err)
	}

	if user.Username != "testuser" {
		t.Errorf("expected username 'testuser', got %q", user.Username)
	}
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_users.NewMockUserRepository(ctrl)
	svc := NewUserService(mockRepo)

	ctx := context.Background()

	user := &models.User{
		Username: "testuser",
		Email:    "updated@example.com",
	}

	mockRepo.EXPECT().
		Update(ctx, user).
		Return(nil)

	err := svc.UpdateUser(ctx, user)
	if err != nil {
		t.Fatalf("UpdateUser failed: %v", err)
	}
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_users.NewMockUserRepository(ctrl)
	svc := NewUserService(mockRepo)

	ctx := context.Background()

	user := &models.User{
		Username: "testuser",
	}

	mockRepo.EXPECT().
		Delete(ctx, user).
		Return(nil)

	err := svc.DeleteUser(ctx, user)
	if err != nil {
		t.Fatalf("DeleteUser failed: %v", err)
	}
}
