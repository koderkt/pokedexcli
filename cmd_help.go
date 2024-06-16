package main

import "fmt"

func cmdHelp(config *Config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, val := range getCommands() {
		fmt.Printf("%s : %s \n", val.name, val.description)
	}
	
	return nil
}
