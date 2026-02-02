package easy

import (
	"testing"
)

func Test_231_isPowerOfTwo(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"Example 1", 1, true},
		{"Example 2", 16, true},
		{"Example 3", 3, false},
		{"Negative", -16, false},
		{"Zero", 0, false},
		{"Large Power of Two", 1024, true},
		{"Large Non-Power of Two", 1023, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _231_isPowerOfTwo(tt.n); got != tt.want {
				t.Errorf("_231_isPowerOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
