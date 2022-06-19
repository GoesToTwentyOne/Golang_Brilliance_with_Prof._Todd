package main

import "fmt"

func main() {
	p := struct {
		first string
		last  string
		age   int
	}{
		first: "Nihad",
		last:  "Hossain",
		age:   10,
	}
	fmt.Println(p)
	fmt.Println(p.first)
}
