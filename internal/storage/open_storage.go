package storage

import (
	"io/ioutil"
	"refactoring/internal/config"
)

func New(store config.ServiceConfiguration) *FileContent {
	f, _ := ioutil.ReadFile(store.FileName)
	return &FileContent{File: f}
}
