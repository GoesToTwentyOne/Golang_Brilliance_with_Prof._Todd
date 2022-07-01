package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	First     string `json:"First"`
	Age       int    `json:"Age"`
	NativeLan string `json:"Native_lan"`
}

func main() {
	var people []Person
	bs := []byte(`[{"First":"Nihad","Age":21,"Native_lan":"English"},{"First":"Alex","Age":23,"Native_lan":"English"},{"First":"Jaconina","Age":20,"Native_lan":"German"}]`)
	dec := json.NewDecoder(strings.NewReader(string(bs)))
	dec.Decode(&people)

	for _, v := range people {

		fmt.Printf("%+v", v)
	}

}
