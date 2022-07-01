package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	First      string
	Age        int
	Native_lan string
}

func main() {
	p := Person{
		First:      "Nihad",
		Age:        21,
		Native_lan: "English",
	}
	p1 := Person{
		First:      "Alex",
		Age:        23,
		Native_lan: "English",
	}
	p2 := Person{
		First:      "Jaconina",
		Age:        20,
		Native_lan: "German",
	}
	pa := []Person{p, p1, p2}
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(pa)

}
