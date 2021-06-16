package main

import "fmt"

func main() {
	a := 24
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	//one way
	b := &a
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", &b)
	//another way
	var c *int = &a
	fmt.Println(c)
	fmt.Println(*c)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", &c)

	//changing value
	*b = 45
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", &b)
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

}
