package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(x)
	x = append(x, 7, 8, 9)
	fmt.Println(x)
	y := []int{65, 66, 777, 777888, 99999, 444, 5655}
	fmt.Println(y)
	x = append(x, y...)
	fmt.Println(x)
}
