package main

import "fmt"

func main() {
	goes(4, 5, 5, 6, 7, 8, 8)
	goes()
	xi := []int{34, 5, 6, 7, 88, 7, 5, 3444}
	goesTo(xi...)

}
func goes(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(cap(x))
	fmt.Println(len(x))
}
func goesTo(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(cap(x))
	fmt.Println(len(x))
}
