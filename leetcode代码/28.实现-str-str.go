package main

/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	
	len_h, len_n := len(haystack), len(needle)
	if len_h < len_n {
		return -1
	}

	for i := range haystack {
		for i2 := range needle {
			if i + i2 < len_h && haystack[i+i2] == needle[i2] {
				if i2 == len_n - 1 {
					return i
				}
			} else {
				break
			}
		}
	}

	return -1
}
// @lc code=end


func main11(){
	strStr("mississippi", "issipi")
}

