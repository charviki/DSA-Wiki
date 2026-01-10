package medium

/*
Title: 删除链表的倒数第 N 个结点
Difficulty: medium
Tags: 链表, 双指针
Link: https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

Description:
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

示例 1：
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]

示例 2：
输入：head = [1], n = 1
输出：[]

示例 3：
输入：head = [1,2], n = 1
输出：[1]

提示：
	链表中结点的数目为 sz
	1 <= sz <= 30
	0 <= Node.val <= 100
	1 <= n <= sz

进阶：你能尝试使用一趟扫描实现吗？
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type _19_ListNode struct {
	Val  int
	Next *_19_ListNode
}

// 快慢指针
// 快指针比慢指针快 n 个节点，当快指针的下一个节点为 null 时，代表慢指针为倒数 n+1 个节点（即倒数 n 个节的前一个节点）
func _19_removeNthFromEnd(head *_19_ListNode, n int) *_19_ListNode {
	dummyNode := &_19_ListNode{Next: head}
	slow, fast := dummyNode, dummyNode
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummyNode.Next
}

func _19_removeNthFromEnd2(head *_19_ListNode, n int) *_19_ListNode {
	lists := []*_19_ListNode{{Next: head}}
	for head != nil {
		lists = append(lists, head)
		head = head.Next
	}
	pre := lists[len(lists)-n-1]
	pre.Next = pre.Next.Next
	return lists[0].Next
}
