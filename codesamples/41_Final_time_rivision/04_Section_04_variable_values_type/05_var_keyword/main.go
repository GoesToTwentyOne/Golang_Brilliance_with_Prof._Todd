package main

import (
	"fmt"
)

var x int

func main() {
	x = 20
	fmt.Println(x)
	foo()
	y := 30
	fmt.Println(y)
}
func foo() {
	fmt.Println(x)
}
