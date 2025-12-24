package main

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/archmagejay/excercise_pt/internal/database"
	"github.com/google/uuid"
)

func commandDaily(s *state, _ ...string) error {
	if s.cfg.IsDailyDue() {
		return dailyEntry(s)
	}

	if latest, err := s.db.GetLatestEntryTimestampForUser(context.Background(), s.cfg.GetUserID()); err != nil {
		return err
	} else if latest.Before(time.Now().Add(-time.Duration(12) * time.Hour)) {
		return dailyEntry(s)
	}

	fmt.Println("You already have an entry in the last 12 hours.\nAre you sure you want to add another? (y/n)")
	if cmdConfirmation(s) {
		return dailyEntry(s)
	}
	return nil
}

func dailyEntry(s *state) error {
	// Make an entry struc with fields that shouldn't change
	entry := database.AddEntryParams{
		ID:     uuid.New(),
		UserID: s.cfg.GetUserID(),
	}

	// Check if the date needs to be input otherwise set it to now
	fmt.Println("Is this an entry for today? (y/n) > ")
	if !cmdConfirmation(s) {
		// Ask what day it is for
		// TODO: Add option for number of days prior as valid input
		fmt.Println("What day is it for? <dd/mm/yy> > ")
	DATE:
		entry_date_string := cmdInput(s)
		entry_date, err := time.ParseInLocation("_2/1/06", entry_date_string, time.Local)
		if err != nil {
			fmt.Println("Unable to parse input please use the format <dd/mm/yy>\ndaily >")
			goto DATE
		}
		// TODO: Create safeguards for ridiculous dates eg dates outside of 1 month from the present
		entry.Date = entry_date
	} else {
		entry.Date = time.Now()
	}

	// Check if the weekly data is due
WEEKLY:
	if weekly, err := s.db.GetLatestWeeklyDataTimestampForUser(context.Background(), s.cfg.GetUserID()); err != nil {
		return err
	} else if weekly.Before(time.Now().AddDate(0, 0, -7)) {
		// Ask for the weekly data
		// TODO: Sanity check inputs
		fmt.Println("Please enter your weight in kg using the format <###.##>: ")
		entry.Weight = cmdInput(s)

		fmt.Println("Please enter your waist measurement in cm: ")
		entry.Waist = cmdInput(s)

		// Confirm correct data entry
		fmt.Printf("Are these fields correct?\n\tWeight: %s kg\n\tWaist: %s cm", entry.Weight, entry.Waist)
		if !cmdConfirmation(s) {
			goto WEEKLY
		}
	}

	// If a saturday ask if a park run was done
PARK_RUN:
	if entry.Date.Weekday() == time.Saturday {
		fmt.Println("Did you complete a park run today? (y/n) > ")
		if cmdConfirmation(s) {
			fmt.Println("What time did you achieve? <mm:ss>")
			// TODO: Sanity check input
			park_run_string := cmdInput(s)
			park_run_string = strings.ReplaceAll(park_run_string, ":", "m") + "s"
			entry.ParkRun = sql.NullString{
				String: park_run_string,
				Valid:  true,
			}
			// Confirm correct data entry
			fmt.Printf("Are these fields correct?\n\tPark Run duration: %s", entry.ParkRun.String)
			if !cmdConfirmation(s) {
				goto PARK_RUN
			}
		}
	}

	// Ask what type of cardio
//CARDIO:
	fmt.Println("Did you use the treadmill? (y/n)\ndaily > ")
	entry.CardioType = !cmdConfirmation(s)
	// TODO: Clean this up
	/* if cmdConfirmation(s) {
		entry.CardioType = false
	} else {
		fmt.Println("The bike? (y/n) > ")
		if cmdConfirmation(s) {
			entry.CardioType = true
			goto PLANK
		} else {
			fmt.Println("No cardio then? (y/n)\ndaily > ")
			if cmdConfirmation(s) {
				entry.CardioType = false
				entry.Cardio = ""
			}
		}
	} */

	// Ask for plank duration
PLANK:
	fmt.Print("How long did you manage to hold the plank? <mm:ss>\ndaily > ")
	// TODO: Sanity check input
	entry_plank := cmdInput(s)
	entry_plank = strings.ReplaceAll(entry_plank, ":", "m") + "s"
	entry.PlankDur = sql.NullString{
		String: entry_plank,
		Valid:  true,
	}
	// Confirm correct data entry
		fmt.Printf("Are these fields correct?\n\tPlank Duration: %s \ndaily >", entry.PlankDur.String)
		if !cmdConfirmation(s) {
			goto PLANK
		}

	// Ask about the plate counts for weights
WEIGHTS:
	prev_plate_count, err := s.db.GetLatestPlateCountForUser(context.Background(), s.cfg.GetUserID())
	if err != nil {
		return err
	}
	entry.PlateCount = prev_plate_count
	fmt.Print("Have you changed any of the plate counts for your exercise? (y/n)\ndaily > ")
	if cmdConfirmation(s) {

		fmt.Println("Which exercises have changed? List them in this format <#[,#]>")
		for i, count := range prev_plate_count {
			fmt.Printf("\t%d. %s: %d\n", i, pcArr[i], count)
		}
		fmt.Print("daily > ")
	WEIGHT_CHANGED:
		changed_str := strings.Split(cmdInput(s), ",")
		if len(changed_str) != 0 {
			changed_idx := []int{}
			for _, str := range changed_str {
				idx, err := strconv.Atoi(str)
				if err != nil {
					fmt.Println("Unable to parse input please use the format <#[,#]>\ndaily > ")
					goto WEIGHT_CHANGED
				}
				changed_idx = append(changed_idx, idx)
			}
			for _, idx := range changed_idx {
				fmt.Printf("Please enter the new plate count for %s > ", pcArr[idx])
			NEW_COUNT:
				new_count, err := strconv.Atoi(cmdInput(s))
				if err != nil {
					fmt.Println("Unable to parse input please use the format <##>\ndaily > ")
					goto NEW_COUNT
				}
				entry.PlateCount[idx] = int32(new_count)
			}
		}
	}
	// Confirm correct data entry
		fmt.Println("Are these fields correct?")
		for i, count := range entry.PlateCount {
			fmt.Printf("\t%s: %d\n", pcArr[i], count)
		}
		fmt.Print("daily > ")
		if !cmdConfirmation(s) {
			goto WEIGHTS
		}

	s.db.AddEntry(context.Background(), entry)
	fmt.Println("New entry added to database!")
	
	// Check if any new goals have been achieved
	// TODO: FIX THIS
	/* if new_achi, err := checkNewAchieved(s, entry); err != nil {
		return err
	} else if new_achi {

	} */

	return nil
}

