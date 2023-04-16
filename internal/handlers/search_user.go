package handlers

import (
	"encoding/json"
	"github.com/go-chi/render"
	"net/http"
)

func (manager *userApiHandler) SearchUsers(w http.ResponseWriter, r *http.Request) {
	s := UserStore{}
	_ = json.Unmarshal(manager.file.File, &s)

	render.JSON(w, r, s.List)
}
