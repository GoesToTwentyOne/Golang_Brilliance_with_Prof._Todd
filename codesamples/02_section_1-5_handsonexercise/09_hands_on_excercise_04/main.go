package main

import "fmt"

type bando int

var x bando

func main() {
	fmt.Println(x)
	fmt.Printf("%T \n", x)
	x = 42
	fmt.Print(x)
}
