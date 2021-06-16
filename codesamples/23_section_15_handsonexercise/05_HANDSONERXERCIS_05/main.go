package main

import (
	"fmt"
	"math"
)

type SQUARE struct {
	length float64
	width  float64
}
type CIRCLE struct {
	radious float64
}

func (S SQUARE) AREA() float64 {
	return S.length * S.width

}
func (C CIRCLE) AREA() float64 {
	return math.Pi * math.Pow(C.radious, 2)
}

type SHAPE interface {
	AREA() float64
}

func INFO(S SHAPE) {
	fmt.Println(S.AREA())
	// fmt.Println(S.(SQUARE).AREA())

}

func main() {
	s := SQUARE{
		length: 45,
		width:  45,
	}
	c := CIRCLE{
		radious: 50,
	}

	INFO(c)
	INFO(s)

}
