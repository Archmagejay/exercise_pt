package main

import (
	"fmt"
)

func commandHelp(s *state, args ...string) error {
	fmt.Print(seperator, "Usage key: \t<> required\t[] optional\t| option seperator\n", "Available commands:\n")
	for _, cmd := range getCommandsOrdered() {
		fmt.Printf("* %s: %s\n", cmd.usage, cmd.description)
	}
	fmt.Print(seperator)
	return nil
}

func commandExit(s *state, args ...string) error {
	shutdown(s)
	return nil
}

func commandClear(s *state, args ...string) error {
	return ErrNotImplemented
}
