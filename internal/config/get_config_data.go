package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

func GetConfigFromFile(filename string) ServiceConfiguration {
	var cfg ServiceConfiguration

	f, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("unable to open config file '%s': %w", filename, err)
		return cfg
	}

	data, err := io.ReadAll(f)
	if err != nil {
		fmt.Errorf("unable to read from config file '%s': %w", filename, err)
		return cfg
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		fmt.Errorf("unable to unmarshall config data: %w", err)
		return cfg
	}
	return cfg
}
