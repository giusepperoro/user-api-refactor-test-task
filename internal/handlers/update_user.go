package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"io/fs"
	"io/ioutil"
	"net/http"
	"refactoring/internal/errors"
)

func (c *UpdateUserRequest) Bind(r *http.Request) error { return nil }

func (manager *userApiHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	s := UserStore{}
	_ = json.Unmarshal(manager.file.File, &s)

	request := UpdateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, errors.ErrInvalidRequest(err))
		return
	}

	id := chi.URLParam(r, "id")

	if _, ok := s.List[id]; !ok {
		_ = render.Render(w, r, errors.ErrInvalidRequest(UserNotFound))
		return
	}

	u := s.List[id]
	u.DisplayName = request.DisplayName
	s.List[id] = u

	b, _ := json.Marshal(&s)
	_ = ioutil.WriteFile(manager.cfg.FileName, b, fs.ModePerm)

	render.Status(r, http.StatusNoContent)
}
