package main

import "fmt"

func main() {
	a := 45 // 0010 1101
	b := 10 // 0000 1010
	var c int

	//bitwise &
	c = a & b
	fmt.Println(c)
	fmt.Printf("%b \n", c)
	//bitwise |
	c = a | b
	fmt.Println(c)
	fmt.Printf("%b \n", c)
	//bitwise ^
	c = a ^ b
	fmt.Println(c)
	fmt.Printf("%b \n", c)

}
