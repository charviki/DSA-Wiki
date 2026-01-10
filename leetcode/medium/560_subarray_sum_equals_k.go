package medium

/*
Title: 和为 K 的子数组
Difficulty: medium
Tags: 数组, 哈希表, 前缀和
Link: https://leetcode.cn/problems/subarray-sum-equals-k/

Description:
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。

子数组是数组中元素的连续非空序列。

示例 1：
输入：nums = [1,1,1], k = 2
输出：2

示例 2：
输入：nums = [1,2,3], k = 3
输出：2

提示：
	1 <= nums.length <= 2 * 104
	-1000 <= nums[i] <= 1000
	-107 <= k <= 107
*/

// 记录前 i 个元素的和出现的次数
// 和为 k 的子数组即 前 i 个元素之和减去前j(0-n之间)个元素之和，即 j 到 i 之间的子数组和为 k
func _560_subarraySum(nums []int, k int) int {
	var (
		cnt = 0
		pre = 0
		m   = map[int]int{
			0: 1,
		}
	)
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		cnt += m[pre-k]
		m[pre]++
	}

	return cnt
}
