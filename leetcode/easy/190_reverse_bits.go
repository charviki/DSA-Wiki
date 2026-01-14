package easy

import (
	"slices"
	"strconv"
)

/*
Title: 颠倒二进制位
Difficulty: easy
Tags: 位运算, 分治
Link: https://leetcode.cn/problems/reverse-bits/

Description:
颠倒给定的 32 位有符号整数的二进制位。

示例 1：
输入：n = 43261596
输出：964176192
解释：
			整数
			二进制

			43261596
			00000010100101000001111010011100

			964176192
			00111001011110000010100101000000

示例 2：
输入：n = 2147483644
输出：1073741822
解释：
			整数
			二进制

			2147483644
			01111111111111111111111111111100

			1073741822
			00111111111111111111111111111110

提示：
	0 <= n <= 231 - 2
	n 为偶数

进阶: 如果多次调用这个函数，你将如何优化你的算法？
*/

// 每次取最低位求和后整体左移
func _190_reverseBits(n int) int {
	val := n & 1
	for i := 1; i < 32; i++ {
		val <<= 1
		n >>= 1
		val |= n & 1
	}
	return val
}

// 字符串颠倒
func _190_reverseBits2(n int) int {
	bts := []byte(strconv.FormatInt(int64(n), 2))
	for len(bts) < 32 {
		bts = append([]byte{'0'}, bts...)
	}
	slices.Reverse(bts)
	res, _ := strconv.ParseInt(string(bts), 2, 32)
	return int(res)
}
