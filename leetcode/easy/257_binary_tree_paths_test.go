package easy

import (
	"reflect"
	"sort"
	"testing"
)

func Test_257_binaryTreePaths(t *testing.T) {
	// Helper to create tree nodes more easily
	node := func(val int) *_257_TreeNode {
		return &_257_TreeNode{Val: val}
	}

	tests := []struct {
		name string
		root *_257_TreeNode
		want []string
	}{
		{
			name: "Example 1",
			root: &_257_TreeNode{
				Val: 1,
				Left: &_257_TreeNode{
					Val:   2,
					Right: node(5),
				},
				Right: node(3),
			},
			want: []string{"1->2->5", "1->3"},
		},
		{
			name: "Example 2",
			root: node(1),
			want: []string{"1"},
		},
		{
			name: "Left Path Only",
			root: &_257_TreeNode{
				Val: 1,
				Left: &_257_TreeNode{
					Val: 2,
					Left: node(3),
				},
			},
			want: []string{"1->2->3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := _257_binaryTreePaths(tt.root)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("_257_binaryTreePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
