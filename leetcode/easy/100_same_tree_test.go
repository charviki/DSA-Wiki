package easy

import "testing"

func Test_100_isSameTree(t *testing.T) {
    tests := []struct {
        name string
        p    *_100_TreeNode
        q    *_100_TreeNode
        want bool
    }{
        {
            name: "Example 1",
            p:    &_100_TreeNode{Val: 1, Left: &_100_TreeNode{Val: 2}, Right: &_100_TreeNode{Val: 3}},
            q:    &_100_TreeNode{Val: 1, Left: &_100_TreeNode{Val: 2}, Right: &_100_TreeNode{Val: 3}},
            want: true,
        },
        {
            name: "Example 2",
            p:    &_100_TreeNode{Val: 1, Left: &_100_TreeNode{Val: 2}},
            q:    &_100_TreeNode{Val: 1, Right: &_100_TreeNode{Val: 2}},
            want: false,
        },
        {
            name: "Example 3",
            p:    &_100_TreeNode{Val: 1, Left: &_100_TreeNode{Val: 2}, Right: &_100_TreeNode{Val: 1}},
            q:    &_100_TreeNode{Val: 1, Left: &_100_TreeNode{Val: 1}, Right: &_100_TreeNode{Val: 2}},
            want: false,
        },
        {
            name: "Both nil",
            p:    nil,
            q:    nil,
            want: true,
        },
        {
            name: "One nil",
            p:    &_100_TreeNode{Val: 1},
            q:    nil,
            want: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := _100_isSameTree(tt.p, tt.q); got != tt.want {
                t.Errorf("isSameTree() = %v, want %v", got, tt.want)
            }
        })
    }
}
