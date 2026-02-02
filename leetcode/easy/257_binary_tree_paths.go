package easy

import "fmt"

/*
Title: 二叉树的所有路径
Difficulty: easy
Tags: 树, 深度优先搜索, 字符串, 回溯, 二叉树
Link: https://leetcode.cn/problems/binary-tree-paths/

Description:
给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。

叶子节点 是指没有子节点的节点。

示例 1：
输入：root = [1,2,3,null,5]
输出：["1->2->5","1->3"]

示例 2：
输入：root = [1]
输出：["1"]

提示：
	树中节点的数目在范围 [1, 100] 内
	-100 <= Node.val <= 100
*/

/**
 * Definition for a binary tree node.
 */
type _257_TreeNode struct {
	Val   int
	Left  *_257_TreeNode
	Right *_257_TreeNode
}

func _257_binaryTreePaths(root *_257_TreeNode) []string {
	res := []string{}
	if root == nil {
		return res
	}
	var search func(node *_257_TreeNode, preStr string)
	search = func(node *_257_TreeNode, preStr string) {
		if node == nil {
			return
		}
		if preStr == "" {
			preStr += fmt.Sprintf("%d", root.Val)
		} else {
			preStr += fmt.Sprintf("->%d", node.Val)
		}

		if node.Left == nil && node.Right == nil {
			res = append(res, preStr)
			return
		}
		search(node.Left, preStr)
		search(node.Right, preStr)
	}
	search(root, "")
	return res
}
