package main

import "fmt"

func main() {
	a := 42
	fmt.Printf("%d \n", a)
	fmt.Printf("%b \n", a)
	fmt.Printf("%#x \n", a)
	fmt.Printf("%d \t %b \t %#x", a, a, a)
}
