package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer goes()

	_, err := os.Open("nofileexist.txt")
	if err != nil {
		log.Println(err)

	}
}
func goes() {
	fmt.Println("defet  execution check")
}
