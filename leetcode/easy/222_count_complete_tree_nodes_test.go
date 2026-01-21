package easy

import "testing"

func Test_222_countNodes(t *testing.T) {
	tests := []struct {
		name string
		root *_222_TreeNode
		want int
	}{
		{"Example 1", _222_buildTree([]interface{}{1, 2, 3, 4, 5, 6}), 6},
		{"Example 2", _222_buildTree([]interface{}{}), 0},
		{"Example 3", _222_buildTree([]interface{}{1}), 1},
		{"Full tree level 2", _222_buildTree([]interface{}{1, 2, 3}), 3},
		{"Single node missing", _222_buildTree([]interface{}{1, 2, 3, 4, 5}), 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _222_countNodes(tt.root); got != tt.want {
				t.Errorf("_222_countNodes() = %v, want %v", got, tt.want)
			}
			if got := _222_countNodes2(tt.root); got != tt.want {
				t.Errorf("_222_countNodes() = %v, want %v", got, tt.want)
			}
			if got := _222_countNodes3(tt.root); got != tt.want {
				t.Errorf("_222_countNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Helper function to build a tree from a slice (level order)
// Use nil for null nodes
func _222_buildTree(nodes []interface{}) *_222_TreeNode {
	if len(nodes) == 0 {
		return nil
	}

	root := &_222_TreeNode{Val: nodes[0].(int)}
	queue := []*_222_TreeNode{root}
	i := 1

	for i < len(nodes) {
		curr := queue[0]
		queue = queue[1:]

		if i < len(nodes) && nodes[i] != nil {
			curr.Left = &_222_TreeNode{Val: nodes[i].(int)}
			queue = append(queue, curr.Left)
		}
		i++

		if i < len(nodes) && nodes[i] != nil {
			curr.Right = &_222_TreeNode{Val: nodes[i].(int)}
			queue = append(queue, curr.Right)
		}
		i++
	}

	return root
}
