package easy

import "slices"

/*
Title: Excel 表列名称
Difficulty: easy
Tags: 数学, 字符串
Link: https://leetcode.cn/problems/excel-sheet-column-title/

Description:
给你一个整数 columnNumber ，返回它在 Excel 表中相对应的列名称。

例如：

A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...

示例 1：
```
输入：columnNumber = 1
输出："A"
```

示例 2：
```
输入：columnNumber = 28
输出："AB"
```

示例 3：
```
输入：columnNumber = 701
输出："ZY"
```

示例 4：
```
输入：columnNumber = 2147483647
输出："FXSHRXW"
```

提示：
	1 <= columnNumber <= 231 - 1
*/

func _168_convertToTitle(columnNumber int) string {
	var res []byte
	for columnNumber > 0 {
		columnNumber--
		res = append(res, byte(columnNumber%26)+'A')
		columnNumber /= 26
	}
	slices.Reverse(res)
	return string(res)
}
