package mathe

import "testing"

func TestAddMul(t *testing.T) {
	type table struct {
		data1  int
		data2  int
		result int
	}
	ta := []table{
		{data1: 4, data2: 2, result: 60},
		{data1: 2, data2: 2, result: 40},
		{data1: 4, data2: 3, result: 70},
		{data1: 3, data2: 2, result: 50},
		{data1: 1, data2: 2, result: 30},
		{data1: 3, data2: 3, result: 60},
		{data1: 7, data2: 1, result: 80},
		{data1: 5, data2: 5, result: 100},
		{data1: 100, data2: 100, result: 2000},
	}
	for _, v := range ta {
		if MathAddMul(v.data1, v.data2) != v.result {
			t.Error("expeted ", MathAddMul(v.data1, v.data2), "got ", v.result)

		}
	}
}
