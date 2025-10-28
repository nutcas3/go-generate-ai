package service

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/example/speedrun-rest-api/db"
	"github.com/jackc/pgx/v5/pgtype"
)

// Helper function to convert time.Time to pgtype.Timestamp
func timeToTimestamp(t time.Time) pgtype.Timestamp {
	return pgtype.Timestamp{
		Time:  t,
		Valid: true,
	}
}

// MockQueries is a mock implementation of db.Queries for testing
type MockQueries struct {
	GetUserByIDFunc    func(ctx context.Context, id int32) (db.User, error)
	GetUserByEmailFunc func(ctx context.Context, email string) (db.User, error)
	ListUsersFunc      func(ctx context.Context, params db.ListUsersParams) ([]db.User, error)
	CountUsersFunc     func(ctx context.Context) (int64, error)
	CreateUserFunc     func(ctx context.Context, params db.CreateUserParams) (db.User, error)
	UpdateUserFunc     func(ctx context.Context, params db.UpdateUserParams) (db.User, error)
	DeleteUserFunc     func(ctx context.Context, id int32) error
}

func (m *MockQueries) GetUserByID(ctx context.Context, id int32) (db.User, error) {
	if m.GetUserByIDFunc != nil {
		return m.GetUserByIDFunc(ctx, id)
	}
	return db.User{}, sql.ErrNoRows
}

func (m *MockQueries) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	if m.GetUserByEmailFunc != nil {
		return m.GetUserByEmailFunc(ctx, email)
	}
	return db.User{}, sql.ErrNoRows
}

func (m *MockQueries) ListUsers(ctx context.Context, params db.ListUsersParams) ([]db.User, error) {
	if m.ListUsersFunc != nil {
		return m.ListUsersFunc(ctx, params)
	}
	return []db.User{}, nil
}

func (m *MockQueries) CountUsers(ctx context.Context) (int64, error) {
	if m.CountUsersFunc != nil {
		return m.CountUsersFunc(ctx)
	}
	return 0, nil
}

func (m *MockQueries) CreateUser(ctx context.Context, params db.CreateUserParams) (db.User, error) {
	if m.CreateUserFunc != nil {
		return m.CreateUserFunc(ctx, params)
	}
	return db.User{}, nil
}

func (m *MockQueries) UpdateUser(ctx context.Context, params db.UpdateUserParams) (db.User, error) {
	if m.UpdateUserFunc != nil {
		return m.UpdateUserFunc(ctx, params)
	}
	return db.User{}, nil
}

func (m *MockQueries) DeleteUser(ctx context.Context, id int32) error {
	if m.DeleteUserFunc != nil {
		return m.DeleteUserFunc(ctx, id)
	}
	return nil
}

func TestGetUserByID_Success(t *testing.T) {
	now := time.Now()
	mockQueries := &MockQueries{
		GetUserByIDFunc: func(ctx context.Context, id int32) (db.User, error) {
			return db.User{
				ID:        1,
				Name:      "John Doe",
				Email:     "john@example.com",
				CreatedAt: timeToTimestamp(now),
				UpdatedAt: timeToTimestamp(now),
			}, nil
		},
	}

	service := NewUserService(mockQueries)
	user, err := service.GetUserByID(context.Background(), 1)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if user.ID != 1 {
		t.Errorf("expected user ID 1, got %d", user.ID)
	}
	if user.Name != "John Doe" {
		t.Errorf("expected name 'John Doe', got %s", user.Name)
	}
}

func TestGetUserByID_NotFound(t *testing.T) {
	mockQueries := &MockQueries{
		GetUserByIDFunc: func(ctx context.Context, id int32) (db.User, error) {
			return db.User{}, sql.ErrNoRows
		},
	}

	service := NewUserService(mockQueries)
	_, err := service.GetUserByID(context.Background(), 999)

	if !errors.Is(err, ErrUserNotFound) {
		t.Errorf("expected ErrUserNotFound, got %v", err)
	}
}

func TestCreateUser_Success(t *testing.T) {
	now := time.Now()
	mockQueries := &MockQueries{
		GetUserByEmailFunc: func(ctx context.Context, email string) (db.User, error) {
			return db.User{}, sql.ErrNoRows // No existing user
		},
		CreateUserFunc: func(ctx context.Context, params db.CreateUserParams) (db.User, error) {
			return db.User{
				ID:        1,
				Name:      params.Name,
				Email:     params.Email,
				CreatedAt: timeToTimestamp(now),
				UpdatedAt: timeToTimestamp(now),
			}, nil
		},
	}

	service := NewUserService(mockQueries)
	user, err := service.CreateUser(context.Background(), "Jane Doe", "jane@example.com")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if user.Name != "Jane Doe" {
		t.Errorf("expected name 'Jane Doe', got %s", user.Name)
	}
	if user.Email != "jane@example.com" {
		t.Errorf("expected email 'jane@example.com', got %s", user.Email)
	}
}

