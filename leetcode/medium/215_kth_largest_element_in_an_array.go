package medium

/*
Title: 数组中的第K个最大元素
Difficulty: medium
Tags: 数组, 分治, 快速选择, 排序, 堆（优先队列）
Link: https://leetcode.cn/problems/kth-largest-element-in-an-array/

Description:
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1:
输入: [3,2,1,5,6,4], k = 2
输出: 5

示例 2:
输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4

提示：
	1 <= k <= nums.length <= 105
	-104 <= nums[i] <= 104
*/

// 快排思路
func _215_findKthLargest(nums []int, k int) int {
	n := len(nums)
	targetIdx := n - k
	lo, hi := 0, n-1
	for {
		pivot := _215_partition(nums, lo, hi)
		switch {
		case pivot < targetIdx:
			lo = pivot + 1
		case pivot > targetIdx:
			hi = pivot - 1
		default:
			return nums[pivot]
		}
	}
}

// 快排改进版
func _215_partition(nums []int, lo, hi int) int {
	if lo >= hi {
		return lo
	}
	pivot := nums[hi]
	i, j := lo, hi-1
	for {
		for ; i <= j && nums[i] < pivot; i++ {
		}
		for ; i <= j && nums[j] > pivot; j-- {
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	nums[i], nums[hi] = nums[hi], nums[i]
	return i
}

// 原快排，这种情况如果数组中大量重复元素，会导致运行时间超限
func _215_partition2(nums []int, lo, hi int) int {
	pivot := nums[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[hi] = nums[hi], nums[i]
	return i
}
