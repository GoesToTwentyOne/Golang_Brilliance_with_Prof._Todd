package main

import "fmt"

func main() {
	ch := make(chan int)
	chr := make(<-chan int) //received
	chs := make(chan<- int) //send
	fmt.Printf("%T \n", ch)
	fmt.Printf("%T \n", chr)
	fmt.Printf("%T \n", chs)
	// specific to specific doesn't assign
	chr = ch
	chs = ch
}
