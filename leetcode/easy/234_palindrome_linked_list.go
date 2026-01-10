package easy

/*
Title: 回文链表
Difficulty: easy
Tags: 栈, 递归, 链表, 双指针
Link: https://leetcode.cn/problems/palindrome-linked-list/

Description:
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。

示例 1：
```
输入：head = [1,2,2,1]
输出：true
```

示例 2：
```
输入：head = [1,2]
输出：false
```

提示：
	链表中节点数目在范围[1, 105] 内
	0 <= Node.val <= 9

进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type _234_ListNode struct {
	Val  int
	Next *_234_ListNode
}

// 链表转数组
func _234_isPalindrome(head *_234_ListNode) bool {
	var vals []int
	for ; head != nil; head = head.Next {
		vals = append(vals, head.Val)
	}
	left, right := 0, len(vals)-1
	for left < right {
		if vals[left] != vals[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 递归
func _234_isPalindrome2(head *_234_ListNode) bool {
	frontNode := head
	var check func(node *_234_ListNode) bool
	check = func(node *_234_ListNode) bool {
		if node != nil {
			if !check(node.Next) {
				return false
			}
			if frontNode.Val != node.Val {
				return false
			}
			frontNode = frontNode.Next
		}
		return true
	}
	return check(head)
}

// 快慢指针找到中间节点
// 反转后半段链表
func _234_isPalindrome3(head *_234_ListNode) bool {
	if head == nil {
		return true
	}
	halfHeadNode := _234_findHalfHeadNode(head)
	reverseHeadNode := _234_reverseList(halfHeadNode.Next)
	p1, p2 := head, reverseHeadNode
	result := true
	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return result
}

func _234_reverseList(head *_234_ListNode) *_234_ListNode {
	var prev, cur *_234_ListNode = nil, head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

func _234_findHalfHeadNode(head *_234_ListNode) *_234_ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
