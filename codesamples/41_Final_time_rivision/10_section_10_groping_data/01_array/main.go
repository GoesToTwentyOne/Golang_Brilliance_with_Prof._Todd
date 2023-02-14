package main

import (
	"fmt"
)

func main() {
	//array size is a part of type
	var x [5]int
	fmt.Println(x)
	x[2] = 23
	fmt.Println(x)
}
