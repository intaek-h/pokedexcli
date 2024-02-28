package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getAllCommands()[commandName]

		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Command not found")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getAllCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "List all available commands",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the program",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "List all available maps",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous maps",
			callback:    commandMapb,
		},
	}
}
