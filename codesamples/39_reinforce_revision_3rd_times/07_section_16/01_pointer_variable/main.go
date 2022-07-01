package main

import "fmt"

func main() {
	a := 21
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)
	b := &a
	fmt.Println(b)
	fmt.Println(*b)
	*b = 44
	fmt.Println(a)
	fmt.Println(*&a)
}
