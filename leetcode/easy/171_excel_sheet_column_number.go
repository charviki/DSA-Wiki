package easy

import "math"

/*
Title: Excel 表列序号
Difficulty: easy
Tags: 数学, 字符串
Link: https://leetcode.cn/problems/excel-sheet-column-number/

Description:
给你一个字符串 columnTitle ，表示 Excel 表格中的列名称。返回 该列名称对应的列序号 。

例如：

A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...

示例 1:
```
输入: columnTitle = "A"
输出: 1
```

示例 2:
```
输入: columnTitle = "AB"
输出: 28
```

示例 3:
```
输入: columnTitle = "ZY"
输出: 701
```

提示：
	1 <= columnTitle.length <= 7
	columnTitle 仅由大写英文组成
	columnTitle 在范围 ["A", "FXSHRXW"] 内
*/

func _171_titleToNumber(columnTitle string) int {
	res := 0
	for i := range len(columnTitle) {
		res = res*26 + int(columnTitle[i]-64)
	}
	return res
}

func _171_titleToNumber2(columnTitle string) int {
	strLen := len(columnTitle)
	if strLen <= 0 {
		return 0
	}
	res := 0
	for i := 2; i <= strLen; i++ {
		res += int(columnTitle[strLen-i]-64) * int(math.Pow(26, float64(i-1)))
	}
	res += int(columnTitle[strLen-1] - 64)
	return res
}
