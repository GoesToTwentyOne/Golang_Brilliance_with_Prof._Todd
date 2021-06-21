package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	//send
	go send(even, odd, quit)
	received(even, odd, quit)

}
func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}

	}
	q <- 0

}
func received(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("the even value ", v)
		case v := <-o:
			fmt.Println("the odd value ", v)
		case v := <-q:
			fmt.Println("the quit value ", v)
			return

		}
	}

}
