package medium

import (
	"reflect"
	"testing"
)

func _92_sliceToList(nums []int) *_92_ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &_92_ListNode{Val: nums[0]}
	curr := head
	for i := 1; i < len(nums); i++ {
		curr.Next = &_92_ListNode{Val: nums[i]}
		curr = curr.Next
	}
	return head
}

func _92_listToSlice(head *_92_ListNode) []int {
	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return nums
}

func Test_92_reverseBetween(t *testing.T) {
	tests := []struct {
		name  string
		head  []int
		left  int
		right int
		want  []int
	}{
		{
			name:  "Example 1",
			head:  []int{1, 2, 3, 4, 5},
			left:  2,
			right: 4,
			want:  []int{1, 4, 3, 2, 5},
		},
		{
			name:  "Example 2",
			head:  []int{5},
			left:  1,
			right: 1,
			want:  []int{5},
		},
		{
			name:  "Reverse full list",
			head:  []int{1, 2, 3},
			left:  1,
			right: 3,
			want:  []int{3, 2, 1},
		},
		{
			name:  "Reverse head only",
			head:  []int{1, 2, 3},
			left:  1,
			right: 1,
			want:  []int{1, 2, 3},
		},
		{
			name:  "Reverse tail",
			head:  []int{1, 2, 3},
			left:  2,
			right: 3,
			want:  []int{1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := _92_sliceToList(tt.head)
			got := _92_reverseBetween(head, tt.left, tt.right)
			if !reflect.DeepEqual(_92_listToSlice(got), tt.want) {
				t.Errorf("_92_reverseBetween() = %v, want %v", _92_listToSlice(got), tt.want)
			}
		})
	}
}
