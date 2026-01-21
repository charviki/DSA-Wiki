package easy

import "math"

/*
Title: 完全二叉树的节点个数
Difficulty: easy
Tags: 位运算, 树, 二分查找, 二叉树
Link: https://leetcode.cn/problems/count-complete-tree-nodes/

Description:
给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。

完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层（从第 0 层开始），则该层包含 1~ 2^h 个节点。

示例 1：
输入：root = [1,2,3,4,5,6]
输出：6

示例 2：
输入：root = []
输出：0

示例 3：
输入：root = [1]
输出：1

提示：
树中节点的数目范围是[0, 5 * 10^4]
0 <= Node.val <= 5 * 10^4
题目数据保证输入的树是 完全二叉树

进阶：遍历树来统计节点是一种时间复杂度为 O(n) 的简单解决方案。你可以设计一个更快的算法吗？
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type _222_TreeNode struct {
	Val   int
	Left  *_222_TreeNode
	Right *_222_TreeNode
}

func _222_countNodes(root *_222_TreeNode) int {
	cnt := 0
	q := []*_222_TreeNode{root}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		if n == nil {
			continue
		}
		q = append(q, n.Left)
		q = append(q, n.Right)
		cnt++
	}
	return cnt
}

func _222_countNodes2(root *_222_TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + _222_countNodes2(root.Left) + _222_countNodes2(root.Right)
}

// 一棵完全二叉树的两棵子树，至少有一棵是满二叉树
// O(logN*logN)
func _222_countNodes3(root *_222_TreeNode) int {
	leftNode, rightNode := root, root
	lh, rh := 0, 0
	for leftNode != nil {
		leftNode = leftNode.Left
		lh++
	}
	for rightNode != nil {
		rightNode = rightNode.Right
		rh++
	}
	if lh == rh {
		return int(math.Pow(2, float64(lh))) - 1
	}
	return 1 + _222_countNodes3(root.Left) + _222_countNodes3(root.Right)
}
