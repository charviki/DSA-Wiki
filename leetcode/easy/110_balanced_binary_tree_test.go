package easy

import "testing"

func Test_110_isBalanced(t *testing.T) {
    tests := []struct {
        name string
        root *_110_TreeNode
        want bool
    }{
        {
            name: "Example 1",
            root: &_110_TreeNode{
                Val: 3,
                Left: &_110_TreeNode{Val: 9},
                Right: &_110_TreeNode{
                    Val:   20,
                    Left:  &_110_TreeNode{Val: 15},
                    Right: &_110_TreeNode{Val: 7},
                },
            },
            want: true,
        },
        {
            name: "Example 2",
            root: &_110_TreeNode{
                Val: 1,
                Left: &_110_TreeNode{
                    Val: 2,
                    Left: &_110_TreeNode{
                        Val:   3,
                        Left:  &_110_TreeNode{Val: 4},
                        Right: &_110_TreeNode{Val: 4},
                    },
                    Right: &_110_TreeNode{Val: 3},
                },
                Right: &_110_TreeNode{Val: 2},
            },
            want: false,
        },
        {
            name: "Example 3",
            root: nil,
            want: true,
        },
        {
            name: "Single node",
            root: &_110_TreeNode{Val: 1},
            want: true,
        },
        {
            name: "Left heavy",
            root: &_110_TreeNode{
                Val: 1,
                Left: &_110_TreeNode{
                    Val: 2,
                    Left: &_110_TreeNode{Val: 3},
                },
            },
            want: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := _110_isBalanced(tt.root); got != tt.want {
                t.Errorf("isBalanced() = %v, want %v", got, tt.want)
            }
        })
    }
}
