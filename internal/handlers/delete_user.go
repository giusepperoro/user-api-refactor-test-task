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

func (manager *userApiHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	s := UserStore{}
	_ = json.Unmarshal(manager.file.File, &s)

	id := chi.URLParam(r, "id")

	if _, ok := s.List[id]; !ok {
		_ = render.Render(w, r, errors.ErrInvalidRequest(UserNotFound))
		return
	}

	delete(s.List, id)

	b, _ := json.Marshal(&s)
	_ = ioutil.WriteFile(manager.cfg.FileName, b, fs.ModePerm)

	render.Status(r, http.StatusNoContent)
}
