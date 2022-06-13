package main

import "fmt"

var x bool

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)
	//predeclare define type  practice
	a := 10
	b := 20
	fmt.Println(a == b)
	fmt.Println(a <= b)
	fmt.Println(a >= b)
	fmt.Println(a != b)
}
