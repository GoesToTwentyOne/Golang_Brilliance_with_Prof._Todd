package main

import (
	"fmt"
)

var x int

type cat_own int

var y cat_own

func main() {
	x = 10
	y = 20
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	x = int(y)
	fmt.Println(y)
	fmt.Printf("%T\n", x)

}
