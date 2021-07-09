package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("Neha.txt")
	if err != nil {
		log.Fatal("print err", err)
	}
	log.SetOutput(f)
	defer f.Close()
	f2, err := os.Open("ore.txt")
	if err != nil {
		log.Fatal(err)
	}
	f2.Close()
}
