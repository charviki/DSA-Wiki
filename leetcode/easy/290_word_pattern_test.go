package easy

import "testing"

func Test_290_wordPattern(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		s       string
		want    bool
	}{
		{"Example 1", "abba", "dog cat cat dog", true},
		{"Example 2", "abba", "dog cat cat fish", false},
		{"Example 3", "aaaa", "dog cat cat dog", false},
		{"Example 4", "abba", "dog dog dog dog", false},
		{"Length Mismatch", "aaa", "aa aa aa aa", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _290_wordPattern(tt.pattern, tt.s); got != tt.want {
				t.Errorf("_290_wordPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
