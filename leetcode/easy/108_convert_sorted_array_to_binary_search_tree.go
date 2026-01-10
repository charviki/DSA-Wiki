package easy

/*
Title: 将有序数组转换为二叉搜索树
Difficulty: easy
Tags: 树, 二叉搜索树, 数组, 分治, 二叉树
Link: https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/

Description:
给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 平衡 二叉搜索树。

示例 1：
```
输入：nums = [-10,-3,0,5,9]
输出：[0,-3,9,-10,null,5]
解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：
```

示例 2：
```
输入：nums = [1,3]
输出：[3,1]
解释：[1,null,3] 和 [3,1] 都是高度平衡二叉搜索树。
```

提示：
	1 <= nums.length <= 104
	-104 <= nums[i] <= 104
	nums 按 严格递增 顺序排列
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type _108_TreeNode struct {
	Val   int
	Left  *_108_TreeNode
	Right *_108_TreeNode
}

func _108_sortedArrayToBST(nums []int) *_108_TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &_108_TreeNode{Val: nums[mid]}
	root.Left = _108_sortedArrayToBST(nums[:mid])
	root.Right = _108_sortedArrayToBST(nums[mid+1:])
	return root
}
