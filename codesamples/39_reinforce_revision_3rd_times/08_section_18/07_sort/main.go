package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{56, 4, 5, 99, 34, 66, 77, 83, 6}
	xs := []string{"Nihad", "Alex", "Gtyu", "Seban", "Todd", "Arindal"}
	sort.Ints(xi)
	sort.Strings(xs)
	fmt.Println(xi)
	fmt.Println(xs)
}
