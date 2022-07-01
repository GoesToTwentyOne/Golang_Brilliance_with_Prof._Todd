package main

import "fmt"

func main() {

	sum := func(f func(x ...int) int, y ...int) int {
		var xi []int
		for _, v := range y {
			if v%2 == 0 {
				xi = append(xi, v)
			}
		}
		return f(xi...)
	}(func(x ...int) int {
		sum := 0
		for _, v := range x {
			sum += v
		}

		return sum
	}, []int{1, 2, 3, 4}...)
	fmt.Println(sum)

}
