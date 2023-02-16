package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
}
type secretAgent struct {
	person
	ltk bool
}

func main() {
	p := []secretAgent{
		{
			person: person{
				First: "Md. Nihad",
				Last:  "Hossain",
			},
			ltk: true,
		},
		{
			person: person{
				First: "Alex",
				Last:  "Goot",
			},
			ltk: false,
		},
	}
	fmt.Println(p)
	bs, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
