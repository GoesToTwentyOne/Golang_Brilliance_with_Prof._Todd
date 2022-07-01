package main

import "fmt"

func main() {
	fmt.Println(sum()([]int{4, 5, 6}...))

}
func sum() func(x ...int) int {
	return func(x ...int) int {
		sum := 0
		for _, v := range x {
			sum += v
		}
		return sum

	}
}
