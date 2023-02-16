package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 5, 3, 2, 1, 6, 666, 77, 8888}
	xs := []string{"Nihad", "Dina", "Handalis", "Alex", "Nadirah"}
	sort.Ints(xi)
	sort.Strings(xs)
	fmt.Println(xi)
	fmt.Println(xs)
}
