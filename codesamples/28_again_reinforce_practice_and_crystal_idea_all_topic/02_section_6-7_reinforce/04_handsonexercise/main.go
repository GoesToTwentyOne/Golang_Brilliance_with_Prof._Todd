package main

import "fmt"

func main() {
	x := 2
	fmt.Printf("%d \t %b \t %#x \n", x, x, x)
	x = 2 << 1
	fmt.Printf("%d \t %b \t %#x", x, x, x)
}
