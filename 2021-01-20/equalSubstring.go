package maximumproduct

import (
	"fmt"
)

// 给你两个长度相同的字符串，s 和 t。
// 将 s 中的第 i 个字符变到 t 中的第 i 个字符需要 |s[i] - t[i]| 的开销（开销可能为 0），也就是两个字符的 ASCII 码值的差的绝对值。
// 于变更字符串的最大预算是 maxCost。在转化字符串时，总开销应当小于等于该预算，这也意味着字符串的转化可能是不完全的。

// 如果你可以将 s 的子字符串转化为它在 t 中对应的子字符串，则返回可以转化的最大长度。
// 如果 s 中没有子字符串可以转化成 t 中对应的子字符串，则返回 0。
func equalSubstring(s string, t string, maxCost int) int {
	fmt.Println("equalSubstring")
	res := 0                 //返回结果
	le := len(s)             //字符串长度
	diffs := make([]int, le) //用来存储每个位置的差值
	for i := 0; i < le; i++ {
		diffs[i] = abs(int(s[i]) - int(t[i]))
	}
	start, end := 0, 0 //开始，结束index
	diffSum := 0       //记录插值的和
	for end < le {
		diffSum += diffs[end]   //差值总和加
		for diffSum > maxCost { //如果差值的和大于maxCost, start指针向右移动
			diffSum -= diffs[start] //差值总和-
			start++
		}
		res = max(res, end-start+1)
		end++ //end指针右移
	}
	return res
}
