package main

import (
	"fmt"
)

var x int

type cat_own int

var y cat_own

func main() {
	x = 10
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	y = 34
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
