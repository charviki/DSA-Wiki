package easy

import (
	"fmt"
	"testing"
)

func Test_70_climbStairs(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{10, 89},
		{20, 10946},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("n=%d", tc.n), func(t *testing.T) {
			result := _70_climbStairs(tc.n)
			if result != tc.expected {
				t.Errorf("climbStairs(%d) = %d; want %d", tc.n, result, tc.expected)
			}
		})
	}
}
