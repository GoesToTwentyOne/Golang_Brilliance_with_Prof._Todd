package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	jsonString := []byte(`[{"First":"Tood","Last":"Mcleod","Age":32},{"First":"Angela","Last":"Duckworth","Age":25}]`)

	var people []Person

	//unmarshaling
	err := json.Unmarshal(jsonString, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)
	fmt.Printf("%+v \n \n", people)
	for i, value := range people {
		fmt.Println("record no ", i)
		fmt.Printf("\t \t %v  %v  %v\n", value.First, value.Last, value.Age)
	}
}
