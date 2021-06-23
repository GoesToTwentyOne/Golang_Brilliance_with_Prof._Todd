package main

import (
	"fmt"
	"os"
)

func main() {
	defer goes()

	_, err := os.Open("nofileexist.txt")
	if err != nil {
		panic(err)

	}
}
func goes() {
	fmt.Println("defet  execution check")
}
