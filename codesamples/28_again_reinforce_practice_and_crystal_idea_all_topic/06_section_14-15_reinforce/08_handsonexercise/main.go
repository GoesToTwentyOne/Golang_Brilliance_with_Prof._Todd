package main

import "fmt"

func main() {
	a := func(x []int) func() int {
		return func() int {
			sum := 0
			for _, v := range x {
				if v%2 == 0 {
					sum += v
				}
			}
			return sum

		}

	}
	n := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("sum of all even ", a(n)())
}
