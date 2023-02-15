package main

import (
	"fmt"
)

func main() {
	x := func(x ...int) {
		sum := 0
		for _, v := range x {
			sum += v
		}
		fmt.Println(sum)
	}
	x([]int{1, 2, 3, 4, 5, 6, 7}...)
}
