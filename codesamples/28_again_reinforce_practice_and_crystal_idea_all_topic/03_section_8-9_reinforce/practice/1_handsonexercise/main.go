package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x)
	x[2] = 43
	fmt.Println(x)

	y := make([]int, 5, 10)
	fmt.Println(y)
	y[1] = 55
	fmt.Println(y)

	z := []int{4, 5, 6, 7, 8, 9} // immuteable
	a := z
	a[0] = 88888
	a = append(a, 78)
	fmt.Println(a)
	fmt.Println(z)

}
