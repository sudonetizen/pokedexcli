package main

import (
    "fmt"
)

func commandPokedex(args ...string) error {
    fmt.Println("Your Pokedex:")
    for _, poke := range myPokeMap {
        fmt.Println("-", poke.Name)
    }
    return nil
}
