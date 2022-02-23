package main

import _ "fmt"

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	//小于0或第一位是0最高位不是0
	if x < 0 || (x/10!=0 && x%10==0){
		return false
	}

	temp := x

	xx := 0
	for temp > xx {
		xx = xx*10 + temp%10
		temp /= 10
		// fmt.Println("xx = ", xx)
		// fmt.Println("temp = ", temp)
	}

	// return x == xx

	// fmt.Println("xx/10 = ",  xx/10)
	// fmt.Println("(xx == temp || xx/10 == temp) = ",  xx == temp || xx/10 == temp)

	//奇数位回文数 xx == temp
	//偶数位回文数 xx/10 == temp
	return xx == temp || xx/10 == temp
}

// @lc code=end

func main22() {
	isPalindrome(121)
}
