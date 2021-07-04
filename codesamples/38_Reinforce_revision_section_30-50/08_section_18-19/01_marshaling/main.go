package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Person struct {
		Name      string
		Age       int
		Dev_tools string
	}
	p1 := Person{
		Name:      "Nihad Hossain",
		Age:       21,
		Dev_tools: "Golang grpc",
	}
	p2 := Person{
		Name:      "Seniya  Hossain",
		Age:       20,
		Dev_tools: "Golang grpc-2",
	}

	per := []Person{p1, p2}
	bs, err := json.Marshal(per)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
