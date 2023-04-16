package handlers

import (
	"errors"
	"refactoring/internal/storage"
	"time"
)

const store = `users.json`

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
}

func NewUserApiHandler(file storage.FileContent) *userApiHandler {
	return &userApiHandler{file: file}
}
