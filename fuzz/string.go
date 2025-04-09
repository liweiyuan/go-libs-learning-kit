package fuzz

import (
	"strings"
	"unicode"
)

// 检查字符串是否是回文
func IsPalindrome(s string) bool {
	// 将字符串转换为小写并去掉空格
	s = strings.ToLower(s) // 确保忽略大小写
	s = removeNonAlphanumeric(s)

	// 判断反转后是否相等
	return s == reverse(s)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// removeNonAlphanumeric 移除非字母字符
func removeNonAlphanumeric(s string) string {
	var result []rune
	for _, r := range s {
		// 只保留字母字符
		if unicode.IsLetter(r) {
			result = append(result, r)
		}
	}
	return string(result)
}
