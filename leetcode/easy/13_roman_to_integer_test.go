package easy

import "testing"

func Test_13_romanToInt(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"III", "III", 3},
		{"IV", "IV", 4},
		{"IX", "IX", 9},
		{"LVIII", "LVIII", 58},
		{"MCMXCIV", "MCMXCIV", 1994},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _13_romanToInt(tt.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
