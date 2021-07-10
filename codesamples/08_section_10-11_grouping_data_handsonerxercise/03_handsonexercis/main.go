package main

import "fmt"

func main() {
	slicein := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(slicein[:5])
	fmt.Println(slicein[5:])
	fmt.Println(slicein[2:7])
	fmt.Println(slicein[1:6])
	fmt.Println(slicein)
}
