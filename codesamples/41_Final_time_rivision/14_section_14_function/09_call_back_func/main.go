package main

import (
	"fmt"
)

func main() {
	xi := []int{1, 2, 3, 4, 5}
	s := sum(xi...)
	fmt.Println(s)
	x := even(sum, xi...)
	fmt.Println(x)

}

func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum

}
func even(f func(x ...int) int, x ...int) int {
	var xi []int
	for _, v := range x {
		if v%2 == 0 {
			xi = append(xi, v)
		}
	}
	return f(xi...)

}
