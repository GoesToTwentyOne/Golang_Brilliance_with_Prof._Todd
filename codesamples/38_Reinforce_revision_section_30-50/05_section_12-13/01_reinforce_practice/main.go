package main

import "fmt"

type person struct {
	first  string
	last   string
	favice []string
}

func main() {
	p1 := person{
		first:  "Nihad",
		last:   "Hossain",
		favice: []string{"vanila", "strawerry", "creame americano"},
	}
	p2 := person{
		first:  "Lirtironiya",
		last:   "Patric",
		favice: []string{"vanila choclate", "strawerry flabour", "creame Latto"},
	}
	fmt.Println(p1)
	fmt.Println(p1.first, p1.last, p1.favice)

	for _, value := range p1.favice {
		fmt.Println(value)

	}
	fmt.Println(p2)
	fmt.Println(p2.first, p2.last, p2.favice)

	for _, value := range p1.favice {
		fmt.Println(value)

	}

}
