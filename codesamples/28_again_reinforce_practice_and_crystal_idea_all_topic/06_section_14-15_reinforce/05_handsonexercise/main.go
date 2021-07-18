package main

import (
	"fmt"
	"math"
)

type SQUARE struct {
	L float64
	W float64
}
type CIRCLE struct {
	r float64
}

func (s SQUARE) AREA() {
	fmt.Println("Squre Area ", s.L*s.W)
}
func (c CIRCLE) AREA() {
	fmt.Println("Squre Area ", math.Pi*c.r*c.r)
}

type SHAPE interface {
	AREA()
}

func INFO(s SHAPE) {
	s.AREA()

}
func main() {
	s := SQUARE{
		L: 4,
		W: 7,
	}
	c := CIRCLE{
		r: 45.56,
	}
	INFO(s)
	INFO(c)

}
