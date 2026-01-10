package medium

import (
	"testing"
)

func Test_128_longestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
		{
			name: "Example 2",
			nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want: 9,
		},
		{
			name: "Example 3",
			nums: []int{1, 0, 1, 2},
			want: 3,
		},
		{
			name: "Empty array",
			nums: []int{},
			want: 0,
		},
		{
			name: "Single element",
			nums: []int{10},
			want: 1,
		},
		{
			name: "Duplicates",
			nums: []int{1, 2, 0, 1},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _128_longestConsecutive(tt.nums); got != tt.want {
				t.Errorf("_128_longestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
