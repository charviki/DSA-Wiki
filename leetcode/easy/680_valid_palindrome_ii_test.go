package simple

import "testing"

func Test_680_validPalindrome(t *testing.T) {
	testCases := []struct {
		s        string
		expected bool
	}{
		{"racecar", true},
		{"racecarx", true},
		{"xracecar", true},
		{"racecars", true},
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", true},
		{"aba", true},
		{"abca", true},
		{"abc", false},
	}

	for _, tc := range testCases {
		result := _680_validPalindrome(tc.s)
		if result != tc.expected {
			t.Errorf("For input '%s', expected %v but got %v", tc.s, tc.expected, result)
		}
	}
}
