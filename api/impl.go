package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"invia/api/restapi"
	"invia/internal/app/services"
)

type RestApiServer struct {
	service services.IService
	log     *slog.Logger
}

func NewRestApiServer(service services.IService, log *slog.Logger) RestApiServer {
	return RestApiServer{
		service: service,
		log:     log,
	}
}

func (s RestApiServer) ListUsers(w http.ResponseWriter, r *http.Request, params restapi.ListUsersParams) {
	s.log.Debug("RestApiServer.ListUsers")
	page := 1
	limit := 10

	if params.Page != nil && *params.Page > 0 {
		page = *params.Page
	}

	if params.Limit != nil && *params.Limit > 0 {
		limit = *params.Limit
	}

	users, err := s.service.ListUsers(r.Context(), page, limit)
	if err != nil {
		s.log.Error("Failed to get users", "error", err)
		http.Error(w, "Failed to get users", http.StatusInternalServerError)

		return
	}

	var resp []restapi.User
	for _, u := range users {
		resp = append(
			resp, restapi.User{
				Id:        u.Id,
				FirstName: u.FirstName,
				LastName:  u.LastName,
				Email:     u.Email,
				IsActive:  u.IsActive,
			},
		)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		s.log.Error("Failed to encode response", "error", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)

		return
	}
}

func (s RestApiServer) CreateUser(w http.ResponseWriter, r *http.Request) {
	s.log.Debug("RestApiServer.CreateUser")
	var newUser restapi.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		s.log.Error("Invalid request payload", "error", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)

		return
	}

	if newUser.Email == nil || newUser.FirstName == nil || newUser.LastName == nil || newUser.Password == nil {
		s.log.Error("Missing required fields", "error", err)
		http.Error(w, "Missing required fields", http.StatusBadRequest)

		return
	}

	now := time.Now()
	newUser.CreatedAt = &now
	newUser.UpdatedAt = &now
	isActive := true
	newUser.IsActive = &isActive

	hashedPassword, err := HashPassword(*newUser.Password)
	if err != nil {
		s.log.Error("Failed to hash password", "error", err)
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)

		return
	}
	newUser.Password = &hashedPassword

	err = s.service.AddUser(r.Context(), &newUser)
	if err != nil {
		s.log.Error("Failed to create user", "error", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(newUser)
	if err != nil {
		s.log.Error("Failed to encode response", "error", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (s RestApiServer) DeleteUser(w http.ResponseWriter, r *http.Request, id int) {
	s.log.Debug("RestApiServer.DeleteUser")
	err := s.service.DeleteUser(r.Context(), id)
	if err != nil {
		s.log.Error("User not found", "error", err)
		http.Error(w, "User not found", http.StatusNotFound)

		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s RestApiServer) GetUserById(w http.ResponseWriter, r *http.Request, id int) {
	s.log.Debug("RestApiServer.GetUserById")
	user, err := s.service.GetUserById(r.Context(), id)
	if err != nil {
		s.log.Error("User not found", "error", err)
		http.Error(w, "User not found", http.StatusNotFound)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		s.log.Error("Failed to encode response", "error", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (s RestApiServer) UpdateUser(w http.ResponseWriter, r *http.Request, id int) {
	s.log.Debug("RestApiServer.UpdateUser")
	var updatedUser restapi.User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		s.log.Error("Invalid request payload", "error", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)

		return
	}

	existingUser, err := s.service.GetUserById(r.Context(), id)
	if err != nil {
		s.log.Error("User not found", "error", err)
		http.Error(w, "User not found", http.StatusNotFound)

		return
	}

	if updatedUser.FirstName != nil {
		existingUser.FirstName = updatedUser.FirstName
	}
	if updatedUser.LastName != nil {
		existingUser.LastName = updatedUser.LastName
	}
	if updatedUser.Email != nil {
		existingUser.Email = updatedUser.Email
	}
	if updatedUser.IsActive != nil {
		existingUser.IsActive = updatedUser.IsActive
	}

	if updatedUser.Password != nil {
		hashedPassword, err := HashPassword(*updatedUser.Password)
		if err != nil {
			s.log.Error("Failed to hash password", "error", err)
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)

			return
		}
		existingUser.Password = &hashedPassword
	}

	now := time.Now()
	existingUser.UpdatedAt = &now

	err = s.service.UpdateUser(r.Context(), id, existingUser)
	if err != nil {
		s.log.Error("Failed to update user", "error", err)
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(existingUser)
	if err != nil {
		s.log.Error("Failed to encode response", "error", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
