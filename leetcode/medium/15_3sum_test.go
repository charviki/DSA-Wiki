package medium

import (
	"testing"
)

func Test_15_3sum(t *testing.T) {
	tests := []struct {
		nums     []int
		expected [][]int
	}{
		{[]int{}, [][]int{}},
		{[]int{1}, [][]int{}},
		{[]int{1, 2}, [][]int{}},
		{[]int{1, 2, 3}, [][]int{}},
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{-1, 0, 1, 0}, [][]int{{-1, 0, 1}}},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{[]int{-2, 0, 1, 1, 2}, [][]int{{-2, 0, 2}, {-2, 1, 1}}},
	}

	for _, test := range tests {
		result := _15_3sum(test.nums)
		if !_15_equal(result, test.expected) {
			t.Errorf("For input %v, expected %v, got %v", test.nums, test.expected, result)
		}
	}
}

// Helper function to compare two slices of slices.
func _15_equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
