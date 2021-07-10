package main

import "fmt"

type goesforward int

var x goesforward
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T \n", x)
	x = 42
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T \n", y)
}
