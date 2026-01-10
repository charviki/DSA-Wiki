package easy

import (
	"reflect"
	"testing"
)

func Test_2164_sortEvenOdd(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{4, 1, 2, 3},
			want: []int{2, 3, 4, 1},
		},
		{
			name: "Example 2",
			nums: []int{2, 1},
			want: []int{2, 1},
		},
		{
			name: "Single element",
			nums: []int{5},
			want: []int{5},
		},
		{
			name: "Three elements",
			nums: []int{5, 2, 8},
			want: []int{5, 2, 8}, // odd idx: [2], even idx: [5, 8]. Sorted: odd [2], even [5, 8]. Result: 5, 2, 8
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)
			if got := _2164_sortEvenOdd(numsCopy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortEvenOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
