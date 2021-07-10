package main

import "fmt"

func main() {
	f := goesAdd(1, 2, 3, 4, 5, 56, 6)
	fmt.Println(f)

}
func goesAdd(n ...int) int {
	sum := 0
	for v := range n {
		sum += v

	}
	return sum
}
