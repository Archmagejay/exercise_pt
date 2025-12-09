package main

import (
	"context"
	"fmt"

	"github.com/archmagejay/excercise_pt/internal/database"
)

func commandUser(s *state, args ...string) error {
	if len(args) == 0 {
		if s.cfg.CurrentUserName != "" {
			user, err := s.db.GetUserByName(context.Background(), s.cfg.CurrentUserName)
			if err.Error() == "sql: no rows in result set" {
				return fmt.Errorf("user <%s> not found in database", s.cfg.CurrentUserName)
			} else if err != nil {
				return err
			}
			printUser(user)
		} else {
			return fmt.Errorf("no user set")
		}
		return nil
	}

	if args[0] == "list" {
		users, err := s.db.ListUsers(context.Background())
		if err != nil {
			return err
		}
		if len(users) == 0 {
			fmt.Println("No users in database")

		}
		for _, user := range users {
			fmt.Println(user)
		}
	} else {
		user, err := s.db.GetUserByName(context.Background(), args[0])
		if err != nil {
			return err
		}
		printUser(user)
	}
	return nil
}

func printUser(u database.User) {
	fmt.Printf("* ID:         %s", u.ID)
	fmt.Printf("* Name:       %s", u.Name)
	fmt.Printf("* Height:     %d", u.Height)
	fmt.Printf("* Start date: %s", u.StartDate.Format("02/01/2006"))
}
