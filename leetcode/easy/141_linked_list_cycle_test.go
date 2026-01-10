package easy

import "testing"

func Test_141_hasCycle(t *testing.T) {
	// Helper function to create a linked list with a cycle
	createCycleList := func(vals []int, pos int) *_141_ListNode {
		if len(vals) == 0 {
			return nil
		}
		head := &_141_ListNode{Val: vals[0]}
		nodes := []*_141_ListNode{head}
		curr := head
		for i := 1; i < len(vals); i++ {
			curr.Next = &_141_ListNode{Val: vals[i]}
			curr = curr.Next
			nodes = append(nodes, curr)
		}
		if pos >= 0 && pos < len(nodes) {
			curr.Next = nodes[pos]
		}
		return head
	}

	tests := []struct {
		name string
		head *_141_ListNode
		want bool
	}{
		{
			name: "Example 1",
			head: createCycleList([]int{3, 2, 0, -4}, 1),
			want: true,
		},
		{
			name: "Example 2",
			head: createCycleList([]int{1, 2}, 0),
			want: true,
		},
		{
			name: "Example 3",
			head: createCycleList([]int{1}, -1),
			want: false,
		},
		{
			name: "Empty list",
			head: createCycleList([]int{}, -1),
			want: false,
		},
		{
			name: "Single node no cycle",
			head: createCycleList([]int{1}, -1),
			want: false,
		},
		{
			name: "Two nodes no cycle",
			head: createCycleList([]int{1, 2}, -1),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _141_hasCycle(tt.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
