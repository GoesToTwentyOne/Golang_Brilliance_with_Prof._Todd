package main

import "fmt"

func main() {
	var x uint8 = 128 ///alias of byte as a byte have 8 bits
	var y int32 = 128 ///alias of rube  as a UTF coding scheme worked with 4 bytes = 32 bits

	fmt.Println(x)
	fmt.Printf("%T \n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
