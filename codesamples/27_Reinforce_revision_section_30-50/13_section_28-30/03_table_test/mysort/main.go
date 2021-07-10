package mysort

func DesLast(n ...int) int {
	temp := 0

	for i := 0; i < len(n); i++ {
		for j := i + 1; j < len(n); j++ {
			if n[i] < n[j] {
				n[i] = n[j]
				temp = n[i]

			}

		}

	}
	return temp

}
