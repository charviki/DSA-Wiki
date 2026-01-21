package easy

/*
Title: 存在重复元素 II
Difficulty: easy
Tags: 数组, 哈希表, 滑动窗口
Link: https://leetcode.cn/problems/contains-duplicate-ii/

Description:
给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。

示例 1：
输入：nums = [1,2,3,1], k = 3
输出：true

示例 2：
输入：nums = [1,0,1,1], k = 1
输出：true

示例 3：
输入：nums = [1,2,3,1,2,3], k = 2
输出：false

提示：
1 <= nums.length <= 10^5
-10^9 <= nums[i] <= 10^9
0 <= k <= 10^5
*/

func _219_containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int][]int)
	for i, v := range nums {
		idxes := m[v]
		for _, j := range idxes {
			if i-j <= k {
				return true
			}
		}
		m[v] = append(m[v], i)
	}
	return false
}
