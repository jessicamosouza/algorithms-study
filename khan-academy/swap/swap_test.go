package swap_test

import (
	"reflect"
	"testing"

	swap "problems/khan-academy/swap"
)

func TestSwap(t *testing.T) {
	testCase := []struct {
		nums   []int
		output []int
	}{
		{
			nums:   []int{7, 9, 4},
			output: []int{4, 7, 9},
		},
		{
			nums:   []int{5, 0, 9, 12, 3, -1},
			output: []int{-1, 0, 3, 5, 9, 12},
		},
		{
			nums:   []int{13, 19, 18, 4, 10},
			output: []int{4, 10, 13, 18, 19},
		},
	}

	for _, test := range testCase {
		result := swap.Swap(test.nums)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("Função retornou %v, porém era para retornar %v, dado os valores de %v", result, test.output, test.nums)
		}
	}
}
