package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonBlob = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	]`)
	type Animal struct {
		Name  string
		Order string
	}
	var unmarshal []Animal
	err := json.Unmarshal(jsonBlob, &unmarshal)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(unmarshal)
	fmt.Printf("%+v", unmarshal)
}
