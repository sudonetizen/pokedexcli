package main

import "testing"


func TestCleanInput(t *testing.T) {
    cases := []struct{
        input string 
        expected []string
    }{
        {input: "  hello world  ", expected: []string{"hello", "world"},},
        {input: "welcome to the earth    ", expected: []string{"welcome", "to", "the", "earth"},},
        {input: "  Human From   Mars", expected: []string{"human", "from", "mars"},},
    }

    for _, c := range cases {
        actual := cleanInput(c.input)
        if len(actual) != len(c.expected) {
            t.Errorf("expected length: %v, actual length: %v", len(c.expected), len(actual))
        } 

        for i := range actual {
            word := actual[i]
            expectedWord := c.expected[i]
            if word != expectedWord { 
                t.Errorf("expected word: %v, actual word: %v", expectedWord, word)
            }
        }

    }


}
