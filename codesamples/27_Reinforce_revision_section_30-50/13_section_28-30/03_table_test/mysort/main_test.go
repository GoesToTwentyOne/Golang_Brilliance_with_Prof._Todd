package mysort

import (
	"testing"
)

func TestDesLast(t *testing.T) {
	type tabletest struct {
		data   []int
		result int
	}
	ta := []tabletest{
		{data: []int{1, 2, 4, 7, 9}, result: 9},
		{data: []int{1, 2, 4, 7, 5}, result: 7},
		{data: []int{1, 2, 4, 7, 10}, result: 10},
		{data: []int{1, 2, 4, 7, 17}, result: 17},
		{data: []int{1, 2, 4, 27, 9}, result: 27},
		{data: []int{1, 2, 4, 7, 59}, result: 59},
		{data: []int{1, 2, 4, 67, 9}, result: 67},
		{data: []int{1, 2, 4, 7, 95454455}, result: 95454455},
		{data: []int{1, 2, 4, 7, 9}, result: 9},
		{data: []int{1, 2, 4, 7, 9}, result: 9},
		{data: []int{1, 2, 4, 7, 9}, result: 9},
		{data: []int{1, 2, 4, 890, 9}, result: 890},
		{data: []int{-1, 2, -4, 7, 9}, result: 9},
		{data: []int{-1, 2, 4, -7, -9}, result: 4},
		{data: []int{1, 2, 4, 7555555555555555, 9}, result: 7555555555555555},
		{data: []int{1, 2, -77777777777774, 7, 9}, result: 9},
		{data: []int{-0, 1, 2, 4, 7, 9}, result: 9},
	}
	for _, v := range ta {
		if DesLast(v.data...) != v.result {
			t.Error("Expected : ", v.result, "Got : ", DesLast(v.data...))
		}
	}

}
