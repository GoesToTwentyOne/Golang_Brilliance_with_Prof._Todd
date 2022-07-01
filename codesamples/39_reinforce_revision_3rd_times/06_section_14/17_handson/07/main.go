package main

import "fmt"

func main() {

	f := func(x ...int) int {
		sum := 0
		for _, v := range x {
			sum += v
		}
		return sum

	}
	fmt.Println(f([]int{123, 54, 78}...))
}
