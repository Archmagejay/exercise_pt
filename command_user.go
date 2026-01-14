package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Archmagejay/exercise_pt/internal/database"
)

func commandUser(s *state, args ...string) error {
	if len(args) == 0 {
		user, err := queryUser(s, false)
		if err != nil {
			return err
		}
		if user.Name == "" {
			return nil
		}
		printUser(user)
		fmt.Print(seperator)
		return nil
	}

	switch args[0] {
	case "list":
		users, err := s.db.ListUsers(context.Background())
		if err != nil {
			return err
		}
		if len(users) == 0 {
			fmt.Println("No users in database")
			return nil
		}
		fmt.Printf("%s*** %v users found ***\n%s", seperator, len(users), seperator)
		for _, user := range users {

			printUser(user)
		}
	case "me":
		user, err := s.db.GetUserByName(context.Background(), s.cfg.CurrentUserName)
		if err == sql.ErrNoRows {

		} else if err != nil {
			return err
		}
		printUser(user)
	case "reset":
		err := resetUsers(s)
		if err != nil {
			return err
		}
	case "remove":
		return removeUser(s)
	default:
		return ErrArgs
	}
	fmt.Print(seperator)
	return nil

}

func queryUser(s *state, headless bool) (database.User, error) {
	if !headless {
		fmt.Print(seperator, "Please enter a username to fetch the details of.\nPress enter to cancel\nUser ", prefix)
	}
INPUT:
	name := cmdInput(s)
	if name == "" {
		return database.User{}, nil
	}
	user, err := s.db.GetUserByName(context.Background(), name)
	if err == sql.ErrNoRows {
		fmt.Printf("User [%s] not found. Try again?\nPlease check the capitalization and spelling\nUser > ", name)
		goto INPUT
	} else if err != nil {
		return database.User{}, err
	}
	return user, nil
}

func resetUsers(s *state) error {
	fmt.Print("Are you sure? (y/n)\n>")

	if cmdConfirmation(s) {
		return s.db.DeleteAllUsers(context.Background())
	}

	return nil
}

func removeUser(s *state) error {
	fmt.Print("Please enter the user you want to remove\nPress enter to cancel\nUser ", prefix)
	user, err := queryUser(s, true)
	if err != nil {
		return err
	}
	err = s.db.DeleteUser(context.Background(), user.Name)
	if err != nil {
		return err
	}
	fmt.Printf("User [%s] has been removed from the database\n", user.Name)
	if user.Name == s.cfg.CurrentUserName {
		s.cfg.SetUser(database.User{})
	}
	return nil
}

func printUser(u database.User) {
	fmt.Print(seperator)
	fmt.Printf("* ID:         %s\n", u.ID)
	fmt.Printf("* Name:       %s\n", u.Name)
	fmt.Printf("* Height:     %d\n", u.Height)
	fmt.Printf("* Start date: %s\n", u.StartDate.Format(time.DateOnly))
}
