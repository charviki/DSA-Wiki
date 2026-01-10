package easy

/*
Title: 杨辉三角 II
Difficulty: easy
Tags: 数组, 动态规划
Link: https://leetcode.cn/problems/pascals-triangle-ii/

Description:
给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。

示例 1:
```
输入: rowIndex = 3
输出: [1,3,3,1]
```

示例 2:
```
输入: rowIndex = 0
输出: [1]
```

示例 3:
```
输入: rowIndex = 1
输出: [1,1]
```

提示:
	0 <= rowIndex <= 33

进阶：
你可以优化你的算法到 O(rowIndex) 空间复杂度吗？
*/

func _119_getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	res[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			res[j] += res[j-1]
		}
	}
	return res
}

func _119_getRow2(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	res[0] = 1
	for i := 1; i <= rowIndex; i++ {
		res[i] = res[i-1] * (rowIndex - i + 1) / i
	}
	return res
}
