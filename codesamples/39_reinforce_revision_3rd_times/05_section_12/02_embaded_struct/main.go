package main

import "fmt"

type emp struct {
	name        string
	designation string
	expert      string
	age         int
	experience  int
}
type office struct {
	emp
	Generator int
	PC_count  int
}

func main() {
	AssusOffice := office{
		emp: emp{
			name:        "Md. Nihad Hossain",
			designation: "Head of Software and Maintanance",
			expert:      "Golang",
			age:         21,
			experience:  3,
		},
		Generator: 10,
		PC_count:  43,
	}
	DellOffice := office{
		emp: emp{
			name:        "Alex Goot",
			designation: "Head of Marketing",
			expert:      "Digital Marketing",
			age:         21,
			experience:  1,
		},
		Generator: 5,
		PC_count:  3,
	}

	fmt.Println(AssusOffice)
	fmt.Println(DellOffice)

}
