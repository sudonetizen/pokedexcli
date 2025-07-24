package main 

import (
    "os"
    "fmt"
    "bufio"
    "strings"
)


type cliCommand struct {
    name string 
    description string
    callback func(args ...string) error
}


var pokedex_commands = map[string]cliCommand{
    "exit": {
        name: "exit", 
        description: "Exit the Pokedex", 
        callback: commandExit,
    },
    "help": {
        name: "help",
        description: "Displays a help message",
        callback: commandHelp,
    },
    "map": {
        name: "map",
        description: "Lists location areas",
        callback: commandMap,
    },
    "mapb": {
        name: "mapb",
        description: "Lists locations areas in reverse",
        callback: commandMapb,
    },
    "explore": {
        name: "explore",
        description: "Information about the location area",
        callback: commandExplore,
    },
    "catch": {
        name: "catch",
        description: "Catch Pokemon",
        callback: commandCatch,
    },
    "inspect": {
        name: "inspect",
        description: "Inspect available pokemon",
        callback: commandInspect,
    },
    "pokedex": {
        name: "pokedex",
        description: "Lists caught pokemons",
        callback: commandPokedex,
    },
}


func cleanInput(text string) []string {
    result := strings.Fields(strings.ToLower(text))
    return result
}


func loopRepl() {
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("Pokedex > ")

        scanner.Scan()
        inputStrings := cleanInput(scanner.Text())
        if len(inputStrings) == 0 {continue}
    
        value, boolean := pokedex_commands[inputStrings[0]]

        if boolean {
            if len(inputStrings) == 1 {
                value.callback()
            } else if len(inputStrings) == 2 {
                value.callback(inputStrings[1])
            }
        } else {fmt.Println("Unknown command")}

    }
}
