package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type Person struct {
		First      string
		Last       string
		Uniform    string
		ShoesColor string
		Height     float64
	}
	P1 := Person{
		First:      "Anylen ",
		Last:       "Seb",
		Uniform:    "Red hot sailor",
		ShoesColor: "Black",
		Height:     6.5,
	}
	P2 := Person{
		First:      "Angela Duck Forth ",
		Last:       "Seb",
		Uniform:    "Black hot sailor under short",
		ShoesColor: "White",
		Height:     6.6,
	}
	people := []Person{P1, P2}
	err := json.NewEncoder(os.Stdout).Encode(people)
	if err != nil {
		fmt.Println(err)
	}

}
