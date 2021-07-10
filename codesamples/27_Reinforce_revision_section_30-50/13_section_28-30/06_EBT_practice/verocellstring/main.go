package vero

import (
	"fmt"
	"strings"
)

func Concat(xs []string) string {
	finals := xs[0]
	for _, v := range xs[1:] {

		finals += " "
		finals += v

	}

	return fmt.Sprint(finals)
}

func Join(xs []string) string {

	return fmt.Sprint(strings.Join(xs, " "))
}

// s := xs[0]
// for _, v := range xs[1:] {
// 	s += " "
// 	s += v
// }
// return s
