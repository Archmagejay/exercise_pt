package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"github.com/archmagejay/excercise_pt/internal/database"
	"github.com/google/uuid"
)

type json_goals []struct {
	Type           string `json:"type"`
	GoalSpeed      []int  `json:"goal_speed,omitempty"`
	GoalDur        []string  `json:"goal_dur,omitempty"`
	GoalPlateCount [][]int  `json:"goal_plate_count,omitempty"`
	GoalDecimal    []float32  `json:"goal_decimal,omitempty"`
	GoalNumber     []int  `json:"goal_number,omitempty"`
}

func commandGoals(s *state, args ...string) error {
	if goals, err := s.db.GetAllGoals(context.Background()); len(goals) == 0 && err == nil {
		if err := defaultGoals(s, ""); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}

func defaultGoals(s *state, _ string) error {
	path, err := os.Getwd()
	if err != nil {
		s.Log(LogFatal, err)
	}
	//TODO: Allow non hard coded goal jsons
	path += "/davids_goals.json"

	f, err := os.Open(path)
	if err != nil {
		s.Log(LogFatal, err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	goal_array := json_goals{}
	err = decoder.Decode(&goal_array)
	if err != nil {
		s.Log(LogError, err)
	}

	for _, g := range goal_array {
		goal := database.AddGoalParams{
			ID: uuid.New(),
			Type: g.Type,
		}
		switch g.Type {
			case "Treadmill":
			case "Bike":
			case "Plates":
			case "Weight":
			case "Waist":
				for i, t := range g.GoalNumber {
					goal.GoalTier = int32(i)
					goal.GoalNumber = sql.NullInt32{
						Int32: int32(t),
						Valid: true,
					}
					s.db.AddGoal(context.Background(), goal)
				}
			case "Plank":
			case "Park Run":
			default:
				s.Log(LogWarning, fmt.Errorf("undefined goal type in default_goals.json"))
		}
	}
	return nil
}