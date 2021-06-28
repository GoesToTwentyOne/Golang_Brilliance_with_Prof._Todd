package main

import "fmt"

func main() {
	a := 2
	fmt.Printf("decimal : %d \t  binary : %b \t hexa-decimal : %#x \n", a, a, a)
	a = 10 << 1
	fmt.Printf("decimal : %d \t  binary : %b \t hexa-decimal : %#x \n", a, a, a)
}
