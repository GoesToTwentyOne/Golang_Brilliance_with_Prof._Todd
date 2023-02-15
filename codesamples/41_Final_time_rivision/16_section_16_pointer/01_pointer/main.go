package main

import (
	"fmt"
)

func main() {
	var a int = 23
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(&a)
	fmt.Printf("%T\n", &a)
	b := &a
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(*&a)
	*b = 25
	fmt.Println(a)
}
