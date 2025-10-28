package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/example/speedrun-rest-api/api"
	"github.com/example/speedrun-rest-api/db"
	"github.com/example/speedrun-rest-api/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Server implements the ServerInterface from oapi-codegen
type Server struct {
	userService *service.UserService
}

// NewServer creates a new Server instance
func NewServer(queries *db.Queries) *Server {
	return &Server{
		userService: service.NewUserService(queries),
	}
}

// GetUser handles GET /users/{id}
// Retrieves a specific user by their ID
func (s *Server) GetUser(w http.ResponseWriter, r *http.Request, id int) {
	ctx := r.Context()
	
	user, err := s.userService.GetUserByID(ctx, int32(id))
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			writeError(w, http.StatusNotFound, "User not found", "USER_NOT_FOUND")
			return
		}
		log.Printf("Error getting user: %v", err)
		writeError(w, http.StatusInternalServerError, "Internal server error", "INTERNAL_ERROR")
		return
	}
	
	// Map database model to API model
	apiUser := dbUserToAPIUser(user)
	writeJSON(w, http.StatusOK, apiUser)
}

// ListUsers handles GET /users
// Retrieves a paginated list of users
func (s *Server) ListUsers(w http.ResponseWriter, r *http.Request, params api.ListUsersParams) {
	ctx := r.Context()
	
	// Set defaults
	limit := int32(10)
	offset := int32(0)
	
	if params.Limit != nil {
		limit = int32(*params.Limit)
	}
	if params.Offset != nil {
		offset = int32(*params.Offset)
	}
	
	users, total, err := s.userService.ListUsers(ctx, limit, offset)
	if err != nil {
		log.Printf("Error listing users: %v", err)
		writeError(w, http.StatusInternalServerError, "Internal server error", "INTERNAL_ERROR")
		return
	}
	
	// Map database models to API models
	apiUsers := make([]api.User, len(users))
	for i, user := range users {
		apiUsers[i] = dbUserToAPIUser(&user)
	}
	
	response := struct {
		Users  []api.User `json:"users"`
		Total  int64      `json:"total"`
		Limit  int32      `json:"limit"`
		Offset int32      `json:"offset"`
	}{
		Users:  apiUsers,
		Total:  total,
		Limit:  limit,
		Offset: offset,
	}
	
	writeJSON(w, http.StatusOK, response)
}

// CreateUser handles POST /users
// Creates a new user with the provided information
func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	var req api.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body", "INVALID_REQUEST")
		return
	}
	
	user, err := s.userService.CreateUser(ctx, req.Name, string(req.Email))
	if err != nil {
		if errors.Is(err, service.ErrDuplicateEmail) {
			writeError(w, http.StatusConflict, "User with this email already exists", "DUPLICATE_EMAIL")
			return
		}
		if errors.Is(err, service.ErrInvalidInput) {
			writeError(w, http.StatusBadRequest, "Invalid input", "INVALID_INPUT")
			return
		}
		log.Printf("Error creating user: %v", err)
		writeError(w, http.StatusInternalServerError, "Internal server error", "INTERNAL_ERROR")
		return
	}
	
	apiUser := dbUserToAPIUser(user)
	writeJSON(w, http.StatusCreated, apiUser)
}

// UpdateUser handles PUT /users/{id}
// Updates an existing user's information
func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request, id int) {
	ctx := r.Context()
	
	var req api.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body", "INVALID_REQUEST")
		return
	}
	
	name := ""
	email := ""
	if req.Name != nil {
		name = *req.Name
	}
	if req.Email != nil {
		email = string(*req.Email)
	}
	
	user, err := s.userService.UpdateUser(ctx, int32(id), name, email)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			writeError(w, http.StatusNotFound, "User not found", "USER_NOT_FOUND")
			return
		}
		if errors.Is(err, service.ErrDuplicateEmail) {
			writeError(w, http.StatusConflict, "Email already in use by another user", "DUPLICATE_EMAIL")
			return
		}
		log.Printf("Error updating user: %v", err)
		writeError(w, http.StatusInternalServerError, "Internal server error", "INTERNAL_ERROR")
		return
	}
	
	apiUser := dbUserToAPIUser(user)
	writeJSON(w, http.StatusOK, apiUser)
}

// DeleteUser handles DELETE /users/{id}
// Deletes a user by their ID
func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request, id int) {
	ctx := r.Context()
	
	err := s.userService.DeleteUser(ctx, int32(id))
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			writeError(w, http.StatusNotFound, "User not found", "USER_NOT_FOUND")
			return
		}
		log.Printf("Error deleting user: %v", err)
		writeError(w, http.StatusInternalServerError, "Internal server error", "INTERNAL_ERROR")
		return
	}
	
	w.WriteHeader(http.StatusNoContent)
}

// SetupRouter creates and configures the HTTP router
func SetupRouter(server *Server) http.Handler {
	r := chi.NewRouter()
	
	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	
	// Register handlers using oapi-codegen
	api.HandlerFromMux(server, r)
	
	return r
}

// Helper functions

// dbUserToAPIUser converts a database User model to an API User model
func dbUserToAPIUser(user *db.User) api.User {
	return api.User{
		Id:        int(user.ID),
		Name:      user.Name,
		Email:     openapi_types.Email(user.Email),
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
	}
}

// writeJSON writes a JSON response
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON: %v", err)
	}
}

// writeError writes an error response
func writeError(w http.ResponseWriter, status int, message, code string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := api.Error{
		Message: message,
		Code:    &code,
	}
	if err := json.NewEncoder(w).Encode(err); err != nil {
		log.Printf("Error encoding error response: %v", err)
	}
}

// Helper to parse int from path parameter
func parseIntParam(r *http.Request, key string) (int, error) {
	param := chi.URLParam(r, key)
	return strconv.Atoi(param)
}
