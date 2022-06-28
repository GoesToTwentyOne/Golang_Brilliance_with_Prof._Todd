package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	priem() float64
}

type rect struct {
	height float64
	width  float64
}
type circ struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.width

}
func (r rect) priem() float64 {
	return 2*r.height + 2*r.width
}
func (c circ) area() float64 {
	return math.Pi * math.Sqrt(c.radius)
}
func (c circ) priem() float64 {
	return 2 * math.Pi * math.Sqrt(c.radius)
}
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.priem())
}
func main() {
	r := rect{height: 4.5, width: 2}
	c := circ{radius: 4.5}
	measure(r)
	measure(c)

}
