package storage

import "net/http"

type FileContent struct {
	File []byte
}

type StManager interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	SearchUsers(w http.ResponseWriter, r *http.Request)
}
