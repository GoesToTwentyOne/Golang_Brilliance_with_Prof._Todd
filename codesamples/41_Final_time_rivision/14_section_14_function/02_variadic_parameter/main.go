package main

import (
	"fmt"
)

func add(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	fmt.Println(sum)
}
func main() {
	add(4, 5, 6, 7, 8, 6, 7, 77, 7, 3)

}
