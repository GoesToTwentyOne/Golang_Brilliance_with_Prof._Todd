package main

import "fmt"

func main() {
	x := []int{1, 2, 4, 5, 6, 7, 8, 9, 10, 11, 12, 44, 88, 66}
	fmt.Println(x)
	x = RemoveIndex(x, 4)
	fmt.Println(x)

}
func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
