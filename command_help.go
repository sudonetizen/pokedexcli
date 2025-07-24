package main

import "fmt"

func commandHelp(args ...string) error {
    text :=`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
map: Lists location areas
mapb: Lists location areas in reverse
explore: Information about the location area
catch: Catch Pokemon
inspect: Inspect available pokemon
pokedex: Lists caught pokemons`
    fmt.Println(text)
    return nil
}
