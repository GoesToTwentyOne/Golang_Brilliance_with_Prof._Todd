package main

import "fmt"

func main() {
	x := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}
	fmt.Println(x)
	for i, v := range x {
		fmt.Printf("%v   %v\n", i, v)
	}
}
