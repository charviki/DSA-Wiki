package easy

import "testing"

func Test_268_missingNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"Example 1", []int{3, 0, 1}, 2},
		{"Example 2", []int{0, 1}, 2},
		{"Example 3", []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
		{"Single Element Missing 1", []int{0}, 1},
		{"Single Element Missing 0", []int{1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _268_missingNumber(tt.nums); got != tt.want {
				t.Errorf("_268_missingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
