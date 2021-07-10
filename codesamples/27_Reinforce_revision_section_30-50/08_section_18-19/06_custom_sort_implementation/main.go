package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name  string
	Age   int
	Level string
}
type ByName []Person

func (byn ByName) Len() int           { return len(byn) }
func (byn ByName) Swap(i, j int)      { byn[i], byn[j] = byn[j], byn[i] }
func (byn ByName) Less(i, j int) bool { return byn[i].Name < byn[j].Name }

func main() {
	p1 := Person{
		Name:  "Nihad Hossain",
		Age:   21,
		Level: "Hats off",
	}
	p2 := Person{
		Name:  "Seneya Leodenous",
		Age:   31,
		Level: "Great job",
	}
	p3 := Person{
		Name:  "Anyel Leodenous",
		Age:   41,
		Level: "Incredible",
	}
	p4 := Person{
		Name:  "Alina  Dcrus",
		Age:   25,
		Level: "lebel up",
	}
	pre := []Person{p1, p2, p3, p4}
	fmt.Println(pre)
	sort.Sort(ByName(pre))
	fmt.Println(pre)

}
