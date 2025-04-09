package fuzz

import (
	"testing"
	"unicode/utf8"
)

func FuzzIsPalindrome(f *testing.F) {

	// 向测试添加初始值
	f.Add("hello")
	f.Add("madam")
	f.Add("racecar")
	f.Add("éçà") // 添加包含非ASCII字符的例子
	f.Add("aA")  // 添加大小写不同的回文字符串
	//f.Add("A man a plan a canal Panama") // 经典的回文

	// Fuzz() 方法接收一个处理函数，每次执行测试时会传入模糊化的输入
	f.Fuzz(func(t *testing.T, input string) {
		// 如果输入包含无效的UTF-8字符，我们跳过
		if !utf8.ValidString(input) {
			t.Skipf("Skipping invalid UTF-8 input: %s", input)
			return
		}

		// 调用 IsPalindrome 函数，并验证其输出
		result := IsPalindrome(input)

		// 我们可以对回文字符串做一个断言，检查输出是否正确
		if result && input != reverse(input) {
			t.Errorf("For input %s, expected palindrome but got false", input)
		}
	})
}
