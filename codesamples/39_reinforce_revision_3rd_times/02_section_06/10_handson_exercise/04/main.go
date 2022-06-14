package main

import "fmt"

func main() {
	a := 44
	fmt.Printf("%d \t %b %x\n", a, a, a)
	b := a << 1
	fmt.Printf("%d \t %b %x\n", b, b, b)
}
