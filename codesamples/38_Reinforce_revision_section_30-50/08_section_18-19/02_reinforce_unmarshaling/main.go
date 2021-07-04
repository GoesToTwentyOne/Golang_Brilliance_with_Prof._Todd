package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	bs := []byte(`[{"Name":"Nihad Hossain","Age":21,"Dev_tools":"Golang grpc"},{"Name":"Seniya  Hossain","Age":20,"Dev_tools":"Golang grpc-2"}]`)
	type Person struct {
		Name      string `json:"Name"`
		Age       int    `json:"Age"`
		Dev_tools string `json:"Dev_tools"`
	}

	var unmar []Person
	err := json.Unmarshal(bs, &unmar)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", unmar)
}
