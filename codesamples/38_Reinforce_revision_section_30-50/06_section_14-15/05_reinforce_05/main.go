package main

import (
	"fmt"
	"math"
)

type squre struct {
	l float64
	w float64
}
type circle struct {
	r float64
}

func (s squre) area() float64 {
	return s.l * s.w

}
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2)

}

type shape interface {
	area() float64
}

func info(s shape) {
	// fmt.Println("Area", s.area())
	//alternative 1
	// switch s.(type) {
	// case squre:
	// 	fmt.Println("squre area :", s.(squre).area())
	// case circle:
	// 	fmt.Println("circle area :", s.(circle).area())
	// }
	//alternative 2
	// switch s.(type) {
	// case squre:
	// 	fmt.Println("squre area :", s.area())
	// case circle:
	// 	fmt.Println("circle area :", s.area())
	// }
	//alternative 3
	switch sr := s.(type) {
	case squre:
		fmt.Println("squre area :", sr.area())
	case circle:
		fmt.Println("circle area :", sr.area())
	}

}

func main() {
	sq := squre{
		l: 5,
		w: 5,
	}
	fmt.Println("value of squre : ", sq)
	cr := circle{
		r: 4.5,
	}
	fmt.Println("value of circle : ", cr)
	info(sq)
	info(cr)

}
