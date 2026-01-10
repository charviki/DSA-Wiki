package easy

/*
Title: 平衡二叉树
Difficulty: easy
Tags: 树, 深度优先搜索, 二叉树
Link: https://leetcode.cn/problems/balanced-binary-tree/

Description:
给定一个二叉树，判断它是否是 平衡二叉树

示例 1：
```
输入：root = [3,9,20,null,null,15,7]
输出：true
```

示例 2：
```
输入：root = [1,2,2,3,3,null,null,4,4]
输出：false
```

示例 3：
```
输入：root = []
输出：true
```

提示：
	树中的节点数在范围 [0, 5000] 内
	-104 <= Node.val <= 104
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type _110_TreeNode struct {
	Val   int
	Left  *_110_TreeNode
	Right *_110_TreeNode
}

func _110_isBalanced(root *_110_TreeNode) bool {
	_, isBalanced := _110_getHeight(root)
	return isBalanced
}

func _110_getHeight(root *_110_TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftHeight, isLeftBalanced := _110_getHeight(root.Left)
	if !isLeftBalanced {
		return 0, false
	}
	rightHeight, isRightBalanced := _110_getHeight(root.Right)
	if !isRightBalanced {
		return 0, false
	}
	diff := leftHeight - rightHeight
	if diff < -1 || diff > 1 {
		return 0, false
	}
	return 1 + max(leftHeight, rightHeight), true
}
