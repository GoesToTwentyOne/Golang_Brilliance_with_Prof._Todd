package main

import "fmt"

func main() {
	ccc := callback(evsum, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println(ccc)

}
func callback(f func(x []int) int, xi []int) int {
	xii := []int{}
	for _, v := range xi {
		if v%2 == 0 {
			xii = append(xii, v)

		}
	}
	return evsum(xii)
}
func evsum(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum

}
