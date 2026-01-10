package medium

import (
	"testing"
)

func Test_215_findKthLargest(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{3, 2, 1, 5, 6, 4},
			k:    2,
			want: 5,
		},
		{
			name: "Example 2",
			nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:    4,
			want: 4,
		},
		{
			name: "Single element",
			nums: []int{1},
			k:    1,
			want: 1,
		},
		{
			name: "K is 1 (Max)",
			nums: []int{10, 20, 30},
			k:    1,
			want: 30,
		},
		{
			name: "K is length (Min)",
			nums: []int{10, 20, 30},
			k:    3,
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy to avoid modifying original nums in multiple runs if implementation sorts in place
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)
			if got := _215_findKthLargest(numsCopy, tt.k); got != tt.want {
				t.Errorf("_215_findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
