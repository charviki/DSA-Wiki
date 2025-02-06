package medium

import "sort"

/*
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

示例 1：

	输入：nums = [1,1,2]
	输出：
		[[1,1,2],
		[1,2,1],
		[2,1,1]]

示例 2：

	输入：nums = [1,2,3]
	输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

提示：

	1 <= nums.length <= 8
	-10 <= nums[i] <= 10
*/
func _47_permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	perm := []int{}
	ans := [][]int{}
	vis := make([]bool, n)
	var backtrace func(int)
	backtrace = func(i int) {
		if i == n {
			ans = append(ans, append([]int{}, perm...))
			return
		}
		for j, v := range nums {
			if vis[j] || j > 0 && !vis[j-1] && v == nums[j-1] {
				continue
			}
			perm = append(perm, v)
			vis[j] = true
			backtrace(i + 1)
			vis[j] = false
			perm = perm[:len(perm)-1]
		}
	}
	backtrace(0)
	return ans
}
