package repository

import (
	"context"
	"testing"

	"github.com/leonardocbsr/golang-grpc-template/pkg/users/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open test database: %v", err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}

	return db
}

func TestUserRepository_Create(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)
	ctx := context.Background()

	user := &models.User{
		Username: "testuser",
		Password: "hashedpassword",
		Email:    "test@example.com",
		Name:     "Test User",
	}

	err := repo.Create(ctx, user)
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	if user.ID == 0 {
		t.Error("expected user ID to be set after creation")
	}
}

func TestUserRepository_GetByUsername(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)
	ctx := context.Background()

	// Create user first
	user := &models.User{
		Username: "testuser",
		Password: "hashedpassword",
		Email:    "test@example.com",
		Name:     "Test User",
	}
	if err := repo.Create(ctx, user); err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	// Get by username
	found, err := repo.GetByUsername(ctx, "testuser")
	if err != nil {
		t.Fatalf("GetByUsername failed: %v", err)
	}

	if found.Username != "testuser" {
		t.Errorf("expected username 'testuser', got %q", found.Username)
	}
}

func TestUserRepository_GetByUsername_NotFound(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)
	ctx := context.Background()

	_, err := repo.GetByUsername(ctx, "nonexistent")
	if err == nil {
		t.Fatal("expected error for nonexistent user")
	}

	st, ok := status.FromError(err)
	if !ok {
		t.Fatalf("expected gRPC status error, got %T", err)
	}

	if st.Code() != codes.NotFound {
		t.Errorf("expected NotFound code, got %v", st.Code())
	}
}

func TestUserRepository_GetByEmail(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)
	ctx := context.Background()

	// Create user first
	user := &models.User{
		Username: "testuser",
		Password: "hashedpassword",
		Email:    "test@example.com",
		Name:     "Test User",
	}
	if err := repo.Create(ctx, user); err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	// Get by email
	found, err := repo.GetByEmail(ctx, "test@example.com")
	if err != nil {
		t.Fatalf("GetByEmail failed: %v", err)
	}

	if found.Email != "test@example.com" {
		t.Errorf("expected email 'test@example.com', got %q", found.Email)
	}
}

func TestUserRepository_GetByEmail_NotFound(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)
	ctx := context.Background()

	_, err := repo.GetByEmail(ctx, "nonexistent@example.com")
	if err == nil {
		t.Fatal("expected error for nonexistent email")
	}

	st, ok := status.FromError(err)
	if !ok {
		t.Fatalf("expected gRPC status error, got %T", err)
	}

	if st.Code() != codes.NotFound {
		t.Errorf("expected NotFound code, got %v", st.Code())
	}
}

func TestUserRepository_GetByID(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)
	ctx := context.Background()

	// Create user first
	user := &models.User{
		Username: "testuser",
		Password: "hashedpassword",
		Email:    "test@example.com",
		Name:     "Test User",
	}
	if err := repo.Create(ctx, user); err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	// Get by ID
	found, err := repo.GetByID(ctx, "1")
	if err != nil {
		t.Fatalf("GetByID failed: %v", err)
	}

	if found.ID != user.ID {
		t.Errorf("expected ID %d, got %d", user.ID, found.ID)
	}
}

func TestUserRepository_GetByID_NotFound(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)
	ctx := context.Background()

	_, err := repo.GetByID(ctx, "999")
	if err == nil {
		t.Fatal("expected error for nonexistent ID")
	}

	st, ok := status.FromError(err)
	if !ok {
		t.Fatalf("expected gRPC status error, got %T", err)
	}

	if st.Code() != codes.NotFound {
		t.Errorf("expected NotFound code, got %v", st.Code())
	}
}

func TestUserRepository_Update(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)
	ctx := context.Background()

	// Create user first
	user := &models.User{
		Username: "testuser",
		Password: "hashedpassword",
		Email:    "test@example.com",
		Name:     "Test User",
	}
	if err := repo.Create(ctx, user); err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	// Update user
	user.Name = "Updated Name"
	if err := repo.Update(ctx, user); err != nil {
		t.Fatalf("Update failed: %v", err)
	}

	// Verify update
	found, err := repo.GetByUsername(ctx, "testuser")
	if err != nil {
		t.Fatalf("GetByUsername failed: %v", err)
	}

	if found.Name != "Updated Name" {
		t.Errorf("expected name 'Updated Name', got %q", found.Name)
	}
}

func TestUserRepository_Delete(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)
	ctx := context.Background()

	// Create user first
	user := &models.User{
		Username: "testuser",
		Password: "hashedpassword",
		Email:    "test@example.com",
		Name:     "Test User",
	}
	if err := repo.Create(ctx, user); err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	// Delete user
	if err := repo.Delete(ctx, user); err != nil {
		t.Fatalf("Delete failed: %v", err)
	}

	// Verify deletion (soft delete - should not be found)
	_, err := repo.GetByUsername(ctx, "testuser")
	if err == nil {
		t.Error("expected error after deletion, got nil")
	}
}
