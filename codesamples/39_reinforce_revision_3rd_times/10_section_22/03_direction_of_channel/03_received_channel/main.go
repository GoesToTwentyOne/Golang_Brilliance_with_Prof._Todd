package main

import "fmt"

func main() {
	ch := make(<-chan int, 2)

	//doesn't work
	//ch <- 21

	fmt.Println(<-ch)
}
