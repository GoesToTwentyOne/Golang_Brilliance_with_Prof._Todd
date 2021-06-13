package main

import "fmt"

func main() {
	tmp := struct {
		name     string
		item     map[string]int
		choclate []string
	}{
		name: "Neaha Kakar",
		item: map[string]int{
			"hiphop": 45,
			"Hotdog": 785,
			"Hindi":  455,
		},
		choclate: []string{
			"Milk ",
			"White",
			"dark",
		},
	}
	fmt.Println(tmp)
	for key, value := range tmp.item {
		fmt.Println(key, value)
	}
	for i, value := range tmp.choclate {
		fmt.Println(i, value)
	}
}
