package medium

import (
	"reflect"
	"testing"
)

func _24_sliceToList(nums []int) *_24_ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &_24_ListNode{Val: nums[0]}
	curr := head
	for i := 1; i < len(nums); i++ {
		curr.Next = &_24_ListNode{Val: nums[i]}
		curr = curr.Next
	}
	return head
}

func _24_listToSlice(head *_24_ListNode) []int {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return nums
}

func Test_24_swapPairs(t *testing.T) {
	tests := []struct {
		name string
		head []int
		want []int
	}{
		{
			name: "Example 1",
			head: []int{1, 2, 3, 4},
			want: []int{2, 1, 4, 3},
		},
		{
			name: "Example 2",
			head: []int{},
			want: []int{},
		},
		{
			name: "Example 3",
			head: []int{1},
			want: []int{1},
		},
		{
			name: "Odd number of elements",
			head: []int{1, 2, 3},
			want: []int{2, 1, 3},
		},
		{
			name: "Two elements",
			head: []int{1, 2},
			want: []int{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := _24_sliceToList(tt.head)
			got := _24_swapPairs(head)
			if !reflect.DeepEqual(_24_listToSlice(got), tt.want) {
				t.Errorf("_24_swapPairs() = %v, want %v", _24_listToSlice(got), tt.want)
			}
		})
	}
}

func Test_24_swapPairs2(t *testing.T) {
	tests := []struct {
		name string
		head []int
		want []int
	}{
		{
			name: "Example 1",
			head: []int{1, 2, 3, 4},
			want: []int{2, 1, 4, 3},
		},
		{
			name: "Example 2",
			head: []int{},
			want: []int{},
		},
		{
			name: "Example 3",
			head: []int{1},
			want: []int{1},
		},
		{
			name: "Odd number of elements",
			head: []int{1, 2, 3},
			want: []int{2, 1, 3},
		},
		{
			name: "Two elements",
			head: []int{1, 2},
			want: []int{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := _24_sliceToList(tt.head)
			got := _24_swapPairs2(head)
			if !reflect.DeepEqual(_24_listToSlice(got), tt.want) {
				t.Errorf("_24_swapPairs() = %v, want %v", _24_listToSlice(got), tt.want)
			}
		})
	}
}
