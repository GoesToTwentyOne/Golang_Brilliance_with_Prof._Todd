package main

import "fmt"

type hotdog int

// var x hotdog = 4

func main() {
	// fmt.Println(x)
	// fmt.Printf("%T \n", x)
	// x.numAdd(4)
	a := hotdog(4)
	a.numAdd(4)

}
func (h hotdog) numAdd(n int) {
	fmt.Println(int(h) + n)

}
