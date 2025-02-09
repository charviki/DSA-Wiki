package medium

import "testing"

func Test_80_removeDuplicates(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{1, 1, 1, 1, 1}, 2},
		{[]int{1, 1, 2, 2, 3, 4, 4, 5}, 8},
		{[]int{1, 1, 1, 2, 2, 3}, 5},
		{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7},
	}

	for _, test := range tests {
		numsCopy := make([]int, len(test.nums))
		copy(numsCopy, test.nums)
		result := _80_removeDuplicates(numsCopy)
		if result != test.expected {
			t.Errorf("removeDuplicates(%v) = %d, want %d", test.nums, result, test.expected)
		}
	}
}

func Test_80_removeDuplicates1(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{1, 1, 1, 1, 1}, 2},
		{[]int{1, 1, 2, 2, 3, 4, 4, 5}, 8},
		{[]int{1, 1, 1, 2, 2, 3}, 5},
		{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7},
	}

	for _, test := range tests {
		numsCopy := make([]int, len(test.nums))
		copy(numsCopy, test.nums)
		result := _80_removeDuplicates1(numsCopy)
		if result != test.expected {
			t.Errorf("removeDuplicates(%v) = %d, want %d", test.nums, result, test.expected)
		}
	}
}
