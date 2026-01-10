package easy

import "testing"

func Test_234_isPalindrome(t *testing.T) {
	// Helper to create list
	createList := func(vals []int) *_234_ListNode {
		if len(vals) == 0 {
			return nil
		}
		head := &_234_ListNode{Val: vals[0]}
		curr := head
		for i := 1; i < len(vals); i++ {
			curr.Next = &_234_ListNode{Val: vals[i]}
			curr = curr.Next
		}
		return head
	}

	tests := []struct {
		name string
		head *_234_ListNode
		want bool
	}{
		{
			name: "Example 1",
			head: createList([]int{1, 2, 2, 1}),
			want: true,
		},
		{
			name: "Example 2",
			head: createList([]int{1, 2}),
			want: false,
		},
		{
			name: "Single node",
			head: createList([]int{1}),
			want: true,
		},
		{
			name: "Empty list",
			head: createList([]int{}),
			want: true,
		},
		{
			name: "Odd length palindrome",
			head: createList([]int{1, 2, 3, 2, 1}),
			want: true,
		},
		{
			name: "Odd length not palindrome",
			head: createList([]int{1, 2, 3, 4, 5}),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _234_isPalindrome(tt.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_234_isPalindrome2(t *testing.T) {
	// Helper to create list
	createList := func(vals []int) *_234_ListNode {
		if len(vals) == 0 {
			return nil
		}
		head := &_234_ListNode{Val: vals[0]}
		curr := head
		for i := 1; i < len(vals); i++ {
			curr.Next = &_234_ListNode{Val: vals[i]}
			curr = curr.Next
		}
		return head
	}

	tests := []struct {
		name string
		head *_234_ListNode
		want bool
	}{
		{
			name: "Example 1",
			head: createList([]int{1, 2, 2, 1}),
			want: true,
		},
		{
			name: "Example 2",
			head: createList([]int{1, 2}),
			want: false,
		},
		{
			name: "Single node",
			head: createList([]int{1}),
			want: true,
		},
		{
			name: "Empty list",
			head: createList([]int{}),
			want: true,
		},
		{
			name: "Odd length palindrome",
			head: createList([]int{1, 2, 3, 2, 1}),
			want: true,
		},
		{
			name: "Odd length not palindrome",
			head: createList([]int{1, 2, 3, 4, 5}),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _234_isPalindrome2(tt.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_234_isPalindrome3(t *testing.T) {
	// Helper to create list
	createList := func(vals []int) *_234_ListNode {
		if len(vals) == 0 {
			return nil
		}
		head := &_234_ListNode{Val: vals[0]}
		curr := head
		for i := 1; i < len(vals); i++ {
			curr.Next = &_234_ListNode{Val: vals[i]}
			curr = curr.Next
		}
		return head
	}

	tests := []struct {
		name string
		head *_234_ListNode
		want bool
	}{
		{
			name: "Example 1",
			head: createList([]int{1, 2, 2, 1}),
			want: true,
		},
		{
			name: "Example 2",
			head: createList([]int{1, 2}),
			want: false,
		},
		{
			name: "Single node",
			head: createList([]int{1}),
			want: true,
		},
		{
			name: "Empty list",
			head: createList([]int{}),
			want: true,
		},
		{
			name: "Odd length palindrome",
			head: createList([]int{1, 2, 3, 2, 1}),
			want: true,
		},
		{
			name: "Odd length not palindrome",
			head: createList([]int{1, 2, 3, 4, 5}),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _234_isPalindrome3(tt.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
