package main

import (
	"fmt"
)

func foo(z *int) {
	fmt.Println("before :", z)
	fmt.Println("before :", *z)
	*z = 666
	fmt.Println("after  :", z)
	fmt.Println("after  :", *z)

}
func main() {
	x := 0
	fmt.Println("before :", x)
	fmt.Println("before :", &x)
	foo(&x)
	fmt.Println("after  :", x)
	fmt.Println("after  :", &x)

}
