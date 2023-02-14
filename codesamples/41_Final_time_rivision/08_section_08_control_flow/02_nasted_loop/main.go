package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Outer loop ", i)
		for j := 0; j < 5; j++ {
			fmt.Println("\t \tInner loop ", j)
		}
	}
}
