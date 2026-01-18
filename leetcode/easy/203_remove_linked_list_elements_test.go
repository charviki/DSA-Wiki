package easy

import (
	"reflect"
	"testing"
)

func Test_203_removeElements(t *testing.T) {
	tests := []struct {
		name string
		head []int
		val  int
		want []int
	}{
		{"Example 1", []int{1, 2, 6, 3, 4, 5, 6}, 6, []int{1, 2, 3, 4, 5}},
		{"Example 2", []int{}, 1, []int{}},
		{"Example 3", []int{7, 7, 7, 7}, 7, []int{}},
		{"All elements match", []int{1, 1, 1}, 1, []int{}},
		{"No elements match", []int{1, 2, 3}, 4, []int{1, 2, 3}},
		{"Head match", []int{1, 2, 3}, 1, []int{2, 3}},
		{"Tail match", []int{1, 2, 3}, 3, []int{1, 2}},
		{"Repeat", []int{7, 7, 7, 7}, 7, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := _203_sliceToList(tt.head)
			gotList := _203_removeElements(head, tt.val)
			got := _203_listToSlice(gotList)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("_203_removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func _203_sliceToList(nums []int) *_203_ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &_203_ListNode{Val: nums[0]}
	curr := head
	for _, num := range nums[1:] {
		curr.Next = &_203_ListNode{Val: num}
		curr = curr.Next
	}
	return head
}

func _203_listToSlice(head *_203_ListNode) []int {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return nums
}
