package main 

import (
    "fmt"
    "strconv"
    "net/http"
    "encoding/json"
    "github.com/sudonetizen/pokecache"
)

var val_num int = 0
var offset *int = &val_num

var page int = 0
var pager *int = &page

type poke struct {
    Name string `json:"name"`
    Url  string `json:"url"`
}

type pokeList struct {
    Count    int    `json:"count"`
    Next     string `json:"next"`
    Previous any    `json:"previous"`
    Results  []poke `json:"results"` 
}

var loCache *pokecache.Cache = pokecache.NewCache(pokecache.Interval)

func commandMap() error {

    url := "https://pokeapi.co/api/v2/location-area/?offset=" + strconv.Itoa(*offset) + "&limit=20"

    cachedData, cachedB := loCache.Get(url)
    if cachedB {
        *offset += 20
        *pager += 1
        fmt.Println(string(cachedData))
        fmt.Printf("you are on the page %v\n", *pager)
        fmt.Printf("cached: %v\n\n", cachedB)
        return nil 
    }

    res, err := http.Get(url)
    if err != nil {return err}
    if res.StatusCode > 299 {return fmt.Errorf("failed")}
    defer res.Body.Close()
    
    pokeMap := pokeList{}
    decoder := json.NewDecoder(res.Body)
    err = decoder.Decode(&pokeMap)
    if err != nil {return err}

    *offset += 20
    *pager += 1
    locations := ""
    for _, i := range pokeMap.Results {
        locations += fmt.Sprintln(i.Name)  
    }
    fmt.Println(locations)
    fmt.Printf("you are on the page %v\n", *pager)
    fmt.Printf("cached: %v\n\n", cachedB)

    loCache.Add(url, []byte(locations))

    return nil
}

func commandMapb() error {
    if *offset >= 40 {*offset -= 40}
    if *pager >= 2 {*pager -= 2}
    commandMap()    
    return nil
}
