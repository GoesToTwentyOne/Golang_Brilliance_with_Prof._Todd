package main

import (
	"encoding/json"
	"fmt"
)

var jsms = []byte(`
	{"First":"Nihad","Last":"Hossain","Age":21,"Class":"14"}
`)

type Perosn struct {
	First string
	Last  string
	Age   int
	Class string
}

var Persons Perosn

func main() {

	err := json.Unmarshal(jsms, &Persons)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", Persons)
}
