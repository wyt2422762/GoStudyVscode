package maximumproduct

import (
	"math"
	"fmt"
)

//返回两个int的最大值
func maxInt(a, b int) int{
	if a > b{
		return a
	}
	return b
}

//求int数组的任意3个数的最大值
func maximumProduct(nums []int) int {
	// 最小的和第二小的
	min1, min2 := math.MaxInt64, math.MaxInt64
	//最大的、第二大的和第三大的
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64
	for i := 0; i < len(nums); i++ {
		x := nums[i]
		if x < min1 {
			min2 = min1
			min1 = x
		} else if x < min2 {
			min2 = x
		}
		if x > max1 {
			max3 = max2
			max2 = max1
			max1 = x
		} else if x > max2 {
			max3 = max2
			max2 = x
		} else if x > max3 {
			max3 = x
		}
	}
	res := maxInt(max1*max2*max3, min1*min2*max1)
	fmt.Println("最大值：", res)
	return res
}
