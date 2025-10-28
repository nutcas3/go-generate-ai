package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/example/speedrun-rest-api/db"
)

var (
	// ErrUserNotFound is returned when a user is not found
	ErrUserNotFound = errors.New("user not found")
	
	// ErrDuplicateEmail is returned when attempting to create a user with an existing email
	ErrDuplicateEmail = errors.New("user with this email already exists")
	
	// ErrInvalidInput is returned when input validation fails
	ErrInvalidInput = errors.New("invalid input")
)

// UserService handles business logic for user operations
type UserService struct {
	queries db.Querier
}

// NewUserService creates a new UserService instance
func NewUserService(queries db.Querier) *UserService {
	return &UserService{
		queries: queries,
	}
}

// GetUserByID retrieves a user by their ID
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - id: The user's unique identifier
//
// Returns:
//   - *db.User: The user object if found
//   - error: ErrUserNotFound if user doesn't exist, or database errors
func (s *UserService) GetUserByID(ctx context.Context, id int32) (*db.User, error) {
	user, err := s.queries.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	
	return &user, nil
}

// ListUsers retrieves a paginated list of users
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - limit: Maximum number of users to return
//   - offset: Number of users to skip
//
// Returns:
//   - []db.User: List of users
//   - int64: Total count of users
//   - error: Database errors if any
func (s *UserService) ListUsers(ctx context.Context, limit, offset int32) ([]db.User, int64, error) {
	users, err := s.queries.ListUsers(ctx, db.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list users: %w", err)
	}
	
	count, err := s.queries.CountUsers(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count users: %w", err)
	}
	
	return users, count, nil
}

// CreateUser creates a new user after performing validation and duplicate checks
//
// This is where business logic lives. We check for duplicate emails,
// validate input, and potentially call external services.
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - name: User's full name
//   - email: User's email address
//
// Returns:
//   - *db.User: The created user object
//   - error: ErrDuplicateEmail, ErrInvalidInput, or database errors
func (s *UserService) CreateUser(ctx context.Context, name, email string) (*db.User, error) {
	// Validate input
	if name == "" || email == "" {
		return nil, ErrInvalidInput
	}
	
	// Check for duplicate email
	existing, err := s.queries.GetUserByEmail(ctx, email)
	if err == nil && existing.ID != 0 {
		return nil, ErrDuplicateEmail
	}
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("failed to check for duplicate email: %w", err)
	}
	
	// Business Logic: In a real application, you might:
	// - Call an email verification service
	// - Check against a blocklist
	// - Apply business rules (e.g., require approval for certain domains)
	// - Send a welcome email
	// - Log to an audit trail
	
	// For demo purposes, we'll add a simple business rule:
	// Users with corporate emails get special handling
	if s.isCorporateEmail(email) {
		// In production, this might trigger an approval workflow
		// or set special flags on the user account
	}
	
	// Create the user
	user, err := s.queries.CreateUser(ctx, db.CreateUserParams{
		Name:  name,
		Email: email,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	
	return &user, nil
}

// UpdateUser updates an existing user's information
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - id: User ID to update
//   - name: New name (optional, empty string means no change)
//   - email: New email (optional, empty string means no change)
//
// Returns:
//   - *db.User: The updated user object
//   - error: ErrUserNotFound, ErrDuplicateEmail, or database errors
func (s *UserService) UpdateUser(ctx context.Context, id int32, name, email string) (*db.User, error) {
	// First, verify the user exists
	existing, err := s.queries.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	
	// Use existing values if not provided
	if name == "" {
		name = existing.Name
	}
	if email == "" {
		email = existing.Email
	}
	
	// Check for duplicate email if email is changing
	if email != existing.Email {
		duplicate, err := s.queries.GetUserByEmail(ctx, email)
		if err == nil && duplicate.ID != 0 && duplicate.ID != id {
			return nil, ErrDuplicateEmail
		}
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("failed to check for duplicate email: %w", err)
		}
	}
	
	// Update the user
	user, err := s.queries.UpdateUser(ctx, db.UpdateUserParams{
		ID:    id,
		Name:  name,
		Email: email,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}
	
	return &user, nil
}

// DeleteUser deletes a user by their ID
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - id: User ID to delete
//
// Returns:
//   - error: ErrUserNotFound if user doesn't exist, or database errors
func (s *UserService) DeleteUser(ctx context.Context, id int32) error {
	// First verify the user exists
	_, err := s.queries.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrUserNotFound
		}
		return fmt.Errorf("failed to get user: %w", err)
	}
	
	// Business Logic: In a real application, you might:
	// - Soft delete instead of hard delete
	// - Check if user has dependent records
	// - Send a deletion confirmation email
	// - Log to an audit trail
	
	err = s.queries.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	
	return nil
}

// isCorporateEmail checks if an email belongs to a corporate domain
// This is an example of business logic that you would implement
func (s *UserService) isCorporateEmail(email string) bool {
	// Simple example - in production, this would be more sophisticated
	corporateDomains := []string{"@company.com", "@enterprise.com"}
	
	for _, domain := range corporateDomains {
		if len(email) > len(domain) && email[len(email)-len(domain):] == domain {
			return true
		}
	}
	
	return false
}

// Helper function to convert time.Time to string for API responses
func formatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}
