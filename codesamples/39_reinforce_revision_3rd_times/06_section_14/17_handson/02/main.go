package main

import "fmt"

func main() {

	fmt.Println(foo([]int{1, 2, 3, 4, 5, 6, 7}...))
	fmt.Println(bar([]int{1, 2, 3, 4, 5, 6, 7}))
}
func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
func bar(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
