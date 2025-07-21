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
    callback func() error
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
           value.callback()
        } else {
            fmt.Println("Unknown command")
        }

    }
}
