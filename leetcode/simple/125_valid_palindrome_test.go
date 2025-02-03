package simple

import (
	"testing"
)

func Test_125_isPalindrome(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{" ", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"aba", true},
		{"abc", false},
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{"No 'x' in Nixon", true},
		{"Was it a car or a cat I saw?", true},
		{"Not a palindrome", false},
	}

	for _, tc := range testCases {
		result := _125_isPalindrome(tc.input)
		if result != tc.expected {
			t.Errorf("isPalindrome(%q) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}
