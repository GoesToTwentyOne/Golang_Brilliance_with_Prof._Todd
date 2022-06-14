package main

import "fmt"

var x = 20
var y = 45.5855
var z = "Nihad"

const (
	a = 10
	b = 34.444
	c = "Hossain"
)
const (
	n int     = 55555555555
	m float32 = 56.555
	o string  = "Golang is first choice"
)

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%v\n", m)
	fmt.Printf("%v\n", n)
	fmt.Printf("%v\n", o)
	fmt.Printf("%T\n", m)
	fmt.Printf("%T\n", n)
	fmt.Printf("%T\n", o)

}
