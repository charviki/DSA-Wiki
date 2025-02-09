package medium

/*
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
说明：

	为什么返回数值是整数，但输出的答案是数组呢？
	请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
	你可以想象内部操作如下:
		// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
		int len = removeDuplicates(nums);
		// 在函数里修改输入数组对于调用者是可见的。
		// 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
		for (int i = 0; i < len; i++) {
			print(nums[i]);
		}

示例 1：

	输入：nums = [1,1,1,2,2,3]
	输出：5, nums = [1,1,2,2,3]
	解释：函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3。 不需要考虑数组中超出新长度后面的元素。

示例 2：

	输入：nums = [0,0,1,1,1,1,2,3,3]
	输出：7, nums = [0,0,1,1,2,3,3]
	解释：函数应返回新长度 length = 7, 并且原数组的前七个元素被修改为 0, 0, 1, 1, 2, 3, 3。不需要考虑数组中超出新长度后面的元素。

提示：

	1 <= nums.length <= 3 * 104
	-104 <= nums[i] <= 104
	nums 已按升序排列
*/

// _80_removeDuplicates
// 思路，两个指针，计算相同值得重复个数，若个数大于2，则将后面的元素往前移动，并更新endIdx
func _80_removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	endIdx := len(nums) - 1
	for i := endIdx; i > 0; {
		j := i - 1
		for ; j >= 0 && nums[j] == nums[i]; j-- {
		}
		if gap := i - j - 2; gap > 0 {
			for k := j + 1; k <= endIdx-gap; k++ {
				nums[k] = nums[k+gap]
			}
			endIdx -= gap
		}
		i = j
	}
	return endIdx + 1
}

// 快慢指针
// 思路：
// slow - 1 为保留数组最后一个元素，fast 为当前检查的元素
// 若 fast 与 slow - 2（最多重复个数）不等，则需交换元素，移动 slow 和 fast
// 否则，继续移动 fast
func _80_removeDuplicates1(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	slow, fast := 2, 2
	for fast < n {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
