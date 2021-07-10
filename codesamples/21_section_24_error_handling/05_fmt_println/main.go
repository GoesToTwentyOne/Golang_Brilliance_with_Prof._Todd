package main

import (
	"fmt"
	"os"
)

func main() {
	defer goes()

	_, err := os.Open("nofileexist.txt")
	if err != nil {
		fmt.Println(err)

	}
}
func goes() {
	fmt.Println("defet  execution check")
}
