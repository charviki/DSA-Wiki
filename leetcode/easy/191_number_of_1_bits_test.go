package easy

import "testing"

func Test_191_hammingWeight(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"Example 1", 11, 3},
		{"Example 2", 128, 1},
		{"Example 3", 2147483645, 30},
		{"Smallest number", 1, 1},
		{"Max Int32", 2147483647, 31},
		{"Power of 2 minus 1", 3, 2},
		{"Arbitrary number", 29, 4}, // 11101 -> 4
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _191_hammingWeight(tt.n); got != tt.want {
				t.Errorf("_191_hammingWeight() = %v, want %v", got, tt.want)
			}
			if got := _191_hammingWeight2(tt.n); got != tt.want {
				t.Errorf("_191_hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
