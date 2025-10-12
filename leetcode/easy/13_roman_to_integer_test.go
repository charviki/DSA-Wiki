package easy

import "testing"

func Test_13_romanToInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"I", 1},
		{"XXVII", 27},
		{"XLIX", 49},
		{"CMXCIX", 999},
		{"MMMDCCCLXXXVIII", 3888},
	}

	for _, tt := range tests {
		result := _13_romanToInt(tt.input)
		if result != tt.expected {
			t.Errorf("input: %q, got: %d, want: %d", tt.input, result, tt.expected)
		}
	}
}
