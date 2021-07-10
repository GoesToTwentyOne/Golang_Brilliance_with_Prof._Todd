package main

import "fmt"

type person struct {
	fname          string
	lname          string
	icecrameFlavor []string
}

func main() {
	p1 := person{
		fname: "Nihad ",
		lname: "Hossain",
		icecrameFlavor: []string{
			"Banana Cream Pie",
			"Banana Nut Bread",
			"Birthday Cake",
		},
	}
	p2 := person{
		fname: "Shalona ",
		lname: "Hossain",
		icecrameFlavor: []string{
			"Black Cherry",
			"Black Raspberry",
			"Butter Brickle",
		},
	}
	fmt.Println(p1)
	fmt.Println(p2)
	for i, value := range p1.icecrameFlavor {
		fmt.Println(i, value)
	}
	fmt.Println("-------------------------------------------------------------------------------------------------")
	for i, value := range p2.icecrameFlavor {
		fmt.Println(i, value)
	}

}
