package simple

import (
	"reflect"
	"testing"
)

func Test_922_sortArrayByParityII(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2}, []int{2}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{2, 1}, []int{2, 1}},
		{[]int{1, 3, 2, 4}, []int{4, 3, 2, 1}},
		{[]int{4, 2, 5, 7}, []int{4, 5, 2, 7}},
		{[]int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
		{[]int{2, 2, 2, 2}, []int{2, 2, 2, 2}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{2, 1, 4, 3, 6, 5}},
	}

	for _, tc := range testCases {
		result := _922_sortArrayByParityII(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("sortArrayByParityII(%v) = %v, want %v", tc.input, result, tc.expected)
		}
	}
}
