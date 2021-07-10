package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(callback(sum, ii))

}
func sum(xi []int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum

}
func callback(c func(xi []int) int, xic []int) int {
	xn := []int{}
	for _, v := range xic {
		if v%2 == 0 {
			xn = append(xn, v)
		}
	}
	return sum(xn)
}
