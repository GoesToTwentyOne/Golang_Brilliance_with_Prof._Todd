package main

import "fmt"

func main() {
	s := []string{"Nihadkhan", "Hossain", "JoyBangla", "Cgu", "BillKenndyisthegoodteacher"}
	// fs := largeString(s...)
	// fmt.Println(fs)
	cs := callback(largeString, s...)
	fmt.Println(cs)

}
func largeString(s ...string) string {
	temp := ""
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				s[i] = s[j]
				temp = s[i]

			}

		}

	}
	return temp

}

func callback(f func(s ...string) string, xs ...string) string {
	return f(xs...)

}
