package easy

import (
	"reflect"
	"testing"
)

func Test_876_middleNode(t *testing.T) {
	// Helper to create list
	createList := func(vals []int) *_876_ListNode {
		if len(vals) == 0 {
			return nil
		}
		head := &_876_ListNode{Val: vals[0]}
		curr := head
		for i := 1; i < len(vals); i++ {
			curr.Next = &_876_ListNode{Val: vals[i]}
			curr = curr.Next
		}
		return head
	}

	// Helper to convert list to slice
	listToSlice := func(head *_876_ListNode) []int {
		var res []int
		for head != nil {
			res = append(res, head.Val)
			head = head.Next
		}
		return res
	}

	tests := []struct {
		name string
		head *_876_ListNode
		want []int
	}{
		{
			name: "Example 1",
			head: createList([]int{1, 2, 3, 4, 5}),
			want: []int{3, 4, 5},
		},
		{
			name: "Example 2",
			head: createList([]int{1, 2, 3, 4, 5, 6}),
			want: []int{4, 5, 6},
		},
		{
			name: "Single node",
			head: createList([]int{1}),
			want: []int{1},
		},
		{
			name: "Two nodes",
			head: createList([]int{1, 2}),
			want: []int{2},
		},
		{
			name: "Three nodes",
			head: createList([]int{1, 2, 3}),
			want: []int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _876_middleNode(tt.head); !reflect.DeepEqual(listToSlice(got), tt.want) {
				t.Errorf("middleNode() = %v, want %v", listToSlice(got), tt.want)
			}
		})
	}
}
