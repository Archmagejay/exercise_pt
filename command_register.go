package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"github.com/archmagejay/excercise_pt/internal/database"
)

func commandRegister(s *state, args ...string) error {
	if len(args) != 1 {
		return ErrArgs
	}

	fmt.Println("Please enter your desired username")

	NAME:
	fmt.Print("> ")
	s.in.Scan()
	name := s.in.Text()

	if u, err := s.db.GetUserByName(context.Background(), name); err != nil {
		s.Log(LogError, err)
	} else {
		fmt.Printf("user [%s] already registered, try again?\n", u.Name)
		goto NAME
	}
	fmt.Printf("Is %s your desired username? (y/n)\n > ", name)
	s.in.Scan()
	if strings.ToLower(s.in.Text()) != "y" {
		goto NAME
	}

	fmt.Println("Please enter your height in centimeters")

	HEIGHT:
	fmt.Print("> ")
	s.in.Scan()
	height, err := strconv.ParseInt(s.in.Text(), 10, 32)
	if err != nil {
		fmt.Println("Use the format: ###")
		goto HEIGHT
	}

	user := database.NewUserParams{
		Name: name,
		Height: int32(height),
	}

	if err := s.db.NewUser(context.Background(), user); err != nil {
		s.Log(LogError, err)
	}

	return nil
}