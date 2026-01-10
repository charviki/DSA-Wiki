package medium

/*
Title: 两两交换链表中的节点
Difficulty: medium
Tags: 递归, 链表
Link: https://leetcode.cn/problems/swap-nodes-in-pairs/

Description:
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

示例 1：
输入：head = [1,2,3,4]
输出：[2,1,4,3]

示例 2：
输入：head = []
输出：[]

示例 3：
输入：head = [1]
输出：[1]

提示：
	链表中节点的数目在范围 [0, 100] 内
	0 <= Node.val <= 100
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type _24_ListNode struct {
	Val  int
	Next *_24_ListNode
}

func _24_swapPairs(head *_24_ListNode) *_24_ListNode {
	sentinel := &_24_ListNode{Next: head}
	pre, cur := sentinel, head
	for cur != nil && cur.Next != nil {
		temp := cur.Next.Next
		pre.Next = cur.Next
		cur.Next.Next = cur
		cur.Next = temp
		pre, cur = cur, temp
	}
	return sentinel.Next
}

func _24_swapPairs2(head *_24_ListNode) *_24_ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = _24_swapPairs2(newHead.Next)
	newHead.Next = head
	return newHead
}
