package main

import "fmt"

type vehicle struct {
	doors  string
	colors string
}
type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	tr := truck{
		vehicle: vehicle{
			doors:  "two",
			colors: "yellow",
		},
		fourWheel: true,
	}
	se := sedan{
		vehicle: vehicle{
			doors:  "two",
			colors: "red",
		},
		luxury: false,
	}
	fmt.Println(tr)
	fmt.Println(se)
	fmt.Println(tr.vehicle)
	fmt.Println(se.vehicle)

}
