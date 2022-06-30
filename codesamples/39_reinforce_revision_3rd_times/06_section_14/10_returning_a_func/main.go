package main

import "fmt"

func main() {
	x := goes()
	fmt.Printf("%T\n", x)
	y := x()
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}

func goes() func() int {
	return func() int {
		return 321
	}
}
