package maximumproduct

import (
	"fmt"
	//"math"
)

//替换后的最长重复字符
//给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，总共可最多替换 k 次。在执行上述操作后，找到包含重复字母的最长子串的长度。
// 示例 1：

// 输入：s = "ABAB", k = 2
// 输出：4
// 解释：用两个'A'替换为两个'B',反之亦然。
// 示例 2：

// 输入：s = "AABABBA", k = 1
// 输出：4
// 解释：
// 将中间的一个'A'替换为'B',字符串变为 "AABBBBA"。
// 子串 "BBBB" 有最长重复字母, 答案为 4。
func characterReplacement(s string, k int) int {
	fmt.Println("characterReplacement")

	le := len(s) //获取字符串的长度
	if le < 2 {
		return le
	}

	runes := []rune(s)         //把字符串转成字符数组 []rune为unicode字符数组，可以存中文， []byte为字节数组，不能存中文，会被截断
	var start, end int         //开始结束index
	map1 := make(map[rune]int) //用来存放每个字符出现的个数
	var maxCount int           //字符出现的最多次数
	var res int                //结果

	for end < le {
		c := runes[end]
		//字符的出现次数
		if _, ok := map1[c]; ok { //判断map又是有对应的key
			map1[c] = map1[c] + 1
		} else {
			map1[c] = 1
		}

		maxCount = max(maxCount, map1[c]) //修改maxCount
		end++                             //结束index向右移动

		if (end - start) > maxCount+k {
			//说明替换k个字符无法让[start, end)子字符串为字符都一样的字符串，这是需要让开始index向右移动
			c = runes[start]
			map1[c] = map1[1] - 1
			start++
		}

		res = max(res, end-start)
	}

	return res
}
