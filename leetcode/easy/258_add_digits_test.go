package easy

import (
	"testing"
)

func Test_258_addDigits(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{"Example 1", 38, 2},
		{"Example 2", 0, 0},
		{"Example 3", 10, 1},
		{"Single Digit", 5, 5},
		{"Multiple Steps", 12345, 6}, // 1+2+3+4+5 = 15 -> 1+5 = 6
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _258_addDigits(tt.num); got != tt.want {
				t.Errorf("_258_addDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
