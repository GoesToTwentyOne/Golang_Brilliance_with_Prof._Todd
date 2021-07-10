package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := foo(xi...)
	fmt.Println(s)

}
func foo(v ...int) int {
	sum := 0
	for _, v := range v {
		sum += v
	}
	return sum

}
