package main

import "fmt"

type goes int

var a goes

func main() {

	a = 45
	fmt.Println(a)
	fmt.Printf("%T", a)
}
