package main

import "fmt"

func main() {
	p1 := struct {
		first    string
		language string
		age      int
		height   float64
	}{
		first:    "Nihad",
		language: "English",
		age:      21,
		height:   5.6,
	}
	fmt.Println(p1)
}
