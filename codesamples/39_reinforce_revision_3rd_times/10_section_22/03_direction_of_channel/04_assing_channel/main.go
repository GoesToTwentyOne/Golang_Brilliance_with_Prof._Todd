package main

import "fmt"

func main() {
	c := make(chan int)
	cs := make(chan<- int)
	cr := make(<-chan int)

	//general to specific
	//does work
	cs = c
	cr = c
	//specific to general
	//doesn't work
	//	c=cs
	//	c=cr

	///conversion

	//general to specific
	//does work
	fmt.Println((chan<- int)(c))
	fmt.Println((<-chan int)(c))

	//specific to general
	//doesn't work
	//fmt.Println((c)(chan<-int))
	//fmt.Println((c)(<-chan int))
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cs)
	fmt.Printf("%T\n", cr)

}
