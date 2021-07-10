package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []string{"Nihad ", "Anyel", "Slena", "Dubai Sheikh"}
	fmt.Println(xi)
	sort.Strings(xi)
	fmt.Println(xi)
}
