package maximumproduct

import (
	"fmt"
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
		r := []rune(s)                              //将字符串转成字符数组
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
