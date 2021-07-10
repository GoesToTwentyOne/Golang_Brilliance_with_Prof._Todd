package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Tood",
		Last:  "Mcleod",
		Age:   32,
	}
	p2 := person{
		First: "Angela",
		Last:  "Duckworth",
		Age:   25,
	}
	people := []person{p1, p2}
	fmt.Println(people)
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
