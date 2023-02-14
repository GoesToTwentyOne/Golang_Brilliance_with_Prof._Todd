package main

import (
	"fmt"
)

func main() {
	for i := 65; i <= 122; i++ {
		fmt.Printf("%v", i)
		for i := 0; i < 3; i++ {
			fmt.Printf("\t\t%#U \n", i)
		}
	}
}
