package easy

/*
Title: 相同的树
Difficulty: easy
Tags: 树, 深度优先搜索, 广度优先搜索, 二叉树
Link: https://leetcode.cn/problems/same-tree/

Description:
给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。

如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

示例 1：
```
输入：p = [1,2,3], q = [1,2,3]
输出：true
```

示例 2：
```
输入：p = [1,2], q = [1,null,2]
输出：false
```

示例 3：
```
输入：p = [1,2,1], q = [1,1,2]
输出：false
```

提示：
	两棵树上的节点数目都在范围 [0, 100] 内
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
type _100_TreeNode struct {
	Val   int
	Left  *_100_TreeNode
	Right *_100_TreeNode
}

func _100_isSameTree(p *_100_TreeNode, q *_100_TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return _100_isSameTree(p.Left, q.Left) && _100_isSameTree(p.Right, q.Right)
}
