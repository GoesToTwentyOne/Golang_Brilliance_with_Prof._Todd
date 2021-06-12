package main

import "fmt"

type person struct {
	Fname string
	Lname string
	Age   int
}

//embeded type
type Programmer struct {
	person
	world_ranking int
	// colletion field
	Age int
}

func main() {
	p1 := Programmer{
		person: person{
			Fname: "Nihad",
			Lname: "Hossain",
			Age:   21,
		},
		world_ranking: 01,
		Age:           431,
	}
	p2 := Programmer{
		person: person{
			Fname: "Jhanker",
			Lname: "Mahbub",
			Age:   31,
		},
		world_ranking: 01,
		Age:           31,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	//colletion field Print
	fmt.Println(p1.Fname, p1.Lname, p1.person.Age, p1.Age)
	fmt.Println(p2.Fname, p2.Lname, p2.person.Age, p2.Age)

}
