package main

import "fmt"

type rect struct {
	height int
	width  int
}

func (r *rect) area() int {
	return r.height * r.width

}
func (r rect) preim() int {
	return 2*r.height + 2*r.width
}

func main() {
	r := rect{height: 32, width: 43}
	fmt.Println(r.area())
	fmt.Println(r.preim())
	r1 := rect{height: 42, width: 85}
	rp := &r1
	fmt.Println(rp.area())
	fmt.Println(rp.preim())

}
