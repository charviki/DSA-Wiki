package easy

import (
	"strings"
)

/*
Title: 反转字符串中的元音字母
Difficulty: easy
Tags: 双指针, 字符串
Link: https://leetcode.cn/problems/reverse-vowels-of-a-string/

Description:
给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。

元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现不止一次。

示例 1：
输入：s = "IceCreAm"
输出："AceCreIm"

解释：
s 中的元音是 ['I', 'e', 'e', 'A']。反转这些元音，s 变为 "AceCreIm".

示例 2：
输入：s = "leetcode"
输出："leotcede"

提示：
1 <= s.length <= 3 * 10^5
s 由 可打印的 ASCII 字符组成
*/

func _345_reverseVowels(s string) string {
	m := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	i, j := 0, len(s)-1
	sArr := []byte(s)
	lowArr := []byte(strings.ToLower(s))
	for i < j {

		for ; i < j && i < len(s)-1 && !m[lowArr[i]]; i++ {
		}
		for ; i < j && j > 0 && !m[lowArr[j]]; j-- {
		}
		if i >= j {
			break
		}
		sArr[i], sArr[j] = sArr[j], sArr[i]
		i++
		j--
	}

	return string(sArr)
}
