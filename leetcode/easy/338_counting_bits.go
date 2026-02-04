package easy

/*
Title: 比特位计数
Difficulty: Easy
Tags: 位运算, 动态规划
Link: https://leetcode.cn/problems/counting-bits/

Description:
给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。

示例 1：
输入：n = 2
输出：[0,1,1]
解释：
0 --> 0
1 --> 1
2 --> 10

示例 2：
输入：n = 5
输出：[0,1,1,2,1,2]
解释：
0 --> 0
1 --> 1
2 --> 10
3 --> 11
4 --> 100
5 --> 101

提示：
0 <= n <= 105

进阶：
- 很容易就能实现时间复杂度为 O(n log n) 的解决方案，你可以在线性时间复杂度 O(n) 内用一趟扫描解决此问题吗？
- 你能不使用任何内置函数解决此问题吗？（如，C++ 中的 __builtin_popcount ）
*/

func _338_countBits(n int) []int {
	bits := make([]int, n+1)
	for i := 0; i <= n; i++ {
		cnt := 0
		for j := i; j > 0; {
			if j&0x01 == 1 {
				cnt++
			}
			j >>= 1
		}
		bits[i] = cnt
	}
	return bits
}

func _338_countBits2(n int) []int {
	bits := make([]int, n+1)
	highBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		bits[i] = bits[i-highBit] + 1
	}
	return bits
}

func _338_countBits3(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bits[i] = i%2 + bits[i/2]
	}
	return bits
}
