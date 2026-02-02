package easy

import "testing"

func Test_263_isUgly(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"Example 1", 6, true},
		{"Example 2", 1, true},
		{"Example 3", 14, false},
		{"Zero", 0, false},
		{"Negative", -6, false},
		{"Power of 2", 8, true},
		{"Product of 2,3,5", 30, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _263_isUgly(tt.n); got != tt.want {
				t.Errorf("_263_isUgly() = %v, want %v", got, tt.want)
			}
		})
	}
}
