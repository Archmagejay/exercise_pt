package main

import "sort"

type cliCommand struct {
	name        string
	description string
	usage       string
	callback    func(*state, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a list of all commands",
			usage:       "help",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the program safely",
			usage:       "exit",
			callback:    commandExit,
		},
		"user": {
			name:        "user",
			description: "\n\tQuery the database for a specific user if no arguments are used\n\tList all users\n\tDetail the current user\n\tReset the user table\n\tRemove a specified user from the database",
			usage:       "user [list|me|reset|remove]",
			callback:    commandUser,
		},
		"register": {
			name:        "register",
			description: "Register a new user",
			usage:       "register",
			callback:    commandRegister,
		},
		"graph": {
			name:        "graph",
			description: "Graph available data",
			usage:       "graph <all|[specific_stat]>",
			callback:    commandGraph,
		},
		"daily": {
			name:        "daily",
			description: "Enter data for today or a specified date",
			usage:       "daily",
			callback:    commandDaily,
		},
		"goals": {
			name:        "goals",
			description: "\n\tList all goals, highlighting achieved ones for the current user if no arguments are used\n\tlist goals within a specific stat",
			usage:       "goals [specific_stat]",
			callback:    commandGoals,
		},
		"export": {
			name:        "export",
			description: "Export the database",
			usage:       "export [user]",
			callback:    commandExport,
		},
		"import": {
			name:        "import",
			description: "Import a database",
			usage:       "import <file_name>",
			callback:    commandImport,
		},
		"change": {
			name:        "change",
			description: "Change a saved value",
			usage:       "change <date>",
			callback:    commandChange,
		},
		"clear": {
			name:        "clear",
			description: "Clear the screen",
			usage:       "clear",
			callback:    commandClear,
		},
	}
}

func getCommandsOrdered() []cliCommand {
	cmds := getCommands()
	var keys []string
	for k := range cmds {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	var ordered = []cliCommand{}
	for _, k := range keys {
		ordered = append(ordered, cmds[k])
	}
	return ordered
}

func altCommands() map[string]string {
	return map[string]string{
		"?":     "help",
		"users": "user list",
	}
}

func cmdConfirmation(s *state) bool {
	s.in.Scan()
	if s.in.Text() == "y" || s.in.Text() == "Y" {
		return true
	}
	return false
}

func cmdInput(s *state) (string) {
	s.in.Scan()
	return s.in.Text()
}