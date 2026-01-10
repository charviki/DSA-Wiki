package medium

/*
Title: 最长连续序列
Difficulty: medium
Tags: 并查集, 数组, 哈希表
Link: https://leetcode.cn/problems/longest-consecutive-sequence/

Description:
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1：
输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。

示例 2：
输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9

示例 3：
输入：nums = [1,0,1,2]
输出：3

提示：
	0 <= nums.length <= 105
	-109 <= nums[i] <= 109
*/

// 外层遍历 n 次
// 内层关键在于 set[num-1] 是否存在的过滤，即从序列的起点开始遍历
// 内层总遍历次数不会超过 n
func _128_longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, v := range nums {
		set[v] = true
	}
	maxCnt := 0
	for num := range set {
		if !set[num-1] {
			curCnt := 1
			for num = num + 1; set[num]; num++ {
				curCnt++
			}
			maxCnt = max(curCnt, maxCnt)
		}
	}
	return maxCnt
}
