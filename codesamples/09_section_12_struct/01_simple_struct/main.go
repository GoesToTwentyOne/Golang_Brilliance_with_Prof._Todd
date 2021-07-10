package main

import "fmt"

type person struct {
	Fname string
	Lname string
	Age   int
}

func main() {
	p1 := person{
		Fname: "Nihad",
		Lname: "Hossain",
		Age:   78,
	}
	p2 := person{
		Fname: "Scotiya Lavera",
		Lname: "Alen",
		Age:   21,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Printf("%#v \n", p2)
	fmt.Println(p1.Fname, p1.Lname)
	fmt.Println(p2.Fname, p2.Lname, p2.Age)

}
