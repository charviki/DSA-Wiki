package easy

import (
	"reflect"
	"testing"
)

func Test_206_reverseList(t *testing.T) {
	// Helper to create list
	createList := func(vals []int) *_206_ListNode {
		if len(vals) == 0 {
			return nil
		}
		head := &_206_ListNode{Val: vals[0]}
		curr := head
		for i := 1; i < len(vals); i++ {
			curr.Next = &_206_ListNode{Val: vals[i]}
			curr = curr.Next
		}
		return head
	}

	// Helper to convert list to slice
	listToSlice := func(head *_206_ListNode) []int {
		var res []int
		for head != nil {
			res = append(res, head.Val)
			head = head.Next
		}
		return res
	}

	tests := []struct {
		name string
		head *_206_ListNode
		want []int
	}{
		{
			name: "Example 1",
			head: createList([]int{1, 2, 3, 4, 5}),
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "Example 2",
			head: createList([]int{1, 2}),
			want: []int{2, 1},
		},
		{
			name: "Example 3",
			head: createList([]int{}),
			want: nil,
		},
		{
			name: "Single node",
			head: createList([]int{1}),
			want: []int{1},
		},
		{
			name: "Three nodes",
			head: createList([]int{1, 2, 3}),
			want: []int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _206_reverseList(tt.head); !reflect.DeepEqual(listToSlice(got), tt.want) {
				t.Errorf("reverseList() = %v, want %v", listToSlice(got), tt.want)
			}
		})
	}
}

func Test_206_reverseList2(t *testing.T) {
	// Helper to create list
	createList := func(vals []int) *_206_ListNode {
		if len(vals) == 0 {
			return nil
		}
		head := &_206_ListNode{Val: vals[0]}
		curr := head
		for i := 1; i < len(vals); i++ {
			curr.Next = &_206_ListNode{Val: vals[i]}
			curr = curr.Next
		}
		return head
	}

	// Helper to convert list to slice
	listToSlice := func(head *_206_ListNode) []int {
		var res []int
		for head != nil {
			res = append(res, head.Val)
			head = head.Next
		}
		return res
	}

	tests := []struct {
		name string
		head *_206_ListNode
		want []int
	}{
		{
			name: "Example 1",
			head: createList([]int{1, 2, 3, 4, 5}),
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "Example 2",
			head: createList([]int{1, 2}),
			want: []int{2, 1},
		},
		{
			name: "Example 3",
			head: createList([]int{}),
			want: nil,
		},
		{
			name: "Single node",
			head: createList([]int{1}),
			want: []int{1},
		},
		{
			name: "Three nodes",
			head: createList([]int{1, 2, 3}),
			want: []int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _206_reverseList2(tt.head); !reflect.DeepEqual(listToSlice(got), tt.want) {
				t.Errorf("reverseList() = %v, want %v", listToSlice(got), tt.want)
			}
		})
	}
}
