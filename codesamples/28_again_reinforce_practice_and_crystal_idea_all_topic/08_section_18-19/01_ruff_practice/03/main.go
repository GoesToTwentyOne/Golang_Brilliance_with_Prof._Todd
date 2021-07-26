package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name         string
	Age          int
	UniformColor string
	UniformTYpe  string
	Shoes        string
	Height       float64
}
type ByHeight []Person

func (h ByHeight) Len() int           { return len(h) }
func (h ByHeight) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h ByHeight) Less(i, j int) bool { return h[i].Height < h[j].Height }

func main() {
	people := []Person{
		{
			Name:         "Alina",
			Age:          21,
			UniformColor: "Red",
			UniformTYpe:  "V Neck Jumper",
			Shoes:        "Black",
			Height:       5.5,
		},
		{
			Name:         "Gorge Patric",
			Age:          19,
			UniformColor: "Black",
			UniformTYpe:  "Short Sleeve",
			Shoes:        "White",
			Height:       5.6,
		},
		{
			Name:         "Patric Angela",
			Age:          16,
			UniformColor: "Creami Black",
			UniformTYpe:  "Short sailor",
			Shoes:        "White",
			Height:       5.4,
		},
		{
			Name:         "Halina",
			Age:          21,
			UniformColor: "Yellow",
			UniformTYpe:  "Red short T-shirt",
			Shoes:        "white",
			Height:       6.0,
		},
		{
			Name:         "James Alen",
			Age:          21,
			UniformColor: "Purple",
			UniformTYpe:  "V Neck Jumper",
			Shoes:        "Golden",
			Height:       6.5,
		},
		{
			Name:         "Angela Norttewa",
			Age:          22,
			UniformColor: "White",
			UniformTYpe:  "V Neck Jumper",
			Shoes:        "Red",
			Height:       6.2,
		},
		{
			Name:         "Malina",
			Age:          27,
			UniformColor: "Red Hot",
			UniformTYpe:  "Short Sailor",
			Shoes:        "Black",
			Height:       6.5,
		},
		{
			Name:         "Catrina Sunnt",
			Age:          22,
			UniformColor: "Black Hot",
			UniformTYpe:  "V Neck Jumper",
			Shoes:        "White",
			Height:       6.5,
		},
	}
	fmt.Println(people)
	sort.Sort(ByHeight(people))
	fmt.Println("-----------------------------------------sort by Height--------------------------------------------")
	for _, v := range people {
		fmt.Println(v.Name, v.Age, v.UniformColor, v.UniformTYpe, v.Shoes, v.Height)

	}
	// fmt.Printf("%+v", people)

}
