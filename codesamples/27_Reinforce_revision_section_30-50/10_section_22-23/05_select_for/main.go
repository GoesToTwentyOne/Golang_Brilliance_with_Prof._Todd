package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	//send
	go send(even, odd, quit)
	//recived
	received(even, odd, quit)

	fmt.Println("all about good execuite")
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
	close(e)
	close(o)

}
func received(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("even ", v)
		case v := <-o:
			fmt.Println("odd ", v)
		case v := <-q:
			fmt.Println("quit", v)
			return

		}
	}
}
