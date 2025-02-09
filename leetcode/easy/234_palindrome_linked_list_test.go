package easy

import "testing"

func Test_234_isPalindrome(t *testing.T) {
	tests := []struct {
		head     *_234_ListNode
		expected bool
	}{
		{nil, true},
		{&_234_ListNode{Val: 1}, true},
		{&_234_ListNode{Val: 1, Next: &_234_ListNode{Val: 1}}, true},
		{&_234_ListNode{Val: 1, Next: &_234_ListNode{Val: 2}}, false},
		{&_234_ListNode{Val: 1, Next: &_234_ListNode{Val: 2, Next: &_234_ListNode{Val: 1}}}, true},
		{&_234_ListNode{Val: 1, Next: &_234_ListNode{Val: 2, Next: &_234_ListNode{Val: 3}}}, false},
		{&_234_ListNode{Val: 1, Next: &_234_ListNode{Val: 2, Next: &_234_ListNode{Val: 2, Next: &_234_ListNode{Val: 1}}}}, true},
		{&_234_ListNode{Val: 1, Next: &_234_ListNode{Val: 2, Next: &_234_ListNode{Val: 3, Next: &_234_ListNode{Val: 2, Next: &_234_ListNode{Val: 1}}}}}, true},
	}

	for _, test := range tests {
		result := _234_isPalindrome(test.head)
		if result != test.expected {
			t.Errorf("For input %v, expected %v but got %v", test.head, test.expected, result)
		}
	}
}

func Test_234_isPalindrome2(t *testing.T) {
	tests := []struct {
		head     *_234_ListNode
		expected bool
	}{
		{nil, true},
		{&_234_ListNode{Val: 1}, true},
		{&_234_ListNode{Val: 1, Next: &_234_ListNode{Val: 1}}, true},
		{&_234_ListNode{Val: 1, Next: &_234_ListNode{Val: 2}}, false},
		{&_234_ListNode{Val: 1, Next: &_234_ListNode{Val: 2, Next: &_234_ListNode{Val: 1}}}, true},
		{&_234_ListNode{Val: 1, Next: &_234_ListNode{Val: 2, Next: &_234_ListNode{Val: 3}}}, false},
		{&_234_ListNode{Val: 1, Next: &_234_ListNode{Val: 2, Next: &_234_ListNode{Val: 2, Next: &_234_ListNode{Val: 1}}}}, true},
		{&_234_ListNode{Val: 1, Next: &_234_ListNode{Val: 2, Next: &_234_ListNode{Val: 3, Next: &_234_ListNode{Val: 2, Next: &_234_ListNode{Val: 1}}}}}, true},
	}

	for _, test := range tests {
		result := _234_isPalindrome2(test.head)
		if result != test.expected {
			t.Errorf("For input %v, expected %v but got %v", test.head, test.expected, result)
		}
	}
}

func Test_234_isPalindrome3(t *testing.T) {
	tests := []struct {
		head     *_234_ListNode
		expected bool
	}{
		{nil, true},
		{&_234_ListNode{Val: 1}, true},
		{&_234_ListNode{Val: 1, Next: &_234_ListNode{Val: 1}}, true},
		{&_234_ListNode{Val: 1, Next: &_234_ListNode{Val: 2}}, false},
		{&_234_ListNode{Val: 1, Next: &_234_ListNode{Val: 2, Next: &_234_ListNode{Val: 1}}}, true},
		{&_234_ListNode{Val: 1, Next: &_234_ListNode{Val: 2, Next: &_234_ListNode{Val: 3}}}, false},
		{&_234_ListNode{Val: 1, Next: &_234_ListNode{Val: 2, Next: &_234_ListNode{Val: 2, Next: &_234_ListNode{Val: 1}}}}, true},
		{&_234_ListNode{Val: 1, Next: &_234_ListNode{Val: 2, Next: &_234_ListNode{Val: 3, Next: &_234_ListNode{Val: 2, Next: &_234_ListNode{Val: 1}}}}}, true},
	}

	for _, test := range tests {
		result := _234_isPalindrome3(test.head)
		if result != test.expected {
			t.Errorf("For input %v, expected %v but got %v", test.head, test.expected, result)
		}
	}
}
