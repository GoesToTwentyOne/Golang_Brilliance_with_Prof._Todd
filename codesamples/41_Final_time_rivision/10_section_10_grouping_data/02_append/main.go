package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8}
	y := []int{44, 55, 66, 77, 88}
	x = append(x, y...)
	fmt.Println(x)
}
