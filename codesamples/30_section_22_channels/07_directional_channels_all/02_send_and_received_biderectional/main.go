package main

import "fmt"

func main() {
	ch := make(chan int)
	chr := make(<-chan int) //received
	chs := make(chan<- int) //send
	fmt.Println("------------------------")
	fmt.Printf("%T \n", ch)
	fmt.Printf("%T \n", chr)
	fmt.Printf("%T \n", chs)
}
