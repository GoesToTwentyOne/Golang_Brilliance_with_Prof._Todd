package main

import (
	"fmt"
)

func main() {
	if x := 21; x == 22 {
		fmt.Println("Is 22")
	} else if x := 23; x == 24 {
		fmt.Println("Is 24")
	} else {
		fmt.Println("isn't 21 or 24")
	}
}
