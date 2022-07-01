package main

import (
	"fmt"
	"math"
)

type SHAPE interface {
	AREA() float64
}
type SQUARE struct {
	height float64
	width  float64
}
type CIRCLE struct {
	radius float64
}

func (s SQUARE) AREA() float64 {
	return (s.height * s.width)
}
func (c CIRCLE) AREA() float64 {
	return (math.Pi * c.radius * c.radius)
}

func INFO(s SHAPE) {
	fmt.Println(s.AREA())
}
func main() {
	sq := SQUARE{
		height: 4.3,
		width:  2.6,
	}
	cr := CIRCLE{
		radius: 54.3,
	}
	INFO(sq)
	INFO(cr)

}
