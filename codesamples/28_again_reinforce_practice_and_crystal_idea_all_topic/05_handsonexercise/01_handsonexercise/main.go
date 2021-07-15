package main

import "fmt"

type person struct {
	first   string
	last    string
	fav_ice []string
}

func main() {

	p1 := person{
		first:   "Slena",
		last:    "Aniya",
		fav_ice: []string{"bnanana", "chocklati", "cokbar", "crime"},
	}
	p2 := person{
		first:   "Martin ",
		last:    "Arel",
		fav_ice: []string{"Weafer", "chocklati", "cokbar", "crime"},
	}
	fmt.Println(p1)
	fmt.Println(p2)

	for _, v := range p1.fav_ice {
		fmt.Println(v)

	}
	fmt.Println("-----------------------------------------------------------------------------------")
	for _, v := range p2.fav_ice {
		fmt.Println(v)

	}
}
