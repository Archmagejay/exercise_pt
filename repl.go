package main

import (
	"fmt"
	"strings"
)

// Reusable text snippets
const (
	seperator = "====================\n"
	welcome   = "Welcome to the Exercise Tracker!\n"
	commands  = "Available commands:\n"
	prefix    = "> "
)

var reservedInputs = map[string]struct{}{
	"":{},
	"cancel":{},
	"y":{},
	"Y":{},
	"n":{},
	"N":{},
}

// The main REPL CLI function
func startRepl(s *state) {
	s.Log(LogInfo, "Program started")
	printWelcome()
	for {
		fmt.Print(prefix)
		s.in.Scan()
		words := cleanInput(s.in.Text())
		// If there is no input wait until there is
		if len(words) == 0 {
			continue
		}
		// Extract the command from input
		commandName := words[0]
		args := []string{}
		// if more than one word input extract the arguments
		if len(words) > 1 {
			args = words[1:]
		}
		// Check if input command matches an alternate
		alt, exists := altCommands()[commandName]
		if exists {
			alt := strings.Fields(alt)
			commandName = alt[0]
			if len(alt) > 1 {
				args = alt[1:]
			}
		}
		// Check if the command exists before attempting to run it
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(s, args...)
			if err == ErrArgs {
				// If the returned error was an ErrArgs print the usage template of the command
				fmt.Println("Usage: ", command.usage)
			} else if err == ErrNotImplemented {
				// If the command is not implemented say so and await futher input
				fmt.Println("This command is not yet implemented try another command")
			} else if err != nil {
				// Otherwise log the error and await new input
				s.Log(LogError, err)
				fmt.Printf("Command: <%s> errored see log for details\n", command.name)
			}
			continue
		} else {
			// If an invalid command is entered print a hint for the help command
			fmt.Printf("Unknown command: %s\n", commandName)
			fmt.Println("Try typing \"help\" for a list of available commands")
			continue
		}
	}
}

// Utility function
//
// Force input to lowercase and then seperate into individual fields
func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	return strings.Fields(lower)
}

func printWelcome() {
	fmt.Print(seperator, welcome, seperator)
}
