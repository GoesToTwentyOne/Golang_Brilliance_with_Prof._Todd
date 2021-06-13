package main

import "fmt"

func main() {
	p := struct {
		fname   string
		lname   string
		licence bool
	}{
		fname:   "Nihad",
		lname:   "Hossain",
		licence: true,
	}
	fmt.Println(p)
}
