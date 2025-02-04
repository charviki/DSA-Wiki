package sort

import (
	"reflect"
	"testing"
)

func TestInsertion(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"EmptySlice", []int{}, []int{}},
		{"SingleElement", []int{1}, []int{1}},
		{"AlreadySorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"ReverseSorted", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"RandomOrder", []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}, []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}},
		{"WithDuplicates", []int{3, 3, 2, 2, 1, 1}, []int{1, 1, 2, 2, 3, 3}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			Insertion(tc.input)
			if !reflect.DeepEqual(tc.input, tc.expected) {
				t.Errorf("Insertion(%v) = %v, want %v", tc.input, tc.input, tc.expected)
			}
		})
	}
}
