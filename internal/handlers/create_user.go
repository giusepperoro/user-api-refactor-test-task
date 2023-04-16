package handlers

import (
	"encoding/json"
	"github.com/go-chi/render"
	"io/fs"
	"io/ioutil"
	"net/http"
	"refactoring/internal/errors"
	"strconv"
	"time"
)

func (c *CreateUserRequest) Bind(r *http.Request) error { return nil }

func (manager *userApiHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	s := UserStore{}
	_ = json.Unmarshal(manager.file.File, &s)

	request := CreateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}

	s.Increment++
	u := User{
		CreatedAt:   time.Now(),
		DisplayName: request.DisplayName,
		Email:       request.DisplayName,
	}

	id := strconv.Itoa(s.Increment)
	s.List[id] = u

	b, _ := json.Marshal(&s)
	_ = ioutil.WriteFile(manager.cfg.FileName, b, fs.ModePerm)

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"user_id": id,
	})
}
