package main

import "fmt"

type person struct {
	first     string
	last      string
	iceflavor []string
}

func main() {
	//copy field is not accepted
	p1 := person{
		first: "Nihad",
		last:  "Hossain",
		iceflavor: []string{
			"Banana Cream Pie",
			"Banana Nut Bread",
			"Birthday Cake",
		},
	}
	p2 := person{
		first: "Shalona",
		last:  "Hasan",
		iceflavor: []string{
			"Black Cherry",
			"Black Raspberry",
			"Butter Brickle",
		},
	}
	p3 := person{
		first: "Alena",
		last:  "jan",
		iceflavor: []string{
			"Black Cherry",
			"Black Raspberry",
			"Butter Brickle",
		},
	}
	// fmt.Println(p1)
	x := map[string]person{
		p1.last: p1,
		p2.last: p2,
		p3.last: p3,
	}
	for _, v := range x {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, value := range v.iceflavor {
			fmt.Println(i, value)
		}

	}

}
