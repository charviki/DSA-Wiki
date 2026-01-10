package easy

import "testing"

func Test_680_validPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Example 1",
			s:    "aba",
			want: true,
		},
		{
			name: "Example 2",
			s:    "abca",
			want: true,
		},
		{
			name: "Example 3",
			s:    "abc",
			want: false,
		},
		{
			name: "Delete first char",
			s:    "eccer",
			want: true,
		},
		{
			name: "Delete last char",
			s:    "abccb",
			want: true,
		},
		{
			name: "No deletion needed",
			s:    "racecar",
			want: true,
		},
		{
			name: "Single char",
			s:    "a",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _680_validPalindrome(tt.s); got != tt.want {
				t.Errorf("validPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
