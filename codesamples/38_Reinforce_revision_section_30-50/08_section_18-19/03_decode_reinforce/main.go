package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	bs := `[{"Name":"Nihad Hossain","Age":21,"Dev_tools":"Golang grpc"},{"Name":"Seniya Hossain","Age":20,"Dev_tools":"Golang grpc-2"}]`
	type Person struct {
		Name      string `json:"Name"`
		Age       int    `json:"Age"`
		Dev_tools string `json:"Dev_tools"`
	}

	var per []Person
	dec := json.NewDecoder(strings.NewReader(bs))
	err := dec.Decode(&per)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(per)

}
