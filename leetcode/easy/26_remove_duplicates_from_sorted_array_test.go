package easy

import "testing"

func Test_26_removeDuplicates(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2, 3}, 3},
		{[]int{1, 1, 1}, 1},
		{[]int{1, 2, 2, 3, 4, 4, 5}, 5},
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}

	for _, tc := range testCases {
		result := _26_removeDuplicates(tc.nums)
		if result != tc.expected {
			t.Errorf("For input %v, expected length %d but got %d", tc.nums, tc.expected, result)
		}
	}
}
