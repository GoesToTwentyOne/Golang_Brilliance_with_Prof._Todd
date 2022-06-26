package main

import "fmt"

type person struct {
	first     string
	last      string
	f_flavour []string
}

func main() {
	p1 := person{
		first:     "Nihad ",
		last:      "Hossain",
		f_flavour: []string{"Cream", "Condence milk", "Vanila", "Ice"},
	}

	p2 := person{
		first:     "Alex ",
		last:      "Goot",
		f_flavour: []string{"Cream", "Condence milk", "Vanila", "Ice"},
	}
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for k, v := range p1.f_flavour {
		fmt.Println(k, v)

	}
	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for k, v := range p2.f_flavour {
		fmt.Println(k, v)

	}
	fmt.Println("///////////////////////////////////////////////////////////////////////////////////////")
	mp := map[string]person{
		p1.first: p1,
		p2.last:  p2,
	}
	for _, v := range mp {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for k, v := range v.f_flavour {
			fmt.Println(k, v)
		}
		fmt.Printf("\n \n")

	}
}
