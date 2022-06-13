package main

import "fmt"

var x int

type cat int

var y cat

func main() {
	fmt.Println()
	x = 10
	y = 20
	fmt.Printf("%#v\n", x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%#v\n", y)
	fmt.Printf("%T\n", y)

}
