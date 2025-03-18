package medium

/*
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
返回容器可以储存的最大水量。
说明：你不能倾斜容器。

示例 1：

	输入：[1,8,6,2,5,4,8,3,7]
	输出：49
	解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

示例 2：

	输入：height = [1,1]
	输出：1

提示：
	n == height.length
	2 <= n <= 105
	0 <= height[i] <= 104

*/

func _11_maxArea(height []int) int {
	maxArea := 0
	for i := range height {
		for j := i + 1; j < len(height); j++ {
			h := min(height[j], height[i])
			area := (j - i) * h
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func _11_maxArea2(height []int) int {
	maxArea := 0
	for i, j := 0, len(height)-1; i < j; {
		h := height[i]
		if height[i] < height[j] {
			i++
		} else {
			h = height[j]
			j--
		}
		area := (j - i + 1) * h
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}
