package maximumproduct

import (
	"fmt"
	"sort"
)

func medianSlidingWindow(nums []int, k int) []float64 {
	fmt.Println("medianSlidingWindow")
	le := len(nums)
	res := make([]float64, le-k+1)

	large := make([]int, k/2) //存储较大的一边的数据
	small := make([]int, k/2+(k%2))
	fmt.Printf("small: %T, larger: %T\n", small, large)

	if le < 1 {
		return res
	}

	for i := 0; i+k <= le; i++ {
		res[i] = midValue(getSortedSubNums(nums, i, i+k-1, small, large))
	}
	return res
}

func getSortedSubNums(nums []int, start, end int, small []int, larger []int) []int {
	res := make([]int, end-start+1)
	for i := start; i <= end; i++ {
		res[i-start] = nums[i]
		//small = append(small, nums[i])
	}
	sort.Ints(res)
	return res
}

//获取排好序的数组的中位数
func midValue(sortedNums []int) float64 {
	le := len(sortedNums)
	if le%2 != 0 { //长度为奇数
		return float64(sortedNums[le/2])
	}
	return float64(sortedNums[le/2-1]+sortedNums[le/2]) / float64(2)
}
