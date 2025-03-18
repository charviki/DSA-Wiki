package easy

import (
	"reflect"
	"testing"
)

func Test_283_moveZeroes(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"EmptyArray", []int{}, []int{}},
		{"NoZeroes", []int{1, 2, 3}, []int{1, 2, 3}},
		{"AllZeroes", []int{0, 0, 0}, []int{0, 0, 0}},
		{"MixedElements", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"SingleZero", []int{0}, []int{0}},
		{"SingleNonZero", []int{1}, []int{1}},
		{"ZeroesAtStart", []int{0, 1, 2, 3}, []int{1, 2, 3, 0}},
		{"ZeroesInMiddle", []int{1, 0, 2, 0, 3}, []int{1, 2, 3, 0, 0}},
		{"ZeroesAtEnd", []int{1, 2, 3, 0}, []int{1, 2, 3, 0}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_283_moveZeroes(tc.input)
			if !reflect.DeepEqual(tc.input, tc.expected) {
				t.Errorf("For input %v, expected %v but got %v", tc.input, tc.expected, tc.input)
			}
		})
	}
}
