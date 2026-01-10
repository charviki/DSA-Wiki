package easy

import (
	"reflect"
	"testing"
)

func Test_905_sortArrayByParity(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		{
			name: "Example 1",
			nums: []int{3, 1, 2, 4},
		},
		{
			name: "Example 2",
			nums: []int{0},
		},
		{
			name: "All even",
			nums: []int{2, 4, 6},
		},
		{
			name: "All odd",
			nums: []int{1, 3, 5},
		},
		{
			name: "Mixed",
			nums: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
            numsCopy := make([]int, len(tt.nums))
            copy(numsCopy, tt.nums)
			got := _905_sortArrayByParity(numsCopy)
			// Verification logic: check if all evens come before all odds
            // and if the elements are the same as input (permutation)
            
            // 1. Check permutation
            if !isPermutation(got, tt.nums) {
                t.Errorf("sortArrayByParity() result is not a permutation of input. got = %v, want permutation of %v", got, tt.nums)
                return
            }

            // 2. Check parity order
            foundOdd := false
            for _, v := range got {
                if v%2 != 0 {
                    foundOdd = true
                } else if foundOdd {
                    // Found even after odd
                    t.Errorf("sortArrayByParity() = %v, even number %d found after odd number", got, v)
                    return
                }
            }
		})
	}
}

func isPermutation(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    m1 := make(map[int]int)
    m2 := make(map[int]int)
    for _, v := range a {
        m1[v]++
    }
    for _, v := range b {
        m2[v]++
    }
    return reflect.DeepEqual(m1, m2)
}
