package maximumproduct

// 给定一个由若干 0 和 1 组成的数组 A，我们最多可以将 K 个值从 0 变成 1 。
// 返回仅包含 1 的最长（连续）子数组的长度。

/* 示例 1：
输入：A = [1,1,1,0,0,0,1,1,1,1,0], K = 2
输出：6
解释：
[1,1,1,0,0,1,1,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 6。 */

/* 示例 2：
输入：A = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
输出：10
解释：
[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 10。 */
func longestOnes(A []int, K int) int {
	res := 0 //返回结果
	le := len(A) //数组长度
	cnt := 0
	left, right := 0, 0 //双指针
	chazhi := 0 //计算两个指针之间的数组之和和全是1的差值
	
	for ; right < le; right++ {
		chazhi += (1 - A[right])
		for chazhi > K && left < right {
			chazhi -= (1 - A[left])
			left++
		}
		cnt = right - left + A[right]
		res = max(res, cnt)
	}

	return res
}
