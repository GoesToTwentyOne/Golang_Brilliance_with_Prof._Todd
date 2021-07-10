package main

import "fmt"

type vehicle struct {
	doors  int
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
	v := vehicle{
		doors:  4,
		colors: "red",
	}
	vTruck := truck{
		vehicle: vehicle{
			doors:  2,
			colors: "Yellow",
		},
		fourWheel: true,
	}
	vSedan := sedan{
		vehicle: vehicle{
			doors:  2,
			colors: "white",
		},
		luxury: true,
	}
	fmt.Println(v)
	fmt.Println(vTruck)
	fmt.Println(vSedan)

	fmt.Println(vTruck.doors)
	fmt.Println(vSedan.doors)
}
