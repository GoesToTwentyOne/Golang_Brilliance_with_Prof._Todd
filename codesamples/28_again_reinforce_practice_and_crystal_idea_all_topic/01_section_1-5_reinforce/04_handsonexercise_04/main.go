package main

import "fmt"

func main() {
	type grit int
	var x grit
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
