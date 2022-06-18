package main

import "fmt"

func main() {
	x := []int{4, 6, 54, 6, 7, 3, 3, 54, 6, 7}
	x = append(x[:4], x[7:]...)
	fmt.Println(x)
}
