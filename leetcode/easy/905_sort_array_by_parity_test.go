package easy

import (
	"reflect"
	"testing"
)

func TestSortArrayByParity(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{2, 4, 6}, []int{2, 4, 6}},
		{[]int{1, 3, 5}, []int{1, 3, 5}},
		{[]int{1}, []int{1}},
		{[]int{2, 3, 4, 5}, []int{2, 4, 3, 5}},
	}

	for _, test := range tests {
		result := _905_sortArrayByParity(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("sortArrayByParity(%v) = %v, want %v", test.input, result, test.expected)
		}
	}
}
