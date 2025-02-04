package sort

// 插入排序
func Insertion(nums []int) {
	for i := 1; i < len(nums); i++ {
		val := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > val {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = val
	}
}
