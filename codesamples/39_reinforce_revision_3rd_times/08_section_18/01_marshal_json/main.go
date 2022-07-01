package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	First string
	Last  string
}
type Student struct {
	Person
	Age   int
	Class string
}

func main() {
	s := Student{
		Person: Person{
			First: "Nihad",
			Last:  "Hossain",
		},
		Age:   21,
		Class: "14",
	}
	bs, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}
	os.Stdout.Write(bs)

}
