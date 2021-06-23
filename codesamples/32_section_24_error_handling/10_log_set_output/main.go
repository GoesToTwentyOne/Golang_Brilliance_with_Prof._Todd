package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("error.txt")
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(f)
	defer f.Close()

	f2, err := os.Open("no_exist_file.txt")
	if err != nil {
		log.Println(err)
	}
	defer f2.Close()
	fmt.Println("error checking with error file")

}
