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
	ma := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	fmt.Println(ma)
	fmt.Println(ma[p1.last])
	ma["Renasa"] = person{
		first: "Kutyula",
		last:  "Henry",
		favice: []string{
			"red vanila",
			"Cholate vanila",
		},
	}
	fmt.Println(ma)
	for key, v := range ma {
		fmt.Println("key : ", key)
		fmt.Println("first last name :", v.first, v.last)
		for i, v := range v.favice {
			fmt.Println(i, v)
		}

	}

}
