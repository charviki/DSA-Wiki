package medium

import (
	"reflect"
	"testing"
)

func _19_sliceToList(nums []int) *_19_ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &_19_ListNode{Val: nums[0]}
	curr := head
	for i := 1; i < len(nums); i++ {
		curr.Next = &_19_ListNode{Val: nums[i]}
		curr = curr.Next
	}
	return head
}

func _19_listToSlice(head *_19_ListNode) []int {
	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return nums
}

func Test_19_removeNthFromEnd(t *testing.T) {
	tests := []struct {
		name string
		head []int
		n    int
		want []int
	}{
		{
			name: "Example 1",
			head: []int{1, 2, 3, 4, 5},
			n:    2,
			want: []int{1, 2, 3, 5},
		},
		{
			name: "Example 2",
			head: []int{1},
			n:    1,
			want: []int{},
		},
		{
			name: "Example 3",
			head: []int{1, 2},
			n:    1,
			want: []int{1},
		},
		{
			name: "Remove head",
			head: []int{1, 2},
			n:    2,
			want: []int{2},
		},
		{
			name: "Remove from 3 elements",
			head: []int{1, 2, 3},
			n:    2,
			want: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := _19_sliceToList(tt.head)
			got := _19_removeNthFromEnd(head, tt.n)
			if !reflect.DeepEqual(_19_listToSlice(got), tt.want) {
				t.Errorf("_19_removeNthFromEnd() = %v, want %v", _19_listToSlice(got), tt.want)
			}
		})
	}
}

func Test_19_removeNthFromEnd2(t *testing.T) {
	tests := []struct {
		name string
		head []int
		n    int
		want []int
	}{
		{
			name: "Example 1",
			head: []int{1, 2, 3, 4, 5},
			n:    2,
			want: []int{1, 2, 3, 5},
		},
		{
			name: "Example 2",
			head: []int{1},
			n:    1,
			want: []int{},
		},
		{
			name: "Example 3",
			head: []int{1, 2},
			n:    1,
			want: []int{1},
		},
		{
			name: "Remove head",
			head: []int{1, 2},
			n:    2,
			want: []int{2},
		},
		{
			name: "Remove from 3 elements",
			head: []int{1, 2, 3},
			n:    2,
			want: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := _19_sliceToList(tt.head)
			got := _19_removeNthFromEnd2(head, tt.n)
			if !reflect.DeepEqual(_19_listToSlice(got), tt.want) {
				t.Errorf("_19_removeNthFromEnd() = %v, want %v", _19_listToSlice(got), tt.want)
			}
		})
	}
}
