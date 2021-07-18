package main

import "fmt"

func main() {
	sum := 0
	x := []int{1, 2, 3, 4, 5, 6, 7, 8}
	func(x []int) {
		for _, v := range x {
			sum += v
		}
		fmt.Println("sum of number :", sum)

	}(x)

}
