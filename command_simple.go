package main

import (
	"errors"
	"fmt"
	"os"
)

func commandHelp(s *state, args ...string) error {
	fmt.Println("====================")
	fmt.Println("Welcome to the Exercise Tracker!")
	fmt.Println("Available commands:")
	for _, cmd := range getCommands() {
		fmt.Printf("* %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("====================")
	return nil
}

func commandExit(s *state, args ...string) error {
	fmt.Println("Closing tracker... Goodbye!")
	err := s.cfg.SaveConfig()
	if err != nil {
		return err
	}
	os.Exit(0)
	return nil
}


func commandClear(s *state, args ...string) error {

	return errors.New("not Implemented")
}