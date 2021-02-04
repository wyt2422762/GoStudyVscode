package maximumproduct

import (
	"fmt"
	"math"
)

//给定 n 个整数，找出平均数最大且长度为 k 的连续子数组，并输出该最大平均数。
// 输入：[1,12,-5,-6,50,3], k = 4
// 输出：12.75
// 解释：最大平均数 (12-5-6+50)/4 = 51/4 = 12.75

// 基础版本
func findMaxAverage(nums []int, k int) float64 {
	fmt.Println("findMaxAverage")
	sum := math.MinInt64
	le := len(nums)
	for i := 0; i <= le-k; i++ {
		sumInner := 0
		for j := 0; j < k; j++ {
			sumInner += nums[i+j]
		}
		sum = max(sum, sumInner)
	}
	return float64(sum) / float64(k)
}

// 优化版本
func findMaxAverage1(nums []int, k int) float64 {
	fmt.Println("findMaxAverage1")
	var sum int
	le := len(nums)
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum := sum
	for i := k;i<le;i++{
		sum = sum - nums[i - k] + nums[i]
		maxSum = max(maxSum, sum)
	}
	return float64(maxSum) / float64(k)
}
