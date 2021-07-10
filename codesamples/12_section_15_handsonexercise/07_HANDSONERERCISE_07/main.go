package main

import "fmt"

func main() {
	a := func(n ...int) {

		sum := 0
		for _, value := range n {
			sum += value

		}
		fmt.Println(sum)

	}
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a(slice1...)
}
