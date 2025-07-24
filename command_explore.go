package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)

type Explore struct {
	EncounterMethodRates []EncounterMethodRates `json:"encounter_method_rates"`
	GameIndex            int                    `json:"game_index"`
	ID                   int                    `json:"id"`
	Location             Location               `json:"location"`
	Name                 string                 `json:"name"`
	Names                []Names                `json:"names"`
	PokemonEncounters    []PokemonEncounters    `json:"pokemon_encounters"`
}
type EncounterMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Version struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type VersionDetails struct {
	Rate    int     `json:"rate"`
	Version Version `json:"version"`
}
type EncounterMethodRates struct {
	EncounterMethod EncounterMethod  `json:"encounter_method"`
	VersionDetails  []VersionDetails `json:"version_details"`
}
type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Language struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Names struct {
	Language Language `json:"language"`
	Name     string   `json:"name"`
}
type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Method struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type EncounterDetails struct {
	Chance          int    `json:"chance"`
	ConditionValues []any  `json:"condition_values"`
	MaxLevel        int    `json:"max_level"`
	Method          Method `json:"method"`
	MinLevel        int    `json:"min_level"`
}
type VersionDetailss struct {
	EncounterDetails []EncounterDetails `json:"encounter_details"`
	MaxChance        int                `json:"max_chance"`
	Version          Version            `json:"version"`
}
type PokemonEncounters struct {
	Pokemon        Pokemon          `json:"pokemon"`
	VersionDetails []VersionDetails `json:"version_details"`
}

////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////

func commandExplore(args ...string) error {
    if len(args) == 0 {
        fmt.Println("provide location name")
    } else {
        fmt.Printf("Exploring %v...\n", args[0] )
        fmt.Println("Found Pokemon:")
        url := "https://pokeapi.co/api/v2/location-area/" + args[0]     
        loc_explore(url)
    }
    return nil
}

func loc_explore(url string) error {
    res, err := http.Get(url)
    if err != nil {return err}
    if res.StatusCode > 299 {fmt.Printf("error: %v\n", res.Status)}
    defer res.Body.Close()

    exploreData := Explore{} 
    decoder := json.NewDecoder(res.Body)
    err = decoder.Decode(&exploreData)
    if err != nil {return err}
    
    for _, x := range exploreData.PokemonEncounters {
        fmt.Println("-", x.Pokemon.Name)
    }    

    return nil
}
