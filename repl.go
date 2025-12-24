package main

import (
	"fmt"
	"strings"
)

// Reusable text snippets
const (
	seperator = "====================\n"
	welcome   = "Welcome to the Exercise Tracker!\n"
	prefix    = "> "
	//timestampFormat = "02/01/06 15:04:05"
	noUserInCfg = "!!! No user detected. Please run the <register> command or <exit> to quit !!!"
)

var badInputs = map[string]struct{}{
	"":       {},
	"cancel": {},
	"y":      {},
	"Y":      {},
	"n":      {},
	"N":      {},
	"exit":   {},
}

var pcArr = []string{
	"Bench Press",
	"Bisep Curls",
	"Lateral Pulldown",
	"Pectoral Fly",
	"Quad Curls",
	"Trapezius Lift",
	"Trisep Curls",
}

var tierMap = map[int32]string{
	0: "Ok",
	1: "Good",
	2: "Great",
	3: "Superb",
}
// The main REPL CLI function
func startRepl(s *state) {
	s.Log(LogInfo, "Program started")

	// Print the welcome message
	fmt.Print(seperator, welcome, seperator)

	for {
		// Check if there is a valid user in the config
		if !s.cfg.IsValidUser() {
			// TODO: May be temporary and be replaced with a startup form later
			fmt.Println(noUserInCfg)
		}
		fmt.Print(prefix)
		s.in.Scan()
		words := cleanInput(s.in.Text())
		// If there is no input wait until there is
		if len(words) == 0 {
			continue
		}
		// If there is no user set in the config do not continue until the register command has been run
		if !s.cfg.IsValidUser() && words[0] != "register" && words[0] != "exit" {
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
