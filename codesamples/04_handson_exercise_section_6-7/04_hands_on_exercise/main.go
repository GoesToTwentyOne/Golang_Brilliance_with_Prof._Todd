package main

import "fmt"

func main() {
	a := 10
	b := a << 1
	fmt.Printf("%d \t %b \t %#x \n", a, a, a)
	fmt.Printf("%d \t %b \t %#x \n", b, b, b)
}
