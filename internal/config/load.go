package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

var ErrInvalidFileType = errors.New("config file must be .json")

func Load(path string) (*Config, error) {
	if filepath.Ext(path) != ".json" {
		return nil, ErrInvalidFileType
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}

	cfg.applyDefaults()
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}
