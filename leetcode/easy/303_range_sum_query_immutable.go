package easy

/*
Title: 区域和检索 - 数组不可变
Difficulty: Easy
Tags: 设计, 数组, 前缀和
Link: https://leetcode.cn/problems/range-sum-query-immutable/

Description:
给定一个整数数组  nums，处理以下类型的多个查询:
1. 计算索引 left 和 right （包含 left 和 right）之间的 nums 元素的 和 ，其中 left <= right

实现 NumArray 类：
- NumArray(int[] nums) 使用数组 nums 初始化对象
- int sumRange(int left, int right) 返回数组 nums 中索引 left 和 right 之间的元素的 总和 ，包含 left 和 right 两点（也就是 nums[left] + nums[left + 1] + ... + nums[right] )

示例 1：
输入：
["NumArray", "sumRange", "sumRange", "sumRange"]
[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
输出：
[null, 1, -1, -3]

解释：
NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
numArray.sumRange(0, 2); // return 1 ((-2) + 0 + 3)
numArray.sumRange(2, 5); // return -1 (3 + (-5) + 2 + (-1))
numArray.sumRange(0, 5); // return -3 ((-2) + 0 + 3 + (-5) + 2 + (-1))

提示：
1 <= nums.length <= 104
-105 <= nums[i] <= 105
0 <= left <= right < nums.length
最多调用 104 次 sumRange 方法
*/

type _303_NumArray struct {
	sum []int
}

func _303_Constructor(nums []int) _303_NumArray {
	sum := make([]int, len(nums)+1)
	for i, v := range nums {
		sum[i+1] = sum[i] + v
	}
	return _303_NumArray{
		sum: sum,
	}
}

func (this *_303_NumArray) SumRange(left int, right int) int {
	return this.sum[right+1] - this.sum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
