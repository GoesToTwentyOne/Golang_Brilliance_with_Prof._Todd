package main

import "fmt"

type love string

var nihad love

func main() {
	fmt.Printf("%#v \n", nihad)
	fmt.Printf("%T ", nihad)
}
