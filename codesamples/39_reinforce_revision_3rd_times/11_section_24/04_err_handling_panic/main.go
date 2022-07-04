package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("name.txt")
	if err != nil {
		//log.Fatal("Error :", err)
		fmt.Println("Error :", err)
	}
	defer f.Close()
	log.SetOutput(f)
	f2, err := os.Open("hmn.txt")
	if err != nil {
		log.Fatal("Error :", err)
	}
	defer f2.Close()

}
