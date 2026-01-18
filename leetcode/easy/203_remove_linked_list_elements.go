package easy

/*
Title: 移除链表元素
Difficulty: easy
Tags: 递归, 链表
Link: https://leetcode.cn/problems/remove-linked-list-elements/

Description:
给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。

示例 1：
输入：head = [1,2,6,3,4,5,6], val = 6
输出：[1,2,3,4,5]

示例 2：
输入：head = [], val = 1
输出：[]

示例 3：
输入：head = [7,7,7,7], val = 7
输出：[]

提示：
列表中的节点数目在范围 [0, 104] 内
1 <= Node.val <= 50
0 <= val <= 50
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type _203_ListNode struct {
	Val  int
	Next *_203_ListNode
}

func _203_removeElements(head *_203_ListNode, val int) *_203_ListNode {
	sentinel := &_203_ListNode{Next: head}
	head = sentinel
	for head != nil && head.Next != nil {
		if head.Next.Val == val {
			head.Next = head.Next.Next
			continue
		}
		head = head.Next
	}
	return sentinel.Next
}
