package easy

import "testing"

func Test_326_isPowerOfThree(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"Example 1", 27, true},
		{"Example 2", 0, false},
		{"Example 3", 9, true},
		{"Example 4", 45, false},
		{"One", 1, true},
		{"Three", 3, true},
		{"Negative", -3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _326_isPowerOfThree(tt.n); got != tt.want {
				t.Errorf("_326_isPowerOfThree() = %v, want %v", got, tt.want)
			}
			if got := _326_isPowerOfThree2(tt.n); got != tt.want {
				t.Errorf("_326_isPowerOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
