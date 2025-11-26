package main


type cliCommand struct {
	name 			string
	description 	string
	callback func(*state, ...string) error
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
		"user": {
			name: "user",
			description: "Switch users",
			callback: commandUser,
		},
		"register": {
			name: "register",
			description: "Register a new user",
			callback: commandRegister,
		},
		"graph": {
			name: "graph",
			description: "Graph available data",
			callback: commandGraph,
		},
		"daily": {
			name: "daily",
			description: "Enter data for a day",
			callback: commandDaily,
		},
		"goals": {
			name: "goals",
			description: "List goals and highlight achieved ones",
			callback: commandGoals,
		},
		"export": {
			name: "export",
			description: "Export the database",
			callback: commandExport,
		},
		"import": {
			name: "import",
			description: "Import a database",
			callback: commandImport,
		},
		"change": {
			name: "change",
			description: "Change a saved value",
			callback: commandChange,
		},
		"clear": {
			name: "clear",
			description: "Clear the screen",
			callback: commandClear,
		},
	}
}

func altCommands() map[string]string {
	return map[string]string {
		"?": "help",
	}
}