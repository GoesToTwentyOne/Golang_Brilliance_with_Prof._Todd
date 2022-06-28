package main

import "fmt"

type person struct {
	first string
	last  string
}
type student struct {
	person
	education string
}

func (s student) speak() {
	fmt.Printf("My name is %v %v . I'm %v student.\n", s.first, s.last, s.education)
}
func main() {
	s1 := student{
		person: person{
			first: "Nihad",
			last:  "Hossain",
		}, education: "First Year",
	}
	s2 := student{
		person: person{
			first: "Rimdny",
			last:  "Dxey",
		}, education: "Third Year",
	}
	s1.speak()
	s2.speak()

}
