package main

import "fmt"

func main() {
	y := 1
	fmt.Println("before pass the address main func", y)
	fmt.Println("before pass the address main func", &y)
	goes(&y)
	fmt.Println("after pass the address main func", y)
	fmt.Println("after pass the address main func", &y)

}
func goes(z *int) {
	fmt.Println("before dereferencing goes func", z)
	fmt.Println("before dereferencing goes func", *z)
	*z = 73
	fmt.Println("after dereferencing goes func", z)
	fmt.Println("after dereferencing goes func", *z)

}
