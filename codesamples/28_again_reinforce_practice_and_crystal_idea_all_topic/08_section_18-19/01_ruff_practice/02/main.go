package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {

	bs := `[{"First":"Anylen ","Last":"Seb","Uniform":"Red hot sailor","ShoesColor":"Black","Height":6.5},{"First":"Angela Duck Forth ","Last":"Seb","Uniform":"Black hot sailor under short","ShoesColor":"White","Height":6.6}]`

	type Peroson struct {
		First      string  `json:"First"`
		Last       string  `json:"Last"`
		Uniform    string  `json:"Uniform"`
		ShoesColor string  `json:"ShoesColor"`
		Height     float64 `json:"Height"`
	}
	var people []Peroson
	err := json.NewDecoder(strings.NewReader(bs)).Decode(&people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", people)
}
