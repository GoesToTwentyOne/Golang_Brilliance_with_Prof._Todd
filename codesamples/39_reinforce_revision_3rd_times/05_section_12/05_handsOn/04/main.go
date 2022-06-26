package main

import "fmt"

func main() {
	p := struct {
		first string
		last  string
		about map[string]string
	}{
		first: "Nihad",
		last:  "Hossain",
		about: map[string]string{
			"Age":   "Twenty",
			"Class": "High",
		},
	}
	for k, v := range p.about {
		fmt.Printf("%v    %v\n", k, v)
	}
	fmt.Println(p)
}
