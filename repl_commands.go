package main

import (
	"fmt"
	"os"

	"github.com/intaek-h/pokedexcli/pokeapi"
)

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getAllCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandExit() error {
	fmt.Println("Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap() error {
	fmt.Println("Getting maps...")
	pokeapi.GetMaps()
	return nil
}
