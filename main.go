package main

import (
	"github.com/koderkt/pokedexcli/internal/pokeapi"
)



type Config struct {
	client   pokeapi.Client
	next     string
	previous string
}

func main() {
	repl(&Config{
		client: pokeapi.NewClient(),
		
	})
}
