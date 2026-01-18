package easy

import "testing"

func Test_205_isIsomorphic(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{"Example 1", "egg", "add", true},
		{"Example 2", "foo", "bar", false},
		{"Example 3", "paper", "title", true},
		{"Single char", "a", "a", true},
		{"Single char mapped", "a", "b", true},
		{"Two to one", "ab", "aa", false},
		{"One to two", "aa", "ab", false},
		{"Complex", "badc", "baba", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _205_isIsomorphic(tt.s, tt.t); got != tt.want {
				t.Errorf("_205_isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
