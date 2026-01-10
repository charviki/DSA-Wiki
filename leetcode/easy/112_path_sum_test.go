package easy

import "testing"

func Test_112_hasPathSum(t *testing.T) {
	tests := []struct {
		name      string
		root      *_112_TreeNode
		targetSum int
		want      bool
	}{
		{
			name: "Example 1",
			root: &_112_TreeNode{
				Val: 5,
				Left: &_112_TreeNode{
					Val: 4,
					Left: &_112_TreeNode{
						Val:   11,
						Left:  &_112_TreeNode{Val: 7},
						Right: &_112_TreeNode{Val: 2},
					},
				},
				Right: &_112_TreeNode{
					Val:  8,
					Left: &_112_TreeNode{Val: 13},
					Right: &_112_TreeNode{
						Val:   4,
						Right: &_112_TreeNode{Val: 1},
					},
				},
			},
			targetSum: 22,
			want:      true,
		},
		{
			name: "Example 2",
			root: &_112_TreeNode{
				Val:   1,
				Left:  &_112_TreeNode{Val: 2},
				Right: &_112_TreeNode{Val: 3},
			},
			targetSum: 5,
			want:      false,
		},
		{
			name:      "Example 3",
			root:      nil,
			targetSum: 0,
			want:      false,
		},
		{
			name:      "One node match",
			root:      &_112_TreeNode{Val: 5},
			targetSum: 5,
			want:      true,
		},
		{
			name:      "One node no match",
			root:      &_112_TreeNode{Val: 5},
			targetSum: 1,
			want:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _112_hasPathSum(tt.root, tt.targetSum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
