package main

import "fmt"

type mytype int

var x mytype
var y int

func main() {
	fmt.Printf("%T \n ", x)
	fmt.Println(x)
	x = 42
	fmt.Printf("%T \n ", x)
	fmt.Println(x)
	y = int(x)
	fmt.Printf("%T \n ", y)
	fmt.Println(y)

}
