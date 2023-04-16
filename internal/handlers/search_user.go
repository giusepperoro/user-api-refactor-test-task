package handlers

import (
	"encoding/json"
	"github.com/go-chi/render"
	"io/ioutil"
	"net/http"
)

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	f, _ := ioutil.ReadFile(store)
	s := UserStore{}
	_ = json.Unmarshal(f, &s)

	render.JSON(w, r, s.List)
}
