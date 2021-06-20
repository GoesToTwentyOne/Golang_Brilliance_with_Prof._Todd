package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	//over flow
	//buffer pipeline one data send 2
	ch <- 21
	ch <- 22
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