func checkNewAchieved(s *state, entry database.AddEntryParams) (bool, error) {
	current_achieved_goals, err := s.db.GetGoalsAchievedByUser(context.Background(), s.cfg.GetUserID())
	if err != nil {
		return false, err
	}
	var cur_ach_goal_map = map[uuid.UUID]int32{}
	var new_ach_map = []database.Goal{}
	if len(current_achieved_goals) != 0 {
		for _, id := range current_achieved_goals {
			tier, err := s.db.GetGoalTierByID(context.Background(), id.GoalID)
			if err != nil {
				return false, err
			}
			cur_ach_goal_map[id.GoalID] = tier
		}
	}

	goals, err := s.db.GetAllGoals(context.Background())
	for _, goal := range goals {
		if _, ok := cur_ach_goal_map[goal.ID]; ok {
			continue
		}
		switch goal.GoalType {
		case database.GoalTypesBike:
			if g_cardio, err := s.db.GetNextGoalTier(context.Background(), database.GetNextGoalTierParams{
				GoalType: database.GoalTypesBike,
				GoalTier: 0,
			}); err != nil {
				return false, err
			} else if entry.Cardio > g_cardio.GoalDecimal.String {
				new_ach_map = append(new_ach_map, goal)
			}
		case database.GoalTypesTreadmill:

		case database.GoalTypesWeight:
		case database.GoalTypesPlank:
		case database.GoalTypesParkRun:
		case database.GoalTypesWaist:
		default:

		}
	}

	return false, nil
}
