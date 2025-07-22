package main

import "fmt"

func commandHelp() error {
    text :=`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
map: Lists location areas
mapb: Lists location areas in reverse`
    fmt.Println(text)
    return nil
}
