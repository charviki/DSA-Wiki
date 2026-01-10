package easy

import (
	"testing"
)

func Test_922_sortArrayByParityII(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		{
			name: "Example 1",
			nums: []int{4, 2, 5, 7},
		},
		{
			name: "Example 2",
			nums: []int{2, 3},
		},
		{
			name: "Larger input",
			nums: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)
			got := _922_sortArrayByParityII(numsCopy)

			// 1. Check permutation
			if !isPermutation(got, tt.nums) {
				t.Errorf("sortArrayByParityII() result is not a permutation of input. got = %v, want permutation of %v", got, tt.nums)
				return
			}

			// 2. Check parity at indices
			for i, v := range got {
				if i%2 == 0 {
					if v%2 != 0 {
						t.Errorf("sortArrayByParityII() = %v, index %d is even but value %d is odd", got, i, v)
						return
					}
				} else {
					if v%2 == 0 {
						t.Errorf("sortArrayByParityII() = %v, index %d is odd but value %d is even", got, i, v)
						return
					}
				}
			}
		})
	}
}
