package main

import "fmt"

func main() {
	birthdate := 1999
	for {
		if birthdate > 2021 {
			break
		}
		fmt.Println("alive ", birthdate)
		birthdate++

	}

}
