package easy

import (
	"reflect"
	"testing"
)

func Test_2164_sortEvenOdd(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{4, 1, 2, 3}, []int{2, 3, 4, 1}},
	}

	for _, test := range tests {
		_2164_sortEvenOdd(test.nums)
		if !reflect.DeepEqual(test.nums, test.expected) {
			t.Errorf("For input %v, expected %v but got %v", test.nums, test.expected, test.nums)
		}
	}
}
