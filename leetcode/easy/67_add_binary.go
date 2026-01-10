package easy

import "strings"

/*
Title: 二进制求和
Difficulty: easy
Tags: 位运算, 数学, 字符串, 模拟
Link: https://leetcode.cn/problems/add-binary/

Description:
给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。

示例 1：
```
输入:a = "11", b = "1"
输出："100"
```

示例 2：
```
输入：a = "1010", b = "1011"
输出："10101"
```

提示：
	1 <= a.length, b.length <= 104
	a 和 b 仅由字符 '0' 或 '1' 组成
	字符串如果不是 "0" ，就不含前导零
*/

func _67_addBinary(a string, b string) string {
	l := max(len(a), len(b))
	a = strings.Repeat("0", l-len(a)) + a
	b = strings.Repeat("0", l-len(b)) + b
	var more byte = 0
	res := ""
	for i := l - 1; i >= 0; i-- {
		more = more + a[i] + b[i] - ('0' + '0')
		res = string(more%2+'0') + res
		more = more / 2
	}
	if more > 0 {
		res = "1" + res
	}
	return res
}
