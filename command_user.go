package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/archmagejay/excercise_pt/internal/database"
)

func commandUser(s *state, args ...string) error {
	if len(args) == 0 {
		fmt.Print(seperator, "Please enter a username to fetch the details of.\nPress enter to cancel\nUser ", prefix)
		INPUT:
		s.in.Scan()
		name := s.in.Text()
		if name == "" {
			return nil
		}
		user, err := s.db.GetUserByName(context.Background(), name)
		if err == sql.ErrNoRows {
			fmt.Printf("User [%s] not found. Is the capitalization and spelling correct?\nUser > ", name)
			goto INPUT
		} else if err != nil {
			return err
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
		if err != nil {
			return err
		}
		printUser(user)
	default:
		return ErrArgs
	}
	fmt.Print(seperator)
	return nil

}

func printUser(u database.User) {
	fmt.Print(seperator)
	fmt.Printf("* ID:         %s\n", u.ID)
	fmt.Printf("* Name:       %s\n", u.Name)
	fmt.Printf("* Height:     %d\n", u.Height)
	fmt.Printf("* Start date: %s\n", u.StartDate.Format(time.DateOnly))

}
