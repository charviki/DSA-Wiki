package medium

/*
Title: 反转链表 II
Difficulty: medium
Tags: 链表
Link: https://leetcode.cn/problems/reverse-linked-list-ii/

Description:
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。

示例 1：
输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]

示例 2：
输入：head = [5], left = 1, right = 1
输出：[5]

提示：
	链表中节点数目为 n
	1 <= n <= 500
	-500 <= Node.val <= 500
	1 <= left <= right <= n

进阶： 你可以使用一趟扫描完成反转吗？
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type _92_ListNode struct {
	Val  int
	Next *_92_ListNode
}

func _92_reverseBetween(head *_92_ListNode, left int, right int) *_92_ListNode {
	dummyNode := &_92_ListNode{Next: head}
	pre := dummyNode
	// 先找到反转前一个结点
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		// 重点在这，将 next 插入到 pre 后面
		next.Next = pre.Next
		pre.Next = next
	}
	return dummyNode.Next
}
