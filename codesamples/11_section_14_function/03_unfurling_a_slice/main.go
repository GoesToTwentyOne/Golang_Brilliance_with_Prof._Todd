package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	goes("Nihad", xi...)

}
func goes(s string, n ...int) {
	fmt.Println(s, n)

}
