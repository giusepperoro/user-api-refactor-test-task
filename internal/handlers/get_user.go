package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

func (manager *userApiHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	s := UserStore{}
	_ = json.Unmarshal(manager.file.File, &s)

	id := chi.URLParam(r, "id")

	render.JSON(w, r, s.List[id])
}
