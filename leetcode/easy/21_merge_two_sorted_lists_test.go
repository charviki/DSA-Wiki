package easy

import (
	"reflect"
	"testing"
)

func Test_21_mergeTwoLists(t *testing.T) {
	// Helper function to create a list from slice
	createList := func(nums []int) *_21_ListNode {
		if len(nums) == 0 {
			return nil
		}
		head := &_21_ListNode{Val: nums[0]}
		current := head
		for i := 1; i < len(nums); i++ {
			current.Next = &_21_ListNode{Val: nums[i]}
			current = current.Next
		}
		return head
	}

	// Helper function to convert list to slice
	listToSlice := func(head *_21_ListNode) []int {
		var res []int
		for head != nil {
			res = append(res, head.Val)
			head = head.Next
		}
		if len(res) == 0 {
			return []int{}
		}
		return res
	}

	tests := []struct {
		name  string
		list1 []int
		list2 []int
		want  []int
	}{
		{
			name:  "Example 1",
			list1: []int{1, 2, 4},
			list2: []int{1, 3, 4},
			want:  []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:  "Example 2",
			list1: []int{},
			list2: []int{},
			want:  []int{},
		},
		{
			name:  "Example 3",
			list1: []int{},
			list2: []int{0},
			want:  []int{0},
		},
		{
			name:  "One empty list",
			list1: []int{1},
			list2: []int{},
			want:  []int{1},
		},
		{
			name:  "Disjoint lists",
			list1: []int{1, 3, 5},
			list2: []int{2, 4, 6},
			want:  []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		fns := []func(*_21_ListNode, *_21_ListNode) *_21_ListNode{
			_21_mergeTwoLists,
			_21_mergeTwoLists2,
		}
		for _, fn := range fns {
			t.Run(tt.name, func(t *testing.T) {
				l1 := createList(tt.list1)
				l2 := createList(tt.list2)
				got := fn(l1, l2)
				// Handle nil result for empty expected list
				gotSlice := listToSlice(got)
				if len(tt.want) == 0 && len(gotSlice) == 0 {
					return
				}
				if !reflect.DeepEqual(gotSlice, tt.want) {
					t.Errorf("mergeTwoLists() = %v, want %v", gotSlice, tt.want)
				}
			})
		}
	}
}
