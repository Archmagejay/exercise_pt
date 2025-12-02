package main

import (
	"fmt"
	"strings"

	//"time"

)



func startRepl(s *state) {


	for {
		// if programState.cfg.LastOpened.Before(time.Now().AddDate(0,0,-1)) {
		//
		// 	fmt.Print("It looks like you havn't entered data for today")
		// }
		fmt.Print("Exercise > ")
		s.in.Scan()
		words := cleanInput(s.in.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		alt, exists := altCommands()[commandName]
		if exists {
			commandName = alt
		}
		command, exists := getCommands()[commandName]
		if exists {
			if err := command.callback(s, args...); err != nil {
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
