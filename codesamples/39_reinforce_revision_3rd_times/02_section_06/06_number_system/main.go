package main

import "fmt"

func main() {
	s := "N"
	fmt.Println(s)
	fmt.Printf("%v\n", s)
	fmt.Printf("%d\n", s[0])
	fmt.Printf("%#x\n", s[0])
	fmt.Printf("%b\n", s[0])
	fmt.Printf("%O\n", s[0])

}
