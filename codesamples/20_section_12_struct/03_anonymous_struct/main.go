package main

import "fmt"

func main() {
	//haven't any name
	p1 := struct {
		Fname string
		Lname string
		Age   int
	}{
		Fname: "Name",
		Lname: "Hossain",
		Age:   78,
	}
	fmt.Println(p1)
}
