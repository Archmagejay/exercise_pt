package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

const configDir = "Exercise_PT"
const configFileName = "config.json"
const db_url string = "postgres://postgres:postgres@localhost:5432/exercise_pt"

var ErrMissingUser = errors.New("no user set")
var ErrDBURL = errors.New("invalid database url")
var ErrTime = errors.New("time not initialized")

type Config struct {
	DBURL                           string    `json:"db_url"`
	CurrentUserName                 string    `json:"current_user_name"`
	LastOpened                      time.Time `json:"last_opened"`
	valid, errUser, errTime, daily bool
}


// Write config file to disk
func (cfg *Config) SaveConfig() error {
	return write(cfg)
}

// Read the config file into memory
// Creating the file if it does not exist
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
		cfg.DBURL, cfg.LastOpened = db_url, time.Now()
		err2 := write(&cfg)
		if err2 != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	return &cfg, nil
}

// Get the filepath to the config file
// Creating the directory if needed
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

// Write the config file to disk
func write(cfg *Config) error {
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

// Validate the config file
func (cfg *Config) Validate() error {
	cfg.valid = false
	cfg.errUser = false
	cfg.errTime = false
	cfg.daily = false
	if cfg.DBURL != db_url {
			return ErrDBURL
		}
	day, err := time.ParseDuration("24h")
	if err != nil {
		fmt.Println("Time package failed")
		return ErrTime
	}
	if cfg.CurrentUserName == "" {
		cfg.errUser = true
		return ErrMissingUser
	}
	if cfg.LastOpened.Before(time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)) {
		cfg.errTime = true
		return ErrTime
	} else if cfg.LastOpened.Before(time.Now().Local().Add(-day)) {
		cfg.daily = true
	}
	cfg.valid = true
	return nil
}

// Set the username of the current user in memory and then write to disk
func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	return write(cfg)
}

// Set the last opened timestamp to now the write to disk
func (cfg *Config) SetTime() error {
	cfg.LastOpened = time.Now()
	return write(cfg)
}

// Check if the config was successfully validated
func (cfg *Config) IsValid() bool {
	cfg.Validate()
	return cfg.valid
}

// Check if the config has a valid username
func (cfg *Config) IsValidUser() bool {
	cfg.Validate()
	return !cfg.errUser && cfg.valid
}

// Check if the config has a valid timestamp
// i.e if the timestamp is from before 2025
func (cfg *Config) IsValidTime() bool {
	cfg.Validate()
	return !cfg.errTime && cfg.valid
}

// Check if a daily entry is required due to the last open time being more than 24 hours ago
func (cfg *Config) IsDailyDue() bool {
	cfg.Validate()
	return !cfg.daily && cfg.valid
}