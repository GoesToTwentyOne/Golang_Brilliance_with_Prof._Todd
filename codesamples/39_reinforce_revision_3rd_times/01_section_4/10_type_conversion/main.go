package main

import "fmt"

var x int

type cat float32

var y cat

func main() {
	x = 10
	y = 32.1
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%T\n", y)
	x = int(y)
	fmt.Println(x)

}
