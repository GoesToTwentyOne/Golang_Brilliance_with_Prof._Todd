package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer goes()
	f, err := os.Open("nn.txt")
	if err != nil {
		log.Fatal("Error :", err)
	}
	f.Close()
}
func goes() {
	fmt.Println("goes")
}
