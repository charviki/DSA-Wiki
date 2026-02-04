package easy

import "testing"

func Test_303_NumArray(t *testing.T) {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := _303_Constructor(nums)

	tests := []struct {
		name  string
		left  int
		right int
		want  int
	}{
		{"Example 1 - Range [0, 2]", 0, 2, 1},   // -2 + 0 + 3 = 1
		{"Example 1 - Range [2, 5]", 2, 5, -1},  // 3 + -5 + 2 + -1 = -1
		{"Example 1 - Range [0, 5]", 0, 5, -3},  // -2 + 0 + 3 + -5 + 2 + -1 = -3
		{"Single Element [0, 0]", 0, 0, -2},     // -2
		{"Single Element [2, 2]", 2, 2, 3},      // 3
		{"Sub Range [1, 4]", 1, 4, 0},           // 0 + 3 + -5 + 2 = 0
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := obj.SumRange(tt.left, tt.right); got != tt.want {
				t.Errorf("SumRange(%d, %d) = %v, want %v", tt.left, tt.right, got, tt.want)
			}
		})
	}
}
