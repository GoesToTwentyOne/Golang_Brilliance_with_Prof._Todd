package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First     string `json:"First"`
	Age       int    `json:"Age"`
	NativeLan string `json:"Native_lan"`
}

func main() {
	bs := []byte(`[{"First":"Nihad","Age":21,"Native_lan":"English"},{"First":"Alex","Age":23,"Native_lan":"English"},{"First":"Jaconina","Age":20,"Native_lan":"German"}]`)
	// people := []Person{}
	var people []Person
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	for i, v := range people {
		fmt.Println("Index :", i)
		fmt.Println("value: ", v)
	}

}
