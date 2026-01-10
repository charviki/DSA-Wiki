package easy

import (
	"reflect"
	"testing"
)

func Test_283_moveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "Example 2",
			nums: []int{0},
			want: []int{0},
		},
		{
			name: "No zeros",
			nums: []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "All zeros",
			nums: []int{0, 0, 0},
			want: []int{0, 0, 0},
		},
		{
			name: "Zeros at start",
			nums: []int{0, 0, 1, 2},
			want: []int{1, 2, 0, 0},
		},
		{
			name: "Zeros at end",
			nums: []int{1, 2, 0, 0},
			want: []int{1, 2, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
            numsCopy := make([]int, len(tt.nums))
            copy(numsCopy, tt.nums)
			_283_moveZeroes(numsCopy)
			if !reflect.DeepEqual(numsCopy, tt.want) {
				t.Errorf("moveZeroes() = %v, want %v", numsCopy, tt.want)
			}
		})
	}
}
