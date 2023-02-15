package main

import (
	"fmt"
)

type squre struct {
	height int
	width  int
}
type circle struct {
	radius float64
}
type shape interface {
	area()
}

func (s squre) area() {
	fmt.Println(s.height * s.width)
}
func (c circle) area() {
	pie := 3.1416
	fmt.Println(pie * c.radius * c.radius)
}
func info(s shape) {
	switch s.(type) {
	case squre:
		fmt.Println("I'm pass by info")
		s.(squre).area()

		break
	case circle:
		fmt.Println("I'm pass by info")
		s.(circle).area()
		break
	default:
		fmt.Print("I can't found you")
	}
}
func main() {
	s := squre{
		height: 40,
		width:  50,
	}
	s2 := squre{
		height: 15,
		width:  20,
	}
	c := circle{
		radius: 45.5,
	}
	info(s)
	info(s2)
	info(c)

}
