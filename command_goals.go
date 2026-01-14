package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Archmagejay/exercise_pt/internal/database"
	"github.com/google/uuid"
)

type json_goals []struct {
	Type           string    `json:"type"`
	GoalDur        []string  `json:"goal_dur,omitempty"`
	GoalPlateCount [][]int32 `json:"goal_plate_count,omitempty"`
	GoalDecimal    []string  `json:"goal_decimal,omitempty"`
	GoalInt        []int32   `json:"goal_number,omitempty"`
}

func commandGoals(s *state, args ...string) error {
	if len(args) > 0 && args[0] != "" {
		if args[0] == "reset" {
			err := s.db.DeleteAllGoals(context.Background())
			if err != nil {
				return err
			}
			fmt.Println("All goals reset")
			s.Log(LogInfo, "goals table cleared")
			return nil
		}
		return ErrNotImplemented
	}
	if goals, err := s.db.GetAllGoals(context.Background()); len(goals) == 0 && err == nil {
		if err := importGoals(s, ""); err != nil {
			return err
		}
	} else if err != nil {
		return err
	} else {
		for _, goal := range goals {
			err = printGoalDebug(goal, false)
			if err != nil {
				s.Log(LogError, err)
			}
		}
	}

	return nil
}

func importGoals(s *state, _ string) error {
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
			GoalType: database.GoalTypes(g.Type),
		}
		switch g.Type {
		case "Treadmill":
			fallthrough
		case "Bike":
			for i, t := range g.GoalDecimal {
				goal.ID = uuid.New()
				goal.GoalTier = int32(i)
				goal.GoalDecimal = sql.NullString{
					String: fmt.Sprint(t),
					Valid:  true,
				}

				err := s.db.AddGoal(context.Background(), goal)
				if err != nil {
					return fmt.Errorf("error adding goal: %v\nerror: %w", goal, err)
				}
			}
		case "Plates":
			for i, t := range g.GoalPlateCount {
				goal.ID = uuid.New()
				goal.GoalTier = int32(i)
				goal.GoalType = database.GoalTypes(pcArr[i])
				if len(t) != 7 {
					s.Log(LogWarning, fmt.Sprintf("Incorrect array length for plate count in array: %d, %v", i+1, t))
					continue
				}
				goal.GoalPlateCount = t

				err := s.db.AddGoal(context.Background(), goal)
				if err != nil {
					return fmt.Errorf("error adding goal: %v\nerror: %w", goal, err)
				}
			}
		case "Weight":
			for i, t := range g.GoalDecimal {
				goal.ID = uuid.New()
				goal.GoalTier = int32(i)
				goal.GoalDecimal = sql.NullString{
					String: fmt.Sprint(t),
					Valid:  true,
				}

				err := s.db.AddGoal(context.Background(), goal)
				if err != nil {
					return fmt.Errorf("error adding goal: %v\nerror: %w", goal, err)
				}
			}
		case "Waist":
			for i, t := range g.GoalInt {
				goal.ID = uuid.New()
				goal.GoalTier = int32(i)
				goal.GoalNumber = sql.NullInt32{
					Int32: t,
					Valid: true,
				}

				err := s.db.AddGoal(context.Background(), goal)
				if err != nil {
					return fmt.Errorf("error adding goal: %v\nerror: %w", goal, err)
				}
			}
		case "Plank":
			fallthrough
		case "Park Run":
			for i, t := range g.GoalDur {
				goal.ID = uuid.New()
				goal.GoalTier = int32(i)
				dur, err := time.ParseDuration(t)
				if err != nil {
					return fmt.Errorf("error parsing time: %v\nerror: %w", t, err)
				}
				goal.GoalDur = sql.NullString{
					String: dur.String(),
					Valid:  true,
				}

				err = s.db.AddGoal(context.Background(), goal)
				if err != nil {
					return fmt.Errorf("error adding goal: %v\nerror: %w", goal, err)
				}
			}
		default:
			s.Log(LogWarning, fmt.Errorf("undefined goal type %s in %s", g.Type, path))
			continue
		}
	}

	fmt.Println("Goals imported successfully")

	s.Log(LogInfo, fmt.Sprintf("Successfully imported goals from: %s", path))
	return nil
}

func printGoalDebug(g database.Goal, debug bool) error {
	//fmt.Print(seperator)
	if debug {
		fmt.Printf("* ID: %v\n", g.ID)
	}
	if g.GoalTier == 0 {
		fmt.Printf("%s* Type: %s\n", seperator, g.GoalType)
	}

	fmt.Printf("* Tier #%d: %s\t", g.GoalTier, tierMap[g.GoalTier])
	var plates bool
	var goal, postfix string

	switch g.GoalType {
	case database.GoalTypesBike: // Distance
		fallthrough
	case database.GoalTypesTreadmill: // Distance
		goal = g.GoalDecimal.String + " kms"
	case database.GoalTypesWeight: // Kilograms
		goal = g.GoalDecimal.String + " kg"
	case database.GoalTypesParkRun: // Duration
		fallthrough
	case database.GoalTypesPlank: // Duration
		time := strings.Split(g.GoalDur.String, ":")
		for i, t := range time {
			if t == "00" {
				continue
			}
			switch i {
			case 0:
				goal += t + "h "
			case 1:
				goal += t + "m "
			case 2:
				goal += t + "s"
			}
		}
	case database.GoalTypesWaist: // Centimeters
		goal = fmt.Sprint(g.GoalNumber.Int32) + " cm"
	default: // Plate Counts
		plates = true
		for i, pc := range g.GoalPlateCount {
			goal += fmt.Sprintf("\n\t* %s:   \t%d Plates", pcArr[i], pc)
		}
	}
	if plates {
		postfix = "s"
	}
	fmt.Printf("* Goal%s: %s\n", postfix, goal)
	return nil
}
