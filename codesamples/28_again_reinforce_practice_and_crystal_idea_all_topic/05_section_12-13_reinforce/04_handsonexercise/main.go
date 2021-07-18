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
		age:   21,
	}
	fmt.Println(p)
}