func TestCreateUser_DuplicateEmail(t *testing.T) {
	mockQueries := &MockQueries{
		GetUserByEmailFunc: func(ctx context.Context, email string) (db.User, error) {
			return db.User{
				ID:    1,
				Email: email,
			}, nil // User already exists
		},
	}

	service := NewUserService(mockQueries)
	_, err := service.CreateUser(context.Background(), "Jane Doe", "existing@example.com")

	if !errors.Is(err, ErrDuplicateEmail) {
		t.Errorf("expected ErrDuplicateEmail, got %v", err)
	}
}

func TestCreateUser_InvalidInput(t *testing.T) {
	tests := []struct {
		name  string
		email string
	}{
		{"", "test@example.com"},
		{"John Doe", ""},
		{"", ""},
	}

	service := NewUserService(&MockQueries{})

	for _, tt := range tests {
		_, err := service.CreateUser(context.Background(), tt.name, tt.email)
		if !errors.Is(err, ErrInvalidInput) {
			t.Errorf("expected ErrInvalidInput for name=%q email=%q, got %v", tt.name, tt.email, err)
		}
	}
}

func TestListUsers_Success(t *testing.T) {
	mockQueries := &MockQueries{
		ListUsersFunc: func(ctx context.Context, params db.ListUsersParams) ([]db.User, error) {
			return []db.User{
				{ID: 1, Name: "User 1", Email: "user1@example.com"},
				{ID: 2, Name: "User 2", Email: "user2@example.com"},
			}, nil
		},
		CountUsersFunc: func(ctx context.Context) (int64, error) {
			return 2, nil
		},
	}

	service := NewUserService(mockQueries)
	users, count, err := service.ListUsers(context.Background(), 10, 0)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(users) != 2 {
		t.Errorf("expected 2 users, got %d", len(users))
	}
	if count != 2 {
		t.Errorf("expected count 2, got %d", count)
	}
}

func TestUpdateUser_Success(t *testing.T) {
	now := time.Now()
	mockQueries := &MockQueries{
		GetUserByIDFunc: func(ctx context.Context, id int32) (db.User, error) {
			return db.User{
				ID:    1,
				Name:  "Old Name",
				Email: "old@example.com",
			}, nil
		},
		GetUserByEmailFunc: func(ctx context.Context, email string) (db.User, error) {
			return db.User{}, sql.ErrNoRows // No duplicate
		},
		UpdateUserFunc: func(ctx context.Context, params db.UpdateUserParams) (db.User, error) {
			return db.User{
				ID:        params.ID,
				Name:      params.Name,
				Email:     params.Email,
				UpdatedAt: timeToTimestamp(now),
			}, nil
		},
	}

	service := NewUserService(mockQueries)
	user, err := service.UpdateUser(context.Background(), 1, "New Name", "new@example.com")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if user.Name != "New Name" {
		t.Errorf("expected name 'New Name', got %s", user.Name)
	}
}

func TestUpdateUser_NotFound(t *testing.T) {
	mockQueries := &MockQueries{
		GetUserByIDFunc: func(ctx context.Context, id int32) (db.User, error) {
			return db.User{}, sql.ErrNoRows
		},
	}

	service := NewUserService(mockQueries)
	_, err := service.UpdateUser(context.Background(), 999, "Name", "email@example.com")

	if !errors.Is(err, ErrUserNotFound) {
		t.Errorf("expected ErrUserNotFound, got %v", err)
	}
}

func TestDeleteUser_Success(t *testing.T) {
	mockQueries := &MockQueries{
		GetUserByIDFunc: func(ctx context.Context, id int32) (db.User, error) {
			return db.User{ID: 1}, nil
		},
		DeleteUserFunc: func(ctx context.Context, id int32) error {
			return nil
		},
	}

	service := NewUserService(mockQueries)
	err := service.DeleteUser(context.Background(), 1)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestDeleteUser_NotFound(t *testing.T) {
	mockQueries := &MockQueries{
		GetUserByIDFunc: func(ctx context.Context, id int32) (db.User, error) {
			return db.User{}, sql.ErrNoRows
		},
	}

	service := NewUserService(mockQueries)
	err := service.DeleteUser(context.Background(), 999)

	if !errors.Is(err, ErrUserNotFound) {
		t.Errorf("expected ErrUserNotFound, got %v", err)
	}
}

func TestIsCorporateEmail(t *testing.T) {
	service := NewUserService(&MockQueries{})

	tests := []struct {
		email    string
		expected bool
	}{
		{"user@company.com", true},
		{"user@enterprise.com", true},
		{"user@gmail.com", false},
		{"user@example.com", false},
	}

	for _, tt := range tests {
		result := service.isCorporateEmail(tt.email)
		if result != tt.expected {
			t.Errorf("isCorporateEmail(%q) = %v, expected %v", tt.email, result, tt.expected)
		}
	}
}
