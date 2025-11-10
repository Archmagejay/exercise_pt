package main

import (
	"fmt"
	"os"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("====================")
	fmt.Println("Welcome to the Exercise Planner and Tracker!")
	fmt.Println("Available commands:")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("====================")
	return nil
}

func commandExit(cfg *config, args ...string) error {

	fmt.Println("Closing planner... Goodbye!")
	os.Exit(0)
	return nil
}