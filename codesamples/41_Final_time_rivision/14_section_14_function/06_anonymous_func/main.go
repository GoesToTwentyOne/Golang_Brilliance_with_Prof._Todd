package main

import "fmt"

func main() {
	func() {
		fmt.Println("I'm anonymous func")
	}()
	func(x ...int) {
		sum := 0
		for _, v := range x {
			sum += v

		}
		fmt.Println(sum)

	}([]int{1, 2, 3, 4, 5, 6, 7, 8}...)

}
