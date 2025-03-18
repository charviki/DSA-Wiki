package medium

import "testing"

func Test_11_maxArea(t *testing.T) {
	tests := []struct {
		heights []int
		want    int
	}{
		{[]int{}, 0},
		{[]int{1}, 0},
		{[]int{1, 1}, 1},
		{[]int{1, 2, 3, 4, 5}, 6},
		{[]int{5, 4, 3, 2, 1}, 6},
		{[]int{1, 2, 3, 4, 5, 4, 3, 2, 1}, 12},
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1}, 8},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 25},
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 25},
	}

	for _, tt := range tests {
		got := _11_maxArea(tt.heights)
		if got != tt.want {
			t.Errorf("maxArea(%v) = %d; want %d", tt.heights, got, tt.want)
		}
	}
}

func Test_11_maxArea2(t *testing.T) {
	tests := []struct {
		heights []int
		want    int
	}{
		{[]int{}, 0},
		{[]int{1}, 0},
		{[]int{1, 1}, 1},
		{[]int{1, 2, 3, 4, 5}, 6},
		{[]int{5, 4, 3, 2, 1}, 6},
		{[]int{1, 2, 3, 4, 5, 4, 3, 2, 1}, 12},
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1}, 8},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 25},
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 25},
	}

	for _, tt := range tests {
		got := _11_maxArea2(tt.heights)
		if got != tt.want {
			t.Errorf("maxArea(%v) = %d; want %d", tt.heights, got, tt.want)
		}
	}
}
