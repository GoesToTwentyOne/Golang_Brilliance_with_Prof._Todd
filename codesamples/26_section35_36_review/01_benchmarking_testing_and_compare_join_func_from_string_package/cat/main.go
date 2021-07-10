package cat

import "strings"

func Mystring(sx []string) string {
	s := ""
	for _, value := range sx {
		s += value
		s += " "

	}
	return s

}
func Join(sx []string) string {
	s := strings.Join(sx, " ")
	return s

}
