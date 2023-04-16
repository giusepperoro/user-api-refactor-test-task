package handlers

import (
	"errors"
	"refactoring/internal/config"
	"refactoring/internal/storage"
	"time"
)

var UserNotFound = errors.New("user_not_found")

type (
	User struct {
		CreatedAt   time.Time `json:"created_at"`
		DisplayName string    `json:"display_name"`
		Email       string    `json:"email"`
	}
	UserList  map[string]User
	UserStore struct {
		Increment int      `json:"increment"`
		List      UserList `json:"list"`
	}
)

type CreateUserRequest struct {
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

type UpdateUserRequest struct {
	DisplayName string `json:"display_name"`
}

type userApiHandler struct {
	file storage.FileContent
	cfg  config.ServiceConfiguration
}

func NewUserApiHandler(file storage.FileContent, cfg config.ServiceConfiguration) *userApiHandler {
	return &userApiHandler{
		file: file,
		cfg:  cfg,
	}
}
