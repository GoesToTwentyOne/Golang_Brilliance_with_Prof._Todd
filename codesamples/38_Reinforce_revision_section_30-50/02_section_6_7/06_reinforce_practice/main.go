package main

import "fmt"

const (
	_io = iota
	a   = 2021 + iota
	b
	c
	d
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}
