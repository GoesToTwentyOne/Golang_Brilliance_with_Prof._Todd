package main

import "fmt"

type grit int

var x grit
var y int

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
