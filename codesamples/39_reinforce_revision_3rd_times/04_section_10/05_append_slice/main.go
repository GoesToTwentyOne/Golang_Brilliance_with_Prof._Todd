package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 78}
	fmt.Println(x)
	x = append(x, 435, 66, 77, 90, 87)
	fmt.Println(x)
	y := []int{7777, 9998, 9990, 555}
	x = append(x, y...)
	fmt.Println(x)
}
