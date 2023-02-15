package main

import (
	"fmt"
)

func unfluringFinal(s string, x ...int) {
	fmt.Println(x)
	fmt.Println(s)
}
func main() {
	xi := []int{1, 2, 4, 5, 5, 6, 6, 6}
	unfluringFinal("Nihad", xi...)
}
