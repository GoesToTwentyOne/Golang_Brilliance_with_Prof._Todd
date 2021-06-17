package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	xs := []string{"Dog", "Egg", "Goat", "God", "Hand", "Jug", "Queen"}
	sort.Ints(xi)
	sort.Strings(xs)

	fmt.Println(xi)
	fmt.Println(xs)
}
