package main

import "fmt"

var x int
var y = 20

func main() {
	fmt.Println("Package level value of y :", y)
	fmt.Println("Pakcage level zero value of x :", x)
	x = 10
	fmt.Println("Pakcage  value of x :", x)
	y := 10
	fmt.Println("Function level value of y :", y)
	println("Goes function value")
	goes()
}
func goes() {
	fmt.Println(y)
	y = 10
	fmt.Println(y)
}
