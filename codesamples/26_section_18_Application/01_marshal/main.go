package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type Person struct {
		First    string
		Last     string
		ColorFav []string
	}
	p := Person{
		First:    "Nihad",
		Last:     "Hossain",
		ColorFav: []string{"Red", "Blue", "Black", "Purple"},
	}
	bs, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf("%+v \n", p)
	// os.Stderr.Write()
	fmt.Println(string(bs))
	os.Stdout.Write(bs)
}
