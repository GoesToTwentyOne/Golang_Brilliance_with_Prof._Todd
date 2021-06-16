package main

import "fmt"

type person struct {
	first string
	last  string
}

func changeme(p *person) {
	fmt.Println("before dereferencing change me func ", p)
	fmt.Println("before dereferencing change me func ", *p)
	p.first = "Nihad"
	p.last = "Hossain"
	// not work
	//*(p).first = "Nihad"

	fmt.Println("after dereferencing change me func ", p)
	fmt.Println("after dereferencing change me func ", *p)

}
func main() {
	p := person{
		first: "Sunny",
		last:  "Baeltyru",
	}
	fmt.Println("before pass value main func", p)
	changeme(&p)
	fmt.Println("after pass value main func ", p)

}
