package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {

}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Exercise > ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		command, exists := getCommands()[commandName]
		if exists {
			if err := command.callback(cfg, args...); err != nil {
				fmt.Printf("Command: %s errored: %s\n", command.name, fmt.Errorf("%w", err))
			}
			continue
		} else {
			fmt.Printf("Unknown command: %s\n", commandName)
			fmt.Println("Try \"help\" for a list of commands")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	return strings.Fields(lower)
}

type cliCommand struct {
	name 			string
	description 	string
	callback func(*config, ...string) error
}

func getCommands() map[string]cliCommand{
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a list of all commands",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exits the program",
			callback: commandExit,
		},
	}
}