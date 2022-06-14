package main

import "fmt"

const (
	a = iota
	b
	c
)
const (
	x = iota
	y = iota
	z = iota
)
const (
	m = iota
	n
	o
	p
	q
	r
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(o)
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)

}
