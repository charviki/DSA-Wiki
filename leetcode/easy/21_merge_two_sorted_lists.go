package easy

/*
Title: 合并两个有序链表
Difficulty: easy
Tags: 递归, 链表
Link: https://leetcode.cn/problems/merge-two-sorted-lists/

Description:
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例 1：
```
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
```

示例 2：
```
输入：l1 = [], l2 = []
输出：[]
```

示例 3：
```
输入：l1 = [], l2 = [0]
输出：[0]
```

提示：
	两个链表的节点数目范围是 [0, 50]
	-100 <= Node.val <= 100
	l1 和 l2 均按 非递减顺序 排列
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type _21_ListNode struct {
	Val  int
	Next *_21_ListNode
}

func _21_mergeTwoLists(list1 *_21_ListNode, list2 *_21_ListNode) *_21_ListNode {
	mergedList := &_21_ListNode{}
	curNode := mergedList
	for list1 != nil && list2 != nil {
		var node *_21_ListNode
		if list1.Val < list2.Val {
			node = list1
			list1 = list1.Next
		} else {
			node = list2
			list2 = list2.Next
		}
		node.Next = nil
		curNode.Next = node
		curNode = curNode.Next
	}

	if list1 != nil {
		curNode.Next = list1
	}
	if list2 != nil {
		curNode.Next = list2
	}

	return mergedList.Next
}

func _21_mergeTwoLists2(list1 *_21_ListNode, list2 *_21_ListNode) *_21_ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = _21_mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = _21_mergeTwoLists(list1, list2.Next)
	return list2
}
