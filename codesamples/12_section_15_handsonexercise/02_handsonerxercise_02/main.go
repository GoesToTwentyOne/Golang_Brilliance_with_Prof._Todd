package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(foo(xi...))
	fmt.Println(bar(xi))

}
func foo(n ...int) int {
	sum := 0
	for _, value := range n {
		sum += value
	}
	return sum

}
func bar(n []int) int {
	sum := 0
	for _, value := range n {
		sum += value
	}
	return sum

}
