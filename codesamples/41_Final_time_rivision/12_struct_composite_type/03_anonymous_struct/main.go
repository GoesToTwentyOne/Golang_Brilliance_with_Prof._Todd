package main

import (
	"fmt"
)

func main() {
	p := struct {
		fisrt string
		last  string
		age   int
	}{
		fisrt: "Nihad",
		last:  "Hossain",
		age:   22,
	}
	fmt.Println(p)
}
