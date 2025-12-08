package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"github.com/archmagejay/excercise_pt/internal/database"
)

func commandRegister(s *state, args ...string) error {
	if len(args) != 1 {
		return ErrArgs
	}

	fmt.Println("Please enter your desired username")
	fmt.Print("> ")
	s.in.Scan()
	name := s.in.Text()
	if u, err := s.db.GetUserByName(context.Background(), name); err == nil {
		log.Println(err)
	} else {
		log.Println(u)
	}


	height, err := strconv.ParseInt(s.in.Text(), 10, 32)
	if err != nil {
		return err
	}

	user := database.NewUserParams{
		Name: name,
		Height: int32(height),
	}

	s.db.NewUser(context.Background(), user)

	return nil
}