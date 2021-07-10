package main

import "fmt"

func main() {
	xn := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	final := even(sum, xn)
	fmt.Println(final)

}
func sum(n ...int) int {
	total := 0
	for _, value := range n {
		total += value
	}
	return total

}
func even(f func(n ...int) int, n []int) int {
	var xi []int
	for _, value := range n {
		if value%2 == 0 {
			xi = append(xi, value)
		}
	}
	return f(xi...)

}
