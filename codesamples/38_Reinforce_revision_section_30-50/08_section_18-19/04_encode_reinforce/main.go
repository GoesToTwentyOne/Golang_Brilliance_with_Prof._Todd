package main

import (
	"encoding/json"
	"fmt"
	"os"
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
	// fmt.Println(per)

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(per); err != nil {
		fmt.Println(err)
	}
	// fmt.Println(per)

}
