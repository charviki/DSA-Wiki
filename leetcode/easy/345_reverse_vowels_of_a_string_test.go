package easy

import "testing"

func Test_345_reverseVowels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{s: "IceCreAm"},
			want: "AceCreIm",
		},
		{
			name: "Example 2",
			args: args{s: "leetcode"},
			want: "leotcede",
		},
		{
			name: "No Vowels",
			args: args{s: "bcdfg"},
			want: "bcdfg",
		},
		{
			name: "All Vowels",
			args: args{s: "aeiou"},
			want: "uoiea",
		},
		{
			name: "Mixed Case",
			args: args{s: "aA"},
			want: "Aa",
		},
		{
			name: "Empty String",
			args: args{s: ""},
			want: "",
		},
		{
			name: "Single Character Vowel",
			args: args{s: "a"},
			want: "a",
		},
		{
			name: "Single Character Consonant",
			args: args{s: "b"},
			want: "b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _345_reverseVowels(tt.args.s); got != tt.want {
				t.Errorf("_345_reverseVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
