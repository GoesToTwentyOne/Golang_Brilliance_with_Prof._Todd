package main

import "fmt"

var a int
var b float64

func main() {
	x := 20
	y := 20.45
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%T\n", y)
	a = 10
	b = 19.44
	fmt.Printf("%v\n", a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%T\n", b)

}
