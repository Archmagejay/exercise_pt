package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/archmagejay/excercise_pt/internal/database"
	"github.com/google/uuid"
)

func commandRegister(s *state, args ...string) error {
	NAME:
	fmt.Println("Please enter your desired username\nPress enter to cancel")

	fmt.Print("register/username > ")
	s.in.Scan()
	name := s.in.Text()
	if name == "" {
		return nil
	}
	if _, ok := badInputs[name]; ok {
		fmt.Print("That name is blacklisted, try again\n")
		goto NAME
	}
	if u, err := s.db.GetUserByName(context.Background(), name); err != nil {
		s.Log(LogError, err)
	} else {
		fmt.Printf("user [%s] already registered, load user [%s]? (y/n)\n> ", u.Name, u.Name)
		if cmdConfirmation(s) {
			s.cfg.SetUser(u)
			fmt.Printf("%sWelcome back [%s]\n%s", seperator, u.Name, seperator)
			return nil
		}
		goto NAME
	}
	fmt.Printf("Is %s your desired username? (y/n)\n > ", name)
	s.in.Scan()
	if strings.ToLower(s.in.Text()) != "y" {
		goto NAME
	}

	fmt.Println("Please enter your height in centimeters")

	HEIGHT:
	fmt.Print("register/height > ")
	s.in.Scan()
	height, err := strconv.ParseInt(s.in.Text(), 10, 32)
	if err != nil {
		fmt.Println("Use the format: ###")
		goto HEIGHT
	}

	user := database.NewUserParams{
		ID: uuid.New(),
		Name: name,
		Height: int32(height),
		StartDate: time.Now(),
	}

	if u, err := s.db.NewUser(context.Background(), user); err != nil {
		s.Log(LogError, err)
	} else {
	fmt.Print(seperator, "New user created: \n")
	fmt.Printf("* Name: %s\n* Height: %d\n* Starting date: %v\n", u.Name, u.Height, u.StartDate.Format(time.DateOnly))
	fmt.Print(seperator)
	s.cfg.SetUser(u)
	}
	return nil
}