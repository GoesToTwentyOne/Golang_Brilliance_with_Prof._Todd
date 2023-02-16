package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
}

func main() {
	s := `[{"First":"Md. Nihad","Last":"Hossain"},{"First":"Alex","Last":"Goot"}]`
	bs := []byte(s)
	var p []person
	err := json.Unmarshal(bs, &p)
	if err != nil {
		fmt.Println(err)
	}
	for i, v := range p {
		fmt.Println(i, "      value   ", v)
	}

}
