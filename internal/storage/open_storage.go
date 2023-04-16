package storage

import (
	"fmt"
	"io/ioutil"
	"refactoring/internal/config"
)

func New(store config.ServiceConfiguration) (*FileContent, error) {
	f, err := ioutil.ReadFile(store.FileName)
	if err != nil {
		return nil, fmt.Errorf("unable to open config file '%s': %w", f, err)
	}
	return &FileContent{File: f}, nil
}
