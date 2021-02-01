package maximumproduct

import (
	"fmt"
	"strings"
	"regexp"
)

//最长回文子串
func longestPalindrome(s string) string {
	fmt.Println("longestPalindrome")
	length := len(s)             //字符串长度
	childLength := 1             //子字符串长度，默认为0
	dp := make([][]bool, length) //用来存放任意两个下标之间的子字符串是否为回文字符串
	for ix, v := range dp {
		if v == nil {
			dp[ix] = make([]bool, length)
		}
	}

	res := "" //返回结果，默认为空字符串
	if length >= 1 {
		r := []rune(s)                               //将字符串转成字符数组
		for ; childLength <= length; childLength++ { //逐个判断子字符串,子字符串的长度从1开始增加
			for i := 0; i < length-childLength+1; i++ {
				j := i + childLength - 1
				fmt.Printf("i: %d, j: %d\n", i, j)
				if childLength == 1 { //子串长度为1，肯定是回文字符串
					dp[i][j] = true
				} else if childLength == 2 { //子串长度为2，只要判断左右两边的字符是否相等即可
					dp[i][j] = (r[i] == r[j])
				} else { //子串长度大于2，需要判断左右两边的字符是否相等，并且左右两边中间的字符串是否为回文字符串
					dp[i][j] = (r[i] == r[j]) && dp[i+1][j-1]
				}
				if dp[i][j] && childLength > len(res) {
					res = string(r[i : j+1])
				}
			}
		}
	}
	return res
}

// 翻转字符串里的单词
// 无空格字符构成一个 单词 。
// 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
// 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
func reverseWords(s string) string {
	s = strings.Trim(s, " ")               //去掉字符串两端的空格，类似于java中的str.trim()
	//正则替换
	reg, _ := regexp.Compile("\\s+");
	s = reg.ReplaceAllString(s, " ")
	s = strings.ReplaceAll(s, "\\s+", " ") //把多个空格改为1个空格

	strs := strings.Split(s, " ")
	for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
		strs[i], strs[j] = strs[j], strs[i]
	}

	return strings.Join(strs, " ")
}
