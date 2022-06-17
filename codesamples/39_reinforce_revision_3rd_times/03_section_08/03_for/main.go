package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("the outer loop : %v\n", i)
		for k := 0; k < 3; k++ {
			fmt.Printf("\t\t inner loop : %v\n", k)

		}
	}
}
