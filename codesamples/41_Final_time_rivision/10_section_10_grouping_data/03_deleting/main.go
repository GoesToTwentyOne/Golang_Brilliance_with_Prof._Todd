package main

import (
	"fmt"
)

func main() {
	x := []int{2, 4, 5, 6, 7, 7, 8, 8, 6}
	x = append(x[:3], x[5:]...)
	fmt.Println(x)
}
