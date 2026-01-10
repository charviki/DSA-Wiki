package easy

import (
	"testing"
)

func Test_108_sortedArrayToBST(t *testing.T) {
	// Helper function to calculate height of the tree
	var height func(root *_108_TreeNode) int
	height = func(root *_108_TreeNode) int {
		if root == nil {
			return 0
		}
		l := height(root.Left)
		r := height(root.Right)
		if l > r {
			return l + 1
		}
		return r + 1
	}

	// Helper function to check if it's a valid BST
	var isValidBST func(root *_108_TreeNode, min, max *int) bool
	isValidBST = func(root *_108_TreeNode, min, max *int) bool {
		if root == nil {
			return true
		}
		if (min != nil && root.Val <= *min) || (max != nil && root.Val >= *max) {
			return false
		}
		val := root.Val
		return isValidBST(root.Left, min, &val) && isValidBST(root.Right, &val, max)
	}

	// Helper to check if tree is balanced
	var isBalanced func(root *_108_TreeNode) bool
	isBalanced = func(root *_108_TreeNode) bool {
		if root == nil {
			return true
		}
		lh := height(root.Left)
		rh := height(root.Right)
		if lh-rh > 1 || rh-lh > 1 {
			return false
		}
		return isBalanced(root.Left) && isBalanced(root.Right)
	}

	tests := []struct {
		name string
		nums []int
	}{
		{
			name: "Example 1",
			nums: []int{-10, -3, 0, 5, 9},
		},
		{
			name: "Example 2",
			nums: []int{1, 3},
		},
		{
			name: "Single element",
			nums: []int{1},
		},
		{
			name: "Two elements",
			nums: []int{1, 2},
		},
		{
			name: "Three elements",
			nums: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := _108_sortedArrayToBST(tt.nums)

			// Check if it's a valid BST
			if !isValidBST(got, nil, nil) {
				t.Errorf("sortedArrayToBST() returned invalid BST")
			}

			// Check if it's balanced
			if !isBalanced(got) {
				t.Errorf("sortedArrayToBST() returned unbalanced tree")
			}
		})
	}
}
