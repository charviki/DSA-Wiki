package medium

import (
	"testing"
)

func Test_560_subarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, 1, 1},
			k:    2,
			want: 2,
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 3},
			k:    3,
			want: 2,
		},
		{
			name: "Negative numbers",
			nums: []int{1, -1, 1},
			k:    1,
			want: 3, // [1], [1, -1, 1], [1]
		},
		{
			name: "No subarray",
			nums: []int{1, 2, 3},
			k:    10,
			want: 0,
		},
		{
			name: "Single element match",
			nums: []int{3},
			k:    3,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _560_subarraySum(tt.nums, tt.k); got != tt.want {
				t.Errorf("_560_subarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
