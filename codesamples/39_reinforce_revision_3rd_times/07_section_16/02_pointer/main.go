package main

import "fmt"

func main() {
	start := 0
	fmt.Printf("before referencing value is :\t%v\n", start)
	fmt.Printf("before referencing address is :\t%v\n", &start)
	point(&start)
	fmt.Printf("after referencing value is :\t%v\n", start)
	fmt.Printf("after referencing address is :\t%v\n", &start)

}
func point(x *int) {
	fmt.Printf("before dereferencing address is :\t%v\n", x)
	fmt.Printf("before dereferencing value is :\t%v\n", *x)
	*x = 5555
	fmt.Printf("after dereferencing address is :\t%v\n", x)
	fmt.Printf("after dereferencing value is :\t%v\n", *x)

}
