package main

import "fmt"

func main() {
	x := 2
	if x == 1 {
		fmt.Println("1")
	} else if x == 2 {
		fmt.Println("2")
	} else if x == 3 {
		fmt.Println("3")
	} else {
		fmt.Println("doesn't match ")
	}
}
