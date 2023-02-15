package main

import (
	"fmt"
)

type stateOfPurpose struct {
	first    string
	last     string
	interest []string
}

func main() {
	p1 := stateOfPurpose{
		first: "Md. Nihad",
		last:  "Hossian",
		interest: []string{
			"Cloud computing",
			"Machine Learning",
			"DevOps",
			"Blockchain",
		},
	}
	p2 := stateOfPurpose{
		first: "Nadira",
		last:  "Khan",
		interest: []string{
			"Cloud computing",
			"Machine Learning",
			"DevOps",
			"Blockchain",
		},
	}
	m := map[string]stateOfPurpose{
		p1.first: p1,
		p1.last:  p2,
	}
	for _, v := range m {
		fmt.Printf(v.first)
		fmt.Println(v.last)
		for i, val := range v.interest {
			fmt.Println(i, val)
		}
	}

}
