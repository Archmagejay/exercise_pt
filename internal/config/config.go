package config

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"time"
)
const configDir = "Exercise_PT"
const configFileName = "config.json"
const db_url = "postgres://postgres:postgres@localhost:5432/exercise_pt"

var ErrMissingUser = errors.New("no user set")
var ErrDBURL = errors.New("invalid database url")
var ErrTime = errors.New("time not initialized")

type Config struct {
	DBURL string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
	LastOpened time.Time `json:"last_opened"`
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	return write(*cfg)
}

func (cfg *Config) SetTime() error {
	cfg.LastOpened = time.Now()
	return write(*cfg)
}

func (cfg *Config) SaveConfig() error {
	return write(*cfg)
}

func Read() (*Config, error) {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile(fullPath, os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err == io.EOF {
		err := write(Config{
			DBURL: db_url,
			CurrentUserName: "",
			LastOpened: time.Now(),
		})
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func getConfigFilePath() (string, error) {
	dir, dirErr := os.UserConfigDir()
	if dirErr != nil {
		return "", dirErr
	}
	fullPath := filepath.Join(dir, configDir, configFileName)

	err := os.MkdirAll(filepath.Dir(fullPath), 0700)
	if err != nil {
		return "", err
	}

	return fullPath, nil
}

func write(cfg Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(cfg); err != nil {
		return err
	}

	return nil
}

func (cfg *Config) Validate() error {
	if cfg.CurrentUserName != "" {
		return ErrMissingUser
	}
	if cfg.DBURL != db_url {
		return ErrDBURL
	}
	if cfg.LastOpened.Before(time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)) {
		return ErrTime
	}
	return nil
}

