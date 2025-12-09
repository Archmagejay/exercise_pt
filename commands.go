package main

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
			description: "Displays a list of all commands",
			usage:       "help",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the program",
			usage:       "exit",
			callback:    commandExit,
		},
		"user": {
			name:        "user",
			description: "Switch to a new user or describe the current user",
			usage:       "user [list|username]",
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
			description: "Enter data for a day",
			usage:       "daily",
			callback:    commandDaily,
		},
		"goals": {
			name:        "goals",
			description: "List goals and highlight achieved ones",
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

func altCommands() map[string]string {
	return map[string]string{
		"?":     "help",
		"users": "user list",
	}
}
