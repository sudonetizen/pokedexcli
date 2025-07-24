package main

import (
    "fmt"
)

func commandInspect(args ...string) error {
    if len(args) == 0 {
        fmt.Println("provide pokemon name")
    } else if len(args) == 1 {
        poke, okk := myPokeMap[args[0]]
        if okk {
            fmt.Println("Name:", poke.Name)
            fmt.Println("Height:", poke.Height)
            fmt.Println("Weight:", poke.Weight)
            fmt.Println("Stats:")
            for _, x := range poke.Stats {
                fmt.Printf("-%v: %v\n", x.Stat.Name, x.BaseStat)
            }
            fmt.Println("Types:")
            for _, x := range poke.Types {
                fmt.Println("-", x.Type.Name)
            }
        } else {fmt.Println("not found")}
    }

    return nil
}
