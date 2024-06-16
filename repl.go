package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl(config *Config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		input := reader.Text()
		commandName := cleanUp(input)
		cmd, exists := getCommands()[commandName[0]]
		if exists {
			err := cmd.callback(config)
			if err != nil {
				fmt.Println("Error: ", err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func cleanUp(input string) []string {
	input = strings.ToLower(input)
	return strings.Fields(input)
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    cmdHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    cmdExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    cmdMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of previous 20 location areas in the Pokemon world",
			callback:    cmdMapb,
		},
	}
}
