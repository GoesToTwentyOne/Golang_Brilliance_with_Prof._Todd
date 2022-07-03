package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 21
		close(c)

	}()

	v, ok := <-c
	if !ok {
		fmt.Println("Not OK : ", v, ok)

	} else {
		fmt.Println("OK:", v, ok)
	}
	v, ok = <-c
	if !ok {
		fmt.Println("Not OK : ", v, ok)

	} else {
		fmt.Println("OK:", v, ok)
	}

}
