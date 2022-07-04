package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("alex.txt")
	if err != nil {
		fmt.Println("Happen Error : ", err)
		log.Println("Happen Error : ", err)
		log.Fatalln("Happen Error : ", err)
	}
	defer f.Close()
}
