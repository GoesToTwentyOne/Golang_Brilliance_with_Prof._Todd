package main

import "fmt"

func main() {
	ch := make(<-chan int)
	go func() {
		// ch <- 23

	}()
	fmt.Println(<-ch)
	fmt.Printf("%T \n", ch)
}
